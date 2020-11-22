package rscp

//  defines the expected Length of the data type or 0 for variable Length
var secretTags = []Tag{
	RSCP_AUTHENTICATION_PASSWORD,
	RSCP_REQ_SET_ENCRYPTION_PASSPHRASE,
}

// Validate validates if the Value matches the expected go type
func (t Tag) IsSecret() bool {
	for _, v := range secretTags {
		if v == t {
			return true
		}
	}
	return false
}
