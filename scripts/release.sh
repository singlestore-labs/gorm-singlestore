#!/bin/bash

# Check if version argument is provided
if [ $# -eq 0 ]; then
    echo "Usage: $0 <version>"
    exit 1
fi

version=$1

# Check if GitHub CLI is installed
if ! command -v gh &> /dev/null
then
    echo "GitHub CLI (gh) is not installed. Please install it to proceed."
    exit 1
fi

# Create a release tag
git tag v$version -m "Release version $version"
git push origin v$version

# Create a release
gh release create v$version --generate-notes

# Publish the package
GOPROXY=proxy.golang.org go list -m github.com/singlestore-labs/gorm-singlestore@v$version

echo "Package released successfully for version $version"