#!/bin/sh

# This script contains helper functions for pushing commits to remote.
message=$1
if [ -z "$message" ]; then
    echo "Usage: $0 <commit message>"
    exit 1
fi

git add .
git commit -m "Exercises: $message"
git push origin main
