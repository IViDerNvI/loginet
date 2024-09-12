package logout

import (
	"fmt"
	"net/url"
)

type logoutHandler struct {
	callback      string `json:"-"`
	login_Method  int    `json:"-"`
	UserAccount   string `json:"user_account"`
	UserPassword  string `json:"user_password"`
	acLogout      int    `json:"-"`
	register_mode int    `json:"-"`
	WlanUserIp    string `json:"wlan_user_ip"`
	wlanUserIpv6  string `json:"-"`
	wlanVlanId    int    `json:"-"`
	wlanUserMac   string `json:"-"`
	wlanAcIp      string `json:"-"`
	wlanAcName    string `json:"-"`
	jsVersion     string `json:"-"`
	v             int    `json:"-"`
	lang          string `json:"-"`
}

type Options func(*logoutHandler)

type LoginResponse struct {
	Result  int    `json:"result"`
	Msg     string `json:"msg"`
	RetCode int    `json:"ret_code"`
}

var defaultScheme = "https"
var defaultHost = "p.njupt.edu.cn:802"
var defaultPath = "/eportal/portal/logout"

func newHandler(opts []Options) *logoutHandler {
	var login = &logoutHandler{
		UserAccount:   "drcom",
		UserPassword:  "123",
		WlanUserIp:    "",
		callback:      "dr1003",
		login_Method:  1,
		wlanAcIp:      "",
		wlanAcName:    "",
		wlanUserIpv6:  "",
		wlanUserMac:   "000000000000",
		jsVersion:     "4.1.3",
		lang:          "zh",
		acLogout:      1,
		register_mode: 1,
		wlanVlanId:    0,
		v:             4323,
	}

	for _, opt := range opts {
		opt(login)
	}

	return login
}

func withWlanUserIp(ipaddr string) Options {
	return func(lh *logoutHandler) {
		lh.WlanUserIp = ipaddr
	}
}

func (lh *logoutHandler) getQuery() string {
	var query string

	query = fmt.Sprintf("%scallback=%s&", query, lh.callback)
	query = fmt.Sprintf("%slogin_method=%d&", query, lh.login_Method)
	query = fmt.Sprintf("%suser_account=%s&", query, lh.UserAccount)
	query = fmt.Sprintf("%suser_password=%s&", query, lh.UserPassword)
	query = fmt.Sprintf("%sac_logout=%d&", query, lh.acLogout)
	query = fmt.Sprintf("%sregister_mode=%d&", query, lh.register_mode)
	query = fmt.Sprintf("%swlan_user_ip=%s&", query, lh.WlanUserIp)
	query = fmt.Sprintf("%swlan_user_ipv6=%s&", query, lh.wlanUserIpv6)
	query = fmt.Sprintf("%swlan_vlan_id=%d&", query, lh.wlanVlanId)
	query = fmt.Sprintf("%swlan_user_mac=%s&", query, lh.wlanUserMac)
	query = fmt.Sprintf("%swlan_ac_ip=%s&", query, lh.wlanAcIp)
	query = fmt.Sprintf("%swlan_ac_name=%s&", query, lh.wlanAcName)
	query = fmt.Sprintf("%sjsVersion=%s&", query, lh.jsVersion)
	query = fmt.Sprintf("%sv=%d&", query, lh.v)
	query = fmt.Sprintf("%slang=%s", query, lh.lang)

	return query
}

func newLogoutUrl(opts ...Options) string {
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
