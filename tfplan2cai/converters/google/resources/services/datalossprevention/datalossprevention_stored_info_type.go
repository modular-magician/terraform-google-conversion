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

package datalossprevention

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// This customizeDiff allows updating the dictionary, regex, and large_custom_dictionary fields, but
// it recreates the resource if changing between these fields. e.g., updating the regex field should
// be allowed, while changing from regex to dictionary should trigger the recreation of the resource.
func storedInfoTypeCustomizeDiffFunc(diff tpgresource.TerraformResourceDiff) error {
	oldDict, newDict := diff.GetChange("dictionary")
	oldRegex, newRegex := diff.GetChange("regex")
	oldLargeCD, newLargeCD := diff.GetChange("large_custom_dictionary")
	if !tpgresource.IsEmptyValue(reflect.ValueOf(oldDict)) && tpgresource.IsEmptyValue(reflect.ValueOf(newDict)) {
		diff.ForceNew("dictionary")
		return nil
	}
	if !tpgresource.IsEmptyValue(reflect.ValueOf(oldRegex)) && tpgresource.IsEmptyValue(reflect.ValueOf(newRegex)) {
		diff.ForceNew("regex")
		return nil
	}
	if !tpgresource.IsEmptyValue(reflect.ValueOf(oldLargeCD)) && tpgresource.IsEmptyValue(reflect.ValueOf(newLargeCD)) {
		diff.ForceNew("large_custom_dictionary")
		return nil
	}
	return nil
}

func storedInfoTypeCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
	return storedInfoTypeCustomizeDiffFunc(diff)
}

const DataLossPreventionStoredInfoTypeAssetType string = "dlp.googleapis.com/StoredInfoType"

func ResourceConverterDataLossPreventionStoredInfoType() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: DataLossPreventionStoredInfoTypeAssetType,
		Convert:   GetDataLossPreventionStoredInfoTypeCaiObject,
	}
}

func GetDataLossPreventionStoredInfoTypeCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//dlp.googleapis.com/{{parent}}/storedInfoTypes/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetDataLossPreventionStoredInfoTypeApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: DataLossPreventionStoredInfoTypeAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/dlp/v2/rest",
				DiscoveryName:        "StoredInfoType",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetDataLossPreventionStoredInfoTypeApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandDataLossPreventionStoredInfoTypeDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	displayNameProp, err := expandDataLossPreventionStoredInfoTypeDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	regexProp, err := expandDataLossPreventionStoredInfoTypeRegex(d.Get("regex"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("regex"); !tpgresource.IsEmptyValue(reflect.ValueOf(regexProp)) && (ok || !reflect.DeepEqual(v, regexProp)) {
		obj["regex"] = regexProp
	}
	dictionaryProp, err := expandDataLossPreventionStoredInfoTypeDictionary(d.Get("dictionary"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dictionary"); !tpgresource.IsEmptyValue(reflect.ValueOf(dictionaryProp)) && (ok || !reflect.DeepEqual(v, dictionaryProp)) {
		obj["dictionary"] = dictionaryProp
	}
	largeCustomDictionaryProp, err := expandDataLossPreventionStoredInfoTypeLargeCustomDictionary(d.Get("large_custom_dictionary"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("large_custom_dictionary"); !tpgresource.IsEmptyValue(reflect.ValueOf(largeCustomDictionaryProp)) && (ok || !reflect.DeepEqual(v, largeCustomDictionaryProp)) {
		obj["largeCustomDictionary"] = largeCustomDictionaryProp
	}

	return resourceDataLossPreventionStoredInfoTypeEncoder(d, config, obj)
}

func resourceDataLossPreventionStoredInfoTypeEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	newObj := make(map[string]interface{})
	newObj["config"] = obj
	storedInfoTypeIdProp, ok := d.GetOk("stored_info_type_id")
	if ok && storedInfoTypeIdProp != nil {
		newObj["storedInfoTypeId"] = storedInfoTypeIdProp
	}
	return newObj, nil
}

func expandDataLossPreventionStoredInfoTypeDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionStoredInfoTypeDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionStoredInfoTypeRegex(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPattern, err := expandDataLossPreventionStoredInfoTypeRegexPattern(original["pattern"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPattern); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pattern"] = transformedPattern
	}

	transformedGroupIndexes, err := expandDataLossPreventionStoredInfoTypeRegexGroupIndexes(original["group_indexes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGroupIndexes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["groupIndexes"] = transformedGroupIndexes
	}

	return transformed, nil
}

func expandDataLossPreventionStoredInfoTypeRegexPattern(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionStoredInfoTypeRegexGroupIndexes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionStoredInfoTypeDictionary(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedWordList, err := expandDataLossPreventionStoredInfoTypeDictionaryWordList(original["word_list"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWordList); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["wordList"] = transformedWordList
	}

	transformedCloudStoragePath, err := expandDataLossPreventionStoredInfoTypeDictionaryCloudStoragePath(original["cloud_storage_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCloudStoragePath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cloudStoragePath"] = transformedCloudStoragePath
	}

	return transformed, nil
}

func expandDataLossPreventionStoredInfoTypeDictionaryWordList(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedWords, err := expandDataLossPreventionStoredInfoTypeDictionaryWordListWords(original["words"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWords); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["words"] = transformedWords
	}

	return transformed, nil
}

func expandDataLossPreventionStoredInfoTypeDictionaryWordListWords(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionStoredInfoTypeDictionaryCloudStoragePath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPath, err := expandDataLossPreventionStoredInfoTypeDictionaryCloudStoragePathPath(original["path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["path"] = transformedPath
	}

	return transformed, nil
}

func expandDataLossPreventionStoredInfoTypeDictionaryCloudStoragePathPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionStoredInfoTypeLargeCustomDictionary(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedOutputPath, err := expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPath(original["output_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOutputPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["outputPath"] = transformedOutputPath
	}

	transformedCloudStorageFileSet, err := expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(original["cloud_storage_file_set"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCloudStorageFileSet); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cloudStorageFileSet"] = transformedCloudStorageFileSet
	}

	transformedBigQueryField, err := expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField(original["big_query_field"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBigQueryField); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["bigQueryField"] = transformedBigQueryField
	}

	return transformed, nil
}

func expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPath, err := expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPathPath(original["path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["path"] = transformedPath
	}

	return transformed, nil
}

func expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPathPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUrl, err := expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetUrl(original["url"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUrl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["url"] = transformedUrl
	}

	return transformed, nil
}

func expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetUrl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTable, err := expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(original["table"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTable); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["table"] = transformedTable
	}

	transformedField, err := expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(original["field"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedField); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["field"] = transformedField
	}

	return transformed, nil
}

func expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProjectId, err := expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	transformedDatasetId, err := expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableDatasetId(original["dataset_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatasetId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["datasetId"] = transformedDatasetId
	}

	transformedTableId, err := expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableTableId(original["table_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTableId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["tableId"] = transformedTableId
	}

	return transformed, nil
}

func expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableDatasetId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableTableId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedName, err := expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldName(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["name"] = transformedName
	}

	return transformed, nil
}

func expandDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
