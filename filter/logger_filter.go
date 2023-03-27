package filter

import (
	"context"
	"fmt"
	"time"

	"github.com/mingo-chen/wheel-minirpc/core"
)

// filter1: func(ctx, req, next) (rsp, error)
// filter2: func(ctx, req, next) (rsp, error)
// filter3: func(ctx, req, next) (rsp, error)
// rsp, err := handler(ctx, req)
// AccessFilter 记录rpc访问日志的过滤器
func AccessFilter(ctx context.Context, req interface{}, chain core.HandlerFunc) (rsp interface{}, err error) {
	t0 := time.Now()
	fmt.Printf("server hand start, req:%+v\n", req)
	rsp, err = chain(ctx, req)
	cost := time.Since(t0)
	// 到时替换为正确的日志组件
	fmt.Printf("server hand done, cost[%d ms], req:%+v, rsp:%+v, err:%+v\n", cost.Milliseconds(), req, rsp, err)
	return rsp, err
}
