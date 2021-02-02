module github.com/sajayantony/nv2-demo

go 1.15

require (
	github.com/docker/cli v0.0.0-20200915230204-cd8016b6bcc5
	github.com/docker/distribution v2.7.1+incompatible
	github.com/docker/docker v1.13.1 // indirect
	github.com/docker/docker-credential-helpers v0.6.3 // indirect
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7
	github.com/notaryproject/notary/v2 v2.0.0-00010101000000-000000000000
	github.com/opencontainers/go-digest v1.0.0
	github.com/opencontainers/image-spec v1.0.1
	github.com/opencontainers/runc v0.1.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/shizhMSFT/docker-generate v0.1.0
	github.com/sirupsen/logrus v1.7.0 // indirect
	github.com/urfave/cli/v2 v2.2.0
	gotest.tools/v3 v3.0.3 // indirect
)

replace github.com/notaryproject/notary/v2 => github.com/shizhMSFT/notary/v2 v2.0.0-20210202124147-4f478202d15a
