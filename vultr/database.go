package vultr

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/vultr/govultr/v3"
)

const defaultTimeout = 60 * time.Minute

func readReplicaSchema(isReadReplica bool) map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		// Required
		"region": {
			Type:             schema.TypeString,
			Required:         true,
			DiffSuppressFunc: IgnoreCase,
		},
		"label": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tag": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},
		// Computed
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"date_created": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"plan": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"plan_disk": {
			Type:     schema.TypeInt,
			Computed: true,
			Optional: true,
		},
		"plan_ram": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"plan_vcpus": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"plan_replicas": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"dbname": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"ferretdb_credentials": {
			Type:     schema.TypeMap,
			Computed: true,
			Optional: true,
			Elem:     schema.TypeString,
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				return d.Get("database_engine") != "ferretpg"
			},
		},
		"host": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"public_host": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},
		"port": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"user": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"password": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"latest_backup": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"database_engine": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"database_engine_version": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"vpc_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"maintenance_dow": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"maintenance_time": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"backup_hour": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},
		"backup_minute": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},
		"cluster_time_zone": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"trusted_ips": {
			Type:     schema.TypeSet,
			Computed: true,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"mysql_sql_modes": {
			Type:     schema.TypeSet,
			Computed: true,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"mysql_slow_query_log": {
			Type:     schema.TypeBool,
			Computed: true,
			Optional: true,
		},
		"mysql_require_primary_key": {
			Type:     schema.TypeBool,
			Computed: true,
			Optional: true,
		},
		"mysql_long_query_time": {
			Type:     schema.TypeInt,
			Computed: true,
			Optional: true,
		},
		"eviction_policy": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},
	}

	if isReadReplica {
		s["database_id"] = &schema.Schema{
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.NoZeroValues,
			ForceNew:     true,
		}
	}

	return s
}

func userACLSchema() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"acl_categories": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"acl_channels": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"acl_commands": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"acl_keys": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}

	return s
}

func flattenUserACL(dbUser *govultr.DatabaseUser) []map[string]interface{} {
	f := []map[string]interface{}{
		{
			"acl_categories": dbUser.AccessControl.ACLCategories,
			"acl_channels":   dbUser.AccessControl.ACLChannels,
			"acl_commands":   dbUser.AccessControl.ACLCommands,
			"acl_keys":       dbUser.AccessControl.ACLKeys,
		},
	}
	return f
}
