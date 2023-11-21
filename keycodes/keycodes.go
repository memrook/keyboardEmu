package keycodes

type KeyCode int

func (k KeyCode) String() (str string) {
	if s, ok := keyNames[k]; ok {
		return s
	}

	return ""
}

// See: https://en.wikipedia.org/wiki/C0_and_C1_control_codes
const (
	KeyNull                   = 0
	KeyStartOfHeading         = 1
	KeyStartOfText            = 2
	KeyExit                   = 3 // ctrl-c
	KeyEndOfTransimission     = 4
	KeyEnquiry                = 5
	KeyAcknowledge            = 6
	KeyBELL                   = 7
	KeyBackspace              = 8
	KeyHorizontalTabulation   = 9
	KeyLineFeed               = 10
	KeyVerticalTabulation     = 11
	KeyFormFeed               = 12
	KeyCarriageReturn         = 13
	KeyShiftOut               = 14
	KeyShiftIn                = 15
	KeyDataLinkEscape         = 16
	KeyDeviceControl1         = 17
	KeyDeviceControl2         = 18
	KeyDeviceControl3         = 19
	KeyDeviceControl4         = 20
	KeyNegativeAcknowledge    = 21
	KeySynchronousIdle        = 22
	KeyEndOfTransmissionBlock = 23
	KeyCancel                 = 24
	KeyEndOfMedium            = 25
	KeySubstitution           = 26
	KeyEscape                 = 27
	KeyFileSeparator          = 28
	KeyGroupSeparator         = 29
	KeyRecordSeparator        = 30
	KeyUnitSeparator          = 31
	KeyDelete                 = 127
)

const (
	Null      KeyCode = KeyNull
	Break     KeyCode = KeyExit
	Enter     KeyCode = KeyCarriageReturn
	Backspace KeyCode = KeyDelete
	Tab       KeyCode = KeyHorizontalTabulation
	Esc       KeyCode = KeyEscape
	Escape    KeyCode = KeyEscape

	CtrlAt KeyCode = KeyNull
	CtrlA  KeyCode = KeyStartOfHeading
	CtrlB  KeyCode = KeyStartOfText
	CtrlC  KeyCode = KeyExit
	CtrlD  KeyCode = KeyEndOfTransimission
	CtrlE  KeyCode = KeyEnquiry
	CtrlF  KeyCode = KeyAcknowledge
	CtrlG  KeyCode = KeyBELL
	CtrlH  KeyCode = KeyBackspace
	CtrlI  KeyCode = KeyHorizontalTabulation
	CtrlJ  KeyCode = KeyLineFeed
	CtrlK  KeyCode = KeyVerticalTabulation
	CtrlL  KeyCode = KeyFormFeed
	CtrlM  KeyCode = KeyCarriageReturn
	CtrlN  KeyCode = KeyShiftOut
	CtrlO  KeyCode = KeyShiftIn
	CtrlP  KeyCode = KeyDataLinkEscape
	CtrlQ  KeyCode = KeyDeviceControl1
	CtrlR  KeyCode = KeyDeviceControl2
	CtrlS  KeyCode = KeyDeviceControl3
	CtrlT  KeyCode = KeyDeviceControl4
	CtrlU  KeyCode = KeyNegativeAcknowledge
	CtrlV  KeyCode = KeySynchronousIdle
	CtrlW  KeyCode = KeyEndOfTransmissionBlock
	CtrlX  KeyCode = KeyCancel
	CtrlY  KeyCode = KeyEndOfMedium
	CtrlZ  KeyCode = KeySubstitution

	CtrlOpenBracket  KeyCode = KeyEscape
	CtrlBackslash    KeyCode = KeyFileSeparator
	CtrlCloseBracket KeyCode = KeyGroupSeparator
	CtrlCaret        KeyCode = KeyRecordSeparator
	CtrlUnderscore   KeyCode = KeyUnitSeparator
	CtrlQuestionMark KeyCode = KeyDelete
)

