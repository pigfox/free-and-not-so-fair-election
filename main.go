package main

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
)

func main() {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"State", "Population", "ElectoralVotes", "Electoral Votes Per Capita"})

	rows := len(states.List)
	counter := 0
	for _, state := range states.List {
		electoralVotesPerCapita := fmt.Sprintf("%f", state.ElectoralVotes/state.Population)
		population := fmt.Sprintf("%v", int(state.Population))
		t.AppendRow([]interface{}{state.Name, population, state.ElectoralVotes, electoralVotesPerCapita})

		if counter < rows-1 {
			t.AppendRow([]interface{}{"---------------------", "-----------", "--------------", "---------------------------"})
		}
		counter++
	}

	t.Render()
}
