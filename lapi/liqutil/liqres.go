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
type GroupResourceList struct {
	Response struct {
		Data []struct {
			Type        string `json:"type"`
			Index       int    `json:"index"`
			Name        string `json:"name"`
			Flags       string `json:"flags"`
			DeviceType  string `json:"device_type"`
			DeviceState string `json:"device_state"`
			FabricType  string `json:"fabric_type"`
			ConnType    string `json:"conn_type"`
			Lanes       int    `json:"lanes"`
			DevID       string `json:"dev_id"`
			SledID      string `json:"sled_id"`
			FabrID      int    `json:"fabr_id"`
			SwitGid     string `json:"swit_gid"`
			PortGid     string `json:"port_gid"`
			FabrGid     string `json:"fabr_gid"`
			PodID       int    `json:"pod_id"`
			Vid         string `json:"vid"`
			Did         string `json:"did"`
			Location    struct {
				Rack  int `json:"rack"`
				Shelf int `json:"shelf"`
				Node  int `json:"node"`
			} `json:"location"`
			Owner struct {
				Rack  int `json:"rack"`
				Shelf int `json:"shelf"`
				Node  int `json:"node"`
			} `json:"owner"`
			Unique     string `json:"unique,omitempty"`
			Hconn      string `json:"hconn,omitempty"`
			CapacityMB int    `json:"capacity(MB),omitempty"`
		} `json:"data"`
		Errors interface{} `json:"errors"`
		Code   int         `json:"code"`
	} `json:"response"`
}
