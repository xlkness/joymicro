// Code generated by protoc-gen-joymicro. DO NOT EDIT.
// source: api.desc/hello3.proto

package proto4

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

var serviceName = "hello3"

func GetServiceName() string {
	return serviceName
}

type Hello3ServiceInterface interface {
	Hello3(context.Context, *Request3) (*Response3, error)
	Hello3All(context.Context, *Request3) (*Response3, error)
	HowAreYou(context.Context, *Request3) (*Response3, error)
	HowAreYouAll(context.Context, *Request3) (*Response3, error)
}

func NewHello3Service(etcdAddrs []string, timeout time.Duration, isPermanent bool) Hello3ServiceInterface {
	c := client.New(serviceName, etcdAddrs, timeout, isPermanent)
	return &hello3Service{
		c: c,
	}
}

type hello3Service struct {
	c *client.Service
}

func (c *hello3Service) Hello3(ctx context.Context, in *Request3) (*Response3, error) {
	out := new(Response3)
	err := c.c.Call(ctx, "Hello3", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hello3Service) Hello3All(ctx context.Context, in *Request3) (*Response3, error) {
	out := new(Response3)
	err := c.c.CallAll(ctx, "Hello3", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hello3Service) HowAreYou(ctx context.Context, in *Request3) (*Response3, error) {
	out := new(Response3)
	err := c.c.Call(ctx, "HowAreYou", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hello3Service) HowAreYouAll(ctx context.Context, in *Request3) (*Response3, error) {
	out := new(Response3)
	err := c.c.CallAll(ctx, "HowAreYou", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type Hello3HandlerInterface interface {
	Hello3(context.Context, *Request3, *Response3) error
	HowAreYou(context.Context, *Request3, *Response3) error
}

func RegisterHello3Handler(s *server.ServicesManager, hdlr Hello3HandlerInterface) error {
	return s.RegisterOneService(serviceName, hdlr)
}

//===============================================Json Handler for Test===============================================

func NewHello3JsonTestService(etcdAddrs []string, timeout time.Duration, isPermanent bool) (reflect.Type, reflect.Value) {
	c := NewHello3Service(etcdAddrs, timeout, isPermanent)
	c1 := &Hello3JsonTestService{c: c}
	return reflect.TypeOf(c1), reflect.ValueOf(c1)
}

type Hello3JsonTestService struct {
	c Hello3ServiceInterface
}

func (c *Hello3JsonTestService) Hello3(ctx context.Context, in string) (*Response3, error) {
	newIn := &Request3{}
	err := json.Unmarshal([]byte(in), newIn)
	if err != nil {
		return nil, err
	}

	return c.c.Hello3(ctx, newIn)
}

func (c *Hello3JsonTestService) Hello3All(ctx context.Context, in string) (*Response3, error) {
	newIn := &Request3{}
	err := json.Unmarshal([]byte(in), newIn)
	if err != nil {
		return nil, err
	}

	return c.c.Hello3All(ctx, newIn)
}

func (c *Hello3JsonTestService) HowAreYou(ctx context.Context, in string) (*Response3, error) {
	newIn := &Request3{}
	err := json.Unmarshal([]byte(in), newIn)
	if err != nil {
		return nil, err
	}

	return c.c.HowAreYou(ctx, newIn)
}

func (c *Hello3JsonTestService) HowAreYouAll(ctx context.Context, in string) (*Response3, error) {
	newIn := &Request3{}
	err := json.Unmarshal([]byte(in), newIn)
	if err != nil {
		return nil, err
	}

	return c.c.HowAreYouAll(ctx, newIn)
}
