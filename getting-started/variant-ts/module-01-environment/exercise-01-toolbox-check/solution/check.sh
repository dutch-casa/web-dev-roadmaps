#!/bin/bash

# Toolbox check
#
# Run this script after installing everything.
# Every check should print a version number.
# If any check says "not found," go back and install that tool.

echo "=== Toolbox Check ==="
echo ""

echo "Bun:"
bun --version
echo ""

echo "Git:"
git --version
echo ""

echo "GitHub CLI:"
gh --version
echo ""

echo "VS Code:"
code --version
echo ""

echo "=== All done ==="
echo ""
echo "If everything above shows a version number, you're ready."
echo "If anything says 'command not found', install it before moving on."
