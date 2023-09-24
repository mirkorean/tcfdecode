package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/LiveRamp/iabconsent"
	"github.com/urfave/cli/v2"
)

func decodeTcf(consent string) string {
	v2, _ := iabconsent.ParseV2(consent)
	tcf2Json, _ := json.Marshal(v2)
	return string(tcf2Json)
}

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func main() {
	app := &cli.App{
		Name:  "tcfdecode",
		Usage: "decode tcf2 strings to json",
		Action: func(cCtx *cli.Context) error {
			if isInputFromPipe() {
				scanner := bufio.NewScanner(os.Stdin)
				for scanner.Scan() {
					fmt.Println(decodeTcf(scanner.Text()))
				}
				if err := scanner.Err(); err != nil {
					fmt.Fprintln(os.Stderr, "reading stdin:", err)
				}
			} else {
				fmt.Println(decodeTcf(cCtx.Args().Get(0)))
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
