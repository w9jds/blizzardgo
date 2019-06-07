package blizzard

import (
	"blizzard-webapi/pkg/blizzard/character"
	"fmt"
	"net/url"
	"strings"
)

// GetCharacterWithTypes returns all the information requested based on the passed in fields
func (api API) GetCharacterWithTypes(realm string, characterName string, fields ...character.FieldType) (*character.Character, error) {
	infoFields := make([]string, len(fields))

	for i := range fields {
		infoFields[i] = fields[i].String()
	}

	return api.GetCharacter(realm, characterName, infoFields...)
}

// GetCharacter gets the requested character with the string fields you request
func (api API) GetCharacter(realm string, characterName string, fields ...string) (*character.Character, error) {
	query := url.Values{
		"fields": {strings.Join(fields, " ")},
		"locale": {api.locale.String()},
	}

	request, err := api.newGet(fmt.Sprintf("/wow/character/%s/%s?%s", realm, characterName, query))
	if err != nil {
		return nil, err
	}

	var character *character.Character
	err = api.do(request, &character)
	if err != nil {
		return nil, err
	}

	return character, nil
}
