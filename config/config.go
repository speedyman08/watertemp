package config

import "flag"

var Debug = flag.Bool("dbg", false, "[-dbg]")
var ResourceIP = flag.String("ip", "10.50.0.116", "[255.255.255.255]")
