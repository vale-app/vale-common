package httpclient

type ClientConfig struct {
	Host    string
	Port    int
	Headers map[string]string
}
