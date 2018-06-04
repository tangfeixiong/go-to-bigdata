package config

type Config struct {
	GrpcAddress    string
	HttpAddress    string
	SecureProtocol bool
	Kubeconfig     string
	LogLevel       int
}
