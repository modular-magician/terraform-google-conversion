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

package secretmanager

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const SecretManagerSecretAssetType string = "secretmanager.googleapis.com/Secret"

func ResourceConverterSecretManagerSecret() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: SecretManagerSecretAssetType,
		Convert:   GetSecretManagerSecretCaiObject,
	}
}

func GetSecretManagerSecretCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//secretmanager.googleapis.com/projects/{{project}}/secrets/{{secret_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetSecretManagerSecretApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: SecretManagerSecretAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/secretmanager/v1/rest",
				DiscoveryName:        "Secret",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetSecretManagerSecretApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	labelsProp, err := expandSecretManagerSecretLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	annotationsProp, err := expandSecretManagerSecretAnnotations(d.Get("annotations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}
	versionAliasesProp, err := expandSecretManagerSecretVersionAliases(d.Get("version_aliases"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("version_aliases"); !tpgresource.IsEmptyValue(reflect.ValueOf(versionAliasesProp)) && (ok || !reflect.DeepEqual(v, versionAliasesProp)) {
		obj["versionAliases"] = versionAliasesProp
	}
	replicationProp, err := expandSecretManagerSecretReplication(d.Get("replication"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("replication"); !tpgresource.IsEmptyValue(reflect.ValueOf(replicationProp)) && (ok || !reflect.DeepEqual(v, replicationProp)) {
		obj["replication"] = replicationProp
	}
	topicsProp, err := expandSecretManagerSecretTopics(d.Get("topics"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("topics"); !tpgresource.IsEmptyValue(reflect.ValueOf(topicsProp)) && (ok || !reflect.DeepEqual(v, topicsProp)) {
		obj["topics"] = topicsProp
	}
	expireTimeProp, err := expandSecretManagerSecretExpireTime(d.Get("expire_time"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("expire_time"); !tpgresource.IsEmptyValue(reflect.ValueOf(expireTimeProp)) && (ok || !reflect.DeepEqual(v, expireTimeProp)) {
		obj["expireTime"] = expireTimeProp
	}
	ttlProp, err := expandSecretManagerSecretTtl(d.Get("ttl"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ttl"); !tpgresource.IsEmptyValue(reflect.ValueOf(ttlProp)) && (ok || !reflect.DeepEqual(v, ttlProp)) {
		obj["ttl"] = ttlProp
	}
	rotationProp, err := expandSecretManagerSecretRotation(d.Get("rotation"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("rotation"); !tpgresource.IsEmptyValue(reflect.ValueOf(rotationProp)) && (ok || !reflect.DeepEqual(v, rotationProp)) {
		obj["rotation"] = rotationProp
	}

	return obj, nil
}

func expandSecretManagerSecretLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandSecretManagerSecretAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandSecretManagerSecretVersionAliases(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandSecretManagerSecretReplication(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAutomatic, err := expandSecretManagerSecretReplicationAutomatic(original["automatic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAutomatic); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["automatic"] = transformedAutomatic
	}

	transformedUserManaged, err := expandSecretManagerSecretReplicationUserManaged(original["user_managed"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUserManaged); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["userManaged"] = transformedUserManaged
	}

	return transformed, nil
}

func expandSecretManagerSecretReplicationAutomatic(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil || !v.(bool) {
		return nil, nil
	}

	return struct{}{}, nil
}

func expandSecretManagerSecretReplicationUserManaged(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedReplicas, err := expandSecretManagerSecretReplicationUserManagedReplicas(original["replicas"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedReplicas); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["replicas"] = transformedReplicas
	}

	return transformed, nil
}

func expandSecretManagerSecretReplicationUserManagedReplicas(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedLocation, err := expandSecretManagerSecretReplicationUserManagedReplicasLocation(original["location"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["location"] = transformedLocation
		}

		transformedCustomerManagedEncryption, err := expandSecretManagerSecretReplicationUserManagedReplicasCustomerManagedEncryption(original["customer_managed_encryption"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCustomerManagedEncryption); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["customerManagedEncryption"] = transformedCustomerManagedEncryption
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandSecretManagerSecretReplicationUserManagedReplicasLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerSecretReplicationUserManagedReplicasCustomerManagedEncryption(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeyName, err := expandSecretManagerSecretReplicationUserManagedReplicasCustomerManagedEncryptionKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	return transformed, nil
}

func expandSecretManagerSecretReplicationUserManagedReplicasCustomerManagedEncryptionKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerSecretTopics(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandSecretManagerSecretTopicsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandSecretManagerSecretTopicsName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerSecretExpireTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerSecretTtl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerSecretRotation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNextRotationTime, err := expandSecretManagerSecretRotationNextRotationTime(original["next_rotation_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNextRotationTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["nextRotationTime"] = transformedNextRotationTime
	}

	transformedRotationPeriod, err := expandSecretManagerSecretRotationRotationPeriod(original["rotation_period"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRotationPeriod); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rotationPeriod"] = transformedRotationPeriod
	}

	return transformed, nil
}

func expandSecretManagerSecretRotationNextRotationTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerSecretRotationRotationPeriod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
