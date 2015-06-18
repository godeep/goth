package resolver

import (
	"github.com/benschw/dns-clb-go/clb"
	"github.com/benschw/dns-clb-go/dns"
	"github.com/benschw/dns-clb-go/ttlcache"
)

var defaultResolver clb.LoadBalancer
var defaultTTL int = 60

func GetTTL() int {
	return defaultTTL
}

func SetTTL(ttl int) {
	defaultTTL = ttl
}

// Get default resolver
func Resolver() clb.LoadBalancer {
	if defaultResolver == nil {
		defaultResolver = NewResolver(defaultTTL)
	}
	return defaultResolver
}

// Get new resolver with given ttl
func NewResolver(ttl int) clb.LoadBalancer {
	lib := dns.NewDefaultLookupLib()
	cache := ttlcache.NewTtlCache(lib, ttl)
	return clb.NewRoundRobinClb(cache)
}
