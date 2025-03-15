// (c) 2025 Renato Torres
// GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

package main

import (
	"os"

	"github.com/alecthomas/kong"
	kongyaml "github.com/alecthomas/kong-yaml"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/renato0307/kassete/internal/config"
	"github.com/renato0307/kassete/internal/logger"
	"github.com/renato0307/kassete/internal/tui"
)

var cli config.Config

func main() {
	ctx := kong.Parse(&cli,
		kong.Name("kassete"),
		kong.Description("kassete, the kubernetes TUI to manage resources in sets"),
		kong.Configuration(kongyaml.Loader, "~/.kassete.yaml", "~/.config/kassete.yaml", "demo.yaml"),
	)

	if cli.Dev {
		cli.LogLevel = logger.LogLevelDebug
		cli.LogFile = "debug.log"
		os.Remove(cli.LogFile) // remove previous log file, if any
	}

	var logFile *os.File
	var err error
	if cli.LogFile != "" {
		logFile, err = tea.LogToFile(cli.LogFile, cli.LogLevel)
		ctx.FatalIfErrorf(err)
	}
	defer func() {
		if logFile != nil {
			logFile.Close()
		}
	}()
	log := logger.InitDefaultLogger(logFile, cli.LogLevel)
	log.Info("starting kassete")
	log.Debug("parsed CLI arguments",
		"config", cli.ConfigFile,
		"dev", cli.Dev,
		"logfile", cli.LogFile,
		"loglevel", cli.LogLevel,
	)
	log.Debug("loaded configuration", "sets", len(cli.Sets))
	for _, set := range cli.Sets {
		log.Debug("loaded set", "name", set.Name, "items", len(set.Items))
	}

	p := tea.NewProgram(tui.NewRootModel(cli), tea.WithAltScreen())
	_, err = p.Run()
	ctx.FatalIfErrorf(err)
}
