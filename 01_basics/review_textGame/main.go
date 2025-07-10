package main

import (
	"errors"
	"fmt"
	"strings"
)

type Room struct {
	description string
	chairItems  []string
	tableItems  []string
	exit        []string
}

type RoomsState interface {
	GetRoom(name string) (Room, error)
	SetRoom(name string, room Room)
}
type RoomStore struct {
	room map[string]Room
}

func (rs *RoomStore) GetRoom(name string) (Room, error) {
	room, ok := rs.room[name]
	if !ok {
		return Room{}, errors.New("неизвестная комната")
	}
	return room, nil
}

func (rs *RoomStore) SetRoom(name string, room Room) {
	rs.room[name] = room
}

type User struct {
	CurrentRoom string
	BackPack    bool
	Inventory   []string
}
type UserState interface {
	GetCurrentRoom() string
	SetCurrentRoom(name string)
	HasBackPack() bool          // имеет ли рюкзак
	PutOnBackPack()             // надеть рюкзак
	AddToInventory(item string) // добавить в иневентраь
	HasItem(item string) bool   // имеет ли предмет

}

func (u *User) GetCurrentRoom() string {
	return u.CurrentRoom
}
func (u *User) SetCurrentRoom(name string) {
	u.CurrentRoom = name
}
func (u *User) HasBackPack() bool {
	return u.BackPack
}
func (u *User) PutOnBackPack() {
	u.BackPack = true
}
func (u *User) AddToInventory(item string) {
	u.Inventory = append(u.Inventory, item)

}
func (u *User) HasItem(item string) bool {
	for _, it := range u.Inventory {
		if it == item {
			return true
		}
	}
	return false
}

type Game struct {
	Rooms        RoomStore
	User         UserState
	KitchenState bool
	DoorState    bool
}
type GameLogic interface {
	Start() error
	Look() string
	Walk(exitName string) string
	PutOn(rooms RoomsState) string
	Take(itemName string, rooms RoomsState) string
	Use(inemName string) string
}

func (g *Game) Start() error {
	if !g.KitchenState && g.User.GetCurrentRoom() != "кухня" {
		return errors.New("игру можно начать только с кухни")
	}
	return nil
}
func (g *Game) Look() string {
	parts := []string{}
	roomName := g.User.GetCurrentRoom()
	room, ok := g.Rooms.GetRoom(roomName)
	if ok != nil {
		return "неизвестная комната"
	}

	if g.KitchenState && roomName == "кухня" {
		parts = append(parts, "ты находишься на кухне")
	}
	if len(room.tableItems) > 0 {
		parts = append(parts, "на столе: "+strings.Join(room.tableItems, ", "))
	}
	if len(room.chairItems) > 0 {
		parts = append(parts, "на стуле: "+strings.Join(room.chairItems, ", "))
	}
	if len(room.chairItems) == 0 && len(room.tableItems) == 0 {
		parts = []string{"пустая комната"}
	}

	switch room.description != "" {
	case roomName == "кухня" && g.User.HasBackPack():
		parts = append(parts, "надо идти в универ")
	case roomName == "кухня" && !g.User.HasBackPack():
		parts = append(parts, strings.TrimSpace(room.description))
	case roomName != "кухня" && len(parts) == 0:
		parts = append(parts, strings.TrimSpace(room.description))
	}
	// if room.description != "" {
	// 	if roomName == "кухня" && g.User.HasBackPack() {
	// 		parts = append(parts, "надо идти в универ")
	// 	} else if roomName == "кухня" && !g.User.HasBackPack() {
	// 		parts = append(parts, strings.TrimSpace(room.description))
	// 	} else if roomName != "кухня" && len(parts) == 0 {
	// 		parts = append(parts, strings.TrimSpace(room.description))
	// 	}
	// }
	res := strings.Join(parts, ", ")
	if len(room.exit) > 0 {
		res += ". можно пройти - " + strings.Join(room.exit, ", ")
	}
	return res

}
func (g *Game) Walk(exitName string) string {
	roomName := g.User.GetCurrentRoom()
	room, err := g.Rooms.GetRoom(roomName)
	if err != nil {
		return "неизвестная комната"
	}
	for _, exit := range room.exit {
		if exit == exitName {
			if exitName == "улица" && !g.DoorState {
				return "дверь закрыта"
			}
			g.User.SetCurrentRoom(exitName)
			newRoom, err := g.Rooms.GetRoom(exitName)
			if err != nil {
				return "неизвестная комната"
			}
			if len(newRoom.tableItems) == 1 && exitName == "кухня" {
				result := "кухня, ничего интересного" + ". можно пройти - " + strings.Join(newRoom.exit, ", ")
				return result

			}
			result := newRoom.description + ". можно пройти - " + strings.Join(newRoom.exit, ", ")
			return result
		}

	}
	return "нет пути в " + exitName
}

