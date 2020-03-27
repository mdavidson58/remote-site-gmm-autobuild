package main

import (
"fmt"
"os"
"strconv"
	"gmmcli/gmmapi"
	
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

gmmapi.Retrieve_gmm_flex_template_list(gmm_api_key, my_org_id) 
gmmapi.Retrieve_gmm_gwy_profiles_list(gmm_api_key, my_org_id)

gmmapi.Retrieve_gmm_gwy_profile(gmm_api_key, my_org_id, "829 Telemedicine")
gmmapi.Retrieve_gmm_gwy_profile(gmm_api_key, my_org_id, "829 Telemedicine dual VPN")
gmmapi.Retrieve_gmm_flex_template(gmm_api_key, my_org_id, "829 Telemedicine Advanced Template")

// Upload a Gateway Profile to GMM
// substitute profile.json with the GMM json file
//gmm_upload_gwy_profile(gmm_api_key,my_org_id,"profile.json")

// Upload a Flexible Template to GMM
// substitute flex.json with the flexible template
//gmm_upload_flex_template(gmm_api_key,my_org_id,"flex.json")


}