#!/bin/bash

# Fix @ characters in label fields
echo "Fixing @ characters in label fields..."

# Find all contact files with label: @ pattern
for file in contacts-data/*.md; do
    if grep -q "^label: @" "$file"; then
        echo "Fixing: $(basename "$file")"
        # Use sed to remove @ after "label: "
        sed -i '' 's/^label: @/label: /' "$file"
    fi
done

echo "Done! Fixed all label fields."