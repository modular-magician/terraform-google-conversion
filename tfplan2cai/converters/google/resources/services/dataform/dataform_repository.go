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

package dataform

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DataformRepositoryAssetType string = "dataform.googleapis.com/Repository"

func ResourceConverterDataformRepository() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DataformRepositoryAssetType,
		Convert:   GetDataformRepositoryCaiObject,
	}
}

func GetDataformRepositoryCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//dataform.googleapis.com/projects/{{project}}/locations/{{region}}/repositories/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDataformRepositoryApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DataformRepositoryAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/dataform/v1beta1/rest",
				DiscoveryName:        "Repository",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDataformRepositoryApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandDataformRepositoryName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	gitRemoteSettingsProp, err := expandDataformRepositoryGitRemoteSettings(d.Get("git_remote_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("git_remote_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(gitRemoteSettingsProp)) && (ok || !reflect.DeepEqual(v, gitRemoteSettingsProp)) {
		obj["gitRemoteSettings"] = gitRemoteSettingsProp
	}
	workspaceCompilationOverridesProp, err := expandDataformRepositoryWorkspaceCompilationOverrides(d.Get("workspace_compilation_overrides"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("workspace_compilation_overrides"); !tpgresource.IsEmptyValue(reflect.ValueOf(workspaceCompilationOverridesProp)) && (ok || !reflect.DeepEqual(v, workspaceCompilationOverridesProp)) {
		obj["workspaceCompilationOverrides"] = workspaceCompilationOverridesProp
	}
	serviceAccountProp, err := expandDataformRepositoryServiceAccount(d.Get("service_account"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service_account"); !tpgresource.IsEmptyValue(reflect.ValueOf(serviceAccountProp)) && (ok || !reflect.DeepEqual(v, serviceAccountProp)) {
		obj["serviceAccount"] = serviceAccountProp
	}

	return obj, nil
}

func expandDataformRepositoryName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryGitRemoteSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUrl, err := expandDataformRepositoryGitRemoteSettingsUrl(original["url"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUrl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["url"] = transformedUrl
	}

	transformedDefaultBranch, err := expandDataformRepositoryGitRemoteSettingsDefaultBranch(original["default_branch"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDefaultBranch); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["defaultBranch"] = transformedDefaultBranch
	}

	transformedAuthenticationTokenSecretVersion, err := expandDataformRepositoryGitRemoteSettingsAuthenticationTokenSecretVersion(original["authentication_token_secret_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAuthenticationTokenSecretVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["authenticationTokenSecretVersion"] = transformedAuthenticationTokenSecretVersion
	}

	transformedSshAuthenticationConfig, err := expandDataformRepositoryGitRemoteSettingsSshAuthenticationConfig(original["ssh_authentication_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSshAuthenticationConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["sshAuthenticationConfig"] = transformedSshAuthenticationConfig
	}

	transformedTokenStatus, err := expandDataformRepositoryGitRemoteSettingsTokenStatus(original["token_status"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTokenStatus); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["tokenStatus"] = transformedTokenStatus
	}

	return transformed, nil
}

func expandDataformRepositoryGitRemoteSettingsUrl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryGitRemoteSettingsDefaultBranch(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryGitRemoteSettingsAuthenticationTokenSecretVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryGitRemoteSettingsSshAuthenticationConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUserPrivateKeySecretVersion, err := expandDataformRepositoryGitRemoteSettingsSshAuthenticationConfigUserPrivateKeySecretVersion(original["user_private_key_secret_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUserPrivateKeySecretVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["userPrivateKeySecretVersion"] = transformedUserPrivateKeySecretVersion
	}

	transformedHostPublicKey, err := expandDataformRepositoryGitRemoteSettingsSshAuthenticationConfigHostPublicKey(original["host_public_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHostPublicKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["hostPublicKey"] = transformedHostPublicKey
	}

	return transformed, nil
}

func expandDataformRepositoryGitRemoteSettingsSshAuthenticationConfigUserPrivateKeySecretVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryGitRemoteSettingsSshAuthenticationConfigHostPublicKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryGitRemoteSettingsTokenStatus(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkspaceCompilationOverrides(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDefaultDatabase, err := expandDataformRepositoryWorkspaceCompilationOverridesDefaultDatabase(original["default_database"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDefaultDatabase); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["defaultDatabase"] = transformedDefaultDatabase
	}

	transformedSchemaSuffix, err := expandDataformRepositoryWorkspaceCompilationOverridesSchemaSuffix(original["schema_suffix"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSchemaSuffix); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["schemaSuffix"] = transformedSchemaSuffix
	}

	transformedTablePrefix, err := expandDataformRepositoryWorkspaceCompilationOverridesTablePrefix(original["table_prefix"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTablePrefix); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["tablePrefix"] = transformedTablePrefix
	}

	return transformed, nil
}

func expandDataformRepositoryWorkspaceCompilationOverridesDefaultDatabase(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkspaceCompilationOverridesSchemaSuffix(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkspaceCompilationOverridesTablePrefix(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
