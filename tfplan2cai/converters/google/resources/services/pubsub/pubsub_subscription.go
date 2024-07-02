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

package pubsub

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func comparePubsubSubscriptionExpirationPolicy(_, old, new string, _ *schema.ResourceData) bool {
	trimmedNew := strings.TrimLeft(new, "0")
	trimmedOld := strings.TrimLeft(old, "0")
	if strings.Contains(trimmedNew, ".") {
		trimmedNew = strings.TrimRight(strings.TrimSuffix(trimmedNew, "s"), "0") + "s"
	}
	if strings.Contains(trimmedOld, ".") {
		trimmedOld = strings.TrimRight(strings.TrimSuffix(trimmedOld, "s"), "0") + "s"
	}
	return trimmedNew == trimmedOld
}

func IgnoreMissingKeyInMap(key string) schema.SchemaDiffSuppressFunc {
	return func(k, old, new string, d *schema.ResourceData) bool {
		log.Printf("[DEBUG] - suppressing diff %q with old %q, new %q", k, old, new)
		if strings.HasSuffix(k, ".%") {
			oldNum, err := strconv.Atoi(old)
			if err != nil {
				log.Printf("[ERROR] could not parse %q as number, no longer attempting diff suppress", old)
				return false
			}
			newNum, err := strconv.Atoi(new)
			if err != nil {
				log.Printf("[ERROR] could not parse %q as number, no longer attempting diff suppress", new)
				return false
			}
			return oldNum+1 == newNum
		} else if strings.HasSuffix(k, "."+key) {
			return old == ""
		}
		return false
	}
}

const PubsubSubscriptionAssetType string = "pubsub.googleapis.com/Subscription"

func ResourceConverterPubsubSubscription() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: PubsubSubscriptionAssetType,
		Convert:   GetPubsubSubscriptionCaiObject,
	}
}

func GetPubsubSubscriptionCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//pubsub.googleapis.com/projects/{{project}}/subscriptions/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetPubsubSubscriptionApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: PubsubSubscriptionAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/pubsub/v1/rest",
				DiscoveryName:        "Subscription",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetPubsubSubscriptionApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandPubsubSubscriptionName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	topicProp, err := expandPubsubSubscriptionTopic(d.Get("topic"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("topic"); !tpgresource.IsEmptyValue(reflect.ValueOf(topicProp)) && (ok || !reflect.DeepEqual(v, topicProp)) {
		obj["topic"] = topicProp
	}
	bigqueryConfigProp, err := expandPubsubSubscriptionBigqueryConfig(d.Get("bigquery_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("bigquery_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(bigqueryConfigProp)) && (ok || !reflect.DeepEqual(v, bigqueryConfigProp)) {
		obj["bigqueryConfig"] = bigqueryConfigProp
	}
	cloudStorageConfigProp, err := expandPubsubSubscriptionCloudStorageConfig(d.Get("cloud_storage_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cloud_storage_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(cloudStorageConfigProp)) && (ok || !reflect.DeepEqual(v, cloudStorageConfigProp)) {
		obj["cloudStorageConfig"] = cloudStorageConfigProp
	}
	pushConfigProp, err := expandPubsubSubscriptionPushConfig(d.Get("push_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("push_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(pushConfigProp)) && (ok || !reflect.DeepEqual(v, pushConfigProp)) {
		obj["pushConfig"] = pushConfigProp
	}
	ackDeadlineSecondsProp, err := expandPubsubSubscriptionAckDeadlineSeconds(d.Get("ack_deadline_seconds"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ack_deadline_seconds"); !tpgresource.IsEmptyValue(reflect.ValueOf(ackDeadlineSecondsProp)) && (ok || !reflect.DeepEqual(v, ackDeadlineSecondsProp)) {
		obj["ackDeadlineSeconds"] = ackDeadlineSecondsProp
	}
	messageRetentionDurationProp, err := expandPubsubSubscriptionMessageRetentionDuration(d.Get("message_retention_duration"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("message_retention_duration"); !tpgresource.IsEmptyValue(reflect.ValueOf(messageRetentionDurationProp)) && (ok || !reflect.DeepEqual(v, messageRetentionDurationProp)) {
		obj["messageRetentionDuration"] = messageRetentionDurationProp
	}
	retainAckedMessagesProp, err := expandPubsubSubscriptionRetainAckedMessages(d.Get("retain_acked_messages"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("retain_acked_messages"); !tpgresource.IsEmptyValue(reflect.ValueOf(retainAckedMessagesProp)) && (ok || !reflect.DeepEqual(v, retainAckedMessagesProp)) {
		obj["retainAckedMessages"] = retainAckedMessagesProp
	}
	expirationPolicyProp, err := expandPubsubSubscriptionExpirationPolicy(d.Get("expiration_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("expiration_policy"); ok || !reflect.DeepEqual(v, expirationPolicyProp) {
		obj["expirationPolicy"] = expirationPolicyProp
	}
	filterProp, err := expandPubsubSubscriptionFilter(d.Get("filter"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("filter"); !tpgresource.IsEmptyValue(reflect.ValueOf(filterProp)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}
	deadLetterPolicyProp, err := expandPubsubSubscriptionDeadLetterPolicy(d.Get("dead_letter_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dead_letter_policy"); ok || !reflect.DeepEqual(v, deadLetterPolicyProp) {
		obj["deadLetterPolicy"] = deadLetterPolicyProp
	}
	retryPolicyProp, err := expandPubsubSubscriptionRetryPolicy(d.Get("retry_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("retry_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(retryPolicyProp)) && (ok || !reflect.DeepEqual(v, retryPolicyProp)) {
		obj["retryPolicy"] = retryPolicyProp
	}
	enableMessageOrderingProp, err := expandPubsubSubscriptionEnableMessageOrdering(d.Get("enable_message_ordering"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_message_ordering"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableMessageOrderingProp)) && (ok || !reflect.DeepEqual(v, enableMessageOrderingProp)) {
		obj["enableMessageOrdering"] = enableMessageOrderingProp
	}
	enableExactlyOnceDeliveryProp, err := expandPubsubSubscriptionEnableExactlyOnceDelivery(d.Get("enable_exactly_once_delivery"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_exactly_once_delivery"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableExactlyOnceDeliveryProp)) && (ok || !reflect.DeepEqual(v, enableExactlyOnceDeliveryProp)) {
		obj["enableExactlyOnceDelivery"] = enableExactlyOnceDeliveryProp
	}
	labelsProp, err := expandPubsubSubscriptionEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return resourcePubsubSubscriptionEncoder(d, config, obj)
}

func resourcePubsubSubscriptionEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	delete(obj, "name")
	return obj, nil
}

func expandPubsubSubscriptionName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.ReplaceVars(d, config, "projects/{{project}}/subscriptions/{{name}}")
}

func expandPubsubSubscriptionTopic(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return "", err
	}

	topic := d.Get("topic").(string)

	re := regexp.MustCompile(`projects\/(.*)\/topics\/(.*)`)
	match := re.FindStringSubmatch(topic)
	if len(match) == 3 {
		return topic, nil
	} else {
		// If no full topic given, we expand it to a full topic on the same project
		fullTopic := fmt.Sprintf("projects/%s/topics/%s", project, topic)
		if err := d.Set("topic", fullTopic); err != nil {
			return nil, fmt.Errorf("Error setting topic: %s", err)
		}
		return fullTopic, nil
	}
}

func expandPubsubSubscriptionBigqueryConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTable, err := expandPubsubSubscriptionBigqueryConfigTable(original["table"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTable); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["table"] = transformedTable
	}

	transformedUseTopicSchema, err := expandPubsubSubscriptionBigqueryConfigUseTopicSchema(original["use_topic_schema"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUseTopicSchema); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["useTopicSchema"] = transformedUseTopicSchema
	}

	transformedUseTableSchema, err := expandPubsubSubscriptionBigqueryConfigUseTableSchema(original["use_table_schema"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUseTableSchema); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["useTableSchema"] = transformedUseTableSchema
	}

	transformedWriteMetadata, err := expandPubsubSubscriptionBigqueryConfigWriteMetadata(original["write_metadata"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWriteMetadata); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["writeMetadata"] = transformedWriteMetadata
	}

	transformedDropUnknownFields, err := expandPubsubSubscriptionBigqueryConfigDropUnknownFields(original["drop_unknown_fields"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDropUnknownFields); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dropUnknownFields"] = transformedDropUnknownFields
	}

	transformedServiceAccountEmail, err := expandPubsubSubscriptionBigqueryConfigServiceAccountEmail(original["service_account_email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAccountEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serviceAccountEmail"] = transformedServiceAccountEmail
	}

	return transformed, nil
}

func expandPubsubSubscriptionBigqueryConfigTable(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionBigqueryConfigUseTopicSchema(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionBigqueryConfigUseTableSchema(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionBigqueryConfigWriteMetadata(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionBigqueryConfigDropUnknownFields(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionBigqueryConfigServiceAccountEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionCloudStorageConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBucket, err := expandPubsubSubscriptionCloudStorageConfigBucket(original["bucket"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBucket); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["bucket"] = transformedBucket
	}

	transformedFilenamePrefix, err := expandPubsubSubscriptionCloudStorageConfigFilenamePrefix(original["filename_prefix"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFilenamePrefix); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["filenamePrefix"] = transformedFilenamePrefix
	}

	transformedFilenameSuffix, err := expandPubsubSubscriptionCloudStorageConfigFilenameSuffix(original["filename_suffix"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFilenameSuffix); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["filenameSuffix"] = transformedFilenameSuffix
	}

	transformedFilenameDatetimeFormat, err := expandPubsubSubscriptionCloudStorageConfigFilenameDatetimeFormat(original["filename_datetime_format"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFilenameDatetimeFormat); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["filenameDatetimeFormat"] = transformedFilenameDatetimeFormat
	}

	transformedMaxDuration, err := expandPubsubSubscriptionCloudStorageConfigMaxDuration(original["max_duration"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxDuration); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxDuration"] = transformedMaxDuration
	}

	transformedMaxBytes, err := expandPubsubSubscriptionCloudStorageConfigMaxBytes(original["max_bytes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxBytes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxBytes"] = transformedMaxBytes
	}

	transformedState, err := expandPubsubSubscriptionCloudStorageConfigState(original["state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["state"] = transformedState
	}

	transformedAvroConfig, err := expandPubsubSubscriptionCloudStorageConfigAvroConfig(original["avro_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAvroConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["avroConfig"] = transformedAvroConfig
	}

	transformedServiceAccountEmail, err := expandPubsubSubscriptionCloudStorageConfigServiceAccountEmail(original["service_account_email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAccountEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serviceAccountEmail"] = transformedServiceAccountEmail
	}

	return transformed, nil
}

func expandPubsubSubscriptionCloudStorageConfigBucket(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionCloudStorageConfigFilenamePrefix(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionCloudStorageConfigFilenameSuffix(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionCloudStorageConfigFilenameDatetimeFormat(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionCloudStorageConfigMaxDuration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionCloudStorageConfigMaxBytes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionCloudStorageConfigState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionCloudStorageConfigAvroConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedWriteMetadata, err := expandPubsubSubscriptionCloudStorageConfigAvroConfigWriteMetadata(original["write_metadata"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWriteMetadata); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["writeMetadata"] = transformedWriteMetadata
	}

	return transformed, nil
}

func expandPubsubSubscriptionCloudStorageConfigAvroConfigWriteMetadata(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionCloudStorageConfigServiceAccountEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionPushConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedOidcToken, err := expandPubsubSubscriptionPushConfigOidcToken(original["oidc_token"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOidcToken); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["oidcToken"] = transformedOidcToken
	}

	transformedPushEndpoint, err := expandPubsubSubscriptionPushConfigPushEndpoint(original["push_endpoint"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPushEndpoint); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pushEndpoint"] = transformedPushEndpoint
	}

	transformedAttributes, err := expandPubsubSubscriptionPushConfigAttributes(original["attributes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAttributes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["attributes"] = transformedAttributes
	}

	transformedNoWrapper, err := expandPubsubSubscriptionPushConfigNoWrapper(original["no_wrapper"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNoWrapper); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["noWrapper"] = transformedNoWrapper
	}

	return transformed, nil
}

func expandPubsubSubscriptionPushConfigOidcToken(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedServiceAccountEmail, err := expandPubsubSubscriptionPushConfigOidcTokenServiceAccountEmail(original["service_account_email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAccountEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serviceAccountEmail"] = transformedServiceAccountEmail
	}

	transformedAudience, err := expandPubsubSubscriptionPushConfigOidcTokenAudience(original["audience"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAudience); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["audience"] = transformedAudience
	}

	return transformed, nil
}

func expandPubsubSubscriptionPushConfigOidcTokenServiceAccountEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionPushConfigOidcTokenAudience(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionPushConfigPushEndpoint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionPushConfigAttributes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandPubsubSubscriptionPushConfigNoWrapper(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedWriteMetadata, err := expandPubsubSubscriptionPushConfigNoWrapperWriteMetadata(original["write_metadata"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["writeMetadata"] = transformedWriteMetadata
	}

	return transformed, nil
}

func expandPubsubSubscriptionPushConfigNoWrapperWriteMetadata(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionAckDeadlineSeconds(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionMessageRetentionDuration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionRetainAckedMessages(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionExpirationPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTtl, err := expandPubsubSubscriptionExpirationPolicyTtl(original["ttl"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTtl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ttl"] = transformedTtl
	}

	return transformed, nil
}

func expandPubsubSubscriptionExpirationPolicyTtl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionFilter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionDeadLetterPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDeadLetterTopic, err := expandPubsubSubscriptionDeadLetterPolicyDeadLetterTopic(original["dead_letter_topic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDeadLetterTopic); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["deadLetterTopic"] = transformedDeadLetterTopic
	}

	transformedMaxDeliveryAttempts, err := expandPubsubSubscriptionDeadLetterPolicyMaxDeliveryAttempts(original["max_delivery_attempts"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxDeliveryAttempts); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxDeliveryAttempts"] = transformedMaxDeliveryAttempts
	}

	return transformed, nil
}

func expandPubsubSubscriptionDeadLetterPolicyDeadLetterTopic(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionDeadLetterPolicyMaxDeliveryAttempts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionRetryPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMinimumBackoff, err := expandPubsubSubscriptionRetryPolicyMinimumBackoff(original["minimum_backoff"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinimumBackoff); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minimumBackoff"] = transformedMinimumBackoff
	}

	transformedMaximumBackoff, err := expandPubsubSubscriptionRetryPolicyMaximumBackoff(original["maximum_backoff"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaximumBackoff); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maximumBackoff"] = transformedMaximumBackoff
	}

	return transformed, nil
}

func expandPubsubSubscriptionRetryPolicyMinimumBackoff(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionRetryPolicyMaximumBackoff(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionEnableMessageOrdering(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionEnableExactlyOnceDelivery(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
