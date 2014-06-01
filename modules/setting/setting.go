// Copyright 2014 Unknown
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package setting

import (
	"strings"

	"github.com/Unknwon/goconfig"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"

	"github.com/gpmgo/gopm-registry/modules/log"
)

var (
	AppName string
	AppVer  string

	Cfg       *goconfig.ConfigFile
	ProdMode  bool
	Languages []string
)

func init() {
	var err error
	Cfg, err = goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		log.Fatal("Fail to load 'conf/app.ini': %v", err)
	}

	AppName = Cfg.MustValue("", "APP_NAME")

	Languages = strings.Split(Cfg.MustValue("app", "LANGUAGES"), "|")
	for _, lang := range Languages {
		if err = i18n.SetMessage(lang, "conf/locale/locale_"+lang+".ini"); err != nil {
			log.Fatal("Fail to load locale file: %v", err)
		}
	}

	// Setting up environment.
	beego.RunMode = Cfg.MustValue("server", "RUN_MODE")
	if beego.RunMode == "prod" {
		ProdMode = true
		log.Info("Production Mode Enabled")
	}
	beego.HttpPort = Cfg.MustInt("server", "HTTP_PORT")
	beego.ViewsPath = "templates"
}
