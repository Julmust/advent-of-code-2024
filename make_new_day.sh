#!/bin/bash

if [ -d ./day_$dn ]; then
    echo Folder ./day_$1 already exists. Exiting
    exit
fi

mkdir ./day_$1
( cd ./day_$1 ; \
    touch example.txt
    touch input.txt
    go mod init aoc_2024/day_$1 ; \
    go work use . ; \
    cp ../new_day_template ./main.go )