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

package networkservices

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const NetworkServicesEdgeCacheKeysetAssetType string = "networkservices.googleapis.com/EdgeCacheKeyset"

func ResourceConverterNetworkServicesEdgeCacheKeyset() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NetworkServicesEdgeCacheKeysetAssetType,
		Convert:   GetNetworkServicesEdgeCacheKeysetCaiObject,
	}
}

func GetNetworkServicesEdgeCacheKeysetCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//networkservices.googleapis.com/projects/{{project}}/locations/global/edgeCacheKeysets/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetNetworkServicesEdgeCacheKeysetApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: NetworkServicesEdgeCacheKeysetAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/networkservices/v1/rest",
				DiscoveryName:        "EdgeCacheKeyset",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetNetworkServicesEdgeCacheKeysetApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkServicesEdgeCacheKeysetDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	publicKeysProp, err := expandNetworkServicesEdgeCacheKeysetPublicKey(d.Get("public_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("public_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(publicKeysProp)) && (ok || !reflect.DeepEqual(v, publicKeysProp)) {
		obj["publicKeys"] = publicKeysProp
	}
	validationSharedKeysProp, err := expandNetworkServicesEdgeCacheKeysetValidationSharedKeys(d.Get("validation_shared_keys"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("validation_shared_keys"); !tpgresource.IsEmptyValue(reflect.ValueOf(validationSharedKeysProp)) && (ok || !reflect.DeepEqual(v, validationSharedKeysProp)) {
		obj["validationSharedKeys"] = validationSharedKeysProp
	}
	labelsProp, err := expandNetworkServicesEdgeCacheKeysetEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandNetworkServicesEdgeCacheKeysetDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesEdgeCacheKeysetPublicKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedId, err := expandNetworkServicesEdgeCacheKeysetPublicKeyId(original["id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["id"] = transformedId
		}

		transformedValue, err := expandNetworkServicesEdgeCacheKeysetPublicKeyValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		transformedManaged, err := expandNetworkServicesEdgeCacheKeysetPublicKeyManaged(original["managed"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedManaged); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["managed"] = transformedManaged
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkServicesEdgeCacheKeysetPublicKeyId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesEdgeCacheKeysetPublicKeyValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesEdgeCacheKeysetPublicKeyManaged(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesEdgeCacheKeysetValidationSharedKeys(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedSecretVersion, err := expandNetworkServicesEdgeCacheKeysetValidationSharedKeysSecretVersion(original["secret_version"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSecretVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["secretVersion"] = transformedSecretVersion
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkServicesEdgeCacheKeysetValidationSharedKeysSecretVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesEdgeCacheKeysetEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
