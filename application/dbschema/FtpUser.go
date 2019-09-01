// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo/param"
	
	"time"
)

// FtpUser FTP用户
type FtpUser struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*FtpUser
	namer   func(string) string
	connID  int
	
	Id          	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	Username    	string  	`db:"username" bson:"username" comment:"用户名" json:"username" xml:"username"`
	Password    	string  	`db:"password" bson:"password" comment:"密码" json:"password" xml:"password"`
	Banned      	string  	`db:"banned" bson:"banned" comment:"是否禁止连接" json:"banned" xml:"banned"`
	Directory   	string  	`db:"directory" bson:"directory" comment:"授权目录(一行一个) " json:"directory" xml:"directory"`
	IpWhitelist 	string  	`db:"ip_whitelist" bson:"ip_whitelist" comment:"IP白名单(一行一个) " json:"ip_whitelist" xml:"ip_whitelist"`
	IpBlacklist 	string  	`db:"ip_blacklist" bson:"ip_blacklist" comment:"IP黑名单(一行一个) " json:"ip_blacklist" xml:"ip_blacklist"`
	Created     	uint    	`db:"created" bson:"created" comment:"创建时间 " json:"created" xml:"created"`
	Updated     	uint    	`db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
	GroupId     	uint    	`db:"group_id" bson:"group_id" comment:"用户组" json:"group_id" xml:"group_id"`
}

func (this *FtpUser) Trans() *factory.Transaction {
	return this.trans
}

func (this *FtpUser) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *FtpUser) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *FtpUser) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *FtpUser) Objects() []*FtpUser {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *FtpUser) NewObjects() *[]*FtpUser {
	this.objects = []*FtpUser{}
	return &this.objects
}

func (this *FtpUser) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *FtpUser) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *FtpUser) Short_() string {
	return "ftp_user"
}

func (this *FtpUser) Struct_() string {
	return "FtpUser"
}

func (this *FtpUser) Name_() string {
	if this.namer != nil {
		return WithPrefix(this.namer(this.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(this.Short_())(this))
}

func (this *FtpUser) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *FtpUser) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *FtpUser) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *FtpUser) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *FtpUser) GroupBy(keyField string, inputRows ...[]*FtpUser) map[string][]*FtpUser {
	var rows []*FtpUser
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string][]*FtpUser{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*FtpUser{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (this *FtpUser) KeyBy(keyField string, inputRows ...[]*FtpUser) map[string]*FtpUser {
	var rows []*FtpUser
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]*FtpUser{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (this *FtpUser) AsKV(keyField string, valueField string, inputRows ...[]*FtpUser) map[string]interface{} {
	var rows []*FtpUser
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]interface{}{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (this *FtpUser) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *FtpUser) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Banned) == 0 { this.Banned = "N" }
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	return
}

func (this *FtpUser) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	this.Updated = uint(time.Now().Unix())
	if len(this.Banned) == 0 { this.Banned = "N" }
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *FtpUser) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *FtpUser) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *FtpUser) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	
	if val, ok := kvset["banned"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["banned"] = "N" } }
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *FtpUser) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		this.Updated = uint(time.Now().Unix())
	if len(this.Banned) == 0 { this.Banned = "N" }
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Banned) == 0 { this.Banned = "N" }
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	return 
}

func (this *FtpUser) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *FtpUser) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *FtpUser) Reset() *FtpUser {
	this.Id = 0
	this.Username = ``
	this.Password = ``
	this.Banned = ``
	this.Directory = ``
	this.IpWhitelist = ``
	this.IpBlacklist = ``
	this.Created = 0
	this.Updated = 0
	this.GroupId = 0
	return this
}

func (this *FtpUser) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Username"] = this.Username
	r["Password"] = this.Password
	r["Banned"] = this.Banned
	r["Directory"] = this.Directory
	r["IpWhitelist"] = this.IpWhitelist
	r["IpBlacklist"] = this.IpBlacklist
	r["Created"] = this.Created
	r["Updated"] = this.Updated
	r["GroupId"] = this.GroupId
	return r
}

func (this *FtpUser) Set(key interface{}, value ...interface{}) factory.Model {
	switch k := key.(type) {
		case map[string]interface{}:
			for kk, vv := range k {
				this.Set(kk, vv)
			}
		default:
			var (
				kk string
				vv interface{}
			)
			if k, y := key.(string); y {
				kk = k
			} else {
				kk = fmt.Sprint(key)
			}
			if len(value) > 0 {
				vv = value[0]
			}
			switch kk {
				case "Id": this.Id = param.AsUint(vv)
				case "Username": this.Username = param.AsString(vv)
				case "Password": this.Password = param.AsString(vv)
				case "Banned": this.Banned = param.AsString(vv)
				case "Directory": this.Directory = param.AsString(vv)
				case "IpWhitelist": this.IpWhitelist = param.AsString(vv)
				case "IpBlacklist": this.IpBlacklist = param.AsString(vv)
				case "Created": this.Created = param.AsUint(vv)
				case "Updated": this.Updated = param.AsUint(vv)
				case "GroupId": this.GroupId = param.AsUint(vv)
			}
	}
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Username"] = this.Username
	r["Password"] = this.Password
	r["Banned"] = this.Banned
	r["Directory"] = this.Directory
	r["IpWhitelist"] = this.IpWhitelist
	r["IpBlacklist"] = this.IpBlacklist
	r["Created"] = this.Created
	r["Updated"] = this.Updated
	r["GroupId"] = this.GroupId
	return r
}

func (this *FtpUser) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["username"] = this.Username
	r["password"] = this.Password
	r["banned"] = this.Banned
	r["directory"] = this.Directory
	r["ip_whitelist"] = this.IpWhitelist
	r["ip_blacklist"] = this.IpBlacklist
	r["created"] = this.Created
	r["updated"] = this.Updated
	r["group_id"] = this.GroupId
	return r
}

func (this *FtpUser) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate(this.Short_(), kvset)
}

func (this *FtpUser) Validate(field string, value interface{}) error {
	return factory.Validate(this.Short_(), field, value)
}

