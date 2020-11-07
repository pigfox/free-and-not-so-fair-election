package main

import (
	"fmt"

	"github.com/tatsushid/go-prettytable"
)

func main() {

	tbl, err := prettytable.NewTable([]prettytable.Column{
		{Header: "State"},
		{Header: "Population"},
		{Header: "ElectoralVotes"},
		{Header: "Electoral Votes Per Capita"},
	}...)

	if err != nil {
		panic(err)
	}

	tbl.Separator = "\t"

	for _, state := range states.List {
		electoralVotesPerCapita := fmt.Sprintf("%f", state.ElectoralVotes/state.Population)
		population := fmt.Sprintf("%v", int(state.Population))
		tbl.AddRow(state.Name, population, state.ElectoralVotes, electoralVotesPerCapita)
	}
	tbl.Print()
}
