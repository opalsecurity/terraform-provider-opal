package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Acceptance tests for linking opal_group / opal_resource to a configuration
// template: create linked, attach, unlink, repoint, rename while linked, and
// reads. They cover the Update ordering (template PUT first), the SDK hook that
// strips template-owned fields while an entity is linked, and the plan-time
// unlink handling in configuration_template_plan.go.

func templateHCL(name, label string) string {
	return fmt.Sprintf(`
resource "opal_configuration_template" "%[2]s" {
  name                   = "%[1]s_%[2]s"
  admin_owner_id         = "%[3]s"
  require_mfa_to_connect = false
  visibility = {
    visibility = "GLOBAL"
  }
}
`, name, label, knownOpalAppAdminOwnerID)
}

func groupOwnConfigHCL(name, groupName, extra string) string {
	return fmt.Sprintf(`
resource "opal_group" "g" {
  name           = "%[2]s"
  app_id         = "%[3]s"
  group_type     = "OPAL_GROUP"
  admin_owner_id = "%[4]s"
  visibility     = "GLOBAL"
  %[5]s
  request_configurations = [{
    allow_requests           = true
    auto_approval            = false
    priority                 = 0
    require_mfa_to_request   = false
    require_support_ticket   = false
    reason_optional          = false
    reviewer_stages = [{
      operator = "AND"
      owner_ids = ["%[4]s"]
    }]
  }]
}
`, name, groupName, knownOpalAppID, knownOpalAppAdminOwnerID, extra)
}

func groupLinkedHCL(name, groupName, tmplRef string) string {
	return fmt.Sprintf(`
resource "opal_group" "g" {
  name                      = "%[2]s"
  app_id                    = "%[3]s"
  group_type                = "OPAL_GROUP"
  configuration_template_id = %[4]s
}
`, name, groupName, knownOpalAppID, tmplRef)
}

func resourceOwnConfigHCL(name, resName, extra string) string {
	return fmt.Sprintf(`
resource "opal_resource" "r" {
  name           = "%[2]s"
  app_id         = "%[3]s"
  resource_type  = "CUSTOM"
  admin_owner_id = "%[4]s"
  visibility     = "GLOBAL"
  %[5]s
  request_configurations = [{
    allow_requests           = true
    auto_approval            = false
    priority                 = 0
    require_mfa_to_request   = false
    require_support_ticket   = false
    reason_optional          = false
    reviewer_stages = [{
      operator = "AND"
      owner_ids = ["%[4]s"]
    }]
  }]
}
`, name, resName, knownCustomAppID, knownCustomAppAdminOwnerID, extra)
}

func resourceLinkedHCL(name, resName, tmplRef string) string {
	return fmt.Sprintf(`
resource "opal_resource" "r" {
  name                      = "%[2]s"
  app_id                    = "%[3]s"
  resource_type             = "CUSTOM"
  configuration_template_id = %[4]s
}
`, name, resName, knownCustomAppID, tmplRef)
}

func templateRef(label string) string {
	return "opal_configuration_template." + label + ".configuration_template_id"
}

func runTemplateLinkCase(t *testing.T, steps ...resource.TestStep) {
	t.Helper()
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps:                    steps,
	})
}

func groupAttachSteps(n, p string) []resource.TestStep {
	return []resource.TestStep{
		{
			Config: p + groupOwnConfigHCL(n, n, ""),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet("opal_group.g", "id"),
				resource.TestCheckNoResourceAttr("opal_group.g", "configuration_template_id"),
			),
		},
		{
			Config: p + groupLinkedHCL(n, n, templateRef("a")),
			Check:  resource.TestCheckResourceAttrSet("opal_group.g", "configuration_template_id"),
		},
	}
}

func resourceAttachSteps(n, p string) []resource.TestStep {
	return []resource.TestStep{
		{
			Config: p + resourceOwnConfigHCL(n, n, ""),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet("opal_resource.r", "id"),
				resource.TestCheckNoResourceAttr("opal_resource.r", "configuration_template_id"),
			),
		},
		{
			Config: p + resourceLinkedHCL(n, n, templateRef("a")),
			Check: resource.TestCheckResourceAttrSet(
				"opal_resource.r",
				"configuration_template_id",
			),
		},
	}
}

// Groups

