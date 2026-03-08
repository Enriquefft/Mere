{
  description = "Mere - Minimal, module-first architecture toolkit for Claude Code";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gopls
            gotools
            go-tools
            just
          ];

          shellHook = ''
            echo "🔧 Mere development environment"
            echo "Go version: $(go version)"
            echo ""
            echo "Available commands:"
            echo "  just build  - Build the mere binary"
            echo "  just test   - Run tests"
            echo "  just install - Install to GOBIN"
            echo "  just help    - Show all commands"
          '';
        };
      }
    );
}
