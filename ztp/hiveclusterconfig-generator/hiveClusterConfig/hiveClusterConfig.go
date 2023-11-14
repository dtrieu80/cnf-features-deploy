package hiveClusterConfig

// HiveClusterConfig
type HiveClusterConfig struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

type Metadata struct {
	Name      string            `yaml:"name"`
	Namespace string            `yaml:"namespace"`
	Labels    map[string]string `yaml:"labels"`
}

// Spec
type Spec struct {
	ClusterName       string            `yaml:"clusterName"`
	BaseDomain        string            `yaml:"baseDomain"`
	OcpVersion        string            `yaml:"ocpVersion"`
	Networking        Networking        `yaml:"networking"`
	ClusterPowerState string            `yaml:"clusterPowerState"`
	AcmClusterSetName string            `yaml:"acmClusterSetName"`
	AcmClusterLabels  map[string]string `yaml:"acmClusterLabels"`
	SshPubkey         string            `yaml:"sshPubkey"`
	Azure             Azure             `yaml:"azure"`
	VSphere           VSphere           `yaml:"vSphere"`
}

// Networking
type Networking struct {
	ClusterNetwork string `yaml:"clusterNetwork"`
	HostPrefix     string `yaml:"hostPrefix"`
	MachineNetwork string `yaml:"machineNetwork"`
	ServiceNetwork string `yaml:"serviceNetwork"`
}

// Azure
type Azure struct {
	ResourceGroup string              `yaml:"resourceGroup"`
	Region        string              `yaml:"region"`
	ControlNodes  CloudControlNodes   `yaml:"controlNodes"`
	ComputeNodes  []CloudComputeNodes `yaml:"computeNodes"`
}

// CloudControlNodes
type CloudControlNodes struct {
	Replicas     int    `yaml:"replicas"`
	Size         string `yaml:"size"`
	OsDiskSizeGB int    `yaml:"osDiskSizeGB"`
}

// CloudComputeNodes
type CloudComputeNodes struct {
	Id           int      `yaml:"id"`
	Replicas     int      `yaml:"replicas"`
	Size         string   `yaml:"size"`
	OsDiskSizeGB int      `yaml:"osDiskSizeGB"`
	Zones        []string `yaml:"zones"`
}

// VSphere
type VSphere struct {
	ApiVIP            string           `yaml:"apiVIP"`
	IngressVIP        string           `yaml:"ingressVIP"`
	ControlNodes      VirtControlNodes `yaml:"controlNodes"`
	ComputeNodes      VirtComputeNodes `yaml:"computeNodes"`
	ProviderSecretRef string           `yaml:"providerSecretRef"`
}

// VirtControlNodes
type VirtControlNodes struct {
	Replicas       int `yaml:"replicas"`
	Cpus           int `yaml:"cpus"`
	CoresPerSocket int `yaml:"coresPerSocket"`
	MemoryMB       int `yaml:"memoryMB"`
	OsDiskSizeGB   int `yaml:"osDiskSizeGB"`
}

// VirtComputeNodes
type VirtComputeNodes struct {
	Id             int      `yaml:"id"`
	Replicas       int      `yaml:"replicas"`
	Cpus           int      `yaml:"cpus"`
	CoresPerSocket int      `yaml:"coresPerSocket"`
	MemoryMB       int      `yaml:"memoryMB"`
	OsDiskSizeGB   int      `yaml:"osDiskSizeGB"`
	Zones          []string `yaml:"zones"`
}
