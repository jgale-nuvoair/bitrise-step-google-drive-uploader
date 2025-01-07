package main

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


func main() {
	fmt.Println("This is the value specified for the input 'service_key_path':", os.Getenv("service_key_path"))
	fmt.Println("This is the value specified for the input 'folder_id':", os.Getenv("folder_id"))
	fmt.Println("This is the value specified for the input 'output_filename':", os.Getenv("output_filename"))
	fmt.Println("This is the value of BITRISE_XCODEBUILD_TEST_LOG_PATH:", os.Getenv("BITRISE_XCODEBUILD_TEST_LOG_PATH"))

	serviceAccount := os.Getenv("service_key_path")
	folderId := os.Getenv("folder_id")
	outputFilename := os.Getenv("output_filename")
	

	// You can use this to find all files of a particular extension and upload them
	// files, err := utils.FindFiles("*.log")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// for _, file := range files {
	// 	fmt.Println("Trying to upload file: ", file)
	// 	UploadFile(serviceAccount, file, folderId)
	// }

	filePath := os.Getenv("BITRISE_XCODEBUILD_TEST_LOG_PATH")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File does not exist:", filePath)
		os.Exit(1)
	}

	UploadFile(serviceAccount, filePath, outputFilename, folderId)

	// --- Exit codes:
	// The exit code of your Step is very important. If you return
	//  with a 0 exit code `bitrise` will register your Step as "successful".
	// Any non zero exit code will be registered as "failed" by `bitrise`.
	os.Exit(0)
}

func serviceAccount(credentialFile string) *http.Client {
	b, err := ioutil.ReadFile(credentialFile)
	if err != nil {
		fmt.Println("credential file for the service account couldn't be read:", credentialFile)
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

func UploadFile(serviceFile string, fileName string, outputFilename string, folderId string) {
	filename := fileName                       // Filename
	baseMimeType := "application/octet-stream" // MimeType
	client := serviceAccount(serviceFile)      // Please set the json file of Service account.

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
	// get base name from filename: baseName := filepath.Base(filename)
	f := &drive.File{Name: outputFilename}
	if folderId != "" {
		f.Parents = []string{folderId}
	}
	res, err := srv.Files.
		Create(f).
		ResumableMedia(context.Background(), file, fileInf.Size(), baseMimeType).
		ProgressUpdater(func(now, size int64) { fmt.Printf("%d, %d\r", now, size) }).
		Do()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Uploaded file %s to with id %s\n", outputFilename, res.Id)
}