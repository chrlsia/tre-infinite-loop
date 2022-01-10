package main

import (
	"bufio"
	"fmt"
	"myapp/mylogger"
	"os"
)

func main()  {
	// read input for the user 5 times and write it 
	// to a log

	// create the reader
	reader:=bufio.NewReader(os.Stdin)

	// create the channel
	ch:=make(chan string)

	go mylogger.ListenForLog(ch)

	// give some instructions
	fmt.Println("Enter something")

	// read info from the user 5 times
	for i := 0; i < 5; i++ {
		fmt.Print("-> ")
		input,_:=reader.ReadString('\n')
		ch<-input
	}
}