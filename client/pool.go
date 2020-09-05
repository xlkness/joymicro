package client

//
//import (
//	"time"
//)
//
//type Pool struct {
//	clients    []*RpcClientManager
//	hashGetFun func(interface{}) int
//}
//
//// NewPool 创建rpc客户端池，支持负载均衡函数，大部分交互不需要用池，
//// 主要是持续大量的通信包会用来做高可用和容灾，一条tcp链接出错不至于
//// 影响全部，例如网关gate和game的交互启动10条rpc通信链
//func NewPool(basePath string, nodeName string, etcdServerAddrs []string, poolNum int,
//	hashGetFun func(interface{}) int, callTimeout time.Duration, isPermanentSocketLink bool) *Pool {
//	pool := &Pool{
//		hashGetFun: hashGetFun,
//	}
//	for i := 0; i < poolNum; i++ {
//		c := New(basePath, nodeName, etcdServerAddrs, callTimeout, isPermanentSocketLink)
//		pool.clients = append(pool.clients, c)
//	}
//	return pool
//}
//
//// Call 单点call，hashParam为hash函数的参数，选择池的客户端，例如playerid
//func (p *Pool) Call(hashParam interface{}, no int, method string,
//	args interface{}, reply interface{}) error {
//	idx := p.hashGetFun(hashParam)
//	c := p.clients[idx]
//	return c.Call(no, method, args, reply)
//}
//
//// LBCall 负载均衡call，hashParam为hash函数的参数，选择池的客户端，例如playerid
//func (p *Pool) LBCall(hashParam interface{}, method string, args interface{}, reply interface{}) error {
//	idx := p.hashGetFun(hashParam)
//	c := p.clients[idx]
//	return c.LBCall(method, args, reply)
//}
//
//// LBCall 负载均衡call，hashParam为hash函数的参数，选择池的客户端，例如playerid
//func (p *Pool) CallAll(hashParam interface{}, method string, args interface{}, reply interface{}) error {
//	idx := p.hashGetFun(hashParam)
//	c := p.clients[idx]
//	return c.CallAll(method, args, reply)
//}
//
//// LBCall 负载均衡Cast，hashParam为pash函数的参数，选择池的客户端，例如playerid
//func (p *Pool) Cast(hashParam interface{}, no int, method string, args interface{}) error {
//	idx := p.hashGetFun(hashParam)
//	c := p.clients[idx]
//	return c.Cast(no, method, args)
//}
//
//// CastAll cast所有
//func (p *Pool) CastAll(hashParam interface{}, method string, args interface{}) error {
//	idx := p.hashGetFun(hashParam)
//	c := p.clients[idx]
//	return c.CastAll(method, args)
//}
//
//// LBCast 负载均衡cast
//func (p *Pool) LBCast(hashParam interface{}, method string, args interface{}) error {
//	idx := p.hashGetFun(hashParam)
//	c := p.clients[idx]
//	return c.LBCast(method, args)
//}
