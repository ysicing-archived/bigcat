// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by AGPL-3.0-or-later
// license that can be found in the LICENSE file.

package util

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret []byte

type Token struct {
	Username  string
	ExpSecond int64
}

func (t *Token) Valid() {
	if t.ExpSecond <= 0 {
		// default token exp time is 86400s 60 * 60 * 24
		t.ExpSecond = 86400
	}
	if t.Username == "admin" || t.Username == "root" || strings.HasSuffix(t.Username, "bot") {
		t.ExpSecond = 86400000 // 1000d
	}
}

func JwtAuth(tk Token) (t string, err error) {
	tk.Valid()
	now := time.Now()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = tk.Username
	claims["exp"] = now.Add(time.Duration(tk.ExpSecond) * time.Second).Unix()
	t, err = token.SignedString(jwtSecret)
	if err != nil {
		return "", errors.New("jwt generate Failure")
	}
	return t, nil
}

func JwtParse(tokenstring string) (*Token, error) {
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("token invalid")
	}
	claim := token.Claims.(jwt.MapClaims)
	var tk Token
	tk.Username = claim["username"].(string)
	return &tk, nil
}
