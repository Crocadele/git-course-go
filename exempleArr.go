package main

import (
    "fmt"
    "time"
)

func print(c chan string){

    for i := 0; i < 10;i++{

        c <- "ping"
    }

}

func print1(c chan string){

    for i := 0;i < 10;i++ {

        c <- "pong" 
        time.Sleep(time.Second * 3)
       

}

}

func printLn(c chan string){

    for i := 0;i < 10;i++ {

    msg := <- c

    fmt.Println(msg)
    time.Sleep(time.Second * 1)
}

}


func main() {
    c := make(chan string)

    go print(c)
    go print1(c)
    go printLn(c)

    var input string
    fmt.Scanln(&input)
}

















