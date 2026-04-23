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

[group('dev')]
[group('docker')]
start-env:
    @docker compose -f deployments/docker-compose.local.yml up -d

[group('dev')]
[group('docker')]
stop-env:
    @docker compose -f deployments/docker-compose.local.yml down
