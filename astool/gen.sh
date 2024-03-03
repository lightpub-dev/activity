#!/bin/bash -x
go build
rm -rf ../streams
./astool  -spec ./activitystreams.jsonld -spec forgefed.jsonld -spec security-v1.jsonld -spec toot.jsonld -path github.com/go-fed/activity ./streams
mv streams ../
