V = 1.0.1
BUILD = ${CURDIR}/build
GO = go

run:
	@echo "Running main.go ..."
	${GO} run main.go

build:
	@echo "Building executable..."
	${GO} build -o ${BUILD}/brainfuck main.go
	@echo "Finished"

clean:
	@echo "Removing executable..."
	rm -f ${BUILD}/brainfuck
	@echo "Completed"
	@echo "Removing build directory..."
	rmdir ${BUILD}
	@echo "Completed"

version:
	@echo "Version :" ${V}