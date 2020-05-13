package transport

import (
	"net"

	"github.com/ghettovoice/gosip/log"
	"github.com/ghettovoice/gosip/sip"
)

type commonOptions struct {
	logger    log.Logger
	msgMapper sip.MessageMapper
}

type LayerOption interface {
	applyLayer(options *layerOptions)
}

type layerOptions struct {
	commonOptions

	dnsResolver *net.Resolver
}

type ProtocolOption interface {
	applyProtocol(options *protocolOptions)
}

type protocolOptions commonOptions

type ConnectionPoolOption interface {
	applyConnectionPool(options *connectionPoolOptions)
}

type connectionPoolOptions commonOptions

type ConnectionHandlerOption interface {
	applyConnectionHandler(options *connectionHandlerOptions)
}

type connectionHandlerOptions commonOptions

type ConnectionOption interface {
	applyConnection(options *connectionOptions)
}

type connectionOptions commonOptions

type ListenerPoolOption interface {
	applyListenerPool(options *listenerPoolOptions)
}

type listenerPoolOptions commonOptions

type ListenerHandlerOption interface {
	applyListenerHandler(options *listenerHandlerOptions)
}

type listenerHandlerOptions commonOptions

func WithLogger(logger log.Logger) interface {
	LayerOption
	ProtocolOption
	ConnectionPoolOption
	ConnectionHandlerOption
	ConnectionOption
	ListenerPoolOption
	ListenerHandlerOption
} {
	return withLogger{logger}
}

type withLogger struct {
	logger log.Logger
}

func (o withLogger) applyLayer(options *layerOptions) {
	if o.logger != nil {
		options.logger = o.logger
	}
}

func (o withLogger) applyProtocol(options *protocolOptions) {
	if o.logger != nil {
		options.logger = o.logger
	}
}

func (o withLogger) applyConnectionPool(options *connectionPoolOptions) {
	if o.logger != nil {
		options.logger = o.logger
	}
}

func (o withLogger) applyConnectionHandler(options *connectionHandlerOptions) {
	if o.logger != nil {
		options.logger = o.logger
	}
}

func (o withLogger) applyConnection(options *connectionOptions) {
	if o.logger != nil {
		options.logger = o.logger
	}
}

func (o withLogger) applyListenerPool(options *listenerPoolOptions) {
	if o.logger != nil {
		options.logger = o.logger
	}
}

func (o withLogger) applyListenerHandler(options *listenerHandlerOptions) {
	if o.logger != nil {
		options.logger = o.logger
	}
}

func WithDnsResolver(resolver *net.Resolver) LayerOption {
	return withDnsResolver{resolver}
}

type withDnsResolver struct {
	resolver *net.Resolver
}

func (o withDnsResolver) applyLayer(options *layerOptions) {
	if o.resolver != nil {
		options.dnsResolver = o.resolver
	}
}

func WithMessageMapper(mapper sip.MessageMapper) interface {
	LayerOption
	ProtocolOption
	ConnectionPoolOption
	ConnectionHandlerOption
} {
	return withMsgMapper{mapper}
}

type withMsgMapper struct {
	mapper sip.MessageMapper
}

func (o withMsgMapper) applyLayer(options *layerOptions) {
	if o.mapper != nil {
		options.msgMapper = o.mapper
	}
}

func (o withMsgMapper) applyProtocol(options *protocolOptions) {
	if o.mapper != nil {
		options.msgMapper = o.mapper
	}
}

func (o withMsgMapper) applyConnectionPool(options *connectionPoolOptions) {
	if o.mapper != nil {
		options.msgMapper = o.mapper
	}
}

func (o withMsgMapper) applyConnectionHandler(options *connectionHandlerOptions) {
	if o.mapper != nil {
		options.msgMapper = o.mapper
	}
}
