package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/author_name/project_name/configs"
	"github.com/author_name/project_name/internal/tgbot"
	"github.com/fatih/color"
)

func banner() string {
	// https://patorjk.com/software/taag/#p=display&f=ANSI%20Shadow&t=%20project_name
	return fmt.Sprintln(
		color.YellowString("==================================================\n"),
		color.HiBlueString(`
     █████╗ ██████╗ ██████╗
    ██╔══██╗██╔══██╗██╔══██╗
    ███████║██████╔╝██████╔╝
    ██╔══██║██╔═══╝ ██╔═══╝
    ██║  ██║██║     ██║
    ╚═╝  ╚═╝╚═╝     ╚═╝ `+" By @author_name\n"),
		color.BlueString("project_description\n"),
		"Credits: https://github.com/author_name/project_name\n",
		"Twitter: https://twitter.com/#\n",
		color.YellowString("=================================================="),
	)
}

type Arguments struct {
	DelayOpt  int
	OutputOpt string
	EnvOpt    string
	ProxyOpt  string
	ModeOpt   string
}

var args Arguments

func init() {
	args = Arguments{}

	flag.Usage = func() {
		h := []string{
			banner(),
			"Usage of: project_name <options> <args>",
			"Options",
			"Args",
		}

		_, _ = fmt.Fprintf(os.Stderr, strings.Join(h, "\n"))
	}
	flag.Parse()
}

func main() {
	fmt.Println(banner())
	configs.AppConfig.PrintInfo()

	bot := tgbot.New()
	// Notify service
	//go bot.SubscribeNotification()
	bot.RunBotController()
}
