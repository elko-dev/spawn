#!/usr/bin/env bash

set -euo pipefail

function encrypt {
    openssl enc -aes-256-cbc -pbkdf2 -e -in local.ci -out env.ci -k "1234"
}

encrypt