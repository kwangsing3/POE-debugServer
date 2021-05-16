package module

import (
	"fmt"
	"math/rand"
	"poe-debugserver/src/model"
	"strconv"
)

// To generate differet data in need (ex. new characters, random items.... etc.)

func NewCharacterList(length int) (result []model.Unit) {
	index := 0
	result = make([]model.Unit, length)
	for i := range result {
		char := model.Character{
			Name:            `Player` + fmt.Sprint(index) + `_DBServer`,
			League:          `Standard`,
			ClassId:         0,
			AscendancyClass: 0,
			Class:           `Scion`,
			Level:           100,
			Experience:      4250334444,
			LastActive:      true,
		}
		result[i].Character = char
		result[i].PassiveTree = GeneratePassiveTree()
		result[i].Items = GenerateItems()
	}
	return
}

func GenerateItems() (result []model.Items) {
	length := 5 + rand.Intn(2)
	inventoryid := []string{`Flask`, `Belt`, `Boots`, `Helm`, `Gloves`, `Weapon`, `Offhand`, `Ring1`, `Ring2`, `Amulet`}
	if length > len(inventoryid) {
		length = len(inventoryid)
	}
	result = make([]model.Items, length)

	// generate explicitMods
	expMods := make([]string, 1+rand.Intn(5))
	for i := range expMods {
		expMods[i] = `RandomExMods` + fmt.Sprintln(i)
	}
	// generate EnchantMods
	mB, _ := strconv.ParseBool(fmt.Sprintln(rand.Intn(1)))
	enMods := map[bool][]string{true: {`RandomEnchantMods`}, false: nil}[mB]
	//influence (max:2)
	influB := []bool{false, false, false, false, false, false}
	inf1 := rand.Intn(6)
	if inf1 < len(influB) && inf1%2 == 0 {
		influB[inf1] = true
	}
	inf2 := rand.Intn(6)
	if inf2 < len(influB) && inf2%2 == 0 {
		influB[inf2] = true
	}
	//isCorrupted
	IsCorrupted, _ := strconv.ParseBool(fmt.Sprintln(rand.Intn(1)))
	//result
	for i := range result {
		result[i] = model.Items{
			Name:         `RandomItems` + fmt.Sprintln(i),
			ExplicitMods: expMods,
			InventoryId:  inventoryid[i],
			EnchantMods:  enMods,
			IsShaper:     influB[0],
			IsElder:      influB[1],
			IsCrusader:   influB[2],
			IsRedeemer:   influB[3],
			IsHunter:     influB[4],
			IsWarlord:    influB[5],
			IsCorrupted:  IsCorrupted,
		}
	}
	return
}
func GeneratePassiveTree() (result model.PassiveTree) {
	index := rand.Intn(2)
	sample := []int32{}

	switch index {
	case 0:
		sample = []int32{1031, 2292, 4184, 4367, 5972, 11420, 14936, 15711, 17579, 19501, 19635, 21075, 21460, 21934, 21958, 22090, 27203, 27929, 33296, 33755, 36226, 36412, 36774, 38805, 38900, 40609, 40705, 44184, 44955, 48362, 53732, 57264}
	case 1:
		sample = []int32{94, 2094, 2185, 6542, 9206, 10843, 12412, 12948, 20807, 21228, 23912, 31222, 32117, 32514, 36221, 36287, 38149, 39665, 39821, 41866, 42964, 43807, 48099, 49459, 50338, 51881, 54142, 55750, 60803, 60942, 61306, 63139, 65210, 65224}
	case 2:
		sample = []int32{1767, 1822, 3656, 4502, 5823, 8001, 9877, 18769, 18865, 21575, 21835, 23507, 24050, 25058, 32710, 35894, 37671, 38129, 38190, 38995, 40644, 41989, 42686, 46136, 47362, 48477, 55247, 58649, 61320, 61653, 63639, 63799, 64501}
	default:
	}
	result.Hashes = sample
	return
}
