package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
)

func doAuth() error {
	dbType := cel.DB.DataType
	fileName := fmt.Sprintf("%d_create_auth_tables", time.Now().UnixMicro())
	upFile := cel.RootPath + "/migrations/" + fileName + ".up.sql"
	downFile := cel.RootPath + "/migrations/" + fileName + ".down.sql"

	log.Println(dbType, upFile, downFile)

	err := copyFilefromTemplate("templates/migrations/auth_tables."+dbType+".sql", upFile)

	if err != nil {
		exitGracefully(err)
	}

	err = copyDataToFile([]byte("drop table if exists users cascade; drop table if exists tokens cascade; drop table if exists remember_tokens;"), downFile)
	if err != nil {
		exitGracefully(err)
	}

	// run migration
	err = doMigrate("up", "")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/data/user.go.txt", cel.RootPath+"/data/user.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/data/token.go.txt", cel.RootPath+"/data/token.go")
	if err != nil {
		exitGracefully(err)
	}

	// copy over middleware
	err = copyFilefromTemplate("templates/middleware/auth.go.txt", cel.RootPath+"/middleware/auth.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFilefromTemplate("templates/middleware/auth-token.go.txt", cel.RootPath+"/middleware/auth-token.go")
	if err != nil {
		exitGracefully(err)
	}

	color.Yellow("  - users, tokens, and remember_tokens migrations created and executed")
	color.Yellow("  - user and token models created")
	color.Yellow("  - auth middleware created")
	color.Yellow("")
	color.Yellow("Don't forget to add user and token models in data/models.go, and to add appropriate middleware to your routes!")

	return nil
}
