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

package securitycenterv2

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const SecurityCenterV2OrganizationNotificationConfigAssetType string = "securitycenter.googleapis.com/OrganizationNotificationConfig"

func ResourceConverterSecurityCenterV2OrganizationNotificationConfig() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: SecurityCenterV2OrganizationNotificationConfigAssetType,
		Convert:   GetSecurityCenterV2OrganizationNotificationConfigCaiObject,
	}
}

func GetSecurityCenterV2OrganizationNotificationConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//securitycenter.googleapis.com/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetSecurityCenterV2OrganizationNotificationConfigApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: SecurityCenterV2OrganizationNotificationConfigAssetType,
			Resource: &cai.AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/securitycenter/v2/rest",
				DiscoveryName:        "OrganizationNotificationConfig",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetSecurityCenterV2OrganizationNotificationConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandSecurityCenterV2OrganizationNotificationConfigDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	pubsubTopicProp, err := expandSecurityCenterV2OrganizationNotificationConfigPubsubTopic(d.Get("pubsub_topic"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("pubsub_topic"); !tpgresource.IsEmptyValue(reflect.ValueOf(pubsubTopicProp)) && (ok || !reflect.DeepEqual(v, pubsubTopicProp)) {
		obj["pubsubTopic"] = pubsubTopicProp
	}
	streamingConfigProp, err := expandSecurityCenterV2OrganizationNotificationConfigStreamingConfig(d.Get("streaming_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("streaming_config"); ok || !reflect.DeepEqual(v, streamingConfigProp) {
		obj["streamingConfig"] = streamingConfigProp
	}

	return obj, nil
}

func expandSecurityCenterV2OrganizationNotificationConfigDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterV2OrganizationNotificationConfigPubsubTopic(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterV2OrganizationNotificationConfigStreamingConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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

	transformedFilter, err := expandSecurityCenterV2OrganizationNotificationConfigStreamingConfigFilter(original["filter"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["filter"] = transformedFilter
	}

	return transformed, nil
}

func expandSecurityCenterV2OrganizationNotificationConfigStreamingConfigFilter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
