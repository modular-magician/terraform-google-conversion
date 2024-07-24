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

package clouddomains

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ClouddomainsRegistrationAssetType string = "domains.googleapis.com/Registration"

func ResourceConverterClouddomainsRegistration() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ClouddomainsRegistrationAssetType,
		Convert:   GetClouddomainsRegistrationCaiObject,
	}
}

func GetClouddomainsRegistrationCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//domains.googleapis.com/projects/{{project}}/locations/{{location}}/registrations/{{domain_name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetClouddomainsRegistrationApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ClouddomainsRegistrationAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/domains/v1beta1/rest",
				DiscoveryName:        "Registration",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetClouddomainsRegistrationApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	domainNoticesProp, err := expandClouddomainsRegistrationDomainNotices(d.Get("domain_notices"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("domain_notices"); !tpgresource.IsEmptyValue(reflect.ValueOf(domainNoticesProp)) && (ok || !reflect.DeepEqual(v, domainNoticesProp)) {
		obj["domainNotices"] = domainNoticesProp
	}
	contactNoticesProp, err := expandClouddomainsRegistrationContactNotices(d.Get("contact_notices"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("contact_notices"); !tpgresource.IsEmptyValue(reflect.ValueOf(contactNoticesProp)) && (ok || !reflect.DeepEqual(v, contactNoticesProp)) {
		obj["contactNotices"] = contactNoticesProp
	}
	yearlyPriceProp, err := expandClouddomainsRegistrationYearlyPrice(d.Get("yearly_price"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("yearly_price"); !tpgresource.IsEmptyValue(reflect.ValueOf(yearlyPriceProp)) && (ok || !reflect.DeepEqual(v, yearlyPriceProp)) {
		obj["yearlyPrice"] = yearlyPriceProp
	}
	managementSettingsProp, err := expandClouddomainsRegistrationManagementSettings(d.Get("management_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("management_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(managementSettingsProp)) && (ok || !reflect.DeepEqual(v, managementSettingsProp)) {
		obj["managementSettings"] = managementSettingsProp
	}
	dnsSettingsProp, err := expandClouddomainsRegistrationDnsSettings(d.Get("dns_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dns_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(dnsSettingsProp)) && (ok || !reflect.DeepEqual(v, dnsSettingsProp)) {
		obj["dnsSettings"] = dnsSettingsProp
	}
	contactSettingsProp, err := expandClouddomainsRegistrationContactSettings(d.Get("contact_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("contact_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(contactSettingsProp)) && (ok || !reflect.DeepEqual(v, contactSettingsProp)) {
		obj["contactSettings"] = contactSettingsProp
	}
	labelsProp, err := expandClouddomainsRegistrationEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	domainNameProp, err := expandClouddomainsRegistrationDomainName(d.Get("domain_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("domain_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(domainNameProp)) && (ok || !reflect.DeepEqual(v, domainNameProp)) {
		obj["domainName"] = domainNameProp
	}

	return resourceClouddomainsRegistrationEncoder(d, config, obj)
}

func resourceClouddomainsRegistrationEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Request body is registration object with additional fields
	// See https://cloud.google.com/domains/docs/reference/rest/v1beta1/projects.locations.registrations/register

	newObj := make(map[string]interface{})

	newObj["domainNotices"] = obj["domainNotices"]
	delete(obj, "domainNotices")
	newObj["contactNotices"] = obj["contactNotices"]
	delete(obj, "contactNotices")
	newObj["yearlyPrice"] = obj["yearlyPrice"]
	delete(obj, "yearlyPrice")

	newObj["registration"] = obj

	return newObj, nil
}

func expandClouddomainsRegistrationDomainNotices(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactNotices(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationYearlyPrice(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCurrencyCode, err := expandClouddomainsRegistrationYearlyPriceCurrencyCode(original["currency_code"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCurrencyCode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["currencyCode"] = transformedCurrencyCode
	}

	transformedUnits, err := expandClouddomainsRegistrationYearlyPriceUnits(original["units"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUnits); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["units"] = transformedUnits
	}

	return transformed, nil
}

func expandClouddomainsRegistrationYearlyPriceCurrencyCode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationYearlyPriceUnits(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationManagementSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRenewalMethod, err := expandClouddomainsRegistrationManagementSettingsRenewalMethod(original["renewal_method"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRenewalMethod); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["renewalMethod"] = transformedRenewalMethod
	}

	transformedPreferredRenewalMethod, err := expandClouddomainsRegistrationManagementSettingsPreferredRenewalMethod(original["preferred_renewal_method"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPreferredRenewalMethod); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["preferredRenewalMethod"] = transformedPreferredRenewalMethod
	}

	transformedTransferLockState, err := expandClouddomainsRegistrationManagementSettingsTransferLockState(original["transfer_lock_state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTransferLockState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["transferLockState"] = transformedTransferLockState
	}

	return transformed, nil
}

func expandClouddomainsRegistrationManagementSettingsRenewalMethod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationManagementSettingsPreferredRenewalMethod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationManagementSettingsTransferLockState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationDnsSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCustomDns, err := expandClouddomainsRegistrationDnsSettingsCustomDns(original["custom_dns"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCustomDns); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["customDns"] = transformedCustomDns
	}

	transformedGlueRecords, err := expandClouddomainsRegistrationDnsSettingsGlueRecords(original["glue_records"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGlueRecords); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["glueRecords"] = transformedGlueRecords
	}

	return transformed, nil
}

func expandClouddomainsRegistrationDnsSettingsCustomDns(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNameServers, err := expandClouddomainsRegistrationDnsSettingsCustomDnsNameServers(original["name_servers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNameServers); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["nameServers"] = transformedNameServers
	}

	transformedDsRecords, err := expandClouddomainsRegistrationDnsSettingsCustomDnsDsRecords(original["ds_records"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDsRecords); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dsRecords"] = transformedDsRecords
	}

	return transformed, nil
}

func expandClouddomainsRegistrationDnsSettingsCustomDnsNameServers(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationDnsSettingsCustomDnsDsRecords(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedKeyTag, err := expandClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsKeyTag(original["key_tag"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKeyTag); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["keyTag"] = transformedKeyTag
		}

		transformedAlgorithm, err := expandClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsAlgorithm(original["algorithm"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAlgorithm); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["algorithm"] = transformedAlgorithm
		}

		transformedDigestType, err := expandClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsDigestType(original["digest_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDigestType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["digestType"] = transformedDigestType
		}

		transformedDigest, err := expandClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsDigest(original["digest"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDigest); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["digest"] = transformedDigest
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsKeyTag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsAlgorithm(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsDigestType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsDigest(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationDnsSettingsGlueRecords(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedHostName, err := expandClouddomainsRegistrationDnsSettingsGlueRecordsHostName(original["host_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedHostName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["hostName"] = transformedHostName
		}

		transformedIpv4Addresses, err := expandClouddomainsRegistrationDnsSettingsGlueRecordsIpv4Addresses(original["ipv4_addresses"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpv4Addresses); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ipv4Addresses"] = transformedIpv4Addresses
		}

		transformedIpv6Addresses, err := expandClouddomainsRegistrationDnsSettingsGlueRecordsIpv6Addresses(original["ipv6_addresses"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpv6Addresses); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ipv6Addresses"] = transformedIpv6Addresses
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandClouddomainsRegistrationDnsSettingsGlueRecordsHostName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationDnsSettingsGlueRecordsIpv4Addresses(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationDnsSettingsGlueRecordsIpv6Addresses(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPrivacy, err := expandClouddomainsRegistrationContactSettingsPrivacy(original["privacy"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPrivacy); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["privacy"] = transformedPrivacy
	}

	transformedRegistrantContact, err := expandClouddomainsRegistrationContactSettingsRegistrantContact(original["registrant_contact"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRegistrantContact); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["registrantContact"] = transformedRegistrantContact
	}

	transformedAdminContact, err := expandClouddomainsRegistrationContactSettingsAdminContact(original["admin_contact"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAdminContact); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["adminContact"] = transformedAdminContact
	}

	transformedTechnicalContact, err := expandClouddomainsRegistrationContactSettingsTechnicalContact(original["technical_contact"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTechnicalContact); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["technicalContact"] = transformedTechnicalContact
	}

	return transformed, nil
}

func expandClouddomainsRegistrationContactSettingsPrivacy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsRegistrantContact(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEmail, err := expandClouddomainsRegistrationContactSettingsRegistrantContactEmail(original["email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["email"] = transformedEmail
	}

	transformedPhoneNumber, err := expandClouddomainsRegistrationContactSettingsRegistrantContactPhoneNumber(original["phone_number"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPhoneNumber); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["phoneNumber"] = transformedPhoneNumber
	}

	transformedFaxNumber, err := expandClouddomainsRegistrationContactSettingsRegistrantContactFaxNumber(original["fax_number"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFaxNumber); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["faxNumber"] = transformedFaxNumber
	}

	transformedPostalAddress, err := expandClouddomainsRegistrationContactSettingsRegistrantContactPostalAddress(original["postal_address"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPostalAddress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["postalAddress"] = transformedPostalAddress
	}

	return transformed, nil
}

func expandClouddomainsRegistrationContactSettingsRegistrantContactEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsRegistrantContactPhoneNumber(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsRegistrantContactFaxNumber(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsRegistrantContactPostalAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRegionCode, err := expandClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressRegionCode(original["region_code"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRegionCode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["regionCode"] = transformedRegionCode
	}

	transformedPostalCode, err := expandClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressPostalCode(original["postal_code"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPostalCode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["postalCode"] = transformedPostalCode
	}

	transformedAdministrativeArea, err := expandClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressAdministrativeArea(original["administrative_area"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAdministrativeArea); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["administrativeArea"] = transformedAdministrativeArea
	}

	transformedLocality, err := expandClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressLocality(original["locality"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocality); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["locality"] = transformedLocality
	}

	transformedOrganization, err := expandClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOrganization(original["organization"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOrganization); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["organization"] = transformedOrganization
	}

	transformedAddressLines, err := expandClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressAddressLines(original["address_lines"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAddressLines); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["addressLines"] = transformedAddressLines
	}

	transformedRecipients, err := expandClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressRecipients(original["recipients"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRecipients); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["recipients"] = transformedRecipients
	}

	return transformed, nil
}

func expandClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressRegionCode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressPostalCode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressAdministrativeArea(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressLocality(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressOrganization(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressAddressLines(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsRegistrantContactPostalAddressRecipients(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsAdminContact(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEmail, err := expandClouddomainsRegistrationContactSettingsAdminContactEmail(original["email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["email"] = transformedEmail
	}

	transformedPhoneNumber, err := expandClouddomainsRegistrationContactSettingsAdminContactPhoneNumber(original["phone_number"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPhoneNumber); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["phoneNumber"] = transformedPhoneNumber
	}

	transformedFaxNumber, err := expandClouddomainsRegistrationContactSettingsAdminContactFaxNumber(original["fax_number"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFaxNumber); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["faxNumber"] = transformedFaxNumber
	}

	transformedPostalAddress, err := expandClouddomainsRegistrationContactSettingsAdminContactPostalAddress(original["postal_address"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPostalAddress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["postalAddress"] = transformedPostalAddress
	}

	return transformed, nil
}

func expandClouddomainsRegistrationContactSettingsAdminContactEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsAdminContactPhoneNumber(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsAdminContactFaxNumber(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsAdminContactPostalAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRegionCode, err := expandClouddomainsRegistrationContactSettingsAdminContactPostalAddressRegionCode(original["region_code"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRegionCode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["regionCode"] = transformedRegionCode
	}

	transformedPostalCode, err := expandClouddomainsRegistrationContactSettingsAdminContactPostalAddressPostalCode(original["postal_code"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPostalCode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["postalCode"] = transformedPostalCode
	}

	transformedAdministrativeArea, err := expandClouddomainsRegistrationContactSettingsAdminContactPostalAddressAdministrativeArea(original["administrative_area"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAdministrativeArea); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["administrativeArea"] = transformedAdministrativeArea
	}

	transformedLocality, err := expandClouddomainsRegistrationContactSettingsAdminContactPostalAddressLocality(original["locality"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocality); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["locality"] = transformedLocality
	}

	transformedOrganization, err := expandClouddomainsRegistrationContactSettingsAdminContactPostalAddressOrganization(original["organization"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOrganization); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["organization"] = transformedOrganization
	}

	transformedAddressLines, err := expandClouddomainsRegistrationContactSettingsAdminContactPostalAddressAddressLines(original["address_lines"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAddressLines); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["addressLines"] = transformedAddressLines
	}

	transformedRecipients, err := expandClouddomainsRegistrationContactSettingsAdminContactPostalAddressRecipients(original["recipients"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRecipients); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["recipients"] = transformedRecipients
	}

	return transformed, nil
}

func expandClouddomainsRegistrationContactSettingsAdminContactPostalAddressRegionCode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsAdminContactPostalAddressPostalCode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsAdminContactPostalAddressAdministrativeArea(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsAdminContactPostalAddressLocality(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsAdminContactPostalAddressOrganization(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsAdminContactPostalAddressAddressLines(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsAdminContactPostalAddressRecipients(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsTechnicalContact(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEmail, err := expandClouddomainsRegistrationContactSettingsTechnicalContactEmail(original["email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["email"] = transformedEmail
	}

	transformedPhoneNumber, err := expandClouddomainsRegistrationContactSettingsTechnicalContactPhoneNumber(original["phone_number"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPhoneNumber); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["phoneNumber"] = transformedPhoneNumber
	}

	transformedFaxNumber, err := expandClouddomainsRegistrationContactSettingsTechnicalContactFaxNumber(original["fax_number"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFaxNumber); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["faxNumber"] = transformedFaxNumber
	}

	transformedPostalAddress, err := expandClouddomainsRegistrationContactSettingsTechnicalContactPostalAddress(original["postal_address"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPostalAddress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["postalAddress"] = transformedPostalAddress
	}

	return transformed, nil
}

func expandClouddomainsRegistrationContactSettingsTechnicalContactEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsTechnicalContactPhoneNumber(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsTechnicalContactFaxNumber(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsTechnicalContactPostalAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRegionCode, err := expandClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressRegionCode(original["region_code"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRegionCode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["regionCode"] = transformedRegionCode
	}

	transformedPostalCode, err := expandClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressPostalCode(original["postal_code"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPostalCode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["postalCode"] = transformedPostalCode
	}

	transformedAdministrativeArea, err := expandClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressAdministrativeArea(original["administrative_area"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAdministrativeArea); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["administrativeArea"] = transformedAdministrativeArea
	}

	transformedLocality, err := expandClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressLocality(original["locality"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocality); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["locality"] = transformedLocality
	}

	transformedOrganization, err := expandClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOrganization(original["organization"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOrganization); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["organization"] = transformedOrganization
	}

	transformedAddressLines, err := expandClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressAddressLines(original["address_lines"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAddressLines); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["addressLines"] = transformedAddressLines
	}

	transformedRecipients, err := expandClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressRecipients(original["recipients"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRecipients); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["recipients"] = transformedRecipients
	}

	return transformed, nil
}

func expandClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressRegionCode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressPostalCode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressAdministrativeArea(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressLocality(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOrganization(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressAddressLines(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressRecipients(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandClouddomainsRegistrationEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandClouddomainsRegistrationDomainName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
