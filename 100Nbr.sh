#!/bin/zsh

python generate_numbers.py
f="100numbers.txt"; while read line; do tr -s "\n" " ""${line}"; done < ${f} | sed 's/.$//'
./push-swap "${f}" | ./checker "${f}"
