package ov

import (
	"encoding/json"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type ServerCertificate struct {
	Category           string              `json:"category,omitempty"`           // "category": "server-certificates",
	CertificateDetails []CertificateDetail `json:"certificateDetails,omitempty"` // "certificateDetails": "",
	CertificateStatus  *CertificateStat    `json:"certificateStatus,omitempty"`  // "certificateStatus": "",
	Created            string              `json:"created,omitempty"`            // "created": "20150831T154835.250Z",
	Description        utils.Nstring       `json:"description,omitempty"`        // "description": "Server Certificate 1",
	ETAG               string              `json:"eTag,omitempty"`               // "eTag": "1441036118675/8",
	Modified           string              `json:"modified,omitempty"`           // "modified": "20150831T154835.250Z",
	Name               string              `json:"name,omitempty"`               // "name": "hostname or IP",
	State              string              `json:"state,omitempty"`              // "state": "Connected",
	Status             string              `json:"status,omitempty"`             // "status": "Critical",
	Type               string              `json:"type,omitempty"`               // "type": "ServerCertificateV2",
	URI                utils.Nstring       `json:"uri,omitempty"`                // "uri": "/rest/server-certificates/e2f0031b-52bd-4223-9ac1-d91cb519d548"
}

type CertificateDetail struct {
	AliasName                string            `json:"aliasName,omitempty"`                //"aliasName":"null"
	AlternativeName          string            `json:"alternativeName,omitempty"`          //"alternativeName":"172.18.13.11"
	Base64Data               utils.Nstring     `json:"base64Data,omitempty"`               //"base64Data":"-----BEGIN CERTIFICATE-----"
	BasicConstraints         string            `json:"basicConstraints,omitempty"`         //"basicConstraints":"Subject Type=End Enti"
	Category                 string            `json:"category,omitempty"`                 //"category":"appliance"
	CertPath                 map[string]string `json:"certPath,omitempty"`                 //"certPath":"null"
	CommonName               string            `json:"commonName,omitempty"`               //"commonName":"172.18.13.11"
	ContactPerson            string            `json:"contactPerson,omitempty"`            //"contactPerson":"null"
	Country                  string            `json:"country,omitempty"`                  //"country":"IN"
	Created                  string            `json:"created,omitempty"`                  //"created":"2020-04-10T06:24:16.229Z"
	CrlDistributionEndPoints []string          `json:"crlDistributionEndPoints,omitempty"` //"crlDistributionEndPoints":"[]"
	Description              string            `json:"description,omitempty"`              //"description":"null"
	Dnqualifier              string            `json:"dnQualifier,omitempty"`              //"dnQualifier":"null"
	Etag                     string            `json:"eTag,omitempty"`                     //"eTag":"2020-04-10T06:24:16.229Z"
	Email                    string            `json:"email,omitempty"`                    //"email":"null"
	EnhancedKeyUsage         string            `json:"enhancedKeyUsage,omitempty"`         //"enhancedKeyUsage":"[1.3.6.1.5.5.7.3.1]"
	ExpiresInDays            string            `json:"expiresInDays,omitempty"`            //"expiresInDays":"36493"
	GivenName                string            `json:"givenName,omitempty"`                //"givenName":"null"
	Initials                 string            `json:"initials,omitempty"`                 //"initials":"null"
	Issuer                   string            `json:"issuer,omitempty"`                   //"issuer":"172.18.13.11"
	KeyUsage                 string            `json:"keyUsage,omitempty"`                 //"keyUsage":"digitalSignature,keyEncipherment,dataEncipherment"
	Locality                 string            `json:"locality,omitempty"`                 //"locality":"BA"
	LocationState            string            `json:"locationState,omitempty"`            //"locationState":"KA"
	Modified                 string            `json:"modified,omitempty"`                 //"modified":"2020-04-10T06:24:16.229Z"
	Name                     string            `json:"name,omitempty"`                     //"name":"null"
	Organization             string            `json:"organization,omitempty"`             //"organization":"HPE"
	OrganizationalUnit       string            `json:"organizationalUnit,omitempty"`       //"organizationalUnit":"EML"
	PublicKey                utils.Nstring     `json:"publicKey,omitempty"`                //"publicKey":"2048 bits RSA Public Key"
	SerialNumber             utils.Nstring     `json:"serialNumber,omitempty"`             //"serialNumber":"93:47:c6:5b:4d:ec:32:6e"
	Sha1Fingerprint          utils.Nstring     `json:"sha1Fingerprint,omitempty"`          //"sha1Fingerprint":"48:32:fc:d8:6a:3f:af:71:b1:fd:5c:21:"
	Sha256Fingerprint        utils.Nstring     `json:"sha256Fingerprint,omitempty"`        //"sha256Fingerprint":"0a:f6:b1:ca:94:c6:3a:75:97:74:e"
	Sha384Fingerprint        utils.Nstring     `json:"sha384Fingerprint,omitempty"`        //"sha384Fingerprint":"5e:3a:21:7c:78:79:6e:c6:5a:03:89:e1:"
	SignatureAlgorithm       string            `json:"signatureAlgorithm,omitempty"`       //"signatureAlgorithm":"SHA256WITHRSA"
	State                    string            `json:"state,omitempty"`                    //"state":"OK"
	Status                   string            `json:"status,omitempty"`                   //"status":"OK"
	Surname                  string            `json:"surname,omitempty"`                  //"surname":"null"
	Type                     string            `json:"type,omitempty"`                     //"type":"CertificateDetailV2"
	Uri                      utils.Nstring     `json:"uri,omitempty"`                      //"uri":"null"
	ValidFrom                string            `json:"validFrom,omitempty"`                //"validFrom":"2020-04-03T08:03:38.000Z"
	ValidUntil               string            `json:"validUntil,omitempty"`               //"validUntil":"2120-03-10T08:03:38.000Z"
	Version                  string            `json:"version,omitempty"`                  //"version":"3"
}
type CertificateStat struct {
	ChainStatus string `json:"chainStatus"` //"chainStatus":"VALID"
	SelfSigned  bool   `json:"selfsigned"`  // "selfsigned":true
	Trusted     bool   `json:"trusted"`     //"trusted":false

}

func (c *OVClient) GetServerCertificateByIp(ip string) (ServerCertificate, error) {
	var (
		serverC ServerCertificate
		uri     = "/rest/certificates/https/remote/" + ip
	)
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	//rest call
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return serverC, err
	}

	log.Debugf("GetServerCertificate %s", data)
	if err := json.Unmarshal([]byte(data), &serverC); err != nil {
		return serverC, err
	}
	//hardware.Client = c
	return serverC, nil

}
func (c *OVClient) GetServerCertificateByName(name string) (ServerCertificate, error) {
	var (
		serverC ServerCertificate
		uri     = "/rest/certificates/servers/" + name
	)
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	//rest call
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return serverC, err
	}

	log.Infof("GetServerCertificate %s", data)
	//data[]
	if err := json.Unmarshal([]byte(data), &serverC); err != nil {
		return serverC, err
	}
	//hardware.Client = c
	return serverC, nil

}

