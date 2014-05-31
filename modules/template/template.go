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

package template

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"

	"github.com/gpmgo/gopm-registry/modules/setting"
)

func init() {
	beego.AddFuncMap("AppName", func() string {
		return setting.AppName
	})
	beego.AddFuncMap("AppVer", func() string {
		return setting.AppVer
	})
	beego.AddFuncMap("i18n", i18n.Tr)
}
