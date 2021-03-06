package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var AllKeysList = []ebiten.Key{
	ebiten.Key0, ebiten.Key1, ebiten.Key2, ebiten.Key3, ebiten.Key4, ebiten.Key5, ebiten.Key6, ebiten.Key7, ebiten.Key8, ebiten.Key9, ebiten.KeyA, ebiten.KeyB, ebiten.KeyC, ebiten.KeyD, ebiten.KeyE, ebiten.KeyF, ebiten.KeyG, ebiten.KeyH, ebiten.KeyI, ebiten.KeyJ, ebiten.KeyK, ebiten.KeyL, ebiten.KeyM, ebiten.KeyN, ebiten.KeyO, ebiten.KeyP, ebiten.KeyQ, ebiten.KeyR, ebiten.KeyS, ebiten.KeyT, ebiten.KeyU, ebiten.KeyV, ebiten.KeyW, ebiten.KeyX, ebiten.KeyY, ebiten.KeyZ, ebiten.KeyApostrophe, ebiten.KeyBackslash, ebiten.KeyBackspace, ebiten.KeyCapsLock, ebiten.KeyComma, ebiten.KeyDelete, ebiten.KeyDown, ebiten.KeyEnd, ebiten.KeyEnter, ebiten.KeyEqual, ebiten.KeyEscape, ebiten.KeyF1, ebiten.KeyF2, ebiten.KeyF3, ebiten.KeyF4, ebiten.KeyF5, ebiten.KeyF6, ebiten.KeyF7, ebiten.KeyF8, ebiten.KeyF9, ebiten.KeyF10, ebiten.KeyF11, ebiten.KeyF12, ebiten.KeyGraveAccent, ebiten.KeyHome, ebiten.KeyInsert, ebiten.KeyKP0, ebiten.KeyKP1, ebiten.KeyKP2, ebiten.KeyKP3, ebiten.KeyKP4, ebiten.KeyKP5, ebiten.KeyKP6, ebiten.KeyKP7, ebiten.KeyKP8, ebiten.KeyKP9, ebiten.KeyKPAdd, ebiten.KeyKPDecimal, ebiten.KeyKPDivide, ebiten.KeyKPEnter, ebiten.KeyKPEqual, ebiten.KeyKPMultiply, ebiten.KeyKPSubtract, ebiten.KeyLeft, ebiten.KeyLeftBracket, ebiten.KeyMenu, ebiten.KeyMinus, ebiten.KeyNumLock, ebiten.KeyPageDown, ebiten.KeyPageUp, ebiten.KeyPause, ebiten.KeyPeriod, ebiten.KeyPrintScreen, ebiten.KeyRight, ebiten.KeyRightBracket, ebiten.KeyScrollLock, ebiten.KeySemicolon, ebiten.KeySlash, ebiten.KeySpace, ebiten.KeyTab, ebiten.KeyUp, ebiten.KeyAlt, ebiten.KeyControl, ebiten.KeyShift, ebiten.KeyMax,
}

