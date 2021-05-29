package codec

import (
	"github.com/gogo/protobuf/proto"
	"google.golang.org/grpc"
)

type TypeProvider interface {
	RegisterTypes(TypeRegistry)
}

type TypeRegistry interface {
	// RegisterInterface associates protoName as the public name for the
	// interface passed in as iface. This is to be used primarily to create
	// a public facing registry of interface implementations for clients.
	// protoName should be a well-chosen public facing name that remains stable.
	// RegisterInterface takes an optional list of impls to be registered
	// as implementations of iface.
	//
	// Ex:
	//   registry.RegisterInterface("cosmos.base.v1beta1.Msg", (*sdk.Msg)(nil))
	RegisterInterface(protoName string, iface interface{}, impls ...proto.Message)

	// RegisterImplementations registers impls as concrete implementations of
	// the interface iface.
	//
	// Ex:
	//  registry.RegisterImplementations((*sdk.Msg)(nil), &MsgSend{}, &MsgMultiSend{})
	RegisterImplementations(iface interface{}, impls ...proto.Message)

	RegisterMsgServiceDesc(sd grpc.ServiceDesc, clientFactory interface{})

	RegisterQueryServiceDesc(sd grpc.ServiceDesc, clientFactory interface{})
}