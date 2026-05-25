{
  description = "Multi Architecture Nix Flake for PHP development";

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

        mkScript =
          name: text:
          let
            script = pkgs.writeShellScriptBin name text;
          in
          script;

        scripts = [ ];

        devPackages = with nixpkgs; [
          pkgs.go
          pkgs.gopls
        ];

      in
      {
        devShells = {
          default = pkgs.mkShell {
            name = "go-dev-shell";
            nativeBuildInputs = scripts;
            packages = devPackages;
            shellHook = ''
              export GOPATH=$PWD/.gopath
              export PATH=$GOPATH/bin:$PATH

              mkdir -p $GOPATH
              echo "Go dev shell ready"
            '';
          };
        };
      }
    );
}