var KeyMapKeyToString = map[ebiten.Key]string{
	ebiten.Key0:            "0",
	ebiten.Key1:            "1",
	ebiten.Key2:            "2",
	ebiten.Key3:            "3",
	ebiten.Key4:            "4",
	ebiten.Key5:            "5",
	ebiten.Key6:            "6",
	ebiten.Key7:            "7",
	ebiten.Key8:            "8",
	ebiten.Key9:            "9",
	ebiten.KeyA:            "a",
	ebiten.KeyB:            "b",
	ebiten.KeyC:            "c",
	ebiten.KeyD:            "d",
	ebiten.KeyE:            "e",
	ebiten.KeyF:            "f",
	ebiten.KeyG:            "g",
	ebiten.KeyH:            "h",
	ebiten.KeyI:            "i",
	ebiten.KeyJ:            "j",
	ebiten.KeyK:            "k",
	ebiten.KeyL:            "l",
	ebiten.KeyM:            "m",
	ebiten.KeyN:            "n",
	ebiten.KeyO:            "o",
	ebiten.KeyP:            "p",
	ebiten.KeyQ:            "q",
	ebiten.KeyR:            "r",
	ebiten.KeyS:            "s",
	ebiten.KeyT:            "t",
	ebiten.KeyU:            "u",
	ebiten.KeyV:            "v",
	ebiten.KeyW:            "w",
	ebiten.KeyX:            "x",
	ebiten.KeyY:            "y",
	ebiten.KeyZ:            "z",
	ebiten.KeyAlt:          "alt",
	ebiten.KeyApostrophe:   "apostrophe",
	ebiten.KeyBackslash:    "backslash",
	ebiten.KeyBackspace:    "backspace",
	ebiten.KeyCapsLock:     "capslock",
	ebiten.KeyComma:        "comma",
	ebiten.KeyControl:      "control",
	ebiten.KeyDelete:       "delete",
	ebiten.KeyDown:         "down",
	ebiten.KeyEnd:          "end",
	ebiten.KeyEnter:        "enter",
	ebiten.KeyEqual:        "equal",
	ebiten.KeyEscape:       "escape",
	ebiten.KeyF1:           "f1",
	ebiten.KeyF2:           "f2",
	ebiten.KeyF3:           "f3",
	ebiten.KeyF4:           "f4",
	ebiten.KeyF5:           "f5",
	ebiten.KeyF6:           "f6",
	ebiten.KeyF7:           "f7",
	ebiten.KeyF8:           "f8",
	ebiten.KeyF9:           "f9",
	ebiten.KeyF10:          "f10",
	ebiten.KeyF11:          "f11",
	ebiten.KeyF12:          "f12",
	ebiten.KeyGraveAccent:  "graveaccent",
	ebiten.KeyHome:         "home",
	ebiten.KeyInsert:       "insert",
	ebiten.KeyKP0:          "kp0",
	ebiten.KeyKP1:          "kp1",
	ebiten.KeyKP2:          "kp2",
	ebiten.KeyKP3:          "kp3",
	ebiten.KeyKP4:          "kp4",
	ebiten.KeyKP5:          "kp5",
	ebiten.KeyKP6:          "kp6",
	ebiten.KeyKP7:          "kp7",
	ebiten.KeyKP8:          "kp8",
	ebiten.KeyKP9:          "kp9",
	ebiten.KeyKPAdd:        "kpadd",
	ebiten.KeyKPDecimal:    "kpdecimal",
	ebiten.KeyKPDivide:     "kpdivide",
	ebiten.KeyKPEnter:      "kpenter",
	ebiten.KeyKPEqual:      "kpequal",
	ebiten.KeyKPMultiply:   "kpmultiply",
	ebiten.KeyKPSubtract:   "kpsubtract",
	ebiten.KeyLeft:         "left",
	ebiten.KeyLeftBracket:  "leftbracket",
	ebiten.KeyMenu:         "menu",
	ebiten.KeyMinus:        "minus",
	ebiten.KeyNumLock:      "numlock",
	ebiten.KeyPageDown:     "pagedown",
	ebiten.KeyPageUp:       "pageup",
	ebiten.KeyPause:        "pause",
	ebiten.KeyPeriod:       "period",
	ebiten.KeyPrintScreen:  "printscreen",
	ebiten.KeyRight:        "right",
	ebiten.KeyRightBracket: "rightbracket",
	ebiten.KeyScrollLock:   "scrolllock",
	ebiten.KeySemicolon:    "semicolon",
	ebiten.KeyShift:        "shift",
	ebiten.KeySlash:        "slash",
	ebiten.KeySpace:        "space",
	ebiten.KeyTab:          "tab",
	ebiten.KeyUp:           "up",
}

