#!/bin/bash

echo Input day number:
read dn

if [ -d ./day_$dn ]; then
    echo Folder ./day_$dn already exists. Exiting
    exit
fi

mkdir ./day_$dn
( cd ./day_$dn ; \
    touch example.txt
    touch input.txt
    go mod init aoc_2024/day_$dn ; \
    go work use . ; \
    echo -e "package main\n\n\nfunc main() {\n}" >> main.go )