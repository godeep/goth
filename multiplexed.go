package client

import (
	"github.com/apache/thrift/lib/go/thrift"
)

// Missing multiplexed protocol factory
type TMultiplexedProtocolFactory struct {
	name string
	thrift.TProtocolFactory
}

func NewTMultiplexedProtocolFactory(
name string,
factory thrift.TProtocolFactory) *TMultiplexedProtocolFactory {
	return &TMultiplexedProtocolFactory{
		name:            name,
		TProtocolFactory: factory,
	}
}

func (p *TMultiplexedProtocolFactory) GetProtocol(
t thrift.TTransport) thrift.TProtocol {
	return thrift.NewTMultiplexedProtocol(p.TProtocolFactory.GetProtocol(t),
		p.name)
}
