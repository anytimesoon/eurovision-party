package migrations

import (
	"github.com/anytimesoon/eurovision-party/pkg/domain"
	"github.com/anytimesoon/eurovision-party/pkg/enum"
)

type initialCountryParams struct {
	Name          string
	Flag          string
	Slug          string
	Participating bool
}

type initialUserParams struct {
	Name    string
	Slug    string
	AuthLvl enum.AuthLvl
}

var initAuth = domain.Auth{}

var initUsers = []initialUserParams{
	{
		Name:    "Euro Host",
		Slug:    "admin",
		AuthLvl: enum.ADMIN,
	},
	{
		Name:    "EuroBot",
		Slug:    "bot",
		AuthLvl: enum.BOT,
	},
}

var initCountries = []initialCountryParams{
	{
		Name:          "Italy",
		Flag:          "🇮🇹",
		Slug:          "italy",
		Participating: true,
	},
	{
		Name:          "France",
		Flag:          "🇫🇷",
		Slug:          "france",
		Participating: true,
	},
	{
		Name:          "Germany",
		Flag:          "🇩🇪",
		Slug:          "germany",
		Participating: true,
	},
	{
		Name:          "Spain",
		Flag:          "🇪🇸",
		Slug:          "spain",
		Participating: true,
	},
	{
		Name:          "United Kingdom",
		Flag:          "🇬🇧",
		Slug:          "united-kingdom",
		Participating: true,
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
		Name: "Ukraine",
		Flag: "🇺🇦",
		Slug: "ukraine",
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
	{
		Name: "Russia",
		Flag: "🇷🇺",
		Slug: "russia",
	},
	{
		Name: "Turkey",
		Flag: "🇹🇷",
		Slug: "turkey",
	},
	{
		Name: "Bosnia and Herzegovina",
		Flag: "🇧🇦",
		Slug: "bosnia-and-herzegovina",
	},
	{
		Name: "Hungary",
		Flag: "🇧🇭",
		Slug: "hungary",
	},
	{
		Name: "Luxembourg",
		Flag: "🇱🇺",
		Slug: "luxembourg",
	},
}
