package main

import (
	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/op"
	"github.com/mearaj/gio-video/ui"
	"log"
	"os"
	"time"
)

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("Gio Player"))
		if err := loop(w); err != nil {
			log.Fatalln(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(window *app.Window) (err error) {
	player := ui.Player{Window: window}

	var ops op.Ops

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			player.Layout(gtx)
			e.Frame(gtx.Ops)
		case key.Event:
			if e.Name == "→" && e.Modifiers == 0x0 && e.State == 0x0 {
				player.RightKeyPressed = true
				player.RightKeyLastUpdated = time.Now().UnixMilli()
			}
			if e.Name == "→" && e.Modifiers == 0x0 && e.State == 0x1 {
				player.RightKeyPressed = false
				player.RightKeyLastUpdated = time.Now().UnixMilli()
			}
			if e.Name == "↑" && e.Modifiers == 0x0 && e.State == 0x0 {
				player.UpKeyPressed = true
				player.UpKeyLastUpdated = time.Now().UnixMilli()
			}
			if e.Name == "↑" && e.Modifiers == 0x0 && e.State == 0x1 {
				player.UpKeyPressed = false
				player.UpKeyLastUpdated = time.Now().UnixMilli()
			}
			if e.Name == "←" && e.Modifiers == 0x0 && e.State == 0x0 {
				player.LeftKeyPressed = true
				player.LeftKeyLastUpdated = time.Now().UnixMilli()
			}
			if e.Name == "←" && e.Modifiers == 0x0 && e.State == 0x1 {
				player.LeftKeyPressed = false
				player.LeftKeyLastUpdated = time.Now().UnixMilli()
			}
			if e.Name == "↓" && e.Modifiers == 0x0 && e.State == 0x0 {
				player.DownKeyPressed = true
				player.DownKeyLastUpdated = time.Now().UnixMilli()
			}
			if e.Name == "↓" && e.Modifiers == 0x0 && e.State == 0x1 {
				player.DownKeyPressed = false
				player.DownKeyLastUpdated = time.Now().UnixMilli()
			}
			if e.Name == "Space" && e.Modifiers == 0x0 && e.State == 0x0 {
				player.SpaceKeyPressed = true
				player.SpaceKeyLastUpdated = time.Now().UnixMilli()
			}
			if e.Name == "Space" && e.Modifiers == 0x0 && e.State == 0x1 {
				player.SpaceKeyPressed = false
				player.SpaceKeyLastUpdated = time.Now().UnixMilli()
			}
			if e.Name == "⎋" && e.Modifiers == 0x0 && e.State == 0x0 {
				player.EscKeyPressed = true
				player.EscKeyLastUpdated = time.Now().UnixMilli()
			}
			if e.Name == "⎋" && e.Modifiers == 0x0 && e.State == 0x1 {
				player.EscKeyPressed = false
				player.EscKeyLastUpdated = time.Now().UnixMilli()
			}
		}
	}
}
