Google Authenticator Go Demo
==============================

This Go Demo can be used to interact with the Google Authenticator mobile app for 2-factor-authentication. This code
can generate secrets, generate codes, validate codes and present a QR-Code for scanning the secret. It implements TOTP
according to [RFC6238](https://tools.ietf.org/html/rfc6238)

For a secure installation you have to make sure that used codes cannot be reused (replay-attack). You also need to
limit the number of verifications, to fight against brute-force attacks. For example you could limit the amount of
verifications to 10 tries within 10 minutes for one IP address (or IPv6 block). It depends on your environment.

Usage:
------

See following example:

```go
package main

import (
	"log"
	"fmt"
	"googleAuthenticator"
)

func main() {
	secret, err := optauth.GenSecretKey()
	if err != nil {
		fmt.Println("Secret generate fail:" + secret)
	} else {
		fmt.Println("Secret is:" + secret)
	}

	otpc := optauth.InitOTPConfig(secret)

	// generate auth url, can generate qrcode, use google auth app scan and bind it
	fmt.Println("Google Charts URL for the QR-Code:", otpc.ProvisionURIWithIssuer("user", "issuer"))


	code := "581466"
	fmt.Println("Checking Code:" + code)

	ok, err := otpc.Authenticate(code)
	if err != nil {
		log.Println(err)
		return
	}

	if ok {
		fmt.Println("auth success")
	} else {
		fmt.Println("auth fail, secret:" + secret + ", code:" + code)
	}
}
```