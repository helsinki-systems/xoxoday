with import <nixpkgs> {};

mkShell {
  name = "xoxoday";
  nativeBuildInputs = [
    go_1_18
    gnumake
  ];

  shellHook = ''
    export GOPATH=$PWD/gopath
  '';
}
