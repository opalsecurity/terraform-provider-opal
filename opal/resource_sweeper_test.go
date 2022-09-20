package opal

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func init() {
	ctx := context.Background()
	resource.AddTestSweepers("opal_resource", &resource.Sweeper{
		Name: "opal_resource",
		F: func(_ string) error {
			client, err := sweeperClient()
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			resources, _, err := client.ResourcesApi.GetResources(ctx).Execute()
			if err != nil {
				return err
			}

			// XXX: Should we paginate? If we run the sweeper often enough, we shouldn't need to.
			for _, resource := range resources.Results {
				if strings.HasPrefix(*resource.Name, "tf_acc_test_") {
					if _, err := client.ResourcesApi.DeleteResource(ctx, resource.ResourceId).Execute(); err != nil {
						log.Printf("Error destroying resource %s during sweep: %s", resource.ResourceId, err)
					}
				}
			}
			return nil
		},
	})
}
