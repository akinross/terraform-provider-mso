package mso

import (
	"fmt"
	"log"
	"strings"

	"github.com/ciscoecosystem/mso-go-client/client"
	"github.com/ciscoecosystem/mso-go-client/container"
	"github.com/ciscoecosystem/mso-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceMSOSchemaSiteServiceGraph() *schema.Resource {
	return &schema.Resource{
		Create: resourceMSOSchemaSiteServiceGraphCreate,
		Read:   resourceMSOSchemaSiteServiceGraphRead,
		Update: resourceMSOSchemaSiteServiceGraphUpdate,
		Delete: resourceMSOSchemaSiteServiceGraphDelete,

		Importer: &schema.ResourceImporter{
			State: resourceMSOSchemaSiteServiceGraphImport,
		},

		SchemaVersion: version,

		CustomizeDiff: func(diff *schema.ResourceDiff, v interface{}) error {

			// Validate Service Node Type from template
			// msoClient := v.(*client.Client)
			// templateServiceGraphNodes, err := getServiceGraphNodes(d *schema.ResourceData, msoClient)
			// if err != nil {
			// 	return err
			// }
			// for _, templateServiceGraphNode := range templateServiceGraphNodes {
			// 	nodeId := templateServiceGraphNode.(map[string]interface{})["nodeId"].(string)
			// 	nodeType, err := getNodeNameFromId(msoClient, nodeId)
			// 	if err != nil {
			// 		return err
			// 	}
			// 	// Validate Service Node Type
			// }

			// Validate Service Node Type
			valA, valB := diff.GetChange("service_node")
			log.Printf("CKECK THIS OUT: %v \n  %v", valA, valB)
			serviceNodeMap := valB.([]interface{})
			for i, val := range serviceNodeMap {
				found := false
				serviceNode := val.(map[string]interface{})
				if serviceNode["service_node_type"] == "other" {
					var other_provider_connector_type_list = []string{"none", "redir"}
					for _, value := range other_provider_connector_type_list {
						if value == serviceNode["provider_connector_type"] {
							found = true
							break
						}
					}
					if !found {
						return fmt.Errorf("The expected value for service_node.%d.provider_connector_type have to be one of [none, redir] when service_node.%d.service_node_type is other, got %s.", i, i, serviceNode["provider_connector_type"])
					}
				}
				if serviceNode["service_node_type"] == "firewall" {
					firewall_provider_connector_type_list := []string{"none", "redir", "snat", "dnat", "snat_dnat"}
					for _, value := range firewall_provider_connector_type_list {
						if value == serviceNode["provider_connector_type"] {
							found = true
							break
						}
					}
					if !found {
						return fmt.Errorf("The expected value for service_node.%d.provider_connector_type have to be one of [none, redir, snat, dnat, snat_dnat] when service_node.%d.service_node_type is firewall, got %s.", i, i, serviceNode["provider_connector_type"])
					}
				}
			}
			return nil
		},

		Schema: (map[string]*schema.Schema{
			"schema_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(1, 1000),
			},
			"template_name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(1, 1000),
			},
			"site_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(1, 1000),
			},
			"service_graph_name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(1, 1000),
			},
			"service_node": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				Description: "Configure service nodes for the service graph.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device_dn": &schema.Schema{
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringLenBetween(1, 1000),
						},
						"service_node_type": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								"other",
								"load-balancer",
								"firewall",
							}, false),
						},
						"consumer_connector_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ValidateFunc: validation.StringInSlice([]string{
								"none",
								"redir",
							}, false),
						},
						"provider_connector_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"consumer_interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"provider_interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		}),
	}
}

func resourceMSOSchemaSiteServiceGraphImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	log.Printf("[DEBUG] %s: Beginning Import", d.Id())

	msoClient := m.(*client.Client)
	get_attribute := strings.Split(d.Id(), "/")
	schemaId := get_attribute[0]
	siteId := get_attribute[2]
	templateName := get_attribute[4]
	graphName := get_attribute[6]

	cont, err := msoClient.GetViaURL(fmt.Sprintf("api/v1/schemas/%s", schemaId))
	if err != nil {
		return nil, err
	}

	graphCont, _, err := getSiteServiceGraphCont(cont, schemaId, templateName, siteId, graphName)
	if err != nil {
		d.SetId("")
		return nil, err
	}

	serviceNodeList, err := setServiceNodeList(graphCont)
	d.Set("service_node", serviceNodeList)

	d.Set("schema_id", schemaId)
	d.Set("template_name", templateName)
	d.Set("site_id", siteId)
	d.Set("service_graph_name", graphName)

	d.SetId(fmt.Sprintf("%s/templates/%s/serviceGraphs/%s", schemaId, templateName, graphName))
	log.Printf("[DEBUG] %s: Import finished successfully", d.Id())
	return []*schema.ResourceData{d}, nil
}

func resourceMSOSchemaSiteServiceGraphCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Begining Creation Site Service Graph")
	msoClient := m.(*client.Client)

	schemaId := d.Get("schema_id").(string)
	templateName := d.Get("template_name").(string)
	siteId := d.Get("site_id").(string)
	graphName := d.Get("service_graph_name").(string)

	cont, err := msoClient.GetViaURL(fmt.Sprintf("api/v1/schemas/%s", schemaId))
	if err != nil {
		return err
	}

	graphCont, _, err := getTemplateServiceGraphCont(cont, templateName, graphName)
	if err != nil {
		return err
	}

	var siteServiceNodeList []interface{}

	if siteServiceNodes, ok := d.GetOk("service_node"); ok {
		siteServiceNodeList, err = createSiteServiceNodeList(msoClient, siteServiceNodes, graphCont)
		if err != nil {
			return err
		}
	}
	serviceNodePath := fmt.Sprintf("/sites/%s-%s/serviceGraphs/%s/serviceNodes", siteId, templateName, graphName)
	siteServiceGraphPayload := models.GetPatchPayloadList("add", serviceNodePath, siteServiceNodeList)
	_, err = msoClient.PatchbyID(fmt.Sprintf("/api/v1/schemas/%s", schemaId), siteServiceGraphPayload)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/sites/%s/template/%s/serviceGraphs/%s", schemaId, siteId, templateName, graphName))
	log.Printf("[DEBUG] %s: Creation finished successfully", d.Id())
	return resourceMSOSchemaSiteServiceGraphRead(d, m)
}

func resourceMSOSchemaSiteServiceGraphRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Read", d.Id())
	msoClient := m.(*client.Client)

	nodeIdSt := d.Id()
	schemaId := d.Get("schema_id").(string)
	templateName := d.Get("template_name").(string)
	siteId := d.Get("site_id").(string)
	graphName := d.Get("service_graph_name").(string)
	cont, err := msoClient.GetViaURL(fmt.Sprintf("api/v1/schemas/%s", schemaId))
	if err != nil {
		return errorForObjectNotFound(err, d.Id(), cont, d)
	}

	graphCont, _, err := getSiteServiceGraphCont(cont, schemaId, templateName, siteId, graphName)
	if err != nil {
		d.SetId("")
		return nil
	}

	serviceNodeList, err := setServiceNodeList(graphCont)
	d.Set("service_node", serviceNodeList)

	d.Set("schema_id", schemaId)
	d.Set("template_name", templateName)
	d.Set("site_id", siteId)
	d.Set("service_graph_name", graphName)

	d.SetId(nodeIdSt)
	return nil
}

func resourceMSOSchemaSiteServiceGraphUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Begining Update Site Service Graph")
	msoClient := m.(*client.Client)

	schemaId := d.Get("schema_id").(string)
	templateName := d.Get("template_name").(string)
	siteId := d.Get("site_id").(string)
	graphName := d.Get("service_graph_name").(string)

	cont, err := msoClient.GetViaURL(fmt.Sprintf("api/v1/schemas/%s", schemaId))
	if err != nil {
		return err
	}

	if d.HasChange("service_node") {
		graphCont, _, err := getTemplateServiceGraphCont(cont, templateName, graphName)
		if err != nil {
			return err
		}

		if siteServiceNodes, ok := d.GetOk("service_node"); ok {
			siteServiceNodeList, err := createSiteServiceNodeList(msoClient, siteServiceNodes, graphCont)
			if err != nil {
				return err
			}

			serviceNodePath := fmt.Sprintf("/sites/%s-%s/serviceGraphs/%s/serviceNodes", siteId, templateName, graphName)
			siteServiceGraphPayload := models.GetPatchPayloadList("replace", serviceNodePath, siteServiceNodeList)
			_, err = msoClient.PatchbyID(fmt.Sprintf("/api/v1/schemas/%s", schemaId), siteServiceGraphPayload)
			if err != nil {
				return err
			}
		}
	}

	d.SetId(d.Id())
	log.Printf("[DEBUG] %s: Update finished successfully", d.Id())
	return resourceMSOSchemaSiteServiceGraphRead(d, m)
}

func resourceMSOSchemaSiteServiceGraphDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[INFO]: Deletion of site Service Graph is not supported by the API.  Site Service Graph will be removed when site is disassociated from the template or when Service Graph is removed at the template level.")
	return nil
}

func createSiteServiceNodeList(msoClient *client.Client, siteServiceNodes interface{}, graphCont *container.Container) ([]interface{}, error) {
	siteServiceNodeList := make([]interface{}, 0, 1)
	for index, serviceNode := range graphCont.S("serviceNodes").Data().([]interface{}) {
		// nodeType, err := getNodeNameFromId(msoClient, serviceNode.(map[string]interface{})["serviceNodeTypeId"].(string))
		// if err != nil {
		// 	return nil, err
		// }

		siteServiceNodeMap := siteServiceNodes.([]interface{})[index].(map[string]interface{})

		// provider_connector_type_value := siteServiceNodeMap["provider_connector_type"]
		// if nodeType == "firewall" {
		// 	provider_connector_type_value = siteServiceNodeMap["firewall_provider_connector_type"]
		// }

		serviceNodeMap := map[string]interface{}{
			"serviceNodeRef": serviceNode.(map[string]interface{})["serviceNodeRef"],
			"device": map[string]interface{}{
				"dn": siteServiceNodeMap["device_dn"],
			},
			"consumerConnectorType": siteServiceNodeMap["consumer_connector_type"],
			"providerConnectorType": siteServiceNodeMap["provider_connector_type"],
			"consumerInterface":     siteServiceNodeMap["consumer_interface"],
			"providerInterface":     siteServiceNodeMap["provider_interface"],
		}
		siteServiceNodeList = append(siteServiceNodeList, serviceNodeMap)
	}
	return siteServiceNodeList, nil
}

func setServiceNodeList(graphCont *container.Container) ([]interface{}, error) {
	serviceNodeList := make([]interface{}, 0, 1)
	for _, val := range graphCont.S("serviceNodes").Data().([]interface{}) {
		serviceNodeMap := map[string]interface{}{
			"device_dn":               val.(map[string]interface{})["device"].(map[string]interface{})["dn"],
			"consumer_connector_type": val.(map[string]interface{})["consumerConnectorType"],
			"provider_connector_type": val.(map[string]interface{})["providerConnectorType"],
			"consumer_interface":      val.(map[string]interface{})["consumerInterface"],
			"provider_interface":      val.(map[string]interface{})["providerInterface"],
		}

		serviceNodeList = append(serviceNodeList, serviceNodeMap)
	}
	return serviceNodeList, nil
}
