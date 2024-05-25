package ui

import (
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"image"
)

func (p *Player) drawHover(gtx C) D {
	gtx.Constraints.Min = gtx.Constraints.Max
	defer clip.Rect(image.Rectangle{Max: gtx.Constraints.Max}).Push(gtx.Ops).Pop()
	defer p.screenClickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return D{Size: gtx.Constraints.Max}
	})
	event.Op(gtx.Ops, &p.lastHoveredTime)
	//pointer.InputOp{
	//	Tag:   &p.lastHoveredTime,
	//	Types: pointer.Move,
	//}.Add(gtx.Ops)
	return D{Size: gtx.Constraints.Max}
}
