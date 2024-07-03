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

type ArgumentsInitParameters struct {

	// Defaults to FIXED_TYPE.
	// Default value is FIXED_TYPE.
	// Possible values are: FIXED_TYPE, ANY_TYPE.
	ArgumentKind *string `json:"argumentKind,omitempty" tf:"argument_kind,omitempty"`

	// A JSON schema for the data type. Required unless argumentKind = ANY_TYPE.
	// ~>NOTE: Because this field expects a JSON string, any changes to the string
	// will create a diff, even if the JSON itself hasn't changed. If the API returns
	// a different value for the same schema, e.g. it switched the order of values
	// or replaced STRUCT field type with RECORD field type, we currently cannot
	// suppress the recurring diff this causes. As a workaround, we recommend using
	// the schema as returned by the API.
	DataType *string `json:"dataType,omitempty" tf:"data_type,omitempty"`

	// Specifies whether the argument is input or output. Can be set for procedures only.
	// Possible values are: IN, OUT, INOUT.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of this argument. Can be absent for function return argument.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ArgumentsObservation struct {

	// Defaults to FIXED_TYPE.
	// Default value is FIXED_TYPE.
	// Possible values are: FIXED_TYPE, ANY_TYPE.
	ArgumentKind *string `json:"argumentKind,omitempty" tf:"argument_kind,omitempty"`

	// A JSON schema for the data type. Required unless argumentKind = ANY_TYPE.
	// ~>NOTE: Because this field expects a JSON string, any changes to the string
	// will create a diff, even if the JSON itself hasn't changed. If the API returns
	// a different value for the same schema, e.g. it switched the order of values
	// or replaced STRUCT field type with RECORD field type, we currently cannot
	// suppress the recurring diff this causes. As a workaround, we recommend using
	// the schema as returned by the API.
	DataType *string `json:"dataType,omitempty" tf:"data_type,omitempty"`

	// Specifies whether the argument is input or output. Can be set for procedures only.
	// Possible values are: IN, OUT, INOUT.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of this argument. Can be absent for function return argument.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ArgumentsParameters struct {

	// Defaults to FIXED_TYPE.
	// Default value is FIXED_TYPE.
	// Possible values are: FIXED_TYPE, ANY_TYPE.
	// +kubebuilder:validation:Optional
	ArgumentKind *string `json:"argumentKind,omitempty" tf:"argument_kind,omitempty"`

	// A JSON schema for the data type. Required unless argumentKind = ANY_TYPE.
	// ~>NOTE: Because this field expects a JSON string, any changes to the string
	// will create a diff, even if the JSON itself hasn't changed. If the API returns
	// a different value for the same schema, e.g. it switched the order of values
	// or replaced STRUCT field type with RECORD field type, we currently cannot
	// suppress the recurring diff this causes. As a workaround, we recommend using
	// the schema as returned by the API.
	// +kubebuilder:validation:Optional
	DataType *string `json:"dataType,omitempty" tf:"data_type,omitempty"`

	// Specifies whether the argument is input or output. Can be set for procedures only.
	// Possible values are: IN, OUT, INOUT.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of this argument. Can be absent for function return argument.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type RemoteFunctionOptionsInitParameters struct {

	// Fully qualified name of the user-provided connection object which holds
	// the authentication information to send requests to the remote service.
	// Format: "projects/{projectId}/locations/{locationId}/connections/{connectionId}"
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta2.Connection
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)
	Connection *string `json:"connection,omitempty" tf:"connection,omitempty"`

	// Reference to a Connection in bigquery to populate connection.
	// +kubebuilder:validation:Optional
	ConnectionRef *v1.Reference `json:"connectionRef,omitempty" tf:"-"`

	// Selector for a Connection in bigquery to populate connection.
	// +kubebuilder:validation:Optional
	ConnectionSelector *v1.Selector `json:"connectionSelector,omitempty" tf:"-"`

	// Endpoint of the user-provided remote service, e.g.
	// https://us-east1-my_gcf_project.cloudfunctions.net/remote_add
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// Max number of rows in each batch sent to the remote service. If absent or if 0,
	// BigQuery dynamically decides the number of rows in a batch.
	MaxBatchingRows *string `json:"maxBatchingRows,omitempty" tf:"max_batching_rows,omitempty"`

	// User-defined context as a set of key/value pairs, which will be sent as function
	// invocation context together with batched arguments in the requests to the remote
	// service. The total number of bytes of keys and values must be less than 8KB.
	// An object containing a list of "key": value pairs. Example:
	// { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	// +mapType=granular
	UserDefinedContext map[string]*string `json:"userDefinedContext,omitempty" tf:"user_defined_context,omitempty"`
}

