package tdef

// TROOP and TOWER constants are used as indicators in the Type field of the units
const (
	TROOP = iota
	TOWER = iota
)

// Tower represents a defensive tower that can be purchased
type Tower struct {
	Type   int
	Cost   int
	Income int
	Damage int
	Enum   int
	Hp     int
	MaxHP  int
	Name   string
	Reach  int
	Speed  int
}

func getTowerPeashooter() *Tower {
	return &Tower{TOWER, 500, -100, 10, 50, 200, 200, "Peashooter", 300, 5}
}

func getTowerFirewall() *Tower {
	return &Tower{TOWER, 1000, -150, 10, 51, 3000, 3000, "Firewall", 300, 5}
}

func getTowerGuardian() *Tower {
	return &Tower{TOWER, 5000, -300, 30, 52, 600, 600, "Guardian", 250, 3}
}

func getTowerBank() *Tower {
	return &Tower{TOWER, 1000, 100, 0, 53, 500, 500, "Bank", 1, 60}
}

func getTowerJunkyard() *Tower {
	return &Tower{TOWER, 1000, 0, 0, 54, 300, 300, "Junkyard", 500, 3}
}

func getTowerStartUp() *Tower {
	return &Tower{TOWER, 3000, 100, 0, 55, 400, 400, "StartUp", 300, 60}
}

func getTowerCorporation() *Tower {
	return &Tower{TOWER, 25000, 200, 0, 56, 1000, 100, "Corporation", 300, 3}
}

func getTowerWarpDrive() *Tower {
	return &Tower{TOWER, 1000, 0, 0, 57, 400, 400, "WarpDrive", 0, 1}
}

func getTowerJammingStation() *Tower {
	return &Tower{TOWER, 3000, -150, 0, 58, 3000, 3000, "JammingStation", 200, 30}
}

func getTowerHotspot() *Tower {
	return &Tower{TOWER, 10000, -500, 0, 59, 400, 400, "Hotstop", 50, 10}
}

// AllTowers contains an array of all the possible Towers
var AllTowers = []*Tower{getTowerPeashooter(), getTowerFirewall(), getTowerGuardian(),
	getTowerBank(), getTowerJunkyard(), getTowerStartUp(), getTowerCorporation(), getTowerWarpDrive(),
	getTowerJammingStation(), getTowerHotspot()}

// Troop represents an offensive unit that can be purchased
type Troop struct {
	Type   int
	Damage int
	Cost   int
	Enum   int
	Hp     int
	MaxHP  int
	Name   string
	Reach  int
	Speed  int
	Stride int
}

func getTroopNut() *Troop {
	return &Troop{TROOP, 10, 200, 0, 100, 100, "Nut", 120, 5, 10}
}

func getTroopBolt() *Troop {
	return &Troop{TROOP, 0, 400, 1, 300, 300, "Bolt", 100, 8, 15}
}

func getTroopGreaseMonkey() *Troop {
	return &Troop{TROOP, 0, 300, 2, 75, 75, "GreaseMonkey", 200, 5, 10}
}

func getTroopWalker() *Troop {
	return &Troop{TROOP, 5, 800, 3, 800, 800, "Walker", 200, 2, 10}
}

func getTroopAimbot() *Troop {
	return &Troop{TROOP, 100, 3000, 4, 100, 100, "Aimbot", 700, 60, 5}
}

func getTroopHardDrive() *Troop {
	return &Troop{TROOP, 50, 2500, 5, 500, 500, "HardDrive", 50, 5, 5}
}

func getTroopScrapheap() *Troop {
	return &Troop{TROOP, 8, 9000, 6, 9000, 9000, "Scrapheap", 120, 5, 5}
}

func getTroopGasGuzzler() *Troop {
	return &Troop{TROOP, 0, 10000, 7, 10000, 10000, "GasGuzzler", 50, 5, 5}
}

func getTroopTerminator() *Troop {
	return &Troop{TROOP, 100, 8000, 8, 80, 80, "Terminator", 120, 5, 6}
}

func getTroopBlackHat() *Troop {
	return &Troop{TROOP, 0, 250, 9, 50, 50, "Blackhat", 50, 1, 6}
}

func getTroopMalware() *Troop {
	return &Troop{TROOP, 30, 6000, 10, 80, 80, "Malware", 200, 5, 4}
}

func getTroopGandhi() *Troop {
	return &Troop{TROOP, 0, 500000, 11, 50, 50, "Gandhi", 50, 1, 6}
}

//AllTroops contains all of the possible Troops
var AllTroops = []*Troop{getTroopNut(), getTroopBolt(), getTroopGreaseMonkey(), getTroopWalker(),
	getTroopAimbot(), getTroopHardDrive(), getTroopScrapheap(), getTroopGasGuzzler(),
	getTroopTerminator(), getTroopBlackHat(), getTroopMalware(), getTroopGandhi()}
