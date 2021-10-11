package oneview

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
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
