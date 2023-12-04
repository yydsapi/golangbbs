// Copyright 2019 golangbbs Core Team.  All rights reserved.
// LICENSE: Use of this source code is governed by AGPL-3.0.
// license that can be found in the LICENSE file.
package webcodev1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	//"encoding/json"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"golangbbs/configs"
	"path/filepath"
	"strings"
)

func BbsReply(c *gin.Context) {
	h := DefaultH(c)
	if configs.CheckUpdateCount(configs.GetFuncName(), 30, c, Csrf_Result) {
		c.JSON(200, gin.H{"info": configs.Translate("tips.morethannrde")})
		return
	}
	myid := c.PostForm("myid")
	myid = strings.TrimSpace(myid)
	mycontent := c.PostForm("mycontent")
	mycontent = strings.TrimSpace(mycontent)
	myid = configs.EscapeWords(myid)
	mycontent = CheckFilterResult(mycontent)
	mycontent = configs.EscapeWords(mycontent)
	myidint, _ := strconv.Atoi(myid)
	userid := "visitor"
	if SessionUserId != "" {
		userid = SessionUserId
	} else {
		c.JSON(200, gin.H{"info": configs.Translate("common.plif")})
		return
	}
	result, err := configs.Db.Exec("insert into bbsreply(pid,replycontent,replyuserid) values(?,?,?)", myidint, mycontent, userid)
	if err != nil {
		configs.LogErr(err)
		c.JSON(200, gin.H{"info": configs.Translate("common.ftudb")})
		return
	}
	_, err = result.RowsAffected()
	result, err = configs.Db.Exec("update bbs set reply_count=reply_count+1 where id=?", myidint)
	result, err = configs.Db.Exec("update myuser set score=score+1 where myusername=?", SessionUserId)
	c.JSON(200, gin.H{"info": "ok"})
	return
	h["Id"] = myid
}

func RaiseBbsById(c *gin.Context) {
	//session := sessions.Default(c)
	h := DefaultH(c)
	IntSucess = 0
	if SessionUserId == "" {
		c.JSON(200, gin.H{"info": configs.Translate("common.plif")})
		return
	}
	if configs.CheckUpdateCount(configs.GetFuncName(), 30, c, Csrf_Result) {
		c.JSON(200, gin.H{"info": configs.Translate("tips.morethannrde")})
		return
	}
	//logrus.Info("RaiseBbsById err ---:", session.Get("RaiseBbsById_"+configs.Remote_IP).(int))
	myid := c.PostForm("myid")
	myid = strings.TrimSpace(myid)
	myraise := c.PostForm("myraise")
	myraise = strings.TrimSpace(myraise)
	myid = configs.EscapeWords(myid)
	myraise = configs.EscapeWords(myraise)
	var itemCount float64
	var itemResult float64
	id, _ := strconv.Atoi(myid)
	err := configs.Db.QueryRow("select raise_count as count from bbs where id=?", id).Scan(&itemCount)
	if err != nil {
		configs.LogErr(err)
	}
	myraiseResult, _ := strconv.ParseFloat(myraise, 64)
	itemResult = (myraiseResult + itemCount) / 2
	itemTrans, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", itemResult), 64)
	_, _ = configs.Db.Exec("update bbs set raise_count=? where id=?", itemTrans, id)
	if SessionUserId != "" {
		_, _ = configs.Db.Exec("update myuser set score=score+0.05 where myusername=?", SessionUserId)
	}
	c.JSON(200, gin.H{"info": "ok"})
	return
	h["Id"] = myid
}

