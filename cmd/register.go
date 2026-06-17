package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var RegisterCmd = &cobra.Command{
	Use:   "register <username> <password>",
	Short: "Register a new user",
	Args:  cobra.ExactArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {

		username := strings.TrimSpace(args[0])
		password := strings.TrimSpace(args[1])

		if err := deps.AuthService.Register(
			username,
			password,
		); err != nil {
			return fmt.Errorf(
				"registration failed: %w",
				err,
			)
		}

		fmt.Println("✓ User registered successfully")

		return nil
	},
}

// package cmd

// import (
// 	"fmt"
// 	"strings"

// 	"github.com/spf13/cobra"
// )

// var RegisteredCmd = &cobra.Command{
// 	Use:   "register <username> <password>",
// 	Short: "Register a new user",
// 	Args:  cobra.ExactArgs(2),
// 	RunE: func(cmd *cobra.Command, args []string) error {

// 		username := args[0]
// 		password := args[1]

// 		// reader := bufio.NewReader(os.Stdin)

// 		fmt.Print("Username: ")
// 		// username, err := reader.ReadString('\n')
// 		// if err != nil {
// 		// 	return err
// 		// }

// 		username = strings.TrimSpace(username)

// 		fmt.Print("Password: ")
// 		// passwordBytes, err := term.ReadPassword(
// 		// 	int(os.Stdin.Fd()),
// 		// )
// 		fmt.Println()

// 		// password := strings.TrimSpace(string(passwordBytes))

// 		if err := deps.AuthService.Register(username, password); err != nil {
// 			return fmt.Errorf("registration failed: %w", err)
// 		}

// 		fmt.Println()
// 		fmt.Println("✓ User registered successfully")

// 		return nil
// 	},
// }
