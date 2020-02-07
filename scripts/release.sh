#!/bin/bash

set -euo pipefail

VERSION=$1
USER=$2
APP_NAME=$3
REPO=$4


SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$SCRIPT_DIR/.."
BIN_OUTPUT="$PROJECT_DIR/release"

github_release() {
    echo "Starting release $VERSION for $REPO $APP_NAME"
    
    github-release release --user "$REPO" \
    --repo "$APP_NAME" --tag "$VERSION" \
    --name "Release $VERSION" \
    --description "$GIT_COMMIT_DESC" \
    --pre-release
    
    echo "Created release $VERSION"

    for GOOS in darwin linux windows; do
        for GOARCH in 386 amd64; do
            ARTIFACT_NAME="$APP_NAME-$GOOS-$GOARCH"
            ARTIFACT_PATH="$BIN_OUTPUT/$ARTIFACT_NAME"
            echo "Releasing artifact $ARTIFACT_PATH"
            github-release upload --user "$REPO" \
                --repo "$APP_NAME" --tag "$VERSION" \
                --name "$ARTIFACT_NAME" \
                --file "$ARTIFACT_PATH"
        done
    done
}

check_to_release() {
    GIT_COMMIT_DESC=$(git log --format=oneline -n 1)
    echo "$GIT_COMMIT_DESC"

    if [[ $GIT_COMMIT_DESC == *"RELEASE"* ]] ; then
        echo "Releasing Spawn..."
        create_tag
        github_release
        
    else 
        echo "No release on this commit"
    fi
}

create_tag() {
    git tag "$VERSION" && git push --tags
    echo "Created git tag"
}

check_to_release
