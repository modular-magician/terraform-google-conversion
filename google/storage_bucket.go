// ----------------------------------------------------------------------------
//
//     This file is copied here by Magic Modules. The code for building up a
//     storage bucket object is copied from the manually implemented
//     provider file:
//     third_party/terraform/resources/resource_storage_bucket.go
//
// ----------------------------------------------------------------------------
package google

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"google.golang.org/api/storage/v1"
)

const StorageBucketAssetType string = "storage.googleapis.com/Bucket"

func resourceConverterStorageBucket() ResourceConverter {
	return ResourceConverter{
		AssetType: StorageBucketAssetType,
		Convert:   GetStorageBucketCaiObject,
	}
}

func GetStorageBucketCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//storage.googleapis.com/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetStorageBucketApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: StorageBucketAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/storage/v1/rest",
				DiscoveryName:        "Bucket",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetStorageBucketApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}

	// Get the bucket and location
	bucket := d.Get("name").(string)
	location := d.Get("location").(string)

	// Create a bucket, setting the labels, location and name.
	sb := &storage.Bucket{
		Name:             bucket,
		Labels:           expandLabels(d),
		Location:         location,
		IamConfiguration: expandIamConfiguration(d),
	}

	if v, ok := d.GetOk("storage_class"); ok {
		sb.StorageClass = v.(string)
	}

	if err := resourceGCSBucketLifecycleCreateOrUpdate(d, sb); err != nil {
		return nil, err
	}

	if v, ok := d.GetOk("versioning"); ok {
		sb.Versioning = expandBucketVersioning(v)
	}

	if v, ok := d.GetOk("website"); ok {
		sb.Website = expandBucketWebsite(v.([]interface{}))
	}

	if v, ok := d.GetOk("retention_policy"); ok {
		sb.RetentionPolicy = expandBucketRetentionPolicy(v.([]interface{}))
	}

	if v, ok := d.GetOk("cors"); ok {
		sb.Cors = expandCors(v.([]interface{}))
	}

	if v, ok := d.GetOk("default_event_based_hold"); ok {
		sb.DefaultEventBasedHold = v.(bool)
	}

	if v, ok := d.GetOk("logging"); ok {
		sb.Logging = expandBucketLogging(v.([]interface{}))
	}

	if v, ok := d.GetOk("encryption"); ok {
		sb.Encryption = expandBucketEncryption(v.([]interface{}))
	}

	if v, ok := d.GetOk("requester_pays"); ok {
		sb.Billing = &storage.BucketBilling{
			RequesterPays: v.(bool),
		}
	}

	m, err := jsonMap(sb)
	if err != nil {
		return nil, err
	}
	m["project"] = project

	return m, nil
}

func expandCors(configured []interface{}) []*storage.BucketCors {
	if len(configured) == 0 || configured[0] == nil {
		return nil
	}
	corsRules := make([]*storage.BucketCors, 0, len(configured))
	for _, raw := range configured {
		data := raw.(map[string]interface{})
		corsRule := storage.BucketCors{
			Origin:         convertStringArr(data["origin"].([]interface{})),
			Method:         convertStringArr(data["method"].([]interface{})),
			ResponseHeader: convertStringArr(data["response_header"].([]interface{})),
			MaxAgeSeconds:  int64(data["max_age_seconds"].(int)),
		}

		corsRules = append(corsRules, &corsRule)
	}
	return corsRules
}

func expandBucketEncryption(configured interface{}) *storage.BucketEncryption {
	encs := configured.([]interface{})
	if len(encs) == 0 || encs[0] == nil {
		return nil
	}
	enc := encs[0].(map[string]interface{})
	keyname := enc["default_kms_key_name"]
	if keyname == nil || keyname.(string) == "" {
		return nil
	}
	bucketenc := &storage.BucketEncryption{
		DefaultKmsKeyName: keyname.(string),
	}
	return bucketenc
}

func expandBucketLogging(configured interface{}) *storage.BucketLogging {
	loggings := configured.([]interface{})
	if len(loggings) == 0 {
		return nil
	}

	logging := loggings[0].(map[string]interface{})

	bucketLogging := &storage.BucketLogging{
		LogBucket:       logging["log_bucket"].(string),
		LogObjectPrefix: logging["log_object_prefix"].(string),
	}

	return bucketLogging
}

func expandBucketVersioning(configured interface{}) *storage.BucketVersioning {
	versionings := configured.([]interface{})
	if len(versionings) == 0 || versionings[0] == nil {
		return nil
	}

	versioning := versionings[0].(map[string]interface{})

	bucketVersioning := &storage.BucketVersioning{}

	bucketVersioning.Enabled = versioning["enabled"].(bool)
	bucketVersioning.ForceSendFields = append(bucketVersioning.ForceSendFields, "Enabled")

	return bucketVersioning
}

