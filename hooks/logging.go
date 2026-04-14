package hooks

import (
	"context"
	"go-ent-demo/entcore"
	"log"
	"time"

	"entgo.io/ent"
)

func LoggingHook() func(next entcore.Mutator) entcore.Mutator {
	return func(next entcore.Mutator) entcore.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			start := time.Now()
			defer func() {
				log.Printf("Op=%s\tType=%s\tTime=%s\tConcreteType=%T\n", m.Op(), m.Type(), time.Since(start), m)
			}()

			return next.Mutate(ctx, m)
		})
	}
}
