/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AuthorityObservation struct {
}

type AuthorityParameters struct {

	// A JSON Web Token (JWT) issuer URI. issuer must start with https:// and // be a valid
	// with length <2000 characters. For example: https://container.googleapis.com/v1/projects/my-project/locations/us-west1/clusters/my-cluster (must be locations rather than zones).googleapis.com/v1/${google_container_cluster.my-cluster.id}".
	// +kubebuilder:validation:Required
	Issuer *string `json:"issuer" tf:"issuer,omitempty"`
}

type EndpointObservation struct {
}

type EndpointParameters struct {

	// If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	GkeCluster []GkeClusterParameters `json:"gkeCluster,omitempty" tf:"gke_cluster,omitempty"`
}

type GkeClusterObservation struct {
}

type GkeClusterParameters struct {

	// Self-link of the GCP resource for the GKE cluster.
	// For example: //container.googleapis.com/projects/my-project/zones/us-west1-a/clusters/my-cluster.
	// It can be at the most 1000 characters in length.googleapis.com/${google_container_cluster.my-cluster.id}" or
	// google_container_cluster.my-cluster.id.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/container/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ResourceLink *string `json:"resourceLink,omitempty" tf:"resource_link,omitempty"`

	// Reference to a Cluster in container to populate resourceLink.
	// +kubebuilder:validation:Optional
	ResourceLinkRef *v1.Reference `json:"resourceLinkRef,omitempty" tf:"-"`

	// Selector for a Cluster in container to populate resourceLink.
	// +kubebuilder:validation:Optional
	ResourceLinkSelector *v1.Selector `json:"resourceLinkSelector,omitempty" tf:"-"`
}

type MembershipObservation struct {

	// an identifier for the resource with format {{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The unique identifier of the membership.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type MembershipParameters struct {

	// Authority encodes how Google will recognize identities from this Membership.
	// See the workload identity documentation for more details:
	// https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Authority []AuthorityParameters `json:"authority,omitempty" tf:"authority,omitempty"`

	// If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Endpoint []EndpointParameters `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// Labels to apply to this membership.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// MembershipSpec defines the desired state of Membership
type MembershipSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MembershipParameters `json:"forProvider"`
}

// MembershipStatus defines the observed state of Membership.
type MembershipStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MembershipObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Membership is the Schema for the Memberships API. Membership contains information about a member cluster.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Membership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MembershipSpec   `json:"spec"`
	Status            MembershipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MembershipList contains a list of Memberships
type MembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Membership `json:"items"`
}

// Repository type metadata.
var (
	Membership_Kind             = "Membership"
	Membership_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Membership_Kind}.String()
	Membership_KindAPIVersion   = Membership_Kind + "." + CRDGroupVersion.String()
	Membership_GroupVersionKind = CRDGroupVersion.WithKind(Membership_Kind)
)

func init() {
	SchemeBuilder.Register(&Membership{}, &MembershipList{})
}