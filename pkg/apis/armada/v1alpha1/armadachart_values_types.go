// Copyright 2019 The Armada Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

type AVPod struct {
	// affinity contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Affinity *runtime.RawExtension `json:"affinity,omitempty"`
	// env contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Env *runtime.RawExtension `json:"env,omitempty"`
	// mount_path contains tbd
	MountPath string `json:"mount_path,omitempty"`
	// lifecycle contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Lifecycle *runtime.RawExtension `json:"lifecycle,omitempty"`
	// replicas contains tbd
	Replicas map[string]int `json:"replicas,omitempty"`
	// resources contains tbd
	Resources map[string]AVPodResources `json:"resources,omitempty"`
	// security_context contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	SecurityContext *runtime.RawExtension `json:"security_context,omitempty"`
}

type AVPodResources struct {
	// curator contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Curator *runtime.RawExtension `json:"curator,omitempty"`
	// fluentbit contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Fluentbit *runtime.RawExtension `json:"fluentbit,omitempty"`
	// limits contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Limits *runtime.RawExtension `json:"limits,omitempty"`
	// requests contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Requests *runtime.RawExtension `json:"requests,omitempty"`
	// image_repo_sync contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	ImageRepoSync *runtime.RawExtension `json:"image_repo_sync,omitempty"`
	// snapshot_repository contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	SnapshortRepository *runtime.RawExtension `json:"snapshot_repository,omitempty"`
	// tests contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Tests *runtime.RawExtension `json:"tests,omitempty"`
}

type AVPodResourceSettings struct {
	// cpu contains tbd
	Cpu string `json:"cpu,omitempty"`
	// memory contains tbd
	Memory string `json:"memory,omitempty"`
}

type AVImages struct {
	// tags contains tbd
	Tags map[string]string `json:"tags,omitempty"`
	// pull_policy contains tbd
	PullPolicy string `json:"pull_policy,omitempty"`
}

type AVVolume struct {
	// chown_on_start contains tbd
	ChownOnStart *bool `json:"chown_on_start,omitempty"`
	// backup contains tbd
	Backup *AVVolumeBackup `json:"backup,omitempty"`
	// class_name contains tbd
	ClassName *string `json:"class_name,omitempty"`
	// size contains tbd
	Size *string `json:"size,omitempty"`
	// enabled contains tbd
	Enabled *bool `json:"enabled,omitempty"`
}

type AVVolumeBackup struct {
	// class_name contains tbd
	ClassName *string `json:"class_name,omitempty"`
	// size contains tbd
	Size *string `json:"size,omitempty"`
	// enabled contains tbd
	Enabled *bool `json:"enabled,omitempty"`
}

type AVSecrets struct {
	// anchor contains tbd
	Anchor *AVSecretAnchor `json:"anchor,omitempty"`
	// etcd contains tbd
	Etcd *AVSecretEtcd `json:"etcd,omitempty"`
	// keyrings contains tbd
	Keyrings *AVSecretKeyrings `json:"keyrings,omitempty"`
	// maas_region contains tbd
	MaasRegion *AVSecretMaasRegion `json:"maas_region,omitempty"`
	// service_account contains tbd
	ServiceAccount *AVSecretServiceAccount `json:"service_account,omitempty"`
	// tls contains tbd
	Tls *AVTls `json:"tls,omitempty"`
}

type AVBootstrap struct {
	// enabled contains tbd
	Enabled bool `json:"enabled,omitempty"`
	// script contains tbd
	Script string `json:"script,omitempty"`
	// ip contains tbd
	Ip string `json:"ip,omitempty"`
}

type AVBootstrapping struct {
	// enabled contains tbd
	Enabled bool `json:"enabled,omitempty"`
	// filename contains tbd
	Filename string `json:"filename,omitempty"`
	// host_directory contains tbd
	HostDirectory string `json:"host_directory,omitempty"`
}

