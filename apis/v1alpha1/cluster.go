// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterSpec defines the desired state of Cluster.
//
// Returns information about a cluster of either the provisioned or the serverless
// type.
type ClusterSpec struct {

	// Information about the brokers.
	// +kubebuilder:validation:Required
	BrokerNodeGroupInfo *BrokerNodeGroupInfo `json:"brokerNodeGroupInfo"`
	// Includes all client authentication related information.
	ClientAuthentication *ClientAuthentication `json:"clientAuthentication,omitempty"`
	// Represents the configuration that you want MSK to use for the cluster.
	ConfigurationInfo *ConfigurationInfo `json:"configurationInfo,omitempty"`
	// Includes all encryption-related information.
	EncryptionInfo *EncryptionInfo `json:"encryptionInfo,omitempty"`
	// Specifies the level of monitoring for the MSK cluster. The possible values
	// are DEFAULT, PER_BROKER, PER_TOPIC_PER_BROKER, and PER_TOPIC_PER_PARTITION.
	EnhancedMonitoring *string `json:"enhancedMonitoring,omitempty"`
	// The version of Apache Kafka.
	// +kubebuilder:validation:Required
	KafkaVersion *string `json:"kafkaVersion"`
	// LoggingInfo details.
	LoggingInfo *LoggingInfo `json:"loggingInfo,omitempty"`
	// The name of the cluster.
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// The number of Apache Kafka broker nodes in the Amazon MSK cluster.
	// +kubebuilder:validation:Required
	NumberOfBrokerNodes *int64 `json:"numberOfBrokerNodes"`
	// The settings for open monitoring.
	OpenMonitoring *OpenMonitoringInfo `json:"openMonitoring,omitempty"`
	// This controls storage mode for supported storage tiers.
	StorageMode *string `json:"storageMode,omitempty"`
	// Create tags when creating the cluster.
	Tags map[string]*string `json:"tags,omitempty"`
}

// ClusterStatus defines the observed state of Cluster
type ClusterStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The state of the cluster. The possible states are ACTIVE, CREATING, DELETING,
	// FAILED, HEALING, MAINTENANCE, REBOOTING_BROKER, and UPDATING.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty"`
}

// Cluster is the Schema for the Clusters API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec,omitempty"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// ClusterList contains a list of Cluster
// +kubebuilder:object:root=true
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
