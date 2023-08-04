package mso

import (
	"fmt"
	"log"
	"strings"

	"github.com/ciscoecosystem/mso-go-client/client"
	"github.com/ciscoecosystem/mso-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceMSOTemplateContract() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMSOTemplateContractRead,

		SchemaVersion: version,

		Schema: (map[string]*schema.Schema{
			"schema_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 1000),
			},
			"template_name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 1000),
			},
			"contract_name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 1000),
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"filter_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"filter_relationships": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filter_schema_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"filter_template_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"filter_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"directives": {
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Computed: true,
						},
					},
				},
			},
		}),
	}
}

func dataSourceMSOTemplateContractRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Read", d.Id())

	msoClient := m.(*client.Client)

	schemaId := d.Get("schema_id").(string)

	cont, err := msoClient.GetViaURL(fmt.Sprintf("api/v1/schemas/%s", schemaId))
	if err != nil {
		return err
	}
	count, err := cont.ArrayCount("templates")
	if err != nil {
		return fmt.Errorf("No Template found")
	}
	stateTemplate := d.Get("template_name").(string)
	stateContract := d.Get("contract_name").(string)

	found := false
	for i := 0; i < count && !found; i++ {
		tempCont, err := cont.ArrayElement(i, "templates")
		if err != nil {
			return err
		}
		apiTemplate := models.StripQuotes(tempCont.S("name").String())

		if apiTemplate == stateTemplate {
			contractCount, err := tempCont.ArrayCount("contracts")
			if err != nil {
				return fmt.Errorf("Unable to get contract list")
			}
			for j := 0; j < contractCount; j++ {
				contractCont, err := tempCont.ArrayElement(j, "contracts")
				if err != nil {
					return err
				}
				apiContract := models.StripQuotes(contractCont.S("name").String())
				if apiContract == stateContract {
					d.SetId(fmt.Sprintf("%s/templates/%s/contracts/%s", schemaId, stateTemplate, stateContract))
					d.Set("contract_name", apiContract)
					d.Set("schema_id", schemaId)
					d.Set("template_name", apiTemplate)
					d.Set("display_name", models.StripQuotes(contractCont.S("displayName").String()))
					d.Set("filter_type", models.StripQuotes(contractCont.S("filterType").String()))
					d.Set("scope", models.StripQuotes(contractCont.S("scope").String()))

					var filterList []interface{}
					count, _ := contractCont.ArrayCount("filterRelationships")
					for i := 0; i < count; i++ {
						filterMap := make(map[string]interface{})
						filterCont, err := contractCont.ArrayElement(i, "filterRelationships")
						if err != nil {
							return fmt.Errorf("Unable to parse the filter Relationships list")
						}
						filRef := filterCont.S("filterRef").Data()
						split := strings.Split(filRef.(string), "/")
						filterMap["filter_schema_id"] = fmt.Sprintf("%s", split[2])
						filterMap["filter_template_name"] = fmt.Sprintf("%s", split[4])
						filterMap["filter_name"] = fmt.Sprintf("%s", split[6])
						filterMap["directives"] = filterCont.S("directives").Data().([]interface{})
						filterMap["action"] = filterCont.S("action").Data().(string)
						filterList = append(filterList, filterMap)
					}
					d.Set("filter_relationships", filterList)

					found = true
					break
				}
			}
		}
	}

	if !found {
		return fmt.Errorf("Unable to find the Contract %s in Template %s of Schema Id %s", stateContract, stateTemplate, schemaId)
	}

	log.Printf("[DEBUG] %s: Read finished successfully", d.Id())
	return nil
}
