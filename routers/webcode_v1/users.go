// Copyright 2019 golangbbs Core Team.  All rights reserved.
// LICENSE: Use of this source code is governed by AGPL-3.0.
// license that can be found in the LICENSE file.
package webcodev1

import (
	"database/sql"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"golangbbs/configs"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"
	"unicode"
)

func SignInGet(c *gin.Context) {
	h := DefaultH(c)
	h["PageTitle"] = "[[[$t('common.signin')]]]"
	c.HTML(http.StatusOK, "users/signin", h)
}

func SignInPost(c *gin.Context) {
	session := sessions.Default(c)
	if configs.CheckUpdateCount(configs.GetFuncName(), 30, c, Csrf_Result) {
		c.JSON(200, gin.H{"info": configs.Translate("tips.morethannrde")})
		return
	}
	myusername := c.PostForm("myregusername")
	myusername = strings.TrimSpace(myusername)
	mypassword := c.PostForm("myregpassword")
	mypassword = strings.TrimSpace(mypassword)

	var intUserLevel int
	if strings.Count(myusername, "")-1 < 5 || strings.Count(mypassword, "")-1 < 5 {
		c.JSON(200, gin.H{"info": configs.Translate("common.inputerr")})
		return
	}
	myusername = strings.ToLower(configs.EscapeWords(myusername))
	mypassword = configs.EscapeWords(mypassword)
	var intAllow int64
	intAllow = -1
	var strUseBlog string
	rows, err := configs.Db.Query("select allow,headicon,blog,score FROM myuser where myusername='" + myusername + "' and mypwd='" + configs.Md5Encrypt(mypassword) + "'")
	if err != nil {
		configs.LogErr(err)
		c.JSON(200, gin.H{"info": configs.Translate("common.ftudb")})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var allow sql.NullInt64
		var headicon sql.NullString
		var blog sql.NullBool
		var userscore sql.NullFloat64
		if err = rows.Scan(&allow, &headicon, &blog, &userscore); err != nil {
			configs.LogErr(err)
			c.JSON(200, gin.H{"info": configs.Translate("common.ftudb")})
			return
		}
		if allow.Valid {
			intAllow = allow.Int64
			intUserLevel = getUserLevel(userscore.Float64)
			if blog.Bool {
				strUseBlog = "1"
			} else {
				strUseBlog = "0"
			}
		} else {
			c.JSON(200, gin.H{"info": configs.Translate("signin.lfuperr")})
		}
	}
	returnURL := c.DefaultQuery("return", "")
	if returnURL == "" {
		returnURL = "/"
	}
	intAllow16, _ := strconv.Atoi(strconv.FormatInt(intAllow, 10))
	switch {
	case intAllow16 == -1:
		c.JSON(200, gin.H{"info": configs.Translate("signin.lfuperr")})
	case intAllow16 == 0:
		c.JSON(200, gin.H{"info": configs.Translate("signin.lfud")})
	case intAllow16 > 0 && intAllow16 < 10:
		session.Set(userIDKey, myusername)
		session.Set(userAdminKey, intAllow16)
		session.Set(userLevelKey, intUserLevel)
		session.Set("userBlog", strUseBlog)
		SessionUserId = myusername
		if strUseBlog == "1" {
			InitBlogmenu("---")
			session.Set("NavBlogMenuStr", NavBlogMenuStr)
		}
		session.Save()
		c.JSON(200, gin.H{"info": "ok", "returnURL": returnURL})
	case intAllow16 > 9 && intAllow < 100:
		session.Set(userIDKey, myusername)
		session.Set(userAdminKey, intAllow16)
		session.Set(userLevelKey, intUserLevel)
		session.Set("userBlog", strUseBlog)
		SessionUserId = myusername
		if strUseBlog == "1" {
			InitBlogmenu("---")
			session.Set("NavBlogMenuStr", NavBlogMenuStr)
		}
		session.Save()
		c.JSON(200, gin.H{"info": "ok", "returnURL": returnURL})
	case intAllow16 > 99:
		session.Set(userIDKey, myusername)
		session.Set(userAdminKey, intAllow16)
		session.Set(userLevelKey, intUserLevel)
		session.Set("userBlog", strUseBlog)
		SessionUserId = myusername
		if strUseBlog == "1" {
			InitBlogmenu("---")
			session.Set("NavBlogMenuStr", NavBlogMenuStr)
		}
		session.Save()
		c.JSON(200, gin.H{"info": "ok", "returnURL": returnURL})
	default:
		c.JSON(200, gin.H{"info": configs.Translate("common.opfailed")})
	}
	//logrus.Info("SignInPost:============>", SessionUserId)
	h := DefaultH(c)
	h["WebTitle"] = Tpl("register")
	return
}
func FindPasswordGet(c *gin.Context) {
	h := DefaultH(c)
	h["Title"] = "FindPassword"
	c.HTML(http.StatusOK, "users/findpassword", h)
}
func FindPasswordPost(c *gin.Context) {
	h := DefaultH(c)
	if configs.CheckUpdateCount(configs.GetFuncName(), 30, c, Csrf_Result) {
		c.JSON(200, gin.H{"info": configs.Translate("tips.morethannrde")})
		return
	}
	h["WebTitle"] = Tpl("find password")
	myusername := c.PostForm("myregusername")
	myusername = strings.TrimSpace(myusername)
	myemail := c.PostForm("myemail")
	myemail = strings.TrimSpace(myemail)
	if strings.Count(myusername, "")-1 < 5 || strings.Count(myemail, "")-1 < 4 {
		c.JSON(200, gin.H{"info": configs.Translate("common.inputerr")})
		return
	}
	myusername = strings.ToLower(configs.EscapeWords(myusername))
	myemail = strings.ToLower(configs.EscapeWords(myemail))
	var itemCount int
	err := configs.Db.QueryRow("select count(*) as count  FROM myuser where myusername='" + myusername + "' and email='" + myemail + "'").Scan(&itemCount)
	if err != nil {
		configs.LogErr(err)
		c.JSON(200, gin.H{"info": configs.Translate("common.ftudb")})
		return
	}
	var mypwd string
	if itemCount == 1 {
		for i := 0; i < 10000; i++ {
			pwd := configs.GetRandomString(1, 10)
			f := configs.Substr(pwd, 0, 1)
			st := []rune(f)
			if unicode.IsLetter(st[0]) {
				mypwd = pwd
				break
			}
		}
	} else {
		c.JSON(200, gin.H{"info": "signin.fineerrumn"})
		return
	}
	_, err = configs.Db.Exec("update myuser set mypwd=? where myusername=? and email=?", configs.Md5Encrypt(mypwd), myusername, myemail)
	if err != nil {
		configs.LogErr(err)
		c.JSON(200, gin.H{"info": configs.Translate("common.opfailed")})
		return
	}
	auth := smtp.PlainAuth("", configs.BbsConfigs.SendMailConfigs.SMTPUser, configs.BbsConfigs.SendMailConfigs.SMTPPassword, configs.BbsConfigs.SendMailConfigs.MailServer)
	to := []string{myemail}
	nickname := configs.BbsConfigs.SendMailConfigs.NickName
	user := configs.BbsConfigs.SendMailConfigs.SMTPUser
	subject := configs.Translate("find password")
	content_type := "Content-Type: text/plain; charset=UTF-8"
	body := configs.Translate("Here is the new temporary password") + ":\n" + mypwd + "\n" + configs.Translate("please login as soon as possible and modify your personal data.") + "\n" + configs.Translate("Click here to log in") + " golangbbs: http://" + configs.BbsConfigs.DominoName + "/signin"
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err = smtp.SendMail(configs.BbsConfigs.SendMailConfigs.MailServer+":25", auth, user, to, msg)
	if err != nil {
		configs.LogErr(err)
		c.JSON(200, gin.H{"info": configs.Translate("common.opfailed")})
		return
	}
	c.JSON(200, gin.H{"info": "ok"})
	return
}

