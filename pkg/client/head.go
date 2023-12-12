package client

import "context"

// Head 请求
type Head struct {
	PathOrApiName string // 请求路径
	Param         *Param // 其他请求参数
	Body          interface{}
	Rsp           interface{}
}

/***************************/

type headContext struct{}

func WithHead(ctx context.Context, head *Head) context.Context {
	return context.WithValue(ctx, headContext{}, head)
}

func HeadFrom(ctx context.Context) *Head {
	t := ctx.Value(headContext{})
	if t != nil {
		return t.(*Head)
	}
	return nil
}

func EnsureHead(ctx context.Context) (context.Context, *Head) {
	head := HeadFrom(ctx)
	if head != nil {
		return ctx, head
	}
	head = &Head{}
	ctx = WithHead(ctx, head)
	return ctx, head
}
