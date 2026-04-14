package schematype

import "entgo.io/ent/dialect"

// 数据库字段类型为char(1)
func SchemaTypeChar1() map[string]string {
	return map[string]string{
		dialect.MySQL:    "char(1)",
		dialect.Postgres: "char(1)",
	}
}
