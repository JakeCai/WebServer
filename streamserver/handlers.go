package main 

import (
	"io"
	"os"
	"net/http"
	"html/template"
	"io/ioutil"
	"log"
	"github.com/julienschmidt/httprouter"
)

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("./videos/upload.html")
 
    t.Execute(w, nil)
}


func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("Entered the streamHandler")
	vid := p.ByName("vid-id")
	videoPath := VIDEO_DIR + vid
	video, err := os.Open(videoPath)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	}
	w.Header().Set("Content-type", "video/mp4")
	http.ServeContent(w, r, "", time.now(), video)
	video.close()
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return 
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error when try to get file: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return 
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	}

	filename := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR + filename, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Upload success")
}
