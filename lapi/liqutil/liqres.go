package liqutil

import "time"

var ApiPath = ":8080/liqid/api/v2/"

type DeviceCount struct {
	Response struct {
		Data []struct {
			CompCnt int `json:"comp_cnt"`
			TargCnt int `json:"targ_cnt"`
			LinkCnt int `json:"link_cnt"`
			GpuCnt  int `json:"gpu_cnt"`
			PlxCnt  int `json:"plx_cnt"`
			FpgaCnt int `json:"fpga_cnt"`
		} `json:"data"`
		Errors interface{} `json:"errors"`
		Code   int         `json:"code"`
	} `json:"response"`
}

type FabricID struct {
	Response struct {
		Data   []int       `json:"data"`
		Errors interface{} `json:"errors"`
		Code   int         `json:"code"`
	} `json:"response"`
}

//Struct for the 2.5 Version information
type Version struct {
	Response struct {
		Data []struct {
			Component      string    `json:"component"`
			Branch         string    `json:"branch"`
			Changeset      string    `json:"changeset"`
			Date           time.Time `json:"date"`
			Version        string    `json:"version"`
			ChangesetShort string    `json:"changeset_short"`
			DateShort      string    `json:"date_short"`
		} `json:"data"`
		Errors interface{} `json:"errors"`
		Code   int         `json:"code"`
	} `json:"response"`
}

//Basic list of groups in the system for a given fabric ID

type GroupList struct {
	Response struct {
		Data []struct {
			GrpID     int    `json:"grp_id"`
			FabrID    int    `json:"fabr_id"`
			PodID     int    `json:"pod_id"`
			GroupName string `json:"group_name"`
		} `json:"data"`
		Errors interface{} `json:"errors"`
		Code   int         `json:"code"`
	} `json:"response"`
}

//Lists the resources inside each group
type GroupResource struct {
	Response struct {
		Data []struct {
			GrpID               int     `json:"grp_id"`
			GroupName           string  `json:"group_name"`
			CPUFrequency        float64 `json:"cpu-frequency"`
			CPUCount            int     `json:"cpu-count"`
			CPULanes            int     `json:"cpu-lanes"`
			CPUCoreCount        int     `json:"cpu-core-count"`
			TotalDram           int     `json:"total-dram"`
			NetworkAdapterCount int     `json:"network-adapter-count"`
			TotalThroughput     string  `json:"total-throughput"`
			StorageDriveCount   int     `json:"storage-drive-count"`
			TotalCapacity       int     `json:"total-capacity"`
			GpuCount            int     `json:"gpu-count"`
			GpuCores            int     `json:"gpu-cores"`
			TotalMachines       int     `json:"total-machines"`
		} `json:"data"`
		Errors interface{} `json:"errors"`
		Code   int         `json:"code"`
	} `json:"response"`
}

//List of machine resources called from liqid/api/v2/machine/details/1
type MachineResources struct {
	Response struct {
		Data []struct {
			CPUSocketsField     string      `json:"cpuSocketsField"`
			MachID              int         `json:"mach_id"`
			MachName            string      `json:"mach_name"`
			CPUThreads          int         `json:"cpu-threads"`
			CPUFrequency        string      `json:"cpu-frequency"`
			CPUCores            int         `json:"cpu-cores"`
			CPUSockets          string      `json:"cpu-sockets"`
			DramMemory          string      `json:"dram-memory"`
			FabricConnect       interface{} `json:"fabric-connect"`
			NetworkAdapterCount int         `json:"network-adapter-count"`
			TotalThroughput     string      `json:"total-throughput"`
			StorageDriveCount   int         `json:"storage-drive-count"`
			TotalCapacity       int         `json:"total-capacity"`
			GpuCount            int         `json:"gpu-count"`
			GpuCores            int         `json:"gpu-cores"`
			OsName              string      `json:"os_name"`
			BootImage           interface{} `json:"boot_image"`
			BootDevice          string      `json:"boot_device"`
			IPAddress           interface{} `json:"ip_address"`
			IpmiAddress         string      `json:"ipmi_address"`
			FpgaCount           int         `json:"fpga-count"`
			FpgaSpeed           string      `json:"fpga-speed"`
			FpgaDramSize        int         `json:"fpga-dram-size"`
			Created             int         `json:"created"`
			Modified            int         `json:"modified"`
		} `json:"data"`
		Errors interface{} `json:"errors"`
		Code   int         `json:"code"`
	} `json:"response"`
}

//Populate this in order to create a new group
type CreateGrpStruct struct {
	FabID   int    `json:"fabr_id"`
	GrpName string `json:"group_name"`
	GrpID   int    `json:"grp_id"`
	PodID   int    `json:"pod_id"`
}

//Populate this in order to create a new machine
type CreateMachStruct struct {
	CompName          string `json:"comp_name"` //pcpu
	ConnectionHistory []struct {
		Atime    int    `json:"atime"`
		DevType  string `json:"dev_type"`
		Dtime    int    `json:"dtime"`
		FabrGid  string `json:"fabr_gid"`
		Free     bool   `json:"free"`
		Name     string `json:"name"`
		OwnerGid string `json:"owner_gid"`
		Udesc    string `json:"udesc"`
	} `json:"connection_history"`
	FabrGid  string `json:"fabr_gid"`
	FabrID   int    `json:"fabr_id"` //fabric ID in decimal
	GrpID    int    `json:"grp_id"`  //group id in decimal
	Index    int    `json:"index"`
	MachID   int    `json:"mach_id"`   //next machine ID
	MachName string `json:"mach_name"` //name of the new machine
	Mtime    int    `json:"mtime"`
	P2P      string `json:"p2p"` //may want to set this to off
	PortGid  string `json:"port_gid"`
	SwitGid  string `json:"swit_gid"`
}

//Used to return the next group or machine ID
type NextMacID struct {
	Response struct {
		Code   int      `json:"code"`
		Data   []string `json:"data"`
		Errors []struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Type    string `json:"type"`
		} `json:"errors"`
	} `json:"response"`
}

type NextGrpID struct {
	Response struct {
		Data   []int       `json:"data"`
		Errors interface{} `json:"errors"`
		Code   int         `json:"code"`
	} `json:"response"`
}
