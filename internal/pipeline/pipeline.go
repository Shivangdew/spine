package pipeline

import (
	"github.com/NARUBROWN/spine"
	"github.com/NARUBROWN/spine/internal/invoker"
	"github.com/NARUBROWN/spine/internal/router"
)

type Pipeline struct {
	invoker *invoker.Invoker
}

func NewPipeline(invoker *invoker.Invoker) *Pipeline {
	return &Pipeline{invoker: invoker}
}

// Execute는 HandlerMeta를 실행합니다.
func (p *Pipeline) Execute(ctx spine.Context, meta router.HandlerMeta) error {
	return p.invoker.Invoke(
		meta.ControllerType,
		meta.MethodName,
		ctx,
	)
}
