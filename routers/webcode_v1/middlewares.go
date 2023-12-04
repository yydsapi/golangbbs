// Copyright 2019 golangbbs Core Team.  All rights reserved.
// LICENSE: Use of this source code is governed by AGPL-3.0.
// license that can be found in the LICENSE file.
package webcodev1

import (
	"net/http"
	"net/url"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
)

//AuthRequired  middleware
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		SessionUserId=""
		session := sessions.Default(c)
		sess := session.Get(userIDKey)	
	    if sess != nil {
	        SessionUserId=strings.TrimSpace(sess.(string))
	    }
	    if SessionUserId=="" {
	    	//if cookie, err := c.Request.Cookie("userID"); err == nil {
	    		//SessionUserId=cookie.Value
	    	//}
	    }
    	SessionUserId=strings.TrimSpace(SessionUserId)
		//logrus.Info("?Session==AuthRequired======?"+SessionUserId+"?")
		if SessionUserId != "" {
			c.Next()
		} else {
			c.Redirect(http.StatusFound, "/signin?return="+url.QueryEscape(c.Request.RequestURI))
			c.Abort()
		}
	}
}
