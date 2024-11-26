#!/bin/bash

# Check if environment variables are set
if [ -z "$PROJECT_NAME" ] || [ -z "$PROJECT_DESCRIPTION" ] || [ -z "$GO_MODULE" ] || [ -z "$AUTHOR_NAME" ] || [ -z "$AUTHOR_EMAIL" ]; then
    echo "Please set the following environment variables:"
    echo "PROJECT_NAME - The name of your project"
    echo "PROJECT_DESCRIPTION - A short description of your project"
    echo "GO_MODULE - The Go module path (e.g., github.com/username/project)"
    echo "AUTHOR_NAME - Your name"
    echo "AUTHOR_EMAIL - Your email"
    exit 1
fi

# Set LC_ALL to handle special characters
export LC_ALL=C

# Function to escape special characters in sed replacement
escape_sed() {
    echo "$1" | sed -e 's/[\/&]/\\&/g'
}

# Escape the GO_MODULE value for sed
ESCAPED_GO_MODULE=$(escape_sed "$GO_MODULE")

echo "Updating Go files..."
# Update Go imports with exact pattern match including quotes
find . -type f -name "*.go" -exec perl -pi -e 's/"{{go_module}}\/([^"]+)"/"'"${ESCAPED_GO_MODULE}"'\/\1"/g' {} +

echo "Updating Proto files..."
# Update Go package declarations in .proto files
find . -type f -name "*.proto" -exec perl -pi -e 's/option go_package = "{{go_module}}([^"]+)"/option go_package = "'"${ESCAPED_GO_MODULE}"'\1"/g' {} +

echo "Updating other template variables..."
# Update other template variables in all files
find . -type f -not -path "*/\.*" -not -path "*/node_modules/*" -not -path "*/dist/*" -exec sed -i '' \
    -e "s|{{project_name}}|${PROJECT_NAME}|g" \
    -e "s|{{project_description}}|${PROJECT_DESCRIPTION}|g" \
    -e "s|{{author_name}}|${AUTHOR_NAME}|g" \
    -e "s|{{author_email}}|${AUTHOR_EMAIL}|g" {} 2>/dev/null || true

echo "Updating go.mod..."
# Update go.mod with new module name
if [ -f "go.mod" ]; then
    sed -i '' "s|module.*|module ${ESCAPED_GO_MODULE}|" go.mod
fi

echo "Updating package.json..."
# Update package.json with new project name and description
if [ -f "package.json" ]; then
    sed -i '' \
        -e "s|\"name\":.*|\"name\": \"${PROJECT_NAME}\",|" \
        -e "s|\"description\":.*|\"description\": \"${PROJECT_DESCRIPTION}\",|" \
        package.json
fi

echo "Template setup completed successfully!"