package teamdefs

type Team struct {
	Name string
	Logo string
	Car  string
}

var (
	allTeams = []Team{
		{
			Name: "Aston Martin",
			Logo: "logo-aston-martin.png",
			Car:  "Aston Martin AMR24",
		},
		{
			Name: "Ferrari",
			Logo: "logo-ferrari.png",
			Car:  "Ferrari F1-75",
		},
		{
			Name: "McLaren",
			Logo: "logo-mclaren.png",
			Car:  "McLaren MCL38",
		},
		{
			Name: "Mercedes-AMG",
			Logo: "logo-mercedes-amg.png",
			Car:  "Mercedes-AMG F1 W11 EQ Performance",
		},
		{
			Name: "Red Bull-Oracle",
			Logo: "logo-red-bull-oracle.png",
			Car:  "Red Bull Racing RB16B",
		},
		{
			Name: "Williams-Martini",
			Logo: "logo-williams-martini.png",
			Car:  "Williams FW42",
		},
	}
)
