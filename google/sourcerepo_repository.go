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
	"bytes"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSourceRepoRepositoryPubSubConfigsHash(v interface{}) int {
	if v == nil {
		return 0
	}

	var buf bytes.Buffer
	m := v.(map[string]interface{})

	buf.WriteString(fmt.Sprintf("%s-", GetResourceNameFromSelfLink(m["topic"].(string))))
	buf.WriteString(fmt.Sprintf("%s-", m["message_format"].(string)))
	if v, ok := m["service_account_email"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	return hashcode.String(buf.String())
}

func GetSourceRepoRepositoryCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//sourcerepo.googleapis.com/projects/{{project}}/repos/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetSourceRepoRepositoryApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "sourcerepo.googleapis.com/Repository",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/sourcerepo/v1/rest",
				DiscoveryName:        "Repository",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetSourceRepoRepositoryApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandSourceRepoRepositoryName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	pubsubConfigsProp, err := expandSourceRepoRepositoryPubsubConfigs(d.Get("pubsub_configs"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("pubsub_configs"); !isEmptyValue(reflect.ValueOf(pubsubConfigsProp)) && (ok || !reflect.DeepEqual(v, pubsubConfigsProp)) {
		obj["pubsubConfigs"] = pubsubConfigsProp
	}

	return obj, nil
}

func expandSourceRepoRepositoryName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return replaceVars(d, config, "projects/{{project}}/repos/{{name}}")
}

func expandSourceRepoRepositoryPubsubConfigs(v interface{}, d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedMessageFormat, err := expandSourceRepoRepositoryPubsubConfigsMessageFormat(original["message_format"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMessageFormat); val.IsValid() && !isEmptyValue(val) {
			transformed["messageFormat"] = transformedMessageFormat
		}

		transformedServiceAccountEmail, err := expandSourceRepoRepositoryPubsubConfigsServiceAccountEmail(original["service_account_email"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedServiceAccountEmail); val.IsValid() && !isEmptyValue(val) {
			transformed["serviceAccountEmail"] = transformedServiceAccountEmail
		}

		transformedTopic, err := expandSourceRepoRepositoryPubsubConfigsTopic(original["topic"], d, config)
		if err != nil {
			return nil, err
		}
		m[transformedTopic] = transformed
	}
	return m, nil
}

func expandSourceRepoRepositoryPubsubConfigsMessageFormat(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandSourceRepoRepositoryPubsubConfigsServiceAccountEmail(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
