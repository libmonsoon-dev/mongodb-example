package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/cheggaaa/pb/v3"

	"mongodb-example/src/usecases/dto"
)

type Data struct {
	Objects []dto.User `json:"objects"`
}

func main() {
	logger := log.New(os.Stdout, "[create-users] ", log.Ltime)

	path := flag.String("path", "./seeds/users_go.json", "Path to json with users")
	maxThreads := flag.Int("threads", 100, "Max insert threads")

	flag.Parse()

	sem := make(chan struct{}, *maxThreads)

	service, err := InitializeUserService()

	if err != nil {
		panic(fmt.Errorf("cannot initialize UserService: %w", err))
	}

	raw, err := ioutil.ReadFile(*path)
	if err != nil {
		panic(fmt.Errorf("canot read %v: %w", *path, err))
	}

	var data Data

	if err := json.Unmarshal(raw, &data); err != nil {
		panic(err)
	}

	bar := pb.StartNew(len(data.Objects))

	for _, user := range data.Objects {
		sem <- struct{}{}
		go func(user dto.User) {
			_, err := service.Store(user)
			if err != nil {
				logger.Printf("Error while saving %#v: %v\n", user, err)
			}
			<-sem
			bar.Increment()
		}(user)

	}
	for i := 0; i < *maxThreads; i++ {
		sem <- struct{}{}
	}
	bar.Finish()
}
