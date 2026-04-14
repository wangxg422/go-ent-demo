package schematype

// ID类型（包含主键、外键）在数据库中的数据类型映射
func SchemaTypeID() map[string]string {
	return map[string]string{
		"mysql": "bigint",
	}
}
