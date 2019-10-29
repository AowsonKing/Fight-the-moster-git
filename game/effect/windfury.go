package effect

import "mashiroc.fun/dragongame/game"

/**
   风怒特效
 */

type Windfury struct {
	game.BaseEffect
	isDo bool
	fun func(game.Card, game.Character, game.Character)
}

func NewWindfury(fun func(game.Card, game.Character, game.Character)) BattleCry {
	return BattleCry{
		BaseEffect: game.BaseEffect{},
		isDo:       false,
		fun:        fun,
	}
}

func (w Windfury) Do(card game.Card, self, other game.Character) {
	if !w.isDo {
		w.fun(card, self, other)
	}
}
