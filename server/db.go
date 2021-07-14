package main

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

// テスト環境(mac docker)
//const CONNECTION_STRING string = "sqlserver://sa:Test1234!@localhost?database=Chi8"

// 本番環境
const CONNECTION_STRING string = "sqlserver://sa:Pass@word1@localhost?database=Chi8"

// 商品画像パスローカル(mac)
//const ITEM_IMAGE_FOLDER string = `/Library/WebServer/Documents/img/items/%s`
//const ITEM_IMAGE_PATH string = `http://chi8.store/img/items/%s/%s`

// 商品画像パス本番(windows)
const ITEM_IMAGE_FOLDER string = `C:\Apache24\htdocs\img\items\%s`
const ITEM_IMAGE_PATH string = `https://chi8.store/img/items/%s/%s`

// 一行を取得
func GetQueryRows(sql string) *sql.Rows {

	db := GetDBConnection()
	rows, err := db.Query(sql)
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	return rows
}

// DB接続
func GetDBConnection() *sql.DB {

	db, err := sql.Open("sqlserver", CONNECTION_STRING)

	if err != nil {
		panic(err.Error())
	}

	return db
}
