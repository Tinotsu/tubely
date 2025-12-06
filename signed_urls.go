package main
//
//import (
//	"time"
//	"log"
//	"strings"
//	"github.com/aws/aws-sdk-go-v2/service/s3"
//	"github.com/bootdotdev/learn-file-storage-s3-golang-starter/internal/database"
//	"context"
//)
//
//func generatePresignedURL(s3Client *s3.Client, bucket, key string, expireTime time.Duration) (string, error) {
//	presignedClient := s3.NewPresignClient(s3Client)
//	ctx := context.Background()
//	expiration := s3.WithPresignExpires(expireTime)
//	objectInput := new(s3.GetObjectInput)
//	objectInput.Bucket = &bucket
//	objectInput.Key = &key
//	objectSigned, err := presignedClient.PresignGetObject(ctx, objectInput, expiration)
//	if err != nil {
//		log.Print("Couldn't get presignedClient: ", err)
//		return "", err
//	}
//	return objectSigned.URL, nil
//}
//
//func (cfg *apiConfig) dbVideoToSignedVideo(video database.Video) (database.Video, error) {
//	if video.VideoURL == nil {
//		//err := errors.New("VideoURL is nil")
//		return video, nil
//	}
//	if !strings.Contains(*video.VideoURL, ",") {
//		//err := errors.New("VideoURL doesn't contains ','")
//		return video, nil
//	}
//	l := strings.Split(*video.VideoURL, ",")
//	bucket := l[0]
//	key := l[1]
//	expireTime := 300*time.Second
//	presignedURL, err := generatePresignedURL(cfg.s3Client, bucket, key, expireTime)
//	if err != nil {
//		log.Print("Couldn't get presignedURL: ", err)
//		return database.Video{}, err
//	}
//	video.VideoURL = &presignedURL
//	return video, nil
//}
