/*
This package provides sample Go functions to achieve the following functionality:
'95	Retrieve GMM API Key
'95	Retrieve GMM Gateway Health Status Summary
'95	List Gateway Profiles in GMM
'95	List Flexible Templates in GMM
'95	Download Gateway Profile from GMM
'95	Upload Gateway Profile to GMM
'95	Download Flexible Template from GMM
'95	Upload Flexible Template to GMM
'95	Associate Flexible Template with Gateway Profile in GMM
'95	Un-claim a Gateway
'95	Claim a Gateway
'95	Retrieve Gateway GPS Data for Last Hour
'95	Name/Rename Gateway within GMM
'95	Modify WiFi SSID and PSK
'95	Modify WGB SSID and PSK
'95 Retrieve GMM Org ID
'95 Retrieve GMM Org Tags
'95 Delete GMM org
'95 Create GMM org
'95 Create GMM Claim Policy
'95 Add User
*/

package gmmapi

import (
"bytes"
"encoding/json"
"fmt"
"io/ioutil"
"net/http"
"os"
"strconv"
"time"

)

// Function to retrieve the GMM API Key
// Need to supply GMM username and password
func Retrieve_gmm_api_key(email string, password string) string {
fmt.Println("Retrieving GMM API Key")

// A Response struct to map the entire response
type Response struct {
Access_token string		`json: "access_token"`
Expires_in int			`json: "expires_in"`
Token_type string		`json: "token_type"`
}

jsonData := map[string]string{"email": email, "password": password}
jsonValue, _ := json.Marshal(jsonData)
request, _ := http.NewRequest("POST", "https://us.ciscokinetic.io/api/v2/users/access_token", bytes.NewBuffer(jsonValue))
request.Header.Set("Content-Type", "application/json")
client := &http.Client{}
r, err := client.Do(request)

if err != nil {
fmt.Printf("API Token retrieve failed with the error %s\n", err)
os.Exit(1)
}

responseData, _ := ioutil.ReadAll(r.Body)
var responseObject Response
e := json.Unmarshal(responseData, &responseObject)
if e != nil {
fmt.Println("Unmarshall Error: ", e)
}

fmt.Println("GMM API Key:", responseObject.Access_token)
return responseObject.Access_token
}

// Function to retrieve Gateway Health Summary
// Need to supply GMM API Key and GMM Org ID
func Retrieve_gmm_gwy_health_summary(gmm_api_key string, org_id int) {

type gwy_status struct {
Summary struct {
Claiming   int `json:"claiming"`
Inactive   int `json:"inactive"`
InProgress int `json:"in_progress"`
Up         int `json:"up"`
Down       int `json:"down"`
Failed     int `json:"failed"`
} `json:"summary"`
}

jsonValue, _ := json.Marshal("")
request, _ := http.NewRequest("GET", "https://us.ciscokinetic.io/api/v2/organizations/" + strconv.Itoa(org_id) + "/gate_ways", bytes.NewBuffer(jsonValue))
token := "Token " + gmm_api_key
request.Header.Set("Authorization", token)
client := &http.Client{}
r, err := client.Do(request)

if err != nil {
fmt.Printf("Retrieve GMM GWY STATUS error %s\n", err)
os.Exit(1)
}

responseData, _ := ioutil.ReadAll(r.Body)
var responseObject gwy_status
e := json.Unmarshal(responseData, &responseObject)
if e != nil {
fmt.Println("Unmarshall Error: ", e)
}

fmt.Println("")
fmt.Println("GMM Gateway Health Summary")
fmt.Println(("--------------------------"))
fmt.Println("Claiming:", responseObject.Summary.Claiming)
fmt.Println("Inactive:", responseObject.Summary.Inactive)
fmt.Println("In_Progress:", responseObject.Summary.InProgress)
fmt.Println("UP:", responseObject.Summary.Up)
fmt.Println("DOWN:", responseObject.Summary.Down)
fmt.Println("FAILED:", responseObject.Summary.Failed)
}

// Function to retrieve a list of gateway profiles in GMM
// Need to supply GMM API Key
func Retrieve_gmm_gwy_profiles_list(gmm_api_key string, org_id int) {

type gwy_profiles struct {
GatewayProfiles []struct {
ID                               int           `json:"id"`
Name                             string        `json:"name"`
} `json:"gateway_profiles"`
Paging struct {
Limit  int `json:"limit"`
Offset int `json:"offset"`
Pages  int `json:"pages"`
Count  int `json:"count"`
Links  struct {
First string `json:"first"`
Last  string `json:"last"`
Next  string `json:"next"`
} `json:"links"`
} `json:"paging"`
}

jsonValue, _ := json.Marshal("")
request, _ := http.NewRequest("GET", "https://us.ciscokinetic.io/api/v2/organizations/" + strconv.Itoa(org_id) + "/gateway_profiles?limit=100", bytes.NewBuffer(jsonValue))
token := "Token " + gmm_api_key
request.Header.Set("Authorization", token)
client := &http.Client{}
r, err := client.Do(request)

if err != nil {
fmt.Printf("Retrieve GMM GWY Profiles error %s\n", err)
os.Exit(1)
}

responseData, _ := ioutil.ReadAll(r.Body)

var responseObject gwy_profiles
e := json.Unmarshal(responseData, &responseObject)
if e != nil {
fmt.Println("Unmarshall Error: ", e)
}

fmt.Println("")
fmt.Println("Total Number of Gateway Profiles in GMM: ", len(responseObject.GatewayProfiles))
fmt.Println("")
fmt.Println("Gateway Profiles in GMM")
fmt.Println("-----------------------")
for i := 0; i < len(responseObject.GatewayProfiles); i++ {
fmt.Println("Profile-ID: ", responseObject.GatewayProfiles[i].ID, " Profile Name: ", responseObject.GatewayProfiles[i].Name)
}
}

