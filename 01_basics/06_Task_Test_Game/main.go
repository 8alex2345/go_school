package main

import (
	"errors"
	"strings"
)

// Метод старт и стоп от игры
type Game struct {
	Rooms        map[string]Room
	User         User
	DoorState    bool
	KitchenState bool
}
type Room struct {
	description string
	tableItems  []string
	chairItems  []string
	exist       []string
}

func InitGame() *Game {
	return &Game{
		Rooms: map[string]Room{
			"кухня": {
				description: "надо собрать рюкзак и идти в универ.",
				tableItems:  []string{"чай"},
				exist:       []string{"Коридор"},
			},
			"коридор": {
				description: "Ничего интересного",
				exist:       []string{"Кухня", "Комната", "Улица"},
			},
			"комната": {
				description: "Ты в своей комнате",
				tableItems:  []string{"ключи", "конспекты"},
				chairItems:  []string{"рюкзак"},
				exist:       []string{"Кухня", "Комната", "Улица"},
			},
			"улица": {
				description: "На улице весна",
				exist:       []string{"Домой"},
			},
		},
		User: User{
			CurrentRoom: "кухня",
		},
		DoorState:    false,
		KitchenState: true,
	}
}
func StartGame(g *Game) error {
	if g.User.CurrentRoom != "кухня" {
		return errors.New("Игру можно начать только с кухни")
	}
	g.KitchenState = true
	return nil
}

type User struct {
	CurrentRoom string
}

func (g *Game) Look() string { // осмотреться
	room, ok := g.Rooms[g.User.CurrentRoom]
	if !ok {
		return "Неизвестная комната"
	}
	result := "Ты находишься на " + g.User.CurrentRoom
	if len(room.chairItems) > 0 {
		result = result + ", на стуле: " + strings.Join(room.chairItems, ", ")
	}
	if len(room.tableItems) > 0 {
		result = result + ", на столе: " + strings.Join(room.tableItems, ", ")
	}
	result += room.description
	if len(room.exist) > 0 {
		result += "Можно пройти - " + strings.Join(room.exist, ", ")
	}
	return result

}
func (u *User) Walk() string { // ходить

}
func (u *User) PutOn() string { // надеть

}
func (u *User) Take() string { // Взять

}
func (u *User) Use() string { // Использовать

}
func hendlCommand(command string, user *User) {

}

func main() {

}
