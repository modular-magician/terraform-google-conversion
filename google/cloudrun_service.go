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
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func revisionNameCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
	autogen := diff.Get("autogenerate_revision_name").(bool)
	if autogen && diff.HasChange("template.0.metadata.0.name") {
		return fmt.Errorf("google_cloud_run_service: `template.metadata.name` cannot be set while `autogenerate_revision_name` is true. Please remove the field or set `autogenerate_revision_name` to false.")
	}
	return nil
}

var cloudRunGoogleProvidedAnnotations = regexp.MustCompile("serving\\.knative\\.dev/(?:(?:creator)|(?:lastModifier))$")

func cloudrunAnnotationDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	// Suppress diffs for the annotations provided by Google
	if cloudRunGoogleProvidedAnnotations.MatchString(k) && new == "" {
		return true
	}

	// Let diff be determined by annotations (above)
	if strings.Contains(k, "annotations.%") {
		return true
	}

	// For other keys, don't suppress diff.
	return false
}

func GetCloudRunServiceCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//cloudrun.googleapis.com/apis/serving.knative.dev/namespaces/{{project}}/services/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetCloudRunServiceApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "cloudrun.googleapis.com/Service",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudrun/v1/rest",
				DiscoveryName:        "Service",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetCloudRunServiceApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	specProp, err := expandCloudRunServiceSpec(nil, d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("spec"); !isEmptyValue(reflect.ValueOf(specProp)) && (ok || !reflect.DeepEqual(v, specProp)) {
		obj["spec"] = specProp
	}
	metadataProp, err := expandCloudRunServiceMetadata(d.Get("metadata"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("metadata"); !isEmptyValue(reflect.ValueOf(metadataProp)) && (ok || !reflect.DeepEqual(v, metadataProp)) {
		obj["metadata"] = metadataProp
	}

	return resourceCloudRunServiceEncoder(d, config, obj)
}

func resourceCloudRunServiceEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	name := d.Get("name").(string)
	if obj["metadata"] == nil {
		obj["metadata"] = make(map[string]interface{})
	}
	metadata := obj["metadata"].(map[string]interface{})
	metadata["name"] = name

	// The only acceptable version/kind right now
	obj["apiVersion"] = "serving.knative.dev/v1"
	obj["kind"] = "Service"
	return obj, nil
}

func expandCloudRunServiceSpec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	transformedTraffic, err := expandCloudRunServiceSpecTraffic(d.Get("traffic"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTraffic); val.IsValid() && !isEmptyValue(val) {
		transformed["traffic"] = transformedTraffic
	}

	transformedTemplate, err := expandCloudRunServiceSpecTemplate(d.Get("template"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTemplate); val.IsValid() && !isEmptyValue(val) {
		transformed["template"] = transformedTemplate
	}

	return transformed, nil
}

func expandCloudRunServiceSpecTraffic(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedRevisionName, err := expandCloudRunServiceSpecTrafficRevisionName(original["revision_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRevisionName); val.IsValid() && !isEmptyValue(val) {
			transformed["revisionName"] = transformedRevisionName
		}

		transformedPercent, err := expandCloudRunServiceSpecTrafficPercent(original["percent"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPercent); val.IsValid() && !isEmptyValue(val) {
			transformed["percent"] = transformedPercent
		}

		transformedLatestRevision, err := expandCloudRunServiceSpecTrafficLatestRevision(original["latest_revision"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedLatestRevision); val.IsValid() && !isEmptyValue(val) {
			transformed["latestRevision"] = transformedLatestRevision
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudRunServiceSpecTrafficRevisionName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceSpecTrafficPercent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceSpecTrafficLatestRevision(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceSpecTemplate(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMetadata, err := expandCloudRunServiceSpecTemplateMetadata(original["metadata"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMetadata); val.IsValid() && !isEmptyValue(val) {
		transformed["metadata"] = transformedMetadata
	}

	transformedSpec, err := expandCloudRunServiceSpecTemplateSpec(original["spec"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSpec); val.IsValid() && !isEmptyValue(val) {
		transformed["spec"] = transformedSpec
	}

	return transformed, nil
}

func expandCloudRunServiceSpecTemplateMetadata(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLabels, err := expandCloudRunServiceSpecTemplateMetadataLabels(original["labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !isEmptyValue(val) {
		transformed["labels"] = transformedLabels
	}

	transformedGeneration, err := expandCloudRunServiceSpecTemplateMetadataGeneration(original["generation"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGeneration); val.IsValid() && !isEmptyValue(val) {
		transformed["generation"] = transformedGeneration
	}

	transformedResourceVersion, err := expandCloudRunServiceSpecTemplateMetadataResourceVersion(original["resource_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResourceVersion); val.IsValid() && !isEmptyValue(val) {
		transformed["resourceVersion"] = transformedResourceVersion
	}

	transformedSelfLink, err := expandCloudRunServiceSpecTemplateMetadataSelfLink(original["self_link"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSelfLink); val.IsValid() && !isEmptyValue(val) {
		transformed["selfLink"] = transformedSelfLink
	}

	transformedUid, err := expandCloudRunServiceSpecTemplateMetadataUid(original["uid"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUid); val.IsValid() && !isEmptyValue(val) {
		transformed["uid"] = transformedUid
	}

	transformedNamespace, err := expandCloudRunServiceSpecTemplateMetadataNamespace(original["namespace"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNamespace); val.IsValid() && !isEmptyValue(val) {
		transformed["namespace"] = transformedNamespace
	}

	transformedAnnotations, err := expandCloudRunServiceSpecTemplateMetadataAnnotations(original["annotations"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAnnotations); val.IsValid() && !isEmptyValue(val) {
		transformed["annotations"] = transformedAnnotations
	}

	transformedName, err := expandCloudRunServiceSpecTemplateMetadataName(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
		transformed["name"] = transformedName
	}

	return transformed, nil
}

func expandCloudRunServiceSpecTemplateMetadataLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudRunServiceSpecTemplateMetadataGeneration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceSpecTemplateMetadataResourceVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceSpecTemplateMetadataSelfLink(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceSpecTemplateMetadataUid(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

// If the property hasn't been explicitly set in config use the project defined by the provider or env.
func expandCloudRunServiceSpecTemplateMetadataNamespace(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	if v == nil {
		project, err := getProject(d, config)
		if err != nil {
			return project, nil
		}
	}
	return v, nil
}

func expandCloudRunServiceSpecTemplateMetadataAnnotations(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudRunServiceSpecTemplateMetadataName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	if d.Get("autogenerate_revision_name") == true {
		return nil, nil
	}
	return v, nil
}

func expandCloudRunServiceSpecTemplateSpec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedContainers, err := expandCloudRunServiceSpecTemplateSpecContainers(original["containers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedContainers); val.IsValid() && !isEmptyValue(val) {
		transformed["containers"] = transformedContainers
	}

	transformedContainerConcurrency, err := expandCloudRunServiceSpecTemplateSpecContainerConcurrency(original["container_concurrency"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedContainerConcurrency); val.IsValid() && !isEmptyValue(val) {
		transformed["containerConcurrency"] = transformedContainerConcurrency
	}

	transformedTimeoutSeconds, err := expandCloudRunServiceSpecTemplateSpecTimeoutSeconds(original["timeout_seconds"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTimeoutSeconds); val.IsValid() && !isEmptyValue(val) {
		transformed["timeoutSeconds"] = transformedTimeoutSeconds
	}

	transformedServiceAccountName, err := expandCloudRunServiceSpecTemplateSpecServiceAccountName(original["service_account_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAccountName); val.IsValid() && !isEmptyValue(val) {
		transformed["serviceAccountName"] = transformedServiceAccountName
	}

	transformedServingState, err := expandCloudRunServiceSpecTemplateSpecServingState(original["serving_state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServingState); val.IsValid() && !isEmptyValue(val) {
		transformed["servingState"] = transformedServingState
	}

	return transformed, nil
}

func expandCloudRunServiceSpecTemplateSpecContainers(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedWorkingDir, err := expandCloudRunServiceSpecTemplateSpecContainersWorkingDir(original["working_dir"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedWorkingDir); val.IsValid() && !isEmptyValue(val) {
			transformed["workingDir"] = transformedWorkingDir
		}

		transformedArgs, err := expandCloudRunServiceSpecTemplateSpecContainersArgs(original["args"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedArgs); val.IsValid() && !isEmptyValue(val) {
			transformed["args"] = transformedArgs
		}

		transformedEnvFrom, err := expandCloudRunServiceSpecTemplateSpecContainersEnvFrom(original["env_from"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEnvFrom); val.IsValid() && !isEmptyValue(val) {
			transformed["envFrom"] = transformedEnvFrom
		}

		transformedImage, err := expandCloudRunServiceSpecTemplateSpecContainersImage(original["image"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedImage); val.IsValid() && !isEmptyValue(val) {
			transformed["image"] = transformedImage
		}

		transformedCommand, err := expandCloudRunServiceSpecTemplateSpecContainersCommand(original["command"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCommand); val.IsValid() && !isEmptyValue(val) {
			transformed["command"] = transformedCommand
		}

		transformedEnv, err := expandCloudRunServiceSpecTemplateSpecContainersEnv(original["env"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEnv); val.IsValid() && !isEmptyValue(val) {
			transformed["env"] = transformedEnv
		}

		transformedPorts, err := expandCloudRunServiceSpecTemplateSpecContainersPorts(original["ports"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPorts); val.IsValid() && !isEmptyValue(val) {
			transformed["ports"] = transformedPorts
		}

		transformedResources, err := expandCloudRunServiceSpecTemplateSpecContainersResources(original["resources"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedResources); val.IsValid() && !isEmptyValue(val) {
			transformed["resources"] = transformedResources
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudRunServiceSpecTemplateSpecContainersWorkingDir(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceSpecTemplateSpecContainersArgs(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceSpecTemplateSpecContainersEnvFrom(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedPrefix, err := expandCloudRunServiceSpecTemplateSpecContainersEnvFromPrefix(original["prefix"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPrefix); val.IsValid() && !isEmptyValue(val) {
			transformed["prefix"] = transformedPrefix
		}

		transformedConfigMapRef, err := expandCloudRunServiceSpecTemplateSpecContainersEnvFromConfigMapRef(original["config_map_ref"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedConfigMapRef); val.IsValid() && !isEmptyValue(val) {
			transformed["configMapRef"] = transformedConfigMapRef
		}

		transformedSecretRef, err := expandCloudRunServiceSpecTemplateSpecContainersEnvFromSecretRef(original["secret_ref"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSecretRef); val.IsValid() && !isEmptyValue(val) {
			transformed["secretRef"] = transformedSecretRef
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudRunServiceSpecTemplateSpecContainersEnvFromPrefix(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceSpecTemplateSpecContainersEnvFromConfigMapRef(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedOptional, err := expandCloudRunServiceSpecTemplateSpecContainersEnvFromConfigMapRefOptional(original["optional"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOptional); val.IsValid() && !isEmptyValue(val) {
		transformed["optional"] = transformedOptional
	}

	transformedLocalObjectReference, err := expandCloudRunServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference(original["local_object_reference"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocalObjectReference); val.IsValid() && !isEmptyValue(val) {
		transformed["localObjectReference"] = transformedLocalObjectReference
	}

	return transformed, nil
}

func expandCloudRunServiceSpecTemplateSpecContainersEnvFromConfigMapRefOptional(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedName, err := expandCloudRunServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReferenceName(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
		transformed["name"] = transformedName
	}

	return transformed, nil
}

func expandCloudRunServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReferenceName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceSpecTemplateSpecContainersEnvFromSecretRef(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLocalObjectReference, err := expandCloudRunServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference(original["local_object_reference"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocalObjectReference); val.IsValid() && !isEmptyValue(val) {
		transformed["localObjectReference"] = transformedLocalObjectReference
	}

	transformedOptional, err := expandCloudRunServiceSpecTemplateSpecContainersEnvFromSecretRefOptional(original["optional"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOptional); val.IsValid() && !isEmptyValue(val) {
		transformed["optional"] = transformedOptional
	}

	return transformed, nil
}

func expandCloudRunServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedName, err := expandCloudRunServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReferenceName(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
		transformed["name"] = transformedName
	}

	return transformed, nil
}

func expandCloudRunServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReferenceName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceSpecTemplateSpecContainersEnvFromSecretRefOptional(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceSpecTemplateSpecContainersImage(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceSpecTemplateSpecContainersCommand(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceSpecTemplateSpecContainersEnv(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandCloudRunServiceSpecTemplateSpecContainersEnvName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedValue, err := expandCloudRunServiceSpecTemplateSpecContainersEnvValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !isEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudRunServiceSpecTemplateSpecContainersEnvName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceSpecTemplateSpecContainersEnvValue(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceSpecTemplateSpecContainersPorts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandCloudRunServiceSpecTemplateSpecContainersPortsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedProtocol, err := expandCloudRunServiceSpecTemplateSpecContainersPortsProtocol(original["protocol"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedProtocol); val.IsValid() && !isEmptyValue(val) {
			transformed["protocol"] = transformedProtocol
		}

		transformedContainerPort, err := expandCloudRunServiceSpecTemplateSpecContainersPortsContainerPort(original["container_port"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedContainerPort); val.IsValid() && !isEmptyValue(val) {
			transformed["containerPort"] = transformedContainerPort
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudRunServiceSpecTemplateSpecContainersPortsName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceSpecTemplateSpecContainersPortsProtocol(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceSpecTemplateSpecContainersPortsContainerPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceSpecTemplateSpecContainersResources(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLimits, err := expandCloudRunServiceSpecTemplateSpecContainersResourcesLimits(original["limits"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLimits); val.IsValid() && !isEmptyValue(val) {
		transformed["limits"] = transformedLimits
	}

	transformedRequests, err := expandCloudRunServiceSpecTemplateSpecContainersResourcesRequests(original["requests"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequests); val.IsValid() && !isEmptyValue(val) {
		transformed["requests"] = transformedRequests
	}

	return transformed, nil
}

func expandCloudRunServiceSpecTemplateSpecContainersResourcesLimits(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudRunServiceSpecTemplateSpecContainersResourcesRequests(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudRunServiceSpecTemplateSpecContainerConcurrency(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceSpecTemplateSpecTimeoutSeconds(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceSpecTemplateSpecServiceAccountName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceSpecTemplateSpecServingState(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceMetadata(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLabels, err := expandCloudRunServiceMetadataLabels(original["labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !isEmptyValue(val) {
		transformed["labels"] = transformedLabels
	}

	transformedGeneration, err := expandCloudRunServiceMetadataGeneration(original["generation"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGeneration); val.IsValid() && !isEmptyValue(val) {
		transformed["generation"] = transformedGeneration
	}

	transformedResourceVersion, err := expandCloudRunServiceMetadataResourceVersion(original["resource_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResourceVersion); val.IsValid() && !isEmptyValue(val) {
		transformed["resourceVersion"] = transformedResourceVersion
	}

	transformedSelfLink, err := expandCloudRunServiceMetadataSelfLink(original["self_link"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSelfLink); val.IsValid() && !isEmptyValue(val) {
		transformed["selfLink"] = transformedSelfLink
	}

	transformedUid, err := expandCloudRunServiceMetadataUid(original["uid"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUid); val.IsValid() && !isEmptyValue(val) {
		transformed["uid"] = transformedUid
	}

	transformedNamespace, err := expandCloudRunServiceMetadataNamespace(original["namespace"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNamespace); val.IsValid() && !isEmptyValue(val) {
		transformed["namespace"] = transformedNamespace
	}

	transformedAnnotations, err := expandCloudRunServiceMetadataAnnotations(original["annotations"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAnnotations); val.IsValid() && !isEmptyValue(val) {
		transformed["annotations"] = transformedAnnotations
	}

	return transformed, nil
}

func expandCloudRunServiceMetadataLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudRunServiceMetadataGeneration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceMetadataResourceVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceMetadataSelfLink(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudRunServiceMetadataUid(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

// If the property hasn't been explicitly set in config use the project defined by the provider or env.
func expandCloudRunServiceMetadataNamespace(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	if v == nil {
		project, err := getProject(d, config)
		if err != nil {
			return project, nil
		}
	}
	return v, nil
}

func expandCloudRunServiceMetadataAnnotations(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
