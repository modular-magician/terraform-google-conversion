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

package cloudscheduler

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Both oidc and oauth headers cannot be set
func validateAuthHeaders(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
	httpBlock := diff.Get("http_target.0").(map[string]interface{})

	if httpBlock != nil {
		oauth := httpBlock["oauth_token"]
		oidc := httpBlock["oidc_token"]

		if oauth != nil && oidc != nil {
			if len(oidc.([]interface{})) > 0 && len(oauth.([]interface{})) > 0 {
				return fmt.Errorf("Error in http_target: only one of oauth_token or oidc_token can be specified, but not both.")
			}
		}
	}

	return nil
}

func authHeaderDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	// If generating an `oauth_token` and `scope` is not provided in the configuration,
	// the default "https://www.googleapis.com/auth/cloud-platform" scope will be used.
	// Similarly, if generating an `oidc_token` and `audience` is not provided in the
	// configuration, the URI specified in target will be used. Although not in the
	// configuration, in both cases the default is returned in the object, but is not in.
	// state. We suppress the diff if the values are these defaults but are not stored in state.

	b := strings.Split(k, ".")
	if b[0] == "http_target" && len(b) > 4 {
		block := b[2]
		attr := b[4]

		if block == "oauth_token" && attr == "scope" {
			if old == tpgresource.CanonicalizeServiceScope("cloud-platform") && new == "" {
				return true
			}
		}

		if block == "oidc_token" && attr == "audience" {
			uri := d.Get(strings.Join(b[0:2], ".") + ".uri")
			if old == uri && new == "" {
				return true
			}
		}

	}

	return false
}

func validateHttpHeaders() schema.SchemaValidateFunc {
	return func(i interface{}, k string) (s []string, es []error) {
		headers := i.(map[string]interface{})
		if _, ok := headers["Content-Length"]; ok {
			es = append(es, fmt.Errorf("Cannot set the Content-Length header on %s", k))
			return
		}
		r := regexp.MustCompile(`(X-Google-|X-AppEngine-).*`)
		for key := range headers {
			if r.MatchString(key) {
				es = append(es, fmt.Errorf("Cannot set the %s header on %s", key, k))
				return
			}
		}

		return
	}
}

const CloudSchedulerJobAssetType string = "cloudscheduler.googleapis.com/Job"

func ResourceConverterCloudSchedulerJob() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: CloudSchedulerJobAssetType,
		Convert:   GetCloudSchedulerJobCaiObject,
	}
}

func GetCloudSchedulerJobCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//cloudscheduler.googleapis.com/projects/{{project}}/locations/{{region}}/jobs/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetCloudSchedulerJobApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: CloudSchedulerJobAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudscheduler/v1/rest",
				DiscoveryName:        "Job",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetCloudSchedulerJobApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandCloudSchedulerJobName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandCloudSchedulerJobDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	scheduleProp, err := expandCloudSchedulerJobSchedule(d.Get("schedule"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("schedule"); !tpgresource.IsEmptyValue(reflect.ValueOf(scheduleProp)) && (ok || !reflect.DeepEqual(v, scheduleProp)) {
		obj["schedule"] = scheduleProp
	}
	timeZoneProp, err := expandCloudSchedulerJobTimeZone(d.Get("time_zone"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("time_zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(timeZoneProp)) && (ok || !reflect.DeepEqual(v, timeZoneProp)) {
		obj["timeZone"] = timeZoneProp
	}
	pausedProp, err := expandCloudSchedulerJobPaused(d.Get("paused"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("paused"); !tpgresource.IsEmptyValue(reflect.ValueOf(pausedProp)) && (ok || !reflect.DeepEqual(v, pausedProp)) {
		obj["paused"] = pausedProp
	}
	attemptDeadlineProp, err := expandCloudSchedulerJobAttemptDeadline(d.Get("attempt_deadline"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("attempt_deadline"); !tpgresource.IsEmptyValue(reflect.ValueOf(attemptDeadlineProp)) && (ok || !reflect.DeepEqual(v, attemptDeadlineProp)) {
		obj["attemptDeadline"] = attemptDeadlineProp
	}
	retryConfigProp, err := expandCloudSchedulerJobRetryConfig(d.Get("retry_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("retry_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(retryConfigProp)) && (ok || !reflect.DeepEqual(v, retryConfigProp)) {
		obj["retryConfig"] = retryConfigProp
	}
	pubsubTargetProp, err := expandCloudSchedulerJobPubsubTarget(d.Get("pubsub_target"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("pubsub_target"); !tpgresource.IsEmptyValue(reflect.ValueOf(pubsubTargetProp)) && (ok || !reflect.DeepEqual(v, pubsubTargetProp)) {
		obj["pubsubTarget"] = pubsubTargetProp
	}
	appEngineHttpTargetProp, err := expandCloudSchedulerJobAppEngineHttpTarget(d.Get("app_engine_http_target"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("app_engine_http_target"); !tpgresource.IsEmptyValue(reflect.ValueOf(appEngineHttpTargetProp)) && (ok || !reflect.DeepEqual(v, appEngineHttpTargetProp)) {
		obj["appEngineHttpTarget"] = appEngineHttpTargetProp
	}
	httpTargetProp, err := expandCloudSchedulerJobHttpTarget(d.Get("http_target"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("http_target"); !tpgresource.IsEmptyValue(reflect.ValueOf(httpTargetProp)) && (ok || !reflect.DeepEqual(v, httpTargetProp)) {
		obj["httpTarget"] = httpTargetProp
	}

	return resourceCloudSchedulerJobEncoder(d, config, obj)
}

func resourceCloudSchedulerJobEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	delete(obj, "paused") // Field doesn't exist in API
	return obj, nil
}

func expandCloudSchedulerJobName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/jobs/{{name}}")
}

func expandCloudSchedulerJobDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobSchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobTimeZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobPaused(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobAttemptDeadline(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobRetryConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRetryCount, err := expandCloudSchedulerJobRetryConfigRetryCount(original["retry_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRetryCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["retryCount"] = transformedRetryCount
	}

	transformedMaxRetryDuration, err := expandCloudSchedulerJobRetryConfigMaxRetryDuration(original["max_retry_duration"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxRetryDuration); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxRetryDuration"] = transformedMaxRetryDuration
	}

	transformedMinBackoffDuration, err := expandCloudSchedulerJobRetryConfigMinBackoffDuration(original["min_backoff_duration"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinBackoffDuration); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minBackoffDuration"] = transformedMinBackoffDuration
	}

	transformedMaxBackoffDuration, err := expandCloudSchedulerJobRetryConfigMaxBackoffDuration(original["max_backoff_duration"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxBackoffDuration); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxBackoffDuration"] = transformedMaxBackoffDuration
	}

	transformedMaxDoublings, err := expandCloudSchedulerJobRetryConfigMaxDoublings(original["max_doublings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxDoublings); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxDoublings"] = transformedMaxDoublings
	}

	return transformed, nil
}

func expandCloudSchedulerJobRetryConfigRetryCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobRetryConfigMaxRetryDuration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobRetryConfigMinBackoffDuration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobRetryConfigMaxBackoffDuration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobRetryConfigMaxDoublings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobPubsubTarget(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTopicName, err := expandCloudSchedulerJobPubsubTargetTopicName(original["topic_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTopicName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["topicName"] = transformedTopicName
	}

	transformedData, err := expandCloudSchedulerJobPubsubTargetData(original["data"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedData); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["data"] = transformedData
	}

	transformedAttributes, err := expandCloudSchedulerJobPubsubTargetAttributes(original["attributes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAttributes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["attributes"] = transformedAttributes
	}

	return transformed, nil
}

func expandCloudSchedulerJobPubsubTargetTopicName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobPubsubTargetData(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobPubsubTargetAttributes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudSchedulerJobAppEngineHttpTarget(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHttpMethod, err := expandCloudSchedulerJobAppEngineHttpTargetHttpMethod(original["http_method"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHttpMethod); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["httpMethod"] = transformedHttpMethod
	}

	transformedAppEngineRouting, err := expandCloudSchedulerJobAppEngineHttpTargetAppEngineRouting(original["app_engine_routing"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAppEngineRouting); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["appEngineRouting"] = transformedAppEngineRouting
	}

	transformedRelativeUri, err := expandCloudSchedulerJobAppEngineHttpTargetRelativeUri(original["relative_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRelativeUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["relativeUri"] = transformedRelativeUri
	}

	transformedBody, err := expandCloudSchedulerJobAppEngineHttpTargetBody(original["body"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBody); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["body"] = transformedBody
	}

	transformedHeaders, err := expandCloudSchedulerJobAppEngineHttpTargetHeaders(original["headers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHeaders); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["headers"] = transformedHeaders
	}

	return transformed, nil
}

func expandCloudSchedulerJobAppEngineHttpTargetHttpMethod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobAppEngineHttpTargetAppEngineRouting(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedService, err := expandCloudSchedulerJobAppEngineHttpTargetAppEngineRoutingService(original["service"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedService); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["service"] = transformedService
	}

	transformedVersion, err := expandCloudSchedulerJobAppEngineHttpTargetAppEngineRoutingVersion(original["version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["version"] = transformedVersion
	}

	transformedInstance, err := expandCloudSchedulerJobAppEngineHttpTargetAppEngineRoutingInstance(original["instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["instance"] = transformedInstance
	}

	return transformed, nil
}

func expandCloudSchedulerJobAppEngineHttpTargetAppEngineRoutingService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobAppEngineHttpTargetAppEngineRoutingVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobAppEngineHttpTargetAppEngineRoutingInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobAppEngineHttpTargetRelativeUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobAppEngineHttpTargetBody(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobAppEngineHttpTargetHeaders(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudSchedulerJobHttpTarget(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUri, err := expandCloudSchedulerJobHttpTargetUri(original["uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["uri"] = transformedUri
	}

	transformedHttpMethod, err := expandCloudSchedulerJobHttpTargetHttpMethod(original["http_method"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHttpMethod); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["httpMethod"] = transformedHttpMethod
	}

	transformedBody, err := expandCloudSchedulerJobHttpTargetBody(original["body"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBody); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["body"] = transformedBody
	}

	transformedHeaders, err := expandCloudSchedulerJobHttpTargetHeaders(original["headers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHeaders); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["headers"] = transformedHeaders
	}

	transformedOauthToken, err := expandCloudSchedulerJobHttpTargetOauthToken(original["oauth_token"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOauthToken); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["oauthToken"] = transformedOauthToken
	}

	transformedOidcToken, err := expandCloudSchedulerJobHttpTargetOidcToken(original["oidc_token"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOidcToken); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["oidcToken"] = transformedOidcToken
	}

	return transformed, nil
}

func expandCloudSchedulerJobHttpTargetUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobHttpTargetHttpMethod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobHttpTargetBody(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobHttpTargetHeaders(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudSchedulerJobHttpTargetOauthToken(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedServiceAccountEmail, err := expandCloudSchedulerJobHttpTargetOauthTokenServiceAccountEmail(original["service_account_email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAccountEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serviceAccountEmail"] = transformedServiceAccountEmail
	}

	transformedScope, err := expandCloudSchedulerJobHttpTargetOauthTokenScope(original["scope"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScope); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scope"] = transformedScope
	}

	return transformed, nil
}

func expandCloudSchedulerJobHttpTargetOauthTokenServiceAccountEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobHttpTargetOauthTokenScope(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobHttpTargetOidcToken(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedServiceAccountEmail, err := expandCloudSchedulerJobHttpTargetOidcTokenServiceAccountEmail(original["service_account_email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAccountEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serviceAccountEmail"] = transformedServiceAccountEmail
	}

	transformedAudience, err := expandCloudSchedulerJobHttpTargetOidcTokenAudience(original["audience"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAudience); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["audience"] = transformedAudience
	}

	return transformed, nil
}

func expandCloudSchedulerJobHttpTargetOidcTokenServiceAccountEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobHttpTargetOidcTokenAudience(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
