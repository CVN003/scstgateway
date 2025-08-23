package core

const (
	SCST_KERNEL_PATH  = "/sys/kernel/scst_tgt/"
	HandlerPath       = SCST_KERNEL_PATH + "handlers/"
	TargetPath        = SCST_KERNEL_PATH + "targets/"
	ISCSI_Target      = TargetPath + "iscsi/"
	QLA2x00T_Target   = TargetPath + "qla2x00t/"
	DEFAULT_SAVE_PATH = "/usr/local/scstgateway/etc/"
)

// SCST_Device represents a virtual lun in the SCST
type SVD struct {
	USN               string `json:"usn"`      // Unique Serial Number
	VID               string `json:"vid"`      // Vendor ID
	PID               string `json:"pid"`      // Product ID
	LunName           string `json:"lun_name"` // LUN Name
	LunDeviceFullPath string `json:"lun_device_full_path"`
	Size              int64  `json:"size"`    // LUN Size in bytes
	Handler           string `json:"handler"` // vdisk_blockio（default)|| vdisk_fileio||vdisk_nullio
	mgmt              string
}

type Lun struct {
	LunID             int    `json:"lun_id"`   // lun id 0 is always dummy device
	LunName           string `json:"lun_name"` // LUN Name
	LunDeviceFullPath string `json:"lun_device_full_path"`
}

// Initiator Group
type IniGroup struct {
	GrpName string   `json:"group_name"`
	Lun     []Lun    `json:"luns"`
	Ini     []string `json:"initiators"`
}

type Target struct {
	Name       string     `json:"name"`
	TargetType string     `json:"target_type"` // iscsi(default) || qla2x00t
	IniGroup   []IniGroup `json:"ini_grps"`
}

// SCSTMapping represents the configuration of the SCST  in memory
type SCSTMapping struct {
	Devs    []SVD    `json:"devs"`
	Targets []Target `json:"targets"`
}
