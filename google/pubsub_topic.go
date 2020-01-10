// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//     ***     DIFF TEST DIFF TEST    ***
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

import "reflect"

func GetPubsubTopicCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//pubsub.googleapis.com/projects/{{project}}/topics/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetPubsubTopicApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "pubsub.googleapis.com/Topic",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/pubsub/v1/rest",
				DiscoveryName:        "Topic",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetPubsubTopicApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandPubsubTopicName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	kmsKeyNameProp, err := expandPubsubTopicKmsKeyName(d.Get("kms_key_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("kms_key_name"); !isEmptyValue(reflect.ValueOf(kmsKeyNameProp)) && (ok || !reflect.DeepEqual(v, kmsKeyNameProp)) {
		obj["kmsKeyName"] = kmsKeyNameProp
	}
	labelsProp, err := expandPubsubTopicLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	messageStoragePolicyProp, err := expandPubsubTopicMessageStoragePolicy(d.Get("message_storage_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("message_storage_policy"); !isEmptyValue(reflect.ValueOf(messageStoragePolicyProp)) && (ok || !reflect.DeepEqual(v, messageStoragePolicyProp)) {
		obj["messageStoragePolicy"] = messageStoragePolicyProp
	}

	return resourcePubsubTopicEncoder(d, config, obj)
}

func resourcePubsubTopicEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	delete(obj, "name")
	return obj, nil
}

func expandPubsubTopicName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return GetResourceNameFromSelfLink(v.(string)), nil
}

func expandPubsubTopicKmsKeyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandPubsubTopicMessageStoragePolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAllowedPersistenceRegions, err := expandPubsubTopicMessageStoragePolicyAllowedPersistenceRegions(original["allowed_persistence_regions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedPersistenceRegions); val.IsValid() && !isEmptyValue(val) {
		transformed["allowedPersistenceRegions"] = transformedAllowedPersistenceRegions
	}

	return transformed, nil
}

func expandPubsubTopicMessageStoragePolicyAllowedPersistenceRegions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
