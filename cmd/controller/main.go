package main

import (
	"fmt"
	"os"

    "github.com/ezebunandu/controller"
)

func main() {
    HueID, HueIPAddress := os.Getenv("HUE_ID"), os.Getenv("HUE_IP_ADDRESS")
    ctl, err := controller.NewController(HueID, HueIPAddress)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    err = ctl.EnsureOn("Lamp Stand 2")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}