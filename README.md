# Brainfuck Interpreter
CLI Tool for Interpreting Brainfuck Code written in Go<br>
Example : A sample [Hello World](./test.bf) file is included in the project's root directory
## Build
Open your terminal and write the following command
```
cd $GOPATH/src
mkdir github.com/soubikbhuiwk007 && cd github.com/soubikbhuiwk007
git clone https://github.com/soubikbhuiwk007/brainfuck.git
cd brainfuck
go build
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

## License

This project is licensed under MIT License, check out the [LICENSE](./LICENSE) file.

### If You found any problem or any fault, feel free to create an issue here