# Push-Swap

### Objectives

**Push-Swap** is a very simple project that uses a Non-Comparative Sorting Algorithm. You have at your disposal a list of int values, two stacks (a and b) and a set of instructions.

You will have to write 2 programs:

- `push-swap`, which calculates and displays on the standard output the smallest program using push-swap instruction language that sorts integer arguments received.

- `checker`, which takes integer arguments and reads instructions on the standard input. Once read, checker executes them and displays OK if integers are sorted. Otherwise, it will display KO.

These are the instructions that you can use to sort the stack :

- `pa` push the top first element of stack `b` to stack `a`
- `pb` push the top first element of stack `a` to stack `b`
- `sa` swap first 2 elements of stack `a`
- `sb` swap first 2 elements of stack `b`
- `ss` execute `sa` and `sb`
- `ra` rotate stack `a` (shift up all elements of stack `a` by 1, the first element becomes the last one)
- `rb` rotate stack `b`
- `rr` execute `ra` and `rb`
- `rra` reverse rotate `a` (shift down all elements of stack `a` by 1, the last element becomes the first one)
- `rrb` reverse rotate `b`
- `rrr` execute `rra` and `rrb`


### Usage

- Display movements to sort the list :

```command
$ ./push-swap "<numbers list>"
```

- Check if movement sort the list :

```command
$ echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" | ./checker "0 9 1 8 2"

// result
OK
```

- To 100 numbers list :

Run conda environment and python script to generate 100 numbers file

```command
$ conda activate py_env

$ python generate_numbers.py
```

Then run Shell script to display numbers 

```command
$ ./100Nbr.sh
```

Paste your numbers list in variable `ARG` and run the program 

```command
$ ARG=("<100 random numbers>"); ./push-swap $ARG | ./checker $ARG
```

