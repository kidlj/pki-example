package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	pool := x509.NewCertPool()
	// This works
	// caCertPaths := []string{"../../pki-example/ca/root-ca.crt", "../../pki-example/ca/signing-ca.crt"}

	// This works too
	caCertPaths := []string{"../../../pki-example/ca/signing-ca.crt"}

	// This is not adequate
	// caCertPaths := []string{"../../pki-example/ca/root-ca.crt"}

	for _, caCertPath := range caCertPaths {
		caCrt, err := ioutil.ReadFile(caCertPath)
		if err != nil {
			fmt.Println("ReadFile err:", err)
			return
		}

		pool.AppendCertsFromPEM(caCrt)
	}

	cliCrt, err := tls.LoadX509KeyPair("../../../pki-example/certs/simple.org.crt", "../../../pki-example/certs/simple.org.key")
	if err != nil {
		fmt.Println("loadx509keypair err:", err)
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{cliCrt},
			RootCAs:      pool,
			// InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://www.simple.org:8080")

	if err != nil {
		// handle error
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
