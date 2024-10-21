#!/bin/bash

# Define the Go file and output binary name
GO_FILE="main.go"
OUTPUT_BINARY="CatFactCLI"

# Build the Go program
echo -e "Building $GO_FILE...\n"
go build -o $OUTPUT_BINARY $GO_FILE

# Check if the build was successful
if [ $? -eq 0 ]; then
    echo -e "Build successful!\n"
    echo -e "Installing $OUTPUT_BINARY\n"
    go install $OUTPUT_BINARY
    
    if [ $? -eq 0 ]; then
        echo -e "Installation successful!\n"
    else
        echo -e "Installation failed.\n"
    fi
else
    echo -e "Build failed.\n"
fi
