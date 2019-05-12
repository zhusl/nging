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

package seaweedfs

import (
	"io"
	"net/url"
	"path"

	"github.com/admpub/goseaweedfs"
	"github.com/admpub/nging/application/registry/upload"
)

const Name = `seaweedfs`

var _ upload.Uploader = &Seaweedfs{}

func init() {
	upload.UploaderRegister(Name, func(typ string) upload.Uploader {
		return NewSeaweedfs(typ)
	})
}

func NewSeaweedfs(typ string) *Seaweedfs {
	a := DefaultConfig.New()
	uploadPath := `public/upload/` + typ
	return &Seaweedfs{
		config:     DefaultConfig,
		instance:   a,
		Type:       typ,
		UploadPath: uploadPath,
	}
}

type Seaweedfs struct {
	config     *Config
	instance   *goseaweedfs.Seaweed
	Type       string
	UploadPath string
}

func (s *Seaweedfs) Engine() string {
	return Name
}

func (f *Seaweedfs) filepath(fname string) string {
	return path.Join(f.UploadPath, fname)
}

func (f *Seaweedfs) Put(dstFile string, src io.Reader, size int64) (string, error) {
	file := f.filepath(dstFile)
	rs, err := f.instance.Filers[0].Upload(src, size, file, f.Type, f.config.TTL)
	if err != nil {
		return "", err
	}
	return rs.FileURL, nil
}

func (f *Seaweedfs) Get(dstFile string) (io.ReadCloser, error) {
	filer := f.instance.Filers[0]
	_, readCloser, err := filer.Download(dstFile)
	return readCloser, err
}

func (f *Seaweedfs) Delete(dstFile string) error {
	filer := f.instance.Filers[0]
	return filer.Delete(dstFile)
}

func (f *Seaweedfs) DeleteDir(dstDir string) error {
	return f.instance.Filers[0].Delete(dstDir, true)
}

func (f *Seaweedfs) apiPut(dstFile string, src io.Reader, size int64) (string, error) {
	_, fID, err := f.instance.Upload(src, dstFile, size, f.Type, f.config.TTL)
	if err != nil {
		return "", err
	}
	view, err := f.instance.LookupFileID(fID, url.Values{}, true)
	if err != nil {
		return view, err
	}
	return view, nil
}

func (f *Seaweedfs) apiGet(fileID string) (io.ReadCloser, error) {
	_, readCloser, err := f.instance.Download(fileID, nil)
	return readCloser, err
}

func (f *Seaweedfs) apiDelete(fileID string) error {
	return f.instance.DeleteFile(fileID, nil)
}