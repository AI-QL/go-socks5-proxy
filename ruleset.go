package main

import (
	"context" // Importing the context package for context management
	"regexp"  // Importing the regexp package for regular expression matching

	"github.com/AI-QL/go-socks5" // Importing the go-socks5 package for SOCKS5 proxy functionality
)

// PermitDestAddrPattern returns a RuleSet which selectively allows addresses
// based on a regular expression pattern.
func PermitDestAddrPattern(pattern string) socks5.RuleSet {
	return &PermitDestAddrPatternRuleSet{pattern} // Return a new instance of PermitDestAddrPatternRuleSet
}

// PermitDestAddrPatternRuleSet is an implementation of the RuleSet which
// enables filtering supported destination addresses based on a regular expression pattern.
type PermitDestAddrPatternRuleSet struct {
	AllowedFqdnPattern string // Regular expression pattern for allowed FQDNs
}

// Allow is a method of PermitDestAddrPatternRuleSet that determines whether a request should be allowed.
// It checks if the destination FQDN matches the allowed pattern and returns the result.
func (p *PermitDestAddrPatternRuleSet) Allow(ctx context.Context, req *socks5.Request) (context.Context, bool) {
	// Use regular expression matching to check if the destination FQDN matches the allowed pattern
	match, _ := regexp.MatchString(p.AllowedFqdnPattern, req.DestAddr.FQDN)
	return ctx, match // Return the context and the result of the match
}