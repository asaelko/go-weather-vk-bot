package main

import (
	"log"
	"sync"

	"flag"
	"github.com/vharitonsky/iniflags"

	"vk/longpoll"
)

var (
	vkApiUrl      *string = flag.String("vk_api_url", "https://api.vk.com/method/", "VK api url")
	vkGroupId     *int    = flag.Int("vk_group_id", 0, "VK group id")
	vkGroupToken  *string = flag.String("vk_group_token", "", "VK group token")
	forecastToken *string = flag.String("forecast_token", "", "Forecast.io token")
)

func main() {
	log.Println("App started..")
	iniflags.Parse()
	log.Println("Config loaded..")

	if err := longpoll.GetCredentials(vkApiUrl, vkGroupToken); err != nil {
		log.Fatalln(err)
	}
	log.Println("Initial Longpoll data received")

	log.Println("Start longpoll receiver")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Fatalln("Recovered from error: ", r)
			}

			wg.Done()
		}()

		longpoll.Pull()
	}()

	wg.Wait()
}
