package user

import "github.com/gorilla/sessions"

const (
	//64位
	cookieStoreAuthKey = "r_y52DdjxDD$Oi8wj@7zyv0gYUDpl7eDJi5fKv2mrHN6s79XhIg3sh1muyy#8vnm"
	//AES encrypt key必须是16或者32位
	cookieStoreEncryptKey = "M*40ScmijZnCjEf4il1alwAVzn39cQH3"
)

var sessionStore *sessions.CookieStore

func init () {
	sessionStore = sessions.NewCookieStore(
		[]byte(cookieStoreAuthKey),
		[]byte(cookieStoreEncryptKey),
	)

	sessionStore.Options = &sessions.Options{
		HttpOnly: true,
		MaxAge:   60 * 15,
	}

}