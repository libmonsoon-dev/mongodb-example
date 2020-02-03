package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"mongodb-example/src/infrastructure"
	"os"

	"github.com/cheggaaa/pb/v3"

	"mongodb-example/src/usecases/dto"
)

type Data struct {
	Objects []dto.User `json:"objects"`
}

func main() {
	path := flag.String("path", "./seeds/users_go.json", "Path to json with users")
	maxThreads := flag.Int("threads", 100, "Max insert threads")

	flag.Parse()

	sem := make(chan struct{}, *maxThreads)

	service, err := Init()
	logger := infrastructure.NewStdLogger()

	if err != nil {
		logger.Errorf("cannot initialize UserService: %w", err)
		os.Exit(1)
	}

	raw, err := ioutil.ReadFile(*path)
	if err != nil {
		logger.Errorf("canot read %v: %w", *path, err)
		os.Exit(1)

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
				logger.Errorln("Error while saving %#v: %v\n", user, err)
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
