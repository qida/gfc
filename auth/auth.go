/*
 * @Author: qida
 * @Date: 2022-01-07 13:49:59
 * @LastEditTime: 2022-01-07 14:15:00
 * @LastEditors: qida
 * @Description:
 * @FilePath: \zxjy_api_crme:\gopath\lib\src\github.com\qida\gfc\auth\auth.go
 * good day
 */
package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Auth struct {
	User  interface{}
	Xxtea *XxTea
	tm    time.Time
	salt  int
}

const (
	f分隔符 = ":"
)

func NewAuth(key string) *Auth {
	return &Auth{Xxtea: NewXxTea(key)}
}

//加密
func (a *Auth) Encrypt(uid int, userAuth *Auth) string {
	userAuth.salt = rand.Intn(1000)
	userAuth.tm = time.Now()
	src, _ := json.Marshal(userAuth)
	encodeString := base64.URLEncoding.EncodeToString(a.Xxtea.Encrypt(false, src))
	return fmt.Sprintf("%d%s%s", uid, f分隔符, encodeString)
}

//解密
func (a *Auth) Decrypt(uid_auth string) (userAuth Auth, err error) {
	if strings.Contains(uid_auth, f分隔符) {
		uid_auth = strings.SplitN(uid_auth, f分隔符, 2)[1]
	}
	var decodeBytes []byte
	if decodeBytes, err = base64.URLEncoding.DecodeString(uid_auth); err == nil {
		err = json.Unmarshal(a.Xxtea.Decrypt(false, decodeBytes), &userAuth)
	}
	return
}