func TestAccGroup_ConfigurationTemplate_CreateOwn_Update_Delete(t *testing.T) {
	n := testAccName("tmpl_g_own")
	p := testAccProviderConfig()
	runTemplateLinkCase(t,
		resource.TestStep{
			Config: p + groupOwnConfigHCL(n, n, ""),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("opal_group.g", "name", n),
				resource.TestCheckResourceAttrSet("opal_group.g", "id"),
			),
		},
		resource.TestStep{
			Config: p + groupOwnConfigHCL(n, n+"_renamed", ""),
			Check:  resource.TestCheckResourceAttr("opal_group.g", "name", n+"_renamed"),
		},
		resource.TestStep{Config: p}, // destroy
	)
}

func TestAccGroup_ConfigurationTemplate_CreateLinked(t *testing.T) {
	n := testAccName("tmpl_g_create")
	p := testAccProviderConfig() + templateHCL(n, "a")
	runTemplateLinkCase(t, resource.TestStep{
		Config: p + groupLinkedHCL(n, n, templateRef("a")),
		Check:  resource.TestCheckResourceAttrSet("opal_group.g", "configuration_template_id"),
	})
}

func TestAccGroup_ConfigurationTemplate_Attach(t *testing.T) {
	n := testAccName("tmpl_g_attach")
	p := testAccProviderConfig() + templateHCL(n, "a")
	runTemplateLinkCase(t, groupAttachSteps(n, p)...)
}

func TestAccGroup_ConfigurationTemplate_UnlinkWithConfig(t *testing.T) {
	n := testAccName("tmpl_g_unlink")
	p := testAccProviderConfig() + templateHCL(n, "a")
	runTemplateLinkCase(t, append(groupAttachSteps(n, p), resource.TestStep{
		Config: p + groupOwnConfigHCL(n, n, ""),
		Check:  resource.TestCheckNoResourceAttr("opal_group.g", "configuration_template_id"),
	})...)
}

func TestAccGroup_ConfigurationTemplate_UnlinkNullWithConfig(t *testing.T) {
	n := testAccName("tmpl_g_unlinknull")
	p := testAccProviderConfig() + templateHCL(n, "a")
	runTemplateLinkCase(t, append(groupAttachSteps(n, p), resource.TestStep{
		Config: p + groupOwnConfigHCL(n, n, "configuration_template_id = null"),
		Check:  resource.TestCheckNoResourceAttr("opal_group.g", "configuration_template_id"),
	})...)
}

// Unlinking without restating own configuration is rejected at plan time:
// exactly one of configuration_template_id and request_configurations/visibility
// must be set (see GroupConfigurationTemplateID).
func TestAccGroup_ConfigurationTemplate_UnlinkNoFields(t *testing.T) {
	n := testAccName("tmpl_g_unlinknofields")
	p := testAccProviderConfig() + templateHCL(n, "a")
	runTemplateLinkCase(t, append(groupAttachSteps(n, p), resource.TestStep{
		Config:      p + groupLinkedHCL(n, n, "null"),
		ExpectError: regexp.MustCompile(`Invalid Attribute Combination`),
	}, resource.TestStep{
		// Valid final config so the post-test destroy can plan.
		Config: p + groupLinkedHCL(n, n, templateRef("a")),
	})...)
}

func TestAccGroup_ConfigurationTemplate_RenameLinked(t *testing.T) {
	n := testAccName("tmpl_g_rename")
	p := testAccProviderConfig() + templateHCL(n, "a")
	runTemplateLinkCase(t, append(groupAttachSteps(n, p), resource.TestStep{
		Config: p + groupLinkedHCL(n, n+"_renamed", templateRef("a")),
		Check:  resource.TestCheckResourceAttr("opal_group.g", "name", n+"_renamed"),
	})...)
}

func TestAccGroup_ConfigurationTemplate_Repoint(t *testing.T) {
	n := testAccName("tmpl_g_repoint")
	p := testAccProviderConfig() + templateHCL(n, "a") + templateHCL(n, "b")
	runTemplateLinkCase(t, append(groupAttachSteps(n, p), resource.TestStep{
		Config: p + groupLinkedHCL(n, n, templateRef("b")),
		Check: resource.TestCheckResourceAttrPair(
			"opal_group.g", "configuration_template_id",
			"opal_configuration_template.b", "configuration_template_id",
		),
	})...)
}

func TestAccGroup_ConfigurationTemplate_ReadDataSource(t *testing.T) {
	n := testAccName("tmpl_g_ds")
	p := testAccProviderConfig() + templateHCL(n, "a")
	cfg := p + groupLinkedHCL(n, n, templateRef("a")) + `
data "opal_group" "g" {
  id = opal_group.g.id
}
`
	runTemplateLinkCase(t, resource.TestStep{
		Config: cfg,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrPair("data.opal_group.g", "id", "opal_group.g", "id"),
			resource.TestCheckResourceAttrPair(
				"data.opal_group.g", "configuration_template_id",
				"opal_group.g", "configuration_template_id",
			),
		),
	})
}

