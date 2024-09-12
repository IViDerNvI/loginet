package loginet

import (
	"fmt"

	myerr "github.com/IViDerNvI/loginet/internal/errors"
	"github.com/IViDerNvI/loginet/internal/loggerapi/login"
	"github.com/IViDerNvI/loginet/internal/loggerapi/logout"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

var (
	v        bool
	username string
	password string
	ipaddr   string
	version  = "0.1.0"
)

var rootCmd = &cobra.Command{
	Use:   "loginet",
	Short: "a login tool of the NJUPT campus network",
	Long:  "loginet is a login tool used to log in or log out of the NJUPT campus network",
	Run: func(cmd *cobra.Command, args []string) {
		if v {
			fmt.Printf("Version: %s\n", version)
		}

	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of loginet",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\n", version)
	},
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to NJUPT-CMCC campus network",
	Run: func(cmd *cobra.Command, args []string) {
		err := login.Login(username, password, ipaddr)

		log.Infof("login by user %s", username)

		if err != nil && !xerrors.Is(err, myerr.ErrSuccess) {
			log.Fatal(err)
			return
		}
		log.Info("login successful, enjoy yourself!")
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out of NJUPT campus network",
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("Logout from %s...", ipaddr)
		err := logout.Logout(ipaddr)
		if err != nil {
			log.Fatal(err)
		}
		log.Info("logout successful!")
	},
}

func init() {
	loginCmd.Flags().StringVarP(&username, "username", "u", "", "Username (required if password is set)")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "Password")
	loginCmd.Flags().StringVarP(&ipaddr, "ipaddr", "i", "", "ipv4 address")

	logoutCmd.Flags().StringVarP(&ipaddr, "ipaddr", "i", "", "ipv4 address")

	rootCmd.Flags().BoolVarP(&v, "version", "v", false, "Show loginet Version")
	rootCmd.AddCommand(versionCmd, loginCmd, logoutCmd)
}
