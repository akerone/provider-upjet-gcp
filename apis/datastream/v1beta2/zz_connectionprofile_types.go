// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BigqueryProfileInitParameters struct {
}

type BigqueryProfileObservation struct {
}

type BigqueryProfileParameters struct {
}

type ConnectionProfileInitParameters struct {

	// BigQuery warehouse profile.
	BigqueryProfile *BigqueryProfileInitParameters `json:"bigqueryProfile,omitempty" tf:"bigquery_profile,omitempty"`

	// Display name.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Forward SSH tunnel connectivity.
	// Structure is documented below.
	ForwardSSHConnectivity *ForwardSSHConnectivityInitParameters `json:"forwardSshConnectivity,omitempty" tf:"forward_ssh_connectivity,omitempty"`

	// Cloud Storage bucket profile.
	// Structure is documented below.
	GcsProfile *GcsProfileInitParameters `json:"gcsProfile,omitempty" tf:"gcs_profile,omitempty"`

	// Labels.
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// MySQL database profile.
	// Structure is documented below.
	MySQLProfile *MySQLProfileInitParameters `json:"mysqlProfile,omitempty" tf:"mysql_profile,omitempty"`

	// Oracle database profile.
	// Structure is documented below.
	OracleProfile *OracleProfileInitParameters `json:"oracleProfile,omitempty" tf:"oracle_profile,omitempty"`

	// PostgreSQL database profile.
	// Structure is documented below.
	PostgresqlProfile *PostgresqlProfileInitParameters `json:"postgresqlProfile,omitempty" tf:"postgresql_profile,omitempty"`

	// Private connectivity.
	// Structure is documented below.
	PrivateConnectivity *PrivateConnectivityInitParameters `json:"privateConnectivity,omitempty" tf:"private_connectivity,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type ConnectionProfileObservation struct {

	// BigQuery warehouse profile.
	BigqueryProfile *BigqueryProfileParameters `json:"bigqueryProfile,omitempty" tf:"bigquery_profile,omitempty"`

	// Display name.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// +mapType=granular
	EffectiveLabels map[string]*string `json:"effectiveLabels,omitempty" tf:"effective_labels,omitempty"`

	// Forward SSH tunnel connectivity.
	// Structure is documented below.
	ForwardSSHConnectivity *ForwardSSHConnectivityObservation `json:"forwardSshConnectivity,omitempty" tf:"forward_ssh_connectivity,omitempty"`

	// Cloud Storage bucket profile.
	// Structure is documented below.
	GcsProfile *GcsProfileObservation `json:"gcsProfile,omitempty" tf:"gcs_profile,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/connectionProfiles/{{connection_profile_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Labels.
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of the location this connection profile is located in.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// MySQL database profile.
	// Structure is documented below.
	MySQLProfile *MySQLProfileObservation `json:"mysqlProfile,omitempty" tf:"mysql_profile,omitempty"`

	// The resource's name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Oracle database profile.
	// Structure is documented below.
	OracleProfile *OracleProfileObservation `json:"oracleProfile,omitempty" tf:"oracle_profile,omitempty"`

	// PostgreSQL database profile.
	// Structure is documented below.
	PostgresqlProfile *PostgresqlProfileObservation `json:"postgresqlProfile,omitempty" tf:"postgresql_profile,omitempty"`

	// Private connectivity.
	// Structure is documented below.
	PrivateConnectivity *PrivateConnectivityObservation `json:"privateConnectivity,omitempty" tf:"private_connectivity,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	// +mapType=granular
	TerraformLabels map[string]*string `json:"terraformLabels,omitempty" tf:"terraform_labels,omitempty"`
}

type ConnectionProfileParameters struct {

	// BigQuery warehouse profile.
	// +kubebuilder:validation:Optional
	BigqueryProfile *BigqueryProfileParameters `json:"bigqueryProfile,omitempty" tf:"bigquery_profile,omitempty"`

	// Display name.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Forward SSH tunnel connectivity.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ForwardSSHConnectivity *ForwardSSHConnectivityParameters `json:"forwardSshConnectivity,omitempty" tf:"forward_ssh_connectivity,omitempty"`

	// Cloud Storage bucket profile.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	GcsProfile *GcsProfileParameters `json:"gcsProfile,omitempty" tf:"gcs_profile,omitempty"`

	// Labels.
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of the location this connection profile is located in.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// MySQL database profile.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	MySQLProfile *MySQLProfileParameters `json:"mysqlProfile,omitempty" tf:"mysql_profile,omitempty"`

	// Oracle database profile.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	OracleProfile *OracleProfileParameters `json:"oracleProfile,omitempty" tf:"oracle_profile,omitempty"`

	// PostgreSQL database profile.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	PostgresqlProfile *PostgresqlProfileParameters `json:"postgresqlProfile,omitempty" tf:"postgresql_profile,omitempty"`

	// Private connectivity.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	PrivateConnectivity *PrivateConnectivityParameters `json:"privateConnectivity,omitempty" tf:"private_connectivity,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type ForwardSSHConnectivityInitParameters struct {

	// Hostname for the SSH tunnel.
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// SSH password.
	// Note: This property is sensitive and will not be displayed in the plan.
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// Port for the SSH tunnel.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// SSH private key.
	// Note: This property is sensitive and will not be displayed in the plan.
	PrivateKeySecretRef *v1.SecretKeySelector `json:"privateKeySecretRef,omitempty" tf:"-"`

	// Username for the SSH tunnel.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type ForwardSSHConnectivityObservation struct {

	// Hostname for the SSH tunnel.
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// Port for the SSH tunnel.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Username for the SSH tunnel.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type ForwardSSHConnectivityParameters struct {

	// Hostname for the SSH tunnel.
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname" tf:"hostname,omitempty"`

	// SSH password.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// Port for the SSH tunnel.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// SSH private key.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	PrivateKeySecretRef *v1.SecretKeySelector `json:"privateKeySecretRef,omitempty" tf:"-"`

	// Username for the SSH tunnel.
	// +kubebuilder:validation:Optional
	Username *string `json:"username" tf:"username,omitempty"`
}

