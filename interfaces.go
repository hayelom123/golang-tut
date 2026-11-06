package main

import ("fmt"
"math")


//=========================== shapes =========================
type Shape interface{
	Area() float64
}


type Circle struct{
	Radius float64
}

func(c Circle) Area()float64{
	return math.Pi*c.Radius* c.Radius
}

type Rectangle struct {
	Width float64
	Height float64
}

func(r Rectangle) Area() float64{
	return r.Width*r.Height
}

func PrintArea(s Shape){
	fmt.Println(s.Area())
}




func main(){

	//circle 
	c:=Circle{Radius:3}

	// rectangle
	r:=Rectangle{Width:4,Height:3}

	PrintArea(c)
	PrintArea(r)

	// logger example 

	consoleLog:=ConsoleLogger{}

	prefixLog:=PrefixLogger{Prefix:"[INFO]"}

	runLog(consoleLog,"Console log")
	runLog(prefixLog,"Prefix log")

	// writer example

	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))


	// increment 
    
	myInt :=IntCounter(0)

	for i:=0; i<10;i++ {
		fmt.Println(myInt.increment())
	}


}


//=========================== logers =========================
type Logger interface{
	Log(message string)
}

type ConsoleLogger struct{
   
}

func(l ConsoleLogger)Log(message string){
	fmt.Println(message)
}

type PrefixLogger struct{
     Prefix string
}

func(l PrefixLogger) Log(message string){
	 fmt.Println(l.Prefix,message)
}

func runLog(l Logger,message string){
	l.Log(message)
}

//=========================== writters =========================

type Writer interface{
	Write([]byte)(int,error)
}

type ConsoleWriter struct{}

func(cw ConsoleWriter) Write(data []byte) (int,error){
	 n,err:=fmt.Println(string(data))
	 return n,err
}

//=========================== Counter =========================

type Incrementer interface{
	increment()int
}

type IntCounter int

func(c *IntCounter) increment()int{
	*c++
	return int(*c)
}