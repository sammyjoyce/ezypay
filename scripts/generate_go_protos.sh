#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Find all proto files recursively
PROTO_FILES=$(find api/protos -name "*.proto")

echo "Generating Go protobuf files..."
for file in $PROTO_FILES; do
    out_dir=$(echo $file | sed 's|api/protos/|internal/gen/|' | xargs dirname)
    mkdir -p "$out_dir"
    
    proto_dir=$(dirname $file)
    
    protoc \
        --go_out=. --go_opt=module={{go_module}} \
        --go-grpc_out=. --go-grpc_opt=module={{go_module}} \
        --proto_path=$proto_dir \
        --proto_path=api/protos \
        $file
    
    echo "Generated Go files for $file"
done

echo "✨ Go protobuf generation completed!"
