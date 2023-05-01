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

package google

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"strings"

	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const SecretManagerSecretVersionAssetType string = "secretmanager.googleapis.com/SecretVersion"

func resourceConverterSecretManagerSecretVersion() ResourceConverter {
	return ResourceConverter{
		AssetType: SecretManagerSecretVersionAssetType,
		Convert:   GetSecretManagerSecretVersionCaiObject,
	}
}

func GetSecretManagerSecretVersionCaiObject(d TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//secretmanager.googleapis.com/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetSecretManagerSecretVersionApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: SecretManagerSecretVersionAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/secretmanager/v1/rest",
				DiscoveryName:        "SecretVersion",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetSecretManagerSecretVersionApiObject(d TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	stateProp, err := expandSecretManagerSecretVersionEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enabled"); !isEmptyValue(reflect.ValueOf(stateProp)) && (ok || !reflect.DeepEqual(v, stateProp)) {
		obj["state"] = stateProp
	}
	payloadProp, err := expandSecretManagerSecretVersionPayload(nil, d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("payload"); !isEmptyValue(reflect.ValueOf(payloadProp)) && (ok || !reflect.DeepEqual(v, payloadProp)) {
		obj["payload"] = payloadProp
	}

	return obj, nil
}

func expandSecretManagerSecretVersionEnabled(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	name := d.Get("name").(string)
	if name == "" {
		return "", nil
	}

	url, err := ReplaceVars(d, config, "{{SecretManagerBasePath}}{{name}}")
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

	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return nil, err
	}

	_, err = transport_tpg.SendRequest(config, "POST", project, url, userAgent, nil)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func expandSecretManagerSecretVersionPayload(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	transformedSecretData, err := expandSecretManagerSecretVersionPayloadSecretData(d.Get("secret_data"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSecretData); val.IsValid() && !isEmptyValue(val) {
		transformed["data"] = transformedSecretData
	}

	return transformed, nil
}

func expandSecretManagerSecretVersionPayloadSecretData(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	return base64.StdEncoding.EncodeToString([]byte(v.(string))), nil
}
