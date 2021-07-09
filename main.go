package main

import (
    "github.com/gopherjs/gopherjs/js"
    "log"
    "time"
    "strconv"
)

var doc, win, canvas, ctx *js.Object
var width, height int

func clearCanvas() {
    ctx.Call("clearRect", 0, 0, width, height)
}

func fitCanvas() {
    width = win.Get("innerWidth").Int()
    height = win.Get("innerHeight").Int()
    canvas.Set("width", width)
    canvas.Set("height", height)
    clearCanvas()
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
    doc.Get("body").Set("style", "background-color:black;touch-action:none;-webkit-overflow-scrolling:none;overflow:hidden;overscroll-behavior:none;padding:0;margin:0")
    doc.Get("body").Call("appendChild", canvas)
    fitCanvas()
}

func main() {
    initGfx()
    ctx.Set("fillStyle", "white")
    for i := 0;; i++ {
        clearCanvas()
        ctx.Set("font", strconv.Itoa(25+i)+"px Monospace")
        textSize := ctx.Call("measureText", "Hello world")
        textWidth := textSize.Get("width").Int()
        textHeight := textSize.Get("height").Int()
        ctx.Call("fillText", "Hello world", width/2-textWidth/2, height/2-textHeight/2)
        if textWidth >= int(0.7*float64(width)) || textHeight >= int(0.7*float64(height)) {
            break
        }
        time.Sleep(time.Millisecond*50)
    }
}
