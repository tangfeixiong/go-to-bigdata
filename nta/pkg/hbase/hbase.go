package hbase

type Config struct {
	mode               string
	Kubeconfig         string
	Kind               string
	CustomResourceName string
	Name               string
	resourceHostname   string
	ServiceName        string
	Namespace          string
	BaseDomain         string
	Port               int
	Dir                string
	ClusterID          string
	NodeType           string
	LogLevel           int
}
