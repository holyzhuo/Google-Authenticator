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
