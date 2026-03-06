# ACC Audio Guard (ORFEU) 🎤🛡️

**Codinome: ORFEU** (O Guardião da Harmonia Sonora)

O **ACC Audio Guard** é uma solução em Go para cancelamento de ruído 100% offline, projetada para ambientes críticos de Drive-Thru. O objetivo é substituir dependências de nuvem por um motor local com baixa latência.

## Status da Fase 1

- [x] Esqueleto da aplicação CLI com ciclo de vida e shutdown gracioso.
- [x] Pipeline de processamento em tempo real (ticker por frame) com fonte sintética para validação.
- [x] Métricas de latência por execução (frames, média, máximo e total de processamento).
- [x] Wrapper RNNoise preparado via CGO com build tag `rnnoise`.
- [ ] Integração de captura de microfone físico (backend real de áudio).
- [ ] Validação de latência fim-a-fim com áudio real (`< 20ms`).

## Arquitetura Inicial

1. **Entrada (`cmd/orfeu`)**: bootstrap e sinais do sistema.
2. **App (`internal/app`)**: valida configuração e orquestra runtime.
3. **Pipeline (`internal/audio`)**: loop real-time por frame + estatísticas.
4. **RNNoise (`internal/rnnoise`)**: contrato + implementação CGO opcional.
5. **Dispositivos (`internal/device`)**: stub para descoberta de hardware.
6. **Logging (`internal/logx`)**: logs com módulo e `hostname`.

## Ambiente Nix (todos os comandos previstos)

Este projeto já inclui:
- `flake.nix` para `nix develop`
- `shell.nix` para `nix-shell`

Ambos expõem os comandos: `go`, `gofmt`, `go test`, `pkg-config`, `wails`, toolchain CGO e libs (`rnnoise`, `portaudio`).

```bash
# Flakes
nix develop

# Sem flakes
nix-shell
```

## Comandos de Desenvolvimento

```bash
# Rodar em modo padrão (dry-run com fonte sintética)
go run ./cmd/orfeu --duration-sec=10

# Rodar testes
go test ./...

# Verificar RNNoise no shell
pkg-config --modversion rnnoise

# Build com RNNoise real (requer librnnoise no ambiente)
go run -tags rnnoise ./cmd/orfeu --dry-run=false --duration-sec=10

# Build GUI (fase futura)
wails build -platform windows/amd64
```

## Estrutura

```text
acc_audio_guard/
  cmd/orfeu/main.go
  internal/app/app.go
  internal/audio/pipeline.go
  internal/audio/pipeline_test.go
  internal/audio/source.go
  internal/config/config.go
  internal/config/config_test.go
  internal/device/device.go
  internal/logx/logx.go
  internal/rnnoise/types.go
  internal/rnnoise/engine_stub.go
  internal/rnnoise/engine_cgo.go
  flake.nix
  shell.nix
```

## Próximos Passos (Fase 1 restante)

1. Integrar captura/injeção real com PortAudio ou Miniaudio.
2. Rodar benchmark com microfone físico + VB-CABLE.
3. Fechar relatório de latência de ponta a ponta com meta `< 20ms`.

## Roadmap (Fases seguintes)

### Fase 2: Integração com Sistema

1. Modo **headless** para execução contínua sem CLI interativa.
2. Execução como serviço do sistema (Windows Service).
3. API local de controle e status (localhost/IPC) para integração com outros módulos.

### Fase 3: Interface e Operação

1. GUI Wails para configuração de dispositivos e diagnóstico.
2. Health checks e telemetria local para suporte operacional.
3. Empacotamento e instalação silenciosa para rollout em escala.

---
*Projeto integrante do ecossistema Sistemas - Marcelo*
