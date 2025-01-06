package main

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/FriendlyUser/bitrise-step-google-drive-uploader/pkg/utils"
)

func main() {
	// print service_key_path
	fmt.Println("This is the value specified for the input 'service_key_path':", os.Getenv("service_key_path"))
	// and print folder_id
	fmt.Println("This is the value specified for the input 'folder_id':", os.Getenv("folder_id"))
	// And Xcode log path
	fmt.Println("This is the value of BITRISE_XCODEBUILD_TEST_LOG_PATH:", os.Getenv("BITRISE_XCODEBUILD_TEST_LOG_PATH"))
	fmt.Println("This is the value of BITRISE_TEST_DEPLOY_DIR:", os.Getenv("BITRISE_TEST_DEPLOY_DIR"))

	

	serviceAccount := os.Getenv("service_key_path")
	folderId := os.Getenv("folder_id")

	// find all files with the extension ending with *.log
	//files, err := utils.FindFiles("**/*.log")
	files, err := utils.FindFiles("*.log")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if os.Getenv("BITRISE_XCODEBUILD_TEST_LOG_PATH") != "" {
		// if it is set, add it to the files slice
		files = append(files, os.Getenv("BITRISE_XCODEBUILD_TEST_LOG_PATH"))
	}

	absFilePath, err := filepath.Abs(os.Getenv("BITRISE_XCODEBUILD_TEST_LOG_PATH"))
	if err != nil {
		fmt.Println("Error getting absolute path: %v", err)
	}
	fmt.Println("absFilePath:", absFilePath)


	filePath := os.Getenv("BITRISE_XCODEBUILD_TEST_LOG_PATH")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File does not exist:", filePath)
		os.Exit(1)
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Could not open file:", filePath)
		log.Println(err)
	}


	utils.UploadFile(serviceAccount, os.Getenv("BITRISE_XCODEBUILD_TEST_LOG_PATH"), folderId)

	// upload all apk files
	//for _, file := range files {
		//fmt.Println("Trying to upload file: ", file)
		//utils.UploadFile(serviceAccount, file, folderId)
	//}

	// --- Exit codes:
	// The exit code of your Step is very important. If you return
	//  with a 0 exit code `bitrise` will register your Step as "successful".
	// Any non zero exit code will be registered as "failed" by `bitrise`.
	os.Exit(0)
}
