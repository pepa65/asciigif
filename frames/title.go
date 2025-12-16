package frames

import (
	"math/rand"
	"time"
)

const titleFrame = `
                        "      "             "      m"" 
  mmm    mmm    mmm   mmm    mmm     mmmm  mmm    mm#mm 
 "   #  #   "  #"  "    #      #    #" "#    #      #   
 m"""#   """m  #        #      #    #   #    #      #   
 "mm"#  "mmm"  "#mm"  mm#mm  mm#mm  "#m"#  mm#mm    #   
                                     m  #               
                                      ""                
`

var titleFrames []string

func generateGlitchedFrames(input string, count int) []string {
	frames := make([]string, count)
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	glitchChars := []rune{'!', '@', '#', '%', '^', '&', '*', '(', ')', '_', '+', '-', '='}
	for i := 0; i < count; i++ {
		glitched := []rune(input)
		for j, _ := range glitched {
			if r.Float32() < 0.05 { // 5% chance to glitch
				if glitched[j] != 10 {
					glitched[j] = glitchChars[r.Intn(len(glitchChars))]
				}
			}
		}
		frames[i] = string(glitched)
	}
	return frames
}

func init() {
	titleFrames = generateGlitchedFrames(titleFrame, 20)
	FrameMap["title"] = DefaultFrameType(titleFrames)
}
