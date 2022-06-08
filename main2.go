package main2

import "fmt"

func main() {
	fmt.Println("Hello World")

	service GreetService{
		//unary
		rpc Greet(GreetRequest) returns (GreetResponse){};

		//server streaming
		rpc GreetManyTimes(GreetRequest) returns (stream GreetResponse){};

		//Client streaming
		rpc LongGreet(stream GreetRequest) returns (stream GreetResponse){};

		//Bi directional streaming
		rpc GreetEveryone(stream GreetRequest) returns (stream GreetResponse){};

	}
}
