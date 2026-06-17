package cmd

import (
	"auth-cli/internal/auth"
	"auth-cli/internal/cli"
	"auth-cli/internal/session"
	"auth-cli/internal/totp"

	"github.com/spf13/cobra"
)

type Dependencies struct {
	AuthService    *auth.Service
	SessionManager *session.Manager
	TOTPService    *totp.Service
	CLIContext     *cli.Context
}

var deps *Dependencies

func SetDependencies(d *Dependencies) {
	deps = d
}

var RootCmd = &cobra.Command{
	Use:           "auth-cli",
	Short:         "Contanerized CLI Authentication System",
	Long:          "A CLI authentication system with MFA, sessions and PostgreSQL persistence.",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.AddCommand(
		RegisterCmd,
		LoginCmd,
		LogoutCmd,
		WhoAmICmd,
		Enable2FACmd,
		Disable2FACmd,
		Verify2FACmd,
		HelpCmd,
	)
}
