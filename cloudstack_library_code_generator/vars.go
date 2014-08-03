package cloudstack_library_code_generator

var ApiToObjectNameMap = map[string]string{
	"activateProject":                             "project",
	"addAccountToProject":                         "accounttoproject",
	"addBaremetalDhcp":                            "baremetaldhcp",
	"addBaremetalHost":                            "baremetalhost",
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
	"addNicToVirtualMachine":                      "nictovirtualmachine",
	"addNiciraNvpDevice":                          "niciranvpdevice",
	"addRegion":                                   "region",
	"addResourceDetail":                           "resourcedetail",
	"addS3":                                       "s3",
	"addSecondaryStorage":                         "secondarystorage",
	"addSrxFirewall":                              "srxfirewall",
	"addSwift":                                    "swift",
	"addTrafficMonitor":                           "trafficmonitor",
	"addTrafficType":                              "traffictype",
	"addUcsManager":                               "ucsmanager",
	"addVmwareDc":                                 "vmwaredc",
	"addVpnUser":                                  "vpnuser",
	"archiveAlerts":                               "alert",
	"archiveEvents":                               "event",
	"assignToGlobalLoadBalancerRule":              "togloballoadbalancerrule",
	"assignToLoadBalancerRule":                    "toloadbalancerrule",
	"assignVirtualMachine":                        "virtualmachine",
	"associateIpAddress":                          "ipaddress",
	"associateLun":                                "lun",
	"associateUcsProfileToBlade":                  "ucsprofiletoblade",
	"attachIso":                                   "iso",
	"attachVolume":                                "volume",
	"authorizeSecurityGroupEgress":                "securitygroupegress",
	"authorizeSecurityGroupIngress":               "securitygroupingress",
	"cancelHostMaintenance":                       "hostmaintenance",
	"cancelStorageMaintenance":                    "storagemaintenance",
	"changeServiceForRouter":                      "serviceforrouter",
	"changeServiceForSystemVm":                    "serviceforsystemvm",
	"changeServiceForVirtualMachine":              "serviceforvirtualmachine",
	"cleanVMReservations":                         "vmreservations",
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
	"createSSHKeyPair":                            "sshkeypair",
	"createSecondaryStagingStore":                 "secondarystagingstore",
	"createSecurityGroup":                         "securitygroup",
	"createServiceOffering":                       "serviceoffering",
	"createSnapshot":                              "snapshot",
	"createSnapshotPolicy":                        "snapshotpolicy",
	"createStaticRoute":                           "staticroute",
	"createStorageNetworkIpRange":                 "storagenetworkiprange",
	"createStoragePool":                           "storagepool",
	"createTags":                                  "tags",
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
	"deleteAccount":                               "account",
	"deleteAccountFromProject":                    "accountfromproject",
	"deleteAffinityGroup":                         "affinitygroup",
	"deleteAlerts":                                "alerts",
	"deleteAutoScalePolicy":                       "autoscalepolicy",
	"deleteAutoScaleVmGroup":                      "autoscalevmgroup",
	"deleteAutoScaleVmProfile":                    "autoscalevmprofile",
	"deleteCiscoAsa1000vResource":                 "ciscoasa1000vresource",
	"deleteCiscoNexusVSM":                         "cisconexusvsm",
	"deleteCiscoVnmcResource":                     "ciscovnmcresource",
	"deleteCluster":                               "cluster",
	"deleteCondition":                             "condition",
	"deleteCounter":                               "counter",
	"deleteDiskOffering":                          "diskoffering",
	"deleteDomain":                                "domain",
	"deleteEgressFirewallRule":                    "egressfirewallrule",
	"deleteEvents":                                "events",
	"deleteExternalFirewall":                      "externalfirewall",
	"deleteExternalLoadBalancer":                  "externalloadbalancer",
	"deleteF5LoadBalancer":                        "f5loadbalancer",
	"deleteFirewallRule":                          "firewallrule",
	"deleteGlobalLoadBalancerRule":                "globalloadbalancerrule",
	"deleteHost":                                  "host",
	"deleteImageStore":                            "imagestore",
	"deleteInstanceGroup":                         "instancegroup",
	"deleteIpForwardingRule":                      "ipforwardingrule",
	"deleteIso":                                   "iso",
	"deleteLBHealthCheckPolicy":                   "lbhealthcheckpolicy",
	"deleteLBStickinessPolicy":                    "lbstickinesspolicy",
	"deleteLoadBalancer":                          "loadbalancer",
	"deleteLoadBalancerRule":                      "loadbalancerrule",
	"deleteNetscalerLoadBalancer":                 "netscalerloadbalancer",
	"deleteNetwork":                               "network",
	"deleteNetworkACL":                            "networkacl",
	"deleteNetworkACLList":                        "networkacllist",
	"deleteNetworkDevice":                         "networkdevice",
	"deleteNetworkOffering":                       "networkoffering",
	"deleteNetworkServiceProvider":                "networkserviceprovider",
	"deleteNiciraNvpDevice":                       "niciranvpdevice",
	"deletePhysicalNetwork":                       "physicalnetwork",
	"deletePod":                                   "pod",
	"deletePool":                                  "pool",
	"deletePortForwardingRule":                    "portforwardingrule",
	"deletePortableIpRange":                       "portableiprange",
	"deletePrivateGateway":                        "privategateway",
	"deleteProject":                               "project",
	"deleteProjectInvitation":                     "projectinvitation",
	"deleteRemoteAccessVpn":                       "remoteaccessvpn",
	"deleteSSHKeyPair":                            "sshkeypair",
	"deleteSecondaryStagingStore":                 "secondarystagingstore",
	"deleteSecurityGroup":                         "securitygroup",
	"deleteServiceOffering":                       "serviceoffering",
	"deleteSnapshot":                              "snapshot",
	"deleteSnapshotPolicies":                      "snapshotpolicies",
	"deleteSrxFirewall":                           "srxfirewall",
	"deleteStaticRoute":                           "staticroute",
	"deleteStorageNetworkIpRange":                 "storagenetworkiprange",
	"deleteStoragePool":                           "storagepool",
	"deleteTags":                                  "tags",
	"deleteTemplate":                              "template",
	"deleteTrafficMonitor":                        "trafficmonitor",
	"deleteTrafficType":                           "traffictype",
	"deleteUcsManager":                            "ucsmanager",
	"deleteUser":                                  "user",
	"deleteVMSnapshot":                            "vmsnapshot",
	"deleteVPC":                                   "vpc",
	"deleteVPCOffering":                           "vpcoffering",
	"deleteVlanIpRange":                           "vlaniprange",
	"deleteVolume":                                "volume",
	"deleteVpnConnection":                         "vpnconnection",
	"deleteVpnCustomerGateway":                    "vpncustomergateway",
	"deleteVpnGateway":                            "vpngateway",
	"deleteZone":                                  "zone",
	"deployVirtualMachine":                        "virtualmachine",
	"destroyLunOnFiler":                           "lunonfiler",
	"destroyRouter":                               "router",
	"destroySystemVm":                             "systemvm",
	"destroyVirtualMachine":                       "virtualmachine",
	"destroyVolumeOnFiler":                        "volumeonfiler",
	"detachIso":                                   "iso",
	"detachVolume":                                "volume",
	"disableAccount":                              "account",
	"disableAutoScaleVmGroup":                     "autoscalevmgroup",
	"disableCiscoNexusVSM":                        "cisconexusvsm",
	"disableStaticNat":                            "staticnat",
	"disableUser":                                 "user",
	"disassociateIpAddress":                       "ipaddress",
	"disassociateUcsProfileFromBlade":             "ucsprofilefromblade",
	"dissociateLun":                               "lun",
	"enableAccount":                               "account",
	"enableAutoScaleVmGroup":                      "autoscalevmgroup",
	"enableCiscoNexusVSM":                         "cisconexusvsm",
	"enableStaticNat":                             "staticnat",
	"enableStorageMaintenance":                    "storagemaintenance",
	"enableUser":                                  "user",
	"extractIso":                                  "iso",
	"extractTemplate":                             "template",
	"extractVolume":                               "volume",
	"findHostsForMigration":                       "hostsformigration",
	"findStoragePoolsForMigration":                "storagepoolsformigration",
	"generateUsageRecords":                        "usagerecords",
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
	"listSupportedNetworkServices":                "supportednetworkservice",
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
	"migrateVirtualMachineWithVolume":             "virtualmachinewithvolume",
	"migrateVolume":                               "volume",
	"modifyPool":                                  "pool",
	"prepareHostForMaintenance":                   "hostformaintenance",
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
	"releaseDedicatedCluster":                     "dedicatedcluster",
	"releaseDedicatedGuestVlanRange":              "dedicatedguestvlanrange",
	"releaseDedicatedHost":                        "dedicatedhost",
	"releaseDedicatedPod":                         "dedicatedpod",
	"releaseDedicatedZone":                        "dedicatedzone",
	"releaseHostReservation":                      "hostreservation",
	"releasePublicIpRange":                        "publiciprange",
	"removeFromGlobalLoadBalancerRule":            "fromgloballoadbalancerrule",
	"removeFromLoadBalancerRule":                  "fromloadbalancerrule",
	"removeIpFromNic":                             "ipfromnic",
	"removeNicFromVirtualMachine":                 "nicfromvirtualmachine",
	"removeRegion":                                "region",
	"removeResourceDetail":                        "resourcedetail",
	"removeVmwareDc":                              "vmwaredc",
	"removeVpnUser":                               "vpnuser",
	"replaceNetworkACLList":                       "networkacllist",
	"resetApiLimit":                               "apilimit",
	"resetPasswordForVirtualMachine":              "passwordforvirtualmachine",
	"resetSSHKeyForVirtualMachine":                "sshkeyforvirtualmachine",
	"resetVpnConnection":                          "vpnconnection",
	"resizeVolume":                                "volume",
	"restartNetwork":                              "network",
	"restartVPC":                                  "vpc",
	"restoreVirtualMachine":                       "virtualmachine",
	"revertToVMSnapshot":                          "vmsnapshot",
	"revokeSecurityGroupEgress":                   "securitygroupegress",
	"revokeSecurityGroupIngress":                  "securitygroupingress",
	"scaleSystemVm":                               "systemvm",
	"scaleVirtualMachine":                         "virtualmachine",
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
	"updateHostPassword":                          "hostpassword",
	"updateHypervisorCapabilities":                "hypervisorcapabilities",
	"updateInstanceGroup":                         "instancegroup",
	"updateIso":                                   "iso",
	"updateIsoPermissions":                        "isopermissions",
	"updateLoadBalancerRule":                      "loadbalancerrule",
	"updateNetwork":                               "network",
	"updateNetworkACLItem":                        "networkaclitem",
	"updateNetworkOffering":                       "networkoffering",
	"updateNetworkServiceProvider":                "networkserviceprovider",
	"updatePhysicalNetwork":                       "physicalnetwork",
	"updatePod":                                   "pod",
	"updatePortForwardingRule":                    "portforwardingrule",
	"updateProject":                               "project",
	"updateProjectInvitation":                     "projectinvitation",
	"updateRegion":                                "region",
	"updateResourceCount":                         "resourcecount",
	"updateResourceLimit":                         "resourcelimit",
	"updateServiceOffering":                       "serviceoffering",
	"updateStorageNetworkIpRange":                 "storagenetworkiprange",
	"updateStoragePool":                           "storagepool",
	"updateTemplate":                              "template",
	"updateTemplatePermissions":                   "templatepermissions",
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
}
