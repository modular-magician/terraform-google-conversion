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

func GetStorageBucketAccessControlCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//storage.googleapis.com/b/{{bucket}}/acl/{{entity}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetStorageBucketAccessControlApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "storage.googleapis.com/BucketAccessControl",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/storage/v1/rest",
				DiscoveryName:        "BucketAccessControl",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetStorageBucketAccessControlApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	bucketProp, err := expandStorageBucketAccessControlBucket(d.Get("bucket"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("bucket"); !isEmptyValue(reflect.ValueOf(bucketProp)) && (ok || !reflect.DeepEqual(v, bucketProp)) {
		obj["bucket"] = bucketProp
	}
	entityProp, err := expandStorageBucketAccessControlEntity(d.Get("entity"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("entity"); !isEmptyValue(reflect.ValueOf(entityProp)) && (ok || !reflect.DeepEqual(v, entityProp)) {
		obj["entity"] = entityProp
	}
	roleProp, err := expandStorageBucketAccessControlRole(d.Get("role"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("role"); !isEmptyValue(reflect.ValueOf(roleProp)) && (ok || !reflect.DeepEqual(v, roleProp)) {
		obj["role"] = roleProp
	}

	return obj, nil
}

func expandStorageBucketAccessControlBucket(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandStorageBucketAccessControlEntity(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandStorageBucketAccessControlRole(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
