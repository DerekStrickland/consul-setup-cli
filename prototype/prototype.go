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

// YamlValues is a strongly typed representations of the possible helm chart values
// that a user could specify during, say a setup wizard process. Ideally, the wizard
// would iterate over each top level struct member, and then ask the user to choose
// and options.  Then recursively, it
type ConsulHelmValues struct {
	Global struct {
		Enabled                   bool          `yaml:"enabled,omitempty"`
		Name                      interface{}   `yaml:"name,omitempty"`
		Domain                    string        `yaml:"domain,omitempty"`
		Image                     string        `yaml:"image,omitempty"`
		ImagePullSecrets          []interface{} `yaml:"imagePullSecrets,omitempty"`
		ImageK8S                  string        `yaml:"imageK8S,omitempty"`
		Datacenter                string        `yaml:"datacenter,omitempty"`
		EnablePodSecurityPolicies bool          `yaml:"enablePodSecurityPolicies,omitempty"`
		GossipEncryption          struct {
			SecretName string `yaml:"secretName,omitempty"`
			SecretKey  string `yaml:"secretKey,omitempty"`
		} `yaml:"gossipEncryption,omitempty"`
		TLS struct {
			Enabled                 bool          `yaml:"enabled,omitempty"`
			EnableAutoEncrypt       bool          `yaml:"enableAutoEncrypt,omitempty"`
			ServerAdditionalDNSSANs []interface{} `yaml:"serverAdditionalDNSSANs,omitempty"`
			ServerAdditionalIPSANs  []interface{} `yaml:"serverAdditionalIPSANs,omitempty"`
			Verify                  bool          `yaml:"verify,omitempty"`
			HTTPSOnly               bool          `yaml:"httpsOnly,omitempty"`
			CaCert                  struct {
				SecretName interface{} `yaml:"secretName,omitempty"`
				SecretKey  interface{} `yaml:"secretKey,omitempty"`
			} `yaml:"caCert,omitempty"`
			CaKey struct {
				SecretName interface{} `yaml:"secretName,omitempty"`
				SecretKey  interface{} `yaml:"secretKey,omitempty"`
			} `yaml:"caKey,omitempty"`
		} `yaml:"tls,omitempty"`
		EnableConsulNamespaces bool `yaml:"enableConsulNamespaces,omitempty"`
		Acls                   struct {
			ManageSystemACLs bool `yaml:"manageSystemACLs,omitempty"`
			BootstrapToken   struct {
				SecretName interface{} `yaml:"secretName,omitempty"`
				SecretKey  interface{} `yaml:"secretKey,omitempty"`
			} `yaml:"bootstrapToken,omitempty"`
			CreateReplicationToken bool `yaml:"createReplicationToken,omitempty"`
			ReplicationToken       struct {
				SecretName interface{} `yaml:"secretName,omitempty"`
				SecretKey  interface{} `yaml:"secretKey,omitempty"`
			} `yaml:"replicationToken,omitempty"`
		} `yaml:"acls,omitempty"`
		Federation struct {
			Enabled                bool `yaml:"enabled,omitempty"`
			CreateFederationSecret bool `yaml:"createFederationSecret,omitempty"`
		} `yaml:"federation,omitempty"`
		LifecycleSidecarContainer struct {
			Resources struct {
				Requests struct {
					Memory string `yaml:"memory,omitempty"`
					CPU    string `yaml:"cpu,omitempty"`
				} `yaml:"requests,omitempty"`
				Limits struct {
					Memory string `yaml:"memory,omitempty"`
					CPU    string `yaml:"cpu,omitempty"`
				} `yaml:"limits,omitempty"`
			} `yaml:"resources,omitempty"`
		} `yaml:"lifecycleSidecarContainer,omitempty"`
		ImageEnvoy string `yaml:"imageEnvoy,omitempty"`
		Openshift  struct {
			Enabled bool `yaml:"enabled,omitempty"`
		} `yaml:"openshift,omitempty"`
	} `yaml:"global,omitempty"`
	Server struct {
		Enabled           string      `yaml:"enabled,omitempty"`
		Image             interface{} `yaml:"image,omitempty"`
		Replicas          int         `yaml:"replicas,omitempty"`
		BootstrapExpect   int         `yaml:"bootstrapExpect,omitempty"`
		EnterpriseLicense struct {
			SecretName interface{} `yaml:"secretName,omitempty"`
			SecretKey  interface{} `yaml:"secretKey,omitempty"`
		} `yaml:"enterpriseLicense,omitempty"`
		Storage      string      `yaml:"storage,omitempty"`
		StorageClass interface{} `yaml:"storageClass,omitempty"`
		Connect      bool        `yaml:"connect,omitempty"`
		Resources    struct {
			Requests struct {
				Memory string `yaml:"memory,omitempty"`
				CPU    string `yaml:"cpu,omitempty"`
			} `yaml:"requests,omitempty"`
			Limits struct {
				Memory string `yaml:"memory,omitempty"`
				CPU    string `yaml:"cpu,omitempty"`
			} `yaml:"limits,omitempty"`
		} `yaml:"resources,omitempty"`
		UpdatePartition  int `yaml:"updatePartition,omitempty"`
		DisruptionBudget struct {
			Enabled        bool        `yaml:"enabled,omitempty"`
			MaxUnavailable interface{} `yaml:"maxUnavailable,omitempty"`
		} `yaml:"disruptionBudget,omitempty"`
		ExtraConfig       string        `yaml:"extraConfig,omitempty"`
		ExtraVolumes      []interface{} `yaml:"extraVolumes,omitempty"`
		Affinity          string        `yaml:"affinity,omitempty"`
		Tolerations       string        `yaml:"tolerations,omitempty"`
		NodeSelector      interface{}   `yaml:"nodeSelector,omitempty"`
		PriorityClassName string        `yaml:"priorityClassName,omitempty"`
		ExtraLabels       interface{}   `yaml:"extraLabels,omitempty"`
		Annotations       interface{}   `yaml:"annotations,omitempty"`
		Service           struct {
			Annotations interface{} `yaml:"annotations,omitempty"`
		} `yaml:"service,omitempty"`
		ExtraEnvironmentVars struct {
		} `yaml:"extraEnvironmentVars,omitempty"`
	} `yaml:"server,omitempty"`
	ExternalServers struct {
		Enabled           bool          `yaml:"enabled,omitempty"`
		Hosts             []interface{} `yaml:"hosts,omitempty"`
		HTTPSPort         int           `yaml:"httpsPort,omitempty"`
		TLSServerName     interface{}   `yaml:"tlsServerName,omitempty"`
		UseSystemRoots    bool          `yaml:"useSystemRoots,omitempty"`
		K8SAuthMethodHost interface{}   `yaml:"k8sAuthMethodHost,omitempty"`
	} `yaml:"externalServers,omitempty"`
	Client struct {
		Enabled               string      `yaml:"enabled,omitempty"`
		Image                 interface{} `yaml:"image,omitempty"`
		Join                  interface{} `yaml:"join,omitempty"`
		DataDirectoryHostPath interface{} `yaml:"dataDirectoryHostPath,omitempty"`
		Grpc                  bool        `yaml:"grpc,omitempty"`
		ExposeGossipPorts     bool        `yaml:"exposeGossipPorts,omitempty"`
		Resources             struct {
			Requests struct {
				Memory string `yaml:"memory,omitempty"`
				CPU    string `yaml:"cpu,omitempty"`
			} `yaml:"requests,omitempty"`
			Limits struct {
				Memory string `yaml:"memory,omitempty"`
				CPU    string `yaml:"cpu,omitempty"`
			} `yaml:"limits,omitempty"`
		} `yaml:"resources,omitempty"`
		ExtraConfig  string        `yaml:"extraConfig,omitempty"`
		ExtraVolumes []interface{} `yaml:"extraVolumes,omitempty"`
		Tolerations  string        `yaml:"tolerations,omitempty"`
		NodeSelector interface{}   `yaml:"nodeSelector,omitempty"`
		Affinity     struct {
		} `yaml:"affinity,omitempty"`
		PriorityClassName    string      `yaml:"priorityClassName,omitempty"`
		Annotations          interface{} `yaml:"annotations,omitempty"`
		ExtraEnvironmentVars struct {
		} `yaml:"extraEnvironmentVars,omitempty"`
		DNSPolicy      interface{} `yaml:"dnsPolicy,omitempty"`
		HostNetwork    bool        `yaml:"hostNetwork,omitempty"`
		UpdateStrategy interface{} `yaml:"updateStrategy,omitempty"`
		SnapshotAgent  struct {
			Enabled      bool `yaml:"enabled,omitempty"`
			Replicas     int  `yaml:"replicas,omitempty"`
			ConfigSecret struct {
				SecretName interface{} `yaml:"secretName,omitempty"`
				SecretKey  interface{} `yaml:"secretKey,omitempty"`
			} `yaml:"configSecret,omitempty"`
			Resources struct {
				Requests struct {
					Memory string `yaml:"memory,omitempty"`
					CPU    string `yaml:"cpu,omitempty"`
				} `yaml:"requests,omitempty"`
				Limits struct {
					Memory string `yaml:"memory,omitempty"`
					CPU    string `yaml:"cpu,omitempty"`
				} `yaml:"limits,omitempty"`
			} `yaml:"resources,omitempty"`
			CaCert interface{} `yaml:"caCert,omitempty"`
		} `yaml:"snapshotAgent,omitempty"`
	} `yaml:"client,omitempty"`
	DNS struct {
		Enabled        string      `yaml:"enabled,omitempty"`
		Type           string      `yaml:"type,omitempty"`
		ClusterIP      interface{} `yaml:"clusterIP,omitempty"`
		Annotations    interface{} `yaml:"annotations,omitempty"`
		AdditionalSpec interface{} `yaml:"additionalSpec,omitempty"`
	} `yaml:"dns,omitempty"`
	UI struct {
		Enabled string `yaml:"enabled,omitempty"`
		Service struct {
			Enabled        bool        `yaml:"enabled,omitempty"`
			Type           interface{} `yaml:"type,omitempty"`
			Annotations    interface{} `yaml:"annotations,omitempty"`
			AdditionalSpec interface{} `yaml:"additionalSpec,omitempty"`
		} `yaml:"service,omitempty"`
	} `yaml:"ui,omitempty"`
	SyncCatalog struct {
		Enabled            bool        `yaml:"enabled,omitempty"`
		Image              interface{} `yaml:"image,omitempty"`
		Default            bool        `yaml:"default,omitempty"`
		PriorityClassName  string      `yaml:"priorityClassName,omitempty"`
		ToConsul           bool        `yaml:"toConsul,omitempty"`
		ToK8S              bool        `yaml:"toK8S,omitempty"`
		K8SPrefix          interface{} `yaml:"k8sPrefix,omitempty"`
		K8SAllowNamespaces []string    `yaml:"k8sAllowNamespaces,omitempty"`
		K8SDenyNamespaces  []string    `yaml:"k8sDenyNamespaces,omitempty"`
		K8SSourceNamespace interface{} `yaml:"k8sSourceNamespace,omitempty"`
		ConsulNamespaces   struct {
			ConsulDestinationNamespace string `yaml:"consulDestinationNamespace,omitempty"`
			MirroringK8S               bool   `yaml:"mirroringK8S,omitempty"`
			MirroringK8SPrefix         string `yaml:"mirroringK8SPrefix,omitempty"`
		} `yaml:"consulNamespaces,omitempty"`
		AddK8SNamespaceSuffix bool        `yaml:"addK8SNamespaceSuffix,omitempty"`
		ConsulPrefix          interface{} `yaml:"consulPrefix,omitempty"`
		K8STag                interface{} `yaml:"k8sTag,omitempty"`
		ConsulNodeName        string      `yaml:"consulNodeName,omitempty"`
		SyncClusterIPServices bool        `yaml:"syncClusterIPServices,omitempty"`
		NodePortSyncType      string      `yaml:"nodePortSyncType,omitempty"`
		ACLSyncToken          struct {
			SecretName interface{} `yaml:"secretName,omitempty"`
			SecretKey  interface{} `yaml:"secretKey,omitempty"`
		} `yaml:"aclSyncToken,omitempty"`
		NodeSelector interface{} `yaml:"nodeSelector,omitempty"`
		Affinity     interface{} `yaml:"affinity,omitempty"`
		Tolerations  interface{} `yaml:"tolerations,omitempty"`
		Resources    struct {
			Requests struct {
				Memory string `yaml:"memory,omitempty"`
				CPU    string `yaml:"cpu,omitempty"`
			} `yaml:"requests,omitempty"`
			Limits struct {
				Memory string `yaml:"memory,omitempty"`
				CPU    string `yaml:"cpu,omitempty"`
			} `yaml:"limits,omitempty"`
		} `yaml:"resources,omitempty"`
		LogLevel            string      `yaml:"logLevel,omitempty"`
		ConsulWriteInterval interface{} `yaml:"consulWriteInterval,omitempty"`
	} `yaml:"syncCatalog,omitempty"`
	ConnectInject struct {
		Enabled      bool        `yaml:"enabled,omitempty"`
		Image        interface{} `yaml:"image,omitempty"`
		Default      bool        `yaml:"default,omitempty"`
		HealthChecks struct {
			Enabled         bool   `yaml:"enabled,omitempty"`
			ReconcilePeriod string `yaml:"reconcilePeriod,omitempty"`
		} `yaml:"healthChecks,omitempty"`
		EnvoyExtraArgs    interface{} `yaml:"envoyExtraArgs,omitempty"`
		PriorityClassName string      `yaml:"priorityClassName,omitempty"`
		ImageConsul       interface{} `yaml:"imageConsul,omitempty"`
		LogLevel          string      `yaml:"logLevel,omitempty"`
		Resources         struct {
			Requests struct {
				Memory string `yaml:"memory,omitempty"`
				CPU    string `yaml:"cpu,omitempty"`
			} `yaml:"requests,omitempty"`
			Limits struct {
				Memory string `yaml:"memory,omitempty"`
				CPU    string `yaml:"cpu,omitempty"`
			} `yaml:"limits,omitempty"`
		} `yaml:"resources,omitempty"`
		NamespaceSelector  interface{}   `yaml:"namespaceSelector,omitempty"`
		K8SAllowNamespaces []string      `yaml:"k8sAllowNamespaces,omitempty"`
		K8SDenyNamespaces  []interface{} `yaml:"k8sDenyNamespaces,omitempty"`
		ConsulNamespaces   struct {
			ConsulDestinationNamespace string `yaml:"consulDestinationNamespace,omitempty"`
			MirroringK8S               bool   `yaml:"mirroringK8S,omitempty"`
			MirroringK8SPrefix         string `yaml:"mirroringK8SPrefix,omitempty"`
		} `yaml:"consulNamespaces,omitempty"`
		Certs struct {
			SecretName interface{} `yaml:"secretName,omitempty"`
			CaBundle   string      `yaml:"caBundle,omitempty"`
			CertName   string      `yaml:"certName,omitempty"`
			KeyName    string      `yaml:"keyName,omitempty"`
		} `yaml:"certs,omitempty"`
		NodeSelector           interface{} `yaml:"nodeSelector,omitempty"`
		Affinity               interface{} `yaml:"affinity,omitempty"`
		Tolerations            interface{} `yaml:"tolerations,omitempty"`
		ACLBindingRuleSelector string      `yaml:"aclBindingRuleSelector,omitempty"`
		OverrideAuthMethodName string      `yaml:"overrideAuthMethodName,omitempty"`
		ACLInjectToken         struct {
			SecretName interface{} `yaml:"secretName,omitempty"`
			SecretKey  interface{} `yaml:"secretKey,omitempty"`
		} `yaml:"aclInjectToken,omitempty"`
		CentralConfig struct {
			Enabled         bool        `yaml:"enabled,omitempty"`
			DefaultProtocol interface{} `yaml:"defaultProtocol,omitempty"`
			ProxyDefaults   string      `yaml:"proxyDefaults,omitempty"`
		} `yaml:"centralConfig,omitempty"`
		SidecarProxy struct {
			Resources struct {
				Requests struct {
					Memory interface{} `yaml:"memory,omitempty"`
					CPU    interface{} `yaml:"cpu,omitempty"`
				} `yaml:"requests,omitempty"`
				Limits struct {
					Memory interface{} `yaml:"memory,omitempty"`
					CPU    interface{} `yaml:"cpu,omitempty"`
				} `yaml:"limits,omitempty"`
			} `yaml:"resources,omitempty"`
		} `yaml:"sidecarProxy,omitempty"`
		InitContainer struct {
			Resources struct {
				Requests struct {
					Memory string `yaml:"memory,omitempty"`
					CPU    string `yaml:"cpu,omitempty"`
				} `yaml:"requests,omitempty"`
				Limits struct {
					Memory string `yaml:"memory,omitempty"`
					CPU    string `yaml:"cpu,omitempty"`
				} `yaml:"limits,omitempty"`
			} `yaml:"resources,omitempty"`
		} `yaml:"initContainer,omitempty"`
	} `yaml:"connectInject,omitempty"`
	Controller struct {
		Enabled   bool   `yaml:"enabled,omitempty"`
		Replicas  int    `yaml:"replicas,omitempty"`
		LogLevel  string `yaml:"logLevel,omitempty"`
		Resources struct {
			Limits struct {
				CPU    string `yaml:"cpu,omitempty"`
				Memory string `yaml:"memory,omitempty"`
			} `yaml:"limits,omitempty"`
			Requests struct {
				CPU    string `yaml:"cpu,omitempty"`
				Memory string `yaml:"memory,omitempty"`
			} `yaml:"requests,omitempty"`
		} `yaml:"resources,omitempty"`
		NodeSelector      interface{} `yaml:"nodeSelector,omitempty"`
		Tolerations       interface{} `yaml:"tolerations,omitempty"`
		Affinity          interface{} `yaml:"affinity,omitempty"`
		PriorityClassName string      `yaml:"priorityClassName,omitempty"`
	} `yaml:"controller,omitempty"`
	MeshGateway struct {
		Enabled    bool   `yaml:"enabled,omitempty"`
		GlobalMode string `yaml:"globalMode,omitempty"`
		Replicas   int    `yaml:"replicas,omitempty"`
		WanAddress struct {
			Source string `yaml:"source,omitempty"`
			Port   int    `yaml:"port,omitempty"`
			Static string `yaml:"static,omitempty"`
		} `yaml:"wanAddress,omitempty"`
		Service struct {
			Enabled        bool        `yaml:"enabled,omitempty"`
			Type           string      `yaml:"type,omitempty"`
			Port           int         `yaml:"port,omitempty"`
			NodePort       interface{} `yaml:"nodePort,omitempty"`
			Annotations    interface{} `yaml:"annotations,omitempty"`
			AdditionalSpec interface{} `yaml:"additionalSpec,omitempty"`
		} `yaml:"service,omitempty"`
		HostNetwork       bool        `yaml:"hostNetwork,omitempty"`
		DNSPolicy         interface{} `yaml:"dnsPolicy,omitempty"`
		ConsulServiceName string      `yaml:"consulServiceName,omitempty"`
		ContainerPort     int         `yaml:"containerPort,omitempty"`
		HostPort          interface{} `yaml:"hostPort,omitempty"`
		Resources         struct {
			Requests struct {
				Memory string `yaml:"memory,omitempty"`
				CPU    string `yaml:"cpu,omitempty"`
			} `yaml:"requests,omitempty"`
			Limits struct {
				Memory string `yaml:"memory,omitempty"`
				CPU    string `yaml:"cpu,omitempty"`
			} `yaml:"limits,omitempty"`
		} `yaml:"resources,omitempty"`
		InitCopyConsulContainer struct {
			Resources struct {
				Requests struct {
					Memory string `yaml:"memory,omitempty"`
					CPU    string `yaml:"cpu,omitempty"`
				} `yaml:"requests,omitempty"`
				Limits struct {
					Memory string `yaml:"memory,omitempty"`
					CPU    string `yaml:"cpu,omitempty"`
				} `yaml:"limits,omitempty"`
			} `yaml:"resources,omitempty"`
		} `yaml:"initCopyConsulContainer,omitempty"`
		Affinity          string      `yaml:"affinity,omitempty"`
		Tolerations       interface{} `yaml:"tolerations,omitempty"`
		NodeSelector      interface{} `yaml:"nodeSelector,omitempty"`
		PriorityClassName string      `yaml:"priorityClassName,omitempty"`
		Annotations       interface{} `yaml:"annotations,omitempty"`
	} `yaml:"meshGateway,omitempty"`
	IngressGateways struct {
		Enabled  bool `yaml:"enabled,omitempty"`
		Defaults struct {
			Replicas int `yaml:"replicas,omitempty"`
			Service  struct {
				Type  string `yaml:"type,omitempty"`
				Ports []struct {
					Port     int         `yaml:"port,omitempty"`
					NodePort interface{} `yaml:"nodePort,omitempty"`
				} `yaml:"ports,omitempty"`
				Annotations    interface{} `yaml:"annotations,omitempty"`
				AdditionalSpec interface{} `yaml:"additionalSpec,omitempty"`
			} `yaml:"service,omitempty"`
			Resources struct {
				Requests struct {
					Memory string `yaml:"memory,omitempty"`
					CPU    string `yaml:"cpu,omitempty"`
				} `yaml:"requests,omitempty"`
				Limits struct {
					Memory string `yaml:"memory,omitempty"`
					CPU    string `yaml:"cpu,omitempty"`
				} `yaml:"limits,omitempty"`
			} `yaml:"resources,omitempty"`
			InitCopyConsulContainer struct {
				Resources struct {
					Requests struct {
						Memory string `yaml:"memory,omitempty"`
						CPU    string `yaml:"cpu,omitempty"`
					} `yaml:"requests,omitempty"`
					Limits struct {
						Memory string `yaml:"memory,omitempty"`
						CPU    string `yaml:"cpu,omitempty"`
					} `yaml:"limits,omitempty"`
				} `yaml:"resources,omitempty"`
			} `yaml:"initCopyConsulContainer,omitempty"`
			Affinity          string      `yaml:"affinity,omitempty"`
			Tolerations       interface{} `yaml:"tolerations,omitempty"`
			NodeSelector      interface{} `yaml:"nodeSelector,omitempty"`
			PriorityClassName string      `yaml:"priorityClassName,omitempty"`
			Annotations       interface{} `yaml:"annotations,omitempty"`
			ConsulNamespace   string      `yaml:"consulNamespace,omitempty"`
		} `yaml:"defaults,omitempty"`
		Gateways []struct {
			Name string `yaml:"name,omitempty"`
		} `yaml:"gateways,omitempty"`
	} `yaml:"ingressGateways,omitempty"`
	TerminatingGateways struct {
		Enabled  bool `yaml:"enabled,omitempty"`
		Defaults struct {
			Replicas     int           `yaml:"replicas,omitempty"`
			ExtraVolumes []interface{} `yaml:"extraVolumes,omitempty"`
			Resources    struct {
				Requests struct {
					Memory string `yaml:"memory,omitempty"`
					CPU    string `yaml:"cpu,omitempty"`
				} `yaml:"requests,omitempty"`
				Limits struct {
					Memory string `yaml:"memory,omitempty"`
					CPU    string `yaml:"cpu,omitempty"`
				} `yaml:"limits,omitempty"`
			} `yaml:"resources,omitempty"`
			InitCopyConsulContainer struct {
				Resources struct {
					Requests struct {
						Memory string `yaml:"memory,omitempty"`
						CPU    string `yaml:"cpu,omitempty"`
					} `yaml:"requests,omitempty"`
					Limits struct {
						Memory string `yaml:"memory,omitempty"`
						CPU    string `yaml:"cpu,omitempty"`
					} `yaml:"limits,omitempty"`
				} `yaml:"resources,omitempty"`
			} `yaml:"initCopyConsulContainer,omitempty"`
			Affinity          string      `yaml:"affinity,omitempty"`
			Tolerations       interface{} `yaml:"tolerations,omitempty"`
			NodeSelector      interface{} `yaml:"nodeSelector,omitempty"`
			PriorityClassName string      `yaml:"priorityClassName,omitempty"`
			Annotations       interface{} `yaml:"annotations,omitempty"`
			ConsulNamespace   string      `yaml:"consulNamespace,omitempty"`
		} `yaml:"defaults,omitempty"`
		Gateways []struct {
			Name string `yaml:"name,omitempty"`
		} `yaml:"gateways,omitempty"`
	} `yaml:"terminatingGateways,omitempty"`
	Tests struct {
		Enabled bool `yaml:"enabled,omitempty"`
	} `yaml:"tests,omitempty"`
}
