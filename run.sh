#!/usr/bin/env bash
cd "$(dirname -- "$0")" || exit 1

_compile_run() {
    cd "$1" || exit 1
    go build -o out main.go && ./out
}

if [[ "$1" =~ [0-9]+ ]]; then
    _compile_run "$1"
else
    read -ra PARTS < <(find . -maxdepth 1 -type d -regex '.*[0-9]+$')
    for part in "${PARTS[@]}"; do
        _compile_run "$part"
    done
fi
