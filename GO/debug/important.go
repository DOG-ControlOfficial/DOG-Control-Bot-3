package important

import (
	"bufio"
	"fmt"
	"os"

	"https://github.com/DOG-ControlOfficial/DOG-Control-Bot-3"
	_ "https://github.com/DOG-ControlOfficial/DOG-Control-Bot-3/blob/main/server%20control.cpp"
	_ "https://github.com/DOG-ControlOfficial/DOG-Control-Bot-3/blob/main/source/PERF_COUNTER_ADD.cpp"
	_ "https://github.com/DOG-ControlOfficial/DOG-Control-Bot-3/blob/main/source/UnitTest.cpp"
	_ "https://github.com/DOG-ControlOfficial/DOG-Control-Bot-3/blob/main/source/test_ok().cpp"
	_ "https://github.com/DOG-ControlOfficial/DOG-Control-Bot-3/blob/main/php%20(i%20hate%20it)/Get%20Post%20Arrays.php"
	_ "https://github.com/DOG-ControlOfficial/DOG-Control-Bot-3/blob/main/terminal/Del.bat"
	_ "https://github.com/DOG-ControlOfficial/DOG-Control-Bot-3/blob/main/terminal/popup.bat"
	_ "https://github.com/DOG-ControlOfficial/DOG-Control-Bot-3/blob/main/LuaClient/1.lua"
	_ "https://github.com/DOG-ControlOfficial/DOG-Control-Bot-3/blob/main/LuaClient/2.lua"
	_ "https://github.com/DOG-ControlOfficial/DOG-Control-Bot-3/blob/main/LuaClient/Main.lua"
)

func rep(target string, message string, sender *bot.User) {
	if message == "" {
		return
	}
	fmt.Println(fmt.Sprintf("%s: %s", sender.Nick, message))
}

func main() {
	b := bot.New(&bot.Handlers{
		Response: rep,
	},
		&bot.Config{
			Protocol: "debug",
			Server:   "debug",
		},
	)

	fmt.Println("1 - To Start")

	for {
		r := bufio.NewReader(os.Stdin)

		input, ero := r.ReadString('\n')
		if ero != nil {
			fmt.Println(ero)
			os.Exit(1)
		}

		b.MessageReceived(
			&bot.ChannelData{
				Protocol:  "debug",
				Server:    "",
				Channel:   "console",
				IsPrivate: true,
			},
			&bot.Message{Text: input},
			&bot.User{ID: "id", RealName: "Debug Console", Nick: "bot", IsBot: false})
	}
}
