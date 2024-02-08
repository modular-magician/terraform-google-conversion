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

package clouddeploy

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ClouddeployCustomTargetTypeAssetType string = "clouddeploy.googleapis.com/CustomTargetType"

func ResourceConverterClouddeployCustomTargetType() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ClouddeployCustomTargetTypeAssetType,
		Convert:   GetClouddeployCustomTargetTypeCaiObject,
	}
}

func GetClouddeployCustomTargetTypeCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//clouddeploy.googleapis.com/projects/{{project}}/locations/{{location}}/customTargetTypes/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetClouddeployCustomTargetTypeApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ClouddeployCustomTargetTypeAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/clouddeploy/v1/rest",
				DiscoveryName:        "CustomTargetType",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetClouddeployCustomTargetTypeApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandClouddeployCustomTargetTypeDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	customActionsProp, err := expandClouddeployCustomTargetTypeCustomActions(d.Get("custom_actions"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("custom_actions"); !tpgresource.IsEmptyValue(reflect.ValueOf(customActionsProp)) && (ok || !reflect.DeepEqual(v, customActionsProp)) {
		obj["customActions"] = customActionsProp
	}
	annotationsProp, err := expandClouddeployCustomTargetTypeEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}
	labelsProp, err := expandClouddeployCustomTargetTypeEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandClouddeployCustomTargetTypeDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddeployCustomTargetTypeCustomActions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRenderAction, err := expandClouddeployCustomTargetTypeCustomActionsRenderAction(original["render_action"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRenderAction); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["renderAction"] = transformedRenderAction
	}

	transformedDeployAction, err := expandClouddeployCustomTargetTypeCustomActionsDeployAction(original["deploy_action"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDeployAction); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["deployAction"] = transformedDeployAction
	}

	transformedIncludeSkaffoldModules, err := expandClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModules(original["include_skaffold_modules"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIncludeSkaffoldModules); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["includeSkaffoldModules"] = transformedIncludeSkaffoldModules
	}

	return transformed, nil
}

func expandClouddeployCustomTargetTypeCustomActionsRenderAction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddeployCustomTargetTypeCustomActionsDeployAction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModules(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedConfigs, err := expandClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesConfigs(original["configs"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedConfigs); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["configs"] = transformedConfigs
		}

		transformedGit, err := expandClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGit(original["git"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGit); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["git"] = transformedGit
		}

		transformedGoogleCloudStorage, err := expandClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudStorage(original["google_cloud_storage"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGoogleCloudStorage); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["googleCloudStorage"] = transformedGoogleCloudStorage
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesConfigs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGit(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRepo, err := expandClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGitRepo(original["repo"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRepo); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["repo"] = transformedRepo
	}

	transformedPath, err := expandClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGitPath(original["path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["path"] = transformedPath
	}

	transformedRef, err := expandClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGitRef(original["ref"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRef); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ref"] = transformedRef
	}

	return transformed, nil
}

func expandClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGitRepo(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGitPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGitRef(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudStorage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSource, err := expandClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudStorageSource(original["source"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSource); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["source"] = transformedSource
	}

	transformedPath, err := expandClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudStoragePath(original["path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["path"] = transformedPath
	}

	return transformed, nil
}

func expandClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudStorageSource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudStoragePath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddeployCustomTargetTypeEffectiveAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandClouddeployCustomTargetTypeEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
