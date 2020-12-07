package prototype

import (
	"github.com/DerekStrickland/consul-setup-cli/reflection"
	"github.com/creasty/defaults"
)

// NewPrototype is a factory method for creating a new prototypical instance of
// the YamlValues struct with default values.
func NewPrototype() (ConsulHelmValues, error) {
	prototype := &ConsulHelmValues{}
	if err := defaults.Set(prototype); err != nil {
		return *prototype, err
	}

	return *prototype, nil
}

// AddDefaultField adds a Zero value instance of the field type
// TODO: wanted this to be an interface, but can't because go.
func AddDefaultField(target interface{}, fieldName string) error {
	var err error
	// if reflection.IsPrimitive(&target, fieldName) {
	// 	if err = reflection.SetField(&target, fieldName, ); err != nil {
	// 		return err
	// 	}
	// 	return nil
	// }

	obj := reflection.NewZeroValue(target, fieldName)
	if err := defaults.Set(obj); err != nil {
		return err
	}

	if err = reflection.SetField(target, fieldName, obj); err != nil {
		return err
	}

	return nil
}

// Gateways
type Gateways struct {
	Name string `yaml:"name"`
}

// Certs
type Certs struct {
	KeyName    string      `yaml:"keyName"`
	SecretName interface{} `yaml:"secretName"`
	CaBundle   string      `yaml:"caBundle"`
	CertName   string      `yaml:"certName"`
}

// ClientResourcesLimits
type ClientResourcesLimits struct {
	Memory string `yaml:"memory"`
	Cpu    string `yaml:"cpu"`
}

// Service
type Service struct {
	Enabled        bool        `yaml:"enabled"`
	Type           interface{} `yaml:"type"`
	Annotations    interface{} `yaml:"annotations"`
	AdditionalSpec interface{} `yaml:"additionalSpec"`
}

// InitCopyConsulContainerResourcesLimits
type InitCopyConsulContainerResourcesLimits struct {
	Memory string `yaml:"memory"`
	Cpu    string `yaml:"cpu"`
}

// LifecycleSidecarContainerResourcesLimits
type LifecycleSidecarContainerResourcesLimits struct {
	Cpu    string `yaml:"cpu"`
	Memory string `yaml:"memory"`
}

// TerminatingGatewaysDefaultsInitCopyConsulContainerResourcesLimits
type TerminatingGatewaysDefaultsInitCopyConsulContainerResourcesLimits struct {
	Memory string `yaml:"memory"`
	Cpu    string `yaml:"cpu"`
}

// ServerService
type ServerService struct {
	Annotations interface{} `yaml:"annotations"`
}

// ClientExtraEnvironmentVars
type ClientExtraEnvironmentVars struct {
}

// MeshGatewayResourcesRequests
type MeshGatewayResourcesRequests struct {
	Memory string `yaml:"memory"`
	Cpu    string `yaml:"cpu"`
}

// MeshGatewayInitCopyConsulContainer
type MeshGatewayInitCopyConsulContainer struct {
	Resources MeshGatewayInitCopyConsulContainerResources `yaml:"resources"`
}

// SyncCatalogResourcesLimits
type SyncCatalogResourcesLimits struct {
	Memory string `yaml:"memory"`
	Cpu    string `yaml:"cpu"`
}

// Controller
type Controller struct {
	PriorityClassName string              `yaml:"priorityClassName"`
	Enabled           bool                `yaml:"enabled"`
	Replicas          int                 `yaml:"replicas"`
	LogLevel          string              `yaml:"logLevel"`
	Resources         ControllerResources `yaml:"resources"`
	NodeSelector      interface{}         `yaml:"nodeSelector"`
	Tolerations       interface{}         `yaml:"tolerations"`
	Affinity          interface{}         `yaml:"affinity"`
}

// BootstrapToken
type BootstrapToken struct {
	SecretName interface{} `yaml:"secretName"`
	SecretKey  interface{} `yaml:"secretKey"`
}

// ConfigSecret
type ConfigSecret struct {
	SecretKey  interface{} `yaml:"secretKey"`
	SecretName interface{} `yaml:"secretName"`
}

// InitContainerResourcesLimits
type InitContainerResourcesLimits struct {
	Memory string `yaml:"memory"`
	Cpu    string `yaml:"cpu"`
}

// SidecarProxy
type SidecarProxy struct {
	Resources SidecarProxyResources `yaml:"resources"`
}

// SidecarProxyResources
type SidecarProxyResources struct {
	Requests SidecarProxyResourcesRequests `yaml:"requests"`
	Limits   SidecarProxyResourcesLimits   `yaml:"limits"`
}

// ControllerResourcesRequests
type ControllerResourcesRequests struct {
	Cpu    string `yaml:"cpu"`
	Memory string `yaml:"memory"`
}

// Global
type Global struct {
	GossipEncryption          GossipEncryption          `yaml:"gossipEncryption"`
	Tls                       Tls                       `yaml:"tls"`
	EnableConsulNamespaces    bool                      `yaml:"enableConsulNamespaces"`
	ImageEnvoy                string                    `yaml:"imageEnvoy"`
	Enabled                   bool                      `yaml:"enabled"`
	LifecycleSidecarContainer LifecycleSidecarContainer `yaml:"lifecycleSidecarContainer"`
	Openshift                 Openshift                 `yaml:"openshift"`
	Name                      interface{}               `yaml:"name"`
	ImageK8S                  string                    `yaml:"imageK8S"`
	Datacenter                string                    `yaml:"datacenter"`
	EnablePodSecurityPolicies bool                      `yaml:"enablePodSecurityPolicies"`
	Image                     string                    `yaml:"image"`
	Acls                      Acls                      `yaml:"acls"`
	Federation                Federation                `yaml:"federation"`
	Domain                    string                    `yaml:"domain"`
}

// ClientResourcesRequests
type ClientResourcesRequests struct {
	Memory string `yaml:"memory"`
	Cpu    string `yaml:"cpu"`
}

// SnapshotAgent
type SnapshotAgent struct {
	CaCert       interface{}            `yaml:"caCert"`
	Enabled      bool                   `yaml:"enabled"`
	Replicas     int                    `yaml:"replicas"`
	ConfigSecret ConfigSecret           `yaml:"configSecret"`
	Resources    SnapshotAgentResources `yaml:"resources"`
}

// LifecycleSidecarContainer
type LifecycleSidecarContainer struct {
	Resources LifecycleSidecarContainerResources `yaml:"resources"`
}

// DefaultsService
type DefaultsService struct {
	AdditionalSpec interface{} `yaml:"additionalSpec"`
	Type           string      `yaml:"type"`
	Ports          []Ports     `yaml:"ports"`
	Annotations    interface{} `yaml:"annotations"`
}

// TerminatingGateways
type TerminatingGateways struct {
	Enabled  bool                          `yaml:"enabled"`
	Defaults TerminatingGatewaysDefaults   `yaml:"defaults"`
	Gateways []TerminatingGatewaysGateways `yaml:"gateways"`
}

// InitContainer
type InitContainer struct {
	Resources InitContainerResources `yaml:"resources"`
}

