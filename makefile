# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=TETRIS

# Determine the OS and architecture
ifeq ($(OS),Windows_NT)
    OSFLAG := windows
    EXTENSION := .exe
else
    OSFLAG := $(shell uname | tr '[:upper:]' '[:lower:]')
    EXTENSION :=
endif

build: $(BINARY_NAME)

$(BINARY_NAME):
	$(GOBUILD) -o $(BINARY_NAME)$(EXTENSION) -v

# Build for Linux
linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)_$(OSFLAG)_amd64$(EXTENSION) -v

# Build for Mac
darwin:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)_$(OSFLAG)_amd64$(EXTENSION) -v

# Build for Windows
windows:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)_windows_amd64$(EXTENSION) -v

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)*$(EXTENSION)
