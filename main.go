package main

import (
    "github.com/gopherjs/gopherjs/js"
    //"log"
)

var date *js.Object
var h1 *js.Object

func fn() {
    h1.Set("innerHTML", date.New())
}

func main() {
    doc := js.Global.Get("document")
    h1 = doc.Call("createElement", "h1")
    date = js.Global.Get("Date")
    h1.Set("innerHTML", date.New())
    h1.Set("id", "clock")
    doc.Get("body").Call("appendChild", h1)
    js.Global.Call("setInterval", fn, 1000)
}
