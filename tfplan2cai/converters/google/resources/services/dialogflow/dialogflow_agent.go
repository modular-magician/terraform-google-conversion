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

package dialogflow

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DialogflowAgentAssetType string = "dialogflow.googleapis.com/Agent"

func ResourceConverterDialogflowAgent() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DialogflowAgentAssetType,
		Convert:   GetDialogflowAgentCaiObject,
	}
}

func GetDialogflowAgentCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//dialogflow.googleapis.com/projects/{{project}}/agent")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDialogflowAgentApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DialogflowAgentAssetType,
			Resource: &cai.AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/dialogflow/v2/rest",
				DiscoveryName:        "Agent",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDialogflowAgentApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowAgentDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	defaultLanguageCodeProp, err := expandDialogflowAgentDefaultLanguageCode(d.Get("default_language_code"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("default_language_code"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultLanguageCodeProp)) && (ok || !reflect.DeepEqual(v, defaultLanguageCodeProp)) {
		obj["defaultLanguageCode"] = defaultLanguageCodeProp
	}
	supportedLanguageCodesProp, err := expandDialogflowAgentSupportedLanguageCodes(d.Get("supported_language_codes"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("supported_language_codes"); !tpgresource.IsEmptyValue(reflect.ValueOf(supportedLanguageCodesProp)) && (ok || !reflect.DeepEqual(v, supportedLanguageCodesProp)) {
		obj["supportedLanguageCodes"] = supportedLanguageCodesProp
	}
	timeZoneProp, err := expandDialogflowAgentTimeZone(d.Get("time_zone"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("time_zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(timeZoneProp)) && (ok || !reflect.DeepEqual(v, timeZoneProp)) {
		obj["timeZone"] = timeZoneProp
	}
	descriptionProp, err := expandDialogflowAgentDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	avatarUriProp, err := expandDialogflowAgentAvatarUri(d.Get("avatar_uri"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("avatar_uri"); !tpgresource.IsEmptyValue(reflect.ValueOf(avatarUriProp)) && (ok || !reflect.DeepEqual(v, avatarUriProp)) {
		obj["avatarUri"] = avatarUriProp
	}
	enableLoggingProp, err := expandDialogflowAgentEnableLogging(d.Get("enable_logging"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_logging"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableLoggingProp)) && (ok || !reflect.DeepEqual(v, enableLoggingProp)) {
		obj["enableLogging"] = enableLoggingProp
	}
	matchModeProp, err := expandDialogflowAgentMatchMode(d.Get("match_mode"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("match_mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(matchModeProp)) && (ok || !reflect.DeepEqual(v, matchModeProp)) {
		obj["matchMode"] = matchModeProp
	}
	classificationThresholdProp, err := expandDialogflowAgentClassificationThreshold(d.Get("classification_threshold"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("classification_threshold"); !tpgresource.IsEmptyValue(reflect.ValueOf(classificationThresholdProp)) && (ok || !reflect.DeepEqual(v, classificationThresholdProp)) {
		obj["classificationThreshold"] = classificationThresholdProp
	}
	apiVersionProp, err := expandDialogflowAgentApiVersion(d.Get("api_version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("api_version"); !tpgresource.IsEmptyValue(reflect.ValueOf(apiVersionProp)) && (ok || !reflect.DeepEqual(v, apiVersionProp)) {
		obj["apiVersion"] = apiVersionProp
	}
	tierProp, err := expandDialogflowAgentTier(d.Get("tier"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("tier"); !tpgresource.IsEmptyValue(reflect.ValueOf(tierProp)) && (ok || !reflect.DeepEqual(v, tierProp)) {
		obj["tier"] = tierProp
	}

	return obj, nil
}

func expandDialogflowAgentDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowAgentDefaultLanguageCode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowAgentSupportedLanguageCodes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowAgentTimeZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowAgentDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowAgentAvatarUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowAgentEnableLogging(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowAgentMatchMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowAgentClassificationThreshold(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowAgentApiVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowAgentTier(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