type AVEndpoints struct {
	// alerts contains tbd
	Alerts *AVEndpointType1 `json:"alerts,omitempty"`
	// armada contains tbd
	Armada *AVEndpointType1 `json:"armada,omitempty"`
	// ceph_object_store contains tbd
	CephObjectStore *AVEndpointType1 `json:"ceph_object_store,omitempty"`
	// cloudformation contains tbd
	Cloudformation *AVEndpointType1 `json:"cloudformation,omitempty"`
	// cloudwatch contains tbd
	Cloudwatch *AVEndpointType1 `json:"cloudwatch,omitempty"`
	// compute_metadata contains tbd
	ComputeMetadata *AVEndpointType1 `json:"compute_metadata,omitempty"`
	// compute_novnc_proxy contains tbd
	ComputeNovncProxy *AVEndpointType1 `json:"compute_novnc_proxy,omitempty"`
	// compute contains tbd
	Compute *AVEndpointType1 `json:"compute,omitempty"`
	// compute_spice_proxy contains tbd
	ComputeSpiceProxy *AVEndpointType1 `json:"compute_spice_proxy,omitempty"`
	// dashboard contains tbd
	Dashboard *AVEndpointType1 `json:"dashboard,omitempty"`
	// deckhand contains tbd
	Deckhand *AVEndpointType1 `json:"deckhand,omitempty"`
	// elasticsearch contains tbd
	Elasticsearch *AVEndpointType1 `json:"elasticsearch,omitempty"`
	// fluentd contains tbd
	Fluentd *AVEndpointType1 `json:"fluentd,omitempty"`
	// grafana contains tbd
	Grafana *AVEndpointType1 `json:"grafana,omitempty"`
	// identity contains tbd
	Identity *AVEndpointType1 `json:"identity,omitempty"`
	// image contains tbd
	Image *AVEndpointType1 `json:"image,omitempty"`
	// image_registry contains tbd
	ImageRegistry *AVEndpointType1 `json:"image_registry,omitempty"`
	// key_manager contains tbd
	KeyManager *AVEndpointType1 `json:"key_manager,omitempty"`
	// kibana contains tbd
	Kibana *AVEndpointType1 `json:"kibana,omitempty"`
	// kube_controller_manager contains tbd
	KubeControllerManager *AVEndpointType1 `json:"kube_controller_manager,omitempty"`
	// kubernetesprovisioner contains tbd
	Kubernetesprovisioner *AVEndpointType1 `json:"kubernetesprovisioner,omitempty"`
	// kube_scheduler contains tbd
	KubeScheduler *AVEndpointType1 `json:"kube_scheduler,omitempty"`
	// kube_state_metrics contains tbd
	KubeStateMetrics *AVEndpointType1 `json:"kube_state_metrics,omitempty"`
	// ldap contains tbd
	Ldap *AVEndpointType1 `json:"ldap,omitempty"`
	// maas_region_ui contains tbd
	MaasRegionUi *AVEndpointType1 `json:"maas_region_ui,omitempty"`
	// monitoring contains tbd
	Monitoring *AVEndpointType1 `json:"monitoring,omitempty"`
	// nagios contains tbd
	Nagios *AVEndpointType1 `json:"nagios,omitempty"`
	// network contains tbd
	Network *AVEndpointType1 `json:"network,omitempty"`
	// node_metrics contains tbd
	NodeMetrics *AVEndpointType1 `json:"node_metrics,omitempty"`
	// object_store contains tbd
	ObjectStore *AVEndpointType1 `json:"object_store,omitempty"`
	// orchestration contains tbd
	Orchestration *AVEndpointType1 `json:"orchestration,omitempty"`
	// physicalprovisioner contains tbd
	Physicalprovisioner *AVEndpointType1 `json:"physicalprovisioner,omitempty"`
	// placement contains tbd
	Placement *AVEndpointType1 `json:"placement,omitempty"`
	// process_exporter_metrics contains tbd
	ProcessExporterMetrics *AVEndpointType1 `json:"process_exporter_metrics,omitempty"`
	// prometheus_elasticsearch_exporter contains tbd
	PrometheusElasticsearchExporter *AVEndpointType1 `json:"prometheus_elasticsearch_exporter,omitempty"`
	// prometheus_fluentd_exporter contains tbd
	PrometheusFluentdExporter *AVEndpointType1 `json:"prometheus_fluentd_exporter,omitempty"`
	// prometheus_mysql_exporter contains tbd
	PrometheusMysqlExporter *AVEndpointType1 `json:"prometheus_mysql_exporter,omitempty"`
	// prometheus_openstack_exporter contains tbd
	PrometheusOpenstackExporter *AVEndpointType1 `json:"prometheus_openstack_exporter,omitempty"`
	// prometheus_rabbitmq_exporter contains tbd
	PrometheusRabbitmqExporter *AVEndpointType1 `json:"prometheus_rabbitmq_exporter,omitempty"`
	// shipyard contains tbd
	Shipyard *AVEndpointType1 `json:"shipyard,omitempty"`
	// volume contains tbd
	Volume *AVEndpointType1 `json:"volume,omitempty"`
	// volumev2 contains tbd
	Volumev2 *AVEndpointType1 `json:"volumev2,omitempty"`
	// volumev3 contains tbd
	Volumev3 *AVEndpointType1 `json:"volumev3,omitempty"`

	// maas_db contains tbd
	MaasDb *AVEndpointType2 `json:"maas_db,omitempty"`
	// oslo_db_api contains tbd
	OsloDbApi *AVEndpointType2 `json:"oslo_db_api,omitempty"`
	// oslo_db_cell0 contains tbd
	OsloDbCell0 *AVEndpointType2 `json:"oslo_db_cell0,omitempty"`
	// oslo_db contains tbd
	OsloDb *AVEndpointType2 `json:"oslo_db,omitempty"`
	// oslo_db_session contains tbd
	OsloDbSession *AVEndpointType2 `json:"oslo_db_session,omitempty"`
	// oslo_messaging contains tbd
	OsloMessaging *AVEndpointType2 `json:"oslo_messaging,omitempty"`
	// postgresql_airflow_celery_db contains tbd
	PostgresqlAirflowCeleryDb *AVEndpointType2 `json:"postgresql_airflow_celery_db,omitempty"`
	// postgresql_airflow_db contains tbd
	PostgresqlAirflowDb *AVEndpointType2 `json:"postgresql_airflow_db,omitempty"`
	// postgresql contains tbd
	Postgresql *AVEndpointType2 `json:"postgresql,omitempty"`
	// postgresql_shipyard_db contains tbd
	PostgresqlShipyardDb *AVEndpointType2 `json:"postgresql_shipyard_db,omitempty"`

	// Added recently
	CephMon     *AVEndpointType1 `json:"ceph_mon,omitempty"`
	CephMgr     *AVEndpointType1 `json:"ceph_mgr,omitempty"`
	Etcd        *AVEndpointType1 `json:"etcd,omitempty"`
	Fluentbit   *AVEndpointType1 `json:"fluentbit,omitempty"`
	Ingress     *AVEndpointType1 `json:"ingress,omitempty"`
	MaasIngress *AVEndpointType1 `json:"maas_ingress,omitempty"`
	MaasRegion  *AVEndpointType1 `json:"maas_region,omitempty"`
	OsloCache   *AVEndpointType2 `json:"oslo_cache,omitempty"`
}

