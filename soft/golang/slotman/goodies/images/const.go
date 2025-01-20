package images

type Team struct {
	Name string
	Logo string
	Car  string
}

var (
	allTeams = []Team{
		{
			Name: "Aston Martin",
			Logo: "team-aston-martin.png",
			Car:  "Aston Martin AMR24",
		},
		{
			Name: "McLaren",
			Logo: "team-mclaren.png",
			Car:  "McLaren MCL38",
		},
		{
			Name: "Mercedes-AMG",
			Logo: "team-mercedes-amg.png",
			Car:  "Mercedes-AMG F1 W11 EQ Performance",
		},
		{
			Name: "Ferrari",
			Logo: "team-ferrari.png",
			Car:  "Ferrari F1-75",
		},
		{
			Name: "Red Bull-Oracle",
			Logo: "team-red-bull-oracle.png",
			Car:  "Red Bull Racing RB16B",
		},
		{
			Name: "Williams-Martini",
			Logo: "team-williams-martini.png",
			Car:  "Williams FW42",
		},
	}
)
