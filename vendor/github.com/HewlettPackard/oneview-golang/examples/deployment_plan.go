package main

import (
	"encoding/json"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/i3s"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	var (
		clientOV             *ov.OVClient
		i3sClient            *i3s.I3SClient
		endpoint             = os.Getenv("ONEVIEW_OV_ENDPOINT")
		i3sc_endpoint        = os.Getenv("ONEVIEW_I3S_ENDPOINT")
		username             = os.Getenv("ONEVIEW_OV_USER")
		password             = os.Getenv("ONEVIEW_OV_PASSWORD")
		domain               = os.Getenv("ONEVIEW_OV_DOMAIN")
		deployment_plan_name = "TestDP"
		new_name             = "RenamedDeploymentPlan"
	)
	api_version, _ := strconv.Atoi(os.Getenv("ONEVIEW_APIVERSION"))
	i3s_apiversion, _ := strconv.Atoi(os.Getenv("ONEVIEW_I3S_APIVERSION"))
	ovc := clientOV.NewOVClient(
		username,
		password,
		domain,
		endpoint,
		false,
		api_version,
		"*")
	ovc.RefreshLogin()
	i3sc := i3sClient.NewI3SClient(i3sc_endpoint, false, i3s_apiversion, ovc.APIKey)

	customAttributes := new([]i3s.CustomAttribute)
	var ca1, ca2, ca3, ca4, ca5, ca6, ca7, ca8, ca9, ca10, ca11 i3s.CustomAttribute
	ca1.Constraints = "{\"options\":[\"English (United States)\",\"French (France)\",\"German (Germany)\",\"Japanese (Japan)\",\"Arabic " +
		"(Saudi Arabia)\",\"Chinese (PRC)\",\"Korean (Korea)\",\"Portuguese (Brazil)\",\"Russian (Russia)\"]}"
	ca1.Editable = true
	ca1.ID = "4509965b-fcdb-4ab2-9e20-1b80294ce94f"
	ca1.Name = "DisplayLanguage"
	ca1.Type = "option"
	ca1.Value = "English (United States)"
	ca1.Visible = true
	*customAttributes = append(*customAttributes, ca1)

	ca2.Constraints = "{\"options\":[\"English (United States)\",\"Arabic (101)\",\"Chinese (Traditional) - US Keyboard\",\"Japanese\",\"Korean\",\"United" +
		" Kingdom Extended\",\"United States - Dvorak\"]}"
	ca2.Editable = true
	ca2.ID = "c6ef28cc-0562-4c1e-8454-8e295f4df00d"
	ca2.Name = "KeyboardLayout"
	ca2.Type = "option"
	ca2.Value = "English (United States)"
	ca2.Visible = true
	*customAttributes = append(*customAttributes, ca2)

	ca3.Constraints = "{}"
	ca3.Editable = true
	ca3.ID = "e084491d-e476-4660-b231-b45a6cf2d42d"
	ca3.Name = "User1Password"
	ca3.Type = "string"
	ca3.Visible = true
	*customAttributes = append(*customAttributes, ca3)

	ca4.Constraints = "{}"
	ca4.Editable = true
	ca4.ID = "056c6f33-6509-4b1c-bb93-17685e631f4d"
	ca4.Name = "User1DisplayName"
	ca4.Type = "string"
	ca4.Visible = true
	*customAttributes = append(*customAttributes, ca4)

	ca5.Constraints = "{\"ipv4static\":true,\"ipv4dhcp\":true,\"ipv4disable\":false,\"parameters\":[\"dns1\",\"dns2\",\"gateway\",\"ipaddress\",\"mac\",\"netmask\",\"vlanid\"]}"
	ca5.Editable = true
	ca5.ID = "1b7ff92d-a4a6-4250-a93a-460209752c20"
	ca5.Name = "ManagementNIC2"
	ca5.Type = "nic"
	ca5.Visible = true
	*customAttributes = append(*customAttributes, ca5)

	ca6.Constraints = "{\"options\":[\"Disallow\",\"Allow (Network Level Authentication)\",\"Allow\"]}"
	ca6.Editable = true
	ca6.ID = "ab0aa2b0-6b4f-433c-a6a6-abcb949ac286"
	ca6.Name = "RemoteDesktop"
	ca6.Type = "option"
	ca6.Value = "Disallow"
	ca6.Visible = true
	*customAttributes = append(*customAttributes, ca6)

	ca7.Constraints = "{}"
	ca7.Editable = true
	ca7.ID = "3c7c8229-1dc6-4656-857d-3392a26585ee"
	ca7.Name = "Hostname"
	ca7.Type = "string"
	ca7.Visible = true
	*customAttributes = append(*customAttributes, ca7)

	ca8.Constraints = "{\"ipv4static\":true,\"ipv4dhcp\":true,\"ipv4disable\":false,\"parameters\":[\"dhcp\",\"dns1\",\"dns2\",\"gateway\",\"ipaddress\",\"mac\",\"netmask\"]}"
	ca8.Editable = true
	ca8.ID = "ec1d95d0-690a-482b-8efd-53bec6e9bfce"
	ca8.Name = "ManagementNIC1"
	ca8.Type = "nic"
	ca8.Visible = true
	*customAttributes = append(*customAttributes, ca8)

	ca9.Constraints = "{\"maxlen\":\"20\"}"
	ca9.Editable = true
	ca9.Description = "Administrator Password"
	ca9.ID = "a881b1af-9034-4c69-a890-5e8c83a13d25"
	ca9.Name = "Password"
	ca9.Type = "password"
	ca9.Visible = true
	*customAttributes = append(*customAttributes, ca9)

	ca10.Constraints = "{\"options\":[\"GMT Standard Time\",\"Arabian Standard Time\",\"AUS Eastern Standard Time\",\"Central Standard Time\",\"China " +
		"Standard Time\",\"Eastern Standard Time\",\"India Standard Time\",\"Mountain Standard Time\",\"Singapore Standard " +
		"Time\",\"Tokyo Standard Time\"]}"
	ca10.Editable = true
	ca10.ID = "0d629c3e-23d0-49e1-950c-ecfcb9d5610d"
	ca10.Name = "TimeZone"
	ca10.Type = "option"
	ca10.Value = "GMT Standard Time"
	ca10.Visible = true
	*customAttributes = append(*customAttributes, ca10)

	ca11.Constraints = "{}"
	ca11.Editable = true
	ca11.ID = "1075adb8-5399-41fe-b6b3-71bea8836615"
	ca11.Name = "User1Name"
	ca11.Type = "string"
	ca11.Visible = true
	*customAttributes = append(*customAttributes, ca11)

	var deploymentPlan i3s.DeploymentPlan
	deploymentPlan.Name = deployment_plan_name
	deploymentPlan.Type = "OEDeploymentPlanV5"
	deploymentPlan.OEBuildPlanURI = "/rest/build-plans/cbaeb42a-9cc7-4673-ae0e-4167c90f006a"
	deploymentPlan.CustomAttributes = *customAttributes
	deploymentPlan.HPProvided = false

	fmt.Println("HPProvided:", deploymentPlan.HPProvided)
	file, _ := json.MarshalIndent(deploymentPlan, "", " ")
	ioutil.WriteFile("inut.json", file, 0644)
	fmt.Println("***********Creating Deployment Plan****************")
	err := i3sc.CreateDeploymentPlan(deploymentPlan)
	if err != nil {
		fmt.Println("Deployment Plan Creation Failed: ", err)
	} else {
		fmt.Println("Deployment Plan created successfully...")
	}

	sort := "name:desc"
	count := "5"
	fmt.Println("**************Get Deployment Plans sorted by name in descending order****************")
	dps, err := i3sc.GetDeploymentPlans(count, "", "", sort, "")
	if err != nil {
		fmt.Println("Error while getting deployment plans:", err)
	} else {
		for i := range dps.Members {
			fmt.Println(dps.Members[i].Name)
		}
	}

	fmt.Println("***********Getting Deployment Plan By Name****************")
	deployment_plan, err := i3sc.GetDeploymentPlanByName(deployment_plan_name)
	if err != nil {
		fmt.Println("Error in getting deployment plan ", err)
	}
	fmt.Println(deployment_plan)

	fmt.Println("***********Updating Deployment Plan****************")
	deployment_plan.Name = new_name
	deployment_plan.GoldenImageUri = "/rest/golden-images/7e709af9-5446-426e-9ca1-df06c63df2cd"
	deployment_plan.Description = utils.NewNstring("Testing Deployment plan")
	err = i3sc.UpdateDeploymentPlan(deployment_plan)
	if err != nil {
		//panic(err)
		fmt.Println("Error whilw updating Deployment Plan:", err)
	} else {
		fmt.Println("Deployment Plan has been updated with name: " + deployment_plan.Name)
	}

	fmt.Println("***********Deleting Deployment Plan****************")
	err = i3sc.DeleteDeploymentPlan(new_name)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Deleteed Deployment Plan successfully...")
	}

}
