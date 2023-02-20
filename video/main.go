package main

import (
	"douyin_backend_microService/video/kitex_gen/videodemo/videoservice"
	"log"
)

func main() {
	svr := videoservice.NewServer(new(VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
