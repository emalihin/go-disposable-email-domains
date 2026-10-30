// Package disposable reports whether an email address or domain belongs to a
// known disposable email provider.
//
// The domain list is sourced from
// https://github.com/disposable-email-domains/disposable-email-domains and is
// embedded at build time. It is refreshed daily and released as a new patch
// version, so upgrading the module picks up the latest list.
package disposable

import (
	_ "embed"
	"slices"
	"strings"
)

//go:embed disposable_email_blocklist.conf
var blocklist string

var domains = parse(blocklist)

func parse(s string) map[string]struct{} {
	m := make(map[string]struct{}, strings.Count(s, "\n")+1)
	for _, line := range strings.Split(s, "\n") {
		line = strings.ToLower(strings.TrimSpace(line))
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		m[line] = struct{}{}
	}
	return m
}

// IsDisposable reports whether the domain part of email is a disposable
// email domain. Inputs without an "@" are treated as invalid and return false.
func IsDisposable(email string) bool {
	i := strings.LastIndexByte(email, '@')
	if i < 0 {
		return false
	}
	return IsDisposableDomain(email[i+1:])
}

// IsDisposableDomain reports whether domain, or any of its parent domains,
// is a disposable email domain. Matching is case-insensitive.
func IsDisposableDomain(domain string) bool {
	domain = strings.TrimSuffix(strings.ToLower(strings.TrimSpace(domain)), ".")
	for domain != "" {
		if _, ok := domains[domain]; ok {
			return true
		}
		i := strings.IndexByte(domain, '.')
		if i < 0 {
			return false
		}
		domain = domain[i+1:]
	}
	return false
}

// Domains returns a sorted copy of all known disposable domains.
func Domains() []string {
	out := make([]string, 0, len(domains))
	for d := range domains {
		out = append(out, d)
	}
	slices.Sort(out)
	return out
}
