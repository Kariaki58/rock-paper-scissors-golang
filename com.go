package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Winner(User, Computer string) string{
	if (User == "scissors") && (Computer == "paper") || (User == "paper") && (Computer == "rock") || (User == "rock") && (Computer == "scissors"){
		return "User Wins"
	}
	if (User == "paper") && (Computer== "scissors") || (User=="rock") && (Computer == "paper") || (User == "scissors") && (Computer== "rock"){
		return "Computer Wins"
	}
	return "Tie"
}
func recovery_fun() {
	a := recover()
	if a != nil{
		fmt.Println("error")
	}
}

func User_choice() int{
	myValue := ""
	fmt.Printf("r for Rock, p for Paper, s for scissors:  ")
	fmt.Scan(&myValue);
	r_value := 0 // default value

	if myValue == "r"{
		r_value = 0
	}else if(myValue == "p"){
		r_value = 1
	}else if(myValue == "s"){
		r_value = 2
	}else{
		defer recovery_fun()
		panic("choice not avialable")
	}
	return r_value
}

func computer_choice() int{

	hidden_choice := rand.NewSource(time.Now().UnixNano())
	rand_choice := rand.New(hidden_choice)
	com_choice := rand_choice.Intn(3);
	return com_choice
}

func main(){
	game := [3]string{"rock", "paper", "scissors"};
	
	play_again := "y"

	for play_again != "n"{
		User := User_choice()

		Computer := computer_choice() 

		str1 := game[User]
		str2 := game[Computer]

		fmt.Println(Winner(str1, str2));
		fmt.Printf("play again? y/yes: n/no: ")
		fmt.Scan(&play_again)
		
	}
}
