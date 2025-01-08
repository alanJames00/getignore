# directory defs
BUILD_DIR=build
SRC_DIR=./cmd/getignore
BIN_NAME=getignore

# Check for OS and append .exe for Windows
ifeq ($(OS),Windows_NT)
    BIN_NAME := $(BIN_NAME).exe
    MKDIR_CMD := if not exist $(BUILD_DIR) mkdir $(BUILD_DIR)
else
    MKDIR_CMD := mkdir -p $(BUILD_DIR)
endif

# Phony target
.PHONY: build

# build the project
build:
	@echo "Building the project"
	@$(MKDIR_CMD)
	go build -o $(BUILD_DIR)/$(BIN_NAME) $(SRC_DIR)
