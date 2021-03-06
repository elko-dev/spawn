#!/usr/bin/env bash

set -euo pipefail
VERSION=$1
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$SCRIPT_DIR/.."
DOWNLOAD_DIR="$PROJECT_DIR/download"
mkdir "$DOWNLOAD_DIR"
wget "https://github.com/elko-dev/spawn/releases/download/$VERSION/spawn-darwin-386" -O "$DOWNLOAD_DIR/spawn"

chmod +x "$DOWNLOAD_DIR/spawn"
