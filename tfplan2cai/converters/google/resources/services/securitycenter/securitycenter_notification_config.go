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

package securitycenter

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const SecurityCenterNotificationConfigAssetType string = "securitycenter.googleapis.com/NotificationConfig"

func ResourceConverterSecurityCenterNotificationConfig() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: SecurityCenterNotificationConfigAssetType,
		Convert:   GetSecurityCenterNotificationConfigCaiObject,
	}
}

func GetSecurityCenterNotificationConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//securitycenter.googleapis.com/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetSecurityCenterNotificationConfigApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: SecurityCenterNotificationConfigAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/securitycenter/v1/rest",
				DiscoveryName:        "NotificationConfig",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetSecurityCenterNotificationConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandSecurityCenterNotificationConfigDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	pubsubTopicProp, err := expandSecurityCenterNotificationConfigPubsubTopic(d.Get("pubsub_topic"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("pubsub_topic"); !tpgresource.IsEmptyValue(reflect.ValueOf(pubsubTopicProp)) && (ok || !reflect.DeepEqual(v, pubsubTopicProp)) {
		obj["pubsubTopic"] = pubsubTopicProp
	}
	streamingConfigProp, err := expandSecurityCenterNotificationConfigStreamingConfig(d.Get("streaming_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("streaming_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(streamingConfigProp)) && (ok || !reflect.DeepEqual(v, streamingConfigProp)) {
		obj["streamingConfig"] = streamingConfigProp
	}

	return obj, nil
}

func expandSecurityCenterNotificationConfigDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterNotificationConfigPubsubTopic(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterNotificationConfigStreamingConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFilter, err := expandSecurityCenterNotificationConfigStreamingConfigFilter(original["filter"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFilter); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["filter"] = transformedFilter
	}

	return transformed, nil
}

func expandSecurityCenterNotificationConfigStreamingConfigFilter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
