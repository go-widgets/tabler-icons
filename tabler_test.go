package tabler

import (
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
	if len(n) == 0 || n[0] != "settings" {
		t.Fatalf("Names = %v", n)
	}
	for _, name := range n {
		if _, ok := SVG(name); !ok {
			t.Errorf("listed %q not readable", name)
		}
	}
}
