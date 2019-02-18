package main
 import "fmt"
 import "flag"

 var name string
 func init()  {
	 flag.StringVar(&name,"name","everyone","The greeting object.")
 }

 func main()  {
	 flag.Parse()
	 fmt.Printf("hello,%s!",name)
 }