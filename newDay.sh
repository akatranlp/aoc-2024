#!/bin/env bash

day=$1

if [ -z "$day" ]; then
    echo "usage: $0 <1 - 24>"
    exit 1
fi

mkdir day$day
cp template/main.go day$day/main.go
cp template/go.mod day$day/go.mod
cp template/input.txt day$day/input.txt
cp template/day1_test.go day$day/day${day}_test.go
cp template/day1.go day$day/day${day}.go
go work use ./day$day

for file in day$day/*; do
    sed -i -e "s/day1/day$day/g" $file
    sed -i -e "s/Day1/Day$day/g" $file
    sed -i -e "s/aoc-day-1/aoc-day-$day/g" $file
done
