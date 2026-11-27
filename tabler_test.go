package tabler

import (
	"slices"
	"strings"
	"testing"
)

func TestSVG(t *testing.T) {
	doc, ok := SVG(Settings)
	if !ok || !strings.Contains(doc, "<svg") || !strings.Contains(doc, `stroke="currentColor"`) {
		t.Fatalf("settings: ok=%v doc=%q", ok, doc)
	}
	for _, bad := range []string{"", "nope", "../LICENSE", "a/b", `a\b`, "settings.svg"} {
		if _, ok := SVG(bad); ok {
			t.Errorf("SVG(%q) ok, want refusal", bad)
		}
	}
}

// Every named icon is embedded, is an SVG that recolours with the ink, and is
// the bare document (no comment header) like the first icon.
func TestEveryConstantIsEmbedded(t *testing.T) {
	all := []string{Settings, PlugConnected, PlugConnectedX, Login, Logout, Pencil, Refresh,
		ExternalLink, X, Trash, ArrowLeft, ArrowRight, Check, Plus}
	if got := len(Names()); got != len(all) {
		t.Fatalf("%d icons embedded, %d named: %v", got, len(all), Names())
	}
	for _, name := range all {
		doc, ok := SVG(name)
		if !ok {
			t.Errorf("%s: not embedded", name)
			continue
		}
		if !strings.HasPrefix(doc, "<svg") || !strings.Contains(doc, `stroke="currentColor"`) ||
			!strings.Contains(doc, `viewBox="0 0 24 24"`) || strings.Contains(doc, "<script") {
			t.Errorf("%s: not a plain currentColor 24x24 SVG: %.60q", name, doc)
		}
	}
}

func TestMustSVG(t *testing.T) {
	if MustSVG(Settings) == "" {
		t.Fatal("empty")
	}
	defer func() {
		if recover() == nil {
			t.Fatal("no panic for an unknown icon")
		}
	}()
	MustSVG("nope")
}

func TestNames(t *testing.T) {
	n := Names()
	if !slices.IsSorted(n) || !slices.Contains(n, Settings) {
		t.Fatalf("Names = %v", n)
	}
	for _, name := range n {
		if _, ok := SVG(name); !ok {
			t.Errorf("listed %q not readable", name)
		}
	}
}
