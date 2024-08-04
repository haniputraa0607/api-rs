package utility

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
)

func serviceAccount(credentialFile string) (*jwt.Config, error) {
	b, err := os.ReadFile(credentialFile)
	if err != nil {
		return nil, err
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
			"https://www.googleapis.com/auth/drive",
		},
		TokenURL: google.JWTTokenURL,
	}
	config.Subject = os.Getenv("GOOGLE_DRIVE_EMAIL")
	return config, nil
}

func SaveToDrive(file *multipart.FileHeader) (id string, err error) {
	config, err := serviceAccount(os.Getenv("GOOGLE_DRIVE_CRED"))
	if err != nil {
		return
	}

	ctx := context.Background()
	client := config.Client(ctx)

	driveService, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create Drive service: %v", err)
	}

	src, err := file.Open()
	if err != nil {
		fmt.Println(1)
		return
	}
	defer src.Close()

	driveFile := &drive.File{Name: file.Filename, Parents: []string{os.Getenv("GOOGLE_DRIVE_FOLDER_ID")}}
	driveFile, err = driveService.Files.Create(driveFile).Media(src).Do()
	if err == nil {
		id = driveFile.Id
	}

	return

}

func GetFileFromDrive(fileID string) (fileContent io.ReadCloser, err error) {
	config, err := serviceAccount(os.Getenv("GOOGLE_DRIVE_CRED"))
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	client := config.Client(ctx)

	driveService, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create Drive service: %v", err)
	}

	resp, err := driveService.Files.Get(fileID).Download()
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func CheckFileExists(driveService *drive.Service, fileID string) (bool, error) {
	_, err := driveService.Files.Get(fileID).Fields("id").Do()
	if err != nil {
		if gErr, ok := err.(*googleapi.Error); ok && gErr.Code == http.StatusNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func DeleteFileFromDrive(fileID string) (err error) {
	config, err := serviceAccount(os.Getenv("GOOGLE_DRIVE_CRED"))
	if err != nil {
		return err
	}

	ctx := context.Background()
	client := config.Client(ctx)

	driveService, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create Drive service: %v", err)
	}

	exists, err := CheckFileExists(driveService, fileID)
	if err != nil {
		return err
	}

	if exists {
		err = driveService.Files.Delete(fileID).Do()
		if err != nil {
			return err
		}
	}

	return err
}
