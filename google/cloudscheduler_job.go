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
	"fmt"
	"reflect"
)

func GetCloudSchedulerJobCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	if obj, err := GetCloudSchedulerJobApiObject(d, config); err == nil {
		return Asset{
			Name: fmt.Sprintf("//cloudscheduler.googleapis.com/%s", obj["selfLink"]),
			Type: "google.cloudscheduler.Job",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudscheduler/v1/rest",
				DiscoveryName:        "Job",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetCloudSchedulerJobApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandCloudSchedulerJobName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandCloudSchedulerJobDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	scheduleProp, err := expandCloudSchedulerJobSchedule(d.Get("schedule"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("schedule"); !isEmptyValue(reflect.ValueOf(scheduleProp)) && (ok || !reflect.DeepEqual(v, scheduleProp)) {
		obj["schedule"] = scheduleProp
	}
	timeZoneProp, err := expandCloudSchedulerJobTimeZone(d.Get("time_zone"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("time_zone"); !isEmptyValue(reflect.ValueOf(timeZoneProp)) && (ok || !reflect.DeepEqual(v, timeZoneProp)) {
		obj["timeZone"] = timeZoneProp
	}
	retryConfigProp, err := expandCloudSchedulerJobRetryConfig(d.Get("retry_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("retry_config"); !isEmptyValue(reflect.ValueOf(retryConfigProp)) && (ok || !reflect.DeepEqual(v, retryConfigProp)) {
		obj["retryConfig"] = retryConfigProp
	}
	pubsubTargetProp, err := expandCloudSchedulerJobPubsubTarget(d.Get("pubsub_target"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("pubsub_target"); !isEmptyValue(reflect.ValueOf(pubsubTargetProp)) && (ok || !reflect.DeepEqual(v, pubsubTargetProp)) {
		obj["pubsubTarget"] = pubsubTargetProp
	}
	appEngineHttpTargetProp, err := expandCloudSchedulerJobAppEngineHttpTarget(d.Get("app_engine_http_target"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("app_engine_http_target"); !isEmptyValue(reflect.ValueOf(appEngineHttpTargetProp)) && (ok || !reflect.DeepEqual(v, appEngineHttpTargetProp)) {
		obj["appEngineHttpTarget"] = appEngineHttpTargetProp
	}
	httpTargetProp, err := expandCloudSchedulerJobHttpTarget(d.Get("http_target"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("http_target"); !isEmptyValue(reflect.ValueOf(httpTargetProp)) && (ok || !reflect.DeepEqual(v, httpTargetProp)) {
		obj["httpTarget"] = httpTargetProp
	}

	return obj, nil
}

func expandCloudSchedulerJobName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	var jobName string
	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}

	region, err := getRegion(d, config)
	if err != nil {
		return nil, err
	}

	if v, ok := d.GetOk("name"); ok {
		jobName = fmt.Sprintf("projects/%s/locations/%s/jobs/%s", project, region, v.(string))
	} else {
		err := fmt.Errorf("The name is missing for the job cannot be empty")
		return nil, err
	}

	return jobName, nil
}

func expandCloudSchedulerJobDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobSchedule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobTimeZone(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobRetryConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
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
	} else if val := reflect.ValueOf(transformedRetryCount); val.IsValid() && !isEmptyValue(val) {
		transformed["retryCount"] = transformedRetryCount
	}

	transformedMaxRetryDuration, err := expandCloudSchedulerJobRetryConfigMaxRetryDuration(original["max_retry_duration"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxRetryDuration); val.IsValid() && !isEmptyValue(val) {
		transformed["maxRetryDuration"] = transformedMaxRetryDuration
	}

	transformedMinBackoffDuration, err := expandCloudSchedulerJobRetryConfigMinBackoffDuration(original["min_backoff_duration"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinBackoffDuration); val.IsValid() && !isEmptyValue(val) {
		transformed["minBackoffDuration"] = transformedMinBackoffDuration
	}

	transformedMaxBackoffDuration, err := expandCloudSchedulerJobRetryConfigMaxBackoffDuration(original["max_backoff_duration"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxBackoffDuration); val.IsValid() && !isEmptyValue(val) {
		transformed["maxBackoffDuration"] = transformedMaxBackoffDuration
	}

	transformedMaxDoublings, err := expandCloudSchedulerJobRetryConfigMaxDoublings(original["max_doublings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxDoublings); val.IsValid() && !isEmptyValue(val) {
		transformed["maxDoublings"] = transformedMaxDoublings
	}

	return transformed, nil
}

func expandCloudSchedulerJobRetryConfigRetryCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobRetryConfigMaxRetryDuration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobRetryConfigMinBackoffDuration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobRetryConfigMaxBackoffDuration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobRetryConfigMaxDoublings(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobPubsubTarget(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
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
	} else if val := reflect.ValueOf(transformedTopicName); val.IsValid() && !isEmptyValue(val) {
		transformed["topicName"] = transformedTopicName
	}

	transformedData, err := expandCloudSchedulerJobPubsubTargetData(original["data"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedData); val.IsValid() && !isEmptyValue(val) {
		transformed["data"] = transformedData
	}

	transformedAttributes, err := expandCloudSchedulerJobPubsubTargetAttributes(original["attributes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAttributes); val.IsValid() && !isEmptyValue(val) {
		transformed["attributes"] = transformedAttributes
	}

	return transformed, nil
}

func expandCloudSchedulerJobPubsubTargetTopicName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobPubsubTargetData(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobPubsubTargetAttributes(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudSchedulerJobAppEngineHttpTarget(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
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
	} else if val := reflect.ValueOf(transformedHttpMethod); val.IsValid() && !isEmptyValue(val) {
		transformed["httpMethod"] = transformedHttpMethod
	}

	transformedAppEngineRouting, err := expandCloudSchedulerJobAppEngineHttpTargetAppEngineRouting(original["app_engine_routing"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAppEngineRouting); val.IsValid() && !isEmptyValue(val) {
		transformed["appEngineRouting"] = transformedAppEngineRouting
	}

	transformedRelativeUri, err := expandCloudSchedulerJobAppEngineHttpTargetRelativeUri(original["relative_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRelativeUri); val.IsValid() && !isEmptyValue(val) {
		transformed["relativeUri"] = transformedRelativeUri
	}

	transformedBody, err := expandCloudSchedulerJobAppEngineHttpTargetBody(original["body"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBody); val.IsValid() && !isEmptyValue(val) {
		transformed["body"] = transformedBody
	}

	transformedHeaders, err := expandCloudSchedulerJobAppEngineHttpTargetHeaders(original["headers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHeaders); val.IsValid() && !isEmptyValue(val) {
		transformed["headers"] = transformedHeaders
	}

	return transformed, nil
}

func expandCloudSchedulerJobAppEngineHttpTargetHttpMethod(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobAppEngineHttpTargetAppEngineRouting(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
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
	} else if val := reflect.ValueOf(transformedService); val.IsValid() && !isEmptyValue(val) {
		transformed["service"] = transformedService
	}

	transformedVersion, err := expandCloudSchedulerJobAppEngineHttpTargetAppEngineRoutingVersion(original["version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !isEmptyValue(val) {
		transformed["version"] = transformedVersion
	}

	transformedInstance, err := expandCloudSchedulerJobAppEngineHttpTargetAppEngineRoutingInstance(original["instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInstance); val.IsValid() && !isEmptyValue(val) {
		transformed["instance"] = transformedInstance
	}

	return transformed, nil
}

func expandCloudSchedulerJobAppEngineHttpTargetAppEngineRoutingService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobAppEngineHttpTargetAppEngineRoutingVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobAppEngineHttpTargetAppEngineRoutingInstance(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobAppEngineHttpTargetRelativeUri(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobAppEngineHttpTargetBody(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobAppEngineHttpTargetHeaders(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudSchedulerJobHttpTarget(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
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
	} else if val := reflect.ValueOf(transformedUri); val.IsValid() && !isEmptyValue(val) {
		transformed["uri"] = transformedUri
	}

	transformedHttpMethod, err := expandCloudSchedulerJobHttpTargetHttpMethod(original["http_method"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHttpMethod); val.IsValid() && !isEmptyValue(val) {
		transformed["httpMethod"] = transformedHttpMethod
	}

	transformedBody, err := expandCloudSchedulerJobHttpTargetBody(original["body"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBody); val.IsValid() && !isEmptyValue(val) {
		transformed["body"] = transformedBody
	}

	transformedHeaders, err := expandCloudSchedulerJobHttpTargetHeaders(original["headers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHeaders); val.IsValid() && !isEmptyValue(val) {
		transformed["headers"] = transformedHeaders
	}

	return transformed, nil
}

func expandCloudSchedulerJobHttpTargetUri(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobHttpTargetHttpMethod(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobHttpTargetBody(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudSchedulerJobHttpTargetHeaders(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
