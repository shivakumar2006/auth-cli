package main

import (
	"auth-cli/cmd"
	authdb "auth-cli/internal/auth"
	"auth-cli/internal/cli"
	"auth-cli/internal/database"
	db "auth-cli/internal/database"
	session "auth-cli/internal/session"
	totpservice "auth-cli/internal/totp"
	"fmt"
	"os"
	"strings"
)

func main() {
	databaseURL := database.BuildDatabaseURL()

	if err := database.RunMigrations(databaseURL); err != nil {
		fmt.Printf("migration error: %v\n", err)
		os.Exit(1)
	}

	// db connection
	postgres, err := db.NewPostgresConnection()
	if err != nil {
		fmt.Printf("failed to connect database : %v\n", err)
		os.Exit(1)
	}
	defer postgres.Close()

	// auth layer
	authStorage := authdb.NewDatabase(
		postgres,
	)

	authService := authdb.NewService(
		authStorage,
	)

	// session
	sessionStorage := session.NewDatabase(
		postgres,
	)

	sessionManager := session.NewManager(
		sessionStorage,
	)

	// totp
	totpService := totpservice.NewService()

	// cli context
	cliContext := cli.NewContext()

	cmd.SetDependencies(
		&cmd.Dependencies{
			AuthService:    authService,
			SessionManager: sessionManager,
			TOTPService:    totpService,
			CLIContext:     cliContext,
		},
	)

	// shell
	shell := cli.NewShell(cliContext, func(input string) {
		input = strings.TrimSpace(input)

		if input == "" {
			return
		}

		if input == "exit" {
			fmt.Println("Goodbye")
			os.Exit(0)
		}

		args := strings.Fields(input)

		cmd.RootCmd.SetArgs(args)

		if err := cmd.Execute(); err != nil {
			fmt.Printf("Error: %v\n", err)
		}

	})

	shell.Start()

}