// InitContainerResourcesRequests
type InitContainerResourcesRequests struct {
	Cpu    string `yaml:"cpu"`
	Memory string `yaml:"memory"`
}

// ConnectInjectConsulNamespaces
type ConnectInjectConsulNamespaces struct {
	MirroringK8SPrefix         string `yaml:"mirroringK8SPrefix"`
	ConsulDestinationNamespace string `yaml:"consulDestinationNamespace"`
	MirroringK8S               bool   `yaml:"mirroringK8S"`
}

// CentralConfig
type CentralConfig struct {
	Enabled         bool        `yaml:"enabled"`
	DefaultProtocol interface{} `yaml:"defaultProtocol"`
	ProxyDefaults   string      `yaml:"proxyDefaults"`
}

// Tls
type Tls struct {
	CaKey             CaKey  `yaml:"caKey"`
	Enabled           bool   `yaml:"enabled"`
	EnableAutoEncrypt bool   `yaml:"enableAutoEncrypt"`
	Verify            bool   `yaml:"verify"`
	HttpsOnly         bool   `yaml:"httpsOnly"`
	CaCert            CaCert `yaml:"caCert"`
}

// ClientResources
type ClientResources struct {
	Requests ClientResourcesRequests `yaml:"requests"`
	Limits   ClientResourcesLimits   `yaml:"limits"`
}

// MeshGatewayService
type MeshGatewayService struct {
	Enabled        bool        `yaml:"enabled"`
	Type           string      `yaml:"type"`
	Port           int         `yaml:"port"`
	NodePort       interface{} `yaml:"nodePort"`
	Annotations    interface{} `yaml:"annotations"`
	AdditionalSpec interface{} `yaml:"additionalSpec"`
}

// IngressGateways
type IngressGateways struct {
	Enabled  bool       `yaml:"enabled"`
	Defaults Defaults   `yaml:"defaults"`
	Gateways []Gateways `yaml:"gateways"`
}

// Defaults
type Defaults struct {
	Resources               Resources               `yaml:"resources"`
	NodeSelector            interface{}             `yaml:"nodeSelector"`
	Annotations             interface{}             `yaml:"annotations"`
	Affinity                string                  `yaml:"affinity"`
	Tolerations             interface{}             `yaml:"tolerations"`
	PriorityClassName       string                  `yaml:"priorityClassName"`
	ConsulNamespace         string                  `yaml:"consulNamespace"`
	Replicas                int                     `yaml:"replicas"`
	Service                 DefaultsService         `yaml:"service"`
	InitCopyConsulContainer InitCopyConsulContainer `yaml:"initCopyConsulContainer"`
}

// Requests
type Requests struct {
	Memory string `yaml:"memory"`
	Cpu    string `yaml:"cpu"`
}

// InitCopyConsulContainer
type InitCopyConsulContainer struct {
	Resources InitCopyConsulContainerResources `yaml:"resources"`
}

// CaKey
type CaKey struct {
	SecretName interface{} `yaml:"secretName"`
	SecretKey  interface{} `yaml:"secretKey"`
}

// MeshGatewayResourcesLimits
type MeshGatewayResourcesLimits struct {
	Memory string `yaml:"memory"`
	Cpu    string `yaml:"cpu"`
}

// WanAddress
type WanAddress struct {
	Static string `yaml:"static"`
	Source string `yaml:"source"`
	Port   int    `yaml:"port"`
}

// Tests
type Tests struct {
	Enabled bool `yaml:"enabled"`
}

// Ui
type Ui struct {
	Enabled string  `yaml:"enabled"`
	Service Service `yaml:"service"`
}

// Ports
type Ports struct {
	Port     int         `yaml:"port"`
	NodePort interface{} `yaml:"nodePort"`
}

// SyncCatalogResources
type SyncCatalogResources struct {
	Requests SyncCatalogResourcesRequests `yaml:"requests"`
	Limits   SyncCatalogResourcesLimits   `yaml:"limits"`
}

// ConnectInjectResourcesLimits
type ConnectInjectResourcesLimits struct {
	Memory string `yaml:"memory"`
	Cpu    string `yaml:"cpu"`
}

// EnterpriseLicense
type EnterpriseLicense struct {
	SecretName interface{} `yaml:"secretName"`
	SecretKey  interface{} `yaml:"secretKey"`
}

// Dns
type Dns struct {
	ClusterIP      interface{} `yaml:"clusterIP"`
	Annotations    interface{} `yaml:"annotations"`
	AdditionalSpec interface{} `yaml:"additionalSpec"`
	Enabled        string      `yaml:"enabled"`
	Type           string      `yaml:"type"`
}

type ConsulHelmValues struct {
	Global              Global              `yaml:"global"`
	Server              Server              `yaml:"server"`
	ExternalServers     ExternalServers     `yaml:"externalServers"`
	Client              Client              `yaml:"client"`
	Dns                 Dns                 `yaml:"dns"`
	Ui                  Ui                  `yaml:"ui"`
	SyncCatalog         SyncCatalog         `yaml:"syncCatalog"`
	ConnectInject       ConnectInject       `yaml:"connectInject"`
	Controller          Controller          `yaml:"controller"`
	MeshGateway         MeshGateway         `yaml:"meshGateway"`
	IngressGateways     IngressGateways     `yaml:"ingressGateways"`
	TerminatingGateways TerminatingGateways `yaml:"terminatingGateways"`
	Tests               Tests               `yaml:"tests"`
}

// Limits
type Limits struct {
	Memory string `yaml:"memory"`
	Cpu    string `yaml:"cpu"`
}

// InitCopyConsulContainerResources
type InitCopyConsulContainerResources struct {
	Requests InitCopyConsulContainerResourcesRequests `yaml:"requests"`
	Limits   InitCopyConsulContainerResourcesLimits   `yaml:"limits"`
}

// TerminatingGatewaysDefaults
type TerminatingGatewaysDefaults struct {
	InitCopyConsulContainer TerminatingGatewaysDefaultsInitCopyConsulContainer `yaml:"initCopyConsulContainer"`
	NodeSelector            interface{}                                        `yaml:"nodeSelector"`
	PriorityClassName       string                                             `yaml:"priorityClassName"`
	Resources               TerminatingGatewaysDefaultsResources               `yaml:"resources"`
	Affinity                string                                             `yaml:"affinity"`
	Tolerations             interface{}                                        `yaml:"tolerations"`
	Annotations             interface{}                                        `yaml:"annotations"`
	ConsulNamespace         string                                             `yaml:"consulNamespace"`
	Replicas                int                                                `yaml:"replicas"`
}

// AclSyncToken
type AclSyncToken struct {
	SecretName interface{} `yaml:"secretName"`
	SecretKey  interface{} `yaml:"secretKey"`
}

// SidecarProxyResourcesRequests
type SidecarProxyResourcesRequests struct {
	Memory interface{} `yaml:"memory"`
	Cpu    interface{} `yaml:"cpu"`
}

// Affinity
type Affinity struct {
}

// MeshGatewayInitCopyConsulContainerResourcesLimits
type MeshGatewayInitCopyConsulContainerResourcesLimits struct {
	Memory string `yaml:"memory"`
	Cpu    string `yaml:"cpu"`
}

