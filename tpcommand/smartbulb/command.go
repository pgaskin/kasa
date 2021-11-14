// Package tpcommand contains types for TP-Link smart bulbs.
package smartbulb

import (
	"github.com/pgaskin/kasa/tpcommand"
	"github.com/pgaskin/kasa/tpcommand/device"
)

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.LightState
// note: TransitionPeriod, DftOnState, Effect, OnOff extracted to allow reuse
type LightState struct {
	Brightness *int    `json:"brightness,omitempty"`
	ColorTemp  *int    `json:"color_temp,omitempty"`
	Hue        *int    `json:"hue,omitempty"`
	Mode       *string `json:"mode,omitempty"`
	OnOff      *int    `json:"on_off,omitempty"`
	Saturation *int    `json:"saturation,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.LightingEffectState
// note: the inheritance of Method is a mistake
type LightingEffectState struct {
	Brightness *int    `json:"brightness,omitempty"`
	Custom     *int    `json:"custom,omitempty"`
	Enable     *int    `json:"enable,omitempty"`
	ID         *string `json:"id,omitempty"`
	Name       *string `json:"name,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.PreferredState
type PreferredState struct {
	LightState
	Index *int `json:"index,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Rule
type Rule struct {
	Day      *int        `json:"day,omitempty"`
	ELight   *LightState `json:"e_light,omitempty"`
	EAct     *int        `json:"eact,omitempty"`
	EMin     *int        `json:"emin,omitempty"`
	Enable   *int        `json:"enable,omitempty"`
	EOffset  *int        `json:"eoffset,omitempty"`
	ETimeOpt *int        `json:"etime_opt,omitempty"`
	ID       *string     `json:"id,omitempty"`
	Month    *int        `json:"month,omitempty"`
	Name     *string     `json:"name,omitempty"`
	Remain   *int        `json:"remain,omitempty"`
	Repeat   *int        `json:"repeat,omitempty"`
	SLight   *LightState `json:"s_light,omitempty"`
	SAct     *int        `json:"sact,omitempty"`
	SMin     *int        `json:"smin,omitempty"` // minutes relative to STimeOpt
	SOffset  *int        `json:"soffset,omitempty"`
	STimeOpt *int        `json:"stime_opt,omitempty"` // none (-1), midnight (0), sunrise (1), sunset (2)
	Wday     *[]int      `json:"wday,omitempty"`
	Year     *int        `json:"year,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.TimeSetting.GetTimeZone
// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.TimeSetting.SetTimeZone
// note: extracted for reuse since Go can do struct composition
type Time struct {
	Hour  *int `json:"hour,omitempty"`
	Mday  *int `json:"mday,omitempty"`
	Min   *int `json:"min,omitempty"`
	Month *int `json:"month,omitempty"`
	Sec   *int `json:"sec,omitempty"`
	Year  *int `json:"year,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand
type SmartBulbCommand struct {
	device.DeviceCommand
	Cloud            *CloudModule            `json:"smartlife.iot.common.cloud,omitempty"`
	Emeter           *EmeterModule           `json:"smartlife.iot.common.emeter,omitempty"`
	LightingService  *LightingServiceModule  `json:"smartlife.iot.smartbulb.lightingservice,omitempty"`
	Schedule         *ScheduleModule         `json:"smartlife.iot.common.schedule,omitempty"`
	SetUpgradeWar    *SetUpgradeWarModule    `json:"SmartBulb_debug,omitempty"`
	SoftAPOnboarding *SoftAPOnboardingModule `json:"smartlife.iot.common.softaponboarding,omitempty"`
	SysInfo          *SysInfoModule          `json:"system,omitempty"`
	TimeSetting      *TimeSettingModule      `json:"smartlife.iot.common.timesetting,omitempty"`
	Weave            *WeaveModule            `json:"smartlife.integration.weave,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Cloud
type CloudModule struct {
	tpcommand.Module
	Bind                  *BindMethod               `json:"bind,omitempty"`
	GetInfo               *GetInfoMethod            `json:"get_info,omitempty"`
	GetIntlFwList         *GetIntlFwListMethod      `json:"get_intl_fw_list,omitempty"`
	SetServerURL          *SetServerURLMethod       `json:"set_n_sefserver_url,omitempty"`
	SetSefServerURL       *SetServerURLMethod       `json:"set_n_server_url,omitempty"`
	LegacySetServerURL    *LegacySetServerURLMethod `json:"set_sefserver_url,omitempty"`
	LegacySetSefServerURL *LegacySetServerURLMethod `json:"set_server_url,omitempty"`
	Unbind                *UnbindMethod             `json:"unbind,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Cloud.Bind
type BindMethod struct {
	tpcommand.Method
	Password *string `json:"password,omitempty"`
	Username *string `json:"username,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Cloud.GetInfo
type GetInfoMethod struct {
	tpcommand.Method
	Binded       *int    `json:"binded,omitempty"`
	CldCnnection *int    `json:"cld_connection,omitempty"`
	FwDIPage     *string `json:"fwDIPage,omitempty"`
	FwNotifyType *int    `json:"fwNotifyType,omitempty"`
	IllegalType  *int    `json:"illegalType,omitempty"`
	Server       *string `json:"server,omitempty"`
	StopConnect  *int    `json:"stopConnect,omitempty"`
	TcspInfo     *string `json:"tcspInfo,omitempty"`
	TcspStatus   *int    `json:"tcspStatus,omitempty"`
	Username     *string `json:"username,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Cloud.GetIntlFwList
type GetIntlFwListMethod struct {
	tpcommand.Method
	FwList *[]map[string]interface{} `json:"fw_list,omitempty"`
}

// Deprecated: com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Cloud.LegacySetServerURL
type LegacySetServerURLMethod struct {
	tpcommand.Method
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Cloud.SetServerURL
type SetServerURLMethod struct {
	tpcommand.Method
	Server *string `json:"server,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Cloud.Unbind
type UnbindMethod struct {
	tpcommand.Method
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Emeter
type EmeterModule struct {
	tpcommand.Module
	EraseEmeterStat *EraseEmeterStatMethod `json:"erase_emeter_stat,omitempty"`
	GetDayStat      *GetDayStatMethod      `json:"get_daystat,omitempty"`
	GetMonthStat    *GetMonthStatMethod    `json:"get_monthstat,omitempty"`
	GetRealTime     *GetRealTimeMethod     `json:"get_realtime,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Emeter.EraseEmeterStat
type EraseEmeterStatMethod struct {
	tpcommand.Method
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Emeter.GetDayStat
// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Schedule.GetDayStat
type GetDayStatMethod struct {
	tpcommand.Method
	DayList *[]map[string]interface{} `json:"day_list,omitempty"`
	Month   *int                      `json:"month,omitempty"`
	Year    *int                      `json:"year,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Emeter.GetMonthStat
// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Schedule.GetMonthStat
type GetMonthStatMethod struct {
	tpcommand.Method
	DayList *[]map[string]interface{} `json:"month_list,omitempty"`
	Year    *int                      `json:"year,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Emeter.GetRealTime
type GetRealTimeMethod struct {
	tpcommand.Method
	PowerMW *int `json:"power_mw,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.LightingService
type LightingServiceModule struct {
	tpcommand.Module
	AdjustLightBrightness *AdjustLightBrightnessMethod `json:"adjust_light_brightness,omitempty"`
	GetDefaultBehavior    *GetDefaultBehaviorMethod    `json:"get_default_behavior,omitempty"`
	GetLightDetails       *GetLightDetailsMethod       `json:"get_light_details,omitempty"`
	GetLightState         *GetLightStateMethod         `json:"get_light_state,omitempty"`
	GetPreferredState     *GetPreferredStateMethod     `json:"get_preferred_state,omitempty"`
	SetDefaultBehavior    *SetDefaultBehaviorMethod    `json:"set_default_behavior,omitempty"`
	SetPreferredState     *SetPreferredStateMethod     `json:"set_preferred_state,omitempty"`
	TransitionLightState  *TransitionLightStateMethod  `json:"transition_light_state,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.LightingService.AdjustLightBrightness
type AdjustLightBrightnessMethod struct {
	tpcommand.Method
	Delta int `json:"delta,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.LightingService.GetDefaultBehavior
type GetDefaultBehaviorMethod struct {
	tpcommand.Method
	HardOn *PreferredState `json:"hard_on,omitempty"`
	SoftOn *PreferredState `json:"soft_on,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.LightingService.GetLightDetails
type GetLightDetailsMethod struct {
	tpcommand.Method
	ColorRenderingIndex    *int `json:"color_rendering_index,omitempty"`
	IncandescentEquivalent *int `json:"incandescent_equivalent,omitempty"`
	LampBeamAngle          *int `json:"lamp_beam_angle,omitempty"`
	MaxLumens              *int `json:"max_lumens,omitempty"`
	MaxVoltage             *int `json:"max_voltage,omitempty"`
	MinVoltage             *int `json:"min_voltage,omitempty"`
	Wattage                *int `json:"wattage,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.LightingService.GetLightState
// note: since Go allows composition, we can deduplicate this against LightState
type GetLightStateMethod struct {
	tpcommand.Method
	DftOnState *LightState `json:"dft_on_state,omitempty"`
	LightState
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.LightingService.GetPreferredState
type GetPreferredStateMethod struct {
	tpcommand.Method
	States *[]PreferredState `json:"states,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.LightingService.SetDefaultBehavior
type SetDefaultBehaviorMethod struct {
	tpcommand.Method
	HardOn *PreferredState `json:"hard_on,omitempty"`
	SoftOn *PreferredState `json:"soft_on,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.LightingService.SetPreferredState
// note: since Go allows composition, we can deduplicate this against LightState
type SetPreferredStateMethod struct {
	tpcommand.Method
	LightState
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.LightingService.TransitionLightState
// note: since Go allows composition, we can deduplicate this against LightState
type TransitionLightStateMethod struct {
	tpcommand.Method
	TransitionPeriod *int        `json:"transition_period,omitempty"`
	IgnoreDefault    *int        `json:"ignore_default,omitempty"`
	DftOnState       *LightState `json:"dft_on_state,omitempty"`
	LightState
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Schedule
type ScheduleModule struct {
	tpcommand.Module
	AddRule          *AddRuleMethod          `json:"add_rule,omitempty"`
	DeleteAllRules   *DeleteAllRulesMethod   `json:"delete_all_rules,omitempty"`
	DeleteRule       *DeleteRuleMethod       `json:"delete_rule,omitempty"`
	EditRule         *EditRuleMethod         `json:"edit_rule,omitempty"`
	EraseRuntimeStat *EraseRuntimeStatMethod `json:"erase_runtime_stat,omitempty"`
	GetDayStat       *GetDayStatMethod       `json:"get_daystat,omitempty"`
	GetMonthStat     *GetMonthStatMethod     `json:"get_monthstat,omitempty"`
	GetNextAction    *GetNextActionMethod    `json:"get_next_action,omitempty"`
	GetRules         *GetRulesMethod         `json:"get_rules,omitempty"`
	SetOverallEnable *SetOverallEnableMethod `json:"set_overall_enable,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Schedule.AddRule
// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Schedule.BaseRule
type AddRuleMethod struct {
	tpcommand.Method
	ConflictID *string `json:"conflict_id,omitempty"`
	Rule
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Schedule.DeleteAllRules
type DeleteAllRulesMethod struct {
	tpcommand.Method
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Schedule.DeleteRule
type DeleteRuleMethod struct {
	tpcommand.Method
	ID *string `json:"id,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Schedule.EditRule
// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Schedule.BaseRule
type EditRuleMethod struct {
	tpcommand.Method
	ConflictID *string `json:"conflict_id,omitempty"`
	Rule
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Schedule.EraseRuntimeStat
type EraseRuntimeStatMethod struct {
	tpcommand.Method
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Schedule.GetNextAction
type GetNextActionMethod struct {
	tpcommand.Method
	Action   *int        `json:"action,omitempty"`
	ID       *string     `json:"id,omitempty"`
	Light    *LightState `json:"light,omitempty"`
	SchdTime *int        `json:"schd_time,omitempty"`
	Type     *int        `json:"type,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Schedule.GetRules
type GetRulesMethod struct {
	tpcommand.Method
	Enable   *int    `json:"enable,omitempty"`
	RuleList *[]Rule `json:"rule_list,omitempty"`
	Version  *int    `json:"version,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.Schedule.SetOverallEnable
type SetOverallEnableMethod struct {
	tpcommand.Method
	Enable *int `json:"enable,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.tpcommon.model.smartlife.iot.common.setupgradewarmodule.SetUpgradeWarModule
type SetUpgradeWarModule struct {
	tpcommand.Module
	SetUpgradeWar *SetUpgradeWarMethod `json:"set_upgrade_war,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.tpcommon.model.smartlife.iot.common.setupgradewarmodule.methods.SetUpgradeWar
type SetUpgradeWarMethod struct {
	tpcommand.Method
	Value *int `json:"value,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.SoftAPOnboarding
type SoftAPOnboardingModule struct {
	tpcommand.Module
	GetScanInfo *GetScanInfoMethod `json:"get_scaninfo,omitempty"`
	GetStaInfo  *GetStaInfoMethod  `json:"get_stainfo,omitempty"`
	SetStaInfo  *SetStaInfoMethod  `json:"set_stainfo,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.SoftAPOnboarding.GetScanInfo
type GetScanInfoMethod struct {
	tpcommand.Method
	ApList        *[]map[string]interface{} `json:"ap_list,omitempty"`
	Refresh       *int                      `json:"refresh,omitempty"`
	WPA3Supported *string                   `json:"wpa3_supported,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.SoftAPOnboarding.GetStaInfo
type GetStaInfoMethod struct {
	tpcommand.Method
	KeyType *int    `json:"key_type,omitempty"`
	RSSI    *int    `json:"rssi,omitempty"`
	SSID    *string `json:"ssid,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.SoftAPOnboarding.SetStaInfo
type SetStaInfoMethod struct {
	tpcommand.Method
	CipherType *int    `json:"cipher_type,omitempty"`
	KeyIndex   *int    `json:"key_index,omitempty"`
	KeyType    *int    `json:"key_type,omitempty"`
	Password   *string `json:"password,omitempty"`
	SSID       *string `json:"ssid,omitempty"`
	WEPMode    *int    `json:"wep_mode,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.SysInfo
type SysInfoModule struct {
	tpcommand.Module
	GetSysInfo *GetSysInfoMethod `json:"get_sysinfo,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.GetSysInfo
type GetSysInfoMethod struct {
	tpcommand.Method
	LEF                 *int                 `json:"LEF,omitempty"`
	ActiveMode          *string              `json:"active_mode,omitempty"`
	Alias               *string              `json:"alias,omitempty"`
	Description         *string              `json:"description,omitempty"`
	DevState            *string              `json:"dev_state,omitempty"`
	DeviceID            *string              `json:"deviceId,omitempty"`
	HwID                *string              `json:"hwId,omitempty"`
	HwVer               *string              `json:"hw_ver,omitempty"`
	IsColor             *int                 `json:"is_color,omitempty"`
	IsDimmable          *int                 `json:"is_dimmable,omitempty"`
	IsVariableColorTemp *int                 `json:"is_variable_color_temp,omitempty"`
	LightState          *LightState          `json:"light_state,omitempty"`
	LightingEffectState *LightingEffectState `json:"lighting_effect_state,omitempty"`
	MicMac              *string              `json:"mic_mac,omitempty"`
	MicType             *string              `json:"mic_type,omitempty"`
	Model               *string              `json:"model,omitempty"`
	OemID               *string              `json:"oemId,omitempty"`
	PreferredState      *[]PreferredState    `json:"preferred_state,omitempty"`
	RelayState          *int                 `json:"relay_state,omitempty"`
	RSSI                *int                 `json:"rssi,omitempty"`
	SwVer               *string              `json:"sw_ver,omitempty"`
	TID                 *string              `json:"tid,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbUtils.getTemperatureRange
// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.iot.devices.DeviceRegistry.Light
func (g GetSysInfoMethod) GetTemperatureRange() (int, int, bool) {
	if g.Model == nil {
		return 0, 0, false
	}
	switch *g.Model {
	case "KL130B", "KL130", "KL135", "KL430", "LB130", "LB230":
		return 2500, 9000, true
	case "KL120":
		return 2700, 5000, true
	case "KL125":
		return 2500, 6500, true
	case "LB120":
		return 2700, 6500, true
	default:
		return 0, 0, false
	}
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.TimeSetting
type TimeSettingModule struct {
	tpcommand.Module
	GetTime     *GetTimeMethod     `json:"get_time,omitempty"`
	GetTimeZone *GetTimeZoneMethod `json:"get_timezone,omitempty"`
	SetTimeZone *SetTimeZoneMethod `json:"set_timezone,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.TimeSetting.GetTime
type GetTimeMethod struct {
	tpcommand.Method
	Time
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.TimeSetting.GetTimeZone
type GetTimeZoneMethod struct {
	tpcommand.Method
	DstOffset *string `json:"dst_offset,omitempty"`
	Index     *int    `json:"index,omitempty"`
	TzStr     *string `json:"tz_str,omitempty"`
	ZoneStr   *string `json:"zone_str,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.light.lball.api.TPSmartBulbCommand.TimeSetting.SetTimeZone
type SetTimeZoneMethod struct {
	tpcommand.Method
	Time
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.tpcommon.model.smartlife.iot.thirdpartyintegration.weave.Weave
type WeaveModule struct {
	tpcommand.Module
	Disable   *WeaveDisableMethod   `json:"disable,omitempty"`
	GetInfo   *WeaveGetInfoMethod   `json:"get_info,omitempty"`
	SetTicket *WeaveSetTicketMethod `json:"set_ticket,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.tpcommon.model.smartlife.iot.thirdpartyintegration.weave.methods.Disable
type WeaveDisableMethod struct {
	tpcommand.Method
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.tpcommon.model.smartlife.iot.thirdpartyintegration.weave.methods.GetInfo
type WeaveGetInfoMethod struct {
	tpcommand.Method
	DaemonRunning *string `json:"daemon_running,omitempty"`
	RegStatus     *int    `json:"reg_status,omitempty"`
	Registration  *string `json:"registration,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.tpcommon.model.smartlife.iot.thirdpartyintegration.weave.methods.SetTicket
type WeaveSetTicketMethod struct {
	tpcommand.Method
}
