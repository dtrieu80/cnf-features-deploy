package hiveClusterConfig

type HiveConfigBuilder struct {
	SourceClusterCRs []interface{}
}

func NewHiveConfigBuilder() (*HiveConfigBuilder, error) {
	hcBuilder := HiveConfigBuilder{}

	return &hcBuilder, nil
}