var KeyMapStringToKey = map[string]ebiten.Key{
	"0":            ebiten.Key0,
	"1":            ebiten.Key1,
	"2":            ebiten.Key2,
	"3":            ebiten.Key3,
	"4":            ebiten.Key4,
	"5":            ebiten.Key5,
	"6":            ebiten.Key6,
	"7":            ebiten.Key7,
	"8":            ebiten.Key8,
	"9":            ebiten.Key9,
	"a":            ebiten.KeyA,
	"b":            ebiten.KeyB,
	"c":            ebiten.KeyC,
	"d":            ebiten.KeyD,
	"e":            ebiten.KeyE,
	"f":            ebiten.KeyF,
	"g":            ebiten.KeyG,
	"h":            ebiten.KeyH,
	"i":            ebiten.KeyI,
	"j":            ebiten.KeyJ,
	"k":            ebiten.KeyK,
	"l":            ebiten.KeyL,
	"m":            ebiten.KeyM,
	"n":            ebiten.KeyN,
	"o":            ebiten.KeyO,
	"p":            ebiten.KeyP,
	"q":            ebiten.KeyQ,
	"r":            ebiten.KeyR,
	"s":            ebiten.KeyS,
	"t":            ebiten.KeyT,
	"u":            ebiten.KeyU,
	"v":            ebiten.KeyV,
	"w":            ebiten.KeyW,
	"x":            ebiten.KeyX,
	"y":            ebiten.KeyY,
	"z":            ebiten.KeyZ,
	"alt":          ebiten.KeyAlt,
	"apostrophe":   ebiten.KeyApostrophe,
	"backslash":    ebiten.KeyBackslash,
	"backspace":    ebiten.KeyBackspace,
	"capslock":     ebiten.KeyCapsLock,
	"comma":        ebiten.KeyComma,
	"control":      ebiten.KeyControl,
	"delete":       ebiten.KeyDelete,
	"down":         ebiten.KeyDown,
	"end":          ebiten.KeyEnd,
	"enter":        ebiten.KeyEnter,
	"equal":        ebiten.KeyEqual,
	"escape":       ebiten.KeyEscape,
	"f1":           ebiten.KeyF1,
	"f2":           ebiten.KeyF2,
	"f3":           ebiten.KeyF3,
	"f4":           ebiten.KeyF4,
	"f5":           ebiten.KeyF5,
	"f6":           ebiten.KeyF6,
	"f7":           ebiten.KeyF7,
	"f8":           ebiten.KeyF8,
	"f9":           ebiten.KeyF9,
	"f10":          ebiten.KeyF10,
	"f11":          ebiten.KeyF11,
	"f12":          ebiten.KeyF12,
	"graveaccent":  ebiten.KeyGraveAccent,
	"home":         ebiten.KeyHome,
	"insert":       ebiten.KeyInsert,
	"kp0":          ebiten.KeyKP0,
	"kp1":          ebiten.KeyKP1,
	"kp2":          ebiten.KeyKP2,
	"kp3":          ebiten.KeyKP3,
	"kp4":          ebiten.KeyKP4,
	"kp5":          ebiten.KeyKP5,
	"kp6":          ebiten.KeyKP6,
	"kp7":          ebiten.KeyKP7,
	"kp8":          ebiten.KeyKP8,
	"kp9":          ebiten.KeyKP9,
	"kpadd":        ebiten.KeyKPAdd,
	"kpdecimal":    ebiten.KeyKPDecimal,
	"kpdivide":     ebiten.KeyKPDivide,
	"kpenter":      ebiten.KeyKPEnter,
	"kpequal":      ebiten.KeyKPEqual,
	"kpmultiply":   ebiten.KeyKPMultiply,
	"kpsubtract":   ebiten.KeyKPSubtract,
	"left":         ebiten.KeyLeft,
	"leftbracket":  ebiten.KeyLeftBracket,
	"menu":         ebiten.KeyMenu,
	"minus":        ebiten.KeyMinus,
	"numlock":      ebiten.KeyNumLock,
	"pagedown":     ebiten.KeyPageDown,
	"pageup":       ebiten.KeyPageUp,
	"pause":        ebiten.KeyPause,
	"period":       ebiten.KeyPeriod,
	"printscreen":  ebiten.KeyPrintScreen,
	"right":        ebiten.KeyRight,
	"rightbracket": ebiten.KeyRightBracket,
	"scrolllock":   ebiten.KeyScrollLock,
	"semicolon":    ebiten.KeySemicolon,
	"shift":        ebiten.KeyShift,
	"slash":        ebiten.KeySlash,
	"space":        ebiten.KeySpace,
	"tab":          ebiten.KeyTab,
	"up":           ebiten.KeyUp,
}
