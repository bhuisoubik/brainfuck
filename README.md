# Brainfuck Interpreter
CLI Tool for Interpreting Brainfuck Code written in Go<br>
Example : A sample [Hello World](./test.bf) file is included in the project's root directory
## Build
Open your terminal and write the following command
```
mkdir $GOPATH/src/github.com/soubikbhuiwk007
cd $GOPATH/src/github.com/soubikbhuiwk007
git clone https://github.com/soubikbhuiwk007/brainfuck.git
cd brainfuck
make build
```
## Features
### Interactive Brainfuck
Run Brainfuck code without creating any file<br>
```
Interactive Brainfuck (v1.0.1)
Type exit to Exit
brainfuck >> bf code
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
Usage : 
    brainfuck <command> <arguments>
Commands :
    run [filename]         :    To Interpret a Brainfuck (*.bf) file
    -v, version            :    To Print the Current Version
    -m, memory [filename]  :    To Run & Print the Memory usage of a Brainfuck (*.bf) file
    -h, help               :    For Help
```
## Install
```
cd $GOPATH/src/github.com/soubikbhuiwk007/brainfuck
go install
```
Check if it was correctly installed<br>
Try running: ```brainfuck -v```<br>
This should output<br>
```
Version: 1.0.1
```
## License

This project is licensed under MIT License, check out the [LICENSE](./LICENSE) file.

### If You found any problem or any fault, you can create an issue here
### Feel free to fork the repo, and contribute to this repo.