// Resources

func TestAccResource_ConfigurationTemplate_CreateOwn_Update_Delete(t *testing.T) {
	n := testAccName("tmpl_r_own")
	p := testAccProviderConfig()
	runTemplateLinkCase(t,
		resource.TestStep{
			Config: p + resourceOwnConfigHCL(n, n, ""),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("opal_resource.r", "name", n),
				resource.TestCheckResourceAttrSet("opal_resource.r", "id"),
			),
		},
		resource.TestStep{
			Config: p + resourceOwnConfigHCL(n, n+"_renamed", ""),
			Check:  resource.TestCheckResourceAttr("opal_resource.r", "name", n+"_renamed"),
		},
		resource.TestStep{Config: p},
	)
}

func TestAccResource_ConfigurationTemplate_CreateLinked(t *testing.T) {
	n := testAccName("tmpl_r_create")
	p := testAccProviderConfig() + templateHCL(n, "a")
	runTemplateLinkCase(t, resource.TestStep{
		Config: p + resourceLinkedHCL(n, n, templateRef("a")),
		Check:  resource.TestCheckResourceAttrSet("opal_resource.r", "configuration_template_id"),
	})
}

func TestAccResource_ConfigurationTemplate_Attach(t *testing.T) {
	n := testAccName("tmpl_r_attach")
	p := testAccProviderConfig() + templateHCL(n, "a")
	runTemplateLinkCase(t, resourceAttachSteps(n, p)...)
}

func TestAccResource_ConfigurationTemplate_UnlinkWithConfig(t *testing.T) {
	n := testAccName("tmpl_r_unlink")
	p := testAccProviderConfig() + templateHCL(n, "a")
	runTemplateLinkCase(t, append(resourceAttachSteps(n, p), resource.TestStep{
		Config: p + resourceOwnConfigHCL(n, n, ""),
		Check:  resource.TestCheckNoResourceAttr("opal_resource.r", "configuration_template_id"),
	})...)
}

func TestAccResource_ConfigurationTemplate_RenameLinked(t *testing.T) {
	n := testAccName("tmpl_r_rename")
	p := testAccProviderConfig() + templateHCL(n, "a")
	runTemplateLinkCase(t, append(resourceAttachSteps(n, p), resource.TestStep{
		Config: p + resourceLinkedHCL(n, n+"_renamed", templateRef("a")),
		Check:  resource.TestCheckResourceAttr("opal_resource.r", "name", n+"_renamed"),
	})...)
}

func TestAccResource_ConfigurationTemplate_Repoint(t *testing.T) {
	n := testAccName("tmpl_r_repoint")
	p := testAccProviderConfig() + templateHCL(n, "a") + templateHCL(n, "b")
	runTemplateLinkCase(t, append(resourceAttachSteps(n, p), resource.TestStep{
		Config: p + resourceLinkedHCL(n, n, templateRef("b")),
		Check: resource.TestCheckResourceAttrPair(
			"opal_resource.r", "configuration_template_id",
			"opal_configuration_template.b", "configuration_template_id",
		),
	})...)
}

func TestAccResource_ConfigurationTemplate_ReadDataSource(t *testing.T) {
	n := testAccName("tmpl_r_ds")
	p := testAccProviderConfig() + templateHCL(n, "a")
	cfg := p + resourceLinkedHCL(n, n, templateRef("a")) + `
data "opal_resource" "r" {
  id = opal_resource.r.id
}
`
	runTemplateLinkCase(t, resource.TestStep{
		Config: cfg,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrPair(
				"data.opal_resource.r",
				"id",
				"opal_resource.r",
				"id",
			),
			resource.TestCheckResourceAttrPair(
				"data.opal_resource.r", "configuration_template_id",
				"opal_resource.r", "configuration_template_id",
			),
		),
	})
}

// Destroying a group that is still linked to a template succeeds.
func TestAccGroup_ConfigurationTemplate_DeleteWhileLinked(t *testing.T) {
	n := testAccName("tmpl_g_dellinked")
	p := testAccProviderConfig() + templateHCL(n, "a")
	runTemplateLinkCase(t,
		resource.TestStep{Config: p + groupLinkedHCL(n, n, templateRef("a"))},
		resource.TestStep{Config: p}, // drop the group, keep the template
	)
}
