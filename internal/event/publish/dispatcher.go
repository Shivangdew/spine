package publish

import (
	"context"
	"log"

	"github.com/NARUBROWN/spine/pkg/event/publish"
)

type EventDispatcher interface {
	Dispatch(ctx context.Context, events []publish.DomainEvent)
}

type DefaultEventDispatcher struct {
	Publishers []EventPublisher
}

func (d *DefaultEventDispatcher) Dispatch(ctx context.Context, events []publish.DomainEvent) {
	for _, e := range events {
		for _, p := range d.Publishers {
			if err := p.Publish(ctx, e); err != nil {
				log.Printf("[EventDispatcher] 이벤트 발행 실패 (%s): %v", e.Name(), err)
			}
		}
	}
}
