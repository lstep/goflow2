package common

import (
	"github.com/jnovack/flag"
	"strings"
	"sync"
)

var (
	selectorVar string
	selector    []string // Hashing fields
	selectorTag string   // Hashing fields

	selectorDeclared     bool
	selectorDeclaredLock = &sync.Mutex{}
)

func SelectorFlag() {
	selectorDeclaredLock.Lock()
	defer selectorDeclaredLock.Unlock()

	if selectorDeclared {
		return
	}
	selectorDeclared = true
	flag.StringVar(&selectorVar, "format.selector", "", "List of fields to do keep in output")
	flag.StringVar(&selectorTag, "format.tag", "", "Use format tag")
}

func ManualSelectorInit() error {
	if selectorVar == "" {
		return nil
	}
	selector = strings.Split(selectorVar, ",")
	return nil
}
