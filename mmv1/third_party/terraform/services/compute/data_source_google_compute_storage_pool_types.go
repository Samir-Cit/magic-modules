package compute

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func DataSourceGoogleComputeStoragePoolTypes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGoogleComputeStoragePoolTypesRead,
		Schema: map[string]*schema.Schema{
			"storage_pool_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Name of the storage pool type.`,
			},
			"zone": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the zone.`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"deprecated": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The deprecation status associated with this storage pool type.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deleted": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `An optional RFC3339 timestamp on or after which the state of this resource is intended to change to DELETED.
This is only informational and the status will not change unless the client explicitly changes it.`,
						},
						"deprecated": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `An optional RFC3339 timestamp on or after which the state of this resource is intended to change to DEPRECATED.
This is only informational and the status will not change unless the client explicitly changes it.`,
						},
						"obsolete": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `An optional RFC3339 timestamp on or after which the state of this resource is intended to change to OBSOLETE.
This is only informational and the status will not change unless the client explicitly changes it.`,
						},
						"replacement": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `The URL of the suggested replacement for a deprecated resource.
The suggested replacement resource must be the same kind of resource as the deprecated resource.`,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `The deprecation state of this resource. This can be ACTIVE, DEPRECATED, OBSOLETE, or DELETED.
Operations which communicate the end of life date for an image, can use ACTIVE.
Operations which create a new resource using a DEPRECATED resource will return successfully,
but with a warning indicating the deprecated resource and recommending its replacement.
Operations which use OBSOLETE or DELETED resources will be rejected and result in an error.`,
						},
					},
				},
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `An optional description of this resource.`,
			},
			"id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The unique identifier for the resource. This identifier is defined by the server.`,
			},
			"kind": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Type of the resource. Always compute#storagePoolType for storage pool types.`,
			},
			"max_pool_provisioned_capacity_gb": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Maximum storage pool size in GB.`,
			},
			"max_pool_provisioned_iops": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Maximum provisioned IOPS.`,
			},
			"max_pool_provisioned_throughput": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Maximum provisioned throughput.`,
			},
			"min_pool_provisioned_capacity_gb": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Minimum storage pool size in GB.`,
			},
			"min_pool_provisioned_iops": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Minimum provisioned IOPS.`,
			},
			"min_pool_provisioned_throughput": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Minimum provisioned throughput.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Name of the resource.`,
			},
			"self_link": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Server-defined URL for the resource.`,
			},
			"self_link_with_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Server-defined URL for this resource with the resource id.`,
			},
			"supported_disk_types": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The list of disk types supported in this storage pool type.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func dataSourceGoogleComputeStoragePoolTypesRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return err
	}
	zone := d.Get("zone").(string)
	storagePoolType := d.Get("storage_pool_type").(string)

	spt, err := config.NewComputeClient(userAgent).StoragePoolTypes.Get(project, zone, storagePoolType).Do()
	if err != nil {
		return transport_tpg.HandleDataSourceNotFoundError(err, d, "GCE storage pool types", fmt.Sprintf("GCE storage pool types in project %s", project))
	}

	if err := d.Set("kind", spt.Kind); err != nil {
		return fmt.Errorf("Error setting kind: %s", err)
	}
	if err := d.Set("id", spt.Id); err != nil {
		return fmt.Errorf("Error setting id: %s", err)
	}
	if err := d.Set("creation_timestamp", spt.CreationTimestamp); err != nil {
		return fmt.Errorf("Error setting creation_timestamp: %s", err)
	}
	if err := d.Set("name", spt.Name); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	if err := d.Set("description", spt.Description); err != nil {
		return fmt.Errorf("Error setting description: %s", err)
	}

	if spt.Deprecated != nil {
		if err := d.Set("state", spt.Deprecated.State); err != nil {
			return fmt.Errorf("Error setting deprecated.state: %s", err)
		}
		if err := d.Set("replacement", spt.Deprecated.Replacement); err != nil {
			return fmt.Errorf("Error setting deprecated.replacement: %s", err)
		}
		if err := d.Set("deprecated", spt.Deprecated.Deprecated); err != nil {
			return fmt.Errorf("Error setting deprecated.deprecated: %s", err)
		}
		if err := d.Set("obsolete", spt.Deprecated.Obsolete); err != nil {
			return fmt.Errorf("Error setting deprecated.obsolete: %s", err)
		}
		if err := d.Set("deleted", spt.Deprecated.Deleted); err != nil {
			return fmt.Errorf("Error setting deprecated.deleted: %s", err)
		}
	}

	if err := d.Set("zone", spt.Zone); err != nil {
		return fmt.Errorf("Error setting zone: %s", err)
	}
	if err := d.Set("self_link", spt.SelfLink); err != nil {
		return fmt.Errorf("Error setting self_link: %s", err)
	}
	if err := d.Set("self_link_with_id", spt.SelfLinkWithId); err != nil {
		return fmt.Errorf("Error setting self_link_with_id: %s", err)
	}

	if err := d.Set("min_pool_provisioned_capacity_gb", spt.MinPoolProvisionedCapacityGb); err != nil {
		return fmt.Errorf("Error setting min_pool_provisioned_capacity_gb: %s", err)
	}
	if err := d.Set("max_pool_provisioned_capacity_gb", spt.MaxPoolProvisionedCapacityGb); err != nil {
		return fmt.Errorf("Error setting max_pool_provisioned_capacity_gb: %s", err)
	}
	if err := d.Set("min_pool_provisioned_iops", spt.MinPoolProvisionedIops); err != nil {
		return fmt.Errorf("Error setting min_pool_provisioned_iops: %s", err)
	}
	if err := d.Set("max_pool_provisioned_iops", spt.MaxPoolProvisionedIops); err != nil {
		return fmt.Errorf("Error setting max_pool_provisioned_iops: %s", err)
	}
	if err := d.Set("min_pool_provisioned_throughput", spt.MinPoolProvisionedThroughput); err != nil {
		return fmt.Errorf("Error setting min_pool_provisioned_throughput: %s", err)
	}
	if err := d.Set("max_pool_provisioned_throughput", spt.MaxPoolProvisionedThroughput); err != nil {
		return fmt.Errorf("Error setting max_pool_provisioned_throughput: %s", err)
	}

	if err := d.Set("supported_disk_types", spt.SupportedDiskTypes); err != nil {
		return fmt.Errorf("Error setting supported_disk_types: %s", err)
	}

	d.SetId(strconv.FormatUint(spt.Id, 10))

	return nil
}
