// Copyright 2019 golangbbs Core Team.  All rights reserved.
// LICENSE: Use of this source code is governed by AGPL-3.0.
// license that can be found in the LICENSE file.
package webcodev1

import (
	"database/sql"
	"encoding/json"
	"golangbbs/configs"
	"html/template"
	"strconv"
	"strings"
	//"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

const InvalidWords = " ,~,!,@,#,　,$,%,^,&,*,(,),_,-,+,=,?,<,>,.,—,，,。,/,\\,|,《,》,？,;,:,：,',‘,；,“,"
const userIDKey = "UserID"
const userLevelKey = "UserLevel"
const userAdminKey = "UserAdmin"
const MenuQuerySQL = "select id,en,cn,pid,mymemo,sort_weight,img,mylink from category_blog where pid =? and allow=1 and member =? order by sort_weight"
const MenuBlogQuerySQL2 = "select id,en,cn,pid,mymemo,sort_weight,img,mylink from category_blog where pid =? and allow=1 and member =? order by sort_weight"
const MenuBlogQuerySQL1 = "select id,en,cn,pid,mymemo,sort_weight,img,mylink from category_blog where pid =? and allow=1  order by sort_weight"

var (
	IntSucess        int
	SessionUserId    string
	SessionUserAdmin int
	SessionUserLevel int
	SessionUseBlog   string
	NavBlogMenuStr   string
	CateGoryBlogStr  string
	TreeBlogStr      string
	IsMobile         bool
	SessionCsrf      string
	Csrf_Result      bool
	HasSensitiveWord bool
	SensitiveWord    map[string]interface{}
	SetFilterWord    map[string]interface{}
	InvalidWord      map[string]interface{}
	MyUser           map[string]interface{}
)

type Bbs struct {
	Id           string        `json:"id"`
	Title        template.HTML `json:"title"`
	Author       string        `json:"author"`
	Add_time     string        `json:"add_time"`
	Reply_count  string        `json:"reply_count"`
	Raise_count  float64       `json:"raise_count"`
	Category     string        `json:"category"`
	Category_cn  string        `json:"category_cn"`
	Content      template.HTML `json:"content"`
	Myphoto      string        `json:"myphoto"`
	Member       string        `json:"member"`
	Edit_time    string        `json:"edit_time"`
	Edit_man     string        `json:"edit_man"`
	Edit_history string        `json:"edit_history"`
	Read_count   string        `json:"read_count"`
	Link         string        `json:"link"`
	Media        string        `json:"media"`
	Img          string        `json:"img"`
	ImgRaise     string        `json:"imgraise"`
	Isprivate    float64       `json:"isprivate"`
	Attachment   string        `json:"attachment"`
	PrivateImg   string        `json:"privateimg"`
	UserHeadIcon string        `json:"userheadicon"`
	UserBlog     string        `json:"userblog"`
	UserLevel    string        `json:"userlevel"`
}
type BbsAttachment struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type Bbsreply struct {
	Id           string        `json:"id"`
	Pid          string        `json:"pid"`
	Replycontent template.HTML `json:"replycontent"`
	Replytime    string        `json:"replytime"`
	Headicon     string        `json:"Headicon"`
	Link         string        `json:"link"`
	Img          string        `json:"img"`
}
type Menu struct {
	En          string  `json:"en"`
	Name        string  `json:"name"`
	ParentId    string  `json:"pid"`
	Id          string  `json:"id"`
	Sort_weight string  `json:"sort_weight"`
	Img         string  `json:"img"`
	Link        string  `json:"link"`
	Memo        string  `json:"memo"`
	Value       string  `json:"value"`
	Label       string  `json:"label"`
	Child       []*Menu `json:"children,omitempty"`
}

func getUserLevel(s float64) (l int) {
	l = 1
	switch {
	case s > 299 && s < 900:
		l = 2
		//fallthrough
	case s > 899 && s < 1800:
		l = 3
	case s > 1799 && s < 3000:
		l = 4
	case s > 2999 && s < 5000:
		l = 5
	case s > 4999 && s < 7000:
		l = 6
	case s > 6999 && s < 10000:
		l = 7
	case s > 9999 && s < 40000:
		l = 8
	case s > 39999:
		l = 9
	default:
		l = 1
	}
	return l
}
func GetMyUserInfo(s string) {
	s = strings.TrimSpace(s)
	MyUser = make(map[string]interface{})
	var headicon string
	var blog bool
	var score int
	if s == "" {
		MyUser["headicon"] = ""
		MyUser["blog"] = false //bool
		MyUser["score"] = 0    //int
	}
	return
	err := configs.Db.QueryRow("select headicon,blog,score from myuser where  myusername=?", s).Scan(&headicon, &blog, &score)
	if err != nil {
		configs.LogErr(err)
		MyUser["headicon"] = ""
		MyUser["blog"] = false //bool
		MyUser["score"] = 0    //int
	}
	MyUser["headicon"] = headicon
	MyUser["blog"] = blog   //bool
	MyUser["score"] = score //int
}
func Tpl(s string) template.HTML {
	return template.HTML(s)
}
func CheckFilterResult(content string) (result string) {
	SensitiveWord = make(map[string]interface{})
	SetFilterWord = make(map[string]interface{})
	InvalidWord = make(map[string]interface{}) //Invalid words, not involved in sensitive word judgment, directly ignored
	HasSensitiveWord = false
	rows, err := configs.Db.Query("select title from filter")
	if err != nil {
		logrus.Info("err:", err)
	}
	defer rows.Close()
	for rows.Next() {
		var title sql.NullString
		if err = rows.Scan(&title); err != nil {
			configs.LogErr(err)
		}
		if title.Valid {
			SetFilterWord[title.String] = nil
		}
	}
	words := strings.Split(InvalidWords, ",")
	for _, v := range words {
		InvalidWord[v] = nil
	}
	AddSensitiveToMap(SetFilterWord)
	strtmp := ChangeSensitiveWords(strings.ToLower(content), SensitiveWord)
	if !HasSensitiveWord || SessionUserId == "limon" {
		strtmp = content
	}
	return strtmp
}
func ShowImg(xx string, yy string, zz float64, a string) string {
	changeShowImg := "/static/img/as.png"
	x, _ := strconv.Atoi(xx)
	y, _ := strconv.Atoi(yy)
	if zz < 5 {
		if x > 15 {
			changeShowImg = "/static/img/as1.png"
		}
		if y > 200 {
			changeShowImg = "/static/img/as2.png"
		}
	} else {
		changeShowImg = "/static/img/ark1.png"
		if zz == 5 {
			changeShowImg = "/static/img/up.png"
		}
	}
	return changeShowImg
}
func ShowRaiseImg(xx string, yy string, zz float64, a string) string {
	changeShowImg := "/static/img/space.png"
	x, _ := strconv.Atoi(xx)
	y, _ := strconv.Atoi(yy)
	if x > 14 {
		changeShowImg = "/static/img/level/hot.png"
	}
	if y > 499 {
		changeShowImg = "/static/img/level/hot.png"
	}
	if zz > 999 {
		changeShowImg = "/static/img/level/jin.png"
	}
	return changeShowImg
}

// DefaultH returns to all pages
func DefaultH(c *gin.Context) gin.H {
	IsMobile = false
	keywords := []string{"Android", "iPhone", "iPod", "iPad", "Windows Phone", "MQQBrowser"}
	for i := 0; i < len(keywords); i++ {
		if strings.Contains(c.Request.UserAgent(), keywords[i]) {
			IsMobile = true
			break
		}
	}
	configs.Lang = "en"
	CkeditLanguage := "en"
	HomePage := "<a href='http://www.golangbbs.com' target=_blank>golangbbs.com</a>"
	Language := strings.TrimSpace(strings.ToLower(c.Request.Header.Get("Accept-Language")))
	if strings.Index(Language, ",") != -1 {
		Language = configs.Substr(Language, 0, strings.Index(Language, ","))
	}

	session := sessions.Default(c)
	//check update
	RightImgSrc := "new" + configs.Lang + ".png"
	sess_imgright := session.Get("imgright")
	if sess_imgright == nil {
		if configs.CheckPortOpen(configs.BbsConfigs.DominoName, 7999) {
			RightImgSrc = "new.png"
		} else {
			RightImgSrc = "new" + configs.Lang + ".png"
		}
		session.Set("imgright", RightImgSrc)
		session.Save()
	} else {
		RightImgSrc = strings.TrimSpace(sess_imgright.(string))
	}

	Csrf_Result = false
	sess_Csrf := session.Get("SessCsrf")
	if sess_Csrf != nil {
		SessionCsrf = strings.TrimSpace(sess_Csrf.(string))
	} else {
		tmp := configs.GetRandomString(1, 30)
		session.Set("SessCsrf", tmp)
		session.Save()
		SessionCsrf = tmp
	}
	if c.Request.Method == "POST" {
		Csrftmp := strings.TrimSpace(c.PostForm("Csrf"))
		//logrus.Info("Csrftmp--------->",Csrftmp)
		if Csrftmp != SessionCsrf {
			Csrf_Result = true
		}
	}
	SessionUserId = ""
	sessuid := session.Get(userIDKey)
	if sessuid != nil {
		SessionUserId = strings.TrimSpace(sessuid.(string))
		SessionUserId = configs.EscapeWords(SessionUserId)
	}
	SessionUserLevel = 0
	sesslev := session.Get(userLevelKey)
	if sesslev != nil {
		SessionUserLevel = sesslev.(int)
	}
	SessionUserAdmin = 0
	sessadmin := session.Get(userAdminKey)
	if sessadmin != nil {
		SessionUserAdmin = sessadmin.(int)
	}
	SessionUseBlog = "0"
	sessuserblog := session.Get("userBlog")
	if sessuserblog != nil {
		SessionUseBlog = strings.TrimSpace(sessuserblog.(string))
		SessionUseBlog = configs.EscapeWords(SessionUseBlog)
	}
	var NavMenuStr string
	NavMenuStr = `<el-menu-item index="555" onclick="location.href='/'"><img src="/static/img/point.png" width=16px height=16px style="margin-top:-2px" border=0> [[[$t('common.home')]]]</el-menu-item>`
	if SessionUserId == "" {
		NavMenuStr += configs.NavMenuStr + `<el-menu-item index="555" onclick="location.href='/signin'">[[[$t('common.signin')]]]</el-menu-item><el-menu-item index="556" onclick="location.href='/signup'">[[[$t('common.signup')]]]</el-menu-item>`
	} else {
		if SessionUseBlog == "1" {
			//logrus.Info("SessionUseBlog -----------------"+SessionUseBlog)
			sessNavBlogMenuStr := session.Get("NavBlogMenuStr")
			if sessNavBlogMenuStr != nil {
				NavBlogMenuStr = strings.TrimSpace(sessNavBlogMenuStr.(string))
			} else {
				InitBlogmenu("---")
				session.Set("NavBlogMenuStr", NavBlogMenuStr)
				session.Save()
			}
			NavMenuStr += configs.NavMenuStr
			NavMenuStr += NavBlogMenuStr
		} else {
			NavMenuStr += configs.NavMenuStr
		}
		NavMenuStr += `<el-submenu index="6660"><template slot="title">[[[$t('common.my')]]](` + SessionUserId + `)</template>`
		NavMenuStr += `<el-menu-item index="6661" onclick="location.href='/bbs/addarticle'">[[[$t('common.publish')]]]</el-menu-item>`
		NavMenuStr += `<el-menu-item index="6662" onclick="location.href='/bbs/personalpublish'">[[[$t('common.myPublish')]]]</el-menu-item>`
		NavMenuStr += `<el-menu-item index="6663" onclick="location.href='/bbs/personalreply'">[[[$t('common.myReply')]]]</el-menu-item>`
		NavMenuStr += `<el-menu-item index="6664" onclick="location.href='/bbs/personaleditblogtree'">[[[$t('common.myBlogConfig')]]]</el-menu-item>`
		NavMenuStr += `<el-menu-item index="6664" onclick="location.href='/bbs/usersinfo'">[[[$t('common.myProfile')]]]</el-menu-item>`
		if SessionUserAdmin > 99 {
			NavMenuStr += `<el-menu-item index="6665" onclick="location.href='/bbs/personaleditmenutree'">[[[$t('common.myMenuConfig')]]]</el-menu-item><el-menu-item index="6666" onclick="location.href='/bbs/usersmanage'">[[[$t('manage.usersmanage')]]]</el-menu-item>`
		}
		NavMenuStr += `</el-submenu><el-menu-item index="667" onclick="confirmLogout()" >[[[$t('common.logout')]]]</el-menu-item>`
	}
	//logrus.Info("NavMenuStr------->>--------->",NavMenuStr)
	configs.Remote_IP = configs.GetCurrentIP(*c.Request)
	//logrus.Info("?Session==routerip="+configs.Remote_IP+"==DefaultH====="+SessionUserId+"?")
	return gin.H{
		"IsMobile":       IsMobile,
		"Title":          "", //page title
		"Context":        c,
		"SessionUserId":  SessionUserId,
		"WebTitle":       "golangbbs",
		"LanguageStr":    Language,
		"IntAdmin":       SessionUserAdmin,
		"NavMenuStr":     template.HTML(NavMenuStr),
		"RightImgSrc":    RightImgSrc,
		"HomePage":       template.HTML(HomePage),
		"CkeditLanguage": CkeditLanguage,
		"Csrf":           SessionCsrf,
	}
}

func MenuCheckSub(m *Menu) bool {
	var itemCount int
	var err error
	if m.En == "---" || m.En == "root" {
		err = configs.Db.QueryRow("select count(*) as count FROM category_blog where allow=1 and pid='" + m.En + "'").Scan(&itemCount)
	} else {
		err = configs.Db.QueryRow("select count(*) as count FROM category_blog where allow=1 and pid='" + m.En + "' and member='" + SessionUserId + "'").Scan(&itemCount)
	}
	if err != nil {
	}
	if itemCount > 0 {
		return true
	} else {
		return false
	}
}
func MenuAppend(m *Menu) {
	child, err := MenuInit(MenuQuerySQL, m.En)
	if err != nil {
		configs.LogErr(err)
		return
	}
	m.Child = child
	for _, v := range m.Child {
		b := MenuCheckSub(v)
		if b {
			NavBlogMenuStr += "<el-submenu index=" + v.Id + "><template slot=\"title\">" + v.Name + "</template>"
			MenuAppendSub(v)

		} else {
			NavBlogMenuStr += "<el-menu-item index=" + v.Id + ">" + v.Name + "</el-menu-item>"
			MenuAppend(v)
		}
	}
}

func MenuAppendSub(m *Menu) {
	child, err := MenuInit(MenuQuerySQL, m.En)
	if err != nil {
		configs.LogErr(err)
		return
	}
	m.Child = child
	for _, v := range m.Child {
		b := MenuCheckSub(v)
		if b {
			NavBlogMenuStr += "<el-submenu index=" + v.Id + "><template slot=\"title\">" + v.Name + "</template>"
			MenuAppendSub(v)
		} else {
			if strings.Index(v.Link, "/") == -1 {
				NavBlogMenuStr += `<el-menu-item index=` + v.Id + ` onclick="location.href='/?page=` + v.Link + `&category=` + v.En + `'">` + v.Name + `</el-menu-item>`
			} else {
				NavBlogMenuStr += `<el-menu-item index=` + v.Id + ` onclick="location.href='` + v.Link + `'">` + v.Name + `</el-menu-item>`
			}
		}

	}
	NavBlogMenuStr += "</el-submenu>"
}

func LoadMenu(strRoot string) (*Menu, error) {
	ms, err := MenuInit(MenuQuerySQL, strRoot)
	if err != nil {
		configs.LogErr(err)
		return nil, err
	}
	root := ms[0]
	MenuAppend(root)
	return root, nil
}
func MenuInit(sql0 string, pid string) ([]*Menu, error) {
	var rows *sql.Rows
	var err error
	if pid == "---" || pid == "root" {
		rows, err = configs.Db.Query(MenuBlogQuerySQL1, pid)
	} else {
		rows, err = configs.Db.Query(MenuBlogQuerySQL2, pid, SessionUserId)
	}
	if err != nil {
		configs.LogErr(err)
		return nil, err
	}
	defer rows.Close()
	ms := make([]*Menu, 0)
	for rows.Next() {
		m := &Menu{}
		var id sql.NullString
		var en sql.NullString
		var cn sql.NullString
		var pid sql.NullString
		var memo sql.NullString
		var sort_weight sql.NullString
		var img sql.NullString
		var link sql.NullString
		if err = rows.Scan(&id, &en, &cn, &pid, &memo, &sort_weight, &img, &link); err != nil {
			configs.LogErr(err)
			return nil, err
		}
		if en.Valid {
			m.Id = id.String
			m.En = en.String
			m.Name = cn.String
			m.ParentId = pid.String
			m.Memo = memo.String
			m.Sort_weight = sort_weight.String
			m.Img = img.String
			m.Link = link.String
			m.Value = en.String
			m.Label = cn.String
		}
		ms = append(ms, m)
	}
	//logrus.Info("current data:", ms)
	return ms, nil
}

func InitBlogmenu(strRoot string) {
	//logrus.Info("InitBlogmenu:============>", SessionUserId)
	NavBlogMenuStr = ""
	CateGoryBlogStr = ""
	TreeBlogStr = ""
	root, err := LoadMenu(strRoot)
	if err != nil {
		//fmt.Println("err:", err)
	}
	/*
	   b, err := json.Marshal(root)
	   if err != nil {
	       configs.LogErr(err)
	   }
	*/
	c, err := json.Marshal(root.Child)
	if err != nil {
		//fmt.Println("err:", err)
	}

	CateGoryBlogStr = string(c)
	TreeBlogStr = "[" + string(c) + "]"
}
