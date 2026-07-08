#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."

failures=0

fail() {
  echo "dependency security check failed: $*" >&2
  failures=$((failures + 1))
}

normalize_version() {
  local version="$1"
  version="${version#v}"
  version="${version%%+*}"
  echo "$version"
}

version_at_least() {
  local actual min first
  actual="$(normalize_version "$1")"
  min="$(normalize_version "$2")"
  first="$(printf '%s\n%s\n' "$min" "$actual" | sort -V | head -n 1)"
  [[ "$first" == "$min" ]]
}

check_min_version() {
  local label actual min
  label="$1"
  actual="$2"
  min="$3"
  if ! version_at_least "$actual" "$min"; then
    fail "$label is $actual, expected >= $min"
  fi
}

go_modules="$(mktemp)"
trap 'rm -f "$go_modules"' EXIT
go list -m all > "$go_modules"

go_version() {
  local module="$1"
  awk -v module="$module" '$1 == module { print $2; exit }' "$go_modules"
}

check_go_min() {
  local module min version
  module="$1"
  min="$2"
  version="$(go_version "$module")"
  if [[ -n "$version" ]]; then
    check_min_version "$module" "$version" "$min"
  fi
}

if grep -q '^github.com/docker/docker ' "$go_modules"; then
  fail "github.com/docker/docker is still in selected Go module graph"
fi

check_go_min github.com/docker/cli v29.6.1
check_go_min github.com/gofiber/fiber/v2 v2.52.13
check_go_min github.com/opencontainers/runc v1.3.6
check_go_min go.opentelemetry.io/otel v1.41.0
check_go_min golang.org/x/crypto v0.53.0
check_go_min golang.org/x/image v0.18.0

check_package_lock_min() {
  local file package min versions
  file="$1"
  package="$2"
  min="$3"
  [[ -f "$file" ]] || return 0
  while IFS= read -r version; do
    [[ -n "$version" ]] || continue
    check_min_version "$file node_modules/$package" "$version" "$min"
  done < <(jq -r --arg package "node_modules/$package" '.packages[$package].version // empty' "$file")
}

check_pnpm_lock_min() {
  local file package min
  file="$1"
  package="$2"
  min="$3"
  [[ -f "$file" ]] || return 0
  while IFS= read -r version; do
    [[ -n "$version" ]] || continue
    check_min_version "$file /$package" "$version" "$min"
  done < <(awk -v package="$package" '
    $1 ~ "^/" package "@" {
      version = $1
      sub("^/" package "@", "", version)
      sub(":$", "", version)
      print version
    }
  ' "$file")
}

check_yarn_lock_min() {
  local file package min
  file="$1"
  package="$2"
  min="$3"
  [[ -f "$file" ]] || return 0
  while IFS= read -r version; do
    [[ -n "$version" ]] || continue
    check_min_version "$file $package" "$version" "$min"
  done < <(awk -v package="$package" '
    $0 ~ "^" package "@" {
      in_package = 1
      next
    }
    in_package && $1 == "version" {
      gsub("\"", "", $2)
      print $2
      in_package = 0
    }
    in_package && /^[^ ]/ {
      in_package = 0
    }
  ' "$file")
}

check_js_min() {
  local package min
  package="$1"
  min="$2"
  check_package_lock_min svelte/package-lock.json "$package" "$min"
  check_pnpm_lock_min svelte/pnpm-lock.yaml "$package" "$min"
  check_yarn_lock_min svelte/yarn.lock "$package" "$min"
}

check_js_min esbuild 0.28.1
check_js_min fast-uri 3.1.1
check_js_min form-data 4.0.6
check_js_min qs 6.14.1
check_js_min uuid 11.1.1
check_js_min ws 8.21.0

if (( failures > 0 )); then
  exit 1
fi

echo "Dependency security check passed."
