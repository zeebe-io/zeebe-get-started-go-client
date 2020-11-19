#!/bin/bash

set -ex -o pipefail

INPUT_VERSION="$1"
INPUT_PUSH_CHANGES="$2"

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

    linesChanged=$(git status --porcelain | wc -l)

    if [ "$linesChanged" -gt 0 ]; then        
        git add go.mod
        git add go.sum
        git add README.md

        git commit -m "chore(project): update zeebe version to $INPUT_VERSION"
        git push
    fi
fi
