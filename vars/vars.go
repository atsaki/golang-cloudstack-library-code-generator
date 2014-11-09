package vars

var ApiToObjectNameMap = map[string]string{
	"activateProject":                             "project",
	"addAccountToProject":                         "result",
	"addBaremetalDhcp":                            "baremetaldhcp",
	"addBaremetalHost":                            "host",
	"addBaremetalPxeKickStartServer":              "baremetalpxekickstartserver",
	"addBaremetalPxePingServer":                   "baremetalpxepingserver",
	"addCiscoAsa1000vResource":                    "ciscoasa1000vresource",
	"addCiscoVnmcResource":                        "ciscovnmcresource",
	"addCluster":                                  "cluster",
	"addExternalFirewall":                         "externalfirewall",
	"addExternalLoadBalancer":                     "externalloadbalancer",
	"addF5LoadBalancer":                           "f5loadbalancer",
	"addHost":                                     "host",
	"addImageStore":                               "imagestore",
	"addIpToNic":                                  "iptonic",
	"addNetscalerLoadBalancer":                    "netscalerloadbalancer",
	"addNetworkDevice":                            "networkdevice",
	"addNetworkServiceProvider":                   "networkserviceprovider",
	"addNicToVirtualMachine":                      "virtualmachine",
	"addNiciraNvpDevice":                          "niciranvpdevice",
	"addRegion":                                   "region",
	"addResourceDetail":                           "result",
	"addS3":                                       "s3",
	"addSecondaryStorage":                         "secondarystorage",
	"addSrxFirewall":                              "srxfirewall",
	"addSwift":                                    "swift",
	"addTrafficMonitor":                           "trafficmonitor",
	"addTrafficType":                              "traffictype",
	"addUcsManager":                               "ucsmanager",
	"addVmwareDc":                                 "vmwaredc",
	"addVpnUser":                                  "vpnuser",
	"archiveAlerts":                               "result",
	"archiveEvents":                               "result",
	"assignToGlobalLoadBalancerRule":              "result",
	"assignToLoadBalancerRule":                    "result",
	"assignVirtualMachine":                        "virtualmachine",
	"associateIpAddress":                          "publicipaddress",
	"associateLun":                                "lun",
	"associateUcsProfileToBlade":                  "ucsprofiletoblade",
	"attachIso":                                   "virtualmachine",
	"attachVolume":                                "volume",
	"authorizeSecurityGroupEgress":                "securitygroupegress",
	"authorizeSecurityGroupIngress":               "securitygroupingress",
	"cancelHostMaintenance":                       "host",
	"cancelStorageMaintenance":                    "storagemaintenance",
	"changeServiceForRouter":                      "router",
	"changeServiceForSystemVm":                    "systemvm",
	"changeServiceForVirtualMachine":              "virtualmachine",
	"cleanVMReservations":                         "result",
	"configureF5LoadBalancer":                     "f5loadbalancer",
	"configureInternalLoadBalancerElement":        "internalloadbalancerelement",
	"configureNetscalerLoadBalancer":              "netscalerloadbalancer",
	"configureSrxFirewall":                        "srxfirewall",
	"configureVirtualRouterElement":               "virtualrouterelement",
	"copyIso":                                     "iso",
	"copyTemplate":                                "template",
	"createAccount":                               "account",
	"createAffinityGroup":                         "affinitygroup",
	"createAutoScalePolicy":                       "autoscalepolicy",
	"createAutoScaleVmGroup":                      "autoscalevmgroup",
	"createAutoScaleVmProfile":                    "autoscalevmprofile",
	"createCondition":                             "condition",
	"createCounter":                               "counter",
	"createDiskOffering":                          "diskoffering",
	"createDomain":                                "domain",
	"createEgressFirewallRule":                    "egressfirewallrule",
	"createFirewallRule":                          "firewallrule",
	"createGlobalLoadBalancerRule":                "globalloadbalancerrule",
	"createInstanceGroup":                         "instancegroup",
	"createInternalLoadBalancerElement":           "internalloadbalancerelement",
	"createIpForwardingRule":                      "ipforwardingrule",
	"createLBHealthCheckPolicy":                   "lbhealthcheckpolicy",
	"createLBStickinessPolicy":                    "lbstickinesspolicy",
	"createLoadBalancer":                          "loadbalancer",
	"createLoadBalancerRule":                      "loadbalancerrule",
	"createLunOnFiler":                            "lunonfiler",
	"createNetwork":                               "network",
	"createNetworkACL":                            "networkacl",
	"createNetworkACLList":                        "networkacllist",
	"createNetworkOffering":                       "networkoffering",
	"createPhysicalNetwork":                       "physicalnetwork",
	"createPod":                                   "pod",
	"createPool":                                  "pool",
	"createPortForwardingRule":                    "portforwardingrule",
	"createPortableIpRange":                       "portableiprange",
	"createPrivateGateway":                        "privategateway",
	"createProject":                               "project",
	"createRemoteAccessVpn":                       "remoteaccessvpn",
	"createSSHKeyPair":                            "keypair",
	"createSecondaryStagingStore":                 "secondarystagingstore",
	"createSecurityGroup":                         "securitygroup",
	"createServiceOffering":                       "serviceoffering",
	"createSnapshot":                              "snapshot",
	"createSnapshotPolicy":                        "snapshotpolicy",
	"createStaticRoute":                           "staticroute",
	"createStorageNetworkIpRange":                 "storagenetworkiprange",
	"createStoragePool":                           "storagepool",
	"createTags":                                  "result",
	"createTemplate":                              "template",
	"createUser":                                  "user",
	"createVMSnapshot":                            "vmsnapshot",
	"createVPC":                                   "vpc",
	"createVPCOffering":                           "vpcoffering",
	"createVirtualRouterElement":                  "virtualrouterelement",
	"createVlanIpRange":                           "vlaniprange",
	"createVolume":                                "volume",
	"createVolumeOnFiler":                         "volumeonfiler",
	"createVpnConnection":                         "vpnconnection",
	"createVpnCustomerGateway":                    "vpncustomergateway",
	"createVpnGateway":                            "vpngateway",
	"createZone":                                  "zone",
	"dedicateCluster":                             "cluster",
	"dedicateGuestVlanRange":                      "guestvlanrange",
	"dedicateHost":                                "host",
	"dedicatePod":                                 "pod",
	"dedicatePublicIpRange":                       "publiciprange",
	"dedicateZone":                                "zone",
	"deleteAccount":                               "result",
	"deleteAccountFromProject":                    "result",
	"deleteAffinityGroup":                         "result",
	"deleteAlerts":                                "result",
	"deleteAutoScalePolicy":                       "result",
	"deleteAutoScaleVmGroup":                      "result",
	"deleteAutoScaleVmProfile":                    "result",
	"deleteCiscoAsa1000vResource":                 "result",
	"deleteCiscoNexusVSM":                         "result",
	"deleteCiscoVnmcResource":                     "result",
	"deleteCluster":                               "result",
	"deleteCondition":                             "result",
	"deleteCounter":                               "result",
	"deleteDiskOffering":                          "result",
	"deleteDomain":                                "result",
	"deleteEgressFirewallRule":                    "result",
	"deleteEvents":                                "result",
	"deleteExternalFirewall":                      "result",
	"deleteExternalLoadBalancer":                  "result",
	"deleteF5LoadBalancer":                        "result",
	"deleteFirewallRule":                          "result",
	"deleteGlobalLoadBalancerRule":                "result",
	"deleteHost":                                  "result",
	"deleteImageStore":                            "result",
	"deleteInstanceGroup":                         "result",
	"deleteIpForwardingRule":                      "result",
	"deleteIso":                                   "result",
	"deleteLBHealthCheckPolicy":                   "result",
	"deleteLBStickinessPolicy":                    "result",
	"deleteLoadBalancer":                          "result",
	"deleteLoadBalancerRule":                      "result",
	"deleteNetscalerLoadBalancer":                 "result",
	"deleteNetwork":                               "result",
	"deleteNetworkACL":                            "result",
	"deleteNetworkACLList":                        "result",
	"deleteNetworkDevice":                         "result",
	"deleteNetworkOffering":                       "result",
	"deleteNetworkServiceProvider":                "result",
	"deleteNiciraNvpDevice":                       "result",
	"deletePhysicalNetwork":                       "result",
	"deletePod":                                   "result",
	"deletePool":                                  "result",
	"deletePortForwardingRule":                    "result",
	"deletePortableIpRange":                       "result",
	"deletePrivateGateway":                        "result",
	"deleteProject":                               "result",
	"deleteProjectInvitation":                     "result",
	"deleteRemoteAccessVpn":                       "result",
	"deleteSSHKeyPair":                            "result",
	"deleteSecondaryStagingStore":                 "result",
	"deleteSecurityGroup":                         "result",
	"deleteServiceOffering":                       "result",
	"deleteSnapshot":                              "result",
	"deleteSnapshotPolicies":                      "result",
	"deleteSrxFirewall":                           "result",
	"deleteStaticRoute":                           "result",
	"deleteStorageNetworkIpRange":                 "result",
	"deleteStoragePool":                           "result",
	"deleteTags":                                  "result",
	"deleteTemplate":                              "result",
	"deleteTrafficMonitor":                        "result",
	"deleteTrafficType":                           "result",
	"deleteUcsManager":                            "result",
	"deleteUser":                                  "result",
	"deleteVMSnapshot":                            "result",
	"deleteVPC":                                   "result",
	"deleteVPCOffering":                           "result",
	"deleteVlanIpRange":                           "result",
	"deleteVolume":                                "result",
	"deleteVpnConnection":                         "result",
	"deleteVpnCustomerGateway":                    "result",
	"deleteVpnGateway":                            "result",
	"deleteZone":                                  "result",
	"deployVirtualMachine":                        "virtualmachine",
	"destroyLunOnFiler":                           "lunonfiler",
	"destroyRouter":                               "router",
	"destroySystemVm":                             "systemvm",
	"destroyVirtualMachine":                       "virtualmachine",
	"destroyVolumeOnFiler":                        "volumeonfiler",
	"detachIso":                                   "virtualmachine",
	"detachVolume":                                "volume",
	"disableAccount":                              "account",
	"disableAutoScaleVmGroup":                     "autoscalevmgroup",
	"disableCiscoNexusVSM":                        "cisconexusvsm",
	"disableStaticNat":                            "result",
	"disableUser":                                 "user",
	"disassociateIpAddress":                       "result",
	"disassociateUcsProfileFromBlade":             "ucsprofilefromblade",
	"dissociateLun":                               "lun",
	"enableAccount":                               "account",
	"enableAutoScaleVmGroup":                      "autoscalevmgroup",
	"enableCiscoNexusVSM":                         "cisconexusvsm",
	"enableStaticNat":                             "result",
	"enableStorageMaintenance":                    "storagemaintenance",
	"enableUser":                                  "user",
	"extractIso":                                  "iso",
	"extractTemplate":                             "template",
	"extractVolume":                               "volume",
	"findHostsForMigration":                       "host",
	"findStoragePoolsForMigration":                "storagepool",
	"generateUsageRecords":                        "result",
	"getApiLimit":                                 "apilimit",
	"getCloudIdentifier":                          "cloudidentifier",
	"getUser":                                     "user",
	"getVMPassword":                               "vmpassword",
	"instantiateUcsTemplateAndAssocaciateToBlade": "ucstemplateandassocaciatetoblade",
	"ldapConfig":                                  "ldapconfig",
	"ldapRemove":                                  "ldapremove",
	"listAccounts":                                "account",
	"listAffinityGroupTypes":                      "affinitygrouptype",
	"listAffinityGroups":                          "affinitygroup",
	"listAlerts":                                  "alert",
	"listApis":                                    "api",
	"listAsyncJobs":                               "asyncjob",
	"listAutoScalePolicies":                       "autoscalepolicy",
	"listAutoScaleVmGroups":                       "autoscalevmgroup",
	"listAutoScaleVmProfiles":                     "autoscalevmprofile",
	"listBaremetalDhcp":                           "baremetaldhcp",
	"listBaremetalPxeServers":                     "baremetalpxeserver",
	"listCapabilities":                            "capability",
	"listCapacity":                                "capacity",
	"listCiscoAsa1000vResources":                  "ciscoasa1000vresource",
	"listCiscoNexusVSMs":                          "cisconexusvsm",
	"listCiscoVnmcResources":                      "ciscovnmcresource",
	"listClusters":                                "cluster",
	"listConditions":                              "condition",
	"listConfigurations":                          "configuration",
	"listCounters":                                "counter",
	"listDedicatedClusters":                       "dedicatedcluster",
	"listDedicatedGuestVlanRanges":                "dedicatedguestvlanrange",
	"listDedicatedHosts":                          "dedicatedhost",
	"listDedicatedPods":                           "dedicatedpod",
	"listDedicatedZones":                          "dedicatedzone",
	"listDeploymentPlanners":                      "deploymentplanner",
	"listDiskOfferings":                           "diskoffering",
	"listDomainChildren":                          "domainchildren",
	"listDomains":                                 "domain",
	"listEgressFirewallRules":                     "egressfirewallrule",
	"listEventTypes":                              "eventtype",
	"listEvents":                                  "event",
	"listExternalFirewalls":                       "externalfirewall",
	"listExternalLoadBalancers":                   "externalloadbalancer",
	"listF5LoadBalancerNetworks":                  "f5loadbalancernetwork",
	"listF5LoadBalancers":                         "f5loadbalancer",
	"listFirewallRules":                           "firewallrule",
	"listGlobalLoadBalancerRules":                 "globalloadbalancerrule",
	"listHosts":                                   "host",
	"listHypervisorCapabilities":                  "hypervisorcapability",
	"listHypervisors":                             "hypervisor",
	"listImageStores":                             "imagestore",
	"listInstanceGroups":                          "instancegroup",
	"listInternalLoadBalancerElements":            "internalloadbalancerelement",
	"listInternalLoadBalancerVMs":                 "internalloadbalancervm",
	"listIpForwardingRules":                       "ipforwardingrule",
	"listIsoPermissions":                          "isopermission",
	"listIsos":                                    "iso",
	"listLBHealthCheckPolicies":                   "lbhealthcheckpolicy",
	"listLBStickinessPolicies":                    "lbstickinesspolicy",
	"listLoadBalancerRuleInstances":               "loadbalancerruleinstance",
	"listLoadBalancerRules":                       "loadbalancerrule",
	"listLoadBalancers":                           "loadbalancer",
	"listLunsOnFiler":                             "lunsonfiler",
	"listNetscalerLoadBalancerNetworks":           "netscalerloadbalancernetwork",
	"listNetscalerLoadBalancers":                  "netscalerloadbalancer",
	"listNetworkACLLists":                         "networkacllist",
	"listNetworkACLs":                             "networkacl",
	"listNetworkDevice":                           "networkdevice",
	"listNetworkIsolationMethods":                 "networkisolationmethod",
	"listNetworkOfferings":                        "networkoffering",
	"listNetworkServiceProviders":                 "networkserviceprovider",
	"listNetworks":                                "network",
	"listNiciraNvpDeviceNetworks":                 "niciranvpdevicenetwork",
	"listNiciraNvpDevices":                        "niciranvpdevice",
	"listNics":                                    "nic",
	"listOsCategories":                            "oscategory",
	"listOsTypes":                                 "ostype",
	"listPhysicalNetworks":                        "physicalnetwork",
	"listPods":                                    "pod",
	"listPools":                                   "pool",
	"listPortForwardingRules":                     "portforwardingrule",
	"listPortableIpRanges":                        "portableiprange",
	"listPrivateGateways":                         "privategateway",
	"listProjectAccounts":                         "projectaccount",
	"listProjectInvitations":                      "projectinvitation",
	"listProjects":                                "project",
	"listPublicIpAddresses":                       "publicipaddress",
	"listRegions":                                 "region",
	"listRemoteAccessVpns":                        "remoteaccessvpn",
	"listResourceDetails":                         "resourcedetail",
	"listResourceLimits":                          "resourcelimit",
	"listRouters":                                 "router",
	"listS3s":                                     "s3",
	"listSSHKeyPairs":                             "sshkeypair",
	"listSecondaryStagingStores":                  "secondarystagingstore",
	"listSecurityGroups":                          "securitygroup",
	"listServiceOfferings":                        "serviceoffering",
	"listSnapshotPolicies":                        "snapshotpolicy",
	"listSnapshots":                               "snapshot",
	"listSrxFirewallNetworks":                     "srxfirewallnetwork",
	"listSrxFirewalls":                            "srxfirewall",
	"listStaticRoutes":                            "staticroute",
	"listStorageNetworkIpRange":                   "storagenetworkiprange",
	"listStoragePools":                            "storagepool",
	"listStorageProviders":                        "storageprovider",
	"listSupportedNetworkServices":                "networkservice",
	"listSwifts":                                  "swift",
	"listSystemVms":                               "systemvm",
	"listTags":                                    "tag",
	"listTemplatePermissions":                     "templatepermission",
	"listTemplates":                               "template",
	"listTrafficMonitors":                         "trafficmonitor",
	"listTrafficTypeImplementors":                 "traffictypeimplementor",
	"listTrafficTypes":                            "traffictype",
	"listUcsBlades":                               "ucsblade",
	"listUcsManagers":                             "ucsmanager",
	"listUcsProfiles":                             "ucsprofile",
	"listUcsTemplates":                            "ucstemplate",
	"listUsageRecords":                            "usagerecord",
	"listUsageTypes":                              "usagetype",
	"listUsers":                                   "user",
	"listVMSnapshot":                              "vmsnapshot",
	"listVPCOfferings":                            "vpcoffering",
	"listVPCs":                                    "vpc",
	"listVirtualMachines":                         "virtualmachine",
	"listVirtualRouterElements":                   "virtualrouterelement",
	"listVlanIpRanges":                            "vlaniprange",
	"listVmwareDcs":                               "vmwaredc",
	"listVolumes":                                 "volume",
	"listVolumesOnFiler":                          "volumesonfiler",
	"listVpnConnections":                          "vpnconnection",
	"listVpnCustomerGateways":                     "vpncustomergateway",
	"listVpnGateways":                             "vpngateway",
	"listVpnUsers":                                "vpnuser",
	"listZones":                                   "zone",
	"lockAccount":                                 "account",
	"lockUser":                                    "user",
	"markDefaultZoneForAccount":                   "defaultzoneforaccount",
	"migrateSystemVm":                             "systemvm",
	"migrateVirtualMachine":                       "virtualmachine",
	"migrateVirtualMachineWithVolume":             "virtualmachine",
	"migrateVolume":                               "volume",
	"modifyPool":                                  "pool",
	"prepareHostForMaintenance":                   "host",
	"prepareTemplate":                             "template",
	"rebootRouter":                                "router",
	"rebootSystemVm":                              "systemvm",
	"rebootVirtualMachine":                        "virtualmachine",
	"reconnectHost":                               "host",
	"recoverVirtualMachine":                       "virtualmachine",
	"refreshUcsBlades":                            "ucsblades",
	"registerIso":                                 "iso",
	"registerSSHKeyPair":                          "sshkeypair",
	"registerTemplate":                            "template",
	"registerUserKeys":                            "userkeys",
	"releaseDedicatedCluster":                     "result",
	"releaseDedicatedGuestVlanRange":              "result",
	"releaseDedicatedHost":                        "result",
	"releaseDedicatedPod":                         "result",
	"releaseDedicatedZone":                        "result",
	"releaseHostReservation":                      "result",
	"releasePublicIpRange":                        "result",
	"removeFromGlobalLoadBalancerRule":            "result",
	"removeFromLoadBalancerRule":                  "result",
	"removeIpFromNic":                             "result",
	"removeNicFromVirtualMachine":                 "virtualmachine",
	"removeRegion":                                "result",
	"removeResourceDetail":                        "result",
	"removeVmwareDc":                              "result",
	"removeVpnUser":                               "result",
	"replaceNetworkACLList":                       "result",
	"resetApiLimit":                               "apilimit",
	"resetPasswordForVirtualMachine":              "virtualmachine",
	"resetSSHKeyForVirtualMachine":                "virtualmachine",
	"resetVpnConnection":                          "vpnconnection",
	"resizeVolume":                                "volume",
	"restartNetwork":                              "network",
	"restartVPC":                                  "vpc",
	"restoreVirtualMachine":                       "virtualmachine",
	"revertToVMSnapshot":                          "virtualmachine",
	"revokeSecurityGroupEgress":                   "result",
	"revokeSecurityGroupIngress":                  "result",
	"scaleSystemVm":                               "systemvm",
	"scaleVirtualMachine":                         "result",
	"startInternalLoadBalancerVM":                 "internalloadbalancervm",
	"startRouter":                                 "router",
	"startSystemVm":                               "systemvm",
	"startVirtualMachine":                         "virtualmachine",
	"stopInternalLoadBalancerVM":                  "internalloadbalancervm",
	"stopRouter":                                  "router",
	"stopSystemVm":                                "systemvm",
	"stopVirtualMachine":                          "virtualmachine",
	"suspendProject":                              "project",
	"updateAccount":                               "account",
	"updateAutoScalePolicy":                       "autoscalepolicy",
	"updateAutoScaleVmGroup":                      "autoscalevmgroup",
	"updateAutoScaleVmProfile":                    "autoscalevmprofile",
	"updateCluster":                               "cluster",
	"updateConfiguration":                         "configuration",
	"updateDefaultNicForVirtualMachine":           "defaultnicforvirtualmachine",
	"updateDiskOffering":                          "diskoffering",
	"updateDomain":                                "domain",
	"updateGlobalLoadBalancerRule":                "globalloadbalancerrule",
	"updateHost":                                  "host",
	"updateHostPassword":                          "result",
	"updateHypervisorCapabilities":                "hypervisorcapabilities",
	"updateInstanceGroup":                         "instancegroup",
	"updateIso":                                   "iso",
	"updateIsoPermissions":                        "result",
	"updateLoadBalancerRule":                      "loadbalancerrule",
	"updateNetwork":                               "network",
	"updateNetworkACLItem":                        "networkaclitem",
	"updateNetworkOffering":                       "networkoffering",
	"updateNetworkServiceProvider":                "networkserviceprovider",
	"updatePhysicalNetwork":                       "physicalnetwork",
	"updatePod":                                   "pod",
	"updatePortForwardingRule":                    "portforwardingrule",
	"updateProject":                               "project",
	"updateProjectInvitation":                     "result",
	"updateRegion":                                "region",
	"updateResourceCount":                         "resourcecount",
	"updateResourceLimit":                         "resourcelimit",
	"updateServiceOffering":                       "serviceoffering",
	"updateStorageNetworkIpRange":                 "storagenetworkiprange",
	"updateStoragePool":                           "storagepool",
	"updateTemplate":                              "template",
	"updateTemplatePermissions":                   "result",
	"updateTrafficType":                           "traffictype",
	"updateUser":                                  "user",
	"updateVMAffinityGroup":                       "vmaffinitygroup",
	"updateVPC":                                   "vpc",
	"updateVPCOffering":                           "vpcoffering",
	"updateVirtualMachine":                        "virtualmachine",
	"updateVolume":                                "volume",
	"updateVpnCustomerGateway":                    "vpncustomergateway",
	"updateZone":                                  "zone",
	"uploadCustomCertificate":                     "customcertificate",
	"uploadVolume":                                "volume",
	"deleteOpenDaylightController":                "opendaylightcontroller",
	"upgradeRouterTemplate":                       "routertemplate",
	"addStratosphereSsp":                          "stratospheressp",
	"addGuestOs":                                  "guestos",
	"revertSnapshot":                              "snapshot",
	"addLdapConfiguration":                        "ldapconfiguration",
	"generateAlert":                               "result",
	"listBigSwitchVnsDevices":                     "bigswitchvnsdevice",
	"addGuestOsMapping":                           "guestosmapping",
	"configureOvsElement":                         "ovselement",
	"updateLBHealthCheckPolicy":                   "lbhealthcheckpolicy",
	"updateFirewallRule":                          "firewallrule",
	"removeGuestOs":                               "result",
	"configurePaloAltoFirewall":                   "paloaltofirewall",
	"updateVpnConnection":                         "vpnconnection",
	"updateVpnGateway":                            "vpngateway",
	"listSslCerts":                                "sslcert",
	"updateGuestOsMapping":                        "guestosmapping",
	"listPaloAltoFirewalls":                       "paloaltofirewall",
	"removeCertFromLoadBalancer":                  "result",
	"listOpenDaylightControllers":                 "opendaylightcontroller",
	"deleteLdapConfiguration":                     "ldapconfiguration",
	"listPaloAltoFirewallNetworks":                "paloaltofirewallnetwork",
	"createServiceInstance":                       "serviceinstance",
	"addOpenDaylightController":                   "opendaylightcontroller",
	"updateEgressFirewallRule":                    "egressfirewallrule",
	"updateGuestOs":                               "guestos",
	"deleteSslCert":                               "result",
	"addBigSwitchVnsDevice":                       "bigswitchvnsdevice",
	"deleteBigSwitchVnsDevice":                    "result",
	"ldapCreateAccount":                           "account",
	"assignCertToLoadBalancer":                    "result",
	"expungeVirtualMachine":                       "result",
	"removeGuestOsMapping":                        "result",
	"updateNetworkACLList":                        "result",
	"uploadSslCert":                               "sslcert",
	"updateLBStickinessPolicy":                    "lbstickinesspolicy",
	"listOvsElements":                             "ovselement",
	"listLdapConfigurations":                      "ldapconfiguration",
	"listLdapUsers":                               "ldapuser",
	"getVirtualMachineUserData":                   "virtualmachineuserdata",
	"updateIpAddress":                             "publicipaddress",
	"addPaloAltoFirewall":                         "paloaltofirewall",
	"importLdapUsers":                             "ldapuser",
	"deletePaloAltoFirewall":                      "result",
	"updateCloudToUseObjectStore":                 "imagestore",
	"listGuestOsMapping":                          "guestosmapping",
	"updateRemoteAccessVpn":                       "remoteaccessvpn",
	"updateLoadBalancer":                          "loadbalancer",
}
