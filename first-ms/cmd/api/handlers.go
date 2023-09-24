package main

import (
	"fmt"
	"net/http"
	"os"
)

func (app *application) video(w http.ResponseWriter, r *http.Request) {

	video, err := os.Open("video.mp4")

	if err != nil {
		app.serverError(w, r, fmt.Errorf("there was an error"))
	}

	defer video.Close()

	videoInfo, err := video.Stat()

	if err != nil {
		app.serverError(w, r, fmt.Errorf("there was an error"))
	}

	http.ServeContent(w, r, videoInfo.Name(), videoInfo.ModTime(), video)

}