func expandBucketWebsite(v interface{}) *storage.BucketWebsite {
	if v == nil {
		return nil
	}
	vs := v.([]interface{})

	if len(vs) < 1 || vs[0] == nil {
		return nil
	}

	website := vs[0].(map[string]interface{})
	w := &storage.BucketWebsite{}

	if v := website["not_found_page"]; v != "" {
		w.NotFoundPage = v.(string)
	}

	if v := website["main_page_suffix"]; v != "" {
		w.MainPageSuffix = v.(string)
	}

	return w
}

func expandBucketRetentionPolicy(configured interface{}) *storage.BucketRetentionPolicy {
	retentionPolicies := configured.([]interface{})
	if len(retentionPolicies) == 0 {
		return nil
	}
	retentionPolicy := retentionPolicies[0].(map[string]interface{})

	bucketRetentionPolicy := &storage.BucketRetentionPolicy{
		RetentionPeriod: int64(retentionPolicy["retention_period"].(int)),
	}

	return bucketRetentionPolicy
}

func resourceGCSBucketLifecycleCreateOrUpdate(d TerraformResourceData, sb *storage.Bucket) error {
	if v, ok := d.GetOk("lifecycle_rule"); ok {
		lifecycle_rules := v.([]interface{})

		sb.Lifecycle = &storage.BucketLifecycle{}
		sb.Lifecycle.Rule = make([]*storage.BucketLifecycleRule, 0, len(lifecycle_rules))

		for _, raw_lifecycle_rule := range lifecycle_rules {
			lifecycle_rule := raw_lifecycle_rule.(map[string]interface{})

			target_lifecycle_rule := &storage.BucketLifecycleRule{}

			if v, ok := lifecycle_rule["action"]; ok {
				if actions := v.(*schema.Set).List(); len(actions) == 1 {
					action := actions[0].(map[string]interface{})

					target_lifecycle_rule.Action = &storage.BucketLifecycleRuleAction{}

					if v, ok := action["type"]; ok {
						target_lifecycle_rule.Action.Type = v.(string)
					}

					if v, ok := action["storage_class"]; ok {
						target_lifecycle_rule.Action.StorageClass = v.(string)
					}
				} else {
					return fmt.Errorf("Exactly one action is required")
				}
			}

			if v, ok := lifecycle_rule["condition"]; ok {
				if conditions := v.(*schema.Set).List(); len(conditions) == 1 {
					condition := conditions[0].(map[string]interface{})

					target_lifecycle_rule.Condition = &storage.BucketLifecycleRuleCondition{}

					if v, ok := condition["age"]; ok {
						target_lifecycle_rule.Condition.Age = int64(v.(int))
					}

					if v, ok := condition["created_before"]; ok {
						target_lifecycle_rule.Condition.CreatedBefore = v.(string)
					}

					if v, ok := condition["matches_storage_class"]; ok {
						matches_storage_classes := v.([]interface{})

						target_matches_storage_classes := make([]string, 0, len(matches_storage_classes))

						for _, v := range matches_storage_classes {
							target_matches_storage_classes = append(target_matches_storage_classes, v.(string))
						}

						target_lifecycle_rule.Condition.MatchesStorageClass = target_matches_storage_classes
					}

					if v, ok := condition["num_newer_versions"]; ok {
						target_lifecycle_rule.Condition.NumNewerVersions = int64(v.(int))
					}
				} else {
					return fmt.Errorf("Exactly one condition is required")
				}
			}

			sb.Lifecycle.Rule = append(sb.Lifecycle.Rule, target_lifecycle_rule)
		}
	} else {
		sb.Lifecycle = &storage.BucketLifecycle{
			ForceSendFields: []string{"Rule"},
		}
	}

	return nil
}

// remove this on next major release of the provider.
func expandIamConfiguration(d TerraformResourceData) *storage.BucketIamConfiguration {
	// We are checking for a change because the last else block is only executed on Create.
	enabled := false
	if d.HasChange("bucket_policy_only") {
		enabled = d.Get("bucket_policy_only").(bool)
	} else if d.HasChange("uniform_bucket_level_access") {
		enabled = d.Get("uniform_bucket_level_access").(bool)
	} else {
		enabled = d.Get("bucket_policy_only").(bool) || d.Get("uniform_bucket_level_access").(bool)
	}

	return &storage.BucketIamConfiguration{
		ForceSendFields: []string{"UniformBucketLevelAccess"},
		UniformBucketLevelAccess: &storage.BucketIamConfigurationUniformBucketLevelAccess{
			Enabled:         enabled,
			ForceSendFields: []string{"Enabled"},
		},
	}
}

// Uncomment once the previous function is removed.
// func expandIamConfiguration(d TerraformResourceData) *storage.BucketIamConfiguration {
// 	return &storage.BucketIamConfiguration{
// 		ForceSendFields: []string{"UniformBucketLevelAccess"},
// 		UniformBucketLevelAccess: &storage.BucketIamConfigurationUniformBucketLevelAccess{
// 			Enabled:         d.Get("uniform_bucket_level_access").(bool),
// 			ForceSendFields: []string{"Enabled"},
// 		},
// 	}
// }
