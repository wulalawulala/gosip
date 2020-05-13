package gosip

import (
	"github.com/ghettovoice/gosip/log"
	"github.com/ghettovoice/gosip/transaction"
	"github.com/ghettovoice/gosip/transport"
)

type commonOptions struct {
	logger log.Logger
}

type ServerOption interface {
	applyServer(options *serverOptions)
}

type serverOptions struct {
	commonOptions

	config    *ServerConfig
	tpFactory transport.LayerFactory
	txFactory transaction.LayerFactory
}

func WithLogger(logger log.Logger) ServerOption {
	return withLogger{logger}
}

type withLogger struct {
	logger log.Logger
}

func (o withLogger) applyServer(options *serverOptions) {
	if o.logger != nil {
		options.logger = o.logger
	}
}

func WithTransactionLayerFactory(factory transaction.LayerFactory) ServerOption {
	return withTxFactory{factory}
}

type withTxFactory struct {
	factory transaction.LayerFactory
}

func (o withTxFactory) applyServer(options *serverOptions) {
	if o.factory != nil {
		options.txFactory = o.factory
	}
}

func WithTransportLayerFactory(factory transport.LayerFactory) ServerOption {
	return withTpFactory{factory}
}

type withTpFactory struct {
	factory transport.LayerFactory
}

func (o withTpFactory) applyServer(options *serverOptions) {
	if o.factory != nil {
		options.tpFactory = o.factory
	}
}
