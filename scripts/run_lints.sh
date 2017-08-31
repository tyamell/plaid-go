#!/bin/bash

TIMEOUT=300s

for i in $@; do
    CGO_ENABLED=0 gometalinter -e format_gen --deadline=$TIMEOUT --disable-all -Evet -Egolint -Egoimports -Eineffassign -Egosimple -Eerrcheck -Eunused -Edeadcode -Emisspell -Egocyclo --cyclo-over=15 $i && echo "ok\t$i"
done