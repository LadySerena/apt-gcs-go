package apt

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/storage"
)

//basic flow
//init gcs client on start
//send out 100 capabilities message to apt on startup
//receive 600 from apt
//send out 200 start message
//get the thing from gcs
//send 201 if happy or 400 if sad
//there's something weird with the encoding per https://github.com/dhaivat/apt-gcs/issues/1

var (
	codeMap = map[int]string{
		100: "Capabilities",
		101: "Log",
		102: "Status",
		200: "URI Start",
		201: "URI Done",
		400: "URI Failure",
		600: "URI Acquire",
		601: "Configuration",
	}
)

type GCSTransport struct {
	client     *storage.Client
	bucket     *storage.BucketHandle
	bucketName string
}

type Message struct {
	Code int
	Content map[string]string
}


func NewGcsTransport(ctx context.Context, projectID string, bucketName string) (*GCSTransport, error) {
	client, clientCreateErr := storage.NewClient(ctx)
	if clientCreateErr != nil {
		return nil, clientCreateErr
	}
	bucketHandle := client.Bucket(bucketName)

	return &GCSTransport{
		client:     client,
		bucket:     bucketHandle,
		bucketName: bucketName,
	}, nil
}

func (t GCSTransport) Close() error {
	return t.client.Close()
}

// encode takes a plaintext message struct and generates a proper plaintext string for example
// $codeNumber: $debugText\n
// $fieldName: $fieldValue\n
// \n
func (m Message) encode() (string, error) {
	var messageBuffer bytes.Buffer
	codeHeader := fmt.Sprintf("%d: %s\n", m.Code, codeMap[m.Code])
	bytesWritten, writeErr := messageBuffer.WriteString(codeHeader)
	if writeErr != nil {
		return "", writeErr
	}
	if bytesWritten != len(codeHeader) {
		return "", errors.New("bytes written does not match header")
	}
	for key, value := range m.Content {
		if key != "" && value != "" {
			content := fmt.Sprintf("%s: %s\n", key, value)
			lengthWritten, contentWriteErr := messageBuffer.WriteString(content)
			if contentWriteErr != nil {
				return "", contentWriteErr
			}
			if lengthWritten != len(content){
				return "", errors.New("write length mismatch")
			}
		}
	}
	messageBuffer.WriteRune('\n')
	return messageBuffer.String(), nil
}