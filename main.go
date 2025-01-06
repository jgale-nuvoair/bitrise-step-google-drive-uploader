package main

import (
	"fmt"
	"os"

	"github.com/FriendlyUser/bitrise-step-google-drive-uploader/pkg/utils"
)

func main() {
	// print service_key_path
	fmt.Println("This is the value specified for the input 'service_key_path':", os.Getenv("service_key_path"))
	// and print folder_id
	fmt.Println("This is the value specified for the input 'folder_id':", os.Getenv("folder_id"))

	serviceAccount := os.Getenv("service_key_path")
	folderId := os.Getenv("folder_id")

	// find all files with the extension ending with *.log
	files, err := utils.FindFiles("**/*.log")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// upload all apk files
	for _, file := range files {
		fmt.Println("Trying to upload file: ", file)
		utils.UploadFile(serviceAccount, file, folderId)
	}

	// --- Exit codes:
	// The exit code of your Step is very important. If you return
	//  with a 0 exit code `bitrise` will register your Step as "successful".
	// Any non zero exit code will be registered as "failed" by `bitrise`.
	os.Exit(0)
}
