// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package vertexai

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var vertexAIPipelineJobGoogleProvidedLabels = []string{
	"vertex-ai-pipelines-run-billing-id",
}

func VertexAIPipelineJobLabelDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	// Suppress diffs for the labels provided by Google
	for _, label := range vertexAIPipelineJobGoogleProvidedLabels {
		if strings.Contains(k, label) && new == "" {
			return true
		}
	}

	// Let diff be determined by labels (above)
	if strings.Contains(k, "labels.%") {
		return true
	}

	// For other keys, don't suppress diff.
	return false
}

func waitForVertexAIPipelineJobReady(d *schema.ResourceData, config *transport_tpg.Config, timeout time.Duration) error {
	return resource.Retry(timeout, func() *resource.RetryError {
		if err := resourceVertexAIPipelineJobRead(d, config); err != nil {
			return resource.NonRetryableError(err)
		}

		name := d.Get("name").(string)
		state := d.Get("state").(string)
		if state == "PIPELINE_STATE_QUEUED" || state == "PIPELINE_STATE_PENDING" || state == "PIPELINE_STATE_CANCELLING" {
			return resource.RetryableError(fmt.Errorf("VertexAIPipelineJob %q has state %q.", name, state))
		} else if state == "PIPELINE_STATE_SUCCEEDED" || state == "PIPELINE_STATE_RUNNING" {
			log.Printf("[DEBUG] VertexAIPipelineJob %q has state %q.", name, state)
			return nil
		} else {
			return resource.NonRetryableError(fmt.Errorf("VertexAIPipelineJob %q has state %q.", name, state))
		}
	})
}

func waitForVertexAIPipelineJobReadyForDeletion(d *schema.ResourceData, config *transport_tpg.Config, timeout time.Duration) error {
	return resource.Retry(timeout, func() *resource.RetryError {
		if err := resourceVertexAIPipelineJobRead(d, config); err != nil {
			return resource.NonRetryableError(err)
		}

		name := d.Get("name").(string)
		state := d.Get("state").(string)
		if state == "PIPELINE_STATE_RUNNING" || state == "PIPELINE_STATE_PENDING" || state == "PIPELINE_STATE_CANCELLING" {
			return resource.RetryableError(fmt.Errorf("VertexAIPipelineJob %q has state %q.", name, state))
		} else {
			log.Printf("[DEBUG] VertexAIPipelineJob %q has state %q.", name, state)
			return nil
		}
	})
}

const VertexAIPipelineJobAssetType string = "{{region}}-aiplatform.googleapis.com/PipelineJob"

func ResourceConverterVertexAIPipelineJob() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: VertexAIPipelineJobAssetType,
		Convert:   GetVertexAIPipelineJobCaiObject,
	}
}

func GetVertexAIPipelineJobCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//{{region}}-aiplatform.googleapis.com/projects/{{project}}/locations/{{location}}/pipelineJobs/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetVertexAIPipelineJobApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: VertexAIPipelineJobAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/{{region}}-aiplatform/v1beta1/rest",
				DiscoveryName:        "PipelineJob",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetVertexAIPipelineJobApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandVertexAIPipelineJobDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	labelsProp, err := expandVertexAIPipelineJobLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	runtimeConfigProp, err := expandVertexAIPipelineJobRuntimeConfig(d.Get("runtime_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("runtime_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(runtimeConfigProp)) && (ok || !reflect.DeepEqual(v, runtimeConfigProp)) {
		obj["runtimeConfig"] = runtimeConfigProp
	}
	encryptionSpecProp, err := expandVertexAIPipelineJobEncryptionSpec(d.Get("encryption_spec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("encryption_spec"); !tpgresource.IsEmptyValue(reflect.ValueOf(encryptionSpecProp)) && (ok || !reflect.DeepEqual(v, encryptionSpecProp)) {
		obj["encryptionSpec"] = encryptionSpecProp
	}
	serviceAccountProp, err := expandVertexAIPipelineJobServiceAccount(d.Get("service_account"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service_account"); !tpgresource.IsEmptyValue(reflect.ValueOf(serviceAccountProp)) && (ok || !reflect.DeepEqual(v, serviceAccountProp)) {
		obj["serviceAccount"] = serviceAccountProp
	}
	networkProp, err := expandVertexAIPipelineJobNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	templateUriProp, err := expandVertexAIPipelineJobTemplateUri(d.Get("template_uri"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("template_uri"); !tpgresource.IsEmptyValue(reflect.ValueOf(templateUriProp)) && (ok || !reflect.DeepEqual(v, templateUriProp)) {
		obj["templateUri"] = templateUriProp
	}

	return obj, nil
}

func expandVertexAIPipelineJobDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIPipelineJobLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandVertexAIPipelineJobRuntimeConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedGcsOutputDirectory, err := expandVertexAIPipelineJobRuntimeConfigGcsOutputDirectory(original["gcs_output_directory"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGcsOutputDirectory); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["gcsOutputDirectory"] = transformedGcsOutputDirectory
	}

	transformedParameterValues, err := expandVertexAIPipelineJobRuntimeConfigParameterValues(original["parameter_values"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedParameterValues); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["parameterValues"] = transformedParameterValues
	}

	transformedFailurePolicy, err := expandVertexAIPipelineJobRuntimeConfigFailurePolicy(original["failure_policy"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFailurePolicy); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["failurePolicy"] = transformedFailurePolicy
	}

	transformedInputArtifacts, err := expandVertexAIPipelineJobRuntimeConfigInputArtifacts(original["input_artifacts"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInputArtifacts); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["inputArtifacts"] = transformedInputArtifacts
	}

	return transformed, nil
}

func expandVertexAIPipelineJobRuntimeConfigGcsOutputDirectory(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIPipelineJobRuntimeConfigParameterValues(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandVertexAIPipelineJobRuntimeConfigFailurePolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIPipelineJobRuntimeConfigInputArtifacts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedVersion, err := expandVertexAIPipelineJobRuntimeConfigInputArtifactsVersion(original["version"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["version"] = transformedVersion
		}

		transformedArtifact, err := tpgresource.ExpandString(original["artifact"], d, config)
		if err != nil {
			return nil, err
		}
		m[transformedArtifact] = transformed
	}
	return m, nil
}

func expandVertexAIPipelineJobRuntimeConfigInputArtifactsVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIPipelineJobEncryptionSpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeyName, err := expandVertexAIPipelineJobEncryptionSpecKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	return transformed, nil
}

func expandVertexAIPipelineJobEncryptionSpecKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIPipelineJobServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIPipelineJobNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIPipelineJobTemplateUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
