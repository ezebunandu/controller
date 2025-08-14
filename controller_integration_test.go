///go:build integration

package controller_test

import (
	"controller"
	"os"
	"testing"
)

var HueID = os.Getenv("HUE_ID")
var HueIPAddress = os.Getenv("HUE_IP_ADDRESS")

func TestNew_ReturnsLoggedInController(t *testing.T){
    t.Parallel()
    if HueID == "" || HueIPAddress == "" {
        t.Fatal("env variables HUE_ID and HUE_IP_ADDRESS must be set")
    }
    ctl, err := controller.NewController(HueID, HueIPAddress)
    if err != nil {
        t.Fatal(err)
    }
    err = ctl.PingBridge()
    if err != nil {
        t.Fatal(err)
    }
}