package main

import (
"fmt"
"os"
"strconv"
	"gmmcli/gmmapi"
"strings"
	
)

func main() {

  if len(os.Args) != 4 {
        fmt.Println("Usage:", os.Args[0], "\"gmmusername\"", "\"gmmpassword\"", "GMMOrgID")
        return
    }


// Username and pass  
  user_name := os.Args[1] 
  user_pass := os.Args[2] 


// substitute 1234 with your org id
my_org_id, err := strconv.Atoi(os.Args[3]) 
if err != nil {
os.Exit(1)
}

// Retrieve GMM API Key
// substitue  uid@enterprise.com and password with your GMM account and password that has API access
gmm_api_key := gmmapi.Retrieve_gmm_api_key(user_name, user_pass)

//gmmapi.Retrieve_gmm_flex_template_list(gmm_api_key, my_org_id) 
//gmmapi.Retrieve_gmm_gwy_profiles_list(gmm_api_key, my_org_id)
//gmmapi.Retrieve_gmm_gwy_profile(gmm_api_key, my_org_id, "829 Telemedicine")
//gmmapi.Retrieve_gmm_gwy_profile(gmm_api_key, my_org_id, "829 Telemedicine dual VPN")
//gmmapi.Retrieve_gmm_flex_template(gmm_api_key, my_org_id, "829 Telemedicine Advanced Template")

//fmt.Println("tag 1 = " + strconv.Itoa(gmmapi.Retrieve_gmm_org_tag_id(gmm_api_key, my_org_id, 0)))

// let's build orgs
for x :=0; x< 2; x++ {
	
Org :="TeleMedicine" + strconv.Itoa(x)

// create the new org 
  new_org_id := gmmapi.Create_gmm_org(gmm_api_key, Org, my_org_id)
  fmt.Println ("org " + strconv.Itoa(new_org_id) + " created")

// Upload a Flexible Template to GMM
// substitute flex.json with the flexible template
dual_radio_flex_template := `
{
     "name": "829 Telemedicine - Dual LTE Adv Template",
      "description": "",
      "template": "ip dhcp pool subtended\n option 43 hex ${custom.WLC_IP_HexFormat}\n!\ncrypto ikev2 authorization policy CVPN \n route set interface\n route accept any distance 70\n!\n\ncrypto ikev2 keyring Flex_key\n !\n peer ${custom.VPN_Headend_IP_1}\n  address ${custom.VPN_Headend_IP_1}\n  identity key-id ${custom.VPN_Headend_IP_1}\n  pre-shared-key ${custom.VPN_1_password}\n !\n peer ${custom.VPN_Headend_IP_2}\n  address ${custom.VPN_Headend_IP_2}\n  identity key-id ${custom.VPN_Headend_IP_2}\n  pre-shared-key ${custom.VPN_2_password}\n !\n!\ncrypto ipsec profile CVPN_IPS_PF\n set ikev2-profile CVPN_I2PF\n!\ncrypto ikev2 profile CVPN_I2PF\n match identity remote key-id ${custom.VPN_Headend_IP_1}\n match identity remote key-id ${custom.VPN_Headend_IP_2}\n identity local email ${gw.sn}@iotspdev.io\n authentication remote pre-share\n authentication local pre-share\n keyring local Flex_key\n dpd 29 2 periodic\n aaa authorization group psk list CVPN CVPN\n!\n\n!\ninterface Tunnel2\n ip address negotiated\n ip mtu 1358\n ip nat outside\n ip virtual-reassembly in\n ip tcp adjust-mss 1318\n tunnel source dynamic\n tunnel mode ipsec ipv4\n tunnel destination dynamic\n tunnel path-mtu-discovery\n tunnel protection ipsec profile CVPN_IPS_PF\n!\n\n!\ninterface Tunnel3\n ip address negotiated\n ip mtu 1358\n ip nat outside\n ip virtual-reassembly in\n ip tcp adjust-mss 1318\n tunnel source dynamic\n tunnel mode ipsec ipv4\n tunnel destination dynamic\n tunnel path-mtu-discovery\n tunnel protection ipsec profile CVPN_IPS_PF\n!\n!\n\ncrypto ikev2 client flexvpn Tunnel2\n  peer 1 ${custom.VPN_Headend_IP_1}\n  source 1 Cellular0/0 track 3\n  client connect Tunnel2\n!\ncrypto ikev2 client flexvpn Tunnel3\n  peer 1 ${custom.VPN_Headend_IP_2}\n  source 1 Cellular1/0 track 2\n  client connect Tunnel3\n!\n!\n!\ntrack 3 interface Cellular0/0 ip routing\n!\ntrack 2 interface Cellular1/0 ip routing\n!\n\nip nat inside source route-map RM_Tu2 interface Tunnel2 overload\nip nat inside source route-map RM_Tu3 interface Tunnel3 overload\nip nat inside source route-map RM_WAN2_ACL interface Cellular0/0 overload\nip nat inside source route-map RM_WAN_ACL interface Cellular1/0 overload\nip route 0.0.0.0 0.0.0.0 Cellular1/0 99 track 10\nip route ${custom.VPN_Headend_IP_2} 255.255.255.255 Cellular1/0 track 20\nip route 0.0.0.0 0.0.0.0 Cellular0/0 99 track 20\nip route ${custom.VPN_Headend_IP_1} 255.255.255.255 Cellular0/0 track 20\n!\n\nip access-list extended filter-internet\n permit esp host ${custom.VPN_Headend_IP_1} any\n permit esp host ${custom.VPN_Headend_IP_2} any\n!\nroute-map RM_WGB_ACL permit 10\n match ip address NAT_ACL\n match interface Vlan50\n!\nroute-map RM_WAN_ACL permit 10\n match ip address NAT_ACL\n match interface Cellular1/0\n!\nroute-map RM_WAN2_ACL permit 10\n match ip address NAT_ACL\n match interface Cellular0/0\n!\nroute-map RM_Tu2 permit 10\n match ip address NAT_ACL\n match interface Tunnel2\n!\nroute-map RM_Tu3 permit 10\n match ip address NAT_ACL\n match interface Tunnel3\n!\n!\nclass-map match-any CLASS_LOW_PRIORITY\n match ip precedence 1  2\nclass-map match-any CLASS_MED_PRIORITY\n match ip precedence 3 4\n match application webex-meeting\n match application ip-camera\n match application telepresence-media\n match application sip\nclass-map match-any CLASS_HIGH_PRIORITY\n match ip precedence 5 \n match application rtp\n match dscp ef \n\n!\npolicy-map LTE_QOS_PMAP\n class class-default\n  shape average 5000000\npolicy-map LTE_QOS_PMAP_INNER\n class CLASS_HIGH_PRIORITY\n  priority 400\n class CLASS_MED_PRIORITY\n  bandwidth 1000\n class CLASS_LOW_PRIORITY\n  bandwidth 1000\n class class-default\n  fair-queue\n  random-detect dscp-based\n!\ninterface Tunnel2\n bandwidth 5000\n service-policy output LTE_QOS_PMAP\n!\ninterface Tunnel3\n bandwidth 5000\n service-policy output LTE_QOS_PMAP",
      "variables": [
        "WLC_IP_HexFormat",
        "VPN_Headend_IP_1",
        "VPN_1_password",
        "VPN_Headend_IP_2",
        "VPN_2_password"
      ],
      "flexible_template_type": "router"
}`


single_radio_flex_template := `
{
      "name": "829 Telemedicine - Single LTE Adv Template",
      "description": "",
      "template": "ip dhcp pool subtended\n option 43 hex ${custom.WLC_IP_HexFormat}\n!\n!\ncrypto ikev2 keyring Flex_key\n !\n peer ${custom.VPN_Headend_IP_1}\n  address ${custom.VPN_Headend_IP_1}\n  identity key-id ${custom.VPN_Headend_IP_1}\n  pre-shared-key ${custom.VPN_1_password}\n !\n!\ncrypto ikev2 profile CVPN_I2PF\n match identity remote key-id ${custom.VPN_Headend_IP_1}\n!\n\n!\n\n!\n!\ncrypto ikev2 client flexvpn Tunnel2\n  peer 1 ${custom.VPN_Headend_IP_1}\n!\n!\nip route ${custom.VPN_Headend_IP_1} 255.255.255.255 ${gw.wan_if} track 3\n!\n\nip access-list extended filter-internet\n permit esp host ${custom.VPN_Headend_IP_1} any\n!\n\nclass-map match-any CLASS_LOW_PRIORITY\n match ip precedence 1  2\nclass-map match-any CLASS_MED_PRIORITY\n match ip precedence 3 4\n match application webex-meeting\n match application ip-camera\n match application telepresence-media\n match application sip\nclass-map match-any CLASS_HIGH_PRIORITY\n match ip precedence 5 \n match application rtp\n match dscp ef \n!\npolicy-map LTE_QOS_PMAP\n class class-default\n  shape average 5000000\npolicy-map LTE_QOS_PMAP_INNER\n class CLASS_HIGH_PRIORITY\n  priority 400\n class CLASS_MED_PRIORITY\n  bandwidth 1000\n class CLASS_LOW_PRIORITY\n  bandwidth 1000\n class class-default\n  fair-queue\n  random-detect dscp-based\n!\ninterface Tunnel2\n bandwidth 5000\n service-policy output LTE_QOS_PMAP\n!",
      "variables": [
        "WLC_IP_HexFormat",
        "VPN_Headend_IP_1",
        "VPN_1_password"
      ],
      "flexible_template_type": "router"
}`

gmmapi.Gmm_upload_flex_template(gmm_api_key,new_org_id,dual_radio_flex_template)
gmmapi.Gmm_upload_flex_template(gmm_api_key,new_org_id,single_radio_flex_template)

//Get the template id
dual_flex_id := gmmapi.Retrieve_gmm_flex_template_id(gmm_api_key, new_org_id, "829 Telemedicine - Dual LTE Adv Template") 
single_flex_id := gmmapi.Retrieve_gmm_flex_template_id(gmm_api_key, new_org_id, "829 Telemedicine - Single LTE Adv Template")

//replace flexible template id with actual flexbile template id
dual_radio_profile := ` 
{
      "name": "829 Telemedicine dual Cell and VPN",
      "wifi_enable": true,
      "wifi_ssid_prefix": "",
      "lan_enable": true,
      "wan_enable": true,
      "automatic_wifi_config": true,
      "wifi_ssid": "",
      "wifi_pre_shared_key": "",
      "asset_names": [],
      "wan_interface_type": "No Change",
      "apn": null,
      "secondary_apn": null,
      "gps_enable": true,
      "wgb_enable": false,
      "wgb_ssid": null,
      "wgb_pre_shared_key": null,
      "wgb_subnet": null,
      "customer_vpn_enable": false,
      "customer_vpn_server_ip": null,
      "customer_vpn_pre_shared_key": null,
      "private_subnet_enable": true,
      "private_subnet": "192.168.10.1/24",
      "private_subnet_dns_ip": "8.8.8.8",
      "private_subnet_exclusion_range_start": "192.168.10.1",
      "private_subnet_exclusion_range_end": "192.168.10.9",
      "default": false,
      "port_forwarding_enable": false,
      "private_subnet_advanced": false,
      "private_subnet_gateway_ip": "192.168.10.1",
      "private_subnet_netmask": "255.255.255.0",
      "ip_dhcp_helper": null,
      "broadcast_ssid_enable": true,
      "vrf_enable": false,
      "model": "IR829",
      "backup_customer_vpn_server_ip": null,
      "backup_customer_vpn_pre_shared_key": null,
      "flexible_template_id": XX,
      "flexible_template_enable": true,
      "flexible_template_advanced": true,
      "flexible_template_variables": [
        {
          "name": "WLC_IP_HexFormat",
          "value": "f104.0a02.0501"
        },
        {
          "name": "VPN_Headend_IP_1",
          "value": "64.102.254.149"
        },
        {
          "name": "VPN_1_password",
          "value": "cisco123"
        },
        {
          "name": "VPN_Headend_IP_2",
          "value": "64.102.254.153"
        },
        {
          "name": "VPN_2_password",
          "value": "cisco123"
        }
      ],
      "advanced_wifi_config": false,
      "unified_ap_enable": true,
      "wlc_primary_ip": "127.0.0.1",
      "wlc_secondary_ip": "",
      "dot1x_enable": false,
      "primary_aaa_server_ip": null,
      "secondary_aaa_server_ip": null,
      "primary_aaa_server_key": null,
      "secondary_aaa_server_key": null,
      "lan_ports": [
        "gi1->true",
        "gi2->true",
        "gi3->true",
        "gi4->true"
      ],
      "recovery_time": 120,
      "dual_lte_active_enable": true,
      "dual_lte_active_load": "equal",
      "dual_lte_active_subnet": null,
      "ip_connectivity_check_sim_0": "8.8.4.4",
      "ip_connectivity_check_sim_1": "9.9.9.9",
      "expansion_module": "f",
      "pluggable_module": null,
      "pluggable_interface_module": null,
      "ap_flexible_template_id": null,
      "ap_flexible_template_enable": false,
      "ap_flexible_template_advanced": false,
      "ap_flexible_template_variables": [],
      "gate_ways_count": 1,
      "port_forwards": []
}`
dual_radio_profile = strings.Replace(dual_radio_profile, "XX", strconv.Itoa(dual_flex_id),1)

single_radio_profile :=` 
{
      "name": "829 Telemedicine single Cell and VPN",
      "wifi_enable": false,
      "wifi_ssid_prefix": null,
      "lan_enable": true,
      "wan_enable": true,
      "automatic_wifi_config": true,
      "wifi_ssid": null,
      "wifi_pre_shared_key": null,
      "asset_names": [],
      "wan_interface_type": "No Change",
      "apn": null,
      "secondary_apn": null,
      "gps_enable": true,
      "wgb_enable": false,
      "wgb_ssid": null,
      "wgb_pre_shared_key": null,
      "wgb_subnet": null,
      "customer_vpn_enable": true,
      "customer_vpn_server_ip": "192.168.254.254",
      "customer_vpn_pre_shared_key": "REPLACEME",
      "private_subnet_enable": true,
      "private_subnet": "192.168.0.1/24",
      "private_subnet_dns_ip": "8.8.8.8",
      "private_subnet_exclusion_range_start": "",
      "private_subnet_exclusion_range_end": "",
      "default": false,
      "port_forwarding_enable": false,
      "private_subnet_advanced": false,
      "private_subnet_gateway_ip": "192.168.0.1",
      "private_subnet_netmask": "255.255.255.0",
      "ip_dhcp_helper": null,
      "broadcast_ssid_enable": true,
      "vrf_enable": false,
      "model": "IR829",
      "backup_customer_vpn_server_ip": "",
      "backup_customer_vpn_pre_shared_key": "",
      "flexible_template_id": XX,
      "flexible_template_enable": true,
      "flexible_template_advanced": true,
      "flexible_template_variables": [
        {
          "name": "WLC_IP_HexFormat",
          "value": "f104.0a02.0501"
        },
        {
          "name": "VPN_Headend_IP_1",
          "value": "64.102.254.149"
        },
        {
          "name": "VPN_1_password",
          "value": "cisco123"
        }
      ],
      "advanced_wifi_config": false,
      "unified_ap_enable": true,
      "wlc_primary_ip": "192.168.254.254",
      "wlc_secondary_ip": "",
      "dot1x_enable": false,
      "primary_aaa_server_ip": null,
      "secondary_aaa_server_ip": null,
      "primary_aaa_server_key": null,
      "secondary_aaa_server_key": null,
      "lan_ports": [
        "gi1->true",
        "gi2->true",
        "gi3->true",
        "gi4->true"
      ],
      "recovery_time": 120,
      "dual_lte_active_enable": false,
      "dual_lte_active_load": null,
      "dual_lte_active_subnet": null,
      "ip_connectivity_check_sim_0": null,
      "ip_connectivity_check_sim_1": null,
      "expansion_module": "f",
      "pluggable_module": null,
      "pluggable_interface_module": null,
      "ap_flexible_template_id": null,
      "ap_flexible_template_enable": false,
      "ap_flexible_template_advanced": false,
      "ap_flexible_template_variables": [],
      "gate_ways_count": 2,
      "port_forwards": []
}`
single_radio_profile = strings.Replace(single_radio_profile, "XX", strconv.Itoa(single_flex_id),1)

// Upload a Gateway Profile to GMM
gmmapi.Gmm_upload_gwy_profile(gmm_api_key,new_org_id,dual_radio_profile)
gmmapi.Gmm_upload_gwy_profile(gmm_api_key,new_org_id,single_radio_profile)

// get profile id for claim policy
dual_profile_id := gmmapi.Retrieve_gmm_profile_id(gmm_api_key, new_org_id, "829 Telemedicine dual Cell and VPN")
single_profile_id := gmmapi.Retrieve_gmm_profile_id(gmm_api_key, new_org_id, "829 Telemedicine single Cell and VPN")

// retrieve the id of the first tag
new_tag_id := gmmapi.Retrieve_gmm_org_tag_id(gmm_api_key, new_org_id, 0)

// this is our tag
tag_string := `  
{
    "default_name": "account 1",
    "custom_name": "Gateway Type",
    "enabled": true,
    "field_type": "Dropdown",
    "allowed_values": [
      "Telemedicine with Dual Radios",
      "Telemedicine with Single Radio"
    ]
 }`

// set the tag
 gmmapi.Update_gmm_org_tags(gmm_api_key, new_tag_id, tag_string)

// sample claim policy
 claim_policy := `
 {
      "name": "Telemedicine",
      "activated": true,
      "tag_id": XX,
      "policy_templates": [
        {
          "value": "Telemedicine with Dual Radios",
          "template_ids": [
            YY
          ]
        },
        {
          "value": "Telemedicine with Single Radio",
          "template_ids": [
            ZZ
          ]
        }
      ] 
 }`

// replcae tag_id with actual tagid
 claim_policy = strings.Replace(claim_policy, "XX", strconv.Itoa(new_tag_id),1)
 // replace template_id with actual profile
 claim_policy = strings.Replace(claim_policy, "YY", strconv.Itoa(dual_profile_id),1)
 claim_policy = strings.Replace(claim_policy, "ZZ", strconv.Itoa(single_profile_id),1)

 	
fmt.Println ("claim policy " + strconv.Itoa(gmmapi.Create_claim_policy(gmm_api_key, new_org_id, claim_policy)) + " created")


}
}