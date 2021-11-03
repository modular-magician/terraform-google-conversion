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
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var accessApprovalCloudProductMapping = map[string]string{
	"appengine.googleapis.com": "App Engine",
	"bigquery.googleapis.com":  "BigQuery",
	"bigtable.googleapis.com":  "Cloud Bigtable",
	"cloudkms.googleapis.com":  "Cloud Key Management Service",
	"compute.googleapis.com":   "Compute Engine",
	"dataflow.googleapis.com":  "Cloud Dataflow",
	"iam.googleapis.com":       "Cloud Identity and Access Management",
	"pubsub.googleapis.com":    "Cloud Pub/Sub",
	"storage.googleapis.com":   "Cloud Storage",
}

func accessApprovalEnrolledServicesHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	cp := m["cloud_product"].(string)
	if n, ok := accessApprovalCloudProductMapping[cp]; ok {
		cp = n
	}
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(cp))) // ToLower just in case
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["enrollment_level"].(string))))
	return hashcode(buf.String())
}

const AccessApprovalFolderSettingsAssetType string = "accessapproval.googleapis.com/FolderSettings"

func resourceConverterAccessApprovalFolderSettings() ResourceConverter {
	return ResourceConverter{
		AssetType: AccessApprovalFolderSettingsAssetType,
		Convert:   GetAccessApprovalFolderSettingsCaiObject,
	}
}

func GetAccessApprovalFolderSettingsCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//accessapproval.googleapis.com/folders/{{folder_id}}/accessApprovalSettings")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetAccessApprovalFolderSettingsApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: AccessApprovalFolderSettingsAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/accessapproval/v1/rest",
				DiscoveryName:        "FolderSettings",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetAccessApprovalFolderSettingsApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	notificationEmailsProp, err := expandAccessApprovalFolderSettingsNotificationEmails(d.Get("notification_emails"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("notification_emails"); !isEmptyValue(reflect.ValueOf(notificationEmailsProp)) && (ok || !reflect.DeepEqual(v, notificationEmailsProp)) {
		obj["notificationEmails"] = notificationEmailsProp
	}
	enrolledServicesProp, err := expandAccessApprovalFolderSettingsEnrolledServices(d.Get("enrolled_services"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enrolled_services"); !isEmptyValue(reflect.ValueOf(enrolledServicesProp)) && (ok || !reflect.DeepEqual(v, enrolledServicesProp)) {
		obj["enrolledServices"] = enrolledServicesProp
	}

	return obj, nil
}

func expandAccessApprovalFolderSettingsNotificationEmails(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandAccessApprovalFolderSettingsEnrolledServices(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedCloudProduct, err := expandAccessApprovalFolderSettingsEnrolledServicesCloudProduct(original["cloud_product"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCloudProduct); val.IsValid() && !isEmptyValue(val) {
			transformed["cloudProduct"] = transformedCloudProduct
		}

		transformedEnrollmentLevel, err := expandAccessApprovalFolderSettingsEnrolledServicesEnrollmentLevel(original["enrollment_level"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEnrollmentLevel); val.IsValid() && !isEmptyValue(val) {
			transformed["enrollmentLevel"] = transformedEnrollmentLevel
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAccessApprovalFolderSettingsEnrolledServicesCloudProduct(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessApprovalFolderSettingsEnrolledServicesEnrollmentLevel(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