type AVEndpointType1 struct {
	// auth contains tbd
	Auth *map[string]AVEndpointAuth `json:"auth,omitempty"`
	// host_fqdn_override contains tbd
	HostFqdnOverride *map[string]string `json:"host_fqdn_override,omitempty"`
	// hosts contains tbd
	Hosts *map[string]string `json:"hosts,omitempty"`
	// name contains tbd
	Name string `json:"name,omitempty"`
	// namespace contains tbd
	Namespace string `json:"namespace,omitempty"`
	// path contains tbd
	Path *map[string]string `json:"path,omitempty"`
	// port contains tbd
	Port *map[string]ArmadaMapInt `json:"port,omitempty"`
	// scheme contains tbd
	Scheme *map[string]string `json:"scheme,omitempty"`
	// type contains tbd
	Type string `json:"type,omitempty"`
}

type AVEndpointType2 struct {
	// auth contains tbd
	Auth *map[string]AVEndpointAuth `json:"auth,omitempty"`
	// host_fqdn_override contains tbd
	HostFqdnOverride *map[string]string `json:"host_fqdn_override,omitempty"`
	// hosts contains tbd
	Hosts *map[string]string `json:"hosts,omitempty"`
	// name contains tbd
	Name string `json:"name,omitempty"`
	// namespace contains tbd
	Namespace string `json:"namespace,omitempty"`
	// path contains tbd
	Path string `json:"path,omitempty"`
	// port contains tbd
	Port *map[string]ArmadaMapInt `json:"port,omitempty"`
	// scheme contains tbd
	Scheme string `json:"scheme,omitempty"`
	// statefulset contains tbd
	StatefuleSet *map[string]string `json:"statefulset,omitempty"`
	// type contains tbd
	Type string `json:"type,omitempty"`
}

