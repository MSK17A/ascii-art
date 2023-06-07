#!/bin/bash

#* test file for multiple diff cases

go run . "\"test\""

echo

go run . "HELLO"

echo

go run . "HeLlo HuMaN"

echo

go run . "1Hello 2There"

echo


go run . "Hello\nThere"

echo

go run . "{Hello & There #}"

echo

go run . 'hello There 1 to 2!' 

echo

go run . "MaD3IrA&LiSboN"

echo

go run . "1a\"#FdwHywR&/()="

echo


go run . "{|}~"

echo 

go run . "1a\"#FdwHywR&/()="

echo


go run . ":;<=>?@"

echo

go run . "[\]^_ 'a"

echo

go run . '\!" #$%&'"'"'()*+,-./'

echo


go run . "RGB"

echo

go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

echo

go run . "abcdefghijklmnopqrstuvwxyz"