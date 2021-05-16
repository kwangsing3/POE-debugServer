package model

type Items struct {
	Verified bool
	Name     string `json:"name"`
	TypeLine string `json:"typeline"`
	Ilvl     int16  `json:"ilvl"`
	League   string `json:"league"`
	//
	W            uint16 `json:"w"`
	H            uint16 `json:"h"`
	Icon         string `json:"icon"`
	StackSize    int16  `json:"stacksize"`
	MaxStackSize int16  `json:"maxstacksize"`
	//
	IsIdentified  bool           `json:"isidentified"`
	IsCorrupted   bool           `json:"iscorrupted"`
	FrameType     int8           `json:"frameType"`
	Sockets       []Socket       `json:"sockets"`
	SocketedItems []SocketedItem `json:"socketedItems"`
	InventoryId   string         `json:"inventoryId"`
	//
	IsShaper   bool `json:"shaper"`
	IsElder    bool `json:"elder"`
	IsCrusader bool `json:"crusader"`
	IsRedeemer bool `json:"redeemer"`
	IsHunter   bool `json:"hunter"`
	IsWarlord  bool `json:"warload"`
	//
	//CosmeticMods
	CraftedMods  []string `json:"craftedMods"`
	EnchantMods  []string `json:"enchantMods"`
	ImplicitMods []string `json:"implicitMods"`
	ExplicitMods []string `json:"explicitMods"`
	UtilityMods  []string `json:"utilityMods"`
}
type Socket struct {
	Group   int8   `json:"group"`   // index, the same group index means there're linked.
	Attr    string `json:"attr"`    // S、D、I(str、dex、int)
	SColour string `json:"scolour"` // R、G、B(str、dex、int)
}
type SocketedItem struct {
	Verified     bool
	Name         string       `json:"name"`
	TypeLine     string       `json:"typeline"`
	Ilvl         int16        `json:"ilvl"`
	League       string       `json:"league"`
	IsSupport    bool         `json:"issupport"`
	Properties   []Properties `json:"properties"`
	Requirements []Properties `json:"requirements"`
	//
	W            uint16 `json:"w"`
	H            uint16 `json:"h"`
	Icon         string `json:"icon"`
	StackSize    int16  `json:"stacksize"`
	MaxStackSize int16  `json:"maxstacksize"`
	//
	IsIdentified bool   `json:"isidentified"`
	IsCorrupted  bool   `json:"iscorrupted"`
	FrameType    int8   `json:"frameType"`
	SocketID     int8   `json:"socketid"`
	Colour       string `json:"colour"`
	//
	IsShaper   bool `json:"shaper"`
	IsElder    bool `json:"elder"`
	IsCrusader bool `json:"crusader"`
	IsRedeemer bool `json:"redeemer"`
	IsHunter   bool `json:"hunter"`
	IsWarlord  bool `json:"warload"`
	//
	//CosmeticMods
	CraftedMods []string `json:"craftedMods"`
	//EnchantMods  []Mods `json:"enchantMods"`
	ImplicitMods []string `json:"implicitMods"`
	ExplicitMods []string `json:"explicitMods"`
	//UtilityMods  []Mods `json:"utilityMods"`
}
type Properties struct {
	Name        string     `json:"name"`
	Values      [][2]int16 `json:"values"`
	DisplayMode int8       `json:"displaymode"`
	Type        int8       `json:"type"`
	Suffix      string     `json:"suffix"`
}
