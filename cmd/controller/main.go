package main

import (
	"fmt"
	"os"

    "controller"
)

// import (
//     hue "github.com/ezebunandu/gohue"
// )

// var HueIPAddress = "192.168.57.231"
// var HueID = "6Q3jiDVwElwoO2XADP-0rvB7BWTvSCuKNJQS9gth"
// func main(){
//     bridge, err := hue.NewBridge(HueIPAddress)
//     if err != nil {
//         panic(err)
//     }
//     if err = bridge.Login(HueID); err != nil {
//         panic(err)
//     }
//     light, err := bridge.GetLightByName("Lamp Stand 2")
//     if err != nil {
//         panic(err)
//     }
//     if err = light.On(); err != nil {
//         panic(err)
//     }
// }

var HueIPAddress = ""
var HueID = ""
func main() {
    ctl, err := controller.New(HueID, HueIPAddress)
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