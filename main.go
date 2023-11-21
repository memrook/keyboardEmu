package main

import (
	"encoding/json"
	"github.com/getlantern/systray"
	"github.com/micmonay/keybd_event"
	"keyboardEmu/icon"
	"log"
	"os"
	"time"
)

type configuration struct {
	EnableOnStartup bool          `json:"enableOnStartup"`
	Duration        time.Duration `json:"duration"`
	//Keys            []string      `json:"keys"`
}

var conf configuration

func init() {
	configFile, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal("can't open the file: ", err)
	}

	if err = json.Unmarshal(configFile, &conf); err != nil {
		log.Fatal("can't unmarshal file: ", err)
	}
}

func main() {
	ch := make(chan bool)
	isActive := conf.EnableOnStartup
	delay := conf.Duration

	//launching routine to pressing keys
	go keyboardEvent(ch)

	//routine for sending status at runtime
	go func() {
		for {
			ch <- isActive
			time.Sleep(delay * time.Second)
			log.Println("send to chan <- ", isActive)
		}
	}()

	//setup tray icon
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
	//if runtime.GOOS == "linux" {
	time.Sleep(2 * time.Second)
	//}

	// Select keys to be pressed
	kb.SetKeys(keybd_event.VK_F5)

	// Set a modifiers to be pressed
	kb.HasSHIFT(false)
	kb.HasALT(false)
	kb.HasCTRL(false)

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

func trayOnReady(isActive *bool) {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("keyboardEmu")
	systray.SetTooltip("Emulating keyboard events")

	mEnable := systray.AddMenuItemCheckbox("Enable", "Enable Refresh", *isActive)
	systray.AddSeparator()
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the app")

	go func() {
		for {
			select {
			case <-mQuitOrig.ClickedCh:
				log.Println("Requesting quit")
				systray.Quit()
				log.Println("Finished quitting")
			case <-mEnable.ClickedCh:
				if *isActive {
					*isActive = false
					mEnable.Uncheck()
				} else {
					*isActive = true
					mEnable.Check()
				}
			}
		}
	}()
}
