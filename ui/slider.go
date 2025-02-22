package ui

import (
	"gioui.org/widget/material"
)

func (p *Player) drawSlider(gtx C) D {
	wg := &p.uiSeekerWidget
	slider := material.Slider(p.th, wg)
	slider.Color = p.th.Fg
	if p.CurrentView() == StateView {
		slider.Color = p.th.Bg
	}
	return slider.Layout(gtx)
}
