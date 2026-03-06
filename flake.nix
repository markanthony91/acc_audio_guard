{
  description = "ACC Audio Guard (ORFEU) development shell";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      systems = [ "x86_64-linux" "aarch64-linux" ];
      forAllSystems = f: nixpkgs.lib.genAttrs systems (system: f system);
    in
    {
      devShells = forAllSystems (system:
        let
          pkgs = import nixpkgs {
            inherit system;
            config.allowUnfree = true;
          };
        in
        {
          default = pkgs.mkShell {
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
              echo "[orfeu] nix shell pronta"
              echo "Comandos previstos:"
              echo "  go run ./cmd/orfeu"
              echo "  go test ./..."
              echo "  go run ./cmd/orfeu --dry-run=false --duration-sec=10 (requer build com tag rnnoise e lib disponível)"
              echo "  wails build -platform windows/amd64"
            '';
          };
        });
    };
}
