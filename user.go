package blizzard

type UserCharacters struct {
	Characters []Character `json:"characters"`
}

type Character struct {
	Name              string        `json:"name"`
	Realm             string        `json:"realm"`
	Battlegroup       string        `json:"battlegroup"`
	Class             uint16        `json:"class"`
	Race              uint16        `json:"race"`
	Gender            uint16        `json:"gender"`
	Level             uint16        `json:"level"`
	AchievementPoints uint32        `json:"achievementPoints"`
	Thumbnail         string        `json:"thumbnail"`
	Spec              CharacterSpec `json:"spec,omitempty"`
	LastModified      uint64        `json:"lastModified"`
}

type CharacterSpec struct {
	Name            string `json:"name"`
	Role            string `json:"role"`
	BackgroundImage string `json:"backgroundImage"`
	Icon            string `json:"icon"`
	Description     string `json:"description"`
	Order           uint32 `json:"order"`
}

func (api API) GetUserCharacters(token AccessToken) {

}
