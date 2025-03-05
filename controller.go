package controller

import (
	"errors"

	hue "github.com/ezebunandu/gohue"
)

type Controller struct {
    bridge *hue.Bridge
}

func New(HueID, HueIPAddress string) (*Controller, error) {
    if HueID == "" || HueIPAddress == ""{
        return nil, errors.New("HueID and HueIPAddress cannot be empty")
    }
    bridge, err := hue.NewBridge(HueIPAddress)
    if err != nil {
        return nil, err
    }
    err = bridge.Login(HueID)
    if err != nil {
        return nil, err
    }
    return &Controller{bridge: bridge}, nil
}

func (ctl *Controller) EnsureOn(lightname string) error {
    light, err := ctl.bridge.GetLightByName(lightname)
    if err != nil {
        return err
    }
    err = light.On()
    if err != nil {
        return err
    }
    return nil
}

func (ctl *Controller) PingBridge() error {
    if err := ctl.bridge.GetInfo(); err != nil {
        return err
    }
    return nil
}