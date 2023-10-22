package model

import "errors"

type os = string

const (
	WIN   os = "win"
	MACOS os = "mac"
	UNIX  os = "unix"
)

var OSMap = map[string]os{
	"WIN":   WIN,
	"MACOS": MACOS,
	"UNIX":  UNIX,
}

type platform = string

const (
	ARM64 platform = "ARM64" // for macbook with M chip

	// todo : maybe contains more details? like AMD64
	x86 platform = "x86"
)

var PlatformMap = map[string]platform{
	"ARM64": ARM64,
	"x86":   x86,
}

type configType = string

const (
	// todo : remote is just a flag, it's not used for now. Should contains more details like ZK,Nasos...
	Remote configType = "remote"
	Local  configType = "local"
)

var ConfigTypeMap = map[string]configType{
	"local": Local,
}

type RuntimeInfo struct {
	IsDebug  bool
	OS       os
	Platform platform
	CfgType  configType
}

func (r *RuntimeInfo) InitByArgs(args []string) error {
	for i := 0; i < len(args); i++ {
		if args[i] == "--debug" {
			r.IsDebug = true
		}

		if args[i] == "-os" {
			if v, ok := OSMap[args[i+1]]; !ok {
				r.OS = v
			} else {
				return errors.New("Invaild args")
			}
		}

		if args[i] == "-platform" {
			if v, ok := PlatformMap[args[i+1]]; !ok {
				r.Platform = v
			} else {
				return errors.New("Invaild args")
			}
		}

		if args[i] == "-cfg" {
			if v, ok := ConfigTypeMap[args[i+1]]; !ok {
				r.CfgType = v
			} else {
				return errors.New("Invaild args")
			}
		}

	}

	return nil
}

var DefaultInfo = RuntimeInfo{
	IsDebug:  false,
	OS:       MACOS,
	Platform: ARM64,
	CfgType:  Local,
}
