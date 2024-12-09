#!/bin/env bash

day=$1
cookie=$2

if [ -z "$day" ] || [ -z "$cookie" ]; then 
    echo "Usage: $0 day cookie"
    exit 1
fi

curl --cookie "session=$cookie" https://adventofcode.com/2024/day/$day/input -o day$day/input.txt