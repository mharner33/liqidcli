package liqutil

type DevCounts struct {
	Servers int64
	Ssd     int64
	Nic     int64
	Gpu     int64
	Plx     int64
	Fpga    int64
}

type Groups struct {
	GroupID   int    `json:"grp_id"`
	FabrID    int    `json:"fabr_id"`
	PodID     int    `json:"pod_id"`
	GroupName string `json:"group_name"`
}

//Struct for the 2.5 Version information
type Version25 struct {
	Response struct {
		Data []struct {
			Component      string `json:"component"`
			Branch         string `json:"branch"`
			Changeset      string `json:"changeset"`
			Date           string `json:"date"` //was time.time
			Version        string `json:"version"`
			ChangesetShort string `json:"changeset_short"`
			DateShort      string `json:"date_short"`
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