// Get a single article
func GetBbsById(c *gin.Context) {
	h := DefaultH(c)
	id := c.Param("id")
	id = strings.TrimSpace(id)
	id = configs.EscapeWords(id)
	searchkey := c.DefaultQuery("searchkey", "")
	searchkey = strings.TrimSpace(searchkey)
	searchkey = configs.EscapeWords(searchkey)
	category := c.DefaultQuery("category", "")
	category = strings.TrimSpace(category)
	category = configs.EscapeWords(category)
	//SELECT a.id, a.title,b.headicon FROM bbs a LEFT JOIN myuser b on a.author=b.myusername where a.id=1148
	rows, err := configs.Db.Query("select a.id,a.title,a.author,a.add_time,a.reply_count,a.raise_count,a.Categoryen,a.content,a.myphoto,a.Edit_man,a.edit_history,a.read_count,a.mylink,a.mymedia,a.reader,a.isprivate,a.attachment,b.headicon,b.blog,b.score from bbs a LEFT JOIN myuser b on a.author=b.myusername where a.allow>0 and a.id=?", id)
	if err != nil {
		configs.LogErr(err)
	}
	//"select id,title from bbs where id > "+ (int)Session["id"] +" and  lb='" + lben + "' order by id limit 0,1";
	defer rows.Close()
	IntSucess = 0
	mediatmp := ""
	fileExt := ""
	attachmentstr := ""
	for rows.Next() {
		var id sql.NullString
		var title sql.NullString
		var author sql.NullString
		var add_time sql.NullString
		var reply_count sql.NullString
		var raise_count sql.NullFloat64
		var categoryen sql.NullString
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
		if err = rows.Scan(&id, &title, &author, &add_time, &reply_count, &raise_count, &categoryen, &content, &myphoto, &edit_man, &edit_history, &read_count, &link, &media, &reader, &isprivate, &attachment, &userheadicon, &userblog, &userscore); err != nil {
			configs.LogErr(err)
		}
		if id.Valid {
			h["Id"] = id.String
			h["Title"] = Tpl(configs.UnEscapeWords(title.String))
			h["Author"] = author.String
			h["Add_time"] = add_time.String
			h["Reply_count"] = reply_count.String
			h["Raise_count"] = raise_count.Float64
			h["Category"] = categoryen.String
			h["Content"] = Tpl(configs.UnEscapeWords(content.String))
			h["myphoto"] = myphoto.String
			h["Edit_man"] = edit_man.String
			h["Edit_history"] = Tpl(configs.UnEscapeWords(edit_history.String))
			h["Read_count"] = read_count.String
			h["Link"] = link.String
			h["Media"] = ""
			if userheadicon.Valid {
				h["Headicon"] = userheadicon.String
			} else {
				h["Headicon"] = "/static/headicon/0.png"
			}
			h["UserBlog"] = userblog.Bool
			h["UserLever"] = strconv.Itoa(getUserLevel(userscore.Float64))
			if strings.TrimSpace(attachment.String) != "" {
				attachmenttmp := strings.Split(attachment.String, "?")
				myFileName := ""
				for i := 0; i < len(attachmenttmp); i++ {
					_, myFileName = filepath.Split(attachmenttmp[i])
					if i == len(attachmenttmp) {
						attachmentstr += "<span><a href='" + attachmenttmp[i] + "'>" + myFileName + "</a></span><br>"
					} else {
						attachmentstr += "<br><span><a href='" + attachmenttmp[i] + "'>" + myFileName + "</a></span>"
					}
				}
			}
			h["Attachment"] = Tpl(attachmentstr)
			if strings.Index(reader.String, ",*,") == -1 {
				if SessionUserId != "" {
					if author.String != SessionUserId && (strings.Index(reader.String, ","+SessionUserId+",") == -1 && isprivate.Bool) {
						h["info"] = configs.Translate("Content read permission error, no login?")
						h["else"] = "id:" + id.String
						c.HTML(http.StatusOK, "showerrinfo", h)
						return
					}
				} else {
					if isprivate.Bool {
						h["info"] = configs.Translate("Content read permission error, no login?")
						h["else"] = "id:" + id.String
						c.HTML(http.StatusOK, "showerrinfo", h)
						return
					}
				}
			}
			mediatmp = strings.TrimSpace(media.String)
			if mediatmp != "" {
				fileExt = filepath.Ext(mediatmp)
				h["Media"] = Tpl("<video src='" + mediatmp + "' width='400' height='250' controls preload></video>")
				if fileExt == ".mp3" || fileExt == ".ogg" || fileExt == ".wav" || fileExt == ".flac" || fileExt == ".ape" {
					h["Media"] = Tpl("<audio class='audio' controls preload autoplay='autoplay' loop='loop'><source src='" + mediatmp + "'></source></audio>")
				}
			}
			if SessionUserId == "" {
				h["ShowTxt"] = Tpl(configs.Translate("please") + " <a href=/signin?return=" + c.Request.URL.String() + ">" + configs.Translate("login") + "</a> " + configs.Translate("to reply") + ".<br><br>")
				h["ShowReply"] = "none"
			} else {
				h["ShowTxt"] = ""
				h["ShowReply"] = ""
			}
			IntSucess = 1
		}
		//ms = append(ms, m)
	}
	allreplycontent := ""
	strheadicon := ""
	replyrows, replyerr := configs.Db.Query("select a.id,a.replycontent,a.replytime,a.replyuserid,a.edit_history,b.headicon,b.blog,b.score from bbsreply a LEFT JOIN myuser b on a.replyuserid=b.myusername where a.allow>0 and a.pid=? order by a.replytime", id)
	if replyerr != nil {
		logrus.Info("GetBbsReply err:", err)
	}
	defer replyrows.Close()
	for replyrows.Next() {
		var id sql.NullString
		var replycontent sql.NullString
		var replytime sql.NullString
		var replyuserid sql.NullString
		var edithistory sql.NullString
		var userheadicon sql.NullString
		var userblog sql.NullBool
		var userscore sql.NullFloat64
		if err = replyrows.Scan(&id, &replycontent, &replytime, &replyuserid, &edithistory, &userheadicon, &userblog, &userscore); err != nil {
			configs.LogErr(err)
		}
		if id.Valid {
			if userheadicon.Valid {
				strheadicon = userheadicon.String
			} else {
				strheadicon = "/static/headicon/0.png"
			}
			if strings.TrimSpace(replycontent.String) == "" {
				allreplycontent += "<div class='clearfix' ><div class='bbscontent' style='float:left'>" + configs.UnEscapeWords(replycontent.String) + "</div></div>"
			} else {
				allreplycontent += "<div class='clearfix' ><div class='bbsreply' style='float:left'>" + configs.UnEscapeWords(replycontent.String) + "</div>"
				allreplycontent += "<div style='width:50%;float:left;'><span align='left'></span>" + configs.UnEscapeWords(edithistory.String) + "</div>"
				allreplycontent += "</div>"
			}
			allreplycontent += "<div class='clearfix' style='border-bottom:1px solid #EBEEF5;'><div style='float:right;margin-right:25px;line-height:30px;'><img src='" + strheadicon + "' width=26px height=26px /> " + replyuserid.String + "<img src='/static/img/level/" + strconv.Itoa(getUserLevel(userscore.Float64)) + ".png' width=20px height=20px /> <font color='#B9B9B9'>" + replytime.String + "</font></div></div>"
		}
		//ms = append(ms, m)
	}
	h["ReplayContent"] = Tpl(allreplycontent)
	var err1 error
	var err2 error
	var prvid string
	var nextid string
	var prvtitle string
	var nexttitle string
	if SessionUserId == "" {
		if searchkey == "" {
			if category == "" {
				err1 = configs.Db.QueryRow("select id,title from bbs where id=(select max(id) from bbs where allow>0 and isprivate=0 and id < ? order by reply_time desc,raise_count desc, id desc)", id).Scan(&nextid, &nexttitle)
				err2 = configs.Db.QueryRow("select id,title from bbs where id=(select min(id) from bbs where allow>0 and isprivate=0 and id > ?  order by reply_time desc,raise_count desc, id desc)", id).Scan(&prvid, &prvtitle)
			} else {
				err1 = configs.Db.QueryRow("select id,title from bbs where id=(select max(id) from bbs where allow>0 and isprivate=0 and categoryen=? and id < ? order by reply_time desc,raise_count desc, id desc)", category, id).Scan(&nextid, &nexttitle)
				err2 = configs.Db.QueryRow("select id,title from bbs where id=(select min(id) from bbs where allow>0 and isprivate=0 and categoryen=? and id > ? order by reply_time desc,raise_count desc, id desc)", category, id).Scan(&prvid, &prvtitle)
			}
		} else {
			err1 = configs.Db.QueryRow("select id,title from bbs where id=(select max(id) from bbs where allow>0 and isprivate=0 and (title like "+configs.SqlSearchString1+" or content like "+configs.SqlSearchString1+") and id < ? order by reply_time desc,raise_count desc, id desc)", searchkey, searchkey, id).Scan(&nextid, &nexttitle)
			err2 = configs.Db.QueryRow("select id,title from bbs where id=(select min(id) from bbs where allow>0 and isprivate=0 and (title like "+configs.SqlSearchString1+" or content like "+configs.SqlSearchString1+") and id > ? order by reply_time desc,raise_count desc, id desc)", searchkey, searchkey, id).Scan(&prvid, &prvtitle)
		}
	} else {
		if searchkey == "" {
			if category == "" {
				err1 = configs.Db.QueryRow("select id,title from bbs where id=(select max(id) from bbs where (isprivate=0 or (isprivate=1 and reader like "+configs.SqlSearchString1+")) and allow > 0 and id < ? order by reply_time desc,raise_count desc, id desc)", ","+SessionUserId+",", id).Scan(&nextid, &nexttitle)
				err2 = configs.Db.QueryRow("select id,title from bbs where id=(select min(id) from bbs where (isprivate=0 or (isprivate=1 and reader like "+configs.SqlSearchString1+")) and allow > 0 and id > ? order by reply_time desc,raise_count desc, id desc)", ","+SessionUserId+",", id).Scan(&prvid, &prvtitle)
			} else {
				err1 = configs.Db.QueryRow("select id,title from bbs where id=(select max(id) from bbs where (isprivate=0 or (isprivate=1 and reader like "+configs.SqlSearchString1+")) and allow > 0 and categoryen=? and id < ? order by reply_time desc,raise_count desc, id desc)", ","+SessionUserId+",", category, id).Scan(&nextid, &nexttitle)
				err2 = configs.Db.QueryRow("select id,title from bbs where id=(select min(id) from bbs where (isprivate=0 or (isprivate=1 and reader like "+configs.SqlSearchString1+")) and allow > 0 and categoryen=? and id > ? order by reply_time desc,raise_count desc, id desc)", ","+SessionUserId+",", category, id).Scan(&prvid, &prvtitle)
			}
		} else {
			err1 = configs.Db.QueryRow("select id,title from bbs where id=(select max(id) from bbs where (isprivate=0 or (isprivate=1 and reader like "+configs.SqlSearchString1+")) and allow > 0 and (title like "+configs.SqlSearchString1+" or content like "+configs.SqlSearchString1+") and id < ? order by reply_time desc,raise_count desc, id desc)", ","+SessionUserId+",", searchkey, searchkey, id).Scan(&nextid, &nexttitle)
			err2 = configs.Db.QueryRow("select id,title from bbs where id=(select min(id) from bbs where (isprivate=0 or (isprivate=1 and reader like "+configs.SqlSearchString1+")) and allow > 0 and (title like "+configs.SqlSearchString1+" or content like "+configs.SqlSearchString1+") and id > ? order by reply_time desc,raise_count desc, id desc)", ","+SessionUserId+",", searchkey, searchkey, id).Scan(&prvid, &prvtitle)
		}
	}
	if err2 != nil {
		configs.LogErr(err)
		prvid = ""
		prvtitle = ""
	}
	if err1 != nil {
		configs.LogErr(err)
		nextid = ""
		nexttitle = ""
	}

	var prvstr string
	var nextstr string
	prvstr = ""
	nextstr = ""
	urlpath := c.Request.URL.String()
	if prvid == "" {
		prvstr = `<font style="margin-right:13px;color:#828282">` + configs.Translate("prev (non-exist)") + `</a>`
	} else {
		prvstr = `<a href="` + strings.Replace(urlpath, id, prvid, -1) + `" style="margin-right:13px;color:#828282">` + configs.Translate("prev") + ` (` + string([]rune(prvtitle)[0:7]) + `..)</a>`
	}

	if nextid == "" {
		nextstr = `<font style="margin-right:20px;color:#828282">` + configs.Translate("next (non-exist)") + `</a>`
	} else {
		nextstr = `<a href="` + strings.Replace(urlpath, id, nextid, -1) + `" style="margin-right:20px;color:#828282">` + configs.Translate("next") + ` (` + string([]rune(nexttitle)[0:7]) + `..)</a>`
	}

	h["PrvPage"] = Tpl(prvstr)
	h["NextPage"] = Tpl(nextstr)
	//logrus.Info(c.Request.URL.Path)    //logrus.Info(c.Request.URL.RawQuery)    //logrus.Info(*(c.Request.URL.String()))
	if IntSucess == 0 {
		h["info"] = configs.Translate("Content read error")
		h["else"] = "id:" + id
		c.HTML(http.StatusOK, "showerrinfo", h)

	} else {
		c.HTML(http.StatusOK, "bbs/show", h)
	}
}

