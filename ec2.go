package main

import "time"

type Instance struct {
	ID                    int64  `json:"id"`
	OwnerID               string `json:"OwnerId"`
	Instances             interface{}
	InstanceId            interface{}
	BlockDeviceMappings   interface{}
	ReservationID         string `json:"ReservationId,omitempty"`
	SecurityGroups        interface{}
	NetworkInterfaces     interface{}
	Attachment            interface{}
	Groups                interface{}
	Hypervisor            string `json:"Hypervisor,omitempty"`
	RootDeviceName        string `json:"RootDeviceName,omitempty"`
	IamInstanceProfile    interface{}
	VirtualizationType    string `json:"VirtualizationType,omitempty"`
	Tags                  interface{}
	AmiLaunchIndex        int  `json:"AmiLaunchIndex,omitempty"`
	EbsOptimized          bool `json:"EbsOptimized,omitempty"`
	State                 interface{}
	Platform              string `json:"Platform,omitempty"`
	PublicDNSName         string `json:"PublicDnsName,omitempty"`
	VpcID                 string `json:"VpcId,omitempty"`
	StateTransitionReason string `json:"StateTransitionReason,omitempty"`
	ImageID               string `json:"ImageId,omitempty"`
	ClientToken           string `json:"ClientToken,omitempty"`
	SubnetID              string `json:"SubnetId,omitempty"`
	InstanceType          string `json:"InstanceType,omitempty"`
	Placement             interface{}
	LaunchTime            time.Time `json:"LaunchTime"`
	Timestamp             time.Time `json:"timestamp"`
}

type Instances []Instance
