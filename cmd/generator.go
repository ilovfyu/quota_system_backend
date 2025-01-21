package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const MySQLDSN = "root:root#123@tcp(127.0.0.1:3306)/quota_system?charset=utf8&parseTime=True&loc=Asia%2FShanghai"

func connectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(MySQLDSN), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("connect db failed, err: %w", err))
	}
	return db
}

func main() {
	g := gen.NewGenerator(gen.Config{

		OutPath: "./dal/query",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(connectDB(MySQLDSN))
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
