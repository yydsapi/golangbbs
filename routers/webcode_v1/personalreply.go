// Copyright 2019 golangbbs Core Team.  All rights reserved.
// LICENSE: Use of this source code is governed by AGPL-3.0.
// license that can be found in the LICENSE file.
package webcodev1

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
    "golangbbs/configs"
    "strconv"
    "strings"
    "time"
)

func PersonalReplyGet(c *gin.Context) {
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
    var st []rune
    if SessionUserId=="" {
        c.JSON(200, gin.H{"info": configs.Translate("common.plif")})
        return
    }
    err = configs.Db.QueryRow("select count(*) as count  from bbsreply where allow>0 and replyuserid=?",SessionUserId).Scan(&itemCount)
        if err != nil {
            configs.LogErr(err)
        }
        h["ItemCounts"]=itemCount
    rows, err = configs.Db.Query("select id,pid,replycontent,replytime,mylink from bbsreply where allow>0 and replyuserid=? order by replytime desc limit ?,?",SessionUserId,(intcurrentpage-1)*intpagesize,pagesize)
    if err != nil {
        configs.LogErr(err)
    }
    defer rows.Close()
    h["CurrentPage"]=intcurrentpage
    h["PageSize"]=intpagesize
    ms := make([]*Bbsreply, 0)
    for rows.Next() {
        m := &Bbsreply{}
        var id sql.NullString
        var pid sql.NullString
        var replycontent sql.NullString
        var replytime sql.NullString
        var link sql.NullString
        if err = rows.Scan(&id,&pid, &replycontent, &replytime,&link); err != nil {
        	configs.LogErr(err)
        }
        if id.Valid {
            m.Id = id.String
            m.Pid = pid.String
            st=[]rune(replycontent.String)
            if len(st)>200{
               m.Replycontent= Tpl(configs.UnEscapeWords(string(st[0:200])+"..."))
            }else{
                m.Replycontent=Tpl(configs.UnEscapeWords(replycontent.String))
            }
            m.Replytime = replytime.String
            m.Link = link.String
            m.Img = "/static/img/reply.png"
        }
        ms = append(ms, m)
    }
    b, err := json.Marshal(ms)
    if err != nil {
       configs.LogErr(err)
    }
	h["Content"]=string(b)
	c.HTML(http.StatusOK, "personalreply/list", h)
}
//modify reply
func EditReplyById(c *gin.Context) {
   h := DefaultH(c)
   id:=c.Param("id")
   h["PageTitle"] = "[[[$t('common.edit')]]]"
    id = strings.TrimSpace(id)
    id = configs.EscapeWords(id)
    rows, err := configs.Db.Query("select id,pid,replycontent,replytime,mylink from bbsreply where id=?",id)
    if err != nil {
        configs.LogErr(err)
    }
    defer rows.Close()
    IntSucess=0
    for rows.Next() {
        //m := &Bbs{}
        var id sql.NullString
        var pid sql.NullString
        var replycontent sql.NullString
        var replytime sql.NullString
        var link sql.NullString
        if err = rows.Scan(&id,&pid, &replycontent, &replytime,&link); err != nil {
            configs.LogErr(err)
        }
        if id.Valid {
            h["Id"] = id.String
            h["Pid"] = pid.String
            h["Replycontent"] = Tpl(configs.UnEscapeWords(replycontent.String))
            h["Replytime"] = replytime.String
            IntSucess=1
        }
        //ms = append(ms, m)
    }
    c.HTML(http.StatusOK, "bbs/personal/editreply", h)
}
func UpdateReplyPost(c *gin.Context) {
    h := DefaultH(c)
    if configs.CheckUpdateCount(configs.GetFuncName(),30,c,Csrf_Result) {
        c.JSON(200, gin.H{"info": configs.Translate("tips.morethannrde")})
        return
    }
    //session := sessions.Default(c)
    h["WebTitle"] = Tpl("modify")
    if SessionUserId=="" {
        c.JSON(200, gin.H{"info": configs.Translate("common.plif")})
        return
    }
    id := c.PostForm("id")
    var itemAuthor string
    useriderr := configs.Db.QueryRow("select replyuserid from bbsreply where id=?",id).Scan(&itemAuthor)
    if useriderr != nil {
        configs.LogErr(useriderr)
        return
    }

    if itemAuthor != SessionUserId && SessionUserAdmin<100 {
        c.JSON(200, gin.H{"info": configs.Translate("common.nopernission")})
        return  
    }
    content:=c.PostForm("content")
    content = strings.TrimSpace(content)
    content= CheckFilterResult(content)
    content = configs.EscapeWords(content)
    timeUnix:=time.Now().Unix()
    formatTimeStr:=time.Unix(timeUnix,0).Format("2006-01-02 15:04:05")
    edit_history:="<font class='edit_history'>-----"+configs.Translate("The content was edited by ")  + SessionUserId + configs.Translate(" at ") + formatTimeStr + "-----</font><br>"
    result,err := configs.Db.Exec("update bbsreply set "+ configs.SqlSearchString2 +",edit_man=?,edit_time=?,replycontent=? where id=?",edit_history,SessionUserId,formatTimeStr,content,id)
    if err != nil{
        configs.LogErr(err)
        c.JSON(http.StatusOK, gin.H{"info": configs.Translate("common.serverError")})
        return
    }
    _,err = result.RowsAffected()
    if err != nil {
        configs.LogErr(err)
        c.JSON(http.StatusOK, gin.H{"info": configs.Translate("common.serverError")})
        return
    }
    returnURL := c.DefaultQuery("return","/")
    c.JSON(200, gin.H{"info": "ok","returnURL": returnURL})
    return
}
func DeleteReplyPost(c *gin.Context) {
    h := DefaultH(c)
    if configs.CheckUpdateCount(configs.GetFuncName(),30,c,Csrf_Result) {
        c.JSON(200, gin.H{"info": configs.Translate("tips.morethannrde")})
        return
    }
    //session := sessions.Default(c)
    h["WebTitle"] = "delete"
    id:=c.PostForm("id")
    if SessionUserId=="" {
        c.JSON(200, gin.H{"info": configs.Translate("common.plif")})
        return
    }

    var itemAuthor string
    useriderr := configs.Db.QueryRow("select replyuserid from bbsreply where id=?",id).Scan(&itemAuthor)
    if useriderr != nil {
        configs.LogErr(useriderr)
        return
    }

    if itemAuthor != SessionUserId && SessionUserAdmin<100 {
        c.JSON(200, gin.H{"info": configs.Translate("common.nopernission")})
        return  
    }

    //result,err := configs.Db.Exec("delete from bbs where id=?",id)
    _,err := configs.Db.Exec("update bbsreply set allow=0 where id=?",id)
    if err != nil{
        configs.LogErr(err)
        c.JSON(200, gin.H{"info": configs.Translate("common.opfailed")})
        return
    }
    c.JSON(200, gin.H{"info": "ok"})
}