// Copyright 2019 golangbbs Core Team.  All rights reserved.
// LICENSE: Use of this source code is governed by AGPL-3.0.
// license that can be found in the LICENSE file.
package main

import (
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/memstore"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "golangbbs/configs"
    "golangbbs/routers/webcode_v1"
    "net/http"
    "os"
    "path/filepath"
)

func main() {
	router := gin.Default()
	//When precompile is 1, it is in precompiled state. SQLite dB and static resources and attachments used are in the current directory
	if configs.BbsConfigs.Precompile == 1 {
		dir, _ := os.Executable()
		exPath := filepath.Dir(dir)
		router.StaticFS("/static", configs.StatikFS)
		router.StaticFS("/documents", http.Dir(exPath+"/documents"))
		configs.LoadStatikTemplates()
	} else {
		router.StaticFS("/static", http.Dir(configs.BbsConfigs.StaticDir)) //better use nginx to serve assets etc
		router.StaticFS("/documents", http.Dir(configs.BbsConfigs.BbsUploadPath))
		configs.LoadTemplates()
	}
	router.SetHTMLTemplate(configs.GetTemplates())
	configs.Initmenu()
	store := memstore.NewStore([]byte(configs.SessionSecret))
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 7 * 86400}) //Also set Secure: true if using SSL
	router.Use(sessions.Sessions("gin-session", store))
	//Periodic tasks
	//gocron.Every(1).Day().Do(something)
	//gocron.Start()
	router.GET("/", webcodev1.IndexGet)
	router.GET("/about", webcodev1.AboutGet)
	router.GET("/favicon.ico", webcodev1.FaviconGet)

	router.GET("/signup", webcodev1.SignUpGet)
	router.POST("/signuppost", webcodev1.SignUpPost)
	router.GET("/signin", webcodev1.SignInGet)
	router.POST("/signinpost", webcodev1.SignInPost)
	router.GET("/logout", webcodev1.LogOut)
	authorized := router.Group("/bbs")
	authorized.GET("/showitem/:id", webcodev1.GetBbsById)
	authorized.POST("/raiseitem", webcodev1.RaiseBbsById)
	authorized.POST("/bbsreply", webcodev1.BbsReply)
	authorized.GET("/photos", webcodev1.PhotoGet)
	authorized.GET("/videos", webcodev1.VideoGet)
	authorized.GET("/musics", webcodev1.MusicGet)
	authorized.GET("/findpassword", webcodev1.FindPasswordGet)
	authorized.POST("/findpasswordpost", webcodev1.FindPasswordPost)
	authorized.Use(webcodev1.AuthRequired())
	{
		authorized.GET("/addarticle", webcodev1.AddArticle)
		authorized.POST("/upload", webcodev1.UploadPost)
		authorized.POST("/uploadatt", webcodev1.UploadPostAtt)
		authorized.POST("/deleteatt", webcodev1.DeleteAtt)
		authorized.POST("/addarticlepost", webcodev1.AddArticlePost)
		authorized.POST("/updatearticlepost", webcodev1.UpdateArticlePost)
		authorized.POST("/updatereplypost", webcodev1.UpdateReplyPost)
		authorized.POST("/deletearticlepost", webcodev1.DeleteArticlePost)
		authorized.POST("/deletereplypost", webcodev1.DeleteReplyPost)
		authorized.GET("/personalpublish", webcodev1.PersonalPublishGet)
		authorized.GET("/edititem/:id", webcodev1.EditBbsById)
		authorized.GET("/editreplyitem/:id", webcodev1.EditReplyById)
		authorized.GET("/personalreply", webcodev1.PersonalReplyGet)
		authorized.GET("/usersinfo", webcodev1.UsersInfo)
		authorized.POST("/personalinfopost", webcodev1.PersonalInfoPost)
		authorized.GET("/personaleditblogtree", webcodev1.PersonalEditBlogTree)
		authorized.POST("/personalmenutreepost", webcodev1.PersonalMenuTreePost)
		authorized.POST("/commonupdatebyid", webcodev1.CommonUpdateById)

		authorized.GET("/personaleditmenutree", webcodev1.ManageEditMenuTree)
		authorized.POST("/personaldeletemenutree", webcodev1.PersonalDeleteTree)
		authorized.GET("/usersmanage", webcodev1.UsersManage)
		authorized.POST("/usersmanagepost", webcodev1.UsersManagePost)
	}

	s := &http.Server{
		Addr:           configs.IpAddr + ":" + configs.HTTPPort,
		Handler:        router,
		ReadTimeout:    configs.ReadTimeout,
		WriteTimeout:   configs.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	logrus.Info("listen port start： " + configs.HTTPPort)
	logrus.Info(s.ListenAndServe())
	logrus.Info("listen port end： " + configs.HTTPPort)

}

//initLogger initializes logrus logger with some defaults
func initLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	if gin.Mode() == gin.DebugMode {
		logrus.SetLevel(logrus.DebugLevel)
	}
}
