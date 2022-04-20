package core

import (
	"github.com/mashnoor/nightwatch/pkg/settings"
	"sync"
)

func initCheckSystem() {
	var wg sync.WaitGroup
	settings.Log.Println("Started service")

	for _, service := range settings.SystemAppConfig.Clusters {
		currentService := service
		//checkLag(&currentService)

		go execute(&currentService, &wg)
		wg.Add(1)

	}

	wg.Wait()
}

func BootApp() {
	settings.LoadAppConfig()
	settings.SetupLogger()
	initCheckSystem()
}
