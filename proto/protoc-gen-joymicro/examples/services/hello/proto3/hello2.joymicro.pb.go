// Code generated by protoc-gen-joymicro. DO NOT EDIT.
// source: api.desc/hello2.proto

package proto3

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	"encoding/json"
	client "joynova.com/joynova/joymicro/client"
	server "joynova.com/joynova/joymicro/service"
	"reflect"
	"time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context

var serviceName = "hello2"

func GetServiceName() string {
	return serviceName
}

type Hello2ServiceInterface interface {
	Hello2(context.Context, *Request2) (*Response2, error)
	Hello2Peer(context.Context, string, *Request2) (*Response2, error)
	Hello2All(context.Context, *Request2) (*Response2, error)
}

func NewHello2Service(etcdAddrs []string, timeout time.Duration, isPermanent bool) Hello2ServiceInterface {
	c := client.New(serviceName, etcdAddrs, timeout, isPermanent)
	// 设置点对点选择器
	c.SetSelector(&client.PeerSelector{})
	return &hello2Service{
		c: c,
	}
}

type hello2Service struct {
	c *client.Service
}

func (c *hello2Service) Hello2(ctx context.Context, in *Request2) (*Response2, error) {
	out := new(Response2)
	err := c.c.Call(ctx, "Hello2", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hello2Service) Hello2Peer(ctx context.Context, peerKey string, in *Request2) (*Response2, error) {
	out := new(Response2)
	ctx = context.WithValue(ctx, "select_key", peerKey)
	err := c.c.Call(ctx, "Hello2", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hello2Service) Hello2All(ctx context.Context, in *Request2) (*Response2, error) {
	out := new(Response2)
	err := c.c.CallAll(ctx, "Hello2", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type Hello2HandlerInterface interface {
	Hello2(context.Context, *Request2, *Response2) error
}

func RegisterHello2Handler(s *server.ServicesManager, hdlr Hello2HandlerInterface) error {
	return s.RegisterOneService(serviceName, hdlr)
}

//===============================================Json Handler for Test===============================================

func NewHello2JsonTestService(etcdAddrs []string, timeout time.Duration, isPermanent bool) (reflect.Type, reflect.Value) {
	c := NewHello2Service(etcdAddrs, timeout, isPermanent)
	c1 := &Hello2JsonTestService{c: c}
	return reflect.TypeOf(c1), reflect.ValueOf(c1)
}

type Hello2JsonTestService struct {
	c Hello2ServiceInterface
}

func (c *Hello2JsonTestService) Hello2(ctx context.Context, in string) (*Response2, error) {
	newIn := &Request2{}
	err := json.Unmarshal([]byte(in), newIn)
	if err != nil {
		return nil, err
	}

	return c.c.Hello2(ctx, newIn)
}

func (c *Hello2JsonTestService) Hello2Peer(ctx context.Context, peerKey string, in string) (*Response2, error) {
	newIn := &Request2{}
	err := json.Unmarshal([]byte(in), newIn)
	if err != nil {
		return nil, err
	}

	return c.c.Hello2Peer(ctx, peerKey, newIn)
}

func (c *Hello2JsonTestService) Hello2All(ctx context.Context, in string) (*Response2, error) {
	newIn := &Request2{}
	err := json.Unmarshal([]byte(in), newIn)
	if err != nil {
		return nil, err
	}

	return c.c.Hello2All(ctx, newIn)
}