// Other keys.
const (
	RuneKey KeyCode = -(iota + 1)
	Up
	Down
	Right
	Left
	ShiftTab
	Home
	End
	PgUp
	PgDown
	Delete
	Space
	CtrlUp
	CtrlDown
	CtrlRight
	CtrlLeft
	ShiftUp
	ShiftDown
	ShiftRight
	ShiftLeft
	CtrlShiftUp
	CtrlShiftDown
	CtrlShiftLeft
	CtrlShiftRight
	F1
	F2
	F3
	F4
	F5
	F6
	F7
	F8
	F9
	F10
	F11
	F12
	F13
	F14
	F15
	F16
	F17
	F18
	F19
	F20
)

var keyNames = map[KeyCode]string{
	// Control keys.
	KeyNull:                   "ctrl+@", // also ctrl+backtick
	KeyStartOfHeading:         "ctrl+a",
	KeyStartOfText:            "ctrl+b",
	KeyExit:                   "ctrl+c",
	KeyEndOfTransimission:     "ctrl+d",
	KeyEnquiry:                "ctrl+e",
	KeyAcknowledge:            "ctrl+f",
	KeyBELL:                   "ctrl+g",
	KeyBackspace:              "ctrl+h",
	KeyHorizontalTabulation:   "tab", // also ctrl+i
	KeyLineFeed:               "ctrl+j",
	KeyVerticalTabulation:     "ctrl+k",
	KeyFormFeed:               "ctrl+l",
	KeyCarriageReturn:         "enter",
	KeyShiftOut:               "ctrl+n",
	KeyShiftIn:                "ctrl+o",
	KeyDataLinkEscape:         "ctrl+p",
	KeyDeviceControl1:         "ctrl+q",
	KeyDeviceControl2:         "ctrl+r",
	KeyDeviceControl3:         "ctrl+s",
	KeyDeviceControl4:         "ctrl+t",
	KeyNegativeAcknowledge:    "ctrl+u",
	KeySynchronousIdle:        "ctrl+v",
	KeyEndOfTransmissionBlock: "ctrl+w",
	KeyCancel:                 "ctrl+x",
	KeyEndOfMedium:            "ctrl+y",
	KeySubstitution:           "ctrl+z",
	KeyEscape:                 "esc",
	KeyFileSeparator:          "ctrl+\\",
	KeyGroupSeparator:         "ctrl+]",
	KeyRecordSeparator:        "ctrl+^",
	KeyUnitSeparator:          "ctrl+_",
	KeyDelete:                 "backspace",

	// Other keys.
	RuneKey:        "runes",
	Up:             "up",
	Down:           "down",
	Right:          "right",
	Space:          "space",
	Left:           "left",
	ShiftTab:       "shift+tab",
	Home:           "home",
	End:            "end",
	PgUp:           "pgup",
	PgDown:         "pgdown",
	Delete:         "delete",
	CtrlUp:         "ctrl+up",
	CtrlDown:       "ctrl+down",
	CtrlRight:      "ctrl+right",
	CtrlLeft:       "ctrl+left",
	ShiftUp:        "shift+up",
	ShiftDown:      "shift+down",
	ShiftRight:     "shift+right",
	ShiftLeft:      "shift+left",
	CtrlShiftUp:    "ctrl+shift+up",
	CtrlShiftDown:  "ctrl+shift+down",
	CtrlShiftLeft:  "ctrl+shift+left",
	CtrlShiftRight: "ctrl+shift+right",
	F1:             "f1",
	F2:             "f2",
	F3:             "f3",
	F4:             "f4",
	F5:             "f5",
	F6:             "f6",
	F7:             "f7",
	F8:             "f8",
	F9:             "f9",
	F10:            "f10",
	F11:            "f11",
	F12:            "f12",
	F13:            "f13",
	F14:            "f14",
	F15:            "f15",
	F16:            "f16",
	F17:            "f17",
	F18:            "f18",
	F19:            "f19",
	F20:            "f20",
}
