package main

import (
	"fmt"
	"os"

	"github.com/FriendlyUser/bitrise-step-google-drive-uploader/pkg/utils"
)

func main() {
	fmt.Println("This is the value specified for the input 'example_step_input':", os.Getenv("example_step_input"))
	// print service_key_path
	fmt.Println("This is the value specified for the input 'service_key_path':", os.Getenv("service_key_path"))
	// and print folder_id
	fmt.Println("This is the value specified for the input 'folder_id':", os.Getenv("folder_id"))

	serviceAccount := os.Getenv("service_key_path")
	folderId := os.Getenv("folder_id")

	// find all files with the extension ending with *.apk
	files, err := utils.FindFiles("**/*.apk")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// upload all apk files
	for _, file := range files {
		utils.UploadFile(serviceAccount, file, folderId)
	}

	//
	// --- Step Outputs: Export Environment Variables for other Steps:
	// You can export Environment Variables for other Steps with
	//  envman, which is automatically installed by `bitrise setup`.
	// A very simple example:
	// cmdLog, err := exec.Command("bitrise", "envman", "add", "--key", "EXAMPLE_STEP_OUTPUT", "--value", "the value you want to share").CombinedOutput()
	// if err != nil {
	// 	fmt.Printf("Failed to expose output with envman, error: %#v | output: %s", err, cmdLog)
	// 	os.Exit(1)
	// }
	// You can find more usage examples on envman's GitHub page
	//  at: https://github.com/bitrise-io/envman

	//
	// --- Exit codes:
	// The exit code of your Step is very important. If you return
	//  with a 0 exit code `bitrise` will register your Step as "successful".
	// Any non zero exit code will be registered as "failed" by `bitrise`.
	os.Exit(0)
}
