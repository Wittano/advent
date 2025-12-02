{ mkShell
, go
, gotools
, act
, nixd
, nixpkgs-fmt
, gcc
, pcre
}: mkShell {
  hardeningDisable = [ "all" ];

  GOROOT = "${go}/share/go";
  DEBUG_ARGS = "";

  nativeBuildInputs = [
    # GO
    go
    gotools

    # Nix
    nixpkgs-fmt
    nixd

    # Github actions
    act
    gcc
    pcre
  ];
}