type RemoteFunctionOptionsObservation struct {

	// Fully qualified name of the user-provided connection object which holds
	// the authentication information to send requests to the remote service.
	// Format: "projects/{projectId}/locations/{locationId}/connections/{connectionId}"
	Connection *string `json:"connection,omitempty" tf:"connection,omitempty"`

	// Endpoint of the user-provided remote service, e.g.
	// https://us-east1-my_gcf_project.cloudfunctions.net/remote_add
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// Max number of rows in each batch sent to the remote service. If absent or if 0,
	// BigQuery dynamically decides the number of rows in a batch.
	MaxBatchingRows *string `json:"maxBatchingRows,omitempty" tf:"max_batching_rows,omitempty"`

	// User-defined context as a set of key/value pairs, which will be sent as function
	// invocation context together with batched arguments in the requests to the remote
	// service. The total number of bytes of keys and values must be less than 8KB.
	// An object containing a list of "key": value pairs. Example:
	// { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	// +mapType=granular
	UserDefinedContext map[string]*string `json:"userDefinedContext,omitempty" tf:"user_defined_context,omitempty"`
}

type RemoteFunctionOptionsParameters struct {

	// Fully qualified name of the user-provided connection object which holds
	// the authentication information to send requests to the remote service.
	// Format: "projects/{projectId}/locations/{locationId}/connections/{connectionId}"
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta2.Connection
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)
	// +kubebuilder:validation:Optional
	Connection *string `json:"connection,omitempty" tf:"connection,omitempty"`

	// Reference to a Connection in bigquery to populate connection.
	// +kubebuilder:validation:Optional
	ConnectionRef *v1.Reference `json:"connectionRef,omitempty" tf:"-"`

	// Selector for a Connection in bigquery to populate connection.
	// +kubebuilder:validation:Optional
	ConnectionSelector *v1.Selector `json:"connectionSelector,omitempty" tf:"-"`

	// Endpoint of the user-provided remote service, e.g.
	// https://us-east1-my_gcf_project.cloudfunctions.net/remote_add
	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// Max number of rows in each batch sent to the remote service. If absent or if 0,
	// BigQuery dynamically decides the number of rows in a batch.
	// +kubebuilder:validation:Optional
	MaxBatchingRows *string `json:"maxBatchingRows,omitempty" tf:"max_batching_rows,omitempty"`

	// User-defined context as a set of key/value pairs, which will be sent as function
	// invocation context together with batched arguments in the requests to the remote
	// service. The total number of bytes of keys and values must be less than 8KB.
	// An object containing a list of "key": value pairs. Example:
	// { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	UserDefinedContext map[string]*string `json:"userDefinedContext,omitempty" tf:"user_defined_context,omitempty"`
}

