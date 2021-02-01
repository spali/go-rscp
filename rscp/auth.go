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
