{
  description = "Multi Architecture Nix Flake for GO development";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
      ...
    }@inputs:

    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs {
          inherit system;
          config.allowUnfree = true;
        };

        packages = import ./nix/packages.nix {
          inherit
            pkgs
            self
            system
            ;
        };

        devShell = import ./nix/devshell.nix {
          inherit pkgs;
        };

      in
      {
        devShells.default = devShell;
        packages = packages;
      }
    );
}
