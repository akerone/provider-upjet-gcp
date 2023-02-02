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

type AnonymousObservation struct {
}

type AnonymousParameters struct {

	// Whether phone number auth is enabled for the project or not.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type EmailObservation struct {
}

type EmailParameters struct {

	// Whether phone number auth is enabled for the project or not.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Whether a password is required for email auth or not. If true, both an email and
	// password must be provided to sign in. If false, a user may sign in via either
	// email/password or email link.
	// +kubebuilder:validation:Optional
	PasswordRequired *bool `json:"passwordRequired,omitempty" tf:"password_required,omitempty"`
}

type HashConfigObservation struct {

	// Different password hash algorithms used in Identity Toolkit.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// Memory cost for hash calculation. Used by scrypt and other similar password derivation algorithms. See https://tools.ietf.org/html/rfc7914 for explanation of field.
	MemoryCost *float64 `json:"memoryCost,omitempty" tf:"memory_cost,omitempty"`

	// How many rounds for hash calculation. Used by scrypt and other similar password derivation algorithms.
	Rounds *float64 `json:"rounds,omitempty" tf:"rounds,omitempty"`

	// Non-printable character to be inserted between the salt and plain text password in base64.
	SaltSeparator *string `json:"saltSeparator,omitempty" tf:"salt_separator,omitempty"`

	// Signer key in base64.
	SignerKey *string `json:"signerKey,omitempty" tf:"signer_key,omitempty"`
}

type HashConfigParameters struct {
}

type PhoneNumberObservation struct {
}

type PhoneNumberParameters struct {

	// Whether phone number auth is enabled for the project or not.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A map of <test phone number, fake code> that can be used for phone auth testing.
	// +kubebuilder:validation:Optional
	TestPhoneNumbers map[string]*string `json:"testPhoneNumbers,omitempty" tf:"test_phone_numbers,omitempty"`
}

type ProjectDefaultConfigObservation struct {

	// an identifier for the resource with format {{project}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Config resource. Example: "projects/my-awesome-project/config"
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Configuration related to local sign in methods.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SignIn []SignInObservation `json:"signIn,omitempty" tf:"sign_in,omitempty"`
}

type ProjectDefaultConfigParameters struct {

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Configuration related to local sign in methods.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SignIn []SignInParameters `json:"signIn,omitempty" tf:"sign_in,omitempty"`
}

type SignInObservation struct {

	// Output only. Hash config information.
	// Structure is documented below.
	HashConfig []HashConfigObservation `json:"hashConfig,omitempty" tf:"hash_config,omitempty"`
}

type SignInParameters struct {

	// Whether to allow more than one account to have the same email.
	// +kubebuilder:validation:Optional
	AllowDuplicateEmails *bool `json:"allowDuplicateEmails,omitempty" tf:"allow_duplicate_emails,omitempty"`

	// Configuration options related to authenticating an anonymous user.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Anonymous []AnonymousParameters `json:"anonymous,omitempty" tf:"anonymous,omitempty"`

	// Configuration options related to authenticating a user by their email address.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Email []EmailParameters `json:"email,omitempty" tf:"email,omitempty"`

	// Configuration options related to authenticated a user by their phone number.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	PhoneNumber []PhoneNumberParameters `json:"phoneNumber,omitempty" tf:"phone_number,omitempty"`
}

// ProjectDefaultConfigSpec defines the desired state of ProjectDefaultConfig
type ProjectDefaultConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectDefaultConfigParameters `json:"forProvider"`
}

// ProjectDefaultConfigStatus defines the observed state of ProjectDefaultConfig.
type ProjectDefaultConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectDefaultConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectDefaultConfig is the Schema for the ProjectDefaultConfigs API. There is no persistent data associated with this resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ProjectDefaultConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectDefaultConfigSpec   `json:"spec"`
	Status            ProjectDefaultConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectDefaultConfigList contains a list of ProjectDefaultConfigs
type ProjectDefaultConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectDefaultConfig `json:"items"`
}

// Repository type metadata.
var (
	ProjectDefaultConfig_Kind             = "ProjectDefaultConfig"
	ProjectDefaultConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectDefaultConfig_Kind}.String()
	ProjectDefaultConfig_KindAPIVersion   = ProjectDefaultConfig_Kind + "." + CRDGroupVersion.String()
	ProjectDefaultConfig_GroupVersionKind = CRDGroupVersion.WithKind(ProjectDefaultConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectDefaultConfig{}, &ProjectDefaultConfigList{})
}