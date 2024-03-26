package main

import (
	"sort"
	"strings"
)

type list []*product

func (l list) Len() int { return len(l) }

func (l list) Less(i, j int) bool { return l[i].title < l[j].title }

func (l list) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

func (l list) String() string {
	if len(l) == 0 {
		return "Sorry. We're waiting for delivery 🚚."

	}

	var str strings.Builder

	for _, p := range l {
		str.WriteString("* ")
		str.WriteString(p.String())
		str.WriteRune('\n')
	}
	return str.String()

}

func (l list) discount(ratio float64) {
	for _, p := range l {
		p.discount(ratio)
	}
}

type byRelease struct {
	list
}

func (br byRelease) Less(i, j int) bool {
	return br.list[i].released.Before(br.list[j].released.Time)
}

func byReleaseDate(l list) sort.Interface {
	return &byRelease{l}
}
