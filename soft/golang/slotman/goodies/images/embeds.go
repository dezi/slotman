package images

import "embed"

//go:embed embeds/team-aston-martin.png
//go:embed embeds/team-ferrari.png
//go:embed embeds/team-mclaren.png
//go:embed embeds/team-mercedes-amg.png
//go:embed embeds/team-red-bull-oracle.png
//go:embed embeds/team-williams-martini.png
var embedFs embed.FS
