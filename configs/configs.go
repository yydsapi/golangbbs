// Copyright 2019 golangbbs Core Team.  All rights reserved.
// LICENSE: Use of this source code is governed by AGPL-3.0.
// license that can be found in the LICENSE file.
package configs

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var IntConfigPath = flag.Int("m", 0, "IntConfigPath")

// read json
type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

type FilterStruct struct {
	content string
}

func (jst *JsonStruct) Load(filename string, v interface{}) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

type DbConfig struct {
	MySqlConnStr  string
	MySqlLimit    int
	Sqlite3DbPath string
}

type SendMailConfig struct {
	SMTPUser     string
	SMTPPassword string
	MailServer   string
	NickName     string
}
type BbsLimit struct {
	UploadMedia int64
	UploadPhoto int64
	UploadElse  int64
}
type BbsConfig struct {
	RunMode         string
	IPAddr          string
	HTTPPort        string
	PageSize        int
	ReadTimeout     int
	WriteTimeout    int
	COOKIE_MAX_AGE  int
	StaticDir       string
	BbsUploadPath   string
	DominoName      string
	Precompile      int
	DbType          string
	DbConfigs       DbConfig
	SendMailConfigs SendMailConfig
	BbsLimits       BbsLimit
}

var (
	RunMode          string
	Lang             string
	IpAddr           string
	HTTPPort         string
	DbType           string
	SqlSearchString1 string
	SqlSearchString2 string
	SqlSearchString3 string
	ReadTimeout      time.Duration
	WriteTimeout     time.Duration
	PageSize         int
	SessionSecret    string
	COOKIE_MAX_AGE   int
	NavMenuStr       string
	UserIdCookie     string
	TreeStr          string
	CateGoryStr      string
	StaticDir        string
	Remote_IP        string
	TreeMenu         *Menu
	BbsConfigs       *BbsConfig
	FilterArray      map[int]FilterStruct
)

func init() {
	SessionSecret = `json:"session_secret"`
	JsonParse := NewJsonStruct()
	flag.Parse()
	switch *IntConfigPath {
	case 0:
		JsonParse.Load("bbs_config_main_i18n_80.json", &BbsConfigs) //    //BbsConfigs:= BbsConfig{}
	case 8001:
		JsonParse.Load("/root/bbs_config_main_i18n_8001.json", &BbsConfigs) //    //BbsConfigs:= BbsConfig{}
	case 8008:
		JsonParse.Load("/root/bbs_config_main_i18n_8008.json", &BbsConfigs) //    //BbsConfigs:= BbsConfig{}
	default:
		JsonParse.Load("bbs_config_main_i18n_80.json", &BbsConfigs) //    //BbsConfigs:= BbsConfig{}
	}
	IpAddr = BbsConfigs.IPAddr
	HTTPPort = BbsConfigs.HTTPPort
	DbType = BbsConfigs.DbType
	RunMode = BbsConfigs.RunMode //#debug or release
	PageSize = BbsConfigs.PageSize
	ReadTimeout = time.Duration(BbsConfigs.ReadTimeout) * time.Second
	WriteTimeout = time.Duration(BbsConfigs.WriteTimeout) * time.Second
	COOKIE_MAX_AGE = BbsConfigs.COOKIE_MAX_AGE * 24 * 3600 // Unit: sec
	StaticDir = BbsConfigs.StaticDir
	//When precompile is 1, it is in precompiled state. SQLite dB and static resources and attachments used are in the current directory
	if BbsConfigs.Precompile == 1 {
		dir, _ := os.Executable()
		exPath := filepath.Dir(dir)
		BbsConfigs.BbsUploadPath = exPath + "/documents"
	}
	FilterArray = make(map[int]FilterStruct)
	Initdb()
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

const (
	MenuQuerySQL = "select id,en,cn,pid,mymemo,sort_weight,img,mylink from category where pid =? and allow=1 order by sort_weight"
)

func GetBreadCrumb(s string) string {
	strBreadCrumb := ""
	switch s {
	case "AddArticle":
		strBreadCrumb = ""
	default:
		strBreadCrumb = ""
	}
	strBreadCrumb = ""
	return strBreadCrumb
}
func MenuCheckSub(m *Menu) bool {
	var itemCount int
	err := Db.QueryRow("select count(*) as count FROM category where pid='" + m.En + "'").Scan(&itemCount)
	if err != nil {
		LogErr(err)
	}
	if itemCount > 0 {
		return true
	} else {
		return false
	}
}

// read menu
func MenuAppend(m *Menu) {
	child, err := MenuInit(MenuQuerySQL, m.En)
	if err != nil {
		LogErr(err)
		return
	}
	m.Child = child
	for _, v := range m.Child {
		b := MenuCheckSub(v)
		if b {
			NavMenuStr += "<el-submenu index=" + v.Id + "><template slot=\"title\">" + v.Name + "</template>"
			MenuAppendSub(v)

		} else {
			NavMenuStr += "<el-menu-item index=" + v.Id + ">" + v.Name + "</el-menu-item>"
			MenuAppend(v)
		}
	}
}

func MenuAppendSub(m *Menu) {
	child, err := MenuInit(MenuQuerySQL, m.En)
	if err != nil {
		LogErr(err)
		return
	}
	m.Child = child
	for _, v := range m.Child {
		b := MenuCheckSub(v)
		if b {
			NavMenuStr += "<el-submenu index=" + v.Id + "><template slot=\"title\">" + v.Name + "</template>"
			MenuAppendSub(v)
		} else {
			if strings.Index(v.Link, "/") == -1 {
				NavMenuStr += `<el-menu-item index=` + v.Id + ` onclick="location.href='/?page=` + v.Link + `&category=` + v.En + `'">` + v.Name + `</el-menu-item>`
			} else {
				NavMenuStr += `<el-menu-item index=` + v.Id + ` onclick="location.href='` + v.Link + `'">` + v.Name + `</el-menu-item>`
			}
		}
	}
	NavMenuStr += "</el-submenu>"
}

func LoadMenu() (*Menu, error) {
	ms, err := MenuInit(MenuQuerySQL, "---")
	if err != nil {
		LogErr(err)
		return nil, err
	}
	root := ms[0]
	MenuAppend(root)
	return root, nil
}
func Tx(vSql string, args ...any) (*sql.Rows, error) {
	stmt, err := Db.Prepare(vSql)
	if err != nil {
		fmt.Println("Tx err:", err)
		return nil, err
	}
	defer stmt.Close()
	return stmt.Query(args...)
}
func MenuInit(sql0 string, pid string) ([]*Menu, error) {
	rows, err := Tx(sql0, pid)
	//rows, err := Db.Query(sql0, pid)
	if err != nil {
		LogErr(err)
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
			logrus.Info("MenuInit error:", err)
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

func Initmenu() {
	//logrus.Info("current ------------------>", "Initmenu")
	root, err := LoadMenu()
	if err != nil {
		LogErr(err)
	}
	b, err := json.Marshal(root)
	if err != nil {
		LogErr(err)
	}
	c, err := json.Marshal(root.Child)
	if err != nil {
		LogErr(err)
	}
	CateGoryStr = string(c)
	TreeStr = "[" + string(b) + "]"
}
