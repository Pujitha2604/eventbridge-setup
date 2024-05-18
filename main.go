package main

import (
    "context"
    "fmt"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/Pujitha2604/eventbridge-setup/parameterstore"
    "github.com/Pujitha2604/eventbridge-setup/mongodb"
)

func handler(ctx context.Context, event events.CloudWatchEvent) error {
    // Extract image tag from the event
    imageTag := getImageTag(event)

    // Save image tag to Parameter Store
    err := parameterstore.SaveImageTagToParameterStore(imageTag)
    if err != nil {
        return err
    }

    // Save details in MongoDB
    err = mongodb.SaveDetailsInMongoDB(imageTag)
    if err != nil {
        return err
    }

    fmt.Printf("Image tag %s saved to Parameter Store and MongoDB\n", imageTag)
    return nil
}

func getImageTag(event events.CloudWatchEvent) string {
    // Implement the logic to extract the image tag from the event here
}

func main() {
    lambda.Start(handler)
}
