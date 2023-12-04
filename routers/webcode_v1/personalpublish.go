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
    "path/filepath"
)

func PersonalPublishGet(c *gin.Context) {
	h := DefaultH(c)
    category:=c.DefaultQuery("category", "jd_ly")
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
    var isprivateimg string
    if SessionUserId=="" {
        c.JSON(200, gin.H{"info": configs.Translate("common.plif")})
        return
    }
    err = configs.Db.QueryRow("select count(*) as count  from bbs where allow>0 and author=?",SessionUserId).Scan(&itemCount)
        if err != nil {
            configs.LogErr(err)
        }
    h["ItemCounts"]=itemCount
    rows, err = configs.Db.Query("select id,title,author,add_time,reply_count,raise_count,categoryen,edit_man,read_count,mylink,mymedia,reader,isprivate,attachment from bbs where allow>0 and author=? order by id desc limit ?,?",SessionUserId,(intcurrentpage-1)*intpagesize,pagesize)
    if err != nil {
        configs.LogErr(err)
    }
    defer rows.Close()
    h["CurrentPage"]=intcurrentpage
    h["PageSize"]=intpagesize
    ms := make([]*Bbs, 0)
    for rows.Next() {
        m := &Bbs{}
        var id sql.NullString
        var title sql.NullString
        var reader sql.NullString
        var author sql.NullString
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
        if err = rows.Scan(&id,&title, &author, &add_time,&reply_count,&raise_count,&categoryen,&edit_man,&read_count,&link,&media,&reader,&isprivate,&attachment); err != nil {
        	configs.LogErr(err)
        }
        if id.Valid {
            if isprivate.Float64 != 0 {
                if strings.Index(reader.String,",*,")==-1 {
                    isprivateimg="isprivate.png"
                }else{
                    isprivateimg="isprivate_share.png"
                }
            }else{
                isprivateimg="space.png"
            }
            m.Id = id.String
            if strings.TrimSpace(attachment.String)!=""{
                if isprivate.Float64 != 0 {
                    if strings.Index(reader.String,",*,")==-1 {
                        m.Title = Tpl("<img src='/static/img/"+isprivateimg+"' alt='' style='width: 16px;height:16px'><img src='/static/img/att.png' style='width: 16px;height:16px'> "+configs.UnEscapeWords(title.String))
                    }else{
                        m.Title = Tpl("<img src='/static/img/"+isprivateimg+"' alt='' style='width: 16px;height:16px'><img src='/static/img/att.png' style='width: 16px;height:16px'> "+configs.UnEscapeWords(title.String))
                    }
                }else{
                    m.Title = Tpl("<img src='/static/img/att.png' style='width: 16px;height:16px'> "+configs.UnEscapeWords(title.String))
                }
            }else{
                if isprivate.Float64 != 0 {
                    m.Title = Tpl("<img src='/static/img/"+isprivateimg+"' alt='' style='width: 16px;height:16px'> "+configs.UnEscapeWords(title.String))
                }else{
                    m.Title = Tpl(configs.UnEscapeWords(title.String))
                }
            }
            m.Author = author.String
            m.Add_time = add_time.String
            m.Reply_count = reply_count.String
            m.Raise_count=raise_count.Float64
            m.Category = categoryen.String
            m.Edit_man = edit_man.String
            m.Read_count = read_count.String
            m.Link = link.String
            m.Isprivate = isprivate.Float64
            m.PrivateImg = "/static/img/"+isprivateimg+".png"
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
    if err != nil {
       configs.LogErr(err)
    }
	h["Content"]=string(b)
	c.HTML(http.StatusOK, "personalpublish/list", h)
}
// modify article
func EditBbsById(c *gin.Context) {
   h := DefaultH(c)
   h["PageTitle"] = "[[[$t('common.edit')]]]"
   if TreeBlogStr!=""{
       h["CateGoryStr"] = strings.TrimRight(configs.CateGoryStr, "]")+","+strings.TrimLeft(CateGoryBlogStr,"[")
   }else{
       h["CateGoryStr"] = configs.CateGoryStr
   }
   id:=c.Param("id")
    id = strings.TrimSpace(id)
    id = configs.EscapeWords(id)
    rows, err := configs.Db.Query("select a.id,a.title,a.author,a.add_time,a.reply_count,a.raise_count,a.categoryenall,a.content,a.myphoto,a.Edit_man,a.edit_history,a.read_count,a.mylink,a.mymedia,a.reader,a.isprivate,a.attachment,b.headicon,b.blog,b.score from bbs a LEFT JOIN myuser b on a.author=b.myusername where a.id=?",id)
    if err != nil {
        configs.LogErr(err)
    }
    defer rows.Close()
    IntSucess=0
    for rows.Next() {
        var id sql.NullString
        var title sql.NullString
        var author sql.NullString
        var add_time sql.NullString
        var reply_count sql.NullString
        var raise_count sql.NullString
        var categoryenall sql.NullString
        var content sql.NullString
        var myphoto sql.NullString
        var edit_man sql.NullString
        var edit_history sql.NullString
        var read_count sql.NullString
        var link sql.NullString
        var media sql.NullString
        var reader sql.NullString
        var isprivate sql.NullBool
        var attachment sql.NullString
        var userheadicon sql.NullString
        var userblog sql.NullBool
        var userscore sql.NullFloat64
        if err = rows.Scan(&id,&title, &author, &add_time,&reply_count,&raise_count,&categoryenall,&content,&myphoto,&edit_man,&edit_history,&read_count,&link,&media,&reader,&isprivate,&attachment,&userheadicon,&userblog,&userscore); err != nil {
            configs.LogErr(err)
        }
        if id.Valid {
            h["Id"] = id.String
            h["Title"] = configs.UnEscapeWords(title.String)
            h["Author"] = author.String
            h["Add_time"] = add_time.String
            h["Reply_count"] = reply_count.String
            h["Raise_count"] = raise_count.String
            h["Categoryenall"] = categoryenall.String
            h["Content"] = configs.UnEscapeWords(content.String)
            h["myphoto"] = myphoto.String            
            h["Edit_man"] = edit_man.String
            h["Edit_history"] = edit_history.String
            h["Read_count"] = read_count.String
            h["Link"] = link.String
            h["isprivate"] = isprivate.Bool
            h["ReaderStr"] = reader.String
            h["UserLevel"] = SessionUserLevel
            h["Media"] = media.String
            if strings.TrimSpace(attachment.String)!=""{
                attach := make([]*BbsAttachment, 0)
                attachmenttmp:=strings.Split(attachment.String,"?")
                myFileName:=""
                for i:=0;i<len(attachmenttmp);i++{
                    _, myFileName = filepath.Split(attachmenttmp[i])
                    att := &BbsAttachment{}
                    att.Name = myFileName
                    att.Url = attachmenttmp[i]
                    attach = append(attach, att)
                }
                b, err := json.Marshal(attach)
                if err != nil {
                   configs.LogErr(err)
                }
                h["FileList"]=string(b)
            }
            IntSucess=1
        }
        //ms = append(ms, m)
    }
    c.HTML(http.StatusOK, "bbs/personal/editarticle", h)
}

func UpdateArticlePost(c *gin.Context) {
    h := DefaultH(c)
    if configs.CheckUpdateCount(configs.GetFuncName(),30,c,Csrf_Result) {
        c.JSON(200, gin.H{"info": configs.Translate("tips.morethannrde")})
        return
    }
    h["WebTitle"] = Tpl("modify")
    if SessionUserId=="" {
        c.JSON(200, gin.H{"info": configs.Translate("common.plif")})
        return
    }

    id:=c.PostForm("id")
    var itemAuthor string
    err := configs.Db.QueryRow("select author from bbs where id=?",id).Scan(&itemAuthor)
    if err != nil {
        configs.LogErr(err)
        return
    }

    if itemAuthor != SessionUserId && SessionUserAdmin<100 {
        c.JSON(200, gin.H{"info": configs.Translate("common.nopernission")})
        return  
    }
    mytitle:=c.PostForm("mytitle")
    mytitle = strings.TrimSpace(mytitle)
    mytitle =strings.Replace(mytitle, "</script>","&lt;/script&gt;",-1)
    mytitle =strings.Replace(mytitle, "<script>","&lt;script&gt;",-1)
    categoryen:=c.PostForm("categoryen")
    categoryen = strings.TrimSpace(categoryen)
    categoryenall:=c.PostForm("categoryenall")
    categoryenall = strings.TrimSpace(categoryenall)
    categorycnall:=c.PostForm("categorycnall")
    categorycnall = strings.TrimSpace(categorycnall)
    content:=c.PostForm("content")
    content = strings.TrimSpace(content)
    mediaurl:=c.PostForm("mediaurltmp")
    mediaurl = strings.TrimSpace(mediaurl)
    postisprivate:=c.PostForm("isprivate")
    postisprivate = strings.TrimSpace(postisprivate)
    atturl:=c.PostForm("attachment")
    atturl = strings.TrimSpace(atturl)
    strshare:=c.PostForm("strshare")
    strshare = strings.TrimSpace(strshare)
    atturl = configs.EscapeWords(atturl)
    strshare = configs.EscapeWords(strshare)
    if strings.Index(strshare,","+SessionUserId+",")==-1 {
        strshare=strshare+","+SessionUserId+","
    }
    //logrus.Info("mytitle :", mytitle+"--"+categoryen+"--"+categoryenall+"--"+categorycnall+"--"+content)
    if (strings.Count(mytitle,"")-1<2 || strings.Count(categoryen,"")-1<2 ){
        c.JSON(200, gin.H{"info": configs.Translate("common.inputerr")})
        return
    }
    mytitle= CheckFilterResult(mytitle)
    mytitle = configs.EscapeWords(mytitle)
    categoryen = configs.EscapeWords(categoryen)
    categoryenall = configs.EscapeWords(categoryenall)
    categorycnall = configs.EscapeWords(categorycnall)
    content= CheckFilterResult(content)
    content = configs.EscapeWords(content)
    mediaurl = configs.EscapeWords(mediaurl)
    postisprivate = configs.EscapeWords(postisprivate)
    isprivate:=1

    if postisprivate=="false" {
        isprivate=0
    }

    timeUnix:=time.Now().Unix()   // timestamps
    formatTimeStr:=time.Unix(timeUnix,0).Format("2006-01-02 15:04:05")
    edit_history:=configs.EscapeWords("<font class='edit_history'>-----"+configs.Translate("The content was edited by ") + SessionUserId + configs.Translate(" at ")+  formatTimeStr +"-----</font><br>")
    result,err := configs.Db.Exec("update bbs set title=?,"+ configs.SqlSearchString2 +",edit_man=?,edit_time=?,categoryen=?,categoryenall=?,categorycnall=?,content=?,mymedia=?,reader=?,isprivate=?,attachment=? where id=?",mytitle,edit_history,SessionUserId,formatTimeStr,categoryen,categoryenall,categorycnall,content,mediaurl,strshare,isprivate,atturl,id)
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

//delete article
func DeleteArticlePost(c *gin.Context) {
    //logrus.Info("SessionUserId--:",SessionUserId)
    h := DefaultH(c)
    if configs.CheckUpdateCount(configs.GetFuncName(),30,c,Csrf_Result) {
        c.JSON(200, gin.H{"info": configs.Translate("tips.morethannrde")})
        return
    }
    h["WebTitle"] = "delete"
    id:=c.PostForm("id")
    if SessionUserId=="" {
        c.JSON(200, gin.H{"info": configs.Translate("common.plif")})
        return
    }

    var itemAuthor string
    err := configs.Db.QueryRow("select author from bbs where id=?",id).Scan(&itemAuthor)
    if err != nil {
        configs.LogErr(err)
        return
    }

    if itemAuthor != SessionUserId && SessionUserAdmin<100 {
        c.JSON(200, gin.H{"info": configs.Translate("common.nopernission")})
        return  
    }

    //result,err := configs.Db.Exec("delete from bbs where id=?",id)
    result,err := configs.Db.Exec("update bbs set allow=0 where id=?",id)
    if err != nil{
        configs.LogErr(err)
        c.JSON(200, gin.H{"info": configs.Translate("common.opfailed")})
        return
    }
    _,err = result.RowsAffected()
    if err != nil {
        configs.LogErr(err)
        c.JSON(200, gin.H{"info": configs.Translate("common.opfailed")})
        return
    }
    c.JSON(200, gin.H{"info": "ok"})
}