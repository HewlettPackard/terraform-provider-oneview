// (C) Copyright 2020 Hewlett Packard Enterprise Development LP
//
// Licensed under the Apache License, Version 2.0 (the "License");
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed
// under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package oneview

import (
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceServerCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCertificateCreate,
		Read:   resourceServerCertificateRead,
		Update: resourceServerCertificateUpdate,
		Delete: resourceServerCertificateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificate_details": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alias_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"alternative_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"base64_data": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"basic_constraints": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"category": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cert_path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"common_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"contact_person": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"country": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"crl_distribution_end_points": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dn_qualifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"etag": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"email": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enhanced_key_usage": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"expires_in_days": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"given_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"initials": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"issuer": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_usage": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"locality": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"location_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"modified": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"organization": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"organizational_unit": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"public_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"serial_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sha1_fingerprint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sha256_fingerprint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sha384_fingerprint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"signature_algorithm": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"surname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"valid_from": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"valid_until": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"certificate_status": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"chain_status": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"self_signed": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"trusted": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceServerCertificateCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	servC := ov.ServerCertificate{
		Category:    d.Get("category").(string),
		Created:     d.Get("created").(string),
		Description: utils.NewNstring(d.Get("description").(string)),
		ETAG:        d.Get("etag").(string),
		Modified:    d.Get("modified").(string),
		Name:        d.Get("name").(string),
		State:       d.Get("state").(string),
		Status:      d.Get("status").(string),
		Type:        d.Get("type").(string),
	}
	CertificateDetailCollect := make([]ov.CertificateDetail, 0)
	rawCertificateDetails := d.Get("certificate_details").(*schema.Set).List()
	for _, raw := range rawCertificateDetails {
		certificateDetails := raw.(map[string]interface{})
		servCCertificateDetails := ov.CertificateDetail{
			AliasName:          certificateDetails["alias_name"].(string),
			AlternativeName:    certificateDetails["alternative_name"].(string),
			Base64Data:         utils.NewNstring(certificateDetails["base64_data"].(string)),
			BasicConstraints:   certificateDetails["basic_constraints"].(string),
			Category:           certificateDetails["category"].(string),
			CommonName:         certificateDetails["common_name"].(string),
			ContactPerson:      certificateDetails["contact_person"].(string),
			Country:            certificateDetails["country"].(string),
			Created:            certificateDetails["created"].(string),
			Description:        certificateDetails["description"].(string),
			Dnqualifier:        certificateDetails["dn_qualifier"].(string),
			Etag:               certificateDetails["etag"].(string),
			Email:              certificateDetails["email"].(string),
			EnhancedKeyUsage:   certificateDetails["enhanced_key_usage"].(string),
			ExpiresInDays:      certificateDetails["expires_in_days"].(string),
			GivenName:          certificateDetails["given_name"].(string),
			Initials:           certificateDetails["initials"].(string),
			Issuer:             certificateDetails["issuer"].(string),
			KeyUsage:           certificateDetails["key_usage"].(string),
			Locality:           certificateDetails["locality"].(string),
			LocationState:      certificateDetails["location_state"].(string),
			Modified:           certificateDetails["modified"].(string),
			Name:               certificateDetails["name"].(string),
			Organization:       certificateDetails["organization"].(string),
			OrganizationalUnit: certificateDetails["organizational_unit"].(string),
			PublicKey:          utils.NewNstring(certificateDetails["public_key"].(string)),
			SerialNumber:       utils.NewNstring(certificateDetails["serial_number"].(string)),
			Sha1Fingerprint:    utils.NewNstring(certificateDetails["sha1_fingerprint"].(string)),
			Sha256Fingerprint:  utils.NewNstring(certificateDetails["sha256_fingerprint"].(string)),
			Sha384Fingerprint:  utils.NewNstring(certificateDetails["sha384_fingerprint"].(string)),
			SignatureAlgorithm: certificateDetails["signature_algorithm"].(string),
			State:              certificateDetails["state"].(string),
			Status:             certificateDetails["status"].(string),
			Surname:            certificateDetails["surname"].(string),
			Type:               certificateDetails["type"].(string),
		}
		CertificateDetailCollect = append(CertificateDetailCollect, servCCertificateDetails)
		servC.CertificateDetails = CertificateDetailCollect
	}
	CertificateStatusList := d.Get("certificate_status").(*schema.Set).List()
	for _, raw := range CertificateStatusList {
		certificatestat := raw.(map[string]interface{})
		servcCertificateStat := ov.CertificateStat{
			ChainStatus: certificatestat["chain_status"].(string),
			SelfSigned:  certificatestat["self_signed"].(bool),
			Trusted:     certificatestat["trusted"].(bool),
		}
		servC.CertificateStatus = &servcCertificateStat
	}

	servCError := config.ovClient.CreateServerCertificate(servC)
	var aliasname string
	for _, servCDetail := range servC.CertificateDetails {
		aliasname = servCDetail.AliasName
	}
	d.SetId(aliasname)
	if servCError != nil {
		d.SetId("")
		return servCError
	}
	return resourceServerCertificateRead(d, meta)
}

func resourceServerCertificateRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	servC, err := config.ovClient.GetServerCertificateByName(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}
	d.Set("category", servC.Category)
	servCCertificateDetails := make([]map[string]interface{}, 0, len(servC.CertificateDetails))
	for _, servCCertificateDetail := range servC.CertificateDetails {
		cdeplist := make([]interface{}, len(servCCertificateDetail.CrlDistributionEndPoints))
		for i, cdep := range servCCertificateDetail.CrlDistributionEndPoints {
			cdeplist[i] = cdep
		}
		servCCertificateDetails = append(servCCertificateDetails, map[string]interface{}{
			"alias_name":          servCCertificateDetail.AliasName,
			"alternative_name":    servCCertificateDetail.AlternativeName,
			"base64_data":         servCCertificateDetail.Base64Data,
			"basic_constraints":   servCCertificateDetail.BasicConstraints,
			"category":            servCCertificateDetail.Category,
			"common_name":         servCCertificateDetail.CommonName,
			"contact_person":      servCCertificateDetail.ContactPerson,
			"country":             servCCertificateDetail.Country,
			"created":             servCCertificateDetail.Created,
			"description":         servCCertificateDetail.Description,
			"dn_qualifier":        servCCertificateDetail.Dnqualifier,
			"etag":                servCCertificateDetail.Etag,
			"email":               servCCertificateDetail.Email,
			"enhanced_key_usage":  servCCertificateDetail.EnhancedKeyUsage,
			"expires_in_days":     servCCertificateDetail.ExpiresInDays,
			"given_name":          servCCertificateDetail.GivenName,
			"initials":            servCCertificateDetail.Initials,
			"issuer":              servCCertificateDetail.Issuer,
			"key_usage":           servCCertificateDetail.KeyUsage,
			"locality":            servCCertificateDetail.Locality,
			"location_state":      servCCertificateDetail.LocationState,
			"modified":            servCCertificateDetail.Modified,
			"name":                servCCertificateDetail.Name,
			"organization":        servCCertificateDetail.Organization,
			"organizational_unit": servCCertificateDetail.OrganizationalUnit,
			"public_key":          servCCertificateDetail.PublicKey,
			"serial_number":       servCCertificateDetail.SerialNumber,
			"sha1_fingerprint":    servCCertificateDetail.Sha1Fingerprint,
			"sha256_fingerprint":  servCCertificateDetail.Sha256Fingerprint,
			"sha384_fingerprint":  servCCertificateDetail.Sha384Fingerprint,
			"signature_algorithm": servCCertificateDetail.SignatureAlgorithm,
			"state":               servCCertificateDetail.State,
			"status":              servCCertificateDetail.Status,
			"surname":             servCCertificateDetail.Surname,
			"type":                servCCertificateDetail.Type,
		})
	}
	d.Set("certificate_details", servCCertificateDetails)
	servCCertificateStatus := make([]map[string]interface{}, 0, 1)
	servCCertificateStatus = append(servCCertificateStatus, map[string]interface{}{
		"chain_status": servC.CertificateStatus.ChainStatus,
		"self_signed":  servC.CertificateStatus.SelfSigned,
		"trusted":      servC.CertificateStatus.Trusted,
	})

	d.Set("certificate_status", servCCertificateStatus)
	d.Set("created", servC.Created)
	d.Set("description", servC.Description)
	d.Set("etag", servC.ETAG)
	d.Set("modified", servC.Modified)
	d.Set("name", servC.Name)
	d.Set("state", servC.State)
	d.Set("status", servC.Status)
	d.Set("type", servC.Type)
	d.Set("uri", servC.URI)
	return nil
}

