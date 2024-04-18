package rscp

// Tag
type Tag uint32

//go:generate go run github.com/alvaroloes/enumer -type=Tag

//
// Contains all to us known tags.
// "undocumented" means in general it is not documented
// in the excel file from the official RscpExample.zip archive.
// most tag specific comments are from the german excel file.
//

// ------------------
// NAMESPACE: System
// 0x00xxxxxx
// -----------------
const (
	// Dieser TAG kapselt eine Authorisierungsanfrage an das S10.
	// Er enthält daher die Daten-Tags AUTHENTICATION_USER und AUTHENTICATION_PASSWORD
	RSCP_REQ_AUTHENTICATION Tag = 0x00000001
	// Benutzername innerhalb eines REQ_AUTHENTICATION
	RSCP_AUTHENTICATION_USER Tag = 0x00000002
	// Passwort innerhalb eines REQ_AUTHENTICATION
	RSCP_AUTHENTICATION_PASSWORD Tag = 0x00000003
	// Die Antwort auf einen REQ_AUTHENTICATION die den erhaltenen Level enthällt.
	// ist die Authorisierung fehlgeschlagen.
	//  NO_AUTH        -   0
	//  USER           -  10
	//  INSTALLER      -  20
	//  PARTNER        -  30
	//  E3DC           -  40
	//  E3DC_ADMIN     -  50
	//  E3DC_ROOT      -  60
	RSCP_AUTHENTICATION Tag = 0x00800001
	RSCP_REQ_USER_LEVEL Tag = 0x00000004
	RSCP_USER_LEVEL     Tag = 0x00800004
	// Setze einen Netzwerk Encryption-Passphrase
	RSCP_REQ_SET_ENCRYPTION_PASSPHRASE Tag = 0x00000005
	RSCP_SET_ENCRYPTION_PASSPHRASE     Tag = 0x00800005
	RSCP_GENERAL_ERROR                 Tag = 0x00FFFFFF

	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	RSCP_REQ_LIST_NON_PAIRED_DEVICE_TYPE_EVCS Tag = 0x00000000

	RSCP_REQ_AUTH_CHALLENGE              Tag = 0x00000006
	RSCP_AUTH_CHALLENGE_INDEX            Tag = 0x00000007
	RSCP_AUTH_CHALLENGE_DATA             Tag = 0x00000008
	RSCP_REQ_SET_PROTOCOL_VERSION        Tag = 0x00000009
	RSCP_REQ_SUPPORTED_PROTOCOL_VERSIONS Tag = 0x0000000A
	RSCP_REQ_TRIGGER_FRAME_DUMP          Tag = 0x0000000D
	RSCP_AUTHENTICATION_TYPE             Tag = 0x00000815
	RSCP_CONFIG_PROCESSED_STATE          Tag = 0x00400001

	RSCP_AUTH_CHALLENGE                                          Tag = 0x00800006
	RSCP_SET_PROTOCOL_VERSION                                    Tag = 0x00800009
	RSCP_SUPPORTED_PROTOCOL_VERSIONS                             Tag = 0x0080000A
	RSCP_TRIGGER_FRAME_DUMP                                      Tag = 0x0080000D
	RSCP_TAG_AVAHIB_RESPONSE_PLAY                                Tag = 0x0080000E
	RSCP_EMOBILITY_GET_OVERLOAD_CHARGING_CURRENT_LIMIT_PER_PHASE Tag = 0x00810021
)

// --------------
// NAMESPACE: EMS
// 0x01xxxxxx
// --------------
const (
	// PV-Leistung des S10s in W
	EMS_REQ_POWER_PV Tag = 0x01000001
	// Batterie-Leistung des S10s in W (-=entladen / +=laden)
	EMS_REQ_POWER_BAT Tag = 0x01000002
	// Hausverbrauchsleistung in W
	EMS_REQ_POWER_HOME Tag = 0x01000003
	// Leistung am Netzeinspeisepunkt in W (-=Einspeisung / +=Bezug)
	EMS_REQ_POWER_GRID Tag = 0x01000004
	// Leistung eines zusätzlich vorhandenen Einspeisers in W
	EMS_REQ_POWER_ADD Tag = 0x01000005
	// Autarkie in %
	EMS_REQ_AUTARKY Tag = 0x01000006
	// Eigenverbrauch in %
	EMS_REQ_SELF_CONSUMPTION Tag = 0x01000007
	// Batterieladezustand in %
	EMS_REQ_BAT_SOC Tag = 0x01000008
	// Abfrage des Betriebsmodus
	EMS_REQ_COUPLING_MODE                 Tag = 0x01000009
	EMS_REQ_STORED_ERRORS                 Tag = 0x0100000A
	EMS_REQ_MODE                          Tag = 0x01000011
	EMS_REQ_BALANCED_PHASES               Tag = 0x01000012
	EMS_REQ_INSTALLED_PEAK_POWER          Tag = 0x01000013
	EMS_REQ_DERATE_AT_PERCENT_VALUE       Tag = 0x01000014
	EMS_REQ_DERATE_AT_POWER_VALUE         Tag = 0x01000015
	EMS_REQ_ERROR_BUZZER_ENABLED          Tag = 0x01000016
	EMS_REQ_SET_BALANCED_PHASES           Tag = 0x01000017
	EMS_REQ_SET_INSTALLED_PEAK_POWER      Tag = 0x01000018
	EMS_REQ_SET_DERATE_PERCENT            Tag = 0x01000019
	EMS_REQ_SET_ERROR_BUZZER_ENABLED      Tag = 0x0100001A
	EMS_REQ_START_ADJUST_BATTERY_VOLTAGE  Tag = 0x0100001B
	EMS_REQ_CANCEL_ADJUST_BATTERY_VOLTAGE Tag = 0x0100001C
	EMS_REQ_ADJUST_BATTERY_VOLTAGE_STATUS Tag = 0x0100001D
	EMS_REQ_CONFIRM_ERRORS                Tag = 0x0100001E
	EMS_REQ_POWER_WB_ALL                  Tag = 0x0100001F
	EMS_REQ_POWER_WB_SOLAR                Tag = 0x01000020
	// Anfragetag ob ein zusätzlicher Leistungsmesser installiert ist, der zusäztliche Quellen misst
	EMS_REQ_EXT_SRC_AVAILABLE Tag = 0x01000021
	// PV-Leistung des S10s in W
	EMS_POWER_PV Tag = 0x01800001
	// Batterie-Leistung des S10s in W (-=entladen / +=laden)
	EMS_POWER_BAT Tag = 0x01800002
	// Hausverbrauchsleistung in W
	EMS_POWER_HOME Tag = 0x01800003
	// Leistung am Netzeinspeisepunkt in W (-=Einspeisung / +=Bezug)
	EMS_POWER_GRID Tag = 0x01800004
	// Leistung eines zusätzlich vorhandenen Einspeisers in W
	EMS_POWER_ADD Tag = 0x01800005
	// Autarkie in %
	EMS_AUTARKY Tag = 0x01800006
	// Eigenverbrauch in %
	EMS_SELF_CONSUMPTION Tag = 0x01800007
	// Batterieladezustand in %
	EMS_BAT_SOC Tag = 0x01800008
	// Betriebsmodus:
	//  0: DC
	//  1: DC-MultiWR
	//  2: AC
	//  3: HYBRID
	//  4: ISLAND
	EMS_COUPLING_MODE Tag = 0x01800009
	// Wenn das EMS im Fehlerzustand ist, wird eine Fehlermeldung übertragen!
	EMS_STORED_ERRORS Tag = 0x0180000A
	// Wenn das EMS im Fehlerzustand ist, wird eine Fehlermeldung übertragen!
	EMS_ERROR_CONTAINER Tag = 0x0180000B
	// Wenn das EMS im Fehlerzustand ist, wird eine Fehlermeldung übertragen!
	EMS_ERROR_TYPE Tag = 0x0180000C
	// Wenn das EMS im Fehlerzustand ist, wird eine Fehlermeldung übertragen!
	EMS_ERROR_SOURCE Tag = 0x0180000D
	// Wenn das EMS im Fehlerzustand ist, wird eine Fehlermeldung übertragen!
	EMS_ERROR_MESSAGE Tag = 0x0180000E
	// Wenn das EMS im Fehlerzustand ist, wird eine Fehlermeldung übertragen!
	EMS_ERROR_CODE Tag = 0x0180000F
	// Wenn das EMS im Fehlerzustand ist, wird eine Fehlermeldung übertragen!
	EMS_ERROR_TIMESTAMP               Tag = 0x01800010
	EMS_MODE                          Tag = 0x01800011
	EMS_BALANCED_PHASES               Tag = 0x01800012
	EMS_INSTALLED_PEAK_POWER          Tag = 0x01800013
	EMS_DERATE_AT_PERCENT_VALUE       Tag = 0x01800014
	EMS_DERATE_AT_POWER_VALUE         Tag = 0x01800015
	EMS_ERROR_BUZZER_ENABLED          Tag = 0x01800016
	EMS_SET_BALANCED_PHASES           Tag = 0x01800017
	EMS_SET_INSTALLED_PEAK_POWER      Tag = 0x01800018
	EMS_SET_DERATE_PERCENT            Tag = 0x01800019
	EMS_SET_ERROR_BUZZER_ENABLED      Tag = 0x0180001A
	EMS_START_ADJUST_BATTERY_VOLTAGE  Tag = 0x0180001B
	EMS_CANCEL_ADJUST_BATTERY_VOLTAGE Tag = 0x0180001C
	EMS_ADJUST_BATTERY_VOLTAGE_STATUS Tag = 0x0180001D
	EMS_CONFIRM_ERRORS                Tag = 0x0180001E
	EMS_POWER_WB_ALL                  Tag = 0x0180001F
	EMS_POWER_WB_SOLAR                Tag = 0x01800020
	EMS_EXT_SRC_AVAILABLE             Tag = 0x01800021
	// Mit diesem TAG kann in die Regelung des S10s eingegriffen werden.
	// Bei DC-Systemen ist die Ladeleistung auf die anliegende PV-Leistung beschränkt,
	// bei AC und Hybrid-Systemen kann die Ladeleistung auch größer der PV-Leistung sein.
	// Achtung: Wenn mit diesem Kommando eingegriffen wird, wird eine eventuell gesetzte Einspeisereduzierung NICHT beachtet!
	// Achtung: Das Kommando muss mindestens alle 30 Sekunden gesetzt werden, ansonsten geht das EMS in den Normalmodus.
	EMS_REQ_SET_POWER Tag = 0x01000030
	// Der Modus in den das S10 gehen soll:
	//  AUTO/NORMAL MODUS    0
	//  IDLE MODUS           1
	//  ENTLADEN MODUS       2
	//  LADEN MODUS          3
	//  NETZ_LADE MODUS      4
	EMS_REQ_SET_POWER_MODE  Tag = 0x01000031
	EMS_REQ_SET_POWER_VALUE Tag = 0x01000032
	// Die Antwort auf einen REQ_SET_POWER. Es werden die empfangenen Werte zurückgespiegelt.
	EMS_SET_POWER Tag = 0x01800030
	// Liefert den aktuellen Status des EMS.
	EMS_REQ_STATUS               Tag = 0x01000040
	EMS_STATUS                   Tag = 0x01800040
	EMS_REQ_USED_CHARGE_LIMIT    Tag = 0x01000041
	EMS_REQ_BAT_CHARGE_LIMIT     Tag = 0x01000042
	EMS_REQ_DCDC_CHARGE_LIMIT    Tag = 0x01000043
	EMS_REQ_USER_CHARGE_LIMIT    Tag = 0x01000044
	EMS_REQ_USED_DISCHARGE_LIMIT Tag = 0x01000045
	EMS_REQ_BAT_DISCHARGE_LIMIT  Tag = 0x01000046
	EMS_REQ_DCDC_DISCHARGE_LIMIT Tag = 0x01000047
	EMS_REQ_USER_DISCHARGE_LIMIT Tag = 0x01000048
	EMS_USED_CHARGE_LIMIT        Tag = 0x01800041
	EMS_BAT_CHARGE_LIMIT         Tag = 0x01800042
	EMS_DCDC_CHARGE_LIMIT        Tag = 0x01800043
	EMS_USER_CHARGE_LIMIT        Tag = 0x01800044
	EMS_USED_DISCHARGE_LIMIT     Tag = 0x01800045
	EMS_BAT_DISCHARGE_LIMIT      Tag = 0x01800046
	EMS_DCDC_DISCHARGE_LIMIT     Tag = 0x01800047
	EMS_USER_DISCHARGE_LIMIT     Tag = 0x01800048
	// Setzt einen Regelungsoffset auf den Batterieleistungssteuerung
	EMS_REQ_SET_POWER_CONTROL_OFFSET Tag = 0x01000060
	// Antwort mit dem tatsächlich gesetzten Offset
	EMS_SET_POWER_CONTROL_OFFSET       Tag = 0x01800060
	EMS_REQ_REMAINING_BAT_CHARGE_POWER Tag = 0x01000071
	// Noch mögliche Ladeleistung nach Abzug der momentanen Ladeleistung vom momentanen Limit
	EMS_REMAINING_BAT_CHARGE_POWER        Tag = 0x01800071
	EMS_REQ_REMAINING_BAT_DISCHARGE_POWER Tag = 0x01000072
	// Noch mögliche Entladeleistung nach Abzug der momentanen Entladeleistung  vom momentanen Limit
	EMS_REMAINING_BAT_DISCHARGE_POWER Tag = 0x01800072
	EMS_REQ_EMERGENCY_POWER_STATUS    Tag = 0x01000073
	// Status:
	//  NOT_POSSIBLE           = 0x00
	//  ACTIVE                 = 0x01
	//  NOT_ACTIVE             = 0x02
	//  NOT_AVAILABLE          = 0x03
	//  SWITCH_IN_ISLAND_STATE = 0x04
	EMS_EMERGENCY_POWER_STATUS Tag = 0x01800073
	// Startet oder stoppt den Notstrommodus
	//  NORMAL_GRID_MODE     = 0x00,
	//  EMERGENCY_MODE       = 0x01,
	//  ISLAND_NO_POWER_MODE = 0x02
	EMS_REQ_SET_EMERGENCY_POWER Tag = 0x01000074
	EMS_SET_EMERGENCY_POWER     Tag = 0x01800074
	// Die verfügbare Solarleistung  wird mit diesem Wert überschrieben! (Dieser Wert wird an die WallBox gesendet)
	EMS_REQ_SET_OVERRIDE_AVAILABLE_POWER Tag = 0x01000075
	EMS_SET_OVERRIDE_AVAILABLE_POWER     Tag = 0x01800075
	// Mode:
	//  1    = Modus aktiviert
	//  0    = Modus deaktiviert
	//  0xFF = Aktivierung nicht möglich (BatteryBeforeCar noch aktiv?)
	EMS_SET_BATTERY_TO_CAR_MODE Tag = 0x01800076
	// Aktiviert, deaktiviert den BatteryToCar Modus
	EMS_REQ_SET_BATTERY_TO_CAR_MODE Tag = 0x01000076
	// 1 = Modus aktiviert / 0 = Modus deaktiviert
	EMS_BATTERY_TO_CAR_MODE Tag = 0x01800077
	// Statusabfrage des BatteryToCar Modus
	EMS_REQ_BATTERY_TO_CAR_MODE Tag = 0x01000077
	// Mode:
	//  1    = Modus aktiviert
	//  0    = Modus deaktiviert
	//  0xFF = Aktivierung nicht möglich (BatteryToCar noch aktiv?)
	EMS_SET_BATTERY_BEFORE_CAR_MODE Tag = 0x01800078
	// Aktiviert, deaktiviert den BatteryBeforeCar Modus
	EMS_REQ_SET_BATTERY_BEFORE_CAR_MODE Tag = 0x01000078
	// Mode:
	//  1 = Modus aktiviert
	//  0 = Modus deaktiviert
	EMS_BATTERY_BEFORE_CAR_MODE Tag = 0x01800079
	// Statusabfrage des BatteryBeforeCar Modus
	EMS_REQ_BATTERY_BEFORE_CAR_MODE   Tag = 0x01000079
	EMS_REQ_GET_IDLE_PERIODS          Tag = 0x01000080
	EMS_GET_IDLE_PERIODS              Tag = 0x01800080
	EMS_REQ_SET_IDLE_PERIODS          Tag = 0x01000081
	EMS_SET_IDLE_PERIODS              Tag = 0x01800081
	EMS_IDLE_PERIOD                   Tag = 0x01000082
	EMS_IDLE_PERIOD_TYPE              Tag = 0x01000083
	EMS_IDLE_PERIOD_DAY               Tag = 0x01000084
	EMS_IDLE_PERIOD_START             Tag = 0x01000085
	EMS_IDLE_PERIOD_END               Tag = 0x01000086
	EMS_IDLE_PERIOD_HOUR              Tag = 0x01000087
	EMS_IDLE_PERIOD_MINUTE            Tag = 0x01000088
	EMS_IDLE_PERIOD_ACTIVE            Tag = 0x01000089
	EMS_REQ_IDLE_PERIOD_CHANGE_MARKER Tag = 0x0100008A
	EMS_IDLE_PERIOD_CHANGE_MARKER     Tag = 0x0180008A
	EMS_REQ_GET_POWER_SETTINGS        Tag = 0x0100008B
	EMS_GET_POWER_SETTINGS            Tag = 0x0180008B
	// Wird zum setzen der Power Settings verwendet. Kann folgende TAGs enthalten:
	//  POWER_LIMITS_USED
	//  MAX_CHARGE_POWER
	//  MAX_DISCHARGE_POWER
	//  MINIMUM_DISCHARGE_POWER
	//  POWERSAVE_ENABLED
	//  WEATHER_REGULATED_CHARGE_ENABLED
	EMS_REQ_SET_POWER_SETTINGS Tag = 0x0100008C
	// Enthält die Antwort auf das Setzen der PowerSettings. Gibt für jeden gesetzen Wert eine entsprechendes Element mit Rückgabecode zurück.
	//
	// Kann die Folgenden TAGS enthalten:
	//  RES_POWER_LIMITS_USED
	//  RES_MAX_CHARGE_POWER
	//  RES_MAX_DISCHARGE_POWER
	//  RES_MINIMUM_DISCHARGE_POWER
	//  RES_POWERSAVE_ENABLED
	//  RES_WEATHER_REGULATED_CHARGE_ENABLED
	EMS_SET_POWER_SETTINGS    Tag = 0x0180008C
	EMS_POWER_LIMITS_USED     Tag = 0x01000100
	EMS_RES_POWER_LIMITS_USED Tag = 0x01800100
	EMS_MAX_CHARGE_POWER      Tag = 0x01000101
	// returns:
	//   1 bei Erfolg, allerdings ist das limit unterhalb des empfohlenden Limits
	//   0 Werte erfolgreich gesetzt
	//  -1 Wert außerhalb des zulässigen Bereichs
	//  -2 setzen momentan nicht möglich, später erneut versuchen
	EMS_RES_MAX_CHARGE_POWER Tag = 0x01800101
	EMS_MAX_DISCHARGE_POWER  Tag = 0x01000102
	// returns:
	//   1 bei Erfolg, allerdings ist das limit unterhalb des empfohlenden Limits
	//   0 Werte erfolgreich gesetzt
	//  -1 Wert außerhalb des zulässigen Bereichs
	//  -2 setzen momentan nicht möglich, später erneut versuchen
	EMS_RES_MAX_DISCHARGE_POWER Tag = 0x01800102
	EMS_DISCHARGE_START_POWER   Tag = 0x01000103
	// returns:
	//   0 Werte erfolgreich gesetzt
	//  -1 Wert außerhalb des zulässigen Bereichs
	//  -2 setzen momentan nicht möglich, später erneut versuchen
	EMS_RES_DISCHARGE_START_POWER            Tag = 0x01800103
	EMS_POWERSAVE_ENABLED                    Tag = 0x01000104
	EMS_RES_POWERSAVE_ENABLED                Tag = 0x01800104
	EMS_WEATHER_REGULATED_CHARGE_ENABLED     Tag = 0x01000105
	EMS_RES_WEATHER_REGULATED_CHARGE_ENABLED Tag = 0x01800105
	EMS_WEATHER_FORECAST_MODE                Tag = 0x01000106 // undocumented response tag
	EMS_RES_WEATHER_FORECAST_MODE            Tag = 0x01800106 // undocumented response tag
	EMS_REQ_SETTINGS_CHANGE_MARKER           Tag = 0x0100008D
	EMS_SETTINGS_CHANGE_MARKER               Tag = 0x0180008D
	EMS_REQ_GET_MANUAL_CHARGE                Tag = 0x0100008E
	EMS_GET_MANUAL_CHARGE                    Tag = 0x0180008E
	EMS_MANUAL_CHARGE_START_COUNTER          Tag = 0x01000150
	EMS_MANUAL_CHARGE_ACTIVE                 Tag = 0x01000151
	EMS_MANUAL_CHARGE_ENERGY_COUNTER         Tag = 0x01000152
	EMS_MANUAL_CHARGE_LASTSTART              Tag = 0x01000153
	EMS_REQ_START_MANUAL_CHARGE              Tag = 0x0100008F
	EMS_START_MANUAL_CHARGE                  Tag = 0x0180008F
	EMS_REQ_START_EMERGENCYPOWER_TEST        Tag = 0x01000090
	// Gibt als Rückantwort die Anzahl der gestarteten Notstromtests zurück
	EMS_START_EMERGENCYPOWER_TEST Tag = 0x01800090
	EMS_REQ_GET_GENERATOR_STATE   Tag = 0x01000091
	// State:
	//  Idle = 0x00
	//  HeatUp = 0x01
	//  HeatUpDone = 0x02
	//  Starting = 0x03
	//  StartingPause = 0x04
	//  Running = 0x05
	//  Stopping = 0x06
	//  Stopped = 0x07
	//  RelaisControlMode = 0x10
	//  Kein Generator vorhanden oder Generatorinterface kommuniziert nicht = 0xFF
	EMS_GET_GENERATOR_STATE Tag = 0x01800091
	// State:
	//  0x01 - Manueller Generatorstop (falls aktuell aktiv) und aktivieren des Normalbetrieb
	//  0x02 - Manueller Generatorstart
	EMS_REQ_SET_GENERATOR_MODE Tag = 0x01000092
	// Gibt als Rückantwort
	//  Erfolgreich = 0x01
	//  Unbekannter Generatormodus = 0xFE
	//  Kein Generator vorhanden oder Generatorinterface kommuniziert nicht = 0xFF
	EMS_SET_GENERATOR_MODE             Tag = 0x01800092
	EMS_REQ_EMERGENCYPOWER_TEST_STATUS Tag = 0x01000093
	EMS_EMERGENCYPOWER_TEST_STATUS     Tag = 0x01800093
	EMS_EPTEST_NEXT_TESTSTART          Tag = 0x01000094
	EMS_EPTEST_START_COUNTER           Tag = 0x01000095
	EMS_EPTEST_RUNNING                 Tag = 0x01000096
	// undocumented request
	EMS_REQ_SYS_STATUS Tag = 0x01000098
	// undocumented response (interpretation unknown)
	EMS_SYS_STATUS        Tag = 0x0100009E
	EMS_REQ_GET_SYS_SPECS Tag = 0x01000097
	// Enthält 1 -x Untercontainer vom Typ SYS_SPEC
	EMS_GET_SYS_SPECS Tag = 0x01800098
	// Enthält die Elemente SYS_SPEC_INDEX, SYS_SPEC_NAME, SYS_SPEC_VALUE und kennzeichnet eine Systemeigenschaft
	EMS_SYS_SPEC Tag = 0x01000099
	// Der Index der Systemeigenschaft
	EMS_SYS_SPEC_INDEX Tag = 0x0100009A
	// Der Name der Systemeigenschaft
	EMS_SYS_SPEC_NAME Tag = 0x0100009B
	// Der Wert der Systemeigenschaft
	EMS_SYS_SPEC_VALUE_INT Tag = 0x0100009C
	// Der Wert der Systemeigenschaft als String
	EMS_SYS_SPEC_VALUE_STRING Tag = 0x0100009D
	// Abfrage ob das S10-EMS betriebsbereit ist.
	EMS_REQ_ALIVE     Tag = 0x01050000
	EMS_ALIVE         Tag = 0x01850000
	EMS_GENERAL_ERROR Tag = 0x01FFFFFF

	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	EMS_REQ_RESCUE_BAT_MODE            Tag = 0x01000022
	EMS_REQ_SET_RESCUE_BAT_MODE        Tag = 0x01000023
	EMS_REQ_IS_PV_DERATING             Tag = 0x01000024
	EMS_REQ_AC_POWER_LIMIT             Tag = 0x01000025
	EMS_REQ_POWER_ADD_SINKS            Tag = 0x01000026
	EMS_REQ_SET_EP_WALLBOX_ALLOW       Tag = 0x01000027
	EMS_REQ_GET_EP_WALLBOX_ALLOW       Tag = 0x01000028
	EMS_REQ_SET_MAX_EP_WALLBOX_POWER_W Tag = 0x01000029
	EMS_REQ_GET_MAX_EP_WALLBOX_POWER_W Tag = 0x0100002A
	EMS_REQ_GET_MIN_EP_WALLBOX_POWER_W Tag = 0x0100002B
	EMS_REQ_SET_MIN_EP_WALLBOX_POWER_W Tag = 0x0100002C
	EMS_REQ_SET_EP_WALLBOX_ENERGY      Tag = 0x0100002D
	EMS_REQ_GET_EP_WALLBOX_ENERGY      Tag = 0x0100002E
	EMS_REQ_SET_IDLE_PERIODS_2         Tag = 0x0100002F

	EMS_REQ_GET_IDLE_PERIODS_2      Tag = 0x01000033
	EMS_REQ_SET_IDLE_PERIODS_ENABLE Tag = 0x01000034
	EMS_REQ_GET_IDLE_PERIODS_ENABLE Tag = 0x01000035
	EMS_REQ_SET_EP_WALLBOX_PRIORITY Tag = 0x01000036
	EMS_REQ_GET_EP_WALLBOX_PRIORITY Tag = 0x01000037

	EMS_POWERSAVE_TIMEOUT Tag = 0x01000107

	EMS_REQ_REMOTE_CONTROL                       Tag = 0x01000200
	EMS_REQ_DEACTIVATE_REMOTE_CONTROL            Tag = 0x01000201
	EMS_REQ_IP_REMOTE_CONTROL                    Tag = 0x01000202
	EMS_REQ_EP_DELAY                             Tag = 0x01000203
	EMS_REQ_SET_EP_DELAY                         Tag = 0x01000204
	EMS_REQ_REMOTE_CONTROL_STATUS                Tag = 0x01000205
	EMS_REQ_IDLE_PERIOD_MIN_SOC_UCB              Tag = 0x01000206
	EMS_REQ_IDLE_PERIOD_MAX_SOC_UCB              Tag = 0x01000207
	EMS_REQ_SET_IDLE_PERIOD_MIN_SOC_UCB          Tag = 0x01000208
	EMS_REQ_SET_IDLE_PERIOD_MAX_SOC_UCB          Tag = 0x01000209
	EMS_REQ_REGULATOR_MODE                       Tag = 0x01000210
	EMS_REQ_SET_REGULATOR_MODE                   Tag = 0x01000211
	EMS_REQ_SUPPORTED_REGULATOR_MODES            Tag = 0x01000212
	EMS_REQ_EMERGENCY_POWER_OVERLOAD_STATUS      Tag = 0x01000213
	EMS_REQ_EMERGENCY_POWER_RETRY                Tag = 0x01000214
	EMS_REQ_DETECT_PHASE_OFFSET                  Tag = 0x01000217
	EMS_REQ_PHASE_DETECTION_STATUS               Tag = 0x01000218
	EMS_REQ_PHASE_OFFSET                         Tag = 0x01000219
	EMS_REQ_ABORT_PHASE_DETECTION                Tag = 0x01000220
	EMS_REQ_REGULATOR_STRATEGY                   Tag = 0x01000221
	EMS_REQ_SET_REGULATOR_STRATEGY               Tag = 0x01000222
	EMS_REQ_POWER_PV_AC_OUT                      Tag = 0x01000223
	EMS_REQ_PV_ENERGY                            Tag = 0x01000224
	EMS_REQ_ENERGY_STORAGE_MODEL                 Tag = 0x01000228
	EMS_REQ_SPECIFICATION_VALUES                 Tag = 0x01000234
	EMS_REQ_EP_RESERVE                           Tag = 0x01000242
	EMS_REQ_SEC_LIMITS                           Tag = 0x01000243
	EMS_REQ_SEC_DEVICE_STATUS                    Tag = 0x01000252
	EMS_REQ_BAT_CURRENT_IN                       Tag = 0x01000258
	EMS_REQ_BAT_CURRENT_OUT                      Tag = 0x01000259
	EMS_REQ_MAX_DC_POWER                         Tag = 0x01000260
	EMS_REQ_AC_REACTIVE_POWER                    Tag = 0x01000261
	EMS_REQ_SET_EP_PARTIAL_GRID                  Tag = 0x01000262
	EMS_REQ_GET_PARTIAL_GRID                     Tag = 0x01000263
	EMS_REQ_ESTIMATED_POWER_LIMITS               Tag = 0x01000264
	EMS_REQ_DESIGN_POWER_LIMITS                  Tag = 0x01000265
	EMS_REQ_SET_CAN_ID_FEED_IN_REDUCTION         Tag = 0x01000266
	EMS_REQ_CAN_ID_FEED_IN_REDUCTION             Tag = 0x01000267
	EMS_REQ_SET_CAN_ID_UNBALANCED_LOAD           Tag = 0x01000268
	EMS_REQ_CAN_ID_UNBALANCED_LOAD               Tag = 0x01000269
	EMS_REQ_SET_WALLBOX_MODE                     Tag = 0x01000270
	EMS_REQ_GET_WALLBOX_MODE                     Tag = 0x01000271
	EMS_REQ_SET_MAX_FUSE_POWER                   Tag = 0x01000272
	EMS_REQ_GET_MAX_FUSE_POWER                   Tag = 0x01000273
	EMS_REQ_SET_CONNECTED_POWER                  Tag = 0x01000274
	EMS_REQ_GET_CONNECTED_POWER                  Tag = 0x01000275
	EMS_REQ_DERATE_AT_CONNECTED_POWER            Tag = 0x01000276
	EMS_REQ_SET_DERATE_AT_CONNECTED_POWER        Tag = 0x01000277
	EMS_REQ_SET_WALLBOX_ENFORCE_POWER_ASSIGNMENT Tag = 0x0100027A
	EMS_REQ_GET_WALLBOX_ENFORCE_POWER_ASSIGNMENT Tag = 0x0100027B
	EMS_REQ_SET_WB_DISCHARGE_BAT_UNTIL           Tag = 0x0100027C
	EMS_REQ_GET_WB_DISCHARGE_BAT_UNTIL           Tag = 0x0100027D
	EMS_REQ_WB_AVAILABLE                         Tag = 0x01000280
	EMS_REQ_WB_PREFERRED_CHARGE_POWER            Tag = 0x01000281
	EMS_REQ_SET_PEAK_SHAVING_POWER               Tag = 0x01000282
	EMS_REQ_GET_PEAK_SHAVING_POWER               Tag = 0x01000283
	EMS_REQ_GET_RUNSCREENVALUES                  Tag = 0x01000284
	EMS_REQ_SET_PEAK_SHAVING_TIMES               Tag = 0x01000286
	EMS_REQ_GET_PEAK_SHAVING_TIMES               Tag = 0x01000287
	EMS_REQ_SET_LIST_ACTOR                       Tag = 0x01000288
	EMS_REQ_GET_LIST_ACTOR                       Tag = 0x01000289
	EMS_REQ_WB_BIC_LOAD_PRICE_POWER_TABLE        Tag = 0x01000307
	EMS_REQ_WB_BIC_PRICE_POWER_TABLE_STATUS      Tag = 0x01000308
	EMS_PARAM_DERATE_POWER_VALUE                 Tag = 0x01040001
	EMS_PARAM_AVAILABLE_POWER                    Tag = 0x01040002
	EMS_PARAM_IP_REMOTE_CONTROL                  Tag = 0x01040004
	EMS_PARAM_POWEROFFSET_VALUE                  Tag = 0x01040005
	EMS_PARAM_POWER_VALUE_L1                     Tag = 0x01040006
	EMS_PARAM_POWER_VALUE_L2                     Tag = 0x01040007
	EMS_PARAM_POWER_VALUE_L3                     Tag = 0x01040008
	EMS_PARAM_SET_POINT                          Tag = 0x01040009
	EMS_PARAM_DERATE_POWER_VALUE_L1              Tag = 0x01040010
	EMS_PARAM_DERATE_POWER_VALUE_L2              Tag = 0x01040011
	EMS_PARAM_DERATE_POWER_VALUE_L3              Tag = 0x01040012
	EMS_PARAM_REMOTE_CONTROL_ACTIVE              Tag = 0x01040013
	EMS_PARAM_TIME_TO_RETRY                      Tag = 0x01040014
	EMS_PARAM_NO_REMAINING_RETRY                 Tag = 0x01040015
	EMS_PARAM_INDEX                              Tag = 0x01040016
	EMS_PARAM_WALLBOX_SETPOINT_L1                Tag = 0x01040017
	EMS_PARAM_WALLBOX_SETPOINT_L2                Tag = 0x01040018
	EMS_PARAM_WALLBOX_SETPOINT_L3                Tag = 0x01040019
	EMS_PARAM_REGULATOR_MODE                     Tag = 0x01040113
	EMS_PARAM_REGULATOR_STRATEGY                 Tag = 0x01040114
	EMS_PARAM_DEACTIVATE_SURPLUS_ACTOR           Tag = 0x01040115
	EMS_PARAM_POWER_GRID_OFFSET_L1               Tag = 0x01040294
	EMS_PARAM_POWER_GRID_OFFSET_L2               Tag = 0x01040295
	EMS_PARAM_POWER_GRID_OFFSET_L3               Tag = 0x01040296
	EMS_PARAM_POWER_GRID_OVERRIDE_L1             Tag = 0x01040297
	EMS_PARAM_POWER_GRID_OVERRIDE_L2             Tag = 0x01040298
	EMS_PARAM_POWER_GRID_OVERRIDE_L3             Tag = 0x01040299

	EMS_PARAM_LIMITS_TOTAL_MAX               Tag = 0x01400265
	EMS_PARAM_LIMITS_TOTAL_MIN               Tag = 0x01400266
	EMS_PARAM_LIMITS_PHASE_MAX_L1            Tag = 0x01400267
	EMS_PARAM_LIMITS_PHASE_MAX_L2            Tag = 0x01400268
	EMS_PARAM_LIMITS_PHASE_MAX_L3            Tag = 0x01400269
	EMS_PARAM_LIMITS_PHASE_MIN_L1            Tag = 0x01400270
	EMS_PARAM_LIMITS_PHASE_MIN_L2            Tag = 0x01400271
	EMS_PARAM_LIMITS_PHASE_MIN_L3            Tag = 0x01400272
	EMS_PARAM_CURR_CHARGED_ENERGY_EP_RESERVE Tag = 0x01400278
	EMS_REQ_SET_TOTAL_EP_RESERVE_W           Tag = 0x01400279
	EMS_IDLE_PERIOD_2                        Tag = 0x0140027A

	EMS_RESCUE_BAT_MODE            Tag = 0x01800022
	EMS_IS_PV_DERATING             Tag = 0x01800024
	EMS_AC_POWER_LIMIT             Tag = 0x01800025
	EMS_POWER_ADD_SINKS            Tag = 0x01800026
	EMS_SET_EP_WALLBOX_ALLOW       Tag = 0x01800027
	EMS_GET_EP_WALLBOX_ALLOW       Tag = 0x01800028
	EMS_SET_MAX_EP_WALLBOX_POWER_W Tag = 0x01800029
	EMS_GET_MAX_EP_WALLBOX_POWER_W Tag = 0x0180002A
	EMS_GET_MIN_EP_WALLBOX_POWER_W Tag = 0x0180002B
	EMS_SET_MIN_EP_WALLBOX_POWER_W Tag = 0x0180002C
	EMS_SET_EP_WALLBOX_ENERGY      Tag = 0x0180002D
	EMS_GET_EP_WALLBOX_ENERGY      Tag = 0x0180002E
	EMS_SET_IDLE_PERIODS_2         Tag = 0x0180002F

	EMS_GET_IDLE_PERIODS_2      Tag = 0x01800033
	EMS_SET_IDLE_PERIODS_ENABLE Tag = 0x01800034
	EMS_GET_IDLE_PERIODS_ENABLE Tag = 0x01800035
	EMS_SET_EP_WALLBOX_PRIORITY Tag = 0x01800036
	EMS_GET_EP_WALLBOX_PRIORITY Tag = 0x01800037

	EMS_RES_POWERSAVE_TIMEOUT                Tag = 0x01800107
	EMS_REMOTE_CONTROL                       Tag = 0x01800200
	EMS_DEACTIVATE_REMOTE_CONTROL            Tag = 0x01800201
	EMS_IP_REMOTE_CONTROL                    Tag = 0x01800202
	EMS_EP_DELAY                             Tag = 0x01800203
	EMS_SET_EP_DELAY                         Tag = 0x01800204
	EMS_REMOTE_CONTROL_STATUS                Tag = 0x01800205
	EMS_IDLE_PERIOD_MIN_SOC_UCB              Tag = 0x01800206
	EMS_IDLE_PERIOD_MAX_SOC_UCB              Tag = 0x01800207
	EMS_SET_IDLE_PERIOD_MIN_SOC_UCB          Tag = 0x01800208
	EMS_SET_IDLE_PERIOD_MAX_SOC_UCB          Tag = 0x01800209
	EMS_REGULATOR_MODE                       Tag = 0x01800210
	EMS_SET_REGULATOR_MODE                   Tag = 0x01800211
	EMS_SUPPORTED_REGULATOR_MODES            Tag = 0x01800212
	EMS_EMERGENCY_POWER_OVERLOAD_STATUS      Tag = 0x01800213
	EMS_EMERGENCY_POWER_RETRY                Tag = 0x01800214
	EMS_DETECT_PHASE_OFFSET                  Tag = 0x01800217
	EMS_PHASE_DETECTION_STATUS               Tag = 0x01800218
	EMS_PHASE_OFFSET                         Tag = 0x01800219
	EMS_ABORT_PHASE_DETECTION                Tag = 0x01800220
	EMS_REGULATOR_STRATEGY                   Tag = 0x01800221
	EMS_SET_REGULATOR_STRATEGY               Tag = 0x01800222
	EMS_POWER_PV_AC_OUT                      Tag = 0x01800223
	EMS_PV_ENERGY                            Tag = 0x01800224
	EMS_PARAM_AC_ENERGY_OUT                  Tag = 0x01800225
	EMS_PARAM_AC_ENERGY_IN                   Tag = 0x01800226
	EMS_PARAM_DC_IN                          Tag = 0x01800227
	EMS_ENERGY_STORAGE_MODEL                 Tag = 0x01800228
	EMS_PARAM_CURR_CHARGED_ENERGY            Tag = 0x01800229
	EMS_PARAM_FULL_CHARGED_ENERGY_EP_RESERVE Tag = 0x01800230
	EMS_PARAM_DESIGN_ENERGY                  Tag = 0x01800231
	EMS_PARAM_FULL_CHARGED_ENERGY            Tag = 0x01800232
	EMS_PARAM_USED_CAPACITY                  Tag = 0x01800233
	EMS_SPECIFICATION_VALUES                 Tag = 0x01800234
	EMS_PARAM_MAX_CHARGE_POWER               Tag = 0x01800235
	EMS_PARAM_MAX_DISCHARGE_POWER            Tag = 0x01800236
	EMS_PARAM_MAX_PV_POWER                   Tag = 0x01800237
	EMS_PARAM_MAX_AC_POWER                   Tag = 0x01800238
	EMS_PARAM_INSTALLED_BAT_CAP              Tag = 0x01800239
	EMS_PARAM_HYBRIT_SUPPORTED               Tag = 0x01800240
	EMS_PARAM_INIT_STATUS                    Tag = 0x01800241
	EMS_EP_RESERVE                           Tag = 0x01800242
	EMS_SEC_LIMITS                           Tag = 0x01800243
	EMS_PARAM_SEL_TOTAL_MAX                  Tag = 0x01800244
	EMS_PARAM_SEL_TOTAL_MIN                  Tag = 0x01800245
	EMS_PARAM_SEL_PHASE_MAX_L1               Tag = 0x01800246
	EMS_PARAM_SEL_PHASE_MAX_L2               Tag = 0x01800247
	EMS_PARAM_SEL_PHASE_MAX_L3               Tag = 0x01800248
	EMS_PARAM_SEL_PHASE_MIN_L1               Tag = 0x01800249
	EMS_PARAM_SEL_PHASE_MIN_L2               Tag = 0x01800250
	EMS_PARAM_SEL_PHASE_MIN_L3               Tag = 0x01800251
	EMS_SEC_DEVICE_STATUS                    Tag = 0x01800252
	EMS_PARAM_PVI_1                          Tag = 0x01800253
	EMS_PARAM_PVI_2                          Tag = 0x01800254
	EMS_PARAM_PVI_3                          Tag = 0x01800255
	EMS_PARAM_DCDC                           Tag = 0x01800256
	EMS_PARAM_BAT                            Tag = 0x01800257
	EMS_BAT_CURRENT_IN                       Tag = 0x01800258
	EMS_BAT_CURRENT_OUT                      Tag = 0x01800259
	EMS_MAX_DC_POWER                         Tag = 0x01800260
	EMS_AC_REACTIVE_POWER                    Tag = 0x01800261
	EMS_GET_PARTIAL_GRID                     Tag = 0x01800263
	EMS_ESTIMATED_POWER_LIMITS               Tag = 0x01800264
	EMS_DESIGN_POWER_LIMITS                  Tag = 0x01800265
	EMS_SET_CAN_ID_FEED_IN_REDUCTION         Tag = 0x01800266
	EMS_CAN_ID_FEED_IN_REDUCTION             Tag = 0x01800267
	EMS_SET_CAN_ID_UNBALANCED_LOAD           Tag = 0x01800268
	EMS_CAN_ID_UNBALANCED_LOAD               Tag = 0x01800269
	EMS_SET_WALLBOX_MODE                     Tag = 0x01800270
	EMS_GET_WALLBOX_MODE                     Tag = 0x01800271
	EMS_SET_MAX_FUSE_POWER                   Tag = 0x01800272
	EMS_GET_MAX_FUSE_POWER                   Tag = 0x01800273
	EMS_SET_CONNECTED_POWER                  Tag = 0x01800274
	EMS_GET_CONNECTED_POWER                  Tag = 0x01800275
	EMS_DERATE_AT_CONNECTED_POWER            Tag = 0x01800276
	EMS_SET_DERATE_AT_CONNECTED_POWER        Tag = 0x01800277
	EMS_SET_WALLBOX_ENFORCE_POWER_ASSIGNMENT Tag = 0x0180027A
	EMS_GET_WALLBOX_ENFORCE_POWER_ASSIGNMENT Tag = 0x0180027B
	EMS_SET_WB_DISCHARGE_BAT_UNTIL           Tag = 0x0180027C
	EMS_GET_WB_DISCHARGE_BAT_UNTIL           Tag = 0x0180027D
	EMS_WB_AVAILABLE                         Tag = 0x01800280
	EMS_WB_PREFERRED_CHARGE_POWER            Tag = 0x01800281
	EMS_SET_PEAK_SHAVING_POWER               Tag = 0x01800282
	EMS_GET_PEAK_SHAVING_POWER               Tag = 0x01800283
	EMS_GET_RUNSCREENVALUES                  Tag = 0x01800284
	EMS_SET_PEAK_SHAVING_TIMES               Tag = 0x01800286
	EMS_GET_PEAK_SHAVING_TIMES               Tag = 0x01800287
	EMS_SET_LIST_ACTOR                       Tag = 0x01800288
	EMS_GET_LIST_ACTOR                       Tag = 0x01800289
	EMS_ACTOR_ITEM                           Tag = 0x01800290
	EMS_ACTOR_ID                             Tag = 0x01800291
	EMS_ACTOR_NAME                           Tag = 0x01800292
	EMS_ACTOR_PRIORITY                       Tag = 0x01800293
	EMS_PERIOD_ITEM                          Tag = 0x01800300
	EMS_PERIOD_ACTIVE                        Tag = 0x01800301
	EMS_PERIOD_NAME                          Tag = 0x01800302
	EMS_PERIOD_WEEKDAYS                      Tag = 0x01800303
	EMS_PERIOD_START                         Tag = 0x01800304
	EMS_PERIOD_STOP                          Tag = 0x01800305
	EMS_PERIOD_POWER                         Tag = 0x01800306
	EMS_WB_BIC_LOAD_PRICE_POWER_TABLE        Tag = 0x01800307
	EMS_WB_BIC_PRICE_POWER_TABLE_STATUS      Tag = 0x01800308

	EMS_SET_TOTAL_EP_RESERVE_W Tag = 0x01C00279
)

// ---------------------
// NAMESPACE: PVInverter
// 0x02xxxxxx
// ---------------------
const (
	// PVI_INDEX & PVI_... Antwort mit allen Daten der REQ_DATA Anfrage
	PVI_DATA Tag = 0x02840000
	// PVI_INDEX & PVI_REQ...  Beinhaltet alle Anfrage-TAGs, der Container MUSS einen Index enthalten
	PVI_REQ_DATA Tag = 0x02040000
	// Index des angefragten Gerätes (0?x), Muss in Anfrage und Antwort zum DATA-Tag vorkommen
	PVI_INDEX Tag = 0x02040001
	// dataType gibt den jeweiligen Daten Typ zurück!
	PVI_VALUE            Tag = 0x02040005
	PVI_GENERAL_ERROR    Tag = 0x02FFFFFF
	PVI_ON_GRID          Tag = 0x02800001
	PVI_REQ_ON_GRID      Tag = 0x02000001
	PVI_STATE            Tag = 0x02800002
	PVI_REQ_STATE        Tag = 0x02000002
	PVI_LAST_ERROR       Tag = 0x02800003
	PVI_REQ_LAST_ERROR   Tag = 0x02000003
	PVI_FLASH_FILE       Tag = 0x02800007
	PVI_REQ_DEVICE_STATE Tag = 0x02060000
	// DEVICE_CONNECTED & DEVICE_WORKING & DEVICE_IN_SERVICE
	PVI_DEVICE_STATE      Tag = 0x02860000
	PVI_DEVICE_CONNECTED  Tag = 0x02860001
	PVI_DEVICE_WORKING    Tag = 0x02860002
	PVI_DEVICE_IN_SERVICE Tag = 0x02860003
	PVI_REQ_TYPE          Tag = 0x02000009
	// 1=SOLU 2=KACO 3=E3DC_E
	PVI_TYPE Tag = 0x02800009
	// PVI_COS_PHI_VALUE & PVI_COS_PHI_IS_AKTIV & PVI_COS_PHI_EXCITED
	PVI_COS_PHI     Tag = 0x02800060
	PVI_REQ_COS_PHI Tag = 0x02000060
	// PVI_COS_PHI_VALUE & PVI_COS_PHI_IS_AKTIV & PVI_COS_PHI_EXCITED
	PVI_REQ_SET_COS_PHI  Tag = 0x02000061
	PVI_COS_PHI_VALUE    Tag = 0x02000062
	PVI_COS_PHI_IS_AKTIV Tag = 0x02000063
	PVI_COS_PHI_EXCITED  Tag = 0x02000064
	// PVI_VOLTAGE_MONITORING_THRESHOLD_TOP &
	// PVI_VOLTAGE_MONITORING_THRESHOLD_BOTTOM &
	// PVI_VOLTAGE_MONITORING_SLOPE_UP &
	// PVI_VOLTAGE_MONITORING_SLOPE_DOWN
	PVI_VOLTAGE_MONITORING                  Tag = 0x02800070
	PVI_REQ_VOLTAGE_MONITORING              Tag = 0x02000070
	PVI_VOLTAGE_MONITORING_THRESHOLD_TOP    Tag = 0x02000072
	PVI_VOLTAGE_MONITORING_THRESHOLD_BOTTOM Tag = 0x02000073
	PVI_VOLTAGE_MONITORING_SLOPE_UP         Tag = 0x02000074
	PVI_VOLTAGE_MONITORING_SLOPE_DOWN       Tag = 0x02000075
	// PVI_FREQUENCY_UNDER & PVI_FREQUENCY_OVER
	PVI_FREQUENCY_UNDER_OVER     Tag = 0x02800080
	PVI_REQ_FREQUENCY_UNDER_OVER Tag = 0x02000080
	PVI_FREQUENCY_UNDER          Tag = 0x02000082
	PVI_FREQUENCY_OVER           Tag = 0x02000083
	// Mode:
	//  IdleMode = 0,
	//  NormalMode = 1,
	//  GridChargeMode = 2,
	//  BackupPowerMode = 3
	PVI_SYSTEM_MODE     Tag = 0x02800085
	PVI_REQ_SYSTEM_MODE Tag = 0x02000085
	// Mode:
	//  PVI ON 1
	//  PVI OFF 0
	//  PVI ON_FORCE 101
	//  PVI OFF_FORCE 100
	PVI_POWER_MODE     Tag = 0x02800087
	PVI_REQ_POWER_MODE Tag = 0x02000087
	// PVI_INDEX & PVI_VALUE
	PVI_TEMPERATURE           Tag = 0x02800100
	PVI_REQ_TEMPERATURE       Tag = 0x02000100
	PVI_TEMPERATURE_COUNT     Tag = 0x02800101
	PVI_REQ_TEMPERATURE_COUNT Tag = 0x02000101
	PVI_MAX_TEMPERATURE       Tag = 0x02800102
	PVI_REQ_MAX_TEMPERATURE   Tag = 0x02000102
	PVI_MIN_TEMPERATURE       Tag = 0x02800103
	PVI_REQ_MIN_TEMPERATURE   Tag = 0x02000103
	PVI_SERIAL_NUMBER         Tag = 0x028ABC01
	PVI_REQ_SERIAL_NUMBER     Tag = 0x020ABC01
	// PVI_VERSION_MAIN |& PVI_VERSION_PIC |& ?.
	PVI_VERSION            Tag = 0x028ABC02
	PVI_REQ_VERSION        Tag = 0x020ABC02
	PVI_VERSION_MAIN       Tag = 0x020ABC03
	PVI_VERSION_PIC        Tag = 0x020ABC04
	PVI_AC_MAX_PHASE_COUNT Tag = 0x028AC000
	// PVI_INDEX & PVI_VALUE
	PVI_AC_POWER Tag = 0x028AC001
	// PVI_INDEX & PVI_VALUE
	PVI_AC_VOLTAGE Tag = 0x028AC002
	// PVI_INDEX & PVI_VALUE
	PVI_AC_CURRENT Tag = 0x028AC003
	// PVI_INDEX & PVI_VALUE
	PVI_AC_APPARENTPOWER Tag = 0x028AC004
	// PVI_INDEX & PVI_VALUE
	PVI_AC_REACTIVEPOWER Tag = 0x028AC005
	// PVI_INDEX & PVI_VALUE
	PVI_AC_ENERGY_ALL Tag = 0x028AC006
	// PVI_INDEX & PVI_VALUE
	PVI_AC_MAX_APPARENTPOWER Tag = 0x028AC007
	// PVI_INDEX & PVI_VALUE
	PVI_AC_ENERGY_DAY Tag = 0x028AC008
	// PVI_INDEX & PVI_VALUE
	PVI_AC_ENERGY_GRID_CONSUMPTION Tag = 0x028AC009
	PVI_REQ_AC_MAX_PHASE_COUNT     Tag = 0x020AC000
	// Value der Anfrage beinhaltet die angefragte Phase
	PVI_REQ_AC_POWER Tag = 0x020AC001
	// Value der Anfrage beinhaltet die angefragte Phase
	PVI_REQ_AC_VOLTAGE Tag = 0x020AC002
	// Value der Anfrage beinhaltet die angefragte Phase
	PVI_REQ_AC_CURRENT Tag = 0x020AC003
	// Value der Anfrage beinhaltet die angefragte Phase
	PVI_REQ_AC_APPARENTPOWER Tag = 0x020AC004
	// Value der Anfrage beinhaltet die angefragte Phase
	PVI_REQ_AC_REACTIVEPOWER Tag = 0x020AC005
	// Value der Anfrage beinhaltet die angefragte Phase
	PVI_REQ_AC_ENERGY_ALL Tag = 0x020AC006
	// Value der Anfrage beinhaltet die angefragte Phase
	PVI_REQ_AC_MAX_APPARENTPOWER Tag = 0x020AC007
	// Value der Anfrage beinhaltet die angefragte Phase
	PVI_REQ_AC_ENERGY_DAY Tag = 0x020AC008
	// Value der Anfrage beinhaltet die angefragte Phase
	PVI_REQ_AC_ENERGY_GRID_CONSUMPTION Tag = 0x020AC009
	PVI_DC_MAX_STRING_COUNT            Tag = 0x028DC000
	// PVI_INDEX & PVI_VALUE
	PVI_DC_POWER Tag = 0x028DC001
	// PVI_INDEX & PVI_VALUE
	PVI_DC_VOLTAGE Tag = 0x028DC002
	// PVI_INDEX & PVI_VALUE
	PVI_DC_CURRENT Tag = 0x028DC003
	// PVI_INDEX & PVI_VALUE
	PVI_DC_MAX_POWER Tag = 0x028DC004
	// PVI_INDEX & PVI_VALUE
	PVI_DC_MAX_VOLTAGE Tag = 0x028DC005
	// PVI_INDEX & PVI_VALUE
	PVI_DC_MIN_VOLTAGE Tag = 0x028DC006
	// PVI_INDEX & PVI_VALUE
	PVI_DC_MAX_CURRENT Tag = 0x028DC007
	// PVI_INDEX & PVI_VALUE
	PVI_DC_MIN_CURRENT Tag = 0x028DC008
	// PVI_INDEX & PVI_VALUE
	PVI_DC_STRING_ENERGY_ALL     Tag = 0x028DC009
	PVI_REQ_DC_MAX_STRING_COUNT  Tag = 0x020DC000
	PVI_REQ_DC_POWER             Tag = 0x020DC001
	PVI_REQ_DC_VOLTAGE           Tag = 0x020DC002
	PVI_REQ_DC_CURRENT           Tag = 0x020DC003
	PVI_REQ_DC_MAX_POWER         Tag = 0x020DC004
	PVI_REQ_DC_MAX_VOLTAGE       Tag = 0x020DC005
	PVI_REQ_DC_MIN_VOLTAGE       Tag = 0x020DC006
	PVI_REQ_DC_MAX_CURRENT       Tag = 0x020DC007
	PVI_REQ_DC_MIN_CURRENT       Tag = 0x020DC008
	PVI_REQ_DC_STRING_ENERGY_ALL Tag = 0x020DC009

	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	PVI_REQ_IS_FLASHING            Tag = 0x02000004
	PVI_REQ_START_FLASHING         Tag = 0x02000005
	PVI_REQ_FLASH_FILE_LIST        Tag = 0x02000006
	PVI_REQ_SERVICE_PROGRESS_STATE Tag = 0x02000008

	PVI_REQ_GET_GAPP_LAST_CHANGE_TIME Tag = 0x0200000A
	PVI_REQ_GET_GAPP_SYNC_TIME        Tag = 0x0200000B
	PVI_REQ_SET_GAPP_SYNC_TIME        Tag = 0x0200000C
	PVI_REQ_LAND_CODE                 Tag = 0x02000010
	PVI_REQ_LAND_CODE_LIST            Tag = 0x02000011
	PVI_REQ_SET_LAND_CODE             Tag = 0x02000012
	PVI_REQ_ERROR_LIST                Tag = 0x02000013
	PVI_ERROR_LIST                    Tag = 0x02000014
	PVI_REQ_STATUS_LIST               Tag = 0x02000015
	PVI_STATUS_LIST                   Tag = 0x02000016
	PVI_REQ_SET_DEVICE_SILENCE        Tag = 0x02000017
	PVI_REQ_DEVICE_SILENCE            Tag = 0x02000018
	PVI_REQ_SELF_TEST                 Tag = 0x02000019
	PVI_REQ_SET_FREE_INVERTER         Tag = 0x02000020
	PVI_REQ_SET_BLOCK_INVERTER        Tag = 0x02000021
	PVI_REQ_UZK_VOLTAGE               Tag = 0x02000050

	PVI_REQ_SET_VOLTAGE_MONITORING Tag = 0x02000071

	PVI_REQ_SET_FREQUENCY_UNDER_OVER Tag = 0x02000081

	PVI_REQ_SET_SYSTEM_MODE Tag = 0x02000084

	PVI_REQ_SET_POWER_MODE Tag = 0x02000086

	PVI_REQ_USED_STRING_COUNT     Tag = 0x02000090
	PVI_REQ_SET_USED_STRING_COUNT Tag = 0x02000091
	PVI_REQ_DERATE_TO_POWER       Tag = 0x02000092

	PVI_REQ_CT_TAR_USV_BOX         Tag = 0x02000104
	PVI_REQ_ROTARY_FIELD_DIRECTION Tag = 0x02000105

	PVI_VERSION_ALL              Tag = 0x020ABC05
	PVI_REQ_GAPP_JSON            Tag = 0x020ABC13
	PVI_REQ_GAPP_JSON_MD5        Tag = 0x020ABC14
	PVI_REQ_SET_RESET_GPIO       Tag = 0x020ABC50
	PVI_REQ_SET_POWEROFF_GPIO    Tag = 0x020ABC51
	PVI_REQ_SET_NIGHTSWITCH_GPIO Tag = 0x020ABC52
	PVI_REQ_SET_PRE_GRID_CHARGE  Tag = 0x020ABC60

	PVI_REQ_AC_FREQUENCY   Tag = 0x020AC00A
	PVI_REQ_INVERTER_COUNT Tag = 0x020CCCC9

	PVI_REQ_AC_ENERGY_PRODUCED_L1 Tag = 0x020DC00A
	PVI_REQ_AC_ENERGY_PRODUCED_L2 Tag = 0x020DC00B
	PVI_REQ_AC_ENERGY_PRODUCED_L3 Tag = 0x020DC00C
	PVI_REQ_AC_ENERGY_CONSUMED_L1 Tag = 0x020DC00D
	PVI_REQ_AC_ENERGY_CONSUMED_L2 Tag = 0x020DC00E
	PVI_REQ_AC_ENERGY_CONSUMED_L3 Tag = 0x020DC00F
	PVI_REQ_ENABLE_FAN_TEST       Tag = 0x020DC010
	PVI_REQ_DISABLE_FAN_TEST      Tag = 0x020DC011
	PVI_REQ_RESET_LAND_NORM       Tag = 0x020DC012
	PVI_REQ_IS_FAN_INSTALLED      Tag = 0x020DC013

	PVI_IS_FLASHING     Tag = 0x02800004
	PVI_FLASH_FILE_LIST Tag = 0x02800006

	PVI_SERVICE_PROGRESS_STATE Tag = 0x02800008

	PVI_GET_GAPP_LAST_CHANGE_TIME Tag = 0x0280000A
	PVI_GET_GAPP_SYNC_TIME        Tag = 0x0280000B
	PVI_SET_GAPP_SYNC_TIME        Tag = 0x0280000C
	PVI_LAND_CODE                 Tag = 0x02800010
	PVI_LAND_CODE_LIST            Tag = 0x02800011
	PVI_ERROR_STRING              Tag = 0x02800013
	PVI_STATUS_STRING             Tag = 0x02800015
	PVI_DEVICE_SILENCE            Tag = 0x02800018
	PVI_SELF_TEST                 Tag = 0x02800019
	PVI_SET_FREE_INVERTER         Tag = 0x02800020
	PVI_SET_BLOCK_INVERTER        Tag = 0x02800021
	PVI_UZK_VOLTAGE               Tag = 0x02800050

	PVI_SET_FREQUENCY_UNDER_OVER Tag = 0x02800081
	PVI_SET_SYSTEM_MODE          Tag = 0x02800084

	PVI_SET_POWER_MODE Tag = 0x02800086

	PVI_USED_STRING_COUNT Tag = 0x02800090
	PVI_DERATE_TO_POWER   Tag = 0x02800092

	PVI_CT_TAR_USV_BOX         Tag = 0x02800104
	PVI_ROTARY_FIELD_DIRECTION Tag = 0x02800105

	PVI_GAPP_JSON     Tag = 0x028ABC13
	PVI_GAPP_JSON_MD5 Tag = 0x028ABC14

	PVI_AC_FREQUENCY   Tag = 0x028AC00A
	PVI_INVERTER_COUNT Tag = 0x028CCCC9

	PVI_AC_ENERGY_PRODUCED_L1 Tag = 0x028DC00A
	PVI_AC_ENERGY_PRODUCED_L2 Tag = 0x028DC00B
	PVI_AC_ENERGY_PRODUCED_L3 Tag = 0x028DC00C
	PVI_AC_ENERGY_CONSUMED_L1 Tag = 0x028DC00D
	PVI_AC_ENERGY_CONSUMED_L2 Tag = 0x028DC00E
	PVI_AC_ENERGY_CONSUMED_L3 Tag = 0x028DC00F
	PVI_RESET_LAND_NORM       Tag = 0x028DC012
	PVI_IS_FAN_INSTALLED      Tag = 0x028DC013
)

// ------------------
// NAMESPACE: Battery
// 0x03xxxxxx
// ------------------
const (
	/*
		BAT_REQ_SPECIFICATION        Tag = 0x03000043 // Source: https://github.com/pvtom/rscp2mqtt/blob/main/RscpTags.h
		BAT_SPECIFICATION            Tag = 0x03800043 // Source: https://github.com/pvtom/rscp2mqtt/blob/main/RscpTags.h
		BAT_SPECIFIED_CAPACITY       Tag = 0x03800125 // Source: https://github.com/pvtom/rscp2mqtt/blob/main/RscpTags.h
		BAT_SPECIFIED_DSCHARGE_POWER Tag = 0x03800126 // Source: https://github.com/pvtom/rscp2mqtt/blob/main/RscpTags.h
		BAT_SPECIFIED_CHARGE_POWER   Tag = 0x03800127 // Source: https://github.com/pvtom/rscp2mqtt/blob/main/RscpTags.h
		BAT_SPECIFIED_MAX_DCB_COUNT  Tag = 0x03800128 // Source: https://github.com/pvtom/rscp2mqtt/blob/main/RscpTags.h
		BAT_ROLE                     Tag = 0x03800129 // Source: https://github.com/pvtom/rscp2mqtt/blob/main/RscpTags.h
	*/

	// Beinhaltet alle Anfrage-TAGs, der Container MUSS einen Index enthalten
	BAT_REQ_DATA Tag = 0x03040000
	// Index des angefragten Gerätes (Im Moment immer 0 bei der Batterie), kann in der Anfrage und in der Antwort vorkommen.
	BAT_INDEX Tag = 0x03040001
	// Antwort mit allen Daten der REQ_DATA Anfrage
	BAT_DATA Tag = 0x03840000
	// Rückgabewert für errechnet SOC Wert
	BAT_RSOC Tag = 0x03800001
	// Rückgabewert für gesamte Batteriespannung
	BAT_MODULE_VOLTAGE Tag = 0x03800002
	// Rückgabewert für gesamten Batteriestrom
	BAT_CURRENT Tag = 0x03800003
	// Rückgabewert für maximale Batteriespannung
	BAT_MAX_BAT_VOLTAGE Tag = 0x03800004
	// Rückgabewert für maximale Batterieladestrom
	BAT_MAX_CHARGE_CURRENT Tag = 0x03800005
	// Rückgabewert für Entladeschlussspannung
	BAT_EOD_VOLTAGE Tag = 0x03800006
	// Rückgabewert für maximale Batterieentladestrom
	BAT_MAX_DISCHARGE_CURRENT Tag = 0x03800007
	// Rückgabewert für Batterieladezyklen
	BAT_CHARGE_CYCLES Tag = 0x03800008
	// Rückgabewert für die Terminalspannung
	BAT_TERMINAL_VOLTAGE Tag = 0x03800009
	// Rückgabewert für Batteriestatus
	BAT_STATUS_CODE Tag = 0x0380000A
	// Rückgabewert für Batteriefehler
	BAT_ERROR_CODE Tag = 0x0380000B
	// Rückgabewert für Batteriebezeichnung
	BAT_DEVICE_NAME Tag = 0x0380000C
	// Rückgabewert für Anzahl der gefundenen DCBs
	BAT_DCB_COUNT Tag = 0x0380000D

	BAT_MAX_DCB_CELL_CURRENT Tag = 0x03800012 // source https://github.com/rxhan/RSCPGui
	BAT_MIN_DCB_CELL_CURRENT Tag = 0x03800013 // source https://github.com/rxhan/RSCPGui
	BAT_MAX_DCB_CELL_VOLTAGE Tag = 0x03800014 // source https://github.com/rxhan/RSCPGui
	BAT_MIN_DCB_CELL_VOLTAGE Tag = 0x03800015 // source https://github.com/rxhan/RSCPGui

	BAT_MAX_DCB_CELL_TEMPERATURE Tag = 0x03800016
	// Ein Container mit allen Temperaturen für die angefragte DCB.
	BAT_MIN_DCB_CELL_TEMPERATURE Tag = 0x03800017
	// Ein Container mit allen Spannungen für die angefragte DCB.
	BAT_DCB_CELL_TEMPERATURE Tag = 0x03800019
	BAT_DCB_CELL_VOLTAGE     Tag = 0x0380001B
	BAT_READY_FOR_SHUTDOWN   Tag = 0x0380001E
	// Dieser Container beinhaltet die Antwort auf ein REQ_INFO. Es beinhaltet immer die folgenden TAGs:
	//  - BAT_RSOC
	//  - BAT_MODULE_VOLTAGE
	//  - BAT_CURRENT
	//  - BAT_MAX_DCB_CELL_TEMPERATURE
	//  - BAT_STATUS_CODE
	//  - BAT_ERROR_CODE
	//  - BAT_CHARGE_CYCLES
	BAT_INFO Tag = 0x03800020
	// Batterietrainingmodus
	//  0 - Nicht im Training
	//  1 - Trainingmodus Entladen
	//  2 - Trainingmodus Laden
	BAT_TRAINING_MODE Tag = 0x03800021
	// Kann nur innerhalb eines REQ_BAT_DATA Container verwendet werden!
	BAT_REQ_RSOC Tag = 0x03000001
	// Kann nur innerhalb eines REQ_BAT_DATA Container verwendet werden!
	BAT_REQ_MODULE_VOLTAGE Tag = 0x03000002
	// Kann nur innerhalb eines REQ_BAT_DATA Container verwendet werden!
	BAT_REQ_CURRENT Tag = 0x03000003
	// Kann nur innerhalb eines REQ_BAT_DATA Container verwendet werden!
	BAT_REQ_MAX_BAT_VOLTAGE Tag = 0x03000004
	// Kann nur innerhalb eines REQ_BAT_DATA Container verwendet werden!
	BAT_REQ_MAX_CHARGE_CURRENT Tag = 0x03000005
	// Kann nur innerhalb eines REQ_BAT_DATA Container verwendet werden!
	BAT_REQ_EOD_VOLTAGE Tag = 0x03000006
	// Kann nur innerhalb eines REQ_BAT_DATA Container verwendet werden!
	BAT_REQ_MAX_DISCHARGE_CURRENT Tag = 0x03000007
	// Kann nur innerhalb eines REQ_BAT_DATA Container verwendet werden!
	BAT_REQ_CHARGE_CYCLES Tag = 0x03000008
	// Kann nur innerhalb eines REQ_BAT_DATA Container verwendet werden!
	BAT_REQ_TERMINAL_VOLTAGE Tag = 0x03000009
	// Kann nur innerhalb eines REQ_BAT_DATA Container verwendet werden!
	BAT_REQ_STATUS_CODE Tag = 0x0300000A
	// Kann nur innerhalb eines REQ_BAT_DATA Container verwendet werden!
	BAT_REQ_ERROR_CODE Tag = 0x0300000B
	// Kann nur innerhalb eines REQ_BAT_DATA Container verwendet werden!
	BAT_REQ_DEVICE_NAME Tag = 0x0300000C
	// Kann nur innerhalb eines REQ_BAT_DATA Container verwendet werden!
	BAT_REQ_DCB_COUNT Tag = 0x0300000D
	// Kann nur innerhalb eines REQ_BAT_DATA Container verwendet werden!
	BAT_REQ_MAX_DCB_CELL_TEMPERATURE Tag = 0x03000016
	// Kann nur innerhalb eines REQ_BAT_DATA Container verwendet werden!
	BAT_REQ_MIN_DCB_CELL_TEMPERATURE  Tag = 0x03000017
	BAT_REQ_DCB_ALL_CELL_TEMPERATURES Tag = 0x03000018 // source https://github.com/rxhan/RSCPGui
	BAT_DCB_ALL_CELL_TEMPERATURES     Tag = 0x03800018 // source https://github.com/rxhan/RSCPGui
	BAT_REQ_DCB_ALL_CELL_VOLTAGES     Tag = 0x0300001A // source https://github.com/rxhan/RSCPGui
	BAT_DCB_ALL_CELL_VOLTAGES         Tag = 0x0380001A // source https://github.com/rxhan/RSCPGui
	// Kann nur innerhalb eines REQ_BAT_DATA Container verwendet werden!
	BAT_REQ_READY_FOR_SHUTDOWN Tag = 0x0300001E
	// Kann nur innerhalb eines REQ_BAT_DATA Container verwendet werden!
	BAT_REQ_INFO Tag = 0x03000020
	// Kann nur innerhalb eines REQ_BAT_DATA Container verwendet werden!
	BAT_REQ_TRAINING_MODE           Tag = 0x03000021
	BAT_REQ_DCB_INFO                Tag = 0x03000042 // source https://github.com/rxhan/RSCPGui
	BAT_DCB_INFO                    Tag = 0x03800042 // source https://github.com/rxhan/RSCPGui
	BAT_DCB_INDEX                   Tag = 0x03800100
	BAT_DCB_LAST_MESSAGE_TIMESTAMP  Tag = 0x03800101
	BAT_DCB_MAX_CHARGE_VOLTAGE      Tag = 0x03800102
	BAT_DCB_MAX_CHARGE_CURRENT      Tag = 0x03800103
	BAT_DCB_END_OF_DISCHARGE        Tag = 0x03800104
	BAT_DCB_MAX_DISCHARGE_CURRENT   Tag = 0x03800105
	BAT_DCB_FULL_CHARGE_CAPACITY    Tag = 0x03800106
	BAT_DCB_REMAINING_CAPACITY      Tag = 0x03800107
	BAT_DCB_SOC                     Tag = 0x03800108
	BAT_DCB_SOH                     Tag = 0x03800109
	BAT_DCB_CYCLE_COUNT             Tag = 0x03800110
	BAT_DCB_CURRENT                 Tag = 0x03800111
	BAT_DCB_VOLTAGE                 Tag = 0x03800112
	BAT_DCB_CURRENT_AVG_30S         Tag = 0x03800113
	BAT_DCB_VOLTAGE_AVG_30S         Tag = 0x03800114
	BAT_DCB_DESIGN_CAPACITY         Tag = 0x03800115
	BAT_DCB_DESIGN_VOLTAGE          Tag = 0x03800116
	BAT_DCB_CHARGE_LOW_TEMPERATURE  Tag = 0x03800117
	BAT_DCB_CHARGE_HIGH_TEMPERATURE Tag = 0x03800118
	BAT_DCB_MANUFACTURE_DATE        Tag = 0x03800119
	BAT_DCB_SERIALNO                Tag = 0x03800120
	BAT_DCB_PROTOCOL_VERSION        Tag = 0x03800121
	BAT_DCB_FW_VERSION              Tag = 0x03800122
	BAT_DCB_DATA_TABLE_VERSION      Tag = 0x03800123
	BAT_DCB_PCB_VERSION             Tag = 0x03800124

	BAT_DCB_NR_SERIES_CELL   Tag = 0x03800300 // source https://github.com/rxhan/RSCPGui
	BAT_DCB_NR_PARALLEL_CELL Tag = 0x03800301 // source https://github.com/rxhan/RSCPGui
	BAT_DCB_MANUFACTURE_NAME Tag = 0x03800302 // source https://github.com/rxhan/RSCPGui
	BAT_DCB_DEVICE_NAME      Tag = 0x03800303 // source https://github.com/rxhan/RSCPGui
	BAT_DCB_SERIALCODE       Tag = 0x03800304 // source https://github.com/rxhan/RSCPGui
	BAT_DCB_NR_SENSOR        Tag = 0x03800305 // source https://github.com/rxhan/RSCPGui
	BAT_DCB_STATUS           Tag = 0x03800306 // source https://github.com/rxhan/RSCPGui
	BAT_DCB_WARNING          Tag = 0x03800307 // source https://github.com/rxhan/RSCPGui
	BAT_DCB_ALARM            Tag = 0x03800308 // source https://github.com/rxhan/RSCPGui
	BAT_DCB_ERROR            Tag = 0x03800309 // source https://github.com/rxhan/RSCPGui

	BAT_REQ_DEVICE_STATE Tag = 0x03060000
	// DEVICE_CONNECTED & DEVICE_WORKING & DEVICE_IN_SERVICE
	BAT_DEVICE_STATE Tag = 0x03860000
	// Kommt nur im BAT_DEVICE_STATE Antwort vor
	BAT_DEVICE_CONNECTED Tag = 0x03860001
	// Kommt nur im BAT_DEVICE_STATE Antwort vor
	BAT_DEVICE_WORKING Tag = 0x03860002
	// Kommt nur im BAT_DEVICE_STATE Antwort vor
	BAT_DEVICE_IN_SERVICE Tag = 0x03860003
	BAT_GENERAL_ERROR     Tag = 0x03FFFFFF

	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	BAT_REQ_RSOC_REAL            Tag = 0x0300000E
	BAT_REQ_ASOC                 Tag = 0x0300000F
	BAT_REQ_FCC                  Tag = 0x03000010
	BAT_REQ_RC                   Tag = 0x03000011
	BAT_REQ_MAX_DCB_CELL_CURRENT Tag = 0x03000012
	BAT_REQ_MIN_DCB_CELL_CURRENT Tag = 0x03000013
	BAT_REQ_MAX_DCB_CELL_VOLTAGE Tag = 0x03000014
	BAT_REQ_MIN_DCB_CELL_VOLTAGE Tag = 0x03000015

	BAT_REQ_OPEN_BREAKER         Tag = 0x0300001C
	BAT_REQ_OPEN_BREAKER_CONFIRM Tag = 0x0300001D

	BAT_REQ_FIRMWARE_VERSION Tag = 0x0300001F

	BAT_REQ_UPDATE_STATUS             Tag = 0x03000022
	BAT_REQ_SET_TRAINING_MODE         Tag = 0x03000023
	BAT_REQ_TIME_LAST_RESPONSE        Tag = 0x03000024
	BAT_REQ_MANUFACTURER_NAME         Tag = 0x03000025
	BAT_REQ_USABLE_CAPACITY           Tag = 0x03000026
	BAT_REQ_USABLE_REMAINING_CAPACITY Tag = 0x03000027
	BAT_REQ_SET_A1_DATA               Tag = 0x03000028
	BAT_REQ_SET_A1_MODE               Tag = 0x03000029
	BAT_REQ_DELETE_DCB_TYPE           Tag = 0x0300002A
	BAT_REQ_DCB_TYPE                  Tag = 0x0300002B
	BAT_REQ_SET_A1_VOLTAGE            Tag = 0x03000030
	BAT_REQ_SET_A1_CURRENT            Tag = 0x03000031
	BAT_REQ_CONTROL_CODE              Tag = 0x03000032
	BAT_REQ_BPM_STATUS                Tag = 0x03000033
	BAT_REQ_DCB_ERROR_LIST            Tag = 0x03000034

	BAT_REQ_SPECIFICATION               Tag = 0x03000043
	BAT_REQ_INTERNALS                   Tag = 0x03000044
	BAT_REQ_DESIGN_CAPACITY             Tag = 0x03000045
	BAT_REQ_DESIGN_VOLTAGE              Tag = 0x03000046
	BAT_REQ_CHARGE_HIGH_TEMP            Tag = 0x03000047
	BAT_REQ_CHARGE_LOW_TEMP             Tag = 0x03000048
	BAT_REQ_MANUFACTURE_DATE            Tag = 0x03000049
	BAT_REQ_SERIALNO                    Tag = 0x03000050
	BAT_REQ_DATA_TABLE_VERSION          Tag = 0x03000051
	BAT_REQ_PROTOCOL_VERSION            Tag = 0x03000052
	BAT_REQ_PCB_VERSION                 Tag = 0x03000053
	BAT_REQ_TOTAL_USE_TIME              Tag = 0x03000054
	BAT_REQ_TOTAL_DISCHARGE_TIME        Tag = 0x03000055
	BAT_REQ_AVAILABLE_BATTERIES         Tag = 0x03000056
	BAT_REQ_OPEN_FET                    Tag = 0x03000060
	BAT_REQ_FET_STATE                   Tag = 0x03000061
	BAT_REQ_BATTERY_SOFT_ON             Tag = 0x03000062
	BAT_REQ_SET_BAT_VOLT_ADJUSTMENT     Tag = 0x03000063
	BAT_REQ_BAT_VOLT_ADJUSTMENT         Tag = 0x03000064
	BAT_REQ_BAT_VOLT_ADJ_READY_INDEX    Tag = 0x03000065
	BAT_REQ_DISCHARGE_UNTIL_EMPTY_STATE Tag = 0x03000092
	BAT_REQ_SET_DISCHARGE_UNTIL_EMPTY   Tag = 0x03000094
	BAT_REQ_CONTROL_STATE               Tag = 0x03000095
	BAT_REQ_INTERNAL_STATE              Tag = 0x03000096
	BAT_REQ_IS_BREAKER_OPEN             Tag = 0x03000097
	BAT_REQ_CLOSE_BREAKER               Tag = 0x03000098

	BAT_PARAM_BAT_VOLT_STATUS       Tag = 0x03400001
	BAT_PARAM_BAT_VOLT_TARGET_VALUE Tag = 0x03400002
	BAT_PARAM_BAT_VOLT_MIN_VOLTAGE  Tag = 0x03400003
	BAT_PARAM_BAT_VOLT_MAX_VOLTAGE  Tag = 0x03400004
	BAT_PARAM_BAT_VOLT_ENABLED      Tag = 0x03400005
	BAT_PARAM_BAT_NUMBER            Tag = 0x03400006
	BAT_PARAM_RACK_NUMBER           Tag = 0x03400007

	BAT_RSOC_REAL Tag = 0x0380000E
	BAT_ASOC      Tag = 0x0380000F
	BAT_FCC       Tag = 0x03800010
	BAT_RC        Tag = 0x03800011

	BAT_OPEN_BREAKER         Tag = 0x0380001C
	BAT_OPEN_BREAKER_CONFIRM Tag = 0x0380001D

	BAT_FIRMWARE_VERSION Tag = 0x0380001F

	BAT_UPDATE_STATUS             Tag = 0x03800022
	BAT_TIME_LAST_RESPONSE        Tag = 0x03800024
	BAT_MANUFACTURER_NAME         Tag = 0x03800025
	BAT_USABLE_CAPACITY           Tag = 0x03800026
	BAT_USABLE_REMAINING_CAPACITY Tag = 0x03800027
	BAT_SET_A1_DATA               Tag = 0x03800028
	BAT_DELETE_DCB_TYPE           Tag = 0x0380002A
	BAT_DCB_TYPE                  Tag = 0x0380002B
	BAT_CONTROL_CODE              Tag = 0x03800032
	BAT_BPM_STATUS                Tag = 0x03800033
	BAT_DCB_ERROR_LIST            Tag = 0x03800034

	BAT_SPECIFICATION               Tag = 0x03800043
	BAT_INTERNALS                   Tag = 0x03800044
	BAT_DESIGN_CAPACITY             Tag = 0x03800045
	BAT_DESIGN_VOLTAGE              Tag = 0x03800046
	BAT_CHARGE_HIGH_TEMP            Tag = 0x03800047
	BAT_CHARGE_LOW_TEMP             Tag = 0x03800048
	BAT_MANUFACTURE_DATE            Tag = 0x03800049
	BAT_SERIALNO                    Tag = 0x03800050
	BAT_DATA_TABLE_VERSION          Tag = 0x03800051
	BAT_PROTOCOL_VERSION            Tag = 0x03800052
	BAT_PCB_VERSION                 Tag = 0x03800053
	BAT_TOTAL_USE_TIME              Tag = 0x03800054
	BAT_TOTAL_DISCHARGE_TIME        Tag = 0x03800055
	BAT_AVAILABLE_BATTERIES         Tag = 0x03800057
	BAT_BATTERY_SPEC                Tag = 0x03800058
	BAT_INSTANCE_DESCRIPTOR         Tag = 0x03800059
	BAT_FET_STATE                   Tag = 0x03800061
	BAT_BATTERY_SOFT_ON             Tag = 0x03800062
	BAT_SET_BAT_VOLT_ADJUSTMENT     Tag = 0x03800063
	BAT_BAT_VOLT_ADJUSTMENT         Tag = 0x03800064
	BAT_BAT_VOLT_ADJ_READY_INDEX    Tag = 0x03800065
	BAT_DISCHARGE_UNTIL_EMPTY_STATE Tag = 0x03800092
	BAT_SET_DISCHARGE_UNTIL_EMPTY   Tag = 0x03800094
	BAT_CONTROL_STATE               Tag = 0x03800095
	BAT_INTERNAL_STATE              Tag = 0x03800096
	BAT_IS_BREAKER_OPEN             Tag = 0x03800097
	BAT_CLOSE_BREAKER               Tag = 0x03800098

	BAT_SPECIFIED_CAPACITY                   Tag = 0x03800125
	BAT_SPECIFIED_DSCHARGE_POWER             Tag = 0x03800126
	BAT_SPECIFIED_CHARGE_POWER               Tag = 0x03800127
	BAT_SPECIFIED_MAX_DCB_COUNT              Tag = 0x03800128
	BAT_ROLE                                 Tag = 0x03800129
	BAT_INTERNAL_CURRENT_AVG30               Tag = 0x03800130
	BAT_INTERNAL_MTV_AVG30                   Tag = 0x03800131
	BAT_INTERNAL_MAX_CHARGE_CURRENT          Tag = 0x03800132
	BAT_INTERNAL_MAX_DISCHARGE_CURRENT       Tag = 0x03800133
	BAT_INTERNAL_MAX_CHARGE_CURR_PER_DCB     Tag = 0x03800134
	BAT_INTERNAL_MAX_DISCHARGE_CURR_PER_DCB  Tag = 0x03800135
	BAT_INTERNAL_MAX_CHARGE_CURR_DATA_LOG    Tag = 0x03800136
	BAT_INTERNAL_MAX_DISCHARGE_CURR_DATA_LOG Tag = 0x03800137
)

// ----------------------
// NAMESPACE: BatteryDcdc
// 0x04xxxxxx
// ----------------------
const (
	// Beinhaltet alle Anfrage-TAGs, der Container MUSS einen Index enthalten
	DCDC_REQ_DATA Tag = 0x04040000
	// Index des angefragten Gerätes (0?n für die FBC Nr oder 0xFF für Gruppe), Kommt in der Anfrage und in der Antwort zum DATA-Tag vor
	DCDC_INDEX Tag = 0x04040001
	// Antwort mit allen Daten der REQ_DATA Anfrage
	DCDC_DATA Tag = 0x04840000
	// As parameter the index of the DCDC is required. Index 0 is for GroupController.
	DCDC_REQ_I_BAT            Tag = 0x04000001
	DCDC_REQ_U_BAT            Tag = 0x04000002
	DCDC_REQ_P_BAT            Tag = 0x04000003
	DCDC_REQ_I_DCL            Tag = 0x04000004
	DCDC_REQ_U_DCL            Tag = 0x04000005
	DCDC_REQ_P_DCL            Tag = 0x04000006
	DCDC_REQ_FIRMWARE_VERSION Tag = 0x04000008
	DCDC_REQ_FPGA_FIRMWARE    Tag = 0x04000009
	DCDC_REQ_SERIAL_NUMBER    Tag = 0x0400000A
	DCDC_REQ_BOARD_VERSION    Tag = 0x0400000B
	DCDC_REQ_FLASH_FILE_LIST  Tag = 0x0400000C
	DCDC_REQ_IS_FLASHING      Tag = 0x0400000E
	DCDC_REQ_FLASH            Tag = 0x0400000F
	DCDC_REQ_STATUS           Tag = 0x04000010
	DCDC_REQ_STATUS_AS_STRING Tag = 0x04000013
	DCDC_I_BAT                Tag = 0x04800001
	DCDC_U_BAT                Tag = 0x04800002
	DCDC_P_BAT                Tag = 0x04800003
	DCDC_I_DCL                Tag = 0x04800004
	DCDC_U_DCL                Tag = 0x04800005
	DCDC_P_DCL                Tag = 0x04800006
	DCDC_FIRMWARE_VERSION     Tag = 0x04800008
	DCDC_FPGA_FIRMWARE        Tag = 0x04800009
	DCDC_SERIAL_NUMBER        Tag = 0x0480000A
	DCDC_BOARD_VERSION        Tag = 0x0480000B
	DCDC_FLASH_FILE_LIST      Tag = 0x0480000C
	DCDC_FLASH_FILE           Tag = 0x0480000D
	DCDC_IS_FLASHING          Tag = 0x0480000E
	DCDC_FLASH                Tag = 0x0480000F
	DCDC_STATUS               Tag = 0x04800010
	DCDC_STATE                Tag = 0x04800011
	DCDC_SUBSTATE             Tag = 0x04800012
	DCDC_STATUS_AS_STRING     Tag = 0x04800013
	DCDC_STATE_AS_STRING      Tag = 0x04800014
	DCDC_SUBSTATE_AS_STRING   Tag = 0x04800015
	DCDC_REQ_DEVICE_STATE     Tag = 0x04060000
	// DEVICE_CONNECTED & DEVICE_WORKING & DEVICE_IN_SERVICE
	DCDC_DEVICE_STATE      Tag = 0x04860000
	DCDC_DEVICE_CONNECTED  Tag = 0x04860001
	DCDC_DEVICE_WORKING    Tag = 0x04860002
	DCDC_DEVICE_IN_SERVICE Tag = 0x04860003
	DCDC_GENERAL_ERROR     Tag = 0x04FFFFFF

	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	DCDC_REQ_SET_BAT_CONNECTED Tag = 0x04000000

	DCDC_REQ_SELF_TEST Tag = 0x04000007

	DCDC_REQ_GET_BAT_CONNECTED                      Tag = 0x04000016
	DCDC_REQ_SET_ZK_SYM_ACTIVE                      Tag = 0x04000017
	DCDC_REQ_GET_ZK_SYM_ACTIVE                      Tag = 0x04000018
	DCDC_REQ_SET_ISO_TEST_REQUESTED                 Tag = 0x04000019
	DCDC_REQ_GET_ISO_TEST_REQUESTED                 Tag = 0x0400001A
	DCDC_REQ_GET_ISO_TEST_SUCCESSFULLY_DONE         Tag = 0x0400001B
	DCDC_REQ_GET_BAT_ON                             Tag = 0x0400001C
	DCDC_REQ_GET_ZK_SYM_ON                          Tag = 0x0400001D
	DCDC_REQ_GET_ISO_ACTIVE                         Tag = 0x0400001E
	DCDC_REQ_GET_RELAYS_ON                          Tag = 0x0400001F
	DCDC_REQ_SET_ACTIVE_BAT_PRECHARGE_GOAL          Tag = 0x04000020
	DCDC_REQ_VERIFY_CORTEX                          Tag = 0x04000021
	DCDC_REQ_FLASH_FPGA                             Tag = 0x04000022
	DCDC_REQ_FLASH_FPGA_FILE_LIST                   Tag = 0x04000023
	DCDC_REQ_SELF_TEST_RESULT                       Tag = 0x04000024
	DCDC_REQ_FLASH_STATUS                           Tag = 0x04000025
	DCDC_REQ_GET_PARAMETER                          Tag = 0x04000026
	DCDC_REQ_SET_PARAMETER                          Tag = 0x04000027
	DCDC_REQ_SET_PID_DEBUG                          Tag = 0x04000028
	DCDC_REQ_COPY_RING_BUFFER                       Tag = 0x04000029
	DCDC_REQ_GET_PID_DEBUG                          Tag = 0x0400002A
	DCDC_REQ_PID_DEBUG_DATA                         Tag = 0x0400002B
	DCDC_REQ_GET_LAST_SET_ACTIVE_BAT_PRECHARGE_GOAL Tag = 0x0400002C
	DCDC_REQ_GET_DESIRED_IC                         Tag = 0x0400002D
	DCDC_REQ_GET_ISO_VALUE                          Tag = 0x0400002E
	DCDC_REQ_BROADCAST_BAT_LIMITS                   Tag = 0x0400002F
	DCDC_REQ_RING_BUFFER                            Tag = 0x04000030
	DCDC_REQ_FREE_RING_BUFFER                       Tag = 0x04000031
	DCDC_REQ_UNICAST_BAT_LIMITS                     Tag = 0x04000032
	DCDC_REQ_SET_POWERSAVE                          Tag = 0x04000033
	DCDC_REQ_GET_POWERSAVE                          Tag = 0x04000034
	DCDC_REQ_GET_HW_CHARACTERISTIC                  Tag = 0x04000035
	DCDC_REQ_U_DCL_DIFF                             Tag = 0x04000036
	DCDC_REQ_DERATING                               Tag = 0x04000037
	DCDC_REQ_DERATING_VALUE                         Tag = 0x04000038
	DCDC_REQ_TARGET_POWER                           Tag = 0x04000039
	DCDC_REQ_DCL_OPERATION_VOLTAGE                  Tag = 0x04000050
	DCDC_REQ_SET_DCL_OPERATION_VOLTAGE              Tag = 0x04000051
	DCDC_REQ_SET_POWER                              Tag = 0x04000071
	DCDC_REQ_SET_IDLE                               Tag = 0x04000072
	DCDC_REQ_HANDLE_ERRORS                          Tag = 0x04000073
	DCDC_REQ_CLEAR_ERRORS                           Tag = 0x04000074
	DCDC_REQ_SEND_COMMAND                           Tag = 0x04000075
	DCDC_REQ_BROADCAST_COMMAND                      Tag = 0x04000076
	DCDC_REQ_ERROR_PENDING                          Tag = 0x04000077
	DCDC_REQ_SET_PVI_TYPE                           Tag = 0x04000078
	DCDC_REQ_PVI_TYPE                               Tag = 0x04000079
	DCDC_REQ_ON_GRID                                Tag = 0x04000080
	DCDC_REQ_SET_ON_GRID                            Tag = 0x04000081
	DCDC_REQ_NEXT_SLAVE_STATE                       Tag = 0x04000082
	DCDC_REQ_ENABLE_NEXT_SLAVE                      Tag = 0x04000083
	DCDC_REQ_DCDC_TYPE                              Tag = 0x04000084
	DCDC_REQ_SEND_KICKSTART                         Tag = 0x04000085
	DCDC_REQ_IS_HEALTHY                             Tag = 0x04000086
	DCDC_REQ_SET_IS_MULTI_GC                        Tag = 0x04000087
	DCDC_REQ_SEND_BAT_PRECHARGE                     Tag = 0x04000088
	DCDC_REQ_SEND_STARTUP_TYPE                      Tag = 0x04000089
	DCDC_REQ_SEND_BAT_CHARGING_STATE                Tag = 0x0400008A
	DCDC_REQ_SEND_STANDBY_RELEASE                   Tag = 0x0400008B

	DCDC_PARAM_FLASH_PROGRESS        Tag = 0x04040010
	DCDC_PARAM_FLASH_TYPE            Tag = 0x04040011
	DCDC_PARAM_FLASHING_ACTIVE       Tag = 0x04040012
	DCDC_PARAM_FLASH_MODE            Tag = 0x04040013
	DCDC_PARAM_FLASH_FILE            Tag = 0x04040014
	DCDC_PARAM_CRC                   Tag = 0x04040015
	DCDC_PARAM_PARAMETER_BLOCK       Tag = 0x04040016
	DCDC_PARAM_PARAMETER_INDEX_FROM  Tag = 0x04040017
	DCDC_PARAM_PARAMETER_INDEX_UNTIL Tag = 0x04040018
	DCDC_PARAM_PARAMETER_VALUE       Tag = 0x04040019
	DCDC_PARAM_RING_BUFFER_ELEMENT   Tag = 0x04040030
	DCDC_PARAM_RB_ID                 Tag = 0x04040031
	DCDC_PARAM_RB_TIME               Tag = 0x04040032
	DCDC_PARAM_RB_I_BAT              Tag = 0x04040033
	DCDC_PARAM_RB_U_BAT              Tag = 0x04040034
	DCDC_PARAM_RB_I_DCL              Tag = 0x04040035
	DCDC_PARAM_RB_U_DCL              Tag = 0x04040036
	DCDC_PARAM_RB_MODE               Tag = 0x04040037
	DCDC_PARAM_RB_SUBSTATE           Tag = 0x04040038
	DCDC_PARAM_RB_SETPOINT           Tag = 0x04040039
	DCDC_PARAM_RB_INDEX_DCDC         Tag = 0x04040040
	DCDC_PARAM_RB_INDEX_FROM         Tag = 0x04040041
	DCDC_PARAM_RB_INDEX_UNTIL        Tag = 0x04040042
	DCDC_PARAM_DCL_OV_UPPER_VOLTAGE  Tag = 0x04040050
	DCDC_PARAM_DCL_OV_LOWER_VOLTAGE  Tag = 0x04040051
	DCDC_PARAM_DCL_OV_INDEX          Tag = 0x04040052
	DCDC_REQ_COUNT_HW_CONTROLLER     Tag = 0x04040060
	DCDC_REQ_ENABLE_FAN_TEST         Tag = 0x04040061
	DCDC_REQ_DISABLE_FAN_TEST        Tag = 0x04040062

	DCDC_SET_BAT_CONNECTED Tag = 0x04800000

	DCDC_SELF_TEST Tag = 0x04800007

	DCDC_GET_BAT_CONNECTED                      Tag = 0x04800016
	DCDC_SET_ZK_SYM_ACTIVE                      Tag = 0x04800017
	DCDC_GET_ZK_SYM_ACTIVE                      Tag = 0x04800018
	DCDC_SET_ISO_TEST_REQUESTED                 Tag = 0x04800019
	DCDC_GET_ISO_TEST_REQUESTED                 Tag = 0x0480001A
	DCDC_GET_ISO_TEST_SUCCESSFULLY_DONE         Tag = 0x0480001B
	DCDC_GET_BAT_ON                             Tag = 0x0480001C
	DCDC_GET_ZK_SYM_ON                          Tag = 0x0480001D
	DCDC_GET_ISO_ACTIVE                         Tag = 0x0480001E
	DCDC_GET_RELAYS_ON                          Tag = 0x0480001F
	DCDC_SET_ACTIVE_BAT_PRECHARGE_GOAL          Tag = 0x04800020
	DCDC_VERIFY_CORTEX                          Tag = 0x04800021
	DCDC_FLASH_FPGA                             Tag = 0x04800022
	DCDC_FLASH_FPGA_FILE_LIST                   Tag = 0x04800023
	DCDC_SELF_TEST_RESULT                       Tag = 0x04800024
	DCDC_FLASH_STATUS                           Tag = 0x04800025
	DCDC_GET_PARAMETER                          Tag = 0x04800026
	DCDC_SET_PARAMETER                          Tag = 0x04800027
	DCDC_COPY_RING_BUFFER                       Tag = 0x04800029
	DCDC_GET_PID_DEBUG                          Tag = 0x0480002A
	DCDC_PID_DEBUG_DATA                         Tag = 0x0480002B
	DCDC_GET_LAST_SET_ACTIVE_BAT_PRECHARGE_GOAL Tag = 0x0480002C
	DCDC_GET_DESIRED_IC                         Tag = 0x0480002D
	DCDC_GET_ISO_VALUE                          Tag = 0x0480002E
	DCDC_BROADCAST_BAT_LIMITS                   Tag = 0x0480002F
	DCDC_RING_BUFFER                            Tag = 0x04800030
	DCDC_FREED_RING_BUFFER                      Tag = 0x04800031
	DCDC_UNICAST_BAT_LIMITS                     Tag = 0x04800032
	DCDC_SET_POWERSAVE                          Tag = 0x04800033
	DCDC_GET_POWERSAVE                          Tag = 0x04800034
	DCDC_GET_HW_CHARACTERISTIC                  Tag = 0x04800035
	DCDC_U_DCL_DIFF                             Tag = 0x04800036
	DCDC_DERATING                               Tag = 0x04800037
	DCDC_DERATING_VALUE                         Tag = 0x04800038
	DCDC_TARGET_POWER                           Tag = 0x04800039
	DCDC_DCL_OPERATION_VOLTAGE                  Tag = 0x04800051
	DCDC_SET_POWER                              Tag = 0x04800071
	DCDC_SET_IDLE                               Tag = 0x04800072
	DCDC_HANDLE_ERRORS                          Tag = 0x04800073
	DCDC_CLEAR_ERRORS                           Tag = 0x04800074
	DCDC_SEND_COMMAND                           Tag = 0x04800075
	DCDC_BROADCAST_COMMAND                      Tag = 0x04800076
	DCDC_ERROR_PENDING                          Tag = 0x04800077
	DCDC_PVI_TYPE                               Tag = 0x04800079
	DCDC_ON_GRID                                Tag = 0x04800080
	DCDC_NEXT_SLAVE_STATE                       Tag = 0x04800082
	DCDC_DCDC_TYPE                              Tag = 0x04800084
	DCDC_SEND_KICKSTART                         Tag = 0x04800085
	DCDC_IS_HEALTHY                             Tag = 0x04800086
	DCDC_REQ_GET_IS_MULTI_GC                    Tag = 0x04800087
	DCDC_SEND_BAT_PRECHARGE                     Tag = 0x04800088
	DCDC_SEND_STARTUP_TYPE                      Tag = 0x04800089
	DCDC_SEND_BAT_CHARGING_STATE                Tag = 0x0480008A
	DCDC_SEND_STANDBY_RELEASE                   Tag = 0x0480008B

	DCDC_COUNT_HW_CONTROLLER Tag = 0x04840060

	DCDC_GET_IS_MULTI_GC Tag = 0x04860004
	DCDC_SET_IS_MULTI_GC Tag = 0x04860005
)

// ---------------------
// NAMESPACE: PowerMeter
// 0x05xxxxxx
// ---------------------
const (
	// Beinhaltet alle Anfrage-TAGs, der Container MUSS einen Index enthalten
	PM_REQ_DATA Tag = 0x05040000
	// Index des angefragten Gerätes (0?x), muss in Anfrage und ist in Antwort enthalten
	PM_INDEX Tag = 0x05040001
	// Antwort mit allen Daten der REQ_DATA Anfrage
	PM_DATA Tag = 0x05840000
	// Kann nur innerhalb eines REQ_PM_DATA Container verwendet werden!
	PM_REQ_POWER_L1 Tag = 0x05000001
	// Kann nur innerhalb eines REQ_PM_DATA Container verwendet werden!
	PM_REQ_POWER_L2 Tag = 0x05000002
	// Kann nur innerhalb eines REQ_PM_DATA Container verwendet werden!
	PM_REQ_POWER_L3 Tag = 0x05000003
	// Kann nur innerhalb eines REQ_PM_DATA Container verwendet werden!
	PM_REQ_ACTIVE_PHASES Tag = 0x05000004
	// Kann nur innerhalb eines REQ_PM_DATA Container verwendet werden!
	PM_REQ_MODE Tag = 0x05000005
	// Kann nur innerhalb eines REQ_PM_DATA Container verwendet werden!
	PM_REQ_ENERGY_L1 Tag = 0x05000006
	// Kann nur innerhalb eines REQ_PM_DATA Container verwendet werden!
	PM_REQ_ENERGY_L2 Tag = 0x05000007
	// Kann nur innerhalb eines REQ_PM_DATA Container verwendet werden!
	PM_REQ_ENERGY_L3 Tag = 0x05000008
	// Kann nur innerhalb eines REQ_PM_DATA Container verwendet werden!
	PM_REQ_DEVICE_ID Tag = 0x05000009
	// Kann nur innerhalb eines REQ_PM_DATA Container verwendet werden!
	PM_REQ_ERROR_CODE Tag = 0x0500000A
	// Kann nur innerhalb eines REQ_PM_DATA Container verwendet werden!
	PM_REQ_SET_PHASE_ELIMINATION Tag = 0x0500000B
	// Kann nur innerhalb eines REQ_PM_DATA Container verwendet werden!
	PM_REQ_GET_PHASE_ELIMINATION Tag = 0x05000018
	// Kann nur innerhalb eines REQ_PM_DATA Container verwendet werden!
	PM_REQ_FIRMWARE_VERSION Tag = 0x0500000C
	// Kann nur innerhalb eines REQ_PM_DATA Container verwendet werden!
	PM_REQ_VOLTAGE_L1 Tag = 0x05000011
	// Kann nur innerhalb eines REQ_PM_DATA Container verwendet werden!
	PM_REQ_VOLTAGE_L2 Tag = 0x05000012
	// Kann nur innerhalb eines REQ_PM_DATA Container verwendet werden!
	PM_REQ_VOLTAGE_L3 Tag = 0x05000013
	// Kann nur innerhalb eines REQ_PM_DATA Container verwendet werden!
	PM_REQ_TYPE Tag = 0x05000014
	// Current power on L1
	PM_POWER_L1 Tag = 0x05800001
	// Current power on L2
	PM_POWER_L2 Tag = 0x05800002
	// Current power on L3
	PM_POWER_L3 Tag = 0x05800003
	// just the three lowest bits of activePhases are used to define
	// what phase is switched on. If the lowest bit is 1 phase1 is active
	// if the lowest bit is 0 phase 1 is inactive ...
	//     static const unsigned char PHASE_1 = 1
	//     static const unsigned char PHASE_2 = 2
	//     static const unsigned char PHASE_3 = 4
	// f.e. if active Phases = 7 -> all phases are active
	PM_ACTIVE_PHASES Tag = 0x05800004
	// used to identify the error bit, if error code is available mode = ERROR_ACTIVE_MODE. ACTIVE_MODE else. Ignore all other modes.
	//     static const unsigned char ACTIVE_MODE = 0
	//     static const unsigned char PASSIVE_MODE = 1
	//     static const unsigned char DIAGNOSE_MODE = 2
	//     static const unsigned char ERROR_ACTIVE_MODE = 3
	//     static const unsigned char ERROR_PASSIVE_MODE = 4
	PM_MODE Tag = 0x05800005
	// Energy counter L1
	PM_ENERGY_L1 Tag = 0x05800006
	// Energy counter L2
	PM_ENERGY_L2 Tag = 0x05800007
	// Energy counter L3
	PM_ENERGY_L3 Tag = 0x05800008
	// ID of that device
	PM_DEVICE_ID Tag = 0x05800009
	// Last reported error code (see mode if error has relevance)
	PM_ERROR_CODE            Tag = 0x0580000A
	PM_SET_PHASE_ELIMINATION Tag = 0x0580000B
	PM_GET_PHASE_ELIMINATION Tag = 0x05800018
	PM_FIRMWARE_VERSION      Tag = 0x0580000C
	// Current voltage on L1 0 if not supported, use ACTIVE_PHASES to detect a broken phase
	PM_VOLTAGE_L1 Tag = 0x05800011
	// Current voltage on L2
	PM_VOLTAGE_L2 Tag = 0x05800012
	// Current voltage on L3
	PM_VOLTAGE_L3 Tag = 0x05800013
	// Leistungsmesser Typ:
	//  PM_TYPE_UNDEFINED               0
	//  PM_TYPE_ROOT                    1
	//  PM_TYPE_ADDITIONAL              2
	//  PM_TYPE_ADDITIONAL_PRODUCTION   3
	//  PM_TYPE_ADDITIONAL_CONSUMPTION  4
	//  PM_TYPE_FARM                    5
	//  PM_TYPE_UNUSED                  6
	//  PM_TYPE_WALLBOX                 7
	// 	PM_TYPE_FARM_ADDITIONAL         8
	PM_TYPE Tag = 0x05800014
	// kann nur innerhalb eines REQ_PM_COMM_STATE Container verwendet werden!
	PM_CS_START_TIME Tag = 0x05800051
	// kann nur innerhalb eines REQ_PM_COMM_STATE Container verwendet werden!
	PM_CS_LAST_TIME Tag = 0x05800052
	// kann nur innerhalb eines REQ_PM_COMM_STATE Container verwendet werden!
	PM_CS_SUCC_FRAMES_ALL Tag = 0x05800053
	// kann nur innerhalb eines REQ_PM_COMM_STATE Container verwendet werden!
	PM_CS_SUCC_FRAMES_100 Tag = 0x05800054
	// kann nur innerhalb eines REQ_PM_COMM_STATE Container verwendet werden!
	PM_CS_EXP_FRAMES_ALL Tag = 0x05800055
	// kann nur innerhalb eines REQ_PM_COMM_STATE Container verwendet werden!
	PM_CS_EXP_FRAMES_100 Tag = 0x05800056
	// kann nur innerhalb eines REQ_PM_COMM_STATE Container verwendet werden!
	PM_CS_ERR_FRAMES_ALL Tag = 0x05800057
	// kann nur innerhalb eines REQ_PM_COMM_STATE Container verwendet werden!
	PM_CS_ERR_FRAMES_100 Tag = 0x05800058
	// kann nur innerhalb eines REQ_PM_COMM_STATE Container verwendet werden!
	PM_CS_UNK_FRAMES Tag = 0x05800059
	// kann nur innerhalb eines REQ_PM_COMM_STATE Container verwendet werden!
	PM_CS_ERR_FRAME     Tag = 0x0580005A
	PM_REQ_DEVICE_STATE Tag = 0x05060000
	// DEVICE_CONNECTED & DEVICE_WORKING & DEVICE_IN_SERVICE
	PM_DEVICE_STATE      Tag = 0x05860000
	PM_DEVICE_CONNECTED  Tag = 0x05860001
	PM_DEVICE_WORKING    Tag = 0x05860002
	PM_DEVICE_IN_SERVICE Tag = 0x05860003
	PM_GENERAL_ERROR     Tag = 0x05FFFFFF

	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	PM_REQ_SET_FOR_EMERGENCY_TEST Tag = 0x0500000D
	PM_REQ_IS_CAN_SILENCE         Tag = 0x0500000E
	PM_REQ_MAX_PHASE_POWER        Tag = 0x0500000F
	PM_REQ_GET_EXTERN_LOAD        Tag = 0x05000010

	PM_REQ_SET_TYPE Tag = 0x05000015

	PM_REQ_COMM_STATE        Tag = 0x05000050
	PM_REQ_CONNECTED_DEVICES Tag = 0x050000A0
	PM_REQ_SET_DEVICE_NAME   Tag = 0x050000B0
	PM_REQ_DEVICE_NAME       Tag = 0x050000B1
	PM_REQ_SET_EXTERN_LOAD   Tag = 0x050000B3
	PM_REQ_INJECT_DATA       Tag = 0x050000B4
	PM_REQ_SET_CAN_ID        Tag = 0x050000B5

	PM_ID1 Tag = 0x05400001
	PM_ID2 Tag = 0x05400002
	PM_ID3 Tag = 0x05400003

	PM_SET_FOR_EMERGENCY_TEST Tag = 0x0580000D
	PM_IS_CAN_SILENCE         Tag = 0x0580000E
	PM_MAX_PHASE_POWER        Tag = 0x0580000F
	PM_GET_EXTERN_LOAD        Tag = 0x05800010

	PM_COMM_STATE Tag = 0x05800050

	PM_CONNECTED_DEVICES Tag = 0x058000A0
	PM_CONNECTED_DEVICE  Tag = 0x058000A1
	PM_DEVICE_NAME       Tag = 0x058000B1
	PM_SET_EXTERN_LOAD   Tag = 0x058000B3
	PM_INJECT_DATA       Tag = 0x058000B4
	PM_SET_CAN_ID        Tag = 0x058000B5
)

// -------------------
// NAMESPACE: Database
// 0x06xxxxxx
// -------------------
const (
	// Muss die TAGs DB_REQ_HISTORY_TIME_START, DB_REQ_HISTORY_TIME_INTERVAL, DB_REQ_HISTORY_TIME_SPAN enthalten
	DB_REQ_HISTORY_DATA_DAY      Tag = 0x06000100
	DB_REQ_HISTORY_TIME_START    Tag = 0x06000101
	DB_REQ_HISTORY_TIME_INTERVAL Tag = 0x06000102
	DB_REQ_HISTORY_TIME_SPAN     Tag = 0x06000103
	// Muss die TAGs DB_REQ_HISTORY_TIME_START, DB_REQ_HISTORY_TIME_INTERVAL, DB_REQ_HISTORY_TIME_SPAN enthalten
	DB_REQ_HISTORY_DATA_WEEK Tag = 0x06000200
	// Muss die TAGs DB_REQ_HISTORY_TIME_START, DB_REQ_HISTORY_TIME_INTERVAL, DB_REQ_HISTORY_TIME_SPAN enthalten
	DB_REQ_HISTORY_DATA_MONTH Tag = 0x06000300
	// Muss die TAGs DB_REQ_HISTORY_TIME_START, DB_REQ_HISTORY_TIME_INTERVAL, DB_REQ_HISTORY_TIME_SPAN enthalten
	DB_REQ_HISTORY_DATA_YEAR Tag = 0x06000400
	// Die Summe zwischen der Energien über den Zeitraum
	DB_SUM_CONTAINER Tag = 0x06800010
	// Meist mehr als einer von diesen Kontainern in einem HISTORY_DATA Kontainer
	DB_VALUE_CONTAINER Tag = 0x06800020
	// Diagrammposition in Prozent
	DB_GRAPH_INDEX         Tag = 0x06800001
	DB_BAT_POWER_IN        Tag = 0x06800002
	DB_BAT_POWER_OUT       Tag = 0x06800003
	DB_DC_POWER            Tag = 0x06800004
	DB_GRID_POWER_IN       Tag = 0x06800005
	DB_GRID_POWER_OUT      Tag = 0x06800006
	DB_CONSUMPTION         Tag = 0x06800007
	DB_PM_0_POWER          Tag = 0x06800008
	DB_PM_1_POWER          Tag = 0x06800009
	DB_BAT_CHARGE_LEVEL    Tag = 0x0680000A
	DB_BAT_CYCLE_COUNT     Tag = 0x0680000B
	DB_CONSUMED_PRODUCTION Tag = 0x0680000C
	DB_AUTARKY             Tag = 0x0680000D
	// Beinhaltet die Container DB_SUM_CONTAINER, VALUE_CONTAINER
	DB_HISTORY_DATA_DAY Tag = 0x06800100
	// Beinhaltet die Container DB_SUM_CONTAINER, VALUE_CONTAINER
	DB_HISTORY_DATA_WEEK Tag = 0x06800200
	// Beinhaltet die Container DB_SUM_CONTAINER, VALUE_CONTAINER
	DB_HISTORY_DATA_MONTH Tag = 0x06800300
	// Beinhaltet die Container DB_SUM_CONTAINER, VALUE_CONTAINER
	DB_HISTORY_DATA_YEAR Tag = 0x06800400
	DB_PAR_TIME_MIN      Tag = 0x06B00000
	DB_PAR_TIME_MAX      Tag = 0x06B00001
	DB_PARAM_ROW         Tag = 0x06B00002
	DB_PARAM_COLUMN      Tag = 0x06B00003
	DB_PARAM_INDEX       Tag = 0x06B00004
	DB_PARAM_VALUE       Tag = 0x06B00005
	DB_PARAM_MAX_ROWS    Tag = 0x06B00006
	DB_PARAM_TIME        Tag = 0x06B00007
	DB_PARAM_VERSION     Tag = 0x06B00008
	DB_PARAM_HEADER      Tag = 0x06B00009

	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	DB_REQ_TEC_ALL            Tag = 0x0600000F
	DB_REQ_TEC_WALLBOX_VALUES Tag = 0x06000016
	DB_REQ_TEC_ROOT_EM_VALUE  Tag = 0x06000017
	DB_REQ_TEC_EM_VALUES      Tag = 0x06000018
	DB_REQ_TEC_BAT            Tag = 0x06000027
	DB_REQ_TEC_PRODUCTION     Tag = 0x06000028
	DB_REQ_TEC_CONSUMPTION    Tag = 0x06000029
	DB_REQ_TEC_DC_POWER       Tag = 0x0600002A

	DB_REQ_HISTORY_UTC_TIME_START Tag = 0x06000104
	DB_REQ_ENERGY_COUNTERS        Tag = 0x06000105
	DB_REQ_SET_IDLE               Tag = 0x0600010A
	DB_REQ_IS_IDLE                Tag = 0x0600010B

	DB_REQ_SYNC_HIST               Tag = 0x06000500
	DB_REQ_VACUUM_HIST             Tag = 0x06000501
	DB_REQ_SYNC_BPU                Tag = 0x06000502
	DB_REQ_VACUUM_BPU              Tag = 0x06000503
	DB_REQ_SYNC_DCB                Tag = 0x06000504
	DB_REQ_VACUUM_DBC              Tag = 0x06000505
	DB_REQ_SYNC_BPU_CONF           Tag = 0x06000506
	DB_REQ_VACUUM_BPU_CONF         Tag = 0x06000507
	DB_REQ_SYNC_DCB_CONF           Tag = 0x06000508
	DB_REQ_VACUUM_DBC_CONF         Tag = 0x06000509
	DB_REQ_SYNC_WALLBOX            Tag = 0x0600050A
	DB_REQ_VACUUM_WALLBOX          Tag = 0x0600050B
	DB_REQ_SYNC_PV_DEBUG           Tag = 0x0600050C
	DB_REQ_VACUUM_PV_DEBUG         Tag = 0x0600050D
	DB_REQ_SYNC_CONFIG             Tag = 0x0600050E
	DB_REQ_VACUUM_CONFIG           Tag = 0x0600050F
	DB_REQ_SET_SYNC_TIME           Tag = 0x06000510
	DB_REQ_PVI_DIAL_RECORDINGS     Tag = 0x06000511
	DB_REQ_SYNC_BAT_DIAGNOSE       Tag = 0x06000512
	DB_REQ_VACUUM_BAT_DIAGNOSE     Tag = 0x06000513
	DB_REQ_SYNC_EXT_LG             Tag = 0x06000514
	DB_REQ_VACUUM_EXT_LG           Tag = 0x06000515
	DB_REQ_CLEAN_DATABASE          Tag = 0x06000516
	DB_REQ_SYNC_ALL                Tag = 0x06000517
	DB_REQ_RESET_TIMESTAMP         Tag = 0x06000518
	DB_TEC_WALLBOX_VALUE           Tag = 0x06400001
	DB_TEC_WALLBOX_ENERGYALL       Tag = 0x06400002
	DB_TEC_WALLBOX_WB_ENERGY_SOLAR Tag = 0x06400003
	DB_TEC_WALLBOX_INDEX           Tag = 0x06400004
	DB_TEC_EM_VALUE                Tag = 0x06400005
	DB_TEC_EM_ENERGY_L1            Tag = 0x06400006
	DB_TEC_EM_ENERGY_L2            Tag = 0x06400007
	DB_TEC_EM_ENERGY_L3            Tag = 0x06400008
	DB_TEC_EM_NET_IN_L1            Tag = 0x06400009
	DB_TEC_EM_NET_IN_L2            Tag = 0x0640000A
	DB_TEC_EM_NET_IN_L3            Tag = 0x0640000B
	DB_TEC_EM_NET_OUT_L1           Tag = 0x0640000C
	DB_TEC_EM_NET_OUT_L2           Tag = 0x0640000D
	DB_TEC_EM_NET_OUT_L3           Tag = 0x0640000E
	DB_TEC_EM_NET_IN               Tag = 0x0640000F
	DB_TEC_EM_NET_OUT              Tag = 0x06400010
	DB_TEC_EM_INDEX                Tag = 0x06400011

	DB_PRODUCTION_POWER Tag = 0x0680000E
	DB_TEC_ALL          Tag = 0x0680000F

	DB_TEC_BAT_POWER_IN     Tag = 0x06800011
	DB_TEC_BAT_POWER_OUT    Tag = 0x06800012
	DB_TEC_PRODUCTION_L1    Tag = 0x06800013
	DB_TEC_PRODUCTION_L2    Tag = 0x06800014
	DB_TEC_PRODUCTION_L3    Tag = 0x06800015
	DB_TEC_WALLBOX_VALUES   Tag = 0x06800016
	DB_TEC_ROOT_EM_VALUE    Tag = 0x06800017
	DB_TEC_EM_VALUES        Tag = 0x06800018
	DB_TEC_BAT_CHARGE_LEVEL Tag = 0x06800019
	DB_TEC_BAT_CURRENT_IN   Tag = 0x0680001A
	DB_TEC_BAT_CURRENT_OUT  Tag = 0x0680001B
	DB_TEC_CONSUMPTION_L1   Tag = 0x0680001C
	DB_TEC_CONSUMPTION_L2   Tag = 0x0680001D
	DB_TEC_CONSUMPTION_L3   Tag = 0x0680001E
	DB_TEC_DC_POWER_S1      Tag = 0x0680001F

	DB_TEC_DC_POWER_S2           Tag = 0x06800021
	DB_TEC_DC_POWER_S3           Tag = 0x06800022
	DB_TEC_ACCURACY              Tag = 0x06800023
	DB_TEC_ACCURACY_LM           Tag = 0x06800024
	DB_TEC_ACCURACY_WB           Tag = 0x06800025
	DB_TEC_BAT_ASOC_CHARGE_LEVEL Tag = 0x06800026
	DB_TEC_BAT                   Tag = 0x06800027
	DB_TEC_PRODUCTION            Tag = 0x06800028
	DB_TEC_CONSUMPTION           Tag = 0x06800029
	DB_TEC_DC_POWER              Tag = 0x0680002A

	DB_ENERGY_COUNTERS Tag = 0x06800105
	DB_SET_IDLE        Tag = 0x0680010A
	DB_IS_IDLE         Tag = 0x0680010B

	DB_SYNC_HIST           Tag = 0x06800500
	DB_VACUUM_HIST         Tag = 0x06800501
	DB_SYNC_BPU            Tag = 0x06800502
	DB_VACUUM_BPU          Tag = 0x06800503
	DB_SYNC_DCB            Tag = 0x06800504
	DB_VACUUM_DCB          Tag = 0x06800505
	DB_SYNC_BPU_CONF       Tag = 0x06800506
	DB_VACUUM_BPU_CONF     Tag = 0x06800507
	DB_SYNC_DCB_CONF       Tag = 0x06800508
	DB_VACUUM_DCB_CONF     Tag = 0x06800509
	DB_SYNC_WALLBOX        Tag = 0x0680050A
	DB_VACUUM_WALLBOX      Tag = 0x0680050B
	DB_SYNC_PV_DEBUG       Tag = 0x0680050C
	DB_VACUUM_PV_DEBUG     Tag = 0x0680050D
	DB_SYNC_CONFIG         Tag = 0x0680050E
	DB_VACUUM_CONFIG       Tag = 0x0680050F
	DB_SET_SYNC_TIME       Tag = 0x06800510
	DB_PVI_DIAL_RECORDINGS Tag = 0x06800511
	DB_SYNC_BAT_DIAGNOSE   Tag = 0x06800512
	DB_VACUUM_BAT_DIAGNOSE Tag = 0x06800513
	DB_SYNC_EXT_LG         Tag = 0x06800514
	DB_VACUUM_EXT_LG       Tag = 0x06800515
	DB_SYNC_ALL            Tag = 0x06800517
	DB_RESET_TIMESTAMP     Tag = 0x06800518

	DB_PARAM_PRODUCTION_L1 Tag = 0x06B00010
	DB_PARAM_PRODUCTION_L2 Tag = 0x06B00011
	DB_PARAM_PRODUCTION_L3 Tag = 0x06B00012
	DB_PARAM_DC_POWER_S1   Tag = 0x06B00013
	DB_PARAM_DC_POWER_S2   Tag = 0x06B00014
	DB_PARAM_DC_POWER_S3   Tag = 0x06B00015
)

// --------------
// NAMESPACE: FMS
// 0x07xxxxxx
// --------------
const ()

// --------------
// NAMESPACE: SRV
// 0x08xxxxxx
// undocumented
// --------------
const (
	SRV_REQ_IS_ONLINE Tag = 0x08000001
	SRV_IS_ONLINE     Tag = 0x08800001
	SRV_REQ_ADD_USER  Tag = 0x08000002
	SRV_ADD_USER      Tag = 0x08800002
	SRV_GENERAL_ERROR Tag = 0x08FFFFFF

	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	SRV_REQ_SET_LOCAL_USER     Tag = 0x08000003
	SRV_REQ_CREATE_USER        Tag = 0x08000004
	SRV_REQ_CREATE_USER_STATUS Tag = 0x08000005
	SRV_NAME                   Tag = 0x08400001
	SRV_STREET                 Tag = 0x08400002
	SRV_STREET_NO              Tag = 0x08400003
	SRV_POSTCODE               Tag = 0x08400004
	SRV_CITY                   Tag = 0x08400005
	SRV_COUNTRY                Tag = 0x08400006
	SRV_FON                    Tag = 0x08400007
	SRV_E_MAIL                 Tag = 0x08400008
	SRV_SURNAME                Tag = 0x08400009

	SRV_SET_LOCAL_USER     Tag = 0x08800003
	SRV_CREATE_USER        Tag = 0x08800004
	SRV_CREATE_USER_STATUS Tag = 0x08800005
)

// --------------
// NAMESPACE: HA
// 0x09xxxxxx
// undocumented
// --------------
const (
	HA_REQ_DATAPOINT_LIST  Tag = 0x09000001
	HA_REQ_ACTUATOR_STATES Tag = 0x09000010
	// Beinhaltet
	// DATAPOINT_INDEX, DATAPOINT_TYPE, DATAPOINT_NAME, DATAPOINT_NAME,
	// DATAPOINT_DESCRIPTIONS, DATAPOINT_DESCRIPTION_VALUE, DATAPOINT_DESCRIPTION_VALUE
	HA_REQ_ADD_ACTUATOR    Tag = 0x09000020
	HA_REQ_REMOVE_ACTUATOR Tag = 0x09000030
	// Beinhaltet DATAPOINT_INDEX, REQ_COMMAND
	HA_REQ_COMMAND_ACTUATOR Tag = 0x09000040
	HA_REQ_COMMAND          Tag = 0x09000041
	// Beinhaltet DATAPOINT_INDEX, DATAPOINT_DESCRIPTIONS
	HA_REQ_DESCRIPTIONS_CHANGE          Tag = 0x09000050
	HA_REQ_CONFIGURATION_CHANGE_COUNTER Tag = 0x09000060
	HA_CONFIGURATION_CHANGE_COUNTER     Tag = 0x09800060
	HA_DATAPOINT_LIST                   Tag = 0x09800001
	// Beinhaltet DATAPOINT_INDEX, DATAPOINT_TYPE, DATAPOINT_NAME, DATAPOINT_DESCRIPTIONS
	HA_DATAPOINT       Tag = 0x09800002
	HA_DATAPOINT_INDEX Tag = 0x09800003
	HA_DATAPOINT_TYPE  Tag = 0x09800004
	HA_DATAPOINT_NAME  Tag = 0x09800005
	// '1' - ON / '2' - OFF / '?' - Unknown / 'G' - Group
	HA_DATAPOINT_STATE Tag = 0x09800011
	// Zeitstempel der letzten Statenachricht
	HA_DATAPOINT_STATE_TIMESTAMP Tag = 0x09800013
	// Verschiedene Bedeutungen je nach DATAPOINTTYPE (z.B:Dimmer Prozente )
	HA_DATAPOINT_STATE_VALUE Tag = 0x09800014
	// Quality:
	//  0x00|0xFF    ///< Not Available, no information
	//  0x01               ///< Empty
	//  0x02               ///< Change it
	//  0x03               ///< Medium
	//  0x04               ///< Good
	//  0x05               ///< New
	//  0x10               ///< mains-powered
	HA_DATAPOINT_SUPPLY_QUALITY Tag = 0x09800015
	// Quality:
	//  0xFF              ///< Not Available, no information
	//  0x00 - 0x64  ///<Wert in Prozent
	HA_DATAPOINT_SIGNAL_QUALITY Tag = 0x09800016
	// Mode:
	//  'A' - Automatic
	//  'M' - Manual
	HA_DATAPOINT_MODE Tag = 0x09800012
	// Beinhaltet mehrere HA_DATAPOINT_DESCRIPTION
	HA_DATAPOINT_DESCRIPTIONS      Tag = 0x09800006
	HA_DATAPOINT_DESCRIPTION       Tag = 0x09800007
	HA_DATAPOINT_DESCRIPTION_NAME  Tag = 0x09800008
	HA_DATAPOINT_DESCRIPTION_VALUE Tag = 0x09800009
	// Beinhaltet eine Liste mit DATAPOINT Container
	HA_ACTUATOR_STATES     Tag = 0x09800010
	HA_ADD_ACTUATOR        Tag = 0x09800020
	HA_REMOVE_ACTUATOR     Tag = 0x09800030
	HA_COMMAND_ACTUATOR    Tag = 0x09800040
	HA_DESCRIPTIONS_CHANGE Tag = 0x09800050
	HA_REQ_DEVICE_STATE    Tag = 0x09060000
	// DEVICE_CONNECTED & DEVICE_WORKING & DEVICE_IN_SERVICE
	HA_DEVICE_STATE      Tag = 0x09860000
	HA_DEVICE_CONNECTED  Tag = 0x09860001
	HA_DEVICE_WORKING    Tag = 0x09860002
	HA_DEVICE_IN_SERVICE Tag = 0x09860003
	HA_GENERAL_ERROR     Tag = 0x09FFFFFF

	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	HA_REQ_POSSIBLE_POWER_METERS Tag = 0x09000070
	HA_REQ_POSSIBLE_ANALOG_MODES Tag = 0x09000080

	HA_POSSIBLE_POWER_METERS Tag = 0x09800070
	HA_POSSIBLE_POWER_METER  Tag = 0x09800071
	HA_POSSIBLE_ANALOG_MODES Tag = 0x09800080
	HA_POSSIBLE_ANALOG_MODE  Tag = 0x09800081
)

// --------------
// NAMESPACE: INFO
// 0x0Axxxxxx
// undocumented
// --------------
const (
	INFO_REQ_SERIAL_NUMBER       Tag = 0x0A000001
	INFO_REQ_PRODUCTION_DATE     Tag = 0x0A000002
	INFO_REQ_MODULES_SW_VERSIONS Tag = 0x0A000003
	INFO_REQ_A35_SERIAL_NUMBER   Tag = 0x0A000007
	INFO_REQ_IP_ADDRESS          Tag = 0x0A000008
	INFO_REQ_SUBNET_MASK         Tag = 0x0A000009
	INFO_REQ_MAC_ADDRESS         Tag = 0x0A00000A
	INFO_REQ_GATEWAY             Tag = 0x0A00000B
	INFO_REQ_DNS                 Tag = 0x0A00000C
	INFO_REQ_DHCP_STATUS         Tag = 0x0A00000D
	INFO_REQ_TIME                Tag = 0x0A00000E
	INFO_REQ_UTC_TIME            Tag = 0x0A00000F
	INFO_REQ_TIME_ZONE           Tag = 0x0A000010
	INFO_REQ_INFO                Tag = 0x0A000011
	INFO_REQ_SET_IP_ADDRESS      Tag = 0x0A000012
	INFO_REQ_SET_SUBNET_MASK     Tag = 0x0A000013
	INFO_REQ_SET_DHCP_STATUS     Tag = 0x0A000014
	INFO_REQ_SET_GATEWAY         Tag = 0x0A000015
	INFO_REQ_SET_DNS             Tag = 0x0A000016
	INFO_REQ_SET_TIME_ZONE       Tag = 0x0A000018
	INFO_REQ_SW_RELEASE          Tag = 0x0A000019
	INFO_SERIAL_NUMBER           Tag = 0x0A800001
	INFO_PRODUCTION_DATE         Tag = 0x0A800002
	// Beinhaltet eine Liste mit INFO_MODULE_SW_VERSION Containern
	INFO_MODULES_SW_VERSIONS Tag = 0x0A800003
	// Beinhaltet die TAGs INFO_MODULE und INFO_VERSION
	INFO_MODULE_SW_VERSION Tag = 0x0A800004
	INFO_MODULE            Tag = 0x0A800005
	INFO_VERSION           Tag = 0x0A800006
	INFO_A35_SERIAL_NUMBER Tag = 0x0A800007
	INFO_IP_ADDRESS        Tag = 0x0A800008
	INFO_SUBNET_MASK       Tag = 0x0A800009
	INFO_MAC_ADDRESS       Tag = 0x0A80000A
	INFO_GATEWAY           Tag = 0x0A80000B
	INFO_DNS               Tag = 0x0A80000C
	INFO_DHCP_STATUS       Tag = 0x0A80000D
	INFO_TIME              Tag = 0x0A80000E
	INFO_UTC_TIME          Tag = 0x0A80000F
	INFO_TIME_ZONE         Tag = 0x0A800010
	// Beinhaltet die TAGs INFO_SERIAL_NUMBER, INFO_PRODUCTION_DATE, INFO_MAC_ADDRESS
	INFO_INFO            Tag = 0x0A800011
	INFO_SET_IP_ADDRESS  Tag = 0x0A800012
	INFO_SET_SUBNET_MASK Tag = 0x0A800013
	INFO_SET_DHCP_STATUS Tag = 0x0A800014
	INFO_SET_GATEWAY     Tag = 0x0A800015
	INFO_SET_DNS         Tag = 0x0A800016
	INFO_SET_TIME        Tag = 0x0A800017
	INFO_SET_TIME_ZONE   Tag = 0x0A800018
	INFO_SW_RELEASE      Tag = 0x0A800019
	INFO_GENERAL_ERROR   Tag = 0x0AFFFFFF

	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	INFO_REQ_SET_GUI_TARGET                   Tag = 0x0A00001A
	INFO_REQ_GUI_TARGET                       Tag = 0x0A00001B
	INFO_REQ_PLATFORM_TYPE                    Tag = 0x0A00001C
	INFO_REQ_IS_CALIBRATED                    Tag = 0x0A00001D
	INFO_REQ_CALIBRATION_CHECK                Tag = 0x0A00001E
	INFO_REQ_RESET_CALIBRATION                Tag = 0x0A00001F
	INFO_REQ_HW_TIME                          Tag = 0x0A000020
	INFO_REQ_SET_TIME_UTC                     Tag = 0x0A000021
	INFO_REQ_SET_HW_TIME                      Tag = 0x0A000022
	INFO_REQ_SET_FACILITY                     Tag = 0x0A000023
	INFO_REQ_GET_FACILITY                     Tag = 0x0A000024
	INFO_REQ_GET_FS_USAGE                     Tag = 0x0A000025
	INFO_REQ_UPNP_STATUS                      Tag = 0x0A000037
	INFO_REQ_SET_UPNP_STATUS                  Tag = 0x0A000038
	INFO_REQ_IS_OVP_POSSIBLE                  Tag = 0x0A000039
	INFO_REQ_IS_RSCP_PASSWORD_SET             Tag = 0x0A00003A
	INFO_REQ_SET_EULA_VERSION                 Tag = 0x0A00003B
	INFO_REQ_SET_EULA_CHECKSUM                Tag = 0x0A00003C
	INFO_REQ_SET_WAIT_FOR_EULA                Tag = 0x0A00003D
	INFO_REQ_ASSEMBLY_SERIAL_NUMBER           Tag = 0x0A00003E
	INFO_REQ_SET_UUID                         Tag = 0x0A00003F
	INFO_REQ_GET_UUID                         Tag = 0x0A000040
	INFO_REQ_GET_SUID                         Tag = 0x0A000041
	INFO_REQ_IS_OVP2_POSSIBLE                 Tag = 0x0A000042
	INFO_LC_REQ_SET_TIME                      Tag = 0x0A000043
	INFO_LC_REQ_GET_IS_TIME_SYNCHRO           Tag = 0x0A000044
	INFO_LC_REQ_GET_SYSTEM_TIME               Tag = 0x0A000045
	INFO_LC_REQ_GET_TIME_TIMEZONE             Tag = 0x0A000046
	INFO_REQ_GET_IS_MULTI_GC_POSSIBLE         Tag = 0x0A000047
	INFO_REQ_GET_VALUES_PRESENTATION_SETTINGS Tag = 0x0A000048
	INFO_LC_PARAM_TOKEN                       Tag = 0x0A400001
	INFO_LC_PARAM_STATUS                      Tag = 0x0A400002
	INFO_LC_PARAM_TIME_SYSTEM                 Tag = 0x0A400003
	INFO_LC_PARAM_TIMEZONE                    Tag = 0x0A400004
	INFO_LC_PARAM_TIME_DIFF                   Tag = 0x0A400005
	INFO_LC_PARAM_TIME_BROWSER                Tag = 0x0A400006
	INFO_LC_PARAM_TIME_SYNCHRO_STATE          Tag = 0x0A400007

	INFO_SET_GUI_TARGET                   Tag = 0x0A80001A
	INFO_GUI_TARGET                       Tag = 0x0A80001B
	INFO_PLATFORM_TYPE                    Tag = 0x0A80001C
	INFO_IS_CALIBRATED                    Tag = 0x0A80001D
	INFO_CALIBRATION_CHECK                Tag = 0x0A80001E
	INFO_RESET_CALIBRATION                Tag = 0x0A80001F
	INFO_HW_TIME                          Tag = 0x0A800020
	INFO_SET_TIME_UTC                     Tag = 0x0A800021
	INFO_SET_HW_TIME                      Tag = 0x0A800022
	INFO_SET_FACILITY                     Tag = 0x0A800023
	INFO_GET_FACILITY                     Tag = 0x0A800024
	INFO_NAME                             Tag = 0x0A800025
	INFO_STREET                           Tag = 0x0A800026
	INFO_STREET_NO                        Tag = 0x0A800027
	INFO_POSTCODE                         Tag = 0x0A800028
	INFO_CITY                             Tag = 0x0A800029
	INFO_FON                              Tag = 0x0A80002A
	INFO_E_MAIL                           Tag = 0x0A80002B
	INFO_COUNTRY                          Tag = 0x0A80002C
	INFO_GET_FS_USAGE                     Tag = 0x0A80002D
	INFO_FS_SIZE                          Tag = 0x0A80002E
	INFO_FS_USED                          Tag = 0x0A80002F
	INFO_FS_AVAILABLE                     Tag = 0x0A800030
	INFO_FS_USE_PERCENT                   Tag = 0x0A800031
	INFO_INODES                           Tag = 0x0A800032
	INFO_INODES_USED                      Tag = 0x0A800033
	INFO_INODES_AVAILABLE                 Tag = 0x0A800034
	INFO_INODES_USE_PERCENT               Tag = 0x0A800035
	INFO_SURNAME                          Tag = 0x0A800036
	INFO_UPNP_STATUS                      Tag = 0x0A800037
	INFO_SET_UPNP_STATUS                  Tag = 0x0A800038
	INFO_IS_OVP_POSSIBLE                  Tag = 0x0A800039
	INFO_IS_RSCP_PASSWORD_SET             Tag = 0x0A80003A
	INFO_SET_EULA_VERSION                 Tag = 0x0A80003B
	INFO_SET_EULA_CHECKSUM                Tag = 0x0A80003C
	INFO_SET_WAIT_FOR_EULA                Tag = 0x0A80003D
	INFO_ASSEMBLY_SERIAL_NUMBER           Tag = 0x0A80003E
	INFO_SET_UUID                         Tag = 0x0A80003F
	INFO_GET_UUID                         Tag = 0x0A800040
	INFO_GET_SUID                         Tag = 0x0A800041
	INFO_IS_OVP2_POSSIBLE                 Tag = 0x0A800042
	INFO_LC_SET_TIME                      Tag = 0x0A800043
	INFO_LC_GET_IS_TIME_SYNCHRO           Tag = 0x0A800044
	INFO_LC_GET_SYSTEM_TIME               Tag = 0x0A800045
	INFO_LC_GET_TIME_TIMEZONE             Tag = 0x0A800046
	INFO_GET_IS_MULTI_GC_POSSIBLE         Tag = 0x0A800047
	INFO_GET_VALUES_PRESENTATION_SETTINGS Tag = 0x0A800048
)

// --------------
// NAMESPACE: EP
// 0x0Bxxxxxx
// undocumented
// --------------
const (
	EP_REQ_IS_READY_FOR_SWITCH Tag = 0x0B000003
	EP_REQ_IS_GRID_CONNECTED   Tag = 0x0B000004
	EP_REQ_IS_ISLAND_GRID      Tag = 0x0B000005
	EP_REQ_IS_INVALID_STATE    Tag = 0x0B000006
	EP_REQ_IS_POSSIBLE         Tag = 0x0B000007
	EP_IS_READY_FOR_SWITCH     Tag = 0x0B800003
	EP_IS_GRID_CONNECTED       Tag = 0x0B800004
	EP_IS_ISLAND_GRID          Tag = 0x0B800005
	EP_IS_INVALID_STATE        Tag = 0x0B800006
	EP_IS_POSSIBLE             Tag = 0x0B800007
	EP_GENERAL_ERROR           Tag = 0x0BFFFFFF

	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	EP_REQ_SWITCH_TO_GRID   Tag = 0x0B000001
	EP_REQ_SWITCH_TO_ISLAND Tag = 0x0B000002

	EP_REQ_LEAVE_INVALID_STATE_TO_ISLAND Tag = 0x0B000008
	EP_REQ_LEAVE_INVALID_STATE_TO_GRID   Tag = 0x0B000009
	EP_SWITCH_TO_GRID                    Tag = 0x0B800001
	EP_SWITCH_TO_ISLAND                  Tag = 0x0B800002

	EP_LEAVE_INVALID_STATE_TO_ISLAND Tag = 0x0B800008
	EP_LEAVE_INVALID_STATE_TO_GRID   Tag = 0x0B800009
)

// --------------
// NAMESPACE: SYS
// 0x0Cxxxxxx
// undocumented
// --------------
const (
	SYS_REQ_SYSTEM_REBOOT Tag = 0x0C000001
	// Erläuterung
	//  0    - Reboot kann momentan nicht durchgeführt -> später nochmal versuchen (im Moment nicht in gebrauch)
	//  1    - Reboot wird durchgeführt
	//  2    - Warten auf andere Services danach wird Reboot autmatisch durchgeführt
	SYS_SYSTEM_REBOOT           Tag = 0x0C800001
	SYS_REQ_IS_SYSTEM_REBOOTING Tag = 0x0C000002
	SYS_IS_SYSTEM_REBOOTING     Tag = 0x0C800002
	SYS_REQ_RESTART_APPLICATION Tag = 0x0C000003
	// Erläuterung
	//  false  - Applikationsneustart kann nicht durchgeführt werden (z.B. Software Update läuft) -> später nochmal versuchen
	//  true   - Applikationsneustart wird durchgeführt
	SYS_RESTART_APPLICATION Tag = 0x0C800003
	SYS_SCRIPT_FILE         Tag = 0x0C800011
	SYS_GENERAL_ERROR       Tag = 0x0CFFFFFF

	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	SYS_REQ_SCRIPT_FILE_LIST           Tag = 0x0C000010
	SYS_REQ_EXECUTE_SCRIPT             Tag = 0x0C000015
	SYS_REQ_SYSTEM_SHUTDOWN            Tag = 0x0C000016
	SYS_REQ_IS_SYSTEM_SHUTTING_DOWN    Tag = 0x0C000017
	SYS_REQ_GUI_PASSWD_RESET           Tag = 0x0C000018
	SYS_REQ_IS_GUI_PASSWD_RESET        Tag = 0x0C000019
	SYS_REQ_GUI_PASSWD_RESET_PERFORMED Tag = 0x0C000020

	SYS_SCRIPT_FILE_LIST Tag = 0x0C800010

	SYS_EXECUTE_SCRIPT             Tag = 0x0C800015
	SYS_IS_SYSTEM_SHUTING_DOWN     Tag = 0x0C800017
	SYS_GUI_PASSWD_RESET           Tag = 0x0C800018
	SYS_IS_GUI_PASSWD_RESET        Tag = 0x0C800019
	SYS_GUI_PASSWD_RESET_PERFORMED Tag = 0x0C800020
)

// -------------
// NAMESPACE: UM
// 0x0Dxxxxxx
// undocumented
// -------------
const (
	UM_REQ_UPDATE_STATUS Tag = 0x0D000001
	// Status:
	//  IDLE = 0x00
	//  UPDATE_CHECK_RUNNING = 0x01
	//  UPDATING_MODULES_AND_FILES  = 0x02
	//  UPDATING_HARDWARE = 0x03
	UM_UPDATE_STATUS         Tag = 0x0D800001
	UM_REQ_CHECK_FOR_UPDATES Tag = 0x0D000003
	// Status:
	//  0 = check nicht möglich (kein Internet?)
	//  1 = check wird ausgeführt, wenn was neues entdeckt wird, wird es installiert
	UM_CHECK_FOR_UPDATES Tag = 0x0D800003
	UM_GENERAL_ERROR     Tag = 0x0DFFFFFF

	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	UM_REQ_UPDATE_DCDC Tag = 0x0D000002

	UM_LC_REQ_SET_START_FILE_TRANSFER Tag = 0x0D000004
	UM_LC_REQ_SET_FILE_TRANSFER       Tag = 0x0D000005
	UM_LC_REQ_SET_END_FILE_TRANSFER   Tag = 0x0D000006
	UM_LC_PARAM_TOKEN                 Tag = 0x0D400001
	UM_LC_PARAM_STATUS                Tag = 0x0D400002
	UM_LC_PARAM_FILE_NAME             Tag = 0x0D400003
	UM_LC_PARAM_FILE_SIZE             Tag = 0x0D400004
	UM_LC_PARAM_FILE_CHECKSUM         Tag = 0x0D400005
	UM_LC_PARAM_TRANSACTION_ID        Tag = 0x0D400006
	UM_LC_PARAM_TRANSFER_TYPE         Tag = 0x0D400007
	UM_LC_PARAM_TRANSACTION_STATUS    Tag = 0x0D400008
	UM_LC_PARAM_BLOCK_ID              Tag = 0x0D400009
	UM_LC_PARAM_UPDATE_STATUS         Tag = 0x0D40000A
	UM_LC_PARAM_CHUNK                 Tag = 0x0D40000B
	UM_LC_PARAM_CHUNK_ID              Tag = 0x0D40000C
	UM_LC_PARAM_CHUNK_DATA            Tag = 0x0D40000D
	UM_LC_PARAM_CHUNK_DATA_LENGTH     Tag = 0x0D40000E

	UM_UPDATE_DCDC Tag = 0x0D800002

	UM_LC_SET_START_FILE_TRANSFER Tag = 0x0D800004
	UM_LC_SET_FILE_TRANSFER       Tag = 0x0D800005
	UM_LC_SET_END_FILE_TRANSFER   Tag = 0x0D800006
)

// -------------
// NAMESPACE: WB
// 0x0Exxxxxx
// undocumented
// -------------
const (
	// Beinhaltet alle Anfrage-TAGs, der Container MUSS einen Index enthalten
	WB_REQ_DATA Tag = 0x0E040000
	// Index des angefragten Gerätes (0?x) 0xFF -> GroupController
	WB_INDEX Tag = 0x0E040001
	// Antwort mit allen Daten der REQ_DATA Anfrage
	WB_DATA Tag = 0x0E840000
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_ENERGY_ALL Tag = 0x0E000001
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_ENERGY_SOLAR Tag = 0x0E000002
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden! Nicht aussagekräftig solange die E-Cars das noch nicht unterstützen
	WB_REQ_SOC Tag = 0x0E000003
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_STATUS Tag = 0x0E000004
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_ERROR_CODE Tag = 0x0E000005
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_MODE Tag = 0x0E000006
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_APP_SOFTWARE Tag = 0x0E000007
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_BOOTLOADER_SOFTWARE Tag = 0x0E000008
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_HW_VERSION Tag = 0x0E000009
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_FLASH_VERSION Tag = 0x0E00000A
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_DEVICE_ID    Tag = 0x0E00000B
	WB_REQ_DEVICE_STATE Tag = 0x0E060000
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_PM_POWER_L1 Tag = 0x0E00000C
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_PM_POWER_L2 Tag = 0x0E00000D
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_PM_POWER_L3 Tag = 0x0E00000E
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_PM_ACTIVE_PHASES Tag = 0x0E00000F
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_PM_MODE Tag = 0x0E000011
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_PM_ENERGY_L1 Tag = 0x0E000012
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_PM_ENERGY_L2 Tag = 0x0E000013
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_PM_ENERGY_L3 Tag = 0x0E000014
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_PM_DEVICE_ID Tag = 0x0E000015
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_PM_ERROR_CODE   Tag = 0x0E000016
	WB_REQ_PM_DEVICE_STATE Tag = 0x0E000029
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_PM_FIRMWARE_VERSION Tag = 0x0E000017
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_DIAG_INFOS Tag = 0x0E00001F
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_DIAG_WARNINGS Tag = 0x0E000020
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_DIAG_ERRORS Tag = 0x0E000021
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_DIAG_TEMP_1 Tag = 0x0E000022
	// Kann nur innerhalb eines REQ_WB_DATA Container verwendet werden!
	WB_REQ_DIAG_TEMP_2 Tag = 0x0E000023
	// Current power on L1
	WB_ENERGY_ALL Tag = 0x0E800001
	// Current power on L2
	WB_ENERGY_SOLAR Tag = 0x0E800002
	// Current power on L3
	WB_SOC Tag = 0x0E800003
	// just the three lowest bits of activePhases are used to define
	// what phase is switched on. If the lowest bit is 1 phase1 is active
	// if the lowest bit is 0 phase 1 is inactive ...
	//     static const unsigned char PHASE_1 = 1
	//     static const unsigned char PHASE_2 = 2
	//     static const unsigned char PHASE_3 = 4
	// f.e. if active Phases = 7 -> all phases are active
	WB_STATUS Tag = 0x0E800004
	// used to identify the error bit, if error code is available mode = ERROR_ACTIVE_MODE. ACTIVE_MODE else. Ignore all other modes.
	//     static const unsigned char ACTIVE_MODE = 0
	//     static const unsigned char PASSIVE_MODE = 1
	//     static const unsigned char DIAGNOSE_MODE = 2
	//     static const unsigned char ERROR_ACTIVE_MODE = 3
	//     static const unsigned char ERROR_PASSIVE_MODE = 4
	WB_ERROR_CODE Tag = 0x0E800005
	// Energy counter L1
	WB_MODE Tag = 0x0E800006
	// Energy counter L2
	WB_APP_SOFTWARE Tag = 0x0E800007
	// Energy counter L3
	WB_BOOTLOADER_SOFTWARE Tag = 0x0E800008
	// ID of that device
	WB_HW_VERSION Tag = 0x0E800009
	// Last reported error code (see mode if error has relevance)
	WB_FLASH_VERSION Tag = 0x0E80000A
	WB_DEVICE_ID     Tag = 0x0E80000B
	// DEVICE_CONNECTED & DEVICE_WORKING & DEVICE_IN_SERVICE
	WB_DEVICE_STATE               Tag = 0x0E860000
	WB_DEVICE_CONNECTED           Tag = 0x0E860001
	WB_DEVICE_WORKING             Tag = 0x0E860002
	WB_DEVICE_IN_SERVICE          Tag = 0x0E860003
	WB_GENERAL_ERROR              Tag = 0x0EFFFFFF
	WB_PM_POWER_L1                Tag = 0x0E80000C
	WB_PM_POWER_L2                Tag = 0x0E80000D
	WB_PM_POWER_L3                Tag = 0x0E80000E
	WB_PM_ACTIVE_PHASES           Tag = 0x0E80000F
	WB_PM_MODE                    Tag = 0x0E800011
	WB_PM_ENERGY_L1               Tag = 0x0E800012
	WB_PM_ENERGY_L2               Tag = 0x0E800013
	WB_PM_ENERGY_L3               Tag = 0x0E800014
	WB_PM_DEVICE_ID               Tag = 0x0E800015
	WB_PM_ERROR_CODE              Tag = 0x0E800016
	WB_PM_DEVICE_STATE            Tag = 0x0E800029
	WB_PM_DEVICE_STATE_CONNECTED  Tag = 0x0E800030
	WB_PM_DEVICE_STATE_WORKING    Tag = 0x0E800031
	WB_PM_DEVICE_STATE_IN_SERVICE Tag = 0x0E800032
	WB_PM_FIRMWARE_VERSION        Tag = 0x0E800017
	WB_DIAG_INFOS                 Tag = 0x0E80001F
	WB_DIAG_WARNINGS              Tag = 0x0E800020
	WB_DIAG_ERRORS                Tag = 0x0E800021
	WB_DIAG_TEMP_1                Tag = 0x0E800022
	WB_DIAG_TEMP_2                Tag = 0x0E800023
	// Beinhaltet WB_INDEX, der Value entscheidet welche Wallbox abgefragt wird
	WB_REQ_AVAILABLE_SOLAR_POWER Tag = 0x0E041000
	WB_POWER                     Tag = 0x0E041001
	WB_STATUS_BIT                Tag = 0x0E041002
	WB_AVAILABLE_SOLAR_POWER     Tag = 0x0E841000
	WB_REQ_SET_MODE              Tag = 0x0E000030
	WB_MODE_PARAM_MODE           Tag = 0x0E040031
	WB_MODE_PARAM_MAX_CURRENT    Tag = 0x0E040032
	// err value, 0 for successfully set mode
	WB_SET_MODE Tag = 0x0E000031
	// Expects EXTERN_DATA (length 6) and EXTERN_DATA_LEN =6
	//  Byte 1:   1-Sonnenmode / 2-Mischmode
	//  Byte 2:   Strombegrenzung für alle Modes, [1 ? 32] A
	//  Byte 3:  PreCharge (1: +5%	// 2: -5%)
	//  Byte 4: > 0: Anzahl Phasen tauschen
	//  Byte 5: > 0: Typ2, Laden abbrechen
	//  Byte 6: > 0: Schuko, Bestätigung für ?AN?
	WB_REQ_SET_EXTERN Tag = 0x0E041010
	// no content
	WB_SET_EXTERN          Tag = 0x0E841010
	WB_EXTERN_DATA         Tag = 0x0E042010
	WB_EXTERN_DATA_LEN     Tag = 0x0E042011
	WB_REQ_EXTERN_DATA_SUN Tag = 0x0E041011
	WB_REQ_EXTERN_DATA_NET Tag = 0x0E041012
	WB_REQ_EXTERN_DATA_ALL Tag = 0x0E041013
	WB_REQ_EXTERN_DATA_ALG Tag = 0x0E041014
	// contains EXTERN_DATA (length 7) and EXTERN_DATA_LEN =7
	//  Byte 1-2: uint16, Sonnenleistung in [W]
	//  Byte 3-6: uint32, Sonnenenergie in [Wh]
	//  Byte 7: uint8, Sonnenmenge in [%] /
	WB_EXTERN_DATA_SUN Tag = 0x0E841011
	// contains EXTERN_DATA (length 7) and EXTERN_DATA_LEN =7
	//  Byte 1-2: uint16, Netzleistung in [W]
	//  Byte 3-6: uint32, Netzenergie in [Wh]
	//  Byte 7: uint8, Netzmenge in [%]
	WB_EXTERN_DATA_NET Tag = 0x0E841012
	// contains EXTERN_DATA (length 7) and EXTERN_DATA_LEN =7
	//  Byte 1-2: uint16, Gesamtleistung in [W]
	//  Byte 3-6: uint32, Gesamtenergie in [Wh]
	//  Byte 7: uint8, Gesamtmenge in [%]
	WB_EXTERN_DATA_ALL Tag = 0x0E841013
	// contains EXTERN_DATA (length 7) and EXTERN_DATA_LEN =7
	//  Byte 1: uint8, PreCharge in [%]
	//  Byte 2: uint8, 1: Sonnenmode, 0: Misch.
	//  Byte 3: uint8, 1: Auto lädt, 0: lädt nicht
	//  Byte 4: uint8, 1: Typ2 verriegelt, 0: entr.
	//  Byte 5: uint8, Anzahl akt. Phasen [0-3]
	//  Byte 6: uint4 low, 1: Schuko belegt
	//          uint4 high, 1: Schuko an
	WB_EXTERN_DATA_ALG Tag = 0x0E841014
	// Set capacity in Wh
	WB_REQ_SET_BAT_CAPACITY Tag = 0x0E041015
	// Expects EXTERN_DATA (length 6) and EXTERN_DATA_LEN =6
	//  Byte 1: User Parameter, uint16 Byte 0
	//  Byte 2: User Parameter, uint16 Byte 1
	//  Byte 3: Maximaler Ladestrom, uint8
	//  Byte 4: Phasenspannung, uint8
	//  Byte 5: Display Sprache, uint8
	//  Byte 6: Display Design, uint8
	WB_REQ_SET_PARAM_1 Tag = 0x0E041018
	// Expects EXTERN_DATA (length 6) and EXTERN_DATA_LEN =6
	//  Byte 1: ISstart, uint8 in [A]
	//  Byte 2: ISmin, uint8 in [A]
	//  Byte 3: ISmax, uint8 in [A]
	//  Byte 4 ? 6: Kein Inhalt
	WB_REQ_SET_PARAM_2  Tag = 0x0E041019
	WB_SET_BAT_CAPACITY Tag = 0x0E841015
	// contains EXTERN_DATA (length 6) and EXTERN_DATA_LEN =6
	//  Byte 1: User Parameter, uint16 Byte 0
	//  Byte 2: User Parameter, uint16 Byte 1
	//  Byte 3: Maximaler Ladestrom, uint8
	//  Byte 4: Phasenspannung, uint8
	//  Byte 5: Display Sprache, uint8
	//  Byte 6: Display Design, uint8
	WB_SET_PARAM_1 Tag = 0x0E841018
	// contains EXTERN_DATA (length 6) and EXTERN_DATA_LEN =6
	//  Byte 1: ISstart, uint8 in [A]
	//  Byte 2: ISmin, uint8 in [A]
	//  Byte 3: ISmax, uint8 in [A]
	//  Byte 4 ? 6: Kein Inhalt /
	WB_SET_PARAM_2 Tag = 0x0E841019
	WB_REQ_PARAM_2 Tag = 0x0E04101A
	// contains EXTERN_DATA (length 6) and EXTERN_DATA_LEN =6
	//  Byte 1: ISstart, uint8 in [A]
	//  Byte 2: ISmin, uint8 in [A]
	//  Byte 3: ISmax, uint8 in [A]
	//  Byte 4 ? 6: Kein Inhalt
	WB_RSP_PARAM_2 Tag = 0x0E84101A
	WB_REQ_PARAM_1 Tag = 0x0E04101B
	// contains EXTERN_DATA (length 6) and EXTERN_DATA_LEN =6
	//  Byte 1: User Parameter, uint16 Byte 0
	//  Byte 2: User Parameter, uint16 Byte 1
	//  Byte 3: Maximaler Ladestrom, uint8
	//  Byte 4: Phasenspannung, uint8
	//  Byte 5: Display Sprache, uint8
	//  Byte 6: Display Design, uint8
	WB_RSP_PARAM_1 Tag = 0x0E84101B

	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	WB_REQ_SET_UPPER_CURRENT_LIMIT Tag = 0x0E000010

	WB_REQ_DIAG_DEVICE_ID      Tag = 0x0E000018
	WB_REQ_DIAG_BAT_CAPACITY   Tag = 0x0E000019
	WB_REQ_DIAG_USER_PARAM     Tag = 0x0E00001A
	WB_REQ_DIAG_MAX_CURRENT    Tag = 0x0E00001B
	WB_REQ_DIAG_PHASE_VOLTAGE  Tag = 0x0E00001C
	WB_REQ_DIAG_DISPLAY_SPEECH Tag = 0x0E00001D
	WB_REQ_DIAG_DESIGN         Tag = 0x0E00001E

	WB_REQ_DIAG_CP_PEGEL     Tag = 0x0E000024
	WB_REQ_DIAG_PP_IN_A      Tag = 0x0E000025
	WB_REQ_DIAG_STATUS_DIODE Tag = 0x0E000026
	WB_REQ_DIAG_DIG_IN_1     Tag = 0x0E000027
	WB_REQ_DIAG_DIG_IN_2     Tag = 0x0E000028

	WB_REQ_AUTH_REQUIRED           Tag = 0x0E00002A
	WB_REQ_SET_AUTH_REQUIRED       Tag = 0x0E00002B
	WB_REQ_SESSION                 Tag = 0x0E00002C
	WB_REQ_SET_RFID_READER_ENABLED Tag = 0x0E00002D
	WB_REQ_RFID_READER_ENABLED     Tag = 0x0E00002E
	WB_REQ_FIRMWARE_VERSION        Tag = 0x0E00002F

	WB_REQ_SET_BIDI_MODE_ACTIVE                Tag = 0x0E000033
	WB_REQ_BIDI_MODE_ACTIVE                    Tag = 0x0E000034
	WB_REQ_BIDI_CONTINGENT                     Tag = 0x0E000035
	WB_REQ_SET_BIDI_CONTINGENT                 Tag = 0x0E000036
	WB_REQ_ALIVE                               Tag = 0x0E000037
	WB_REQ_SET_AUTO_PHASE_SWITCH_ENABLED       Tag = 0x0E000038
	WB_REQ_AUTO_PHASE_SWITCH_ENABLED           Tag = 0x0E000039
	WB_REQ_BIDI_MIN_ENERGY_REQUEST             Tag = 0x0E00003A
	WB_REQ_BIDI_MAX_ENERGY_REQUEST             Tag = 0x0E00003B
	WB_REQ_BIDI_TARGET_ENERGY_REQUEST          Tag = 0x0E00003C
	WB_REQ_BIDI_ENERGY_REQUESTS                Tag = 0x0E00003D
	WB_REQ_CHARGE_IN_FALLBACK_MODE_SELECTABLE  Tag = 0x0E00003E
	WB_REQ_CHARGE_IN_FALLBACK_MODE_ALLOWED     Tag = 0x0E00003F
	WB_REQ_PM_MAX_PHASE_POWER                  Tag = 0x0E000040
	WB_REQ_SET_DEVICE_NAME                     Tag = 0x0E000041
	WB_REQ_DEVICE_NAME                         Tag = 0x0E000042
	WB_REFRESH_LOGGEDMESSAGES                  Tag = 0x0E000043
	WB_REQ_GET_LOGGEDMESSAGES                  Tag = 0x0E000044
	WB_REQ_SET_CHARGE_IN_FALLBACK_MODE_ALLOWED Tag = 0x0E000045
	WB_REQ_MAX_CURRENT_FALLBACK                Tag = 0x0E000046
	WB_REQ_SET_MAX_CURRENT_FALLBACK            Tag = 0x0E000047
	WB_REQ_SET_LED_COLOR                       Tag = 0x0E000048
	WB_REQ_DC_CHARGER_STATE                    Tag = 0x0E000049

	WB_REQ_SET_ENERGY_ALL   Tag = 0x0E041016
	WB_REQ_SET_ENERGY_SOLAR Tag = 0x0E041017

	WB_REQ_CONNECTED_DEVICES          Tag = 0x0E04101C
	WB_REQ_SET_SOC                    Tag = 0x0E04101D
	WB_REQ_STATION_AVAILABLE          Tag = 0x0E04101E
	WB_REQ_SET_STATION_AVAILABLE      Tag = 0x0E04101F
	WB_REQ_SET_PW                     Tag = 0x0E041020
	WB_REQ_SET_STATION_ENABLED        Tag = 0x0E041021
	WB_REQ_MAC_ADDRESS                Tag = 0x0E041022
	WB_REQ_PROXIMITY_PLUG             Tag = 0x0E041023
	WB_REQ_PREFERRED_CHARGE_POWER     Tag = 0x0E041024
	WB_REQ_CHARGE_FULL                Tag = 0x0E041025
	WB_REQ_SET_CHARGE_FULL            Tag = 0x0E041026
	WB_REQ_ACTIVE_CHARGE_STRATEGY     Tag = 0x0E041027
	WB_REQ_SET_ACTIVE_CHARGE_STRATEGY Tag = 0x0E041028
	WB_REQ_PARAMETER_LIST             Tag = 0x0E041029
	WB_REQ_STATION_ENABLED            Tag = 0x0E04102A
	WB_REQ_SET_PARAMETER_LIST         Tag = 0x0E041030
	WB_REQ_GATEWAY                    Tag = 0x0E041031
	WB_REQ_SUBNET_MASK                Tag = 0x0E041032
	WB_REQ_IP_ADDR                    Tag = 0x0E041033
	WB_REQ_DHCP_ENABLED               Tag = 0x0E041034
	WB_REQ_SET_DHCP_ENABLED           Tag = 0x0E041035
	WB_REQ_WALLBOX_TYPE               Tag = 0x0E041036
	WB_REQ_UPDATE_NETWORK_CONFIG      Tag = 0x0E041037
	WB_REQ_SUN_MODE_ACTIVE            Tag = 0x0E041038
	WB_REQ_SET_SUN_MODE_ACTIVE        Tag = 0x0E041039
	WB_REQ_NUMBER                     Tag = 0x0E04103A
	WB_REQ_NUMBER_PHASES              Tag = 0x0E04103B
	WB_REQ_SET_NUMBER_PHASES          Tag = 0x0E04103C
	WB_REQ_ABORT_CHARGING             Tag = 0x0E04103D
	WB_REQ_SET_ABORT_CHARGING         Tag = 0x0E04103E
	WB_REQ_SCHUKO_AVAILABLE           Tag = 0x0E041040
	WB_REQ_IS_SCHUKO_LOCKED           Tag = 0x0E041041
	WB_REQ_SET_SCHUKO_LOCKED          Tag = 0x0E041042
	WB_REQ_MAX_POWER_PER_PHASE        Tag = 0x0E041043
	WB_REQ_MIN_POWER_PER_PHASE        Tag = 0x0E041044
	WB_REQ_UPPER_CURRENT_LIMIT        Tag = 0x0E041045
	WB_REQ_LOWER_CURRENT_LIMIT        Tag = 0x0E041046
	WB_REQ_MAX_CHARGE_CURRENT         Tag = 0x0E041047
	WB_REQ_MIN_CHARGE_CURRENT         Tag = 0x0E041048
	WB_REQ_SET_MAX_CHARGE_CURRENT     Tag = 0x0E041049
	WB_REQ_SET_MIN_CHARGE_CURRENT     Tag = 0x0E04104A
	WB_PARAM_INDEX                    Tag = 0x0E04104B
	WB_REQ_CHARGE_STOP_HYSTERESIS     Tag = 0x0E04104C
	WB_REQ_SET_CHARGE_STOP_HYSTERESIS Tag = 0x0E04104D
	WB_REQ_GET_KEY_LOCK_MODE          Tag = 0x0E04104E
	WB_REQ_SET_KEY_LOCK_MODE          Tag = 0x0E04104F
	WB_REQ_KEY_STATE                  Tag = 0x0E041050
	WB_SERIAL                         Tag = 0x0E041051
	WB_REQ_MAX_CHARGE_POWER           Tag = 0x0E041052
	WB_REQ_MIN_CHARGE_POWER           Tag = 0x0E041053

	WB_PARAM_USR Tag = 0x0E042012
	WB_PARAM_PW  Tag = 0x0E042013

	WB_REQ_SET_BIC_MODE                     Tag = 0x0E0F0001
	WB_REQ_GET_BIC_MODE                     Tag = 0x0E0F0002
	WB_REQ_GET_CHARGE_PLAN_TEXT             Tag = 0x0E0F0003
	WB_STRING_PARAMETER                     Tag = 0x0E440010
	WB_PREFERRED_CHARGE_POWER               Tag = 0x0E741024
	WB_SESSION_END_TIME                     Tag = 0x0E741025
	WB_SESSION_START_TIME                   Tag = 0x0E741026
	WB_SESSION_STATUS                       Tag = 0x0E741027
	WB_SESSION_ACTIVE_CHARGE_TIME           Tag = 0x0E741028
	WB_SESSION_INACTIVE_TIME                Tag = 0x0E741029
	WB_SESSION_CHARGED_ENERGY               Tag = 0x0E74102A
	WB_SESSION_CHARGED_SUN_ENERGY           Tag = 0x0E74102B
	WB_SESSION_METER_ENERGY_START           Tag = 0x0E74102C
	WB_SESSION_METER_ENERGY_END             Tag = 0x0E74102D
	WB_SESSION_ID                           Tag = 0x0E74102E
	WB_SESSION_VEHICLE_ID                   Tag = 0x0E74102F
	WB_SESSION_AUTH_DATA                    Tag = 0x0E741030
	WB_SESSION_RECEIPT_SIGNATURE            Tag = 0x0E741031
	WB_SESSION_RECEIPT_DATA                 Tag = 0x0E741032
	WB_SESSION_AUTH_TYPE                    Tag = 0x0E741033
	WB_SESSION_WALLBOX_ID                   Tag = 0x0E741035
	WB_BIDI_CONTINGENT_USED_KWH_YEAR        Tag = 0x0E741036
	WB_BIDI_CONTINGENT_USED_HOURS_YEAR      Tag = 0x0E741037
	WB_BIDI_CONTINGENT_MAX_KWH_LIFETIME     Tag = 0x0E741038
	WB_BIDI_CONTINGENT_MAX_HOURS_LIFETIME   Tag = 0x0E741039
	WB_BIDI_CONTINGENT_MAX_HOURS_YEAR       Tag = 0x0E74103A
	WB_BIDI_CONTINGENT_MAX_KWH_YEAR         Tag = 0x0E74103B
	WB_BIDI_CONTINGENT_CAR_ID               Tag = 0x0E74103C
	WB_BIDI_CONTINGENT_CAR_LIFETIME         Tag = 0x0E74103D
	WB_SESSION_AUTH_DATA_SWAPPED            Tag = 0x0E74103E
	WB_SESSION_ACTIVE_DISCHARGE_TIME        Tag = 0x0E74103F
	WB_SESSION_DISCHARGED_ENERGY            Tag = 0x0E741040
	WB_SESSION_DISCHARGE_METER_ENERGY_START Tag = 0x0E741041
	WB_SESSION_DISCHARGE_METER_ENERGY_STOP  Tag = 0x0E741042

	WB_SET_UPPER_CURRENT_LIMIT Tag = 0x0E800010

	WB_DIAG_DEVICE_ID      Tag = 0x0E800018
	WB_DIAG_BAT_CAPACITY   Tag = 0x0E800019
	WB_DIAG_USER_PARAM     Tag = 0x0E80001A
	WB_DIAG_MAX_CURRENT    Tag = 0x0E80001B
	WB_DIAG_PHASE_VOLTAGE  Tag = 0x0E80001C
	WB_DIAG_DISPLAY_SPEECH Tag = 0x0E80001D
	WB_DIAG_DESIGN         Tag = 0x0E80001E

	WB_DIAG_CP_PEGEL     Tag = 0x0E800024
	WB_DIAG_PP_IN_A      Tag = 0x0E800025
	WB_DIAG_STATUS_DIODE Tag = 0x0E800026
	WB_DIAG_DIG_IN_1     Tag = 0x0E800027
	WB_DIAG_DIG_IN_2     Tag = 0x0E800028

	WB_AUTH_REQUIRED           Tag = 0x0E80002A
	WB_SET_AUTH_REQUIRED       Tag = 0x0E80002B
	WB_SESSION                 Tag = 0x0E80002C
	WB_SET_RFID_READER_ENABLED Tag = 0x0E80002D
	WB_RFID_READER_ENABLED     Tag = 0x0E80002E
	WB_FIRMWARE_VERSION        Tag = 0x0E80002F

	WB_SET_BIDI_MODE_ACTIVE                Tag = 0x0E800033
	WB_BIDI_MODE_ACTIVE                    Tag = 0x0E800034
	WB_BIDI_CONTINGENT                     Tag = 0x0E800035
	WB_SET_BIDI_CONTINGENT                 Tag = 0x0E800036
	WB_ALIVE                               Tag = 0x0E800037
	WB_SET_AUTO_PHASE_SWITCH_ENABLED       Tag = 0x0E800038
	WB_AUTO_PHASE_SWITCH_ENABLED           Tag = 0x0E800039
	WB_BIDI_MIN_ENERGY_REQUEST             Tag = 0x0E80003A
	WB_BIDI_MAX_ENERGY_REQUEST             Tag = 0x0E80003B
	WB_BIDI_TARGET_ENERGY_REQUEST          Tag = 0x0E80003C
	WB_BIDI_ENERGY_REQUESTS                Tag = 0x0E80003D
	WB_CHARGE_IN_FALLBACK_MODE_SELECTABLE  Tag = 0x0E80003E
	WB_CHARGE_IN_FALLBACK_MODE_ALLOWED     Tag = 0x0E80003F
	WB_PM_MAX_PHASE_POWER                  Tag = 0x0E800040
	WB_DEVICE_NAME                         Tag = 0x0E800042
	WB_SET_CHARGE_IN_FALLBACK_MODE_ALLOWED Tag = 0x0E800045
	WB_MAX_CURRENT_FALLBACK                Tag = 0x0E800046
	WB_SET_MAX_CURRENT_FALLBACK            Tag = 0x0E800047
	WB_SET_LED_COLOR                       Tag = 0x0E800048
	WB_DC_CHARGER_STATE                    Tag = 0x0E800049

	WB_SET_ENERGY_ALL   Tag = 0x0E841016
	WB_SET_ENERGY_SOLAR Tag = 0x0E841017

	WB_CONNECTED_DEVICES          Tag = 0x0E84101C
	WB_SET_SOC                    Tag = 0x0E84101D
	WB_STATION_AVAILABLE          Tag = 0x0E84101E
	WB_SET_STATION_AVAILABLE      Tag = 0x0E84101F
	WB_SET_PW                     Tag = 0x0E841020
	WB_SET_STATION_ENABLED        Tag = 0x0E841021
	WB_MAC_ADDRESS                Tag = 0x0E841022
	WB_PROXIMITY_PLUG             Tag = 0x0E841023
	WB_CHARGE_FULL                Tag = 0x0E841025
	WB_SET_CHARGE_FULL            Tag = 0x0E841026
	WB_ACTIVE_CHARGE_STRATEGY     Tag = 0x0E841027
	WB_SET_ACTIVE_CHARGE_STRATEGY Tag = 0x0E841028
	WB_PARAMETER_LIST             Tag = 0x0E841029
	WB_STATION_ENABLED            Tag = 0x0E84102A
	WB_SET_PARAMETER_LIST         Tag = 0x0E841030
	WB_GATEWAY                    Tag = 0x0E841031
	WB_SUBNET_MASK                Tag = 0x0E841032
	WB_IP_ADDR                    Tag = 0x0E841033
	WB_DHCP_ENABLED               Tag = 0x0E841034
	WB_SET_DHCP_ENABLED           Tag = 0x0E841035
	WB_WALLBOX_TYPE               Tag = 0x0E841036
	WB_UPDATE_NETWORK_CONFIG      Tag = 0x0E841037
	WB_SUN_MODE_ACTIVE            Tag = 0x0E841038
	WB_SET_SUN_MODE_ACTIVE        Tag = 0x0E841039
	WB_NUMBER                     Tag = 0x0E84103A
	WB_NUMBER_PHASES              Tag = 0x0E84103B
	WB_SET_NUMBER_PHASES          Tag = 0x0E84103C
	WB_ABORT_CHARGING             Tag = 0x0E84103D
	WB_SET_ABORT_CHARGING         Tag = 0x0E84103F
	WB_SCHUKO_AVAILABLE           Tag = 0x0E841040
	WB_IS_SCHUKO_LOCKED           Tag = 0x0E841041
	WB_SET_SCHUKO_LOCKED          Tag = 0x0E841042
	WB_MAX_POWER_PER_PHASE        Tag = 0x0E841043
	WB_MIN_POWER_PER_PHASE        Tag = 0x0E841044
	WB_UPPER_CURRENT_LIMIT        Tag = 0x0E841045
	WB_LOWER_CURRENT_LIMIT        Tag = 0x0E841046
	WB_MAX_CHARGE_CURRENT         Tag = 0x0E841047
	WB_MIN_CHARGE_CURRENT         Tag = 0x0E841048
	WB_SET_MAX_CHARGE_CURRENT     Tag = 0x0E841049
	WB_SET_MIN_CHARGE_CURRENT     Tag = 0x0E84104A
	WB_CHARGE_STOP_HYSTERESIS     Tag = 0x0E84104C
	WB_SET_CHARGE_STOP_HYSTERESIS Tag = 0x0E84104D
	WB_GET_KEY_LOCK_MODE          Tag = 0x0E84104E
	WB_SET_KEY_LOCK_MODE          Tag = 0x0E84104F
	WB_KEY_STATE                  Tag = 0x0E841050
	WB_REQ_SERIAL                 Tag = 0x0E841051
	WB_MAX_CHARGE_POWER           Tag = 0x0E841052
	WB_MIN_CHARGE_POWER           Tag = 0x0E841053

	WB_SET_BIC_MODE          Tag = 0x0E8F0001
	WB_GET_BIC_MODE          Tag = 0x0E8F0002
	WB_GET_CHARGE_PLAN_TEXT  Tag = 0x0E8F0003
	WB_SESSION_DATA_SAVED    Tag = 0x0E8F0004
	WB_SESSION_DATA_WAIT_FOR Tag = 0x0E8F0005
	WB_GET_LOGGEDMESSAGES    Tag = 0x0E8F0006
)

// ---------------
// NAMESPACE: PTDB
// 0x0Fxxxxxx
// undocumented
// ---------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	PTDB_REQ_SET_STD_PROPS Tag = 0x0F000001
	PTDB_REQ_SET_PROP      Tag = 0x0F000002
	PTDB_REQ_PROP          Tag = 0x0F000003
	PTDB_PARAM_TABLE       Tag = 0x0F400001
	PTDB_PARAM_KEY         Tag = 0x0F400002
	PTDB_PARAM_VALUE       Tag = 0x0F400003
	PTDB_SET_STD_PROPS     Tag = 0x0F800001
	PTDB_SET_PROP          Tag = 0x0F800002
	PTDB_PROP              Tag = 0x0F800003
)

// --------------
// NAMESPACE: LED
// 0x10xxxxxx
// undocumented
// --------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	LED_REQ_SET_BAR_SWITCHED_ON_STATE Tag = 0x10000001
	LED_REQ_BAR_SWITCHED_ON_STATE     Tag = 0x10000002
	LED_REQ_INTENSITY                 Tag = 0x10000003
	LED_REQ_SET_INTENSITY             Tag = 0x10000004
	LED_REQ_COLOR                     Tag = 0x10000005
	LED_REQ_SET_COLOR                 Tag = 0x10000006
	LED_REQ_HW_INFO                   Tag = 0x10000007
	LED_REQ_STORE_CONFIG              Tag = 0x10000008
	LED_REQ_DEVICE_STATE              Tag = 0x10060000
	LED_SET_BAR_SWITCHED_ON_STATE     Tag = 0x10800001
	LED_BAR_SWITCHED_ON_STATE         Tag = 0x10800002
	LED_INTENSITY                     Tag = 0x10800003
	LED_SET_INTENSITY                 Tag = 0x10800004
	LED_COLOR                         Tag = 0x10800005
	LED_SET_COLOR                     Tag = 0x10800006
	LED_HW_INFO                       Tag = 0x10800007
	LED_CONFIG_STORED                 Tag = 0x10800009
	LED_DEVICE_STATE                  Tag = 0x10860000
	LED_INDEX                         Tag = 0x10860001
	LED_RED                           Tag = 0x10860002
	LED_GREEN                         Tag = 0x10860003
	LED_BLUE                          Tag = 0x10860004
	LED_FW_VERSION                    Tag = 0x10860005
	LED_BL_VERSION                    Tag = 0x10860006
	LED_DEVICE_CONNECTED              Tag = 0x10860007
	LED_DEVICE_WORKING                Tag = 0x10860008
	LED_DEVICE_IN_SERVICE             Tag = 0x10860009
)

// ---------------
// NAMESPACE: DIAG
// 0x11xxxxxx
// undocumented
// ---------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	DIAG_REQ_CURRENT_ISSUES  Tag = 0x11000000
	DIAG_REQ_REPORTED_ISSUES Tag = 0x11000001
	DIAG_CURRENT_ISSUES      Tag = 0x11800000
	DIAG_REPORTED_ISSUES     Tag = 0x11800001
	DIAG_ISSUE               Tag = 0x11860000
	DIAG_ERR_CODE            Tag = 0x11860001
	DIAG_ENDURE_TIME         Tag = 0x11860002
	DIAG_TIME_ARISED         Tag = 0x11860003 //nolint:misspell
	DIAG_ERR_MSG             Tag = 0x11860004
)

// --------------
// NAMESPACE: SGR
// 0x12xxxxxx
// undocumented
// --------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	SGR_REQ_STATE              Tag = 0x12000001
	SGR_REQ_READY_TO_USE       Tag = 0x12000002
	SGR_REQ_HW_PROVIDER_LIST   Tag = 0x12000003
	SGR_REQ_SET_AKTIV          Tag = 0x12000006
	SGR_REQ_SET_COOLDOWN_START Tag = 0x12000007
	SGR_REQ_COOLDOWN_END       Tag = 0x12000008
	SGR_REQ_SET_USED_POWER     Tag = 0x12000009
	SGR_REQ_USED_POWER         Tag = 0x12000010
	SGR_REQ_SET_STATE          Tag = 0x12000011
	SGR_REQ_SET_GLOBAL_OFF     Tag = 0x12000012
	SGR_REQ_GLOBAL_OFF         Tag = 0x12000013
	SGR_REQ_DATA               Tag = 0x12040000
	SGR_INDEX                  Tag = 0x12040001
	SGR_STATE                  Tag = 0x12800001
	SGR_READY_TO_USE           Tag = 0x12800002
	SGR_HW_PROVIDER_LIST       Tag = 0x12800003
	SGR_HW_PROVIDER            Tag = 0x12800004
	SGR_NAME                   Tag = 0x12800005
	SGR_AKTIV                  Tag = 0x12800006
	SGR_COOLDOWN_END           Tag = 0x12800008
	SGR_USED_POWER             Tag = 0x12800009
	SGR_GLOBAL_OFF             Tag = 0x12800012
	SGR_DATA                   Tag = 0x12840000
)

// --------------
// NAMESPACE: MBS
// 0x13xxxxxx
// undocumented
// --------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	MBS_REQ_MODBUS_ENABLED         Tag = 0x13000001
	MBS_REQ_MODBUS_CONNECTORS      Tag = 0x13000002
	MBS_REQ_ENABLE_CONNECTOR       Tag = 0x13000003
	MBS_REQ_DISABLE_CONNECTOR      Tag = 0x13000004
	MBS_REQ_CHANGE_SETTING         Tag = 0x13000005
	MBS_REQ_SET_MODBUS_ENABLED     Tag = 0x13700001
	MBS_MODBUS_ENABLED             Tag = 0x13800001
	MBS_MODBUS_CONNECTORS          Tag = 0x13800002
	MBS_CHANGE_SETTING             Tag = 0x13800005
	MBS_REQ_CHANGE_SETTING_ERR     Tag = 0x13800006
	MBS_MODBUS_CONNECTOR_CONTAINER Tag = 0x13810002
	MBS_MODBUS_CONNECTOR_NAME      Tag = 0x13810003
	MBS_MODBUS_CONNECTOR_ID        Tag = 0x13810004
	MBS_MODBUS_CONNECTOR_ENABLED   Tag = 0x13810005
	MBS_MODBUS_CONNECTOR_SETUP     Tag = 0x13810006
	MBS_MODBUS_SETUP_NAME          Tag = 0x13810007
	MBS_MODBUS_SETUP_TYPE          Tag = 0x13810008
	MBS_MODBUS_SETUP_VALUE         Tag = 0x13810009
	MBS_MODBUS_SETUP_VALUES        Tag = 0x1381000A
	MBS_MODBUS_SETUP_VALUE_STRING  Tag = 0x1381000B
)

// -------------
// NAMESPACE: EH
// 0x14xxxxxx
// undocumented
// -------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	EH_REQ_UNREPORTED_ERRORS Tag = 0x14000001
	EH_REQ_MARK_REPORTED     Tag = 0x14000002
	EH_REQ_GET_SAVED_ERRORS  Tag = 0x14000003
	EH_PARAM_ROW             Tag = 0x14040000
	EH_PARAM_ROW_ID          Tag = 0x14040001
	EH_PARAM_ROW_TIME        Tag = 0x14040002
	EH_PARAM_ROW_CODE        Tag = 0x14040003
	EH_PARAM_ROW_TYPE        Tag = 0x14040004
	EH_PARAM_ROW_CLEARED     Tag = 0x14040005
	EH_PARAM_ROW_ERR_SRC     Tag = 0x14040006
	EH_PARAM_ROW_MSG         Tag = 0x14040007
	EH_UNREPORTED_ERRORS     Tag = 0x14800001
	EH_MARKED_REPORTED       Tag = 0x14800002
	EH_GET_SAVED_ERRORS      Tag = 0x14800003
)

// ----------------
// NAMESPACE: UPNPC
// 0x15xxxxxx
// undocumented
// ----------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	UPNPC_REQ_DEFAULT_LIST     Tag = 0x15000001
	UPNPC_REQ_SET_DEFAULT_LIST Tag = 0x15000002
	UPNPC_REQ_SERVICE_LIST     Tag = 0x15000003
	UPNPC_REQ_DEFAULT_LIST_REV Tag = 0x15000006
	UPNPC_REQ_SERVICE_LIST_REV Tag = 0x15000007
	UPNPC_PARAM_DEVICE_ENTRY   Tag = 0x15040000
	UPNPC_PARAM_SERIALNO       Tag = 0x15040001
	UPNPC_PARAM_IP_ADR         Tag = 0x15040002
	UPNPC_PARAM_PORT           Tag = 0x15040003
	UPNPC_PARAM_NAME           Tag = 0x15040004
	UPNPC_PARAM_LOCATION       Tag = 0x15040005
	UPNPC_DEFAULT_LIST         Tag = 0x15800001
	UPNPC_SET_DEFAULT_LIST     Tag = 0x15800002
	UPNPC_SERVICE_LIST         Tag = 0x15800003
	UPNPC_DEFAULT_LIST_REV     Tag = 0x15800006
	UPNPC_SERVICE_LIST_REV     Tag = 0x15800007
)

// --------------
// NAMESPACE: KNX
// 0x16xxxxxx
// undocumented
// --------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	KNX_REQ_SET             Tag = 0x16000001
	KNX_MAC                 Tag = 0x16000002
	KNX_IP                  Tag = 0x16000003
	KNX_SOFTWAREVERSION     Tag = 0x16000004
	KNX_FIRMWAREVERSION     Tag = 0x16000005
	KNX_REQ_STORED_ERRORS   Tag = 0x16000006
	KNX_REQ_STORE_FILTER    Tag = 0x16000007
	KNX_REQ_SHOW_KNX_SCREEN Tag = 0x16000008
	KNX_ERROR_CONTAINER     Tag = 0x16400001
	KNX_ERROR_TYPE          Tag = 0x16400002
	KNX_ERROR_SOURCE        Tag = 0x16400003
	KNX_ERROR_MESSAGE       Tag = 0x16400004
	KNX_ERROR_CODE          Tag = 0x16400005
	KNX_ERROR_TIMESTAMP     Tag = 0x16400006
	KNX_STORE_FILTER        Tag = 0x16400007
	KNX_FILTER_CATEGORY     Tag = 0x16400008
	KNX_FILTER_ENUM         Tag = 0x16400009
	KNX_RSP_SET             Tag = 0x16800001
	KNX_STORED_ERRORS       Tag = 0x16800002
)

// ----------------
// NAMESPACE: EMSHB
// 0x17xxxxxx
// undocumented
// ----------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	EMSHB_REQ_HB_DATA         Tag = 0x17000001
	EMSHB_PARAM_VERSION       Tag = 0x17040000
	EMSHB_PARAM_BAT_S1        Tag = 0x17040001
	EMSHB_PARAM_BAT_S2        Tag = 0x17040002
	EMSHB_PARAM_BAT_S3        Tag = 0x17040003
	EMSHB_PARAM_LM1           Tag = 0x17040004
	EMSHB_PARAM_LM2           Tag = 0x17040005
	EMSHB_PARAM_LM3           Tag = 0x17040006
	EMSHB_PARAM_AC_L1         Tag = 0x17040007
	EMSHB_PARAM_AC_L2         Tag = 0x17040008
	EMSHB_PARAM_AC_L3         Tag = 0x17040009
	EMSHB_PARAM_C_L1          Tag = 0x17040010
	EMSHB_PARAM_C_L2          Tag = 0x17040011
	EMSHB_PARAM_C_L3          Tag = 0x17040012
	EMSHB_PARAM_SOC           Tag = 0x17040013
	EMSHB_PARAM_SYS_STATUS    Tag = 0x17040014
	EMSHB_PARAM_WB            Tag = 0x17040015
	EMSHB_PARAM_WB_INDEX      Tag = 0x17040016
	EMSHB_PARAM_WB_L1         Tag = 0x17040017
	EMSHB_PARAM_WB_L2         Tag = 0x17040018
	EMSHB_PARAM_WB_L3         Tag = 0x17040019
	EMSHB_PARAM_WB_L1_ACTIVE  Tag = 0x17040020
	EMSHB_PARAM_WB_L2_ACTIVE  Tag = 0x17040021
	EMSHB_PARAM_WB_L3_ACTIVE  Tag = 0x17040022
	EMSHB_PARAM_PV_S1         Tag = 0x17040023
	EMSHB_PARAM_PV_S2         Tag = 0x17040024
	EMSHB_PARAM_PV_S3         Tag = 0x17040025
	EMSHB_PARAM_LM            Tag = 0x17040026
	EMSHB_PARAM_ID            Tag = 0x17040027
	EMSHB_PARAM_L1            Tag = 0x17040028
	EMSHB_PARAM_L2            Tag = 0x17040029
	EMSHB_PARAM_L3            Tag = 0x17040030
	EMSHB_PARAM_LM_ALIVE_FLAG Tag = 0x17040031
	EMSHB_PARAM_WB_ALIVE_FLAG Tag = 0x17040032
	EMSHB_PARAM_WB_SOLAR_L1   Tag = 0x17040033
	EMSHB_PARAM_WB_SOLAR_L2   Tag = 0x17040034
	EMSHB_PARAM_WB_SOLAR_L3   Tag = 0x17040035
	EMSHB_HB_DATA             Tag = 0x17800001
)

// ---------------
// NAMESPACE: MYPV
// 0x18xxxxxx
// undocumented
// ---------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	MYPV_REQ_FIND_DEVICES           Tag = 0x18000003
	MYPV_REQ_REMOVE_DEVICES         Tag = 0x18000006
	MYPV_REQ_INSTANT_BOOST          Tag = 0x18000007
	MYPV_DEVICE                     Tag = 0x18000100
	MYPV_DEVICE_SERIAL              Tag = 0x18000101
	MYPV_DEVICE_ENABLED             Tag = 0x18000102
	MYPV_DEVICE_IP                  Tag = 0x18000103
	MYPV_DEVICE_TEMPERATURE_CURRENT Tag = 0x18000104
	MYPV_DEVICE_TEMPERATURE_MAXIMUM Tag = 0x18000105
	MYPV_DEVICE_POWER               Tag = 0x18000106
	MYPV_DEVICE_STATUS              Tag = 0x18000107
	MYPV_DEVICE_CONTROL_MODE        Tag = 0x18000108
	MYPV_DEVICE_TYPE                Tag = 0x18000109
	MYPV_DEVICE_TIMESPAN_IBOOST     Tag = 0x18000110
	MYPV_DEVICE_BOOST_LIST          Tag = 0x18000200
	MYPV_DEVICE_BOOST_ITEM          Tag = 0x18000300
	MYPV_DEVICE_BOOST_START         Tag = 0x18000301
	MYPV_DEVICE_BOOST_STOP          Tag = 0x18000302
	MYPV_DEVICE_BOOST_TEMPERATURE   Tag = 0x18000303
	MYPV_DEVICE_BOOST_ACTIVE        Tag = 0x18000304
	MYPV_DEVICE_BOOST_WEEKDAYS      Tag = 0x18000305
	MYPV_DEVICE_BOOST_NAME          Tag = 0x18000306
	MYPV_REQ_LIST_DEVICES           Tag = 0x18200004
	MYPV_REQ_WRITE_DEVICES          Tag = 0x18300004
	MYPV_RSP_FIND_DEVICES           Tag = 0x18800003
	MYPV_RSP_REMOVE_DEVICES         Tag = 0x18800006
	MYPV_RSP_INSTANT_BOOST          Tag = 0x18800007
	MYPV_RSP_LIST_DEVICES           Tag = 0x18A00004
	MYPV_RSP_WRITE_DEVICES          Tag = 0x18B00004
)

// ---------------
// NAMESPACE: GPIO
// 0x19xxxxxx
// undocumented
// ---------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	GPIO_REQ_SET             Tag = 0x19000001
	GPIO_REQ_GET             Tag = 0x19000002
	GPIO_REQ_LIST            Tag = 0x19000003
	GPIO_REQ_LIST_ALIAS      Tag = 0x19000004
	GPIO_REQ_LIST_REAL       Tag = 0x19000005
	GPIO_REQ_CONNECT         Tag = 0x19000006
	GPIO_REQ_CHANGECOUNTER   Tag = 0x19000007
	GPIO_REQ_REMOTE_EXCHANGE Tag = 0x19000008
	GPIO_RSP_SET             Tag = 0x19800001
	GPIO_RSP_GET             Tag = 0x19800002
	GPIO_RSP_LIST            Tag = 0x19800003
	GPIO_RSP_LIST_ALIAS      Tag = 0x19800004
	GPIO_RSP_LIST_REAL       Tag = 0x19800005
	GPIO_RSP_CONNECT         Tag = 0x19800006
	GPIO_RSP_CHANGECOUNTER   Tag = 0x19800007
	GPIO_RSP_REMOTE_EXCHANGE Tag = 0x19800008
	GPIO_TUPEL               Tag = 0x19860001
	GPIO_NUMBER              Tag = 0x19860002
	GPIO_NAME                Tag = 0x19860003
	GPIO_VALUE               Tag = 0x19860004
	GPIO_KEY                 Tag = 0x19860005
	GPIO_AVAILABLE           Tag = 0x19860006
	GPIO_MODE                Tag = 0x19860007
	GPIO_SUPPORTED           Tag = 0x19860008
	GPIO_CONNECTEDTO         Tag = 0x19860009
	GPIO_USERLEVEL           Tag = 0x1986000A
	GPIO_INACTIVE            Tag = 0x1986000B
	GPIO_SETTINGS            Tag = 0x1986000C
	GPIO_REMOTE_LINKED       Tag = 0x1986000D
)

// ---------------
// NAMESPACE: FARM
// 0x1Axxxxxx
// undocumented
// ---------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	FARM_REQ_CONNECTED_DEVICES      Tag = 0x1A000001
	FARM_REQ_CONNECTED_DEVICES_REV  Tag = 0x1A000003
	FARM_REQ_LIST_AVAILABLE_DEVICES Tag = 0x1A000004
	FARM_PARAM_DEVICE               Tag = 0x1A040000
	FARM_PARAM_SERIALNO             Tag = 0x1A040001
	FARM_PARAM_CNAME                Tag = 0x1A040002
	FARM_CONNECTED_DEVICES_REV      Tag = 0x1A040003
	FARM_PARAM_SHORT_ID             Tag = 0x1A400001
	FARM_CONNECTED_DEVICES          Tag = 0x1A800001
	FARM_LIST_AVAILABLE_DEVICES     Tag = 0x1A800002
)

// -------------
// NAMESPACE: SE
// 0x1Bxxxxxx
// undocumented
// -------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	SE_REQ_SE_COUNT                        Tag = 0x1B000001
	SE_REQ_SE_DATA                         Tag = 0x1B000002
	SE_REQ_SET_POWER                       Tag = 0x1B000003
	SE_REQ_SET_DERATING                    Tag = 0x1B000004
	SE_REQ_SET_COUPLE_MODE                 Tag = 0x1B000005
	SE_REQ_COUPLE_MODE                     Tag = 0x1B000006
	SE_REQ_SET_BRI                         Tag = 0x1B000007
	SE_REQ_GET_BRI                         Tag = 0x1B000008
	SE_REQ_EP_RESERVE                      Tag = 0x1B000009
	SE_REQ_GET_PROTECTION_STRATEGY         Tag = 0x1B00000A
	SE_REQ_SET_EP_RESERVE                  Tag = 0x1B000010
	SE_REQ_GET_ESTIMATED_POWER_LIMIT       Tag = 0x1B000011
	SE_REQ_DESIGN_LIMIT                    Tag = 0x1B000012
	SE_SET_RESTART_SWITCH_TO_EP            Tag = 0x1B000013
	SE_EP_STATUS                           Tag = 0x1B000014
	SE_REQ_DIAGNOSIS_ERRORS                Tag = 0x1B000015
	SE_REQ_RESET_POWERSAVE_TIMEOUT         Tag = 0x1B000027
	SE_REQ_EMERGENCY_POWER_OVERLOAD_STATUS Tag = 0x1B000028
	SE_REQ_EMERGENCY_POWER_RETRY           Tag = 0x1B000029
	SE_REQ_IS_EMERGENCYPOWER_POSSIBLE      Tag = 0x1B000030
	SE_REQ_SET_PROTECTION_STRATEGY         Tag = 0x1B000031
	SE_PARAM_INDEX                         Tag = 0x1B040000
	SE_PARAM_DCDC_STATUS                   Tag = 0x1B040001
	SE_PARAM_BAT_STATUS                    Tag = 0x1B040002
	SE_PARAM_CTRL_STATE                    Tag = 0x1B040003
	SE_PARAM_PVPOWER                       Tag = 0x1B040004
	SE_PARAM_PVENERGY                      Tag = 0x1B040005
	SE_PARAM_BATTERYPOWER                  Tag = 0x1B040006
	SE_PARAM_BATCAPACITY                   Tag = 0x1B040007
	SE_PARAM_LIMITS                        Tag = 0x1B040008
	SE_PARAM_DESIREDPOWER                  Tag = 0x1B040009
	SE_PARAM_DESIREDDERATING               Tag = 0x1B040010
	SE_PARAM_INT                           Tag = 0x1B040011
	SE_PARAM_UINT                          Tag = 0x1B040012
	SE_PARAM_FLOAT                         Tag = 0x1B040013
	SE_PARAM_EMERGENCYMODE                 Tag = 0x1B040014
	SE_PARAM_PVI1_STATUS                   Tag = 0x1B040020
	SE_PARAM_PVI2_STATUS                   Tag = 0x1B040021
	SE_PARAM_PVI3_STATUS                   Tag = 0x1B040022
	SE_PARAM_EP_RESERVE                    Tag = 0x1B040023
	SE_PARAM_TIME_LAST_FULL                Tag = 0x1B040024
	SE_PARAM_TIME_LAST_EMPTY               Tag = 0x1B040025
	SE_PARAM_LAST_SOC                      Tag = 0x1B040026
	SE_PARAM_EP_STATUS                     Tag = 0x1B040027
	SE_PARAM_TIME_TO_RETRY                 Tag = 0x1B040030
	SE_PARAM_NO_REMAINING_RETRY            Tag = 0x1B040031
	SE_PARAM_EP_RESERVE_W                  Tag = 0x1B040033
	SE_PARAM_EP_RESERVE_MAX_W              Tag = 0x1B040034
	SE_PARAM_BOOL                          Tag = 0x1B400001
	SE_PARAM_BRI_INDEX                     Tag = 0x1B400002
	SE_SE_COUNT                            Tag = 0x1B800001
	SE_SE_DATA                             Tag = 0x1B800002
	SE_SET_POWER                           Tag = 0x1B800003
	SE_SET_DERATE                          Tag = 0x1B800004
	SE_SET_COUPLE_MODE                     Tag = 0x1B800005
	SE_COUPLE_MODE                         Tag = 0x1B800006
	SE_SET_BRI                             Tag = 0x1B800007
	SE_GET_BRI                             Tag = 0x1B800008
	SE_EP_RESERVE                          Tag = 0x1B800009
	SE_GET_PROTECTION_STRATEGY             Tag = 0x1B80000A
	SE_GET_ESTIMATED_POWER_LIMIT           Tag = 0x1B800011
	SE_DESIGN_LIMIT                        Tag = 0x1B800012
	SE_REQ_SET_RESTART_SWITCH_TO_EP        Tag = 0x1B800013
	SE_REQ_EP_STATUS                       Tag = 0x1B800014
	SE_DIAGNOSIS_ERRORS                    Tag = 0x1B800015
	SE_EMERGENCY_POWER_OVERLOAD_STATUS     Tag = 0x1B800028
	SE_EMERGENCY_POWER_RETRY               Tag = 0x1B800029
	SE_IS_EMERGENCYPOWER_POSSIBLE          Tag = 0x1B800030
	SE_SET_PROTECTION_STRATEGY             Tag = 0x1B800031
)

// --------------
// NAMESPACE: QPI
// 0x1Cxxxxxx
// undocumented
// --------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	QPI_REQ_INVERTER_COUNT                Tag = 0x1C000001
	QPI_REQ_INVERTER_DATA                 Tag = 0x1C000002
	QPI_REQ_UPDATE_FIRMWARE               Tag = 0x1C000003
	QPI_REQ_UPDATE_STATUS                 Tag = 0x1C000004
	QPI_REQ_INVERTER_SET_VALUES           Tag = 0x1C000005
	QPI_REQ_RESET_STATE_1_COUNTER         Tag = 0x1C000006
	QPI_REQ_STATE_1_COUNTER               Tag = 0x1C000007
	QPI_REQ_INVERTER_SET_POWER            Tag = 0x1C000008
	QPI_REQ_SET_BAT_INFO                  Tag = 0x1C000009
	QPI_REQ_GET_PARAM                     Tag = 0x1C00000A
	QPI_REQ_SET_PARAM                     Tag = 0x1C00000B
	QPI_REQ_ERR_LIST                      Tag = 0x1C00000C
	QPI_REQ_ERR_LIST_4105                 Tag = 0x1C00000D
	QPI_REQ_CLEAR_ERR_HIST                Tag = 0x1C00000E
	QPI_REQ_STATE_0                       Tag = 0x1C00000F
	QPI_REQ_DEBUG_DATA                    Tag = 0x1C000010
	QPI_REQ_INVERTER_COUNT_DETAIL         Tag = 0x1C000011
	QPI_REQ_SELECTED_INVERTER_DATA        Tag = 0x1C000012
	QPI_REQ_ERR_HIST                      Tag = 0x1C000013
	QPI_REQ_HW_INFO                       Tag = 0x1C000014
	QPI_REQ_VERSION_RESET                 Tag = 0x1C000015
	QPI_REQ_SET_DESIRED_CURRENT_FOR_EP    Tag = 0x1C000016
	QPI_REQ_GET_DESIRED_CURRENT_FOR_EP    Tag = 0x1C000017
	QPI_REQ_SETTINGS_EP_ENABLED           Tag = 0x1C000030
	QPI_REQ_SET_SETTINGS_EP_ENABLED       Tag = 0x1C00003A
	QPI_REQ_SETTINGS_VDE_2510_ENABLED     Tag = 0x1C00003B
	QPI_REQ_SET_SETTINGS_VDE_2510_ENABLED Tag = 0x1C00003C
	QPI_PARAM_INDEX                       Tag = 0x1C040000
	QPI_PARAM_U_BAT                       Tag = 0x1C040001
	QPI_PARAM_I_BAT                       Tag = 0x1C040002
	QPI_PARAM_U_AC                        Tag = 0x1C040003
	QPI_PARAM_I_AC                        Tag = 0x1C040004
	QPI_PARAM_PHI                         Tag = 0x1C040005
	QPI_PARAM_POWER                       Tag = 0x1C040006
	QPI_PARAM_APP_POWER                   Tag = 0x1C040007
	QPI_PARAM_REA_POWER                   Tag = 0x1C040008
	QPI_PARAM_FILE_NAME                   Tag = 0x1C040009
	QPI_PARAM_PROGRESS                    Tag = 0x1C040010
	QPI_PARAM_CHILD                       Tag = 0x1C040011
	QPI_PARAM_POWER_L1                    Tag = 0x1C040012
	QPI_PARAM_POWER_L2                    Tag = 0x1C040013
	QPI_PARAM_POWER_L3                    Tag = 0x1C040014
	QPI_PARAM_TIME                        Tag = 0x1C040015
	QPI_PARAM_STATE_1_COUNT               Tag = 0x1C040016
	QPI_PARAM_CMD                         Tag = 0x1C040017
	QPI_PARAM_STATE_1_PS                  Tag = 0x1C400018
	QPI_PARAM_U_MAX                       Tag = 0x1C400019
	QPI_PARAM_U_MIN                       Tag = 0x1C400020
	QPI_PARAM_I_MAX                       Tag = 0x1C400021
	QPI_PARAM_I_MIN                       Tag = 0x1C400022
	QPI_PARAM_BLOCK                       Tag = 0x1C400023
	QPI_PARAM_ITEM                        Tag = 0x1C400024
	QPI_PARAM_VALUE                       Tag = 0x1C400025
	QPI_PARAM_ERR                         Tag = 0x1C400030
	QPI_PARAM_ERR_STR                     Tag = 0x1C400031
	QPI_PARAM_ERR_DATE                    Tag = 0x1C400032
	QPI_PARAM_ERR_CODE                    Tag = 0x1C400033
	QPI_PARAM_ERR_F_MIN                   Tag = 0x1C400034
	QPI_PARAM_ERR_F_MAX                   Tag = 0x1C400035
	QPI_PARAM_ERR_U_MIN                   Tag = 0x1C400036
	QPI_PARAM_ERR_U_MAX                   Tag = 0x1C400037
	QPI_PARAM_ERR_U_L1_PHI                Tag = 0x1C400038
	QPI_PARAM_ERR_U_L1_RMS                Tag = 0x1C400039
	QPI_PARAM_ERR_U_L2_PHI                Tag = 0x1C400040
	QPI_PARAM_ERR_U_L2_RMS                Tag = 0x1C400041
	QPI_PARAM_ERR_U_L3_PHI                Tag = 0x1C400042
	QPI_PARAM_ERR_U_L3_RMS                Tag = 0x1C400043
	QPI_PARAM_ERR_U_INV_PHI               Tag = 0x1C400044
	QPI_PARAM_ERR_U_INV_RMS               Tag = 0x1C400045
	QPI_PARAM_ERR_I_DCL_RMS               Tag = 0x1C400046
	QPI_PARAM_ERR_U_DCL                   Tag = 0x1C400047
	QPI_PARAM_ERR_I_LOAD_PHI              Tag = 0x1C400048
	QPI_PARAM_ERR_I_LOAD_RMS              Tag = 0x1C400049
	QPI_PARAM_ERR_REGULATOR_OUT           Tag = 0x1C400050
	QPI_PARAM_ERR_STATE_MASHINE_E_STATE   Tag = 0x1C400051
	QPI_PARAM_ERR_STATE_ACT_REG           Tag = 0x1C400052
	QPI_PARAM_ERR_TMP_0                   Tag = 0x1C400053
	QPI_PARAM_ERR_TMP_1                   Tag = 0x1C400054
	QPI_PARAM_ERR_TMP_2                   Tag = 0x1C400055
	QPI_PARAM_ERR_TMP_3                   Tag = 0x1C400056
	QPI_PARAM_ERR_CO_PRO                  Tag = 0x1C400057
	QPI_PARAM_ERR_TREATMENT               Tag = 0x1C400058
	QPI_PARAM_ERR_I_LOAD_T0               Tag = 0x1C400059
	QPI_PARAM_ERR_U_INV_TO                Tag = 0x1C400060
	QPI_PARAM_U_AC_L1                     Tag = 0x1C400061
	QPI_PARAM_U_AC_L2                     Tag = 0x1C400062
	QPI_PARAM_U_AC_L3                     Tag = 0x1C400063
	QPI_PARAM_I_AC_L1                     Tag = 0x1C400064
	QPI_PARAM_I_AC_L2                     Tag = 0x1C400065
	QPI_PARAM_I_AC_L3                     Tag = 0x1C400066
	QPI_PARAM_APP_POWER_L1                Tag = 0x1C400067
	QPI_PARAM_APP_POWER_L2                Tag = 0x1C400068
	QPI_PARAM_APP_POWER_L3                Tag = 0x1C400069
	QPI_PARAM_REA_POWER_L1                Tag = 0x1C400070
	QPI_PARAM_REA_POWER_L2                Tag = 0x1C400071
	QPI_PARAM_REA_POWER_L3                Tag = 0x1C400072
	QPI_PARAM_STATE_0_STATE               Tag = 0x1C400073
	QPI_PARAM_STATE_0_ERR_COUNT_ALL       Tag = 0x1C400074
	QPI_PARAM_STATE_0_ERR_COUNT_ACTIVE    Tag = 0x1C400075
	QPI_PARAM_STATE_0_OP_STATE            Tag = 0x1C400076
	QPI_PARAM_CONF_STATE                  Tag = 0x1C400079
	QPI_PARAM_ACTIVATED                   Tag = 0x1C400080
	QPI_PARAM_INVERTER_GROUP              Tag = 0x1C400081
	QPI_PARAM_COUNT_DETAIL                Tag = 0x1C400082
	QPI_PARAM_DEBUG_SM                    Tag = 0x1C400083
	QPI_PARAM_DEBUG_ACTUAL_REG            Tag = 0x1C400084
	QPI_PARAM_DEBUG_U_DCL                 Tag = 0x1C400085
	QPI_PARAM_DEBUG_I_DCL_RMS             Tag = 0x1C400086
	QPI_PARAM_DEBUG_I_LOAD_RMS            Tag = 0x1C400087
	QPI_PARAM_DEBUG_I_LOAD_T0             Tag = 0x1C400088
	QPI_PARAM_DEBUG_U_INV_RMS             Tag = 0x1C400089
	QPI_PARAM_DEBUG_U_INV_T0              Tag = 0x1C400090
	QPI_PARAM_DEBUG_U_L1_RMS              Tag = 0x1C400091
	QPI_PARAM_DEBUG_U_L2_RMS              Tag = 0x1C400092
	QPI_PARAM_DEBUG_U_L3_RMS              Tag = 0x1C400093
	QPI_PARAM_DEBUG_U_L1_T0               Tag = 0x1C400094
	QPI_PARAM_DEBUG_U_L2_T0               Tag = 0x1C400095
	QPI_PARAM_DEBUG_U_L3_T0               Tag = 0x1C400096
	QPI_PARAM_DEBUG_TMP_0                 Tag = 0x1C400097
	QPI_PARAM_DEBUG_TMP_1                 Tag = 0x1C400098
	QPI_PARAM_DEBUG_TMP_2                 Tag = 0x1C400099
	QPI_PARAM_DEBUG_TMP_3                 Tag = 0x1C400100
	QPI_PARAM_DEBUG_F_LINE                Tag = 0x1C400101
	QPI_PARAM_DEBUG_I_DCL_AVG             Tag = 0x1C400102
	QPI_PARAM_DEBUG_U_L1_PHI              Tag = 0x1C400103
	QPI_PARAM_DEBUG_U_L2_PHI              Tag = 0x1C400104
	QPI_PARAM_DEBUG_U_L3_PHI              Tag = 0x1C400105
	QPI_PARAM_DEBUG_INV_PHI               Tag = 0x1C400106
	QPI_PARAM_DEBUG_I_LOAD_PHI            Tag = 0x1C400107
	QPI_PARAM_NUMBER_CHILDS               Tag = 0x1C400108
	QPI_PARAM_INVERTER_STORED_SI_GROUP    Tag = 0x1C400109
	QPI_PARAM_DEBUG_U_L1_RMS_COPRO        Tag = 0x1C40010A
	QPI_PARAM_DEBUG_U_L2_RMS_COPRO        Tag = 0x1C40010B
	QPI_PARAM_DEBUG_U_L3_RMS_COPRO        Tag = 0x1C40010C
	QPI_PARAM_DEBUG_F_LINE_COPRO          Tag = 0x1C40010D
	QPI_PARAM_SW_VERSION_DATE             Tag = 0x1C40010E
	QPI_PARAM_SW_VERSION                  Tag = 0x1C40010F
	QPI_PARAM_INVERTER_EP_LINE            Tag = 0x1C400110
	QPI_PARAM_COPRO_SW_VERSION_DATE       Tag = 0x1C400111
	QPI_PARAM_COPRO_SW_VERSION            Tag = 0x1C400112
	QPI_PARAM_COPRO_SW_SVN                Tag = 0x1C400113
	QPI_PARAM_HW_VERSION_MAIN             Tag = 0x1C400114
	QPI_PARAM_HW_VERSION_COPRO            Tag = 0x1C400115
	QPI_PARAM_HW_VERSION_PCB_CODE         Tag = 0x1C400116
	QPI_PARAM_BOARD_SERIAL                Tag = 0x1C400117
	QPI_PARAM_MODULE_SERIAL               Tag = 0x1C400118
	QPI_PARAM_ERR_F_LINE                  Tag = 0x1C400119
	QPI_PARAM_ERR_OPT_STATE               Tag = 0x1C40011A
	QPI_PARAM_RT_RESULT                   Tag = 0x1C40011B
	QPI_PARAM_DOOR_SW_OPEN                Tag = 0x1C40011C
	QPI_PARAM_FAN_REQESTED                Tag = 0x1C40011D
	QPI_PARAM_DEBUG_COPRO_STATE           Tag = 0x1C40011E
	QPI_PARAM_ERR_I_DCL_T0                Tag = 0x1C40011F
	QPI_PARAM_ERR_COPRO_U_L1_RMS          Tag = 0x1C400120
	QPI_PARAM_ERR_COPRO_U_L2_RMS          Tag = 0x1C400121
	QPI_PARAM_ERR_COPRO_U_L3_RMS          Tag = 0x1C400122
	QPI_PARAM_ERR_COPRO_U_INV_RMS         Tag = 0x1C400123
	QPI_PARAM_ERR_COPRO_REL_STATE         Tag = 0x1C400124
	QPI_PARAM_ERR_COPRO_MODE              Tag = 0x1C400125
	QPI_PARAM_ERR_COPRO_F_LINE            Tag = 0x1C400126
	QPI_PARAM_ERR_COPRO_PEN_TV_STATE      Tag = 0x1C400127
	QPI_PARAM_ERR_COPRO_ERR_DATA          Tag = 0x1C400128
	QPI_PARAM_ERR_COPRO_TRIPP_LN_OUT      Tag = 0x1C400129
	QPI_PARAM_ERR_COPRO_TRIPP_LN_IN       Tag = 0x1C40012A
	QPI_PARAM_ERR_COPRO_U_PEN_AVG         Tag = 0x1C40012B
	QPI_PARAM_ERR_I_PRIM                  Tag = 0x1C40012C
	QPI_PARAM_SW_COUNTRY                  Tag = 0x1C40012D
	QPI_PARAM_SELECTEC_COUNTRY            Tag = 0x1C40012E
	QPI_PARAM_COPRO_SW_COUNTRY            Tag = 0x1C40012F
	QPI_PARAM_COPRO_SELECTEC_COUNTRY      Tag = 0x1C400130
	QPI_PARAM_MAX_AC_APPARENT_POWER       Tag = 0x1C400131
	QPI_PARAM_DEBUG_U_INV_RMS_COPRO       Tag = 0x1C400132
	QPI_PARAM_DEBUG_I_PRIM_RMS            Tag = 0x1C400133
	QPI_PARAM_SW_SVN                      Tag = 0x1C400134
	QPI_INVERTER_COUNT                    Tag = 0x1C800001
	QPI_INVERTER_DATA                     Tag = 0x1C800002
	QPI_UPDATE_FIRMWARE                   Tag = 0x1C800003
	QPI_UPDATE_STATUS                     Tag = 0x1C800004
	QPI_INVERTER_SET_VALUES               Tag = 0x1C800005
	QPI_RESET_STATE_1_COUNTER             Tag = 0x1C800006
	QPI_STATE_1_COUNTER                   Tag = 0x1C800007
	QPI_INVERTER_SET_POWER                Tag = 0x1C800008
	QPI_SET_BAT_INFO                      Tag = 0x1C800009
	QPI_GET_PARAM                         Tag = 0x1C80000A
	QPI_SET_PARAM                         Tag = 0x1C80000B
	QPI_ERR_LIST                          Tag = 0x1C80000C
	QPI_ERR_LIST_4105                     Tag = 0x1C80000D
	QPI_CLEAR_ERR_HIST                    Tag = 0x1C80000E
	QPI_STATE_0                           Tag = 0x1C80000F
	QPI_DEBUG_DATA                        Tag = 0x1C800010
	QPI_INVERTER_COUNT_DETAIL             Tag = 0x1C800011
	QPI_SELECTED_INVERTER_DATA            Tag = 0x1C800012
	QPI_ERR_HIST                          Tag = 0x1C800013
	QPI_HW_INFO                           Tag = 0x1C800014
	QPI_VERSION_RESET                     Tag = 0x1C800015
	QPI_SET_DESIRED_CURRENT_FOR_EP        Tag = 0x1C800016
	QPI_GET_DESIRED_CURRENT_FOR_EP        Tag = 0x1C800017
	QPI_SETTINGS_EP_ENABLED               Tag = 0x1C800030
	QPI_SET_SETTINGS_EP_ENABLED           Tag = 0x1C80003A
	QPI_SETTINGS_VDE_2510_ENABLED         Tag = 0x1C80003B
	QPI_SET_SETTINGS_VDE_2510_ENABLED     Tag = 0x1C80003C
)

// ---------------
// NAMESPACE: GAPP
// 0x1Dxxxxxx
// undocumented
// ---------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	GAPP_REQ_DEV_COUNT                               Tag = 0x1D000001
	GAPP_REQ_SERIALNO                                Tag = 0x1D000002
	GAPP_REQ_SUPPORTED_REACTIVE_POWER_FUNCTIONS      Tag = 0x1D000003
	GAPP_REQ_ENABLED_REACTIVE_POWER_FUNCTIONS        Tag = 0x1D000004
	GAPP_REQ_SET_REACTIVE_POWER_FUNCTIONS            Tag = 0x1D000005
	GAPP_REQ_SET_REACTIVE_POWER_FUNCTIONS_PARAMETER  Tag = 0x1D000006
	GAPP_REQ_REACTIVE_POWER_FUNCTIONS_PARAMETER      Tag = 0x1D000007
	GAPP_REQ_SUPPORTED_ACTIVE_POWER_FUNCTIONS        Tag = 0x1D000008
	GAPP_REQ_ENABLED_ACTIVE_POWER_FUNCTIONS_PU       Tag = 0x1D000009
	GAPP_REQ_SET_ACTIVE_POWER_FUNCTIONS_PU           Tag = 0x1D00000A
	GAPP_REQ_SET_ACTIVE_POWER_FUNCTIONS_PU_PARAMETER Tag = 0x1D00000B
	GAPP_REQ_ACTIVE_POWER_FUNCTION_PU_PARAMETER      Tag = 0x1D00000C
	GAPP_REQ_REACTIVE_POWER_SETTINGS_EQUAL           Tag = 0x1D00000D
	GAPP_REQ_ACTIVE_POWER_SETTINGS_EQUAL             Tag = 0x1D00000E
	GAPP_REQ_SUPPORTED_GRID_PROTECTION_FUNCTIONS     Tag = 0x1D00000F
	GAPP_REQ_GRID_PROTECTION_FUNCTION_PARAMETER      Tag = 0x1D000010
	GAPP_REQ_SET_GRID_PROTECTION_FUNCTION_PARAMETER  Tag = 0x1D000011
	GAPP_PARAM_INDEX                                 Tag = 0x1D040000
	GAPP_PARAM_SERIALNO                              Tag = 0x1D040001
	GAPP_PARAM_REACTIVE_POWER_FUNCTION               Tag = 0x1D040002
	GAPP_PARAM_ACTIVE_POWER_FUNCTION_PU              Tag = 0x1D040003
	GAPP_PARAM_GAPP_PARAMETER                        Tag = 0x1D040004
	GAPP_PARAM_GAPP_PARAMETER_FUNCTION               Tag = 0x1D040005
	GAPP_PARAM_GAPP_PARAMETER_VALUE_LIST             Tag = 0x1D040006
	GAPP_PARAM_GAPP_PARAMETER_VALUE_LIST_ENTRY       Tag = 0x1D040007
	GAPP_PARAM_GAPP_PARAMETER_SCALE_FACTOR_X         Tag = 0x1D040008
	GAPP_PARAM_GAPP_PARAMETER_SCALE_FACTOR_Y         Tag = 0x1D040009
	GAPP_PARAM_GAPP_PARAMETER_VALUE_MAX              Tag = 0x1D040010
	GAPP_PARAM_GAPP_PARAMETER_VALUE_MIN              Tag = 0x1D040011
	GAPP_PARAM_GAPP_PARAMETER_VALUE                  Tag = 0x1D040012
	GAPP_PARAM_GAPP_PARAMETER_VALUE_INIT             Tag = 0x1D040013
	GAPP_PARAM_SUPPORTED_REACTIVE_POWER_FUNCTIONS    Tag = 0x1D040014
	GAPP_PARAM_SUPPORTED_ACTIVE_POWER_FUNCTIONS      Tag = 0x1D040015
	GAPP_PARAM_SUCCESS                               Tag = 0x1D040016
	GAPP_PARAM_GAPP_PARAMETER_HAS_Y                  Tag = 0x1D040017
	GAPP_PARAM_GAPP_PARAMETER_COUNT_MIN              Tag = 0x1D040018
	GAPP_PARAM_GAPP_PARAMETER_COUNT_MAX              Tag = 0x1D040019
	GAPP_PARAM_GAPP_PARAMETER_COUNT_USED             Tag = 0x1D040020
	GAPP_PARAM_SUPPORTED_GRID_PROTECTION_FUNCTIONS   Tag = 0x1D040021
	GAPP_DEV_COUNT                                   Tag = 0x1D800001
	GAPP_SERIALNO                                    Tag = 0x1D800002
	GAPP_SUPPORTED_REACTIVE_POWER_FUNCTIONS          Tag = 0x1D800003
	GAPP_ENABLED_REACTIVE_POWER_FUNCTIONS            Tag = 0x1D800004
	GAPP_SET_REACTIVE_POWER_FUNCTIONS                Tag = 0x1D800005
	GAPP_SET_REACTIVE_POWER_FUNCTIONS_PARAMETER      Tag = 0x1D800006
	GAPP_REACTIVE_POWER_FUNCTIONS_PARAMETER          Tag = 0x1D800007
	GAPP_SUPPORTED_ACTIVE_POWER_FUNCTIONS            Tag = 0x1D800008
	GAPP_ENABLED_ACTIVE_POWER_FUNCTIONS_PU           Tag = 0x1D800009
	GAPP_SET_ACTIVE_POWER_FUNCTIONS_PU               Tag = 0x1D80000A
	GAPP_SET_ACTIVE_POWER_FUNCTIONS_PU_PARAMETER     Tag = 0x1D80000B
	GAPP_ACTIVE_POWER_FUNCTION_PU_PARAMETER          Tag = 0x1D80000C
	GAPP_REACTIVE_POWER_SETTINGS_EQUAL               Tag = 0x1D80000D
	GAPP_ACTIVE_POWER_SETTINGS_EQUAL                 Tag = 0x1D80000E
	GAPP_SUPPORTED_GRID_PROTECTION_FUNCTIONS         Tag = 0x1D80000F
	GAPP_GRID_PROTECTION_FUNCTION_PARAMETER          Tag = 0x1D800010
	GAPP_SET_GRID_PROTECTION_FUNCTION_PARAMETER      Tag = 0x1D800011
)

// ----------------
// NAMESPACE: EMSPR
// 0x1Exxxxxx
// undocumented
// ----------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	EMSPR_REQ_GET_EQUATIONS          Tag = 0x1E000001
	EMSPR_REQ_SET_EQUATIONS          Tag = 0x1E000002
	EMSPR_REQ_GET_ACTIVE             Tag = 0x1E000003
	EMSPR_REQ_GET_PINCOUNT           Tag = 0x1E000004
	EMSPR_REQ_SET_PINCOUNT           Tag = 0x1E000005
	EMSPR_REQ_CHANGECOUNTER          Tag = 0x1E000006
	EMSPR_REQ_GET_INVERTERENABLING   Tag = 0x1E000007
	EMSPR_REQ_SET_INVERTERENABLING   Tag = 0x1E000008
	EMSPR_REQ_GET_INVERTERENABLEWAIT Tag = 0x1E00000B
	EMSPR_REQ_SET_INVERTERENABLEWAIT Tag = 0x1E00000C
	EMSPR_RSP_GET_EQUATIONS          Tag = 0x1E800001
	EMSPR_RSP_SET_EQUATIONS          Tag = 0x1E800002
	EMSPR_RSP_GET_ACTIVE             Tag = 0x1E800003
	EMSPR_RSP_GET_PINCOUNT           Tag = 0x1E800004
	EMSPR_RSP_SET_PINCOUNT           Tag = 0x1E800005
	EMSPR_RSP_CHANGECOUNTER          Tag = 0x1E800006
	EMSPR_RSP_GET_INVERTERENABLING   Tag = 0x1E800007
	EMSPR_RSP_SET_INVERTERENABLING   Tag = 0x1E800008
	EMSPR_RSP_GET_INVERTERENABLEWAIT Tag = 0x1E80000B
	EMSPR_RSP_SET_INVERTERENABLEWAIT Tag = 0x1E80000C
	EMSPR_EQUATION                   Tag = 0x1E860001
	EMSPR_INPUT                      Tag = 0x1E860002
	EMSPR_MASK                       Tag = 0x1E860003
	EMSPR_INVALID                    Tag = 0x1E860004
	EMSPR_OUTPUT                     Tag = 0x1E860005
	EMSPR_ISACTIVE                   Tag = 0x1E860006
	EMSPR_FAILURESTATE               Tag = 0x1E860007
	EMSPR_INVERTERENABLED            Tag = 0x1E860008
)

// ----------------
// NAMESPACE: IOBOX
// 0x1Fxxxxxx
// undocumented
// ----------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	IOBOX_REQ_IDENTIFY         Tag = 0x1F000001
	IOBOX_REQ_GETCOUNT         Tag = 0x1F000002
	IOBOX_REQ_SEARCH           Tag = 0x1F000003
	IOBOX_REQ_SERIAL_NUMBER    Tag = 0x1F000004
	IOBOX_REQ_STATE            Tag = 0x1F000005
	IOBOX_REQ_LIST             Tag = 0x1F000006
	IOBOX_REQ_REMOVE           Tag = 0x1F000007
	IOBOX_REQ_DATA             Tag = 0x1F040000
	IOBOX_INDEX                Tag = 0x1F040001
	IOBOX_ITEM                 Tag = 0x1F400001
	IOBOX_ITEM_INDEX           Tag = 0x1F400002
	IOBOX_ITEM_NAME            Tag = 0x1F400003
	IOBOX_ITEM_SERIAL_NUMBER   Tag = 0x1F400004
	IOBOX_ITEM_STATUS          Tag = 0x1F400005
	IOBOX_ITEM_SWVERSION       Tag = 0x1F400006
	IOBOX_ITEM_ADDITIONAL_INFO Tag = 0x1F400007
	IOBOX_RSP_IDENTIFY         Tag = 0x1F800001
	IOBOX_RSP_GETCOUNT         Tag = 0x1F800002
	IOBOX_RSP_SEARCH           Tag = 0x1F800003
	IOBOX_RSP_SERIAL_NUMBER    Tag = 0x1F800004
	IOBOX_RSP_STATE            Tag = 0x1F800005
	IOBOX_LIST                 Tag = 0x1F800006
	IOBOX_REMOVE               Tag = 0x1F800007
	IOBOX_DATA                 Tag = 0x1F840000
)

// --------------
// NAMESPACE: WBD
// 0x20xxxxxx
// undocumented
// --------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	WBD_REQ_START_SCAN     Tag = 0x20000001
	WBD_REQ_IS_SCANNING    Tag = 0x20000002
	WBD_REQ_CREATE_WB      Tag = 0x20000003
	WBD_REQ_CANCEL_SCAN    Tag = 0x20000004
	WBD_REQ_DELETE_WALLBOX Tag = 0x20000005
	WBD_START_SCAN         Tag = 0x20800001
	WBD_IS_SCANNING        Tag = 0x20800002
	WBD_CREATE_WB          Tag = 0x20800003
	WBD_CANCEL_CAN         Tag = 0x20800004
	WBD_DELETE_WALL_BOX    Tag = 0x20800005
)

// ---------------
// NAMESPACE: REFU
// 0x21xxxxxx
// undocumented
// ---------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	REFU_REQ_START_SCAN        Tag = 0x21000001
	REFU_REQ_IS_SCANNING       Tag = 0x21000002
	REFU_REQ_CREATE_INV        Tag = 0x21000003
	REFU_REQ_CANCEL_SCAN       Tag = 0x21000004
	REFU_REQ_DELETE_INVERTER   Tag = 0x21000005
	REFU_REQ_NO_INVERTERS      Tag = 0x21000006
	REFU_REQ_CONNECTED_DEVICES Tag = 0x21000008
	REFU_PARAM_MAC             Tag = 0x21400001
	REFU_PARAM_IP              Tag = 0x21400002
	REFU_PARAM_ALIVE           Tag = 0x21400003
	REFU_PARAM_INDEX           Tag = 0x21400004
	REFU_PARAM_DHCP            Tag = 0x21400005
	REFU_START_SCAN            Tag = 0x21800001
	REFU_IS_SCANNING           Tag = 0x21800002
	REFU_CREATE_INV            Tag = 0x21800003
	REFU_CANCEL_CAN            Tag = 0x21800004
	REFU_DELETE_INVERTER       Tag = 0x21800005
	REFU_NO_INVERTERS          Tag = 0x21800006
	REFU_CONNECTED_DEVICES     Tag = 0x21800008
)

// --------------
// NAMESPACE: OVP
// 0x22xxxxxx
// undocumented
// --------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	OVP_REQ_STATUS   Tag = 0x22000001
	OVP_REQ_RESET    Tag = 0x22000002
	OVP_PARAM_INDEX  Tag = 0x22400001
	OVP_PARAM_STATUS Tag = 0x22800001
	OVP_PARAM_RESET  Tag = 0x22800002
	OVP_STATUS       Tag = 0x22800003
	OVP_RESET        Tag = 0x22800004
)

// ------------------
// NAMESPACE: NETWORK
// 0x23xxxxxx
// undocumented
// ------------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	NETWORK_REQ_PING                    Tag = 0x23000001
	NETWORK_REQ_INFO                    Tag = 0x23000002
	NETWORK_REQ_DIAGNOSE                Tag = 0x23000003
	NETWORK_REQ_SET_IP                  Tag = 0x23000004
	NETWORK_REQ_SET_SUBNETMASK          Tag = 0x23000005
	NETWORK_REQ_SET_GATEWAY             Tag = 0x23000006
	NETWORK_REQ_SET_DNS_SERVER          Tag = 0x23000007
	NETWORK_REQ_SET_DHCP                Tag = 0x23000008
	NETWORK_REQ_MAKE_INFO_FILE          Tag = 0x23000009
	NETWORK_REQ_CHECK_INTERNET_SPEED    Tag = 0x2300000A
	NETWORK_PARAM_IP                    Tag = 0x23400001
	NETWORK_PARAM_SUBNETMASK            Tag = 0x23400002
	NETWORK_PARAM_GATEWAY               Tag = 0x23400003
	NETWORK_PARAM_DNS_SERVER            Tag = 0x23400004
	NETWORK_PARAM_DHCP                  Tag = 0x23400005
	NETWORK_PARAM_MAC_ADDRESS           Tag = 0x23400006
	NETWORK_PARAM_E3DC_SERVER_REACHABLE Tag = 0x23400007
	NETWORK_PARAM_DNS_SERVER_REACHABLE  Tag = 0x23400008
	NETWORK_PARAM_DNS_LOOKUP_WORKING    Tag = 0x23400009
	NETWORK_PARAM_GATEWAY_REACHABLE     Tag = 0x2340000A
	NETWORK_PARAM_SYSTEM_IP_VALID       Tag = 0x2340000B
	NETWORK_PARAM_CABLE_CONNECTED       Tag = 0x2340000C
	NETWORK_PARAM_INTERNET_SPEED        Tag = 0x2340000D
	NETWORK_PING                        Tag = 0x23800001
	NETWORK_INFO                        Tag = 0x23800002
	NETWORK_DIAGNOSE                    Tag = 0x23800003
	NETWORK_SET_IP                      Tag = 0x23800004
	NETWORK_SET_SUBNETMASK              Tag = 0x23800005
	NETWORK_SET_GATEWAY                 Tag = 0x23800006
	NETWORK_SET_DNS_SERVER              Tag = 0x23800007
	NETWORK_SET_DHCP                    Tag = 0x23800008
	NETWORK_MAKE_INFO_FILE              Tag = 0x23800009
	NETWORK_CHECK_INTERNET_SPEED        Tag = 0x2380000A
)

// -----------------
// NAMESPACE: WBAUTH
// 0x24xxxxxx
// undocumented
// -----------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	WBAUTH_REQ_VERIFY_SESSION  Tag = 0x24000000
	WBAUTH_WALLBOX_ID          Tag = 0x24400001
	WBAUTH_SESSION_ID          Tag = 0x24400002
	WBAUTH_VERIFICATION_ENTRY  Tag = 0x24400003
	WBAUTH_ENTRY_TYPE          Tag = 0x24400004
	WBAUTH_ENTRY_DATA          Tag = 0x24400005
	WBAUTH_VERIFICATION_RESULT Tag = 0x24400006
	WBAUTH_ENTRY_SWAPPED       Tag = 0x24400007
	WBAUTH_VERIFY_SESSION      Tag = 0x24800000
)

// ---------------
// NAMESPACE: PLAY
// 0x25xxxxxx
// undocumented
// ---------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	PLAY_REQ_DATA  Tag = 0x25000001
	PLAY_SEND_TIME Tag = 0x25400001
	PLAY_DATA      Tag = 0x25800001
)

// --------------
// NAMESPACE: GDI
// 0x26xxxxxx
// undocumented
// --------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	GDI_KEY_2                         Tag = 0x26000000
	GDI_REQ_PUSH_DESC                 Tag = 0x26000001
	GDI_PARAM_CONTENT_DESC            Tag = 0x26400001
	GDI_PUSH_DATA                     Tag = 0x26400002
	GDI_PUSH_DATA_OKAY                Tag = 0x26400003
	GDI_PARAM_CONTENT                 Tag = 0x26400004
	GDI_METADATA                      Tag = 0x26400005
	GDI_DATADESCRIPTIONS              Tag = 0x26400006
	GDI_DATALOGGING                   Tag = 0x26400007
	GDI_LOGGINGTYPE                   Tag = 0x26400008
	GDI_HASHSUM                       Tag = 0x26400009
	GDI_DATATYPE                      Tag = 0x2640000A
	GDI_INDEX                         Tag = 0x2640000B
	GDI_NAME                          Tag = 0x2640000C
	GDI_DATATYPE_BOOL                 Tag = 0x26500001
	GDI_DATATYPE_INT8                 Tag = 0x26500002
	GDI_DATATYPE_UINT8                Tag = 0x26500003
	GDI_DATATYPE_INT16                Tag = 0x26500004
	GDI_DATATYPE_UINT16               Tag = 0x26500005
	GDI_DATATYPE_UINT32               Tag = 0x26500006
	GDI_DATATYPE_INT32                Tag = 0x26500007
	GDI_DATATYPE_INT64                Tag = 0x26500008
	GDI_DATATYPE_UINT64               Tag = 0x26500009
	GDI_DATATYPE_FLOAT32              Tag = 0x2650000A
	GDI_DATATYPE_DOUBLE64             Tag = 0x2650000B
	GDI_DATATYPE_BITFIELD             Tag = 0x2650000C
	GDI_DATATYPE_CSTRING              Tag = 0x2650000D
	GDI_CONTAINER                     Tag = 0x2650000E
	GDI_DATATYPE_TIMESTAMP            Tag = 0x2650000F
	GDI_DATATYPE_BYTEARRAY            Tag = 0x26500010
	GDI_DATATYPE_COMPRESSED_BYTEARRAY Tag = 0x26500011
	GDI_DATA_INDEXKEY                 Tag = 0x26500100
	GDI_INDEXKEY                      Tag = 0x26500101
	GDI_LOKAL_STORING                 Tag = 0x26500104
	GDI_LOKAL_STORING_AVG             Tag = 0x26500105
	GDI_LOKAL_STORING_MOD             Tag = 0x26500106
	GDI_LOKAL_STORING_AVG_MOD         Tag = 0x26500107
	GDI_PUSH_STORING                  Tag = 0x26500108
	GDI_PUSH_STORING_AVG              Tag = 0x26500109
	GDI_PUSH_STORING_MOD              Tag = 0x2650010A
	GDI_PUSH_STORING_AVG_MOD          Tag = 0x2650010B
	GDI_DATA_TIMEKEY                  Tag = 0x2650010C
	GDI_DATA_MANIFEST                 Tag = 0x2650010D
	GDI_PAKET_NUMBER                  Tag = 0x2650010E
	GDI_PUSH_INTERMISSION_IN_SEC      Tag = 0x2650010F
	GDI_DESC_VERSION                  Tag = 0x26500110
	GDI_LOGDESCRIPTIONS               Tag = 0x26500111
	GDI_DIAGNOSEDESCRIPTIONS          Tag = 0x26500112
	GDI_KEYS                          Tag = 0x26500113
	GDI_KEY_PRIMARY                   Tag = 0x26500114
	GDI_OFFSET                        Tag = 0x26500115
	GDI_LENGTH                        Tag = 0x26500116
	GDI_SNAPSHOT                      Tag = 0x26500117
	GDI_RESOLUTION                    Tag = 0x26500118
	GDI_SIZE                          Tag = 0x26500119
	GDI_TRIGGER_ON_INTERVAL           Tag = 0x2650011A
	GDI_TRIGGER_ON_CHANGE             Tag = 0x2650011B
	GDI_DELAY                         Tag = 0x2650011C
	GDI_ENTITY_MAX                    Tag = 0x2650011D
	GDI_USE_PUSH_SERVICE              Tag = 0x2650011E
	GDI_DB_MAX                        Tag = 0x2650011F
	GDI_USE_P_DB                      Tag = 0x26500120
	GDI_FIX_INTERVAL                  Tag = 0x26500121
	GDI_TRIGGER_ON_TIME               Tag = 0x26500122
	GDI_DATA_LENGTH                   Tag = 0x26500123
	GDI_TRIGGER_ON_NULL               Tag = 0x26500124
	GDI_USE_NP_DB                     Tag = 0x26500125
	GDI_NO_TRIGGER_ON_NULL            Tag = 0x26500126
	GDI_MODIFICATIONS                 Tag = 0x26500127
	GDI_MODIFICATION_MIN              Tag = 0x26500128
	GDI_MODIFICATION_MAX              Tag = 0x26500129
	GDI_MODIFICATION_AVG              Tag = 0x2650012A
	GDI_MODIFICATION_AND              Tag = 0x2650012B
	GDI_MODIFICATION_OR               Tag = 0x2650012C
	GDI_MODIFICATION_INTEGRATE        Tag = 0x2650012D
	GDI_MODIFICATION_AUTOINC          Tag = 0x2650012E
	GDI_DESCRIPTION                   Tag = 0x2650012F
	GDI_DESCRIPTION_WITH_SIZE         Tag = 0x26500130
	GDI_GDI_TRIGGER_ON_CHANGE         Tag = 0x26500131
	GDI_DATA_MODIFICATION             Tag = 0x26500132
	GDI_PUSH_CONFIG                   Tag = 0x26500133
	GDI_PUSH_DELAY                    Tag = 0x26500134
	GDI_PUSH_CHANNEL                  Tag = 0x26500135
	GDI_PUSH_DESC                     Tag = 0x26800001
)

// --------------
// NAMESPACE: SCM
// 0x27xxxxxx
// undocumented
// --------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	SCM_REQ_SET_CONFIG                    Tag = 0x27000001
	SCM_CONFIG_PROCESSED                  Tag = 0x27000002
	SCM_CONFIG_PROCESSED_STATE            Tag = 0x27000003
	SCM_REQ_CONFIG_PROCESSED_STATE        Tag = 0x27000004
	SCM_REQ_ACCEPTED_CONFIG_FILE_VERSION  Tag = 0x27000005
	SCM_REQ_SET_CONFIG_RESET              Tag = 0x27000006
	SCM_LC_REQ_GET_TYPE_MEASURE           Tag = 0x27000007
	SCM_LC_REQ_GET_POWER_MGT_CONFIG       Tag = 0x27000008
	SCM_LC_REQ_SET_TYPE_MEASURE           Tag = 0x27000009
	SCM_LC_REQ_SET_POWER_MGT_CONFIG       Tag = 0x2700000A
	SCM_CONFIG_ID                         Tag = 0x27400001
	SCM_CONFIG_FILE_VALID                 Tag = 0x27400002
	SCM_LC_PARAM_POWER_MGT_TYPE           Tag = 0x27400003
	SCM_LC_PARAM_TOKEN                    Tag = 0x27400004
	SCM_LC_PARAM_MAIN_FUSE                Tag = 0x27400005
	SCM_LC_PARAM_MAX_CURRENT              Tag = 0x27400006
	SCM_LC_PARAM_PHASE_COUNT              Tag = 0x27400007
	SCM_LC_PARAM_MEASURE_TYPE             Tag = 0x27400008
	SCM_LC_PARAM_CT_RATIO                 Tag = 0x27400009
	SCM_LC_PARAM_METER_ADDRESS            Tag = 0x2740000A
	SCM_LC_PARAM_STATUS                   Tag = 0x2740000B
	SCM_LC_PARAM_MAIN_FUSE_DERATED        Tag = 0x2740000C
	SCM_PARAM_EXT_LOAD_SHEDDING_ACTIVE    Tag = 0x2740000D
	SCM_PARAM_LOAD_SHEDDING_INPUT_ADDRESS Tag = 0x2740000E
	SCM_SET_CONFIG                        Tag = 0x27800001
	SCM_ACCEPTED_CONFIG_FILE_VERSION      Tag = 0x27800002
	SCM_SET_CONFIG_RESET                  Tag = 0x27800006
	SCM_LC_GET_TYPE_MEASURE               Tag = 0x27800007
	SCM_LC_GET_POWER_MGT_CONFIG           Tag = 0x27800008
	SCM_LC_SET_TYPE_MEASURE               Tag = 0x27800009
	SCM_LC_SET_POWER_MGT_CONFIG           Tag = 0x2780000A
)

// ----------------
// NAMESPACE: EEBUS
// 0x28xxxxxx
// undocumented
// ----------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	EEBUS_REQ_LIST_AVAILABLE_DEVICES                                       Tag = 0x28000001
	EEBUS_REQ_ASSOCIATE_DEVICE                                             Tag = 0x28000002
	EEBUS_REQ_LIST_CONNECTED_DEVICES                                       Tag = 0x28000003
	EEBUS_REQ_SET_SEARCH_ACTIVE                                            Tag = 0x28000004
	EEBUS_REQ_ASSOCIATE_DEVICE_TYPE_EVCS                                   Tag = 0x28000005
	EEBUS_REQ_LIST_NON_PAIRED_DEVICE_TYPE_EVCS                             Tag = 0x28000006
	EEBUS_REQ_APPLY_WB_PRECONFIG                                           Tag = 0x28000007
	EEBUS_EMOBILITY_WB_REQ_GET_EVSE_LIST_USECASE                           Tag = 0x28010000
	EEBUS_EMOBILITY_WB_EV_REQ_GET_EV_LIST_USECASE                          Tag = 0x28010001
	EEBUS_EMOBILITY_WB_REQ_GET_EVSE_MANUFACTURER_INFO                      Tag = 0x28010002
	EEBUS_EMOBILITY_WB_EV_REQ_GET_EV_MANUFACTURER_INFO                     Tag = 0x28010003
	EEBUS_EMOBILITY_WB_EV_REQ_GET_CHARGING_POWER_LIMITS                    Tag = 0x28010004
	EEBUS_EMOBILITY_WB_EV_REQ_GET_IDENTIFICATION                           Tag = 0x28010005
	EEBUS_EMOBILITY_WB_EV_REQ_GET_SLEEP_MODE                               Tag = 0x28010006
	EEBUS_EMOBILITY_WB_EV_REQ_GET_ASYMMETRIC_CHARGING_SUPPORTED            Tag = 0x28010007
	EEBUS_EMOBILITY_WB_EV_REQ_GET_COMMUNICATION_STANDARD                   Tag = 0x28010008
	EEBUS_EMOBILITY_WB_EV_REQ_GET_CURRENT_MEASUREMENT_CONSTRAINTS          Tag = 0x28010010
	EEBUS_EMOBILITY_WB_EV_REQ_GET_POWER_MEASUREMENT_CONSTRAINTS            Tag = 0x28010011
	EEBUS_EMOBILITY_WB_EV_REQ_GET_ENERGY_MEASUREMENT_CONSTRAINTS           Tag = 0x28010012
	EEBUS_EMOBILITY_WB_EV_REQ_GET_CHARGING_MEASUREMENT_CONSTRAINTS         Tag = 0x28010013
	EEBUS_EMOBILITY_WB_EV_REQ_GET_CHARGING_CURRENT_MEASUREMENT             Tag = 0x28010014
	EEBUS_EMOBILITY_WB_EV_REQ_GET_CHARGING_POWER_MEASUREMENT               Tag = 0x28010015
	EEBUS_EMOBILITY_WB_EV_REQ_GET_CHARGING_ENERGY_MEASUREMENT              Tag = 0x28010017
	EEBUS_EMOBILITY_WB_EV_REQ_GET_CHARGING_MEASUREMENT                     Tag = 0x28010018
	EEBUS_EMOBILITY_WB_EV_REQ_GET_CHARGING_CURRENT_CONSTRAINTS             Tag = 0x28010020
	EEBUS_EMOBILITY_WB_EV_REQ_GET_OVERLOAD_CHARGING_CURRENT_LIMIT          Tag = 0x28010021
	EEBUS_EMOBILITY_WB_EV_REQ_GET_DATA                                     Tag = 0x2801FFFE
	EEBUS_EMOBILITY_WB_REQ_GET_DATA                                        Tag = 0x2801FFFF
	EEBUS_PARAM_DEVICE                                                     Tag = 0x28400001
	EEBUS_PARAM_SHIP_ID                                                    Tag = 0x28400002
	EEBUS_PARAM_SKI                                                        Tag = 0x28400003
	EEBUS_PARAM_DEVICE_ID                                                  Tag = 0x28400004
	EEBUS_PARAM_EEBUS_ID                                                   Tag = 0x28400005
	EEBUS_PARAM_GENERIC_AC_PHASE                                           Tag = 0x28400006
	EEBUS_PARAM_GENERIC_MANUFACTURER_INFO                                  Tag = 0x28400007
	EEBUS_PARAM_GENERIC_MEASUREMENT_CONSTRAINTS                            Tag = 0x28400008
	EEBUS_PARAM_GENERIC_USECASE_NAME                                       Tag = 0x28400009
	EEBUS_PARAM_GENERIC_USECASE_ENABLED                                    Tag = 0x2840000A
	EEBUS_PARAM_GENERIC_USECASE                                            Tag = 0x2840000B
	EEBUS_PARAM_GENERIC_MANUFACTURER_INFO_DEVICE_NAME                      Tag = 0x2840000C
	EEBUS_PARAM_GENERIC_MANUFACTURER_INFO_DEVICE_CODE                      Tag = 0x2840000D
	EEBUS_PARAM_GENERIC_MANUFACTURER_INFO_SERIAL_NUMBER                    Tag = 0x2840000E
	EEBUS_PARAM_GENERIC_MANUFACTURER_INFO_HARDWARE_REVISION                Tag = 0x2840000F
	EEBUS_PARAM_GENERIC_MANUFACTURER_INFO_SOFTWARE_REVISION                Tag = 0x28400010
	EEBUS_PARAM_GENERIC_MANUFACTURER_INFO_VENDOR_NAME                      Tag = 0x28400011
	EEBUS_PARAM_GENERIC_MANUFACTURER_INFO_VENDOR_CODE                      Tag = 0x28400012
	EEBUS_PARAM_GENERIC_MANUFACTURER_INFO_BRAND_NAME                       Tag = 0x28400013
	EEBUS_PARAM_GENERIC_MANUFACTURER_INFO_MANUFACTURER_LABEL               Tag = 0x28400014
	EEBUS_PARAM_GENERIC_MANUFACTURER_INFO_MANUFACTURER_DESCRIPTION         Tag = 0x28400015
	EEBUS_PARAM_GENERIC_MANUFACTURER_INFO_MANUFACTURER_NODE_IDENTIFICATION Tag = 0x28400016
	EEBUS_PARAM_GENERIC_MEASUREMENT_CONSTRAINTS_MIN                        Tag = 0x28400017
	EEBUS_PARAM_GENERIC_MEASUREMENT_CONSTRAINTS_MAX                        Tag = 0x28400018
	EEBUS_PARAM_GENERIC_MEASUREMENT_CONSTRAINTS_STEP_SIZE                  Tag = 0x28400019
	EEBUS_PARAM_GENERIC_SIGNIFICANT_NUMBER                                 Tag = 0x2840001A
	EEBUS_PARAM_GENERIC_SCALE_NUMBER                                       Tag = 0x2840001B
	EEBUS_PARAM_GENERIC_LOAD_LIMIT_ACTIVATED                               Tag = 0x2840001C
	EEBUS_PARAM_GENERIC_LOAD_LIMIT_VALUE                                   Tag = 0x2840001D
	EEBUS_PARAM_EMOBILITY_EV_ID                                            Tag = 0x28410000
	EEBUS_PARAM_EMOBILITY_EV_IDENTIFICATION_VALUE                          Tag = 0x28410001
	EEBUS_PARAM_EMOBILITY_EV_IDENTIFICATION_TYPE                           Tag = 0x28410002
	EEBUS_PARAM_EMOBILITY_POWER_LIMITS_MIN                                 Tag = 0x28410003
	EEBUS_PARAM_EMOBILITY_POWER_LIMITS_MAX                                 Tag = 0x28410004
	EEBUS_PARAM_EMOBILITY_POWER_LIMITS_STANDBY_POWER                       Tag = 0x28410005
	EEBUS_PARAM_EMOBILITY_CHARGING_MEASUREMENT_CONSTRAINT_PER_PHASE        Tag = 0x28410006
	EEBUS_PARAM_EMOBILITY_CHARGING_MEASUREMENT                             Tag = 0x28410007
	EEBUS_PARAM_EMOBILITY_CHARGING_MEASUREMENT_PER_PHASE                   Tag = 0x28410008
	EEBUS_PARAM_EMOBILITY_CHARGING_CURRENT_CONSTRAINTS_MIN                 Tag = 0x28410009
	EEBUS_PARAM_EMOBILITY_CHARGING_CURRENT_CONSTRAINTS_MAX                 Tag = 0x2841000A
	EEBUS_PARAM_EMOBILITY_CHARGING_CURRENT_CONSTRAINTS_DISCRETE_VALUE      Tag = 0x2841000B
	EEBUS_PARAM_EMOBILITY_CHARGING_CURRENT_CONSTRAINTS_PER_PHASE           Tag = 0x2841000C
	EEBUS_PARAM_EMOBILITY_SLEEP_MODE                                       Tag = 0x2841000D
	EEBUS_PARAM_EMOBILITY_ASYMMETRIC_CHARGING_SUPPORTED                    Tag = 0x2841000E
	EEBUS_PARAM_EMOBILITY_COMMUNICATION_STANDARD                           Tag = 0x2841000F
	EEBUS_PARAM_EMOBILITY_CHARGING_CURRENT_LIMIT_PER_PHASE                 Tag = 0x28410010
	EEBUS_LIST_AVAILABLE_DEVICES                                           Tag = 0x28800001
	EEBUS_LIST_CONNECTED_DEVICES                                           Tag = 0x28800002
	EEBUS_LIST_ASSOCIATED_DEVICE                                           Tag = 0x28800003
	EEBUS_LIST_NON_PAIRED_DEVICE_TYPE_EVCS                                 Tag = 0x28800004
	EEBUS_SET_SEARCH_ACTIVE                                                Tag = 0x28800005
	EEBUS_ASSOCIATE_DEVICE_TYPE_EVCS                                       Tag = 0x28800006
	EEBUS_LIST_USECASE                                                     Tag = 0x28800008
	EEBUS_EMOBILITY_WB_GET_EVSE_LIST_USECASE                               Tag = 0x28810000
	EEBUS_EMOBILITY_WB_EV_GET_EV_LIST_USECASE                              Tag = 0x28810001
	EEBUS_EMOBILITY_WB_GET_EVSE_MANUFACTURER_INFO                          Tag = 0x28810002
	EEBUS_EMOBILITY_WB_EV_GET_EV_MANUFACTURER_INFO                         Tag = 0x28810003
	EEBUS_EMOBILITY_WB_EV_GET_CHARGING_POWER_LIMITS                        Tag = 0x28810004
	EEBUS_EMOBILITY_WB_EV_GET_IDENTIFICATION                               Tag = 0x28810005
	EEBUS_EMOBILITY_WB_EV_GET_SLEEP_MODE                                   Tag = 0x28810006
	EEBUS_EMOBILITY_WB_EV_GET_ASYMMETRIC_CHARGING_SUPPORTED                Tag = 0x28810007
	EEBUS_EMOBILITY_WB_EV_GET_COMMUNICATION_STANDARD                       Tag = 0x28810008
	EEBUS_EMOBILITY_WB_EV_GET_CURRENT_MEASUREMENT_CONSTRAINTS              Tag = 0x28810010
	EEBUS_EMOBILITY_WB_EV_GET_POWER_MEASUREMENT_CONSTRAINTS                Tag = 0x28810011
	EEBUS_EMOBILITY_WB_EV_GET_ENERGY_MEASUREMENT_CONSTRAINTS               Tag = 0x28810012
	EEBUS_EMOBILITY_WB_EV_GET_CHARGING_MEASUREMENT_CONSTRAINTS             Tag = 0x28810013
	EEBUS_EMOBILITY_WB_EV_GET_CHARGING_CURRENT_MEASUREMENT                 Tag = 0x28810014
	EEBUS_EMOBILITY_WB_EV_GET_CHARGING_POWER_MEASUREMENT                   Tag = 0x28810015
	EEBUS_EMOBILITY_WB_EV_GET_CHARGING_ENERGY_MEASUREMENT                  Tag = 0x28810017
	EEBUS_EMOBILITY_WB_EV_GET_CHARGING_MEASUREMENT                         Tag = 0x28810018
	EEBUS_EMOBILITY_WB_EV_GET_CHARGING_CURRENT_CONSTRAINTS                 Tag = 0x28810020
	EEBUS_EMOBILITY_WB_EV_GET_OVERLOAD_CHARGING_CURRENT_LIMIT              Tag = 0x28810021
	EEBUS_EMOBILITY_WB_EV_GET_DATA                                         Tag = 0x2881FFFE
	EEBUS_EMOBILITY_WB_GET_DATA                                            Tag = 0x2881FFFF
)

// ---------------
// NAMESPACE: SDSA
// 0x29xxxxxx
// undocumented
// ---------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	SDSA_REQ_STATUS_FILE Tag = 0x29000001
	SDSA_STATUS_FILE     Tag = 0x29800001
)

// --------------
// NAMESPACE: ETH
// 0x2Axxxxxx
// undocumented
// --------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	ETH_REQ_SET_SEARCH_ACTIVE               Tag = 0x2A000001
	ETH_REQ_LIST_NON_CONNECTED_DEVICES      Tag = 0x2A000002
	ETH_REQ_CONNECT_TO_DEVICE               Tag = 0x2A000003
	ETH_REQ_IS_SCANNING                     Tag = 0x2A000004
	ETH_REQ_DISCONNECT_FROM_DEVICE          Tag = 0x2A000005
	ETH_REQ_LIST_CONNECTED_DEVICES          Tag = 0x2A000006
	ETH_REQ_APPLY_WB_PRECONFIG              Tag = 0x2A000007
	ETH_PARAM_SEARCH_ACTIVE                 Tag = 0x2A400001
	ETH_PARAM_AUTO_ACCEPT                   Tag = 0x2A400002
	ETH_PARAM_EXTERN_CONFIGURATOR_DEVICE_ID Tag = 0x2A400003
	ETH_PARAM_AUTO_ACCEPT_VALUES            Tag = 0x2A400004
	ETH_PARAM_DEVICE                        Tag = 0x2A400005
	ETH_PARAM_DEVICE_NAME                   Tag = 0x2A400006
	ETH_PARAM_SERIAL_NUMBER                 Tag = 0x2A400007
	ETH_PARAM_LOCAL_ID                      Tag = 0x2A400008
	ETH_PARAM_MAC_ADDRESS                   Tag = 0x2A400009
	ETH_PARAM_IP_ADDRESS                    Tag = 0x2A40000A
	ETH_PARAM_TIMEOUT                       Tag = 0x2A40000B
	ETH_PARAM_DISCONNECTED                  Tag = 0x2A40000C
	ETH_PARAM_CONNECTION_SUCCESS            Tag = 0x2A40000D
	ETH_PARAM_DEVICE_ID                     Tag = 0x2A40000E
	ETH_PARAM_NUMBER_OF_DEVICES             Tag = 0x2A40000F
	ETH_SET_SEARCH_ACTIVE                   Tag = 0x2A800001
	ETH_LIST_NON_CONNECTED_DEVICES          Tag = 0x2A800002
	ETH_CONNECT_TO_DEVICE                   Tag = 0x2A800003
	ETH_IS_SCANNING                         Tag = 0x2A800004
	ETH_DISCONNECT_FROM_DEVICE              Tag = 0x2A800005
	ETH_LIST_CONNECTED_DEVICES              Tag = 0x2A800006
)

// --------------
// NAMESPACE: LCT
// 0x2Bxxxxxx
// undocumented
// --------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	LCT_PARAM_DURATION           Tag = 0x2B000000
	LCT_REQ_START_NGROK          Tag = 0x2B000001
	LCT_REQ_GET_IS_NGROK_RUNNING Tag = 0x2B000002
	LCT_REQ_GET_IS_GTOU_ACCEPTED Tag = 0x2B000003
	LCT_REQ_SET_GTOU_ACCEPTED    Tag = 0x2B000004
	LCT_SD_REQ_GET_MEM_INFO      Tag = 0x2B000005
	LCT_SD_REQ_GET_MOF_HEAP      Tag = 0x2B000006
	LCT_SD_REQ_GET_UP_TIME       Tag = 0x2B000007
	LCT_PARAM_STATUS             Tag = 0x2B400001
	LCT_PARAM_NGROK_HOSTNAME     Tag = 0x2B400002
	LCT_PARAM_GTOU_ACCEPTED      Tag = 0x2B400003
	LCT_PARAM_TOKEN              Tag = 0x2B400004
	LCT_SD_PARAM_MEM_FREE        Tag = 0x2B400005
	LCT_SD_PARAM_MEM_CACHED      Tag = 0x2B400006
	LCT_SD_PARAM_MEM_WD_FREE     Tag = 0x2B400007
	LCT_SD_PARAM_MEM_BUFFER      Tag = 0x2B400008
	LCT_SD_PARAM_SYNC_TIME       Tag = 0x2B400009
	LCT_START_NGROK              Tag = 0x2B800001
	LCT_GET_IS_NGROK_RUNNING     Tag = 0x2B800002
	LCT_GET_IS_GTOU_ACCEPTED     Tag = 0x2B800003
	LCT_SET_GTOU_ACCEPTED        Tag = 0x2B800004
	LCT_SD_GET_MEM_INFO          Tag = 0x2B800005
	LCT_SD_GET_MOF_HEAP          Tag = 0x2B800006
	LCT_SD_GET_UP_TIME           Tag = 0x2B800007
)

// ----------------------------
// NAMESPACE: HG_OCPP_WB_CFG_LC
// 0x2Cxxxxxx
// undocumented
// ----------------------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	HG_OCPP_WB_CFG_LC_REQ_SET_START_SCAN              Tag = 0x2C000001
	HG_OCPP_WB_CFG_LC_REQ_GET_NON_CONFIGURED_HAG_EVCS Tag = 0x2C000002
	HG_OCPP_WB_CFG_LC_REQ_SET_CONFIGURE_HAG_EVCS      Tag = 0x2C000003
	HG_OCPP_WB_CFG_LC_REQ_SET_CONFIGURE_ALL_HAG_EVCS  Tag = 0x2C000004
	HG_OCPP_WB_CFG_LC_REQ_SET_START_AUTO_CONFIG       Tag = 0x2C000005
	HG_OCPP_WB_CFG_LC_REQ_GET_LLM_URI                 Tag = 0x2C000006
	HG_OCPP_WB_CFG_LC_PARAM_TOKEN                     Tag = 0x2C400001
	HG_OCPP_WB_CFG_LC_PARAM_STATUS                    Tag = 0x2C400002
	HG_OCPP_WB_CFG_LC_PARAM_EVCS_DEVICE               Tag = 0x2C400003
	HG_OCPP_WB_CFG_LC_PARAM_EVCS_DEVICE_HOSTNAME      Tag = 0x2C400004
	HG_OCPP_WB_CFG_LC_PARAM_EVCS_DEVICE_IP            Tag = 0x2C400005
	HG_OCPP_WB_CFG_LC_PARAM_EVCS_DEVICE_STATUS        Tag = 0x2C400006
	HG_OCPP_WB_CFG_LC_PARAM_RUN                       Tag = 0x2C400007
	HG_OCPP_WB_CFG_LC_PARAM_MAX_ACCEPTED              Tag = 0x2C400008
	HG_OCPP_WB_CFG_LC_PARAM_ACCEPTED                  Tag = 0x2C400009
	HG_OCPP_WB_CFG_LC_PARAM_ACCEPTED_HAGER            Tag = 0x2C40000A
	HG_OCPP_WB_CFG_LC_PARAM_UNDERWAY                  Tag = 0x2C40000B
	HG_OCPP_WB_CFG_LC_PARAM_LLM_URI                   Tag = 0x2C40000C
	HG_OCPP_WB_CFG_LC_SET_START_SCAN                  Tag = 0x2C800001
	HG_OCPP_WB_CFG_LC_GET_NON_CONFIGURED_HAG_EVCS     Tag = 0x2C800002
	HG_OCPP_WB_CFG_LC_SET_CONFIGURE_HAG_EVCS          Tag = 0x2C800003
	HG_OCPP_WB_CFG_LC_SET_CONFIGURE_ALL_HAG_EVCS      Tag = 0x2C800004
	HG_OCPP_WB_CFG_LC_SET_START_AUTO_CONFIG           Tag = 0x2C800005
	HG_OCPP_WB_CFG_LC_GET_LLM_URI                     Tag = 0x2C800006
)

// ---------------------
// NAMESPACE: OCPP_WB_LC
// 0x2Dxxxxxx
// undocumented
// ---------------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	OCPP_WB_LC_REQ_SET_ACCEPTANCE_RULE                              Tag = 0x2D000001
	OCPP_WB_LC_REQ_GET_NON_ACCEPTED_EVCS                            Tag = 0x2D000002
	OCPP_WB_LC_REQ_GET_ACCEPTED_EVCS                                Tag = 0x2D000003
	OCPP_WB_LC_REQ_SET_ACCEPT_EVCS                                  Tag = 0x2D000004
	OCPP_WB_LC_REQ_SET_ACCEPT_ALL_EVCS                              Tag = 0x2D000005
	OCPP_WB_LC_REQ_SET_BADGE_STRATEGY                               Tag = 0x2D000006
	OCPP_WB_LC_REQ_SET_BADGE_ACQUISITION_MODE_START                 Tag = 0x2D000007
	OCPP_WB_LC_REQ_SET_BADGE_ACQUISITION_MODE_STOP                  Tag = 0x2D000008
	OCPP_WB_LC_REQ_GET_ACCEPTANCE_RULE                              Tag = 0x2D000009
	OCPP_WB_LC_REQ_GET_BADGE_STRATEGY                               Tag = 0x2D00000A
	OCPP_WB_LC_REQ_GET_ACCEPTED_EVCS_EXTENDED                       Tag = 0x2D00000B
	OCPP_WB_LC_REQ_GET_EVCS_AVAILABLE_PARAMETERS_PROFILES           Tag = 0x2D00000C
	OCPP_WB_LC_REQ_GET_EVCS_DEVICE_ADVANCED_SETTINGS                Tag = 0x2D00000D
	OCPP_WB_LC_REQ_SET_EVCS_DEVICES_PARAMETERS                      Tag = 0x2D00000E
	OCPP_WB_LC_REQ_SET_REMOVE_WB                                    Tag = 0x2D00000F
	OCPP_WB_LC_REQ_SET_FACTORY_RESET                                Tag = 0x2D000010
	OCPP_WB_LC_REQ_GET_WALLBOXES_COUNT                              Tag = 0x2D000011
	OCPP_WB_LC_REQ_GET_MAX_WALLBOXES_COUNT                          Tag = 0x2D000012
	OCPP_WB_LC_REQ_GET_AVAILABLE_CPO_LIST                           Tag = 0x2D000013
	OCPP_WB_LC_REQ_SET_CPO_BUTTON_STATE                             Tag = 0x2D000014
	OCPP_WB_LC_REQ_GET_CPO_BUTTON_STATE                             Tag = 0x2D000015
	OCPP_WB_LC_REQ_SET_CLEAN_CPO_SETTINGS                           Tag = 0x2D000016
	OCPP_WB_LC_REQ_SET_CPO_MODE                                     Tag = 0x2D000017
	OCPP_WB_LC_REQ_SET_CPO_SETTINGS                                 Tag = 0x2D000018
	OCPP_WB_LC_REQ_GET_CPO_SETTINGS                                 Tag = 0x2D000019
	OCPP_WB_LC_REQ_GET_EVCS_DEVICE_CPO_CONNECTABLE                  Tag = 0x2D00001A
	OCPP_WB_LC_REQ_SET_EVCS_DEVICE_CPO_CONNECTABLE                  Tag = 0x2D00001B
	OCPP_WB_LC_REQ_GET_CPO_MODE                                     Tag = 0x2D00001C
	OCPP_WB_LC_REQ_GET_EVCS_CHARGING_SESSION_HISTORY_NB             Tag = 0x2D00001D
	OCPP_WB_LC_REQ_GET_EVCS_CHARGING_SESSION_HISTORY                Tag = 0x2D00001E
	OCPP_WB_LC_REQ_GET_EVCS_CHARGING_SESSIONS_ACTIVE                Tag = 0x2D00001F
	OCPP_WB_LC_REQ_GET_BADGES_PER_EVCS                              Tag = 0x2D000020
	OCPP_WB_LC_REQ_SET_BADGES_PER_EVCS                              Tag = 0x2D000021
	OCPP_WB_LC_REQ_GET_EVCS_LIST_PER_BADGE                          Tag = 0x2D000022
	OCPP_WB_LC_REQ_SET_EVCS_LIST_PER_BADGE                          Tag = 0x2D000023
	OCPP_WB_LC_PARAM_TOKEN                                          Tag = 0x2D400001
	OCPP_WB_LC_PARAM_ACCEPTANCE_RULE_VAL                            Tag = 0x2D400002
	OCPP_WB_LC_PARAM_STATUS                                         Tag = 0x2D400003
	OCPP_WB_LC_PARAM_EVCS_DEVICE                                    Tag = 0x2D400004
	OCPP_WB_LC_PARAM_EVCS_DEVICE_HOSTNAME                           Tag = 0x2D400005
	OCPP_WB_LC_PARAM_EVCS_DEVICE_IP                                 Tag = 0x2D400006
	OCPP_WB_LC_PARAM_EVCS_DEVICE_OCPP_ID                            Tag = 0x2D400007
	OCPP_WB_LC_PARAM_EVCS_DEVICE_STATUS                             Tag = 0x2D400008
	OCPP_WB_LC_PARAM_BADGE_STRATEGY                                 Tag = 0x2D400009
	OCPP_WB_LC_PARAM_EVCS_DEVICE_CONNECTOR                          Tag = 0x2D40000A
	OCPP_WB_LC_PARAM_EVCS_DEVICE_CONNECTOR_ID                       Tag = 0x2D40000B
	OCPP_WB_LC_PARAM_EVCS_DEVICE_CONNECTOR_PLUG_TYPE                Tag = 0x2D40000C
	OCPP_WB_LC_PARAM_EVCS_DEVICE_CONNECTOR_MAX_CURRENT_PER_PHASE    Tag = 0x2D40000D
	OCPP_WB_LC_PARAM_EVCS_DEVICE_CONNECTOR_MIN_CURRENT_PER_PHASE    Tag = 0x2D40000E
	OCPP_WB_LC_PARAM_EVCS_DEVICE_CONNECTOR_AUTH_OFFLINE_CHARGE_SESS Tag = 0x2D40000F
	OCPP_WB_LC_PARAM_EVCS_DEVICE_OCPP_SETTINGS                      Tag = 0x2D400010
	OCPP_WB_LC_PARAM_EVCS_DEVICE_OCPP_ACCEPT_UNKNOWN_OFFLINE_AUTH   Tag = 0x2D400011
	OCPP_WB_LC_PARAM_EVCS_DEVICE_OCPP_AUTHORIZATION_CACHE_ENABLED   Tag = 0x2D400012
	OCPP_WB_LC_PARAM_EVCS_DEVICE_OCPP_REMOTE_TRANSACTION_ENABLED    Tag = 0x2D400013
	OCPP_WB_LC_PARAM_EVCS_DEVICE_OCPP_BLINKS_NUMBER                 Tag = 0x2D400014
	OCPP_WB_LC_PARAM_EVCS_DEVICE_OCPP_CLOCK_ALIGNED_DATA_INTERVAL   Tag = 0x2D400015
	OCPP_WB_LC_PARAM_EVCS_DEVICE_OCPP_CONNECTION_TIMEOUT            Tag = 0x2D400016
	OCPP_WB_LC_PARAM_EVCS_DEVICE_OCPP_CONNECTOR_PHASE_ROTATION      Tag = 0x2D400017
	OCPP_WB_LC_PARAM_EVCS_DEVICE_OCPP_HEARTBEAT_INTERVAL            Tag = 0x2D400018
	OCPP_WB_LC_PARAM_EVCS_DEVICE_OCPP_LIGHT_INTENSITY               Tag = 0x2D400019
	OCPP_WB_LC_PARAM_EVCS_DEVICE_OCPP_LOCAL_AUTHORIZE_OFFLINE       Tag = 0x2D40001A
	OCPP_WB_LC_PARAM_EVCS_DEVICE_OCPP_LOCAL_PRE_AUTHORIZE           Tag = 0x2D40001B
	OCPP_WB_LC_PARAM_EVCS_DEVICE_OCPP_MAX_ENERGY_ON_INVALID_ID      Tag = 0x2D40001C
	OCPP_WB_LC_PARAM_EVCS_DEVICE_PHASE_MAPPING                      Tag = 0x2D40001D
	OCPP_WB_LC_PARAM_EVCS_DEVICE_CLUSTER_NAME                       Tag = 0x2D40001E
	OCPP_WB_LC_PARAM_EVCS_DEVICE_NAME                               Tag = 0x2D40001F
	OCPP_WB_LC_PARAM_EVCS_DEVICE_CONFIG_STATUS                      Tag = 0x2D400020
	OCPP_WB_LC_PARAM_EVCS_PARAMETERS_PROFILE                        Tag = 0x2D400021
	OCPP_WB_LC_PARAM_EVCS_DEVICE_DEFAULT_CHARGING_PROFILE           Tag = 0x2D400022
	OCPP_WB_LC_PARAM_EVCS_DEVICE_PLUG_TYPE_PER_CONNECTOR            Tag = 0x2D400023
	OCPP_WB_LC_PARAM_EVCS_DEVICE_ID                                 Tag = 0x2D400024
	OCPP_WB_LC_PARAM_EVCS_DEVICE_STATE                              Tag = 0x2D400025
	OCPP_WB_LC_PARAM_WB_COUNT                                       Tag = 0x2D400026
	OCPP_WB_LC_PARAM_EVCS_DEVICE_CONNECTOR_COUNT                    Tag = 0x2D400027
	OCPP_WB_LC_PARAM_CPO_INFO                                       Tag = 0x2D400028
	OCPP_WB_LC_PARAM_CPO_NAME                                       Tag = 0x2D400029
	OCPP_WB_LC_PARAM_CPO_URI                                        Tag = 0x2D40002A
	OCPP_WB_LC_PARAM_CPO_BUTTON                                     Tag = 0x2D40002B
	OCPP_WB_LC_PARAM_CPO_MODE                                       Tag = 0x2D40002C
	OCPP_WB_LC_PARAM_SETTINGS                                       Tag = 0x2D40002D
	OCPP_WB_LC_PARAM_CREDENTIALS                                    Tag = 0x2D40002E
	OCPP_WB_LC_PARAM_LOGIN                                          Tag = 0x2D40002F
	OCPP_WB_LC_PARAM_PASSWORD                                       Tag = 0x2D400030
	OCPP_WB_LC_PARAM_AUTH_TYPE                                      Tag = 0x2D400031
	OCPP_WB_LC_PARAM_CPO_ADDRESS                                    Tag = 0x2D400032
	OCPP_WB_LC_PARAM_CPO_PORT                                       Tag = 0x2D400033
	OCPP_WB_LC_PARAM_CPO_PATH                                       Tag = 0x2D400034
	OCPP_WB_LC_PARAM_CPO_UUID                                       Tag = 0x2D400035
	OCPP_WB_LC_PARAM_EVCS_CHARGING_SESSION_START_DATE               Tag = 0x2D400036
	OCPP_WB_LC_PARAM_EVCS_CHARGING_SESSION_STOP_DATE                Tag = 0x2D400037
	OCPP_WB_LC_PARAM_EVCS_CHARGING_SESSION_HISTORY_NB               Tag = 0x2D400038
	OCPP_WB_LC_PARAM_EVCS_CHARGING_SESSION_HISTORY                  Tag = 0x2D400039
	OCPP_WB_LC_PARAM_EVCS_BADGE_ID                                  Tag = 0x2D40003A
	OCPP_WB_LC_PARAM_EVCS_BADGE_COMMENT                             Tag = 0x2D40003B
	OCPP_WB_LC_PARAM_EVCS_ENERGY_CHARGED                            Tag = 0x2D40003C
	OCPP_WB_LC_PARAM_EVCS_END_DATE                                  Tag = 0x2D40003D
	OCPP_WB_LC_PARAM_EVCS_TRANSACTION_ID                            Tag = 0x2D40003E
	OCPP_WB_LC_PARAM_IS_EVCS_CHARGING_SESSIONS_ACTIVE               Tag = 0x2D40003F
	OCPP_WB_LC_PARAM_EVCS_DEVICE_CPO_CONFIG_STATUS                  Tag = 0x2D400041
	OCPP_WB_LC_PARAM_EVCS_DEVICE_ASSOCIATED                         Tag = 0x2D400042
	OCPP_WB_LC_PARAM_EVCS_DEVICE_BADGE                              Tag = 0x2D400044
	OCPP_WB_LC_PARAM_EVCS_BADGE_ASSOCIATED                          Tag = 0x2D400045
	OCPP_WB_LC_PARAM_EVCS_BADGE_MASK                                Tag = 0x2D400048
	OCPP_WB_LC_PARAM_EVCS_BADGE_MASK_INDEX                          Tag = 0x2D400049
	OCPP_WB_LC_SET_ACCEPTANCE_RULE                                  Tag = 0x2D800001
	OCPP_WB_LC_GET_NON_ACCEPTED_EVCS                                Tag = 0x2D800002
	OCPP_WB_LC_GET_ACCEPTED_EVCS                                    Tag = 0x2D800003
	OCPP_WB_LC_SET_ACCEPT_EVCS                                      Tag = 0x2D800004
	OCPP_WB_LC_SET_ACCEPT_ALL_EVCS                                  Tag = 0x2D800005
	OCPP_WB_LC_SET_BADGE_STRATEGY                                   Tag = 0x2D800006
	OCPP_WB_LC_SET_BADGE_ACQUISITION_MODE_START                     Tag = 0x2D800007
	OCPP_WB_LC_SET_BADGE_ACQUISITION_MODE_STOP                      Tag = 0x2D800008
	OCPP_WB_LC_GET_ACCEPTANCE_RULE                                  Tag = 0x2D800009
	OCPP_WB_LC_GET_BADGE_STRATEGY                                   Tag = 0x2D80000A
	OCPP_WB_LC_GET_ACCEPTED_EVCS_EXTENDED                           Tag = 0x2D80000B
	OCPP_WB_LC_GET_EVCS_AVAILABLE_PARAMETERS_PROFILES               Tag = 0x2D80000C
	OCPP_WB_LC_GET_EVCS_DEVICE_ADVANCED_SETTINGS                    Tag = 0x2D80000D
	OCPP_WB_LC_SET_EVCS_DEVICES_PARAMETERS                          Tag = 0x2D80000E
	OCPP_WB_LC_SET_REMOVE_WB                                        Tag = 0x2D80000F
	OCPP_WB_LC_SET_FACTORY_RESET                                    Tag = 0x2D800010
	OCPP_WB_LC_GET_WALLBOXES_COUNT                                  Tag = 0x2D800011
	OCPP_WB_LC_GET_MAX_WALLBOXES_COUNT                              Tag = 0x2D800012
	OCPP_WB_LC_GET_AVAILABLE_CPO_LIST                               Tag = 0x2D800013
	OCPP_WB_LC_SET_CPO_BUTTON_STATE                                 Tag = 0x2D800014
	OCPP_WB_LC_GET_CPO_BUTTON_STATE                                 Tag = 0x2D800015
	OCPP_WB_LC_SET_CLEAN_CPO_SETTINGS                               Tag = 0x2D800016
	OCPP_WB_LC_SET_CPO_MODE                                         Tag = 0x2D800017
	OCPP_WB_LC_SET_CPO_SETTINGS                                     Tag = 0x2D800018
	OCPP_WB_LC_GET_CPO_SETTINGS                                     Tag = 0x2D800019
	OCPP_WB_LC_GET_EVCS_DEVICE_CPO_CONNECTABLE                      Tag = 0x2D80001A
	OCPP_WB_LC_SET_EVCS_DEVICE_CPO_CONNECTABLE                      Tag = 0x2D80001B
	OCPP_WB_LC_GET_CPO_MODE                                         Tag = 0x2D80001C
	OCPP_WB_LC_GET_EVCS_CHARGING_SESSION_HISTORY_NB                 Tag = 0x2D80001D
	OCPP_WB_LC_GET_EVCS_CHARGING_SESSION_HISTORY                    Tag = 0x2D80001E
	OCPP_WB_LC_GET_EVCS_CHARGING_SESSIONS_ACTIVE                    Tag = 0x2D80001F
	OCPP_WB_LC_GET_BADGES_PER_EVCS                                  Tag = 0x2D800020
	OCPP_WB_LC_SET_BADGES_PER_EVCS                                  Tag = 0x2D800021
	OCPP_WB_LC_GET_EVCS_LIST_PER_BADGE                              Tag = 0x2D800022
	OCPP_WB_LC_SET_EVCS_LIST_PER_BADGE                              Tag = 0x2D800023
)

// --------------------------
// NAMESPACE: WB_BADGE_MGT_LC
// 0x2Exxxxxx
// undocumented
// --------------------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	WB_BADGE_MGT_LC_REQ_SET_ADD_BADGE            Tag = 0x2E000001
	WB_BADGE_MGT_LC_REQ_SET_REMOVE_BADGE         Tag = 0x2E000002
	WB_BADGE_MGT_LC_REQ_GET_BADGE                Tag = 0x2E000003
	WB_BADGE_MGT_LC_REQ_SET_BADGE_PARAM          Tag = 0x2E000004
	WB_BADGE_MGT_LC_REQ_GET_BADGES_COUNT         Tag = 0x2E000005
	WB_BADGE_MGT_LC_PARAM_TOKEN                  Tag = 0x2E400001
	WB_BADGE_MGT_LC_PARAM_BADGE                  Tag = 0x2E400002
	WB_BADGE_MGT_LC_PARAM_BADGE_ID               Tag = 0x2E400003
	WB_BADGE_MGT_LC_PARAM_BADGE_ROLE             Tag = 0x2E400004
	WB_BADGE_MGT_LC_PARAM_BADGE_EXP_DATE         Tag = 0x2E400005
	WB_BADGE_MGT_LC_PARAM_BADGE_COMMENT          Tag = 0x2E400006
	WB_BADGE_MGT_LC_PARAM_BADGE_BLOCKED          Tag = 0x2E400007
	WB_BADGE_MGT_LC_PARAM_BADGE_MAIL             Tag = 0x2E400008
	WB_BADGE_MGT_LC_PARAM_STATUS                 Tag = 0x2E400009
	WB_BADGE_MGT_LC_PARAM_FILTER_BLOCKED         Tag = 0x2E40000A
	WB_BADGE_MGT_LC_PARAM_FILTER_NON_BLOCKED     Tag = 0x2E40000B
	WB_BADGE_MGT_LC_PARAM_FILTER_ROLE_EQ         Tag = 0x2E40000C
	WB_BADGE_MGT_LC_PARAM_FILTER_ROLE_NEQ        Tag = 0x2E40000D
	WB_BADGE_MGT_LC_PARAM_FILTER_BADGE_ID        Tag = 0x2E40000E
	WB_BADGE_MGT_LC_PARAM_FILTER_COMMENT         Tag = 0x2E40000F
	WB_BADGE_MGT_LC_PARAM_FILTER_MAIL            Tag = 0x2E400010
	WB_BADGE_MGT_LC_PARAM_FILTER_EXP_DATE_BEFORE Tag = 0x2E400011
	WB_BADGE_MGT_LC_PARAM_FILTER_EXP_DATE_AFTER  Tag = 0x2E400012
	WB_BADGE_MGT_LC_PARAM_BADGES_COUNT           Tag = 0x2E400013
	WB_BADGE_MGT_LC_PARAM_BADGE_MASK_INDEX       Tag = 0x2E400014
	WB_BADGE_MGT_LC_SET_ADD_BADGE                Tag = 0x2E800001
	WB_BADGE_MGT_LC_SET_REMOVE_BADGE             Tag = 0x2E800002
	WB_BADGE_MGT_LC_GET_BADGE                    Tag = 0x2E800003
	WB_BADGE_MGT_LC_SET_BADGE_PARAM              Tag = 0x2E800004
	WB_BADGE_MGT_LC_GET_BADGES_COUNT             Tag = 0x2E800005
)

// ------------------------
// NAMESPACE: LC_USR_MGT_LC
// 0x2Fxxxxxx
// undocumented
// ------------------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	LC_USR_MGT_LC_REQ_GET_LOGIN                 Tag = 0x2F000001
	LC_USR_MGT_LC_REQ_GET_PROFILE_BY_TOKEN      Tag = 0x2F000002
	LC_USR_MGT_LC_REQ_SET_ADD_USER_OR_CHG_PWD   Tag = 0x2F000003
	LC_USR_MGT_LC_REQ_SET_UPDATE_USER_ACCOUNT   Tag = 0x2F000004
	LC_USR_MGT_LC_REQ_SET_ENABLE_PUBLIC_ACCESS  Tag = 0x2F000005
	LC_USR_MGT_LC_REQ_SET_DISABLE_PUBLIC_ACCESS Tag = 0x2F000006
	LC_USR_MGT_LC_REQ_SET_REMOVE_USER           Tag = 0x2F000007
	LC_USR_MGT_LC_REQ_GET_PUBLIC_ACCESS_STATE   Tag = 0x2F000008
	LC_USR_MGT_LC_REQ_SET_DISCONNECT_USER       Tag = 0x2F000009
	LC_USR_MGT_LC_REQ_GET_USERS_LIST            Tag = 0x2F00000A
	LC_USR_MGT_LC_REQ_GET_USER_RELATED_INFO     Tag = 0x2F00000B
	LC_USR_MGT_LC_REQ_GET_LOGIN_PUBLIC_USER     Tag = 0x2F00000C
	LC_USR_MGT_LC_REQ_GET_USERS_COUNT           Tag = 0x2F00000D
	LC_USR_MGT_LC_PARAM_TOKEN                   Tag = 0x2F400001
	LC_USR_MGT_LC_PARAM_USER                    Tag = 0x2F400002
	LC_USR_MGT_LC_PARAM_LOGIN                   Tag = 0x2F400003
	LC_USR_MGT_LC_PARAM_PWD                     Tag = 0x2F400004
	LC_USR_MGT_LC_PARAM_STATUS                  Tag = 0x2F400005
	LC_USR_MGT_LC_PARAM_ROLE                    Tag = 0x2F400006
	LC_USR_MGT_LC_PARAM_NAME                    Tag = 0x2F400007
	LC_USR_MGT_LC_PARAM_FIRST_NAME              Tag = 0x2F400008
	LC_USR_MGT_LC_PARAM_STREET                  Tag = 0x2F400009
	LC_USR_MGT_LC_PARAM_ZIP                     Tag = 0x2F40000A
	LC_USR_MGT_LC_PARAM_CITY                    Tag = 0x2F40000B
	LC_USR_MGT_LC_PARAM_COUNTRY                 Tag = 0x2F40000C
	LC_USR_MGT_LC_PARAM_EMAIL                   Tag = 0x2F40000D
	LC_USR_MGT_LC_PARAM_PHONE                   Tag = 0x2F40000E
	LC_USR_MGT_LC_PARAM_EMAIL_DISPLAYED         Tag = 0x2F40000F
	LC_USR_MGT_LC_PARAM_PHONE_DISPLAYED         Tag = 0x2F400010
	LC_USR_MGT_LC_PARAM_PUBLIC_ACCESS_STATE     Tag = 0x2F400011
	LC_USR_MGT_LC_PARAM_USERS_COUNT             Tag = 0x2F400012
	LC_USR_MGT_LC_GET_LOGIN                     Tag = 0x2F800001
	LC_USR_MGT_LC_GET_PROFILE_BY_TOKEN          Tag = 0x2F800002
	LC_USR_MGT_LC_SET_ADD_USER_OR_CHG_PWD       Tag = 0x2F800003
	LC_USR_MGT_LC_SET_UPDATE_USER_ACCOUNT       Tag = 0x2F800004
	LC_USR_MGT_LC_SET_ENABLE_PUBLIC_ACCESS      Tag = 0x2F800005
	LC_USR_MGT_LC_SET_DISABLE_PUBLIC_ACCESS     Tag = 0x2F800006
	LC_USR_MGT_LC_SET_REMOVE_USER               Tag = 0x2F800007
	LC_USR_MGT_LC_GET_PUBLIC_ACCESS_STATE       Tag = 0x2F800008
	LC_USR_MGT_LC_SET_DISCONNECT_USER           Tag = 0x2F800009
	LC_USR_MGT_LC_GET_USERS_LIST                Tag = 0x2F80000A
	LC_USR_MGT_LC_GET_USER_RELATED_INFO         Tag = 0x2F80000B
	LC_USR_MGT_LC_GET_LOGIN_PUBLIC_USER         Tag = 0x2F80000C
	LC_USR_MGT_LC_GET_USERS_COUNT               Tag = 0x2F80000D
)

// -----------------------
// NAMESPACE: DASHBOARD_LC
// 0x30xxxxxx
// undocumented
// -----------------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	DASHBOARD_LC_REQ_GET_SYSTEM_INFO                   Tag = 0x30000001
	DASHBOARD_LC_REQ_GET_CONSUMPTIONS                  Tag = 0x30000002
	DASHBOARD_LC_REQ_GET_CHARGING_SESSIONS             Tag = 0x30000003
	DASHBOARD_LC_REQ_GET_SUPPORT_CONTACTS              Tag = 0x30000004
	DASHBOARD_LC_PARAM_TOKEN                           Tag = 0x30400001
	DASHBOARD_LC_PARAM_POWER_MGT_TYPE                  Tag = 0x30400002
	DASHBOARD_LC_PARAM_MAX_CURRENT_PER_PHASE           Tag = 0x30400003
	DASHBOARD_LC_PARAM_WALLBOXES_COUNT                 Tag = 0x30400004
	DASHBOARD_LC_PARAM_BADGE_STRATEGY                  Tag = 0x30400005
	DASHBOARD_LC_PARAM_BADGES_COUNT                    Tag = 0x30400006
	DASHBOARD_LC_PARAM_STATUS                          Tag = 0x30400007
	DASHBOARD_LC_PARAM_CONSUMPTION_TOTAL_EVSE          Tag = 0x30400008
	DASHBOARD_LC_PARAM_CONSUMPTION_TOTAL_EVSE_PHASE_L1 Tag = 0x30400009
	DASHBOARD_LC_PARAM_CONSUMPTION_TOTAL_EVSE_PHASE_L2 Tag = 0x3040000A
	DASHBOARD_LC_PARAM_CONSUMPTION_TOTAL_EVSE_PHASE_L3 Tag = 0x3040000B
	DASHBOARD_LC_PARAM_CONSUMPTION_OVERALL             Tag = 0x3040000C
	DASHBOARD_LC_PARAM_CONSUMPTION_OVERALL_PHASE_L1    Tag = 0x3040000D
	DASHBOARD_LC_PARAM_CONSUMPTION_OVERALL_PHASE_L2    Tag = 0x3040000E
	DASHBOARD_LC_PARAM_CONSUMPTION_OVERALL_PHASE_L3    Tag = 0x3040000F
	DASHBOARD_LC_PARAM_WB_SESSION                      Tag = 0x30400010
	DASHBOARD_LC_PARAM_WB_LABEL                        Tag = 0x30400011
	DASHBOARD_LC_PARAM_WB_HOSTNAME                     Tag = 0x30400012
	DASHBOARD_LC_PARAM_WB_MAC_ADDRESS                  Tag = 0x30400013
	DASHBOARD_LC_PARAM_WB_STATE                        Tag = 0x30400014
	DASHBOARD_LC_PARAM_WB_SESSION_ACTIVE_SEC           Tag = 0x30400015
	DASHBOARD_LC_PARAM_WB_SESSION_INACTIVE_SEC         Tag = 0x30400016
	DASHBOARD_LC_PARAM_WB_SESSION_BADGE_ID             Tag = 0x30400017
	DASHBOARD_LC_PARAM_WB_CONNECTOR                    Tag = 0x30400018
	DASHBOARD_LC_PARAM_WB_CONNECTOR_ID                 Tag = 0x30400019
	DASHBOARD_LC_PARAM_WB_CONNECTOR_TYPE               Tag = 0x3040001A
	DASHBOARD_LC_PARAM_WB_CONNECTOR_STATUS             Tag = 0x3040001B
	DASHBOARD_LC_PARAM_WB_CONNECTOR_ERROR_MSG          Tag = 0x3040001C
	DASHBOARD_LC_PARAM_WB_CONNECTOR_PHASE_L1           Tag = 0x3040001D
	DASHBOARD_LC_PARAM_WB_CONNECTOR_PHASE_L2           Tag = 0x3040001E
	DASHBOARD_LC_PARAM_WB_CONNECTOR_PHASE_L3           Tag = 0x3040001F
	DASHBOARD_LC_PARAM_WB_CONNECTOR_PHASE_LIMIT        Tag = 0x30400020
	DASHBOARD_LC_PARAM_WB_CONNECTOR_ENERGY_CHARGED     Tag = 0x30400021
	DASHBOARD_LC_PARAM_CONTACT                         Tag = 0x30400022
	DASHBOARD_LC_PARAM_CONTACT_NAME                    Tag = 0x30400023
	DASHBOARD_LC_PARAM_CONTACT_FIRST_NAME              Tag = 0x30400024
	DASHBOARD_LC_PARAM_CONTACT_PHONE                   Tag = 0x30400025
	DASHBOARD_LC_PARAM_CONTACT_EMAIL                   Tag = 0x30400026
	DASHBOARD_LC_PARAM_CONTACT_ROLE                    Tag = 0x30400027
	DASHBOARD_LC_PARAM_CONTACT_TYPE                    Tag = 0x30400028
	DASHBOARD_LC_GET_SYSTEM_INFO                       Tag = 0x30800001
	DASHBOARD_LC_GET_CONSUMPTIONS                      Tag = 0x30800002
	DASHBOARD_LC_GET_CHARGING_SESSIONS                 Tag = 0x30800003
	DASHBOARD_LC_GET_SUPPORT_CONTACTS                  Tag = 0x30800004
)

// ---------------
// NAMESPACE: UMRC
// 0xF0xxxxxx
// undocumented
// ---------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	UMRC_REQ_FILE_LIST           Tag = 0xF0000104
	UMRC_REQ_MD5_HASH            Tag = 0xF0000105
	UMRC_REQ_FILE_CONTENT        Tag = 0xF0000106
	UMRC_REQ_CREATE_FOLDER       Tag = 0xF0000107
	UMRC_REQ_APPEND_FILE_CONTENT Tag = 0xF0000108
	UMRC_REQ_RENAME_FILE         Tag = 0xF0000109
	UMRC_REQ_DELETE_FILE         Tag = 0xF000010D
	UMRC_REQ_DELETE_FOLDER       Tag = 0xF000010E
	UMRC_REQ_SHOW_MSG            Tag = 0xF000010F
	UMRC_UM_AVAILABLE            Tag = 0xF0000110
	UMRC_REQ_CHK_SW              Tag = 0xF000A001
	UMRC_REQ_RESTART             Tag = 0xF000A005
	UMRC_REQ_MD5                 Tag = 0xF000A006
	UMRC_MD5                     Tag = 0xF000B005
	UMRC_FOLDER_NAME             Tag = 0xF000B007
	UMRC_FILE                    Tag = 0xF000B008
	UMRC_TYPE                    Tag = 0xF000B009
	UMRC_SIZE                    Tag = 0xF000B00A
	UMRC_VALUE                   Tag = 0xF000B00C
	UMRC_INDEX_FROM              Tag = 0xF000B00D
	UMRC_INDEX_UNTIL             Tag = 0xF000B00E
	UMRC_DATA_LEN                Tag = 0xF000B00F
	UMRC_DATA                    Tag = 0xF000B010
	UMRC_SRC_FILE                Tag = 0xF000B011
	UMRC_DST_FILE                Tag = 0xF000B012
	UMRC_FILE_LIST               Tag = 0xF0080104
	UMRC_MD5_HASH                Tag = 0xF0080105
	UMRC_FILE_CONTENT            Tag = 0xF0080106
	UMRC_FOLDER_CREATED          Tag = 0xF0080107
	UMRC_FILE_CONTENT_APPENDED   Tag = 0xF0080108
	UMRC_FILE_RENAMED            Tag = 0xF008010A
	UMRC_FILE_DELETED            Tag = 0xF008010D
	UMRC_FOLDER_DELETED          Tag = 0xF008010E
	UMRC_MSG_SHOWN               Tag = 0xF008010F
	UMRC_MD5_RSP                 Tag = 0xF008A006
)

// --------------
// NAMESPACE: LOG
// 0xF4xxxxxx
// undocumented
// --------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	LOG_REQ_LOGGER        Tag = 0xF400A001
	LOG_REQ_SET_LEVEL     Tag = 0xF400A002
	LOG_REQ_LOG           Tag = 0xF400A003
	LOG_NOTIFY_REGISTERED Tag = 0xF400A004
	LOG_PAR_LEVEL         Tag = 0xF400B001
	LOG_PAR_NAME          Tag = 0xF400B002
	LOG_PAR_CONTENT       Tag = 0xF400B003
	LOG_PAR_TIME          Tag = 0xF400B004
	LOG_PAR_SERVER_ID     Tag = 0xF400B005
	LOG_LOGGER            Tag = 0xF408A001
	LOG_SET_LEVEL         Tag = 0xF408A002
)

// --------------
// NAMESPACE: DCL
// 0xF5xxxxxx
// undocumented
// --------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	DCL_REQ_RM_ENTRY           Tag = 0xF5000AF0
	DCL_REQ_CL_LIST            Tag = 0xF500A001
	DCL_REQ_INSERT_CL_MSG      Tag = 0xF500A002
	DCL_REQ_REGISTER_ON_DEVICE Tag = 0xF500A003
	DCL_SESSION_ID             Tag = 0xF500B001
	DCL_GROUP_ID               Tag = 0xF500B002
	DCL_USER_ID                Tag = 0xF500B003
	DCL_TIME                   Tag = 0xF500B004
	DCL_USERNAME               Tag = 0xF500B005
	DCL_TICKET_REF             Tag = 0xF500B006
	DCL_DESCRIPTION            Tag = 0xF500B007
	DCL_SERIALNO               Tag = 0xF500B008
	DCL_LOCAL_ADDRESS          Tag = 0xF500B009
	DCL_ADD_CL_ELEMENT         Tag = 0xF508A001
	DCL_ADD_CL_SESSION         Tag = 0xF508A002
	DCL_REGISTER_ON_DEVICE     Tag = 0xF508A003
)

// -------------
// NAMESPACE: CL
// 0xF6xxxxxx
// undocumented
// -------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	CL_CLIENT_LIST_AVAILABLE Tag = 0xF600A001
	CL_REQ_CLIENT_LIST       Tag = 0xF600A002
	CL_REQ_CONNECTED_USER    Tag = 0xF600A004
	CL_REQ_INET_ADDR         Tag = 0xF600A005
	CL_REQ_CONNECTION_TIME   Tag = 0xF600A006
	CL_REQ_DIAG              Tag = 0xF600A007
	CL_REQ_UPDATE            Tag = 0xF600A008
	CL_FILTER_TYPE           Tag = 0xF600B001
	CL_LIST_FILTER           Tag = 0xF600B002
	CL_DEVICE_TYPE           Tag = 0xF600B003
	CL_SERIALNO              Tag = 0xF600B004
	CL_PRODUCTION_DATE       Tag = 0xF600B005
	CL_MAC_ADDRESS           Tag = 0xF600B006
	CL_IP_ADDRESS            Tag = 0xF600B007
	CL_SUBNET_MASK           Tag = 0xF600B008
	CL_GATEWAY               Tag = 0xF600B009
	CL_DNS                   Tag = 0xF600B010
	CL_DHCP_STATUS           Tag = 0xF600B011
	CL_SYSTEM_TIME           Tag = 0xF600B012
	CL_TIME_ZONE             Tag = 0xF600B013
	CL_UTC_TIME              Tag = 0xF600B014
	CL_A35_SERIALNO          Tag = 0xF600B015
	CL_REG_ID                Tag = 0xF600B016
	CL_IS_ONLINE             Tag = 0xF600B017
	CL_USERNAME              Tag = 0xF600B018
	CL_PARAM_DIAG            Tag = 0xF600B019
	CL_LIST_LIMIT            Tag = 0xF600B01A
	CL_ADD_CLIENT            Tag = 0xF608A002
	CL_REMOVE_CLIENT         Tag = 0xF608A003
	CL_CONNECTED_USER        Tag = 0xF608A004
	CL_INET_ADDR             Tag = 0xF608A005
	CL_CONNECTION_TIME       Tag = 0xF608A006
	CL_DIAG                  Tag = 0xF608A007
)

// ---------------
// NAMESPACE: DBRC
// 0xF7xxxxxx
// undocumented
// ---------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	DBRC_REQ_SET_SERVICE_PRIO             Tag = 0xF7000001
	DBRC_REQ_REMOVE_SERVICE_PRIO          Tag = 0xF7000002
	DBRC_REQ_SW_RELEASE_LIST              Tag = 0xF7000003
	DBRC_DEVICE_LIST_AVAILABLE            Tag = 0xF700A000
	DBRC_REQ_DEVICE_LIST                  Tag = 0xF700A001
	DBRC_REQ_USER_DATA                    Tag = 0xF700A002
	DBRC_REQ_ADDRESS                      Tag = 0xF700A003
	DBRC_REQ_MY_PERMISSIONS               Tag = 0xF700A004
	DBRC_REQ_LOCK_DEVICE                  Tag = 0xF700A005
	DBRC_REQ_UNLOCK_DEVICE                Tag = 0xF700A006
	DBRC_REQ_LOCK_STATE                   Tag = 0xF700A007
	DBRC_REQ_UPDATE_MACROS                Tag = 0xF700A008
	DBRC_REQ_MACROS                       Tag = 0xF700A009
	DBRC_REQ_BAT_ANALYSED                 Tag = 0xF700A021
	DBRC_REQ_BAT_ANALYSED2                Tag = 0xF700A022
	DBRC_REQ_BAT_ANALYSED3                Tag = 0xF700A023
	DBRC_REQ_BAT_ANALYSED4                Tag = 0xF700A024
	DBRC_SERIALNO                         Tag = 0xF700B001
	DBRC_INIT_CHECK                       Tag = 0xF700B002
	DBRC_TFO                              Tag = 0xF700B003
	DBRC_INSTALL_DATE                     Tag = 0xF700B004
	DBRC_OWNER_ID                         Tag = 0xF700B005
	DBRC_ADDRESS_ID                       Tag = 0xF700B006
	DBRC_INSTALLER_ID                     Tag = 0xF700B007
	DBRC_SW_RELEASE                       Tag = 0xF700B008
	DBRC_DEVICE_TYPE                      Tag = 0xF700B009
	DBRC_FILTER_TYPE                      Tag = 0xF700B00A
	DBRC_FILTER                           Tag = 0xF700B00B
	DBRC_COUPLING_MODE                    Tag = 0xF700B010
	DBRC_TUNNEL_PORT                      Tag = 0xF700B011
	DBRC_M_SPEC_ID                        Tag = 0xF700B012
	DBRC_M_SPEC_CATEGORY                  Tag = 0xF700B013
	DBRC_M_SPEC_USER_REF                  Tag = 0xF700B014
	DBRC_M_SPEC_ORDER_INDEX               Tag = 0xF700B015
	DBRC_M_DEF_SMD_ID                     Tag = 0xF700B016
	DBRC_M_DEF_NAME                       Tag = 0xF700B017
	DBRC_M_DEF_LEVEL                      Tag = 0xF700B018
	DBRC_M_DEF_NR_PARAM                   Tag = 0xF700B019
	DBRC_M_DEF_DESCR                      Tag = 0xF700B020
	DBRC_M_DEF_CMD                        Tag = 0xF700B021
	DBRC_M_DEF_SPEC_REF                   Tag = 0xF700B022
	DBRC_M_DEF_ORDER_INDEX                Tag = 0xF700B023
	DBRC_M_DEF_SYS_CMD                    Tag = 0xF700B024
	DBRC_DB_ENTRY_HEADER                  Tag = 0xF700B041
	DBRC_DB_ENTRY_ROWS                    Tag = 0xF700B042
	DBRC_DB_ENTRY_ROW                     Tag = 0xF700B043
	DBRC_VALUE                            Tag = 0xF700B044
	DBRC_BAT_ANALYSE_ID                   Tag = 0xF700B045
	DBRC_DONE                             Tag = 0xF700B046
	DBRC_BPM_TYPE                         Tag = 0xF700B047
	DBRC_DCB_TYPE                         Tag = 0xF700B048
	DBRC_NR_DCB                           Tag = 0xF700B049
	DBRC_ANALYSE_TYPE                     Tag = 0xF700B050
	DBRC_SN_FROM                          Tag = 0xF700B051
	DBRC_SN_UNTIL                         Tag = 0xF700B052
	DBRC_INSTALLED_FROM                   Tag = 0xF700B053
	DBRC_INSTALLED_UNTIL                  Tag = 0xF700B054
	DBRC_PARAM_ID                         Tag = 0xF7040001
	DBRC_PARAM_UD_USERNAME                Tag = 0xF7040002
	DBRC_PARAM_UD_NAME                    Tag = 0xF7040003
	DBRC_PARAM_UD_PRENAME                 Tag = 0xF7040004
	DBRC_PARAM_UD_IPIN                    Tag = 0xF7040005
	DBRC_PARAM_UD_LEVEL                   Tag = 0xF7040006
	DBRC_PARAM_UD_ID_PRIVATE_ADDR         Tag = 0xF7040007
	DBRC_PARAM_UD_ID_FIRM_ADDR            Tag = 0xF7040008
	DBRC_PARAM_ADR_PLZ                    Tag = 0xF7040009
	DBRC_PARAM_ADR_PLACE                  Tag = 0xF704000A
	DBRC_PARAM_ADR_GPS_LATITUDE           Tag = 0xF704000B
	DBRC_PARAM_ADR_GPS_LONGITUE           Tag = 0xF704000C
	DBRC_PARAM_ADR_PHONE                  Tag = 0xF704000D
	DBRC_PARAM_ADR_EMAIL                  Tag = 0xF704000E
	DBRC_PARAM_ADR_ADDRESS                Tag = 0xF704000F
	DBRC_PARAM_ADR_COUNTRY                Tag = 0xF7040010
	DBRC_PERMISSION                       Tag = 0xF7040011
	DBRC_PARAM_DCDC_TYPE                  Tag = 0xF7040012
	DBRC_PARAM_LAST_IP                    Tag = 0xF7040013
	DBRC_PARAM_PVI_TYPE                   Tag = 0xF7040014
	DBRC_PARAM_PVI_HW                     Tag = 0xF7040015
	DBRC_ADD_DEVICE                       Tag = 0xF708A001
	DBRC_ADD_USER                         Tag = 0xF708A002
	DBRC_ADD_ADDRESS                      Tag = 0xF708A003
	DBRC_MY_PERMISSIONS                   Tag = 0xF708A004
	DBRC_DEV_LOCK_STATE                   Tag = 0xF708A007
	DBRC_MACRO_SPEC                       Tag = 0xF708A009
	DBRC_MACRO_DEV                        Tag = 0xF708A00A
	DBRC_BAT_ANALYSED                     Tag = 0xF708A021
	DBRC_BAT_ANALYSED2                    Tag = 0xF708A022
	DBRC_BAT_ANALYSED3                    Tag = 0xF708A023
	DBRC_BAT_ANALYSED4                    Tag = 0xF708A024
	DBRC_SERVICE_PRIORITY                 Tag = 0xF7400001
	DBRC_SW_RELEASE_ENTRY                 Tag = 0xF7400002
	DBRC_PARAM_SW_RELEASE_NAME            Tag = 0xF7400003
	DBRC_PARAM_SW_RELEASE_UPDATE_PLATFORM Tag = 0xF7400004
	DBRC_UPDATE_PLATFORM                  Tag = 0xF7400005
	DBRC_PARAM_SERVICE_PRIO_EXPIRATION    Tag = 0xF7400006
	DBRC_PARAM_SERIALNO                   Tag = 0xF7400007
	DBRC_SET_SERVICE_PRIO                 Tag = 0xF7800001
	DBRC_REMOVE_SERVICE_PRIO              Tag = 0xF7800002
	DBRC_SW_RELEASE_LIST                  Tag = 0xF7800003
	DBRC_USERNAME                         Tag = 0xF7800004
)

// -----------------
// NAMESPACE: SERVER
// 0xF8xxxxxx
// undocumented
// -----------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	SERVER_REGISTER_CONNECTION        Tag = 0xF800A001
	SERVER_UNREGISTER_CONNECTION      Tag = 0xF800A002
	SERVER_REQ_RSCP_CMD               Tag = 0xF800A003
	SERVER_REQ_PING                   Tag = 0xF800A004
	SERVER_REQ_NEW_VIRTUAL_CONNECTION Tag = 0xF800A005
	SERVER_REQ_IS_RC_SERVER_CONNECTED Tag = 0xF800A00A
	SERVER_CONNECTION_ID              Tag = 0xF800B001
	SERVER_AUTH_LEVEL                 Tag = 0xF800B002
	SERVER_STATUS                     Tag = 0xF800B003
	SERVER_RSCP_DATA_LEN              Tag = 0xF800B004
	SERVER_RSCP_DATA                  Tag = 0xF800B005
	SERVER_TYPE                       Tag = 0xF800B006
	SERVER_HASH_CODE                  Tag = 0xF800B007
	SERVER_USER                       Tag = 0xF800B008
	SERVER_PASSWD                     Tag = 0xF800B009
	SERVER_IDENTIFIER                 Tag = 0xF800B010
	SERVER_CONNECTION_REGISTERED      Tag = 0xF808A001
	SERVER_CONNECTION_UNREGISTERED    Tag = 0xF808A002
	SERVER_RSCP_CMD_RESP              Tag = 0xF808A003
	SERVER_PING                       Tag = 0xF808A004
	SERVER_IS_RC_SERVER_CONNECTED     Tag = 0xF808A00A
)

// ------------------
// NAMESPACE: SYS_CMD
// 0xF9xxxxxx
// undocumented
// ------------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	SYS_CMD_REQ_FB                         Tag = 0xF9000001
	SYS_CMD_REQ_SYSTEM_REBOOT              Tag = 0xF9000005
	SYS_CMD_REQ_SYSTEM_KILL                Tag = 0xF9000006
	SYS_CMD_REQ_SERVER_SOCKETS             Tag = 0xF9000007
	SYS_CMD_REQ_SYS_EXECUTE                Tag = 0xF9000010
	SYS_CMD_REQ_SYS_EXECUTE_CANCEL         Tag = 0xF9000011
	SYS_CMD_REQ_SHELL_REGISTER             Tag = 0xF9000012
	SYS_CMD_REQ_SHELL_UNREGISTER           Tag = 0xF9000013
	SYS_CMD_REQ_SHELL_IN                   Tag = 0xF9000014
	SYS_CMD_REQ_SHELL_OUT                  Tag = 0xF9000015
	SYS_CMD_REQ_REGISTER_REMOTE_DEBUG      Tag = 0xF9000016
	SYS_CMD_REQ_UNREGISTER_REMOTE_DEBUG    Tag = 0xF9000017
	SYS_CMD_REQ_ALL_DEBUG_DEVICES          Tag = 0xF9000018
	SYS_CMD_REQ_ADD_DEBUG_DEVICE           Tag = 0xF9000019
	SYS_CMD_REQ_SET_LEVEL                  Tag = 0xF9000020
	SYS_CMD_REQ_REMOVE_DEBUG_DEVICE        Tag = 0xF9000021
	SYS_CMD_REQ_LOG_MSG                    Tag = 0xF9000022
	SYS_CMD_REQ_INLINE_SYS_EXECUTE         Tag = 0xF9000023
	SYS_CMD_REQ_VIRTUAL_CONNECTIONS        Tag = 0xF9000024
	SYS_CMD_REQ_SYS_EXECUTE_INTERACTIVE    Tag = 0xF9000026
	SYS_CMD_REQ_REGISTER_SYS_OBSERVER      Tag = 0xF9000100
	SYS_CMD_REQ_UNREGISTER_SYS_OBSERVER    Tag = 0xF9000101
	SYS_CMD_REQ_CURRENT_FOLDER             Tag = 0xF9000102
	SYS_CMD_REQ_LIST_ACTIVATOR_STATUS      Tag = 0xF9000103
	SYS_CMD_REQ_FILE_LIST                  Tag = 0xF9000104
	SYS_CMD_REQ_MD5_HASH                   Tag = 0xF9000105
	SYS_CMD_REQ_FILE_CONTENT               Tag = 0xF9000106
	SYS_CMD_REQ_CREATE_FOLDER              Tag = 0xF9000107
	SYS_CMD_REQ_APPEND_FILE_CONTENT        Tag = 0xF9000108
	SYS_CMD_REQ_RENAME_FILE                Tag = 0xF9000109
	SYS_CMD_REQ_CREATE_SYM_LINK            Tag = 0xF900010B
	SYS_CMD_REQ_DELETE_FILE                Tag = 0xF900010D
	SYS_CMD_REQ_DELETE_FOLDER              Tag = 0xF900010E
	SYS_CMD_REQ_REGISTER_PROC_OBS          Tag = 0xF9000200
	SYS_CMD_REQ_PROC_ID                    Tag = 0xF9000203
	SYS_CMD_REQ_PROC_FILE                  Tag = 0xF9000204
	SYS_CMD_REQ_PROC_LIST                  Tag = 0xF9000205
	SYS_CMD_REQ_UNREGISTER_PROC_OBS        Tag = 0xF9000206
	SYS_CMD_REQ_REGISTER_PUSH_SERVICE      Tag = 0xF9000207
	SYS_CMD_PUSH_SERVICE_ID                Tag = 0xF9000208
	SYS_CMD_REQ_PUSH_MESSAGE               Tag = 0xF9000209
	SYS_CMD_PUSH_SERVICE_IDRL              Tag = 0xF9000210
	SYS_CMD_PUSH_MESSAGE_CONTENT           Tag = 0xF9000211
	SYS_CMD_REQ_UNREGISTER_PUSH_SERVICE    Tag = 0xF9000212
	SYS_CMD_NAME                           Tag = 0xF900B001
	SYS_CMD_ID                             Tag = 0xF900B002
	SYS_CMD_LEVEL                          Tag = 0xF900B003
	SYS_CMD_MESSAGE                        Tag = 0xF900B004
	SYS_CMD_STATUS                         Tag = 0xF900B005
	SYS_CMD_VERSION                        Tag = 0xF900B006
	SYS_CMD_FOLDER_NAME                    Tag = 0xF900B007
	SYS_CMD_FILE                           Tag = 0xF900B008
	SYS_CMD_TYPE                           Tag = 0xF900B009
	SYS_CMD_SIZE                           Tag = 0xF900B00A
	SYS_CMD_VALUE                          Tag = 0xF900B00C
	SYS_CMD_INDEX_FROM                     Tag = 0xF900B00D
	SYS_CMD_INDEX_UNTIL                    Tag = 0xF900B00E
	SYS_CMD_DATA_LEN                       Tag = 0xF900B00F
	SYS_CMD_DATA                           Tag = 0xF900B010
	SYS_CMD_SRC_FILE                       Tag = 0xF900B011
	SYS_CMD_DST_FILE                       Tag = 0xF900B012
	SYS_CMD_LAST_ACCESS                    Tag = 0xF900B013
	SYS_CMD_LAST_MODIFICATION              Tag = 0xF900B014
	SYS_CMD_LAST_STATUS_CHANGE             Tag = 0xF900B015
	SYS_CMD_PROC_LIST_ENTRY                Tag = 0xF900B016
	SYS_CMD_PROC_LIST_CURR_FOLDER          Tag = 0xF900B017
	SYS_CMD_REQ_INJECT_TOUCH_EVENT         Tag = 0xF900B018
	SYS_CMD_SYSTEM_REBOOT                  Tag = 0xF9080005
	SYS_CMD_SYSTEM_KILL                    Tag = 0xF9080006
	SYS_CMD_SERVER_SOCKETS                 Tag = 0xF9080007
	SYS_CMD_SSOCKET                        Tag = 0xF9080008
	SYS_CMD_SYS_EXECUTE                    Tag = 0xF9080010
	SYS_CMD_SYS_EXECUTE_CANCEL             Tag = 0xF9080011
	SYS_CMD_SHELL_REGISTER                 Tag = 0xF9080012
	SYS_CMD_SHELL_UNREGISTER               Tag = 0xF9080013
	SYS_CMD_SHELL_IN                       Tag = 0xF9080014
	SYS_CMD_SHELL_OUT                      Tag = 0xF9080015
	SYS_CMD_REMOTE_DEBUG_REGISTERED        Tag = 0xF9080016
	SYS_CMD_REMOTE_DEBUG_UNREGISTERED      Tag = 0xF9080017
	SYS_CMD_INLINE_SYS_EXECUTE             Tag = 0xF9080023
	SYS_CMD_VIRTUAL_CONNECTIONS            Tag = 0xF9080024
	SYS_CMD_VIRTUAL_CONNECTION             Tag = 0xF9080025
	SYS_CMD_SYS_EXECUTE_INTERACTIVE        Tag = 0xF9080026
	SYS_CMD_SYS_OBSERVER_REGISTERED        Tag = 0xF9080100
	SYS_CMD_SYS_OBSERVER_UNREGISTERED      Tag = 0xF9080101
	SYS_CMD_CURRENT_FOLDER                 Tag = 0xF9080102
	SYS_CMD_LIST_ACTIVATOR_STATUS          Tag = 0xF9080103
	SYS_CMD_FILE_LIST                      Tag = 0xF9080104
	SYS_CMD_MD5_HASH                       Tag = 0xF9080105
	SYS_CMD_FILE_CONTENT                   Tag = 0xF9080106
	SYS_CMD_FOLDER_CREATED                 Tag = 0xF9080107
	SYS_CMD_FILE_CONTENT_APPENDED          Tag = 0xF9080108
	SYS_CMD_FILE_RENAMED                   Tag = 0xF908010A
	SYS_CMD_SYM_LINK_CREATED               Tag = 0xF908010B
	SYS_CMD_FILE_DELETED                   Tag = 0xF908010D
	SYS_CMD_FOLDER_DELETED                 Tag = 0xF908010E
	SYS_CMD_PROC_OBS_REGISTERED            Tag = 0xF9080200
	SYS_CMD_PROC_ID                        Tag = 0xF9080203
	SYS_CMD_PROC_FILE                      Tag = 0xF9080204
	SYS_CMD_PROC_LIST                      Tag = 0xF9080205
	SYS_CMD_PROC_OBS_UNREGISTERED          Tag = 0xF9080206
	SYS_CMD_ACTIVATOR_STATUS               Tag = 0xF9090103
	SYS_CMD_PARAM_FB_SCREEN_WIDTH          Tag = 0xF9400001
	SYS_CMD_PARAM_FB_SCREEN_HEIGHT         Tag = 0xF9400002
	SYS_CMD_PARAM_FB_BPS                   Tag = 0xF9400003
	SYS_CMD_PARAM_FB_RAW_DATA              Tag = 0xF9400004
	SYS_CMD_PARAM_FB_SCALE                 Tag = 0xF9400005
	SYS_CMD_PARAM_INJECT_TOUCH_EVENT_POS_Y Tag = 0xF9400006
	SYS_CMD_PARAM_INJECT_TOUCH_EVENT_POS_X Tag = 0xF9400007
	SYS_CMD_PARAM_FB_ALIGNMENT             Tag = 0xF9400008
	SYS_CMD_PARAM_FB_INTERPOLATE           Tag = 0xF9400009
	SYS_CMD_FB                             Tag = 0xF9800001
	SYS_CMD_REGISTER_PUSH_SERVICE          Tag = 0xF9800207
	SYS_CMD_PUSH_MESSAGE                   Tag = 0xF9800209
	SYS_CMD_UNREGISTER_PUSH_SERVICE        Tag = 0xF9800212
)

// ---------------------
// NAMESPACE: DB_SERVICE
// 0xFAxxxxxx
// undocumented
// ---------------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	DB_SERVICE_REQ_S10_SYS_INFO               Tag = 0xFA000001
	DB_SERVICE_REQ_SET_S10_SYS_INFO           Tag = 0xFA000002
	DB_SERVICE_NOTIFY_REGISTERED              Tag = 0xFA000003
	DB_SERVICE_REQ_SET_RELEASE_LATEST         Tag = 0xFA000004
	DB_SERVICE_REQ_SYNC_CONFIG_TO_TARGET      Tag = 0xFA000005
	DB_SERVICE_PARAM_SERIALNO                 Tag = 0xFA00B001
	DB_SERVICE_PARAM_TIMESTAMP_INIT_CHECK     Tag = 0xFA00B002
	DB_SERVICE_PARAM_TIMESTAMP_FIRST_ONLINE   Tag = 0xFA00B003
	DB_SERVICE_PARAM_TIMESTAMP_INSTALL_DATE   Tag = 0xFA00B004
	DB_SERVICE_PARAM_SW_RELEASE               Tag = 0xFA00B005
	DB_SERVICE_PARAM_ACC_KEY_STATUS           Tag = 0xFA00B006
	DB_SERVICE_PARAM_DEVICE_TYPE              Tag = 0xFA00B007
	DB_SERVICE_PARAM_COUPLING_MODE            Tag = 0xFA00B008
	DB_SERVICE_PARAM_TUNNEL_PORT              Tag = 0xFA00B009
	DB_SERVICE_PARAM_SW_UPDATE_DISABLED       Tag = 0xFA00B010
	DB_SERVICE_PARAM_HW_UPDATE_DISABLED       Tag = 0xFA00B011
	DB_SERVICE_PARAM_COMMENT                  Tag = 0xFA00B012
	DB_SERVICE_PARAM_LAST_UPDATE              Tag = 0xFA00B013
	DB_SERVICE_PARAM_HIDE_FROM_DL             Tag = 0xFA00B014
	DB_SERVICE_PARAM_ERR_CODE                 Tag = 0xFA00B015
	DB_SERVICE_PARAM_TIMESTAMP_LAST_ONLINE    Tag = 0xFA00B016
	DB_SERVICE_PARAM_REPORTED_SW_RELEASE      Tag = 0xFA00B017
	DB_SERVICE_PARAM_REPORTED_SW_RELEASE_DATE Tag = 0xFA00B018
	DB_SERVICE_PARAM_SW_UPDATE_LAST_CONNECT   Tag = 0xFA00B019
	DB_SERVICE_PARAM_UPDATE_PLATFORM          Tag = 0xFA00B01A
	DB_SERVICE_PARAM_HIDDEN                   Tag = 0xFA00B01B
	DB_SERVICE_REQ_REMOVABLE_DEVICES          Tag = 0xFA00C001
	DB_SERVICE_REMOVABLE_DEVICE_NAME          Tag = 0xFA00C002
	DB_SERVICE_REQ_REMOVE_DEVICE              Tag = 0xFA00C003
	DB_SERVICE_REQ_TABLE_DATA                 Tag = 0xFA00D001
	DB_SERVICE_TABLE_NAME                     Tag = 0xFA00D002
	DB_SERVICE_TABLE_MAX_COLUMNS              Tag = 0xFA00D003
	DB_SERVICE_TABLE_INCLUDE_METADATA         Tag = 0xFA00D004
	DB_SERVICE_TABLE_COLUMNS                  Tag = 0xFA00D005
	DB_SERVICE_TABLE_COLUMN                   Tag = 0xFA00D006
	DB_SERVICE_TABLE_COLUMN_NAME              Tag = 0xFA00D007
	DB_SERVICE_TABLE_COLUMN_TYPE              Tag = 0xFA00D008
	DB_SERVICE_TABLE_DATA_ROWS                Tag = 0xFA00D009
	DB_SERVICE_TABLE_DATA_CONTENT             Tag = 0xFA00D00A
	DB_SERVICE_REQ_RESYNC                     Tag = 0xFA00E001
	DB_SERVICE_RESYNC_FROM                    Tag = 0xFA00E002
	DB_SERVICE_REQ_SYNC_CONFIG_PROPERTIES     Tag = 0xFA00E003
	DB_SERVICE_S10_SYS_INFO                   Tag = 0xFA800001
	DB_SERVICE_SET_S10_SYS_INFO               Tag = 0xFA800002
	DB_SERVICE_SET_RELEASE_LATEST             Tag = 0xFA800004
	DB_SERVICE_SYNC_CONFIG_TO_TARGET          Tag = 0xFA800005
	DB_SERVICE_REMOVABLE_DEVICES              Tag = 0xFA80C001
	DB_SERVICE_REMOVE_DEVICE                  Tag = 0xFA80C003
	DB_SERVICE_TABLE_DATA                     Tag = 0xFA80D001
	DB_SERVICE_RESYNC                         Tag = 0xFA80E001
)

// ----------------------
// NAMESPACE: DB_RECOVERY
// 0xFBxxxxxx
// undocumented
// ----------------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	DB_RECOVERY_FINISH                 Tag = 0xFB000000
	DB_RECOVERY_REQ_RECOVERY           Tag = 0xFB000001
	DB_RECOVERY_SOURCE                 Tag = 0xFB000002
	DB_RECOVERY_TARGET                 Tag = 0xFB000003
	DB_RECOVERY_RECOVERY               Tag = 0xFB000004
	DB_RECOVERY_ERROR                  Tag = 0xFB000005
	DB_RECOVERY_LOG                    Tag = 0xFB000006
	DB_RECOVERY_REQ_RECOVERY_CONFIRM   Tag = 0xFB000007
	DB_RECOVERY_RECOVERY_CONFIRM       Tag = 0xFB000008
	DB_RECOVERY_REQ_RECOVERY_CANCEL    Tag = 0xFB000009
	DB_RECOVERY_REQ_CANDIDATE_LIST     Tag = 0xFB00000A
	DB_RECOVERY_REQ_STATUS             Tag = 0xFB00000B
	DB_RECOVERY_REQ_FINISH             Tag = 0xFB00000C
	DB_RECOVERY_RECOVERY_CANCEL        Tag = 0xFB000010
	DB_RECOVERY_PARAM_USER             Tag = 0xFB400001
	DB_RECOVERY_PARAM_PASSWORD         Tag = 0xFB400002
	DB_RECOVERY_PARAM_CANDIDATE_SERIAL Tag = 0xFB400003
	DB_RECOVERY_STATUS                 Tag = 0xFB400004
	DB_RECOVERY_PARAM_ID               Tag = 0xFB400005
	DB_RECOVERY_CANDIDATE_LIST         Tag = 0xFB80000A
)

// ---------------------
// NAMESPACE: GROUP_CTRL
// 0xFCxxxxxx
// undocumented
// ---------------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	GROUP_CTRL_REQ_STATUS         Tag = 0xFC000001
	GROUP_CTRL_GROUP_ID           Tag = 0xFC000002
	GROUP_CTRL_READY              Tag = 0xFC000003
	GROUP_CTRL_P_OPERATION_POINT  Tag = 0xFC000004
	GROUP_CTRL_P_ACTUAL           Tag = 0xFC000005
	GROUP_CTRL_FORECAST_60MINUTES Tag = 0xFC000006
	GROUP_CTRL_REQ_CONTROL        Tag = 0xFC000007
	GROUP_CTRL_P_TARGET           Tag = 0xFC000008
	GROUP_CTRL_ACTIVE             Tag = 0xFC000009
	GROUP_CTRL_AWARD              Tag = 0xFC00000A
	GROUP_CTRL_TIME               Tag = 0xFC00000B
	GROUP_CTRL_STATUS             Tag = 0xFC800001
	GROUP_CTRL_CONTROL            Tag = 0xFC800007
)

// ----------------
// NAMESPACE: ADMIN
// 0xFDxxxxxx
// undocumented
// ----------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	ADMIN_MESSAGE                               Tag = 0xFD000001
	ADMIN_ERROR                                 Tag = 0xFD000002
	ADMIN_REQ_HISTORY_VALUES_AGGREGATION_RECALC Tag = 0xFD000003
	ADMIN_REQ_HISTORY_VALUES_AGGREGATION_STOP   Tag = 0xFD000004
	ADMIN_REQ_POWER_METER_AGGREGATION_RECALC    Tag = 0xFD000005
	ADMIN_REQ_POWER_METER_AGG_RECALC_STOP       Tag = 0xFD000006
	ADMIN_REQ_SYSTEMERROR                       Tag = 0xFD000007
	ADMIN_REQ_ASSEMBLY_SERIAL_MAPPING           Tag = 0xFD000009
	ADMIN_REQ_ASSEMBLY_SERIAL_HISTORY           Tag = 0xFD00000A
	ADMIN_REQ_EXEC_CMD_TARGETS                  Tag = 0xFD00000B
	ADMIN_REQ_EXEC_CMD_REQUEST                  Tag = 0xFD00000C
	ADMIN_PARAM_SN                              Tag = 0xFD400001
	ADMIN_PARAM_FROM                            Tag = 0xFD400002
	ADMIN_PARAM_TO                              Tag = 0xFD400003
	ADMIN_PARAM_PM_TYPE                         Tag = 0xFD400004
	ADMIN_PARAM_PM_ID                           Tag = 0xFD400005
	ADMIN_PARAM_ERRORCODE                       Tag = 0xFD400006
	ADMIN_PARAM_ERRORNAME                       Tag = 0xFD400007
	ADMIN_PARAM_ERRORDATE                       Tag = 0xFD400008
	ADMIN_PARAM_ERRORCOUNT                      Tag = 0xFD400009
	ADMIN_PARAM_ERRORDESCRIPTION                Tag = 0xFD40000A
	ADMIN_PARAM_SERIALNO                        Tag = 0xFD40000B
	ADMIN_PARAM_ASSEMBLY_SERIAL                 Tag = 0xFD40000C
	ADMIN_PARAM_S10_SERIAL                      Tag = 0xFD40000D
	ADMIN_PARAM_TIMESTAMP                       Tag = 0xFD40000E
	ADMIN_ASSEMBLY_SERIAL_HISTORY_ENTRY         Tag = 0xFD40000F
	ADMIN_PARAM_EXEC_CMD_REQUEST                Tag = 0xFD400010
	ADMIN_PARAM_EXEC_CMD_RESPONSE               Tag = 0xFD400011
	ADMIN_PARAM_EXEC_CMD_HASH                   Tag = 0xFD400012
	ADMIN_PARAM_EXEC_CMD_TARGET                 Tag = 0xFD400013
	ADMIN_HISTORY_VALUES_AGGREGATION_RECALC     Tag = 0xFD800003
	ADMIN_HISTORY_VALUES_AGGREGATION_STOP       Tag = 0xFD800004
	ADMIN_POWER_METER_AGGREGATION_RECALC        Tag = 0xFD800005
	ADMIN_POWER_METER_AGG_RECALC_STOP           Tag = 0xFD800006
	ADMIN_SYSTEMERROR                           Tag = 0xFD800007
	ADMIN_SYSTEMERROR_CODES                     Tag = 0xFD800008
	ADMIN_ASSEMBLY_SERIAL_MAPPING               Tag = 0xFD800009
	ADMIN_ASSEMBLY_SERIAL_HISTORY               Tag = 0xFD80000A
	ADMIN_EXEC_CMD_TARGETS                      Tag = 0xFD80000B
	ADMIN_EXEC_CMD_RESPONSE                     Tag = 0xFD80000C
)

// ----------------------
// NAMESPACE: FINAL_CHECK
// 0xFExxxxxx
// undocumented
// ----------------------
const (
	//
	// undocumented from portal.
	// most if not all will have unknown data type,
	// so you need to know the data type to make the requests.
	//
	FINAL_CHECK_PARAM_TEST_PASSED_ON         Tag = 0xFE000000
	FINAL_CHECK_REPORT                       Tag = 0xFE000001
	FINAL_CHECK_REQ_SERIAL_FOR_ASSEMBLY      Tag = 0xFE000002
	FINAL_CHECK_REQ_SERIAL_STATUS            Tag = 0xFE000004
	FINAL_CHECK_REQ_SUBMIT_TEST_PROTOCOL     Tag = 0xFE000005
	FINAL_CHECK_PARAM_SN                     Tag = 0xFE400001
	FINAL_CHECK_PARAM_STATUS                 Tag = 0xFE400002
	FINAL_CHECK_PARAM_START                  Tag = 0xFE400003
	FINAL_CHECK_PARAM_END                    Tag = 0xFE400004
	FINAL_CHECK_PARAM_META                   Tag = 0xFE400005
	FINAL_CHECK_PARAM_LOG                    Tag = 0xFE400006
	FINAL_CHECK_PARAM_TICKET                 Tag = 0xFE400007
	FINAL_CHECK_DEVICE_PROPERTY              Tag = 0xFE400008
	FINAL_CHECK_DEVICE_PROPERTY_MODULE_NAME  Tag = 0xFE400009
	FINAL_CHECK_DEVICE_PROPERTY_PROP_NAME    Tag = 0xFE40000A
	FINAL_CHECK_DEVICE_PROPERTY_ACTUAL_VALUE Tag = 0xFE40000B
	FINAL_CHECK_PARAM_ASSEMBLY_SERIAL        Tag = 0xFE40000C
	FINAL_CHECK_PARAM_SERIAL_NUMBER          Tag = 0xFE40000D
	FINAL_CHECK_ERROR_TAG_REFERENCE          Tag = 0xFE40000E
	FINAL_CHECK_ERROR_TEXT                   Tag = 0xFE40000F
	FINAL_CHECK_PARAM_TEST_STATUS            Tag = 0xFE400011
	FINAL_CHECK_PARAM_SUBMITTED_ON           Tag = 0xFE400012
	FINAL_CHECK_PARAM_PROTOCOL               Tag = 0xFE400013
	FINAL_CHECK_PARAM_PROTOCOL_SAVED         Tag = 0xFE400014
	FINAL_CHECK_PARAM_TEST_PASSED            Tag = 0xFE400016
	FINAL_CHECK_REPORT_SAVED                 Tag = 0xFE800001
	FINAL_CHECK_SERIAL_FOR_ASSEMBLY          Tag = 0xFE800002
	FINAL_CHECK_ERROR_MESSAGE                Tag = 0xFE800003
	FINAL_CHECK_SERIAL_STATUS                Tag = 0xFE800004
	FINAL_CHECK_SUBMIT_TEST_PROTOCOL         Tag = 0xFE800005
)
