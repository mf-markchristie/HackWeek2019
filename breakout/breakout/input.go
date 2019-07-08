package breakout

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func (i *Input) Update() {
	if !i.keyboardConfig.IsInitialized() {
		return
	}

	if i.virtualKeyboardKeyStates == nil {
		i.virtualKeyboardKeyStates = map[virtualKeyboardKey]int{}
	}
	for _, b := range virtualKeyboardKey {
		if !i.keyboardConfig.IsKeyPressed(b) {
			i.virtualKeyboardKeyStates[b] = 0
			continue
		}
		i.virtualKeyboardKeyStates[b]++
	}
}
