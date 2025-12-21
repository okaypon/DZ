package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Item struct {
	Name    string
	Type    string
	Attack  int
	Defence int
	Health  int
}

var weapons = []Item{
	{"Алмазный меч", "оружие", 15, 0, 0},
	{"Глок пистолет", "оружие", 12, 0, 0},
	{"Дубинка", "оружие", 8, 0, 0},
	{"Арбалет", "оружие", 18, 0, 0},
	{"Лук", "оружие", 10, 0, 0},
}

var armors = []Item{
	{"Незеритовая броня", "броня", 0, 20, 0},
	{"Щит", "броня", 0, 15, 0},
	{"Алмазная броня", "броня", 0, 12, 0},
	{"Железная броня", "броня", 0, 8, 0},
	{"Золотая броня", "броня", 0, 10, 0},
}

var consumables = []Item{
	{"Зелье здоровья", "расходник", 0, 0, 30},
	{"Стимулятор", "расходник", 0, 0, 20},
	{"Зелье регенерации", "расходник", 0, 0, 15},
	{"Зелье исцеления", "расходник", 0, 0, 25},
	{"Золотое яблоко", "расходник", 0, 0, 10},
}

type Player struct {
	Name      string
	Health    int
	Inventory []Item
	Weapon    *Item
	Armor     *Item
}

type Enemy struct {
	Name   string
	Health int
	Item   *Item
}

func (p *Player) Equip() {
	if len(p.Inventory) == 0 {
		fmt.Println("Инвентарь пуст")
		return
	}

	fmt.Println("Инвентарь:")
	for i, item := range p.Inventory {
		if item.Type == "оружие" {
			fmt.Println(i, "-", item.Name, "(оружие, урон:", item.Attack, ")")
		} else if item.Type == "броня" {
			fmt.Println(i, "-", item.Name, "(броня, защита:", item.Defence, ")")
		} else {
			fmt.Println(i, "-", item.Name, "(расходник, лечение:", item.Health, "HP)")
		}
	}

	fmt.Print("Выберите предмет: ")
	var choice int
	fmt.Scan(&choice)

	if choice < 0 || choice >= len(p.Inventory) {
		fmt.Println("Неверный выбор")
		return
	}

	item := p.Inventory[choice]

	if item.Type == "расходник" {
		p.Health += item.Health
		fmt.Println("Использован", item.Name, "+", item.Health, "HP")
		fmt.Println("Здоровье:", p.Health)
		p.Inventory = append(p.Inventory[:choice], p.Inventory[choice+1:]...)
	} else if item.Type == "оружие" {
		if p.Weapon != nil {
			p.Inventory = append(p.Inventory, *p.Weapon)
		}
		p.Weapon = &item
		p.Inventory = append(p.Inventory[:choice], p.Inventory[choice+1:]...)
		fmt.Println("Экипировано оружие:", item.Name)
	} else if item.Type == "броня" {
		if p.Armor != nil {
			p.Inventory = append(p.Inventory, *p.Armor)
		}
		p.Armor = &item
		p.Inventory = append(p.Inventory[:choice], p.Inventory[choice+1:]...)
		fmt.Println("Экипирована броня:", item.Name)
	}
}

func (p *Player) ShowStats() {
	fmt.Println("\nИгрок:", p.Name)
	fmt.Println("Здоровье:", p.Health)

	if p.Weapon != nil {
		fmt.Println("Оружие:", p.Weapon.Name, "(урон +", p.Weapon.Attack, ")")
	} else {
		fmt.Println("Оружие: нет")
	}

	if p.Armor != nil {
		fmt.Println("Броня:", p.Armor.Name, "(защита +", p.Armor.Defence, ")")
	} else {
		fmt.Println("Броня: нет")
	}

	fmt.Println("Предметов в инвентаре:", len(p.Inventory))
}

func Fight(player *Player, enemy *Enemy) {
	fmt.Println("\nБой:", player.Name, "против", enemy.Name)

	playerAttack := 10
	if player.Weapon != nil {
		playerAttack += player.Weapon.Attack
	}

	enemyDefense := 0
	if player.Armor != nil {
		enemyDefense = player.Armor.Defence / 5
	}

	for player.Health > 0 && enemy.Health > 0 {
		enemy.Health -= playerAttack
		fmt.Println(player.Name, "атакует на", playerAttack, "урона")

		if enemy.Health <= 0 {
			fmt.Println(enemy.Name, "побежден!")
			break
		}

		enemyDamage := 10 - enemyDefense
		if enemyDamage < 1 {
			enemyDamage = 1
		}
		player.Health -= enemyDamage
		fmt.Println(enemy.Name, "атакует на", enemyDamage, "урона")
		fmt.Println("У", player.Name, "осталось", player.Health, "HP")
	}

	if player.Health > 0 {
		fmt.Println(player.Name, "победил!")
		if enemy.Item != nil {
			player.Inventory = append(player.Inventory, *enemy.Item)
			fmt.Println("Получен предмет:", enemy.Item.Name)
		}
	} else {
		fmt.Println(player.Name, "проиграл")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	player := &Player{
		Name:   "Арбус",
		Health: 100,
	}

	player.Inventory = []Item{
		weapons[rand.Intn(len(weapons))],
		armors[rand.Intn(len(armors))],
		consumables[rand.Intn(len(consumables))],
	}

	enemy := &Enemy{
		Name:   "Зомби",
		Health: 60,
	}

	if rand.Intn(2) == 1 {
		items := [][]Item{weapons, armors, consumables}
		itemType := rand.Intn(3)
		itemList := items[itemType]
		enemy.Item = &itemList[rand.Intn(len(itemList))]
	}

	fmt.Println("ИГРА НАЧАЛАСЬ")
	player.ShowStats()

	fmt.Println("\nЭкипировка")
	player.Equip()
	player.Equip()

	Fight(player, enemy)

	fmt.Println("\nИГРА ЗАКОНЧЕНА")
	player.ShowStats()

	if len(player.Inventory) > 0 {
		fmt.Println("\nОстальные предметы:")
		for i, item := range player.Inventory {
			fmt.Println(i, "-", item.Name)
		}
	}
}
