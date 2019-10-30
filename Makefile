#-------------------------------------------------------------------------------
# Copyright (c) 2016 Mateus Furquim
#-------------------------------------------------------------------------------

#-------------------------------------------------------------------------------
# Nome do projeto
# https://stackoverflow.com/questions/18136918/how-to-get-current-relative-directory-of-your-makefile
#-------------------------------------------------------------------------------
mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
NAME := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))

SRC_DIR := ./src
BIN_DIR := ./bin

TARGET := $(BIN_DIR)/$(NAME)

.PHONY: all run clean build dirs

all: dirs
	@$(MAKE) $(TARGET)
	@$(MAKE) run
	@$(MAKE) clean

$(TARGET): $(SRC_DIR)/main.go
	@go build -o $@ $^

run: $(TARGET)
	@$^ ${ARGS}

dirs:
	@mkdir -p $(BIN_DIR)

clean:
	@rm -f $(TARGET) &> /dev/null
