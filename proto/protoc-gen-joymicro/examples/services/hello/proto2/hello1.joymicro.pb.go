// Code generated by protoc-gen-joymicro. DO NOT EDIT.
// source: api.desc/hello1.proto

package proto2

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

var serviceName = "hello1"

func GetServiceName() string {
	return serviceName
}

type Hello1ServiceInterface interface {
	Hello1(context.Context, string, *Request1) (*Response1, error)
	Hello1All(context.Context, *Request1) (*Response1, error)
}

func NewHello1Service(etcdAddrs []string, timeout time.Duration, isPermanent bool) Hello1ServiceInterface {
	c := client.New(serviceName, etcdAddrs, timeout, isPermanent)
	// 设置一致性hash
	c.SetSelector(client.NewConsistentHashSelector())
	return &hello1Service{
		c: c,
	}
}

type hello1Service struct {
	c *client.Service
}

func (c *hello1Service) Hello1(ctx context.Context, key string, in *Request1) (*Response1, error) {
	out := new(Response1)
	ctx = context.WithValue(ctx, "select_key", key)
	err := c.c.Call(ctx, "Hello1", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hello1Service) Hello1All(ctx context.Context, in *Request1) (*Response1, error) {
	out := new(Response1)
	err := c.c.CallAll(ctx, "Hello1", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type Hello1HandlerInterface interface {
	Hello1(context.Context, *Request1, *Response1) error
}

func RegisterHello1Handler(s *server.ServicesManager, hdlr Hello1HandlerInterface) error {
	return s.RegisterOneService(serviceName, hdlr)
}

//===============================================Json Handler for Test===============================================

func NewHello1JsonTestService(etcdAddrs []string, timeout time.Duration, isPermanent bool) (reflect.Type, reflect.Value) {
	c := NewHello1Service(etcdAddrs, timeout, isPermanent)
	c1 := &Hello1JsonTestService{c: c}
	return reflect.TypeOf(c1), reflect.ValueOf(c1)
}

type Hello1JsonTestService struct {
	c Hello1ServiceInterface
}

func (c *Hello1JsonTestService) Hello1(ctx context.Context, key string, in string) (*Response1, error) {
	newIn := &Request1{}
	err := json.Unmarshal([]byte(in), newIn)
	if err != nil {
		return nil, err
	}

	return c.c.Hello1(ctx, key, newIn)
}

func (c *Hello1JsonTestService) Hello1All(ctx context.Context, in string) (*Response1, error) {
	newIn := &Request1{}
	err := json.Unmarshal([]byte(in), newIn)
	if err != nil {
		return nil, err
	}

	return c.c.Hello1All(ctx, newIn)
}