// Function to retrieve a List of Flexible Templates in GMM
// Need to supply GMM API Key and GMM Org ID
func Retrieve_gmm_flex_template_list(gmm_api_key string, org_id int) {

type flex_templates_list struct {
FlexibleTemplates []struct {
ID          int      `json:"id"`
Name        string   `json:"name"`
Description string   `json:"description"`
Template    string   `json:"template"`
Variables   []string `json:"variables"`
} `json:"flexible_templates"`
Paging struct {
Limit  int `json:"limit"`
Offset int `json:"offset"`
Pages  int `json:"pages"`
Count  int `json:"count"`
Links  struct {
First string `json:"first"`
Last  string `json:"last"`
Prev  string `json:"prev"`
Next  string `json:"next"`
} `json:"links"`
} `json:"paging"`
}

jsonValue, _ := json.Marshal("")
request, _ := http.NewRequest("GET", "https://us.ciscokinetic.io/api/v2/organizations/" + strconv.Itoa(org_id) + "/flexible_templates", bytes.NewBuffer(jsonValue))
token := "Token " + gmm_api_key
request.Header.Set("Authorization", token)
client := &http.Client{}
r, err := client.Do(request)

if err != nil {
fmt.Printf("Retrieve GMM GWY Profiles error %s\n", err)
os.Exit(1)
}

responseData, _ := ioutil.ReadAll(r.Body)

var responseObject flex_templates_list
e := json.Unmarshal(responseData, &responseObject)
if e != nil {
fmt.Println("Unmarshall Error: ", e)
}

fmt.Println("")
fmt.Println("Total Number of Flexible Templates in GMM: ", len(responseObject.FlexibleTemplates))
fmt.Println("")
fmt.Println("Flexible Templates in GMM")
fmt.Println("-------------------------")
for i := 0; i < len(responseObject.FlexibleTemplates); i++ {
fmt.Println("Flex-Template-ID: ", responseObject.FlexibleTemplates[i].ID, " Flex Template Name: ", responseObject.FlexibleTemplates[i].Name)
}
}

// Function to retrieve GMM Gateway ID correponding to a particular Gateway S/N
// Need to supply GMM API Key and Gateway S/N
func Retrieve_gmm_gwy_id(gmm_api_key string, org_id int, gwy_sn string) (id int) {

type gwy_status struct {
GateWays []struct {
ID                      int           `json:"id"`
UUID                    string        `json:"uuid"`
} `json:"gate_ways"`
Paging struct {
Limit  int `json:"limit"`
Offset int `json:"offset"`
Pages  int `json:"pages"`
Count  int `json:"count"`
Links  struct {
First string `json:"first"`
Last  string `json:"last"`
} `json:"links"`
} `json:"paging"`
}

jsonValue, _ := json.Marshal("")
request, _ := http.NewRequest("GET", "https://us.ciscokinetic.io/api/v2/organizations/" + strconv.Itoa(org_id) + "/gate_ways", bytes.NewBuffer(jsonValue))
token := "Token " + gmm_api_key
request.Header.Set("Authorization", token)
client := &http.Client{}
r, err := client.Do(request)

if err != nil {
fmt.Printf("Retrieve GMM GWY STATUS error %s\n", err)
os.Exit(1)
}

responseData, _ := ioutil.ReadAll(r.Body)
var responseObject gwy_status
e := json.Unmarshal(responseData, &responseObject)
if e != nil {
fmt.Println("Unmarshall Error: ", e)
}

gwy_id := 0
for i:= 0; i < len(responseObject.GateWays); i++ {
if responseObject.GateWays[i].UUID == gwy_sn {
gwy_id = responseObject.GateWays[i].ID
}
}

return gwy_id
}

// Function to Unclaim a Gateway from GMM
// Need to supply GMM API Key, GMM Org ID and the Gateway S/N
func Gmm_unclaim_gwy(gmm_api_key string, org_id int, gwy_sn string) {

type unclaim struct {
Id   int
UUID string
Name string
}

// Retrieving corresponding gateway ID
gwy_id := Retrieve_gmm_gwy_id(gmm_api_key, org_id, gwy_sn)

if gwy_id == 0 {
fmt.Println("")
fmt.Println("Gateway " + gwy_sn + " could not be unclaimed since it's currently not claimed within GMM")
return
}

jsonValue, _ := json.Marshal("")
url := "https://us.ciscokinetic.io/api/v2/claims/" + strconv.Itoa(gwy_id)
request, _ := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonValue))
token := "Token " + gmm_api_key
request.Header.Set("Authorization", token)
client := &http.Client{}
r, err := client.Do(request)

