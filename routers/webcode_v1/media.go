// Copyright 2019 golangbbs Core Team.  All rights reserved.
// LICENSE: Use of this source code is governed by AGPL-3.0.
// license that can be found in the LICENSE file.
package webcodev1

import (
	"database/sql"
	"golangbbs/configs"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

func PhotoGet(c *gin.Context) {
	h := DefaultH(c)
	dirfromurl := c.DefaultQuery("dir", "")
	dirfromurl = strings.TrimSpace(dirfromurl)
	dirfromurl = configs.EscapeWords(dirfromurl)
	h["WebTitle"] = "Welcome"
	rootpath := configs.BbsConfigs.BbsUploadPath + "/PicturesDisplay/photos"
	if configs.BbsConfigs.BbsUploadPath == "" {
		dir, _ := os.Executable()
		exPath := filepath.Dir(dir)
		rootpath = exPath + "/documents/PicturesDisplay/photos"
	}
	strurl := ""
	if dirfromurl != "" {
		rootpath += "/" + strings.Replace(dirfromurl, ".", "/", -1)
	}
	logrus.Info("rootpath:", rootpath)
	//strurl = "/documents/PicturesDisplay/photos/"
	strurl = strings.Replace(rootpath, configs.BbsConfigs.BbsUploadPath, "/documents", -1) + "/"
	f, _ := os.Open(rootpath)
	names, _ := f.Readdirnames(-1)
	f.Close()
	strphoto := ""
	strdir := ""
	sort.Strings(names)
	strreturn := ""
	if strings.Index(dirfromurl, ".") != -1 {
		fileSuffix := filepath.Ext(dirfromurl)
		strreturn = strings.TrimSuffix(dirfromurl, fileSuffix)
	}
	//logrus.Info("strreturn:",strreturn)
	if dirfromurl != "" {
		if strreturn == "" {
			strdir += `<a href="/bbs/photos" title="` + configs.Translate("upper level") + `"><img src="/static/img/return.png" class="imgreturn"></a> `
		} else {
			strdir += `<a href="/bbs/photos/?dir=` + strreturn + `" title="` + configs.Translate("upper level") + `"><img src="/static/img/return.png" class="imgreturn"></a> `
		}
	} else {
		strdir += `<a href="/" title="` + configs.Translate("upper level") + `"><img src="/static/img/return.png" class="imgreturn"></a> `
	}
	for _, filename := range names { // loop through the files
		fpath := filepath.Join(rootpath, filename) // Splicing full path
		//logrus.Info("fpath------->",fpath)
		fio, _ := os.Lstat(fpath) // Construct file structure
		if fio.IsDir() {          // If the current file traversed is a directory, enter the directory for recursion
			//logrus.Info(dirfromurl,strings.Index(dirfromurl,"."))
			logrus.Info(string([]rune(filename)[0:2]))
			if string([]rune(filename)[0:2]) != "my" {
				if dirfromurl == "" {
					strdir += `<a href="/bbs/photos/?dir=` + filename + `" title="` + filename + `"><img src="/static/img/photo.png"  class="imgreturn"></a>`
				} else {
					if strings.Index(dirfromurl, ".") != -1 {
						strdir += `<a href="/bbs/photos/?dir=` + dirfromurl + "." + filename + `" title="` + filename + `"><img src="/static/img/photo.png"  class="imgreturn"></a>`
					} else {
						strdir += `<a href="/bbs/photos/?dir=` + dirfromurl + "." + filename + `"><img src="/static/img/photo.png"  class="imgreturn"></a>`
					}
				}
			}
		} else {
			strphoto += `<a href="` + strurl + filename + `" title="` + filename + `" data-gallery=""><img src="` + strurl + filename + `" style="width:60px;height:45px"></a>`
		}
	}
	//logrus.Info("strdir------>",strdir)
	h["UrlPhoto"] = Tpl(strphoto)
	h["UrlDir"] = Tpl(strdir)
	c.HTML(http.StatusOK, "index/photo", h)
}
func VideoGet(c *gin.Context) {
	strurl := ""
	strurlone := ""
	h := DefaultH(c)
	var rows *sql.Rows
	var err error
	rows, err = configs.Db.Query("select mymedia from bbs where allow>0 and isprivate=0 and (mymedia like "+configs.SqlSearchString1+" or mymedia like "+configs.SqlSearchString1+" or mymedia like "+configs.SqlSearchString1+") order by "+configs.SqlSearchString3+" limit 20", "mp4", "mkv", "avi")
	if err != nil {
		configs.LogErr(err)
	}
	defer rows.Close()
	for rows.Next() {
		var link sql.NullString
		if err = rows.Scan(&link); err != nil {
			configs.LogErr(err)
		}
		if link.Valid {
			if strurlone == "" {
				strurlone = link.String
			} else {
				if strurl == "" {
					strurl = strings.Replace(link.String, "/", ":", -1)
				} else {
					strurl += "," + strings.Replace(link.String, "/", ":", -1)
				}
			}
		}

	}
	h["UrlOne"] = strurlone
	h["Url"] = strurl
	c.HTML(http.StatusOK, "index/video", h)
}

func MusicGet(c *gin.Context) {
	strurl := ""
	strurlone := ""
	h := DefaultH(c)
	var rows *sql.Rows
	var err error
	rows, err = configs.Db.Query("select mymedia from bbs where allow>0 and isprivate=0 and (mymedia like "+configs.SqlSearchString1+" or mymedia like "+configs.SqlSearchString1+" or mymedia like "+configs.SqlSearchString1+") order by "+configs.SqlSearchString3+" limit 20", "mp3", "wav", "ape")
	if err != nil {
		configs.LogErr(err)
	}
	defer rows.Close()
	for rows.Next() {
		var link sql.NullString
		if err = rows.Scan(&link); err != nil {
			configs.LogErr(err)
		}
		if link.Valid {
			if strurlone == "" {
				strurlone = link.String
			} else {
				if strurl == "" {
					strurl = strings.Replace(link.String, "/", ":", -1)
				} else {
					strurl += "," + strings.Replace(link.String, "/", ":", -1)
				}
			}
		}

	}
	h["UrlOne"] = strurlone
	h["Url"] = strurl
	c.HTML(http.StatusOK, "index/music", h)
}
