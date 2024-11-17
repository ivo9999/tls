package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "it's working")
}

func showCommonName(w http.ResponseWriter, req *http.Request) {
	var commonName string
	if req.TLS != nil && len(req.TLS.PeerCertificates) > 0 {
		commonName = req.TLS.VerifiedChains[0][0].Subject.CommonName
	}
	fmt.Fprintf(w, "your common name: %s\n", commonName)
}

func main() {
	caBytes, err := os.ReadFile("ca.crt")
	if err != nil {
		return
	}
	ca := x509.NewCertPool()
	if !ca.AppendCertsFromPEM(caBytes) {
		return
	}

	server := http.Server{
		Addr: ":443",
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequireAndVerifyClientCert,
			ClientCAs:  ca,
			MaxVersion: tls.VersionTLS13,
		},
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/common-name", showCommonName)
	err = server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		log.Fatal("error: ", err)
	}
}
