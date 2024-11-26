#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Directory configurations
PROTO_DIR="../api/protos/hello"
GENERATED_CODE_DIR="./app/api/generated/protos/hello"

# Ensure the proto directory exists
if [ ! -d "$PROTO_DIR" ]; then
  echo "Error: Proto directory $PROTO_DIR does not exist."
  exit 1
fi

# Create the output directory if it doesn't exist
mkdir -p "$GENERATED_CODE_DIR"

# Clean the output directory
rm -rf "${GENERATED_CODE_DIR:?}"/*

# Debug information
echo "Proto directory: $PROTO_DIR"
echo "Found proto files:"
find "$PROTO_DIR" -name "*.proto" -type f

# Generate TypeScript code from proto files
PROTO_FILES=$(find "$PROTO_DIR" -name "*.proto" -type f)

grpc_tools_node_protoc \
  --plugin=protoc-gen-ts_proto=./node_modules/.bin/protoc-gen-ts_proto \
  --ts_proto_out=$GENERATED_CODE_DIR \
  --ts_proto_opt=outputServices=nice-grpc,outputServices=generic-definitions,useExactTypes=false,esModuleInterop=true,importSuffix=.js \
  --proto_path=../api/protos/hello \
  $PROTO_FILES

echo "Proto generation completed successfully!"

# Make the generated code more TypeScript-friendly
for file in "$GENERATED_CODE_DIR"/*.ts; do
  echo "// @ts-nocheck" | cat - "$file" > temp && mv temp "$file"
done

echo " ✨ TypeScript protobuf code generated in $GENERATED_CODE_DIR"
