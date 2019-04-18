// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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
	"bytes"
	"fmt"
	"log"
	"reflect"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGoogleComputeBackendServiceBackendHash(v interface{}) int {
	if v == nil {
		return 0
	}

	var buf bytes.Buffer
	m := v.(map[string]interface{})
	log.Printf("[DEBUG] hashing %v", m)

	if group, err := getRelativePath(m["group"].(string)); err != nil {
		log.Printf("[WARN] Error on retrieving relative path of instance group: %s", err)
		buf.WriteString(fmt.Sprintf("%s-", m["group"].(string)))
	} else {
		buf.WriteString(fmt.Sprintf("%s-", group))
	}

	if v, ok := m["balancing_mode"]; ok {
		if v == nil {
			v = ""
		}

		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}
	if v, ok := m["capacity_scaler"]; ok {
		if v == nil {
			v = 0.0
		}

		buf.WriteString(fmt.Sprintf("%f-", v.(float64)))
	}
	if v, ok := m["description"]; ok {
		if v == nil {
			v = ""
		}

		log.Printf("[DEBUG] writing description %s", v)
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}
	if v, ok := m["max_rate"]; ok {
		if v == nil {
			v = 0
		}

		buf.WriteString(fmt.Sprintf("%d-", int64(v.(int))))
	}
	if v, ok := m["max_rate_per_instance"]; ok {
		if v == nil {
			v = 0.0
		}

		buf.WriteString(fmt.Sprintf("%f-", v.(float64)))
	}
	if v, ok := m["max_connections"]; ok {
		if v == nil {
			v = 0
		}

		switch v := v.(type) {
		case float64:
			// The Golang JSON library can't tell int values apart from floats,
			// because MM doesn't give fields strong types. Since another value
			// in this block was a real float, it assumed this was a float too.
			// It's not.
			// Note that math.Round in Go is from float64 -> float64, so it will
			// be a noop. int(floatVal) truncates extra parts, so if the float64
			// representation of an int falls below the real value we'll have
			// the wrong value. eg if 3 was represented as 2.999999, that would
			// convert to 2. So we add 0.5, ensuring that we'll truncate to the
			// correct value. This wouldn't remain true if we were far enough
			// from 0 that we were off by > 0.5, but no float conversion *could*
			// work correctly in that case. 53-bit floating types as the only
			// numeric type was not a good idea, thanks Javascript.
			var vInt int
			if v < 0 {
				vInt = int(v - 0.5)
			} else {
				vInt = int(v + 0.5)
			}

			log.Printf("[DEBUG] writing float value %f as integer value %v", v, vInt)
			buf.WriteString(fmt.Sprintf("%d-", vInt))
		default:
			buf.WriteString(fmt.Sprintf("%d-", int64(v.(int))))
		}
	}
	if v, ok := m["max_connections_per_instance"]; ok {
		if v == nil {
			v = 0
		}

		switch v := v.(type) {
		case float64:
			// The Golang JSON library can't tell int values apart from floats,
			// because MM doesn't give fields strong types. Since another value
			// in this block was a real float, it assumed this was a float too.
			// It's not.
			// Note that math.Round in Go is from float64 -> float64, so it will
			// be a noop. int(floatVal) truncates extra parts, so if the float64
			// representation of an int falls below the real value we'll have
			// the wrong value. eg if 3 was represented as 2.999999, that would
			// convert to 2. So we add 0.5, ensuring that we'll truncate to the
			// correct value. This wouldn't remain true if we were far enough
			// from 0 that we were off by > 0.5, but no float conversion *could*
			// work correctly in that case. 53-bit floating types as the only
			// numeric type was not a good idea, thanks Javascript.
			var vInt int
			if v < 0 {
				vInt = int(v - 0.5)
			} else {
				vInt = int(v + 0.5)
			}

			log.Printf("[DEBUG] writing float value %f as integer value %v", v, vInt)
			buf.WriteString(fmt.Sprintf("%d-", vInt))
		default:
			buf.WriteString(fmt.Sprintf("%d-", int64(v.(int))))
		}
	}
	if v, ok := m["max_rate_per_instance"]; ok {
		if v == nil {
			v = 0.0
		}

		buf.WriteString(fmt.Sprintf("%f-", v.(float64)))
	}

	log.Printf("[DEBUG] computed hash value of %v from %v", hashcode.String(buf.String()), buf.String())
	return hashcode.String(buf.String())
}

func GetComputeBackendServiceCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/backendServices/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetComputeBackendServiceApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "compute.googleapis.com/BackendService",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "BackendService",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetComputeBackendServiceApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	affinityCookieTtlSecProp, err := expandComputeBackendServiceAffinityCookieTtlSec(d.Get("affinity_cookie_ttl_sec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("affinity_cookie_ttl_sec"); !isEmptyValue(reflect.ValueOf(affinityCookieTtlSecProp)) && (ok || !reflect.DeepEqual(v, affinityCookieTtlSecProp)) {
		obj["affinityCookieTtlSec"] = affinityCookieTtlSecProp
	}
	backendsProp, err := expandComputeBackendServiceBackend(d.Get("backend"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("backend"); !isEmptyValue(reflect.ValueOf(backendsProp)) && (ok || !reflect.DeepEqual(v, backendsProp)) {
		obj["backends"] = backendsProp
	}
	cdnPolicyProp, err := expandComputeBackendServiceCdnPolicy(d.Get("cdn_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cdn_policy"); !isEmptyValue(reflect.ValueOf(cdnPolicyProp)) && (ok || !reflect.DeepEqual(v, cdnPolicyProp)) {
		obj["cdnPolicy"] = cdnPolicyProp
	}
	connectionDrainingProp, err := expandComputeBackendServiceConnectionDraining(d, config)
	if err != nil {
		return nil, err
	} else if !isEmptyValue(reflect.ValueOf(connectionDrainingProp)) {
		obj["connectionDraining"] = connectionDrainingProp
	}
	fingerprintProp, err := expandComputeBackendServiceFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("fingerprint"); !isEmptyValue(reflect.ValueOf(fingerprintProp)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	descriptionProp, err := expandComputeBackendServiceDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	enableCDNProp, err := expandComputeBackendServiceEnableCDN(d.Get("enable_cdn"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_cdn"); !isEmptyValue(reflect.ValueOf(enableCDNProp)) && (ok || !reflect.DeepEqual(v, enableCDNProp)) {
		obj["enableCDN"] = enableCDNProp
	}
	healthChecksProp, err := expandComputeBackendServiceHealthChecks(d.Get("health_checks"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("health_checks"); !isEmptyValue(reflect.ValueOf(healthChecksProp)) && (ok || !reflect.DeepEqual(v, healthChecksProp)) {
		obj["healthChecks"] = healthChecksProp
	}
	iapProp, err := expandComputeBackendServiceIap(d.Get("iap"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("iap"); ok || !reflect.DeepEqual(v, iapProp) {
		obj["iap"] = iapProp
	}
	loadBalancingSchemeProp, err := expandComputeBackendServiceLoadBalancingScheme(d.Get("load_balancing_scheme"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("load_balancing_scheme"); !isEmptyValue(reflect.ValueOf(loadBalancingSchemeProp)) && (ok || !reflect.DeepEqual(v, loadBalancingSchemeProp)) {
		obj["loadBalancingScheme"] = loadBalancingSchemeProp
	}
	nameProp, err := expandComputeBackendServiceName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	portNameProp, err := expandComputeBackendServicePortName(d.Get("port_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("port_name"); !isEmptyValue(reflect.ValueOf(portNameProp)) && (ok || !reflect.DeepEqual(v, portNameProp)) {
		obj["portName"] = portNameProp
	}
	protocolProp, err := expandComputeBackendServiceProtocol(d.Get("protocol"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("protocol"); !isEmptyValue(reflect.ValueOf(protocolProp)) && (ok || !reflect.DeepEqual(v, protocolProp)) {
		obj["protocol"] = protocolProp
	}
	securityPolicyProp, err := expandComputeBackendServiceSecurityPolicy(d.Get("security_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("security_policy"); !isEmptyValue(reflect.ValueOf(securityPolicyProp)) && (ok || !reflect.DeepEqual(v, securityPolicyProp)) {
		obj["securityPolicy"] = securityPolicyProp
	}
	sessionAffinityProp, err := expandComputeBackendServiceSessionAffinity(d.Get("session_affinity"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("session_affinity"); !isEmptyValue(reflect.ValueOf(sessionAffinityProp)) && (ok || !reflect.DeepEqual(v, sessionAffinityProp)) {
		obj["sessionAffinity"] = sessionAffinityProp
	}
	timeoutSecProp, err := expandComputeBackendServiceTimeoutSec(d.Get("timeout_sec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("timeout_sec"); !isEmptyValue(reflect.ValueOf(timeoutSecProp)) && (ok || !reflect.DeepEqual(v, timeoutSecProp)) {
		obj["timeoutSec"] = timeoutSecProp
	}

	return resourceComputeBackendServiceEncoder(d, config, obj)
}

func resourceComputeBackendServiceEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// The BackendService API's Update / PUT API is badly formed and behaves like
	// a PATCH field for at least IAP. When sent a `null` `iap` field, the API
	// doesn't disable an existing field. To work around this, we need to emulate
	// the old Terraform behaviour of always sending the block (at both update and
	// create), and force sending each subfield as empty when the block isn't
	// present in config.

	iapVal := obj["iap"]
	if iapVal == nil {
		data := map[string]interface{}{}
		data["enabled"] = false
		data["oauth2ClientId"] = ""
		data["oauth2ClientSecret"] = ""
		obj["iap"] = data
	} else {
		iap := iapVal.(map[string]interface{})
		iap["enabled"] = true
		obj["iap"] = iap
	}

	return obj, nil
}

func expandComputeBackendServiceAffinityCookieTtlSec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceBackend(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedBalancingMode, err := expandComputeBackendServiceBackendBalancingMode(original["balancing_mode"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedBalancingMode); val.IsValid() && !isEmptyValue(val) {
			transformed["balancingMode"] = transformedBalancingMode
		}

		transformedCapacityScaler, err := expandComputeBackendServiceBackendCapacityScaler(original["capacity_scaler"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCapacityScaler); val.IsValid() && !isEmptyValue(val) {
			transformed["capacityScaler"] = transformedCapacityScaler
		}

		transformedDescription, err := expandComputeBackendServiceBackendDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedGroup, err := expandComputeBackendServiceBackendGroup(original["group"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGroup); val.IsValid() && !isEmptyValue(val) {
			transformed["group"] = transformedGroup
		}

		transformedMaxConnections, err := expandComputeBackendServiceBackendMaxConnections(original["max_connections"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMaxConnections); val.IsValid() && !isEmptyValue(val) {
			transformed["maxConnections"] = transformedMaxConnections
		}

		transformedMaxConnectionsPerInstance, err := expandComputeBackendServiceBackendMaxConnectionsPerInstance(original["max_connections_per_instance"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMaxConnectionsPerInstance); val.IsValid() && !isEmptyValue(val) {
			transformed["maxConnectionsPerInstance"] = transformedMaxConnectionsPerInstance
		}

		transformedMaxRate, err := expandComputeBackendServiceBackendMaxRate(original["max_rate"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMaxRate); val.IsValid() && !isEmptyValue(val) {
			transformed["maxRate"] = transformedMaxRate
		}

		transformedMaxRatePerInstance, err := expandComputeBackendServiceBackendMaxRatePerInstance(original["max_rate_per_instance"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMaxRatePerInstance); val.IsValid() && !isEmptyValue(val) {
			transformed["maxRatePerInstance"] = transformedMaxRatePerInstance
		}

		transformedMaxUtilization, err := expandComputeBackendServiceBackendMaxUtilization(original["max_utilization"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMaxUtilization); val.IsValid() && !isEmptyValue(val) {
			transformed["maxUtilization"] = transformedMaxUtilization
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeBackendServiceBackendBalancingMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceBackendCapacityScaler(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceBackendDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceBackendGroup(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceBackendMaxConnections(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceBackendMaxConnectionsPerInstance(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceBackendMaxRate(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceBackendMaxRatePerInstance(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceBackendMaxUtilization(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceCdnPolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCacheKeyPolicy, err := expandComputeBackendServiceCdnPolicyCacheKeyPolicy(original["cache_key_policy"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCacheKeyPolicy); val.IsValid() && !isEmptyValue(val) {
		transformed["cacheKeyPolicy"] = transformedCacheKeyPolicy
	}

	transformedSignedUrlCacheMaxAgeSec, err := expandComputeBackendServiceCdnPolicySignedUrlCacheMaxAgeSec(original["signed_url_cache_max_age_sec"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSignedUrlCacheMaxAgeSec); val.IsValid() && !isEmptyValue(val) {
		transformed["signedUrlCacheMaxAgeSec"] = transformedSignedUrlCacheMaxAgeSec
	}

	return transformed, nil
}

func expandComputeBackendServiceCdnPolicyCacheKeyPolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIncludeHost, err := expandComputeBackendServiceCdnPolicyCacheKeyPolicyIncludeHost(original["include_host"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIncludeHost); val.IsValid() && !isEmptyValue(val) {
		transformed["includeHost"] = transformedIncludeHost
	}

	transformedIncludeProtocol, err := expandComputeBackendServiceCdnPolicyCacheKeyPolicyIncludeProtocol(original["include_protocol"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIncludeProtocol); val.IsValid() && !isEmptyValue(val) {
		transformed["includeProtocol"] = transformedIncludeProtocol
	}

	transformedIncludeQueryString, err := expandComputeBackendServiceCdnPolicyCacheKeyPolicyIncludeQueryString(original["include_query_string"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIncludeQueryString); val.IsValid() && !isEmptyValue(val) {
		transformed["includeQueryString"] = transformedIncludeQueryString
	}

	transformedQueryStringBlacklist, err := expandComputeBackendServiceCdnPolicyCacheKeyPolicyQueryStringBlacklist(original["query_string_blacklist"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedQueryStringBlacklist); val.IsValid() && !isEmptyValue(val) {
		transformed["queryStringBlacklist"] = transformedQueryStringBlacklist
	}

	transformedQueryStringWhitelist, err := expandComputeBackendServiceCdnPolicyCacheKeyPolicyQueryStringWhitelist(original["query_string_whitelist"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedQueryStringWhitelist); val.IsValid() && !isEmptyValue(val) {
		transformed["queryStringWhitelist"] = transformedQueryStringWhitelist
	}

	return transformed, nil
}

func expandComputeBackendServiceCdnPolicyCacheKeyPolicyIncludeHost(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceCdnPolicyCacheKeyPolicyIncludeProtocol(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceCdnPolicyCacheKeyPolicyIncludeQueryString(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceCdnPolicyCacheKeyPolicyQueryStringBlacklist(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeBackendServiceCdnPolicyCacheKeyPolicyQueryStringWhitelist(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeBackendServiceCdnPolicySignedUrlCacheMaxAgeSec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceConnectionDraining(d TerraformResourceData, config *Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	// Note that nesting flattened objects won't work because we don't handle them properly here.
	transformedConnection_draining_timeout_sec, err := expandComputeBackendServiceConnectionDrainingConnection_draining_timeout_sec(d.Get("connection_draining_timeout_sec"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConnection_draining_timeout_sec); val.IsValid() && !isEmptyValue(val) {
		transformed["drainingTimeoutSec"] = transformedConnection_draining_timeout_sec
	}

	return transformed, nil
}

func expandComputeBackendServiceConnectionDrainingConnection_draining_timeout_sec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceFingerprint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceEnableCDN(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceHealthChecks(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeBackendServiceIap(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedOauth2ClientId, err := expandComputeBackendServiceIapOauth2ClientId(original["oauth2_client_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOauth2ClientId); val.IsValid() && !isEmptyValue(val) {
		transformed["oauth2ClientId"] = transformedOauth2ClientId
	}

	transformedOauth2ClientSecret, err := expandComputeBackendServiceIapOauth2ClientSecret(original["oauth2_client_secret"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["oauth2ClientSecret"] = transformedOauth2ClientSecret
	}

	transformedOauth2ClientSecretSha256, err := expandComputeBackendServiceIapOauth2ClientSecretSha256(original["oauth2_client_secret_sha256"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOauth2ClientSecretSha256); val.IsValid() && !isEmptyValue(val) {
		transformed["oauth2ClientSecretSha256"] = transformedOauth2ClientSecretSha256
	}

	return transformed, nil
}

func expandComputeBackendServiceIapOauth2ClientId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceIapOauth2ClientSecret(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceIapOauth2ClientSecretSha256(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceLoadBalancingScheme(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServicePortName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceProtocol(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceSecurityPolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceSessionAffinity(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendServiceTimeoutSec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
