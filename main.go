package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func main() {
	var moveUp bool
	exec.Command("adb", "shell", "input", "tap", "540", "200").Run() //Dismissing any overlays, Tap in the middle upper screen
	for {
		fmt.Println("Checking")
		exec.Command("adb", "shell", "uiautomator", "dump").Run()
		out, _ := exec.Command("adb", "shell", "cat", "/sdcard/window_dump.xml").Output()
		xmlStr := string(out)
		if !strings.Contains(xmlStr, `package="com.badoo.mobile"`) {
			exec.Command("adb", "shell", "am", "start", "-n", "com.badoo.mobile/.android.BadooActivity").Run()
			fmt.Println("Changing Apps back to Badoo 📲")
		}
		if strings.Contains(xmlStr, `text="Watch Now"`) {
			exec.Command("adb", "shell", "am", "start", "-n", "com.badoo.mobile/.android.BadooActivity").Run()
			fmt.Println("Changing Apps back to Badoo 📲")
		}
		if strings.Contains(xmlStr, "We’ve updated your filters to show you more people from further away") {
			exec.Command("adb", "shell", "input", "tap", "540", "1661").Run()
			fmt.Println("Geting rid of the filter pop-up")
		}

		if strings.Contains(xmlStr, "Average rating 4.7 stars in 28 thousand reviews") {
			exec.Command("adb", "shell", "input", "swipe", "540", "960", "540", "760", "200").Run()
			fmt.Println(" Glam Ad, so pressing Back 🔙")
			continue
		}

		if strings.Contains(xmlStr, "Google Play") {
			exec.Command("adb", "shell", "input", "tap", "251", "181").Run()
			fmt.Println("Clicking Google Play 🛝...")
			exec.Command("adb", "shell", "am", "start", "-n", "com.badoo.mobile/.android.BadooActivity").Run()
			time.Sleep(time.Second)
			fmt.Println("Changing Apps back to Badoo 📲")
		}

		if strings.Contains(xmlStr, "[847,128][1080,432]") {
			time.Sleep(time.Second)
			exec.Command("adb", "shell", "input", "tap", "964", "280").Run()
			fmt.Println(" Continue 🪙")
			continue
		}

		if strings.Contains(xmlStr, "Get  Crush each day") {
			fmt.Println("👋 Program End 🔚")
			return
		}
		if strings.Contains(xmlStr, "We don’t have a video to display right now, please come back soon!") {
			exec.Command("adb", "shell", "input", "tap", "877", "1311").Run()
			fmt.Println("👋 Program End 🔚")
			return
		}
		if strings.Contains(xmlStr, `text="Watch A Video"`) {
			exec.Command("adb", "shell", "input", "tap", "540", "1657").Run()
			fmt.Println("Dismissing pop-up 🌤️")
		}

		if strings.Contains(xmlStr, "Don&#39;t miss out! Install") || strings.Contains(xmlStr, "Shapedly Shop Now") || strings.Contains(xmlStr, `text="AI Hub: Photo, Video Generator"`) || strings.Contains(xmlStr, "SHOP NOW 👉") {
			exec.Command("adb", "shell", "am", "start", "-n", "com.badoo.mobile/.android.BadooActivity").Run()
			fmt.Println("Yalla Ad found")
			time.Sleep(time.Millisecond * 500)
		}

		if strings.Contains(xmlStr, `text="Free 10 credits"`) {
			exec.Command("adb", "shell", "input", "tap", "540", "1757").Run()
			fmt.Println("Free 10 credit's, Maybe Later 🎬")
		}

		if strings.Contains(xmlStr, "Keeta") {
			//exec.Command("adb", "shell", "input", "swipe", "900", "1000", "100", "1000", "300").Run()
			fmt.Println("Found Keeta ad")
		}

		if strings.Contains(xmlStr, `com.badoo.mobile:id/ad_container`) {
			fmt.Println("Ad bar detected at bottom")
			moveUp = true
		} else {
			moveUp = false
		}
		if strings.Contains(xmlStr, "Encounters") && moveUp {
			fmt.Println("Profile page with banner on bottom")
			time.Sleep(time.Millisecond * 500)
			exec.Command("adb", "shell", "input", "tap", "540", "1661").Run()
			fmt.Println("Clicking the Crush button")
		} else if strings.Contains(xmlStr, "Encounters") {
			fmt.Println("Profile page found")
			exec.Command("adb", "shell", "input", "tap", "540", "1863").Run()
			fmt.Println("Clicking the Crush button")
			time.Sleep(time.Millisecond * 500)
		} else if strings.Contains(xmlStr, "Show them you really like them") {
			fmt.Println("Sending for Free")
			exec.Command("adb", "shell", "input", "tap", "540", "1853").Run()
			fmt.Println("Clicking the send for free button")
			time.Sleep(time.Millisecond * 500)
		} else if strings.Contains(xmlStr, "_ad") || strings.Contains(xmlStr, "UnityAdsCache") || strings.Contains(xmlStr, "Unity Ads") || strings.HasPrefix(xmlStr, "mbridge_") || strings.HasPrefix(xmlStr, "left to be rewarded") || strings.Contains(xmlStr, "Cocos3dGameContainer") || strings.Contains(xmlStr, `resource-id="ad_video"`) || strings.Contains(xmlStr, "GameCanvas") || (strings.Contains(xmlStr, "Install") && strings.Contains(xmlStr, `resource-id="ad_video"`) || (strings.Contains(xmlStr, "Crush") && strings.Contains(xmlStr, "Awesome"))) {
			fmt.Println("📽️ Ad found🍿")
			if strings.Contains(xmlStr, "Close button") {
				exec.Command("adb", "shell", "input", "tap", "981", "195").Run()
				fmt.Println("✅ Closing the Ad")
			} else if strings.Contains(xmlStr, "[972,108][1069,206]") {
				exec.Command("adb", "shell", "input", "tap", "1020", "157").Run()
				fmt.Println("Clicking Next ⏭️")
			} else if strings.Contains(xmlStr, "[951,125][1053,226]") {
				exec.Command("adb", "shell", "input", "tap", "1002", "175").Run()
				fmt.Println("Clicking Next ⏭️WhiteOut Survival Game ☃️")
			} else if strings.Contains(xmlStr, "[962,149][1046,233]") {
				exec.Command("adb", "shell", "input", "tap", "1004", "191").Run()
				fmt.Println("Clicking Next ⏭️WhiteOut Survival Game ☃️")
			} else if strings.Contains(xmlStr, "[884,131][1046,293]") {
				exec.Command("adb", "shell", "input", "tap", "1004", "191").Run()
				fmt.Println("Clicking close❎ Royal Kingdom Ad 🫅")
			} else if strings.Contains(xmlStr, "[921,135][1026,243]") {
				exec.Command("adb", "shell", "input", "tap", "973", "189").Run()
				fmt.Println("Clicking close❎ Toki🫅")
			} else if strings.Contains(xmlStr, "[34,132][1046,382]") {
				exec.Command("adb", "shell", "input", "tap", "540", "257").Run()
				fmt.Println("Clicking close❎ Toki🫅")
			} else if strings.Contains(xmlStr, "[948,162][1015,229]") {
				exec.Command("adb", "shell", "input", "tap", "981", "195").Run()
			} else if strings.Contains(xmlStr, "Continue") {
				exec.Command("adb", "shell", "input", "tap", "856", "280").Run()
				fmt.Println("Clicking continue  ⚽")
			} else if !strings.Contains(xmlStr, `clickable="true"`) {
				exec.Command("powershell", "-Command", "adb shell input swipe 540 1200 540 1000 500; adb shell input swipe 540 1000 540 1100 500").Run()
				fmt.Println("Tapping middle screen  👇")
			} else if strings.Contains(xmlStr, "image_0-fullscreen-overlay") {
				exec.Command("adb", "shell", "input", "tap", "540", "1200").Run()
				fmt.Println("✅ Tapping fullscreen overlay")
			}

		} else if strings.Contains(xmlStr, `text="Sponsored"`) && strings.Contains(xmlStr, `com.badoo.mobile:id/encounters_stackView`) {
			exec.Command("adb", "shell", "input", "swipe", "900", "500", "100", "50").Run()
			fmt.Println("Swiping left on Ad 🛵...")

		} else if strings.Contains(xmlStr, "It’s a match!") {
			exec.Command("adb", "shell", "input", "tap", "95", "183").Run()
			fmt.Println("🥳It's a match! And closing..🍾")
		} else if strings.Contains(xmlStr, "Average rating 4.7 stars in 28 thousand reviews") {
			exec.Command("adb", "shell", "input", "keyevent", "KEYCODE_BACK").Run()
			fmt.Println(" Glam Screw Ad, so pressing Back 🔙")

		} else {
			exec.Command("powershell", "-Command", "adb shell input swipe 540 1200 540 1000 500; adb shell input swipe 540 1000 540 1100 500").Run()
			fmt.Println("🤔 Didn't match any checks 🤷")

		}
		time.Sleep(time.Second)
	}
}