type AVEndpointAuth struct {
	// access_key contains tbd
	AccessKey string `json:"access_key,omitempty"`
	// bind contains tbd
	Bind string `json:"bind,omitempty"`
	// bind_dn contains tbd
	BindDn string `json:"bind_dn,omitempty"`
	// database contains tbd
	Database string `json:"database,omitempty"`
	// domain_name contains tbd
	DomainName string `json:"domain_name,omitempty"`
	// email contains tbd
	Email string `json:"email,omitempty"`
	// memcache_secret_key contains tbd
	MemcacheSecretKey string `json:"memcache_secret_key,omitempty"`
	// password contains tbd
	Password string `json:"password,omitempty"`
	// project_domain_name contains tbd
	ProjectDomainName string `json:"project_domain_name,omitempty"`
	// project_name contains tbd
	ProjectName string `json:"project_name,omitempty"`
	// region_name contains tbd
	RegionName string `json:"region_name,omitempty"`
	// role contains tbd
	Role *string `json:"role,omitempty"`
	// secret_key contains tbd
	SecretKey string `json:"secret_key,omitempty"`
	// tls contains tbd
	Tls *AVTls `json:"tls,omitempty"`
	// tmpurlkey contains tbd
	Tmpurlkey string `json:"tmpurlkey,omitempty"`
	// user_domain_name contains tbd
	UserDomainName string `json:"user_domain_name,omitempty"`
	// username contains tbd
	Username string `json:"username,omitempty"`
}

type AVEndpointPort struct {
	// default contains tbd
	Default int `json:"default,omitempty"`
	// internal contains tbd
	Internal int `json:"internal,omitempty"`
	// nodeport contains tbd
	Nodeport int `json:"nodeport,omitempty"`
	// podport contains tbd
	Podport int `json:"podport,omitempty"`
	// public contains tbd
	Public int `json:"public,omitempty"`
}

type AVTls struct {
	// ca contains tbd
	Ca string `json:"ca,omitempty"`
	// crt contains tbd
	Crt string `json:"crt,omitempty"`
	// cert contains tbd
	Cert string `json:"cert,omitempty"`
	// key contains tbd
	Key string `json:"key,omitempty"`
	// client contains tbd
	Client *AVTlsCa `json:"client,omitempty"`
	// peer contains tbd
	Peer *AVTlsCa `json:"peer,omitempty"`
}

type AVTlsCa struct {
	// ca contains tbd
	Ca string `json:"ca,omitempty"`
}

type AVSecretAnchor struct {
	// tls contains tbd
	Tls *AVTls `json:"tls,omitempty"`
}

type AVSecretMaasRegion struct {
	// value contains tbd
	Value string `json:"value,omitempty"`
}

type AVSecretKeyrings struct {
	// admin contains tbd
	Admin string `json:"admin,omitempty"`
}

type AVSecretServiceAccount struct {
	// private_key contains tbd
	PrivateKey string `json:"private_key,omitempty"`
	// public contains tbd
	PublicKey string `json:"public_key,omitempty"`
}

type AVSecretEtcd struct {
	// tls contains tbd
	Tls *AVTls `json:"tls,omitempty"`
}

