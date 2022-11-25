package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	drive "google.golang.org/api/drive/v3"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

// ServiceAccount : Use Service account
func ServiceAccount(credentialFile string) *http.Client {
	b, err := ioutil.ReadFile(credentialFile)
	if err != nil {
		log.Fatal(err)
	}
	var c = struct {
		Email      string `json:"client_email"`
		PrivateKey string `json:"private_key"`
	}{}
	json.Unmarshal(b, &c)
	config := &jwt.Config{
		Email:      c.Email,
		PrivateKey: []byte(c.PrivateKey),
		Scopes: []string{
			drive.DriveScope,
		},
		TokenURL: google.JWTTokenURL,
	}
	client := config.Client(oauth2.NoContext)
	return client
}

func UploadFile(serviceFile string, fileName string, mimeType string) {
	filename := fileName                       // Filename
	baseMimeType := "application/octet-stream" // MimeType
	client := ServiceAccount(serviceFile)      // Please set the json file of Service account.

	srv, err := drive.New(client)
	if err != nil {
		log.Fatalln(err)
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	fileInf, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	// to add parents, we may need [folder_id]
	f := &drive.File{Name: filename}
	res, err := srv.Files.
		Create(f).
		ResumableMedia(context.Background(), file, fileInf.Size(), baseMimeType).
		ProgressUpdater(func(now, size int64) { fmt.Printf("%d, %d\r", now, size) }).
		Do()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", res.Id)
}
