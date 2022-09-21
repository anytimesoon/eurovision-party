package migrations

import "eurovision/pkg/domain"

type initialCountryParams struct {
	Name string
	Flag string
	Slug string
}

type initialUserParams struct {
	Name    string
	Email   string
	Slug    string
	AuthLvl uint8
}

var initAuth = domain.Auth{}

var initUsers = []initialUserParams{
	{
		Name:    "admin",
		Email:   "admin@admin.com",
		Slug:    "admin",
		AuthLvl: 1,
	},
	{
		Name:    "bot",
		Email:   "bot@bot.com",
		Slug:    "bot",
		AuthLvl: 2,
	},
}

var initCountries = []initialCountryParams{
	{
		Name: "Italy",
		Flag: "🇮🇹",
		Slug: "italy",
	},
	{
		Name: "France",
		Flag: "🇫🇷",
		Slug: "france",
	},
	{
		Name: "Germany",
		Flag: "🇩🇪",
		Slug: "germany",
	},
	{
		Name: "Spain",
		Flag: "🇪🇸",
		Slug: "spain",
	},
	{
		Name: "United Kingdom",
		Flag: "🇬🇧",
		Slug: "united-kingdom",
	},
	{
		Name: "Albania",
		Flag: "🇦🇱",
		Slug: "albania",
	},
	{
		Name: "Latvia",
		Flag: "🇱🇻",
		Slug: "latvia",
	},
	{
		Name: "Lithuania",
		Flag: "🇱🇹",
		Slug: "lithuania",
	},
	{
		Name: "Switzerland",
		Flag: "🇨🇭",
		Slug: "switzerland",
	},
	{
		Name: "Slovenia",
		Flag: "🇸🇮",
		Slug: "slovenia",
	},
	{
		Name: "Ukrain",
		Flag: "🇺🇦",
		Slug: "ukrain",
	},
	{
		Name: "Bulgaria",
		Flag: "🇧🇬",
		Slug: "bulgaria",
	},
	{
		Name: "Netherlands",
		Flag: "🇳🇱",
		Slug: "netherlands",
	},
	{
		Name: "Moldova",
		Flag: "🇲🇩",
		Slug: "moldova",
	},
	{
		Name: "Portugal",
		Flag: "🇵🇹",
		Slug: "portugal",
	},
	{
		Name: "Croatia",
		Flag: "🇭🇷",
		Slug: "croatia",
	},
	{
		Name: "Denmark",
		Flag: "🇩🇰",
		Slug: "denmark",
	},
	{
		Name: "Austria",
		Flag: "🇦🇹",
		Slug: "austria",
	},
	{
		Name: "Iceland",
		Flag: "🇮🇸",
		Slug: "iceland",
	},
	{
		Name: "Greece",
		Flag: "🇬🇷",
		Slug: "greece",
	},
	{
		Name: "Norway",
		Flag: "🇳🇴",
		Slug: "norway",
	},
	{
		Name: "Armenia",
		Flag: "🇦🇲",
		Slug: "armenia",
	},
	{
		Name: "Finland",
		Flag: "🇫🇮",
		Slug: "finland",
	},
	{
		Name: "Israel",
		Flag: "🇮🇱",
		Slug: "israel",
	},
	{
		Name: "Serbia",
		Flag: "🇷🇸",
		Slug: "serbia",
	},
	{
		Name: "Azerbaijan",
		Flag: "🇦🇿",
		Slug: "azerbaijan",
	},
	{
		Name: "Georgia",
		Flag: "🇬🇪",
		Slug: "georgia",
	},
	{
		Name: "Malta",
		Flag: "🇲🇹",
		Slug: "malta",
	},
	{
		Name: "San Marino",
		Flag: "🇸🇲",
		Slug: "san-marino",
	},
	{
		Name: "Australia",
		Flag: "🇦🇺",
		Slug: "australia",
	},
	{
		Name: "Cyprus",
		Flag: "🇨🇾",
		Slug: "cyprus",
	},
	{
		Name: "Ireland",
		Flag: "🇮🇪",
		Slug: "ireland",
	},
	{
		Name: "North Macedonia",
		Flag: "🇲🇰",
		Slug: "north-macedonia",
	},
	{
		Name: "Estonia",
		Flag: "🇪🇪",
		Slug: "estonia",
	},
	{
		Name: "Romania",
		Flag: "🇷🇴",
		Slug: "romania",
	},
	{
		Name: "Poland",
		Flag: "🇵🇱",
		Slug: "poland",
	},
	{
		Name: "Montenegro",
		Flag: "🇲🇪",
		Slug: "montenegro",
	},
	{
		Name: "Belgium",
		Flag: "🇧🇪",
		Slug: "belgium",
	},
	{
		Name: "Sweden",
		Flag: "🇸🇪",
		Slug: "sweden",
	},
	{
		Name: "Czech Republic",
		Flag: "🇨🇿",
		Slug: "czech-republic",
	},
}
