package nsptopoviewer

import "encoding/json"

type IETFNetworkStruct struct {
	Response struct {
		Status    int `json:"status"`
		StartRow  int `json:"startRow"`
		EndRow    int `json:"endRow"`
		TotalRows int `json:"totalRows"`
		Data      struct {
			TeTopologiesAugment interface{} `json:"teTopologiesAugment"`
			Network             []struct {
				NetworkID         string `json:"networkId"`
				ServerProvided    bool   `json:"serverProvided"`
				TeTopologyAugment struct {
					Te struct {
						ProviderID struct {
							Uint32 int `json:"uint32"`
						} `json:"providerId"`
						ClientID     interface{} `json:"clientId"`
						TeTopologyID struct {
							String string `json:"string"`
						} `json:"teTopologyId"`
						Config interface{} `json:"config"`
						State  struct {
							TeTopologyConfig struct {
								Schedules             interface{} `json:"schedules"`
								Preference            int         `json:"preference"`
								OptimizationCriterion struct {
								} `json:"optimizationCriterion"`
							} `json:"teTopologyConfig"`
						} `json:"state"`
					} `json:"te"`
				} `json:"teTopologyAugment"`
				NetworkTypes struct {
					TeTopologyType struct {
						TeTopology struct {
						} `json:"teTopology"`
					} `json:"teTopologyType"`
				} `json:"networkTypes"`
				SupportingNetwork interface{} `json:"supportingNetwork"`
				Node              []struct {
					NodeID        string `json:"nodeId"`
					TeNodeAugment struct {
						Te struct {
							TeNodeID struct {
								DottedQuad struct {
									String string `json:"string"`
								} `json:"dottedQuad"`
							} `json:"teNodeId"`
							Config interface{} `json:"config"`
							State  struct {
								TeNodeConfig       interface{} `json:"teNodeConfig"`
								TeNodeStateDerived struct {
									InformationSource      interface{} `json:"informationSource"`
									InformationSourceState interface{} `json:"informationSourceState"`
									OperStatus             string      `json:"operStatus"`
									IsMultiAccessDr        interface{} `json:"isMultiAccessDr"`
									InformationSourceEntry interface{} `json:"informationSourceEntry"`
								} `json:"teNodeStateDerived"`
							} `json:"state"`
							TunnelTerminationPoint interface{} `json:"tunnelTerminationPoint"`
						} `json:"te"`
					} `json:"teNodeAugment"`
					SupportingNode   interface{} `json:"supportingNode"`
					TerminationPoint []struct {
						TpID                      string `json:"tpId"`
						TeTerminationPointAugment struct {
							Te struct {
								TeTpID struct {
									Uint32    interface{} `json:"uint32"`
									IPAddress struct {
										Ipv6Address interface{} `json:"ipv6Address"`
										Ipv4Address struct {
											String string `json:"string"`
										} `json:"ipv4Address"`
									} `json:"ipAddress"`
								} `json:"teTpId"`
								Config interface{} `json:"config"`
								State  struct {
									TeTerminationPointConfig struct {
										Schedules                        interface{} `json:"schedules"`
										InterfaceSwitchingCapabilityList struct {
											InterfaceSwitchingCapability struct {
												SwitchingPsc1 struct {
													TimeDivisionMultiplexCapable interface{} `json:"timeDivisionMultiplexCapable"`
													MaxLspBandwidth              interface{} `json:"maxLspBandwidth"`
													SwitchingCapability          struct {
													} `json:"switchingCapability"`
													Encoding interface{} `json:"encoding"`
												} `json:"SwitchingPsc1"`
											} `json:"interfaceSwitchingCapability"`
										} `json:"interfaceSwitchingCapabilityList"`
										InterLayerLockID interface{} `json:"interLayerLockId"`
									} `json:"teTerminationPointConfig"`
								} `json:"state"`
							} `json:"te"`
						} `json:"teTerminationPointAugment"`
						SupportingTerminationPoint interface{} `json:"supportingTerminationPoint"`
					} `json:"terminationPoint"`
				} `json:"node"`
				Link []struct {
					LinkID        string `json:"linkId"`
					TeLinkAugment struct {
						Te struct {
							Config struct {
								TeLinkConfig struct {
									TeLinkAttributes struct {
										Schedules                          interface{} `json:"schedules"`
										AccessType                         interface{} `json:"accessType"`
										ExternalDomain                     interface{} `json:"externalDomain"`
										IsAbstract                         interface{} `json:"isAbstract"`
										Name                               interface{} `json:"name"`
										Underlay                           interface{} `json:"underlay"`
										AdminStatus                        interface{} `json:"adminStatus"`
										PerformanceMetricThrottleContainer interface{} `json:"performanceMetricThrottleContainer"`
										TeLinkInfoAttributes               struct {
											InterfaceSwitchingCapabilityList interface{} `json:"interfaceSwitchingCapabilityList"`
											TeLinkConnectivityAttributes     struct {
												MaxLinkBandwidth     interface{} `json:"maxLinkBandwidth"`
												MaxResvLinkBandwidth json.Number `json:"maxResvLinkBandwidth"`
												UnreservedBandwidth  interface{} `json:"unreservedBandwidth"`
												TeDefaultMetric      interface{} `json:"teDefaultMetric"`
												PerformanceMetric    struct {
													TePerformanceMetric interface{} `json:"tePerformanceMetric"`
													Measurement         struct {
														PerformanceMetricAttributes struct {
															UnidirectionalDelay              int         `json:"unidirectionalDelay"`
															UnidirectionalMinDelay           interface{} `json:"unidirectionalMinDelay"`
															UnidirectionalMaxDelay           interface{} `json:"unidirectionalMaxDelay"`
															UnidirectionalDelayVariation     interface{} `json:"unidirectionalDelayVariation"`
															UnidirectionalPacketLoss         interface{} `json:"unidirectionalPacketLoss"`
															UnidirectionalResidualBandwidth  interface{} `json:"unidirectionalResidualBandwidth"`
															UnidirectionalAvailableBandwidth interface{} `json:"unidirectionalAvailableBandwidth"`
															UnidirectionalUtilizedBandwidth  interface{} `json:"unidirectionalUtilizedBandwidth"`
														} `json:"performanceMetricAttributes"`
													} `json:"measurement"`
													Normality interface{} `json:"normality"`
												} `json:"performanceMetric"`
												TeSrlgs struct {
													Value []interface{} `json:"value"`
												} `json:"teSrlgs"`
											} `json:"teLinkConnectivityAttributes"`
											LinkIndex           interface{} `json:"linkIndex"`
											AdministrativeGroup interface{} `json:"administrativeGroup"`
											LinkProtectionType  interface{} `json:"linkProtectionType"`
										} `json:"teLinkInfoAttributes"`
									} `json:"teLinkAttributes"`
									BundledLinks   interface{} `json:"bundledLinks"`
									ComponentLinks interface{} `json:"componentLinks"`
									TeLinkTemplate interface{} `json:"teLinkTemplate"`
									Template       interface{} `json:"template"`
								} `json:"teLinkConfig"`
							} `json:"config"`
							State struct {
								TeLinkConfig struct {
									TeLinkAttributes struct {
										Schedules                          interface{} `json:"schedules"`
										AccessType                         string      `json:"accessType"`
										ExternalDomain                     interface{} `json:"externalDomain"`
										IsAbstract                         bool        `json:"isAbstract"`
										Name                               string      `json:"name"`
										Underlay                           interface{} `json:"underlay"`
										AdminStatus                        string      `json:"adminStatus"`
										PerformanceMetricThrottleContainer interface{} `json:"performanceMetricThrottleContainer"`
										TeLinkInfoAttributes               struct {
											InterfaceSwitchingCapabilityList interface{} `json:"interfaceSwitchingCapabilityList"`
											TeLinkConnectivityAttributes     struct {
												MaxLinkBandwidth     interface{} `json:"maxLinkBandwidth"`
												MaxResvLinkBandwidth json.Number `json:"maxResvLinkBandwidth"`
												UnreservedBandwidth  struct {
													Num0 json.Number `json:"0"`
												} `json:"unreservedBandwidth"`
												TeDefaultMetric   int         `json:"teDefaultMetric"`
												PerformanceMetric interface{} `json:"performanceMetric"`
												TeSrlgs           struct {
													Value []interface{} `json:"value"`
												} `json:"teSrlgs"`
											} `json:"teLinkConnectivityAttributes"`
											LinkIndex           interface{} `json:"linkIndex"`
											AdministrativeGroup struct {
												AdminGroup struct {
													Binary []int `json:"binary"`
												} `json:"adminGroup"`
												ExtendedAdminGroup interface{} `json:"extendedAdminGroup"`
											} `json:"administrativeGroup"`
											LinkProtectionType interface{} `json:"linkProtectionType"`
										} `json:"teLinkInfoAttributes"`
									} `json:"teLinkAttributes"`
									BundledLinks   interface{} `json:"bundledLinks"`
									ComponentLinks interface{} `json:"componentLinks"`
									TeLinkTemplate interface{} `json:"teLinkTemplate"`
									Template       interface{} `json:"template"`
								} `json:"teLinkConfig"`
								TeLinkStateDerived struct {
									InformationSource      interface{} `json:"informationSource"`
									InformationSourceState interface{} `json:"informationSourceState"`
									OperStatus             string      `json:"operStatus"`
									IsTransitional         interface{} `json:"isTransitional"`
									InformationSourceEntry interface{} `json:"informationSourceEntry"`
									Recovery               interface{} `json:"recovery"`
									Underlay               interface{} `json:"underlay"`
								} `json:"teLinkStateDerived"`
							} `json:"state"`
						} `json:"te"`
					} `json:"teLinkAugment"`
					Source struct {
						SourceNode string `json:"sourceNode"`
						SourceTp   string `json:"sourceTp"`
					} `json:"source"`
					Destination struct {
						DestNode string `json:"destNode"`
						DestTp   string `json:"destTp"`
					} `json:"destination"`
					SupportingLink interface{} `json:"supportingLink"`
				} `json:"link"`
			} `json:"network"`
		} `json:"data"`
	} `json:"response"`
}

