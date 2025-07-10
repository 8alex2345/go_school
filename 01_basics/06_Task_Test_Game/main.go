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
	UserState   string
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
				description: "надо собрать рюкзак и идти в универ",
				tableItems:  []string{"чай"},
				exist:       []string{"коридор"},
			},
			"коридор": {
				description: "ничего интересного",
				exist:       []string{"кухня", "комната", "улица"},
			},
			"комната": {
				description: "ты в своей комнате",
				tableItems:  []string{"ключи", "конспекты"},
				chairItems:  []string{"рюкзак"},
				exist:       []string{"коридор"},
			},
			"улица": {
				description: "на улице весна",
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
	result := []string{}
	room, ok := g.Rooms[g.User.CurrentRoom]
	if !ok {
		return "неизвестная комната"
	} else if g.User.CurrentRoom == "кухня" && g.KitchenState {
		result = append(result, "ты находишься на кухне")
	}

	if len(room.tableItems) > 0 {
		result = append(result, "на столе: "+strings.Join(room.tableItems, ", "))
	}

	if len(room.chairItems) > 0 {
		result = append(result, "на стуле: "+strings.Join(room.chairItems, ", "))
	}

	if len(room.chairItems) == 0 && len(room.tableItems) == 0 {
		result = []string{"пустая комната"}
	}

	if room.description != "" {
		if g.User.CurrentRoom == "кухня" && g.User.BackPack {
			result = append(result, "надо идти в универ")
		} else if g.User.CurrentRoom == "кухня" && !g.User.BackPack {
			result = append(result, strings.TrimSpace(room.description))
		} else if g.User.CurrentRoom != "кухня" && len(result) == 0 {
			result = append(result, strings.TrimSpace(room.description))
		}
	}
	res := strings.Join(result, ", ")
	if len(room.exist) > 0 {
		res += ". можно пройти - " + strings.Join(room.exist, ", ")
	}
	return res

}
func (g *Game) Walk(existName string) string { // ходить
	room, ok := g.Rooms[g.User.CurrentRoom]
	if !ok {
		return "неизвестная комната"
	}
	for _, exit := range room.exist {
		if exit == existName {
			if existName == "улица" && !g.DoorState {
				return "дверь закрыта"
			}
			g.User.CurrentRoom = existName
			newRoom, ok := g.Rooms[existName]
			if !ok {
				return "неизвестная комната"
			}
			if len(newRoom.tableItems) == 1 && existName == "кухня" {
				result := "кухня, ничего интересного" + ". можно пройти - " + strings.Join(newRoom.exist, ", ")
				return result

			}
			result := newRoom.description + ". можно пройти - " + strings.Join(newRoom.exist, ", ")
			return result
		}
	}

	return "нет пути в " + existName
}
func (g *Game) PutOn() string { // надеть
	room, ok := g.Rooms[g.User.CurrentRoom]
	if !ok {
		return "неизвестная комната"
	}
	for i, item := range room.chairItems {
		if item == "рюкзак" {
			room.chairItems = append(room.chairItems[:i], room.chairItems[i+1:]...)
			g.Rooms[g.User.CurrentRoom] = room
			g.User.BackPack = true
			res := "вы надели: " + item
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
			res := "предмет добавлен в инвентарь: " + item
			return res
		}
	}
	return "нет такого"

}
func (g *Game) Use(itemName string) string { // Использовать
	words := strings.Fields(itemName)
	if itemName != "ключи дверь" {
		if words[0] != "ключи" {
			return "нет предмета в инвентаре - " + words[0]
		}
		return "не к чему применить"
	}

	for _, item := range g.User.Inventory {
		if item == "ключи" {
			g.DoorState = true
			return "дверь открыта"
		}

	}
	return "нет предмета в инвентаре - " + words[0]

}
func handleCommand(command string, g *Game) string {
	words := strings.Fields(command)
	if len(words) == 0 {
		return "пустая строка"
	}
	switch words[0] {
	case "осмотреться":
		return g.Look()
	case "идти":
		if len(words) > 1 {
			return g.Walk(strings.Join(words[1:], " "))
		}
		return "не указано куда идти"
	case "надеть":
		if len(words) > 1 {
			return g.PutOn()
		}
		return "не указано что надеть"
	case "взять":
		if len(words) > 1 {
			return g.Take(strings.Join(words[1:], " "))
		}
		return "не указано что взять"

	case "применить":
		if len(words) > 1 {
			return g.Use(strings.Join(words[1:], " "))
		}
		return "не указано что применить"
	default:
		return "неизвестная команда"

	}
}

func main() {
	game := InitGame()
	fmt.Println("Игра началась")
	var input string
	for {
		fmt.Scan(&input)
		if input == "выход" {
			break
		}
		fmt.Println(handleCommand(input, game))
	}
}
