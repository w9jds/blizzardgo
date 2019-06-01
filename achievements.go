package blizzard

type Key struct {
	href string
}

type Index struct {
	key  Key
	id   uint32
	name string
}

func (api API) GetAchievementCategoryIndexes() {
	// path := "/data/wow/achievement-category/"
}

func (api API) GetAchievementIndexes() {
	// path := "/data/wow/achievement-category/%d"
}

func (api API) GetCategory(id uint32) {
	// path := "/data/wow/achievement/index"
}

func (api API) GetAchievement(id uint32) {
	// path := "/data/wow/achievement/%d"
}
