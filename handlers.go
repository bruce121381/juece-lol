package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"time"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	videoId := p.ByName("vid-id")
	videoLink := VIDEO_DIR + videoId
	video, err := os.Open(videoLink)
	if err != nil {
		log.Printf("open video is failed, err is %s", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
	defer video.Close()

}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
