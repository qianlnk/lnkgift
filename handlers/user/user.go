/**
 * FileName:		user.go
 * Description:		all user's handler
 * Author:			Qianno.Xie
 * Email:			qianlnk@163.com
**/
package user

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"io"
	"io/ioutil"
	"lnkgift/database/mysql"
	"net/http"
	"net/url"
)

var MD5Prefix = "1234567890"

//请求校验
func authVerify(req url.Values) bool {
	ts := req.Get("ts")
	key := req.Get("key")

	if ts == "" || key == "" {
		return false
	}

	h := md5.New()
	io.WriteString(h, MD5Prefix+ts)
	tmpkey := hex.EncodeToString(h.Sum(nil))

	return key == tmpkey
}

func Login(res http.ResponseWriter, req *http.Request) {
	fmt.Println("login...")
	res.Header().Set("Connection", "close")
	res.WriteHeader(200)
	//校验
	if authVerify(req.URL.Query()) == false {
		res.WriteHeader(404)
		fmt.Fprintf(res, "Verify failed, please check your MD5Prefix!")
		return
	}

	//读取body
	type user struct {
		Name   string
		Passwd string
	}
	var loginUser user
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(403)
		fmt.Fprintf(res, "Read body failed, please check your body infomation!")
		return
	}

	json.Unmarshal(body, &loginUser)

	executor := func(ctx context.Context, db *sqlx.DB, dest interface{}) error {
		err := db.Get(&dest, "select 1 from user where email = ? and passed = ?", loginUser.Name, loginUser.Passwd)
		if err != nil {
			fmt.Println("err=", err)
			return err
		}
		return nil
	}

	var exist int
	err = mysql.Invoke(context.TODO(), executor, &exist)
	if err != nil {
		res.WriteHeader(403)
		fmt.Fprintf(res, err.Error())
		return
	}

	if exist == 1 {
		fmt.Fprintf(res, "YES")
	} else {
		fmt.Fprintf(res, "NO")
	}
}