if err != nil {
fmt.Printf("Unclaim Gateway error %s\n", err)
os.Exit(1)
}

responseData, _ := ioutil.ReadAll(r.Body)
var responseObject unclaim
e := json.Unmarshal(responseData, &responseObject)
if e != nil {
fmt.Println("Unmarshall Error: ", e)
} else {
fmt.Println("")
fmt.Println("Gateway", responseObject.UUID, "with ID", responseObject.Id, "Unclaimed")
}

time.Sleep(250000000000)
}

// Function to Name/Re-name a Gateway in GMM
// Need to supply GMM API Key, GMM Org ID, Gateway S/N and the Gateway Name to be configured
func Gmm_rename_gwy(gmm_api_key string, org_id int, gwy_sn string, gwy_name string) {

data := []byte(`{ "gate_way": { "name": "` + gwy_name + `" } }`)

// Retrieving corresponding gateway ID
gwy_id := Retrieve_gmm_gwy_id(gmm_api_key, org_id, gwy_sn)

url := "https://us.ciscokinetic.io/api/v2/gate_ways/" + strconv.Itoa(gwy_id)
request, _ := http.NewRequest("PUT", url, bytes.NewBuffer(data))
token := "Token " + gmm_api_key
request.Header.Set("Authorization", token)
request.Header.Set("Content-Type", "application/json")
client := &http.Client{}
r, err := client.Do(request)

if err != nil {
fmt.Printf("Renaming gateway failed with error %s\n", err)
os.Exit(1)
}

responseData, _ := ioutil.ReadAll(r.Body)

fmt.Println()
fmt.Println("Rename Gateway Successful: " + string(responseData))
}

// Function to retrieve Gateway GPS Data for last hour
// Need to supply GMM API Key, GMM Org ID and the Gateway S/N
// Returns Gateway GPS Data as a JSON blob
func Retrieve_gmm_gwy_gps(gmm_api_key string, org_id int, gwy_sn string) (gps_history string) {

jsonValue, _ := json.Marshal("")

// Retrieving corresponding gateway ID
gwy_id := Retrieve_gmm_gwy_id(gmm_api_key, org_id, gwy_sn)

now := time.Now()
from_time := (now.Unix() - 3600) * 1000
to_time := now.Unix() * 1000
url := "https://us.ciscokinetic.io/api/v2/gate_ways/" + strconv.Itoa(gwy_id) + "/gps_history?from_time=" + strconv.FormatInt(from_time, 10) + "&to_time=" + strconv.FormatInt(to_time, 10)

request, _ := http.NewRequest("GET", url, bytes.NewBuffer(jsonValue))
token := "Token " + gmm_api_key
request.Header.Set("Authorization", token)
client := &http.Client{}
r, err := client.Do(request)

if err != nil {
fmt.Printf("Retrieve GMM GWY GPS History error %s\n", err)
os.Exit(1)
}

responseData, _ := ioutil.ReadAll(r.Body)
return string(responseData)
}

// Function to retrieve a particular Gateway Profile
// Need to supply GMM API Key and Profile Name
// Saves Gateway Profile as JSON file in the /tmp directory
func Retrieve_gmm_gwy_profile(gmm_api_key string, org_id int, profile_name string) {

// Retrieve Gateway Profile ID from GMM
profile_id := Retrieve_gmm_profile_id(gmm_api_key, org_id, profile_name)

jsonValue, _ := json.Marshal("")
url := "https://us.ciscokinetic.io/api/v2/gateway_profiles/" + strconv.Itoa(profile_id)
request, _ := http.NewRequest("GET", url, bytes.NewBuffer(jsonValue))
token := "Token " + gmm_api_key
request.Header.Set("Authorization", token)
client := &http.Client{}
r, err := client.Do(request)

if err != nil {
fmt.Printf("Retrieve GMM GWY Profile error %s\n", err)
os.Exit(1)
}

responseData, _ := ioutil.ReadAll(r.Body)
fmt.Println("Retrieve GWY Profiles = " +string(responseData))

filename := "/tmp/" + profile_name + ".json"
ioutil.WriteFile(filename, responseData, 0644)
}

// Function to Upload a Gateway Profile to GMM
// Need to supply GMM API Key, GMM Org ID, Profile as JSON File
// The Gateway Profile JSON file needs to be in the same directory as this script
func Gmm_upload_gwy_profile(gmm_api_key string, org_id int, gw_profile string) {

byteValue := []byte(gw_profile)

url := "https://us.ciscokinetic.io/api/v2/organizations/" + strconv.Itoa(org_id) + "/gateway_profiles"
request, _ := http.NewRequest("POST", url, bytes.NewBuffer(byteValue))
token := "Token " + gmm_api_key
request.Header.Set("Authorization", token)
request.Header.Set("Content-Type", "application/json")
client := &http.Client{}
r, err := client.Do(request)

if err != nil {
fmt.Printf("Upload of Gateway Profile failed with error %s\n", err)
os.Exit(1)
}

responseData, _ := ioutil.ReadAll(r.Body)
fmt.Println("Gateway Profile Uploaded into GMM: " + string(responseData))
}

