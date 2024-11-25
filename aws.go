package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// AWS contains session variables
type AWS struct {
	session *session.Session
}

// Open starts AWS connections
func (a *AWS) Open(region string) error {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String(region)},
	})
	if err != nil {
		return err
	}

	a.session = sess
	return nil
}

// Close destroys all AWS connections
func (a *AWS) Close() {
	a.session = nil
}

// mainAWS demonstrates how to use the AWS struct
func mainAWS() {
	awsInstance := &AWS{}
	err := awsInstance.Open("us-west-2")
	if err != nil {
		log.Fatalf("Failed to open AWS session: %v", err)
	}
	defer awsInstance.Close()

	// Your code here
}
