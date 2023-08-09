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

package compute

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeBackendBucketAssetType string = "compute.googleapis.com/BackendBucket"

func ResourceConverterComputeBackendBucket() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: ComputeBackendBucketAssetType,
		Convert:   GetComputeBackendBucketCaiObject,
	}
}

func GetComputeBackendBucketCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/backendBuckets/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetComputeBackendBucketApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: ComputeBackendBucketAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "BackendBucket",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetComputeBackendBucketApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	bucketNameProp, err := expandComputeBackendBucketBucketName(d.Get("bucket_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("bucket_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(bucketNameProp)) && (ok || !reflect.DeepEqual(v, bucketNameProp)) {
		obj["bucketName"] = bucketNameProp
	}
	cdnPolicyProp, err := expandComputeBackendBucketCdnPolicy(d.Get("cdn_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cdn_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(cdnPolicyProp)) && (ok || !reflect.DeepEqual(v, cdnPolicyProp)) {
		obj["cdnPolicy"] = cdnPolicyProp
	}
	compressionModeProp, err := expandComputeBackendBucketCompressionMode(d.Get("compression_mode"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("compression_mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(compressionModeProp)) && (ok || !reflect.DeepEqual(v, compressionModeProp)) {
		obj["compressionMode"] = compressionModeProp
	}
	edgeSecurityPolicyProp, err := expandComputeBackendBucketEdgeSecurityPolicy(d.Get("edge_security_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("edge_security_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(edgeSecurityPolicyProp)) && (ok || !reflect.DeepEqual(v, edgeSecurityPolicyProp)) {
		obj["edgeSecurityPolicy"] = edgeSecurityPolicyProp
	}
	customResponseHeadersProp, err := expandComputeBackendBucketCustomResponseHeaders(d.Get("custom_response_headers"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("custom_response_headers"); !tpgresource.IsEmptyValue(reflect.ValueOf(customResponseHeadersProp)) && (ok || !reflect.DeepEqual(v, customResponseHeadersProp)) {
		obj["customResponseHeaders"] = customResponseHeadersProp
	}
	descriptionProp, err := expandComputeBackendBucketDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	enableCdnProp, err := expandComputeBackendBucketEnableCdn(d.Get("enable_cdn"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_cdn"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableCdnProp)) && (ok || !reflect.DeepEqual(v, enableCdnProp)) {
		obj["enableCdn"] = enableCdnProp
	}
	nameProp, err := expandComputeBackendBucketName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	return resourceComputeBackendBucketEncoder(d, config, obj)
}

func resourceComputeBackendBucketEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// This custom encoder helps prevent sending 0 for clientTtl, defaultTtl and
	// maxTtl in API calls to update these values  when unset in the provider
	// (doing so results in an API level error)
	c, cdnPolicyOk := d.GetOk("cdn_policy")

	// Only apply during updates
	if !cdnPolicyOk || obj["cdnPolicy"] == nil {
		return obj, nil
	}

	currentCdnPolicies := c.([]interface{})

	// state does not contain cdnPolicy, so we can return early here as well
	if len(currentCdnPolicies) == 0 {
		return obj, nil
	}

	futureCdnPolicy := obj["cdnPolicy"].(map[string]interface{})
	currentCdnPolicy := currentCdnPolicies[0].(map[string]interface{})

	cacheMode, ok := futureCdnPolicy["cache_mode"].(string)
	// Fallback to state if doesn't exist in object
	if !ok {
		cacheMode = currentCdnPolicy["cache_mode"].(string)
	}

	switch cacheMode {
	case "USE_ORIGIN_HEADERS":
		if _, ok := futureCdnPolicy["clientTtl"]; ok {
			delete(futureCdnPolicy, "clientTtl")
		}
		if _, ok := futureCdnPolicy["defaultTtl"]; ok {
			delete(futureCdnPolicy, "defaultTtl")
		}
		if _, ok := futureCdnPolicy["maxTtl"]; ok {
			delete(futureCdnPolicy, "maxTtl")
		}
	}

	return obj, nil
}

func expandComputeBackendBucketBucketName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCacheKeyPolicy, err := expandComputeBackendBucketCdnPolicyCacheKeyPolicy(original["cache_key_policy"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCacheKeyPolicy); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cacheKeyPolicy"] = transformedCacheKeyPolicy
	}

	transformedSignedUrlCacheMaxAgeSec, err := expandComputeBackendBucketCdnPolicySignedUrlCacheMaxAgeSec(original["signed_url_cache_max_age_sec"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["signedUrlCacheMaxAgeSec"] = transformedSignedUrlCacheMaxAgeSec
	}

	transformedDefaultTtl, err := expandComputeBackendBucketCdnPolicyDefaultTtl(original["default_ttl"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["defaultTtl"] = transformedDefaultTtl
	}

	transformedMaxTtl, err := expandComputeBackendBucketCdnPolicyMaxTtl(original["max_ttl"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxTtl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxTtl"] = transformedMaxTtl
	}

	transformedClientTtl, err := expandComputeBackendBucketCdnPolicyClientTtl(original["client_ttl"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["clientTtl"] = transformedClientTtl
	}

	transformedNegativeCaching, err := expandComputeBackendBucketCdnPolicyNegativeCaching(original["negative_caching"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["negativeCaching"] = transformedNegativeCaching
	}

	transformedNegativeCachingPolicy, err := expandComputeBackendBucketCdnPolicyNegativeCachingPolicy(original["negative_caching_policy"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNegativeCachingPolicy); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["negativeCachingPolicy"] = transformedNegativeCachingPolicy
	}

	transformedCacheMode, err := expandComputeBackendBucketCdnPolicyCacheMode(original["cache_mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCacheMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cacheMode"] = transformedCacheMode
	}

	transformedServeWhileStale, err := expandComputeBackendBucketCdnPolicyServeWhileStale(original["serve_while_stale"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["serveWhileStale"] = transformedServeWhileStale
	}

	transformedRequestCoalescing, err := expandComputeBackendBucketCdnPolicyRequestCoalescing(original["request_coalescing"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["requestCoalescing"] = transformedRequestCoalescing
	}

	transformedBypassCacheOnRequestHeaders, err := expandComputeBackendBucketCdnPolicyBypassCacheOnRequestHeaders(original["bypass_cache_on_request_headers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBypassCacheOnRequestHeaders); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["bypassCacheOnRequestHeaders"] = transformedBypassCacheOnRequestHeaders
	}

	return transformed, nil
}

func expandComputeBackendBucketCdnPolicyCacheKeyPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedQueryStringWhitelist, err := expandComputeBackendBucketCdnPolicyCacheKeyPolicyQueryStringWhitelist(original["query_string_whitelist"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["queryStringWhitelist"] = transformedQueryStringWhitelist
	}

	transformedIncludeHttpHeaders, err := expandComputeBackendBucketCdnPolicyCacheKeyPolicyIncludeHttpHeaders(original["include_http_headers"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["includeHttpHeaders"] = transformedIncludeHttpHeaders
	}

	return transformed, nil
}

func expandComputeBackendBucketCdnPolicyCacheKeyPolicyQueryStringWhitelist(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyCacheKeyPolicyIncludeHttpHeaders(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicySignedUrlCacheMaxAgeSec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyDefaultTtl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyMaxTtl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyClientTtl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyNegativeCaching(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyNegativeCachingPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedCode, err := expandComputeBackendBucketCdnPolicyNegativeCachingPolicyCode(original["code"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["code"] = transformedCode
		}

		transformedTtl, err := expandComputeBackendBucketCdnPolicyNegativeCachingPolicyTtl(original["ttl"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["ttl"] = transformedTtl
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeBackendBucketCdnPolicyNegativeCachingPolicyCode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyNegativeCachingPolicyTtl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyCacheMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyServeWhileStale(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyRequestCoalescing(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCdnPolicyBypassCacheOnRequestHeaders(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedHeaderName, err := expandComputeBackendBucketCdnPolicyBypassCacheOnRequestHeadersHeaderName(original["header_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedHeaderName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["headerName"] = transformedHeaderName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeBackendBucketCdnPolicyBypassCacheOnRequestHeadersHeaderName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCompressionMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketEdgeSecurityPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketCustomResponseHeaders(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketEnableCdn(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
