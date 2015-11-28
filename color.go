// Package termboxutil provides useful operations that augment nsf's termbox-go
// library.
package termboxutil

import (
	"image/color"

	"github.com/nsf/termbox-go"
)

var palette16 = palette256[:16]

var palette256 = color.Palette{
	color.RGBA{0x00, 0x00, 0x00, 0xff},
	color.RGBA{0x80, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0x80, 0x00, 0xff},
	color.RGBA{0x80, 0x80, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0x80, 0xff},
	color.RGBA{0x80, 0x00, 0x80, 0xff},
	color.RGBA{0x00, 0x80, 0x80, 0xff},
	color.RGBA{0xc0, 0xc0, 0xc0, 0xff},

	color.RGBA{0x80, 0x80, 0x80, 0xff},
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0xff, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.RGBA{0xff, 0x00, 0xff, 0xff},
	color.RGBA{0x00, 0xff, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0xff, 0xff},

	color.RGBA{0x00, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0x5f, 0xff},
	color.RGBA{0x00, 0x00, 0x87, 0xff},
	color.RGBA{0x00, 0x00, 0xaf, 0xff},
	color.RGBA{0x00, 0x00, 0xd7, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.RGBA{0x00, 0x5f, 0x00, 0xff},
	color.RGBA{0x00, 0x5f, 0x5f, 0xff},
	color.RGBA{0x00, 0x5f, 0x87, 0xff},
	color.RGBA{0x00, 0x5f, 0xaf, 0xff},
	color.RGBA{0x00, 0x5f, 0xd7, 0xff},
	color.RGBA{0x00, 0x5f, 0xff, 0xff},
	color.RGBA{0x00, 0x87, 0x00, 0xff},
	color.RGBA{0x00, 0x87, 0x5f, 0xff},
	color.RGBA{0x00, 0x87, 0x87, 0xff},
	color.RGBA{0x00, 0x87, 0xaf, 0xff},
	color.RGBA{0x00, 0x87, 0xd7, 0xff},
	color.RGBA{0x00, 0x87, 0xff, 0xff},
	color.RGBA{0x00, 0xaf, 0x00, 0xff},
	color.RGBA{0x00, 0xaf, 0x5f, 0xff},
	color.RGBA{0x00, 0xaf, 0x87, 0xff},
	color.RGBA{0x00, 0xaf, 0xaf, 0xff},
	color.RGBA{0x00, 0xaf, 0xd7, 0xff},
	color.RGBA{0x00, 0xaf, 0xff, 0xff},
	color.RGBA{0x00, 0xd7, 0x00, 0xff},
	color.RGBA{0x00, 0xd7, 0x5f, 0xff},
	color.RGBA{0x00, 0xd7, 0x87, 0xff},
	color.RGBA{0x00, 0xd7, 0xaf, 0xff},
	color.RGBA{0x00, 0xd7, 0xd7, 0xff},
	color.RGBA{0x00, 0xd7, 0xff, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x5f, 0xff},
	color.RGBA{0x00, 0xff, 0x87, 0xff},
	color.RGBA{0x00, 0xff, 0xaf, 0xff},
	color.RGBA{0x00, 0xff, 0xd7, 0xff},
	color.RGBA{0x00, 0xff, 0xff, 0xff},
	color.RGBA{0x5f, 0x00, 0x00, 0xff},
	color.RGBA{0x5f, 0x00, 0x5f, 0xff},
	color.RGBA{0x5f, 0x00, 0x87, 0xff},
	color.RGBA{0x5f, 0x00, 0xaf, 0xff},
	color.RGBA{0x5f, 0x00, 0xd7, 0xff},
	color.RGBA{0x5f, 0x00, 0xff, 0xff},
	color.RGBA{0x5f, 0x5f, 0x00, 0xff},
	color.RGBA{0x5f, 0x5f, 0x5f, 0xff},
	color.RGBA{0x5f, 0x5f, 0x87, 0xff},
	color.RGBA{0x5f, 0x5f, 0xaf, 0xff},
	color.RGBA{0x5f, 0x5f, 0xd7, 0xff},
	color.RGBA{0x5f, 0x5f, 0xff, 0xff},
	color.RGBA{0x5f, 0x87, 0x00, 0xff},
	color.RGBA{0x5f, 0x87, 0x5f, 0xff},
	color.RGBA{0x5f, 0x87, 0x87, 0xff},
	color.RGBA{0x5f, 0x87, 0xaf, 0xff},
	color.RGBA{0x5f, 0x87, 0xd7, 0xff},
	color.RGBA{0x5f, 0x87, 0xff, 0xff},
	color.RGBA{0x5f, 0xaf, 0x00, 0xff},
	color.RGBA{0x5f, 0xaf, 0x5f, 0xff},
	color.RGBA{0x5f, 0xaf, 0x87, 0xff},
	color.RGBA{0x5f, 0xaf, 0xaf, 0xff},
	color.RGBA{0x5f, 0xaf, 0xd7, 0xff},
	color.RGBA{0x5f, 0xaf, 0xff, 0xff},
	color.RGBA{0x5f, 0xd7, 0x00, 0xff},
	color.RGBA{0x5f, 0xd7, 0x5f, 0xff},
	color.RGBA{0x5f, 0xd7, 0x87, 0xff},
	color.RGBA{0x5f, 0xd7, 0xaf, 0xff},
	color.RGBA{0x5f, 0xd7, 0xd7, 0xff},
	color.RGBA{0x5f, 0xd7, 0xff, 0xff},
	color.RGBA{0x5f, 0xff, 0x00, 0xff},
	color.RGBA{0x5f, 0xff, 0x5f, 0xff},
	color.RGBA{0x5f, 0xff, 0x87, 0xff},
	color.RGBA{0x5f, 0xff, 0xaf, 0xff},
	color.RGBA{0x5f, 0xff, 0xd7, 0xff},
	color.RGBA{0x5f, 0xff, 0xff, 0xff},
	color.RGBA{0x87, 0x00, 0x00, 0xff},
	color.RGBA{0x87, 0x00, 0x5f, 0xff},
	color.RGBA{0x87, 0x00, 0x87, 0xff},
	color.RGBA{0x87, 0x00, 0xaf, 0xff},
	color.RGBA{0x87, 0x00, 0xd7, 0xff},
	color.RGBA{0x87, 0x00, 0xff, 0xff},
	color.RGBA{0x87, 0x5f, 0x00, 0xff},
	color.RGBA{0x87, 0x5f, 0x5f, 0xff},
	color.RGBA{0x87, 0x5f, 0x87, 0xff},
	color.RGBA{0x87, 0x5f, 0xaf, 0xff},
	color.RGBA{0x87, 0x5f, 0xd7, 0xff},
	color.RGBA{0x87, 0x5f, 0xff, 0xff},
	color.RGBA{0x87, 0x87, 0x00, 0xff},
	color.RGBA{0x87, 0x87, 0x5f, 0xff},
	color.RGBA{0x87, 0x87, 0x87, 0xff},
	color.RGBA{0x87, 0x87, 0xaf, 0xff},
	color.RGBA{0x87, 0x87, 0xd7, 0xff},
	color.RGBA{0x87, 0x87, 0xff, 0xff},
	color.RGBA{0x87, 0xaf, 0x00, 0xff},
	color.RGBA{0x87, 0xaf, 0x5f, 0xff},
	color.RGBA{0x87, 0xaf, 0x87, 0xff},
	color.RGBA{0x87, 0xaf, 0xaf, 0xff},
	color.RGBA{0x87, 0xaf, 0xd7, 0xff},
	color.RGBA{0x87, 0xaf, 0xff, 0xff},
	color.RGBA{0x87, 0xd7, 0x00, 0xff},
	color.RGBA{0x87, 0xd7, 0x5f, 0xff},
	color.RGBA{0x87, 0xd7, 0x87, 0xff},
	color.RGBA{0x87, 0xd7, 0xaf, 0xff},
	color.RGBA{0x87, 0xd7, 0xd7, 0xff},
	color.RGBA{0x87, 0xd7, 0xff, 0xff},
	color.RGBA{0x87, 0xff, 0x00, 0xff},
	color.RGBA{0x87, 0xff, 0x5f, 0xff},
	color.RGBA{0x87, 0xff, 0x87, 0xff},
	color.RGBA{0x87, 0xff, 0xaf, 0xff},
	color.RGBA{0x87, 0xff, 0xd7, 0xff},
	color.RGBA{0x87, 0xff, 0xff, 0xff},
	color.RGBA{0xaf, 0x00, 0x00, 0xff},
	color.RGBA{0xaf, 0x00, 0x5f, 0xff},
	color.RGBA{0xaf, 0x00, 0x87, 0xff},
	color.RGBA{0xaf, 0x00, 0xaf, 0xff},
	color.RGBA{0xaf, 0x00, 0xd7, 0xff},
	color.RGBA{0xaf, 0x00, 0xff, 0xff},
	color.RGBA{0xaf, 0x5f, 0x00, 0xff},
	color.RGBA{0xaf, 0x5f, 0x5f, 0xff},
	color.RGBA{0xaf, 0x5f, 0x87, 0xff},
	color.RGBA{0xaf, 0x5f, 0xaf, 0xff},
	color.RGBA{0xaf, 0x5f, 0xd7, 0xff},
	color.RGBA{0xaf, 0x5f, 0xff, 0xff},
	color.RGBA{0xaf, 0x87, 0x00, 0xff},
	color.RGBA{0xaf, 0x87, 0x5f, 0xff},
	color.RGBA{0xaf, 0x87, 0x87, 0xff},
	color.RGBA{0xaf, 0x87, 0xaf, 0xff},
	color.RGBA{0xaf, 0x87, 0xd7, 0xff},
	color.RGBA{0xaf, 0x87, 0xff, 0xff},
	color.RGBA{0xaf, 0xaf, 0x00, 0xff},
	color.RGBA{0xaf, 0xaf, 0x5f, 0xff},
	color.RGBA{0xaf, 0xaf, 0x87, 0xff},
	color.RGBA{0xaf, 0xaf, 0xaf, 0xff},
	color.RGBA{0xaf, 0xaf, 0xd7, 0xff},
	color.RGBA{0xaf, 0xaf, 0xff, 0xff},
	color.RGBA{0xaf, 0xd7, 0x00, 0xff},
	color.RGBA{0xaf, 0xd7, 0x5f, 0xff},
	color.RGBA{0xaf, 0xd7, 0x87, 0xff},
	color.RGBA{0xaf, 0xd7, 0xaf, 0xff},
	color.RGBA{0xaf, 0xd7, 0xd7, 0xff},
	color.RGBA{0xaf, 0xd7, 0xff, 0xff},
	color.RGBA{0xaf, 0xff, 0x00, 0xff},
	color.RGBA{0xaf, 0xff, 0x5f, 0xff},
	color.RGBA{0xaf, 0xff, 0x87, 0xff},
	color.RGBA{0xaf, 0xff, 0xaf, 0xff},
	color.RGBA{0xaf, 0xff, 0xd7, 0xff},
	color.RGBA{0xaf, 0xff, 0xff, 0xff},
	color.RGBA{0xd7, 0x00, 0x00, 0xff},
	color.RGBA{0xd7, 0x00, 0x5f, 0xff},
	color.RGBA{0xd7, 0x00, 0x87, 0xff},
	color.RGBA{0xd7, 0x00, 0xaf, 0xff},
	color.RGBA{0xd7, 0x00, 0xd7, 0xff},
	color.RGBA{0xd7, 0x00, 0xff, 0xff},
	color.RGBA{0xd7, 0x5f, 0x00, 0xff},
	color.RGBA{0xd7, 0x5f, 0x5f, 0xff},
	color.RGBA{0xd7, 0x5f, 0x87, 0xff},
	color.RGBA{0xd7, 0x5f, 0xaf, 0xff},
	color.RGBA{0xd7, 0x5f, 0xd7, 0xff},
	color.RGBA{0xd7, 0x5f, 0xff, 0xff},
	color.RGBA{0xd7, 0x87, 0x00, 0xff},
	color.RGBA{0xd7, 0x87, 0x5f, 0xff},
	color.RGBA{0xd7, 0x87, 0x87, 0xff},
	color.RGBA{0xd7, 0x87, 0xaf, 0xff},
	color.RGBA{0xd7, 0x87, 0xd7, 0xff},
	color.RGBA{0xd7, 0x87, 0xff, 0xff},
	color.RGBA{0xdf, 0xaf, 0x00, 0xff},
	color.RGBA{0xdf, 0xaf, 0x5f, 0xff},
	color.RGBA{0xdf, 0xaf, 0x87, 0xff},
	color.RGBA{0xdf, 0xaf, 0xaf, 0xff},
	color.RGBA{0xdf, 0xaf, 0xdf, 0xff},
	color.RGBA{0xdf, 0xaf, 0xff, 0xff},
	color.RGBA{0xdf, 0xdf, 0x00, 0xff},
	color.RGBA{0xdf, 0xdf, 0x5f, 0xff},
	color.RGBA{0xdf, 0xdf, 0x87, 0xff},
	color.RGBA{0xdf, 0xdf, 0xaf, 0xff},
	color.RGBA{0xdf, 0xdf, 0xdf, 0xff},
	color.RGBA{0xdf, 0xdf, 0xff, 0xff},
	color.RGBA{0xdf, 0xff, 0x00, 0xff},
	color.RGBA{0xdf, 0xff, 0x5f, 0xff},
	color.RGBA{0xdf, 0xff, 0x87, 0xff},
	color.RGBA{0xdf, 0xff, 0xaf, 0xff},
	color.RGBA{0xdf, 0xff, 0xdf, 0xff},
	color.RGBA{0xdf, 0xff, 0xff, 0xff},
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0xff, 0x00, 0x5f, 0xff},
	color.RGBA{0xff, 0x00, 0x87, 0xff},
	color.RGBA{0xff, 0x00, 0xaf, 0xff},
	color.RGBA{0xff, 0x00, 0xdf, 0xff},
	color.RGBA{0xff, 0x00, 0xff, 0xff},
	color.RGBA{0xff, 0x5f, 0x00, 0xff},
	color.RGBA{0xff, 0x5f, 0x5f, 0xff},
	color.RGBA{0xff, 0x5f, 0x87, 0xff},
	color.RGBA{0xff, 0x5f, 0xaf, 0xff},
	color.RGBA{0xff, 0x5f, 0xdf, 0xff},
	color.RGBA{0xff, 0x5f, 0xff, 0xff},
	color.RGBA{0xff, 0x87, 0x00, 0xff},
	color.RGBA{0xff, 0x87, 0x5f, 0xff},
	color.RGBA{0xff, 0x87, 0x87, 0xff},
	color.RGBA{0xff, 0x87, 0xaf, 0xff},
	color.RGBA{0xff, 0x87, 0xdf, 0xff},
	color.RGBA{0xff, 0x87, 0xff, 0xff},
	color.RGBA{0xff, 0xaf, 0x00, 0xff},
	color.RGBA{0xff, 0xaf, 0x5f, 0xff},
	color.RGBA{0xff, 0xaf, 0x87, 0xff},
	color.RGBA{0xff, 0xaf, 0xaf, 0xff},
	color.RGBA{0xff, 0xaf, 0xdf, 0xff},
	color.RGBA{0xff, 0xaf, 0xff, 0xff},
	color.RGBA{0xff, 0xdf, 0x00, 0xff},
	color.RGBA{0xff, 0xdf, 0x5f, 0xff},
	color.RGBA{0xff, 0xdf, 0x87, 0xff},
	color.RGBA{0xff, 0xdf, 0xaf, 0xff},
	color.RGBA{0xff, 0xdf, 0xdf, 0xff},
	color.RGBA{0xff, 0xdf, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0x00, 0xff},
	color.RGBA{0xff, 0xff, 0x5f, 0xff},
	color.RGBA{0xff, 0xff, 0x87, 0xff},
	color.RGBA{0xff, 0xff, 0xaf, 0xff},
	color.RGBA{0xff, 0xff, 0xdf, 0xff},
	color.RGBA{0xff, 0xff, 0xff, 0xff},

	color.RGBA{0x08, 0x08, 0x08, 0xff},
	color.RGBA{0x12, 0x12, 0x12, 0xff},
	color.RGBA{0x1c, 0x1c, 0x1c, 0xff},
	color.RGBA{0x26, 0x26, 0x26, 0xff},
	color.RGBA{0x30, 0x30, 0x30, 0xff},
	color.RGBA{0x3a, 0x3a, 0x3a, 0xff},
	color.RGBA{0x44, 0x44, 0x44, 0xff},
	color.RGBA{0x4e, 0x4e, 0x4e, 0xff},
	color.RGBA{0x58, 0x58, 0x58, 0xff},
	color.RGBA{0x62, 0x62, 0x62, 0xff},
	color.RGBA{0x6c, 0x6c, 0x6c, 0xff},
	color.RGBA{0x76, 0x76, 0x76, 0xff},
	color.RGBA{0x80, 0x80, 0x80, 0xff},
	color.RGBA{0x8a, 0x8a, 0x8a, 0xff},
	color.RGBA{0x94, 0x94, 0x94, 0xff},
	color.RGBA{0x9e, 0x9e, 0x9e, 0xff},
	color.RGBA{0xa8, 0xa8, 0xa8, 0xff},
	color.RGBA{0xb2, 0xb2, 0xb2, 0xff},
	color.RGBA{0xbc, 0xbc, 0xbc, 0xff},
	color.RGBA{0xc6, 0xc6, 0xc6, 0xff},
	color.RGBA{0xd0, 0xd0, 0xd0, 0xff},
	color.RGBA{0xda, 0xda, 0xda, 0xff},
	color.RGBA{0xe4, 0xe4, 0xe4, 0xff},
	color.RGBA{0xee, 0xee, 0xee, 0xff},
}

