package seeder

import "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/model"

var Categories = []*model.Category{
	{
		Name:  "Elektronik",
		Level: 1,
	},
	{
		Name:   "TV & Aksesoris",
		Level:  2,
		Parent: 1,
	},
	{
		Name:   "TV",
		Level:  3,
		Parent: 2,
	},
	{
		Name:   "Kelistrikan",
		Level:  2,
		Parent: 1,
	},
	{
		Name:   "Saklar",
		Level:  3,
		Parent: 2,
	},
}
