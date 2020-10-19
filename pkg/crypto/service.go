package crypto

import (
	"crypto/x509"

	"github.com/docker/libtrust"
	"github.com/notaryproject/notary/v2"
	x509nv2 "github.com/notaryproject/notary/v2/signature/x509"
	"github.com/notaryproject/notary/v2/simple"
)

// GetSigningService returns a signing service
func GetSigningService(keyPath, certPath string) (notary.SigningService, error) {
	var (
		key       libtrust.PrivateKey
		certs     []*x509.Certificate
		rootCerts *x509.CertPool
		err       error
	)
	if keyPath != "" {
		key, err = x509nv2.ReadPrivateKeyFile(keyPath)
		if err != nil {
			return nil, err
		}
	}
	if certPath != "" {
		certs, err = x509nv2.ReadCertificateFile(certPath)
		if err != nil {
			return nil, err
		}
		rootCerts = x509.NewCertPool()
		for _, cert := range certs {
			rootCerts.AddCert(cert)
		}
	}
	return simple.NewSigningService(key, certs, certs, rootCerts)
}
