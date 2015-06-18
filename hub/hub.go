package client

import (
	"bitbucket.org/tamtam-im/tamtam-proto/go/client/pool"
	"bitbucket.org/tamtam-im/tamtam-proto/go/client/resolver"
	"fmt"
	"github.com/benschw/dns-clb-go/clb"
	"gopkg.in/tamtam-im/logx.v4"
	"sync"
)

var defaultTTransportHub *TTransportHub

// Transport pool hub
type TTransportHub struct {
	// Provided resolver
	resolver       clb.LoadBalancer
	transportPools map[string]*pool.TTransportPool
	mu             *sync.Mutex
}

// Default transport hub with default resolver
func DefaultTTransportHub() *TTransportHub {
	if defaultTTransportHub == nil {
		defaultTTransportHub = NewTTransportHub(resolver.Resolver())
	}
	return defaultTTransportHub
}

// Create new transport hub
func NewTTransportHub(resolver clb.LoadBalancer) *TTransportHub {
	return &TTransportHub{resolver,
		map[string]*pool.TTransportPool{},
		&sync.Mutex{}}
}

// Take transport from hub by client options and discovery
func (h *TTransportHub) Take(
	clientOptions *TOptions,
	discovery string,
) (transport *pool.TPooledTransport, err error) {
	p := h.getPool(clientOptions, discovery)
	sock, err := p.Get(clientOptions.Ctx)
	transport = &pool.TPooledTransport{
		clientOptions.TransportFactory.GetTransport(sock),
		sock,
		p,
	}
	return
}

func (h *TTransportHub) getPool(
	clientOptions *TOptions,
	discovery string) (res *pool.TTransportPool) {
	d := clientOptions.GetDiscovery(discovery)
	sign := fmt.Sprintf("%s:%T", d, clientOptions.TransportFactory)

	h.mu.Lock()
	res, ok := h.transportPools[sign]
	if !ok {
		res = pool.NewTTransportPool(
			h.resolver, d,
			clientOptions.DialTimeout,
			clientOptions.PoolCapacity, clientOptions.PoolMax,
			clientOptions.PoolTTL)
		h.transportPools[sign] = res
		logx.Debugf("pool %p created in hub %p", res, h)
	}
	h.mu.Unlock()
	return
}
