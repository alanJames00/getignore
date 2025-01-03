# directory defs
BUILD_DIR=build
SRC_DIR=./cmd/getignore
BIN_NAME=getignore

# build the project
build:
	@echo "Building the project"
	mkdir ${BUILD_DIR}
	go build -o ${BUILD_DIR}/${BIN_NAME} ${SRC_DIR}
