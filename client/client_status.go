package client

type Status int8

// String converter function for ClientStatus
func (status Status) String() string {
	return StatusMeta[status].Name
}

// all errors as constant
//nolint: golint,stylecheck
const (
	STATUS_OK                        Status = 0
	STATUS_ERR_INVALID_INPUT         Status = -1
	STATUS_ERR_NO_MEMORY             Status = -2
	STATUS_ERR_INVALID_MAGIC         Status = -3
	STATUS_ERR_PROT_VERSION_MISMATCH Status = -4
	STATUS_ERR_INVALID_FRAME_LENGTH  Status = -5
	STATUS_ERR_INVALID_CRC           Status = -6
	STATUS_ERR_DATA_LIMIT_EXCEEDED   Status = -7
)

type StatusMetaType struct {
	Name string
}

// TODO: implement as go errors

// StatusMeta contains meta informations about each error
var StatusMeta = map[Status]StatusMetaType{
	STATUS_OK:                {"STATUS_OK"},
	STATUS_ERR_INVALID_INPUT: {"STATUS_ERR_INVALID_INPUT"},
	STATUS_ERR_NO_MEMORY:     {"STATUS_ERR_NO_MEMORY"},
	//STATUS_ERR_INVALID_MAGIC:         {"STATUS_ERR_INVALID_MAGIC"},
	//STATUS_ERR_PROT_VERSION_MISMATCH: {"STATUS_ERR_PROT_VERSION_MISMATCH"},
	//STATUS_ERR_INVALID_FRAME_LENGTH: {"STATUS_ERR_INVALID_FRAME_LENGTH"},
	//STATUS_ERR_INVALID_CRC:          {"STATUS_ERR_INVALID_CRC"},
	//STATUS_ERR_DATA_LIMIT_EXCEEDED: {"STATUS_ERR_DATA_LIMIT_EXCEEDED"},
}
