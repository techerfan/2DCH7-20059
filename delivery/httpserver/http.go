package httpserver

type HttpPort interface {
	Start(port string) error
	Shutdown() error
}
