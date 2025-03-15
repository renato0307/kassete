// (c) 2025 Renato Torres
// GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

package config

type Config struct {
	ConfigFile string `short:"c" long:"config" help:"Path to the configuration file, if not set, a demo configuration will be used" env:"CONFIG"`
	Dev        bool   `long:"dev" help:"Run in development mode, sets log level to debug and logs to debug.log" env:"DEV"`
	LogLevel   string `long:"loglevel" help:"Log level, the possible values are debug or info" default:"info" env:"LOG_LEVEL"`
	LogFile    string `long:"logfile" help:"Log file, if not set, logging will be disabled" default:"" env:"LOG_FILE"`
	Sets       []Set  `long:"sets" help:"Sets of resources to manage"`
}

type Set struct {
	Name  string `long:"name"`
	Items []Item `long:"items"`
}

type Item struct {
	Name string `long:"name"`
	Type string `long:"type"`
}
