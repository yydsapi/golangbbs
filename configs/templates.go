// Copyright 2019 golangbbs Core Team.  All rights reserved.
// LICENSE: Use of this source code is governed by AGPL-3.0.
// license that can be found in the LICENSE file.
package configs

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"
	"github.com/rakyll/statik/fs"
)

var tmpl *template.Template

//loads templates from templates directory
func LoadTemplates() {
	tmpl = template.New("").Funcs(template.FuncMap{
		"now":      now,
		"noescape": noescape,
	})
	fn := func(path string, f os.FileInfo, err error) error {
		if f.IsDir() != true && strings.HasSuffix(f.Name(), ".tpl") {
			var err error
			tmpl, err = tmpl.ParseFiles(path)
			if err != nil {
				return err
			}
		}
		return nil
	}

	if err := filepath.Walk("templates", fn); err != nil {
		panic(err)
	}
}

func LoadStatikTemplates() {
	tmpl = template.New("").Funcs(template.FuncMap{
		"now":      now,
		"noescape": noescape,
	})
	fn := func(path string, f os.FileInfo, err error) error {
		if f.IsDir() != true && strings.HasSuffix(f.Name(), ".tpl") {
			var err error
			tmp,_:=fs.ReadFile(StatikFS,path)
			tmpl, err = tmpl.Parse(string(tmp))
			if err != nil {
				return err
			}
		}
		return nil
	}
	if err := fs.Walk(StatikFS,"/templates", fn); err != nil {
		panic(err)
	}
}

func GetTemplates() *template.Template {
	return tmpl
}

func now() time.Time {
	return time.Now()
}

func noescape(content string) template.HTML {
	return template.HTML(content)
}
