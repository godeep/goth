package pool

import (
	"github.com/apache/thrift/lib/go/thrift"
)

// Pooled thrift transport
type TPooledTransport struct {
	thrift.TTransport
	Sock *thrift.TSocket
	TransportPool *TTransportPool
}

func (t *TPooledTransport) Release() {
	t.TransportPool.Put(t.Sock)
}
