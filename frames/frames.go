package frames

type FrameType struct {
	GetFrame  func(int) string
	GetLength func() int
}

// Create a function that returns the next frame, based on length
func DefaultGetFrame(frames []string) func(int) string {
	return func(i int) string {
		return frames[i%(len(frames))]
	}
}

// Create a function that returns frame length
func DefaultGetLength(frames []string) func() int {
	return func() int {
		return len(frames)
	}
}

// Given frames, create a FrameType with those frames
func DefaultFrameType(frames []string) FrameType {
	return FrameType{
		GetFrame:  DefaultGetFrame(frames),
		GetLength: DefaultGetLength(frames),
	}
}

var FrameMap = map[string]FrameType{
// LABEL             VARIABLE
// (a-di.eu/LABEL)   ([Default]FrameType)
	"batman":          Batman,
	"batmanrobin":     BatmanRobin,
	"clock":           Clock,
	"coin":            Coin,
	"donut":           Donut,
	"dvd":             Dvd,
	"forrest":         Forrest,
	"hes":             HES,
	"knot":            Knot,
	"nyan":            Nyan,
	"parrot":          Parrot,
	"rick":            Rick,
	"spidy":           Spidy,
	"purdue":          Purdue,
	"as":              As,
	"bomb":            Bomb,
	"maxwell":         Maxwell,
	"earth":           Earth,
	"kitty":           Kitty,
	"india":           India,
	"brittany":        Brittany,
	"blueshirt":       Blueshirt,
	"sixtyseven":      Sixtyseven,
	"macarenavisbal":  MacarenaVisbal,
	"badapple":        BadApple,
	"gandalf":         Gandalf,
	"love":            Love,
	"gon":             Gon,
	"skull":           Skull,
	"gina":            Gina,
	"sodapop":         SodaPop,
	"darkfountain":    DarkFountain,
}