type GcsProfileInitParameters struct {

	// The Cloud Storage bucket name.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The root path inside the Cloud Storage bucket.
	RootPath *string `json:"rootPath,omitempty" tf:"root_path,omitempty"`
}

type GcsProfileObservation struct {

	// The Cloud Storage bucket name.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The root path inside the Cloud Storage bucket.
	RootPath *string `json:"rootPath,omitempty" tf:"root_path,omitempty"`
}

type GcsProfileParameters struct {

	// The Cloud Storage bucket name.
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// The root path inside the Cloud Storage bucket.
	// +kubebuilder:validation:Optional
	RootPath *string `json:"rootPath,omitempty" tf:"root_path,omitempty"`
}

type MySQLProfileInitParameters struct {

	// Hostname for the MySQL connection.
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// Password for the MySQL connection.
	// Note: This property is sensitive and will not be displayed in the plan.
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Port for the MySQL connection.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// SSL configuration for the MySQL connection.
	// Structure is documented below.
	SSLConfig *SSLConfigInitParameters `json:"sslConfig,omitempty" tf:"ssl_config,omitempty"`

	// Username for the MySQL connection.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type MySQLProfileObservation struct {

	// Hostname for the MySQL connection.
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// Port for the MySQL connection.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// SSL configuration for the MySQL connection.
	// Structure is documented below.
	SSLConfig *SSLConfigObservation `json:"sslConfig,omitempty" tf:"ssl_config,omitempty"`

	// Username for the MySQL connection.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type MySQLProfileParameters struct {

	// Hostname for the MySQL connection.
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname" tf:"hostname,omitempty"`

	// Password for the MySQL connection.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Port for the MySQL connection.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// SSL configuration for the MySQL connection.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SSLConfig *SSLConfigParameters `json:"sslConfig,omitempty" tf:"ssl_config,omitempty"`

	// Username for the MySQL connection.
	// +kubebuilder:validation:Optional
	Username *string `json:"username" tf:"username,omitempty"`
}

