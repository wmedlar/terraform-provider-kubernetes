package kubernetes

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceKubernetesPodV0() *schema.Resource {
	schemaV1 := dataSourceKubernetesPodSchemaV1()
	schemaV0 := patchPodSpecWithResourcesFieldV0(schemaV1)
	return &schema.Resource{Schema: schemaV0}
}

func dataSourceKubernetesPodUpgradeV0(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	log.Println("[DEBUG] migrating state for data.kubernetes_pod")
	return upgradePodSpecWithResourcesFieldV0(ctx, rawState, meta)
}
