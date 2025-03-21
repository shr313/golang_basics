package exercise

import (
	"fmt"
)

// func Exercise() {

//1. Creat a PLayer Struct ( Name And Inventory)
//2. also creat an iten struct( Name  and Type)
//3 Methods
//  PickUpitem: Add an Item
//  DropItem: Remove an Item
//  UseItem: Depending On the item It will Print Different Things
// Define the Item struct

type Item struct {
	Name string
	Type string // For example: "ball", "racket", "shoes"
}

// Define the Player struct
type Player struct {
	Name      string
	Inventory []Item // Player's inventory is a slice of items (sports equipment)
}

// Method to pick up an item and add it to the inventory
func (p *Player) PickUpItem(item Item) {
	p.Inventory = append(p.Inventory, item)
	fmt.Printf("%s picked up a %s.\n", p.Name, item.Name)
}

// Method to drop an item from the inventory
func (p *Player) DropItem(itemName string) {
	for i, item := range p.Inventory {
		if item.Name == itemName {
			// Remove item from the slice
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			fmt.Printf("%s dropped the %s.\n", p.Name, itemName)
			return
		}
	}
	fmt.Printf("%s not found in inventory.\n", itemName)
}

// Method to use an item (depending on the type, print different messages)
func (p *Player) UseItem(itemName string) {
	for _, item := range p.Inventory {
		if item.Name == itemName {
			switch item.Type {
			case "ball":
				fmt.Printf("%s is playing with the %s!\n", p.Name, item.Name)
			case "racket":
				fmt.Printf("%s is practicing with the %s!\n", p.Name, item.Name)
			case "shoes":
				fmt.Printf("%s is wearing the %s to run!\n", p.Name, item.Name)
			default:
				fmt.Printf("%s used the %s.\n", p.Name, item.Name)
			}
			return
		}
	}
	fmt.Printf("%s not found in inventory.\n", itemName)
}

func GameExercise() {
	// Create a player
	player := Player{Name: "John"}

	// Create some sports equipment items
	basketball := Item{Name: "Basketball", Type: "ball"}
	soccerBall := Item{Name: "Soccer Ball", Type: "ball"}
	tennisRacket := Item{Name: "Tennis Racket", Type: "racket"}
	runningShoes := Item{Name: "Running Shoes", Type: "shoes"}

	// Player picks up items
	player.PickUpItem(basketball)
	player.PickUpItem(soccerBall)
	player.PickUpItem(tennisRacket)
	player.PickUpItem(runningShoes)

	// Player uses an item
	player.UseItem("Basketball")
	player.UseItem("Soccer Ball")
	player.UseItem("Tennis Racket")
	player.UseItem("Running Shoes")

	// Player drops an item
	player.DropItem("Tennis Racket")
	player.DropItem("Running Shoes")
}
