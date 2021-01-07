package main

import (
	"fmt"
)


type Player struct {
    Name string
    Hp int
    Attack int
}

func newPlayer(name string, hp int, attack int) *Player {
    return &Player{
        Name: name,
        Hp: hp,
        Attack: attack,
    }
}

func main(){
    p1 := newPlayer(
        "Tom",
        100,
        50,
    )
    p2 := newPlayer(
        "Jerry",
        50,
        100,
    )
    
    fmt.Printf("%+v\n", p1)
    fmt.Printf("%+v\n", p2)
    fmt.Printf("%p, %p\n", p1, p2)
}