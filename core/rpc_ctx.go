package core

import (
	"context"
	"fmt"
)

type ctxKey string

var (
	rpcCtxKey = ctxKey("rpc_ctx")
)

func WithCtx(ctx context.Context, rpcCtx RpcCtx) context.Context {
	return context.WithValue(ctx, rpcCtxKey, rpcCtx)
}

func RpcContext(ctx context.Context) RpcCtx {
	val := ctx.Value(rpcCtxKey)
	if val == nil {
		val = NewDefaultRpcCtx()
	}

	return val.(RpcCtx)
}

// RpcCtx rpc请求上下文
type RpcCtx interface {
	GetName() string      // 获取RPCName
	WithName(name string) // 设置RPCName

	Ext() map[string]any         // 获取扩展信息
	WithExt(key string, val any) // 设置扩展字段
}

type defaultRpcCtx struct {
	name string
	ext  map[string]any
}

func NewDefaultRpcCtx() *defaultRpcCtx {
	return &defaultRpcCtx{
		ext: map[string]any{},
	}
}

func (ctx *defaultRpcCtx) GetName() string {
	return ctx.name
}

func (ctx *defaultRpcCtx) WithName(name string) {
	ctx.name = name
}

func (ctx *defaultRpcCtx) Ext() map[string]any {
	return ctx.ext
}

func (ctx *defaultRpcCtx) WithExt(key string, val any) {
	ctx.ext[key] = val
}

func ExtVal(ctx RpcCtx, key string) (any, error) {
	ext := ctx.Ext()
	val, ok := ext[key]
	if !ok {
		return 0, fmt.Errorf("%s not exist", key)
	}

	return val, nil
}

func ExtIntVal(ctx RpcCtx, key string) (int, error) {
	if val, err := ExtVal(ctx, key); err != nil {
		return 0, err
	} else {
		v, _ := val.(int)
		return v, nil
	}
}

func ExtUint32Val(ctx RpcCtx, key string) uint32 {
	if val, err := ExtVal(ctx, key); err != nil {
		return 0
	} else {
		v, _ := val.(uint32)
		return v
	}
}

func ExtInt32Val(ctx RpcCtx, key string) int32 {
	if val, err := ExtVal(ctx, key); err != nil {
		return 0
	} else {
		v, _ := val.(int32)
		return v
	}
}

func ExtStringVal(ctx RpcCtx, key string) string {
	if val, err := ExtVal(ctx, key); err != nil {
		return ""
	} else {
		v, _ := val.(string)
		return v
	}
}