type RoutineInitParameters_2 struct {

	// Input/output argument of a function or a stored procedure.
	// Structure is documented below.
	Arguments []ArgumentsInitParameters `json:"arguments,omitempty" tf:"arguments,omitempty"`

	// If set to DATA_MASKING, the function is validated and made available as a masking function. For more information, see https://cloud.google.com/bigquery/docs/user-defined-functions#custom-mask
	// Possible values are: DATA_MASKING.
	DataGovernanceType *string `json:"dataGovernanceType,omitempty" tf:"data_governance_type,omitempty"`

	// The body of the routine. For functions, this is the expression in the AS clause.
	// If language=SQL, it is the substring inside (but excluding) the parentheses.
	DefinitionBody *string `json:"definitionBody,omitempty" tf:"definition_body,omitempty"`

	// The description of the routine if defined.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The determinism level of the JavaScript UDF if defined.
	// Possible values are: DETERMINISM_LEVEL_UNSPECIFIED, DETERMINISTIC, NOT_DETERMINISTIC.
	DeterminismLevel *string `json:"determinismLevel,omitempty" tf:"determinism_level,omitempty"`

	// Optional. If language = "JAVASCRIPT", this field stores the path of the
	// imported JAVASCRIPT libraries.
	ImportedLibraries []*string `json:"importedLibraries,omitempty" tf:"imported_libraries,omitempty"`

	// The language of the routine.
	// Possible values are: SQL, JAVASCRIPT, PYTHON, JAVA, SCALA.
	Language *string `json:"language,omitempty" tf:"language,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Remote function specific options.
	// Structure is documented below.
	RemoteFunctionOptions *RemoteFunctionOptionsInitParameters `json:"remoteFunctionOptions,omitempty" tf:"remote_function_options,omitempty"`

	// Optional. Can be set only if routineType = "TABLE_VALUED_FUNCTION".
	// If absent, the return table type is inferred from definitionBody at query time in each query
	// that references this routine. If present, then the columns in the evaluated table result will
	// be cast to match the column types specificed in return table type, at query time.
	ReturnTableType *string `json:"returnTableType,omitempty" tf:"return_table_type,omitempty"`

	// A JSON schema for the return type. Optional if language = "SQL"; required otherwise.
	// If absent, the return type is inferred from definitionBody at query time in each query
	// that references this routine. If present, then the evaluated result will be cast to
	// the specified returned type at query time. ~>NOTE: Because this field expects a JSON
	// string, any changes to the string will create a diff, even if the JSON itself hasn't
	// changed. If the API returns a different value for the same schema, e.g. it switche
	// d the order of values or replaced STRUCT field type with RECORD field type, we currently
	// cannot suppress the recurring diff this causes. As a workaround, we recommend using
	// the schema as returned by the API.
	ReturnType *string `json:"returnType,omitempty" tf:"return_type,omitempty"`

	// The type of routine.
	// Possible values are: SCALAR_FUNCTION, PROCEDURE, TABLE_VALUED_FUNCTION.
	RoutineType *string `json:"routineType,omitempty" tf:"routine_type,omitempty"`

	// Optional. If language is one of "PYTHON", "JAVA", "SCALA", this field stores the options for spark stored procedure.
	// Structure is documented below.
	SparkOptions *SparkOptionsInitParameters `json:"sparkOptions,omitempty" tf:"spark_options,omitempty"`
}

type RoutineObservation_2 struct {

	// Input/output argument of a function or a stored procedure.
	// Structure is documented below.
	Arguments []ArgumentsObservation `json:"arguments,omitempty" tf:"arguments,omitempty"`

	// The time when this routine was created, in milliseconds since the
	// epoch.
	CreationTime *float64 `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	// If set to DATA_MASKING, the function is validated and made available as a masking function. For more information, see https://cloud.google.com/bigquery/docs/user-defined-functions#custom-mask
	// Possible values are: DATA_MASKING.
	DataGovernanceType *string `json:"dataGovernanceType,omitempty" tf:"data_governance_type,omitempty"`

	// The ID of the dataset containing this routine
	DatasetID *string `json:"datasetId,omitempty" tf:"dataset_id,omitempty"`

	// The body of the routine. For functions, this is the expression in the AS clause.
	// If language=SQL, it is the substring inside (but excluding) the parentheses.
	DefinitionBody *string `json:"definitionBody,omitempty" tf:"definition_body,omitempty"`

	// The description of the routine if defined.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The determinism level of the JavaScript UDF if defined.
	// Possible values are: DETERMINISM_LEVEL_UNSPECIFIED, DETERMINISTIC, NOT_DETERMINISTIC.
	DeterminismLevel *string `json:"determinismLevel,omitempty" tf:"determinism_level,omitempty"`

	// an identifier for the resource with format projects/{{project}}/datasets/{{dataset_id}}/routines/{{routine_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Optional. If language = "JAVASCRIPT", this field stores the path of the
	// imported JAVASCRIPT libraries.
	ImportedLibraries []*string `json:"importedLibraries,omitempty" tf:"imported_libraries,omitempty"`

	// The language of the routine.
	// Possible values are: SQL, JAVASCRIPT, PYTHON, JAVA, SCALA.
	Language *string `json:"language,omitempty" tf:"language,omitempty"`

	// The time when this routine was modified, in milliseconds since the
	// epoch.
	LastModifiedTime *float64 `json:"lastModifiedTime,omitempty" tf:"last_modified_time,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Remote function specific options.
	// Structure is documented below.
	RemoteFunctionOptions *RemoteFunctionOptionsObservation `json:"remoteFunctionOptions,omitempty" tf:"remote_function_options,omitempty"`

	// Optional. Can be set only if routineType = "TABLE_VALUED_FUNCTION".
	// If absent, the return table type is inferred from definitionBody at query time in each query
	// that references this routine. If present, then the columns in the evaluated table result will
	// be cast to match the column types specificed in return table type, at query time.
	ReturnTableType *string `json:"returnTableType,omitempty" tf:"return_table_type,omitempty"`

	// A JSON schema for the return type. Optional if language = "SQL"; required otherwise.
	// If absent, the return type is inferred from definitionBody at query time in each query
	// that references this routine. If present, then the evaluated result will be cast to
	// the specified returned type at query time. ~>NOTE: Because this field expects a JSON
	// string, any changes to the string will create a diff, even if the JSON itself hasn't
	// changed. If the API returns a different value for the same schema, e.g. it switche
	// d the order of values or replaced STRUCT field type with RECORD field type, we currently
	// cannot suppress the recurring diff this causes. As a workaround, we recommend using
	// the schema as returned by the API.
	ReturnType *string `json:"returnType,omitempty" tf:"return_type,omitempty"`

	// The type of routine.
	// Possible values are: SCALAR_FUNCTION, PROCEDURE, TABLE_VALUED_FUNCTION.
	RoutineType *string `json:"routineType,omitempty" tf:"routine_type,omitempty"`

	// Optional. If language is one of "PYTHON", "JAVA", "SCALA", this field stores the options for spark stored procedure.
	// Structure is documented below.
	SparkOptions *SparkOptionsObservation `json:"sparkOptions,omitempty" tf:"spark_options,omitempty"`
}

type RoutineParameters_2 struct {

	// Input/output argument of a function or a stored procedure.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Arguments []ArgumentsParameters `json:"arguments,omitempty" tf:"arguments,omitempty"`

	// If set to DATA_MASKING, the function is validated and made available as a masking function. For more information, see https://cloud.google.com/bigquery/docs/user-defined-functions#custom-mask
	// Possible values are: DATA_MASKING.
	// +kubebuilder:validation:Optional
	DataGovernanceType *string `json:"dataGovernanceType,omitempty" tf:"data_governance_type,omitempty"`

	// The ID of the dataset containing this routine
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta2.Dataset
	// +kubebuilder:validation:Optional
	DatasetID *string `json:"datasetId,omitempty" tf:"dataset_id,omitempty"`

	// Reference to a Dataset in bigquery to populate datasetId.
	// +kubebuilder:validation:Optional
	DatasetIDRef *v1.Reference `json:"datasetIdRef,omitempty" tf:"-"`

	// Selector for a Dataset in bigquery to populate datasetId.
	// +kubebuilder:validation:Optional
	DatasetIDSelector *v1.Selector `json:"datasetIdSelector,omitempty" tf:"-"`

	// The body of the routine. For functions, this is the expression in the AS clause.
	// If language=SQL, it is the substring inside (but excluding) the parentheses.
	// +kubebuilder:validation:Optional
	DefinitionBody *string `json:"definitionBody,omitempty" tf:"definition_body,omitempty"`

	// The description of the routine if defined.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The determinism level of the JavaScript UDF if defined.
	// Possible values are: DETERMINISM_LEVEL_UNSPECIFIED, DETERMINISTIC, NOT_DETERMINISTIC.
	// +kubebuilder:validation:Optional
	DeterminismLevel *string `json:"determinismLevel,omitempty" tf:"determinism_level,omitempty"`

	// Optional. If language = "JAVASCRIPT", this field stores the path of the
	// imported JAVASCRIPT libraries.
	// +kubebuilder:validation:Optional
	ImportedLibraries []*string `json:"importedLibraries,omitempty" tf:"imported_libraries,omitempty"`

	// The language of the routine.
	// Possible values are: SQL, JAVASCRIPT, PYTHON, JAVA, SCALA.
	// +kubebuilder:validation:Optional
	Language *string `json:"language,omitempty" tf:"language,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Remote function specific options.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	RemoteFunctionOptions *RemoteFunctionOptionsParameters `json:"remoteFunctionOptions,omitempty" tf:"remote_function_options,omitempty"`

	// Optional. Can be set only if routineType = "TABLE_VALUED_FUNCTION".
	// If absent, the return table type is inferred from definitionBody at query time in each query
	// that references this routine. If present, then the columns in the evaluated table result will
	// be cast to match the column types specificed in return table type, at query time.
	// +kubebuilder:validation:Optional
	ReturnTableType *string `json:"returnTableType,omitempty" tf:"return_table_type,omitempty"`

	// A JSON schema for the return type. Optional if language = "SQL"; required otherwise.
	// If absent, the return type is inferred from definitionBody at query time in each query
	// that references this routine. If present, then the evaluated result will be cast to
	// the specified returned type at query time. ~>NOTE: Because this field expects a JSON
	// string, any changes to the string will create a diff, even if the JSON itself hasn't
	// changed. If the API returns a different value for the same schema, e.g. it switche
	// d the order of values or replaced STRUCT field type with RECORD field type, we currently
	// cannot suppress the recurring diff this causes. As a workaround, we recommend using
	// the schema as returned by the API.
	// +kubebuilder:validation:Optional
	ReturnType *string `json:"returnType,omitempty" tf:"return_type,omitempty"`

	// The type of routine.
	// Possible values are: SCALAR_FUNCTION, PROCEDURE, TABLE_VALUED_FUNCTION.
	// +kubebuilder:validation:Optional
	RoutineType *string `json:"routineType,omitempty" tf:"routine_type,omitempty"`

	// Optional. If language is one of "PYTHON", "JAVA", "SCALA", this field stores the options for spark stored procedure.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SparkOptions *SparkOptionsParameters `json:"sparkOptions,omitempty" tf:"spark_options,omitempty"`
}

