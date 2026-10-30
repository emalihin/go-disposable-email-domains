# go-disposable-email-domains

Go package for detecting disposable (throwaway) email addresses, using the
community-maintained list from
[disposable-email-domains/disposable-email-domains](https://github.com/disposable-email-domains/disposable-email-domains).

The list is embedded in the package, so there is no network access or file I/O
at runtime. A GitHub Action syncs the upstream list daily and, when it changes,
commits it, tags a new patch release, and publishes it to the Go module proxy.

## Install

```sh
go get github.com/emalihin/go-disposable-email-domains@latest
```

Re-run the command (or use Dependabot/Renovate) to pick up newer lists.

## Usage

```go
import disposable "github.com/emalihin/go-disposable-email-domains"

disposable.IsDisposable("someone@0-mail.com")     // true
disposable.IsDisposable("someone@sub.0-mail.com") // true (parent domains match)
disposable.IsDisposable("someone@gmail.com")      // false

disposable.IsDisposableDomain("0-mail.com")       // true
disposable.Domains()                              // sorted []string of all domains
```

Matching is case-insensitive and ignores a trailing dot. `IsDisposable` uses the
part after the last `@`; it does not validate the address otherwise.

## Updating manually

```sh
./scripts/update.sh && go test ./...
```

## License

The domain list is published by its maintainers under
[CC0 1.0](https://github.com/disposable-email-domains/disposable-email-domains/blob/main/LICENSE.txt).
