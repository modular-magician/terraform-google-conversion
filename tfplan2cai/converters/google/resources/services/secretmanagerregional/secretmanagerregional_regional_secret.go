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

package secretmanagerregional

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const SecretManagerRegionalRegionalSecretAssetType string = "secretmanager.googleapis.com/RegionalSecret"

func ResourceConverterSecretManagerRegionalRegionalSecret() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: SecretManagerRegionalRegionalSecretAssetType,
		Convert:   GetSecretManagerRegionalRegionalSecretCaiObject,
	}
}

func GetSecretManagerRegionalRegionalSecretCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//secretmanager.googleapis.com/projects/{{project}}/locations/{{location}}/secrets/{{secret_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetSecretManagerRegionalRegionalSecretApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: SecretManagerRegionalRegionalSecretAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/secretmanager/v1/rest",
				DiscoveryName:        "RegionalSecret",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetSecretManagerRegionalRegionalSecretApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	versionAliasesProp, err := expandSecretManagerRegionalRegionalSecretVersionAliases(d.Get("version_aliases"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("version_aliases"); !tpgresource.IsEmptyValue(reflect.ValueOf(versionAliasesProp)) && (ok || !reflect.DeepEqual(v, versionAliasesProp)) {
		obj["versionAliases"] = versionAliasesProp
	}
	customerManagedEncryptionProp, err := expandSecretManagerRegionalRegionalSecretCustomerManagedEncryption(d.Get("customer_managed_encryption"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("customer_managed_encryption"); !tpgresource.IsEmptyValue(reflect.ValueOf(customerManagedEncryptionProp)) && (ok || !reflect.DeepEqual(v, customerManagedEncryptionProp)) {
		obj["customerManagedEncryption"] = customerManagedEncryptionProp
	}
	topicsProp, err := expandSecretManagerRegionalRegionalSecretTopics(d.Get("topics"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("topics"); !tpgresource.IsEmptyValue(reflect.ValueOf(topicsProp)) && (ok || !reflect.DeepEqual(v, topicsProp)) {
		obj["topics"] = topicsProp
	}
	rotationProp, err := expandSecretManagerRegionalRegionalSecretRotation(d.Get("rotation"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("rotation"); !tpgresource.IsEmptyValue(reflect.ValueOf(rotationProp)) && (ok || !reflect.DeepEqual(v, rotationProp)) {
		obj["rotation"] = rotationProp
	}
	expireTimeProp, err := expandSecretManagerRegionalRegionalSecretExpireTime(d.Get("expire_time"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("expire_time"); !tpgresource.IsEmptyValue(reflect.ValueOf(expireTimeProp)) && (ok || !reflect.DeepEqual(v, expireTimeProp)) {
		obj["expireTime"] = expireTimeProp
	}
	ttlProp, err := expandSecretManagerRegionalRegionalSecretTtl(d.Get("ttl"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ttl"); !tpgresource.IsEmptyValue(reflect.ValueOf(ttlProp)) && (ok || !reflect.DeepEqual(v, ttlProp)) {
		obj["ttl"] = ttlProp
	}
	versionDestroyTtlProp, err := expandSecretManagerRegionalRegionalSecretVersionDestroyTtl(d.Get("version_destroy_ttl"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("version_destroy_ttl"); !tpgresource.IsEmptyValue(reflect.ValueOf(versionDestroyTtlProp)) && (ok || !reflect.DeepEqual(v, versionDestroyTtlProp)) {
		obj["versionDestroyTtl"] = versionDestroyTtlProp
	}
	labelsProp, err := expandSecretManagerRegionalRegionalSecretEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	annotationsProp, err := expandSecretManagerRegionalRegionalSecretEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	return obj, nil
}

func expandSecretManagerRegionalRegionalSecretVersionAliases(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandSecretManagerRegionalRegionalSecretCustomerManagedEncryption(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeyName, err := expandSecretManagerRegionalRegionalSecretCustomerManagedEncryptionKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	return transformed, nil
}

func expandSecretManagerRegionalRegionalSecretCustomerManagedEncryptionKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerRegionalRegionalSecretTopics(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandSecretManagerRegionalRegionalSecretTopicsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandSecretManagerRegionalRegionalSecretTopicsName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerRegionalRegionalSecretRotation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNextRotationTime, err := expandSecretManagerRegionalRegionalSecretRotationNextRotationTime(original["next_rotation_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNextRotationTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["nextRotationTime"] = transformedNextRotationTime
	}

	transformedRotationPeriod, err := expandSecretManagerRegionalRegionalSecretRotationRotationPeriod(original["rotation_period"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRotationPeriod); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rotationPeriod"] = transformedRotationPeriod
	}

	return transformed, nil
}

func expandSecretManagerRegionalRegionalSecretRotationNextRotationTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerRegionalRegionalSecretRotationRotationPeriod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerRegionalRegionalSecretExpireTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerRegionalRegionalSecretTtl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerRegionalRegionalSecretVersionDestroyTtl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerRegionalRegionalSecretEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandSecretManagerRegionalRegionalSecretEffectiveAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