// SignUpGet handles
func SignUpGet(c *gin.Context) {
	h := DefaultH(c)
	h["PageTitle"] = "[[[$t('common.signup')]]]"
	c.HTML(http.StatusOK, "users/signupget", h)
}

func SignUpPost(c *gin.Context) {
	session := sessions.Default(c)
	h := DefaultH(c)
	if configs.CheckUpdateCount(configs.GetFuncName(), 30, c, Csrf_Result) {
		c.JSON(200, gin.H{"info": configs.Translate("tips.morethannrde")})
		return
	}
	h["WebTitle"] = "register"
	myusername := c.PostForm("myusername")
	myemail := c.PostForm("myemail")
	mysex := c.PostForm("mysex")
	mypassword := c.PostForm("mypassword")
	myicon := c.PostForm("myicon")
	myusername = strings.ToLower(strings.TrimSpace(myusername))
	myemail = strings.ToLower(strings.TrimSpace(myemail))
	mysex = strings.TrimSpace(mysex)
	mypassword = strings.TrimSpace(mypassword)
	myicon = strings.TrimSpace(myicon)
	myusername = CheckFilterResult(myusername)
	myusername = configs.EscapeWords(myusername)
	if strings.Count(myusername, "")-1 < 4 || strings.Count(myemail, "")-1 < 4 || strings.Count(mypassword, "")-1 < 4 {
		c.JSON(200, gin.H{"info": configs.Translate("common.inputerr")})
		return
	}
	myusername = strings.ToLower(configs.EscapeWords(myusername))
	myemail = strings.ToLower(configs.EscapeWords(myemail))
	mysex = configs.EscapeWords(mysex)
	mypassword = configs.EscapeWords(mypassword)
	myicon = configs.EscapeWords(myicon)
	if myicon == "" {
		myicon = "/static/headicon/1.png"
	}
	var itemCount int
	err := configs.Db.QueryRow("select count(*) as count  FROM myuser where myusername='" + myusername + "' or email='" + myemail + "'").Scan(&itemCount)
	if err != nil {
		configs.LogErr(err)
		c.JSON(200, gin.H{"info": configs.Translate("common.ftudb")})
		return
	}

	if itemCount == 0 {
		_, err := configs.Db.Exec("insert into myuser(myusername,mypwd,sex,email,headicon) values (?,?,?,?,?)", myusername, configs.Md5Encrypt(mypassword), mysex, myemail, myicon)
		if err != nil {
			configs.LogErr(err)
			c.JSON(200, gin.H{"info": configs.Translate("common.opfailed")})
			return
		}
		session.Set(userIDKey, myusername)
		session.Set(userAdminKey, 1)
		session.Save()
		returnURL := c.DefaultQuery("return", "/")
		c.JSON(200, gin.H{"info": "ok", "returnURL": returnURL})
		return
	} else {
		c.JSON(200, gin.H{"info": configs.Translate("signin.uerep")})
		return
	}
}

func LogOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(userIDKey)
	session.Delete("userAdminKey")
	session.Delete("userLevelKey")
	session.Delete("userBlog")
	session.Delete("NavBlogMenuStr")
	configs.UserIdCookie = ""
	session.Save()
	c.JSON(200, gin.H{"info": "ok"})
	return
}
func UsersInfo(c *gin.Context) {
	h := DefaultH(c)
	rows, err := configs.Db.Query("select id,sex,email,headicon,score from myuser where allow>0 and myusername=?", SessionUserId)
	if err != nil {
		configs.LogErr(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id sql.NullString
		var sex sql.NullString
		var email sql.NullString
		var headicon sql.NullString
		var userscore sql.NullFloat64
		if err = rows.Scan(&id, &sex, &email, &headicon, &userscore); err != nil {
			configs.LogErr(err)
		}
		if id.Valid {
			h["Id"] = id.String
			h["Sex"] = sex.String
			h["Email"] = email.String
			h["Headicon"] = headicon.String
			h["UserLevel"] = "/static/img/level/" + strconv.Itoa(SessionUserLevel) + ".png"
			h["UsersScore"] = configs.TruncateNaive(userscore.Float64, 0.01)
		}
		//ms = append(ms, m)
	}
	h["Userid"] = SessionUserId
	h["info"] = "ok"
	c.HTML(http.StatusOK, "users/usersinfo", h)
}
func PersonalInfoPost(c *gin.Context) {
	h := DefaultH(c)
	if configs.CheckUpdateCount(configs.GetFuncName(), 30, c, Csrf_Result) {
		c.JSON(200, gin.H{"info": configs.Translate("tips.morethannrde")})
		return
	}
	h["WebTitle"] = "modify"
	mysex := c.PostForm("mysex")
	mysex = strings.TrimSpace(mysex)
	mypassword := c.PostForm("mypassword")
	mypassword = strings.TrimSpace(mypassword)

	myicon := c.PostForm("myicon")
	myicon = strings.TrimSpace(myicon)
	mysex = configs.EscapeWords(mysex)
	mypassword = configs.EscapeWords(mypassword)
	myicon = configs.EscapeWords(myicon)
	if strings.Count(myicon, "")-1 < 4 {
		c.JSON(200, gin.H{"info": configs.Translate("common.inputerr")})
		return
	}
	if mypassword != "" && strings.Count(mypassword, "")-1 < 4 {
		c.JSON(200, gin.H{"info": configs.Translate("common.inputerr")})
		return
	}
	var itemCount int
	err := configs.Db.QueryRow("select count(*) as count  FROM myuser where myusername='" + SessionUserId + "' and allow>0").Scan(&itemCount)
	if err != nil {
		configs.LogErr(err)
		c.JSON(200, gin.H{"info": configs.Translate("common.ftudb")})
		return
	}
	if itemCount == 1 {
		var err error
		if mypassword == "" {
			_, err = configs.Db.Exec("update myuser set sex=?,headicon=? where myusername=?", mysex, myicon, SessionUserId)
		} else {
			_, err = configs.Db.Exec("update myuser set sex=?,mypwd=?,headicon=? where myusername=?", mysex, configs.Md5Encrypt(mypassword), myicon, SessionUserId)
		}
		if err != nil {
			configs.LogErr(err)
			c.JSON(200, gin.H{"info": configs.Translate("common.opfailed")})
			return
		}
		c.JSON(200, gin.H{"info": "ok"})
		return
	} else {
		c.JSON(200, gin.H{"info": configs.Translate("common.ftudb")})
		return
	}
}
func PersonalEditBlogTree(c *gin.Context) {
	var itemBlog bool
	err := configs.Db.QueryRow("select blog from myuser where myusername=?", SessionUserId).Scan(&itemBlog)
	if err != nil {
		logrus.Info("PersonalEditBlogTree select err:", err)
		return
	}
	session := sessions.Default(c)
	if itemBlog == false {
		session.Set("userBlog", "0")
		session.Save()
	} else {
		session.Set("userBlog", "1")
		session.Save()
	}
	InitBlogmenu("---")
	session.Set("NavBlogMenuStr", NavBlogMenuStr)
	session.Save()
	h := DefaultH(c)
	h["SessionUseBlog"] = itemBlog
	h["Content"] = TreeBlogStr
	h["Title"] = "Welcome to BlogMenu edit"
	c.HTML(http.StatusOK, "users/editblogtree", h)
}

func PersonalMenuTreePost(c *gin.Context) {
	h := DefaultH(c)
	if configs.CheckUpdateCount(configs.GetFuncName(), 30, c, Csrf_Result) {
		c.JSON(200, gin.H{"info": configs.Translate("tips.morethannrde")})
		return
	}
	h["WebTitle"] = Tpl("")
	pid := c.PostForm("pid")
	pid = strings.TrimSpace(pid)
	cid := c.PostForm("cid")
	cid = strings.TrimSpace(cid)
	categoryname := c.PostForm("categoryname")
	categoryname = strings.TrimSpace(categoryname)
	weight := c.PostForm("weight")
	weight = strings.TrimSpace(weight)
	pid = configs.EscapeWords(pid)
	cid = configs.EscapeWords(cid)
	categoryname = CheckFilterResult(categoryname)
	categoryname = configs.EscapeWords(categoryname)
	weight = configs.EscapeWords(weight)
	dbid := c.PostForm("dbid")
	dbid = strings.TrimSpace(dbid)
	dbid = configs.EscapeWords(dbid)
	dbname := ""
	switch dbid {
	case "menu":
		dbname = "category"
	case "blog":
		dbname = "category_blog"
	default:
		dbname = ""
	}
	if dbname == "" {
		c.JSON(200, gin.H{"info": configs.Translate("common.opfailed")})
		return
	}
	//var result sql.Result
	var err error
	if cid == "" {
		en := configs.GetRandomString(0, 8) + "_" + pid
		logrus.Info(pid)
		_, err = configs.Db.Exec("insert into "+dbname+"(en,cn,pid,member,sort_weight) values (?,?,?,?,?)", en, categoryname, pid, SessionUserId, weight)
	} else {
		_, err = configs.Db.Exec("update "+dbname+" set cn=?,sort_weight=? where id=?", categoryname, weight, cid)
	}
	if err != nil {
		c.JSON(200, gin.H{"info": configs.Translate("common.opfailed")})
		return
	}
	c.JSON(200, gin.H{"info": "ok"})
	return
}

func PersonalDeleteTree(c *gin.Context) {
	//logrus.Info("SessionUserId--:",SessionUserId)
	//h := DefaultH(c)
	if configs.CheckUpdateCount(configs.GetFuncName(), 30, c, Csrf_Result) {
		c.JSON(200, gin.H{"info": configs.Translate("tips.morethannrde")})
		return
	}
	//session := sessions.Default(c)
	id := c.PostForm("id")
	dbid := c.PostForm("dbid")
	dbid = strings.TrimSpace(dbid)
	dbid = configs.EscapeWords(dbid)
	dbname := ""
	switch dbid {
	case "menu":
		dbname = "category"
	case "blog":
		dbname = "category_blog"
	default:
		dbname = ""
	}
	if dbname == "" {
		c.JSON(200, gin.H{"info": configs.Translate("common.opfailed")})
		return
	}
	//logrus.Info("SessionUserId:",SessionUserId)
	if SessionUserId == "" {
		c.JSON(200, gin.H{"info": configs.Translate("common.plif")})
		return
	}

	itemAuthor := ""
	if dbname == "category_blog" {
		useriderr := configs.Db.QueryRow("select member from category_blog where id=?", id).Scan(&itemAuthor)
		if useriderr != nil {
			logrus.Info("DeleteArticlePost select err:", useriderr)
			return
		}
		if itemAuthor != SessionUserId && SessionUserAdmin < 100 {
			c.JSON(200, gin.H{"info": configs.Translate("common.nopernission")})
			return
		}
	} else {
		if SessionUserAdmin < 100 {
			c.JSON(200, gin.H{"info": configs.Translate("common.nopernission")})
			return
		}
	}
	_, err := configs.Db.Exec("update "+dbname+" set allow=0 where id=?", id)
	if err != nil {
		configs.LogErr(err)
		c.JSON(200, gin.H{"info": configs.Translate("common.opfailed")})
		return
	}
	_, err = configs.Db.Exec("update bbs set allow=0 where bbs.categoryen=(select en from "+dbname+" WHERE id=?)", id)
	if err != nil {
		configs.LogErr(err)
		c.JSON(200, gin.H{"info": configs.Translate("common.opfailed")})
		return
	}
	c.JSON(200, gin.H{"info": "ok"})
}

func CommonUpdateById(c *gin.Context) {
	h := DefaultH(c)
	if configs.CheckUpdateCount(configs.GetFuncName(), 30, c, Csrf_Result) {
		c.JSON(200, gin.H{"info": configs.Translate("tips.morethannrde")})
		return
	}
	h["WebTitle"] = "update"
	id := c.PostForm("id")
	id = strings.TrimSpace(id)
	id = configs.EscapeWords(id)
	sort := c.PostForm("sort")
	sort = strings.TrimSpace(sort)
	sort = configs.EscapeWords(sort)
	if SessionUserId == "" {
		c.JSON(200, gin.H{"info": configs.Translate("common.plif")})
		return
	}
	switch sort {
	case "updateuserblog":
		_, err := configs.Db.Exec("update myuser set blog=? where myusername=?", id, SessionUserId)
		if err != nil {
			configs.LogErr(err)
			c.JSON(200, gin.H{"info": configs.Translate("common.opfailed")})
			return
		}
		c.JSON(200, gin.H{"info": "ok"})
		return

	default:
		c.JSON(200, gin.H{"info": configs.Translate("common.unknownerr")})
		return
	}
}