func AddArticle(c *gin.Context) {
	h := DefaultH(c)
	h["PageTitle"] = "[[[$t('common.publish')]]]"
	category := c.DefaultQuery("category", "")
	category = strings.TrimSpace(category)
	Categoryenall := "[]"
	if category != "" {
		Categoryenall = ""
		str := configs.SpliteCategory(category, "_")
		arr := strings.Split(str, ",")
		for i := len(arr) - 1; i >= 0; i-- {
			if i == len(arr)-1 {
				Categoryenall = arr[i]
			} else {
				Categoryenall = Categoryenall + "," + arr[i]
			}
		}
	}
	logrus.Info(Categoryenall)
	var score float64
	err := configs.Db.QueryRow("select score from myuser where  myusername=?", SessionUserId).Scan(&score)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"info": configs.Translate("common.serverError")})
		return
	} else {
		h["UserLevel"] = SessionUserLevel
	}
	if TreeBlogStr != "" {
		h["CateGoryStr"] = strings.TrimRight(configs.CateGoryStr, "]") + "," + strings.TrimLeft(CateGoryBlogStr, "[")
	} else {
		h["CateGoryStr"] = configs.CateGoryStr
	}
	h["Categoryenall"] = Categoryenall
	c.HTML(http.StatusOK, "bbs/addarticle", h)
}

