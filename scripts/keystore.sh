#!/usr/bin/env bash

set -euo pipefail

keytool -genkey -keystore my.keystore -keyalg RSA -keysize 2048 \
        -validity 10000 -alias app -dname "cn=Unknown, ou=Unknown, o=Unknown, c=Unknown" \
        -storepass abcdef12 -keypass abcdef12