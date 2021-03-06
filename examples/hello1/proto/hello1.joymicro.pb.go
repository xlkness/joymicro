// Code generated by protoc-gen-joymicro. DO NOT EDIT.
// source: hello1.proto

package proto

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

var serviceName = "hello"

func GetServiceName() string {
	return serviceName
}

type HelloServiceInterface interface {
	Hello(context.Context, *Request) (*Response, error)
	HelloPeer(context.Context, string, *Request) (*Response, error)
	HelloAll(context.Context, *Request) (*Response, error)
}

func NewHelloService(etcdAddrs []string, timeout time.Duration, isPermanent bool) HelloServiceInterface {
	c := client.New(serviceName, etcdAddrs, timeout, isPermanent)
	// 设置点对点选择器
	c.SetSelector(&client.PeerSelector{})
	return &helloService{
		c: c,
	}
}

type helloService struct {
	c *client.Service
}

func (c *helloService) Hello(ctx context.Context, in *Request) (*Response, error) {
	out := new(Response)
	err := c.c.Call(ctx, "Hello", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloService) HelloPeer(ctx context.Context, peerKey string, in *Request) (*Response, error) {
	out := new(Response)
	ctx = context.WithValue(ctx, "select_key", peerKey)
	err := c.c.Call(ctx, "Hello", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloService) HelloAll(ctx context.Context, in *Request) (*Response, error) {
	out := new(Response)
	err := c.c.CallAll(ctx, "Hello", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type HelloHandlerInterface interface {
	Hello(context.Context, *Request, *Response) error
}

func RegisterHelloHandler(s *server.ServicesManager, hdlr HelloHandlerInterface) error {
	return s.RegisterOneService(serviceName, hdlr)
}

//===============================================Json Handler for Test===============================================

func NewHelloJsonTestService(etcdAddrs []string, timeout time.Duration, isPermanent bool) (reflect.Type, reflect.Value) {
	c := NewHelloService(etcdAddrs, timeout, isPermanent)
	c1 := &HelloJsonTestService{c: c}
	return reflect.TypeOf(c1), reflect.ValueOf(c1)
}

type HelloJsonTestService struct {
	c HelloServiceInterface
}

func (c *HelloJsonTestService) Hello(ctx context.Context, in string) (*Response, error) {
	newIn := &Request{}
	err := json.Unmarshal([]byte(in), newIn)
	if err != nil {
		return nil, err
	}

	return c.c.Hello(ctx, newIn)
}

func (c *HelloJsonTestService) HelloPeer(ctx context.Context, peerKey string, in string) (*Response, error) {
	newIn := &Request{}
	err := json.Unmarshal([]byte(in), newIn)
	if err != nil {
		return nil, err
	}

	return c.c.HelloPeer(ctx, peerKey, newIn)
}

func (c *HelloJsonTestService) HelloAll(ctx context.Context, in string) (*Response, error) {
	newIn := &Request{}
	err := json.Unmarshal([]byte(in), newIn)
	if err != nil {
		return nil, err
	}

	return c.c.HelloAll(ctx, newIn)
}
