#!/bin/bash
echo Building domainChecker...
go build -o domainChecker

echo Building synonyms...
cd ../synonym
go build -o ../domainChecker/lib/synonyms

echo Building available...
cd ../avail
go build -o ../domainChecker/lib/available

echo Building sprinkle...
cd ../sprinkl
go build -o ../domainChecker/lib/sprinkle


echo Building coolify...
cd ../coolify
go build -o ../domainChecker/lib/coolify


echo Building domainify...
cd ../dominit
go build -o ../domainChecker/lib/domainify

echo Done.