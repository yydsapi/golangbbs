// Copyright 2019 golangbbs Core Team.  All rights reserved.
// LICENSE: Use of this source code is governed by AGPL-3.0.
// license that can be found in the LICENSE file.
package webcodev1

import (
	"net/http"
	"strings"
	"encoding/json"
    "golangbbs/configs"
	//"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
    "strconv"
)

type User struct {
    Id    string   `json:"id"`
    MyUserName     string  `json:"myusername"`
    Regtime  string   `json:"regtime"`
    Warn  int64   `json:"warn"`
    Sex     string  `json:"sex"`
    Email     string  `json:"email"`
    Allow  int64   `json:"allow"`
    Score  float64   `json:"score"`
    HeadIcon  string   `json:"headicon"`
    LevelIcon  string   `json:"levelicon"`
}

func ManageEditMenuTree(c *gin.Context) {
    configs.NavMenuStr=""
    configs.Initmenu()
    if SessionUserAdmin<100 {
        c.JSON(200, gin.H{"info": configs.Translate("common.plif")})
        return
    }
    h := DefaultH(c)
    h["Content"] = configs.TreeStr
    h["Title"] = "Welcome to menu edit"
	c.HTML(http.StatusOK, "manage/editmenutree", h)
}

//UsersManage handles GET
func UsersManage(c *gin.Context) {
	h := DefaultH(c)
    currentpage:=c.DefaultQuery("currentpage", "1")
    currentpage = strings.TrimSpace(currentpage)
    currentpage = configs.EscapeWords(currentpage)
    pagesize:=c.DefaultQuery("pagesize", "30")
    pagesize = strings.TrimSpace(pagesize)
    pagesize = configs.EscapeWords(pagesize)
    intcurrentpage,_:=strconv.Atoi(currentpage)
    intpagesize,_:=strconv.Atoi(pagesize)
    var itemCount int
    var rows *sql.Rows
    var err error
    if SessionUserAdmin<100 {
        c.JSON(200, gin.H{"info": configs.Translate("common.plif")})
        return
    }
    err = configs.Db.QueryRow("select count(*) as count  from myuser").Scan(&itemCount)
        if err != nil {
            configs.LogErr(err)
        }
    h["ItemCounts"]=itemCount
    rows, err = configs.Db.Query("select id,myusername,regtime,warn,sex,email,allow,headicon,score from myuser order by id desc limit ?,?",(intcurrentpage-1)*intpagesize,pagesize)
    if err != nil {
        configs.LogErr(err)
    }
    defer rows.Close()
    var intUserLevel int
    h["CurrentPage"]=intcurrentpage
    h["PageSize"]=intpagesize
    ms := make([]*User, 0)
    for rows.Next() {
        m := &User{}
        var id sql.NullString
        var myusername sql.NullString
        var regtime sql.NullString
        var warn sql.NullInt64
        var sex sql.NullString
        var email sql.NullString
        var allow sql.NullInt64
        var headicon sql.NullString
        var score sql.NullFloat64
        if err = rows.Scan(&id,&myusername,&regtime,&warn,&sex,&email,&allow,&headicon,&score); err != nil {
        	configs.LogErr(err)
        }
        if id.Valid {
            m.Id = id.String
            m.MyUserName = myusername.String
            m.Regtime = regtime.String
            m.Warn=warn.Int64
            m.Sex = sex.String
            m.Email = email.String
			m.Allow = allow.Int64
            m.Score = configs.TruncateNaive(score.Float64,0.01)
            m.HeadIcon = headicon.String
            m.LevelIcon = "/static/img/level/"+strconv.Itoa(intUserLevel)+".png"
        }
        ms = append(ms, m)
    }
    b, err := json.Marshal(ms)
    if err != nil {
       configs.LogErr(err)
    }
	h["Content"]=string(b)
	c.HTML(http.StatusOK, "manage/usersmanage", h)
}

func UsersManagePost(c *gin.Context) {
	h := DefaultH(c)
	if configs.CheckUpdateCount(configs.GetFuncName(),30,c,Csrf_Result) {
        c.JSON(200, gin.H{"info": configs.Translate("tips.morethannrde")})
        return
    }
	h["WebTitle"] = "modify"
	cid:=c.PostForm("cid")
	cid = strings.TrimSpace(cid)
	mypassword:=c.PostForm("mypassword")
	mypassword = strings.TrimSpace(mypassword)
	allow:=c.PostForm("allow")
	allow = strings.TrimSpace(allow)
	score:=c.PostForm("score")
	score = strings.TrimSpace(score)
	cid = configs.EscapeWords(cid)
	mypassword = configs.EscapeWords(mypassword)
	score = configs.EscapeWords(score)
	allow = configs.EscapeWords(allow)
	if (mypassword!="" && strings.Count(mypassword,"")-1<4 ){
		c.JSON(200, gin.H{"info": configs.Translate("common.inputerr")})
        return
	}
	    var itemCount int
		err := configs.Db.QueryRow("select count(*) as count  FROM myuser where id="+cid).Scan(&itemCount)
	    if err != nil {
			c.JSON(200, gin.H{"info": configs.Translate("common.ftudb")})
		    return
	        panic(err.Error()) // proper error handling instead of panic in your app
	    }
	    if itemCount ==1 {
    	    var err error
	    	if mypassword=="" {
				_,err = configs.Db.Exec("update myuser set score=?,allow=? where id=?",score,allow,cid)
	    	}else{
	    		_,err = configs.Db.Exec("update myuser set mypwd=?,score=?,allow=? where id=?",configs.Md5Encrypt(mypassword),score,allow,cid)
	    	}
		    if err != nil {
				c.JSON(200, gin.H{"info": configs.Translate("common.opfailed")})
		        return
		    }
				c.JSON(200, gin.H{"info": "ok"})
		        return
	    }else{
			c.JSON(200, gin.H{"info": configs.Translate("common.ftudb")})
	        return
	    }
}