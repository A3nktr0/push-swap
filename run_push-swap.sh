#!/bin/zsh

echo "push empty"
./push-swap
echo

echo "normal push"
./push-swap "2 1 3 6 5 8"
echo 

echo "ordered push"
./push-swap "0 1 2 3 4 5"
echo 

echo "bad input"
./push-swap "0 one 2 3"
echo 

echo "duplicate"
./push-swap "1 2 2 3"
echo

echo "random five"
./push-swap "56 4 32 6 1"
echo 

echo "random five 2"
./push-swap "8 9 14 6 3"
