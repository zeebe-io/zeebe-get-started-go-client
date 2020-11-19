#!/bin/bash

set -ex -o pipefail

echo "$INPUT_VERSION"
echo "$INPUT_PUSH_CHANGES"

if [ -z "$INPUT_VERSION" ] | [ -z "$INPUT_PUSH_CHANGES" ]; then
    echo -e "Please provide the version and should push changes\nExample: $0 0.24.1 true"
    exit 1
fi

./update-version.sh "$INPUT_VERSION"

if [ "$INPUT_PUSH_CHANGES" = true ] ; then
    git config --global user.email "github-actions[bot]@users.noreply.github.com"
    git config --global user.name "github-actions[bot]"
    git add go.mod
    git add go.sum
    git commit -m "chore(project): update zeebe version to $INPUT_VERSION"
    git push
fi
