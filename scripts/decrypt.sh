#!/usr/bin/env bash

set -euo pipefail

# decrypt a file called env.ci using the encryption key with environment variable GITLAB_CI_ENV_KEY
function decrypt {
    openssl enc -d -aes-256-cbc -pbkdf2 -in env.ci -k $GITLAB_CI_ENV_KEY
}