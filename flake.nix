{
  description = "Cost function genie";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs = { self, nixpkgs }:
  let
    system = "x86_64-linux";
    pkgs = nixpkgs.legacyPackages.${system};
  in {
    devShells.${system}.default = pkgs.mkShell {
      buildInputs = with pkgs; [
        go # Compiler
        gopls # LSP
        delve # Debugger
      ];

	  shellHook = ''
      echo "Starting gopls..."
      ${pkgs.gopls}/bin/gopls > /tmp/gopls.log 2>&1 &
	  '';
    };
  };
}
