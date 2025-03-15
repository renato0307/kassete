// (c) 2025 Renato Torres
// GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

package main

import (
	"os"

	"github.com/alecthomas/kong"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/renato0307/kassete/internal/config"
	"github.com/renato0307/kassete/internal/logger"
	"github.com/renato0307/kassete/internal/tui"
)

var cli struct {
	ConfigFile string `short:"c" long:"config" help:"Path to the configuration file, if not set, a demo configuration will be used" env:"CONFIG"`
	LogLevel   string `long:"loglevel" help:"Log level, the possible values are debug or info" default:"info" env:"LOG_LEVEL"`
	LogFile    string `long:"logfile" help:"Log file, if not set, logging will be disabled" default:"" env:"LOG_FILE"`
}

func main() {
	ctx := kong.Parse(&cli, kong.Name("kassete"), kong.Description("kassete, the kubernetes TUI to manage resources in sets"))

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
	log.Debug("parsed CLI arguments", "config", cli.ConfigFile)

	cfg := config.Config{}
	if cli.ConfigFile == "" {
		log.Info("no configuration file provided, using demo configuration")
		cfg = config.Test()
	}

	p := tea.NewProgram(tui.NewRootModel(cfg))
	_, err = p.Run()
	ctx.FatalIfErrorf(err)
}
