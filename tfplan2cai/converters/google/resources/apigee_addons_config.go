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

package google

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const ApigeeAddonsConfigAssetType string = "apigee.googleapis.com/AddonsConfig"

func resourceConverterApigeeAddonsConfig() ResourceConverter {
	return ResourceConverter{
		AssetType: ApigeeAddonsConfigAssetType,
		Convert:   GetApigeeAddonsConfigCaiObject,
	}
}

func GetApigeeAddonsConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//apigee.googleapis.com/organizations/{{org}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetApigeeAddonsConfigApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: ApigeeAddonsConfigAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/apigee/v1/rest",
				DiscoveryName:        "AddonsConfig",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetApigeeAddonsConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	addonsConfigProp, err := expandApigeeAddonsConfigAddonsConfig(d.Get("addons_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("addons_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(addonsConfigProp)) && (ok || !reflect.DeepEqual(v, addonsConfigProp)) {
		obj["addonsConfig"] = addonsConfigProp
	}

	return obj, nil
}

func expandApigeeAddonsConfigAddonsConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAdvancedApiOpsConfig, err := expandApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfig(original["advanced_api_ops_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAdvancedApiOpsConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["advancedApiOpsConfig"] = transformedAdvancedApiOpsConfig
	}

	transformedIntegrationConfig, err := expandApigeeAddonsConfigAddonsConfigIntegrationConfig(original["integration_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIntegrationConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["integrationConfig"] = transformedIntegrationConfig
	}

	transformedMonetizationConfig, err := expandApigeeAddonsConfigAddonsConfigMonetizationConfig(original["monetization_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMonetizationConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["monetizationConfig"] = transformedMonetizationConfig
	}

	transformedApiSecurityConfig, err := expandApigeeAddonsConfigAddonsConfigApiSecurityConfig(original["api_security_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedApiSecurityConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["apiSecurityConfig"] = transformedApiSecurityConfig
	}

	transformedConnectorsPlatformConfig, err := expandApigeeAddonsConfigAddonsConfigConnectorsPlatformConfig(original["connectors_platform_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConnectorsPlatformConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["connectorsPlatformConfig"] = transformedConnectorsPlatformConfig
	}

	return transformed, nil
}

func expandApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfigEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	return transformed, nil
}

func expandApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfigEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeAddonsConfigAddonsConfigIntegrationConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandApigeeAddonsConfigAddonsConfigIntegrationConfigEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	return transformed, nil
}

func expandApigeeAddonsConfigAddonsConfigIntegrationConfigEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeAddonsConfigAddonsConfigMonetizationConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandApigeeAddonsConfigAddonsConfigMonetizationConfigEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	return transformed, nil
}

func expandApigeeAddonsConfigAddonsConfigMonetizationConfigEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeAddonsConfigAddonsConfigApiSecurityConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandApigeeAddonsConfigAddonsConfigApiSecurityConfigEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	transformedExpiresAt, err := expandApigeeAddonsConfigAddonsConfigApiSecurityConfigExpiresAt(original["expires_at"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpiresAt); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["expiresAt"] = transformedExpiresAt
	}

	return transformed, nil
}

func expandApigeeAddonsConfigAddonsConfigApiSecurityConfigEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeAddonsConfigAddonsConfigApiSecurityConfigExpiresAt(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeAddonsConfigAddonsConfigConnectorsPlatformConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandApigeeAddonsConfigAddonsConfigConnectorsPlatformConfigEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	transformedExpiresAt, err := expandApigeeAddonsConfigAddonsConfigConnectorsPlatformConfigExpiresAt(original["expires_at"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpiresAt); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["expiresAt"] = transformedExpiresAt
	}

	return transformed, nil
}

func expandApigeeAddonsConfigAddonsConfigConnectorsPlatformConfigEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeAddonsConfigAddonsConfigConnectorsPlatformConfigExpiresAt(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
