package main

import (
	"context"
	"fmt"
	"go-ent-demo/ent"
	"go-ent-demo/ent/sysuser"
	"log"

	_ "github.com/mattn/go-sqlite3"
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

func main() {
	clientProd := NewClient()
	defer clientProd.Close()
	Migration(clientProd)
	
	client := clientProd.Debug()

	ctx := context.Background()

	_, err := client.SysDept.Create().SetID(1).SetParentID(0).SetDeptCode("develop").SetDeptName("develop department").Save(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = client.SysRole.Create().SetID(1).SetRoleCode("role1").Save(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = client.SysRole.Create().SetID(2).SetRoleCode("role2").Save(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = client.SysUser.Create().SetID(1).SetUserName("zhangsan").SetSex(1).SetDeptID(1).AddSysRoleIDs(1, 2).Save(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = client.SysUser.Update().Where(sysuser.UserNameEQ("zhangsan")).SetNickName("zhangsan").Save(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	query_user, err := client.SysUser.Query().Where(sysuser.UserNameEQ("zhangsan")).Only(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(query_user)
}