// Resources
type Resources struct {
	Requests Requests `yaml:"requests"`
	Limits   Limits   `yaml:"limits"`
}

// ConnectInjectResourcesRequests
type ConnectInjectResourcesRequests struct {
	Memory string `yaml:"memory"`
	Cpu    string `yaml:"cpu"`
}

// LifecycleSidecarContainerResourcesRequests
type LifecycleSidecarContainerResourcesRequests struct {
	Memory string `yaml:"memory"`
	Cpu    string `yaml:"cpu"`
}

// ExtraEnvironmentVars
type ExtraEnvironmentVars struct {
}

// ExternalServers
type ExternalServers struct {
	Enabled           bool        `yaml:"enabled"`
	HttpsPort         int         `yaml:"httpsPort"`
	TlsServerName     interface{} `yaml:"tlsServerName"`
	UseSystemRoots    bool        `yaml:"useSystemRoots"`
	K8sAuthMethodHost interface{} `yaml:"k8sAuthMethodHost"`
}

// SnapshotAgentResources
type SnapshotAgentResources struct {
	Requests SnapshotAgentResourcesRequests `yaml:"requests"`
	Limits   SnapshotAgentResourcesLimits   `yaml:"limits"`
}

// SnapshotAgentResourcesLimits
type SnapshotAgentResourcesLimits struct {
	Memory string `yaml:"memory"`
	Cpu    string `yaml:"cpu"`
}

// ServerResourcesLimits
type ServerResourcesLimits struct {
	Memory string `yaml:"memory"`
	Cpu    string `yaml:"cpu"`
}

// InitCopyConsulContainerResourcesRequests
type InitCopyConsulContainerResourcesRequests struct {
	Memory string `yaml:"memory"`
	Cpu    string `yaml:"cpu"`
}

// TerminatingGatewaysDefaultsInitCopyConsulContainerResources
type TerminatingGatewaysDefaultsInitCopyConsulContainerResources struct {
	Requests TerminatingGatewaysDefaultsInitCopyConsulContainerResourcesRequests `yaml:"requests"`
	Limits   TerminatingGatewaysDefaultsInitCopyConsulContainerResourcesLimits   `yaml:"limits"`
}

// SyncCatalog
type SyncCatalog struct {
	NodeSelector          interface{}          `yaml:"nodeSelector"`
	Tolerations           interface{}          `yaml:"tolerations"`
	Default               bool                 `yaml:"default"`
	K8sDenyNamespaces     []string             `yaml:"k8sDenyNamespaces"`
	K8sTag                interface{}          `yaml:"k8sTag"`
	ConsulNodeName        string               `yaml:"consulNodeName"`
	Resources             SyncCatalogResources `yaml:"resources"`
	ConsulWriteInterval   interface{}          `yaml:"consulWriteInterval"`
	PriorityClassName     string               `yaml:"priorityClassName"`
	ToK8S                 bool                 `yaml:"toK8S"`
	K8sAllowNamespaces    []string             `yaml:"k8sAllowNamespaces"`
	ConsulPrefix          interface{}          `yaml:"consulPrefix"`
	NodePortSyncType      string               `yaml:"nodePortSyncType"`
	LogLevel              string               `yaml:"logLevel"`
	K8sPrefix             interface{}          `yaml:"k8sPrefix"`
	K8sSourceNamespace    interface{}          `yaml:"k8sSourceNamespace"`
	ConsulNamespaces      ConsulNamespaces     `yaml:"consulNamespaces"`
	AddK8SNamespaceSuffix bool                 `yaml:"addK8SNamespaceSuffix"`
	AclSyncToken          AclSyncToken         `yaml:"aclSyncToken"`
	Affinity              interface{}          `yaml:"affinity"`
	Enabled               bool                 `yaml:"enabled"`
	Image                 interface{}          `yaml:"image"`
	ToConsul              bool                 `yaml:"toConsul"`
	SyncClusterIPServices bool                 `yaml:"syncClusterIPServices"`
}

// ConnectInject
type ConnectInject struct {
	InitContainer          InitContainer                 `yaml:"initContainer"`
	Enabled                bool                          `yaml:"enabled"`
	Default                bool                          `yaml:"default"`
	EnvoyExtraArgs         interface{}                   `yaml:"envoyExtraArgs"`
	NamespaceSelector      interface{}                   `yaml:"namespaceSelector"`
	ConsulNamespaces       ConnectInjectConsulNamespaces `yaml:"consulNamespaces"`
	Tolerations            interface{}                   `yaml:"tolerations"`
	AclInjectToken         AclInjectToken                `yaml:"aclInjectToken"`
	PriorityClassName      string                        `yaml:"priorityClassName"`
	LogLevel               string                        `yaml:"logLevel"`
	Certs                  Certs                         `yaml:"certs"`
	Affinity               interface{}                   `yaml:"affinity"`
	OverrideAuthMethodName string                        `yaml:"overrideAuthMethodName"`
	ImageConsul            interface{}                   `yaml:"imageConsul"`
	Resources              ConnectInjectResources        `yaml:"resources"`
	NodeSelector           interface{}                   `yaml:"nodeSelector"`
	SidecarProxy           SidecarProxy                  `yaml:"sidecarProxy"`
	Image                  interface{}                   `yaml:"image"`
	HealthChecks           HealthChecks                  `yaml:"healthChecks"`
	K8sAllowNamespaces     []string                      `yaml:"k8sAllowNamespaces"`
	AclBindingRuleSelector string                        `yaml:"aclBindingRuleSelector"`
	CentralConfig          CentralConfig                 `yaml:"centralConfig"`
}

// AclInjectToken
type AclInjectToken struct {
	SecretName interface{} `yaml:"secretName"`
	SecretKey  interface{} `yaml:"secretKey"`
}

// Acls
type Acls struct {
	ManageSystemACLs       bool             `yaml:"manageSystemACLs"`
	BootstrapToken         BootstrapToken   `yaml:"bootstrapToken"`
	CreateReplicationToken bool             `yaml:"createReplicationToken"`
	ReplicationToken       ReplicationToken `yaml:"replicationToken"`
}

// ReplicationToken
type ReplicationToken struct {
	SecretName interface{} `yaml:"secretName"`
	SecretKey  interface{} `yaml:"secretKey"`
}

// MeshGatewayResources
type MeshGatewayResources struct {
	Requests MeshGatewayResourcesRequests `yaml:"requests"`
	Limits   MeshGatewayResourcesLimits   `yaml:"limits"`
}

// MeshGatewayInitCopyConsulContainerResources
type MeshGatewayInitCopyConsulContainerResources struct {
	Requests MeshGatewayInitCopyConsulContainerResourcesRequests `yaml:"requests"`
	Limits   MeshGatewayInitCopyConsulContainerResourcesLimits   `yaml:"limits"`
}

// TerminatingGatewaysDefaultsInitCopyConsulContainer
type TerminatingGatewaysDefaultsInitCopyConsulContainer struct {
	Resources TerminatingGatewaysDefaultsInitCopyConsulContainerResources `yaml:"resources"`
}

