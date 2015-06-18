package hub
import (
	"time"
	"golang.org/x/net/context"
	"github.com/apache/thrift/lib/go/thrift"
)

var defaultTOptions *TOptions

// Thrift transport hub options
type TOptions struct {
	// Default context
	Ctx context.Context

	// DNS Discovery base: service.<base>. Default is ".thrift"
	DiscoveryBase string

	// DNS discovery environment. Default is ""
	DiscoveryEnv string

	// Dial timeout
	DialTimeout time.Duration

	// Pool capacity. 16 by default
	PoolCapacity int

	// Pool cap. 256 by default
	PoolMax int

	// Pool TTL. Hour by default
	PoolTTL time.Duration

	// Transport factory overlay.
	// thrift.NewTTransportFactory() used by default
	TransportFactory thrift.TTransportFactory

	// Thrift Protocol Factory factory :-)
	// By default returns thrift.NewTBinaryProtocolFactoryDefault()
	ProtoFactory func() thrift.TProtocolFactory

	// Hub factory to use different hub
	HubFactory func() *TTransportHub
}

// Get clone of default Thrift client options or just return first
// given. This logic is little funny. But very convenient for per-client
// options.
func GetTOptions(options ...*TOptions) (res *TOptions) {
	if len(options) > 0 {
		return options[0]
	}
	if defaultTOptions == nil {
		defaultTOptions = &TOptions{
			DiscoveryBase: ".thrift",
			DiscoveryEnv: "",
			TransportFactory: thrift.NewTFramedTransportFactory(
				thrift.NewTTransportFactory()),
			DialTimeout: time.Minute * 3,
			PoolCapacity: 32,
			PoolMax: 256,
			PoolTTL: time.Hour,
			ProtoFactory: func() thrift.TProtocolFactory {
				return thrift.NewTBinaryProtocolFactoryDefault()
			},
			Ctx: context.Background(),
			HubFactory: DefaultTTransportHub,
		}
	}
	clone := *defaultTOptions
	res = &clone
	return
}

// Replace default Thrift client options
func SetTOptions(options *TOptions) {
	defaultTOptions = options
}

// Get discovery dns
func (o *TOptions) GetDiscovery(discovery string) string {
	return o.DiscoveryEnv + "." + discovery + "." + o.DiscoveryBase
}

func (o *TOptions) SetCtx(ctx context.Context) *TOptions {
	o.Ctx = ctx
	return o
}

func (o *TOptions) SetDiscovery(env, base string) *TOptions {
	o.DiscoveryEnv = env
	o.DiscoveryBase = base
	return o
}

// Set discovery base. "thrift.tam.tam" for example.
func (o *TOptions) SetDiscoveryBase(discoveryBase string) *TOptions {
	o.DiscoveryBase = discoveryBase
	return o
}

// Set discovery environment. "stage"
func (o *TOptions) SetDiscoveryEnv(discoveryEnv string) *TOptions {
	o.DiscoveryEnv = discoveryEnv
	return o
}

func (o *TOptions) SetDialTimeout(timeout time.Duration) *TOptions {
	o.DialTimeout = timeout
	return o
}

func (o *TOptions) SetPoolCapacity(capacity int) *TOptions {
	o.PoolCapacity = capacity
	return o
}

func (o *TOptions) SetPoolMax(max int) *TOptions {
	o.PoolMax = max
	return o
}

func (o *TOptions) SetPoolTTL(ttl time.Duration) *TOptions {
	o.PoolTTL = ttl
	return o
}

func (o *TOptions) SetTransportFactory(factory thrift.TTransportFactory) *TOptions {
	o.TransportFactory = factory
	return o
}

func (o *TOptions) SetProtoFactory(factory func() thrift.TProtocolFactory) *TOptions {
	o.ProtoFactory = factory
	return o
}

func (o *TOptions) SetHubFactory(factory func() *TTransportHub) *TOptions {
	o.HubFactory = factory
	return o
}
