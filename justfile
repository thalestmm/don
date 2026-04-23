_help:
    @just --list

[group('dev')]
watch:
    @air -c config/air.toml

[group('dev')]
push message="chore: update":
    @git add .
    @git commit -m "{{ message }}"
    @git push

[group('ci')]
fmt:
    @go fmt ./...

[group('ci')]
lint: fmt
    @golangci-lint run

# Run SQLC code generation
[group('ci')]
gen:
    @cp config/sqlc.yaml .
    @sqlc generate
    @rm sqlc.yaml

[group('dev')]
[group('docker')]
start-env:
    @docker compose -f deployments/docker-compose.local.yml up -d

[group('dev')]
[group('docker')]
stop-env:
    @docker compose -f deployments/docker-compose.local.yml down

[group('dev')]
[group('docker')]
docker-build:
    @docker build -t don:dev . -f build/Dockerfile
