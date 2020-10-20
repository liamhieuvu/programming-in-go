package font

import (
	"fmt"
	"log"
)

const defFamily = "Courier"
const defSize = 10

type Font struct {
	family string
	size   int
}

func New(family string, size int) *Font {
	return &Font{saneFamily(defFamily, family), saneSize(defSize, size)}
}

func (font *Font) Family() string { return font.family }

func (font *Font) SetFamily(family string) {
	font.family = saneFamily(font.family, family)
}

func (font *Font) Size() int { return font.size }

func (font *Font) SetSize(size int) {
	font.size = saneSize(font.size, size)
}

func (font *Font) String() string {
	return fmt.Sprintf("{font-family: %q; font-size: %dpt;}", font.family,
		font.size)
}

func saneFamily(def, family string) string {
	if family == "" {
		log.Printf("Invalid family. Change to '%s'", def)
		return def
	}
	return family
}

func saneSize(oldSize, newSize int) int {
	if newSize > 144 || newSize < 5 {
		log.Printf("Invalid size (%d). Change to %d", oldSize)
		return oldSize
	}
	return newSize
}
