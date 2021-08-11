{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.gopls
    pkgs.goimports

    # keep this line if you use bash
    pkgs.bashInteractive
  ];
}
