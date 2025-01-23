#!/bin/bash

# Create the base folders
mkdir -p solutions/easy
mkdir -p solutions/medium
mkdir -p solutions/hard

# Create top-level files
cat <<EOL > README.md
# LeetCode Solutions

This repository contains my solutions to LeetCode problems written in Go.

## Structure

The solutions are categorized by difficulty:

- **Easy**: \`solutions/easy\`
- **Medium**: \`solutions/medium\`
- **Hard**: \`solutions/hard\`

Each solution has:
- \`main.go\`: The Go implementation.
- \`main_test.go\`: Unit tests for the solution.
- \`README.md\`: Explanation of the solution.

## How to Run

1. Navigate to a solution's directory.
2. Run the solution: \`go run main.go\`
3. Run tests: \`go test\`

---

Happy coding!
EOL

# Create .gitignore
cat <<EOL > .gitignore
# Binaries
*.exe
*.out

# Vendor directory
vendor/

# IDE files
.idea/
.vscode/
EOL

# Create a placeholder LICENSE file (MIT License template)
cat <<EOL > LICENSE
MIT License

Copyright (c) $(date +%Y)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
EOL

echo "Repository structure initialized successfully."

