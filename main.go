package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/micmonay/keybd_event"
	"keyboardEmu/icon"
	"log"
	"runtime"
	"time"
)

func main() {
	ch := make(chan bool)
	isActive := true

	go keyboardEvent(ch)

	go func() {
		for {
			ch <- isActive
			time.Sleep(5 * time.Second)
			log.Println("send to chan <- ", isActive)
		}
	}()

	trayOnExit := func() {
		log.Println("called trayOnExit")
	}
	systray.Run(func() { trayOnReady(&isActive) }, trayOnExit)
}

func keyboardEvent(ch <-chan bool) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		var counter int
		try(err, &counter, 5)
	}

	// For linux, it is very important to wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	// Select keys to be pressed
	kb.SetKeys(keybd_event.VK_F5)

	// Set shift to be pressed
	kb.HasSHIFT(false)

	for v := range ch {
		if v {
			// Press the selected keys
			err = kb.Launching()
			if err != nil {
				var counter int
				try(err, &counter, 5)
			}
		}
	}

	// Here, the program will generate "ABAB" as if they were pressed on the keyboard.
}

func try(err error, counter *int, tries int) {
	if tries > 0 && tries == *counter {
		log.Fatal(err)
	}
	log.Printf("err %d/%d: %v", counter, tries, err)
}

func trayOnReady(mEnableStatus *bool) {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("keyboardEmu")
	systray.SetTooltip("Emulate keyboard events")

	mEnable := systray.AddMenuItemCheckbox("Enable", "Enable App", true)
	//mEnableStatus := true
	systray.AddSeparator()
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		for {
			select {
			case <-mQuitOrig.ClickedCh:
				fmt.Println("Requesting quit")
				systray.Quit()
				fmt.Println("Finished quitting")
			case <-mEnable.ClickedCh:
				if *mEnableStatus {
					*mEnableStatus = false
					mEnable.Uncheck()
				} else {
					*mEnableStatus = true
					mEnable.Check()
				}
			}
		}
	}()
}
