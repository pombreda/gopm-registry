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
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/howeyc/fsnotify"

	"github.com/gpmgo/gopm-registry/modules/log"
	"github.com/gpmgo/gopm-registry/modules/setting"
)

func monitorI18nLocale() {
	log.Trace("Monitor i18n locale files enabled")

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("Fail to init locale watcher: %v", err)
	}

	go func() {
		for {
			select {
			case event := <-watcher.Event:
				switch filepath.Ext(event.Name) {
				case ".ini":
					if err := i18n.ReloadLangs(); err != nil {
						log.Error("Fail to relaod locale file reloaded: %v", err)
					}
					log.Trace("Locale file reloaded: %s", strings.TrimPrefix(event.Name, "conf/locale/"))
				}
			}
		}
	}()

	if err := watcher.WatchFlags("conf/locale", fsnotify.FSN_MODIFY); err != nil {
		log.Fatal("Fail to start locale watcher: %v", err)
	}
}

func init() {
	beego.AddFuncMap("AppName", func() string {
		return setting.AppName
	})
	beego.AddFuncMap("AppVer", func() string {
		return setting.AppVer
	})
	beego.AddFuncMap("i18n", i18n.Tr)

	if !setting.ProdMode {
		monitorI18nLocale()
	}
}
