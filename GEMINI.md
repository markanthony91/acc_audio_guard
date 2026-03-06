# ACC Audio Guard - GEMINI References

## Fonte principal
- [README.md](./README.md)

## Entradas e runtime
- [cmd/orfeu/main.go](./cmd/orfeu/main.go)
- [internal/app/app.go](./internal/app/app.go)

## Pipeline e áudio
- [internal/audio/pipeline.go](./internal/audio/pipeline.go)
- [internal/audio/source.go](./internal/audio/source.go)
- [internal/audio/pipeline_test.go](./internal/audio/pipeline_test.go)

## Configuração
- [internal/config/config.go](./internal/config/config.go)
- [internal/config/config_test.go](./internal/config/config_test.go)

## RNNoise
- [internal/rnnoise/types.go](./internal/rnnoise/types.go)
- [internal/rnnoise/engine_stub.go](./internal/rnnoise/engine_stub.go)
- [internal/rnnoise/engine_cgo.go](./internal/rnnoise/engine_cgo.go)
- [internal/rnnoise/engine_stub_test.go](./internal/rnnoise/engine_stub_test.go)

## Ambiente
- [flake.nix](./flake.nix)
- [shell.nix](./shell.nix)
