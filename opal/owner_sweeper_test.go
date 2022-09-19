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
	resource.AddTestSweepers("opal_owner", &resource.Sweeper{
		Name: "opal_owner",
		F: func(_ string) error {
			client, err := sweeperClient()
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			owners, _, err := client.OwnersApi.GetOwners(ctx).Execute()
			if err != nil {
				return err
			}

			// XXX: Should we paginate? If we run the sweeper often enough, we shouldn't need to.
			for _, owner := range owners.Results {
				if strings.HasPrefix(*owner.Name, "tf_acc_test_") {
					if _, err := client.OwnersApi.DeleteOwner(ctx, owner.OwnerId).Execute(); err != nil {
						log.Printf("Error destroying owner %s during sweep: %s", owner.OwnerId, err)
					}
				}
			}
			return nil
		},
	})
}