// TerminatingGatewaysDefaultsResources
type TerminatingGatewaysDefaultsResources struct {
	Requests TerminatingGatewaysDefaultsResourcesRequests `yaml:"requests"`
	Limits   TerminatingGatewaysDefaultsResourcesLimits   `yaml:"limits"`
}

// TerminatingGatewaysDefaultsResourcesLimits
type TerminatingGatewaysDefaultsResourcesLimits struct {
	Memory string `yaml:"memory"`
	Cpu    string `yaml:"cpu"`
}

// CaCert
type CaCert struct {
	SecretName interface{} `yaml:"secretName"`
	SecretKey  interface{} `yaml:"secretKey"`
}

// TerminatingGatewaysDefaultsResourcesRequests
type TerminatingGatewaysDefaultsResourcesRequests struct {
	Memory string `yaml:"memory"`
	Cpu    string `yaml:"cpu"`
}

// TerminatingGatewaysGateways
type TerminatingGatewaysGateways struct {
}

// InitContainerResources
type InitContainerResources struct {
	Requests InitContainerResourcesRequests `yaml:"requests"`
	Limits   InitContainerResourcesLimits   `yaml:"limits"`
}

// Server
type Server struct {
	Image                interface{}          `yaml:"image"`
	Replicas             int                  `yaml:"replicas"`
	EnterpriseLicense    EnterpriseLicense    `yaml:"enterpriseLicense"`
	Tolerations          string               `yaml:"tolerations"`
	Enabled              string               `yaml:"enabled"`
	Connect              bool                 `yaml:"connect"`
	Resources            ServerResources      `yaml:"resources"`
	DisruptionBudget     DisruptionBudget     `yaml:"disruptionBudget"`
	ExtraConfig          string               `yaml:"extraConfig"`
	Affinity             string               `yaml:"affinity"`
	NodeSelector         interface{}          `yaml:"nodeSelector"`
	Annotations          interface{}          `yaml:"annotations"`
	StorageClass         interface{}          `yaml:"storageClass"`
	Service              ServerService        `yaml:"service"`
	UpdatePartition      int                  `yaml:"updatePartition"`
	PriorityClassName    string               `yaml:"priorityClassName"`
	ExtraLabels          interface{}          `yaml:"extraLabels"`
	ExtraEnvironmentVars ExtraEnvironmentVars `yaml:"extraEnvironmentVars"`
	Storage              string               `yaml:"storage"`
	BootstrapExpect      int                  `yaml:"bootstrapExpect"`
}

// ServerResources
type ServerResources struct {
	Limits   ServerResourcesLimits   `yaml:"limits"`
	Requests ServerResourcesRequests `yaml:"requests"`
}

// Client
type Client struct {
	DataDirectoryHostPath interface{}                `yaml:"dataDirectoryHostPath"`
	ExposeGossipPorts     bool                       `yaml:"exposeGossipPorts"`
	ExtraConfig           string                     `yaml:"extraConfig"`
	Affinity              Affinity                   `yaml:"affinity"`
	UpdateStrategy        interface{}                `yaml:"updateStrategy"`
	Resources             ClientResources            `yaml:"resources"`
	PriorityClassName     string                     `yaml:"priorityClassName"`
	DnsPolicy             interface{}                `yaml:"dnsPolicy"`
	SnapshotAgent         SnapshotAgent              `yaml:"snapshotAgent"`
	Enabled               string                     `yaml:"enabled"`
	Join                  interface{}                `yaml:"join"`
	Annotations           interface{}                `yaml:"annotations"`
	HostNetwork           bool                       `yaml:"hostNetwork"`
	Image                 interface{}                `yaml:"image"`
	Grpc                  bool                       `yaml:"grpc"`
	Tolerations           string                     `yaml:"tolerations"`
	NodeSelector          interface{}                `yaml:"nodeSelector"`
	ExtraEnvironmentVars  ClientExtraEnvironmentVars `yaml:"extraEnvironmentVars"`
}

// ConsulNamespaces
type ConsulNamespaces struct {
	ConsulDestinationNamespace string `yaml:"consulDestinationNamespace"`
	MirroringK8S               bool   `yaml:"mirroringK8S"`
	MirroringK8SPrefix         string `yaml:"mirroringK8SPrefix"`
}

// ConnectInjectResources
type ConnectInjectResources struct {
	Requests ConnectInjectResourcesRequests `yaml:"requests"`
	Limits   ConnectInjectResourcesLimits   `yaml:"limits"`
}

// HealthChecks
type HealthChecks struct {
	Enabled         bool   `yaml:"enabled"`
	ReconcilePeriod string `yaml:"reconcilePeriod"`
}

// DisruptionBudget
type DisruptionBudget struct {
	Enabled        bool        `yaml:"enabled"`
	MaxUnavailable interface{} `yaml:"maxUnavailable"`
}

// SidecarProxyResourcesLimits
type SidecarProxyResourcesLimits struct {
	Memory interface{} `yaml:"memory"`
	Cpu    interface{} `yaml:"cpu"`
}

// LifecycleSidecarContainerResources
type LifecycleSidecarContainerResources struct {
	Requests LifecycleSidecarContainerResourcesRequests `yaml:"requests"`
	Limits   LifecycleSidecarContainerResourcesLimits   `yaml:"limits"`
}

// Openshift
type Openshift struct {
	Enabled bool `yaml:"enabled"`
}

// Federation
type Federation struct {
	Enabled                bool `yaml:"enabled"`
	CreateFederationSecret bool `yaml:"createFederationSecret"`
}

// ServerResourcesRequests
type ServerResourcesRequests struct {
	Cpu    string `yaml:"cpu"`
	Memory string `yaml:"memory"`
}

// MeshGatewayInitCopyConsulContainerResourcesRequests
type MeshGatewayInitCopyConsulContainerResourcesRequests struct {
	Memory string `yaml:"memory"`
	Cpu    string `yaml:"cpu"`
}

// TerminatingGatewaysDefaultsInitCopyConsulContainerResourcesRequests
type TerminatingGatewaysDefaultsInitCopyConsulContainerResourcesRequests struct {
	Memory string `yaml:"memory"`
	Cpu    string `yaml:"cpu"`
}

// SyncCatalogResourcesRequests
type SyncCatalogResourcesRequests struct {
	Memory string `yaml:"memory"`
	Cpu    string `yaml:"cpu"`
}

// ControllerResources
type ControllerResources struct {
	Requests ControllerResourcesRequests `yaml:"requests"`
	Limits   ControllerResourcesLimits   `yaml:"limits"`
}

// ControllerResourcesLimits
type ControllerResourcesLimits struct {
	Cpu    string `yaml:"cpu"`
	Memory string `yaml:"memory"`
}

// GossipEncryption
type GossipEncryption struct {
	SecretKey  string `yaml:"secretKey"`
	SecretName string `yaml:"secretName"`
}

// SnapshotAgentResourcesRequests
type SnapshotAgentResourcesRequests struct {
	Memory string `yaml:"memory"`
	Cpu    string `yaml:"cpu"`
}

