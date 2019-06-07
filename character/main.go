package character

import (
	"blizzard-webapi/pkg/timestamp"
)

type NewsFeed struct {
}

type Guild struct {
	Name              string      `json:"name"`
	Realm             string      `json:"realm"`
	Battlegroup       string      `json:"battlegroup"`
	Members           uint32      `json:"members"`
	AchievementPoints uint32      `json:"achievementPoints"`
	Emblem            GuildEmblem `json:"emblem"`
}

type GuildEmblem struct {
	Icon              uint32 `json:"icon"`
	IconColor         string `json:"iconColor"`
	IconColorID       uint32 `json:"iconColorId"`
	Border            uint32 `json:"border"`
	BorderColor       string `json:"borderColor"`
	BorderColorID     uint32 `json:"borderColorId"`
	BackgroundColor   string `json:"backgroundColor"`
	BackgroundColorID uint32 `json:"backgroundColorId"`
}

type Achievements struct {
	CompletedIDs        []uint32              `json:"achievementsCompleted"`
	CriteriaIDs         []uint32              `json:"criteria"`
	CriteriaQuantity    []uint64              `json:"criteriaQuantity"`
	CompletedTimestamps []timestamp.Timestamp `json:"achievementsCompletedTimestamp"`
	CriteriaTimestamp   []timestamp.Timestamp `json:"criteriaTimestamp"`
	CriteriaCreated     []timestamp.Timestamp `json:"criteriaCreated"`
}

type Appearance struct {
	FaceVariation    uint16 `json:"faceVariation"`
	SkinColor        uint16 `json:"skinColor"`
	HairVariation    uint16 `json:"hairVariation"`
	HairColor        uint16 `json:"hairColor"`
	FeatureVariation uint16 `json:"featureVariation"`
	ShowHelm         bool   `json:"showHelm"`
	ShowCloak        bool   `json:"showCloak"`

	CustomDisplayOptions []uint32 `json:"customDisplayOptions"`
}

type HunterPets struct {
}

type EquippedItems struct {
	AverageItemLevel         uint16 `json:"averageItemLevel"`
	AverageItemLevelEquipped uint32 `json:"averageItemLevelEquipped"`

	Head     *EquippedItem `json:"head,omitempty"`
	Neck     *EquippedItem `json:"neck,omitempty"`
	Shoulder *EquippedItem `json:"shoulder,omitempty"`
	Back     *EquippedItem `json:"back,omitempty"`
	Chest    *EquippedItem `json:"chest,omitempty"`
	Wrist    *EquippedItem `json:"wrist,omitempty"`
	Hands    *EquippedItem `json:"hands,omitempty"`
	Waist    *EquippedItem `json:"waist,omitempty"`
	Legs     *EquippedItem `json:"legs,omitempty"`
	Feet     *EquippedItem `json:"feet,omitempty"`
	Finger1  *EquippedItem `json:"finger1,omitempty"`
	Finger2  *EquippedItem `json:"finder2,omitempty"`
	Trinket1 *EquippedItem `json:"trinket1,omitempty"`
	Trinket2 *EquippedItem `json:"trinket2,omitemtpy"`
	MainHand *EquippedItem `json:"mainHand,omitempty"`
}

type EquippedItem struct {
}

type Stat struct {
	Stat   uint32 `json:"stat"`
	Amount uint32 `json:"amount"`
}

type Character struct {
	LastModified        timestamp.Timestamp `json:"lastModified"`
	Name                string              `json:"name"`
	Realm               string              `json:"realm"`
	Battlegroup         string              `json:"battlegroup"`
	Class               uint16              `json:"class"`
	Race                uint16              `json:"race"`
	Gender              uint8               `json:"gender"`
	Level               uint8               `json:"level"`
	AchievementPoints   uint32              `json:"achievementPoints"`
	Thumbnail           string              `json:"thumbnail"`
	CalcClass           string              `json:"calcClass"`
	Faction             uint8               `json:"faction"`
	TotalHonorableKills uint64              `json:"totalHonorableKills"`
	Achievements        *Achievements       `json:"achievements,omitempty"`
	Appearance          *Appearance         `json:"appearance,omitempty"`
	NewsFeed            *NewsFeed           `json:"feed,omitempty"`
	Guild               *Guild              `json:"guild,omitempty"`
	HunterPets          *HunterPets         `json:"hunterPets,omitempty"`
	EquippedItems       *EquippedItems      `json:"items,omitempty"`
	Mounts              *interface{}        `json:"mounts,omitempty"`
	Pets                *interface{}        `json:"pets,omitempty"`
	PetSlots            *interface{}        `json:"petSlots,omitempty"`
	Professions         *interface{}        `json:"professions,omitempty"`
	Progression         *interface{}        `json:"progression,omitempty"`
	PVP                 *interface{}        `json:"pvp,omitempty"`
	Quests              []uint32            `json:"quests,omitempty"`
	Reputation          *interface{}        `json:"reputation,omitempty"`
	Statistics          *interface{}        `json:"statistics,omitempty"`
	Stats               *interface{}        `json:"stats,omitempty"`
	Talents             *interface{}        `json:"talents,omitempty"`
	Titles              *interface{}        `json:"titles,omitempty"`
	Audit               *interface{}        `json:"audit,omitempty"`
}
