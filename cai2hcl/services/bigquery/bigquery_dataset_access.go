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
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var bigqueryAccessRoleToPrimitiveMap = map[string]string{
	"roles/bigquery.dataOwner":  "OWNER",
	"roles/bigquery.dataEditor": "WRITER",
	"roles/bigquery.dataViewer": "READER",
}

func resourceBigQueryDatasetAccessRoleDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	if primitiveRole, ok := bigqueryAccessRoleToPrimitiveMap[new]; ok {
		return primitiveRole == old
	}
	return false
}

// we want to diff suppress any iam_members that are configured as `iam_member`, but stored in state as a different member type
func resourceBigQueryDatasetAccessIamMemberDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	if primitiveRole, ok := bigqueryAccessRoleToPrimitiveMap[new]; ok {
		return primitiveRole == old
	}

	if d.Get("api_updated_member") == true {
		expectedIamMember := d.Get("iam_member").(string)
		parts := strings.SplitAfter(expectedIamMember, ":")

		strippedIamMember := parts[0]
		if len(parts) > 1 {
			strippedIamMember = parts[1]
		}

		if memberInState := d.Get("user_by_email").(string); memberInState != "" {
			return strings.ToUpper(memberInState) == strings.ToUpper(strippedIamMember)
		}

		if memberInState := d.Get("group_by_email").(string); memberInState != "" {
			return strings.ToUpper(memberInState) == strings.ToUpper(strippedIamMember)
		}

		if memberInState := d.Get("domain").(string); memberInState != "" {
			return strings.ToUpper(memberInState) == strings.ToUpper(strippedIamMember)
		}

		if memberInState := d.Get("special_group").(string); memberInState != "" {
			return strings.ToUpper(memberInState) == strings.ToUpper(strippedIamMember)
		}
	}

	return false
}

// this function will go through a response's access list and see if the iam_member has been reassigned to a different member_type
// if it has, it will return the member type, and the member
func resourceBigQueryDatasetAccessReassignIamMemberInNestedObjectList(d *schema.ResourceData, meta interface{}, items []interface{}) (member_type string, member interface{}, err error) {
	expectedRole, err := expandNestedBigQueryDatasetAccessRole(d.Get("role"), d, meta.(*transport_tpg.Config))
	if err != nil {
		return "", nil, err
	}
	expectedFlattenedRole := flattenNestedBigQueryDatasetAccessRole(expectedRole, d, meta.(*transport_tpg.Config))

	expectedIamMember, err := expandNestedBigQueryDatasetAccessIamMember(d.Get("iam_member"), d, meta.(*transport_tpg.Config))
	if err != nil {
		return "", nil, err
	}
	expectedFlattenedIamMember := flattenNestedBigQueryDatasetAccessIamMember(expectedIamMember, d, meta.(*transport_tpg.Config))

	parts := strings.SplitAfter(expectedFlattenedIamMember.(string), ":")

	expectedStrippedIamMember := parts[0]
	if len(parts) > 1 {
		expectedStrippedIamMember = parts[1]
	}

	// Search list for this resource.
	for _, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		itemRole := flattenNestedBigQueryDatasetAccessRole(item["role"], d, meta.(*transport_tpg.Config))
		// IsEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(tpgresource.IsEmptyValue(reflect.ValueOf(itemRole)) && tpgresource.IsEmptyValue(reflect.ValueOf(expectedFlattenedRole))) && !reflect.DeepEqual(itemRole, expectedFlattenedRole) {
			log.Printf("[DEBUG] Skipping item with role= %#v, looking for %#v)", itemRole, expectedFlattenedRole)
			continue
		}

		itemUserByEmail := flattenNestedBigQueryDatasetAccessUserByEmail(item["userByEmail"], d, meta.(*transport_tpg.Config))
		if reflect.DeepEqual(itemUserByEmail, expectedStrippedIamMember) {
			log.Printf("[DEBUG] Iam Member changed to userByEmail= %#v)", itemUserByEmail)
			return "user_by_email", itemUserByEmail, nil
		}
		itemGroupByEmail := flattenNestedBigQueryDatasetAccessGroupByEmail(item["groupByEmail"], d, meta.(*transport_tpg.Config))
		if reflect.DeepEqual(itemGroupByEmail, expectedStrippedIamMember) {
			log.Printf("[DEBUG] Iam Member changed to groupByEmail= %#v)", itemGroupByEmail)
			return "group_by_email", itemGroupByEmail, nil
		}
		itemDomain := flattenNestedBigQueryDatasetAccessDomain(item["domain"], d, meta.(*transport_tpg.Config))
		if reflect.DeepEqual(itemDomain, expectedStrippedIamMember) {
			log.Printf("[DEBUG] Iam Member changed to domain= %#v)", itemDomain)
			return "domain", itemDomain, nil
		}
		itemSpecialGroup := flattenNestedBigQueryDatasetAccessSpecialGroup(item["specialGroup"], d, meta.(*transport_tpg.Config))
		if reflect.DeepEqual(itemSpecialGroup, expectedStrippedIamMember) {
			log.Printf("[DEBUG] Iam Member changed to specialGroup= %#v)", itemSpecialGroup)
			return "special_group", itemSpecialGroup, nil
		}
		itemIamMember := flattenNestedBigQueryDatasetAccessIamMember(item["iamMember"], d, meta.(*transport_tpg.Config))
		if reflect.DeepEqual(itemIamMember, expectedFlattenedIamMember) {
			log.Printf("[DEBUG] Iam Member stayed as iamMember= %#v)", itemIamMember)
			return "", nil, nil
		}
		continue
	}
	log.Printf("[DEBUG] Did not find item for resource %q)", d.Id())
	return "", nil, nil
}

