package logout

import (
	"github.com/IViDerNvI/loginet/internal/util"

	"golang.org/x/xerrors"
)

func Logout(ipaddr string) error {
	var err error
	if ipaddr == "" {
		ipaddr, err = util.GetIpAddr()
		if err != nil {
			return xerrors.Errorf("Logout failed: %w", err)
		}
	}

	var url = newLogoutUrl(
		withWlanUserIp(ipaddr),
	)

	_, err = util.UrlRequest(url)
	if err != nil {

		return xerrors.Errorf("Logout failed: %w", err)
	}

	return nil
}
