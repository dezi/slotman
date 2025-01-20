package pilots

import "embed"

//go:embed embeds/profile-dennis.jpg
//go:embed embeds/profile-kim.jpg
//go:embed embeds/profile-lukas.jpg
//go:embed embeds/profile-omar.jpg
//go:embed embeds/profile-patrick.jpg
//go:embed embeds/profile-susi.jpg
var embedFs embed.FS
