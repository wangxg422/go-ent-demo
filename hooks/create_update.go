package hooks

import (
	"go-ent-demo/entcore"

	"context"

	"entgo.io/ent"
)

// 为模型自动设置创建者、创建时间/更新者、更新时间
func CreateUpdateHook() func(next entcore.Mutator) entcore.Mutator {
	return func(next entcore.Mutator) entcore.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			op := m.Op()
			// 非创建和更新操作直接放行
			if op != ent.OpCreate && op != ent.OpUpdate && op != ent.OpUpdateOne {
				return next.Mutate(ctx, m)
			}

			//traceCtx := trace.FromContext(ctx)
			// subject, err := traceCtx.GetSubject()
			// if err != nil {
			// 	log.Warnf("[ORM CreateUpdateHook] get subject failed: %s", err)
			// }

			// switch op {
			// case ent.OpCreate:
			// 	// CreatedAt
			// 	if setter, ok := m.(interface{ SetCreatedAt(types.CustomTime) }); ok {
			// 		setter.SetCreatedAt(types.NewCustomTimeNow())
			// 	}

			// 	// CreatedBy
			// 	if subject != nil {
			// 		if setter, ok := m.(interface{ SetCreatedBy(string) }); ok {
			// 			setter.SetCreatedBy(subject.UserName)
			// 		}
			// 	}
			// case ent.OpUpdate, ent.OpUpdateOne:
			// 	// UpdatedAt
			// 	if setter, ok := m.(interface{ SetUpdatedAt(types.CustomTime) }); ok {
			// 		setter.SetUpdatedAt(types.NewCustomTimeNow())
			// 	}

			// 	// UpdatedBy
			// 	if subject != nil {
			// 		if setter, ok := m.(interface{ SetUpdatedBy(string) }); ok {
			// 			setter.SetUpdatedBy(subject.UserName)
			// 		}
			// 	}
			// }

			return next.Mutate(ctx, m)
		})
	}
}
