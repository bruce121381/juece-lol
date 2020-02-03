package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

/**
播放视频
*/
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

/**
上传视频
*/
func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//上传视频最大大小
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)

	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		log.Printf("Bad Request, err is %s", err)
		sendErrorResponse(w, http.StatusBadRequest, "Fail is too large")
		return
	}

	file, _, err := r.FormFile("file") //<file name = "file"
	if err != nil {
		log.Printf("file read is failed, err is %s", err)
		sendErrorResponse(w, http.StatusInternalServerError, "failed read file")
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file is failed, err is %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
	}

	fileName := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+fileName, data, 0666)
	if err != nil {
		log.Printf("file write error, err is %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "file open failed")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Upload is success")
}
