//go:generate go-extpoints . AdapterFactory
package bridge

import (
	"net/url"

	dockerapi "github.com/fsouza/go-dockerclient"
	"strings"
)

type AdapterFactory interface {
	New(uri *url.URL) RegistryAdapter
}

type RegistryAdapter interface {
	Ping() error
	Register(service *Service) error
	Deregister(service *Service) error
	Refresh(service *Service) error
	Services() ([]*Service, error)
}

type Config struct {
	HostIp          string
	Internal        bool
	ForceTags       string
	RefreshTtl      int
	RefreshInterval int
	DeregisterCheck string
	Cleanup         bool
}

type Service struct {
	ServiceName string
	serviceID   string

	OuterIP     string

	OuterPorts  []string
	InnerPorts  []string

	TTL         int
}

func (s *Service) GetRegisterPath() string {
	return "/" + s.ServiceName + "/" + s.serviceID
}

func (s *Service) GetRegisterData() string {
	ports := make([]string, len(s.OuterPorts))
	for i := 0; i < len(s.OuterPorts); i++ {
		ports[i] = s.OuterPorts[i] + ":" + s.InnerPorts[i]
	}
	return s.OuterIP + "," + strings.Join(ports, ",")
}

//
//type Service struct {
//	ID    string
//	Name  string
//	Port  int
//	IP    string
//	Tags  []string
//	Attrs map[string]string
//	TTL   int
//
//	Origin ServicePort
//}

type DeadContainer struct {
	TTL      int
	Services []*Service
}

type ServicePort struct {
	HostPort          string
	HostIP            string
	ExposedPort       string
	ExposedIP         string
	PortType          string
	ContainerHostname string
	ContainerID       string
	ContainerName     string
	container         *dockerapi.Container
}