type SparkOptionsInitParameters struct {

	// Archive files to be extracted into the working directory of each executor. For more information about Apache Spark, see Apache Spark.
	ArchiveUris []*string `json:"archiveUris,omitempty" tf:"archive_uris,omitempty"`

	// Fully qualified name of the user-provided Spark connection object.
	// Format: "projects/{projectId}/locations/{locationId}/connections/{connectionId}"
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta2.Connection
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)
	Connection *string `json:"connection,omitempty" tf:"connection,omitempty"`

	// Reference to a Connection in bigquery to populate connection.
	// +kubebuilder:validation:Optional
	ConnectionRef *v1.Reference `json:"connectionRef,omitempty" tf:"-"`

	// Selector for a Connection in bigquery to populate connection.
	// +kubebuilder:validation:Optional
	ConnectionSelector *v1.Selector `json:"connectionSelector,omitempty" tf:"-"`

	// Custom container image for the runtime environment.
	ContainerImage *string `json:"containerImage,omitempty" tf:"container_image,omitempty"`

	// Files to be placed in the working directory of each executor. For more information about Apache Spark, see Apache Spark.
	FileUris []*string `json:"fileUris,omitempty" tf:"file_uris,omitempty"`

	// JARs to include on the driver and executor CLASSPATH. For more information about Apache Spark, see Apache Spark.
	JarUris []*string `json:"jarUris,omitempty" tf:"jar_uris,omitempty"`

	// The fully qualified name of a class in jarUris, for example, com.example.wordcount.
	// Exactly one of mainClass and main_jar_uri field should be set for Java/Scala language type.
	MainClass *string `json:"mainClass,omitempty" tf:"main_class,omitempty"`

	// The main file/jar URI of the Spark application.
	// Exactly one of the definitionBody field and the mainFileUri field must be set for Python.
	// Exactly one of mainClass and mainFileUri field should be set for Java/Scala language type.
	MainFileURI *string `json:"mainFileUri,omitempty" tf:"main_file_uri,omitempty"`

	// Configuration properties as a set of key/value pairs, which will be passed on to the Spark application.
	// For more information, see Apache Spark and the procedure option list.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	// +mapType=granular
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// Python files to be placed on the PYTHONPATH for PySpark application. Supported file types: .py, .egg, and .zip. For more information about Apache Spark, see Apache Spark.
	PyFileUris []*string `json:"pyFileUris,omitempty" tf:"py_file_uris,omitempty"`

	// Runtime version. If not specified, the default runtime version is used.
	RuntimeVersion *string `json:"runtimeVersion,omitempty" tf:"runtime_version,omitempty"`
}

