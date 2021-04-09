GO = go

all: 
	$(GO) build ./cmd/brainfuck
	$(GO) build ./cmd/ibf

brainfuck:
	$(GO) build ./cmd/brainfuck

ibf:
	$(GO) build ./cmd/ibf

install:
	$(GO) install ./cmd/brainfuck
	$(GO) install ./cmd/ibf

clean:
	rm -f ibf
	rm -f brainfuck
