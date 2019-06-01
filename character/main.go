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

	Head     *Item `json:"head,omitempty"`
	Neck     *Item `json:"neck,omitempty"`
	Shoulder *Item `json:"shoulder,omitempty"`
	Back     *Item `json:"back,omitempty"`
	Chest    *Item `json:"chest,omitempty"`
	Wrist    *Item `json:"wrist,omitempty"`
	Hands    *Item `json:"hands,omitempty"`
	Waist    *Item `json:"waist,omitempty"`
	Legs     *Item `json:"legs,omitempty"`
	Feet     *Item `json:"feet,omitempty"`
	Finger1  *Item `json:"finger1,omitempty"`
	Finger2  *Item `json:"finder2,omitempty"`
	Trinket1 *Item `json:"trinket1,omitempty"`
	Trinket2 *Item `json:"trinket2,omitemtpy"`
	MainHand *Item `json:"mainHand,omitempty"`
}

type Item struct {

	// "mainHand": {
	//     "id": 163872,
	//     "name": "Honorbound War Staff",
	//     "icon": "inv_staff_2h_warfrontshorde_c_01",
	//     "quality": 4,
	//     "itemLevel": 370,
	//     "tooltipParams": {
	//         "timewalkerLevel": 120,
	//         "azeritePower0": 0,
	//         "azeritePower1": 0,
	//         "azeritePower2": 0,
	//         "azeritePower3": 0,
	//         "azeritePowerLevel": 0,
	//         "azeritePower4": 0
	//     },
	//     "stats": [
	//         {
	//             "stat": 40,
	//             "amount": 90
	//         },
	//         {
	//             "stat": 36,
	//             "amount": 140
	//         },
	//         {
	//             "stat": 5,
	//             "amount": 286
	//         },
	//         {
	//             "stat": 62,
	//             "amount": 99
	//         },
	//         {
	//             "stat": 7,
	//             "amount": 473
	//         }
	//     ],
	//     "armor": 0,
	//     "weaponInfo": {
	//         "damage": {
	//             "min": 329,
	//             "max": 445,
	//             "exactMin": 329,
	//             "exactMax": 445
	//         },
	//         "weaponSpeed": 3.6,
	//         "dps": 107.5
	//     },
	//     "context": "raid-mythic",
	//     "bonusLists": [
	//         5126,
	//         41,
	//         1562,
	//         4786
	//     ],
	//     "artifactId": 0,
	//     "displayInfoId": 168844,
	//     "artifactAppearanceId": 0,
	//     "artifactTraits": [],
	//     "relics": [],
	//     "appearance": {
	//         "itemAppearanceModId": 0
	//     },
	//     "azeriteItem": {
	//         "azeriteLevel": 0,
	//         "azeriteExperience": 0,
	//         "azeriteExperienceRemaining": 0
	//     },
	//     "azeriteEmpoweredItem": {
	//         "azeritePowers": []
	//     }
	// }
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
	Quests              []uint32            `json:"quests,omitempty"`
}
