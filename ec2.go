package main

import (
	"encoding/json"
	"time"
)

type Instance struct {
	ID                     int64 `json:"id"`
	InstanceId             *json.RawMessage
	Architecture           *json.RawMessage
	AvailablityZone        *json.RawMessage
	BlockDeviceMappings    *json.RawMessage
	ClientToken            *json.RawMessage
	DNSName                *json.RawMessage
	GroupId                *json.RawMessage
	HostId                 *json.RawMessage
	Hypervisor             *json.RawMessage
	IamInstanceProfile     *json.RawMessage
	ImageId                *json.RawMessage
	InstanceLifecycle      *json.RawMessage
	InstanceStateCode      *json.RawMessage
	InstanceStateName      *json.RawMessage
	InstanceType           *json.RawMessage
	InstanceGroupId        *json.RawMessage
	InstanceGroupName      *json.RawMessage
	IPAddress              *json.RawMessage
	KernelId               *json.RawMessage
	KeyName                *json.RawMessage
	LaunchIndex            *json.RawMessage
	LaunchTime             *json.RawMessage
	MontoringState         *json.RawMessage
	NetworkInterfaces      *json.RawMessage
	OwnerId                *json.RawMessage
	PlacementGroupName     *json.RawMessage
	Platform               *json.RawMessage
	PrivateDNSNAme         *json.RawMessage
	PrivateIPAddress       *json.RawMessage
	ProductCodes           *json.RawMessage
	RamdiskId              *json.RawMessage
	Reason                 *json.RawMessage
	RequesterId            *json.RawMessage
	ReservationId          *json.RawMessage
	RootDeviceName         *json.RawMessage
	RootDeviceType         *json.RawMessage
	SourceDestinationCheck *json.RawMessage
	SubnetId               *json.RawMessage
	TagKey                 *json.RawMessage
	TagValue               *json.RawMessage
	Tenancy                *json.RawMessage
	VirtualizationType     *json.RawMessage
	VpcId                  *json.RawMessage
	Timestamp              time.Time `json:"timestamp"`
}

type Instances []Instance
