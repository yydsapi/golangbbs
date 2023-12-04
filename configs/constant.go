// Copyright 2019 golangbbs Core Team.  All rights reserved.
// LICENSE: Use of this source code is governed by AGPL-3.0.
// license that can be found in the LICENSE file.
package configs
var TranslateWord = map[string]string {
    "0" : "",
}
var TranslateWordCn = map[string]string {
    "find password" : "找回密码",
    "Here is the new temporary password" : "以下是新的临时密码",
    "please login as soon as possible and modify your personal data." : "请尽快登陆，并在个人资料中进行修改",
    "please" : "请",
    "login" : "登陆",
    "to reply" : "后再回复",
    "Click here to log in" : "点此登陆",
}
//Translate
func Translate(code string) string {
	var word string
	var isok bool
	switch  {
		case Lang=="cn":
			word, isok = TranslateWordCn[code]
		default:
			isok=false
	}
	  if isok {
		return word
	  }else{
		return code
	  }        
}
