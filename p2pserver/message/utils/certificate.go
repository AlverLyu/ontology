package utils

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func verifyCert(certPEM string) error {
	const rootPEM = `
-----BEGIN CERTIFICATE-----
MIIBjDCCATOgAwIBAgIUaJfiACtw29PSuHPVhd9fwt6HwVUwCgYIKoZIzj0EAwIw
IzEhMB8GA1UEAxMYZmFicmljLWNhLXNlcnZlci1yb290Y2ExMB4XDTE5MDUzMDA2
NTcwMFoXDTM0MDUyNjA2NTcwMFowIzEhMB8GA1UEAxMYZmFicmljLWNhLXNlcnZl
ci1yb290Y2ExMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEjW+kCikul2GyL4qe
Yaks5FfvbO06lBUeN/SvcQV/o5QnBh0EqzXittRGkK98IfhHyKyZg/cxoQVc3WoG
WauwVKNFMEMwDgYDVR0PAQH/BAQDAgEGMBIGA1UdEwEB/wQIMAYBAf8CAQAwHQYD
VR0OBBYEFCNxRMFl7acm295rVmxsIO/yiwHdMAoGCCqGSM49BAMCA0cAMEQCIFrq
b8WLu4Yq7nP9mhV2hyPyQVImV786HDaAW4v8SQLJAiA0jOoCeXllhrdFaJX7iM2m
GmSgyzbIM3Z08Ah5hbf44g==
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
