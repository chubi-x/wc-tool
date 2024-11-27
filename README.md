# Word count tool

Command line tool to count words, lines, characters, and bytes in a given file.

## Options

The tool accepts a total of 4 flags, each to count a different entity. They are:

- `-c` counts bytes
- `-l` counts lines
- `-w` counts words
- `-m` counts characters

The command can receive input from a file or from stdin.

## Usage

you can build the binary with `go build wc.go` and use it as follows:

`./wc [-option] [-filename]`

Or through Stdin using a pipe: `cat test.txt | ./wc [-option]`

Built as part of John Cricket's [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-wc/)
