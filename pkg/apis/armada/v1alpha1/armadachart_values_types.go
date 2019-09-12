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

import ()

type AVPod struct {
	// affinity contains tbd
	Affinity *AVPodAffinity `json:"affinity,omitempty"`
	// env contains tbd
	Env *AVPodEnv `json:"env,omitempty"`
	// mount_path contains tbd
	MountPath string `json:"mount_path,omitempty"`
	// lifecycle contains tbd
	Lifecycle *AVPodLifecycle `json:"lifecycle,omitempty"`
	// replicas contains tbd
	Replicas map[string]int `json:"replicas,omitempty"`
	// resources contains tbd
	Resources map[string]AVPodResources `json:"resources,omitempty"`
	// security_context contains tbd
	SecurityContext *AVPodSecurityContext `json:"security_context,omitempty"`
}

type AVPodSecurityContext struct {
}

type AVPodResources struct {
	// curator contains tbd
	Curator *AVPodResourceCurator `json:"curator,omitempty"`
	// fluentbit contains tbd
	Fluentbit *AVPodResourceFluentbit `json:"fluentbit,omitempty"`
	// limits contains tbd
	Limits *AVPodResourceSettings `json:"limits,omitempty"`
	// requests contains tbd
	Requests *AVPodResourceSettings `json:"requests,omitempty"`
	// image_repo_sync contains tbd
	ImageRepoSync *AVPodResourceImageRepoSync `json:"image_repo_sync,omitempty"`
	// snapshot_repository contains tbd
	SnapshortRepository *AVPodResourceSnapshotRepository `json:"snapshot_repository,omitempty"`
	// tests contains tbd
	Tests *AVPodResourceTest `json:"tests,omitempty"`
}

type AVPodResourceCurator struct {
}

type AVPodResourceFluentbit struct {
}

type AVPodResourceSettings struct {
	// cpu contains tbd
	Cpu string `json:"cpu,omitempty"`
	// memory contains tbd
	Memory string `json:"memory,omitempty"`
}

type AVPodResourceImageRepoSync struct {
}

type AVPodResourceTest struct {
}

type AVPodResourceSnapshotRepository struct {
}

type AVPodLifecycle struct {
}

type AVPodReplicas struct {
}

type AVPodAffinity struct {
}

type AVPodEnv struct {
}

type AVStorageclass struct {
}

type AVNetworking struct {
}

type AVLivenessprobe struct {
}

type AVImages struct {
	// tags contains tbd
	Tags map[string]string `json:"tags,omitempty"`
	// pull_policy contains tbd
	PullPolicy string `json:"pull_policy,omitempty"`
}

type AVData struct {
}

type AVDependencies struct {
}

type AVGlobal struct {
}

type AVJobs struct {
}

type AVNodes struct {
}

