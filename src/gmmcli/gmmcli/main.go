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
	
Org :="RemoteSite" + strconv.Itoa(x)

// create the new org 
  new_org_id := gmmapi.Create_gmm_org(gmm_api_key, Org, my_org_id)
  fmt.Println ("org " + strconv.Itoa(new_org_id) + " created")

// Upload a Flexible Template to GMM
// substitute flex.json with the flexible template
flex_template_807 := `
{
      "name": "807 Remote Access Adv Template",
      "description": "",
      "template": "!\n/* Customer must configure custom subnet. Example: Set Custom subnet to 192.168.3.0/24\n/* Note do not use gmm_ip_01, it is the IP address of the router\n\n<#assign gmm_subnet = gw.ip_prefix + \".\" + gw.ip_suffix>\n<#assign gmm_ip_01 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 16)>\n<#assign gmm_ip_02 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 17)>\n<#assign gmm_ip_03 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 18)>\n<#assign gmm_ip_04 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 19)>\n<#assign gmm_ip_05 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 20)>\n<#assign gmm_ip_06 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 21)>\n<#assign gmm_ip_07 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 22)>\n<#assign gmm_ip_08 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 23)>\n<#assign gmm_ip_09 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 24)>\n<#assign gmm_ip_10 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 25)>\n<#assign gmm_ip_11 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 26)>\n<#assign gmm_ip_12 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 27)>\n<#assign gmm_ip_13 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 28)>\n<#assign gmm_ip_14 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 29)>\n\ninterface FastEthernet1\nip address ${ custom.LAN_interface_address } ${gw.lan_netmask}\n\n/* Define IP addresse in custom subnet\n/*\n\n/* 32 ip total\n/* 1st 8 ip reserved for IOX network\n/* 2nd 8 ip reserved for AP network\n/* 2nd half of 32 ip(16 ip) reserved for device subnet\n\n/* assign Loopback0 to device subnet block (since not used because custom subnet is configured)\ninterface Loopback0\nip address ${ gmm_ip_01 } 255.255.255.240\nip nat outside\n\n/* any local IP from device subnet block going to ASA\nip access-list extended nat-on-a-stick\npermit ip ${ gmm_subnet } 0.0.0.15 10.7.0.0 0.0.255.255\nip access-list extended remote-access\npermit ip ${ gmm_subnet } 0.0.0.15 any\nip access-list extended remote-asa\npermit ip 10.7.0.0 0.0.255.255 any\nroute-map nat-on-a-stick permit 10\nmatch ip address nat-on-a-stick\nset interface Loopback0\n\n/* configure block for number of simultaneous connections down to the device subnet from ASA\nip nat pool remote-access ${ custom.nat_inside_address } ${ custom.nat_inside_address} netmask ${ gw.lan_netmask }\n\n/* add additional NAT statement as needed based on number of devices (up to 13 devices)\nip nat inside source static ${ custom.device_ip_1 } ${ gmm_ip_02 } route-map remote-asa\n\nip nat outside source list remote-asa pool remote-access add-route\n\nint FastEthernet1\nip policy route-map nat-on-a-stick\n\nint Tunnel1\nip nat outside",
      "variables": [
        "LAN_interface_address",
        "nat_inside_address",
        "device_ip_1"
      ],
      "flexible_template_type": "router"

}
`


