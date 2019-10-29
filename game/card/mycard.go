// Time : 2019/10/19 20:28
// Author : MashiroC

// card
package card

import (
	"fmt"
	"mashiroc.fun/dragongame/game"
	"mashiroc.fun/dragongame/game/effect"
)

// mycard.go something

type WildBoar struct {
	game.BaseCard
}

func NewWildBoar() game.Card {
	return &WildBoar{
		*game.NewCard("石牙野猪", 1, 1, 1, effect.Charge{}),
	}
}

type ElvenArcher struct {
	game.BaseCard
}

func NewElvenArcher() game.Card {
	return &ElvenArcher{
		*game.NewCard("精灵弓箭手",
			1,
			1,
			1,
			effect.NewBattleCry(
				func(card game.Card, player game.Character, dragon game.Character) {
					fmt.Println("请选择一个对方目标：")
					var flag int
					fmt.Scanf("%d", &flag)
					var target game.Card
					if flag == 0 {
						target = dragon
					} else {
						target = dragon.Servant()[flag-1]
					}
					target.ReduceHp(1)
				}),
		),
	}
}
type WitchDoctor struct{
	game.BaseCard
}
func NewWitchDoctor() game.Card{
	return &WitchDoctor{
		*game.NewCard("巫医",
			3,
			1,
			3,
			effect.NewBattleCry(
				func(card game.Card, player game.Character, dragon game.Character) {
					fmt.Println("请选择一个对方目标：")
					var flag int
					fmt.Scanf("%d", &flag)
					var target game.Card
					if flag == 0 {
						target = player
					} else {
						target = player.Servant()[flag-1]
					}
					target.ReduceHp(-1)

				}),
		),
	}
}
type WolfRider struct{
	game.BaseCard
}
func NewWolfRider()game.Card{
	return &WolfRider{
		*game.NewCard("狼骑士", 3, 1, 3, effect.Charge{}),
	}
}
type Treatment struct{
	game.BaseCard
}
func NewTreatment() game.Card{
	return &Treatment{
		*game.NewCard("治疗术",
			1,
			1,
			1,
			effect.NewBattleCry(
				func(card game.Card, player game.Character, dragon game.Character) {
					fmt.Println("请选择一个己方目标：")
					var flag int
					fmt.Scanf("%d", &flag)
					var target game.Card
					if flag == 0 {
						target = player
					} else {
						target = player.Servant()[flag-1]
					}
					target.ReduceHp(-2)
					defer card.ReduceHp(1)
				}),
		),
	}
}
type Windfury struct{
	game.BaseCard
}
func NewWindfury() game.Card {
	return &Windfury{
		*game.NewCard("风怒",
			1,
			2,
			0,
			effect.NewWindfury(
				func(card game.Card, player game.Character, dragon game.Character) {
					fmt.Println("请选择一个己方目标：")
					var flag int
					fmt.Scanf("%d", &flag)
					var target game.Card
					if flag == 0 {
						target = player
					} else {
						target = player.Servant()[flag-1]
					}
					target.SetAttendNum(2)
					defer card.ReduceHp(1)
				}),
		),
	}
}





type FireBall struct {
	game.BaseCard
}
func NewFireBall()game.Card{
	return &FireBall{
		*game.NewCard("火球术",
			1,
			1,
			1,
			effect.NewBattleCry(
				func(card game.Card, player game.Character, dragon game.Character) {
					fmt.Println("请选择一个对方目标：")
					var flag int
					fmt.Scanf("%d", &flag)
					var target game.Card
					if flag == 0 {
						target = dragon
					} else {
						target = dragon.Servant()[flag-1]
					}
					target.ReduceHp(99)
					defer card.ReduceHp(1)

				}),
		),
	}
}