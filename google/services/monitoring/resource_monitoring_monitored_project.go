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

package monitoring

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceMonitoringMonitoredProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceMonitoringMonitoredProjectCreate,
		Read:   resourceMonitoringMonitoredProjectRead,
		Delete: resourceMonitoringMonitoredProjectDelete,

		Importer: &schema.ResourceImporter{
			State: resourceMonitoringMonitoredProjectImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		SchemaVersion: 1,

		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceMonitoringMonitoredProjectResourceV0().CoreConfigSchema().ImpliedType(),
				Upgrade: ResourceMonitoringMonitoredProjectUpgradeV0,
				Version: 0,
			},
		},

		Schema: map[string]*schema.Schema{
			"metrics_scope": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareResourceNames,
				Description:      `Required. The resource name of the existing Metrics Scope that will monitor this project. Example: locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}`,
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareResourceNames,
				Description:      `Immutable. The resource name of the 'MonitoredProject'. On input, the resource name includes the scoping project ID and monitored project ID. On output, it contains the equivalent project numbers. Example: 'locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}/projects/{MONITORED_PROJECT_ID_OR_NUMBER}'`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The time when this 'MonitoredProject' was created.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceMonitoringMonitoredProjectCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandNestedMonitoringMonitoredProjectName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	obj, err = resourceMonitoringMonitoredProjectEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{MonitoringBasePath}}v1/locations/global/metricsScopes/{{metrics_scope}}/projects")
	if err != nil {
		return err
	}
	url = strings.ReplaceAll(url, "projects/projects/", "projects/")

	log.Printf("[DEBUG] Creating new MonitoredProject: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "POST",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		Timeout:              d.Timeout(schema.TimeoutCreate),
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IsMonitoringPermissionError},
	})
	if err != nil {
		return fmt.Errorf("Error creating MonitoredProject: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "locations/global/metricsScopes/{{metrics_scope}}/projects/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	id = strings.ReplaceAll(id, "projects/projects/", "projects/")
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating MonitoredProject %q: %#v", d.Id(), res)

	return resourceMonitoringMonitoredProjectRead(d, meta)
}

func resourceMonitoringMonitoredProjectRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{MonitoringBasePath}}v1/locations/global/metricsScopes/{{metrics_scope}}")
	if err != nil {
		return err
	}
	url = strings.ReplaceAll(url, "projects/projects/", "projects/")

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	name := d.Get("name").(string)
	name = tpgresource.GetResourceNameFromSelfLink(name)
	d.Set("name", name)
	metricsScope := d.Get("metrics_scope").(string)
	metricsScope = tpgresource.GetResourceNameFromSelfLink(metricsScope)
	d.Set("metrics_scope", metricsScope)
	url, err = tpgresource.ReplaceVars(d, config, "{{MonitoringBasePath}}v1/locations/global/metricsScopes/{{metrics_scope}}")
	if err != nil {
		return err
	}
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "GET",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IsMonitoringPermissionError},
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("MonitoringMonitoredProject %q", d.Id()))
	}

	res, err = flattenNestedMonitoringMonitoredProject(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing MonitoringMonitoredProject because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	res, err = resourceMonitoringMonitoredProjectDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing MonitoringMonitoredProject because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("name", flattenNestedMonitoringMonitoredProjectName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading MonitoredProject: %s", err)
	}
	if err := d.Set("create_time", flattenNestedMonitoringMonitoredProjectCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading MonitoredProject: %s", err)
	}

	return nil
}

func resourceMonitoringMonitoredProjectDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{MonitoringBasePath}}v1/locations/global/metricsScopes/{{metrics_scope}}/projects/{{name}}")
	if err != nil {
		return err
	}
	url = strings.ReplaceAll(url, "projects/projects/", "projects/")

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting MonitoredProject %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "DELETE",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		Timeout:              d.Timeout(schema.TimeoutDelete),
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IsMonitoringPermissionError},
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "MonitoredProject")
	}

	log.Printf("[DEBUG] Finished deleting MonitoredProject %q: %#v", d.Id(), res)
	return nil
}

func resourceMonitoringMonitoredProjectImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	name := d.Get("name").(string)
	name = tpgresource.GetResourceNameFromSelfLink(name)
	d.Set("name", name)
	metricsScope := d.Get("metrics_scope").(string)
	metricsScope = tpgresource.GetResourceNameFromSelfLink(metricsScope)
	d.Set("metrics_scope", metricsScope)
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"locations/global/metricsScopes/(?P<metrics_scope>[^/]+)/projects/(?P<name>[^/]+)",
		"v1/locations/global/metricsScopes/(?P<metrics_scope>[^/]+)/projects/(?P<name>[^/]+)",
		"(?P<metrics_scope>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "locations/global/metricsScopes/{{metrics_scope}}/projects/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	id = strings.ReplaceAll(id, "projects/projects/", "projects/")
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNestedMonitoringMonitoredProjectName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedMonitoringMonitoredProjectCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNestedMonitoringMonitoredProjectName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceMonitoringMonitoredProjectEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	name := d.Get("name").(string)
	name = tpgresource.GetResourceNameFromSelfLink(name)
	d.Set("name", name)
	metricsScope := d.Get("metrics_scope").(string)
	metricsScope = tpgresource.GetResourceNameFromSelfLink(metricsScope)
	d.Set("metrics_scope", metricsScope)
	obj["name"] = fmt.Sprintf("locations/global/metricsScopes/%s/projects/%s", metricsScope, name)
	return obj, nil
}

func flattenNestedMonitoringMonitoredProject(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["monitoredProjects"]
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
		return nil, fmt.Errorf("expected list or map for value monitoredProjects. Actual value: %v", v)
	}

	_, item, err := resourceMonitoringMonitoredProjectFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceMonitoringMonitoredProjectFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedName, err := expandNestedMonitoringMonitoredProjectName(d.Get("name"), d, meta.(*transport_tpg.Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedName := flattenNestedMonitoringMonitoredProjectName(expectedName, d, meta.(*transport_tpg.Config))

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		// Decode list item before comparing.
		item, err := resourceMonitoringMonitoredProjectDecoder(d, meta, item)
		if err != nil {
			return -1, nil, err
		}

		itemName := flattenNestedMonitoringMonitoredProjectName(item["name"], d, meta.(*transport_tpg.Config))
		// IsEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(tpgresource.IsEmptyValue(reflect.ValueOf(itemName)) && tpgresource.IsEmptyValue(reflect.ValueOf(expectedFlattenedName))) && !reflect.DeepEqual(itemName, expectedFlattenedName) {
			log.Printf("[DEBUG] Skipping item with name= %#v, looking for %#v)", itemName, expectedFlattenedName)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}
func resourceMonitoringMonitoredProjectDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*transport_tpg.Config)
	name := res["name"].(string)
	name = tpgresource.GetResourceNameFromSelfLink(name)
	if name != "" {
		project, err := config.NewResourceManagerClient(config.UserAgent).Projects.Get(name).Do()
		if err != nil {
			return nil, err
		}
		res["name"] = project.ProjectId
	}
	return res, nil
}

func resourceMonitoringMonitoredProjectResourceV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metrics_scope": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareResourceNames,
				Description:      `Required. The resource name of the existing Metrics Scope that will monitor this project. Example: locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}`,
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareResourceNames,
				Description:      `Immutable. The resource name of the 'MonitoredProject'. On input, the resource name includes the scoping project ID and monitored project ID. On output, it contains the equivalent project numbers. Example: 'locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}/projects/{MONITORED_PROJECT_ID_OR_NUMBER}'`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The time when this 'MonitoredProject' was created.`,
			},
		},
		UseJSONNumber: true,
	}
}

func ResourceMonitoringMonitoredProjectUpgradeV0(_ context.Context, rawState map[string]any, meta any) (map[string]any, error) {
	log.Printf("[DEBUG] Attributes before migration: %#v", rawState)

	rawState["id"] = strings.TrimPrefix(rawState["id"].(string), "v1/")

	log.Printf("[DEBUG] Attributes after migration: %#v", rawState)
	return rawState, nil
}
