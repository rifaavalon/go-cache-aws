package main

import "time"

type EC2Instance struct {
	ID                    int64     `json:"id"`
	OwnerID               string    `json:"owner_id"`
	Instances             []string  `json:"instances,omitempty"`
	InstanceId            string    `json:"instance_id,omitempty"`
	BlockDeviceMappings   []string  `json:"block_device_mappings,omitempty"`
	ReservationID         string    `json:"reservation_id,omitempty"`
	SecurityGroups        []string  `json:"security_groups,omitempty"`
	NetworkInterfaces     []string  `json:"network_interfaces,omitempty"`
	Attachment            string    `json:"attachment,omitempty"`
	Groups                []string  `json:"groups,omitempty"`
	Hypervisor            string    `json:"hypervisor,omitempty"`
	RootDeviceName        string    `json:"root_device_name,omitempty"`
	IamInstanceProfile    string    `json:"iam_instance_profile,omitempty"`
	VirtualizationType    string    `json:"virtualization_type,omitempty"`
	Tags                  []string  `json:"tags,omitempty"`
	AmiLaunchIndex        int       `json:"ami_launch_index,omitempty"`
	EbsOptimized          bool      `json:"ebs_optimized,omitempty"`
	State                 string    `json:"state,omitempty"`
	Platform              string    `json:"platform,omitempty"`
	PublicDNSName         string    `json:"public_dns_name,omitempty"`
	VpcID                 string    `json:"vpc_id,omitempty"`
	StateTransitionReason string    `json:"state_transition_reason,omitempty"`
	ImageID               string    `json:"image_id,omitempty"`
	ClientToken           string    `json:"client_token,omitempty"`
	SubnetID              string    `json:"subnet_id,omitempty"`
	InstanceType          string    `json:"instance_type,omitempty"`
	Placement             string    `json:"placement,omitempty"`
	LaunchTime            time.Time `json:"launch_time"`
	Timestamp             time.Time `json:"timestamp"`
}

type Instances []EC2Instance
