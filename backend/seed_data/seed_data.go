package seed_data

import (
	"ScavengerHunt/backend/models"
)

var UsersSeed = []models.User{
	{UserID: 1, Password: "asdflkj", Email: "twerasdfs@gmail.com", Username: "Horatio"},
	{UserID: 2, Password: "asdflkj", Email: "clarlasdfs@gmail.com", Username: "Mercedes"},
	{UserID: 3, Password: "asdflkj", Email: "t2wasdfasdfls@gmail.com", Username: "Drakmar"},
	{UserID: 4, Password: "asdflkj", Email: "garasdfble@gmail.com", Username: "Groucho"},
}

var TeamsSeed = []models.Team{
	{
		TeamID: 1,
		Name:   "Apricot Apricot",
		HuntID: 1,
	},
	{
		TeamID: 2,
		Name:   "Upricot Downricot",
		HuntID: 1,
	},
	{
		TeamID: 3,
		Name:   "Leftricot Rightricot",
		HuntID: 1,
	},
}

var ScavengerHuntSeed = []models.Hunt{
	{
		HuntID:      1,
		Title:       "Gilded Age",
		MaxTeamSize: 3,
		CreatedBy:   1,
		Description: "Blurb blurb blurb",
	},
	{
		HuntID:      2,
		Title:       "Silvered Age",
		MaxTeamSize: 3,
		CreatedBy:   1,
		Description: "Blurb blurb blurb",
	},
	{
		HuntID:      3,
		Title:       "Bronzered Age",
		MaxTeamSize: 4,
		CreatedBy:   2,
		Description: "Blurb blurb blurb",
	},
	{
		HuntID:      4,
		Title:       "Gilded Age II: The Return",
		MaxTeamSize: 3,
		CreatedBy:   1,
		Description: "Blurb blurb blurb",
	},
	{
		HuntID:      5,
		Title:       "Silvered Age: Big Screen",
		MaxTeamSize: 3,
		CreatedBy:   1,
		Description: "Blurb blurb blurb",
	},
	{
		HuntID:      6,
		Title:       "Bronzered Age II: Back in the Habit",
		MaxTeamSize: 4,
		CreatedBy:   2,
		Description: "Blurb blurb blurb",
	},
	{
		HuntID:      7,
		Title:       "Gilded Cinematic Universe",
		MaxTeamSize: 3,
		CreatedBy:   3,
		Description: "Blurb blurb blurb",
	},
}

var ScavengerHuntCluesSeed = []models.Clue{
	{
		ClueID:         1,
		Name:           "Brooklyn Bridge",
		Text:           "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Venenatis urna cursus eget nunc. Non blandit massa enim nec dui nunc. Senectus et netus et malesuada fames. Id porta nibh venenatis cras. Sit amet massa vitae tortor condimentum. Hac habitasse platea dictumst quisque sagittis purus sit. Tincidunt ornare massa eget egestas purus viverra accumsan in nisl. Hac habitasse platea dictumst vestibulum rhoncus. Egestas tellus rutrum tellus pellentesque eu.",
		HuntID:         1,
		Value:          20,
		MaxSubmissions: 3,
	},

	{
		ClueID:         2,
		Name:           "Manhattan Bridge",
		Text:           "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Venenatis urna cursus eget nunc. Non blandit massa enim nec dui nunc. Senectus et netus et malesuada fames. Id porta nibh venenatis cras. Sit amet massa vitae tortor condimentum. Hac habitasse platea dictumst quisque sagittis purus sit. Tincidunt ornare massa eget egestas purus viverra accumsan in nisl. Hac habitasse platea dictumst vestibulum rhoncus. Egestas tellus rutrum tellus pellentesque eu.",
		HuntID:         1,
		Value:          20,
		MaxSubmissions: 3,
	},
	{
		ClueID:         3,
		Name:           "Bridge to Terabithia",
		Text:           "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Venenatis urna cursus eget nunc. Non blandit massa enim nec dui nunc. Senectus et netus et malesuada fames. Id porta nibh venenatis cras. Sit amet massa vitae tortor condimentum. Hac habitasse platea dictumst quisque sagittis purus sit. Tincidunt ornare massa eget egestas purus viverra accumsan in nisl. Hac habitasse platea dictumst vestibulum rhoncus. Egestas tellus rutrum tellus pellentesque eu.",
		HuntID:         1,
		Value:          20,
		MaxSubmissions: 3,
	},
	{
		ClueID:         4,
		Name:           "London Bridge",
		Text:           "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Venenatis urna cursus eget nunc. Non blandit massa enim nec dui nunc. Senectus et netus et malesuada fames. Id porta nibh venenatis cras. Sit amet massa vitae tortor condimentum. Hac habitasse platea dictumst quisque sagittis purus sit. Tincidunt ornare massa eget egestas purus viverra accumsan in nisl. Hac habitasse platea dictumst vestibulum rhoncus. Egestas tellus rutrum tellus pellentesque eu.",
		HuntID:         1,
		Value:          20,
		MaxSubmissions: 3,
	},
	{
		ClueID:         5,
		Name:           "II Bridge II London",
		Text:           "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Venenatis urna cursus eget nunc. Non blandit massa enim nec dui nunc. Senectus et netus et malesuada fames. Id porta nibh venenatis cras. Sit amet massa vitae tortor condimentum. Hac habitasse platea dictumst quisque sagittis purus sit. Tincidunt ornare massa eget egestas purus viverra accumsan in nisl. Hac habitasse platea dictumst vestibulum rhoncus. Egestas tellus rutrum tellus pellentesque eu.",
		HuntID:         2,
		Value:          20,
		MaxSubmissions: 3,
	},
	{
		ClueID:         6,
		Name:           "Sky Bridge",
		Text:           "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Venenatis urna cursus eget nunc. Non blandit massa enim nec dui nunc. Senectus et netus et malesuada fames. Id porta nibh venenatis cras. Sit amet massa vitae tortor condimentum. Hac habitasse platea dictumst quisque sagittis purus sit. Tincidunt ornare massa eget egestas purus viverra accumsan in nisl. Hac habitasse platea dictumst vestibulum rhoncus. Egestas tellus rutrum tellus pellentesque eu.",
		HuntID:         2,
		Value:          5,
		MaxSubmissions: 3,
	},
	{
		ClueID:         7,
		Name:           "Bifrost",
		Text:           "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Venenatis urna cursus eget nunc. Non blandit massa enim nec dui nunc. Senectus et netus et malesuada fames. Id porta nibh venenatis cras. Sit amet massa vitae tortor condimentum. Hac habitasse platea dictumst quisque sagittis purus sit. Tincidunt ornare massa eget egestas purus viverra accumsan in nisl. Hac habitasse platea dictumst vestibulum rhoncus. Egestas tellus rutrum tellus pellentesque eu.",
		HuntID:         2,
		Value:          15,
		MaxSubmissions: 3,
	},
}
