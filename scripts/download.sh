#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$SCRIPT_DIR/.."
DOWNLOAD_DIR="$PROJECT_DIR/download"
mkdir "$DOWNLOAD_DIR"
wget https://github.com/elko-dev/spawn/releases/download/v0.5.0/spawn-darwin-386 -O "$DOWNLOAD_DIR/spawn"