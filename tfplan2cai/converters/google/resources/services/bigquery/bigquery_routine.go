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

package bigquery

import (
	"encoding/json"
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const BigQueryRoutineAssetType string = "bigquery.googleapis.com/Routine"

func ResourceConverterBigQueryRoutine() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: BigQueryRoutineAssetType,
		Convert:   GetBigQueryRoutineCaiObject,
	}
}

func GetBigQueryRoutineCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//bigquery.googleapis.com/projects/{{project}}/datasets/{{dataset_id}}/routines/{{routine_id}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetBigQueryRoutineApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: BigQueryRoutineAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/bigquery/v2/rest",
				DiscoveryName:        "Routine",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetBigQueryRoutineApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	routineReferenceProp, err := expandBigQueryRoutineRoutineReference(nil, d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("routine_reference"); !tpgresource.IsEmptyValue(reflect.ValueOf(routineReferenceProp)) && (ok || !reflect.DeepEqual(v, routineReferenceProp)) {
		obj["routineReference"] = routineReferenceProp
	}
	routineTypeProp, err := expandBigQueryRoutineRoutineType(d.Get("routine_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("routine_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(routineTypeProp)) && (ok || !reflect.DeepEqual(v, routineTypeProp)) {
		obj["routineType"] = routineTypeProp
	}
	languageProp, err := expandBigQueryRoutineLanguage(d.Get("language"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("language"); !tpgresource.IsEmptyValue(reflect.ValueOf(languageProp)) && (ok || !reflect.DeepEqual(v, languageProp)) {
		obj["language"] = languageProp
	}
	argumentsProp, err := expandBigQueryRoutineArguments(d.Get("arguments"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("arguments"); !tpgresource.IsEmptyValue(reflect.ValueOf(argumentsProp)) && (ok || !reflect.DeepEqual(v, argumentsProp)) {
		obj["arguments"] = argumentsProp
	}
	returnTypeProp, err := expandBigQueryRoutineReturnType(d.Get("return_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("return_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(returnTypeProp)) && (ok || !reflect.DeepEqual(v, returnTypeProp)) {
		obj["returnType"] = returnTypeProp
	}
	returnTableTypeProp, err := expandBigQueryRoutineReturnTableType(d.Get("return_table_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("return_table_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(returnTableTypeProp)) && (ok || !reflect.DeepEqual(v, returnTableTypeProp)) {
		obj["returnTableType"] = returnTableTypeProp
	}
	importedLibrariesProp, err := expandBigQueryRoutineImportedLibraries(d.Get("imported_libraries"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("imported_libraries"); !tpgresource.IsEmptyValue(reflect.ValueOf(importedLibrariesProp)) && (ok || !reflect.DeepEqual(v, importedLibrariesProp)) {
		obj["importedLibraries"] = importedLibrariesProp
	}
	definitionBodyProp, err := expandBigQueryRoutineDefinitionBody(d.Get("definition_body"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("definition_body"); !tpgresource.IsEmptyValue(reflect.ValueOf(definitionBodyProp)) && (ok || !reflect.DeepEqual(v, definitionBodyProp)) {
		obj["definitionBody"] = definitionBodyProp
	}
	descriptionProp, err := expandBigQueryRoutineDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	determinismLevelProp, err := expandBigQueryRoutineDeterminismLevel(d.Get("determinism_level"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("determinism_level"); !tpgresource.IsEmptyValue(reflect.ValueOf(determinismLevelProp)) && (ok || !reflect.DeepEqual(v, determinismLevelProp)) {
		obj["determinismLevel"] = determinismLevelProp
	}

	return obj, nil
}

func expandBigQueryRoutineRoutineReference(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {

	transformed := make(map[string]interface{})
	transformed["datasetId"] = d.Get("dataset_id")
	project, _ := tpgresource.GetProject(d, config)
	transformed["projectId"] = project
	transformed["routineId"] = d.Get("routine_id")

	return transformed, nil
}

func expandBigQueryRoutineRoutineType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryRoutineLanguage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryRoutineArguments(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandBigQueryRoutineArgumentsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedArgumentKind, err := expandBigQueryRoutineArgumentsArgumentKind(original["argument_kind"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedArgumentKind); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["argumentKind"] = transformedArgumentKind
		}

		transformedMode, err := expandBigQueryRoutineArgumentsMode(original["mode"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["mode"] = transformedMode
		}

		transformedDataType, err := expandBigQueryRoutineArgumentsDataType(original["data_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDataType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["dataType"] = transformedDataType
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandBigQueryRoutineArgumentsName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryRoutineArgumentsArgumentKind(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryRoutineArgumentsMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryRoutineArgumentsDataType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	b := []byte(v.(string))
	if len(b) == 0 {
		return nil, nil
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func expandBigQueryRoutineReturnType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	b := []byte(v.(string))
	if len(b) == 0 {
		return nil, nil
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func expandBigQueryRoutineReturnTableType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	b := []byte(v.(string))
	if len(b) == 0 {
		return nil, nil
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func expandBigQueryRoutineImportedLibraries(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryRoutineDefinitionBody(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryRoutineDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryRoutineDeterminismLevel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
