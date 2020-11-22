package rscp

// all auth levels as constant
//nolint: golint,stylecheck
const (
	AUTH_LEVEL_NO_AUTH   uint8 = 0
	AUTH_LEVEL_USER      uint8 = 10
	AUTH_LEVEL_INSTALLER uint8 = 20
	AUTH_LEVEL_SERVICE   uint8 = 30
	AUTH_LEVEL_ADMIN     uint8 = 40
	AUTH_LEVEL_E3DC      uint8 = 50
	AUTH_LEVEL_E3DC_ROOT uint8 = 60
)

type AuthLevelMetaType struct {
	Name string
}

// AuthLevelMeta contains meta informations about each auth level
var AuthLevelMeta = map[uint8]AuthLevelMetaType{
	AUTH_LEVEL_NO_AUTH:   {"NO_AUTH"},
	AUTH_LEVEL_USER:      {"USER"},
	AUTH_LEVEL_INSTALLER: {"INSTALLER"},
	AUTH_LEVEL_SERVICE:   {"SERVICE"},
	AUTH_LEVEL_ADMIN:     {"ADMIN"},
	AUTH_LEVEL_E3DC:      {"E3DC"},
	AUTH_LEVEL_E3DC_ROOT: {"E3DC_ROOT"},
}
