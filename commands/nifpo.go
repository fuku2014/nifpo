package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	prompt "github.com/c-bata/go-prompt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var rootCommandSuggestions = []prompt.Suggest{}
var commandFlagSuggestions = make(map[string][]prompt.Suggest)
var globalFlagSuggestions = []prompt.Suggest{
	prompt.Suggest{
		Text:        "--access-key",
		Description: "NIFCLOUD API ACCESS KEY",
	},
	prompt.Suggest{
		Text:        "--secret-key",
		Description: "NIFCLOUD API SECRET KEY",
	},
	prompt.Suggest{
		Text:        "--region",
		Description: "NIFCLOUD Region",
	},
}

// Nifpo is a top level command instance
var Nifpo = &cobra.Command{
	Short: "The NIFCLOUD Command Prompt is a unified tool to manage your NIFCLOUD services.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()

		defer fmt.Println("Bye!")

		p := prompt.New(
			executor,
			completer,
			prompt.OptionTitle("nifpo: interactive nifcloud client"),
			prompt.OptionPrefix(">> "),
		)

		p.Run()
	},
}

func init() {
	Nifpo.PersistentFlags().BoolP("debug", "", false, "Enable debug mode")
	Nifpo.PersistentFlags().StringP("region", "", "", "NIFCLOUD Region (default NIFCLOUD_DEFAULT_REGION environment variable)")
	Nifpo.PersistentFlags().StringP("access-key", "", "", "NIFCLOUD API ACCESS KEY (default NIFCLOUD_ACCESS_KEY_ID environment variable")
	Nifpo.PersistentFlags().StringP("secret-key", "", "", "NIFCLOUD API SECRET KEY (default NIFCLOUD_SECRET_ACCESS_KEY environment variable")

	viper.BindPFlag("access-key", Nifpo.PersistentFlags().Lookup("access-key"))
	viper.BindPFlag("secret-key", Nifpo.PersistentFlags().Lookup("secret-key"))
	viper.BindPFlag("region", Nifpo.PersistentFlags().Lookup("region"))
	viper.BindPFlag("debug", Nifpo.PersistentFlags().Lookup("debug"))

	Nifpo.AddCommand(version())
	Nifpo.AddCommand(computing())

	rootCommandSuggestions = getCommandSuggestions(Nifpo)
	getOptionSuggestions(Nifpo)
}

func executor(s string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	} else if s == "quit" || s == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}

	commandToExecute := exec.Command("/bin/sh", "-c", "nifpo "+s)
	commandToExecute.Stdin = os.Stdin
	commandToExecute.Stdout = os.Stdout
	commandToExecute.Stderr = os.Stderr
	if err := commandToExecute.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
	return
}

func completer(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	args := strings.Split(d.TextBeforeCursor(), " ")
	w := d.GetWordBeforeCursor()

	for i := range args {
		if args[i] == "|" {
			return []prompt.Suggest{}
		}
	}

	if strings.HasPrefix(w, "-") {
		return optionCompleter(args, strings.HasPrefix(w, "--"))
	}

	return argumentsCompleter(excludeOptions(args))
}

func argumentsCompleter(args []string) []prompt.Suggest {
	// nifprompt
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(rootCommandSuggestions, args[0], true)
	}

	// service name
	if len(args) == 2 {
		switch args[0] {
		case "computing":
			suggestions := getCommandSuggestions(computing())
			return prompt.FilterHasPrefix(suggestions, args[1], true)
		}
	}

	return []prompt.Suggest{}
}

func getCommandSuggestions(cmd *cobra.Command) []prompt.Suggest {
	suggestions := []prompt.Suggest{}
	for _, command := range cmd.Commands() {
		var suggestion = prompt.Suggest{
			Text:        command.Name(),
			Description: command.Short,
		}
		suggestions = append(suggestions, suggestion)
	}
	return suggestions
}

func getOptionSuggestions(cmd *cobra.Command) {
	suggestions := []prompt.Suggest{}
	for _, command := range cmd.Commands() {
		command.Flags().VisitAll(func(flag *pflag.Flag) {
			suggestions = append(suggestions, prompt.Suggest{
				Text:        "--" + flag.Name,
				Description: flag.Usage,
			})
		})
		commandFlagSuggestions[command.Name()] = suggestions
		getOptionSuggestions(command)
	}
}

func excludeOptions(args []string) []string {
	ret := make([]string, 0, len(args))
	for i := range args {
		if !strings.HasPrefix(args[i], "-") {
			ret = append(ret, args[i])
		}
	}
	return ret
}

func optionCompleter(args []string, long bool) []prompt.Suggest {
	l := len(args)
	if l <= 2 {
		return globalFlagSuggestions
	}

	commandArgs := excludeOptions(args)
	command := commandArgs[1]
	suggests := commandFlagSuggestions[command]

	if long {
		return prompt.FilterContains(
			prompt.FilterHasPrefix(suggests, "--", false),
			strings.TrimLeft(args[l-1], "--"),
			true,
		)
	}
	return prompt.FilterContains(suggests, strings.TrimLeft(args[l-1], "-"), true)
}
