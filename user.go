package blizzard

// UserCharacters is the response to /wow/user/characters listing your characters
type UserCharacters struct {
	Characters []Character `json:"characters"`
}

// Character is a struct base for each character that you own
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

// CharacterSpec is the information about that character's currently select spec
type CharacterSpec struct {
	Name            string `json:"name"`
	Role            string `json:"role"`
	BackgroundImage string `json:"backgroundImage"`
	Icon            string `json:"icon"`
	Description     string `json:"description"`
	Order           uint32 `json:"order"`
}

// GetUserCharacters returns a list of all characters on the token's account
func (api API) GetUserCharacters(token AccessToken) (*UserCharacters, error) {
	request, err := api.newGet("/wow/user/characters")
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+token.AccessToken)
	var characters *UserCharacters
	err = api.do(request, err)
	if err != nil {
		return nil, err
	}

	return characters, nil
}