type AVConf struct {
	// acconfig contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Acconfig *runtime.RawExtension `json:"acconfig,omitempty"`
	// agent contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Agent *runtime.RawExtension `json:"agent,omitempty"`
	// anchor contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Anchor *runtime.RawExtension `json:"anchor,omitempty"`
	// apache contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Apache *runtime.RawExtension `json:"apache,omitempty"`
	// api_metadata contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	ApiMetadata *runtime.RawExtension `json:"api_metadata,omitempty"`
	// armada contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Armada *runtime.RawExtension `json:"armada,omitempty"`
	// auto_bridge_add contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	AutoBridgeAdd *runtime.RawExtension `json:"auto_bridge_add,omitempty"`
	// cache contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Cache *runtime.RawExtension `json:"cache,omitempty"`
	// ceph contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Ceph *runtime.RawExtension `json:"ceph,omitempty"`
	// cni_network_config contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	CniNetworkConfig *runtime.RawExtension `json:"cni_network_config,omitempty"`
	// conductor contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Conductor *runtime.RawExtension `json:"conductor,omitempty"`
	// config contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Config *runtime.RawExtension `json:"config,omitempty"`
	// consoleauth contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Consoleauth *runtime.RawExtension `json:"consoleauth,omitempty"`
	// controllers contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Controllers *runtime.RawExtension `json:"controllers,omitempty"`
	// coredns contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Coredns *runtime.RawExtension `json:"coredns,omitempty"`
	// curator contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Curator *runtime.RawExtension `json:"curator,omitempty"`
	// deckhand contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Deckhand *runtime.RawExtension `json:"deckhand,omitempty"`
	// defaults contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Defaults *runtime.RawExtension `json:"defaults,omitempty"`
	// drydock contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Drydock *runtime.RawExtension `json:"drydock,omitempty"`
	// elasticsearch contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Elasticsearch *runtime.RawExtension `json:"elasticsearch,omitempty"`
	// encryption_provider contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	EncryptionProvider *runtime.RawExtension `json:"encryption_provider,omitempty"`
	// eventconfig contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Eventconfig *runtime.RawExtension `json:"eventconfig,omitempty"`
	// exec contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Exec *runtime.RawExtension `json:"exec,omitempty"`
	// features contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Features *runtime.RawExtension `json:"features,omitempty"`
	// fluentbit contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Fluentbit *runtime.RawExtension `json:"fluentbit,omitempty"`
	// fluentd contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Fluentd *runtime.RawExtension `json:"fluentd,omitempty"`
	// httpd contains tbd
	Httpd string `json:"httpd,omitempty"`
	// ingress contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Ingress *runtime.RawExtension `json:"ingress,omitempty"`
	// job contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Job *runtime.RawExtension `json:"job,omitempty"`
	// keystone contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Keystone *runtime.RawExtension `json:"keystone,omitempty"`
	// ldap contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Ldap *runtime.RawExtension `json:"ldap,omitempty"`
	// logging contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Logging *runtime.RawExtension `json:"logging,omitempty"`
	// maas contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Maas *runtime.RawExtension `json:"maas,omitempty"`
	// metadata_agent contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	MetadataAgent *runtime.RawExtension `json:"metadata_agent,omitempty"`
	// neutron contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Neutron *runtime.RawExtension `json:"neutron,omitempty"`
	// node contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Node *runtime.RawExtension `json:"node,omitempty"`
	// nova contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Nova *runtime.RawExtension `json:"nova,omitempty"`
	// novncproxy contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Novncproxy *runtime.RawExtension `json:"novncproxy,omitempty"`
	// openstack_version contains tbd
	OpenstackVersion string `json:"openstack_version,omitempty"`
	// osapi contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Osapi *runtime.RawExtension `json:"osapi,omitempty"`
	// osd contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Osd *runtime.RawExtension `json:"osd,omitempty"`
	// overrides contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Overrides *runtime.RawExtension `json:"overrides,omitempty"`
	// parsers contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Parsers *runtime.RawExtension `json:"parsers,omitempty"`
	// paste contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Paste *runtime.RawExtension `json:"paste,omitempty"`
	// placement contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Placement *runtime.RawExtension `json:"placement,omitempty"`
	// plugins contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Plugins *runtime.RawExtension `json:"plugins,omitempty"`
	// pool contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Pool *runtime.RawExtension `json:"pool,omitempty"`
	// postgresql contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Postgresql *runtime.RawExtension `json:"postgresql,omitempty"`
	// prometheus contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Prometheus *runtime.RawExtension `json:"prometheus,omitempty"`
	// provisioning contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Provisioning *runtime.RawExtension `json:"provisioning,omitempty"`
	// qemu contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Qemu *runtime.RawExtension `json:"qemu,omitempty"`
	// rgw_ks contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	RgwKs *runtime.RawExtension `json:"rgw_ks,omitempty"`
	// rgw_s3 contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	RgwS3 *runtime.RawExtension `json:"rgw_s3,omitempty"`
	// scheduler contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Scheduler *runtime.RawExtension `json:"scheduler,omitempty"`
	// security contains tbd
	Security string `json:"security,omitempty"`
	// shipyard contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Shipyard *runtime.RawExtension `json:"shipyard,omitempty"`
	// software contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Software *runtime.RawExtension `json:"software,omitempty"`
	// spiceproxy contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Spiceproxy *runtime.RawExtension `json:"spiceproxy,omitempty"`
	// ssh contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Ssh *runtime.RawExtension `json:"ssh,omitempty"`
	// storage contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Storage *runtime.RawExtension `json:"storage,omitempty"`
	// sysctl contains tbd
	Sysctl *map[string]string `json:"sysctl,omitempty"`
	// test contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Test *runtime.RawExtension `json:"test,omitempty"`
	// uamlite contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Uamlite *runtime.RawExtension `json:"uamlite,omitempty"`
}

