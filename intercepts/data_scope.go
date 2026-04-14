package interceptors

import (
	"context"
	"go-ent-demo/entcore"

	"entgo.io/ent"
)

func DataScopeQuery() entcore.Interceptor {
	return entcore.InterceptFunc(func(next entcore.Querier) entcore.Querier {
		return ent.QuerierFunc(func(ctx context.Context, q ent.Query) (ent.Value, error) {
			// traceCtx := trace.FromContext(ctx)
			// subject, err := traceCtx.GetSubject()
			// if err != nil {
			// 	log.Warnf("[ORM CreateUpdateHook] get subject failed: %s", err)
			// }
			// var dataScope perm.DataScope = "0"
			// dataScopeCustomDeptIDs := []types.ID{}

			// // 获取 Query Builder 的 SQL selector（Generic，无实体绑定）
			// selector, ok := q.(interface{ QueryBuilder() *sql.Selector })
			// if !ok {
			// 	return next.Query(ctx, q)
			// }

			// s := selector.QueryBuilder()

			// switch dataScope {

			// case perm.DataScopeAll:
			// 	// 不限制
			// 	return next.Query(ctx, q)

			// case perm.DataScopeSelf:
			// 	// 仅本人
			// 	s.Where(sql.EQ(s.C("created_by"), ""))

			// case perm.DataScopeDept:
			// 	// 本部门
			// 	s.Where(sql.EQ(s.C("dept_id"), ""))

			// case perm.DataScopeDeptAndSub:
			// 	// 子部门需要你已有 dept tree，示例用一个 SQL IN 子查询占位
			// 	s.Where(
			// 		sql.P(func(b *sql.Builder) {
			// 			b.WriteString(fmt.Sprintf(
			// 				"%s IN (SELECT id FROM sys_dept WHERE path LIKE (SELECT CONCAT(path, '%%') FROM sys_dept WHERE id = %d))",
			// 				s.C("dept_id"),
			// 				0,
			// 			))
			// 		}),
			// 	)

			// case perm.DataScopeCustom:
			// 	if len(dataScopeCustomDeptIDs) > 0 {
			// 		s.Where(sql.In(s.C("dept_id"), dataScopeCustomDeptIDs))
			// 	} else {
			// 		// 无数据权限
			// 		s.Where(sql.In(s.C("dept_id"), -1))
			// 	}
			// }

			return next.Query(ctx, q)
		})
	})
}