const BigQueryDatasetAccessAssetType string = "bigquery.googleapis.com/DatasetAccess"

const BigQueryDatasetAccessAssetNameRegex string = "projects/(?P<project>[^/]+)/datasets/(?P<dataset_id>[^/]+)"

type BigQueryDatasetAccessConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewBigQueryDatasetAccessConverter(name string, schema map[string]*schema.Schema) common.Converter {
	return &BigQueryDatasetAccessConverter{
		name:   name,
		schema: schema,
	}
}

func (c *BigQueryDatasetAccessConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
	var blocks []*common.HCLResourceBlock
	config := common.NewConfig()

	for _, asset := range assets {
		if asset == nil {
			continue
		}
		if asset.Resource != nil && asset.Resource.Data != nil {
			block, err := c.convertResourceData(asset, config)
			if err != nil {
				return nil, err
			}
			blocks = append(blocks, block)
		}
	}
	return blocks, nil
}

func (c *BigQueryDatasetAccessConverter) convertResourceData(asset *caiasset.Asset, config *transport_tpg.Config) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	assetResourceData := asset.Resource.Data

	hcl, _ := resourceBigQueryDatasetAccessRead(assetResourceData, config)

	ctyVal, err := common.MapToCtyValWithSchema(hcl, c.schema)
	if err != nil {
		return nil, err
	}

	resourceName := assetResourceData["name"].(string)

	return &common.HCLResourceBlock{
		Labels: []string{c.name, resourceName},
		Value:  ctyVal,
	}, nil
}

func resourceBigQueryDatasetAccessRead(resource map[string]interface{}, config *transport_tpg.Config) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var resource_data *schema.ResourceData = nil

	result["role"] = flattenNestedBigQueryDatasetAccessRole(resource["role"], resource_data, config)
	result["user_by_email"] = flattenNestedBigQueryDatasetAccessUserByEmail(resource["userByEmail"], resource_data, config)
	result["group_by_email"] = flattenNestedBigQueryDatasetAccessGroupByEmail(resource["groupByEmail"], resource_data, config)
	result["domain"] = flattenNestedBigQueryDatasetAccessDomain(resource["domain"], resource_data, config)
	result["special_group"] = flattenNestedBigQueryDatasetAccessSpecialGroup(resource["specialGroup"], resource_data, config)
	result["iam_member"] = flattenNestedBigQueryDatasetAccessIamMember(resource["iamMember"], resource_data, config)
	result["view"] = flattenNestedBigQueryDatasetAccessView(resource["view"], resource_data, config)
	result["dataset"] = flattenNestedBigQueryDatasetAccessDataset(resource["dataset"], resource_data, config)
	result["routine"] = flattenNestedBigQueryDatasetAccessRoutine(resource["routine"], resource_data, config)

	return result, nil
}

func flattenNestedBigQueryDatasetAccessRole(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedBigQueryDatasetAccessUserByEmail(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedBigQueryDatasetAccessGroupByEmail(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedBigQueryDatasetAccessDomain(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedBigQueryDatasetAccessSpecialGroup(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedBigQueryDatasetAccessIamMember(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedBigQueryDatasetAccessView(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["dataset_id"] =
		flattenNestedBigQueryDatasetAccessViewDatasetId(original["datasetId"], d, config)
	transformed["project_id"] =
		flattenNestedBigQueryDatasetAccessViewProjectId(original["projectId"], d, config)
	transformed["table_id"] =
		flattenNestedBigQueryDatasetAccessViewTableId(original["tableId"], d, config)
	return []interface{}{transformed}
}
func flattenNestedBigQueryDatasetAccessViewDatasetId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedBigQueryDatasetAccessViewProjectId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedBigQueryDatasetAccessViewTableId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedBigQueryDatasetAccessDataset(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["dataset"] =
		flattenNestedBigQueryDatasetAccessDatasetDataset(original["dataset"], d, config)
	transformed["target_types"] =
		flattenNestedBigQueryDatasetAccessDatasetTargetTypes(original["targetTypes"], d, config)
	return []interface{}{transformed}
}
func flattenNestedBigQueryDatasetAccessDatasetDataset(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["dataset_id"] =
		flattenNestedBigQueryDatasetAccessDatasetDatasetDatasetId(original["datasetId"], d, config)
	transformed["project_id"] =
		flattenNestedBigQueryDatasetAccessDatasetDatasetProjectId(original["projectId"], d, config)
	return []interface{}{transformed}
}
func flattenNestedBigQueryDatasetAccessDatasetDatasetDatasetId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedBigQueryDatasetAccessDatasetDatasetProjectId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedBigQueryDatasetAccessDatasetTargetTypes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedBigQueryDatasetAccessRoutine(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["dataset_id"] =
		flattenNestedBigQueryDatasetAccessRoutineDatasetId(original["datasetId"], d, config)
	transformed["project_id"] =
		flattenNestedBigQueryDatasetAccessRoutineProjectId(original["projectId"], d, config)
	transformed["routine_id"] =
		flattenNestedBigQueryDatasetAccessRoutineRoutineId(original["routineId"], d, config)
	return []interface{}{transformed}
}
func flattenNestedBigQueryDatasetAccessRoutineDatasetId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedBigQueryDatasetAccessRoutineProjectId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedBigQueryDatasetAccessRoutineRoutineId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
