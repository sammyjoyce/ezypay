#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Find all proto directories recursively
PROTO_DIRS=$(find api/protos -type d -mindepth 1)

cd web

echo "Generating Web protobuf files..."
for proto_dir in $PROTO_DIRS; do
    dir_name=$(basename $proto_dir)
    out_dir="./app/api/generated/protos/$dir_name"
    
    # Create and clean output directory
    mkdir -p "$out_dir"
    rm -rf "${out_dir:?}"/*
    
    echo "Generating TypeScript files for $proto_dir..."
    
    ./node_modules/.bin/grpc_tools_node_protoc \
        --plugin=protoc-gen-ts_proto=./node_modules/.bin/protoc-gen-ts_proto \
        --ts_proto_out="$out_dir" \
        --ts_proto_opt=outputServices=nice-grpc,outputServices=generic-definitions,useExactTypes=false,esModuleInterop=true,importSuffix=.js \
        --proto_path="../$proto_dir" \
        --proto_path="../api/protos" \
        "../$proto_dir"/*.proto
    
    # Add ts-nocheck to generated files
    for file in "$out_dir"/*.ts; do
        if [ -f "$file" ]; then
            echo "// @ts-nocheck" | cat - "$file" > temp && mv temp "$file"
        fi
    done
    
    echo "Generated TypeScript files for $proto_dir"
done

cd ..
echo "✨ Web protobuf generation completed!"