type ServiceLayerStruct struct {
	NspServiceIntentIntent []struct {
		At struct {
			NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
			NspModelIdentifier   string `json:"nsp-model:identifier"`
		} `json:"@"`
		IntentType         string `json:"intent-type"`
		ServiceName        string `json:"service-name"`
		IntentTypeVersion  string `json:"intent-type-version"`
		OlcState           string `json:"olc-state"`
		TemplateName       string `json:"template-name"`
		IntentSpecificData struct {
			EvpnEpipeEvpnEpipe struct {
				At struct {
					NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
					NspModelIdentifier   string `json:"nsp-model:identifier"`
				} `json:"@"`
				NeServiceID int         `json:"ne-service-id"`
				Mtu         int         `json:"mtu"`
				EvpnType    string      `json:"evpn-type"`
				CustomerID  int         `json:"customer-id"`
				Description string      `json:"description"`
				AdminState  string      `json:"admin-state"`
				JobID       interface{} `json:"job-id"`
				SiteA       struct {
					At struct {
						NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
						NspModelIdentifier   string `json:"nsp-model:identifier"`
					} `json:"@"`
					DeviceID    string      `json:"device-id"`
					SiteName    interface{} `json:"site-name"`
					Description interface{} `json:"description"`
					Evi         int         `json:"evi"`
					Ecmp        int         `json:"ecmp"`
					LocalAc     struct {
						At struct {
							NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
							NspModelIdentifier   string `json:"nsp-model:identifier"`
						} `json:"@"`
						Name   string `json:"name"`
						EthTag int    `json:"eth-tag"`
					} `json:"local-ac"`
					RemoteAc struct {
						At struct {
							NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
							NspModelIdentifier   string `json:"nsp-model:identifier"`
						} `json:"@"`
						Name   string `json:"name"`
						EthTag int    `json:"eth-tag"`
					} `json:"remote-ac"`
					Mpls struct {
						At struct {
							NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
							NspModelIdentifier   string `json:"nsp-model:identifier"`
						} `json:"@"`
						ForceVcForwarding string `json:"force-vc-forwarding"`
						BgpInstance       struct {
							At struct {
								NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
								NspModelIdentifier   string `json:"nsp-model:identifier"`
							} `json:"@"`
							RouteDistinguisher string `json:"route-distinguisher"`
							RouteTarget        []struct {
								At struct {
									NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
									NspModelIdentifier   string `json:"nsp-model:identifier"`
								} `json:"@"`
								TargetType  string `json:"target-type"`
								TargetValue string `json:"target-value"`
							} `json:"route-target"`
						} `json:"bgp-instance"`
						AutoBindTunnel struct {
							At struct {
								NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
								NspModelIdentifier   string `json:"nsp-model:identifier"`
							} `json:"@"`
							Resolution                 string      `json:"resolution"`
							EnforceStrictTunnelTagging interface{} `json:"enforce-strict-tunnel-tagging"`
							ResolutionFilter           struct {
								At struct {
									NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
									NspModelIdentifier   string `json:"nsp-model:identifier"`
								} `json:"@"`
								Bgp    interface{} `json:"bgp"`
								Gre    interface{} `json:"gre"`
								Ldp    interface{} `json:"ldp"`
								Rsvp   bool        `json:"rsvp"`
								SrIsis interface{} `json:"sr-isis"`
								SrOspf interface{} `json:"sr-ospf"`
								SrTe   bool        `json:"sr-te"`
							} `json:"resolution-filter"`
						} `json:"auto-bind-tunnel"`
					} `json:"mpls"`
					SapDetails struct {
						At struct {
							NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
							NspModelIdentifier   string `json:"nsp-model:identifier"`
						} `json:"@"`
						Sap []struct {
							At struct {
								NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
								NspModelIdentifier   string `json:"nsp-model:identifier"`
							} `json:"@"`
							PortID       string `json:"port-id"`
							InnerVlanTag int    `json:"inner-vlan-tag"`
							OuterVlanTag int    `json:"outer-vlan-tag"`
							AdminState   string `json:"admin-state"`
							Description  string `json:"description"`
							EnableQos    bool   `json:"enable-qos"`
							EnableFilter bool   `json:"enable-filter"`
							Ingress      struct {
								At struct {
									NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
									NspModelIdentifier   string `json:"nsp-model:identifier"`
								} `json:"@"`
							} `json:"ingress"`
							Egress struct {
								At struct {
									NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
									NspModelIdentifier   string `json:"nsp-model:identifier"`
								} `json:"@"`
							} `json:"egress"`
						} `json:"sap"`
					} `json:"sap-details"`
				} `json:"site-a"`
				SiteB struct {
					At struct {
						NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
						NspModelIdentifier   string `json:"nsp-model:identifier"`
					} `json:"@"`
					DeviceID    string      `json:"device-id"`
					SiteName    interface{} `json:"site-name"`
					Description interface{} `json:"description"`
					Evi         int         `json:"evi"`
					Ecmp        int         `json:"ecmp"`
					LocalAc     struct {
						At struct {
							NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
							NspModelIdentifier   string `json:"nsp-model:identifier"`
						} `json:"@"`
						Name   string `json:"name"`
						EthTag int    `json:"eth-tag"`
					} `json:"local-ac"`
					RemoteAc struct {
						At struct {
							NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
							NspModelIdentifier   string `json:"nsp-model:identifier"`
						} `json:"@"`
						Name   string `json:"name"`
						EthTag int    `json:"eth-tag"`
					} `json:"remote-ac"`
					Mpls struct {
						At struct {
							NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
							NspModelIdentifier   string `json:"nsp-model:identifier"`
						} `json:"@"`
						ForceVcForwarding string `json:"force-vc-forwarding"`
						BgpInstance       struct {
							At struct {
								NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
								NspModelIdentifier   string `json:"nsp-model:identifier"`
							} `json:"@"`
							RouteDistinguisher string `json:"route-distinguisher"`
							RouteTarget        []struct {
								At struct {
									NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
									NspModelIdentifier   string `json:"nsp-model:identifier"`
								} `json:"@"`
								TargetType  string `json:"target-type"`
								TargetValue string `json:"target-value"`
							} `json:"route-target"`
						} `json:"bgp-instance"`
						AutoBindTunnel struct {
							At struct {
								NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
								NspModelIdentifier   string `json:"nsp-model:identifier"`
							} `json:"@"`
							Resolution                 string      `json:"resolution"`
							EnforceStrictTunnelTagging interface{} `json:"enforce-strict-tunnel-tagging"`
							ResolutionFilter           struct {
								At struct {
									NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
									NspModelIdentifier   string `json:"nsp-model:identifier"`
								} `json:"@"`
								Bgp    interface{} `json:"bgp"`
								Gre    interface{} `json:"gre"`
								Ldp    interface{} `json:"ldp"`
								Rsvp   bool        `json:"rsvp"`
								SrIsis interface{} `json:"sr-isis"`
								SrOspf interface{} `json:"sr-ospf"`
								SrTe   bool        `json:"sr-te"`
							} `json:"resolution-filter"`
						} `json:"auto-bind-tunnel"`
					} `json:"mpls"`
					SapDetails struct {
						At struct {
							NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
							NspModelIdentifier   string `json:"nsp-model:identifier"`
						} `json:"@"`
						Sap []struct {
							At struct {
								NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
								NspModelIdentifier   string `json:"nsp-model:identifier"`
							} `json:"@"`
							PortID       string `json:"port-id"`
							InnerVlanTag int    `json:"inner-vlan-tag"`
							OuterVlanTag int    `json:"outer-vlan-tag"`
							AdminState   string `json:"admin-state"`
							Description  string `json:"description"`
							EnableQos    bool   `json:"enable-qos"`
							EnableFilter bool   `json:"enable-filter"`
							Ingress      struct {
								At struct {
									NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
									NspModelIdentifier   string `json:"nsp-model:identifier"`
								} `json:"@"`
							} `json:"ingress"`
							Egress struct {
								At struct {
									NspModelSchemaNodeid string `json:"nsp-model:schema-nodeid"`
									NspModelIdentifier   string `json:"nsp-model:identifier"`
								} `json:"@"`
							} `json:"egress"`
						} `json:"sap"`
					} `json:"sap-details"`
				} `json:"site-b"`
			} `json:"evpn-epipe:evpn-epipe"`
		} `json:"intent-specific-data"`
	} `json:"nsp-service-intent:intent"`
}
