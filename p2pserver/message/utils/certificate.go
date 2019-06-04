package utils

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func verifyCert(certPEM string) error {
	const rootPEM = `
-----BEGIN CERTIFICATE-----
MIIBdjCCAR2gAwIBAgIUKf0VsCNrb4KcUorO7H3Sv6cwZzkwCgYIKoZIzj0EAwIw
GDEWMBQGA1UEAxMNRE5BLWNhLXNlcnZlcjAeFw0xOTA1MzAwODQyMDBaFw0zNDA1
MjYwODQyMDBaMBgxFjAUBgNVBAMTDUROQS1jYS1zZXJ2ZXIwWTATBgcqhkjOPQIB
BggqhkjOPQMBBwNCAARIgV6DZNaH0KWe5dbfn7qRkkvjIxkNAzhr+4CWh4yyjknU
6y+0Q/n1V/th1pofJ0vO0NSjJH5pzkKweJmzsvYqo0UwQzAOBgNVHQ8BAf8EBAMC
AQYwEgYDVR0TAQH/BAgwBgEB/wIBADAdBgNVHQ4EFgQUEnyuo+ksVZGnkQTGfuMX
7QFXxcIwCgYIKoZIzj0EAwIDRwAwRAIgRIhyb5wLZ/sAZHFtUUsmO/KT0IeFzzwL
/8bdNTbvWWgCIB4SBqvG9wSi/SgSurPp5zsGXmYft85f3z98lu4e3Pe4
-----END CERTIFICATE-----`

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	if !ok {
		return errors.New("failed to parse root certificate")
	}

	block, _ := pem.Decode([]byte(certPEM))
	if block == nil {
		return errors.New("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return errors.New("failed to parse certificate: " + err.Error())
	}

	opts := x509.VerifyOptions{
		Roots: roots,
	}

	if _, err := cert.Verify(opts); err != nil {
		return errors.New("failed to verify certificate: " + err.Error())
	}

	return nil
}
