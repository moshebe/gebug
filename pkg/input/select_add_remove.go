package input

import (
	"strings"

	"github.com/manifoldco/promptui"
)

const (
	defaultDoneItem     = `✔ Done`
	defaultRemovePrefix = `✖️ `
)

// SelectWithAddAndRemove provides a ui prompt for a list that supports both adding to and removing from a list
type SelectWithAddAndRemove struct {
	prompt       *promptui.SelectWithAdd
	field        *[]string
	doneItem     string
	removePrefix string
}

// NewSelectWithAddAndRemove constructs SelectWithAddAndRemove
func NewSelectWithAddAndRemove(field *[]string, prompt *promptui.SelectWithAdd) *SelectWithAddAndRemove {
	return &SelectWithAddAndRemove{
		prompt:       prompt,
		field:        field,
		doneItem:     defaultDoneItem,
		removePrefix: defaultRemovePrefix,
	}
}

func (s *SelectWithAddAndRemove) remove(items []string, i int) []string {
	if len(items) <= 1 {
		return []string{}
	}

	items[i] = items[len(items)-1]
	return items[:len(items)-1]
}

func (s *SelectWithAddAndRemove) add(items []string, item string) []string {
	return append(items[:len(items)-1], s.removePrefix+item, s.doneItem)
}

func (s *SelectWithAddAndRemove) appendRemovePrefix(items []string) {
	n := len(items)
	for i := range items {
		if i != n-1 {
			items[i] = s.removePrefix + items[i]
		}
	}
}

func (s *SelectWithAddAndRemove) cleanResults(items []string) []string {
	items = items[:len(items)-1]
	for i := range items {
		items[i] = strings.Replace(items[i], s.removePrefix, "", 1)
	}
	return items
}

// Run runs the prompt
func (s *SelectWithAddAndRemove) Run() error {
	items := append(*s.field, s.doneItem)
	s.appendRemovePrefix(items)

	var result string

	for result != s.doneItem {
		var err error
		var index int
		s.prompt.Items = items
		index, result, err = s.prompt.Run()
		if err != nil {
			return err
		}
		if result == s.doneItem {
			continue
		}

		if index < 0 {
			items = s.add(items, result)
		} else {
			items = s.remove(items, index)
		}
	}

	*s.field = s.cleanResults(items)
	return nil
}
