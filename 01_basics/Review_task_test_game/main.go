package main

import (
	"errors"
	"fmt"
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
type User struct {
	CurrentRoom string
	BackPack    bool
	Inventory   []string
}

func StartGame(g *Game) error {
	if g.User.CurrentRoom != "кухня" {
		return errors.New("игру можно начать только с кухни")
	}
	g.KitchenState = true
	return nil
}

func InitGame() *Game {
	return &Game{
		Rooms: map[string]Room{
			"кухня": {
				description: ", надо собрать рюкзак и идти в универ.",
				tableItems:  []string{"чай"},
				exist:       []string{"коридор"},
			},
			"коридор": {
				description: " Ничего интересного ",
				exist:       []string{"кухня", "комната", "улица"},
			},
			"комната": {
				description: " Ты в своей комнате ",
				tableItems:  []string{"ключи", "конспекты"},
				chairItems:  []string{"рюкзак"},
				exist:       []string{"кухня", "комната", "улица"},
			},
			"улица": {
				description: " на улице весна ",
				exist:       []string{"домой"},
			},
		},
		User: User{
			CurrentRoom: "кухня",
		},
		DoorState:    false,
		KitchenState: true,
	}
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
func (g *Game) Walk(existName string) string { // ходить
	room, ok := g.Rooms[g.User.CurrentRoom]
	if !ok {
		return "Неизвестная комната"
	}
	for _, exit := range room.exist {
		if exit == existName {
			g.User.CurrentRoom = existName

			return g.Look()
		}
	}

	return "нет пути в " + existName
}
func (g *Game) PutOn() string { // надеть
	room, ok := g.Rooms[g.User.CurrentRoom]
	if !ok {
		return "Неизвестная комната"
	}
	for i, item := range room.chairItems {
		if item == "рюкзак" {
			room.chairItems = append(room.chairItems[:i], room.chairItems[i+1:]...)
			g.Rooms[g.User.CurrentRoom] = room
			g.User.BackPack = true
			res := "Вы надели " + item
			return res
		}
	}
	return "здесь нет рюкзака"
}
func (g *Game) Take(itemName string) string { // Взять
	if !g.User.BackPack {
		return "некуда класть"
	}
	room, ok := g.Rooms[g.User.CurrentRoom]
	if !ok {
		return "неизвестная комната"
	}
	for i, item := range room.tableItems {
		if item == itemName {
			g.User.Inventory = append(g.User.Inventory, item)
			room.tableItems = append(room.tableItems[:i], room.tableItems[i+1:]...)
			g.Rooms[g.User.CurrentRoom] = room
			res := "Предмет добавлен в инвентарь: " + item
			return res
		}
	}
	return "Нет такого"

}
func (g *Game) Use(itemName string) string { // Использовать
	if itemName != "Применить ключи дверь" {
		return "не к чему применять"
	}
	for _, item := range g.User.Inventory {
		if item == "ключи" {
			g.DoorState = true
			return "Дверь открыта"
		}
	}
	return "Нет предмета в инвентаре"
}
func hendleCommand(g *Game, command string) string {
	commands := map[string]func() string{
		"осмотреться": g.Look,
		"идти коридор": func() string {
			return g.Walk("коридор")
		},
		"идти комната": func() string {
			return g.Walk("комната")
		},
		"надеть рюкзак": g.PutOn,
		"взять ключи": func() string {
			return g.Take("ключи")
		},
		"взять конспекты": func() string {
			return g.Take("конспекты")
		},
		"применить ключи дверь": func() string {
			return g.Use("применить ключи дверь")
		},
		"идти улица": func() string {
			return g.Walk("улица")
		},
	}
	if cmd, ok := commands[command]; ok {
		return cmd()
	}
	return "не известная команда"

}

func main() {
	game := InitGame()
	fmt.Println(hendleCommand(game, "осмотреться"))
	fmt.Println(hendleCommand(game, "идти коридор"))
}
