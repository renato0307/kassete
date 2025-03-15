// (c) 2025 Renato Torres
// GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	sets []Set
}

type Set struct {
	Name  string
	Items []Item
}

type Item struct {
	Name string
	Type string
}

func FromFile(path string) (Config, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("could not config read file: %w", err)
	}
	return New(string(content))
}

func New(yamlConfig string) (Config, error) {
	config := Config{}
	err := yaml.Unmarshal([]byte(yamlConfig), config)
	if err != nil {
		return Config{}, fmt.Errorf("could not unmarshal config: %w", err)
	}
	return Config{}, nil
}

func Test() Config {
	config, _ := New(`
		sets:
		- name: "set1"
			items:
			- name: "item1"
			  type: "type1"
			- name: "item2"
			  type: "type2"
		- name: "set2"
			items:
			- name: "item3"
			  type: "type3"
			- name: "item4"
			  type: "type4"
	`)

	return config
}
