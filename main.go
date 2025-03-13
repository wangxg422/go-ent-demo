package main

import (
	"context"
	"fmt"
	"log"

	"go-ent-demo/ent"
	"go-ent-demo/ent/user"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
    clientProd, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
    if err != nil {
        log.Fatalf("failed opening connection to sqlite: %v", err)
    }
    defer clientProd.Close()
    // Run the auto migration tool.
    if err := clientProd.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }

	client := clientProd.Debug();

	ctx := context.Background();


	_,err = client.Dept.Create().SetID(1).SetParentID(0).SetDeptCode("develop").SetDeptName("develop department").Save(ctx);
	if err != nil {
		log.Fatalln(err)
	}

	_,err = client.Role.Create().SetID(1).SetRoleCode("role1").Save(ctx);
	if err != nil {
		log.Fatalln(err)
	}

	_,err = client.Role.Create().SetID(2).SetRoleCode("role2").Save(ctx);
	if err != nil {
		log.Fatalln(err)
	}

	_,err = client.User.Create().SetID(1).SetUserName("zhangsan").SetSex(1).SetDeptID(1).AddSysRoleIDs(1, 2).Save(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	query_user,err := client.User.Query().Where(user.UserNameEQ("zhangsan")).Only(ctx);
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(query_user)
}