flex_template_1101 := `
{
      "name": "1101 Remote Access Adv Template",
      "description": "",
      "template": "/* Customer must configure custom subnet. Example: Set Custom subnet to 192.168.3.0/24\n/* Note do not use gmm_ip_01, it is the IP address of the router\n\n<#assign gmm_subnet = gw.ip_prefix + \".\" + gw.ip_suffix>\n<#assign gmm_ip_01 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 16)>\n<#assign gmm_ip_02 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 17)>\n<#assign gmm_ip_03 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 18)>\n<#assign gmm_ip_04 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 19)>\n<#assign gmm_ip_05 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 20)>\n<#assign gmm_ip_06 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 21)>\n<#assign gmm_ip_07 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 22)>\n<#assign gmm_ip_08 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 23)>\n<#assign gmm_ip_09 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 24)>\n<#assign gmm_ip_10 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 25)>\n<#assign gmm_ip_11 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 26)>\n<#assign gmm_ip_12 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 27)>\n<#assign gmm_ip_13 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 28)>\n<#assign gmm_ip_14 = gw.ip_prefix + \".\" + (gw.ip_suffix?number + 29)>\n\ninterface Vlan1\nip address ${ custom.LAN_interface_address } ${gw.lan_netmask}\n\n/* Define IP addresse in custom subnet\n/*\n\n/* 32 ip total\n/* 1st 8 ip reserved for IOX network\n/* 2nd 8 ip reserved for AP network\n/* 2nd half of 32 ip(16 ip) reserved for device subnet\n\n/* assign Loopback0 to device subnet block (since not used because custom subnet is configured)\ninterface Loopback0\nip address ${ gmm_ip_01 } 255.255.255.240\nip nat outside\n\n/* any local IP from device subnet block going to ASA\nip access-list extended nat-on-a-stick\npermit ip ${ gmm_subnet } 0.0.0.15 10.7.0.0 0.0.255.255\nip access-list extended remote-access\npermit ip ${ gmm_subnet } 0.0.0.15 any\nip access-list extended remote-asa\npermit ip 10.7.0.0 0.0.255.255 any\nroute-map nat-on-a-stick permit 10\nmatch ip address nat-on-a-stick\nset interface Loopback0\n\n/* configure block for number of simultaneous connections down to the device subnet from ASA\nip nat pool remote-access ${ custom.nat_inside_address } ${ custom.nat_inside_address} netmask ${ gw.lan_netmask }\n\n/* add additional NAT statement as needed based on number of devices (up to 13 devices)\nip nat inside source static ${ custom.device_ip_1 } ${ gmm_ip_02 } route-map remote-asa\n\nip nat outside source list remote-asa pool remote-access add-route\n\nint Vlan1\nip policy route-map nat-on-a-stick\n\nint Tunnel1\nip nat outside\n\n!\ninterface Async0/2/0\nno ip address\nencapsulation relay-line\n!\nline 0/2/0\ntransport input telnet\ntransport output none\nstopbits 1\n!",
      "variables": [
        "LAN_interface_address",
        "nat_inside_address",
        "device_ip_1"
      ],
      "flexible_template_type": "router"

}
`

gmmapi.Gmm_upload_flex_template(gmm_api_key,new_org_id,flex_template_807)
gmmapi.Gmm_upload_flex_template(gmm_api_key,new_org_id,flex_template_1101)

//Get the template id
flex_id_807 := gmmapi.Retrieve_gmm_flex_template_id(gmm_api_key, new_org_id, "829 Telemedicine - Dual LTE Adv Template") 
flex_id_1101 := gmmapi.Retrieve_gmm_flex_template_id(gmm_api_key, new_org_id, "829 Telemedicine - Single LTE Adv Template")

//replace flexible template id with actual flexbile template id
profile_807 := `
{
      "name": "807 Remote Access Template",
      "wifi_enable": false,
      "wifi_ssid_prefix": null,
      "lan_enable": true,
      "wan_enable": true,
      "organization_id": 2835,
      "automatic_wifi_config": true,
      "wifi_ssid": null,
      "wifi_pre_shared_key": null,
      "asset_names": [],
      "wan_interface_type": "No Change",
      "apn": null,
      "secondary_apn": null,
      "gps_enable": false,
      "wgb_enable": false,
      "wgb_ssid": null,
      "wgb_pre_shared_key": null,
      "wgb_subnet": null,
      "customer_vpn_enable": false,
      "customer_vpn_server_ip": null,
      "customer_vpn_pre_shared_key": null,
      "private_subnet_enable": true,
      "private_subnet": "192.168.3.1/24",
      "private_subnet_dns_ip": "",
      "private_subnet_exclusion_range_start": "192.168.3.1",
      "private_subnet_exclusion_range_end": "192.168.3.254",
      "default": false,
      "port_forwarding_enable": false,
      "private_subnet_advanced": false,
      "private_subnet_gateway_ip": "192.168.3.1",
      "private_subnet_netmask": "255.255.255.0",
      "ip_dhcp_helper": null,
      "broadcast_ssid_enable": true,
      "vrf_enable": false,
      "model": "IR807",
      "backup_customer_vpn_server_ip": null,
      "backup_customer_vpn_pre_shared_key": null,
      "flexible_template_id": XX,
      "flexible_template_enable": true,
      "flexible_template_advanced": true,
      "flexible_template_variables": [],
      "advanced_wifi_config": false,
      "unified_ap_enable": false,
      "wlc_primary_ip": null,
      "wlc_secondary_ip": null,
      "dot1x_enable": false,
      "primary_aaa_server_ip": null,
      "secondary_aaa_server_ip": null,
      "primary_aaa_server_key": null,
      "secondary_aaa_server_key": null,
      "lan_ports": [
        "FastEthernet1->true"
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
      "gate_ways_count": 1,
      "port_forwards": []
}`

profile_807 = strings.Replace(profile_807, "XX", strconv.Itoa(flex_id_807),1)

