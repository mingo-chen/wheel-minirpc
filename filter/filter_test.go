package filter

import (
	"context"
	"fmt"
	"testing"

	"github.com/mingo-chen/wheel-minirpc/core"
)

var echoHandler = func(ctx context.Context, req interface{}) (rsp interface{}, err error) {
	fmt.Printf("echo handler...\n")
	return nil, nil
}

var noopFilter = func(ctx context.Context, req interface{}, chain core.HandlerFunc) (rsp interface{}, err error) {
	fmt.Printf("noop filter before\n")
	rsp, err = chain(ctx, req)
	fmt.Printf("noop filter after\n")
	return
}

func Test_filter(t *testing.T) {
	ctx := context.TODO()
	var req interface{}

	filters := []Filter{AccessFilter, noopFilter}
	rsp, err := FilterChain(filters, echoHandler)(ctx, req)
	t.Logf("rsp:%+v, err:%+v", rsp, err)
}