var palette216 = palette256[16 : 216+16]

var paletteGray = append(append(color.Palette{palette256[16]}, palette256[232:]...), palette256[231])

// ColorNormal converts a color to a termbox attribute for use with
// termbox.OutputNormal.
func ColorNormal(c color.Color) termbox.Attribute {
	i := palette16.Index(c)
	return termbox.Attribute((i%8)+1) | (termbox.Attribute(i/8) * termbox.AttrBold)
}

// Color256 converts a color to a termbox attribute for use with
// termbox.Output256.
func Color256(c color.Color) termbox.Attribute {
	return termbox.Attribute(palette256[16:].Index(c) + 1 + 16)
}

// Color216 converts a color to a termbox attribute for use with
// termbox.Output216.
func Color216(c color.Color) termbox.Attribute {
	return termbox.Attribute(palette216.Index(c) + 1)
}

// ColorGray converts a color to a termbox attribute for use with
// termbox.OutputGrayscale.
func ColorGray(c color.Color) termbox.Attribute {
	return termbox.Attribute(paletteGray.Index(c) + 1)
}

// Color converts a color to a termbox attribute using the current output mode.
func Color(c color.Color) termbox.Attribute {
	switch termbox.SetOutputMode(termbox.OutputCurrent) {
	case termbox.OutputNormal:
		return ColorNormal(c)
	case termbox.Output256:
		return Color256(c)
	case termbox.Output216:
		return Color216(c)
	case termbox.OutputGrayscale:
		return ColorGray(c)
	default:
		panic("unexpected output mode")
	}
}