profile_1101 :=` 
{
      "name": "1101 Remote Access for Manufacturing",
      "wifi_enable": false,
      "wifi_ssid_prefix": null,
      "lan_enable": true,
      "wan_enable": true,
      "organization_id": 2835,
      "automatic_wifi_config": true,
      "wifi_ssid": null,
      "wifi_pre_shared_key": null,
      "asset_names": [],
      "wan_interface_type": "No Change",
      "apn": null,
      "secondary_apn": null,
      "gps_enable": false,
      "wgb_enable": false,
      "wgb_ssid": null,
      "wgb_pre_shared_key": null,
      "wgb_subnet": null,
      "customer_vpn_enable": false,
      "customer_vpn_server_ip": null,
      "customer_vpn_pre_shared_key": null,
      "private_subnet_enable": true,
      "private_subnet": "192.168.3.1/24",
      "private_subnet_dns_ip": "",
      "private_subnet_exclusion_range_start": "192.168.3.1",
      "private_subnet_exclusion_range_end": "192.168.3.254",
      "default": false,
      "port_forwarding_enable": false,
      "private_subnet_advanced": false,
      "private_subnet_gateway_ip": "192.168.3.1",
      "private_subnet_netmask": "255.255.255.0",
      "ip_dhcp_helper": null,
      "broadcast_ssid_enable": true,
      "vrf_enable": false,
      "model": "IR1101",
      "backup_customer_vpn_server_ip": null,
      "backup_customer_vpn_pre_shared_key": null,
      "flexible_template_id": XX,
      "flexible_template_enable": true,
      "flexible_template_advanced": true,
      "flexible_template_variables": [],
      "advanced_wifi_config": false,
      "unified_ap_enable": false,
      "wlc_primary_ip": null,
      "wlc_secondary_ip": null,
      "dot1x_enable": false,
      "primary_aaa_server_ip": null,
      "secondary_aaa_server_ip": null,
      "primary_aaa_server_key": null,
      "secondary_aaa_server_key": null,
      "lan_ports": [
        "FastEthernet0/0/1->true",
        "FastEthernet0/0/2->true",
        "FastEthernet0/0/3->true",
        "FastEthernet0/0/4->true"
      ],
      "recovery_time": 120,
      "dual_lte_active_enable": false,
      "dual_lte_active_load": null,
      "dual_lte_active_subnet": null,
      "ip_connectivity_check_sim_0": null,
      "ip_connectivity_check_sim_1": null,
      "expansion_module": "",
      "pluggable_module": "",
      "pluggable_interface_module": "P-LTE",
      "ap_flexible_template_id": null,
      "ap_flexible_template_enable": false,
      "ap_flexible_template_advanced": false,
      "ap_flexible_template_variables": [],
      "gate_ways_count": 1,
      "port_forwards": []
}
`
profile_1101 = strings.Replace(profile_1101, "XX", strconv.Itoa(flex_id_1101),1)

// Upload a Gateway Profile to GMM
gmmapi.Gmm_upload_gwy_profile(gmm_api_key,new_org_id,profile_807)
gmmapi.Gmm_upload_gwy_profile(gmm_api_key,new_org_id,profile_1101)

// get profile id for claim policy
profile_id_807 := gmmapi.Retrieve_gmm_profile_id(gmm_api_key, new_org_id, "807 Remote Access Template")
profile_id_1101 := gmmapi.Retrieve_gmm_profile_id(gmm_api_key, new_org_id, "1101 Remote Access for Manufacturing")

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
      "IR807",
      "IR1101"
    ]
 }`

// set the tag
 gmmapi.Update_gmm_org_tags(gmm_api_key, new_tag_id, tag_string)

// sample claim policy
 claim_policy := `
 {
      "name": "Remote Access",
      "activated": true,
      "tag_id": XX,
      "policy_templates": [
        {
          "value": "Remote Access with IR807",
          "template_ids": [
            YY
          ]
        },
        {
          "value": "Remote Access with IR1101",
          "template_ids": [
            ZZ
          ]
        }
      ] 
 }`

// replcae tag_id with actual tagid
 claim_policy = strings.Replace(claim_policy, "XX", strconv.Itoa(new_tag_id),1)
 // replace template_id with actual profile
 claim_policy = strings.Replace(claim_policy, "YY", strconv.Itoa(profile_id_807),1)
 claim_policy = strings.Replace(claim_policy, "ZZ", strconv.Itoa(profile_id_1101),1)

 	
fmt.Println ("claim policy " + strconv.Itoa(gmmapi.Create_claim_policy(gmm_api_key, new_org_id, claim_policy)) + " created")


}
}