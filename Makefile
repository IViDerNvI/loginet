BIN_NAME := loginet
SRC_FILES := $(shell find . -name "*.go")
BIN_DIR := bin
BUILD_DIR := build
GIT_FILE := .gitignore .git/
READ_FILE := README_zh.md License
GO_FILE := $(wildcard *go*)
CONFIG_FILE := config.json
ALL_FILE := $(GIT_FILE) $(READ_FILE) $(GO_FILE) $(CONFIG_FILE)

.DEFAULT_GOAL := help

.PHONY: help build clean

help: Makefile
	@echo "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' $^ | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"


## build 	build loginet into bin directory
## install 	install loginet into your PATH
## clean 	remove build files

check_brew_installation:

ifeq (, $(shell command -v brew))
	@echo "brew is not installed, downloading..."
	/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
#   /bin/zsh -c "$(curl -fsSL https://gitee.com/cunkai/HomebrewCN/raw/master/Homebrew.sh)" 国内源
	@echo "brew install successfully!"
endif

check_go_installation: check_brew_installation

ifeq (,$(shell command -v go))
	@echo "Go is not installed, downloading..."
	brew install golang
	@echo "Go has been installed successfully!"
else
	@echo "Go is already installed."
endif


build: check_go_installation $(BIN_DIR)/$(BIN_NAME)

$(BIN_DIR)/$(BIN_NAME): $(SRC_FILES)

	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BIN_NAME) main.go
	@echo "构建完毕"

install: build
	sudo mv $(BIN_DIR)/$(BIN_NAME) /usr/local/bin

clean:
	rm -rf $(BIN_DIR) $(BUILD_DIR)
	@echo "清理完毕"

tar: $(ALL_FILE)
	tar -a -cf loginet.tar.gz $(ALL_FILE)
