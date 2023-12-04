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
    "path/filepath"
)
func FaviconGet(c *gin.Context) {
    c.Redirect(http.StatusSeeOther, "/static/favicon.ico")
}
//HomeGet handles
func IndexGet(c *gin.Context) {
    h := DefaultH(c)
    searchkey:=c.DefaultQuery("searchkey", "")
    searchkey = strings.TrimSpace(searchkey)
    searchkey = configs.EscapeWords(searchkey)
    category:=c.DefaultQuery("category", "")
    category = strings.TrimSpace(category)
    category = configs.EscapeWords(category)
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
    var sqlstr string
    var sqlcountstr1 string
    var sqlcountstr2 string
    sqlstr="select a.id,a.title,a.author,a.add_time,a.reply_count,a.raise_count,a.categoryen,a.edit_man,a.read_count,a.mylink,a.mymedia,a.reader,a.isprivate,a.attachment,b.headicon,b.blog,b.score from bbs a LEFT JOIN myuser b on a.author=b.myusername"
    sqlcountstr1="select count(*) as count  from bbs where allow>0 and isprivate=0"
    sqlcountstr2="select count(*) as count  from bbs where (isprivate=0 or (isprivate=1 and reader like "+ configs.SqlSearchString1 +")) and allow>0"
    if SessionUserId=="" {
        if searchkey=="" {
            if category==""{
                err = configs.Db.QueryRow(sqlcountstr1).Scan(&itemCount)
                rows, err = configs.Db.Query(sqlstr+" where a.allow>0 and a.isprivate=0 order by a.reply_time desc,a.raise_count desc, a.id desc limit ?,?",(intcurrentpage-1)*intpagesize,pagesize)
            }else{
                err = configs.Db.QueryRow(sqlcountstr1 + " and categoryen=?",category).Scan(&itemCount)
                rows, err = configs.Db.Query(sqlstr+" where a.allow>0 and a.isprivate=0 and a.categoryen=? order by a.reply_time desc,a.raise_count desc, a.id desc limit ?,?",category,(intcurrentpage-1)*intpagesize,pagesize)
            }
        }else{
            err = configs.Db.QueryRow(sqlcountstr1 + " and ( title like "+ configs.SqlSearchString1 +" or content like "+ configs.SqlSearchString1 +")",searchkey,searchkey).Scan(&itemCount)
            rows, err = configs.Db.Query(sqlstr+" where a.allow>0 and a.isprivate=0 and (a.title like "+ configs.SqlSearchString1 +" or a.content like "+ configs.SqlSearchString1 +") order by a.reply_time desc,a.raise_count desc, a.id desc limit ?,?",searchkey,searchkey,(intcurrentpage-1)*intpagesize,pagesize)          
        }
    }else{
        if searchkey=="" {
            if category==""{
                err = configs.Db.QueryRow(sqlcountstr2,","+SessionUserId+",").Scan(&itemCount)
                rows, err = configs.Db.Query(sqlstr+" where (a.isprivate=0 or (a.isprivate=1 and a.reader like "+ configs.SqlSearchString1 +")) and a.allow>0 order by a.reply_time desc,a.raise_count desc, a.id desc  limit ?,?",","+SessionUserId+",",(intcurrentpage-1)*intpagesize,pagesize)
            }else{
                err = configs.Db.QueryRow(sqlcountstr2+" and categoryen=?",","+SessionUserId+",",category).Scan(&itemCount)
                rows, err = configs.Db.Query(sqlstr+" where (a.isprivate=0 or (a.isprivate=1 and a.reader like "+ configs.SqlSearchString1 +")) and a.allow>0 and a.categoryen=? order by a.reply_time desc,a.raise_count desc, a.id desc  limit ?,?",","+SessionUserId+",",category,(intcurrentpage-1)*intpagesize,pagesize)
            }
        }else{
                err = configs.Db.QueryRow(sqlcountstr2+" and (title like "+ configs.SqlSearchString1 +" or content like "+ configs.SqlSearchString1 +")",","+SessionUserId+",",searchkey,searchkey).Scan(&itemCount)
                rows,err = configs.Db.Query(sqlstr+" where (a.isprivate=0 or (a.isprivate=1 and a.reader like "+ configs.SqlSearchString1 +")) and a.allow>0 and (a.title like "+ configs.SqlSearchString1 +" or a.content like "+ configs.SqlSearchString1 +") order by a.reply_time desc,a.raise_count desc, a.id desc limit ?,?",","+SessionUserId+",",searchkey,searchkey,(intcurrentpage-1)*intpagesize,pagesize)
        } 
    }
    configs.LogErr(err)
    h["ItemCounts"]=itemCount
    defer rows.Close()
    h["CurrentPage"]=intcurrentpage
    h["PageSize"]=intpagesize
    ms := make([]*Bbs, 0)
    for rows.Next() {
        m := &Bbs{}
        var id sql.NullString
        var title sql.NullString
        var author sql.NullString
        var reader sql.NullString
        var add_time sql.NullString
        var reply_count sql.NullString
        var raise_count sql.NullFloat64
        var categoryen sql.NullString
        var edit_man sql.NullString
        var read_count sql.NullString
        var link sql.NullString
        var media sql.NullString
        var isprivate sql.NullFloat64
        var attachment sql.NullString
        var userheadicon sql.NullString
        var userblog sql.NullBool
        var userscore sql.NullFloat64
        if err = rows.Scan(&id,&title, &author, &add_time,&reply_count,&raise_count,&categoryen,&edit_man,&read_count,&link,&media,&reader,&isprivate,&attachment,&userheadicon,&userblog,&userscore); err != nil {
            configs.LogErr(err)
        }
        if id.Valid {
            m.Id = id.String
            if strings.TrimSpace(attachment.String)!=""{
                m.Title = Tpl(`<img src='/static/img/att.png' alt='attachment' style='width: 16px;height:16px'> `+configs.UnEscapeWords(title.String))
            }else{
                m.Title = Tpl(configs.UnEscapeWords(title.String))
            }
            m.Author = author.String
            m.Add_time = add_time.String
            if IsMobile{
                m.Add_time = add_time.String [0:10]
            }            
            m.Reply_count = reply_count.String
            m.Raise_count=raise_count.Float64
            m.Category = categoryen.String
            m.Edit_man = edit_man.String
            m.Read_count = read_count.String
            m.Link = link.String
            m.Isprivate = isprivate.Float64
            h["UserBlogTmp"] = userblog.Bool
            if isprivate.Float64 != 0 {
                if strings.Index(reader.String,",*,")==-1 {
                    m.PrivateImg = "/static/img/isprivate.png"
                }else{
                    m.PrivateImg = "/static/img/isprivate_share.png"
                }
            }else{
                m.PrivateImg = "/static/img/space.png"
            }
            m.UserHeadIcon = userheadicon.String
            m.UserLevel = "/static/img/level/"+strconv.Itoa(getUserLevel(userscore.Float64))+".png"
            m.ImgRaise=ShowRaiseImg(reply_count.String,read_count.String,m.Raise_count,"")
            m.Img=ShowImg(reply_count.String,read_count.String,m.Raise_count,"")
            if strings.TrimSpace(media.String)!=""{
                fileExt := filepath.Ext(strings.TrimSpace(strings.ToLower(media.String)))
                if fileExt==".mp4" || fileExt==".mkv" || fileExt==".avi" || fileExt==".rmvb" {
                    m.Img="/static/img/film.png"
                }else{
                    m.Img="/static/img/music.png"
                }
        }

        }
        ms = append(ms, m)
    }
    b, err := json.Marshal(ms)
    configs.LogErr(err)
    h["Searchkey"]=searchkey
    h["Content"]=string(b)
    c.HTML(http.StatusOK, "index/list", h)
}

func AboutGet(c *gin.Context) {
    h := DefaultH(c)
    c.HTML(http.StatusOK, "index/about", h)
}
