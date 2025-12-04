package main

import (
	"fmt"
	"net/http"
	"io"
	"encoding/base64"

	"github.com/bootdotdev/learn-file-storage-s3-golang-starter/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUploadThumbnail(w http.ResponseWriter, r *http.Request) {
	videoIDString := r.PathValue("videoID")
	videoID, err := uuid.Parse(videoIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID", err)
		return
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}


	fmt.Println("uploading thumbnail for video", videoID, "by user", userID)

	// TODO: implement the upload here

	const maxMemory = 10 << 20

	err = r.ParseMultipartForm(maxMemory)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't parse maxMemory", err)
		return
	}

	data, fileHeader, err := r.FormFile("thumbnail")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't form file", err)
		return
	}

	mediaType := fileHeader.Header.Get("Content-Type")

	dataBytes, err := io.ReadAll(data)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't read data", err)
		return
	}

	metadata, err := cfg.db.GetVideo(videoID)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't read data", err)
		return
	}

//	thumbnailImage := *new(thumbnail)
//	thumbnailImage.data = dataBytes
//	thumbnailImage.mediaType = mediaType
//
//	videoThumbnails[videoID] = thumbnailImage
//
//	port := os.Getenv("PORT")
//	url := fmt.Sprintf("http://localhost:%s/api/thumbnails/%s", port, fmt.Sprintln(videoID))
//	metadata.ThumbnailURL = &url
//	err = cfg.db.UpdateVideo(metadata)
//	if err != nil {
//		respondWithError(w, http.StatusInternalServerError, "Couldn't read data", err)
//		return
//	}

	// Encoding

	dataString := base64.StdEncoding.EncodeToString(dataBytes)
	dataURL := fmt.Sprintf("data:%s;base64,%s", mediaType, dataString)

	metadata.ThumbnailURL = &dataURL
	err = cfg.db.UpdateVideo(metadata)

	respondWithJSON(w, http.StatusOK, metadata)
}
