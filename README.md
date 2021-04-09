# Brainfuck Interpreter
CLI Tool for Interpreting Brainfuck Code written in Go<br>
## Build
### Prerequisite
```
1. Go sdk (^1.16.0)
2. make
```
Open your terminal and write the following command
```shell
$ mkdir $GOPATH/src/github.com/soubikbhuiwk007
$ cd $GOPATH/src/github.com/soubikbhuiwk007
$ git clone https://github.com/soubikbhuiwk007/brainfuck.git
$ cd brainfuck
$ make
```
## Features
### Interactive Brainfuck
Run Brainfuck code without creating any file<br>
Open Terminal, and type `ibf`
```
Interactive Brainfuck (v1.0.0)
Type exit to Exit
ibf >> bf code
```
### Memory View
View the Memory usage by the brainfuck code<br>
Command: ```brainfuck -m test.bf```
```
Hello World!

Memory:  [0 87 100 33 10]
```
## Usage
```

Brainfuck

Usage:
        brainfuck <command> || <file>

Command:
        -v, version     :       Print Version
        -m <file>       :       Memory view
        -h, help        :       Help (this)

```
## Install
Go Project's root directory and in terminal type:
```shell
$ make install
```
Check if it was correctly installed<br>
Try running: ```brainfuck -v```<br>
This should output<br>
```
Version: 1.0.2
```
## License

This project is licensed under MIT License, check out the [LICENSE](./LICENSE) file.

#### If You found any problem, you can create an issue here
#### Feel free to fork the repo, and contribute to this repo.
