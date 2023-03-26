package core

import "context"

// HandlerFunc 请求处理的逻辑
type HandlerFunc func(ctx context.Context, req interface{}) (rsp interface{}, err error)
