// Package device contains types for TP-Link smart devices.
package device

import "github.com/pgaskin/kasa/tpcommand"

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.tpcommon.model.smartlife.iot.TPDeviceCommand
type DeviceCommand struct {
	Context *ContextModule `json:"context,omitempty"`
	System  *SystemModule  `json:"system,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.tpcommon.model.smartlife.iot.common.Context
type ContextModule struct {
	tpcommand.Module
	Source *string `json:"source,omitempty"`
	TID    *int    `json:"tid,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.tpcommon.model.smartlife.iot.common.system.System
type SystemModule struct {
	tpcommand.Module
	DownloadFirmware *DownloadFirmwareMethod `json:"download_firmware,omitempty"`
	GetDownloadState *GetDownloadStateMethod `json:"get_download_state,omitempty"`
	Reboot           *RebootMethod           `json:"reboot,omitempty"`
	Reset            *ResetMethod            `json:"reset,omitempty"`
	SetDevAlias      *SetDevAliasMethod      `json:"set_dev_alias,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.tpcommon.model.smartlife.iot.common.system.methods.DownloadFirmware
type DownloadFirmwareMethod struct {
	tpcommand.Method
	AutoFlash  *bool   `json:"auto_flash,omitempty"`
	AutoReboot *bool   `json:"auto_reboot,omitempty"`
	URL        *string `json:"url,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.tpcommon.model.smartlife.iot.common.system.methods.GetDownloadState
type GetDownloadStateMethod struct {
	tpcommand.Method
	FlashTime  *int `json:"flash_time,omitempty"`
	Ratio      *int `json:"ratio,omitempty"`
	RebootTime *int `json:"reboot_time,omitempty"`
	Status     *int `json:"status,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.tpcommon.model.smartlife.iot.common.system.methods.Reboot
type RebootMethod struct {
	tpcommand.Method
	Delay *int `json:"delay,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.tpcommon.model.smartlife.iot.common.system.methods.Reset
type ResetMethod struct {
	tpcommand.Method
	NoReboot *int `json:"noReboot,omitempty"`
}

// com.tplink.kasa_android@2.35.0.1021/com.tplinkra.tpcommon.model.smartlife.iot.common.system.methods.SetDevAlias
type SetDevAliasMethod struct {
	tpcommand.Method
	Alias *string `json:"alias,omitempty"`
}
