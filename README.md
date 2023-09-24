# tcfdecoder

A simple cli to decode TCF2 strings to JSON. Needed something to pipe a lot of consent strings into for further insights.

E.g. checking for distribution of TCF2 versions
`cat a-lot-of-consent-strings.txt | tcfdecode | jq .TCFPolicyVersion | sort | uniq -c | sort -nr`

There is an official node cli from the IAB but it does output to multiple line and is not up to date with. 
Using Liveramp's Golang implementation of the IAB Specs from: https://github.com/LiveRamp/iabconsent