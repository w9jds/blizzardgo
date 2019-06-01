package character

// FieldType is an enum to specify what kind of info you want back about a character
type FieldType int

const (
	// FieldAchievements adds the achievements field to the character information response
	FieldAchievements FieldType = iota
	// FieldAppearance adds the appearance field to the character information response
	FieldAppearance FieldType = iota + 1
	// FieldNewsFeed adds the feed field to the character information response
	FieldNewsFeed FieldType = iota + 2
	// FieldGuild adds the guild field to the character information response
	FieldGuild FieldType = iota + 3
	// FieldHunterPets adds the hunterPets field to the character information response
	FieldHunterPets FieldType = iota + 4
	// FieldEquippedItems adds the items field to the character information response
	FieldEquippedItems FieldType = iota + 5
	// FieldMounts adds the mounts field to the character information response
	FieldMounts FieldType = iota + 6
	// FieldPets adds the pets field to the character information response
	FieldPets FieldType = iota + 7
	// FieldPetSlots adds the petSlots field to the character information response
	FieldPetSlots FieldType = iota + 8
	// FieldProfessions adds the professions field to the character information response
	FieldProfessions FieldType = iota + 9
	// FieldProgression adds the progression field to the character information response
	FieldProgression FieldType = iota + 10
	// FieldPVP adds the pvp field to the character information response
	FieldPVP FieldType = iota + 11
	// FieldCompletedQuests adds the quests field to the character information response
	FieldCompletedQuests FieldType = iota + 12
	// FieldReputation adds the reputation field to the character information response
	FieldReputation FieldType = iota + 13
	// FieldStatistics adds the statistics field to the character information response
	FieldStatistics FieldType = iota + 14
	// FieldStats adds the stats field to the character information response
	FieldStats FieldType = iota + 15
	// FieldTalents adds the talents field to the character information response
	FieldTalents FieldType = iota + 16
	// FieldTitles adds the titles field to the character information response
	FieldTitles FieldType = iota + 17
	// FieldAudit adds the audit field to the character information response
	FieldAudit FieldType = iota + 18
)

func (field FieldType) String() string {
	return [...]string{
		"achievements",
		"appearance",
		"feed",
		"guild",
		"hunterPets",
		"items",
		"mounts",
		"pets",
		"petSlots",
		"professions",
		"progression",
		"pvp",
		"quests",
		"reputation",
		"statistics",
		"stats",
		"talents",
		"titles",
		"audit",
	}[field]
}
