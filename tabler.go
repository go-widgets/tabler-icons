// Copyright (c) 2026 the go-widgets/tabler-icons authors. All rights reserved.
// The Go code is BSD-3-Clause; the embedded SVGs are Tabler Icons, MIT
// (see LICENSE and README.md).

// Package tabler embeds a small, curated subset of the Tabler Icons
// (https://tabler.io/icons, MIT, outline style) and maps an icon name to its
// SVG document.
//
// The set is deliberately tiny: a program links only the icons listed in
// icons/, not the ~5000-icon pack. README.md says how to add one.
//
// Render an icon through the toolkit, which rasterises once per (icon, ink,
// size) and caches it:
//
//	draw := toolkit.SVGIcon(tabler.MustSVG(tabler.Settings))
package tabler

import (
	"embed"
	"fmt"
	"io/fs"
	"sort"
	"strings"
)

//go:embed icons/*.svg
var iconFS embed.FS

// Names of the embedded icons.
const (
	// Settings is a cog (Tabler "settings").
	Settings = "settings"
)

// SVG returns the SVG document of the icon called name; ok is false for a name
// outside the embedded set. The SVG draws with stroke="currentColor", so a
// rasteriser's ink recolours it.
func SVG(name string) (doc string, ok bool) {
	// The embedded FS refuses any path that is not a plain file under icons/
	// ("..", absolute, backslash names), so no separate name filter is needed.
	b, err := fs.ReadFile(iconFS, "icons/"+name+".svg")
	if err != nil {
		return "", false
	}
	return string(b), true
}

// MustSVG is [SVG] that panics on an unknown name; for package-level
// initialisation with the name constants.
func MustSVG(name string) string {
	doc, ok := SVG(name)
	if !ok {
		panic(fmt.Sprintf("tabler: no icon %q", name))
	}
	return doc
}

// Names lists the embedded icon names, sorted.
func Names() []string {
	ents, _ := fs.ReadDir(iconFS, "icons")
	out := make([]string, 0, len(ents))
	for _, e := range ents {
		out = append(out, strings.TrimSuffix(e.Name(), ".svg"))
	}
	sort.Strings(out)
	return out
}
