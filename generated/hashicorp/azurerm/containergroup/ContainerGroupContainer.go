package containergroup


type ContainerGroupContainer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/container_group#cpu ContainerGroup#cpu}.
	Cpu *float64 `field:"required" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/container_group#image ContainerGroup#image}.
	Image *string `field:"required" json:"image" yaml:"image"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/container_group#memory ContainerGroup#memory}.
	Memory *float64 `field:"required" json:"memory" yaml:"memory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/container_group#name ContainerGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/container_group#commands ContainerGroup#commands}.
	Commands *[]*string `field:"optional" json:"commands" yaml:"commands"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/container_group#environment_variables ContainerGroup#environment_variables}.
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// gpu block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/container_group#gpu ContainerGroup#gpu}
	Gpu *ContainerGroupContainerGpu `field:"optional" json:"gpu" yaml:"gpu"`
	// liveness_probe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/container_group#liveness_probe ContainerGroup#liveness_probe}
	LivenessProbe *ContainerGroupContainerLivenessProbe `field:"optional" json:"livenessProbe" yaml:"livenessProbe"`
	// ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/container_group#ports ContainerGroup#ports}
	Ports interface{} `field:"optional" json:"ports" yaml:"ports"`
	// readiness_probe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/container_group#readiness_probe ContainerGroup#readiness_probe}
	ReadinessProbe *ContainerGroupContainerReadinessProbe `field:"optional" json:"readinessProbe" yaml:"readinessProbe"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/container_group#secure_environment_variables ContainerGroup#secure_environment_variables}.
	SecureEnvironmentVariables *map[string]*string `field:"optional" json:"secureEnvironmentVariables" yaml:"secureEnvironmentVariables"`
	// volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/container_group#volume ContainerGroup#volume}
	Volume interface{} `field:"optional" json:"volume" yaml:"volume"`
}

