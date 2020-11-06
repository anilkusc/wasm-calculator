package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

/*
func printMessage(this js.Value, inputs []js.Value) interface{} {
	message := inputs[0].String()

	document := js.Global().Get("document")
	p := document.Call("createElement", "p")
	p.Set("innerHTML", message)
	document.Get("body").Call("appendChild", p)
	return nil
}*/
/*
func add(this js.Value, inputs []js.Value) interface{} {
	js.Global().Set("output", js.ValueOf(inputs[0].Int()+inputs[1].Int()))
	//fmt.Println(js.ValueOf(inputs[0].Int() + inputs[1].Int()).String())
	fmt.Println(inputs[0].Int() + inputs[1].Int())
	//fmt.Println(inputs[0].Int() + inputs[1].Int())
	var number = js.Global().Get("document").Call("getElementById", "number")
	newNumber := 2
	number.Set("innerHTML", strconv.Itoa(newNumber))

	return nil
}
*/

func add(a int, b int) int {
	return a + b
}
func sub(a int, b int) int {
	return a - b
}
func multi(a int, b int) int {
	return a * b
}
func divide(a int, b int) int {
	return a / b
}
func operate(this js.Value, inputs []js.Value) interface{} {

	window := js.Global()
	doc := window.Get("document")
	body := doc.Get("body")
	total := doc.Call("getElementById", "total")

	firstNumber := inputs[0].String()
	secondNumber := inputs[1].String()
	operator := inputs[2].String()
	first, _ := strconv.Atoi(firstNumber)
	second, _ := strconv.Atoi(secondNumber)
	var result int
	switch operator {
	case "+":
		result = add(first, second)
		break
	case "-":
		result = sub(first, second)
		break
	case "*":
		result = multi(first, second)
		break
	case "/":
		result = divide(first, second)
		break
	default:
		fmt.Println("Error")
		result = 100
		break
	}
	//js.Global().Get("document").Call("getElementById", "total").Set("innerHTML", strconv.Itoa(result))
	//js.Global().Get("body").Call("appendChild", "total")
	total.Set("innerHTML", strconv.Itoa(result))
	body.Call("appendChild", total)
	//fmt.Println(result)
	/*hello := doc.Call("createElement", "h1")
	hello.Set("innerText", strconv.Itoa(result))
	doc.Get("body").Call("appendChild", hello)*/
	return nil
}

func registerCallbacks() {

	js.Global().Set("operate", js.FuncOf(operate))
}

func main() {

	c := make(chan bool)
	registerCallbacks()
	<-c
}
