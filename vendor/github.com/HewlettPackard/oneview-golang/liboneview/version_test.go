package liboneview

import (
	"testing"

	// "github.com/docker/machine/drivers/oneview/icsp"
	// "github.com/docker/machine/drivers/oneview/ov"
	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
)

// TestCalculateVersion
func TestCalculateVersion(t *testing.T) {
	var v Version
	v = v.CalculateVersion(120, 108)
	log.Debugf("v => %+v, %d", v, v.Integer())
	assert.True(t, API_VER1.EqualV(v))  // should be ver1
	assert.False(t, API_VER2.EqualV(v)) // should not be ver2

	v = v.CalculateVersion(300, 108)
	log.Debugf("v => %+v, %d", v, v.Integer())
	assert.False(t, API_VER1.EqualV(v))       // should not be ver1
	assert.False(t, API_VER2.EqualV(v))       // should not be ver2
	assert.True(t, API_VER_UNKNOWN.EqualV(v)) // should be unkown

	v = v.CalculateVersion(200, 108)
	assert.False(t, API_VER1.EqualV(v)) // should not be ver1
	assert.True(t, API_VER2.EqualV(v))  // should be ver2

	v = v.CalculateVersion(201, 108)
	assert.False(t, API_VER1.EqualV(v))       // should not be ver1
	assert.True(t, API_VER_UNKNOWN.EqualV(v)) // should unkown
}

// Test GetAPIVersion
func TestCheckVersion(t *testing.T) {
	// TODO: removing refrences to icsp/ov libs because we get this error:
	// 	# github.com/HewlettPackard/oneview-golang/liboneview
	// import cycle not allowed in test
	// package github.com/HewlettPackard/oneview-golang/liboneview (test)
	// 	imports github.com/docker/machine/drivers/oneview/ov
	// 	imports github.com/HewlettPackard/oneview-golang/liboneview
	//
	// FAIL	github.com/HewlettPackard/oneview-golang/liboneview [setup failed]

	// var (
	// 	ClientICSP *icsp.ICSPClient
	// 	ClientOV   *ov.OVClient
	// )
	// if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
	//
	// 	/*Client := &icsp.ICSPClient{
	// 		rest.Client{
	// 			User:     os.Getenv("ONEVIEW_ICSP_USER"),
	// 			Password: os.Getenv("ONEVIEW_ICSP_PASSWORD"),
	// 			Domain:   os.Getenv("ONEVIEW_ICSP_DOMAIN"),
	// 			Endpoint: os.Getenv("ONEVIEW_ICSP_ENDPOINT"),
	// 			// ConfigDir:
	// 			SSLVerify:  false,
	// 			APIVersion: 108,
	// 			APIKey:     "none",
	// 		},
	// 	}*/
	// 	log.Debugf("%+v", os.Getenv("ONEVIEW_ICSP_USER"))
	// 	log.Debugf("%+v", os.Getenv("ONEVIEW_ICSP_PASSWORD"))
	// 	log.Debugf("%+v", os.Getenv("ONEVIEW_ICSP_ENDPOINT"))
	// 	c := ClientICSP.NewICSPClient(
	// 		os.Getenv("ONEVIEW_ICSP_USER"),
	// 		os.Getenv("ONEVIEW_ICSP_PASSWORD"),
	// 		os.Getenv("ONEVIEW_ICSP_DOMAIN"),
	// 		os.Getenv("ONEVIEW_ICSP_ENDPOINT"),
	// 		false,
	// 		108)
	// 	icspVer, _ := c.GetAPIVersion()
	// 	log.Debugf("%d", icspVer.CurrentVersion)
	//
	// 	ovc := ClientOV.NewOVClient(
	// 		os.Getenv("ONEVIEW__OV_USER"),
	// 		os.Getenv("ONEVIEW_OV_PASSWORD"),
	// 		os.Getenv("ONEVIEW_OV_DOMAIN"),
	// 		os.Getenv("ONEVIEW_OV_ENDPOINT"),
	// 		false,
	// 		120)
	// 	ovVer, _ := ovc.GetAPIVersion()
	// 	log.Debugf("%d", ovVer.CurrentVersion)
	// 	verCombo := ovVer.CurrentVersion + icspVer.CurrentVersion
	// 	assert.True(t, IsVersionValid(verCombo))
	// 	// fmt.Printf("after GetAPIVersion: %s -> (err) %s", data.CurrentVersion, err)
	// 	// assert.Error(t,err, fmt.Sprintf("Error caught as expected: %s",err))
	// 	//assert.NoError(t, err, "GetAPIVersion threw error -> %s", err)
	// 	//assert.Equal(t, true, d.Tc.EqualFaceI(d.Tc.GetExpectsData(d.Env, "CurrentVersion"), data.CurrentVersion))
	// 	//assert.Equal(t, true, d.Tc.EqualFaceI(d.Tc.GetExpectsData(d.Env, "MinimumVersion"), data.MinimumVersion))
	//
	// } else {
	// quickie unit test
	isValid := IsVersionValid(400)
	assert.False(t, isValid)
	// ov,120 + icsp,108
	isValid = IsVersionValid(228)
	assert.True(t, isValid)
	// }

}
