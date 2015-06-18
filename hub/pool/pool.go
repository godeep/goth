package pool

import (
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/benschw/dns-clb-go/clb"
	"github.com/tamtam-im/goth/hub/pool/resource_pool"
	"golang.org/x/net/context"
	"gopkg.in/tamtam-im/logx.v4"
	"io"
	"time"
)

type closerWrapper struct {
	io.Closer
}

func (c *closerWrapper) Close() {
	c.Closer.Close()
}

// thrift.TTransport pool
type TTransportPool struct {
	// Provided resolver
	resolver clb.LoadBalancer

	// Pool discovery
	discovery string

	// Dial timeout
	dialTimeout time.Duration

	// TTransport overlay
	transportFactory thrift.TTransportFactory

	*resource_pool.ResourcePool
}

func NewTTransportPool(
	resolver clb.LoadBalancer,
	discovery string,
	timeout time.Duration,
	cap, max int,
	ttl time.Duration,
) (res *TTransportPool) {
	res = &TTransportPool{
		resolver:    resolver,
		discovery:   discovery,
		dialTimeout: timeout,
	}
	res.ResourcePool = resource_pool.NewResourcePool(res.factory, cap, max, ttl)
	return
}

func (p *TTransportPool) Get(ctx context.Context) (
	sock *thrift.TSocket, err error) {

	w, err := p.ResourcePool.Get(ctx)
	if err != nil {
		return
	}
	sock = w.(*closerWrapper).Closer.(*thrift.TSocket)
	logx.Debugf("socket %p taken from pool %p (%s)",
		sock, p, p.discovery)
	return
}

func (p *TTransportPool) Put(sock *thrift.TSocket) {
	if sock.IsOpen() {
		p.ResourcePool.Put(&closerWrapper{sock})
		logx.Debugf("socket %p returned to pool %p (%s)",
			sock, p, p.discovery)
	} else {
		p.ResourcePool.Put(nil)
		logx.Debugf("transport %p ejected from pool %p (%s)",
			sock, p, p.discovery)
	}
}

func (p *TTransportPool) factory() (res resource_pool.Resource, err error) {
	resolved, err := p.resolver.GetAddress(p.discovery)
	if err != nil {
		return
	}
	sock, err := thrift.NewTSocketTimeout(resolved.String(), p.dialTimeout)
	if err != nil {
		return
	}
	err = sock.Open()
	if err != nil {
		return
	}
	res = &closerWrapper{sock}
	logx.Debugf("socket %p created in pool %p (%s)",
		sock, p, p.discovery)
	return
}
