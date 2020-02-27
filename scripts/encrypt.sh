#!/usr/bin/env bash

set -euo pipefail
    echo "Started"

function encrypt {
    echo "Started"

    openssl enc -aes-256-cbc -pbkdf2 -e -in local.ci -out env.ci -k "1234"
    echo "Finished"

}

encrypt