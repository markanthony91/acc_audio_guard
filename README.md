# ACC Audio Guard (ORFEU) 🎤🛡️

**Codinome: ORFEU** (O Guardião da Harmonia Sonora)

O **ACC Audio Guard** é uma solução em Go para cancelamento de ruído 100% offline, projetada especificamente para ambientes críticos de Drive-Thru no Burger King. Ele substitui dependências de nuvem (como o Krisp) por um motor de processamento local baseado em redes neurais leves.

## 🚀 Visão Geral
Em ambientes de alta rotatividade e instabilidade de rede, a dependência de autenticação em nuvem para ferramentas de áudio é um risco operacional. O **ORFEU** atua como um driver de áudio inteligente que filtra ruídos de motores, vento e cozinha, entregando voz limpa para o sistema de VoIP (MicroSIP), sem nunca precisar de internet.

## 🏗️ Arquitetura Técnica
O sistema funciona como uma ponte de áudio entre o hardware físico e um driver virtual:

1.  **Captura (Input):** O motor em Go captura o áudio bruto do microfone físico (Jabra/Logitech).
2.  **Processamento (Engine):** O áudio passa pela biblioteca **RNNoise** (Recurrent Neural Network Noise Suppression), que limpa o ruído preservando a voz.
3.  **Injeção (Output):** O áudio limpo é injetado em um **Cabo de Áudio Virtual** (VB-CABLE).
4.  **Consumo:** O MicroSIP é configurado para escutar o "Cabo Virtual", recebendo o áudio já tratado.

## 🛠️ Stack Tecnológica
- **Linguagem:** Go (Golang) 1.23+
- **Filtro:** RNNoise (C-based RNN) via CGO.
- **Interface:** Wails (Go + Vite/React) para a GUI de seleção de dispositivos.
- **Áudio:** PortAudio ou Miniaudio para baixa latência.
- **SO Alvo:** Windows 10/11 (BK Kiosks).

## 🗺️ Roadmap de Desenvolvimento

### Fase 1: Fundação & Engine (Atual) 🏗️
- [ ] Implementar wrapper Go para RNNoise (CGO).
- [ ] Teste de captura e processamento em tempo real (CLI).
- [ ] Validação de latência (Meta: < 20ms).

### Fase 2: Roteamento de Áudio 🎧
- [ ] Integração com PortAudio/Miniaudio.
- [ ] Mapeamento automático de dispositivos de entrada/saída.
- [ ] Verificação de presença do Driver Virtual (VB-CABLE).

### Fase 3: Interface de Usuário (GUI) 🖥️
- [ ] Desenvolvimento da GUI via Wails.
- [ ] Seletores dinâmicos de Microfone/Saída.
- [ ] Indicador visual de "Noise Floor" e "Voice Activity".

### Fase 4: Estabilização & Deploy 🛡️
- [ ] Implementação como Windows Service (Daemon).
- [ ] Monitoramento de saúde (Auto-restart em caso de falha de driver).
- [ ] Empacotamento para instalação silenciosa (MSI/EXE).

## 🛠️ Comandos de Desenvolvimento
```bash
# Entrar no ambiente
nix develop

# Executar em modo desenvolvimento
go run main.go

# Build para Windows
wails build -platform windows/amd64
```

---
*Projeto integrante do ecossistema Sistemas - Marcelo*
