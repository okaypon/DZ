package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Golova = 1
	Puzo   = 2
	Nogi   = 3
)

func BodyPart(b int) string {
	switch b {
	case Golova:
		return "голова"
	case Puzo:
		return "живот"
	case Nogi:
		return "ноги"
	default:
		return "живот"
	}
}

type Item struct {
	Name    string
	Type    string
	Attack  int
	Defence int
	PlusHP  int
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

type Personaj interface {
	Hit() int
	Block() int
	GetWeapon() *Weapon
	RandomWeapon()
}

type Weapon struct {
	Name   string
	MinDmg int
	MaxDmg int
}

type Player struct {
	Name      string
	HP        int
	Strength  int
	hit       int
	block     int
	weapon    *Weapon
	Inventory []Item
	Equipment map[string]*Item
}

func (p *Player) GetWeapon() *Weapon {
	return p.weapon
}

func (p *Player) Hit() int {
	return p.hit
}

func (p *Player) Block() int {
	return p.block
}

func (p *Player) RandomWeapon() {
	weapons := []Weapon{
		{"Glock-18", 15, 25},
		{"AK-47", 28, 100},
		{"MAC-10", 28, 69},
		{"AWP", 84, 100},
		{"Scout", 60, 100},
	}
	p.weapon = &weapons[rand.Intn(len(weapons))]
}

func (p *Player) TakeOff() {
	if len(p.Equipment) == 0 {
		fmt.Println("Нет надетых предметов")
		return
	}

	fmt.Println("Надетые предметы:")
	i := 1
	equipList := make(map[int]string)
	for equipType, item := range p.Equipment {
		fmt.Println(i, "-", item.Name, "(", equipType, ")")
		equipList[i] = equipType
		i++
	}

	fmt.Print("Выберите предмет для снятия (0 - отмена): ")
	var choice int
	fmt.Scan(&choice)

	if choice == 0 {
		return
	}

	if equipType, exists := equipList[choice]; exists {
		item := p.Equipment[equipType]
		p.Inventory = append(p.Inventory, *item)
		delete(p.Equipment, equipType)
		fmt.Println("Снят предмет:", item.Name)
	} else {
		fmt.Println("Неверный выбор")
	}
}

func (p *Player) Equip() {
	if len(p.Inventory) == 0 {
		fmt.Println("Инвентарь пуст")
		return
	}

	fmt.Println("Инвентарь:")
	for i, item := range p.Inventory {
		fmt.Print(i, "-", item.Name)
		if item.Type == "оружие" {
			fmt.Println(" (оружие, урон:", item.Attack, ")")
		} else if item.Type == "броня" {
			fmt.Println(" (броня, защита:", item.Defence, ")")
		} else {
			fmt.Println(" (расходник, лечение:", item.PlusHP, "HP)")
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
		p.HP += item.PlusHP
		if p.HP > 100 {
			p.HP = 100
		}
		fmt.Println("Использован", item.Name, "+", item.PlusHP, "HP")
		fmt.Println("Здоровье:", p.HP)
		p.Inventory = append(p.Inventory[:choice], p.Inventory[choice+1:]...)
	} else if item.Type == "оружие" {
		if p.Equipment["оружие"] != nil {
			fmt.Println("Уже экипировано оружие:", p.Equipment["оружие"].Name)
			fmt.Println("Сначала снимите текущее оружие")
			return
		}
		p.Equipment["оружие"] = &item
		p.Inventory = append(p.Inventory[:choice], p.Inventory[choice+1:]...)
		fmt.Println("Экипировано оружие:", item.Name)
	} else if item.Type == "броня" {
		if p.Equipment["броня"] != nil {
			fmt.Println("Уже экипирована броня:", p.Equipment["броня"].Name)
			fmt.Println("Сначала снимите текущую броню")
			return
		}
		p.Equipment["броня"] = &item
		p.Inventory = append(p.Inventory[:choice], p.Inventory[choice+1:]...)
		fmt.Println("Экипирована броня:", item.Name)
	}
}

func generateRoundStory(round int, p *Player, e *Enemy) string {
	switch round {
	case 1:
		return "ПЕРВЫЙ КОНТАКТ - Бойцы занимают позиции на карте de_dust2..."
	case 3:
		return "ЭКО-РАУНД - Противники экономят, но все равно сходятся в бою..."
	case 5:
		return "ФОРС-БУЙ - Обе команды покупают нормальное оружие за раунд..."
	case 7:
		return "КЛАТЧ-СИТУАЦИЯ - Все зависит от игроков..."
	case 9:
		return "МАТЧ-ПОЙНТ - Следующий раунд может решить исход всей игры..."
	}
	stories := []string{
		"ТЕРРОРИСТЫ ВЫСТАВЛЕНИЯ A - Контр-террористы проверяют углы",
		"ВРЫВ B",
		"СНАЙПЕРСКАЯ ДУЭЛЬ - " + p.weapon.Name + " против " + e.weapon.Name + " на миде",
		"ЗАКЛАДКА БОМБЫ - Террористы движутся к точке A",
		"ОБЕЗВРЕЖИВАНИЕ - Контр-террористы ищут бомбу",
		"ФЛЭШБЭК - Противники ослеплены светошумовыми",
	}

	return stories[rand.Intn(len(stories))]
}

func describeRoundResult(p *Player, e *Enemy, pHit, eHit int, pDamage, eDamage int) {
	if pDamage > 0 {
		fmt.Println(p.Name, "наносит противнику из", p.weapon.Name, "[", pDamage, "урона]")
	}

	if eDamage > 0 {
		fmt.Println(e.Name, "наносит противнику из", e.weapon.Name, "[", eDamage, "урона]")
	}
}

type Enemy struct {
	Name     string
	HP       int
	Strength int
	hit      int
	block    int
	weapon   *Weapon
	Item     *Item
}

func (e *Enemy) GetWeapon() *Weapon {
	return e.weapon
}

func (e *Enemy) Hit() int {
	return e.hit
}

func (e *Enemy) Block() int {
	return e.block
}

func PlayerHitBlock(p *Player) {
	var attack, defence int
	fmt.Println("Выберите часть тела для атаки")
	fmt.Println("1-Голова, 2-Живот, 3-Ноги")
	fmt.Scan(&attack)

	if attack < 1 || attack > 3 {
		attack = Puzo
		fmt.Println("Неверный выбор! Установлено значение по умолчанию: живот")
	}
	p.hit = attack

	fmt.Println("Выберите часть тела для защиты")
	fmt.Println("1-Голова, 2-Живот, 3-Ноги")
	fmt.Scan(&defence)

	if defence < 1 || defence > 3 {
		defence = Puzo
		fmt.Println("Неверный выбор! Установлено значение по умолчанию: живот")
	}
	p.block = defence
}

func EnemyHitBlock(e *Enemy) {
	e.hit = rand.Intn(3) + 1
	e.block = rand.Intn(3) + 1
}

func (e *Enemy) RandomWeapon() {
	weapons := []Weapon{
		{"USP-S", 15, 25},
		{"M4A4", 28, 89},
		{"M4A1-S", 28, 89},
		{"FAMAS", 25, 84},
		{"AWP", 84, 100},
	}
	e.weapon = &weapons[rand.Intn(len(weapons))]
}

func Round(p *Player, e *Enemy) {
	playerDamage := calculateDamage(p.Strength, p.weapon)
	if p.Equipment["оружие"] != nil {
		playerDamage += p.Equipment["оружие"].Attack
	}

	if p.Hit() != e.Block() {
		e.HP -= playerDamage
		fmt.Println(p.Name, "наносит", playerDamage, "урона", e.Name)
	} else {
		fmt.Println(e.Name, "блокирует удар", p.Name)
	}

	if e.HP > 0 && e.Hit() != p.Block() {
		enemyDamage := calculateDamage(e.Strength, e.weapon)

		if p.Equipment["броня"] != nil {
			armorReduction := p.Equipment["броня"].Defence / 5
			enemyDamage -= armorReduction
			if enemyDamage < 0 {
				enemyDamage = 0
			}
		}

		p.HP -= enemyDamage
		fmt.Println(e.Name, "наносит", enemyDamage, "урона", p.Name)
	} else if e.HP > 0 {
		fmt.Println(p.Name, "блокирует удар", e.Name)
	}
}

func calculateDamage(strength int, weapon *Weapon) int {
	baseDamage := rand.Intn(weapon.MaxDmg-weapon.MinDmg+1) + weapon.MinDmg
	return baseDamage
}

func Status(p *Player, e *Enemy) {
	fmt.Println("HP:", p.Name, p.HP, e.Name, e.HP)
}

func Winner(p *Player, e *Enemy) {
	if p.HP > 0 {
		fmt.Println("Победитель:", p.Name)
		fmt.Println("Проигравший:", e.Name)
	} else {
		fmt.Println("Победитель:", e.Name)
		fmt.Println("Проигравший:", p.Name)
	}
}

func fight(p *Player, e *Enemy) {
	rand.Seed(time.Now().UnixNano())
	round := 1
	fmt.Println("ИГРА 1х1 НАЧАЛАСЬ:")
	fmt.Println(p.Name, "vs", e.Name)

	p.RandomWeapon()
	e.RandomWeapon()

	p.Inventory = []Item{
		weapons[rand.Intn(len(weapons))],
		armors[rand.Intn(len(armors))],
		consumables[rand.Intn(len(consumables))],
	}
	p.Equipment = make(map[string]*Item)

	if rand.Intn(2) == 1 {
		items := [][]Item{weapons, armors, consumables}
		itemType := rand.Intn(3)
		itemList := items[itemType]
		e.Item = &itemList[rand.Intn(len(itemList))]
	}

	for p.HP > 0 && e.HP > 0 {
		fmt.Println("Раунд - ", round)

		if story := generateRoundStory(round, p, e); story != "" {
			fmt.Println(story)
		}

		if rand.Float32() < 0.3 {
			p.RandomWeapon()
			fmt.Println(p.Name, "с оружием: ", p.weapon.Name, "(", p.weapon.MinDmg, "-", p.weapon.MaxDmg, "урона)")
		}

		var command string
		fmt.Println("Выберите действие: атака, экип, снять")
		fmt.Scan(&command)

		if command == "экип" {
			p.Equip()
			continue
		} else if command == "снять" {
			p.TakeOff()
			continue
		}

		PlayerHitBlock(p)
		EnemyHitBlock(e)

		fmt.Println(p.Name, "атакует", BodyPart(p.Hit()), "и защищает", BodyPart(p.Block()))
		fmt.Println(e.Name, "атакует", BodyPart(e.Hit()), "и защищает", BodyPart(e.Block()))

		pHPBefore := p.HP
		eHPBefore := e.HP

		Round(p, e)

		pDamage := pHPBefore - p.HP
		eDamage := eHPBefore - e.HP

		if eDamage > 0 || pDamage > 0 {
			describeRoundResult(p, e, p.Hit(), e.Hit(), pDamage, eDamage)
		}

		fmt.Println("СТАТУС ПОСЛЕ РАУНДА:")
		fmt.Println(p.Name, ":", p.HP, "HP")
		fmt.Println(e.Name, ":", e.HP, "HP")

		round++
		time.Sleep(2 * time.Second)
	}

	fmt.Println("ИГРА ОКОНЧЕНА:")
	Winner(p, e)

	fmt.Println("СТАТИСТИКА БОЯ:")
	fmt.Println("Количество раундов:", round-1)
	if p.HP > 0 {
		fmt.Println(p.Name, "победил с", p.HP, "HP")
		fmt.Println("Оружие победителя:", p.weapon.Name)

		if e.Item != nil {
			p.Inventory = append(p.Inventory, *e.Item)
			fmt.Println("Трофей от", e.Name, ":", e.Item.Name)
		}
	} else {
		fmt.Println(e.Name, "победил с", e.HP, "HP")
		fmt.Println("Оружие победителя:", e.weapon.Name)
	}

	if len(p.Inventory) > 0 {
		fmt.Println("Ваш инвентарь:")
		for i, item := range p.Inventory {
			fmt.Println(i, "-", item.Name)
		}
	}

	if len(p.Equipment) > 0 {
		fmt.Println("Экипировка:")
		for equipType, item := range p.Equipment {
			fmt.Println(equipType, ":", item.Name)
		}
	}
}

func main() {
	player := &Player{
		Name:     "Ама",
		HP:       100,
		Strength: 0,
	}

	enemy := &Enemy{
		Name:     "Дамир",
		HP:       100,
		Strength: 0,
	}

	fight(player, enemy)
}

