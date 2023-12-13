# tcfdecoder

A simple cli to decode TCF2 strings to JSON. Needed something to pipe a lot of consent strings into for further insights.

E.g. checking for distribution of TCF2 versions
```
cat a-lot-of-consent-strings.txt | tcfdecode | jq .TCFPolicyVersion | sort | uniq -c
```

There is an official node cli from the IAB but it does output to multiple lines and is not up to date.
