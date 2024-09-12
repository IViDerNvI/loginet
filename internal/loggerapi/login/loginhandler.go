package login

import (
	"fmt"
	"net/url"
)

type loginHandler struct {
	UserAccount  string `json:"user_account"`
	UserPassword string `json:"user_password"`
	WlanUserIp   string `json:"wlan_user_ip"`
	IpAddr       string `json:"ipaddr"`
	wlanUserIpv6 string `json:"-"`
	wlanUserMac  string `json:"-"`
	wlanAcIp     string `json:"-"`
	wlanAcName   string `json:"-"`
	jsVersion    string `json:"-"`
	terminalType int    `json:"-"`
	lang         string `json:"-"`
	callback     string `json:"-"`
	login_Method int    `json:"-"`
	v            int    `json:"-"`
	lang_2       string `json:"-"`
}

type Options func(*loginHandler)

type LoginResponse struct {
	Result  int    `json:"result"`
	Msg     string `json:"msg"`
	RetCode int    `json:"ret_code"`
}

var defaultScheme = "https"
var defaultHost = "p.njupt.edu.cn:802"
var defaultPath = "/eportal/portal/login"

func newHandler(opts []Options) *loginHandler {
	var login = &loginHandler{
		UserAccount:  "",
		UserPassword: "",
		WlanUserIp:   "",
		callback:     "dr1003",
		login_Method: 1,
		wlanAcIp:     "",
		wlanAcName:   "",
		wlanUserIpv6: "",
		wlanUserMac:  "000000000000",
		jsVersion:    "4.1.3",
		terminalType: 1,
		lang:         "zh-cn",
		v:            6938,
		lang_2:       "zh",
	}

	for _, opt := range opts {
		opt(login)
	}

	return login
}

func withUserAccount(account string) Options {
	return func(lh *loginHandler) {
		lh.UserAccount = fmt.Sprintf("%%2C0%%2C%s%%40cmcc", account)
	}
}

func withUserPassword(password string) Options {
	return func(lh *loginHandler) {
		lh.UserPassword = password
	}
}

func withWlanUserIp(ipaddr string) Options {
	return func(lh *loginHandler) {
		lh.WlanUserIp = ipaddr
	}
}

func (lh *loginHandler) getQuery() string {
	var query string

	query = fmt.Sprintf("%scallback=%s&", query, lh.callback)
	query = fmt.Sprintf("%slogin_method=%d&", query, lh.login_Method)
	query = fmt.Sprintf("%suser_account=%s&", query, lh.UserAccount)
	query = fmt.Sprintf("%suser_password=%s&", query, lh.UserPassword)
	query = fmt.Sprintf("%swlan_user_ip=%s&", query, lh.WlanUserIp)
	query = fmt.Sprintf("%swlan_user_ipv6=%s&", query, lh.WlanUserIp)
	query = fmt.Sprintf("%swlan_user_mac=%s&", query, lh.wlanUserMac)
	query = fmt.Sprintf("%swlan_ac_ip=%s&", query, lh.wlanAcIp)
	query = fmt.Sprintf("%swlan_ac_name=%s&", query, lh.wlanAcName)
	query = fmt.Sprintf("%sjsVersion=%s&", query, lh.jsVersion)
	query = fmt.Sprintf("%sterminal_type=%d&", query, lh.terminalType)
	query = fmt.Sprintf("%slang=%s&", query, lh.lang)
	query = fmt.Sprintf("%sv=%d&", query, lh.v)
	query = fmt.Sprintf("%slang=%s&", query, lh.lang_2)

	return query
}

func newLoginUrl(opts ...Options) string {
	handler := newHandler(opts)
	query := handler.getQuery()

	var url = url.URL{
		Scheme:   defaultScheme,
		Host:     defaultHost,
		Path:     defaultPath,
		RawQuery: query,
	}

	return url.String()
}