// Function to Upload a Flexible Template to GMM
// Need to supply GMM API Key, GMM Org ID, Flexible Template as JSON File
// The Flex Template JSON file needs to be in the same directory as this script
func Gmm_upload_flex_template(gmm_api_key string, org_id int, flex_template string) {

//jsonFile, err := os.Open(flex_template_filename)

//if err != nil {
//fmt.Println(err)
//os.Exit(1)
//}

//fmt.Println("")
//fmt.Println("Successfully opened " + flex_template_filename)

//byteValue, _ := ioutil.ReadAll(jsonFile)
byteValue := []byte(flex_template)


url := "https://us.ciscokinetic.io/api/v2/organizations/" + strconv.Itoa(org_id) + "/flexible_templates"
request, _ := http.NewRequest("POST", url, bytes.NewBuffer(byteValue))
token := "Token " + gmm_api_key
request.Header.Set("Authorization", token)
request.Header.Set("Content-Type", "application/json")
client := &http.Client{}
r, err := client.Do(request)

if err != nil {
fmt.Printf("Upload of Flexible Template failed with error %s\n", err)
os.Exit(1)
}

responseData, _ := ioutil.ReadAll(r.Body)
fmt.Println("Flexible Template Uploaded into GMM: " + string(responseData))
}

// Function to retrieve Gateway Profile ID corresponding to Profile Name
// Need to supply GMM API Key, GMM Org ID and the Profile Name
// Returns Profile ID
func Retrieve_gmm_profile_id(gmm_api_key string, org_id int, profile_name string) (pid int) {

type gwy_profiles struct {
GatewayProfiles []struct {
ID                               int           `json:"id"`
Name                             string        `json:"name"`
} `json:"gateway_profiles"`
Paging struct {
Limit  int `json:"limit"`
Offset int `json:"offset"`
Pages  int `json:"pages"`
Count  int `json:"count"`
Links  struct {
First string `json:"first"`
Last  string `json:"last"`
Next  string `json:"next"`
} `json:"links"`
} `json:"paging"`
}

jsonValue, _ := json.Marshal("")
request, _ := http.NewRequest("GET", "https://us.ciscokinetic.io/api/v2/organizations/" + strconv.Itoa(org_id) + "/gateway_profiles?limit=100", bytes.NewBuffer(jsonValue))
token := "Token " + gmm_api_key
request.Header.Set("Authorization", token)
client := &http.Client{}
r, err := client.Do(request)

if err != nil {
fmt.Printf("Retrieve GMM GWY Profiles error %s\n", err)
os.Exit(1)
}

responseData, _ := ioutil.ReadAll(r.Body)
var responseObject gwy_profiles
e := json.Unmarshal(responseData, &responseObject)
if e != nil {
fmt.Println("Unmarshall Error: ", e)
os.Exit(1)
}

profile_id := 0
for i := 0; i < len(responseObject.GatewayProfiles); i++ {
if responseObject.GatewayProfiles[i].Name == profile_name {
profile_id = responseObject.GatewayProfiles[i].ID
}
}

fmt.Println("")
fmt.Println("Profile ID for Gateway Profile " + profile_name + " is: " + strconv.Itoa(profile_id))
return profile_id
}

// Function to retrieve Flexible Template ID
// Need to supply GMM API Key, GMM Org ID and the Flexible Template Name
// Returns the GMM Flexible Template ID
func Retrieve_gmm_flex_template_id(gmm_api_key string, org_id int, flex_template_name string) (ftid int) {

type flex_template_list struct {
FlexibleTemplates []struct {
ID          int      `json:"id"`
Name        string   `json:"name"`
Description string   `json:"description"`
Template    string   `json:"template"`
Variables   []string `json:"variables"`
} `json:"flexible_templates"`
Paging struct {
Limit  int `json:"limit"`
Offset int `json:"offset"`
Pages  int `json:"pages"`
Count  int `json:"count"`
Links  struct {
First string `json:"first"`
Last  string `json:"last"`
Prev  string `json:"prev"`
Next  string `json:"next"`
} `json:"links"`
} `json:"paging"`
}

jsonValue, _ := json.Marshal("")
request, _ := http.NewRequest("GET", "https://us.ciscokinetic.io/api/v2/organizations/" + strconv.Itoa(org_id) + "/flexible_templates", bytes.NewBuffer(jsonValue))
token := "Token " + gmm_api_key
request.Header.Set("Authorization", token)
client := &http.Client{}
r, err := client.Do(request)

if err != nil {
fmt.Printf("Retrieve GMM GWY Profiles error %s\n", err)
os.Exit(1)
}

responseData, _ := ioutil.ReadAll(r.Body)
var responseObject flex_template_list
e := json.Unmarshal(responseData, &responseObject)
if e != nil {
fmt.Println("Unmarshall Error: ", e)
os.Exit(1)
}

flex_template_id := 0
for i := 0; i < len(responseObject.FlexibleTemplates); i++ {
if responseObject.FlexibleTemplates[i].Name == flex_template_name {
flex_template_id = responseObject.FlexibleTemplates[i].ID
}
}

fmt.Println("")
fmt.Println("Flexible Template ID for Flexible Template " + flex_template_name + " is: " + strconv.Itoa(flex_template_id))
return flex_template_id
}

