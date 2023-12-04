// Copyright 2019 golangbbs Core Team.  All rights reserved.
// LICENSE: Use of this source code is governed by AGPL-3.0.
// license that can be found in the LICENSE file.
package webcodev1

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
    "golangbbs/configs"
	"github.com/gin-gonic/gin"
	"strings"
	"strconv"
)

func UploadPost(c *gin.Context) {
	h := DefaultH(c)
	h["Title"] = "upload"
	if SessionUserLevel<5 {
		c.JSON(200, gin.H{"uploaded": false, "info": configs.Translate("Your current level LV ")+strconv.Itoa(SessionUserLevel)+", "+configs.Translate("LV5 or above to upload pictures and attachments")})
	    return
	}
    Csrf_Result=false
    if configs.CheckUpdateCount(configs.GetFuncName(),30,c,Csrf_Result) {
        c.JSON(200, gin.H{"info": configs.Translate("tips.morethannrde")})
        return
    }
	mpartFile, mpartHeader, err := c.Request.FormFile("upload")
	if err != nil {
		mpartFile, mpartHeader, err = c.Request.FormFile("file")
		if err != nil {
			configs.LogErr(err)
			c.JSON(200, gin.H{"uploaded": false, "info": configs.Translate("common.opfailed")})
			return
		}
	}
	defer mpartFile.Close()
	fileExt := filepath.Ext(mpartHeader.Filename)
	fileExt=strings.TrimSpace(strings.ToLower(fileExt))
	if fileExt==".mp4" || fileExt==".mp3" || fileExt==".ogg" || fileExt==".wav" || fileExt==".flac" || fileExt==".ape" || fileExt==".mkv" || fileExt==".avi" || fileExt==".rmvb" {
		err = c.Request.ParseMultipartForm(configs.BbsConfigs.BbsLimits.UploadMedia << 20)
	}else if fileExt==".jpg" || fileExt==".png" || fileExt==".gif" || fileExt==".ico" {
		err = c.Request.ParseMultipartForm(configs.BbsConfigs.BbsLimits.UploadPhoto << 20)
	}else{
		err = c.Request.ParseMultipartForm(configs.BbsConfigs.BbsLimits.UploadElse << 20)
	}
	if err != nil {
		configs.LogErr(err)
		c.JSON(200, gin.H{"uploaded": false, "info":"Media limit "+strconv.FormatInt(configs.BbsConfigs.BbsLimits.UploadMedia,10)+"M,Photo limit "+strconv.FormatInt(configs.BbsConfigs.BbsLimits.UploadPhoto,10)+",Else limit "+strconv.FormatInt(configs.BbsConfigs.BbsLimits.UploadPhoto,10)+"M"})
		return
	}
	uri, err := saveFile(mpartHeader, mpartFile)
	if err != nil {
		configs.LogErr(err)
		c.JSON(200, gin.H{"uploaded": false, "info": configs.Translate("common.opfailed")})
		return
	}

	c.JSON(200, gin.H{"uploaded": true, "url": uri})
}
//saves file to disc
func saveFile(fh *multipart.FileHeader, f multipart.File) (string, error) {
	fileExt := filepath.Ext(fh.Filename)
	fileExt=strings.TrimSpace(strings.ToLower(fileExt))
	newName := fmt.Sprint(time.Now().Unix()) + fileExt //unique file name ;D
	//	newName := fmt.Sprint(time.Now().UnixNano()) + fileExt //unique file name ;D
	//time.Sleep(time.Duration(300)*time.Millisecond)
	uri := ""
	fullPath := ""
	fullName := ""
	if fileExt==".mp4" || fileExt==".mp3" || fileExt==".ogg" || fileExt==".wav" || fileExt==".flac" || fileExt==".ape" || fileExt==".mkv" || fileExt==".avi" || fileExt==".rmvb" {
		uri = "/documents/videos/" + SessionUserId +"/" +newName
		fullPath=configs.BbsConfigs.BbsUploadPath+"/videos/"+SessionUserId
		
	}else if fileExt==".jpg" || fileExt==".bmp" || fileExt==".gif" || fileExt==".png"{
		uri = "/documents/images/" + SessionUserId +"/" +newName
		fullPath = configs.BbsConfigs.BbsUploadPath+"/images/"+SessionUserId
	}else{
		uri = "/documents/attachments/" + SessionUserId +"/" +newName
		fullPath = configs.BbsConfigs.BbsUploadPath+"/attachments/"+SessionUserId
	}
	exist, err := configs.PathExists(fullPath)
    if err != nil {
        configs.LogErr(err)
    }
    if !exist {
    	err := os.Mkdir(fullPath, 0777)
        if err != nil {
            configs.LogErr(err)
        }
	}
	fullName = filepath.Join(fullPath, newName)
	file, err := os.OpenFile(fullName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err = io.Copy(file, f)
	if err != nil {
		return "", err
	}
	return uri, nil
}

func UploadPostAtt(c *gin.Context) {
	h := DefaultH(c)
	h["Title"] = "upload"

	if SessionUserLevel<5 {
		c.JSON(200, gin.H{"uploaded": false, "info": configs.Translate("Your current level LV ")+strconv.Itoa(SessionUserLevel)+", "+configs.Translate("LV5 or above to upload pictures and attachments")})
	    return
	}
    //logrus.Info(Csrf_Result)
    if configs.CheckUpdateCount(configs.GetFuncName(),30,c,Csrf_Result) {
        c.JSON(200, gin.H{"info": configs.Translate("tips.morethannrde")})
        return
    }
	mpartFile, mpartHeader, err := c.Request.FormFile("upload")
	if err != nil {
		mpartFile, mpartHeader, err = c.Request.FormFile("file")
		if err != nil {
			configs.LogErr(err)
			c.JSON(200, gin.H{"uploaded": false, "info": configs.Translate("common.opfailed")})
			return
		}
	}
	defer mpartFile.Close()
	fileExt := filepath.Ext(mpartHeader.Filename)
	fileExt=strings.TrimSpace(strings.ToLower(fileExt))
	if fileExt==".mp4" || fileExt==".mp3" || fileExt==".ogg" || fileExt==".wav" || fileExt==".flac" || fileExt==".ape" || fileExt==".mkv" || fileExt==".avi" || fileExt==".rmvb" {
		err = c.Request.ParseMultipartForm(configs.BbsConfigs.BbsLimits.UploadMedia << 20)
	}else if fileExt==".jpg" || fileExt==".png" || fileExt==".gif" || fileExt==".ico" {
		err = c.Request.ParseMultipartForm(configs.BbsConfigs.BbsLimits.UploadPhoto << 20)
	}else{
		err = c.Request.ParseMultipartForm(configs.BbsConfigs.BbsLimits.UploadElse << 20)
	}
	if err != nil {
		configs.LogErr(err)
		c.JSON(200, gin.H{"uploaded": false, "info":"Media limit "+strconv.FormatInt(configs.BbsConfigs.BbsLimits.UploadMedia,10)+"M,Photo limit "+strconv.FormatInt(configs.BbsConfigs.BbsLimits.UploadPhoto,10)+",Else limit "+strconv.FormatInt(configs.BbsConfigs.BbsLimits.UploadPhoto,10)+"M"})
		return
	}
	uri, err := saveFileAtt(mpartHeader, mpartFile)
	if err != nil {
		configs.LogErr(err)
		c.JSON(200, gin.H{"uploaded": false, "info": configs.Translate("common.opfailed")})
		return
	}

	c.JSON(200, gin.H{"uploaded": true, "url": uri})
}
//saves att file to disc
func saveFileAtt(fh *multipart.FileHeader, f multipart.File) (string, error) {
	uri := ""
	tmpPath := ""
	fullPath := ""
	fullName := ""
	newPath:=fmt.Sprint(time.Now().UnixNano())
		uri = "/documents/attachments/" + SessionUserId +"/" +newPath+"/"+fh.Filename
		tmpPath = configs.BbsConfigs.BbsUploadPath+"/attachments/"+SessionUserId
		fullPath = configs.BbsConfigs.BbsUploadPath+"/attachments/"+SessionUserId+"/"+newPath
	existtmp, err := configs.PathExists(tmpPath)
    if err != nil {
        configs.LogErr(err)
    }
    if !existtmp {
    	err = os.Mkdir(tmpPath, 0777)
        if err != nil {
            configs.LogErr(err)
        }
	}
	exist, err := configs.PathExists(fullPath)
    if err != nil {
        configs.LogErr(err)
    }
    if !exist {
    	err = os.Mkdir(fullPath, 0777)
        if err != nil {
            configs.LogErr(err)
        }
	}
	fullName = filepath.Join(fullPath, fh.Filename)
	file, err := os.OpenFile(fullName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err = io.Copy(file, f)
	if err != nil {
		return "", err
	}
	return uri, nil
}
func DeleteAtt(c *gin.Context) {
	 h := DefaultH(c)
	    if configs.CheckUpdateCount(configs.GetFuncName(),30,c,Csrf_Result) {
        c.JSON(200, gin.H{"info": configs.Translate("tips.morethannrde")})
        return
    }
    att:=c.PostForm("att")
    att = strings.TrimSpace(att)
    if att!=""{
    	err := os.Remove(strings.Replace(att,"/documents/",configs.BbsConfigs.BbsUploadPath,-1))
    	if err != nil {
        	configs.LogErr(err)
    	}
    }
    h["WebTitle"] = "delete"
	c.JSON(200, gin.H{"info": "ok","returnURL": "returnURL"})
}
