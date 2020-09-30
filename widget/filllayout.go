package widget

import "image"

type fillLayout struct {
	padding Insets
	dirty   bool
}

type FillLayoutOpt func(f *fillLayout)

const FillLayoutOpts = fillLayoutOpts(true)

type fillLayoutOpts bool

func NewFillLayout(opts ...FillLayoutOpt) Layouter {
	f := &fillLayout{}

	for _, o := range opts {
		o(f)
	}

	return f
}

func (o fillLayoutOpts) Padding(i Insets) FillLayoutOpt {
	return func(f *fillLayout) {
		f.padding = i
	}
}

func (f *fillLayout) PreferredSize(widgets []PreferredSizeLocateableWidget) (int, int) {
	px, py := f.padding.Dx(), f.padding.Dy()

	if len(widgets) == 0 {
		return px, py
	}

	w, h := widgets[0].PreferredSize()
	return w + px, h + py
}

func (f *fillLayout) Layout(widgets []PreferredSizeLocateableWidget, rect image.Rectangle) {
	if !f.dirty {
		return
	}

	defer func() {
		f.dirty = false
	}()

	if len(widgets) == 0 {
		return
	}

	widgets[0].SetLocation(f.padding.Apply(rect))
}

func (f *fillLayout) MarkDirty() {
	f.dirty = true
}