// Function to retrieve a particular Flexible Template
// Need to supply GMM API Key, GMM Org ID and the Flexible Template Name
// Saves Flexible Template as JSON file in the /tmp directory
func Retrieve_gmm_flex_template(gmm_api_key string, org_id int, ft_name string) {

// Retrieve Flexible Template ID from GMM
ft_id := Retrieve_gmm_flex_template_id(gmm_api_key, org_id, ft_name)

jsonValue, _ := json.Marshal("")
url := "https://us.ciscokinetic.io/api/v2/flexible_templates/" + strconv.Itoa(ft_id)
request, _ := http.NewRequest("GET", url, bytes.NewBuffer(jsonValue))
token := "Token " + gmm_api_key
request.Header.Set("Authorization", token)
client := &http.Client{}
r, err := client.Do(request)

if err != nil {
fmt.Printf("Retrieve GMM Flexible Template error %s\n", err)
os.Exit(1)
}

responseData, _ := ioutil.ReadAll(r.Body)
fmt.Println(string(responseData))

filename := "/tmp/" + ft_name + ".json"
ioutil.WriteFile(filename, responseData, 0644)
}

// Function to Modify WiFi SSID and PSK
// Need to supply GMM API Key, GMM Org ID, Gateway Profile Name, New WiFi SSID and/or New WiFi PSK
func Gmm_modify_gwy_wifi(gmm_api_key string, org_id int, profile_name string, wifi_ssid string, wifi_psk string) {

data := []byte(`{ "wifi_ssid": "` + wifi_ssid + `", "wifi_pre_shared_key": "` + wifi_psk + `" }`)

profile_id := Retrieve_gmm_profile_id(gmm_api_key, org_id, profile_name)

url := "https://us.ciscokinetic.io/api/v2/gateway_profiles/" + strconv.Itoa(profile_id)
request, _ := http.NewRequest("PUT", url, bytes.NewBuffer(data))
token := "Token " + gmm_api_key
request.Header.Set("Authorization", token)
request.Header.Set("Content-Type", "application/json")
client := &http.Client{}
r, err := client.Do(request)

if err != nil {
fmt.Printf("Modifying gateway WiFi settings failed with error %s\n", err)
os.Exit(1)
}

responseData, _ := ioutil.ReadAll(r.Body)
fmt.Println("Modifying Gateway WiFi settings : " + string(responseData))
}

// Function to Modify WGB SSID and PSK
// Need to supply GMM API Key, GMM Org ID, Gateway Profile Name, New WGB SSID and/or New WGB PSK
func Gmm_modify_gwy_wgb(gmm_api_key string, org_id int, profile_name string, wgb_ssid string, wgb_psk string) {

data := []byte(`{ "wgb_ssid": "` + wgb_ssid + `", "wgb_pre_shared_key": "` + wgb_psk + `" }`)

profile_id := Retrieve_gmm_profile_id(gmm_api_key, org_id, profile_name)

url := "https://us.ciscokinetic.io/api/v2/gateway_profiles/" + strconv.Itoa(profile_id)
request, _ := http.NewRequest("PUT", url, bytes.NewBuffer(data))
token := "Token " + gmm_api_key
request.Header.Set("Authorization", token)
request.Header.Set("Content-Type", "application/json")
client := &http.Client{}
r, err := client.Do(request)

if err != nil {
fmt.Printf("Modifying gateway WGB settings failed with error %s\n", err)
os.Exit(1)
}

responseData, _ := ioutil.ReadAll(r.Body)
fmt.Println("Modifying Gateway WGB settings : " + string(responseData))
}

// Function to claim a gateway in GMM
// Need to supply GMM API Key, GMM Org ID, Gateway S/N, Gateway Model and Gateway Profile
func Gmm_claim_gwy(gmm_api_key string, org_id int, gwy_sn string, model string, profile_name string) {

profile_id := Retrieve_gmm_profile_id(gmm_api_key, org_id, profile_name)

payload := `{ "claim_ids": ["` + gwy_sn + `"], "gateway_profile_id": ` + strconv.Itoa(profile_id) + `, "model": "` + model + `" }`

data := []byte(payload)

url := "https://us.ciscokinetic.io/api/v2/organizations/" + strconv.Itoa(org_id) + "/claims"
request, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
token := "Token " + gmm_api_key
request.Header.Set("Authorization", token)
request.Header.Set("Content-Type", "application/json")
client := &http.Client{}
r, err := client.Do(request)

if err != nil {
fmt.Printf("Claiming Gateway failed with error %s\n", err)
os.Exit(1)
}

responseData, _ := ioutil.ReadAll(r.Body)
fmt.Println("")
fmt.Println("Claiming Gateway : " + string(responseData))
}

