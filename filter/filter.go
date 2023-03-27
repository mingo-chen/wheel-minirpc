package filter

import (
	"context"
	"fmt"

	"github.com/mingo-chen/wheel-minirpc/core"
)

// Filter 过滤器
type Filter func(ctx context.Context, req interface{}, next core.HandlerFunc) (rsp interface{}, err error)

// FilterChain Filter+HandlerFunc 组成的处理链
var FilterChain = func(filters []Filter, baseHandler core.HandlerFunc) core.HandlerFunc {
	return func(ctx context.Context, req interface{}) (rsp interface{}, err error) {
		handler := baseHandler // 从最底的业务实现处理器开始，层层包装
		for i := len(filters) - 1; i >= 0; i-- {
			curHandler, curFilter := handler, filters[i]
			handler = func(ctx context.Context, req interface{}) (rsp interface{}, err error) {
				return curFilter(ctx, req, curHandler)
			}
		}

		rsp, err = handler(ctx, req) // 最终的处理器
		return
	}
}

var filterHub = map[string]Filter{}

// RegistFilter 注册过滤器
func RegistFilter(name string, filter Filter) {
	if _, ok := filterHub[name]; ok {
		panic(fmt.Sprintf("filter[%s] has regist", name))
	}

	filterHub[name] = filter
}

// Get 按名字取出预先注册的过滤器
func Get(name string) Filter {
	return filterHub[name]
}
