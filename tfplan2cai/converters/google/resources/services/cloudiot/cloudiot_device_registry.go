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

package cloudiot

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func expandCloudIotDeviceRegistryHTTPConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHTTPEnabledState, err := expandCloudIotDeviceRegistryHTTPEnabledState(original["http_enabled_state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHTTPEnabledState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["httpEnabledState"] = transformedHTTPEnabledState
	}

	return transformed, nil
}

func expandCloudIotDeviceRegistryHTTPEnabledState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudIotDeviceRegistryMqttConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMqttEnabledState, err := expandCloudIotDeviceRegistryMqttEnabledState(original["mqtt_enabled_state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMqttEnabledState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mqttEnabledState"] = transformedMqttEnabledState
	}

	return transformed, nil
}

func expandCloudIotDeviceRegistryMqttEnabledState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudIotDeviceRegistryStateNotificationConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPubsubTopicName, err := expandCloudIotDeviceRegistryStateNotificationConfigPubsubTopicName(original["pubsub_topic_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPubsubTopicName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pubsubTopicName"] = transformedPubsubTopicName
	}

	return transformed, nil
}

func expandCloudIotDeviceRegistryStateNotificationConfigPubsubTopicName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudIotDeviceRegistryCredentials(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))

	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedPublicKeyCertificate, err := expandCloudIotDeviceRegistryCredentialsPublicKeyCertificate(original["public_key_certificate"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPublicKeyCertificate); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["publicKeyCertificate"] = transformedPublicKeyCertificate
		}

		req = append(req, transformed)
	}

	return req, nil
}

func expandCloudIotDeviceRegistryCredentialsPublicKeyCertificate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFormat, err := expandCloudIotDeviceRegistryPublicKeyCertificateFormat(original["format"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFormat); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["format"] = transformedFormat
	}

	transformedCertificate, err := expandCloudIotDeviceRegistryPublicKeyCertificateCertificate(original["certificate"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCertificate); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["certificate"] = transformedCertificate
	}

	return transformed, nil
}

func expandCloudIotDeviceRegistryPublicKeyCertificateFormat(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudIotDeviceRegistryPublicKeyCertificateCertificate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func flattenCloudIotDeviceRegistryCredentials(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	log.Printf("[DEBUG] Flattening device resitry credentials: %q", d.Id())
	if v == nil {
		log.Printf("[DEBUG] The credentials array is nil: %q", d.Id())
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		log.Printf("[DEBUG] Original credential: %+v", original)
		if len(original) < 1 {
			log.Printf("[DEBUG] Excluding empty credential that the API returned. %q", d.Id())
			continue
		}
		log.Printf("[DEBUG] Credentials array before appending a new credential: %+v", transformed)
		transformed = append(transformed, map[string]interface{}{
			"public_key_certificate": flattenCloudIotDeviceRegistryCredentialsPublicKeyCertificate(original["publicKeyCertificate"], d, config),
		})
		log.Printf("[DEBUG] Credentials array after appending a new credential: %+v", transformed)
	}
	return transformed
}

func flattenCloudIotDeviceRegistryCredentialsPublicKeyCertificate(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	log.Printf("[DEBUG] Flattening device resitry credentials public key certificate: %q", d.Id())
	if v == nil {
		log.Printf("[DEBUG] The public key certificate is nil: %q", d.Id())
		return v
	}

	original := v.(map[string]interface{})
	log.Printf("[DEBUG] Original public key certificate: %+v", original)
	transformed := make(map[string]interface{})

	transformedPublicKeyCertificateFormat := flattenCloudIotDeviceRegistryPublicKeyCertificateFormat(original["format"], d, config)
	transformed["format"] = transformedPublicKeyCertificateFormat

	transformedPublicKeyCertificateCertificate := flattenCloudIotDeviceRegistryPublicKeyCertificateCertificate(original["certificate"], d, config)
	transformed["certificate"] = transformedPublicKeyCertificateCertificate

	log.Printf("[DEBUG] Transformed public key certificate: %+v", transformed)

	return transformed
}

func flattenCloudIotDeviceRegistryPublicKeyCertificateFormat(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudIotDeviceRegistryPublicKeyCertificateCertificate(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudIotDeviceRegistryHTTPConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHTTPEnabledState := flattenCloudIotDeviceRegistryHTTPConfigHTTPEnabledState(original["httpEnabledState"], d, config)
	transformed["http_enabled_state"] = transformedHTTPEnabledState

	return transformed
}

func flattenCloudIotDeviceRegistryHTTPConfigHTTPEnabledState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudIotDeviceRegistryMqttConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMqttEnabledState := flattenCloudIotDeviceRegistryMqttConfigMqttEnabledState(original["mqttEnabledState"], d, config)
	transformed["mqtt_enabled_state"] = transformedMqttEnabledState

	return transformed
}

func flattenCloudIotDeviceRegistryMqttConfigMqttEnabledState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudIotDeviceRegistryStateNotificationConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	log.Printf("[DEBUG] Flattening state notification config: %+v", v)
	if v == nil {
		return v
	}

	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPubsubTopicName := flattenCloudIotDeviceRegistryStateNotificationConfigPubsubTopicName(original["pubsubTopicName"], d, config)
	if val := reflect.ValueOf(transformedPubsubTopicName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		log.Printf("[DEBUG] pubsub topic name is not null: %v", d.Get("pubsub_topic_name"))
		transformed["pubsub_topic_name"] = transformedPubsubTopicName
	}

	return transformed
}

