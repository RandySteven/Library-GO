#!/bin/bash

# Directory containing generated mocks
MOCKS_DIR="./mocks"

# Base package directory
BASE_PKG_DIR="./pkg"

# Function to organize mocks
organize_mocks() {
  find "$MOCKS_DIR" -type f -name "*.go" | while read -r mock_file; do
    # Extract the package name from the file content
    pkg_name=$(grep -E '^package ' "$mock_file" | awk '{print $2}')

    # Find the original package directory
    pkg_path=$(find "$BASE_PKG_DIR" -type d -name "$pkg_name" -print -quit)

    if [[ -n "$pkg_path" ]]; then
      # Create corresponding directory in mocks and move the file
      mkdir -p "$MOCKS_DIR/$pkg_name"
      mv "$mock_file" "$MOCKS_DIR/$pkg_name/"
    else
      echo "Package $pkg_name not found for mock $mock_file"
    fi
  done
}

organize_mocks
