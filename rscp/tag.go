package rscp

// Tag datatype
type Tag uint32

// all tags as constant
//
//nolint:revive,stylecheck
//go:generate go run github.com/alvaroloes/enumer -type=Tag -json
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
	HA_REQ_DATAPOINT_LIST        Tag = 0x09000001
	HA_REQ_ACTUATOR_STATES       Tag = 0x09000010
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
	HA_DEVICE_STATE              Tag = 0x09860000
	HA_DEVICE_CONNECTED          Tag = 0x09860001
	HA_DEVICE_WORKING            Tag = 0x09860002
	HA_DEVICE_IN_SERVICE         Tag = 0x09860003
	HA_GENERAL_ERROR             Tag = 0x09FFFFFF
	SRV_REQ_IS_ONLINE            Tag = 0x08000001
	SRV_IS_ONLINE                Tag = 0x08800001
	SRV_REQ_ADD_USER             Tag = 0x08000002
	SRV_ADD_USER                 Tag = 0x08800002
	SRV_GENERAL_ERROR            Tag = 0x08FFFFFF
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
	INFO_INFO                  Tag = 0x0A800011
	INFO_SET_IP_ADDRESS        Tag = 0x0A800012
	INFO_SET_SUBNET_MASK       Tag = 0x0A800013
	INFO_SET_DHCP_STATUS       Tag = 0x0A800014
	INFO_SET_GATEWAY           Tag = 0x0A800015
	INFO_SET_DNS               Tag = 0x0A800016
	INFO_SET_TIME              Tag = 0x0A800017
	INFO_SET_TIME_ZONE         Tag = 0x0A800018
	INFO_SW_RELEASE            Tag = 0x0A800019
	INFO_GENERAL_ERROR         Tag = 0x0AFFFFFF
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
	DB_HISTORY_DATA_YEAR  Tag = 0x06800400
	DB_PAR_TIME_MIN       Tag = 0x06B00000
	DB_PAR_TIME_MAX       Tag = 0x06B00001
	DB_PARAM_ROW          Tag = 0x06B00002
	DB_PARAM_COLUMN       Tag = 0x06B00003
	DB_PARAM_INDEX        Tag = 0x06B00004
	DB_PARAM_VALUE        Tag = 0x06B00005
	DB_PARAM_MAX_ROWS     Tag = 0x06B00006
	DB_PARAM_TIME         Tag = 0x06B00007
	DB_PARAM_VERSION      Tag = 0x06B00008
	DB_PARAM_HEADER       Tag = 0x06B00009
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
	UM_REQ_UPDATE_STATUS    Tag = 0x0D000001
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
)
