// +build go1.2

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

// Registry server for Gopm (hosting/delivering of packages).
package main

import (
	"github.com/astaxie/beego"

	"github.com/gpmgo/gopm-registry/modules/log"
	"github.com/gpmgo/gopm-registry/modules/setting"
	_ "github.com/gpmgo/gopm-registry/modules/template"
	"github.com/gpmgo/gopm-registry/routers"
)

const APP_VER = "0.0.0.0601"

func init() {
	setting.AppVer = APP_VER
}

func main() {
	log.Info("%s %s", setting.AppName, APP_VER)

	beego.Router("/", &routers.HomeRouter{}, "get:Home")
	beego.Router("/VERSION.json", &routers.VersionRouter{})

	beego.Run()
}
