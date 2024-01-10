#!/bin/bash

# Check if filename is provided
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 filename.txt"
    exit 1
fi

filename=$1

# Check if file exists
if [ ! -f "$filename" ]; then
    echo "File $filename not found"
    exit 1
fi

# Process each subdomain in the file
while IFS= read -r subdomain; do
    modified_subdomain=$(echo "$subdomain" | sed 's/^[^.]*\.//')
    echo "$modified_subdomain"
done < "$filename"
