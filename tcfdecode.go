package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/LiveRamp/iabconsent"
	"github.com/urfave/cli/v2"
)

func decodeTcf(consent string) (string, error) {
	if consent == "" {
		return "", errors.New("missing TCF string")
	}

	v2, err := iabconsent.ParseV2(consent)
	if err != nil {
		return "", err
	}

	tcf2JSON, err := json.Marshal(v2)
	if err != nil {
		return "", err
	}

	return string(tcf2JSON), nil
}

func isInputFromPipe() bool {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func main() {
	app := &cli.App{
		Name:  "tcfdecode",
		Usage: "decode tcf2 strings to json",
		Action: func(cCtx *cli.Context) error {
			if isInputFromPipe() {
				scanner := bufio.NewScanner(os.Stdin)
				line := 0
				failed := false
				for scanner.Scan() {
					line++
					decoded, err := decodeTcf(scanner.Text())
					if err != nil {
						fmt.Fprintf(os.Stderr, "line %d: %v\n", line, err)
						failed = true
						continue
					}
					fmt.Println(decoded)
				}
				if err := scanner.Err(); err != nil {
					return fmt.Errorf("reading stdin: %w", err)
				}
				if failed {
					return cli.Exit("one or more TCF strings failed to decode", 1)
				}
				return nil
			}

			decoded, err := decodeTcf(cCtx.Args().Get(0))
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}
			fmt.Println(decoded)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
