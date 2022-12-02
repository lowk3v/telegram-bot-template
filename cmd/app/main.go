package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"

	"github.com/author_name/project_name/configs"
	"github.com/author_name/project_name/pkg/app"
)

var args app.Arguments

func banner() string {
	// https://patorjk.com/software/taag/#p=display&f=ANSI%20Shadow&t=app
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
		"Credits: https://github.com/author_name/app\n",
		"Twitter: https://twitter.com/#\n",
		color.YellowString("=================================================="),
	)
}

func init() {
	args = app.Arguments{}

	// delay time between requests
	flag.IntVar(&args.DelayOpt, "delay", 200, "DelayOpt between requests (ms)")
	flag.IntVar(&args.DelayOpt, "d", 200, "DelayOpt between requests (ms)")

	// output folder path
	flag.StringVar(&args.OutputOpt, "output", "contracts", "Specified output folder path")
	flag.StringVar(&args.OutputOpt, "o", "contracts", "Specified output folder path")

	// set proxy options
	flag.StringVar(&args.ProxyOpt, "proxy", "", "Specified proxy options")
	flag.StringVar(&args.ProxyOpt, "x", "", "Specified proxy options")

	flag.Usage = func() {
		h := []string{
			banner(),
			"Usage of: APP_ENV=develop && app <options> <args>",
			"Options",
			"  -d, --delay <delay>       	DelayOpt between issuing requests (ms)",
			"  -o, --output <dir>        	Directory to save responses in (will be created)",
			"  -x, --proxy <proxyURL>    	Use the provided HTTP proxy",
			"",
			"Args",
		}

		_, _ = fmt.Fprintf(os.Stderr, strings.Join(h, "\n"))
	}
	flag.Parse()
}

func main() {
	fmt.Println(banner())
	configs.Secret.PrintInfo()
	app.Run(&args)
}
