// +build bindata

/*
   Nging is a toolbox for webmasters
   Copyright (C) 2018-present  Wenhui Shen <swh@admpub.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package bindata

import (
	"net/http"
	"strings"
	"os"

	"github.com/admpub/nging/application/initialize/backend"

	assetfs "github.com/admpub/go-bindata-assetfs"
	"github.com/admpub/nging/application/cmd/event"
	"github.com/admpub/nging/application/library/modal"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/middleware/bindata"
	"github.com/webx-top/image"
	"github.com/admpub/nging/application/registry/upload/helper"
)

func NewAssetFS(prefix string) *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    prefix,
	}
}

var (
	StaticAssetPrefix       string
	BackendTmplAssetPrefix  = "template/backend"
	FrontendTmplAssetPrefix = "template/frontend"
	StaticAssetFS           *assetfs.AssetFS
	BackendTmplAssetFS      *assetfs.AssetFS
	FrontendTmplAssetFS     *assetfs.AssetFS
	LanguageAssetFSFunc     = func(dir string) http.FileSystem {
		return NewAssetFS(dir)
	}
)

func Initialize() {
	event.Bindata = true
	if StaticAssetFS == nil {
		StaticAssetFS = NewAssetFS(StaticAssetPrefix)
	}
	if BackendTmplAssetFS == nil {
		BackendTmplAssetFS = NewAssetFS(BackendTmplAssetPrefix)
	}
	if FrontendTmplAssetFS == nil {
		FrontendTmplAssetFS = NewAssetFS(FrontendTmplAssetPrefix)
	}
	event.StaticMW = bindata.Static("/public/assets/", StaticAssetFS)
	event.FaviconHandler = func(c echo.Context) error {
		return c.File(event.FaviconPath, StaticAssetFS)
	}
	event.BackendTmplMgr = bindata.NewTmplManager(BackendTmplAssetFS)
	event.FrontendTmplMgr = bindata.NewTmplManager(FrontendTmplAssetFS)
	modal.ReadConfigFile = func(file string) ([]byte, error) {
		file = strings.TrimPrefix(file, backend.TemplateDir)
		return event.BackendTmplMgr.GetTemplate(file)
	}
	image.WatermarkOpen = func(file string) (image.FileReader, error) {
		f, err := image.DefaultHTTPSystemOpen(file)
		if err != nil {
			if os.IsNotExist(err) {
				if strings.HasPrefix(file, echo.Wd()) {
					file = strings.TrimPrefix(file, echo.Wd())
					return StaticAssetFS.Open(file)
				}
				if strings.HasPrefix(file, helper.DefaultUploadURLPath) || strings.HasPrefix(file, `/public/assets/`) {
					return StaticAssetFS.Open(file)
				}
			}
		}
		return f, err
	}
	event.LangFSFunc = LanguageAssetFSFunc
}
