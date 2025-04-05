swagger:
	swagger generate server -f ./docs/api.yaml -t ./internal/api --main-package ../../../ -A bysykkel

run:
	@env -S "`cat $(wildcard ./.env)`" go run main.go

generate-mocks:
	@mockgen -source=./internal/api/handlers/stations.go -destination=./internal/api/handlers/mocks/bysykkel_mock.go -package=mocks

test:
	go test -v -bench -coverpkg=./internal/...,./cmd/... -coverprofile=coverage.cov ./... -json -v | tparse
