{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  packages = with pkgs; [
    git
    gnumake
    go
    gopls
    gotools
    delve
    gcc
    pkg-config
    rnnoise
    portaudio
    nodejs_22
    wails
  ];

  CGO_ENABLED = "1";
  PKG_CONFIG_PATH = pkgs.lib.makeSearchPath "lib/pkgconfig" [ pkgs.rnnoise pkgs.portaudio ];

  shellHook = ''
    echo "[orfeu] nix-shell pronta"
    echo "Comandos previstos disponíveis: go, gofmt, go test, pkg-config, wails"
  '';
}
