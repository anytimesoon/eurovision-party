package initialData

import (
	domain "eurovision/pkg/domain"

	_ "github.com/go-sql-driver/mysql"
)

var InitCountries = []domain.Country{
	{
		Name: "Italy",
		Flag: "🇮🇹",
	},
	{
		Name: "France",
		Flag: "🇫🇷",
	},
	{
		Name: "Germany",
		Flag: "🇩🇪",
	},
	{
		Name: "Spain",
		Flag: "🇪🇸",
	},
	{
		Name: "United Kingdom",
		Flag: "🇬🇧",
	},
	{
		Name: "Albania",
		Flag: "🇦🇱",
	},
	{
		Name: "Latvia",
		Flag: "🇱🇻",
	},
	{
		Name: "Lithuania",
		Flag: "🇱🇹",
	},
	{
		Name: "Switzerland",
		Flag: "🇨🇭",
	},
	{
		Name: "Slovenia",
		Flag: "🇸🇮",
	},
	{
		Name: "Ukrain",
		Flag: "🇺🇦",
	},
	{
		Name: "Bulgaria",
		Flag: "🇧🇬",
	},
	{
		Name: "Netherlands",
		Flag: "🇳🇱",
	},
	{
		Name: "Moldova",
		Flag: "🇲🇩",
	},
	{
		Name: "Portugal",
		Flag: "🇵🇹",
	},
	{
		Name: "Croatia",
		Flag: "🇭🇷",
	},
	{
		Name: "Denmark",
		Flag: "🇩🇰",
	},
	{
		Name: "Austria",
		Flag: "🇦🇹",
	},
	{
		Name: "Iceland",
		Flag: "🇮🇸",
	},
	{
		Name: "Greece",
		Flag: "🇬🇷",
	},
	{
		Name: "Norway",
		Flag: "🇳🇴",
	},
	{
		Name: "Armenia",
		Flag: "🇦🇲",
	},
	{
		Name: "Finland",
		Flag: "🇫🇮",
	},
	{
		Name: "Israel",
		Flag: "🇮🇱",
	},
	{
		Name: "Serbia",
		Flag: "🇷🇸",
	},
	{
		Name: "Azerbaijan",
		Flag: "🇦🇿",
	},
	{
		Name: "Georgia",
		Flag: "🇬🇪",
	},
	{
		Name: "Malta",
		Flag: "🇲🇹",
	},
	{
		Name: "San Marino",
		Flag: "🇸🇲",
	},
	{
		Name: "Australia",
		Flag: "🇦🇺",
	},
	{
		Name: "Cyprus",
		Flag: "🇨🇾",
	},
	{
		Name: "Ireland",
		Flag: "🇮🇪",
	},
	{
		Name: "North Macedonia",
		Flag: "🇲🇰",
	},
	{
		Name: "Estonia",
		Flag: "🇪🇪",
	},
	{
		Name: "Romania",
		Flag: "🇷🇴",
	},
	{
		Name: "Poland",
		Flag: "🇵🇱",
	},
	{
		Name: "Montenegro",
		Flag: "🇲🇪",
	},
	{
		Name: "Belgium",
		Flag: "🇧🇪",
	},
	{
		Name: "Sweden",
		Flag: "🇸🇪",
	},
	{
		Name: "Czech Republic",
		Flag: "🇨🇿",
	},
}

const (
	Username string = "eurovision"
	Password string = "P,PO)+{l4!C{ff"
	Hostname string = "127.0.0.1:3306"
	DBName   string = "eurovision"
)
