/*
Package acceptance includes all test cases of resources and data sources that
imported directly from huaweicloud.
*/
package acceptance

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/FlexibleEngineCloud/terraform-provider-flexibleengine/flexibleengine"
)

var (
	OS_DEPRECATED_ENVIRONMENT = os.Getenv("OS_DEPRECATED_ENVIRONMENT")
	OS_AVAILABILITY_ZONE      = os.Getenv("OS_AVAILABILITY_ZONE")
	OS_REGION_NAME            = os.Getenv("OS_REGION_NAME")
	OS_ACCESS_KEY             = os.Getenv("OS_ACCESS_KEY")
	OS_SECRET_KEY             = os.Getenv("OS_SECRET_KEY")

	OS_VPC_ID     = os.Getenv("OS_VPC_ID")
	OS_NETWORK_ID = os.Getenv("OS_NETWORK_ID")
	OS_SUBNET_ID  = os.Getenv("OS_SUBNET_ID")

	OS_FLAVOR_ID    = os.Getenv("OS_FLAVOR_ID")
	OS_IMAGE_ID     = os.Getenv("OS_IMAGE_ID")
	OS_KEYPAIR_NAME = os.Getenv("OS_KEYPAIR_NAME")
	OS_FGS_BUCKET   = os.Getenv("OS_FGS_BUCKET")
)

// TestAccProviderFactories is a static map containing only the main provider instance
var TestAccProviderFactories map[string]func() (*schema.Provider, error)

// testAccProvider is the "main" provider instance
var testAccProvider *schema.Provider

func init() {
	testAccProvider = flexibleengine.Provider()

	TestAccProviderFactories = map[string]func() (*schema.Provider, error){
		"flexibleengine": func() (*schema.Provider, error) {
			return testAccProvider, nil
		},
	}
}

func testAccPreCheckRequiredEnvVars(t *testing.T) {
	if OS_REGION_NAME == "" {
		t.Fatal("OS_REGION_NAME must be set for acceptance tests")
	}

	if OS_AVAILABILITY_ZONE == "" {
		t.Fatal("OS_AVAILABILITY_ZONE must be set for acceptance tests")
	}
}

func testAccPreCheck(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)

	// Do not run the test if this is a deprecated testing environment.
	if OS_DEPRECATED_ENVIRONMENT != "" {
		t.Skip("This environment only runs deprecated tests")
	}
}

func testAccPreCheckDeprecated(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)

	if OS_DEPRECATED_ENVIRONMENT == "" {
		t.Skip("This environment does not support deprecated tests")
	}
}

func testAccPreCheckAdminOnly(t *testing.T) {
	v := os.Getenv("OS_ADMIN")
	if v != "admin" {
		t.Skip("Skipping test because it requires the admin user")
	}
}

func testAccPreCheckOBS(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)

	if OS_ACCESS_KEY == "" || OS_SECRET_KEY == "" {
		t.Skip("OS_ACCESS_KEY and OS_SECRET_KEY must be set for OBS acceptance tests")
	}
}