func AddArticlePost(c *gin.Context) {
	h := DefaultH(c)
	if configs.CheckUpdateCount(configs.GetFuncName(), 30, c, Csrf_Result) {
		c.JSON(200, gin.H{"info": configs.Translate("tips.morethannrde")})
		return
	}
	h["WebTitle"] = Tpl("publish")
	if SessionUserId == "" {
		c.JSON(200, gin.H{"info": configs.Translate("common.plif")})
		return
	}
	mytitle := c.PostForm("mytitle")
	mytitle = strings.TrimSpace(mytitle)
	mytitle = strings.Replace(mytitle, "</script>", "&lt;/script&gt;", -1)
	mytitle = strings.Replace(mytitle, "<script>", "&lt;script&gt;", -1)
	categoryen := c.PostForm("categoryen")
	categoryen = strings.TrimSpace(categoryen)
	categoryenall := c.PostForm("categoryenall")
	categoryenall = strings.TrimSpace(categoryenall)
	categorycnall := c.PostForm("categorycnall")
	categorycnall = strings.TrimSpace(categorycnall)
	content := c.PostForm("content")
	content = strings.TrimSpace(content)
	mediaurl := c.PostForm("mediaurltmp")
	mediaurl = strings.TrimSpace(mediaurl)
	atturl := c.PostForm("attachment")
	atturl = strings.TrimSpace(atturl)
	strshare := c.PostForm("strshare")
	strshare = strings.TrimSpace(strshare)
	postisprivate := c.PostForm("isprivate")
	postisprivate = strings.TrimSpace(postisprivate)
	author := SessionUserId
	//logrus.Info("mytitle :", mytitle+"--"+categoryen+"--"+categoryenall+"--"+categorycnall+"--"+content)
	if strings.Count(mytitle, "")-1 < 2 || strings.Count(categoryen, "")-1 < 2 {
		c.JSON(200, gin.H{"info": configs.Translate("common.inputerr")})
		return
	}
	mytitle = CheckFilterResult(mytitle)
	mytitle = configs.EscapeWords(mytitle)
	categoryen = configs.EscapeWords(categoryen)
	categoryenall = configs.EscapeWords(categoryenall)
	categorycnall = configs.EscapeWords(categorycnall)
	content = CheckFilterResult(content)
	content = configs.EscapeWords(content)
	mediaurl = configs.EscapeWords(mediaurl)
	atturl = configs.EscapeWords(atturl)
	strshare = configs.EscapeWords(strshare)
	postisprivate = configs.EscapeWords(postisprivate)
	if strings.Index(strshare, ","+SessionUserId+",") == -1 {
		strshare = strshare + "," + SessionUserId + ","
	}
	isprivate := 1
	if postisprivate == "false" {
		isprivate = 0
	}
	var itemCount int
	counterr := configs.Db.QueryRow("select count(*) as count  from bbs where title=? and author=?", mytitle, author).Scan(&itemCount)
	if counterr != nil {
		configs.LogErr(counterr)
	}
	if itemCount > 0 {
		c.JSON(200, gin.H{"info": configs.Translate("publish.titlerep")})
		return
	}

	result, err := configs.Db.Exec("insert into bbs(title,author,categoryen,categoryenall,categorycnall,content,mymedia,reader,isprivate,attachment) values(?,?,?,?,?,?,?,?,?,?)", mytitle, author, categoryen, categoryenall, categorycnall, content, mediaurl, strshare, isprivate, atturl)
	if err != nil {
		configs.LogErr(err)
		c.JSON(200, gin.H{"info": configs.Translate("common.opfailed")})
		return
	}
	_, err = result.RowsAffected()
	if err != nil {
		configs.LogErr(err)
		c.JSON(200, gin.H{"info": configs.Translate("common.opfailed")})
		return
	}
	if SessionUserId != "" {
		_, err = configs.Db.Exec("update myuser set score=score+5 where myusername=?", SessionUserId)
		if err != nil {
			configs.LogErr(err)
			c.JSON(200, gin.H{"info": configs.Translate("common.opfailed")})
			return
		}
	}
	returnURL := c.DefaultQuery("return", "/")
	c.JSON(200, gin.H{"info": "ok", "returnURL": returnURL})
	return
}
