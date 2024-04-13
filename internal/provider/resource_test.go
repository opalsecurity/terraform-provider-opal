package provider

// TestAccResource_Read tests the reading of an Opal resource.
func TestAccResource_Read(t *testing.T) {
	t.Parallel()
	baseName, resourceName := generateBaseNameAndResourceName()
	config := generateSimpleOpalResourceConfig(baseName, resourceName)

	resource.Test(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_0_0),
		},
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		CheckDestroy:             testAccCheckResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: GenerateSimpleResource(&config),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", config.Name),
					resource.TestCheckResourceAttr(resourceName, "visibility", "GLOBAL"),
				),
			},
		},
	})
}

// TestAccResource_Update tests the updating of an Opal resource.
func TestAccResource_Update(t *testing.T) {
	t.Parallel()
	baseName, resourceName := generateBaseNameAndResourceName()
	config := generateSimpleOpalResourceConfig(baseName, resourceName)
	updatedConfig := generateSimpleOpalResourceConfig(baseName+"_updated", resourceName)

	resource.Test(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_0_0),
		},
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		CheckDestroy:             testAccCheckResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: GenerateSimpleResource(&config),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", config.Name),
					resource.TestCheckResourceAttr(resourceName, "visibility", "GLOBAL"),
				),
			},
			{
				Config: GenerateSimpleResource(&updatedConfig),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedConfig.Name),
					resource.TestCheckResourceAttr(resourceName, "visibility", "GLOBAL"),
				),
			},
		},
	})
}

// TestAccResource_Delete tests the deletion of an Opal resource.
func TestAccResource_Delete(t *testing.T) {
	t.Parallel()
	baseName, resourceName := generateBaseNameAndResourceName()
	config := generateSimpleOpalResourceConfig(baseName, resourceName)

	resource.Test(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_0_0),
		},
		ProtoV6ProviderFactories: testAccProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		CheckDestroy:             testAccCheckResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: GenerateSimpleResource(&config),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", config.Name),
					resource.TestCheckResourceAttr(resourceName, "visibility", "GLOBAL"),
				),
			},
			{
				Config: "",
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResourceDestroy,
				),
			},
		},
	})
}
