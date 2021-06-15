package ov

import (
	"encoding/json"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type Ipv4Subnet struct {
	AllocatorUri        utils.Nstring          `json:"allocatorUri,omitempty"`
	AssociatedResources []AssociatedResSubnets `json:"associatedResources,omitempty"`
	Category            string                 `json:"category,omitempty"`
	CollectorUri        utils.Nstring          `json:"collectorUri"`
	Created             string                 `json:"created,omitempty"`
	DnsServers          []utils.Nstring        `json:"dnsServers,omitempty"`
	Domain              string                 `json:"domain,omitempty"`
	ETAG                string                 `json:"eTag,omitempty"`
	Gateway             string                 `json:"gateway"`
	Modified            string                 `json:"modified,omitempty"`
	Name                string                 `json:"name,omitempty"`
	NetworkId           string                 `json:"networkId"`
	RangeUris           []utils.Nstring        `json:"rangeUris,omitempty"`
	SubnetMask          string                 `json:"subnetmask"`
	Type                string                 `json:"type,omitempty"`
	URI                 utils.Nstring          `json:"uri,omitempty"`
}

type AssociatedResSubnets struct {
	AssociationType  string        `json:"associationType,omitempty"`
	ResourceCategory string        `json:"resourceCategory,omitempty"`
	ResourceName     string        `json:"resourceName,omitempty"`
	ResourceUri      utils.Nstring `json:"resourceUri,omitempty"`
}

type SubnetList struct {
	Total       int           `json:"total,omitempty"`       // "total": 1,
	Count       int           `json:"count,omitempty"`       // "count": 1,
	Start       int           `json:"start,omitempty"`       // "start": 0,
	PrevPageURI utils.Nstring `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI utils.Nstring `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         utils.Nstring `json:"uri,omitempty"`         // "uri": "/rest/id-pools/ipv4/subnets?filter=networkId='10.10'"
	Members     []Ipv4Subnet  `json:"members,omitempty"`     // "members":[]
}

type SubnetAllocatorList struct {
	Count  int             `json:"count,omitempty"`
	ETAG   string          `json:"eTag,omitempty"`
	Valid  bool            `json:"valid,omitempty"`
	IdList []utils.Nstring `json:"idList,omitempty"`
}

type SubnetCollectorList struct {
	ETAG   string          `json:"eTag,omitempty"`
	IdList []utils.Nstring `json:"idList,omitempty"`
}

func (c *OVClient) GetIPv4SubnetbyId(id string) (Ipv4Subnet, error) {
	var (
		uri        = "/rest/id-pools/ipv4/subnets/"
		ipv4Subnet Ipv4Subnet
	)

	uri = uri + id
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return ipv4Subnet, err
	}

	log.Debugf("Get Ipv4 Subnets %s", data)
	if err := json.Unmarshal([]byte(data), &ipv4Subnet); err != nil {
		return ipv4Subnet, err
	}
	return ipv4Subnet, nil
}

func (c *OVClient) GetSubnetByNetworkId(nwId string) (Ipv4Subnet, error) {
	var (
		subnet Ipv4Subnet
	)
	subnets, err := c.GetIPv4Subnets("", "", fmt.Sprintf("networkId='%s'", nwId), "")
	if subnets.Total > 0 {
		return subnets.Members[0], err
	} else {
		return subnet, err
	}
}

func (c *OVClient) GetIPv4Subnets(start string, count string, filter string, sort string) (SubnetList, error) {
	var (
		uri     = "/rest/id-pools/ipv4/subnets/"
		q       map[string]interface{}
		subnets SubnetList
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
		return subnets, err
	}

	log.Debugf("Get All Subnets", data)
	if err := json.Unmarshal([]byte(data), &subnets); err != nil {
		return subnets, err
	}
	return subnets, nil
}

func (c *OVClient) CreateIPv4Subnet(subnet Ipv4Subnet) error {
	var (
		uri = "/rest/id-pools/ipv4/subnets/"
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	log.Debugf("REST : %s \n %+v\n", uri, subnet)
	data, err := c.RestAPICall(rest.POST, uri, subnet)
	if err != nil {
		log.Errorf("Error submitting subnet creation request: %s", err)
		return err
	}

	log.Debugf("Response New ipv4 Subnet %s", data)
	if err := json.Unmarshal(data, &subnet); err != nil {
		log.Errorf("Error with task un-marshal: %s", err)
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func (c *OVClient) AllocateIpv4Subnet(id string, subnet SubnetAllocatorList) (SubnetAllocatorList, error) {
	var (
		uri             = "/rest/id-pools/ipv4/subnets/" + id + "/allocator"
		subnetAllocator SubnetAllocatorList
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	log.Debugf("REST : %s \n %+v\n", uri, subnet)
	data, err := c.RestAPICall(rest.PUT, uri, subnet)
	if err != nil {
		log.Errorf("Error submitting ipv4 allocator request: %s", err)
		return subnetAllocator, err
	}

	log.Debugf("Response of ipv4 allocator %s", data)
	if err := json.Unmarshal([]byte(data), &subnetAllocator); err != nil {
		log.Errorf("Error with task un-marshal: %s", err)
		return subnetAllocator, err
	}

	return subnetAllocator, nil
}

func (c *OVClient) CollectIpv4Subnet(id string, subnet SubnetCollectorList) (SubnetCollectorList, error) {
	var (
		uri             = "/rest/id-pools/ipv4/subnets/" + id + "/collector"
		subnetCollector SubnetCollectorList
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	log.Debugf("REST : %s \n %+v\n", uri, subnet)
	data, err := c.RestAPICall(rest.PUT, uri, subnet)
	if err != nil {
		log.Errorf("Error submitting ipv4 collection request: %s", err)
		return subnetCollector, err
	}

	log.Debugf("Response of ipv4 Subnet Ids Collector %s", data)
	if err := json.Unmarshal([]byte(data), &subnetCollector); err != nil {
		log.Errorf("Error with task un-marshal: %s", err)
		return subnetCollector, err
	}

	return subnetCollector, nil
}

func (c *OVClient) DeleteIpv4Subnet(id string) error {
	var (
		subnet Ipv4Subnet
		err    error
	)

	subnet, err = c.GetIPv4SubnetbyId(id)
	if err != nil {
		return err
	}
	if subnet.NetworkId != "" {
		log.Infof("URI:%s", subnet.URI)
		log.Debugf("REST : %s \n %+v\n", subnet.URI, subnet)

		if subnet.URI == "" {
			log.Warn("Unable to post delete, no uri found.")
			return err
		}
		_, err := c.RestAPICall(rest.DELETE, subnet.URI.String(), nil)
		if err != nil {
			log.Errorf("Error submitting subnet delete request: %s", err)
			return err
		}

		return nil
	} else {
		log.Infof("ipv4 Subnet could not be found to delete, %s, skipping delete ...", subnet.NetworkId)
	}
	return nil
}

func (c *OVClient) UpdateIpv4Subnet(id string, subnet Ipv4Subnet) error {
	var (
		uri = "/rest/id-pools/ipv4/subnets/" + id
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	log.Debugf("REST : %s \n %+v\n", uri, subnet)
	data, err := c.RestAPICall(rest.PUT, uri, subnet)
	if err != nil {
		log.Errorf("Error submitting update ipv4 Subnet request: %s", err)
		return err
	}

	log.Debugf("Response update ipv4 Subnet %s", data)
	if err := json.Unmarshal([]byte(data), &subnet); err != nil {
		log.Errorf("Error with task un-marshal: %s", err)
		return err
	}

	return nil
}