// Function to associate flexible template

func Gmm_associate_flex_template(gmm_api_key string, org_id int, profile_name string, flex_template_name string) {

pid := Retrieve_gmm_profile_id(gmm_api_key, org_id, profile_name)
ftid := Retrieve_gmm_flex_template_id(gmm_api_key, org_id, flex_template_name)

data := []byte(`{"flexible_template_id": ` + strconv.Itoa(ftid) + `, "flexible_template_enable": true, "flexible_template_advanced": false, "flexible_template_variables": [{ "name": "", "value": "none"}]}`)

url := "https://us.ciscokinetic.io/api/v2/gateway_profiles/" + strconv.Itoa(pid)
request, _ := http.NewRequest("PUT", url, bytes.NewBuffer(data))
token := "Token " + gmm_api_key
request.Header.Set("Authorization", token)
request.Header.Set("Content-Type", "application/json")
client := &http.Client{}
r, err := client.Do(request)

if err != nil {
fmt.Printf("Associating Flex Template with Base Template failed with error %s\n", err)
os.Exit(1)
}

responseData, _ := ioutil.ReadAll(r.Body)
fmt.Println("Associated Flexible Template : " + string(responseData))
}

// Function to retrieve Gateway tags for an oganization.  Returns tag_id
func Retrieve_gmm_org_tag_id(gmm_api_key string, org_id int, varnum int) (tag_id int) {

type org_tag struct {
ID          int     	`json:"id"`
Default_Name 	string	`json:"default_name"`
Custom_Name    	string	`json:"custom_name"`
Org_ID		int			`json:"organization_id"`
created_at  string 		`json:"created_at"`
updated_at  string 		`json:"updated_at"`
Enabled     bool		`json:"enabled"`
Field_Type  string		`json:"enabled"`
Allowed_Values   []string	`json:"allowed_values"`
data_route_key  string		`json:"data_route_key"`
guid        string			`json:"guid"`

}

if varnum > 5 {
	fmt.Println("Error: only 6 total variables available")
	return -1
}

type Org_Tag struct {
	Collection []org_tag
}

jsonValue, _ := json.Marshal("")

url := "https://us.ciscokinetic.io/api/v2/organizations/" + strconv.Itoa(org_id) + "/tags?offset=" + strconv.Itoa(varnum) + "&limit=" + strconv.Itoa(1)

request, _ := http.NewRequest("GET", url, bytes.NewBuffer(jsonValue))
token := "Token " + gmm_api_key
request.Header.Set("Authorization", token)
client := &http.Client{}
r, err := client.Do(request)

if err != nil {
fmt.Printf("Retrieve GMM Tags error %s\n", err)
os.Exit(1)
}

responseData, _ := ioutil.ReadAll(r.Body)
responseObject := make([]org_tag,10)
e := json.Unmarshal(responseData, &responseObject)
if e != nil {
fmt.Println("Unmarshall Error: ", e)
  os.Exit(1)
}

tagid := 0
counter := 0
for i:=0; i<varnum; i++ {
	if responseObject[i].Org_ID ==  org_id  {
		if counter == 0 {
			tagid = responseObject[i].ID
			fmt.Println("Retrieved Tag ID = " + strconv.Itoa(responseObject[i].ID))
		}
		counter = counter + 1
	}
}
return tagid
//fmt.Println("Retrieved Tag ID = " + strconv.Itoa(responseObject[varnum].ID))
//return responseObject[varnum].ID


}

// Create a new org in GMM
func Create_gmm_org(gmm_api_key string, orgname string, parentOrg int) (orgid int) {

type org_info struct {
ID          int     	`json:"id"`
Owner_ID 	int			`json:"owner_id"`
created_at  string 		`json:"created_at"`
updated_at  string 		`json:"updated_at"`
field_director_id	int	`json:"field_director_id"`
field_director_object_ref 	[]string `json:"field_director_object_ref"`
fog_director_id	int		`json:"fog_director_id"`
ancestry	string		`json:"ancestry"`
fnd_template_file	string	`json:"fnd_template_file"`
tag_id		int			`json:"tag_id_for_data_routing"`
fnd_template_id	int		`json:"fnd_template_id"`
subscriber_uid	string	`json:"subscriber_uid"`
guid		string		`json:"guid"`
betas		[]string	`json:"betas"`
dcm_server	string		`json:"dcm_server"`
data_url	string		`json:"data_exchange_url"`
gw_show		string		`json:"gateway_show_commands"`
gw_debug	string		`json:"gateway_debug_commands"`
features    string		`json:"features"`

}

fmt.Println("Create GMM Organization with name of " + orgname + " and parent of " + strconv.Itoa(parentOrg))

// A Response struct to map the entire response
type Response struct {
Access_token string		`json: "access_token"`
Expires_in int			`json: "expires_in"`
Token_type string		`json: "token_type"`
}

jsonData := `{ "organization": { "name": "` + orgname + `", "parent_id": ` + strconv.Itoa(parentOrg) + `, "user_account": "false" }}`
jsonValue := []byte(jsonData)

request, _ := http.NewRequest("POST", "https://us.ciscokinetic.io/api/v2/organizations/", bytes.NewBuffer(jsonValue))
token := "Token " + gmm_api_key
request.Header.Set("Authorization", token)
request.Header.Set("Content-Type", "application/json")
client := &http.Client{}
r, err := client.Do(request)

if err != nil {
fmt.Printf("Org creation failed with the error %s\n", err)
os.Exit(1)
}

responseData, _ := ioutil.ReadAll(r.Body)
var responseObject org_info
e := json.Unmarshal(responseData, &responseObject)
if e != nil {
fmt.Println("Unmarshall Error: ", e)
os.Exit(1)
}

fmt.Println("Organization " + orgname + " created with ID = " + strconv.Itoa(responseObject.ID))
return responseObject.ID

}

