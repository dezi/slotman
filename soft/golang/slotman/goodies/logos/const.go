package logos

type Team struct {
	Name string
	Logo string
}

var (
	allTeams = []Team{
		{
			Name: "Ferrari",
			Logo: "team-ferrari.png",
		},
		{
			Name: "Red Bull-Oracle",
			Logo: "team-red-bull-oracle.png",
		},
		{
			Name: "Mercedes-AMG",
			Logo: "team-mercedes-amg.png",
		},
		{
			Name: "Williams-Martini",
			Logo: "team-williams-martini.png",
		},
		{
			Name: "Aston Martin",
			Logo: "team-aston-martin.png",
		},
	}
)
