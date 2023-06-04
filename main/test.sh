#!/bin/bash

#* test file for multiple diff cases

go run . "MaD3IrA&LiSboN"

echo 

go run . "1a\"#FdwHywR&/()="

echo

go run . "{|}~"

echo

go run . "[\]^_ 'a"

echo

go run . "RGB"

echo

go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

echo

go run . "abcdefghijklmnopqrstuvwxyz"