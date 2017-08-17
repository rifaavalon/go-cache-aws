package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

//AWS contains session and SQS variables
type AWS struct {
	session     *session.Session
	acreds      *credentials.Credentials
	assumedRole bool
}

//Open starts AWS connections
func (a *AWS) Open(region string) error {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String(region)},
	})
	if err != nil {
		log.Fatal(err)
	}

	a.session = sess

	return err
}

//Close destroys all AWS connecitons
func (a *AWS) Close() {
	a.session = nil
}
