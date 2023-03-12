package filter

import (
	"context"
	"fmt"
	"testing"
)

var echoHandler = func(ctx context.Context, req interface{}) (rsp interface{}, err error) {
	fmt.Printf("echo handler...\n")
	return nil, nil
}

var noopFilter = func(ctx context.Context, req interface{}, chain HandlerFunc) (rsp interface{}, err error) {
	fmt.Printf("noop filter before\n")
	rsp, err = chain(ctx, req)
	fmt.Printf("noop filter after\n")
	return
}

func Test_filter(t *testing.T) {
	ctx := context.TODO()
	var req interface{}

	filters := []Filter{LoggerFilter, noopFilter}
	var next = echoHandler
	for i := len(filters) - 1; i >= 0; i-- {
		curF, curFilter := next, filters[i]
		next = func(ctx context.Context, req interface{}) (rsp interface{}, err error) {
			return curFilter(ctx, req, curF)
		}
	}

	rsp, err := next(ctx, req)
	t.Logf("rsp:%+v, err:%+v", rsp, err)
}
