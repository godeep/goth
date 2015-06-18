package server

import "github.com/apache/thrift/lib/go/thrift"

func GuessProtoFactory(what string) (res thrift.TProtocolFactory) {
	switch what {
	case "compact":
		res = thrift.NewTCompactProtocolFactory()
	case "simplejson":
		res = thrift.NewTSimpleJSONProtocolFactory()
	case "json":
		res = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		res = thrift.NewTBinaryProtocolFactoryDefault()
	default:
		res = thrift.NewTBinaryProtocolFactoryDefault()
	}
	return
}

func GuessTransportFactory(framed bool, buffer int) (res thrift.TTransportFactory) {
	if buffer != 0 {
		res = thrift.NewTBufferedTransportFactory(s.Buffer)
	} else {
		res = thrift.NewTTransportFactory()
	}

	if framed {
		res = thrift.NewTFramedTransportFactory(res)
	}
	return
}
