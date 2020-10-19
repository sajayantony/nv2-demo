package docker

import (
	"os/exec"

	"github.com/docker/distribution/manifest/schema2"
	"github.com/opencontainers/go-digest"
	oci "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/shizhMSFT/docker-generate/pkg/manifest"
)

// GenerateManifest generate manifest from docker save
func GenerateManifest(reference string) ([]byte, error) {
	cmd := exec.Command("docker", "save", reference)
	reader, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	manifest, err := manifest.GenerateSchema2FromDockerSave(reader)
	if err != nil {
		return nil, err
	}
	_, payload, err := manifest.Payload()
	return payload, err
}

// GenerateManifestOCIDescriptor generate manifest descriptor from docker save
func GenerateManifestOCIDescriptor(reference string) (oci.Descriptor, error) {
	manifest, err := GenerateManifest(reference)
	if err != nil {
		return oci.Descriptor{}, err
	}
	return oci.Descriptor{
		MediaType: schema2.MediaTypeManifest,
		Digest:    digest.FromBytes(manifest),
		Size:      int64(len(manifest)),
	}, nil
}
