package login

import (
	"github.com/IViDerNvI/loginet/internal/loggerapi"
	"github.com/IViDerNvI/loginet/internal/util"
	"golang.org/x/xerrors"
)

func Login(account, password, ipaddr string) error {
	if account == "" {
		return xerrors.New("No Account")
	} else if password == "" {
		return xerrors.New("No Password")
	}

	var err error
	if ipaddr == "" {
		ipaddr, err = util.GetIpAddr()
		if err != nil {
			return xerrors.Errorf("Login failed: %w", err)
		}
	}

	var url = newLoginUrl(
		withUserAccount(account),
		withUserPassword(password),
		withWlanUserIp(ipaddr),
	)

	respon, err := util.UrlRequest(url)
	if err != nil {
		return xerrors.Errorf("Login failed: %w", err)
	}

	return loggerapi.GetLoginResult(respon)
}