type AVDevelopment struct {
	// enabled contains tbd
	Enabled bool `json:"enabled,omitempty"`
}

type AVNetwork struct {
	// api contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Api *runtime.RawExtension `json:"api,omitempty"`
	// backend contains tbd
	Backend []string `json:"backend,omitempty"`
	// cluster contains tbd
	Cluster string `json:"cluster,omitempty"`
	// drydock contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Drydock *runtime.RawExtension `json:"drydock,omitempty"`
	// host_namespace contains tbd
	HostNamespace bool `json:"host_namespace,omitempty"`
	// ingress contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Ingress *runtime.RawExtension `json:"ingress,omitempty"`
	// interface contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Interface *runtime.RawExtension `json:"interface,omitempty"`
	// kubernetes_netloc contains tbd
	KubernetesNetloc string `json:"kubernetes_netloc,omitempty"`
	// kubernetes_service_ip contains tbd
	KubernetesServiceIp string `json:"kubernetes_service_ip,omitempty"`
	// maas_ingress contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	MaasIngress *runtime.RawExtension `json:"maas_ingress,omitempty"`
	// pod_cidr contains tbd
	PodCidr string `json:"pod_cidr,omitempty"`
	// public contains tbd
	Public string `json:"public,omitempty"`
	// region_api contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	RegionApi *runtime.RawExtension `json:"region_api,omitempty"`
	// region_proxy contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	RegionProxy *runtime.RawExtension `json:"region_proxy,omitempty"`
	// service_cidr contains tbd
	ServiceCidr string `json:"service_cidr,omitempty"`
	// service_client contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	ServiceClient *runtime.RawExtension `json:"service_client,omitempty"`
	// service_peer contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	ServicePeer *runtime.RawExtension `json:"service_peer,omitempty"`
	// vip contains tbd
	// +k8s:openapi-gen:schema-type-format=object
	Vip *runtime.RawExtension `json:"vip,omitempty"`
}

type AVService struct {
	// ip contains tbd
	Ip string `json:"ip,omitempty"`
	// name contains tbd
	Name string `json:"name,omitempty"`
}
