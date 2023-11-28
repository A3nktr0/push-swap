#!/bin/zsh

echo "empty checker"
./checker
echo 

echo "bad instruction"
./checker
echo

echo "normal checker, incomplete instructions"
echo -e "sa\npb\nrrr\n" | ./checker "0 9 1 8 2 7 3 6 4 5"
echo

echo "normal checker, good instructions"
echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" | ./checker "0 9 1 8 2"
echo

echo "using args"
ARG=("4 67 3 87 23"); ./push-swap $ARG | ./checker $ARG
echo
