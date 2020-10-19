# docker-nv2
Docker CLI plugin to demonstrate Notary v2 integration with Docker CLI.

## Build and Install
This plugin requires [golang](https://golang.org/dl/) with version `>= 1.15`.

To build and install, run
```
go build -o ~/.docker/cli-plugins/docker-nv2 ./cmd/docker-nv2
```

## Instructions
### Create Alias
For better demonstration experience, it is suggested to create the following alias in your shell:
```bash
alias docker="docker nv2"
```
or if you are using PowerShell on Windows,
```powershell
function docker { cmd /c docker nv2 $args }
```

### Example Run
On the producer machine:
```bash
docker notary --enabled
docker build -t $image .
docker notary sign --key identity.pem --cert identity.crt $image
docker push $image
```

On the consumer machine:
```bash
docker notary --enabled
docker pull $image
```
It may fail since the producer machine may use a self-signed certificate, or invalid certificates detected.
