{
  description = "A flake for my advent of code solutions in Go";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs = {
    nixpkgs,
  }: let
    system = "x86_64-linux";
    pkgs = import nixpkgs {
      inherit system;
    };

    aocCookie = pkgs.writeShellApplication {
      name = "aoc-cookie";

      runtimeInputs = with pkgs; [
        coreutils
        curl
      ];

      text = ''
        OPERATION="''${1:-help}"

        if [[ "$OPERATION" == "help" ]]; then
          echo "Usage: $0 get|invalidate|help" >&2
          echo "" >&2
          echo "  get         Get the session cookie for the current session" >&2
          echo "  invalidate  Invalidate the session cookie for the current session" >&2
          echo "  help        Print this help message" >&2
          exit 0
        fi

        CACHE_DIR="''${XDG_CACHE_HOME:-$HOME/.cache}/advent-of-code"
        mkdir -p "$CACHE_DIR"

        SESSION_FILE="$CACHE_DIR/session"

        if [[ "$OPERATION" == "get" ]]; then
          if [ ! -f "$SESSION_FILE" ]; then
            echo "Please enter your session cookie: " >&2
            read -r SESSION
            echo "$SESSION" > "$SESSION_FILE"
          fi

          cat "$SESSION_FILE"
          exit 0
        fi

        if [[ "$OPERATION" == "invalidate" ]]; then
          rm -f "$SESSION_FILE"
          exit 0
        fi
      '';
    };

    aocInput = pkgs.writeShellApplication {
      name = "aoc-input";

      runtimeInputs = [
        aocCookie
        pkgs.coreutils
        pkgs.curl
      ];

      text = ''
        CACHE_DIR="''${XDG_CACHE_HOME:-$HOME/.cache}/advent-of-code-2024"
        mkdir -p "$CACHE_DIR"

        DAY="$1"
        OUTPUT_FILE="$CACHE_DIR/day$1.txt"

        curl_with_opts() {
          curl -s https://adventofcode.com/2024/day/"$DAY"/input "$@"
        }

        while [ ! -f "$OUTPUT_FILE" ]; do
          COOKIE="$(aoc-cookie get)"

          curl_with_cookie() {
            curl_with_opts -H "Cookie: session=$COOKIE" "$@"
          }

          STATUS_CODE="$(curl_with_cookie -o /dev/null -w '%{http_code}')"
          if [ "$STATUS_CODE" == "400" ] || [ "$STATUS_CODE" == "500" ]; then
            aoc-cookie invalidate
          elif [ "$STATUS_CODE" == "404" ]; then
            exit 1
          else
            curl_with_cookie -o "$OUTPUT_FILE"
            chmod 444 "$OUTPUT_FILE"
          fi
        done

        echo "$OUTPUT_FILE"
      '';
    };

    mkDay = _day: let
      day = "day${toString _day}";
    in
      pkgs.writeShellApplication {
        name = "aoc-${day}";

        runtimeInputs = with pkgs; [
          go
          wl-clipboard
        ];

        text = ''
          cd "${day}"
          go test
          go run main.go "$@"
        '';
      };
  in {
    devShell.${system} = pkgs.mkShell {
      packages = with pkgs; [
        go
        wl-clipboard
      ];
    };

    formatter.${system} = pkgs.alejandra;

    packages.${system} = {
      init = pkgs.writeShellApplication {
        name = "aoc-init";

        runtimeInputs = [
          aocInput
          pkgs.coreutils
        ];

        text = ''
          if [ -d "day$1" ]; then
            echo "Day $1 already exists" >&2
            exit 1
          fi

          INPUT_FILE="$(aoc-input "$1")"

          cp -r template "day$1"

          touch "day$1/example.txt"
          cat "$INPUT_FILE" > "day$1/input.txt"
        '';
      };

      day1 = mkDay 1;
      day2 = mkDay 2;
      day3 = mkDay 3;
    };
  };
}