// Function to update  tags for an oganization.  
func Update_gmm_org_tags(gmm_api_key string, tag_id int, tag string)  {



data := []byte(tag)

url := "https://us.ciscokinetic.io/api/v2/tags/" + strconv.Itoa(tag_id)
request, _ := http.NewRequest("PUT", url, bytes.NewBuffer(data))
token := "Token " + gmm_api_key
request.Header.Set("Authorization", token)
request.Header.Set("Content-Type", "application/json")
client := &http.Client{}
r, err := client.Do(request)

if err != nil {
fmt.Printf("Update GMM Tags error %s\n", err)
os.Exit(1)
}

responseData, _ := ioutil.ReadAll(r.Body)
fmt.Printf("Tag Update Response Data = " + string(responseData))
}

func Create_claim_policy(gmm_api_key string, org_id int, claimpolicy string) (cp_id int) {

type claim_policy struct {
ID			int			`json:"id"`
guid		string		`json:"guid"`
name  		string 		`json:"name"`
activated	bool 		`json:"activated"`
tag_id		int			`json:"tag_id`
policy_templates 	[]string `json:"policy_templates"`

}

// A Response struct to map the entire response
type Response struct {
Access_token string		`json: "access_token"`
Expires_in int			`json: "expires_in"`
Token_type string		`json: "token_type"`
}

jsonValue := []byte(claimpolicy)
uri := "https://us.ciscokinetic.io/api/v2/organizations/"+ strconv.Itoa(org_id) + "/gateway_claim_policies"
request, _ := http.NewRequest("POST", uri , bytes.NewBuffer(jsonValue))
token := "Token " + gmm_api_key
request.Header.Set("Authorization", token)
request.Header.Set("Content-Type", "application/json")
client := &http.Client{}
r, err := client.Do(request)

if err != nil {
fmt.Printf("claim policy creation failed with the error %s\n", err)
os.Exit(1)
}

responseData, _ := ioutil.ReadAll(r.Body)
var responseObject claim_policy
e := json.Unmarshal(responseData, &responseObject)
if e != nil {
fmt.Println("Unmarshall Error: ", e)
os.Exit(1)
}

fmt.Println("Created Claim Policy " + responseObject.name + " in org " + strconv.Itoa(org_id))
return responseObject.ID

}




// retrieve org ID
func Retrieve_gmm_org_id(gmm_api_key string, parent_org_id int, org_name string) (id int) {

type Org_profiles struct {
Organizations []struct {
Id                      int     `json:"id"`
Name					string	`json:"name"`
Owner_id                int     `json:"owner_id"`
Ancestry				string	`json:"ancestry"`
Ancestry_depth			int		`json:"acestry_depth"`
Tag_id					string	`json:"tag_id_for_data_routing"`
Betas					[]string	`json:"betas"`
Data_exchange_url		string	`json:"data_exchange_url"`
Gateway_show_commands	string	`json:"gateway_show_commands"`
Features				string 	`json:"features"`
} `json:"organizations"`
Paging struct {
Limit  int `json:"limit"`
Offset int `json:"offset"`
Pages  int `json:"pages"`
Count  int `json:"count"`
Links  struct {
First string `json:"first"`
Last  string `json:"last"`
} `json:"links"`
} `json:"paging"`
Child_organization_level	string	`json:"child_organization_level"`
}

jsonValue, _ := json.Marshal("")
request, _ := http.NewRequest("GET", "https://us.ciscokinetic.io/api/v2/organizations/" + strconv.Itoa(parent_org_id) + "/child_organizations?limit=200", bytes.NewBuffer(jsonValue))
token := "Token " + gmm_api_key
request.Header.Set("Authorization", token)
client := &http.Client{}
r, err := client.Do(request)

if err != nil {
fmt.Printf("Retrieve GMM ORG STATUS error %s\n", err)
os.Exit(1)
}


responseData, _ := ioutil.ReadAll(r.Body)
var responseObject Org_profiles
e := json.Unmarshal(responseData, &responseObject)
if e != nil {
fmt.Println("Unmarshall Error: ", e)
  os.Exit(1)
}

org_id := 0
for i:= 0; i < len(responseObject.Organizations); i++ {
	if responseObject.Organizations[i].Name == org_name {

		org_id = responseObject.Organizations[i].Id
		fmt.Println("Found Org Name " + org_name + " with org id " + strconv.Itoa(org_id))
	} 
}

return org_id
}

