#!/bin/bash

# Check for required arguments
if [ "$#" -ne 2 ]; then
  echo "Usage: $0 <proto_file_without_extension> <go|descriptor>"
  echo "Example: $0 circleoftrustmembers go"
  exit 1
fi

PROTO_NAME="$1"
OUTPUT_TYPE="$2"

# Static paths
PROTO_DIR="./"  # Directory containing .proto files
OUTPUT_DIR="./$1" # Desired output directory for generated files
mkdir $1
PROTO_FILE="$PROTO_NAME.proto"  # File name relative to PROTO_DIR

# Check if protoc is installed
if ! command -v protoc &>/dev/null; then
  echo "Error: protoc is not installed. Please install it before running this script."
  exit 1
fi

# Check if the .proto file exists
if [ ! -f "./$PROTO_FILE" ]; then
  echo "Error: File '$PROTO_DIR/$PROTO_FILE' not found."
  exit 1
fi

# Generate code based on the output type
if [ "$OUTPUT_TYPE" == "go" ]; then
  echo "Generating Go code for '$PROTO_FILE' into $OUTPUT_DIR..."
  protoc --proto_path="$PROTO_DIR" \
         --go_out="$OUTPUT_DIR" --go_opt=paths=source_relative \
         --go-grpc_out="$OUTPUT_DIR" --go-grpc_opt=paths=source_relative \
         "$PROTO_DIR/$PROTO_FILE"
elif [ "$OUTPUT_TYPE" == "descriptor" ]; then
  echo "Generating descriptor set for '$PROTO_FILE' into $OUTPUT_DIR..."
  protoc --proto_path="$PROTO_DIR" \
         --descriptor_set_out="$OUTPUT_DIR/$PROTO_NAME.pb" "$PROTO_DIR/$PROTO_FILE"
else
  echo "Error: Invalid output type '$OUTPUT_TYPE'. Use 'go' or 'descriptor'."
  exit 1
fi

echo "Protobuf generation completed successfully in $OUTPUT_DIR."