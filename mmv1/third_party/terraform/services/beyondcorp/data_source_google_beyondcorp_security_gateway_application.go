package beyondcorp

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func DataSourceGoogleBeyondcorpSecurityGatewayApplication() *schema.Resource {

	dsSchema := tpgresource.DatasourceSchemaFromResourceSchema(ResourceBeyondcorpSecurityGatewayApplication().Schema)
	tpgresource.AddRequiredFieldsToSchema(dsSchema, "application_id")
	tpgresource.AddRequiredFieldsToSchema(dsSchema, "security_gateways_id")

	tpgresource.AddOptionalFieldsToSchema(dsSchema, "project")

	return &schema.Resource{
		Read:   dataSourceGoogleBeyondcorpSecurityGatewayApplicationRead,
		Schema: dsSchema,
	}
}

func dataSourceGoogleBeyondcorpSecurityGatewayApplicationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)

	name := d.Get("application_id").(string)
	security_gateways_id := d.Get("security_gateways_id").(string)

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return err
	}

	id := fmt.Sprintf("projects/%s/locations/global/securityGateways/%s/applications/%s", project, security_gateways_id, name)
	d.SetId(id)

	err = resourceBeyondcorpSecurityGatewayApplicationRead(d, meta)
	if err != nil {
		return err
	}

	if err := tpgresource.SetDataSourceLabels(d); err != nil {
		return err
	}

	if d.Id() == "" {
		return fmt.Errorf("%s not found", id)
	}

	return nil
}
