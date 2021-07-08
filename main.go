package main

import (
    "github.com/gopherjs/gopherjs/js"
    "log"
)

var doc, win, canvas, ctx *js.Object
var width, height int

func fitCanvas() {
    width = win.Get("innerWidth").Int()-20
    height = win.Get("innerHeight").Int()-20
    canvas.Set("width", width)
    canvas.Set("height", height)
    ctx.Set("fillStyle", "black")
    ctx.Call("fillRect", 0, 0, width, height)
}

func handleResize(event *js.Object) {
    fitCanvas()
    log.Printf("Resize: %dx%d", width, height)
}

func initGfx() {
    doc = js.Global.Get("document")
    canvas = doc.Call("createElement", "canvas")
    ctx = canvas.Call("getContext", "2d")
    win = js.Global.Get("window")
    win.Call("addEventListener", "resize", handleResize)
    doc.Get("body").Set("style", "background-color:grey")
    doc.Get("body").Call("appendChild", canvas)
    fitCanvas()
}

func main() {
    initGfx()
}
