package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Users holds the schema definition for the Users entity.
type Users struct {
	ent.Schema
}

// Fields of the Users.
func (Users) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "bigint", // Override MySQL.
		}).Unique(),

		field.String("mobile").SchemaType(map[string]string{
			dialect.MySQL: "varchar(11)", // Override MySQL.
		}).Comment("手机号码，用户唯一标识"),

		field.String("password").SchemaType(map[string]string{
			dialect.MySQL: "varchar(100)", // Override MySQL.
		}),

		field.String("nick_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(25)", // Override MySQL.
		}).Optional().Comment("用户昵称"),

		field.Time("birthday").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Optional().Comment("出生日日期"),

		field.String("gender").SchemaType(map[string]string{
			dialect.MySQL: "varchar(16)", // Override MySQL.
		}).Optional().Default("male").Comment("female:女,male:男"),

		field.Int32("role").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Optional().Default(1).Comment("1:普通用户，2:管理员"),

		field.Time("add_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime(3)", // Override MySQL.
		}).Optional(),

		field.Time("update_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime(3)", // Override MySQL.
		}).Optional(),

		field.Time("deleted_at").SchemaType(map[string]string{
			dialect.MySQL: "datetime(3)", // Override MySQL.
		}).Optional(),

		field.Int8("is_deleted_at").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)", // Override MySQL.
		}).Optional(),
	}

}

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return nil
}
