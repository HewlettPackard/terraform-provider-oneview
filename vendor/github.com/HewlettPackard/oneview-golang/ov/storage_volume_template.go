package ov

import (
	"encoding/json"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

// The Marshal() will omit the bool attibutes below if they are false.
// Please remove the omitempty option and use it as and when required.

type StorageVolumeTemplate struct {
	Category                    string              `json:"category,omitempty"`
	Created                     string              `json:"created,omitempty"`
	CompatibleStorageSystemsUri utils.Nstring       `json:"compatibleStorageSystemsUri,omitempty"`
	Description                 utils.Nstring       `json:"description,omitempty"`
	ETAG                        string              `json:"eTag,omitempty"`
	IsRoot                      bool                `json:"isRoot,omitempty"`
	ScopesURI                   utils.Nstring       `json:"scopesUri,omitempty"`
	Name                        string              `json:"name,omitempty"`
	State                       string              `json:"state,omitempty"`
	Status                      string              `json:"status,omitempty"`
	Type                        string              `json:"type,omitempty"`
	URI                         utils.Nstring       `json:"uri,omitempty"`
	InitialScopeUris            []utils.Nstring     `json:"initialScopeUris,omitempty"`
	RootTemplateUri             utils.Nstring       `json:"rootTemplateUri,omitempty"`
	StoragePoolUri              utils.Nstring       `json:"storagePoolUri,omitempty"`
	Version                     string              `json:"version,omitempty"`
	Uuid                        string              `json:"uuid,omitempty"`
	Family                      string              `json:"family,omitempty"`
	TemplateProperties          *TemplateProperties `json:"properties,omitempty"`
}

type TemplateProperties struct {
	Name                          *TemplatePropertyDatatypeStructString `json:"name,omitempty"`
	StoragePool                   *TemplatePropertyDatatypeStructString `json:"storagePool,omitempty"`
	Size                          *TemplatePropertyDatatypeStructInt    `json:"size,omitempty"`
	ProvisioningType              *TemplatePropertyDatatypeStructString `json:"provisioningType,omitempty"`
	SnapshotPool                  *TemplatePropertyDatatypeStructString `json:"snapshotPool,omitempty"`
	DataTransferLimit             *TemplatePropertyDatatypeStructInt    `json:"dataTransferLimit,omitempty"`
	IsDeduplicated                *TemplatePropertyDatatypeStructBool   `json:"isDeduplicated,omitempty"`
	IsEncrypted                   *TemplatePropertyDatatypeStructBool   `json:"isEncrypted,omitempty"`
	IsPinned                      *TemplatePropertyDatatypeStructBool   `json:"isPinned,omitempty"`
	IopsLimit                     *TemplatePropertyDatatypeStructInt    `json:"iopsLimit,omitempty"`
	Folder                        *TemplatePropertyDatatypeStructString `json:"folder,omitempty"`
	TemplateVersion               *TemplatePropertyDatatypeStructString `json:"templateVersion,omitempty"`
	PerformancePolicy             *TemplatePropertyDatatypeStructString `json:"performancePolicy,omitempty"`
	VolumetSet                    *TemplatePropertyDatatypeStructString `json:"volumeSet,omitempty"`
	Description                   *TemplatePropertyDatatypeStructString `json:"description,omitempty"`
	IsAdaptiveOptimizationEnabled *TemplatePropertyDatatypeStructBool   `json:"isAdaptiveOptimizationEnabled,omitempty"`
	IsCompressed                  *TemplatePropertyDatatypeStructBool   `json:"isCompressed,omitempty"`
	DataProtectionLevel           *TemplatePropertyDatatypeStructString `json:"dataProtectionLevel,omitempty"`
	IsShareable                   *TemplatePropertyDatatypeStructBool   `json:"isShareable,omitempty"`
}

// Struct for properties whose default value is a string
type TemplatePropertyDatatypeStructString struct {
	Meta        *Meta         `json:"meta,omitempty"`
	Type        string        `json:"type,omitempty"`
	Title       string        `json:"title,omitempty"`
	Required    bool          `json:"required"`
	Maxlength   int           `json:"maxLength,omitempty"`
	Minlength   int           `json:"minLength,omitempty"`
	Description utils.Nstring `json:"description"`
	Enum        []string      `json:"enum,omitempty"`
	Default     string        `json:"default,omitempty"`
	Minimum     int           `json:"minimum,omitempty"`
	Format      string        `json:"format,omitempty"`
}

// Struct for properties whose default value is an int
type TemplatePropertyDatatypeStructInt struct {
	Meta        *Meta         `json:"meta,omitempty"`
	Type        string        `json:"type,omitempty"`
	Title       string        `json:"title,omitempty"`
	Required    bool          `json:"required"`
	Description utils.Nstring `json:"description"`
	Enum        []string      `json:"enum,omitempty"`
	Default     int           `json:"default"`
	Maximum     int           `json:"maximum,omitempty"`
	Minimum     int           `json:"minimum,omitempty"`
	Format      string        `json:"format,omitempty"`
}

// Struct for properties whose default value is a bool
type TemplatePropertyDatatypeStructBool struct {
	Meta        *Meta         `json:"meta,omitempty"`
	Type        string        `json:"type,omitempty"`
	Title       string        `json:"title,omitempty"`
	Required    bool          `json:"required"`
	Description utils.Nstring `json:"description"`
	Enum        []string      `json:"enum,omitempty"`
	Default     bool          `json:"default"`
	Format      string        `json:"format,omitempty"`
}

type Meta struct {
	Locked       bool   `json:"locked"`
	SemanticType string `json:"semanticType,omitempty"`
	CreateOnly   bool   `json:"createOnly,omitempty"`
}

type StorageVolumeTemplateList struct {
	Total       int                     `json:"total,omitempty"`       // "total": 1,
	Count       int                     `json:"count,omitempty"`       // "count": 1,
	Start       int                     `json:"start,omitempty"`       // "start": 0,
	PrevPageURI utils.Nstring           `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI utils.Nstring           `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         utils.Nstring           `json:"uri,omitempty"`         // "uri": "/rest/server-profiles?filter=connectionTemplateUri%20matches%7769cae0-b680-435b-9b87-9b864c81657fsort=name:asc"
	Members     []StorageVolumeTemplate `json:"members,omitempty"`     // "members":[]
}

func (c *OVClient) GetStorageVolumeTemplateByName(name string) (StorageVolumeTemplate, error) {
	var (
		sVolTemplate StorageVolumeTemplate
	)
	sVolTemplates, err := c.GetStorageVolumeTemplates(fmt.Sprintf("name matches '%s'", name), "name:asc", "", "")
	if sVolTemplates.Total > 0 {
		return sVolTemplates.Members[0], err
	} else {
		return sVolTemplate, err
	}
}

func (c *OVClient) GetStorageVolumeTemplates(filter string, sort string, start string, count string) (StorageVolumeTemplateList, error) {
	var (
		uri           = "/rest/storage-volume-templates"
		q             map[string]interface{}
		sVolTemplates StorageVolumeTemplateList
	)
	q = make(map[string]interface{})
	if len(filter) > 0 {
		q["filter"] = filter
	}

	if sort != "" {
		q["sort"] = sort
	}

	if start != "" {
		q["start"] = start
	}

	if count != "" {
		q["count"] = count
	}

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	// Setup query
	if len(q) > 0 {
		c.SetQueryString(q)
	}

	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return sVolTemplates, err
	}

	log.Debugf("GetStorageVolumeTemplates %s", data)
	if err := json.Unmarshal([]byte(data), &sVolTemplates); err != nil {
		return sVolTemplates, err
	}
	return sVolTemplates, nil
}

