package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

func main() {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serial, _ := rand.Int(rand.Reader, max)
	certTmpl := x509.Certificate{
		SerialNumber: serial,
		Subject: pkix.Name{
			Organization:       []string{"Oak Labs"},
			OrganizationalUnit: []string{"R&D"},
			CommonName:         "chang.lab.oak.san",
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().Add(10 * 365 * 24 * time.Hour),
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses: []net.IP{net.ParseIP("127.0.0.1")},
	}

	// Create private key and save to file
	pvtKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	pvtKeyFile, _ := os.Create("key.pem")
	_ = pem.Encode(pvtKeyFile, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(pvtKey),
	})

	// Create certificate using private key and save to file
	derBytes, _ := x509.CreateCertificate(rand.Reader, &certTmpl, &certTmpl, &pvtKey.PublicKey, pvtKey)
	certFile, _ := os.Create("cert.pem")
	_ = pem.Encode(certFile, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derBytes,
	})

	_ = pvtKeyFile.Close()
	_ = certFile.Close()
}
