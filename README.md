# bitrise-step-google-drive-uploader

## Google Service Account

Follow these steps to create a Service Account, and get the credentials file:
https://www.perplexity.ai/search/how-to-create-a-google-drive-s-.iyDFF9VSByMpZmCdYgJAg

1. Upload this file to Bitrise > your-project > Project Settings > Files
2. Set the Download URL to be named BITRISEIO_GOOGLE_DRIVE_SERVICE_ACCOUNT_CREDENTIALS_URL

## Google Drive Folder ID
Using the Google Drive API we can upload files to a specified google drive folder using a service account and sharing the folder to that service account.

```
https://drive.google.com/drive/folders/{folderId}
```

Get the Folder ID here, and add it as an Environment Variable named GOOGLE_DRIVE_FOLDER_ID under Bitrise > your-project > Workflow Editor > Env Vars.


## Set up Bitrise Workflows

Modify the bitrise.yml to add the following steps:

```
    - file-downloader:
        inputs:
        - source: "$BITRISEIO_GOOGLE_DRIVE_SERVICE_ACCOUNT_CREDENTIALS_URL"
        - destination: "./credentials.json"
    - git::https://github.com/jgale-nuvoair/bitrise-step-google-drive-uploader@main:
        inputs:
        - folder_id: "$GOOGLE_DRIVE_FOLDER_ID"
        - service_key_path: credentials.json
```

This sets up the Bitrise [File Downloader](https://bitrise.io/integrations/steps/file-downloader) step to download the JSON file during your Bitrise workflow, using the Download URL from previous step.

It configures the uploader step with the Google Drive folder environment variable, and the downloaded credentials file.

