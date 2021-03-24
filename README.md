# Brainfuck Interpreter
CLI Tool for Interpreting Brainfuck Code written in Go<br>
Example : A sample [Hello World](./test.bf) brainfuck file is included in the project's root directory<br>
NOTE: You need to have go1.16 installed on your system.
## Build
In terminal write the following commands
```
mkdir $GOPATH/src/github.com/soubikbhuiwk007
cd $GOPATH/src/github.com/soubikbhuiwk007
git clone https://github.com/soubikbhuiwk007/brainfuck.git
cd brainfuck
make build
```
## Run
```
cd $GOPATH/src/github.com/soubikbhuiwk007
make run
```
## Install
You can install this brainfuck cli to use it globally across your system.
```
cd $GOPATH/src/github.com/soubikbhuiwk007/brainfuck
go install
```
To check if it was properly installed on your system, type the following in terminal
```
brainfuck -v
```
This should print:
```
Version: 1.0.0
```
## Usage
```
Usage : 
    brainfuck <command> <arguments>
Commands :
    run [filename]         :    To Interpret a Brainfuck (*.bf) file
    -v, version            :    To Print the Current Version
    -m, memory [filename]  :    To Run & Print the Memory usage of a Brainfuck (*.bf) file
    -h, help               :    For Help
```

## Features
### 1. Interactive Brainfuck
Here you can run brainfuck code without creating any file<br>
**Usage**: Just launch brainfuck after building the project
```
Interactive Brainfuck (v1.0.0)
Type exit to Exit
brainfuck >> bf code
```
### 2. Print Memory
This is used for running & printing the memory values in the same time<br>
**Usage**: After building the project, run
```
brainfuck -m <filename>
```
Example ```brainfuck -m test.bf```:
This will output the following
```
Hello World!

Memory:  [0 87 100 33 10]
```
## Makefile
```make build``` to build the project<br>
```make run``` to run the main.go<br>
```make clean``` to clean (remove executables) the project<br>
```make version``` to print the version<br>
## License

This project is licensed under MIT License, check out the [LICENSE](./LICENSE) file.

#### If You found any problem, you can create an issue here
#### Feel free to fork the repo, and contribute to it
