// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by AGPL-3.0-or-later
// license that can be found in the LICENSE file.

package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string `gorm:"type:varchar(50);not null;index:user_idx" json:"username"`
	Password   string `gorm:"type:varchar(150);not null" json:"password"`
	Department string `gorm:"type:varchar(50);" json:"department"`
	RealName   string `gorm:"type:varchar(50);" json:"real_name"`
	Email      string `gorm:"type:varchar(50);" json:"email"`
}

func init() {
	migrate(User{})
}

func (User) TableName() string {
	return "user"
}