type SparkOptionsObservation struct {

	// Archive files to be extracted into the working directory of each executor. For more information about Apache Spark, see Apache Spark.
	ArchiveUris []*string `json:"archiveUris,omitempty" tf:"archive_uris,omitempty"`

	// Fully qualified name of the user-provided Spark connection object.
	// Format: "projects/{projectId}/locations/{locationId}/connections/{connectionId}"
	Connection *string `json:"connection,omitempty" tf:"connection,omitempty"`

	// Custom container image for the runtime environment.
	ContainerImage *string `json:"containerImage,omitempty" tf:"container_image,omitempty"`

	// Files to be placed in the working directory of each executor. For more information about Apache Spark, see Apache Spark.
	FileUris []*string `json:"fileUris,omitempty" tf:"file_uris,omitempty"`

	// JARs to include on the driver and executor CLASSPATH. For more information about Apache Spark, see Apache Spark.
	JarUris []*string `json:"jarUris,omitempty" tf:"jar_uris,omitempty"`

	// The fully qualified name of a class in jarUris, for example, com.example.wordcount.
	// Exactly one of mainClass and main_jar_uri field should be set for Java/Scala language type.
	MainClass *string `json:"mainClass,omitempty" tf:"main_class,omitempty"`

	// The main file/jar URI of the Spark application.
	// Exactly one of the definitionBody field and the mainFileUri field must be set for Python.
	// Exactly one of mainClass and mainFileUri field should be set for Java/Scala language type.
	MainFileURI *string `json:"mainFileUri,omitempty" tf:"main_file_uri,omitempty"`

	// Configuration properties as a set of key/value pairs, which will be passed on to the Spark application.
	// For more information, see Apache Spark and the procedure option list.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	// +mapType=granular
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// Python files to be placed on the PYTHONPATH for PySpark application. Supported file types: .py, .egg, and .zip. For more information about Apache Spark, see Apache Spark.
	PyFileUris []*string `json:"pyFileUris,omitempty" tf:"py_file_uris,omitempty"`

	// Runtime version. If not specified, the default runtime version is used.
	RuntimeVersion *string `json:"runtimeVersion,omitempty" tf:"runtime_version,omitempty"`
}