func flattenCloudIotDeviceRegistryStateNotificationConfigPubsubTopicName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func ValidateCloudIotDeviceRegistryID(v interface{}, k string) (warnings []string, errors []error) {
	value := v.(string)
	if strings.HasPrefix(value, "goog") {
		errors = append(errors, fmt.Errorf(
			"%q (%q) can not start with \"goog\"", k, value))
	}
	if !regexp.MustCompile(verify.CloudIoTIdRegex).MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"%q (%q) doesn't match regexp %q", k, value, verify.CloudIoTIdRegex))
	}
	return
}

func validateCloudIotDeviceRegistrySubfolderMatch(v interface{}, k string) (warnings []string, errors []error) {
	value := v.(string)
	if strings.HasPrefix(value, "/") {
		errors = append(errors, fmt.Errorf(
			"%q (%q) can not start with '/'", k, value))
	}
	return
}

const CloudIotDeviceRegistryAssetType string = "cloudiot.googleapis.com/DeviceRegistry"

func ResourceConverterCloudIotDeviceRegistry() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: CloudIotDeviceRegistryAssetType,
		Convert:   GetCloudIotDeviceRegistryCaiObject,
	}
}

func GetCloudIotDeviceRegistryCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//cloudiot.googleapis.com/projects/{{project}}/locations/{{region}}/registries/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetCloudIotDeviceRegistryApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: CloudIotDeviceRegistryAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudiot/v1/rest",
				DiscoveryName:        "DeviceRegistry",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetCloudIotDeviceRegistryApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	idProp, err := expandCloudIotDeviceRegistryName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(idProp)) && (ok || !reflect.DeepEqual(v, idProp)) {
		obj["id"] = idProp
	}
	eventNotificationConfigsProp, err := expandCloudIotDeviceRegistryEventNotificationConfigs(d.Get("event_notification_configs"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("event_notification_configs"); !tpgresource.IsEmptyValue(reflect.ValueOf(eventNotificationConfigsProp)) && (ok || !reflect.DeepEqual(v, eventNotificationConfigsProp)) {
		obj["eventNotificationConfigs"] = eventNotificationConfigsProp
	}
	logLevelProp, err := expandCloudIotDeviceRegistryLogLevel(d.Get("log_level"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("log_level"); !tpgresource.IsEmptyValue(reflect.ValueOf(logLevelProp)) && (ok || !reflect.DeepEqual(v, logLevelProp)) {
		obj["logLevel"] = logLevelProp
	}

	return resourceCloudIotDeviceRegistryEncoder(d, config, obj)
}

func resourceCloudIotDeviceRegistryEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*transport_tpg.Config)

	log.Printf("[DEBUG] Resource data before encoding extra schema entries %q: %#v", d.Id(), obj)

	log.Printf("[DEBUG] Encoding state notification config: %q", d.Id())
	stateNotificationConfigProp, err := expandCloudIotDeviceRegistryStateNotificationConfig(d.Get("state_notification_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("state_notification_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(stateNotificationConfigProp)) && (ok || !reflect.DeepEqual(v, stateNotificationConfigProp)) {
		log.Printf("[DEBUG] Encoding %q. Setting stateNotificationConfig: %#v", d.Id(), stateNotificationConfigProp)
		obj["stateNotificationConfig"] = stateNotificationConfigProp
	}

	log.Printf("[DEBUG] Encoding HTTP config: %q", d.Id())
	httpConfigProp, err := expandCloudIotDeviceRegistryHTTPConfig(d.Get("http_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("http_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(httpConfigProp)) && (ok || !reflect.DeepEqual(v, httpConfigProp)) {
		log.Printf("[DEBUG] Encoding %q. Setting httpConfig: %#v", d.Id(), httpConfigProp)
		obj["httpConfig"] = httpConfigProp
	}

	log.Printf("[DEBUG] Encoding MQTT config: %q", d.Id())
	mqttConfigProp, err := expandCloudIotDeviceRegistryMqttConfig(d.Get("mqtt_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("mqtt_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(mqttConfigProp)) && (ok || !reflect.DeepEqual(v, mqttConfigProp)) {
		log.Printf("[DEBUG] Encoding %q. Setting mqttConfig: %#v", d.Id(), mqttConfigProp)
		obj["mqttConfig"] = mqttConfigProp
	}

	log.Printf("[DEBUG] Encoding credentials: %q", d.Id())
	credentialsProp, err := expandCloudIotDeviceRegistryCredentials(d.Get("credentials"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("credentials"); !tpgresource.IsEmptyValue(reflect.ValueOf(credentialsProp)) && (ok || !reflect.DeepEqual(v, credentialsProp)) {
		log.Printf("[DEBUG] Encoding %q. Setting credentials: %#v", d.Id(), credentialsProp)
		obj["credentials"] = credentialsProp
	}

	log.Printf("[DEBUG] Resource data after encoding extra schema entries %q: %#v", d.Id(), obj)

	return obj, nil
}

func expandCloudIotDeviceRegistryName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudIotDeviceRegistryEventNotificationConfigs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedSubfolderMatches, err := expandCloudIotDeviceRegistryEventNotificationConfigsSubfolderMatches(original["subfolder_matches"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSubfolderMatches); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["subfolderMatches"] = transformedSubfolderMatches
		}

		transformedPubsubTopicName, err := expandCloudIotDeviceRegistryEventNotificationConfigsPubsubTopicName(original["pubsub_topic_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPubsubTopicName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["pubsubTopicName"] = transformedPubsubTopicName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudIotDeviceRegistryEventNotificationConfigsSubfolderMatches(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudIotDeviceRegistryEventNotificationConfigsPubsubTopicName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudIotDeviceRegistryLogLevel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
