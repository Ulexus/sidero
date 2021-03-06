// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package v1alpha1

import (
	"reflect"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BMC defines data about how to talk to the node via ipmitool.
type BMC struct {
	Endpoint string `json:"endpoint"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
}

// ManagementAPI defines data about how to talk to the node via simple HTTP API.
type ManagementAPI struct {
	Endpoint string `json:"endpoint"`
}

type SystemInformation struct {
	Manufacturer string `json:"manufacturer,omitempty"`
	ProductName  string `json:"productName,omitempty"`
	Version      string `json:"version,omitempty"`
	SerialNumber string `json:"serialNumber,omitempty"`
	SKUNumber    string `json:"skuNumber,omitempty"`
	Family       string `json:"family,omitempty"`
}

func (a *SystemInformation) PartialEqual(b *SystemInformation) bool {
	return PartialEqual(a, b)
}

type CPUInformation struct {
	Manufacturer string `json:"manufacturer,omitempty"`
	Version      string `json:"version,omitempty"`
}

func (a *CPUInformation) PartialEqual(b *CPUInformation) bool {
	return PartialEqual(a, b)
}

func PartialEqual(a, b interface{}) bool {
	old := reflect.ValueOf(a)
	new := reflect.ValueOf(b)

	if old.Kind() == reflect.Ptr {
		old = old.Elem()
	}

	if new.Kind() == reflect.Ptr {
		new = new.Elem()
	}

	for i := 0; i < old.NumField(); i++ {
		if old.Field(i).IsZero() {
			// Skip zero values, since that indicates that the user did not supply
			// the field, and does not want to compare it.
			continue
		}

		f1 := old.Field(i).Interface()
		f2 := new.Field(i).Interface()

		if f1 != f2 {
			return false
		}
	}

	return true
}

// ServerSpec defines the desired state of Server.
type ServerSpec struct {
	EnvironmentRef    *corev1.ObjectReference `json:"environmentRef,omitempty"`
	Hostname          string                  `json:"hostname,omitempty"`
	SystemInformation *SystemInformation      `json:"system,omitempty"`
	CPU               *CPUInformation         `json:"cpu,omitempty"`
	BMC               *BMC                    `json:"bmc,omitempty"`
	ManagementAPI     *ManagementAPI          `json:"managementApi,omitempty"`
	ConfigPatches     []ConfigPatches         `json:"configPatches,omitempty"`
	Accepted          bool                    `json:"accepted"`
}

// ConditionPowerCycle is used to control the powercycle flow.
const ConditionPowerCycle clusterv1.ConditionType = "PowerCycle"

// ServerStatus defines the observed state of Server.
type ServerStatus struct {
	Ready      bool                  `json:"ready"`
	InUse      bool                  `json:"inUse"`
	IsClean    bool                  `json:"isClean"`
	Conditions []clusterv1.Condition `json:"conditions,omitempty"`
	Addresses  []corev1.NodeAddress  `json:"addresses,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:printcolumn:name="Hostname",type="string",JSONPath=".spec.hostname",description="server hostname"
// +kubebuilder:printcolumn:name="Accepted",type="boolean",JSONPath=".spec.accepted",description="indicates if the server is accepted"
// +kubebuilder:printcolumn:name="Allocated",type="boolean",JSONPath=".status.inUse",description="indicates that the server has been allocated"
// +kubebuilder:printcolumn:name="Clean",type="boolean",JSONPath=".status.isClean",description="indicates if the server is clean or not"

// Server is the Schema for the servers API.
type Server struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServerSpec   `json:"spec,omitempty"`
	Status ServerStatus `json:"status,omitempty"`
}

func (s *Server) GetConditions() clusterv1.Conditions {
	return s.Status.Conditions
}

func (s *Server) SetConditions(conditions clusterv1.Conditions) {
	s.Status.Conditions = conditions
}

// +kubebuilder:object:root=true

// ServerList contains a list of Server.
type ServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Server `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Server{}, &ServerList{})
}
