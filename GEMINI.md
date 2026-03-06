# ACC Audio Guard - GEMINI.md
**Codinome: ORFEU** (O Guardião da Harmonia Sonora)

Este agente é uma solução em Go para cancelamento de ruído local (offline-first) em ambientes críticos Burger King (Windows 11 22H2).

## 🚀 Padrões de Implementação
- **Tecnologia:** Go (Golang) + CGO (RNNoise).
- **Lógica:** Redes neurais leves para supressão de ruído ambiental.
- **Interface:** Wails v2 (GUI moderna e minimalista).
- **Modo:** Windows Service (Background Process) + GUI (System Tray).
- **GitHub:** SEMPRE preencher a descrição ("About") do repositório via `gh repo edit`.
- **Objetivo:** Filtrar áudio de microfones físicos (Jabra/Logitech) e injetar em drivers virtuais (VB-CABLE).

## 📊 Estatísticas
| Métrica | Valor |
|---------|-------|
| Versão | 0.1.0 |
| Latência Alvo | < 20ms |
| RAM Alvo | ~10-15MB |
| Testes | Unitários e Stress de Áudio |

## 🗺️ Roadmap & Progresso
- [x] Definição de arquitetura e tecnologia (Fase 1).
- [ ] Implementação do wrapper RNNoise em Go.
- [ ] Teste de conceito de áudio em tempo real.
- [ ] GUI Wails para configuração de dispositivos.
- [ ] Empacotamento para Windows.

---
*Atualizado em 06 de Março de 2026*
