package router

import (
	"poe-debugserver/src/model"
)

//this place gives different type of sample for serveral request.
//using io to load from file?  yeah, it's feasible, but lazy :P
//I still keep place to load file, try replace variables instead of default.
func sandbox() {
	// load file here
}

var currentLeague = []model.League{
	{
		ID:          `Standard`,
		Realm:       `pc`,
		URL:         `https:\/\/www.pathofexile.com\/forum\/view-thread\/71278`,
		Startat:     `2013-01-23T21:00:00Z`,
		Endat:       `null`,
		Description: `The default game mode.`,
		Registerat:  `2019-09-06T19:00:00Z`,
		Delveevent:  true,
		Rules:       []model.Rule{},
	},
	{
		ID:          `Hardcore`,
		Realm:       `pc`,
		URL:         `https:\/\/www.pathofexile.com\/forum\/view-thread\/71276`,
		Startat:     `2013-01-23T21:00:00Z`,
		Endat:       `null`,
		Description: `A character killed in the Hardcore league is moved to the Standard league.`,
		Registerat:  `2019-09-06T19:00:00Z`,
		Delveevent:  true,
		Rules: []model.Rule{
			{
				ID:          `Hardcore`,
				Name:        `Hardcore`,
				Description: `A character killed in Hardcore is moved to its parent league.`,
			},
		},
	},
	{
		ID:          `SSF Standard`,
		Realm:       `pc`,
		URL:         `https:\/\/www.pathofexile.com\/forum\/view-thread\/1841357`,
		Startat:     `2013-01-23T21:00:00Z`,
		Endat:       `null`,
		Description: `SSF Standard.`,
		Registerat:  `2019-09-06T19:00:00Z`,
		Delveevent:  true,
		Rules: []model.Rule{
			{
				ID:          `NoParties`,
				Name:        `Solo`,
				Description: `You may not party in this league.`,
			},
		},
	},
	{
		ID:          `SSF Hardcore`,
		Realm:       `pc`,
		URL:         `https:\/\/www.pathofexile.com\/forum\/view-thread\/1841353`,
		Startat:     `2013-01-23T21:00:00Z`,
		Endat:       `null`,
		Description: `SSF Hardcore`,
		Registerat:  `2019-09-06T19:00:00Z`,
		Delveevent:  true,
		Rules: []model.Rule{
			{
				ID:          `Hardcore`,
				Name:        `Hardcore`,
				Description: `A character killed in Hardcore is moved to its parent league.`,
			},
			{
				ID:          `NoParties`,
				Name:        `Solo`,
				Description: `You may not party in this league.`,
			},
		},
	},
	{
		ID:          `Ultimatum`,
		Realm:       `pc`,
		URL:         `https:\/\/www.pathofexile.com\/forum\/view-thread\/3080795`,
		Startat:     `2013-01-23T21:00:00Z`,
		Endat:       `2025-12-07T21:00:00Z`,
		Description: `Encounter an emissary of the Vaal entity Chaos and decide whether to risk it all.\n\nThis is the default Path of Exile league.`,
		Registerat:  `2021-04-16T17:30:00Z`,
		Delveevent:  true,
		Rules:       []model.Rule{},
	},
	{
		ID:          `Hardcore Ultimatum`,
		Realm:       `pc`,
		URL:         `https:\/\/www.pathofexile.com\/forum\/view-thread\/3080803`,
		Startat:     `2013-01-23T21:00:00Z`,
		Endat:       `2025-12-07T21:00:00Z`,
		Description: `Encounter an emissary of the Vaal entity Chaos and decide whether to risk it all.\n\nA character killed in Hardcore Ultimatum becomes a Standard character.`,
		Registerat:  `2019-09-06T19:00:00Z`,
		Delveevent:  true,
		Rules: []model.Rule{
			{
				ID:          `Hardcore`,
				Name:        `Hardcore`,
				Description: `A character killed in Hardcore is moved to its parent league.`,
			},
		},
	},
	{
		ID:          `SSF Ultimatum`,
		Realm:       `pc`,
		URL:         `https:\/\/www.pathofexile.com\/forum\/view-thread\/3080811`,
		Startat:     `2013-01-23T21:00:00Z`,
		Endat:       `2025-12-07T21:00:00Z`,
		Description: `SSF Ultimatum`,
		Registerat:  `2021-04-16T17:30:00Z`,
		Delveevent:  true,
		Rules: []model.Rule{
			{
				ID:          `NoParties`,
				Name:        `Solo`,
				Description: `You may not party in this league.`,
			},
		},
	},
	{
		ID:          `SSF Ultimatum HC`,
		Realm:       `pc`,
		URL:         `https:\/\/www.pathofexile.com\/forum\/view-thread\/3080819`,
		Startat:     `2021-04-16T20:00:00Z`,
		Endat:       `2025-12-07T21:00:00Z`,
		Description: `SSF Ultimatum HC`,
		Registerat:  `2021-04-16T17:30:00Z`,
		Delveevent:  true,
		Rules: []model.Rule{
			{
				ID:          `Hardcore`,
				Name:        `Hardcore`,
				Description: `A character killed in Hardcore is moved to its parent league.`,
			},
			{
				ID:          `NoParties`,
				Name:        `Solo`,
				Description: `You may not party in this league.`,
			},
		},
	},
}
