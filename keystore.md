keytool -genkey -v \
-keystore ${RELEASE_KEY}.keystore \
-alias ${ALIAS_NAME} -keyalg RSA \
-keysize 2048 -validity 10000
