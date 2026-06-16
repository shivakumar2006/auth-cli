package cli

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
)

type Executor func(command string)

type Shell struct {
	ctx      *Context
	executor Executor
}

func NewShell(ctx *Context, executor Executor) *Shell {
	return &Shell{
		ctx:      ctx,
		executor: executor,
	}
}

func (s *Shell) Start() {
	fmt.Println("===================================")
	fmt.Println("      Auth CLI Login System")
	fmt.Println("===================================")
	fmt.Println("Type 'help' to see commands")
	fmt.Println()

	p := prompt.New(
		func(in string) {
			in = strings.TrimSpace(in)

			if in == "" {
				return
			}

			s.executor(in)
		},
		Completer(s.ctx),

		prompt.OptionPrefix("auth-cli > "),
		prompt.OptionTitle("Auth CLI"),
	)

	p.Run()
}
