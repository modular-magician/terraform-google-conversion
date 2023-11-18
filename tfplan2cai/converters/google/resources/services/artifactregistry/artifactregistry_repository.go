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

package artifactregistry

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func upstreamPoliciesDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	o, n := d.GetChange("virtual_repository_config.0.upstream_policies")
	oldPolicies, ok := o.([]any)
	if !ok {
		return false
	}
	newPolicies, ok := n.([]any)
	if !ok {
		return false
	}

	var oldHashes, newHashes []interface{}
	for _, policy := range oldPolicies {
		data, ok := policy.(map[string]any)
		if !ok {
			return false
		}
		hashStr := fmt.Sprintf("[id:%v priority:%v repository:%v]", data["id"], data["priority"], data["repository"])
		oldHashes = append(oldHashes, hashStr)
	}
	for _, policy := range newPolicies {
		data, ok := policy.(map[string]any)
		if !ok {
			return false
		}
		hashStr := fmt.Sprintf("[id:%v priority:%v repository:%v]", data["id"], data["priority"], data["repository"])
		newHashes = append(newHashes, hashStr)
	}

	oldSet := schema.NewSet(schema.HashString, oldHashes)
	newSet := schema.NewSet(schema.HashString, newHashes)
	return oldSet.Equal(newSet)
}

const ArtifactRegistryRepositoryAssetType string = "artifactregistry.googleapis.com/Repository"

func ResourceConverterArtifactRegistryRepository() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ArtifactRegistryRepositoryAssetType,
		Convert:   GetArtifactRegistryRepositoryCaiObject,
	}
}

func GetArtifactRegistryRepositoryCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//artifactregistry.googleapis.com/projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetArtifactRegistryRepositoryApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ArtifactRegistryRepositoryAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/artifactregistry/v1/rest",
				DiscoveryName:        "Repository",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetArtifactRegistryRepositoryApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	formatProp, err := expandArtifactRegistryRepositoryFormat(d.Get("format"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("format"); !tpgresource.IsEmptyValue(reflect.ValueOf(formatProp)) && (ok || !reflect.DeepEqual(v, formatProp)) {
		obj["format"] = formatProp
	}
	descriptionProp, err := expandArtifactRegistryRepositoryDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	kmsKeyNameProp, err := expandArtifactRegistryRepositoryKmsKeyName(d.Get("kms_key_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("kms_key_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(kmsKeyNameProp)) && (ok || !reflect.DeepEqual(v, kmsKeyNameProp)) {
		obj["kmsKeyName"] = kmsKeyNameProp
	}
	dockerConfigProp, err := expandArtifactRegistryRepositoryDockerConfig(d.Get("docker_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("docker_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(dockerConfigProp)) && (ok || !reflect.DeepEqual(v, dockerConfigProp)) {
		obj["dockerConfig"] = dockerConfigProp
	}
	mavenConfigProp, err := expandArtifactRegistryRepositoryMavenConfig(d.Get("maven_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("maven_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(mavenConfigProp)) && (ok || !reflect.DeepEqual(v, mavenConfigProp)) {
		obj["mavenConfig"] = mavenConfigProp
	}
	modeProp, err := expandArtifactRegistryRepositoryMode(d.Get("mode"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(modeProp)) && (ok || !reflect.DeepEqual(v, modeProp)) {
		obj["mode"] = modeProp
	}
	virtualRepositoryConfigProp, err := expandArtifactRegistryRepositoryVirtualRepositoryConfig(d.Get("virtual_repository_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("virtual_repository_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(virtualRepositoryConfigProp)) && (ok || !reflect.DeepEqual(v, virtualRepositoryConfigProp)) {
		obj["virtualRepositoryConfig"] = virtualRepositoryConfigProp
	}
	cleanupPoliciesProp, err := expandArtifactRegistryRepositoryCleanupPolicies(d.Get("cleanup_policies"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cleanup_policies"); !tpgresource.IsEmptyValue(reflect.ValueOf(cleanupPoliciesProp)) && (ok || !reflect.DeepEqual(v, cleanupPoliciesProp)) {
		obj["cleanupPolicies"] = cleanupPoliciesProp
	}
	remoteRepositoryConfigProp, err := expandArtifactRegistryRepositoryRemoteRepositoryConfig(d.Get("remote_repository_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("remote_repository_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(remoteRepositoryConfigProp)) && (ok || !reflect.DeepEqual(v, remoteRepositoryConfigProp)) {
		obj["remoteRepositoryConfig"] = remoteRepositoryConfigProp
	}
	cleanupPolicyDryRunProp, err := expandArtifactRegistryRepositoryCleanupPolicyDryRun(d.Get("cleanup_policy_dry_run"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cleanup_policy_dry_run"); !tpgresource.IsEmptyValue(reflect.ValueOf(cleanupPolicyDryRunProp)) && (ok || !reflect.DeepEqual(v, cleanupPolicyDryRunProp)) {
		obj["cleanupPolicyDryRun"] = cleanupPolicyDryRunProp
	}
	labelsProp, err := expandArtifactRegistryRepositoryEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return resourceArtifactRegistryRepositoryEncoder(d, config, obj)
}

func resourceArtifactRegistryRepositoryEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*transport_tpg.Config)
	if _, ok := d.GetOk("location"); !ok {
		location, err := tpgresource.GetRegionFromSchema("region", "zone", d, config)
		if err != nil {
			return nil, fmt.Errorf("Cannot determine location: set in this resource, or set provider-level 'region' or 'zone'.")
		}
		if err := d.Set("location", location); err != nil {
			return nil, fmt.Errorf("Error setting location: %s", err)
		}
	}
	return obj, nil
}

func expandArtifactRegistryRepositoryFormat(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryDockerConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedImmutableTags, err := expandArtifactRegistryRepositoryDockerConfigImmutableTags(original["immutable_tags"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImmutableTags); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["immutableTags"] = transformedImmutableTags
	}

	return transformed, nil
}

func expandArtifactRegistryRepositoryDockerConfigImmutableTags(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryMavenConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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

	transformedAllowSnapshotOverwrites, err := expandArtifactRegistryRepositoryMavenConfigAllowSnapshotOverwrites(original["allow_snapshot_overwrites"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowSnapshotOverwrites); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowSnapshotOverwrites"] = transformedAllowSnapshotOverwrites
	}

	transformedVersionPolicy, err := expandArtifactRegistryRepositoryMavenConfigVersionPolicy(original["version_policy"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVersionPolicy); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["versionPolicy"] = transformedVersionPolicy
	}

	return transformed, nil
}

func expandArtifactRegistryRepositoryMavenConfigAllowSnapshotOverwrites(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryMavenConfigVersionPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryVirtualRepositoryConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUpstreamPolicies, err := expandArtifactRegistryRepositoryVirtualRepositoryConfigUpstreamPolicies(original["upstream_policies"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUpstreamPolicies); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["upstreamPolicies"] = transformedUpstreamPolicies
	}

	return transformed, nil
}

func expandArtifactRegistryRepositoryVirtualRepositoryConfigUpstreamPolicies(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedId, err := expandArtifactRegistryRepositoryVirtualRepositoryConfigUpstreamPoliciesId(original["id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["id"] = transformedId
		}

		transformedRepository, err := expandArtifactRegistryRepositoryVirtualRepositoryConfigUpstreamPoliciesRepository(original["repository"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRepository); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["repository"] = transformedRepository
		}

		transformedPriority, err := expandArtifactRegistryRepositoryVirtualRepositoryConfigUpstreamPoliciesPriority(original["priority"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPriority); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["priority"] = transformedPriority
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandArtifactRegistryRepositoryVirtualRepositoryConfigUpstreamPoliciesId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryVirtualRepositoryConfigUpstreamPoliciesRepository(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryVirtualRepositoryConfigUpstreamPoliciesPriority(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryCleanupPolicies(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedAction, err := expandArtifactRegistryRepositoryCleanupPoliciesAction(original["action"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAction); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["action"] = transformedAction
		}

		transformedCondition, err := expandArtifactRegistryRepositoryCleanupPoliciesCondition(original["condition"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCondition); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["condition"] = transformedCondition
		}

		transformedMostRecentVersions, err := expandArtifactRegistryRepositoryCleanupPoliciesMostRecentVersions(original["most_recent_versions"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMostRecentVersions); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["mostRecentVersions"] = transformedMostRecentVersions
		}

		transformedId, err := tpgresource.ExpandString(original["id"], d, config)
		if err != nil {
			return nil, err
		}
		m[transformedId] = transformed
	}
	return m, nil
}

func expandArtifactRegistryRepositoryCleanupPoliciesAction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryCleanupPoliciesCondition(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTagState, err := expandArtifactRegistryRepositoryCleanupPoliciesConditionTagState(original["tag_state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTagState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["tagState"] = transformedTagState
	}

	transformedTagPrefixes, err := expandArtifactRegistryRepositoryCleanupPoliciesConditionTagPrefixes(original["tag_prefixes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTagPrefixes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["tagPrefixes"] = transformedTagPrefixes
	}

	transformedVersionNamePrefixes, err := expandArtifactRegistryRepositoryCleanupPoliciesConditionVersionNamePrefixes(original["version_name_prefixes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVersionNamePrefixes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["versionNamePrefixes"] = transformedVersionNamePrefixes
	}

	transformedPackageNamePrefixes, err := expandArtifactRegistryRepositoryCleanupPoliciesConditionPackageNamePrefixes(original["package_name_prefixes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPackageNamePrefixes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["packageNamePrefixes"] = transformedPackageNamePrefixes
	}

	transformedOlderThan, err := expandArtifactRegistryRepositoryCleanupPoliciesConditionOlderThan(original["older_than"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOlderThan); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["olderThan"] = transformedOlderThan
	}

	transformedNewerThan, err := expandArtifactRegistryRepositoryCleanupPoliciesConditionNewerThan(original["newer_than"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNewerThan); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["newerThan"] = transformedNewerThan
	}

	return transformed, nil
}

func expandArtifactRegistryRepositoryCleanupPoliciesConditionTagState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryCleanupPoliciesConditionTagPrefixes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryCleanupPoliciesConditionVersionNamePrefixes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryCleanupPoliciesConditionPackageNamePrefixes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryCleanupPoliciesConditionOlderThan(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryCleanupPoliciesConditionNewerThan(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryCleanupPoliciesMostRecentVersions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPackageNamePrefixes, err := expandArtifactRegistryRepositoryCleanupPoliciesMostRecentVersionsPackageNamePrefixes(original["package_name_prefixes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPackageNamePrefixes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["packageNamePrefixes"] = transformedPackageNamePrefixes
	}

	transformedKeepCount, err := expandArtifactRegistryRepositoryCleanupPoliciesMostRecentVersionsKeepCount(original["keep_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKeepCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["keepCount"] = transformedKeepCount
	}

	return transformed, nil
}

func expandArtifactRegistryRepositoryCleanupPoliciesMostRecentVersionsPackageNamePrefixes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryCleanupPoliciesMostRecentVersionsKeepCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryRemoteRepositoryConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDescription, err := expandArtifactRegistryRepositoryRemoteRepositoryConfigDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedAptRepository, err := expandArtifactRegistryRepositoryRemoteRepositoryConfigAptRepository(original["apt_repository"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAptRepository); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["aptRepository"] = transformedAptRepository
	}

	transformedDockerRepository, err := expandArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepository(original["docker_repository"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDockerRepository); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dockerRepository"] = transformedDockerRepository
	}

	transformedMavenRepository, err := expandArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepository(original["maven_repository"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMavenRepository); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mavenRepository"] = transformedMavenRepository
	}

	transformedNpmRepository, err := expandArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepository(original["npm_repository"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNpmRepository); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["npmRepository"] = transformedNpmRepository
	}

	transformedPythonRepository, err := expandArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepository(original["python_repository"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPythonRepository); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pythonRepository"] = transformedPythonRepository
	}

	transformedYumRepository, err := expandArtifactRegistryRepositoryRemoteRepositoryConfigYumRepository(original["yum_repository"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedYumRepository); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["yumRepository"] = transformedYumRepository
	}

	transformedUpstreamCredentials, err := expandArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentials(original["upstream_credentials"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUpstreamCredentials); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["upstreamCredentials"] = transformedUpstreamCredentials
	}

	return transformed, nil
}

func expandArtifactRegistryRepositoryRemoteRepositoryConfigDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryRemoteRepositoryConfigAptRepository(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPublicRepository, err := expandArtifactRegistryRepositoryRemoteRepositoryConfigAptRepositoryPublicRepository(original["public_repository"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPublicRepository); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["publicRepository"] = transformedPublicRepository
	}

	return transformed, nil
}

func expandArtifactRegistryRepositoryRemoteRepositoryConfigAptRepositoryPublicRepository(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRepositoryBase, err := expandArtifactRegistryRepositoryRemoteRepositoryConfigAptRepositoryPublicRepositoryRepositoryBase(original["repository_base"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRepositoryBase); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["repositoryBase"] = transformedRepositoryBase
	}

	transformedRepositoryPath, err := expandArtifactRegistryRepositoryRemoteRepositoryConfigAptRepositoryPublicRepositoryRepositoryPath(original["repository_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRepositoryPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["repositoryPath"] = transformedRepositoryPath
	}

	return transformed, nil
}

func expandArtifactRegistryRepositoryRemoteRepositoryConfigAptRepositoryPublicRepositoryRepositoryBase(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryRemoteRepositoryConfigAptRepositoryPublicRepositoryRepositoryPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepository(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPublicRepository, err := expandArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryPublicRepository(original["public_repository"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPublicRepository); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["publicRepository"] = transformedPublicRepository
	}

	return transformed, nil
}

func expandArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryPublicRepository(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepository(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPublicRepository, err := expandArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepositoryPublicRepository(original["public_repository"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPublicRepository); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["publicRepository"] = transformedPublicRepository
	}

	return transformed, nil
}

func expandArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepositoryPublicRepository(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepository(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPublicRepository, err := expandArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepositoryPublicRepository(original["public_repository"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPublicRepository); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["publicRepository"] = transformedPublicRepository
	}

	return transformed, nil
}

func expandArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepositoryPublicRepository(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepository(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPublicRepository, err := expandArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryPublicRepository(original["public_repository"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPublicRepository); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["publicRepository"] = transformedPublicRepository
	}

	return transformed, nil
}

func expandArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryPublicRepository(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryRemoteRepositoryConfigYumRepository(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPublicRepository, err := expandArtifactRegistryRepositoryRemoteRepositoryConfigYumRepositoryPublicRepository(original["public_repository"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPublicRepository); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["publicRepository"] = transformedPublicRepository
	}

	return transformed, nil
}

func expandArtifactRegistryRepositoryRemoteRepositoryConfigYumRepositoryPublicRepository(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRepositoryBase, err := expandArtifactRegistryRepositoryRemoteRepositoryConfigYumRepositoryPublicRepositoryRepositoryBase(original["repository_base"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRepositoryBase); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["repositoryBase"] = transformedRepositoryBase
	}

	transformedRepositoryPath, err := expandArtifactRegistryRepositoryRemoteRepositoryConfigYumRepositoryPublicRepositoryRepositoryPath(original["repository_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRepositoryPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["repositoryPath"] = transformedRepositoryPath
	}

	return transformed, nil
}

func expandArtifactRegistryRepositoryRemoteRepositoryConfigYumRepositoryPublicRepositoryRepositoryBase(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryRemoteRepositoryConfigYumRepositoryPublicRepositoryRepositoryPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentials(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUsernamePasswordCredentials, err := expandArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentialsUsernamePasswordCredentials(original["username_password_credentials"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUsernamePasswordCredentials); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["usernamePasswordCredentials"] = transformedUsernamePasswordCredentials
	}

	return transformed, nil
}

func expandArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentialsUsernamePasswordCredentials(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUsername, err := expandArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentialsUsernamePasswordCredentialsUsername(original["username"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUsername); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["username"] = transformedUsername
	}

	transformedPasswordSecretVersion, err := expandArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentialsUsernamePasswordCredentialsPasswordSecretVersion(original["password_secret_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPasswordSecretVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["passwordSecretVersion"] = transformedPasswordSecretVersion
	}

	return transformed, nil
}

func expandArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentialsUsernamePasswordCredentialsUsername(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentialsUsernamePasswordCredentialsPasswordSecretVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryCleanupPolicyDryRun(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandArtifactRegistryRepositoryEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
