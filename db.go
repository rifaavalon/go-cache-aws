package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoURI          = "mongodb://localhost:27017"
	databaseName      = "gocache"
	collectionName    = "instances"
	connectionTimeout = 10 * time.Second
)

var currentInstanceID int64

func MongoConnect() (*mongo.Client, context.Context, context.CancelFunc) {
	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), connectionTimeout)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	HandleError(err)

	return client, ctx, cancel
}

func (a *AWS) init() error {
	client, ctx, cancel := MongoConnect()
	defer client.Disconnect(ctx)
	defer cancel()

	svc := ec2.New(a.session)

	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String(""),
				Values: []*string{
					aws.String(""),
				},
			},
		},
	}
	result, err := svc.DescribeInstances(input)
	if err != nil {
		return fmt.Errorf("failed to describe instances: %w", err)
	}

	collection := client.Database(databaseName).Collection(collectionName)
	for _, reservations := range result.Reservations {
		for _, instance := range reservations.Instances {
			fmt.Println("instance:", *instance)
			ec2Instance := EC2Instance{
				// Map fields from AWS SDK instance to EC2Instance struct
			}
			_, err := collection.InsertOne(ctx, ec2Instance)
			if err != nil {
				log.Printf("Failed to insert instance: %v", err)
			}
		}
	}
	return nil
}

func FindAll() Instances {
	var instances Instances

	client, ctx, cancel := MongoConnect()
	defer client.Disconnect(ctx)
	defer cancel()

	collection := client.Database(databaseName).Collection(collectionName)
	cursor, err := collection.Find(ctx, bson.M{})
	HandleError(err)
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var instance EC2Instance
		err := cursor.Decode(&instance)
		HandleError(err)
		instances = append(instances, instance)
	}
	return instances
}

func FindInstance(id int) EC2Instance {
	var instance EC2Instance

	client, ctx, cancel := MongoConnect()
	defer client.Disconnect(ctx)
	defer cancel()

	collection := client.Database(databaseName).Collection(collectionName)
	err := collection.FindOne(ctx, bson.M{"id": id}).Decode(&instance)
	HandleError(err)

	fmt.Println("GET OK")

	return instance
}

func CreateInstance(p EC2Instance) {
	currentime := time.Now()

	currentInstanceID := currentime.Unix()

	p.ID = currentInstanceID
	p.Timestamp = time.Now()

	client, ctx, cancel := MongoConnect()
	defer client.Disconnect(ctx)
	defer cancel()

	collection := client.Database(databaseName).Collection(collectionName)
	_, err := collection.InsertOne(ctx, p)
	HandleError(err)

	fmt.Println("Instance created")
}

func HandleError(err error) {
	if err != nil {
		log.Printf("Error: %v", err)
	}
}
