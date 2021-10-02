package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type IFileUpload interface {
	Upload(string, []graphql.Upload)
}

type fileupload struct{}

func NewFileUpload() IFileUpload {
	return &fileupload{}
}

func (service *fileupload) Upload(folder string, files []graphql.Upload) {
	// Step 1: Get the Google Drive service Client
	client := getClient("google_service_account_key.json")

	driveService, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve drive Client %v", err)
	}
	// Step 2. Create the directory
	dir, err := createDir(driveService, folder, os.Getenv("GOOGLE_DRIVE_FOLDER_ID"))

	if err != nil {
		panic(fmt.Sprintf("Could not create dir: %v\n", err))
	}

	// Step 3: create the file and upload

	for i := 0; i < len(files); i++ {
		input := files[i]
		_, err = createFile(driveService, input.Filename, input.ContentType, input.File, dir.Id)
	}

	if err != nil {
		panic(fmt.Sprintf("Could not create file: %v\n", err))
	}
}

func getClient(secretFile string) *http.Client {
	b, err := ioutil.ReadFile(secretFile)
	if err != nil {
		log.Fatal("error while reading the credential file", err)
	}
	var s = struct {
		Email      string `json:"client_email"`
		PrivateKey string `json:"private_key"`
	}{}
	json.Unmarshal(b, &s)
	config := &jwt.Config{
		Email:      s.Email,
		PrivateKey: []byte(s.PrivateKey),
		Scopes: []string{
			drive.DriveScope,
		},
		TokenURL: google.JWTTokenURL,
	}
	client := config.Client(context.Background())
	return client
}

func createDir(service *drive.Service, name string, parentId string) (*drive.File, error) {
	d := &drive.File{
		Name:     name,
		MimeType: "application/vnd.google-apps.folder",
		Parents:  []string{parentId},
	}

	// folders, err := service.Files.List().Do()

	// fmt.Printf("%+v\n", folders)

	file, err := service.Files.Create(d).Do()

	if err != nil {
		log.Println("Could not create dir: " + err.Error())
		return nil, err
	}

	return file, nil
}

func createFile(service *drive.Service, name string, mimeType string, content io.Reader, parentId string) (*drive.File, error) {
	f := &drive.File{
		MimeType: mimeType,
		Name:     name,
		Parents:  []string{parentId},
	}
	file, err := service.Files.Create(f).Media(content).Do()

	if err != nil {
		log.Println("Could not create file: " + err.Error())
		return nil, err
	}

	return file, nil
}
