// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by AGPL-3.0-or-later
// license that can be found in the LICENSE file.

package model

import (
	"github.com/ergoapi/util/color"
	"github.com/ergoapi/util/exid"
	"github.com/ergoapi/util/expass"
	"github.com/ergoapi/zlog"
	"github.com/ysicing/bigcat/pkg/util/config"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string `gorm:"type:varchar(50);not null;index:user_idx" json:"username"`
	Password   string `gorm:"type:varchar(150);not null" json:"-"`
	Department string `gorm:"type:varchar(50);" json:"department"`
	RealName   string `gorm:"type:varchar(50);" json:"real_name"`
	Email      string `gorm:"type:varchar(50);" json:"email"`
	TOTPSecret string `json:"-"`
	Banned     bool   `json:"banned"`
	Type       string `gorm:"type:varchar(50);default:internal" json:"type"` // ldap,internal,github
	Token      string `gorm:"type:varchar(150);" json:"-"`
}

func init() {
	migrate(User{})
}

func (User) TableName() string {
	return "user"
}

func (u *User) Save() error {
	var uu User
	err := db.Model(User{}).Where("username = ?", u.Username).Last(&uu).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if uu.ID > 0 {
		uu.Banned = u.Banned
		uu.Email = u.Email
		return db.Save(&uu).Error
	}
	return db.Create(&u).Error
}

// InitAdmin init
func InitAdmin() {
	val, err := ConfigsGet("initadmin")
	if err != nil {
		zlog.Fatal("cannot query initadmin err: %v", err)
	}
	if val != "" {
		zlog.Info(color.SGreen("exist initadmin %v success...", val))
		return
	}
	salt, _ := ConfigsGet("salt")
	user := config.GetString("server.admin.user", "admin")
	adminuser := User{
		Username:   user,
		Password:   expass.SaltPbkdf2Pass(salt, config.GetString("server.admin.pass", "admin")),
		RealName:   user,
		Department: "未知",
		Email:      config.GetString("server.admin.mail", "root@example.com"),
		Banned:     false,
		Token:      exid.GenUID(user),
	}
	if err := adminuser.Save(); err != nil {
		zlog.Fatal("init admin in mysql err: %v", err)
	}
	err = ConfigsSet("initadmin", "done")
	if err != nil {
		zlog.Fatal("init initadmin in mysql err: %v", err)
	}
	zlog.Info(color.SGreen("init  admin %v success...", user))
}
