package input

import (
	"gebug/pkg/config"
	"github.com/manifoldco/promptui"
	"strings"
)

const (
	doneItem     = `✔ Done`
	removePrefix = `✖️ `
)

type PromptExposePort struct {
	*config.Config
}

func removeItem(items []string, i int) []string {
	if len(items) <= 1 {
		return []string{}
	}

	items[i] = items[len(items)-1]
	return items[:len(items)-1]
}

func addItem(items []string, item string) []string {
	return append(items[:len(items)-1], removePrefix+item, doneItem)
}

func addRemovePrefix(items []string) {
	n := len(items)
	for i := range items {
		if i != n-1 {
			items[i] = removePrefix + items[i]
		}
	}
}

func cleanResults(items []string) []string {
	items = items[:len(items)-1]
	for i := range items {
		items[i] = strings.Replace(items[i], removePrefix, "", 1)
	}
	return items
}

func (p *PromptExposePort) Run() error {
	items := append(p.ExposePorts, doneItem)
	addRemovePrefix(items)

	var result string

	for result != doneItem {
		prompt := &promptui.SelectWithAdd{
			Label:    "Define ports to expose. Press existing choices to delete (e.g: 8080[:8080])",
			Items:    items,
			AddLabel: "Add port",
		}

		cleanResults(items)
		var err error
		var index int
		index, result, err = prompt.Run()
		if err != nil {
			return err
		}
		if result == doneItem {
			continue
		}

		if index < 0 {
			items = addItem(items, result)
		} else {
			items = removeItem(items, index)
		}
	}

	p.ExposePorts = cleanResults(items)
	return nil
}
