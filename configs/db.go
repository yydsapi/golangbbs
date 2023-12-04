// Copyright 2019 golangbbs Core Team.  All rights reserved.
// LICENSE: Use of this source code is governed by AGPL-3.0.
// license that can be found in the LICENSE file.
package configs

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rakyll/statik/fs"
	"github.com/sirupsen/logrus"
	//_ "golangbbs/statik"
	"net/http"
	"os"
	"path/filepath"
)

var Db *sql.DB // global variable to share it between main and the HTTP handler
var err error
var StatikFS http.FileSystem

func Initdb() {
	if BbsConfigs.DbType == "mysql" && BbsConfigs.Precompile != 1 {
		SqlSearchString1 = "concat('%',?,'%')"
		SqlSearchString2 = "edit_history= CONCAT(edit_history,?)"
		SqlSearchString3 = "rand()"
		Db, err = sql.Open("mysql", BbsConfigs.DbConfigs.MySqlConnStr)
		if err != nil {
			LogErr(err)
		}
		Db.SetMaxOpenConns(2000)
		Db.SetMaxIdleConns(1000)
		var version string
		Db.QueryRow("SELECT VERSION()").Scan(&version)
		logrus.Info("Connected to Mysql ok, version: " + version)
		err = Db.Ping() // This DOES open a connection if necessary. This makes sure the database is accessible
		if err != nil {
			LogErr(err)
		}
	} else {
		SqlSearchString1 = "'%' || ? || '%'"
		SqlSearchString2 = "edit_history= edit_history || ?"
		SqlSearchString3 = "random()"
		dir, _ := os.Executable()
		exPath := filepath.Dir(dir)
		if BbsConfigs.Precompile == 1 {
			Db, err = sql.Open("sqlite3", exPath+"/sqlite.db")
		} else {
			Db, err = sql.Open("sqlite3", BbsConfigs.DbConfigs.Sqlite3DbPath)
		}
		if err != nil {
			LogErr(err)
		} else {
			if BbsConfigs.Precompile == 1 {
				logrus.Info("connection sqlite3 db ok, DbPath:" + exPath + "/sqlite.db")
			} else {
				logrus.Info("connection sqlite3 db ok, DbPath:" + BbsConfigs.DbConfigs.Sqlite3DbPath)
			}
		}
	}
	if BbsConfigs.Precompile == 1 {
		StatikFS, err = fs.New()
		if err != nil {
			LogErr(err)
		}
	}
	//defer Db.Close()
}
