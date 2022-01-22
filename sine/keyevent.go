package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type KeyEventSys struct {
	keyLedger        map[ebiten.Key]int64
	keyEventsPress   map[ebiten.Key][]KeyEventCB
	keyEventsRelease map[ebiten.Key][]KeyEventCB
}

type KeyEventCB func(*Game)

func NewKeyEventSys() *KeyEventSys {
	var kvs = KeyEventSys{}
	kvs.keyLedger = make(map[ebiten.Key]int64, len(AllKeysList))
	for i := 0; i < len(AllKeysList); i++ {
		kvs.keyLedger[AllKeysList[i]] = -1
	}

	kvs.keyEventsPress = make(map[ebiten.Key][]KeyEventCB)
	kvs.keyEventsRelease = make(map[ebiten.Key][]KeyEventCB)
	return &kvs
}

func (kvs *KeyEventSys) UpdateInput(g *Game) {
	for i := 0; i < len(AllKeysList); i++ {
		curKey := AllKeysList[i]
		lastPress := kvs.keyLedger[curKey]
		pressHandlers, phExists := kvs.keyEventsPress[curKey]
		releaseHandlers, rhExists := kvs.keyEventsRelease[curKey]
		if phExists || rhExists {
			if lastPress > -1 {
				if !ebiten.IsKeyPressed(curKey) {
					// Event Release
					for _, cb := range releaseHandlers {
						cb(g)
					}
					kvs.keyLedger[curKey] = -1
				}
			} else {
				if ebiten.IsKeyPressed(curKey) {
					// Event Press
					for _, cb := range pressHandlers {
						cb(g)
					}
					kvs.keyLedger[curKey] = g.frame
				}
			}
		}
	}
}

func (kvs *KeyEventSys) AddPressHandler(key ebiten.Key, cb KeyEventCB) {
	if _, ok := kvs.keyEventsPress[key]; !ok {
		kvs.keyEventsPress[key] = append(kvs.keyEventsPress[key], cb)
	}
}

func (kvs *KeyEventSys) AddReleaseHandler(key ebiten.Key, cb KeyEventCB) {
	if _, ok := kvs.keyEventsRelease[key]; !ok {
		kvs.keyEventsRelease[key] = append(kvs.keyEventsRelease[key], cb)
	}
}
