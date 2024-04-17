package rscp

type AuthLevel uint8

// all auth levels as constant
//
//nolint:revive,stylecheck
//go:generate go run github.com/alvaroloes/enumer -type=AuthLevel -json
const (
	AUTH_LEVEL_NO_AUTH   AuthLevel = 0
	AUTH_LEVEL_USER      AuthLevel = 10
	AUTH_LEVEL_INSTALLER AuthLevel = 20
	AUTH_LEVEL_SERVICE   AuthLevel = 30
	AUTH_LEVEL_ADMIN     AuthLevel = 40
	AUTH_LEVEL_E3DC      AuthLevel = 50
	AUTH_LEVEL_E3DC_ROOT AuthLevel = 60
)