type OracleProfileInitParameters struct {

	// Connection string attributes
	// +mapType=granular
	ConnectionAttributes map[string]*string `json:"connectionAttributes,omitempty" tf:"connection_attributes,omitempty"`

	// Database for the Oracle connection.
	DatabaseService *string `json:"databaseService,omitempty" tf:"database_service,omitempty"`

	// Hostname for the Oracle connection.
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// Password for the Oracle connection.
	// Note: This property is sensitive and will not be displayed in the plan.
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Port for the Oracle connection.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Username for the Oracle connection.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type OracleProfileObservation struct {

	// Connection string attributes
	// +mapType=granular
	ConnectionAttributes map[string]*string `json:"connectionAttributes,omitempty" tf:"connection_attributes,omitempty"`

	// Database for the Oracle connection.
	DatabaseService *string `json:"databaseService,omitempty" tf:"database_service,omitempty"`

	// Hostname for the Oracle connection.
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// Port for the Oracle connection.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Username for the Oracle connection.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type OracleProfileParameters struct {

	// Connection string attributes
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ConnectionAttributes map[string]*string `json:"connectionAttributes,omitempty" tf:"connection_attributes,omitempty"`

	// Database for the Oracle connection.
	// +kubebuilder:validation:Optional
	DatabaseService *string `json:"databaseService" tf:"database_service,omitempty"`

	// Hostname for the Oracle connection.
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname" tf:"hostname,omitempty"`

	// Password for the Oracle connection.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Port for the Oracle connection.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Username for the Oracle connection.
	// +kubebuilder:validation:Optional
	Username *string `json:"username" tf:"username,omitempty"`
}

type PostgresqlProfileInitParameters struct {

	// Database for the PostgreSQL connection.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/sql/v1beta1.Database
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// Reference to a Database in sql to populate database.
	// +kubebuilder:validation:Optional
	DatabaseRef *v1.Reference `json:"databaseRef,omitempty" tf:"-"`

	// Selector for a Database in sql to populate database.
	// +kubebuilder:validation:Optional
	DatabaseSelector *v1.Selector `json:"databaseSelector,omitempty" tf:"-"`

	// Hostname for the PostgreSQL connection.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/sql/v1beta2.DatabaseInstance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("public_ip_address",true)
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// Reference to a DatabaseInstance in sql to populate hostname.
	// +kubebuilder:validation:Optional
	HostnameRef *v1.Reference `json:"hostnameRef,omitempty" tf:"-"`

	// Selector for a DatabaseInstance in sql to populate hostname.
	// +kubebuilder:validation:Optional
	HostnameSelector *v1.Selector `json:"hostnameSelector,omitempty" tf:"-"`

	// Password for the PostgreSQL connection.
	// Note: This property is sensitive and will not be displayed in the plan.
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Port for the PostgreSQL connection.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Username for the PostgreSQL connection.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/sql/v1beta2.User
	Username *string `json:"username,omitempty" tf:"username,omitempty"`

	// Reference to a User in sql to populate username.
	// +kubebuilder:validation:Optional
	UsernameRef *v1.Reference `json:"usernameRef,omitempty" tf:"-"`

	// Selector for a User in sql to populate username.
	// +kubebuilder:validation:Optional
	UsernameSelector *v1.Selector `json:"usernameSelector,omitempty" tf:"-"`
}

type PostgresqlProfileObservation struct {

	// Database for the PostgreSQL connection.
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// Hostname for the PostgreSQL connection.
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// Port for the PostgreSQL connection.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Username for the PostgreSQL connection.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type PostgresqlProfileParameters struct {

	// Database for the PostgreSQL connection.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/sql/v1beta1.Database
	// +kubebuilder:validation:Optional
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// Reference to a Database in sql to populate database.
	// +kubebuilder:validation:Optional
	DatabaseRef *v1.Reference `json:"databaseRef,omitempty" tf:"-"`

	// Selector for a Database in sql to populate database.
	// +kubebuilder:validation:Optional
	DatabaseSelector *v1.Selector `json:"databaseSelector,omitempty" tf:"-"`

	// Hostname for the PostgreSQL connection.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/sql/v1beta2.DatabaseInstance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("public_ip_address",true)
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// Reference to a DatabaseInstance in sql to populate hostname.
	// +kubebuilder:validation:Optional
	HostnameRef *v1.Reference `json:"hostnameRef,omitempty" tf:"-"`

	// Selector for a DatabaseInstance in sql to populate hostname.
	// +kubebuilder:validation:Optional
	HostnameSelector *v1.Selector `json:"hostnameSelector,omitempty" tf:"-"`

	// Password for the PostgreSQL connection.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Port for the PostgreSQL connection.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Username for the PostgreSQL connection.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/sql/v1beta2.User
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`

	// Reference to a User in sql to populate username.
	// +kubebuilder:validation:Optional
	UsernameRef *v1.Reference `json:"usernameRef,omitempty" tf:"-"`

	// Selector for a User in sql to populate username.
	// +kubebuilder:validation:Optional
	UsernameSelector *v1.Selector `json:"usernameSelector,omitempty" tf:"-"`
}

type PrivateConnectivityInitParameters struct {

	// A reference to a private connection resource. Format: projects/{project}/locations/{location}/privateConnections/{name}
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/datastream/v1beta2.PrivateConnection
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	PrivateConnection *string `json:"privateConnection,omitempty" tf:"private_connection,omitempty"`

	// Reference to a PrivateConnection in datastream to populate privateConnection.
	// +kubebuilder:validation:Optional
	PrivateConnectionRef *v1.Reference `json:"privateConnectionRef,omitempty" tf:"-"`

	// Selector for a PrivateConnection in datastream to populate privateConnection.
	// +kubebuilder:validation:Optional
	PrivateConnectionSelector *v1.Selector `json:"privateConnectionSelector,omitempty" tf:"-"`
}

type PrivateConnectivityObservation struct {

	// A reference to a private connection resource. Format: projects/{project}/locations/{location}/privateConnections/{name}
	PrivateConnection *string `json:"privateConnection,omitempty" tf:"private_connection,omitempty"`
}

type PrivateConnectivityParameters struct {

	// A reference to a private connection resource. Format: projects/{project}/locations/{location}/privateConnections/{name}
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/datastream/v1beta2.PrivateConnection
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PrivateConnection *string `json:"privateConnection,omitempty" tf:"private_connection,omitempty"`

	// Reference to a PrivateConnection in datastream to populate privateConnection.
	// +kubebuilder:validation:Optional
	PrivateConnectionRef *v1.Reference `json:"privateConnectionRef,omitempty" tf:"-"`

	// Selector for a PrivateConnection in datastream to populate privateConnection.
	// +kubebuilder:validation:Optional
	PrivateConnectionSelector *v1.Selector `json:"privateConnectionSelector,omitempty" tf:"-"`
}

type SSLConfigInitParameters struct {

	// PEM-encoded certificate of the CA that signed the source database
	// server's certificate.
	// Note: This property is sensitive and will not be displayed in the plan.
	CACertificateSecretRef *v1.SecretKeySelector `json:"caCertificateSecretRef,omitempty" tf:"-"`

	// PEM-encoded certificate that will be used by the replica to
	// authenticate against the source database server. If this field
	// is used then the 'clientKey' and the 'caCertificate' fields are
	// mandatory.
	// Note: This property is sensitive and will not be displayed in the plan.
	ClientCertificateSecretRef *v1.SecretKeySelector `json:"clientCertificateSecretRef,omitempty" tf:"-"`

	// PEM-encoded private key associated with the Client Certificate.
	// If this field is used then the 'client_certificate' and the
	// 'ca_certificate' fields are mandatory.
	// Note: This property is sensitive and will not be displayed in the plan.
	ClientKeySecretRef *v1.SecretKeySelector `json:"clientKeySecretRef,omitempty" tf:"-"`
}

type SSLConfigObservation struct {

	// (Output)
	// Indicates whether the clientKey field is set.
	CACertificateSet *bool `json:"caCertificateSet,omitempty" tf:"ca_certificate_set,omitempty"`

	// (Output)
	// Indicates whether the clientCertificate field is set.
	ClientCertificateSet *bool `json:"clientCertificateSet,omitempty" tf:"client_certificate_set,omitempty"`

	// (Output)
	// Indicates whether the clientKey field is set.
	ClientKeySet *bool `json:"clientKeySet,omitempty" tf:"client_key_set,omitempty"`
}

type SSLConfigParameters struct {

	// PEM-encoded certificate of the CA that signed the source database
	// server's certificate.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	CACertificateSecretRef *v1.SecretKeySelector `json:"caCertificateSecretRef,omitempty" tf:"-"`

	// PEM-encoded certificate that will be used by the replica to
	// authenticate against the source database server. If this field
	// is used then the 'clientKey' and the 'caCertificate' fields are
	// mandatory.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	ClientCertificateSecretRef *v1.SecretKeySelector `json:"clientCertificateSecretRef,omitempty" tf:"-"`

	// PEM-encoded private key associated with the Client Certificate.
	// If this field is used then the 'client_certificate' and the
	// 'ca_certificate' fields are mandatory.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	ClientKeySecretRef *v1.SecretKeySelector `json:"clientKeySecretRef,omitempty" tf:"-"`
}

// ConnectionProfileSpec defines the desired state of ConnectionProfile
type ConnectionProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectionProfileParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ConnectionProfileInitParameters `json:"initProvider,omitempty"`
}

// ConnectionProfileStatus defines the observed state of ConnectionProfile.
type ConnectionProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectionProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ConnectionProfile is the Schema for the ConnectionProfiles API. A set of reusable connection configurations to be used as a source or destination for a stream.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ConnectionProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	Spec   ConnectionProfileSpec   `json:"spec"`
	Status ConnectionProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionProfileList contains a list of ConnectionProfiles
type ConnectionProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConnectionProfile `json:"items"`
}

// Repository type metadata.
var (
	ConnectionProfile_Kind             = "ConnectionProfile"
	ConnectionProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConnectionProfile_Kind}.String()
	ConnectionProfile_KindAPIVersion   = ConnectionProfile_Kind + "." + CRDGroupVersion.String()
	ConnectionProfile_GroupVersionKind = CRDGroupVersion.WithKind(ConnectionProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&ConnectionProfile{}, &ConnectionProfileList{})
}
