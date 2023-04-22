package main

import (
	"flag"
	"fmt"
	"github.com/timofef/tinyURL/internal/pkg/tinyURL/repository"
	"github.com/timofef/tinyURL/internal/pkg/tinyURL/usecase"
	"github.com/timofef/tinyURL/internal/tinyURL/utils"
	"sync"
)

func main() {
	// Parse storage option flag
	useInMemoryStorage := *flag.Bool("in-memory", false, "use in-memory storage")
	flag.Parse()

	// Init repository
	var repo repository.IRepository
	if useInMemoryStorage {
		repo = &repository.TinyUrlInMemRepository{
			Mux: sync.RWMutex{},
			DB:  make(map[string]string),
		}
		// log
		fmt.Println("Using in-memory storage")
	} else {
		/*repo = &repository.TinyUrlSqlRepository{
			Mux: sync.RWMutex{},
			DB: make(map[string]string),
		}*/
	}

	// Init usecase
	usecase := usecase.TinyUrlUsecase{
		BaseUrl:         "http://base.com/",
		Repository:      repo,
		GenerateTinyUrl: utils.GenerateString,
	}

	// Init server

	// Serve

}
