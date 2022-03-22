provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  i3s_endpoint  = var.i3s_endpoint
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Updates resource created from main.tf 

resource "oneview_deployment_plan" "dp" {
  name              = "Testdp"
  description       = "Testing creation of Deployment Plan"
  oe_build_plan_uri = "/rest/build-plans/97c12c1a-a048-4b52-bcdb-de7f20b5295f"
  hp_provided       = false
  type              = "OEDeploymentPlanV5"
  custom_attributes {
    constraints = "{\"options\":[\"English (United States)\",\"French (France)\",\"German (Germany)\",\"Japanese (Japan)\",\"Arabic (Saudi Arabia)\",\"Chinese (PRC)\",\"Korean (Korea)\",\"Portuguese (Brazil)\",\"Russian (Russia)\"]}"
    editable    = true
    id          = "4509965b-fcdb-4ab2-9e20-1b80294ce94f"
    name        = "DisplayLanguage"
    type        = "option"
    value       = "English (United States)"
    visible     = true
  }
  custom_attributes {
    constraints = "{\"options\":[\"English (United States)\",\"Arabic (101)\",\"Chinese (Traditional) - US Keyboard\",\"Japanese\",\"Korean\",\"United Kingdom Extended\",\"United States - Dvorak\"]}"
    editable    = true
    id          = "c6ef28cc-0562-4c1e-8454-8e295f4df00d"
    name        = "KeyboardLayout"
    type        = "option"
    value       = "English (United States)"
    visible     = true
  }
  custom_attributes {
    constraints = "{}"
    editable    = true
    id          = "e084491d-e476-4660-b231-b45a6cf2d42d"
    name        = "User1Password"
    type        = "string"
    visible     = true
  }
  custom_attributes {
    constraints = "{}"
    editable    = true
    id          = "056c6f33-6509-4b1c-bb93-17685e631f4d"
    name        = "User1DisplayName"
    type        = "string"
    visible     = true
  }
  custom_attributes {
    constraints = "{\"ipv4static\":true,\"ipv4dhcp\":true,\"ipv4disable\":false,\"parameters\":[\"mac\"]}"
    editable    = true
    id          = "8303c168-3e23-4025-9e63-2bddd644b461"
    name        = "ManagementNIC2"
    type        = "nic"
    visible     = true
  }
  custom_attributes {
    constraints = "{\"options\":[\"Disallow\",\"Allow (Network Level Authentication)\",\"Allow\"]}"
    editable    = true
    id          = "ab0aa2b0-6b4f-433c-a6a6-abcb949ac286"
    name        = "RemoteDesktop"
    type        = "option"
    value       = "Disallow"
    visible     = true
  }
  custom_attributes {
    constraints = "{}"
    editable    = true
    id          = "3c7c8229-1dc6-4656-857d-3392a26585ee"
    name        = "Hostname"
    type        = "string"
    visible     = true
  }
  custom_attributes {
    constraints = "{\"ipv4static\":true,\"ipv4dhcp\":true,\"ipv4disable\":false,\"parameters\":[\"dhcp\",\"dns1\",\"dns2\",\"gateway\",\"ipaddress\",\"mac\",\"netmask\"]}"
    editable    = true
    id          = "ec1d95d0-690a-482b-8efd-53bec6e9bfce"
    name        = "ManagementNIC1"
    type        = "nic"
    visible     = true
  }
  custom_attributes {
    constraints = "{\"maxlen\":\"20\"}"
    editable    = true
    description = "Administrator Password"
    id          = "a881b1af-9034-4c69-a890-5e8c83a13d25"
    name        = "Password"
    type        = "password"
    visible     = true
  }
  custom_attributes {
    constraints = "{\"options\":[\"Disallow\",\"Allow (Network Level Authentication)\",\"Allow\"]}"
    editable    = true
    id          = "ab0aa2b0-6b4f-433c-a6a6-abcb949ac286"
    name        = "RemoteDesktop"
    type        = "option"
    value       = "Disallow"
    visible     = true
  }
  custom_attributes {
    constraints = "{\"options\":[\"GMT Standard Time\",\"Arabian Standard Time\",\"AUS Eastern Standard Time\",\"Central Standard Time\",\"China Standard Time\",\"Eastern Standard Time\",\"India Standard Time\",\"Mountain Standard Time\",\"Singapore Standard Time\",\"Tokyo Standard Time\"]}"
    editable    = true
    id          = "0d629c3e-23d0-49e1-950c-ecfcb9d5610d"
    name        = "TimeZone"
    type        = "option"
    value       = "GMT Standard Time"
    visible     = true
  }
  custom_attributes {
    constraints = "{}"
    editable    = true
    id          = "1075adb8-5399-41fe-b6b3-71bea8836615"
    name        = "User1Name"
    type        = "string"
    visible     = true
  }
}

