#!/bin/bash

# Replace template variables in files
find . -type f -not -path "*/\.*" -not -path "*/node_modules/*" -not -path "*/dist/*" -exec sed -i '' \
    -e "s|{{project_name}}|$PROJECT_NAME|g" \
    -e "s|{{project_description}}|$PROJECT_DESCRIPTION|g" \
    -e "s|{{go_module}}|$GO_MODULE|g" \
    -e "s|{{author_name}}|$AUTHOR_NAME|g" \
    -e "s|{{author_email}}|$AUTHOR_EMAIL|g" {} +

# Update go.mod with new module name
if [ -f "go.mod" ]; then
    sed -i '' "s|module.*|module $GO_MODULE|" go.mod
fi

# Update package.json with new project name and description
if [ -f "package.json" ]; then
    sed -i '' \
        -e "s|\"name\":.*|\"name\": \"$PROJECT_NAME\",|" \
        -e "s|\"description\":.*|\"description\": \"$PROJECT_DESCRIPTION\",|" \
        package.json
fi
