# Go example app using Gin web framework



### Building

```bash
go build -o api
./api
```

### Accessing the app

Either via browser on <http://localhost:8080> or via `curl`:

```bash
curl -X GET http://localhost:8080/api/accounts
curl -X POST --header "Content-Type: application/json" --data '{"name":"Test Account","email":"test@foo.com"}' http://localhost:8080/api/accounts
curl -X GET http://localhost:8080/api/accounts/1
```

## Development

### Setup environment

- Install golang 1.22.x
- Install wire `go install github.com/google/wire/cmd/wire@latest`

### Update dependency wiring (if needed)

```bash
go generate ./...
```

## Testing

```bash
go test ./... -coverprofile=coverage.out
```

## Deploy to Tanzu Platform

```bash
kubectl apply -f tekton-pipeline.yaml
tanzu apps workload apply gin-gonic-api --git-repo https://github.com/markusrt/gin-gonic-api --git-branch main --type web --label apps.tanzu.vmware.com/has-tests=true --tail
tanzu apps workload apply go-gin --git-repo https://github.com/markusrt/go-gin --git-branch main --type web --build-env "BP_KEEP_FILES=templates/*"  --label apps.tanzu.vmware.com/has-tests=true --tail
```

tanzu apps workload create go-gin --app go-gin --type web --image ghcr.io/markusrt/gin-gonic-api:main@sha256:78a374ff16d3153fd19463aad878cab819ae9d27b3db2d862aabd04607ae37e1