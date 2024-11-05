// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

package compute

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceComputeRegionNetworkEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeRegionNetworkEndpointCreate,
		Read:   resourceComputeRegionNetworkEndpointRead,
		Delete: resourceComputeRegionNetworkEndpointDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeRegionNetworkEndpointImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
			tpgresource.DefaultProviderRegion,
		),

		Schema: map[string]*schema.Schema{
			"port": {
				Type:         schema.TypeInt,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IntAtLeast(1),
				Description:  `Port number of network endpoint.`,
			},
			"region_network_endpoint_group": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareResourceNames,
				Description:      `The network endpoint group this endpoint is part of.`,
			},
			"fqdn": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `Fully qualified domain name of network endpoint.

This can only be specified when network_endpoint_type of the NEG is INTERNET_FQDN_PORT.`,
				AtLeastOneOf: []string{"fqdn", "ip_address"},
			},
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `IPv4 address external endpoint.

This can only be specified when network_endpoint_type of the NEG is INTERNET_IP_PORT.`,
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `Region where the containing network endpoint group is located.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeRegionNetworkEndpointCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	portProp, err := expandNestedComputeRegionNetworkEndpointPort(d.Get("port"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("port"); !tpgresource.IsEmptyValue(reflect.ValueOf(portProp)) && (ok || !reflect.DeepEqual(v, portProp)) {
		obj["port"] = portProp
	}
	ipAddressProp, err := expandNestedComputeRegionNetworkEndpointIpAddress(d.Get("ip_address"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_address"); !tpgresource.IsEmptyValue(reflect.ValueOf(ipAddressProp)) && (ok || !reflect.DeepEqual(v, ipAddressProp)) {
		obj["ipAddress"] = ipAddressProp
	}
	fqdnProp, err := expandNestedComputeRegionNetworkEndpointFqdn(d.Get("fqdn"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fqdn"); !tpgresource.IsEmptyValue(reflect.ValueOf(fqdnProp)) && (ok || !reflect.DeepEqual(v, fqdnProp)) {
		obj["fqdn"] = fqdnProp
	}

	obj, err = resourceComputeRegionNetworkEndpointEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "networkEndpoint/{{project}}/{{region}}/{{region_network_endpoint_group}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/networkEndpointGroups/{{region_network_endpoint_group}}/attachNetworkEndpoints")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new RegionNetworkEndpoint: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionNetworkEndpoint: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating RegionNetworkEndpoint: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{project}}/{{region}}/{{region_network_endpoint_group}}/{{ip_address}}/{{fqdn}}/{{port}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating RegionNetworkEndpoint", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create RegionNetworkEndpoint: %s", err)
	}

	log.Printf("[DEBUG] Finished creating RegionNetworkEndpoint %q: %#v", d.Id(), res)

	return resourceComputeRegionNetworkEndpointRead(d, meta)
}

func resourceComputeRegionNetworkEndpointRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/networkEndpointGroups/{{region_network_endpoint_group}}/listNetworkEndpoints")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionNetworkEndpoint: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeRegionNetworkEndpoint %q", d.Id()))
	}

	res, err = flattenNestedComputeRegionNetworkEndpoint(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing ComputeRegionNetworkEndpoint because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	res, err = resourceComputeRegionNetworkEndpointDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing ComputeRegionNetworkEndpoint because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpoint: %s", err)
	}

	region, err := tpgresource.GetRegion(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("region", region); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpoint: %s", err)
	}

	if err := d.Set("port", flattenNestedComputeRegionNetworkEndpointPort(res["port"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpoint: %s", err)
	}
	if err := d.Set("ip_address", flattenNestedComputeRegionNetworkEndpointIpAddress(res["ipAddress"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpoint: %s", err)
	}
	if err := d.Set("fqdn", flattenNestedComputeRegionNetworkEndpointFqdn(res["fqdn"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionNetworkEndpoint: %s", err)
	}

	return nil
}

func resourceComputeRegionNetworkEndpointDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionNetworkEndpoint: %s", err)
	}
	billingProject = project

	lockName, err := tpgresource.ReplaceVars(d, config, "networkEndpoint/{{project}}/{{region}}/{{region_network_endpoint_group}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/networkEndpointGroups/{{region_network_endpoint_group}}/detachNetworkEndpoints")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	toDelete := make(map[string]interface{})

	// Port
	portProp, err := expandNestedComputeRegionNetworkEndpointPort(d.Get("port"), d, config)
	if err != nil {
		return err
	}
	if portProp != 0 {
		toDelete["port"] = portProp
	}

	// IP address
	ipAddressProp, err := expandNestedComputeRegionNetworkEndpointIpAddress(d.Get("ip_address"), d, config)
	if err != nil {
		return err
	}
	if ipAddressProp != "" {
		toDelete["ipAddress"] = ipAddressProp
	}

	// FQDN
	fqdnProp, err := expandNestedComputeRegionNetworkEndpointFqdn(d.Get("fqdn"), d, config)
	if err != nil {
		return err
	}
	if fqdnProp != "" {
		toDelete["fqdn"] = fqdnProp
	}

	obj = map[string]interface{}{
		"networkEndpoints": []map[string]interface{}{toDelete},
	}

	log.Printf("[DEBUG] Deleting RegionNetworkEndpoint %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "RegionNetworkEndpoint")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting RegionNetworkEndpoint", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting RegionNetworkEndpoint %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeRegionNetworkEndpointImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	// instance is optional, so use * instead of + when reading the import id
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/networkEndpointGroups/(?P<region_network_endpoint_group>[^/]+)/(?P<ip_address>[^/]*)/(?P<fqdn>[^/]*)/(?P<port>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<region_network_endpoint_group>[^/]+)/(?P<ip_address>[^/]*)/(?P<fqdn>[^/]*)/(?P<port>[^/]+)",
		"(?P<region>[^/]+)/(?P<region_network_endpoint_group>[^/]+)/(?P<ip_address>[^/]*)/(?P<fqdn>[^/]*)/(?P<port>[^/]+)",
		"(?P<region_network_endpoint_group>[^/]+)/(?P<ip_address>[^/]*)/(?P<fqdn>[^/]*)/(?P<port>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{project}}/{{region}}/{{region_network_endpoint_group}}/{{ip_address}}/{{fqdn}}/{{port}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNestedComputeRegionNetworkEndpointPort(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles int given in float64 format
	if floatVal, ok := v.(float64); ok {
		return int(floatVal)
	}
	return v
}

func flattenNestedComputeRegionNetworkEndpointIpAddress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedComputeRegionNetworkEndpointFqdn(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNestedComputeRegionNetworkEndpointPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedComputeRegionNetworkEndpointIpAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedComputeRegionNetworkEndpointFqdn(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceComputeRegionNetworkEndpointEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Network Endpoint Group is a URL parameter only, so replace self-link/path with resource name only.
	if err := d.Set("region_network_endpoint_group", tpgresource.GetResourceNameFromSelfLink(d.Get("region_network_endpoint_group").(string))); err != nil {
		return nil, fmt.Errorf("Error setting region_network_endpoint_group: %s", err)
	}

	wrappedReq := map[string]interface{}{
		"networkEndpoints": []interface{}{obj},
	}
	return wrappedReq, nil
}

func flattenNestedComputeRegionNetworkEndpoint(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["items"]
	if !ok || v == nil {
		return nil, nil
	}

	switch v.(type) {
	case []interface{}:
		break
	case map[string]interface{}:
		// Construct list out of single nested resource
		v = []interface{}{v}
	default:
		return nil, fmt.Errorf("expected list or map for value items. Actual value: %v", v)
	}

	_, item, err := resourceComputeRegionNetworkEndpointFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceComputeRegionNetworkEndpointFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedIpAddress, err := expandNestedComputeRegionNetworkEndpointIpAddress(d.Get("ip_address"), d, meta.(*transport_tpg.Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedIpAddress := flattenNestedComputeRegionNetworkEndpointIpAddress(expectedIpAddress, d, meta.(*transport_tpg.Config))
	expectedFqdn, err := expandNestedComputeRegionNetworkEndpointFqdn(d.Get("fqdn"), d, meta.(*transport_tpg.Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedFqdn := flattenNestedComputeRegionNetworkEndpointFqdn(expectedFqdn, d, meta.(*transport_tpg.Config))
	expectedPort, err := expandNestedComputeRegionNetworkEndpointPort(d.Get("port"), d, meta.(*transport_tpg.Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedPort := flattenNestedComputeRegionNetworkEndpointPort(expectedPort, d, meta.(*transport_tpg.Config))

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		// Decode list item before comparing.
		item, err := resourceComputeRegionNetworkEndpointDecoder(d, meta, item)
		if err != nil {
			return -1, nil, err
		}

		itemIpAddress := flattenNestedComputeRegionNetworkEndpointIpAddress(item["ipAddress"], d, meta.(*transport_tpg.Config))
		// IsEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(tpgresource.IsEmptyValue(reflect.ValueOf(itemIpAddress)) && tpgresource.IsEmptyValue(reflect.ValueOf(expectedFlattenedIpAddress))) && !reflect.DeepEqual(itemIpAddress, expectedFlattenedIpAddress) {
			log.Printf("[DEBUG] Skipping item with ipAddress= %#v, looking for %#v)", itemIpAddress, expectedFlattenedIpAddress)
			continue
		}
		itemFqdn := flattenNestedComputeRegionNetworkEndpointFqdn(item["fqdn"], d, meta.(*transport_tpg.Config))
		// IsEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(tpgresource.IsEmptyValue(reflect.ValueOf(itemFqdn)) && tpgresource.IsEmptyValue(reflect.ValueOf(expectedFlattenedFqdn))) && !reflect.DeepEqual(itemFqdn, expectedFlattenedFqdn) {
			log.Printf("[DEBUG] Skipping item with fqdn= %#v, looking for %#v)", itemFqdn, expectedFlattenedFqdn)
			continue
		}
		itemPort := flattenNestedComputeRegionNetworkEndpointPort(item["port"], d, meta.(*transport_tpg.Config))
		// IsEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(tpgresource.IsEmptyValue(reflect.ValueOf(itemPort)) && tpgresource.IsEmptyValue(reflect.ValueOf(expectedFlattenedPort))) && !reflect.DeepEqual(itemPort, expectedFlattenedPort) {
			log.Printf("[DEBUG] Skipping item with port= %#v, looking for %#v)", itemPort, expectedFlattenedPort)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}
func resourceComputeRegionNetworkEndpointDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	v, ok := res["networkEndpoint"]
	if !ok || v == nil {
		return res, nil
	}

	return v.(map[string]interface{}), nil
}
