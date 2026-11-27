# tabler-icons

A curated subset of [Tabler Icons](https://tabler.io/icons) (MIT, outline
style, 24x24, `stroke="currentColor"`) for go-widgets programs, one Go API:

```go
doc, ok := tabler.SVG(tabler.Settings) // name -> SVG bytes
draw := toolkit.SVGIcon(tabler.MustSVG(tabler.Settings))
```

`toolkit.SVGIcon` (go-widgets/toolkit) rasterises through go-gfx/gfx
`svg.Rasterize`, recolours to the ink, and caches one raster per
(icon, ink, size) — never per frame.

## Size

Only the icons in `icons/` are embedded (today: `settings`, `plug-connected`,
`plug-connected-x`, `login-2`, `logout-2`, `pencil`, `refresh`, `external-link`,
`x`, `trash`, `arrow-left`, `arrow-right`, `check`, `plus`; about 6 KB), not the
~5000-icon set. Source: `@tabler/icons` 3.35.0, `icons/outline/`.

## Adding an icon

1. Copy `icons/outline/<name>.svg` from `@tabler/icons` 3.35.0 into `icons/`.
2. Add `const Name = "<name>"` in `tabler.go`.
3. `go test ./...`.

## License

Icons: MIT, (c) 2020-2025 Pawel Kuna — see `LICENSE`. Go code: BSD-3-Clause.
