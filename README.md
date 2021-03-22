# Brainfuck Interpreter (v1.0.0)
CLI Tool for Interpreting Brainfuck Code written in Go<br>
Example : A sample [Hello World](./test.bf) brainfuck file is included in the project's root directory
## Build
In terminal write the following commands
```
cd $GOPATH/src
mkdir github.com/soubikbhuiwk007 && cd github.com/soubikbhuiwk007
git clone https://github.com/soubikbhuiwk007/brainfuck.git
cd brainfuck
go build
```
## Install
You can install this brainfuck cli to use it globally across your system.
```
cd $GOPATH/src/github.com/soubikbhuiwk007/brainfuck
go install
```
To check if it was properly installed on your system, you type the following in terminal
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
**Usage**: Just launch brainfuck.exe after building the project
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
## License

This project is licensed under MIT License, check out the [LICENSE](./LICENSE) file.

#### If You found any problem, feel free to create an issue here