type SparkOptionsParameters struct {

	// Archive files to be extracted into the working directory of each executor. For more information about Apache Spark, see Apache Spark.
	// +kubebuilder:validation:Optional
	ArchiveUris []*string `json:"archiveUris,omitempty" tf:"archive_uris,omitempty"`

	// Fully qualified name of the user-provided Spark connection object.
	// Format: "projects/{projectId}/locations/{locationId}/connections/{connectionId}"
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta2.Connection
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)
	// +kubebuilder:validation:Optional
	Connection *string `json:"connection,omitempty" tf:"connection,omitempty"`

	// Reference to a Connection in bigquery to populate connection.
	// +kubebuilder:validation:Optional
	ConnectionRef *v1.Reference `json:"connectionRef,omitempty" tf:"-"`

	// Selector for a Connection in bigquery to populate connection.
	// +kubebuilder:validation:Optional
	ConnectionSelector *v1.Selector `json:"connectionSelector,omitempty" tf:"-"`

	// Custom container image for the runtime environment.
	// +kubebuilder:validation:Optional
	ContainerImage *string `json:"containerImage,omitempty" tf:"container_image,omitempty"`

	// Files to be placed in the working directory of each executor. For more information about Apache Spark, see Apache Spark.
	// +kubebuilder:validation:Optional
	FileUris []*string `json:"fileUris,omitempty" tf:"file_uris,omitempty"`

	// JARs to include on the driver and executor CLASSPATH. For more information about Apache Spark, see Apache Spark.
	// +kubebuilder:validation:Optional
	JarUris []*string `json:"jarUris,omitempty" tf:"jar_uris,omitempty"`

	// The fully qualified name of a class in jarUris, for example, com.example.wordcount.
	// Exactly one of mainClass and main_jar_uri field should be set for Java/Scala language type.
	// +kubebuilder:validation:Optional
	MainClass *string `json:"mainClass,omitempty" tf:"main_class,omitempty"`

	// The main file/jar URI of the Spark application.
	// Exactly one of the definitionBody field and the mainFileUri field must be set for Python.
	// Exactly one of mainClass and mainFileUri field should be set for Java/Scala language type.
	// +kubebuilder:validation:Optional
	MainFileURI *string `json:"mainFileUri,omitempty" tf:"main_file_uri,omitempty"`

	// Configuration properties as a set of key/value pairs, which will be passed on to the Spark application.
	// For more information, see Apache Spark and the procedure option list.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// Python files to be placed on the PYTHONPATH for PySpark application. Supported file types: .py, .egg, and .zip. For more information about Apache Spark, see Apache Spark.
	// +kubebuilder:validation:Optional
	PyFileUris []*string `json:"pyFileUris,omitempty" tf:"py_file_uris,omitempty"`

	// Runtime version. If not specified, the default runtime version is used.
	// +kubebuilder:validation:Optional
	RuntimeVersion *string `json:"runtimeVersion,omitempty" tf:"runtime_version,omitempty"`
}

// RoutineSpec defines the desired state of Routine
type RoutineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoutineParameters_2 `json:"forProvider"`
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
	InitProvider RoutineInitParameters_2 `json:"initProvider,omitempty"`
}

// RoutineStatus defines the observed state of Routine.
type RoutineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoutineObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Routine is the Schema for the Routines API. A user-defined function or a stored procedure that belongs to a Dataset
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Routine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.definitionBody) || (has(self.initProvider) && has(self.initProvider.definitionBody))",message="spec.forProvider.definitionBody is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.routineType) || (has(self.initProvider) && has(self.initProvider.routineType))",message="spec.forProvider.routineType is a required parameter"
	Spec   RoutineSpec   `json:"spec"`
	Status RoutineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoutineList contains a list of Routines
type RoutineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Routine `json:"items"`
}

// Repository type metadata.
var (
	Routine_Kind             = "Routine"
	Routine_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Routine_Kind}.String()
	Routine_KindAPIVersion   = Routine_Kind + "." + CRDGroupVersion.String()
	Routine_GroupVersionKind = CRDGroupVersion.WithKind(Routine_Kind)
)

func init() {
	SchemeBuilder.Register(&Routine{}, &RoutineList{})
}
