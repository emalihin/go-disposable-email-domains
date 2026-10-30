#!/usr/bin/env bash
# Downloads the latest upstream blocklist into the repo root.
set -euo pipefail

URL="https://raw.githubusercontent.com/disposable-email-domains/disposable-email-domains/main/disposable_email_blocklist.conf"
DEST="$(cd "$(dirname "$0")/.." && pwd)/disposable_email_blocklist.conf"

tmp="$(mktemp)"
trap 'rm -f "$tmp"' EXIT

curl -fsSL --retry 3 "$URL" -o "$tmp"

# Guard against truncated or empty downloads replacing a good list.
lines="$(wc -l < "$tmp")"
if [ "$lines" -lt 1000 ]; then
  echo "downloaded list has only $lines lines; refusing to update" >&2
  exit 1
fi

mv "$tmp" "$DEST"
trap - EXIT
echo "updated $DEST ($lines lines)"
