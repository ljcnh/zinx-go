package znet

import "github.com/ljcnh/zinx-go/ziface"

// 基类

type BaseRouter struct{}

func (r *BaseRouter) PreHandle(request ziface.IRequest) {}

func (r *BaseRouter) Handle(request ziface.IRequest) {}

func (r *BaseRouter) PostHandle(request ziface.IRequest) {}