func (c *OVClient) CreateServerCertificate(serverC ServerCertificate) error {
	log.Infof("Initializing adding of ServerCertificate %s.", serverC.Name)
	var (
		uri = "/rest/certificates/servers/"
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, serverC)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.POST, uri, serverC)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new add ServerCertificate request: %s", err)
		return err
	}

	log.Debugf("Response New ServerCertificate %s", data)
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return err
	}

	err = t.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (c *OVClient) DeleteServerCertificate(name string) error {

	var (
		uri = "/rest/certificates/servers/" + name
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("task -> %+v", t)

	data, err := c.RestAPICall(rest.DELETE, uri, nil)
	if err != nil {
		log.Errorf("Error submitting delete server certificate request: %s", err)
		t.TaskIsDone = true
		return err
	}

	log.Debugf("Response delete server certificate %s", data)
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return err
	}
	err = t.Wait()
	if err != nil {
		return err
	}
	return nil
}

func (c *OVClient) UpdateServerCertificate(serverC ServerCertificate) error {
	log.Infof("Initializing update of server certificate for %s.", serverC.Name)
	var (
		uri = serverC.URI.String()
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, serverC)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, serverC)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update server certificate request: %s", err)
		return err
	}

	log.Debugf("Response update ServerCertificate %s", data)
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return err
	}

	err = t.Wait()
	if err != nil {
		return err
	}

	return nil
}
