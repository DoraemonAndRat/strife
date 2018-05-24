package strife

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/ttf"
)

var fontLoaderInitialized bool = false

type fontStyle int

const (
	plain fontStyle = iota
	underline
	bold
	strike
)

type glyph struct {
	tex *sdl.Texture
	dim []int32
}

type glyphInfo struct {
	val   rune
	col   uint32
	style fontStyle
}

func (g *glyphInfo) asKey() string {
	return string(g.val) + fmt.Sprintf("%d", g.col) + fmt.Sprintf("%d", int(g.style))
}

func encode(col uint32, style fontStyle, val rune) glyphInfo {
	return glyphInfo{
		val: val, col: col, style: style,
	}
}

type Font struct {
	*ttf.Font
	texCache map[string]*glyph
}

func (r *Font) hasGlyph(g glyphInfo) (*glyph, bool) {
	if val, ok := r.texCache[g.asKey()]; ok {
		return val, true
	}
	return nil, false
}

func (f *Font) cache(g glyphInfo, texture *sdl.Texture, dim []int32) *glyph {
	// todo cache collision?
	glyph := &glyph{texture, dim}
	texture.SetBlendMode(sdl.BLENDMODE_BLEND)
	f.texCache[g.asKey()] = glyph
	return glyph
}

func LoadFont(path string, size int) (*Font, error) {
	if !fontLoaderInitialized {
		ttf.Init()
	}

	font, err := ttf.OpenFont(path, size)
	if err != nil {
		return nil, fmt.Errorf("Failed to load font at '%s'\n", path)
	}

	return &Font{
		font,
		map[string]*glyph{},
	}, nil
}

func (f *Font) Destroy() {
	for _, glyph := range f.texCache {
		glyph.tex.Destroy()
	}
	f.Font.Close()
}