//function to delete org
func Gmm_delete_org(gmm_api_key string, parent_org_id int, org_name string ) {

type Org_info struct {
ID          int     	`json:"id"`
Name		string		`json:"name"`
Owner_ID 	int			`json:"owner_id"`
Created_at  string 		`json:"created_at"`
Updated_at  string 		`json:"updated_at"`
Field_director_id	int	`json:"field_director_id"`
Field_director_object_ref struct {} `json:"field_director_object_ref"`
Fog_director_id	int		`json:"fog_director_id"`
Ancestry	string		`json:"ancestry"`
Ancestry_depth 	int		`json:"ancestry_depth"`
Fnd_template_file	string	`json:"fnd_template_file"`
Tag_id		int			`json:"tag_id_for_data_routing"`
Fnd_template_id	int		`json:"fnd_template_id"`
Subscriber_uid	string	`json:"subscriber_uid"`
Guid		string		`json:"guid"`
Betas		[]string	`json:"betas"`
Dcm			string		`json:"dcm_server"`
Dxc_url		string		`json:"data_exchange_url"`
Gw_show		string		`json:"gateway_show_commands"`
Gw_debug	string		`json:"gateway_debug_commands"`
Alerts		int			`json:"alert_email_count"`
Alert_date	string		`json:"alerts_limit_email_sent_on"`
Features    string		`json:"features"`
FirstAlert	string		`json:"first_alert_for_day"`
Ui_Column	string		`json:"ui_column"`
Cust_Id		string		`json:"customer_identifier"`
Cust_Type	string		`json:"customer_type"`
Notes		string		`json:"notes"`

}

// Retrieving corresponding gateway ID
org_id := Retrieve_gmm_org_id(gmm_api_key, parent_org_id,  org_name )

if org_id == 0 {
fmt.Println("")
fmt.Println("Gateway " + org_name + " could not be deleted since doesn't exist")
return
}

jsonValue, _ := json.Marshal("")
url := "https://us.ciscokinetic.io/api/v2/organizations/" + strconv.Itoa(org_id)
request, _ := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonValue))
token := "Token " + gmm_api_key
request.Header.Set("Authorization", token)
client := &http.Client{}
r, err := client.Do(request)

if err != nil {
fmt.Printf("Delete org error %s\n", err)
os.Exit(1)
}

responseData, _ := ioutil.ReadAll(r.Body)
var responseObject Org_info
e := json.Unmarshal(responseData, &responseObject)
if e != nil {
fmt.Println("Unmarshall Error: ", e)
os.Exit(1)
}

if r.StatusCode != 200 {
		fmt.Println("Org " + org_name + " with ID " + strconv.Itoa(org_id) + " could not be deleted.  Status code: " + strconv.Itoa(r.StatusCode))
		os.Exit(1)
}

fmt.Println("")
fmt.Println("Org " + org_name + " with ID " + strconv.Itoa(org_id) + " deleted")


}

// Add user to org
func Gmm_add_user(gmm_api_key string, org_id int,  user_info string ) (user_id int) {

type User_response struct {
Id          int     	`json:"id"`
Org_Id		int			`json:"organization_id"`
Role	 	string		`json:"role"`
Created_at  string 		`json:"created_at"`
Updated_at  string 		`json:"updated_at"`
User struct {
	Confirmed_at	string	`json:"confirmed_at"`
	Email			string	`json:"mail"`
	Id				int		`json:"id"`
	Name			string	`json:"name"`
	
} `json:"user"`

}

/*
type user_info struct {
	membership struct {
Email		string			`json:"email"`
Name  		string 			`json:"name"`
Role		string 			`json:"role"`
	} `json:"membership"`

}
*/

if org_id == 0 {
fmt.Println("")
fmt.Println( strconv.Itoa(org_id) + " doesn't exist")
return
}


jsonValue := []byte(user_info)
uri := "https://us.ciscokinetic.io/api/v2/organizations/"+ strconv.Itoa(org_id) + "/memberships"
request, _ := http.NewRequest("POST", uri , bytes.NewBuffer(jsonValue))
token := "Token " + gmm_api_key
request.Header.Set("Authorization", token)
request.Header.Set("Content-Type", "application/json")
client := &http.Client{}
r, err := client.Do(request)


if err != nil {
fmt.Printf("Add user error %s\n", err)
os.Exit(1)
}

responseData, _ := ioutil.ReadAll(r.Body)
var responseObject User_response
e := json.Unmarshal(responseData, &responseObject)
if e != nil {
fmt.Println("Unmarshall Error: ", e)
os.Exit(1)
}

if r.StatusCode != 200 {
		fmt.Println("User could not be added.  Status code: " + strconv.Itoa(r.StatusCode))
		os.Exit(1)
}

fmt.Println("")
fmt.Println("User " + responseObject.User.Email + " added")

return responseObject.Id
}
