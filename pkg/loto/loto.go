package loto

import (
	"log"
	"math/rand"
	"time"
)

type User struct {
	UserID   string
	ID       []int64
	PriceWin Price
}

type Price []string

//постановка задачи линейного программирования
func (u *User) WinPrice(id []int64) (index int, indexprice int) {
	rand.Seed(int64(time.Now().Second()))
	if len(id) < 2 {
		log.Print("error: len: id >= 2")
	}
	index = rand.Intn(len(id)) + 1
	indexprice = rand.Intn(len(u.PriceWin) + 1)
	return index, indexprice
}
