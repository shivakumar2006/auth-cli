package cli

import "github.com/c-bata/go-prompt"

func Completer(
	ctx *Context,
) func(prompt.Document) []prompt.Suggest {

	return func(
		d prompt.Document,
	) []prompt.Suggest {

		var suggestions []prompt.Suggest

		if !ctx.IsLoggedIn() {

			suggestions = []prompt.Suggest{
				{
					Text:        "register",
					Description: "Create a new user",
				},
				{
					Text:        "login",
					Description: "Login to your account",
				},
				{
					Text:        "help",
					Description: "Show available commands",
				},
				{
					Text:        "exit",
					Description: "Exit the application",
				},
			}

		} else {

			suggestions = []prompt.Suggest{
				{
					Text:        "whoami",
					Description: "Show current user details",
				},
				{
					Text:        "enable-2fa",
					Description: "Enable MFA",
				},
				{
					Text:        "disable-2fa",
					Description: "Disable MFA",
				},
				{
					Text:        "logout",
					Description: "Logout current session",
				},
				{
					Text:        "help",
					Description: "Show available commands",
				},
				{
					Text:        "exit",
					Description: "Exit application",
				},
			}
		}

		return prompt.FilterHasPrefix(
			suggestions,
			d.GetWordBeforeCursor(),
			true,
		)
	}
}
