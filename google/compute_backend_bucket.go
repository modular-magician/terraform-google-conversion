// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Nathan is editing this to generate diffs in lots of files.
//     He won't submit this change.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import "reflect"

func GetComputeBackendBucketCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/backendBuckets/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetComputeBackendBucketApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "compute.googleapis.com/BackendBucket",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "BackendBucket",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetComputeBackendBucketApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	bucketNameProp, err := expandComputeBackendBucketBucketName(d.Get("bucket_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("bucket_name"); !isEmptyValue(reflect.ValueOf(bucketNameProp)) && (ok || !reflect.DeepEqual(v, bucketNameProp)) {
		obj["bucketName"] = bucketNameProp
	}
	cdnPolicyProp, err := expandComputeBackendBucketCdnPolicy(d.Get("cdn_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cdn_policy"); !isEmptyValue(reflect.ValueOf(cdnPolicyProp)) && (ok || !reflect.DeepEqual(v, cdnPolicyProp)) {
		obj["cdnPolicy"] = cdnPolicyProp
	}
	descriptionProp, err := expandComputeBackendBucketDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	enableCdnProp, err := expandComputeBackendBucketEnableCdn(d.Get("enable_cdn"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_cdn"); !isEmptyValue(reflect.ValueOf(enableCdnProp)) && (ok || !reflect.DeepEqual(v, enableCdnProp)) {
		obj["enableCdn"] = enableCdnProp
	}
	nameProp, err := expandComputeBackendBucketName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	return obj, nil
}

func expandComputeBackendBucketBucketName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSignedUrlCacheMaxAgeSec, err := expandComputeBackendBucketCdnPolicySignedUrlCacheMaxAgeSec(original["signed_url_cache_max_age_sec"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSignedUrlCacheMaxAgeSec); val.IsValid() && !isEmptyValue(val) {
		transformed["signedUrlCacheMaxAgeSec"] = transformedSignedUrlCacheMaxAgeSec
	}

	return transformed, nil
}

func expandComputeBackendBucketCdnPolicySignedUrlCacheMaxAgeSec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketEnableCdn(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
