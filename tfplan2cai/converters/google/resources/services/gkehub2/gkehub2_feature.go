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

package gkehub2

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const GKEHub2FeatureAssetType string = "gkehub.googleapis.com/Feature"

func ResourceConverterGKEHub2Feature() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: GKEHub2FeatureAssetType,
		Convert:   GetGKEHub2FeatureCaiObject,
	}
}

func GetGKEHub2FeatureCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//gkehub.googleapis.com/projects/{{project}}/locations/{{location}}/features/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetGKEHub2FeatureApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: GKEHub2FeatureAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/gkehub/v1beta/rest",
				DiscoveryName:        "Feature",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetGKEHub2FeatureApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	specProp, err := expandGKEHub2FeatureSpec(d.Get("spec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("spec"); !tpgresource.IsEmptyValue(reflect.ValueOf(specProp)) && (ok || !reflect.DeepEqual(v, specProp)) {
		obj["spec"] = specProp
	}
	fleetDefaultMemberConfigProp, err := expandGKEHub2FeatureFleetDefaultMemberConfig(d.Get("fleet_default_member_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("fleet_default_member_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(fleetDefaultMemberConfigProp)) && (ok || !reflect.DeepEqual(v, fleetDefaultMemberConfigProp)) {
		obj["fleetDefaultMemberConfig"] = fleetDefaultMemberConfigProp
	}
	labelsProp, err := expandGKEHub2FeatureEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandGKEHub2FeatureSpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMulticlusteringress, err := expandGKEHub2FeatureSpecMulticlusteringress(original["multiclusteringress"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMulticlusteringress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["multiclusteringress"] = transformedMulticlusteringress
	}

	transformedFleetobservability, err := expandGKEHub2FeatureSpecFleetobservability(original["fleetobservability"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFleetobservability); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fleetobservability"] = transformedFleetobservability
	}

	return transformed, nil
}

func expandGKEHub2FeatureSpecMulticlusteringress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedConfigMembership, err := expandGKEHub2FeatureSpecMulticlusteringressConfigMembership(original["config_membership"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConfigMembership); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["configMembership"] = transformedConfigMembership
	}

	return transformed, nil
}

func expandGKEHub2FeatureSpecMulticlusteringressConfigMembership(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FeatureSpecFleetobservability(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLoggingConfig, err := expandGKEHub2FeatureSpecFleetobservabilityLoggingConfig(original["logging_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLoggingConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["loggingConfig"] = transformedLoggingConfig
	}

	return transformed, nil
}

func expandGKEHub2FeatureSpecFleetobservabilityLoggingConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDefaultConfig, err := expandGKEHub2FeatureSpecFleetobservabilityLoggingConfigDefaultConfig(original["default_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDefaultConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["defaultConfig"] = transformedDefaultConfig
	}

	transformedFleetScopeLogsConfig, err := expandGKEHub2FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(original["fleet_scope_logs_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFleetScopeLogsConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fleetScopeLogsConfig"] = transformedFleetScopeLogsConfig
	}

	return transformed, nil
}

func expandGKEHub2FeatureSpecFleetobservabilityLoggingConfigDefaultConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMode, err := expandGKEHub2FeatureSpecFleetobservabilityLoggingConfigDefaultConfigMode(original["mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mode"] = transformedMode
	}

	return transformed, nil
}

func expandGKEHub2FeatureSpecFleetobservabilityLoggingConfigDefaultConfigMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMode, err := expandGKEHub2FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigMode(original["mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mode"] = transformedMode
	}

	return transformed, nil
}

func expandGKEHub2FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FeatureFleetDefaultMemberConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMesh, err := expandGKEHub2FeatureFleetDefaultMemberConfigMesh(original["mesh"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMesh); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mesh"] = transformedMesh
	}

	transformedConfigmanagement, err := expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagement(original["configmanagement"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConfigmanagement); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["configmanagement"] = transformedConfigmanagement
	}

	return transformed, nil
}

func expandGKEHub2FeatureFleetDefaultMemberConfigMesh(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedManagement, err := expandGKEHub2FeatureFleetDefaultMemberConfigMeshManagement(original["management"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedManagement); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["management"] = transformedManagement
	}

	return transformed, nil
}

func expandGKEHub2FeatureFleetDefaultMemberConfigMeshManagement(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagement(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedConfigSync, err := expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSync(original["config_sync"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConfigSync); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["configSync"] = transformedConfigSync
	}

	return transformed, nil
}

func expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSync(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSourceFormat, err := expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncSourceFormat(original["source_format"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSourceFormat); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["sourceFormat"] = transformedSourceFormat
	}

	transformedGit, err := expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGit(original["git"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGit); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["git"] = transformedGit
	}

	transformedOci, err := expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOci(original["oci"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOci); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["oci"] = transformedOci
	}

	return transformed, nil
}

func expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncSourceFormat(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGit(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSyncRepo, err := expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitSyncRepo(original["sync_repo"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSyncRepo); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["syncRepo"] = transformedSyncRepo
	}

	transformedSyncBranch, err := expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitSyncBranch(original["sync_branch"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSyncBranch); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["syncBranch"] = transformedSyncBranch
	}

	transformedPolicyDir, err := expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitPolicyDir(original["policy_dir"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPolicyDir); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["policyDir"] = transformedPolicyDir
	}

	transformedSyncRev, err := expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitSyncRev(original["sync_rev"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSyncRev); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["syncRev"] = transformedSyncRev
	}

	transformedSecretType, err := expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitSecretType(original["secret_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSecretType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["secretType"] = transformedSecretType
	}

	transformedHttpsProxy, err := expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitHttpsProxy(original["https_proxy"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHttpsProxy); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["httpsProxy"] = transformedHttpsProxy
	}

	transformedGcpServiceAccountEmail, err := expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitGcpServiceAccountEmail(original["gcp_service_account_email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGcpServiceAccountEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["gcpServiceAccountEmail"] = transformedGcpServiceAccountEmail
	}

	transformedSyncWaitSecs, err := expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitSyncWaitSecs(original["sync_wait_secs"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSyncWaitSecs); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["syncWaitSecs"] = transformedSyncWaitSecs
	}

	return transformed, nil
}

func expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitSyncRepo(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitSyncBranch(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitPolicyDir(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitSyncRev(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitSecretType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitHttpsProxy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitGcpServiceAccountEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitSyncWaitSecs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOci(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSyncRepo, err := expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciSyncRepo(original["sync_repo"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSyncRepo); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["syncRepo"] = transformedSyncRepo
	}

	transformedPolicyDir, err := expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciPolicyDir(original["policy_dir"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPolicyDir); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["policyDir"] = transformedPolicyDir
	}

	transformedSecretType, err := expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciSecretType(original["secret_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSecretType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["secretType"] = transformedSecretType
	}

	transformedGcpServiceAccountEmail, err := expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciGcpServiceAccountEmail(original["gcp_service_account_email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGcpServiceAccountEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["gcpServiceAccountEmail"] = transformedGcpServiceAccountEmail
	}

	transformedSyncWaitSecs, err := expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciSyncWaitSecs(original["sync_wait_secs"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSyncWaitSecs); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["syncWaitSecs"] = transformedSyncWaitSecs
	}

	transformedVersion, err := expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciVersion(original["version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["version"] = transformedVersion
	}

	return transformed, nil
}

func expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciSyncRepo(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciPolicyDir(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciSecretType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciGcpServiceAccountEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciSyncWaitSecs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FeatureEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