func (u *User) PutOn(rooms RoomsState) string {
	roomName := u.GetCurrentRoom()
	room, err := rooms.GetRoom(roomName)
	if err != nil {
		return "неизвестная комната"
	}
	for i, item := range room.chairItems {
		if item == "рюкзак" {
			room.chairItems = append(room.chairItems[:i], room.chairItems[i+1:]...)
			rooms.SetRoom(roomName, room)
			u.PutOnBackPack()
			return "вы надели: " + item
		}
	}
	return "здесь нет рюкзака"
}

func (u *User) Take(itemName string, rooms RoomsState) string {
	if !u.HasBackPack() {
		return "некуда класть"
	}
	roomName := u.GetCurrentRoom()
	room, err := rooms.GetRoom(roomName)
	if err != nil {
		return "неизвестная комната"
	}
	for i, item := range room.tableItems {
		if item == itemName {
			u.AddToInventory(item)
			room.tableItems = append(room.tableItems[:i], room.tableItems[i+1:]...)
			rooms.SetRoom(roomName, room)
			return "предмет добавлен в инвентарь: " + item
		}
	}
	return "нет такого"
}

func (g *Game) Use(itemName string) string {
	words := strings.Fields(itemName)
	if itemName != "ключи дверь" {
		if words[0] != "ключи" {
			return "нет предмета в инвентаре - " + words[0]
		}
		return "не к чему применить"
	}
	if g.User.HasItem("ключи") {
		g.DoorState = true
		return "дверь открыта"

	}
	return "нет предмета в инвентаре - " + words[0]
}

func InitGame() *Game {
	rooms := map[string]Room{
		"кухня": {
			description: "надо собрать рюкзак и идти в универ",
			tableItems:  []string{"чай"},
			exit:        []string{"коридор"},
		},
		"коридор": {
			description: "ничего интересного",
			exit:        []string{"кухня", "комната", "улица"},
		},
		"комната": {
			description: "ты в своей комнате",
			tableItems:  []string{"ключи", "конспекты"},
			chairItems:  []string{"рюкзак"},
			exit:        []string{"коридор"},
		},
		"улица": {
			description: "на улице весна",
			exit:        []string{"домой"},
		},
	}
	user := &User{
		CurrentRoom: "кухня",
		BackPack:    false,
		Inventory:   []string{},
	}
	return &Game{
		Rooms:        RoomStore{room: rooms},
		User:         user,
		KitchenState: true,
		DoorState:    false,
	}
}
func handleCommand(command string, g *Game, u *User) string {
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
			return u.PutOn(&g.Rooms)
		}
		return "не указано что надеть"
	case "взять":
		if len(words) > 1 {
			return u.Take(strings.Join(words[1:], " "), &g.Rooms)
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
	user := game.User.(*User)
	fmt.Println("Игра началась")
	var input string
	for {
		fmt.Scan(&input)
		if input == "выход" {
			break
		}
		fmt.Println(handleCommand(input, game, user))
	}
}
