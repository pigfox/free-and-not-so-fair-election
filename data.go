/*
Population data from:
https://en.wikipedia.org/wiki/List_of_states_and_territories_of_the_United_States_by_population

Number of eletoral votes from:
https://state.1keydata.com/state-electoral-votes.php
*/

package main

type States struct {
	List []State
}

type State struct {
	Name           string
	Population     float64
	ElectoralVotes float64
}

var states States

func init() {
	states.List = []State{
		{Name: "Alabama", Population: 4903185, ElectoralVotes: 9},
		{Name: "Alaska", Population: 731545, ElectoralVotes: 3},
		{Name: "Arizona", Population: 7278717, ElectoralVotes: 11},
		{Name: "Arkansas", Population: 3017804, ElectoralVotes: 6},
		{Name: "California", Population: 39512223, ElectoralVotes: 55},
		{Name: "Colorado", Population: 5758736, ElectoralVotes: 9},
		{Name: "Connecticut", Population: 3565287, ElectoralVotes: 7},
		{Name: "Delaware", Population: 973764, ElectoralVotes: 3},
		{Name: "Florida", Population: 21477737, ElectoralVotes: 29},
		{Name: "Georgia", Population: 10617423, ElectoralVotes: 16},
		{Name: "Hawaii", Population: 1415872, ElectoralVotes: 4},
		{Name: "Idaho", Population: 1787065, ElectoralVotes: 4},
		{Name: "Illinois", Population: 12671821, ElectoralVotes: 20},
		{Name: "Indiana", Population: 6732219, ElectoralVotes: 11},
		{Name: "Iowa", Population: 3155070, ElectoralVotes: 6},
		{Name: "Kansas", Population: 2913314, ElectoralVotes: 6},
		{Name: "Kentucky", Population: 4467673, ElectoralVotes: 8},
		{Name: "Louisiana", Population: 4648794, ElectoralVotes: 8},
		{Name: "Maine", Population: 1344212, ElectoralVotes: 4},
		{Name: "Maryland", Population: 6045680, ElectoralVotes: 10},
		{Name: "Massachusetts", Population: 6892503, ElectoralVotes: 11},
		{Name: "Michigan", Population: 9986857, ElectoralVotes: 16},
		{Name: "Minnesota", Population: 5639632, ElectoralVotes: 10},
		{Name: "Mississippi", Population: 2976149, ElectoralVotes: 6},
		{Name: "Missouri", Population: 6137428, ElectoralVotes: 10},
		{Name: "Montana", Population: 1068778, ElectoralVotes: 3},
		{Name: "Nebraska", Population: 1934408, ElectoralVotes: 5},
		{Name: "Nevada", Population: 3080156, ElectoralVotes: 6},
		{Name: "New Hampshire", Population: 1359711, ElectoralVotes: 4},
		{Name: "New Jersey", Population: 8882190, ElectoralVotes: 14},
		{Name: "New Mexico", Population: 2096829, ElectoralVotes: 5},
		{Name: "New York", Population: 19453561, ElectoralVotes: 29},
		{Name: "North Carolina", Population: 10488084, ElectoralVotes: 15},
		{Name: "North Dakota", Population: 762062, ElectoralVotes: 3},
		{Name: "Ohio", Population: 11689100, ElectoralVotes: 18},
		{Name: "Oklahoma", Population: 3956971, ElectoralVotes: 7},
		{Name: "Oregon", Population: 4217737, ElectoralVotes: 7},
		{Name: "Pennsylvania", Population: 12801989, ElectoralVotes: 20},
		{Name: "Rhode Island", Population: 1059361, ElectoralVotes: 4},
		{Name: "South Carolina", Population: 5148714, ElectoralVotes: 9},
		{Name: "South Dakota", Population: 884659, ElectoralVotes: 3},
		{Name: "Tennessee", Population: 6829174, ElectoralVotes: 11},
		{Name: "Texas", Population: 28995881, ElectoralVotes: 38},
		{Name: "Utah", Population: 3205958, ElectoralVotes: 6},
		{Name: "Vermont", Population: 623989, ElectoralVotes: 3},
		{Name: "Virginia", Population: 8535519, ElectoralVotes: 13},
		{Name: "Wasington", Population: 7614893, ElectoralVotes: 12},
		{Name: "West Virginia", Population: 1792147, ElectoralVotes: 5},
		{Name: "Wisconsin", Population: 5822434, ElectoralVotes: 10},
		{Name: "Wyoming", Population: 578759, ElectoralVotes: 3},
	}
}
