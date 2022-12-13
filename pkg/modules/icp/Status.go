package icp

type Status string

const (
	Empty    Status = ""
	Enabled  Status = "enabled"
	Disabled Status = "disabled"
	Expired  Status = "expired"
)
