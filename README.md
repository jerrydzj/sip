# shici

`shici` is a Go CLI for mood-based Chinese poetry.

## Development

Run directly:

```bash
go run ./cmd/shici <mood>
```

Build binary:

```bash
go build -o bin/shici ./cmd/shici
./bin/shici <mood>
```

Tooling targets:

```bash
make fmt
make test
make vet
make lint
make build
```
