package crud

import (
	"context"
	"go-ent-demo/ent"
	"go-ent-demo/ent/sysrole"
	"go-ent-demo/ent/sysuser"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewClient() *ent.Client {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
		panic("failed opening connection to sqlite")
	}
	return client
}

func Migration(client *ent.Client) {
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func TestEntRelationsWithAssert(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)

	// 使用内存 SQLite 数据库
	_client := NewClient()
	client := _client.Debug()
	defer client.Close()

	// 自动创建表
	Migration(client)

	// --- 1. 创建部门 ---
	dept, err := client.SysDept.
		Create().
		SetDeptName("研发部").
		Save(ctx)
	assert.NoError(err)
	assert.NotNil(dept)

	// --- 2. 创建角色 ---
	roleAdmin, err := client.SysRole.Create().SetRoleName("管理员").Save(ctx)
	assert.NoError(err)
	roleUser, err := client.SysRole.Create().SetRoleName("普通用户").Save(ctx)
	assert.NoError(err)

	// --- 3. 创建用户，并关联部门 ---
	user1, err := client.SysUser.
		Create().
		SetUserName("Alice").
		SetSysDept(dept).
		AddSysRoles(roleAdmin, roleUser). // 多对多关联角色
		Save(ctx)
	assert.NoError(err)
	assert.NotNil(user1)

	user2, err := client.SysUser.
		Create().
		SetUserName("Bob").
		SetSysDept(dept).
		AddSysRoles(roleUser).
		Save(ctx)
	assert.NoError(err)
	assert.NotNil(user2)

	// --- 4. 测试查询用户的部门 ---
	u1, err := client.SysUser.Query().Where(sysuser.IDEQ(user1.ID)).QuerySysDept().Only(ctx)
	assert.NoError(err)
	assert.Equal("研发部", u1.DeptName)

	// --- 5. 测试查询用户的角色 ---
	roles, err := client.SysUser.Query().Where(sysuser.IDEQ(user1.ID)).QuerySysRoles().All(ctx)
	assert.NoError(err)
	roleNames := []string{}
	for _, r := range roles {
		roleNames = append(roleNames, r.RoleName)
	}
	assert.ElementsMatch([]string{"管理员", "普通用户"}, roleNames)

	// --- 6. 测试查询角色的用户 ---
	usersWithAdmin, err := client.SysRole.Query().Where(sysrole.IDEQ(roleAdmin.ID)).QuerySysUsers().All(ctx)
	assert.NoError(err)
	userNames := []string{}
	for _, u := range usersWithAdmin {
		userNames = append(userNames, u.UserName)
	}
	assert.ElementsMatch([]string{"Alice"}, userNames)
}