// MeshGateway
type MeshGateway struct {
	Resources               MeshGatewayResources               `yaml:"resources"`
	Annotations             interface{}                        `yaml:"annotations"`
	Enabled                 bool                               `yaml:"enabled"`
	HostNetwork             bool                               `yaml:"hostNetwork"`
	PriorityClassName       string                             `yaml:"priorityClassName"`
	Affinity                string                             `yaml:"affinity"`
	NodeSelector            interface{}                        `yaml:"nodeSelector"`
	WanAddress              WanAddress                         `yaml:"wanAddress"`
	DnsPolicy               interface{}                        `yaml:"dnsPolicy"`
	HostPort                interface{}                        `yaml:"hostPort"`
	InitCopyConsulContainer MeshGatewayInitCopyConsulContainer `yaml:"initCopyConsulContainer"`
	GlobalMode              string                             `yaml:"globalMode"`
	Replicas                int                                `yaml:"replicas"`
	ContainerPort           int                                `yaml:"containerPort"`
	Tolerations             interface{}                        `yaml:"tolerations"`
	Service                 MeshGatewayService                 `yaml:"service"`
	ConsulServiceName       string                             `yaml:"consulServiceName"`
}

// // YamlValues is a strongly typed representations of the possible helm chart values
// // that a user could specify during, say a setup wizard process. Ideally, the wizard
// // would iterate over each top level struct member, and then ask the user to choose
// // and options.  Then recursively, it
// type ConsulHelmValues struct {
// 	Global struct {
// 		Enabled                   bool          `yaml:"enabled,omitempty"`
// 		Name                      interface{}   `yaml:"name,omitempty"`
// 		Domain                    string        `yaml:"domain,omitempty"`
// 		Image                     string        `yaml:"image,omitempty"`
// 		ImagePullSecrets          []interface{} `yaml:"imagePullSecrets,omitempty"`
// 		ImageK8S                  string        `yaml:"imageK8S,omitempty"`
// 		Datacenter                string        `yaml:"datacenter,omitempty"`
// 		EnablePodSecurityPolicies bool          `yaml:"enablePodSecurityPolicies,omitempty"`
// 		GossipEncryption          struct {
// 			SecretName string `yaml:"secretName,omitempty"`
// 			SecretKey  string `yaml:"secretKey,omitempty"`
// 		} `yaml:"gossipEncryption,omitempty"`
// 		TLS struct {
// 			Enabled                 bool          `yaml:"enabled,omitempty"`
// 			EnableAutoEncrypt       bool          `yaml:"enableAutoEncrypt,omitempty"`
// 			ServerAdditionalDNSSANs []interface{} `yaml:"serverAdditionalDNSSANs,omitempty"`
// 			ServerAdditionalIPSANs  []interface{} `yaml:"serverAdditionalIPSANs,omitempty"`
// 			Verify                  bool          `yaml:"verify,omitempty"`
// 			HTTPSOnly               bool          `yaml:"httpsOnly,omitempty"`
// 			CaCert                  struct {
// 				SecretName interface{} `yaml:"secretName,omitempty"`
// 				SecretKey  interface{} `yaml:"secretKey,omitempty"`
// 			} `yaml:"caCert,omitempty"`
// 			CaKey struct {
// 				SecretName interface{} `yaml:"secretName,omitempty"`
// 				SecretKey  interface{} `yaml:"secretKey,omitempty"`
// 			} `yaml:"caKey,omitempty"`
// 		} `yaml:"tls,omitempty"`
// 		EnableConsulNamespaces bool `yaml:"enableConsulNamespaces,omitempty"`
// 		Acls                   struct {
// 			ManageSystemACLs bool `yaml:"manageSystemACLs,omitempty"`
// 			BootstrapToken   struct {
// 				SecretName interface{} `yaml:"secretName,omitempty"`
// 				SecretKey  interface{} `yaml:"secretKey,omitempty"`
// 			} `yaml:"bootstrapToken,omitempty"`
// 			CreateReplicationToken bool `yaml:"createReplicationToken,omitempty"`
// 			ReplicationToken       struct {
// 				SecretName interface{} `yaml:"secretName,omitempty"`
// 				SecretKey  interface{} `yaml:"secretKey,omitempty"`
// 			} `yaml:"replicationToken,omitempty"`
// 		} `yaml:"acls,omitempty"`
// 		Federation struct {
// 			Enabled                bool `yaml:"enabled,omitempty"`
// 			CreateFederationSecret bool `yaml:"createFederationSecret,omitempty"`
// 		} `yaml:"federation,omitempty"`
// 		LifecycleSidecarContainer struct {
// 			Resources struct {
// 				Requests struct {
// 					Memory string `yaml:"memory,omitempty"`
// 					CPU    string `yaml:"cpu,omitempty"`
// 				} `yaml:"requests,omitempty"`
// 				Limits struct {
// 					Memory string `yaml:"memory,omitempty"`
// 					CPU    string `yaml:"cpu,omitempty"`
// 				} `yaml:"limits,omitempty"`
// 			} `yaml:"resources,omitempty"`
// 		} `yaml:"lifecycleSidecarContainer,omitempty"`
// 		ImageEnvoy string `yaml:"imageEnvoy,omitempty"`
// 		Openshift  struct {
// 			Enabled bool `yaml:"enabled,omitempty"`
// 		} `yaml:"openshift,omitempty"`
// 	} `yaml:"global,omitempty"`
// 	Server struct {
// 		Enabled           string      `yaml:"enabled,omitempty"`
// 		Image             interface{} `yaml:"image,omitempty"`
// 		Replicas          int         `yaml:"replicas,omitempty"`
// 		BootstrapExpect   int         `yaml:"bootstrapExpect,omitempty"`
// 		EnterpriseLicense struct {
// 			SecretName interface{} `yaml:"secretName,omitempty"`
// 			SecretKey  interface{} `yaml:"secretKey,omitempty"`
// 		} `yaml:"enterpriseLicense,omitempty"`
// 		Storage      string      `yaml:"storage,omitempty"`
// 		StorageClass interface{} `yaml:"storageClass,omitempty"`
// 		Connect      bool        `yaml:"connect,omitempty"`
// 		Resources    struct {
// 			Requests struct {
// 				Memory string `yaml:"memory,omitempty"`
// 				CPU    string `yaml:"cpu,omitempty"`
// 			} `yaml:"requests,omitempty"`
// 			Limits struct {
// 				Memory string `yaml:"memory,omitempty"`
// 				CPU    string `yaml:"cpu,omitempty"`
// 			} `yaml:"limits,omitempty"`
// 		} `yaml:"resources,omitempty"`
// 		UpdatePartition  int `yaml:"updatePartition,omitempty"`
// 		DisruptionBudget struct {
// 			Enabled        bool        `yaml:"enabled,omitempty"`
// 			MaxUnavailable interface{} `yaml:"maxUnavailable,omitempty"`
// 		} `yaml:"disruptionBudget,omitempty"`
// 		ExtraConfig       string        `yaml:"extraConfig,omitempty"`
// 		ExtraVolumes      []interface{} `yaml:"extraVolumes,omitempty"`
// 		Affinity          string        `yaml:"affinity,omitempty"`
// 		Tolerations       string        `yaml:"tolerations,omitempty"`
// 		NodeSelector      interface{}   `yaml:"nodeSelector,omitempty"`
// 		PriorityClassName string        `yaml:"priorityClassName,omitempty"`
// 		ExtraLabels       interface{}   `yaml:"extraLabels,omitempty"`
// 		Annotations       interface{}   `yaml:"annotations,omitempty"`
// 		Service           struct {
// 			Annotations interface{} `yaml:"annotations,omitempty"`
// 		} `yaml:"service,omitempty"`
// 		ExtraEnvironmentVars struct {
// 		} `yaml:"extraEnvironmentVars,omitempty"`
// 	} `yaml:"server,omitempty"`
// 	ExternalServers struct {
// 		Enabled           bool          `yaml:"enabled,omitempty"`
// 		Hosts             []interface{} `yaml:"hosts,omitempty"`
// 		HTTPSPort         int           `yaml:"httpsPort,omitempty"`
// 		TLSServerName     interface{}   `yaml:"tlsServerName,omitempty"`
// 		UseSystemRoots    bool          `yaml:"useSystemRoots,omitempty"`
// 		K8SAuthMethodHost interface{}   `yaml:"k8sAuthMethodHost,omitempty"`
// 	} `yaml:"externalServers,omitempty"`
// 	Client struct {
// 		Enabled               string      `yaml:"enabled,omitempty"`
// 		Image                 interface{} `yaml:"image,omitempty"`
// 		Join                  interface{} `yaml:"join,omitempty"`
// 		DataDirectoryHostPath interface{} `yaml:"dataDirectoryHostPath,omitempty"`
// 		Grpc                  bool        `yaml:"grpc,omitempty"`
// 		ExposeGossipPorts     bool        `yaml:"exposeGossipPorts,omitempty"`
// 		Resources             struct {
// 			Requests struct {
// 				Memory string `yaml:"memory,omitempty"`
// 				CPU    string `yaml:"cpu,omitempty"`
// 			} `yaml:"requests,omitempty"`
// 			Limits struct {
// 				Memory string `yaml:"memory,omitempty"`
// 				CPU    string `yaml:"cpu,omitempty"`
// 			} `yaml:"limits,omitempty"`
// 		} `yaml:"resources,omitempty"`
// 		ExtraConfig  string        `yaml:"extraConfig,omitempty"`
// 		ExtraVolumes []interface{} `yaml:"extraVolumes,omitempty"`
// 		Tolerations  string        `yaml:"tolerations,omitempty"`
// 		NodeSelector interface{}   `yaml:"nodeSelector,omitempty"`
// 		Affinity     struct {
// 		} `yaml:"affinity,omitempty"`
// 		PriorityClassName    string      `yaml:"priorityClassName,omitempty"`
// 		Annotations          interface{} `yaml:"annotations,omitempty"`
// 		ExtraEnvironmentVars struct {
// 		} `yaml:"extraEnvironmentVars,omitempty"`
// 		DNSPolicy      interface{} `yaml:"dnsPolicy,omitempty"`
// 		HostNetwork    bool        `yaml:"hostNetwork,omitempty"`
// 		UpdateStrategy interface{} `yaml:"updateStrategy,omitempty"`
// 		SnapshotAgent  struct {
// 			Enabled      bool `yaml:"enabled,omitempty"`
// 			Replicas     int  `yaml:"replicas,omitempty"`
// 			ConfigSecret struct {
// 				SecretName interface{} `yaml:"secretName,omitempty"`
// 				SecretKey  interface{} `yaml:"secretKey,omitempty"`
// 			} `yaml:"configSecret,omitempty"`
// 			Resources struct {
// 				Requests struct {
// 					Memory string `yaml:"memory,omitempty"`
// 					CPU    string `yaml:"cpu,omitempty"`
// 				} `yaml:"requests,omitempty"`
// 				Limits struct {
// 					Memory string `yaml:"memory,omitempty"`
// 					CPU    string `yaml:"cpu,omitempty"`
// 				} `yaml:"limits,omitempty"`
// 			} `yaml:"resources,omitempty"`
// 			CaCert interface{} `yaml:"caCert,omitempty"`
// 		} `yaml:"snapshotAgent,omitempty"`
// 	} `yaml:"client,omitempty"`
// 	DNS struct {
// 		Enabled        string      `yaml:"enabled,omitempty"`
// 		Type           string      `yaml:"type,omitempty"`
// 		ClusterIP      interface{} `yaml:"clusterIP,omitempty"`
// 		Annotations    interface{} `yaml:"annotations,omitempty"`
// 		AdditionalSpec interface{} `yaml:"additionalSpec,omitempty"`
// 	} `yaml:"dns,omitempty"`
// 	UI struct {
// 		Enabled string `yaml:"enabled,omitempty"`
// 		Service struct {
// 			Enabled        bool        `yaml:"enabled,omitempty"`
// 			Type           interface{} `yaml:"type,omitempty"`
// 			Annotations    interface{} `yaml:"annotations,omitempty"`
// 			AdditionalSpec interface{} `yaml:"additionalSpec,omitempty"`
// 		} `yaml:"service,omitempty"`
// 	} `yaml:"ui,omitempty"`
// 	SyncCatalog struct {
// 		Enabled            bool        `yaml:"enabled,omitempty"`
// 		Image              interface{} `yaml:"image,omitempty"`
// 		Default            bool        `yaml:"default,omitempty"`
// 		PriorityClassName  string      `yaml:"priorityClassName,omitempty"`
// 		ToConsul           bool        `yaml:"toConsul,omitempty"`
// 		ToK8S              bool        `yaml:"toK8S,omitempty"`
// 		K8SPrefix          interface{} `yaml:"k8sPrefix,omitempty"`
// 		K8SAllowNamespaces []string    `yaml:"k8sAllowNamespaces,omitempty"`
// 		K8SDenyNamespaces  []string    `yaml:"k8sDenyNamespaces,omitempty"`
// 		K8SSourceNamespace interface{} `yaml:"k8sSourceNamespace,omitempty"`
// 		ConsulNamespaces   struct {
// 			ConsulDestinationNamespace string `yaml:"consulDestinationNamespace,omitempty"`
// 			MirroringK8S               bool   `yaml:"mirroringK8S,omitempty"`
// 			MirroringK8SPrefix         string `yaml:"mirroringK8SPrefix,omitempty"`
// 		} `yaml:"consulNamespaces,omitempty"`
// 		AddK8SNamespaceSuffix bool        `yaml:"addK8SNamespaceSuffix,omitempty"`
// 		ConsulPrefix          interface{} `yaml:"consulPrefix,omitempty"`
// 		K8STag                interface{} `yaml:"k8sTag,omitempty"`
// 		ConsulNodeName        string      `yaml:"consulNodeName,omitempty"`
// 		SyncClusterIPServices bool        `yaml:"syncClusterIPServices,omitempty"`
// 		NodePortSyncType      string      `yaml:"nodePortSyncType,omitempty"`
// 		ACLSyncToken          struct {
// 			SecretName interface{} `yaml:"secretName,omitempty"`
// 			SecretKey  interface{} `yaml:"secretKey,omitempty"`
// 		} `yaml:"aclSyncToken,omitempty"`
// 		NodeSelector interface{} `yaml:"nodeSelector,omitempty"`
// 		Affinity     interface{} `yaml:"affinity,omitempty"`
// 		Tolerations  interface{} `yaml:"tolerations,omitempty"`
// 		Resources    struct {
// 			Requests struct {
// 				Memory string `yaml:"memory,omitempty"`
// 				CPU    string `yaml:"cpu,omitempty"`
// 			} `yaml:"requests,omitempty"`
// 			Limits struct {
// 				Memory string `yaml:"memory,omitempty"`
// 				CPU    string `yaml:"cpu,omitempty"`
// 			} `yaml:"limits,omitempty"`
// 		} `yaml:"resources,omitempty"`
// 		LogLevel            string      `yaml:"logLevel,omitempty"`
// 		ConsulWriteInterval interface{} `yaml:"consulWriteInterval,omitempty"`
// 	} `yaml:"syncCatalog,omitempty"`
// 	ConnectInject struct {
// 		Enabled      bool        `yaml:"enabled,omitempty"`
// 		Image        interface{} `yaml:"image,omitempty"`
// 		Default      bool        `yaml:"default,omitempty"`
// 		HealthChecks struct {
// 			Enabled         bool   `yaml:"enabled,omitempty"`
// 			ReconcilePeriod string `yaml:"reconcilePeriod,omitempty"`
// 		} `yaml:"healthChecks,omitempty"`
// 		EnvoyExtraArgs    interface{} `yaml:"envoyExtraArgs,omitempty"`
// 		PriorityClassName string      `yaml:"priorityClassName,omitempty"`
// 		ImageConsul       interface{} `yaml:"imageConsul,omitempty"`
// 		LogLevel          string      `yaml:"logLevel,omitempty"`
// 		Resources         struct {
// 			Requests struct {
// 				Memory string `yaml:"memory,omitempty"`
// 				CPU    string `yaml:"cpu,omitempty"`
// 			} `yaml:"requests,omitempty"`
// 			Limits struct {
// 				Memory string `yaml:"memory,omitempty"`
// 				CPU    string `yaml:"cpu,omitempty"`
// 			} `yaml:"limits,omitempty"`
// 		} `yaml:"resources,omitempty"`
// 		NamespaceSelector  interface{}   `yaml:"namespaceSelector,omitempty"`
// 		K8SAllowNamespaces []string      `yaml:"k8sAllowNamespaces,omitempty"`
// 		K8SDenyNamespaces  []interface{} `yaml:"k8sDenyNamespaces,omitempty"`
// 		ConsulNamespaces   struct {
// 			ConsulDestinationNamespace string `yaml:"consulDestinationNamespace,omitempty"`
// 			MirroringK8S               bool   `yaml:"mirroringK8S,omitempty"`
// 			MirroringK8SPrefix         string `yaml:"mirroringK8SPrefix,omitempty"`
// 		} `yaml:"consulNamespaces,omitempty"`
// 		Certs struct {
// 			SecretName interface{} `yaml:"secretName,omitempty"`
// 			CaBundle   string      `yaml:"caBundle,omitempty"`
// 			CertName   string      `yaml:"certName,omitempty"`
// 			KeyName    string      `yaml:"keyName,omitempty"`
// 		} `yaml:"certs,omitempty"`
// 		NodeSelector           interface{} `yaml:"nodeSelector,omitempty"`
// 		Affinity               interface{} `yaml:"affinity,omitempty"`
// 		Tolerations            interface{} `yaml:"tolerations,omitempty"`
// 		ACLBindingRuleSelector string      `yaml:"aclBindingRuleSelector,omitempty"`
// 		OverrideAuthMethodName string      `yaml:"overrideAuthMethodName,omitempty"`
// 		ACLInjectToken         struct {
// 			SecretName interface{} `yaml:"secretName,omitempty"`
// 			SecretKey  interface{} `yaml:"secretKey,omitempty"`
// 		} `yaml:"aclInjectToken,omitempty"`
// 		CentralConfig struct {
// 			Enabled         bool        `yaml:"enabled,omitempty"`
// 			DefaultProtocol interface{} `yaml:"defaultProtocol,omitempty"`
// 			ProxyDefaults   string      `yaml:"proxyDefaults,omitempty"`
// 		} `yaml:"centralConfig,omitempty"`
// 		SidecarProxy struct {
// 			Resources struct {
// 				Requests struct {
// 					Memory interface{} `yaml:"memory,omitempty"`
// 					CPU    interface{} `yaml:"cpu,omitempty"`
// 				} `yaml:"requests,omitempty"`
// 				Limits struct {
// 					Memory interface{} `yaml:"memory,omitempty"`
// 					CPU    interface{} `yaml:"cpu,omitempty"`
// 				} `yaml:"limits,omitempty"`
// 			} `yaml:"resources,omitempty"`
// 		} `yaml:"sidecarProxy,omitempty"`
// 		InitContainer struct {
// 			Resources struct {
// 				Requests struct {
// 					Memory string `yaml:"memory,omitempty"`
// 					CPU    string `yaml:"cpu,omitempty"`
// 				} `yaml:"requests,omitempty"`
// 				Limits struct {
// 					Memory string `yaml:"memory,omitempty"`
// 					CPU    string `yaml:"cpu,omitempty"`
// 				} `yaml:"limits,omitempty"`
// 			} `yaml:"resources,omitempty"`
// 		} `yaml:"initContainer,omitempty"`
// 	} `yaml:"connectInject,omitempty"`
// 	Controller struct {
// 		Enabled   bool   `yaml:"enabled,omitempty"`
// 		Replicas  int    `yaml:"replicas,omitempty"`
// 		LogLevel  string `yaml:"logLevel,omitempty"`
// 		Resources struct {
// 			Limits struct {
// 				CPU    string `yaml:"cpu,omitempty"`
// 				Memory string `yaml:"memory,omitempty"`
// 			} `yaml:"limits,omitempty"`
// 			Requests struct {
// 				CPU    string `yaml:"cpu,omitempty"`
// 				Memory string `yaml:"memory,omitempty"`
// 			} `yaml:"requests,omitempty"`
// 		} `yaml:"resources,omitempty"`
// 		NodeSelector      interface{} `yaml:"nodeSelector,omitempty"`
// 		Tolerations       interface{} `yaml:"tolerations,omitempty"`
// 		Affinity          interface{} `yaml:"affinity,omitempty"`
// 		PriorityClassName string      `yaml:"priorityClassName,omitempty"`
// 	} `yaml:"controller,omitempty"`
// 	MeshGateway struct {
// 		Enabled    bool   `yaml:"enabled,omitempty"`
// 		GlobalMode string `yaml:"globalMode,omitempty"`
// 		Replicas   int    `yaml:"replicas,omitempty"`
// 		WanAddress struct {
// 			Source string `yaml:"source,omitempty"`
// 			Port   int    `yaml:"port,omitempty"`
// 			Static string `yaml:"static,omitempty"`
// 		} `yaml:"wanAddress,omitempty"`
// 		Service struct {
// 			Enabled        bool        `yaml:"enabled,omitempty"`
// 			Type           string      `yaml:"type,omitempty"`
// 			Port           int         `yaml:"port,omitempty"`
// 			NodePort       interface{} `yaml:"nodePort,omitempty"`
// 			Annotations    interface{} `yaml:"annotations,omitempty"`
// 			AdditionalSpec interface{} `yaml:"additionalSpec,omitempty"`
// 		} `yaml:"service,omitempty"`
// 		HostNetwork       bool        `yaml:"hostNetwork,omitempty"`
// 		DNSPolicy         interface{} `yaml:"dnsPolicy,omitempty"`
// 		ConsulServiceName string      `yaml:"consulServiceName,omitempty"`
// 		ContainerPort     int         `yaml:"containerPort,omitempty"`
// 		HostPort          interface{} `yaml:"hostPort,omitempty"`
// 		Resources         struct {
// 			Requests struct {
// 				Memory string `yaml:"memory,omitempty"`
// 				CPU    string `yaml:"cpu,omitempty"`
// 			} `yaml:"requests,omitempty"`
// 			Limits struct {
// 				Memory string `yaml:"memory,omitempty"`
// 				CPU    string `yaml:"cpu,omitempty"`
// 			} `yaml:"limits,omitempty"`
// 		} `yaml:"resources,omitempty"`
// 		InitCopyConsulContainer struct {
// 			Resources struct {
// 				Requests struct {
// 					Memory string `yaml:"memory,omitempty"`
// 					CPU    string `yaml:"cpu,omitempty"`
// 				} `yaml:"requests,omitempty"`
// 				Limits struct {
// 					Memory string `yaml:"memory,omitempty"`
// 					CPU    string `yaml:"cpu,omitempty"`
// 				} `yaml:"limits,omitempty"`
// 			} `yaml:"resources,omitempty"`
// 		} `yaml:"initCopyConsulContainer,omitempty"`
// 		Affinity          string      `yaml:"affinity,omitempty"`
// 		Tolerations       interface{} `yaml:"tolerations,omitempty"`
// 		NodeSelector      interface{} `yaml:"nodeSelector,omitempty"`
// 		PriorityClassName string      `yaml:"priorityClassName,omitempty"`
// 		Annotations       interface{} `yaml:"annotations,omitempty"`
// 	} `yaml:"meshGateway,omitempty"`
// 	IngressGateways struct {
// 		Enabled  bool `yaml:"enabled,omitempty"`
// 		Defaults struct {
// 			Replicas int `yaml:"replicas,omitempty"`
// 			Service  struct {
// 				Type  string `yaml:"type,omitempty"`
// 				Ports []struct {
// 					Port     int         `yaml:"port,omitempty"`
// 					NodePort interface{} `yaml:"nodePort,omitempty"`
// 				} `yaml:"ports,omitempty"`
// 				Annotations    interface{} `yaml:"annotations,omitempty"`
// 				AdditionalSpec interface{} `yaml:"additionalSpec,omitempty"`
// 			} `yaml:"service,omitempty"`
// 			Resources struct {
// 				Requests struct {
// 					Memory string `yaml:"memory,omitempty"`
// 					CPU    string `yaml:"cpu,omitempty"`
// 				} `yaml:"requests,omitempty"`
// 				Limits struct {
// 					Memory string `yaml:"memory,omitempty"`
// 					CPU    string `yaml:"cpu,omitempty"`
// 				} `yaml:"limits,omitempty"`
// 			} `yaml:"resources,omitempty"`
// 			InitCopyConsulContainer struct {
// 				Resources struct {
// 					Requests struct {
// 						Memory string `yaml:"memory,omitempty"`
// 						CPU    string `yaml:"cpu,omitempty"`
// 					} `yaml:"requests,omitempty"`
// 					Limits struct {
// 						Memory string `yaml:"memory,omitempty"`
// 						CPU    string `yaml:"cpu,omitempty"`
// 					} `yaml:"limits,omitempty"`
// 				} `yaml:"resources,omitempty"`
// 			} `yaml:"initCopyConsulContainer,omitempty"`
// 			Affinity          string      `yaml:"affinity,omitempty"`
// 			Tolerations       interface{} `yaml:"tolerations,omitempty"`
// 			NodeSelector      interface{} `yaml:"nodeSelector,omitempty"`
// 			PriorityClassName string      `yaml:"priorityClassName,omitempty"`
// 			Annotations       interface{} `yaml:"annotations,omitempty"`
// 			ConsulNamespace   string      `yaml:"consulNamespace,omitempty"`
// 		} `yaml:"defaults,omitempty"`
// 		Gateways []struct {
// 			Name string `yaml:"name,omitempty"`
// 		} `yaml:"gateways,omitempty"`
// 	} `yaml:"ingressGateways,omitempty"`
// 	TerminatingGateways struct {
// 		Enabled  bool `yaml:"enabled,omitempty"`
// 		Defaults struct {
// 			Replicas     int           `yaml:"replicas,omitempty"`
// 			ExtraVolumes []interface{} `yaml:"extraVolumes,omitempty"`
// 			Resources    struct {
// 				Requests struct {
// 					Memory string `yaml:"memory,omitempty"`
// 					CPU    string `yaml:"cpu,omitempty"`
// 				} `yaml:"requests,omitempty"`
// 				Limits struct {
// 					Memory string `yaml:"memory,omitempty"`
// 					CPU    string `yaml:"cpu,omitempty"`
// 				} `yaml:"limits,omitempty"`
// 			} `yaml:"resources,omitempty"`
// 			InitCopyConsulContainer struct {
// 				Resources struct {
// 					Requests struct {
// 						Memory string `yaml:"memory,omitempty"`
// 						CPU    string `yaml:"cpu,omitempty"`
// 					} `yaml:"requests,omitempty"`
// 					Limits struct {
// 						Memory string `yaml:"memory,omitempty"`
// 						CPU    string `yaml:"cpu,omitempty"`
// 					} `yaml:"limits,omitempty"`
// 				} `yaml:"resources,omitempty"`
// 			} `yaml:"initCopyConsulContainer,omitempty"`
// 			Affinity          string      `yaml:"affinity,omitempty"`
// 			Tolerations       interface{} `yaml:"tolerations,omitempty"`
// 			NodeSelector      interface{} `yaml:"nodeSelector,omitempty"`
// 			PriorityClassName string      `yaml:"priorityClassName,omitempty"`
// 			Annotations       interface{} `yaml:"annotations,omitempty"`
// 			ConsulNamespace   string      `yaml:"consulNamespace,omitempty"`
// 		} `yaml:"defaults,omitempty"`
// 		Gateways []struct {
// 			Name string `yaml:"name,omitempty"`
// 		} `yaml:"gateways,omitempty"`
// 	} `yaml:"terminatingGateways,omitempty"`
// 	Tests struct {
// 		Enabled bool `yaml:"enabled,omitempty"`
// 	} `yaml:"tests,omitempty"`
// }