type AVStorage struct {
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

type AVAnchor struct {
}

type AVApiserver struct {
}

type AVMonitoring struct {
}

type AVKubeService struct {
}

type AVCephMgrModulesConfig struct {
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
	Acconfig *map[string]ArmadaMapString `json:"acconfig,omitempty"`
	// agent contains tbd
	Agent *map[string]ArmadaMapString `json:"agent,omitempty"`
	// anchor contains tbd
	Anchor *map[string]ArmadaMapString `json:"anchor,omitempty"`
	// apache contains tbd
	Apache *map[string]ArmadaMapString `json:"apache,omitempty"`
	// api_metadata contains tbd
	ApiMetadata *map[string]ArmadaMapString `json:"api_metadata,omitempty"`
	// armada contains tbd
	Armada *map[string]ArmadaMapString `json:"armada,omitempty"`
	// auto_bridge_add contains tbd
	AutoBridgeAdd *map[string]ArmadaMapString `json:"auto_bridge_add,omitempty"`
	// cache contains tbd
	Cache *map[string]ArmadaMapString `json:"cache,omitempty"`
	// ceph contains tbd
	Ceph *map[string]ArmadaMapString `json:"ceph,omitempty"`
	// cni_network_config contains tbd
	CniNetworkConfig *map[string]ArmadaMapString `json:"cni_network_config,omitempty"`
	// conductor contains tbd
	Conductor *map[string]ArmadaMapString `json:"conductor,omitempty"`
	// config contains tbd
	Config *map[string]ArmadaMapString `json:"config,omitempty"`
	// consoleauth contains tbd
	Consoleauth *map[string]ArmadaMapString `json:"consoleauth,omitempty"`
	// controllers contains tbd
	Controllers *map[string]ArmadaMapString `json:"controllers,omitempty"`
	// coredns contains tbd
	Coredns *map[string]ArmadaMapString `json:"coredns,omitempty"`
	// curator contains tbd
	Curator *map[string]ArmadaMapString `json:"curator,omitempty"`
	// deckhand contains tbd
	Deckhand *map[string]ArmadaMapString `json:"deckhand,omitempty"`
	// defaults contains tbd
	Defaults *map[string]ArmadaMapString `json:"defaults,omitempty"`
	// drydock contains tbd
	Drydock *map[string]ArmadaMapString `json:"drydock,omitempty"`
	// elasticsearch contains tbd
	Elasticsearch *map[string]ArmadaMapString `json:"elasticsearch,omitempty"`
	// encryption_provider contains tbd
	EncryptionProvider *map[string]ArmadaMapString `json:"encryption_provider,omitempty"`
	// eventconfig contains tbd
	Eventconfig *map[string]ArmadaMapString `json:"eventconfig,omitempty"`
	// exec contains tbd
	Exec *map[string]ArmadaMapString `json:"exec,omitempty"`
	// features contains tbd
	Features *map[string]ArmadaMapString `json:"features,omitempty"`
	// fluentbit contains tbd
	Fluentbit *map[string]ArmadaMapString `json:"fluentbit,omitempty"`
	// fluentd contains tbd
	Fluentd *map[string]ArmadaMapString `json:"fluentd,omitempty"`
	// httpd contains tbd
	Httpd *map[string]ArmadaMapString `json:"httpd,omitempty"`
	// ingress contains tbd
	Ingress *map[string]ArmadaMapString `json:"ingress,omitempty"`
	// job contains tbd
	Job *map[string]ArmadaMapString `json:"job,omitempty"`
	// keystone contains tbd
	Keystone *map[string]ArmadaMapString `json:"keystone,omitempty"`
	// ldap contains tbd
	Ldap *map[string]ArmadaMapString `json:"ldap,omitempty"`
	// logging contains tbd
	Logging *map[string]ArmadaMapString `json:"logging,omitempty"`
	// maas contains tbd
	Maas *map[string]ArmadaMapString `json:"maas,omitempty"`
	// metadata_agent contains tbd
	MetadataAgent *map[string]ArmadaMapString `json:"metadata_agent,omitempty"`
	// neutron contains tbd
	Neutron *map[string]ArmadaMapString `json:"neutron,omitempty"`
	// node contains tbd
	Node *map[string]ArmadaMapString `json:"node,omitempty"`
	// nova contains tbd
	Nova *map[string]ArmadaMapString `json:"nova,omitempty"`
	// novncproxy contains tbd
	Novncproxy *map[string]ArmadaMapString `json:"novncproxy,omitempty"`
	// openstack_version contains tbd
	OpenstackVersion *map[string]ArmadaMapString `json:"openstack_version,omitempty"`
	// osapi contains tbd
	Osapi *map[string]ArmadaMapString `json:"osapi,omitempty"`
	// osd contains tbd
	Osd *map[string]ArmadaMapString `json:"osd,omitempty"`
	// overrides contains tbd
	Overrides *map[string]ArmadaMapString `json:"overrides,omitempty"`
	// parsers contains tbd
	Parsers *map[string]ArmadaMapString `json:"parsers,omitempty"`
	// paste contains tbd
	Paste *map[string]ArmadaMapString `json:"paste,omitempty"`
	// placement contains tbd
	Placement *map[string]ArmadaMapString `json:"placement,omitempty"`
	// plugins contains tbd
	Plugins *map[string]ArmadaMapString `json:"plugins,omitempty"`
	// pool contains tbd
	Pool *map[string]ArmadaMapString `json:"pool,omitempty"`
	// postgresql contains tbd
	Postgresql *map[string]ArmadaMapString `json:"postgresql,omitempty"`
	// prometheus contains tbd
	Prometheus *map[string]ArmadaMapString `json:"prometheus,omitempty"`
	// provisioning contains tbd
	Provisioning *map[string]ArmadaMapString `json:"provisioning,omitempty"`
	// qemu contains tbd
	Qemu *map[string]ArmadaMapString `json:"qemu,omitempty"`
	// rgw_ks contains tbd
	RgwKs *map[string]ArmadaMapString `json:"rgw_ks,omitempty"`
	// rgw_s3 contains tbd
	RgwS3 *map[string]ArmadaMapString `json:"rgw_s3,omitempty"`
	// scheduler contains tbd
	Scheduler *map[string]ArmadaMapString `json:"scheduler,omitempty"`
	// security contains tbd
	Security *map[string]ArmadaMapString `json:"security,omitempty"`
	// shipyard contains tbd
	Shipyard *map[string]ArmadaMapString `json:"shipyard,omitempty"`
	// software contains tbd
	Software *map[string]ArmadaMapString `json:"software,omitempty"`
	// spiceproxy contains tbd
	Spiceproxy *map[string]ArmadaMapString `json:"spiceproxy,omitempty"`
	// ssh contains tbd
	Ssh *map[string]ArmadaMapString `json:"ssh,omitempty"`
	// storage contains tbd
	Storage *map[string]ArmadaMapString `json:"storage,omitempty"`
	// sysctl contains tbd
	Sysctl *map[string]ArmadaMapString `json:"sysctl,omitempty"`
	// test contains tbd
	Test *map[string]ArmadaMapString `json:"test,omitempty"`
	// uamlite contains tbd
	Uamlite *map[string]ArmadaMapString `json:"uamlite,omitempty"`
}

type AVDevelopment struct {
    // enabled contains tbd
    Enabled bool `json:"enabled,omitempty"`
}

type AVNetwork struct {
    // backend contains tbd
    Backend *map[string]ArmadaMapString `json:"backend,omitempty"`
    // cluster contains tbd
    Cluster *map[string]ArmadaMapString `json:"cluster,omitempty"`
    // drydock contains tbd
    Drydock *map[string]ArmadaMapString `json:"drydock,omitempty"`
    // host_namespace contains tbd
    HostNamespace *map[string]ArmadaMapString `json:"host_namespace,omitempty"`
    // ingress contains tbd
    Ingress *map[string]ArmadaMapString `json:"ingress,omitempty"`
    // interface contains tbd
    Interface *map[string]ArmadaMapString `json:"interface,omitempty"`
    // kubernetes_netloc contains tbd
    KubernetesNetloc *map[string]ArmadaMapString `json:"kubernetes_netloc,omitempty"`
    // kubernetes_service_ip contains tbd
    KubernetesServiceIp *map[string]ArmadaMapString `json:"kubernetes_service_ip,omitempty"`
    // maas_ingress contains tbd
    MaasIngress *map[string]ArmadaMapString `json:"maas_ingress,omitempty"`
    // pod_cidr contains tbd
    PodCidr *map[string]ArmadaMapString `json:"pod_cidr,omitempty"`
    // public contains tbd
    Public *map[string]ArmadaMapString `json:"public,omitempty"`
    // region_proxy contains tbd
    RegionProxy *map[string]ArmadaMapString `json:"region_proxy,omitempty"`
    // service_cidion_api contains tbd
    ServiceCidionApi *map[string]ArmadaMapString `json:"service_cidion_api,omitempty"`
    // service_peer contains tbd
    ServicePeer *map[string]ArmadaMapString `json:"service_peer,omitempty"`
    // vip contains tbd
    Vip *map[string]ArmadaMapString `json:"vip,omitempty"`
}

type AVService struct {
    // ip contains tbd
    Ip *map[string]ArmadaMapString `json:"ip,omitempty"`
    // name contains tbd
    Name *map[string]ArmadaMapString `json:"name,omitempty"`
}
