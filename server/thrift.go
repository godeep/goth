package server

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"gopkg.in/tamtam-im/logx.v4"
)

// Configurable thrift server
type TServer struct {

	// Bind `address:port` (:9090)
	Bind string

	// Protocol factory
	ProtoFactory thrift.TProtocolFactory

	// Transport factory
	TransportFactory thrift.TTransportFactory

	// Processor
	Processor thrift.TProcessor

	server    thrift.TServer
}

// New thrift server with defaults and common
func NewThriftServer(processor thrift.TProcessor) (res *TServer) {
	res = &TServer{
		ProtoFactory: thrift.NewTBinaryProtocolFactoryDefault(),
		TransportFactory: thrift.NewTTransportFactory(),
		Processor: processor,
	}
	return
}

func (s TServer) String() string {
	return fmt.Sprintf("%s (%T %T)",
		s.Bind, s.ProtoFactory, s.TransportFactory)
}

// Serve requests
func (s *TServer) Serve() error {
	var transport thrift.TServerTransport
	transport, err := thrift.NewTServerSocket(s.Bind)
	if err != nil {
		return err
	}
	s.server = thrift.NewTSimpleServer4(s.GetProcessor(), transport,
		s.TransportFactory, s.ProtoFactory)

	logx.Infof("serving %s", s)
	return s.server.Serve()
}

func (s *TServer) GetProcessor() thrift.TProcessor {
	return s.Processor
}

func (s *TServer) Stop() error {
	logx.Info("bye")
	return s.server.Stop()
}
