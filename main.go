package main

import (
	"context"
	"entsoftdelete/ent"
	"entsoftdelete/ent/task"
	"entsoftdelete/ent/user"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&_fk=1", ent.Debug(), ent.Log(func(s ...any) {
		fmt.Println(s...)
	}))
	if err != nil {
		panic(err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	ctx := context.Background()

	{

		_, err = client.User.Create().SetName("a").Save(ctx)
		if err != nil {
			panic(err)
		}
		u, err := client.User.Query().Where(user.Name("a")).First(ctx)
		if err != nil {
			panic(err)
		}
		if err := client.User.DeleteOne(u).Exec(ctx); err != nil {
			panic(err)
		}

		_, err = client.User.Query().Where(user.Name("a")).First(ctx)
		if err == nil {
			panic("found no soft delete user")
		} else {
			if !ent.IsNotFound(err) {
				panic(err)
			}
		}

		nu, err := client.User.Query().Real().Where(user.Name("a")).First(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Println(nu)

		if err := client.User.DeleteOne(nu).Real().Exec(ctx); err != nil {
			panic(err)
		}
		_, err = client.User.Query().Real().Where(user.Name("a")).First(ctx)
		if err == nil {
			panic("found no delete user")
		} else {
			if !ent.IsNotFound(err) {
				panic(err)
			}
		}
	}
	{

		_, err = client.Task.Create().SetName("a").Save(ctx)
		if err != nil {
			panic(err)
		}
		u, err := client.Task.Query().Where(task.Name("a")).First(ctx)
		if err != nil {
			panic(err)
		}
		if err := client.Task.DeleteOne(u).Exec(ctx); err != nil {
			panic(err)
		}
		_, err = client.Task.Query().Where(task.Name("a")).First(ctx)
		if err == nil {
			panic("found no soft delete user")
		} else {
			if !ent.IsNotFound(err) {
				panic(err)
			}
		}

	}
}
