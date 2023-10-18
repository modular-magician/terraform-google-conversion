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

package cloudfunctions2

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const Cloudfunctions2functionAssetType string = "cloudfunctions.googleapis.com/function"

func ResourceConverterCloudfunctions2function() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: Cloudfunctions2functionAssetType,
		Convert:   GetCloudfunctions2functionCaiObject,
	}
}

func GetCloudfunctions2functionCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//cloudfunctions.googleapis.com/projects/{{project}}/locations/{{location}}/functions/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetCloudfunctions2functionApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: Cloudfunctions2functionAssetType,
			Resource: &cai.AssetResource{
				Version:              "v2beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudfunctions/v2beta/rest",
				DiscoveryName:        "function",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetCloudfunctions2functionApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandCloudfunctions2functionName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandCloudfunctions2functionDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	buildConfigProp, err := expandCloudfunctions2functionBuildConfig(d.Get("build_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("build_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(buildConfigProp)) && (ok || !reflect.DeepEqual(v, buildConfigProp)) {
		obj["buildConfig"] = buildConfigProp
	}
	serviceConfigProp, err := expandCloudfunctions2functionServiceConfig(d.Get("service_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(serviceConfigProp)) && (ok || !reflect.DeepEqual(v, serviceConfigProp)) {
		obj["serviceConfig"] = serviceConfigProp
	}
	eventTriggerProp, err := expandCloudfunctions2functionEventTrigger(d.Get("event_trigger"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("event_trigger"); !tpgresource.IsEmptyValue(reflect.ValueOf(eventTriggerProp)) && (ok || !reflect.DeepEqual(v, eventTriggerProp)) {
		obj["eventTrigger"] = eventTriggerProp
	}
	kmsKeyNameProp, err := expandCloudfunctions2functionKmsKeyName(d.Get("kms_key_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("kms_key_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(kmsKeyNameProp)) && (ok || !reflect.DeepEqual(v, kmsKeyNameProp)) {
		obj["kmsKeyName"] = kmsKeyNameProp
	}
	labelsProp, err := expandCloudfunctions2functionEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandCloudfunctions2functionName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/functions/{{name}}")
}

func expandCloudfunctions2functionDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRuntime, err := expandCloudfunctions2functionBuildConfigRuntime(original["runtime"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRuntime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["runtime"] = transformedRuntime
	}

	transformedEntryPoint, err := expandCloudfunctions2functionBuildConfigEntryPoint(original["entry_point"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEntryPoint); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["entryPoint"] = transformedEntryPoint
	}

	transformedSource, err := expandCloudfunctions2functionBuildConfigSource(original["source"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSource); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["source"] = transformedSource
	}

	transformedWorkerPool, err := expandCloudfunctions2functionBuildConfigWorkerPool(original["worker_pool"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWorkerPool); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["workerPool"] = transformedWorkerPool
	}

	transformedEnvironmentVariables, err := expandCloudfunctions2functionBuildConfigEnvironmentVariables(original["environment_variables"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnvironmentVariables); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["environmentVariables"] = transformedEnvironmentVariables
	}

	transformedDockerRepository, err := expandCloudfunctions2functionBuildConfigDockerRepository(original["docker_repository"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDockerRepository); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dockerRepository"] = transformedDockerRepository
	}

	return transformed, nil
}

func expandCloudfunctions2functionBuildConfigRuntime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigEntryPoint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigSource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedStorageSource, err := expandCloudfunctions2functionBuildConfigSourceStorageSource(original["storage_source"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStorageSource); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["storageSource"] = transformedStorageSource
	}

	transformedRepoSource, err := expandCloudfunctions2functionBuildConfigSourceRepoSource(original["repo_source"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRepoSource); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["repoSource"] = transformedRepoSource
	}

	return transformed, nil
}

func expandCloudfunctions2functionBuildConfigSourceStorageSource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBucket, err := expandCloudfunctions2functionBuildConfigSourceStorageSourceBucket(original["bucket"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBucket); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["bucket"] = transformedBucket
	}

	transformedObject, err := expandCloudfunctions2functionBuildConfigSourceStorageSourceObject(original["object"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedObject); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["object"] = transformedObject
	}

	transformedGeneration, err := expandCloudfunctions2functionBuildConfigSourceStorageSourceGeneration(original["generation"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGeneration); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["generation"] = transformedGeneration
	}

	return transformed, nil
}

func expandCloudfunctions2functionBuildConfigSourceStorageSourceBucket(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigSourceStorageSourceObject(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigSourceStorageSourceGeneration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigSourceRepoSource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProjectId, err := expandCloudfunctions2functionBuildConfigSourceRepoSourceProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	transformedRepoName, err := expandCloudfunctions2functionBuildConfigSourceRepoSourceRepoName(original["repo_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRepoName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["repoName"] = transformedRepoName
	}

	transformedBranchName, err := expandCloudfunctions2functionBuildConfigSourceRepoSourceBranchName(original["branch_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBranchName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["branchName"] = transformedBranchName
	}

	transformedTagName, err := expandCloudfunctions2functionBuildConfigSourceRepoSourceTagName(original["tag_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTagName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["tagName"] = transformedTagName
	}

	transformedCommitSha, err := expandCloudfunctions2functionBuildConfigSourceRepoSourceCommitSha(original["commit_sha"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCommitSha); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["commitSha"] = transformedCommitSha
	}

	transformedDir, err := expandCloudfunctions2functionBuildConfigSourceRepoSourceDir(original["dir"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDir); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dir"] = transformedDir
	}

	transformedInvertRegex, err := expandCloudfunctions2functionBuildConfigSourceRepoSourceInvertRegex(original["invert_regex"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInvertRegex); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["invertRegex"] = transformedInvertRegex
	}

	return transformed, nil
}

func expandCloudfunctions2functionBuildConfigSourceRepoSourceProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigSourceRepoSourceRepoName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigSourceRepoSourceBranchName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigSourceRepoSourceTagName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigSourceRepoSourceCommitSha(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigSourceRepoSourceDir(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigSourceRepoSourceInvertRegex(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigWorkerPool(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionBuildConfigEnvironmentVariables(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudfunctions2functionBuildConfigDockerRepository(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedService, err := expandCloudfunctions2functionServiceConfigService(original["service"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedService); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["service"] = transformedService
	}

	transformedTimeoutSeconds, err := expandCloudfunctions2functionServiceConfigTimeoutSeconds(original["timeout_seconds"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTimeoutSeconds); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["timeoutSeconds"] = transformedTimeoutSeconds
	}

	transformedAvailableMemory, err := expandCloudfunctions2functionServiceConfigAvailableMemory(original["available_memory"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAvailableMemory); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["availableMemory"] = transformedAvailableMemory
	}

	transformedMaxInstanceRequestConcurrency, err := expandCloudfunctions2functionServiceConfigMaxInstanceRequestConcurrency(original["max_instance_request_concurrency"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxInstanceRequestConcurrency); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxInstanceRequestConcurrency"] = transformedMaxInstanceRequestConcurrency
	}

	transformedAvailableCpu, err := expandCloudfunctions2functionServiceConfigAvailableCpu(original["available_cpu"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAvailableCpu); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["availableCpu"] = transformedAvailableCpu
	}

	transformedEnvironmentVariables, err := expandCloudfunctions2functionServiceConfigEnvironmentVariables(original["environment_variables"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnvironmentVariables); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["environmentVariables"] = transformedEnvironmentVariables
	}

	transformedMaxInstanceCount, err := expandCloudfunctions2functionServiceConfigMaxInstanceCount(original["max_instance_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxInstanceCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxInstanceCount"] = transformedMaxInstanceCount
	}

	transformedMinInstanceCount, err := expandCloudfunctions2functionServiceConfigMinInstanceCount(original["min_instance_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinInstanceCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minInstanceCount"] = transformedMinInstanceCount
	}

	transformedVpcConnector, err := expandCloudfunctions2functionServiceConfigVpcConnector(original["vpc_connector"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVpcConnector); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["vpcConnector"] = transformedVpcConnector
	}

	transformedVpcConnectorEgressSettings, err := expandCloudfunctions2functionServiceConfigVpcConnectorEgressSettings(original["vpc_connector_egress_settings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVpcConnectorEgressSettings); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["vpcConnectorEgressSettings"] = transformedVpcConnectorEgressSettings
	}

	transformedIngressSettings, err := expandCloudfunctions2functionServiceConfigIngressSettings(original["ingress_settings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIngressSettings); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ingressSettings"] = transformedIngressSettings
	}

	transformedServiceAccountEmail, err := expandCloudfunctions2functionServiceConfigServiceAccountEmail(original["service_account_email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAccountEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serviceAccountEmail"] = transformedServiceAccountEmail
	}

	transformedAllTrafficOnLatestRevision, err := expandCloudfunctions2functionServiceConfigAllTrafficOnLatestRevision(original["all_traffic_on_latest_revision"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllTrafficOnLatestRevision); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allTrafficOnLatestRevision"] = transformedAllTrafficOnLatestRevision
	}

	transformedSecretEnvironmentVariables, err := expandCloudfunctions2functionServiceConfigSecretEnvironmentVariables(original["secret_environment_variables"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSecretEnvironmentVariables); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["secretEnvironmentVariables"] = transformedSecretEnvironmentVariables
	}

	transformedSecretVolumes, err := expandCloudfunctions2functionServiceConfigSecretVolumes(original["secret_volumes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSecretVolumes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["secretVolumes"] = transformedSecretVolumes
	}

	return transformed, nil
}

func expandCloudfunctions2functionServiceConfigService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigTimeoutSeconds(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigAvailableMemory(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigMaxInstanceRequestConcurrency(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigAvailableCpu(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigEnvironmentVariables(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudfunctions2functionServiceConfigMaxInstanceCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigMinInstanceCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigVpcConnector(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigVpcConnectorEgressSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigIngressSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigServiceAccountEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigAllTrafficOnLatestRevision(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigSecretEnvironmentVariables(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedKey, err := expandCloudfunctions2functionServiceConfigSecretEnvironmentVariablesKey(original["key"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["key"] = transformedKey
		}

		transformedProjectId, err := expandCloudfunctions2functionServiceConfigSecretEnvironmentVariablesProjectId(original["project_id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["projectId"] = transformedProjectId
		}

		transformedSecret, err := expandCloudfunctions2functionServiceConfigSecretEnvironmentVariablesSecret(original["secret"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSecret); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["secret"] = transformedSecret
		}

		transformedVersion, err := expandCloudfunctions2functionServiceConfigSecretEnvironmentVariablesVersion(original["version"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["version"] = transformedVersion
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudfunctions2functionServiceConfigSecretEnvironmentVariablesKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigSecretEnvironmentVariablesProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigSecretEnvironmentVariablesSecret(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigSecretEnvironmentVariablesVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigSecretVolumes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedMountPath, err := expandCloudfunctions2functionServiceConfigSecretVolumesMountPath(original["mount_path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMountPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["mountPath"] = transformedMountPath
		}

		transformedProjectId, err := expandCloudfunctions2functionServiceConfigSecretVolumesProjectId(original["project_id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["projectId"] = transformedProjectId
		}

		transformedSecret, err := expandCloudfunctions2functionServiceConfigSecretVolumesSecret(original["secret"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSecret); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["secret"] = transformedSecret
		}

		transformedVersions, err := expandCloudfunctions2functionServiceConfigSecretVolumesVersions(original["versions"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedVersions); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["versions"] = transformedVersions
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudfunctions2functionServiceConfigSecretVolumesMountPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigSecretVolumesProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigSecretVolumesSecret(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigSecretVolumesVersions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedVersion, err := expandCloudfunctions2functionServiceConfigSecretVolumesVersionsVersion(original["version"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["version"] = transformedVersion
		}

		transformedPath, err := expandCloudfunctions2functionServiceConfigSecretVolumesVersionsPath(original["path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["path"] = transformedPath
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudfunctions2functionServiceConfigSecretVolumesVersionsVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionServiceConfigSecretVolumesVersionsPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionEventTrigger(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTriggerRegion, err := expandCloudfunctions2functionEventTriggerTriggerRegion(original["trigger_region"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTriggerRegion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["triggerRegion"] = transformedTriggerRegion
	}

	transformedEventType, err := expandCloudfunctions2functionEventTriggerEventType(original["event_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEventType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["eventType"] = transformedEventType
	}

	transformedEventFilters, err := expandCloudfunctions2functionEventTriggerEventFilters(original["event_filters"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEventFilters); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["eventFilters"] = transformedEventFilters
	}

	transformedPubsubTopic, err := expandCloudfunctions2functionEventTriggerPubsubTopic(original["pubsub_topic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPubsubTopic); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pubsubTopic"] = transformedPubsubTopic
	}

	transformedServiceAccountEmail, err := expandCloudfunctions2functionEventTriggerServiceAccountEmail(original["service_account_email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAccountEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serviceAccountEmail"] = transformedServiceAccountEmail
	}

	transformedRetryPolicy, err := expandCloudfunctions2functionEventTriggerRetryPolicy(original["retry_policy"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRetryPolicy); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["retryPolicy"] = transformedRetryPolicy
	}

	return transformed, nil
}

func expandCloudfunctions2functionEventTriggerTriggerRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionEventTriggerEventType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionEventTriggerEventFilters(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedAttribute, err := expandCloudfunctions2functionEventTriggerEventFiltersAttribute(original["attribute"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAttribute); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["attribute"] = transformedAttribute
		}

		transformedValue, err := expandCloudfunctions2functionEventTriggerEventFiltersValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		transformedOperator, err := expandCloudfunctions2functionEventTriggerEventFiltersOperator(original["operator"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedOperator); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["operator"] = transformedOperator
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudfunctions2functionEventTriggerEventFiltersAttribute(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionEventTriggerEventFiltersValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionEventTriggerEventFiltersOperator(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionEventTriggerPubsubTopic(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionEventTriggerServiceAccountEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionEventTriggerRetryPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudfunctions2functionEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
