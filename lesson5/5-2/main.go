package main

import (
	// "log"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// log
	log.Print("Hay, I'm a log!")

	// fatal
	// log.Fatal("Hay, I'm an error log!")
	// fmt.Print("Can you see me?")

	// panic
	// log.Panic("Hay, I'm panic!")
	// fmt.Print("Can you see me?")

	// prefix
	// log.SetPrefix("main(): ")
	// log.Print("Hey, I'm a log!")
	// log.Fatal("Hay, I'm an error log!")

	// ファイルへの記録
	// file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer file.Close()

	// log.SetOutput(file)
	// log.Print("Hey, I'm a log!")

	// ログ記録フレームワーク
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Print("Hay! I'm a log message!")

	log.Debug().Int("EmployeeID", 1001).Msg("Getting employee information")
	log.Debug().Str("Name", "John").Send()
}
