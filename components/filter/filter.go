package filter

import (
	"context"
	"fmt"
)

// HandlerFunc 请求处理接口
type HandlerFunc func(ctx context.Context, req interface{}) (rsp interface{}, err error)

// Filter 过滤器
type Filter func(ctx context.Context, req interface{}, next HandlerFunc) (rsp interface{}, err error)

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
