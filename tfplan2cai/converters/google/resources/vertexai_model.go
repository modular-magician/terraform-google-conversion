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

import "reflect"

const VertexAIModelAssetType string = "{{region}}-aiplatform.googleapis.com/Model"

func resourceConverterVertexAIModel() ResourceConverter {
	return ResourceConverter{
		AssetType: VertexAIModelAssetType,
		Convert:   GetVertexAIModelCaiObject,
	}
}

func GetVertexAIModelCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//{{region}}-aiplatform.googleapis.com/projects/{{project}}/locations/{{region}}/models/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetVertexAIModelApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: VertexAIModelAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/{{region}}-aiplatform/v1/rest",
				DiscoveryName:        "Model",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetVertexAIModelApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandVertexAIModelName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	versionAliasesProp, err := expandVertexAIModelVersionAliases(d.Get("version_aliases"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("version_aliases"); !isEmptyValue(reflect.ValueOf(versionAliasesProp)) && (ok || !reflect.DeepEqual(v, versionAliasesProp)) {
		obj["versionAliases"] = versionAliasesProp
	}
	displayNameProp, err := expandVertexAIModelDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandVertexAIModelDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	versionDescriptionProp, err := expandVertexAIModelVersionDescription(d.Get("version_description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("version_description"); !isEmptyValue(reflect.ValueOf(versionDescriptionProp)) && (ok || !reflect.DeepEqual(v, versionDescriptionProp)) {
		obj["versionDescription"] = versionDescriptionProp
	}
	containerSpecProp, err := expandVertexAIModelContainerSpec(d.Get("container_spec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("container_spec"); !isEmptyValue(reflect.ValueOf(containerSpecProp)) && (ok || !reflect.DeepEqual(v, containerSpecProp)) {
		obj["containerSpec"] = containerSpecProp
	}
	artifactUriProp, err := expandVertexAIModelArtifactUri(d.Get("artifact_uri"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("artifact_uri"); !isEmptyValue(reflect.ValueOf(artifactUriProp)) && (ok || !reflect.DeepEqual(v, artifactUriProp)) {
		obj["artifactUri"] = artifactUriProp
	}
	labelsProp, err := expandVertexAIModelLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	encryptionSpecProp, err := expandVertexAIModelEncryptionSpec(d.Get("encryption_spec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("encryption_spec"); !isEmptyValue(reflect.ValueOf(encryptionSpecProp)) && (ok || !reflect.DeepEqual(v, encryptionSpecProp)) {
		obj["encryptionSpec"] = encryptionSpecProp
	}

	return resourceVertexAIModelEncoder(d, config, obj)
}

func resourceVertexAIModelEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	newObj := make(map[string]interface{})
	if name, ok := obj["name"]; ok {
		newObj["modelId"] = name
		delete(obj, "name")
	}
	newObj["model"] = obj
	return newObj, nil
}

func expandVertexAIModelName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIModelVersionAliases(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIModelDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIModelDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIModelVersionDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIModelContainerSpec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedImageUri, err := expandVertexAIModelContainerSpecImageUri(original["image_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImageUri); val.IsValid() && !isEmptyValue(val) {
		transformed["imageUri"] = transformedImageUri
	}

	transformedCommand, err := expandVertexAIModelContainerSpecCommand(original["command"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCommand); val.IsValid() && !isEmptyValue(val) {
		transformed["command"] = transformedCommand
	}

	transformedArgs, err := expandVertexAIModelContainerSpecArgs(original["args"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedArgs); val.IsValid() && !isEmptyValue(val) {
		transformed["args"] = transformedArgs
	}

	transformedEnv, err := expandVertexAIModelContainerSpecEnv(original["env"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnv); val.IsValid() && !isEmptyValue(val) {
		transformed["env"] = transformedEnv
	}

	transformedPorts, err := expandVertexAIModelContainerSpecPorts(original["ports"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPorts); val.IsValid() && !isEmptyValue(val) {
		transformed["ports"] = transformedPorts
	}

	transformedPredictRoute, err := expandVertexAIModelContainerSpecPredictRoute(original["predict_route"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPredictRoute); val.IsValid() && !isEmptyValue(val) {
		transformed["predictRoute"] = transformedPredictRoute
	}

	transformedHealthRoute, err := expandVertexAIModelContainerSpecHealthRoute(original["health_route"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHealthRoute); val.IsValid() && !isEmptyValue(val) {
		transformed["healthRoute"] = transformedHealthRoute
	}

	return transformed, nil
}

func expandVertexAIModelContainerSpecImageUri(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIModelContainerSpecCommand(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIModelContainerSpecArgs(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIModelContainerSpecEnv(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandVertexAIModelContainerSpecEnvName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedValue, err := expandVertexAIModelContainerSpecEnvValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !isEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandVertexAIModelContainerSpecEnvName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIModelContainerSpecEnvValue(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIModelContainerSpecPorts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedContainerPort, err := expandVertexAIModelContainerSpecPortsContainerPort(original["container_port"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedContainerPort); val.IsValid() && !isEmptyValue(val) {
			transformed["containerPort"] = transformedContainerPort
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandVertexAIModelContainerSpecPortsContainerPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIModelContainerSpecPredictRoute(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIModelContainerSpecHealthRoute(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIModelArtifactUri(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIModelLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandVertexAIModelEncryptionSpec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeyName, err := expandVertexAIModelEncryptionSpecKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !isEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	return transformed, nil
}

func expandVertexAIModelEncryptionSpecKmsKeyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