func resourceServerCertificateUpdate(d *schema.ResourceData, meta interface{}) error {

	config := meta.(*Config)
	servC := ov.ServerCertificate{
		Category:    d.Get("category").(string),
		Created:     d.Get("created").(string),
		Description: utils.NewNstring(d.Get("description").(string)),
		ETAG:        d.Get("etag").(string),
		Modified:    d.Get("modified").(string),
		Name:        d.Get("name").(string),
		State:       d.Get("state").(string),
		Status:      d.Get("status").(string),
		Type:        d.Get("type").(string),
		URI:         utils.NewNstring(d.Get("uri").(string)),
	}
	CertificateDetailCollect := make([]ov.CertificateDetail, 0)
	rawCertificateDetails := d.Get("certificate_details").(*schema.Set).List()
	for _, raw := range rawCertificateDetails {
		certificateDetails := raw.(map[string]interface{})
		servCCertificateDetails := ov.CertificateDetail{
			AliasName:          certificateDetails["alias_name"].(string),
			AlternativeName:    certificateDetails["alternative_name"].(string),
			Base64Data:         utils.NewNstring(certificateDetails["base64_data"].(string)),
			BasicConstraints:   certificateDetails["basic_constraints"].(string),
			Category:           certificateDetails["category"].(string),
			CommonName:         certificateDetails["common_name"].(string),
			ContactPerson:      certificateDetails["contact_person"].(string),
			Country:            certificateDetails["country"].(string),
			Created:            certificateDetails["created"].(string),
			Description:        certificateDetails["description"].(string),
			Dnqualifier:        certificateDetails["dn_qualifier"].(string),
			Etag:               certificateDetails["etag"].(string),
			Email:              certificateDetails["email"].(string),
			EnhancedKeyUsage:   certificateDetails["enhanced_key_usage"].(string),
			ExpiresInDays:      certificateDetails["expires_in_days"].(string),
			GivenName:          certificateDetails["given_name"].(string),
			Initials:           certificateDetails["initials"].(string),
			Issuer:             certificateDetails["issuer"].(string),
			KeyUsage:           certificateDetails["key_usage"].(string),
			Locality:           certificateDetails["locality"].(string),
			LocationState:      certificateDetails["location_state"].(string),
			Modified:           certificateDetails["modified"].(string),
			Name:               certificateDetails["name"].(string),
			Organization:       certificateDetails["organization"].(string),
			OrganizationalUnit: certificateDetails["organizational_unit"].(string),
			PublicKey:          utils.NewNstring(certificateDetails["public_key"].(string)),
			SerialNumber:       utils.NewNstring(certificateDetails["serial_number"].(string)),
			Sha1Fingerprint:    utils.NewNstring(certificateDetails["sha1_fingerprint"].(string)),
			Sha256Fingerprint:  utils.NewNstring(certificateDetails["sha256_fingerprint"].(string)),
			Sha384Fingerprint:  utils.NewNstring(certificateDetails["sha384_fingerprint"].(string)),
			SignatureAlgorithm: certificateDetails["signature_algorithm"].(string),
			State:              certificateDetails["state"].(string),
			Status:             certificateDetails["status"].(string),
			Surname:            certificateDetails["surname"].(string),
			Type:               certificateDetails["type"].(string),
		}
		CertificateDetailCollect = append(CertificateDetailCollect, servCCertificateDetails)
		servC.CertificateDetails = CertificateDetailCollect
	}
	CertificateStatusList := d.Get("certificate_status").(*schema.Set).List()
	for _, raw := range CertificateStatusList {
		certificatestat := raw.(map[string]interface{})
		servcCertificateStat := ov.CertificateStat{
			ChainStatus: certificatestat["chain_status"].(string),
			SelfSigned:  certificatestat["self_signed"].(bool),
			Trusted:     certificatestat["trusted"].(bool),
		}
		servC.CertificateStatus = &servcCertificateStat
	}

	servCError := config.ovClient.UpdateServerCertificate(servC)
	var aliasname string
	for _, servCDetail := range servC.CertificateDetails {
		aliasname = servCDetail.AliasName
	}
	d.SetId(aliasname)
	if servCError != nil {
		d.SetId("")
		return servCError
	}
	return resourceServerCertificateRead(d, meta)

}
func resourceServerCertificateDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	err := config.ovClient.DeleteServerCertificate(d.Id())
	if err != nil {
		return err
	}
	return nil
}
