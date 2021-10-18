package oneview

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"log"
)

// UpdateScopeUris - Method is exported to resources for updating resource's scope uris
func UpdateScopeUris(meta interface{}, val []interface{}, uri string) error {
	config := meta.(*Config)
	rawVal := []string{}
	for _, scope := range val {
		rawVal = append(rawVal, scope.(string))
	}
	scopes := ov.Scope{
		ResourceUri: uri,
		ScopeUris:   rawVal,
	}
	err := config.ovClient.UpdateScopeForResource(scopes)
	if err != nil {
		return fmt.Errorf("unable to update scopes on the resource: %s", err)
	}
	return nil
}

// contains checks if a string is present in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// readResourceUris - Method is specific to Scope Resource.
// input parameters: Scope uri , resource Uris, type of utility i.e. AddScopeUris or RemovedScopeUris
// returns: resource uris valid for addedResourceUris or removeResourceUris
func readResourceUris(meta interface{}, uri string, rawVal []interface{}, utility bool) []string {
	resourceUris := []string{}
	config := meta.(*Config)
	for _, resourceUri := range rawVal {
		scopeUris, err := config.ovClient.GetScopeFromResource(resourceUri.(string))
		if err != nil {
			log.Printf("unable to fetch scope uris from resourceUris: %s", err)
		} else {
			if utility == true {
				// adds if the scope is present in the resource - This works for AddedScopeUris
				if contains(scopeUris.ScopeUris, uri) {
					resourceUris = append(resourceUris, resourceUri.(string))
				}
			} else {
				// adds if the scope is not present in the resource - This works for RemovedScopeUris
				if !(contains(scopeUris.ScopeUris, uri)) {
					resourceUris = append(resourceUris, resourceUri.(string))
				}
			}
		}
	}
	return resourceUris
}
