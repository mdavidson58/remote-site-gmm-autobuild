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
	
Org :="MyNewTestOrg" + strconv.Itoa(x)

// create the new org 
  new_org_id := gmmapi.Create_gmm_org(gmm_api_key, Org, my_org_id)
  fmt.Println ("org " + strconv.Itoa(new_org_id) + " created")

// Upload a Flexible Template to GMM
// substitute flex.json with the flexible template
gmmapi.Gmm_upload_flex_template(gmm_api_key,new_org_id,"flex.json")

//Get the template id
flex_id := gmmapi.Retrieve_gmm_flex_template_id(gmm_api_key, new_org_id, "829 Telemedicine Advanced Template") 

//replace flexible template id with actual flexbile template id
gw_profile := `{"name":"829 Telemedicine dual VPN","wifi_enable":true,"wifi_ssid_prefix":"","lan_enable":true,"wan_enable":true,"organization_id":2835,"automatic_wifi_config":false,"wifi_ssid":"829-Z0L1","wifi_pre_shared_key":"Cisco,123","asset_names":[],"wan_interface_type":"No Change","apn":null,"secondary_apn":null,"gps_enable":false,"wgb_enable":false,"wgb_ssid":null,"wgb_pre_shared_key":null,"wgb_subnet":null,"customer_vpn_enable":true,"customer_vpn_server_ip":"64.102.254.149","customer_vpn_pre_shared_key":"cisco123","private_subnet_enable":false,"private_subnet":null,"private_subnet_dns_ip":null,"private_subnet_exclusion_range_start":null,"private_subnet_exclusion_range_end":null,"default":false,"port_forwarding_enable":false,"private_subnet_advanced":false,"private_subnet_gateway_ip":null,"private_subnet_netmask":null,"ip_dhcp_helper":null,"broadcast_ssid_enable":true,"vrf_enable":false,"model":"IR829","backup_customer_vpn_server_ip":"64.102.254.153","backup_customer_vpn_pre_shared_key":"cisco123","flexible_template_id":XX,"flexible_template_enable":true,"flexible_template_advanced":false,"flexible_template_variables":[],"advanced_wifi_config":false,"unified_ap_enable":false,"wlc_primary_ip":"","wlc_secondary_ip":"","dot1x_enable":false,"primary_aaa_server_ip":null,"secondary_aaa_server_ip":null,"primary_aaa_server_key":null,"secondary_aaa_server_key":null,"lan_ports":["gi1-\u003etrue","gi2-\u003etrue","gi3-\u003etrue","gi4-\u003etrue"],"recovery_time":120,"dual_lte_active_enable":false,"dual_lte_active_load":null,"dual_lte_active_subnet":null,"ip_connectivity_check_sim_0":null,"ip_connectivity_check_sim_1":null,"expansion_module":"f","pluggable_module":null,"pluggable_interface_module":null,"ap_flexible_template_id":null,"ap_flexible_template_enable":false,"ap_flexible_template_advanced":false,"ap_flexible_template_variables":[],"gate_ways_count":0,"port_forwards":[]}`
gw_profile = strings.Replace(gw_profile, "XX", strconv.Itoa(flex_id),1)

// Upload a Gateway Profile to GMM
gmmapi.Gmm_upload_gwy_profile(gmm_api_key,new_org_id,gw_profile)

// get profile id for claim policy
profile_id := gmmapi.Retrieve_gmm_profile_id(gmm_api_key, new_org_id, "829 Telemedicine dual VPN")

// retrieve the id of the first tag
new_tag_id := gmmapi.Retrieve_gmm_org_tag_id(gmm_api_key, new_org_id, 0)

// this is our tag
tag_string := `{
  "custom_name": "Tag1",
  "default_name": "account 1",
  "field_type": "Dropdown",
  "enabled": true,
  "allowed_values": [
      "Yes",
      "No"
    ]

}`

// set the tag
 gmmapi.Update_gmm_org_tags(gmm_api_key, new_tag_id, tag_string)

// sample claim policy
 claim_policy := `{
    
      "name": "My new claim policy",
      "activated": false,
      "tag_id": XX,
      "policy_templates": [
        {
          "value": "Yes",
          "template_ids": [YY]
    
        }
      
     ]
   }`

// replcae tag_id with actual tagid
 claim_policy = strings.Replace(claim_policy, "XX", strconv.Itoa(new_tag_id),1)
 // replace template_id with actual profile
 claim_policy = strings.Replace(claim_policy, "YY", strconv.Itoa(profile_id),1)
 	
fmt.Println ("claim policy " + strconv.Itoa(gmmapi.Create_claim_policy(gmm_api_key, new_org_id, claim_policy)) + " created")


}
}