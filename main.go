package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"

	ie2utilities "github.com/insightengine2/ie2-utilities/utils"
)

// env variable names
const ENV_METADATA_FILENAME = "ENV_METADATA_FILENAME"
const ENV_EMBEDDING_URL = "ENV_EMBEDDING_URL"
const ENV_EMBEDDING_API_KEY = "ENV_EMBEDDING_API_KEY"
const ENV_EMBEDDING_API_HEADER = "ENV_EMBEDDING_API_HEADER"

// default values
const DEF_METADATA_FILENAME = "metadata.json"
const DEF_EMBEDDING_API_HEADER = "x-api-key"

type filevent struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

func HandleRequest(context context.Context, ev filevent) error {

	log.Print("Testing autodeployment via CICD")

	log.Print("Parsing event...")
	log.Print("Starting new Session...")

	conf, err := config.LoadDefaultConfig(context)

	if err != nil {
		log.Printf("Failed to load default config: %s", err)
		return err
	}

	log.Print("Retrieving env variables.")

	// get env variables
	var metadatafilename = os.Getenv(ENV_METADATA_FILENAME)
	var embeddingurl = os.Getenv(ENV_EMBEDDING_URL)
	var embeddingapikey = os.Getenv(ENV_EMBEDDING_API_KEY)
	var embeddingapiheader = os.Getenv(ENV_EMBEDDING_API_HEADER)

	// required
	if len(embeddingurl) <= 0 {
		msg := fmt.Sprintf("embedding url is empty or undefined: please define a value for the env variable %s", ENV_EMBEDDING_URL)
		log.Print(msg)
		log.Print("Exiting Text2Vec...")
		return errors.New(msg)
	}

	// optional
	if len(metadatafilename) <= 0 {
		log.Print("Metadata env variable is NOT set.")
		log.Printf("Using default: %s.", DEF_METADATA_FILENAME)
		metadatafilename = DEF_METADATA_FILENAME
	}

	// optional
	if len(embeddingapikey) <= 0 {
		msg := "embedding api key is empty or undefined: using an empty string"
		log.Print(msg)
		embeddingapikey = ""
	}

	// optional
	if len(embeddingapiheader) <= 0 {
		msg := fmt.Sprintf("embedding api header name is empty or undefined: using the default value %s", DEF_EMBEDDING_API_HEADER)
		log.Print(msg)
		embeddingapiheader = DEF_EMBEDDING_API_HEADER
	}

	accountid, err := ie2utilities.AWSGetAccountId(&conf, &context)

	if err != nil {
		log.Print("Failed to get AWS Account Id.")
		return err
	}

	log.Printf("Using Account Id %s...", accountid)

	log.Print("Logging event details...")
	log.Print(ev)

	if strings.Contains(ev.Key, metadatafilename) {
		log.Print("File uploaded is metadata...ignoring.")
		return nil
	}

	log.Print("Exiting Text2Vec...")

	return nil
}

// entry point to your lambda
func main() {
	lambda.Start(HandleRequest)
}
