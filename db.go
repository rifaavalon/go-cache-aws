package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/garyburd/redigo/redis"
)

var currentInstanceID int64

func RedisConnect() redis.Conn {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
	}
	HandleError(err)

	return c
}

func (a *AWS) init() {
	c := RedisConnect()
	defer c.Close()

	svc := ec2.New(a.session)

	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			&ec2.Filter{
				Name: aws.String(""),
				Values: []*string{
					aws.String(""),
				},
			},
		},
	}
	result, err := svc.DescribeInstances(input)

	if err != nil {
		fmt.Println("Error", err)
	}
	for _, reservations := range result.Reservations {
		for _, instance := range reservations.Instances {
			fmt.Println("instance:", *instance.InstanceId)
			output, err1 := json.Marshal(*instance)
			if err1 != nil {
				fmt.Println(err1)
			}
			c.Do("SET", fmt.Sprintf("instance:%s", *instance.InstanceId), output)
		}
	}
}

func FindAll() Instances {
	var instances Instances

	c := RedisConnect()
	defer c.Close()

	keys, err := c.Do("KEYS", "instance:*")
	HandleError(err)

	for _, k := range keys.([]interface{}) {
		var instance Instance
		reply, err := c.Do("GET", k.([]byte))
		HandleError(err)

		if err := json.Unmarshal(reply.([]byte), &instance); err != nil {
			panic(err)
		}
		instances = append(instances, instance)
	}
	return instances
}

func FindInstance(id int) Instance {
	var instance Instance

	c := RedisConnect()
	defer c.Close()

	reply, err := c.Do("GET", "instance:"+strconv.Itoa(id))
	HandleError(err)

	fmt.Println("GET OK")

	if err = json.Unmarshal(reply.([]byte), &instance); err != nil {
		panic(err)
	}
	return instance
}

func CreateInstance(p Instance) {
	currentime := time.Now()

	currentInstanceID := currentime.Unix()

	p.ID = currentInstanceID
	p.Timestamp = time.Now()

	c := RedisConnect()
	defer c.Close()

	p.ID = currentInstanceID
	p.Timestamp = time.Now()

	b, err := json.Marshal(p)
	HandleError(err)

	reply, err := c.Do("SET", "instance:"+strconv.Itoa(int(p.ID)), b)
	HandleError(err)

	fmt.Println("GET", reply)
}
