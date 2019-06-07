package blizzard

type Key struct {
	href string
}

type Index struct {
	key  Key
	id   uint32
	name string
}

// func (api API) GetAchievementCategoryIndexes() {
// 	request, err := api.newGet("/data/wow/achievement-category")
// 	if err != nil {
// 		return nil, err
// 	}
// }

// func (api API) GetAchievementIndexes(categoryID uint32) {
// 	request, err := api.newGet(fmt.Sprintf("/data/wow/achievement-category/%d", categoryID))
// 	if err != nil {
// 		return nil, err
// 	}
// }

// func (api API) GetCategory() {
// 	request, err := api.newGet("/data/wow/achievement/index")
// 	if err != nil {
// 		return nil, err
// 	}
// }

// func (api API) GetAchievement(achievementID uint32) {
// 	request, err := api.newGet(fmt.Sprintf("/data/wow/achievement/%d", achievementID))
// 	if err != nil {
// 		return nil, err
// 	}
// }
