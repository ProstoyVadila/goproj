package info

import (
	"github.com/ProstoyVadila/goproj/pkg/output"
	"github.com/elliotchance/orderedmap/v2"
)

const (
	Version = "1.1.6"
	Author  = "Vadim Gorbachev"
	Repo    = "github.com/ProstoyVadila/goproj"
)

func Show() {
	output.Show(show())
}

func show() (*orderedmap.OrderedMap[string, any], string) {
	msg := "Goproj package info"
	omap := orderedmap.NewOrderedMap[string, any]()

	omap.Set("Version: %s", Version)
	omap.Set("Author: %s", Author)
	omap.Set("Source: %s", Repo)

	return omap, msg
}
