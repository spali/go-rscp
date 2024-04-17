package rscp

type ClientStatus int8

// String converter function for ClientStatus
func (status ClientStatus) String() string {
	return statusMeta[status].Name
}

// all errors as constant
//
//nolint:revive,stylecheck
const (
	STATUS_OK                        ClientStatus = 0
	STATUS_ERR_INVALID_INPUT         ClientStatus = -1
	STATUS_ERR_NO_MEMORY             ClientStatus = -2
	STATUS_ERR_INVALID_MAGIC         ClientStatus = -3
	STATUS_ERR_PROT_VERSION_MISMATCH ClientStatus = -4
	STATUS_ERR_INVALID_FRAME_LENGTH  ClientStatus = -5
	STATUS_ERR_INVALID_CRC           ClientStatus = -6
	STATUS_ERR_DATA_LIMIT_EXCEEDED   ClientStatus = -7
)

type statusMetaType struct {
	Name string
}

// TODO: implement as go errors

// statusMeta contains meta informations about each error
var statusMeta = map[ClientStatus]statusMetaType{
	STATUS_OK:                {"STATUS_OK"},
	STATUS_ERR_INVALID_INPUT: {"STATUS_ERR_INVALID_INPUT"},
	STATUS_ERR_NO_MEMORY:     {"STATUS_ERR_NO_MEMORY"},
	//STATUS_ERR_INVALID_MAGIC:         {"STATUS_ERR_INVALID_MAGIC"},
	//STATUS_ERR_PROT_VERSION_MISMATCH: {"STATUS_ERR_PROT_VERSION_MISMATCH"},
	//STATUS_ERR_INVALID_FRAME_LENGTH: {"STATUS_ERR_INVALID_FRAME_LENGTH"},
	//STATUS_ERR_INVALID_CRC:          {"STATUS_ERR_INVALID_CRC"},
	//STATUS_ERR_DATA_LIMIT_EXCEEDED: {"STATUS_ERR_DATA_LIMIT_EXCEEDED"},
}
