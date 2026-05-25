{
  pkgs,
}:
let
  mkScript = name: text: pkgs.writeShellScriptBin name text;
  scripts = [ ];

  shellHook = ''
    export GOPATH=$PWD/.gopath
    export PATH=$GOPATH/bin:$PATH
    mkdir -p $GOPATH
  '';

  devPackages = with pkgs; [
    go
    gopls
  ];

in

pkgs.mkShell {
  name = "go-dev-shell";
  nativeBuildInputs = scripts;
  packages = devPackages;
  shellHook = shellHook;
}
