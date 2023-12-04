// Copyright 2019 golangbbs Core Team.  All rights reserved.
// LICENSE: Use of this source code is governed by AGPL-3.0.
// license that can be found in the LICENSE file.
package configs

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"html"
	"math"
	"math/rand"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
	//"reflect"
)

type FileNode struct {
	Name      string      `json:"name"`
	Path      string      `json:"path"`
	FileNodes []*FileNode `json:"children"`
}

// List all directories and files in the current directory
func WalkPath(path string, info os.FileInfo, node *FileNode) {
	files := ListFiles(path)
	for _, filename := range files {
		fpath := filepath.Join(path, filename)
		fio, _ := os.Lstat(fpath)
		child := FileNode{filename, fpath, []*FileNode{}}
		node.FileNodes = append(node.FileNodes, &child)
		if fio.IsDir() {
			WalkPath(fpath, fio, &child)
		}
	}
	return
}

func ListFiles(dirname string) []string {
	f, _ := os.Open(dirname)
	names, _ := f.Readdirnames(-1)
	f.Close()
	sort.Strings(names)
	return names
}
func FilterReplace(content string, FS map[int]FilterStruct) string {
	for _, value := range FS {
		content = strings.Replace(content, value.content, "*^*", -1)
	}
	return content
}
func LogErr(err error) {
	if err != nil {
		logrus.Error(err)
	}
}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func GetRandomString(m int, n int) string {
	var str string
	if m == 0 {
		str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	} else {
		str = "0123456789_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetCurrentIP(r http.Request) string {
	// Here you can also use the first value of the X-Forwarded-For request header as the user's IP
	// However, it should be noted that the IP represented by both request headers may be forged
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		// Get IP directly when the request header does not exist or there is no proxy
		ip = strings.Split(r.RemoteAddr, ":")[0]

	}
	return ip
}
func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dirctory string) string {
	return Substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		//log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func WriteCookie(cname string, cvalue string, c *gin.Context) {
	uid_cookie := &http.Cookie{
		Name:     cname,
		Value:    cvalue,
		Path:     "/",
		HttpOnly: false,
		MaxAge:   COOKIE_MAX_AGE,
	}

	http.SetCookie(c.Writer, uid_cookie)
}

func EscapeWords(cvalue string) string {
	result := html.EscapeString(cvalue)
	return result
}
func ChangeWord(cvalue string) string {
	result := strings.Replace(cvalue, "'", "&#39;", -1)
	result = strings.Replace(result, "\"", "&#34;", -1)
	return result
}
func UnEscapeWords(cvalue string) string {
	result := html.UnescapeString(cvalue)
	return result
}
func Md5Encrypt(data string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(data))
	cipherStr := md5Ctx.Sum(nil)
	encryptedData := hex.EncodeToString(cipherStr)
	return encryptedData
}
func CheckUpdateCount(sessname string, updatecount int, c *gin.Context, Csrf_Result bool) bool {
	session := sessions.Default(c)
	CheckUpdateCountBool := false
	sessname = sessname + "Count_" + Remote_IP
	sessCount := session.Get(sessname)
	if sessCount == nil {
		session.Set(sessname, 1)
	} else {
		s, _ := sessCount.(int)
		session.Set(sessname, (s + 1))
		if sessCount.(int) > updatecount {
			CheckUpdateCountBool = true
		}
	}
	session.Save()
	return CheckUpdateCountBool || Csrf_Result
}
func GetFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func GetCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

func SpliteCategory(s string, m string) string {
	strCategory := ""
	if strings.Index(s, m) != -1 {
		strtmp := Substr(s, strings.Index(s, m)+1, len(s))
		strCategory += s + "," + SpliteCategory(strtmp, m)

	} else {
		strCategory += s
	}
	return strCategory
}

func SpliteCategorys(s string, m string) string {
	strCategory := ""
	if strings.Index(s, m) != -1 {
		i := strings.LastIndex(s, m)
		strtmp := Substr(s, i+1, len(s))
		strCategory += strtmp + "," + strtmp + "_" + SpliteCategorys(Substr(s, 0, i), m)

	} else {
		strCategory = s + "," + strCategory
	}
	return strCategory
}

func TruncateNaive(f float64, unit float64) float64 {
	if BbsConfigs.DbType == "mysql" {
		return f
	} else {
		return math.Trunc(f/unit) * unit
	}
	//return math.Trunc(value*1e2+0.5) * 1e-2   //Add 0.5 for rounding. If you want to keep several decimal places, change 2
}

func CheckPortOpen(dominoname string, port int) bool {
	var isopen bool
	isopen = false
	/*
			ip, err := net.ResolveIPAddr("ip", dominoname)
		    if err != nil {
		        logrus.Info(dominoname+" can't ParseToIP: ", err)
		    }else{
			    tcpAddr := net.TCPAddr{
			        IP:   net.ParseIP(ip.String()),
			        Port: port,
			    }
			    conn, err := net.DialTCP("tcp", nil, &tcpAddr)
			    if err == nil {
			        isopen = true
			        conn.Close()
			    }
			}
	*/
	connTimeout := 2 * time.Second
	conn, err := net.DialTimeout("tcp", dominoname+":"+strconv.Itoa(port), connTimeout) // 1s timeout
	if err == nil {
		isopen = true
		conn.Close()
	}
	return isopen
}
