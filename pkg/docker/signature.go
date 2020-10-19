package docker

import (
	"context"
	"net"

	"github.com/docker/distribution/reference"
	"github.com/notaryproject/notary/v2"
	"github.com/notaryproject/notary/v2/registry"
)

// GetSignatureRepository returns a signature repository
func GetSignatureRepository(ctx context.Context, ref string) (notary.SignatureRepository, error) {
	named, err := reference.ParseNamed(ref)
	if err != nil {
		return nil, err
	}
	hostname, repository := reference.SplitHostname(named)

	tr, err := Transport(hostname)
	if err != nil {
		return nil, err
	}

	insecure := false
	if host, _, _ := net.SplitHostPort(hostname); host == "localhost" {
		insecure = true
	}
	client := registry.NewClient(tr, hostname, insecure)

	return client.Repository(ctx, repository), nil
}