func (c *OVClient) CreateStorageVolumeTemplate(sVolTemplate StorageVolumeTemplate) error {
	log.Infof("Initializing creation of storage volume for %s.", sVolTemplate.Name)
	var (
		uri = "/rest/storage-volume-templates"
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	log.Debugf("REST : %s \n %+v\n", uri, sVolTemplate)

	data, err := c.RestAPICall(rest.POST, uri, sVolTemplate)

	if err != nil {
		log.Errorf("Error submitting new storage volume template request: %s", err)
		return err
	}

	log.Debugf("Response New StorageVolumeTemplate %s", data)

	return nil
}

func (c *OVClient) DeleteStorageVolumeTemplate(name string) error {
	var (
		sVolTemplate StorageVolumeTemplate
		err          error
		uri          string
	)

	sVolTemplate, err = c.GetStorageVolumeTemplateByName(name)
	if err != nil {
		return err
	}
	if sVolTemplate.Name != "" {
		log.Debugf("REST : %s \n %+v\n", sVolTemplate.URI, sVolTemplate)
		uri = sVolTemplate.URI.String()
		if uri == "" {
			log.Warn("Unable to post delete, no uri found.")
			return err
		}
		_, err := c.RestAPICall(rest.DELETE, uri, nil)
		if err != nil {
			log.Errorf("Error submitting delete storage volume template request: %s", err)
			return err
		}

		return nil
	} else {
		log.Infof("StorageVolumeTemplate could not be found to delete, %s, skipping delete ...", name)
	}
	return nil
}

func (c *OVClient) UpdateStorageVolumeTemplate(sVolTemplate StorageVolumeTemplate) error {
	log.Infof("Initializing update of storage volume to %s.", sVolTemplate.Name)
	var (
		uri = sVolTemplate.URI.String()
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	log.Debugf("REST : %s \n %+v\n", uri, sVolTemplate)
	_, err := c.RestAPICall(rest.PUT, uri, sVolTemplate)
	if err != nil {
		log.Errorf("Error submitting update StorageVolumeTemplate request: %s", err)
		return err
	}
	return nil
}
