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

package secretmanager

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"strings"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const SecretManagerSecretVersionAssetType string = "secretmanager.googleapis.com/SecretVersion"

func ResourceConverterSecretManagerSecretVersion() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: SecretManagerSecretVersionAssetType,
		Convert:   GetSecretManagerSecretVersionCaiObject,
	}
}

func GetSecretManagerSecretVersionCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//secretmanager.googleapis.com/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetSecretManagerSecretVersionApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: SecretManagerSecretVersionAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/secretmanager/v1/rest",
				DiscoveryName:        "SecretVersion",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetSecretManagerSecretVersionApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	stateProp, err := expandSecretManagerSecretVersionEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(stateProp)) && (ok || !reflect.DeepEqual(v, stateProp)) {
		obj["state"] = stateProp
	}
	payloadProp, err := expandSecretManagerSecretVersionPayload(nil, d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("payload"); !tpgresource.IsEmptyValue(reflect.ValueOf(payloadProp)) && (ok || !reflect.DeepEqual(v, payloadProp)) {
		obj["payload"] = payloadProp
	}

	return obj, nil
}

func expandSecretManagerSecretVersionEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	name := d.Get("name").(string)
	if name == "" {
		return "", nil
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecretManagerBasePath}}{{name}}")
	if err != nil {
		return nil, err
	}

	if v == true {
		url = fmt.Sprintf("%s:enable", url)
	} else {
		url = fmt.Sprintf("%s:disable", url)
	}

	parts := strings.Split(name, "/")
	project := parts[1]

	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return nil, err
	}

	_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   project,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func expandSecretManagerSecretVersionPayload(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	transformedSecretData, err := expandSecretManagerSecretVersionPayloadSecretData(d.Get("secret_data"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSecretData); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["data"] = transformedSecretData
	}

	return transformed, nil
}

func expandSecretManagerSecretVersionPayloadSecretData(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	if d.Get("is_secret_data_base64").(bool) {
		return v, nil
	}
	return base64.StdEncoding.EncodeToString([]byte(v.(string))), nil
}
