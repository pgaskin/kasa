package tpcommand

import (
	"fmt"
	"strconv"
	"strings"
)

// Error is a TP-Link API error with a message.
//
// Based on:
// - com.tplink.kasa_android@2.35.0.1021/com.tplinkra.iot.ErrorConstants
// - com.tplink.kasa_android@2.35.0.1021/com.tplinkra.tpcommon.model.smartlife.iot.Method
type Error struct {
	Code    ErrorCode
	Message string
}

func (err Error) Error() string {
	var b strings.Builder
	b.WriteString("tp-link api error")
	if err.Code.IsError() {
		b.WriteByte(' ')
		b.WriteString(err.Code.String())
	} else if !(err.Code == 0 && err.Message != "") {
		panic(fmt.Errorf("%#v is not an error", err))
	}
	if err.Message != "" {
		b.WriteByte(':')
		b.WriteByte(' ')
		b.WriteString(err.Message)
	}
	return b.String()
}

// ErrorCode is a TP-Link API error.
//
// Based on:
// - com.tplink.kasa_android@2.35.0.1021/com.tplinkra.iot.ErrorConstants
// - com.tplink.kasa_android@2.35.0.1021/com.tplinkra.tpcommon.model.smartlife.iot.Method
type ErrorCode int

func (err ErrorCode) IsError() bool {
	return err < 0
}

func (err ErrorCode) String() string {
	if !err.IsError() {
		return "no error"
	}
	var b strings.Builder
	b.WriteString(strconv.Itoa(int(err)))
	if ns := err.Names(); len(ns) != 0 {
		b.WriteByte(' ')
		b.WriteByte('(')
		for i, n := range ns {
			if i != 0 {
				b.WriteByte(' ')
				b.WriteByte('|')
				b.WriteByte('|')
				b.WriteByte(' ')
			}
			b.WriteString(n)
		}
		b.WriteByte(')')
	}
	return b.String()
}

func (c ErrorCode) ToError() error {
	return Error{c, ""}
}

func (c ErrorCode) WithMessage(msg string) error {
	return Error{c, msg}
}

func (c ErrorCode) WithMessagef(format string, a ...interface{}) error {
	return c.WithMessage(fmt.Sprintf(format, a...))
}

const (
	ACCOUNT_FEATURES_DATASET_CORRUPTED                         ErrorCode = -94005
	ACCOUNT_FEATURES_DEVICE_ADDRESS_REQUIRED                   ErrorCode = -94011
	ACCOUNT_FEATURES_DEVICE_CONTEXT_REQUIRED                   ErrorCode = -94007
	ACCOUNT_FEATURES_DEVICE_ID_NOT_MATCH                       ErrorCode = -94008
	ACCOUNT_FEATURES_DEVICE_MODEL_REQUIRED                     ErrorCode = -94010
	ACCOUNT_FEATURES_DEVICE_TYPE_REQUIRED                      ErrorCode = -94009
	ACCOUNT_FEATURES_FAILED_DUE_TO_DB_ERRORS                   ErrorCode = -94001
	ACCOUNT_FEATURES_INVALID_DUPLICATE_FEATURE_IDS             ErrorCode = -94003
	ACCOUNT_FEATURES_INVALID_FEATURE_ID_REQUIRED               ErrorCode = -94006
	ACCOUNT_FEATURES_INVALID_MULTIPLE_KEYS                     ErrorCode = -94002
	ACCOUNT_FEATURES_NOT_FOUND                                 ErrorCode = -94004
	ACCOUNT_MFA_ENABLED                                        ErrorCode = -20677
	API_GATEWAY_READ_REQUEST_BODY_FAILED                       ErrorCode = -97300
	APP_ACCOUNT_ALREADY_ACTIVATED                              ErrorCode = -20623
	APP_ACCOUNT_ALREADY_EXISTS                                 ErrorCode = -20603
	APP_ACCOUNT_INACTIVE                                       ErrorCode = -20602
	APP_ACCOUNT_IS_LOCKED                                      ErrorCode = -20661
	APP_ACCOUNT_IS_NOT_BINDED_TO_DEVICE                        ErrorCode = -20580
	APP_ACCOUNT_LOGGED_IN_OTHER_PLACES                         ErrorCode = -20675
	APP_ACCOUNT_NOT_FOUND                                      ErrorCode = -20600
	APP_ACCOUNT_RESEND_EMAIL_EXCEED_LIMIT                      ErrorCode = -21002
	APP_CREDENTIAL_MISMATCH                                    ErrorCode = -20601
	APP_DEVICE_ASSOCIATED_WITH_ANOTHER_ACCOUNT                 ErrorCode = -20506
	APP_DEVICE_ASSOCIATION_LIMIT_EXCEEDED                      ErrorCode = -20508
	APP_DEVICE_ID_AND_FIRMWARE_ID_MISMATCH                     ErrorCode = -20703
	APP_DEVICE_IS_OFFLINE                                      ErrorCode = -20571
	APP_DEVICE_NOT_ASSOCIATED_WITH_ANY_ACCOUNT                 ErrorCode = -20507
	APP_DEVICE_NOT_FOUND                                       ErrorCode = -20501
	APP_EMAIL_ALREADY_EXISTS                                   ErrorCode = -20621
	APP_FIRMWARE_ID_NOT_FOUND                                  ErrorCode = -20505
	APP_HARDWARE_ID_AND_OEM_ID_MISMATCH                        ErrorCode = -20704
	APP_HARDWARE_ID_NOT_FOUND                                  ErrorCode = -20504
	APP_INVAID_USERNAME                                        ErrorCode = -20202
	APP_INVALID_BADGE_VALUE                                    ErrorCode = -20813
	APP_INVALID_DEVICE                                         ErrorCode = -20511
	APP_INVALID_DEVICE_ALIAS                                   ErrorCode = -20572
	APP_INVALID_DEVICE_TOKEN                                   ErrorCode = -20812
	APP_INVALID_EMAIL                                          ErrorCode = -20200
	APP_INVALID_NEW_PASSWORD                                   ErrorCode = -20616
	APP_INVALID_NICKNAME                                       ErrorCode = -20620
	APP_INVALID_PASSWORD                                       ErrorCode = -20615
	APP_INVALID_PHONE                                          ErrorCode = -20201
	APP_NOT_AUTHORIZED                                         ErrorCode = -20509
	APP_NO_SUCH_APP                                            ErrorCode = -23001
	APP_OEM_ID_NOT_FOUND                                       ErrorCode = -20510
	APP_OWNER_ACCOUNT_NOT_FOUND                                ErrorCode = -20617
	APP_PARAMETER_NOT_FOUND                                    ErrorCode = -20104
	APP_PHONE_ALREADY_EXISTS                                   ErrorCode = -20619
	APP_PHONE_NOT_SUPPORTED                                    ErrorCode = -20811
	APP_REFRESH_TOKEN_EXPIRED                                  ErrorCode = -20655
	APP_REFRESH_TOKEN_NOT_FOUND                                ErrorCode = -20656
	APP_REQUEST_TIMEOUT                                        ErrorCode = -20002
	APP_TOKEN_CREDENTIAL_MISMATCH                              ErrorCode = -20652
	APP_TOKEN_EXPIRED                                          ErrorCode = -20651
	APP_USERNAME_ALREADY_EXISTS                                ErrorCode = -20622
	APP_USER_ACCOUNT_NOT_FOUND                                 ErrorCode = -20618
	APP_VERSION_TOO_OLD                                        ErrorCode = -23003
	AUTH_ACCESS_TOKEN_NOT_FOUND                                ErrorCode = -99110
	AUTH_ACCOUNT_COUNTRY_NOT_FOUND                             ErrorCode = -99140
	AUTH_ACCOUNT_NOT_FOUND                                     ErrorCode = -20600
	AUTH_ACCOUNT_SETTING_NOT_FOUND                             ErrorCode = -99126
	AUTH_AUTHORIZATION_CODE_ALREADY_USED                       ErrorCode = -99124
	AUTH_AUTHORIZATION_NOT_FOUND                               ErrorCode = -99111
	AUTH_BACKEND_UNAVAILABLE                                   ErrorCode = -99103
	AUTH_CLIENTID_NOT_SUPPORTED                                ErrorCode = -99133
	AUTH_CLIENT_NOT_FOUND                                      ErrorCode = -99104
	AUTH_CLOUD_SERVICES_NOT_INITIALIZED                        ErrorCode = -99102
	AUTH_CLOUD_SERVICE_TOKEN_NOT_INITIALIZED                   ErrorCode = -99121
	AUTH_CLOUD_TOKEN_NOT_FOUND                                 ErrorCode = -99115
	AUTH_COUNTRY_CODE_INVALID                                  ErrorCode = -99127
	AUTH_DUPLICATE_CLOUDSERVICETOKEN_CREATION                  ErrorCode = -99145
	AUTH_DUPLICATE_USER_CREATION                               ErrorCode = -99144
	AUTH_EXPIRED_AUTHORIZATION_CODE                            ErrorCode = -99125
	AUTH_EXTERNAL_NETWORK_ACCESS_TOKEN_ERROR                   ErrorCode = -99114
	AUTH_EXTERNAL_NETWORK_ACCESS_TOKEN_NOT_FOUND               ErrorCode = -99143
	AUTH_EXTERNAL_NETWORK_ACCESS_TOKEN_URL_NOT_FOUND           ErrorCode = -99142
	AUTH_EXTERNAL_NETWORK_APP_LINK_FAILED                      ErrorCode = -99136
	AUTH_EXTERNAL_NETWORK_INSUFFICIENT_PERMISSIONS             ErrorCode = -99118
	AUTH_EXTERNAL_NETWORK_NOT_LINKED                           ErrorCode = -99113
	AUTH_EXTERNAL_NETWORK_UNAUTHORIZED                         ErrorCode = -99116
	AUTH_FORBIDDEN                                             ErrorCode = -99106
	AUTH_FORBIDDEN_EMAIL_DO_NOT_MATCH                          ErrorCode = -99141
	AUTH_GEOLOCATION_NOT_FOUND                                 ErrorCode = -99129
	AUTH_GRANT_NOT_SUPPORTED                                   ErrorCode = -99123
	AUTH_INTEGRATION_LIMIT_REACHED                             ErrorCode = -99132
	AUTH_INVALID_CLIENT                                        ErrorCode = -99137
	AUTH_INVALID_CLIENT_AND_TOKEN                              ErrorCode = -99105
	AUTH_INVALID_GRANT                                         ErrorCode = -99122
	AUTH_INVALID_STATE                                         ErrorCode = -99138
	AUTH_LOCALE_INVALID                                        ErrorCode = -99134
	AUTH_NETWORK_MISMATCH                                      ErrorCode = -99131
	AUTH_RATE_LIMIT_EXCEEDED                                   ErrorCode = -99119
	AUTH_REDIRECT_URL_EXISTS                                   ErrorCode = -99108
	AUTH_REDIRECT_URL_NOT_FOUND                                ErrorCode = -99109
	AUTH_REGION_NOT_FOUND                                      ErrorCode = -99128
	AUTH_TERMINAL_ASSOCIATED_WITH_OTHER_ACCOUNT                ErrorCode = -99147
	AUTH_TERMINAL_ID_NOT_FOUND                                 ErrorCode = -99117
	AUTH_THIRD_PARTY_INTERNAL_ERROR                            ErrorCode = -99148
	AUTH_TIMEZONE_NOT_FOUND                                    ErrorCode = -99120
	AUTH_UNAUTHORIZED                                          ErrorCode = -20651
	AUTH_UNSUPPORTED_INTERNAL_NETWORK                          ErrorCode = -99130
	AUTH_USER_NOT_FOUND                                        ErrorCode = -99112
	AUTH_USER_PROFILE_ASSOCIATION_EXISTS                       ErrorCode = -99146
	AUTH_WHITELISTING_BLACKLISTED                              ErrorCode = -99139
	AUTH_WHITELISTING_NOT_SUPPORTED                            ErrorCode = -99135
	BO_FEEDBACK_NOT_FOUND                                      ErrorCode = -99301
	BO_NOT_INITIALIZED                                         ErrorCode = -99300
	CACHE_CONNECTION_FAILURE                                   ErrorCode = -99801
	CACHE_GENERAL_FAILURE                                      ErrorCode = -99800
	COMPLIANCE_SERVER_EXECUTION_CONFIG_DUPLICATED              ErrorCode = -96003
	COMPLIANCE_SERVER_EXECUTION_CONFIG_WRONG_TYPE              ErrorCode = -96004
	COMPLIANCE_SERVER_ILLEGAL_SHARD_MESSAGE                    ErrorCode = -96002
	COMPLIANCE_SERVER_MESSAGE_NOT_SUPPORTED                    ErrorCode = -96001
	COMPLIANCE_SERVER_ORCHESTRATION_FAILED                     ErrorCode = -96007
	COMPLIANCE_SERVER_ORCHESTRATION_TIMEOUT                    ErrorCode = -96006
	COMPLIANCE_SERVER_REGION_MISSING                           ErrorCode = -96008
	COMPLIANCE_SERVER_UNKNOW_ORCHESTRATION_RESPONSE            ErrorCode = -96005
	COTURN_IP_NOT_FOUND                                        ErrorCode = -97500
	CS_ACTIVITY_CENTER_DELAYED                                 ErrorCode = -99717
	CS_ACTIVITY_CENTER_NOT_FOUND                               ErrorCode = -99718
	CS_DEVICE_NOT_SUBSCRIBED                                   ErrorCode = -99700
	CS_ELIGIBLE_DEVICE_ALREADY_EXISTING                        ErrorCode = -99715
	CS_ELIGIBLE_DEVICE_NOT_FOUND                               ErrorCode = -99713
	CS_FEATURE_INCOMPLETE                                      ErrorCode = -99719
	CS_FORBIDDEN                                               ErrorCode = -99709
	CS_FREE_CLOUDSTORAGE_DEVICE_NOT_FOUND                      ErrorCode = -99714
	CS_GENERIC_ERROR                                           ErrorCode = -99712
	CS_INVALID_DATA_TYPE                                       ErrorCode = -99703
	CS_INVALID_TIME_RANGE                                      ErrorCode = -99704
	CS_INVALID_URL                                             ErrorCode = -99710
	CS_MEDIA_NOT_FOUND                                         ErrorCode = -99701
	CS_NOT_FOUND                                               ErrorCode = -99705
	CS_OUTDATED_SUBSCRIPTION_EVENT                             ErrorCode = -99708
	CS_PREPARE_HLS_FAILED                                      ErrorCode = -99711
	CS_PREPARE_VIDEO_FILE_FAILED                               ErrorCode = -99706
	CS_QUOTA_NOT_FOUND                                         ErrorCode = -99716
	CS_STORAGE_PLAN_NOT_FOUND                                  ErrorCode = -99702
	CS_VIDEO_NOT_READY                                         ErrorCode = -99707
	DB_DAO_NOT_INITIALIZED                                     ErrorCode = -99401
	DB_GENERAL_ERROR                                           ErrorCode = -99400
	DB_OBJECT_ALREADY_EXISTING                                 ErrorCode = -99403
	DB_OBJECT_LIMIT_REACHED                                    ErrorCode = -99404
	DB_OBJECT_NOT_FOUND                                        ErrorCode = -99402
	DC_GENERIC_ERROR                                           ErrorCode = -99900
	DC_INVALID_DEVICE_STATE_VALUE                              ErrorCode = -99902
	DEVICE_CACHE_BUILT_IN_PROGRESS                             ErrorCode = -99101
	DEVICE_CACHE_DEVICE_NOT_FOUND                              ErrorCode = -99100
	DEVICE_CACHE_FAILED_TO_RETRIEVE                            ErrorCode = -99102
	DEVICE_CACHE_NOT_INITIALIZED                               ErrorCode = -99103
	DEVICE_META_MODULE_NOT_FOUND                               ErrorCode = -91100
	DEV_ERROR_IN_USE_BY_OTHER_CLIENT                           ErrorCode = -44
	DEV_OWNERSHIP_VIOLATION                                    ErrorCode = -99920
	DEV_ROLE_UPDATE_NOT_SUPPORTED                              ErrorCode = -99921
	DEV_TOKEN_ERROR                                            ErrorCode = -25002
	DEV_TOKEN_EXPIRED                                          ErrorCode = -25001
	DG_GENERIC_ERROR                                           ErrorCode = -99950
	DG_INVALID_ACTION                                          ErrorCode = -99951
	DG_UNSUPPORTED_DEVICE_FILTER                               ErrorCode = -99952
	DIRECTED_CUSTOMER_ACCOUNT_TRANSFER_IN_PROGRESS             ErrorCode = -96102
	DIRECTED_CUSTOMER_EMAIL_TAKEN                              ErrorCode = -96100
	DIRECTED_CUSTOMER_NETWORK_DCID_EXISTS                      ErrorCode = -96101
	EC_SYNC_INVENTORY_IN_PROGRESS                              ErrorCode = -98100
	ES_GENERIC_ERROR                                           ErrorCode = -97400
	ES_REQUEST_FAILED                                          ErrorCode = -97401
	ES_REQUEST_NOT_SUPPORTED                                   ErrorCode = -97402
	FEATURE_REGISTRY_EMPTY_FEATURE_LIST                        ErrorCode = -96202
	FEATURE_REGISTRY_FEATURE_NOT_FOUND                         ErrorCode = -96201
	FEATURE_REGISTRY_GENERIC_ERROR                             ErrorCode = -96200
	FFS_GENERAL_EXCEPTION                                      ErrorCode = -98600
	FFS_SESSION_NOT_FOUND                                      ErrorCode = -98601
	GEOFENCE_FAIL_TO_PROCESS_USER_PLACE                        ErrorCode = -92003
	GEOFENCE_FAIL_TO_PROCESS_USER_PROFILE                      ErrorCode = -92002
	GEOFENCE_LAST_KNOWN_PLACE_INFO_OUT_OF_DATE                 ErrorCode = -92007
	GEOFENCE_USER_PLACE_NOT_FOUND                              ErrorCode = -92001
	GEOFENCE_USER_PLACE_REQUEST_INVALID                        ErrorCode = -92008
	GEOFENCE_USER_PLACE_WEBHOOK_INVALID                        ErrorCode = -92005
	GEOFENCE_USER_PROFILE_CONTAINS_UNKNOWN_PLACEID             ErrorCode = -92009
	GEOFENCE_USER_PROFILE_NOT_FOUND                            ErrorCode = -92000
	GEOFENCE_USER_PROFILE_WEBHOOK_INVALID                      ErrorCode = -92004
	GEOFENCE_USER_PROFLE_REQUEST_INVALID                       ErrorCode = -92006
	HK_APPLE_AUTH_ENTITY_ILLEGAL_STATE                         ErrorCode = -98025
	HK_APPLE_AUTH_ENTITY_IN_EXPECTED_STATE_ALREADY             ErrorCode = -98027
	HK_APPLE_BAD_REQUEST                                       ErrorCode = -98023
	HK_APPLE_BAD_URL                                           ErrorCode = -98012
	HK_APPLE_DUPLICATE_AUTH_ENTITY                             ErrorCode = -98028
	HK_APPLE_ENTITY_DOWNLOADS_NOT_AVAILABLE                    ErrorCode = -98018
	HK_APPLE_ENTITY_GENERATION_IN_PROGRESS                     ErrorCode = -98009
	HK_APPLE_EXCEEDED_MAX_PAYLOAD_COUNT                        ErrorCode = -98029
	HK_APPLE_HTTP_METHOD_NOT_ALLOWED                           ErrorCode = -98013
	HK_APPLE_INSUFFICIENT_AUTH_ENTITIES                        ErrorCode = -98017
	HK_APPLE_INTERNAL_SERVER_ERROR                             ErrorCode = -98022
	HK_APPLE_INVALID_AUTH_ENTITY                               ErrorCode = -98019
	HK_APPLE_INVALID_AUTH_ENTITY_TYPE                          ErrorCode = -98026
	HK_APPLE_INVALID_OR_UNSUPPORTED_PPID                       ErrorCode = -98021
	HK_APPLE_INVALID_PPID                                      ErrorCode = -98016
	HK_APPLE_INVALID_REQUEST_ID                                ErrorCode = -98015
	HK_APPLE_JSON_ERROR                                        ErrorCode = -98008
	HK_APPLE_MISMATCHED_PLAN_ID                                ErrorCode = -98020
	HK_APPLE_MISMATCHED_UUID                                   ErrorCode = -98024
	HK_APPLE_MISSING_MANDATORY_PARAMETERS                      ErrorCode = -98030
	HK_APPLE_SERVICE_UNAVAILABLE                               ErrorCode = -98010
	HK_APPLE_UNAUTHORIZED_USER                                 ErrorCode = -98011
	HK_APPLE_UNKNOWN_SERVER_ERROR                              ErrorCode = -98014
	HK_ERROR_DEVICE_MODEL_NOT_SUPPORTED                        ErrorCode = -98005
	HK_ERROR_DOWNLOAD_TOKENS                                   ErrorCode = -98003
	HK_ERROR_FILE_PARSE                                        ErrorCode = -98002
	HK_ERROR_HARDWARE_VERSION_NOT_SUPPORTED                    ErrorCode = -98004
	HK_ERROR_TOKENS_NOT_AVAILABLE                              ErrorCode = -98006
	HK_PREVIOUS_RETRIEVE_TOKENS_IS_BEING_PROCESSED             ErrorCode = -98032
	HK_PRODUCT_PLAN_EXISTS                                     ErrorCode = -98000
	HK_PRODUCT_PLAN_NOT_FOUND                                  ErrorCode = -98001
	HK_REGISTRATION_IN_PROGRESS                                ErrorCode = -98035
	HK_REQUEST_TOKEN_NOT_FOUND                                 ErrorCode = -98007
	HK_TOKENS_NOT_YET_READY_FOR_DOWNLOAD                       ErrorCode = -98033
	HK_TOKEN_REQUEST_NOT_APPLICABLE_FOR_FACTORY                ErrorCode = -98034
	HK_UNDOCUMENTED_APPLE_ERROR                                ErrorCode = -98031
	INCORRECT_APP_SERVER_URL                                   ErrorCode = -20212
	INCORRECT_VERIFICATION_CODE                                ErrorCode = -20607
	IOT_ACTIVITY_NOT_FOUND                                     ErrorCode = -99032
	IOT_ASSERTION_EXCEPTION                                    ErrorCode = -99020
	IOT_BACKEND_SERVICE_UNAVAILABLE                            ErrorCode = -99066
	IOT_CACHE_OUT_OF_SERVICE                                   ErrorCode = -99026
	IOT_CAPABILITY_NOT_SUPPORTED                               ErrorCode = -99001
	IOT_CIRCUIT_BREAKER_OPEN                                   ErrorCode = -99067
	IOT_DEPENDENT_SERVICE_NOT_INITIALIZED                      ErrorCode = -99007
	IOT_DESERIALIZATION_EXCEPTION                              ErrorCode = -99011
	IOT_DEVICE_ASSOCIATED_WITH_ANOTHER_PARENT_DEVICE           ErrorCode = -99047
	IOT_DEVICE_ERROR                                           ErrorCode = -99024
	IOT_DEVICE_LOGS_NOT_FOUND                                  ErrorCode = -99045
	IOT_DEVICE_REGION_ERROR                                    ErrorCode = -99030
	IOT_EVENT_NOT_SUPPORTED_EXCEPTION                          ErrorCode = -99043
	IOT_EVENT_RUNTIME_EXCEPTION                                ErrorCode = -99042
	IOT_EXTERNAL_NETWORK_SERVICE_DOES_NOT_EXIST                ErrorCode = -99029
	IOT_EXTERNAL_NETWORK_SERVICE_NOT_AVAILABLE                 ErrorCode = -99028
	IOT_FAILED_TO_INITIALIZE_CAPABILITIES                      ErrorCode = -99005
	IOT_FIRMWARE_UPGRADE_FAILED                                ErrorCode = -99014
	IOT_FIRMWARE_UPGRADE_IN_PROGRESS                           ErrorCode = -99015
	IOT_GENERAL_EXCEPTION                                      ErrorCode = -99000
	IOT_GENERAL_SCENE_EXCEPTION                                ErrorCode = -99018
	IOT_GOOGLE_WEAVE_MODEL_MANIFEST_ALREADY_EXISTS             ErrorCode = -99022
	IOT_GOOGLE_WEAVE_MODEL_MANIFEST_NOT_FOUND                  ErrorCode = -99023
	IOT_INITIALIZATION_FAILED                                  ErrorCode = -99003
	IOT_INTEGRATION_EXCEPTION                                  ErrorCode = -99016
	IOT_INVALID_CONFIGURATION                                  ErrorCode = -99009
	IOT_INVALID_EVENT_EXCEPTION                                ErrorCode = -99041
	IOT_INVALID_MATCHER_EXCLUDE_PATH                           ErrorCode = -99013
	IOT_INVALID_TIMEZONE_ID                                    ErrorCode = -99031
	IOT_IO_EXCEPTION                                           ErrorCode = -99006
	IOT_JWT_REFRESH_TOKEN_DECODE_FAILED                        ErrorCode = -99057
	IOT_JWT_REFRESH_TOKEN_GENERATE_FAILED                      ErrorCode = -99055
	IOT_JWT_TOKEN_DECODE_FAILED                                ErrorCode = -99056
	IOT_JWT_TOKEN_GENERATE_FAILED                              ErrorCode = -99054
	IOT_JWT_TOKEN_INVALID                                      ErrorCode = -99053
	IOT_KINESIS_STREAM_PRODUCER_INIT_FAILED                    ErrorCode = -99046
	IOT_LOCATION_NOT_LINKED                                    ErrorCode = -99027
	IOT_MESSAGEBROKER_INITIALIZATION_ERROR                     ErrorCode = -99017
	IOT_OUT_OF_TIME_EXCEPTION                                  ErrorCode = -99038
	IOT_REGION_ENDPOINT_NOT_AVAILABLE                          ErrorCode = -99048
	IOT_REQUEST_VERSION_NOT_SUPPORTED                          ErrorCode = -99052
	IOT_ROUTER_RULE_NOT_FOUND                                  ErrorCode = -99025
	IOT_SCENE_DEVICE_NOT_FOUND                                 ErrorCode = -99019
	IOT_SCENE_NOT_FOUND                                        ErrorCode = -99012
	IOT_SERIALIZATION_EXCEPTION                                ErrorCode = -99010
	IOT_SHARED_VIDEOS_NOT_AVAILABLE                            ErrorCode = -99060
	IOT_SSL_INITIALIZATION_ERROR                               ErrorCode = -99021
	IOT_SYSTEM_SHUTDOWN_INITIATED                              ErrorCode = -99061
	IOT_TECHNICAL_EXCEPTION                                    ErrorCode = -99008
	IOT_TIMEZONE_NOT_FOUND                                     ErrorCode = -99002
	IOT_TRY_LOCK_TIMEOUT                                       ErrorCode = -99050
	IOT_UNIDENTIFIED_JWT_TOKEN_ISSUER                          ErrorCode = -99051
	IOT_UNKNOWN_DEVICE_TYPE                                    ErrorCode = -99004
	IOT_UNRECOGNIZED_VIDEO_ANALYTICS_CLASSIFICATION            ErrorCode = -99049
	K8S_CLIENT_INIT_WRONG_TYPE                                 ErrorCode = -95001
	K8S_CLIENT_REQUEST_ERROR                                   ErrorCode = -95002
	KC_GENERIC_ERROR                                           ErrorCode = -99200
	LIGHTING_EFFECTS_NOT_FOUND                                 ErrorCode = -98700
	LIGHTING_EFFECTS_PREDEFINED_EFFECT_ALREADY_EXISTS          ErrorCode = -98702
	LIGHTING_EFFECTS_PREDEFINED_EFFECT_TEMPLATE_ALREADY_EXISTS ErrorCode = -98704
	LIGHTING_EFFECTS_PREDEFINED_EFFECT_TEMPLATE_NOT_FOUND      ErrorCode = -98703
	LIGHTING_EFFECTS_VALIDATION_FAILED                         ErrorCode = -98701
	MC_METHOD_NOT_ALLOWED                                      ErrorCode = -98202
	MC_RESOURCE_NOT_FOUND                                      ErrorCode = -98200
	MC_UNDOCUMENTED_ERROR                                      ErrorCode = -98201
	METRICS_PROVIDER_INVALID                                   ErrorCode = -98802
	METRICS_PROVIDER_NOT_FOUND                                 ErrorCode = -98801
	METRICS_PROVIDER_NOT_REGISTERED                            ErrorCode = -98800
	MFA_PROCESS_CLOSED                                         ErrorCode = -20683
	MFA_PROCESS_EXPIRED                                        ErrorCode = -20681
	PC_SEVER_ERROR                                             ErrorCode = -98301
	PC_SYSTEM_UPDATING                                         ErrorCode = -98300
	PC_UNDOCUMENTED_ERROR                                      ErrorCode = -98302
	REQUEST_TOO_MUCH_CODE                                      ErrorCode = -20662
	SA_CONDITION_NOT_EVALUATED                                 ErrorCode = -98911
	SA_EXCEPTION                                               ErrorCode = -98910
	SA_EXECUTION_ALREADY_IN_PROGRESS                           ErrorCode = -98903
	SA_EXECUTION_PLAN_GENERATION_FAILED                        ErrorCode = -98906
	SA_EXECUTION_PLAN_NOT_AVAILABLE                            ErrorCode = -98905
	SA_FAILED_TO_SATISFY_CONDITION                             ErrorCode = -98908
	SA_GEO_LOCATION_NOT_FOUND                                  ErrorCode = -99129
	SA_INVALID_CONDITION                                       ErrorCode = -98907
	SA_RANGE_KEY_MISSING                                       ErrorCode = -98901
	SA_RECOVERING_FAILED_EXECUTION                             ErrorCode = -98904
	SA_RULES_QUOTA_EXCEEDED                                    ErrorCode = -98900
	SA_SKIP_BY_CONFIG                                          ErrorCode = -98909
	SA_SMART_ACTION_NOT_INITIALIZED                            ErrorCode = -98902
	SA_TIMEZONE_NOT_FOUND                                      ErrorCode = -99120
	SC_MODULE_NOT_REGISTERED                                   ErrorCode = -97201
	SC_PAYLOAD_NOT_FOUND                                       ErrorCode = -97202
	SC_SCHEDULE_NOT_FOUND                                      ErrorCode = -97200
	SGW_ACCOUNT_ALREADY_EXISTS                                 ErrorCode = -99501
	SGW_ACCOUNT_ALREADY_INACTIVE                               ErrorCode = -99659
	SGW_ACCOUNT_BAD_REQUEST                                    ErrorCode = -99202
	SGW_ACCOUNT_BILLING_INFO_REQUIRED                          ErrorCode = -99512
	SGW_ACCOUNT_CLOSED                                         ErrorCode = -99635
	SGW_ACCOUNT_IMMUTABLE_SUBSCRIPTION                         ErrorCode = -99204
	SGW_ACCOUNT_INTERNAL_SERVER_ERROR                          ErrorCode = -99203
	SGW_ACCOUNT_INVALID_API_KEY                                ErrorCode = -99205
	SGW_ACCOUNT_INVALID_API_VERSION                            ErrorCode = -99206
	SGW_ACCOUNT_INVALID_CONFIGURATION                          ErrorCode = -99507
	SGW_ACCOUNT_INVALID_CONTENT_TYPE                           ErrorCode = -99207
	SGW_ACCOUNT_INVALID_DATA                                   ErrorCode = -99502
	SGW_ACCOUNT_INVALID_PERMISSIONS                            ErrorCode = -9208
	SGW_ACCOUNT_INVALID_TOKEN                                  ErrorCode = -99209
	SGW_ACCOUNT_INVALID_TRANSITION                             ErrorCode = -99511
	SGW_ACCOUNT_MANDATORY_FIELD                                ErrorCode = -99506
	SGW_ACCOUNT_MISSING_FEATURE                                ErrorCode = -99216
	SGW_ACCOUNT_NOT_CLOSED                                     ErrorCode = -99636
	SGW_ACCOUNT_NOT_FOUND                                      ErrorCode = -99504
	SGW_ACCOUNT_RATE_LIMITED                                   ErrorCode = -99217
	SGW_ACCOUNT_SIMULTANEOUS_REQUEST                           ErrorCode = -99210
	SGW_ACCOUNT_STATE_INVALID                                  ErrorCode = -99505
	SGW_ACCOUNT_TRANSACTION                                    ErrorCode = -99211
	SGW_ACCOUNT_UNACCEPTABLE_VALUE                             ErrorCode = -99508
	SGW_ACCOUNT_UNAUTHORIZED                                   ErrorCode = -99212
	SGW_ACCOUNT_UNAVAILABLE_IN_API_VERSION                     ErrorCode = -99213
	SGW_ACCOUNT_UNKNOWN_API_VERSION                            ErrorCode = -99214
	SGW_ACCOUNT_VALIDATION                                     ErrorCode = -99215
	SGW_ACCOUNT_VALUE_NOT_A_NUMBER                             ErrorCode = -99510
	SGW_ACCOUNT_VALUE_NOT_INCLUDED_IN_LIST                     ErrorCode = -99509
	SGW_ACCOUNT_VERSION_NOT_FOUND                              ErrorCode = -99655
	SGW_BILLING_INFO_ALREADY_EXISTS                            ErrorCode = -99544
	SGW_BILLING_INFO_BAD_REQUEST                               ErrorCode = -99250
	SGW_BILLING_INFO_CONFIGURATION                             ErrorCode = -99548
	SGW_BILLING_INFO_IMMUTABLE_SUBSCRIPTION                    ErrorCode = -99252
	SGW_BILLING_INFO_INTERNAL_SERVER_ERROR                     ErrorCode = -99251
	SGW_BILLING_INFO_INVALID_API_KEY                           ErrorCode = -99253
	SGW_BILLING_INFO_INVALID_API_VERSION                       ErrorCode = -99254
	SGW_BILLING_INFO_INVALID_CONTENT_TYPE                      ErrorCode = -99255
	SGW_BILLING_INFO_INVALID_DATA                              ErrorCode = -99545
	SGW_BILLING_INFO_INVALID_PERMISSIONS                       ErrorCode = -99256
	SGW_BILLING_INFO_INVALID_TOKEN                             ErrorCode = -99257
	SGW_BILLING_INFO_MANDATORY_FIELD                           ErrorCode = -99546
	SGW_BILLING_INFO_MISSING_FEATURE                           ErrorCode = -99265
	SGW_BILLING_INFO_NOT_FOUND                                 ErrorCode = -99543
	SGW_BILLING_INFO_RATE_LIMITED                              ErrorCode = -99266
	SGW_BILLING_INFO_REQUIRED                                  ErrorCode = -99542
	SGW_BILLING_INFO_SIMULTANEOUS_REQUEST                      ErrorCode = -99258
	SGW_BILLING_INFO_TOKEN_INVALID                             ErrorCode = -99541
	SGW_BILLING_INFO_TRANSACTION                               ErrorCode = -99260
	SGW_BILLING_INFO_UNAUTHORIZED                              ErrorCode = -99261
	SGW_BILLING_INFO_UNAVAILABLE_IN_API_VERSION                ErrorCode = -99262
	SGW_BILLING_INFO_UNKNOWN_API_VERSION                       ErrorCode = -99263
	SGW_BILLING_INFO_UPDATE_DECLINE                            ErrorCode = -99268
	SGW_BILLING_INFO_VALIDATION                                ErrorCode = -99264
	SGW_BILLING_INFO_VALUE_NOT_A_NUMBER                        ErrorCode = -99547
	SGW_COUPON_NOT_FOUND                                       ErrorCode = -99633
	SGW_COUPON_NOT_REDEEMABLE                                  ErrorCode = -99634
	SGW_DEVICES_ALREADY_IN_MAX_BIND_NUMBER                     ErrorCode = -99666
	SGW_DEVICE_ALREADY_IN_FREE_TRIAL                           ErrorCode = -99657
	SGW_DEVICE_ALREADY_IN_PREMIUM_SUBSCRIPTION                 ErrorCode = -99665
	SGW_FREE_TRIAL_PERIOD_INACTIVE                             ErrorCode = -99656
	SGW_GENERIC_ERROR                                          ErrorCode = -99500
	SGW_INVALID_EMAIL                                          ErrorCode = -99513
	SGW_INVALID_PAYMENT_TOKEN                                  ErrorCode = -99526
	SGW_INVALID_TRANSACTION                                    ErrorCode = -99530
	SGW_JOB_IN_PROGRESS                                        ErrorCode = -99650
	SGW_NOT_SUPPORTED                                          ErrorCode = -99651
	SGW_PAYMENT_ACH_TRANSACTIONS_NOT_SUPPORTED                 ErrorCode = -99610
	SGW_PAYMENT_API_ERROR                                      ErrorCode = -99627
	SGW_PAYMENT_APPROVED                                       ErrorCode = -99549
	SGW_PAYMENT_APPROVED_FRAUD_REVIEW                          ErrorCode = -99550
	SGW_PAYMENT_AUTHORIZATION_ALREADY_CAPTURED                 ErrorCode = -99623
	SGW_PAYMENT_AUTHORIZATION_AMOUNT_DEPLETED                  ErrorCode = -99624
	SGW_PAYMENT_AUTHORIZATION_EXPIRED                          ErrorCode = -99622
	SGW_PAYMENT_CALL_ISSUER                                    ErrorCode = -99555
	SGW_PAYMENT_CALL_ISSUER_UPDATE_CARDHOLDER_DATA             ErrorCode = -99556
	SGW_PAYMENT_CANNOT_REFUND_UNSETTLED_TRANSACTIONS           ErrorCode = -99618
	SGW_PAYMENT_CANNOT_VOID_PAYMENT_AUTHORIZATION              ErrorCode = -99616
	SGW_PAYMENT_CARDHOLDER_REQUESTED_STOP                      ErrorCode = -99579
	SGW_PAYMENT_CARD_NOT_ACTIVATED                             ErrorCode = -99631
	SGW_PAYMENT_CARD_TYPE_NOT_ACCEPTED                         ErrorCode = -99574
	SGW_PAYMENT_CONTACT_GATEWAY                                ErrorCode = -99603
	SGW_PAYMENT_CURRENCY_NOT_SUPPORTED                         ErrorCode = -99606
	SGW_PAYMENT_CUSTOMER_CANCELED_TRANSACTION                  ErrorCode = -99578
	SGW_PAYMENT_CVV_REQUIRED                                   ErrorCode = -99605
	SGW_PAYMENT_DECLINED                                       ErrorCode = -99551
	SGW_PAYMENT_DECLINED_CARD_NUMBER                           ErrorCode = -99564
	SGW_PAYMENT_DECLINED_EXCEPTION                             ErrorCode = -99560
	SGW_PAYMENT_DECLINED_EXPIRATION_DATE                       ErrorCode = -99569
	SGW_PAYMENT_DECLINED_MISSING_DATA                          ErrorCode = -99561
	SGW_PAYMENT_DECLINED_SECURITY_CODE                         ErrorCode = -99559
	SGW_PAYMENT_DEPOSIT_REFERENCED_CHARGEBACK                  ErrorCode = -99632
	SGW_PAYMENT_DUPLICATE_TRANSACTION                          ErrorCode = -99628
	SGW_PAYMENT_EXCEEDS_DAILY_LIMIT                            ErrorCode = -99570
	SGW_PAYMENT_EXPIRED_CARD                                   ErrorCode = -99568
	SGW_PAYMENT_FRAUD_ADDRESS                                  ErrorCode = -99584
	SGW_PAYMENT_FRAUD_ADDRESS_RECURLY                          ErrorCode = -99593
	SGW_PAYMENT_FRAUD_ADVANCED_VERIFICATION                    ErrorCode = -99590
	SGW_PAYMENT_FRAUD_GATEWAY                                  ErrorCode = -99588
	SGW_PAYMENT_FRAUD_GENERIC                                  ErrorCode = -99592
	SGW_PAYMENT_FRAUD_IP_ADDRESS                               ErrorCode = -99587
	SGW_PAYMENT_FRAUD_MANUAL_DECISION                          ErrorCode = -99595
	SGW_PAYMENT_FRAUD_RISK_CHECK                               ErrorCode = -99594
	SGW_PAYMENT_FRAUD_SECURITY_CODE                            ErrorCode = -99585
	SGW_PAYMENT_FRAUD_STOLEN_CARD                              ErrorCode = -99586
	SGW_PAYMENT_FRAUD_TOO_MANY_ATTEMPTS                        ErrorCode = -99589
	SGW_PAYMENT_FRAUD_VELOCITY                                 ErrorCode = -99591
	SGW_PAYMENT_GATEWAY_ERROR                                  ErrorCode = -99602
	SGW_PAYMENT_GATEWAY_TIMEOUT                                ErrorCode = -99601
	SGW_PAYMENT_GATEWAY_TOKEN_NOT_FOUND                        ErrorCode = -99567
	SGW_PAYMENT_GATEWAY_UNAVAILABLE                            ErrorCode = -99598
	SGW_PAYMENT_INSUFFICIENT_FUNDS                             ErrorCode = -99552
	SGW_PAYMENT_INVALID_ACCOUNT_NUMBER                         ErrorCode = -99566
	SGW_PAYMENT_INVALID_CARD_NUMBER                            ErrorCode = -99565
	SGW_PAYMENT_INVALID_DATA                                   ErrorCode = -99562
	SGW_PAYMENT_INVALID_EMAIL                                  ErrorCode = -99563
	SGW_PAYMENT_INVALID_GATEWAY_CONFIGURATION                  ErrorCode = -99596
	SGW_PAYMENT_INVALID_ISSUER                                 ErrorCode = -99573
	SGW_PAYMENT_INVALID_LOGIN                                  ErrorCode = -99597
	SGW_PAYMENT_INVALID_MERCHANT_TYPE                          ErrorCode = -99571
	SGW_PAYMENT_INVALID_TRANSACTION                            ErrorCode = -99572
	SGW_PAYMENT_ISSUER_UNAVAILABLE                             ErrorCode = -99600
	SGW_PAYMENT_NO_BILLING_INFORMATION                         ErrorCode = -99580
	SGW_PAYMENT_NO_GATEWAY                                     ErrorCode = -99609
	SGW_PAYMENT_PARTIAL_CREDITS_NOT_SUPPORTED                  ErrorCode = -99617
	SGW_PAYMENT_PAYMENT_NOT_ACCEPTED                           ErrorCode = -99575
	SGW_PAYMENT_PAYPAL_ACCOUNT_ISSUE                           ErrorCode = -99583
	SGW_PAYMENT_PAYPAL_DECLINED_USE_ALTERNATE                  ErrorCode = -99558
	SGW_PAYMENT_PAYPAL_HARD_DECLINE                            ErrorCode = -99582
	SGW_PAYMENT_PAYPAL_INVALID_BILLING_AGREEMENT               ErrorCode = -99581
	SGW_PAYMENT_PAYPAL_PRIMARY_DECLINED                        ErrorCode = -99557
	SGW_PAYMENT_PROCESSOR_UNAVAILABLE                          ErrorCode = -99599
	SGW_PAYMENT_RECURLY_ERROR                                  ErrorCode = -99625
	SGW_PAYMENT_RECURLY_FAILED_TO_GET_TOKEN                    ErrorCode = -99629
	SGW_PAYMENT_RECURLY_TOKEN_NOT_FOUND                        ErrorCode = -99630
	SGW_PAYMENT_RESTRICTED_CARD                                ErrorCode = -99576
	SGW_PAYMENT_RESTRICTED_CARD_CHARGEBACK                     ErrorCode = -99577
	SGW_PAYMENT_SSL_ERROR                                      ErrorCode = -99607
	SGW_PAYMENT_TEMPORARY_HOLD                                 ErrorCode = -99553
	SGW_PAYMENT_THREE_D_SECURE_NOT_SUPPORTED                   ErrorCode = -99611
	SGW_PAYMENT_TOO_MANY_ATTEMPTS                              ErrorCode = -99554
	SGW_PAYMENT_TOTAL_CREDIT_EXCEEDS_CAPTURE                   ErrorCode = -99621
	SGW_PAYMENT_TRANSACTION_ALREADY_VOIDED                     ErrorCode = -99614
	SGW_PAYMENT_TRANSACTION_CANNOT_BE_REFUNDED                 ErrorCode = -99619
	SGW_PAYMENT_TRANSACTION_CANNOT_BE_VOIDED                   ErrorCode = -99620
	SGW_PAYMENT_TRANSACTION_FAILED_TO_SETTLE                   ErrorCode = -99615
	SGW_PAYMENT_TRANSACTION_NOT_FOUND                          ErrorCode = -99612
	SGW_PAYMENT_TRANSACTION_SETTLED                            ErrorCode = -99613
	SGW_PAYMENT_TRY_AGAIN                                      ErrorCode = -99604
	SGW_PAYMENT_UNKNOWN                                        ErrorCode = -99626
	SGW_PAYMENT_ZERO_DOLLAR_AUTH_NOT_SUPPORTED                 ErrorCode = -99608
	SGW_PLAN_ALREADY_EXISTS                                    ErrorCode = -99515
	SGW_PLAN_BAD_REQUEST                                       ErrorCode = -99218
	SGW_PLAN_DATE_IN_PAST                                      ErrorCode = -99522
	SGW_PLAN_END_TIME_INVALID                                  ErrorCode = -99521
	SGW_PLAN_FEATURE_ALREADY_EXISTS                            ErrorCode = -99652
	SGW_PLAN_FEATURE_NOT_FOUND                                 ErrorCode = -99653
	SGW_PLAN_FIELD_CANNOT_BE_BLANK                             ErrorCode = -99518
	SGW_PLAN_IMMUTABLE_SUBSCRIPTION                            ErrorCode = -99220
	SGW_PLAN_INTERNAL_SERVER_ERROR                             ErrorCode = -99219
	SGW_PLAN_INVALID_API_KEY                                   ErrorCode = -99221
	SGW_PLAN_INVALID_API_VERSION                               ErrorCode = -99222
	SGW_PLAN_INVALID_CONFIGURATION                             ErrorCode = -99524
	SGW_PLAN_INVALID_CONTENT_TYPE                              ErrorCode = -99223
	SGW_PLAN_INVALID_DATA                                      ErrorCode = -99516
	SGW_PLAN_INVALID_PERMISSIONS                               ErrorCode = -99224
	SGW_PLAN_INVALID_TOKEN                                     ErrorCode = -99225
	SGW_PLAN_INVALID_TRANSITION                                ErrorCode = -99523
	SGW_PLAN_MANDATORY_FIELD                                   ErrorCode = -99517
	SGW_PLAN_MISSING_FEATURE                                   ErrorCode = -99232
	SGW_PLAN_NOT_FOUND                                         ErrorCode = -99514
	SGW_PLAN_RATE_LIMITED                                      ErrorCode = -99233
	SGW_PLAN_SIMULTANEOUS_REQUEST                              ErrorCode = -99226
	SGW_PLAN_TRANSACTION                                       ErrorCode = -99227
	SGW_PLAN_UNACCEPTABLE_VALUE                                ErrorCode = -99519
	SGW_PLAN_UNAUTHORIZED                                      ErrorCode = -99228
	SGW_PLAN_UNAVAILABLE_IN_API_VERSION                        ErrorCode = -99229
	SGW_PLAN_UNKNOWN_API_VERSION                               ErrorCode = -99230
	SGW_PLAN_VALIDATION                                        ErrorCode = -99231
	SGW_PLAN_VALUE_NOT_A_NUMBER                                ErrorCode = -99520
	SGW_PLAN_VALUE_NOT_INCLUDED_IN_LIST                        ErrorCode = -99525
	SGW_REFUND_FAILED                                          ErrorCode = -99267
	SGW_SESSION_ALREADY_EXISTS                                 ErrorCode = -99654
	SGW_SESSION_NOT_FOUND                                      ErrorCode = -99654
	SGW_SUBSCRIPTION_ACCOUNT_BILLING_INFO_REQUIRED             ErrorCode = -99537
	SGW_SUBSCRIPTION_ALREADY_EXISTS                            ErrorCode = -99528
	SGW_SUBSCRIPTION_BAD_REQUEST                               ErrorCode = -99234
	SGW_SUBSCRIPTION_IMMUTABLE_SUBSCRIPTION                    ErrorCode = -99236
	SGW_SUBSCRIPTION_INACTIVE                                  ErrorCode = -99667
	SGW_SUBSCRIPTION_INTERNAL_SERVER_ERROR                     ErrorCode = -99235
	SGW_SUBSCRIPTION_INVALID_API_KEY                           ErrorCode = -99237
	SGW_SUBSCRIPTION_INVALID_API_VERSION                       ErrorCode = -99238
	SGW_SUBSCRIPTION_INVALID_CONFIGURATION                     ErrorCode = -99535
	SGW_SUBSCRIPTION_INVALID_CONTENT_TYPE                      ErrorCode = -99239
	SGW_SUBSCRIPTION_INVALID_DATA                              ErrorCode = -99531
	SGW_SUBSCRIPTION_INVALID_PERMISSIONS                       ErrorCode = -99240
	SGW_SUBSCRIPTION_INVALID_TOKEN                             ErrorCode = -99241
	SGW_SUBSCRIPTION_INVALID_TRANSITION                        ErrorCode = -99534
	SGW_SUBSCRIPTION_MANDATORY_FIELD                           ErrorCode = -99532
	SGW_SUBSCRIPTION_MISSING_FEATURE                           ErrorCode = -99248
	SGW_SUBSCRIPTION_NOT_FOUND                                 ErrorCode = -99527
	SGW_SUBSCRIPTION_NOT_MODIFIABLE                            ErrorCode = -99529
	SGW_SUBSCRIPTION_NOT_PREMIUM                               ErrorCode = -99664
	SGW_SUBSCRIPTION_RATE_LIMITED                              ErrorCode = -99249
	SGW_SUBSCRIPTION_SIMULTANEOUS_REQUEST                      ErrorCode = -99242
	SGW_SUBSCRIPTION_STATE_INVALID                             ErrorCode = -99540
	SGW_SUBSCRIPTION_TRANSACTION                               ErrorCode = -99243
	SGW_SUBSCRIPTION_TRANSACTION_DECLINE                       ErrorCode = -99269
	SGW_SUBSCRIPTION_TRANSACTION_DECLINE_DUE_TO_3DS            ErrorCode = -99270
	SGW_SUBSCRIPTION_UNAUTHORIZED                              ErrorCode = -99244
	SGW_SUBSCRIPTION_UNAVAILABLE_IN_API_VERSION                ErrorCode = -99245
	SGW_SUBSCRIPTION_UNKNOWN_API_VERSION                       ErrorCode = -99246
	SGW_SUBSCRIPTION_VALIDATION                                ErrorCode = -99247
	SGW_SUBSCRIPTION_VALUES_UNCHANGED                          ErrorCode = -99658
	SGW_SUBSCRIPTION_VALUE_NOT_A_NUMBER                        ErrorCode = -99533
	SGW_SUBSCRIPTION_VALUE_NOT_INCLUDED_IN_LIST                ErrorCode = -99536
	SGW_TAX_ORDER_TRANSACTION_ID_NOT_FOUND                     ErrorCode = -99662
	SGW_TAX_RECORD_ORDER_HAS_BEEN_CREATED                      ErrorCode = -99660
	SGW_TAX_RECORD_REFUND_ORDER_HAS_BEEN_CREATED               ErrorCode = -99661
	SGW_TAX_REFUND_ORDER_REFENRENCE_TRANSACTION_ID_NOT_FOUND   ErrorCode = -99663
	SGW_TGW_BAD_REQUEST                                        ErrorCode = -99639
	SGW_TGW_FORBIDDEN                                          ErrorCode = -99641
	SGW_TGW_GENERIC_EXCEPTION                                  ErrorCode = -99638
	SGW_TGW_GONE                                               ErrorCode = -99645
	SGW_TGW_INTERNAL_SERVER_ERROR                              ErrorCode = -99648
	SGW_TGW_METHOD_NOT_ALLOWED                                 ErrorCode = -99643
	SGW_TGW_NOT_ACCEPTABLE                                     ErrorCode = -99644
	SGW_TGW_NOT_FOUND                                          ErrorCode = -99642
	SGW_TGW_SERVICE_UNAVAILABLE                                ErrorCode = -99649
	SGW_TGW_TOO_MANY_REQUESTS                                  ErrorCode = -99647
	SGW_TGW_UNAUTHORIZED                                       ErrorCode = -99640
	SGW_TGW_UNPROCESSABLE_ENTITY                               ErrorCode = -99646
	SGW_USER_DATA_NOT_DELETED                                  ErrorCode = -99637
	TERMINAL_NOT_BOUND                                         ErrorCode = -23024
	TOO_MUCH_LOGIN_WITH_ONE_CODE                               ErrorCode = -20676
	UNACCEPTABLE_SUBSCRIPTION_PLAN                             ErrorCode = -99201
	URL_ACCOUNT_NOT_MATCH                                      ErrorCode = -97904
	URL_INVALID_ENDPOINT                                       ErrorCode = -97907
	URL_INVALID_TARGET_ERROR                                   ErrorCode = -97901
	URL_IS_USED                                                ErrorCode = -97906
	URL_NOT_FOUND_OR_EXPIRED                                   ErrorCode = -97905
	URL_SHORTURL_EXPIRED                                       ErrorCode = -97903
	URL_SHORTURL_NOT_FOUND                                     ErrorCode = -97902
	USER_PLACE_GEOFENCE_FAILED                                 ErrorCode = -91002
	USER_PLACE_GEOFENCE_SERVER_NOT_AVAILABLE                   ErrorCode = -91003
	USER_PLACE_NOT_FOUND                                       ErrorCode = -91001
	USER_PROFILE_CONTAINS_INVALID_PLACE_ID                     ErrorCode = -93007
	USER_PROFILE_GEOFENCE_FAILED                               ErrorCode = -93003
	USER_PROFILE_GEOFENCE_SERVER_NOT_AVAILABLE                 ErrorCode = -93005
	USER_PROFILE_NOT_FOUND                                     ErrorCode = -93002
	USER_PROFILE_PLACE_ID_VALIDATION_FAILED                    ErrorCode = -93006
	USER_PROFILE_TERMINAL_ASSOCIATION_EXISTS                   ErrorCode = -93001
	USER_PROFILE_TERMINAL_ID_NOT_FOUND                         ErrorCode = -93004
	USER_PROFILE_TERMINAL_NOT_FOUND                            ErrorCode = -93000
	VA_VIDEO_ANALYTICS_NOT_ENABLED                             ErrorCode = -98110
	VA_VIDEO_ANALYTICS_WOWZA_ERROR                             ErrorCode = -98111
	VA_VIDEO_SUMMARY_AUTH_FAILED                               ErrorCode = -98106
	VA_VIDEO_SUMMARY_EVENT_FAILED                              ErrorCode = -98101
	VA_VIDEO_SUMMARY_INVALID_PAGINATOR                         ErrorCode = -98105
	VA_VIDEO_SUMMARY_IN_PROGRESS                               ErrorCode = -98102
	VA_VIDEO_SUMMARY_NOT_ENABLED                               ErrorCode = -98103
	VA_VIDEO_SUMMARY_NOT_GENERATED                             ErrorCode = -98100
	VA_VIDEO_SUMMARY_WRONG_TYPE                                ErrorCode = -98104
	WEBRTC_CLIENT_EXISTS                                       ErrorCode = -97102
	WEBRTC_CLIENT_NOT_FOUND                                    ErrorCode = -97103
	WEBRTC_CONFERENCE_NOT_FOUND                                ErrorCode = -97101
	WEBRTC_CONFERENCE_NOT_INITIALIZED                          ErrorCode = -97100
	WEBRTC_KURENTO_SERVICE_UNAVAILABLE                         ErrorCode = -97104
	WEBRTC_RECORDING_ALREADY                                   ErrorCode = -97106
	WEBRTC_RECORDING_NOT_READY                                 ErrorCode = -97105
	WEBRTC_RECOVERING                                          ErrorCode = -97107
	WEBSOCKET_ACCOUNT_NOT_BOUND                                ErrorCode = -97912
	WEBSOCKET_CONNECTION_INFO_NOT_FOUND                        ErrorCode = -97908
	WEBSOCKET_FAILED_TO_CONNECT_TO_APIGATEWAY                  ErrorCode = -97910
	WEBSOCKET_FAILED_TO_CREATE_CACHE_PROVIDER                  ErrorCode = -97911
	WEBSOCKET_FAILED_TO_NOTIFY_ALL                             ErrorCode = -97914
	WEBSOCKET_FAILED_TO_PROXY                                  ErrorCode = -97913
	WEBSOCKET_PASSTHROUGH_TIMEOUT                              ErrorCode = -97909
	WL_UNKNOWN_ROLE_TYPE                                       ErrorCode = -97000
	WL_UNKNOWN_STATUS_TYPE                                     ErrorCode = -97001
	ZB_UNDOCUMENTED_ERROR                                      ErrorCode = -98500
)

func (c ErrorCode) Names() []string {
	var names []string
	if c == ACCOUNT_FEATURES_DATASET_CORRUPTED {
		names = append(names, "ACCOUNT_FEATURES_DATASET_CORRUPTED")
	}
	if c == ACCOUNT_FEATURES_DEVICE_ADDRESS_REQUIRED {
		names = append(names, "ACCOUNT_FEATURES_DEVICE_ADDRESS_REQUIRED")
	}
	if c == ACCOUNT_FEATURES_DEVICE_CONTEXT_REQUIRED {
		names = append(names, "ACCOUNT_FEATURES_DEVICE_CONTEXT_REQUIRED")
	}
	if c == ACCOUNT_FEATURES_DEVICE_ID_NOT_MATCH {
		names = append(names, "ACCOUNT_FEATURES_DEVICE_ID_NOT_MATCH")
	}
	if c == ACCOUNT_FEATURES_DEVICE_MODEL_REQUIRED {
		names = append(names, "ACCOUNT_FEATURES_DEVICE_MODEL_REQUIRED")
	}
	if c == ACCOUNT_FEATURES_DEVICE_TYPE_REQUIRED {
		names = append(names, "ACCOUNT_FEATURES_DEVICE_TYPE_REQUIRED")
	}
	if c == ACCOUNT_FEATURES_FAILED_DUE_TO_DB_ERRORS {
		names = append(names, "ACCOUNT_FEATURES_FAILED_DUE_TO_DB_ERRORS")
	}
	if c == ACCOUNT_FEATURES_INVALID_DUPLICATE_FEATURE_IDS {
		names = append(names, "ACCOUNT_FEATURES_INVALID_DUPLICATE_FEATURE_IDS")
	}
	if c == ACCOUNT_FEATURES_INVALID_FEATURE_ID_REQUIRED {
		names = append(names, "ACCOUNT_FEATURES_INVALID_FEATURE_ID_REQUIRED")
	}
	if c == ACCOUNT_FEATURES_INVALID_MULTIPLE_KEYS {
		names = append(names, "ACCOUNT_FEATURES_INVALID_MULTIPLE_KEYS")
	}
	if c == ACCOUNT_FEATURES_NOT_FOUND {
		names = append(names, "ACCOUNT_FEATURES_NOT_FOUND")
	}
	if c == ACCOUNT_MFA_ENABLED {
		names = append(names, "ACCOUNT_MFA_ENABLED")
	}
	if c == API_GATEWAY_READ_REQUEST_BODY_FAILED {
		names = append(names, "API_GATEWAY_READ_REQUEST_BODY_FAILED")
	}
	if c == APP_ACCOUNT_ALREADY_ACTIVATED {
		names = append(names, "APP_ACCOUNT_ALREADY_ACTIVATED")
	}
	if c == APP_ACCOUNT_ALREADY_EXISTS {
		names = append(names, "APP_ACCOUNT_ALREADY_EXISTS")
	}
	if c == APP_ACCOUNT_INACTIVE {
		names = append(names, "APP_ACCOUNT_INACTIVE")
	}
	if c == APP_ACCOUNT_IS_LOCKED {
		names = append(names, "APP_ACCOUNT_IS_LOCKED")
	}
	if c == APP_ACCOUNT_IS_NOT_BINDED_TO_DEVICE {
		names = append(names, "APP_ACCOUNT_IS_NOT_BINDED_TO_DEVICE")
	}
	if c == APP_ACCOUNT_LOGGED_IN_OTHER_PLACES {
		names = append(names, "APP_ACCOUNT_LOGGED_IN_OTHER_PLACES")
	}
	if c == APP_ACCOUNT_NOT_FOUND {
		names = append(names, "APP_ACCOUNT_NOT_FOUND")
	}
	if c == APP_ACCOUNT_RESEND_EMAIL_EXCEED_LIMIT {
		names = append(names, "APP_ACCOUNT_RESEND_EMAIL_EXCEED_LIMIT")
	}
	if c == APP_CREDENTIAL_MISMATCH {
		names = append(names, "APP_CREDENTIAL_MISMATCH")
	}
	if c == APP_DEVICE_ASSOCIATED_WITH_ANOTHER_ACCOUNT {
		names = append(names, "APP_DEVICE_ASSOCIATED_WITH_ANOTHER_ACCOUNT")
	}
	if c == APP_DEVICE_ASSOCIATION_LIMIT_EXCEEDED {
		names = append(names, "APP_DEVICE_ASSOCIATION_LIMIT_EXCEEDED")
	}
	if c == APP_DEVICE_ID_AND_FIRMWARE_ID_MISMATCH {
		names = append(names, "APP_DEVICE_ID_AND_FIRMWARE_ID_MISMATCH")
	}
	if c == APP_DEVICE_IS_OFFLINE {
		names = append(names, "APP_DEVICE_IS_OFFLINE")
	}
	if c == APP_DEVICE_NOT_ASSOCIATED_WITH_ANY_ACCOUNT {
		names = append(names, "APP_DEVICE_NOT_ASSOCIATED_WITH_ANY_ACCOUNT")
	}
	if c == APP_DEVICE_NOT_FOUND {
		names = append(names, "APP_DEVICE_NOT_FOUND")
	}
	if c == APP_EMAIL_ALREADY_EXISTS {
		names = append(names, "APP_EMAIL_ALREADY_EXISTS")
	}
	if c == APP_FIRMWARE_ID_NOT_FOUND {
		names = append(names, "APP_FIRMWARE_ID_NOT_FOUND")
	}
	if c == APP_HARDWARE_ID_AND_OEM_ID_MISMATCH {
		names = append(names, "APP_HARDWARE_ID_AND_OEM_ID_MISMATCH")
	}
	if c == APP_HARDWARE_ID_NOT_FOUND {
		names = append(names, "APP_HARDWARE_ID_NOT_FOUND")
	}
	if c == APP_INVAID_USERNAME {
		names = append(names, "APP_INVAID_USERNAME")
	}
	if c == APP_INVALID_BADGE_VALUE {
		names = append(names, "APP_INVALID_BADGE_VALUE")
	}
	if c == APP_INVALID_DEVICE {
		names = append(names, "APP_INVALID_DEVICE")
	}
	if c == APP_INVALID_DEVICE_ALIAS {
		names = append(names, "APP_INVALID_DEVICE_ALIAS")
	}
	if c == APP_INVALID_DEVICE_TOKEN {
		names = append(names, "APP_INVALID_DEVICE_TOKEN")
	}
	if c == APP_INVALID_EMAIL {
		names = append(names, "APP_INVALID_EMAIL")
	}
	if c == APP_INVALID_NEW_PASSWORD {
		names = append(names, "APP_INVALID_NEW_PASSWORD")
	}
	if c == APP_INVALID_NICKNAME {
		names = append(names, "APP_INVALID_NICKNAME")
	}
	if c == APP_INVALID_PASSWORD {
		names = append(names, "APP_INVALID_PASSWORD")
	}
	if c == APP_INVALID_PHONE {
		names = append(names, "APP_INVALID_PHONE")
	}
	if c == APP_NOT_AUTHORIZED {
		names = append(names, "APP_NOT_AUTHORIZED")
	}
	if c == APP_NO_SUCH_APP {
		names = append(names, "APP_NO_SUCH_APP")
	}
	if c == APP_OEM_ID_NOT_FOUND {
		names = append(names, "APP_OEM_ID_NOT_FOUND")
	}
	if c == APP_OWNER_ACCOUNT_NOT_FOUND {
		names = append(names, "APP_OWNER_ACCOUNT_NOT_FOUND")
	}
	if c == APP_PARAMETER_NOT_FOUND {
		names = append(names, "APP_PARAMETER_NOT_FOUND")
	}
	if c == APP_PHONE_ALREADY_EXISTS {
		names = append(names, "APP_PHONE_ALREADY_EXISTS")
	}
	if c == APP_PHONE_NOT_SUPPORTED {
		names = append(names, "APP_PHONE_NOT_SUPPORTED")
	}
	if c == APP_REFRESH_TOKEN_EXPIRED {
		names = append(names, "APP_REFRESH_TOKEN_EXPIRED")
	}
	if c == APP_REFRESH_TOKEN_NOT_FOUND {
		names = append(names, "APP_REFRESH_TOKEN_NOT_FOUND")
	}
	if c == APP_REQUEST_TIMEOUT {
		names = append(names, "APP_REQUEST_TIMEOUT")
	}
	if c == APP_TOKEN_CREDENTIAL_MISMATCH {
		names = append(names, "APP_TOKEN_CREDENTIAL_MISMATCH")
	}
	if c == APP_TOKEN_EXPIRED {
		names = append(names, "APP_TOKEN_EXPIRED")
	}
	if c == APP_USERNAME_ALREADY_EXISTS {
		names = append(names, "APP_USERNAME_ALREADY_EXISTS")
	}
	if c == APP_USER_ACCOUNT_NOT_FOUND {
		names = append(names, "APP_USER_ACCOUNT_NOT_FOUND")
	}
	if c == APP_VERSION_TOO_OLD {
		names = append(names, "APP_VERSION_TOO_OLD")
	}
	if c == AUTH_ACCESS_TOKEN_NOT_FOUND {
		names = append(names, "AUTH_ACCESS_TOKEN_NOT_FOUND")
	}
	if c == AUTH_ACCOUNT_COUNTRY_NOT_FOUND {
		names = append(names, "AUTH_ACCOUNT_COUNTRY_NOT_FOUND")
	}
	if c == AUTH_ACCOUNT_NOT_FOUND {
		names = append(names, "AUTH_ACCOUNT_NOT_FOUND")
	}
	if c == AUTH_ACCOUNT_SETTING_NOT_FOUND {
		names = append(names, "AUTH_ACCOUNT_SETTING_NOT_FOUND")
	}
	if c == AUTH_AUTHORIZATION_CODE_ALREADY_USED {
		names = append(names, "AUTH_AUTHORIZATION_CODE_ALREADY_USED")
	}
	if c == AUTH_AUTHORIZATION_NOT_FOUND {
		names = append(names, "AUTH_AUTHORIZATION_NOT_FOUND")
	}
	if c == AUTH_BACKEND_UNAVAILABLE {
		names = append(names, "AUTH_BACKEND_UNAVAILABLE")
	}
	if c == AUTH_CLIENTID_NOT_SUPPORTED {
		names = append(names, "AUTH_CLIENTID_NOT_SUPPORTED")
	}
	if c == AUTH_CLIENT_NOT_FOUND {
		names = append(names, "AUTH_CLIENT_NOT_FOUND")
	}
	if c == AUTH_CLOUD_SERVICES_NOT_INITIALIZED {
		names = append(names, "AUTH_CLOUD_SERVICES_NOT_INITIALIZED")
	}
	if c == AUTH_CLOUD_SERVICE_TOKEN_NOT_INITIALIZED {
		names = append(names, "AUTH_CLOUD_SERVICE_TOKEN_NOT_INITIALIZED")
	}
	if c == AUTH_CLOUD_TOKEN_NOT_FOUND {
		names = append(names, "AUTH_CLOUD_TOKEN_NOT_FOUND")
	}
	if c == AUTH_COUNTRY_CODE_INVALID {
		names = append(names, "AUTH_COUNTRY_CODE_INVALID")
	}
	if c == AUTH_DUPLICATE_CLOUDSERVICETOKEN_CREATION {
		names = append(names, "AUTH_DUPLICATE_CLOUDSERVICETOKEN_CREATION")
	}
	if c == AUTH_DUPLICATE_USER_CREATION {
		names = append(names, "AUTH_DUPLICATE_USER_CREATION")
	}
	if c == AUTH_EXPIRED_AUTHORIZATION_CODE {
		names = append(names, "AUTH_EXPIRED_AUTHORIZATION_CODE")
	}
	if c == AUTH_EXTERNAL_NETWORK_ACCESS_TOKEN_ERROR {
		names = append(names, "AUTH_EXTERNAL_NETWORK_ACCESS_TOKEN_ERROR")
	}
	if c == AUTH_EXTERNAL_NETWORK_ACCESS_TOKEN_NOT_FOUND {
		names = append(names, "AUTH_EXTERNAL_NETWORK_ACCESS_TOKEN_NOT_FOUND")
	}
	if c == AUTH_EXTERNAL_NETWORK_ACCESS_TOKEN_URL_NOT_FOUND {
		names = append(names, "AUTH_EXTERNAL_NETWORK_ACCESS_TOKEN_URL_NOT_FOUND")
	}
	if c == AUTH_EXTERNAL_NETWORK_APP_LINK_FAILED {
		names = append(names, "AUTH_EXTERNAL_NETWORK_APP_LINK_FAILED")
	}
	if c == AUTH_EXTERNAL_NETWORK_INSUFFICIENT_PERMISSIONS {
		names = append(names, "AUTH_EXTERNAL_NETWORK_INSUFFICIENT_PERMISSIONS")
	}
	if c == AUTH_EXTERNAL_NETWORK_NOT_LINKED {
		names = append(names, "AUTH_EXTERNAL_NETWORK_NOT_LINKED")
	}
	if c == AUTH_EXTERNAL_NETWORK_UNAUTHORIZED {
		names = append(names, "AUTH_EXTERNAL_NETWORK_UNAUTHORIZED")
	}
	if c == AUTH_FORBIDDEN {
		names = append(names, "AUTH_FORBIDDEN")
	}
	if c == AUTH_FORBIDDEN_EMAIL_DO_NOT_MATCH {
		names = append(names, "AUTH_FORBIDDEN_EMAIL_DO_NOT_MATCH")
	}
	if c == AUTH_GEOLOCATION_NOT_FOUND {
		names = append(names, "AUTH_GEOLOCATION_NOT_FOUND")
	}
	if c == AUTH_GRANT_NOT_SUPPORTED {
		names = append(names, "AUTH_GRANT_NOT_SUPPORTED")
	}
	if c == AUTH_INTEGRATION_LIMIT_REACHED {
		names = append(names, "AUTH_INTEGRATION_LIMIT_REACHED")
	}
	if c == AUTH_INVALID_CLIENT {
		names = append(names, "AUTH_INVALID_CLIENT")
	}
	if c == AUTH_INVALID_CLIENT_AND_TOKEN {
		names = append(names, "AUTH_INVALID_CLIENT_AND_TOKEN")
	}
	if c == AUTH_INVALID_GRANT {
		names = append(names, "AUTH_INVALID_GRANT")
	}
	if c == AUTH_INVALID_STATE {
		names = append(names, "AUTH_INVALID_STATE")
	}
	if c == AUTH_LOCALE_INVALID {
		names = append(names, "AUTH_LOCALE_INVALID")
	}
	if c == AUTH_NETWORK_MISMATCH {
		names = append(names, "AUTH_NETWORK_MISMATCH")
	}
	if c == AUTH_RATE_LIMIT_EXCEEDED {
		names = append(names, "AUTH_RATE_LIMIT_EXCEEDED")
	}
	if c == AUTH_REDIRECT_URL_EXISTS {
		names = append(names, "AUTH_REDIRECT_URL_EXISTS")
	}
	if c == AUTH_REDIRECT_URL_NOT_FOUND {
		names = append(names, "AUTH_REDIRECT_URL_NOT_FOUND")
	}
	if c == AUTH_REGION_NOT_FOUND {
		names = append(names, "AUTH_REGION_NOT_FOUND")
	}
	if c == AUTH_TERMINAL_ASSOCIATED_WITH_OTHER_ACCOUNT {
		names = append(names, "AUTH_TERMINAL_ASSOCIATED_WITH_OTHER_ACCOUNT")
	}
	if c == AUTH_TERMINAL_ID_NOT_FOUND {
		names = append(names, "AUTH_TERMINAL_ID_NOT_FOUND")
	}
	if c == AUTH_THIRD_PARTY_INTERNAL_ERROR {
		names = append(names, "AUTH_THIRD_PARTY_INTERNAL_ERROR")
	}
	if c == AUTH_TIMEZONE_NOT_FOUND {
		names = append(names, "AUTH_TIMEZONE_NOT_FOUND")
	}
	if c == AUTH_UNAUTHORIZED {
		names = append(names, "AUTH_UNAUTHORIZED")
	}
	if c == AUTH_UNSUPPORTED_INTERNAL_NETWORK {
		names = append(names, "AUTH_UNSUPPORTED_INTERNAL_NETWORK")
	}
	if c == AUTH_USER_NOT_FOUND {
		names = append(names, "AUTH_USER_NOT_FOUND")
	}
	if c == AUTH_USER_PROFILE_ASSOCIATION_EXISTS {
		names = append(names, "AUTH_USER_PROFILE_ASSOCIATION_EXISTS")
	}
	if c == AUTH_WHITELISTING_BLACKLISTED {
		names = append(names, "AUTH_WHITELISTING_BLACKLISTED")
	}
	if c == AUTH_WHITELISTING_NOT_SUPPORTED {
		names = append(names, "AUTH_WHITELISTING_NOT_SUPPORTED")
	}
	if c == BO_FEEDBACK_NOT_FOUND {
		names = append(names, "BO_FEEDBACK_NOT_FOUND")
	}
	if c == BO_NOT_INITIALIZED {
		names = append(names, "BO_NOT_INITIALIZED")
	}
	if c == CACHE_CONNECTION_FAILURE {
		names = append(names, "CACHE_CONNECTION_FAILURE")
	}
	if c == CACHE_GENERAL_FAILURE {
		names = append(names, "CACHE_GENERAL_FAILURE")
	}
	if c == COMPLIANCE_SERVER_EXECUTION_CONFIG_DUPLICATED {
		names = append(names, "COMPLIANCE_SERVER_EXECUTION_CONFIG_DUPLICATED")
	}
	if c == COMPLIANCE_SERVER_EXECUTION_CONFIG_WRONG_TYPE {
		names = append(names, "COMPLIANCE_SERVER_EXECUTION_CONFIG_WRONG_TYPE")
	}
	if c == COMPLIANCE_SERVER_ILLEGAL_SHARD_MESSAGE {
		names = append(names, "COMPLIANCE_SERVER_ILLEGAL_SHARD_MESSAGE")
	}
	if c == COMPLIANCE_SERVER_MESSAGE_NOT_SUPPORTED {
		names = append(names, "COMPLIANCE_SERVER_MESSAGE_NOT_SUPPORTED")
	}
	if c == COMPLIANCE_SERVER_ORCHESTRATION_FAILED {
		names = append(names, "COMPLIANCE_SERVER_ORCHESTRATION_FAILED")
	}
	if c == COMPLIANCE_SERVER_ORCHESTRATION_TIMEOUT {
		names = append(names, "COMPLIANCE_SERVER_ORCHESTRATION_TIMEOUT")
	}
	if c == COMPLIANCE_SERVER_REGION_MISSING {
		names = append(names, "COMPLIANCE_SERVER_REGION_MISSING")
	}
	if c == COMPLIANCE_SERVER_UNKNOW_ORCHESTRATION_RESPONSE {
		names = append(names, "COMPLIANCE_SERVER_UNKNOW_ORCHESTRATION_RESPONSE")
	}
	if c == COTURN_IP_NOT_FOUND {
		names = append(names, "COTURN_IP_NOT_FOUND")
	}
	if c == CS_ACTIVITY_CENTER_DELAYED {
		names = append(names, "CS_ACTIVITY_CENTER_DELAYED")
	}
	if c == CS_ACTIVITY_CENTER_NOT_FOUND {
		names = append(names, "CS_ACTIVITY_CENTER_NOT_FOUND")
	}
	if c == CS_DEVICE_NOT_SUBSCRIBED {
		names = append(names, "CS_DEVICE_NOT_SUBSCRIBED")
	}
	if c == CS_ELIGIBLE_DEVICE_ALREADY_EXISTING {
		names = append(names, "CS_ELIGIBLE_DEVICE_ALREADY_EXISTING")
	}
	if c == CS_ELIGIBLE_DEVICE_NOT_FOUND {
		names = append(names, "CS_ELIGIBLE_DEVICE_NOT_FOUND")
	}
	if c == CS_FEATURE_INCOMPLETE {
		names = append(names, "CS_FEATURE_INCOMPLETE")
	}
	if c == CS_FORBIDDEN {
		names = append(names, "CS_FORBIDDEN")
	}
	if c == CS_FREE_CLOUDSTORAGE_DEVICE_NOT_FOUND {
		names = append(names, "CS_FREE_CLOUDSTORAGE_DEVICE_NOT_FOUND")
	}
	if c == CS_GENERIC_ERROR {
		names = append(names, "CS_GENERIC_ERROR")
	}
	if c == CS_INVALID_DATA_TYPE {
		names = append(names, "CS_INVALID_DATA_TYPE")
	}
	if c == CS_INVALID_TIME_RANGE {
		names = append(names, "CS_INVALID_TIME_RANGE")
	}
	if c == CS_INVALID_URL {
		names = append(names, "CS_INVALID_URL")
	}
	if c == CS_MEDIA_NOT_FOUND {
		names = append(names, "CS_MEDIA_NOT_FOUND")
	}
	if c == CS_NOT_FOUND {
		names = append(names, "CS_NOT_FOUND")
	}
	if c == CS_OUTDATED_SUBSCRIPTION_EVENT {
		names = append(names, "CS_OUTDATED_SUBSCRIPTION_EVENT")
	}
	if c == CS_PREPARE_HLS_FAILED {
		names = append(names, "CS_PREPARE_HLS_FAILED")
	}
	if c == CS_PREPARE_VIDEO_FILE_FAILED {
		names = append(names, "CS_PREPARE_VIDEO_FILE_FAILED")
	}
	if c == CS_QUOTA_NOT_FOUND {
		names = append(names, "CS_QUOTA_NOT_FOUND")
	}
	if c == CS_STORAGE_PLAN_NOT_FOUND {
		names = append(names, "CS_STORAGE_PLAN_NOT_FOUND")
	}
	if c == CS_VIDEO_NOT_READY {
		names = append(names, "CS_VIDEO_NOT_READY")
	}
	if c == DB_DAO_NOT_INITIALIZED {
		names = append(names, "DB_DAO_NOT_INITIALIZED")
	}
	if c == DB_GENERAL_ERROR {
		names = append(names, "DB_GENERAL_ERROR")
	}
	if c == DB_OBJECT_ALREADY_EXISTING {
		names = append(names, "DB_OBJECT_ALREADY_EXISTING")
	}
	if c == DB_OBJECT_LIMIT_REACHED {
		names = append(names, "DB_OBJECT_LIMIT_REACHED")
	}
	if c == DB_OBJECT_NOT_FOUND {
		names = append(names, "DB_OBJECT_NOT_FOUND")
	}
	if c == DC_GENERIC_ERROR {
		names = append(names, "DC_GENERIC_ERROR")
	}
	if c == DC_INVALID_DEVICE_STATE_VALUE {
		names = append(names, "DC_INVALID_DEVICE_STATE_VALUE")
	}
	if c == DEVICE_CACHE_BUILT_IN_PROGRESS {
		names = append(names, "DEVICE_CACHE_BUILT_IN_PROGRESS")
	}
	if c == DEVICE_CACHE_DEVICE_NOT_FOUND {
		names = append(names, "DEVICE_CACHE_DEVICE_NOT_FOUND")
	}
	if c == DEVICE_CACHE_FAILED_TO_RETRIEVE {
		names = append(names, "DEVICE_CACHE_FAILED_TO_RETRIEVE")
	}
	if c == DEVICE_CACHE_NOT_INITIALIZED {
		names = append(names, "DEVICE_CACHE_NOT_INITIALIZED")
	}
	if c == DEVICE_META_MODULE_NOT_FOUND {
		names = append(names, "DEVICE_META_MODULE_NOT_FOUND")
	}
	if c == DEV_ERROR_IN_USE_BY_OTHER_CLIENT {
		names = append(names, "DEV_ERROR_IN_USE_BY_OTHER_CLIENT")
	}
	if c == DEV_OWNERSHIP_VIOLATION {
		names = append(names, "DEV_OWNERSHIP_VIOLATION")
	}
	if c == DEV_ROLE_UPDATE_NOT_SUPPORTED {
		names = append(names, "DEV_ROLE_UPDATE_NOT_SUPPORTED")
	}
	if c == DEV_TOKEN_ERROR {
		names = append(names, "DEV_TOKEN_ERROR")
	}
	if c == DEV_TOKEN_EXPIRED {
		names = append(names, "DEV_TOKEN_EXPIRED")
	}
	if c == DG_GENERIC_ERROR {
		names = append(names, "DG_GENERIC_ERROR")
	}
	if c == DG_INVALID_ACTION {
		names = append(names, "DG_INVALID_ACTION")
	}
	if c == DG_UNSUPPORTED_DEVICE_FILTER {
		names = append(names, "DG_UNSUPPORTED_DEVICE_FILTER")
	}
	if c == DIRECTED_CUSTOMER_ACCOUNT_TRANSFER_IN_PROGRESS {
		names = append(names, "DIRECTED_CUSTOMER_ACCOUNT_TRANSFER_IN_PROGRESS")
	}
	if c == DIRECTED_CUSTOMER_EMAIL_TAKEN {
		names = append(names, "DIRECTED_CUSTOMER_EMAIL_TAKEN")
	}
	if c == DIRECTED_CUSTOMER_NETWORK_DCID_EXISTS {
		names = append(names, "DIRECTED_CUSTOMER_NETWORK_DCID_EXISTS")
	}
	if c == EC_SYNC_INVENTORY_IN_PROGRESS {
		names = append(names, "EC_SYNC_INVENTORY_IN_PROGRESS")
	}
	if c == ES_GENERIC_ERROR {
		names = append(names, "ES_GENERIC_ERROR")
	}
	if c == ES_REQUEST_FAILED {
		names = append(names, "ES_REQUEST_FAILED")
	}
	if c == ES_REQUEST_NOT_SUPPORTED {
		names = append(names, "ES_REQUEST_NOT_SUPPORTED")
	}
	if c == FEATURE_REGISTRY_EMPTY_FEATURE_LIST {
		names = append(names, "FEATURE_REGISTRY_EMPTY_FEATURE_LIST")
	}
	if c == FEATURE_REGISTRY_FEATURE_NOT_FOUND {
		names = append(names, "FEATURE_REGISTRY_FEATURE_NOT_FOUND")
	}
	if c == FEATURE_REGISTRY_GENERIC_ERROR {
		names = append(names, "FEATURE_REGISTRY_GENERIC_ERROR")
	}
	if c == FFS_GENERAL_EXCEPTION {
		names = append(names, "FFS_GENERAL_EXCEPTION")
	}
	if c == FFS_SESSION_NOT_FOUND {
		names = append(names, "FFS_SESSION_NOT_FOUND")
	}
	if c == GEOFENCE_FAIL_TO_PROCESS_USER_PLACE {
		names = append(names, "GEOFENCE_FAIL_TO_PROCESS_USER_PLACE")
	}
	if c == GEOFENCE_FAIL_TO_PROCESS_USER_PROFILE {
		names = append(names, "GEOFENCE_FAIL_TO_PROCESS_USER_PROFILE")
	}
	if c == GEOFENCE_LAST_KNOWN_PLACE_INFO_OUT_OF_DATE {
		names = append(names, "GEOFENCE_LAST_KNOWN_PLACE_INFO_OUT_OF_DATE")
	}
	if c == GEOFENCE_USER_PLACE_NOT_FOUND {
		names = append(names, "GEOFENCE_USER_PLACE_NOT_FOUND")
	}
	if c == GEOFENCE_USER_PLACE_REQUEST_INVALID {
		names = append(names, "GEOFENCE_USER_PLACE_REQUEST_INVALID")
	}
	if c == GEOFENCE_USER_PLACE_WEBHOOK_INVALID {
		names = append(names, "GEOFENCE_USER_PLACE_WEBHOOK_INVALID")
	}
	if c == GEOFENCE_USER_PROFILE_CONTAINS_UNKNOWN_PLACEID {
		names = append(names, "GEOFENCE_USER_PROFILE_CONTAINS_UNKNOWN_PLACEID")
	}
	if c == GEOFENCE_USER_PROFILE_NOT_FOUND {
		names = append(names, "GEOFENCE_USER_PROFILE_NOT_FOUND")
	}
	if c == GEOFENCE_USER_PROFILE_WEBHOOK_INVALID {
		names = append(names, "GEOFENCE_USER_PROFILE_WEBHOOK_INVALID")
	}
	if c == GEOFENCE_USER_PROFLE_REQUEST_INVALID {
		names = append(names, "GEOFENCE_USER_PROFLE_REQUEST_INVALID")
	}
	if c == HK_APPLE_AUTH_ENTITY_ILLEGAL_STATE {
		names = append(names, "HK_APPLE_AUTH_ENTITY_ILLEGAL_STATE")
	}
	if c == HK_APPLE_AUTH_ENTITY_IN_EXPECTED_STATE_ALREADY {
		names = append(names, "HK_APPLE_AUTH_ENTITY_IN_EXPECTED_STATE_ALREADY")
	}
	if c == HK_APPLE_BAD_REQUEST {
		names = append(names, "HK_APPLE_BAD_REQUEST")
	}
	if c == HK_APPLE_BAD_URL {
		names = append(names, "HK_APPLE_BAD_URL")
	}
	if c == HK_APPLE_DUPLICATE_AUTH_ENTITY {
		names = append(names, "HK_APPLE_DUPLICATE_AUTH_ENTITY")
	}
	if c == HK_APPLE_ENTITY_DOWNLOADS_NOT_AVAILABLE {
		names = append(names, "HK_APPLE_ENTITY_DOWNLOADS_NOT_AVAILABLE")
	}
	if c == HK_APPLE_ENTITY_GENERATION_IN_PROGRESS {
		names = append(names, "HK_APPLE_ENTITY_GENERATION_IN_PROGRESS")
	}
	if c == HK_APPLE_EXCEEDED_MAX_PAYLOAD_COUNT {
		names = append(names, "HK_APPLE_EXCEEDED_MAX_PAYLOAD_COUNT")
	}
	if c == HK_APPLE_HTTP_METHOD_NOT_ALLOWED {
		names = append(names, "HK_APPLE_HTTP_METHOD_NOT_ALLOWED")
	}
	if c == HK_APPLE_INSUFFICIENT_AUTH_ENTITIES {
		names = append(names, "HK_APPLE_INSUFFICIENT_AUTH_ENTITIES")
	}
	if c == HK_APPLE_INTERNAL_SERVER_ERROR {
		names = append(names, "HK_APPLE_INTERNAL_SERVER_ERROR")
	}
	if c == HK_APPLE_INVALID_AUTH_ENTITY {
		names = append(names, "HK_APPLE_INVALID_AUTH_ENTITY")
	}
	if c == HK_APPLE_INVALID_AUTH_ENTITY_TYPE {
		names = append(names, "HK_APPLE_INVALID_AUTH_ENTITY_TYPE")
	}
	if c == HK_APPLE_INVALID_OR_UNSUPPORTED_PPID {
		names = append(names, "HK_APPLE_INVALID_OR_UNSUPPORTED_PPID")
	}
	if c == HK_APPLE_INVALID_PPID {
		names = append(names, "HK_APPLE_INVALID_PPID")
	}
	if c == HK_APPLE_INVALID_REQUEST_ID {
		names = append(names, "HK_APPLE_INVALID_REQUEST_ID")
	}
	if c == HK_APPLE_JSON_ERROR {
		names = append(names, "HK_APPLE_JSON_ERROR")
	}
	if c == HK_APPLE_MISMATCHED_PLAN_ID {
		names = append(names, "HK_APPLE_MISMATCHED_PLAN_ID")
	}
	if c == HK_APPLE_MISMATCHED_UUID {
		names = append(names, "HK_APPLE_MISMATCHED_UUID")
	}
	if c == HK_APPLE_MISSING_MANDATORY_PARAMETERS {
		names = append(names, "HK_APPLE_MISSING_MANDATORY_PARAMETERS")
	}
	if c == HK_APPLE_SERVICE_UNAVAILABLE {
		names = append(names, "HK_APPLE_SERVICE_UNAVAILABLE")
	}
	if c == HK_APPLE_UNAUTHORIZED_USER {
		names = append(names, "HK_APPLE_UNAUTHORIZED_USER")
	}
	if c == HK_APPLE_UNKNOWN_SERVER_ERROR {
		names = append(names, "HK_APPLE_UNKNOWN_SERVER_ERROR")
	}
	if c == HK_ERROR_DEVICE_MODEL_NOT_SUPPORTED {
		names = append(names, "HK_ERROR_DEVICE_MODEL_NOT_SUPPORTED")
	}
	if c == HK_ERROR_DOWNLOAD_TOKENS {
		names = append(names, "HK_ERROR_DOWNLOAD_TOKENS")
	}
	if c == HK_ERROR_FILE_PARSE {
		names = append(names, "HK_ERROR_FILE_PARSE")
	}
	if c == HK_ERROR_HARDWARE_VERSION_NOT_SUPPORTED {
		names = append(names, "HK_ERROR_HARDWARE_VERSION_NOT_SUPPORTED")
	}
	if c == HK_ERROR_TOKENS_NOT_AVAILABLE {
		names = append(names, "HK_ERROR_TOKENS_NOT_AVAILABLE")
	}
	if c == HK_PREVIOUS_RETRIEVE_TOKENS_IS_BEING_PROCESSED {
		names = append(names, "HK_PREVIOUS_RETRIEVE_TOKENS_IS_BEING_PROCESSED")
	}
	if c == HK_PRODUCT_PLAN_EXISTS {
		names = append(names, "HK_PRODUCT_PLAN_EXISTS")
	}
	if c == HK_PRODUCT_PLAN_NOT_FOUND {
		names = append(names, "HK_PRODUCT_PLAN_NOT_FOUND")
	}
	if c == HK_REGISTRATION_IN_PROGRESS {
		names = append(names, "HK_REGISTRATION_IN_PROGRESS")
	}
	if c == HK_REQUEST_TOKEN_NOT_FOUND {
		names = append(names, "HK_REQUEST_TOKEN_NOT_FOUND")
	}
	if c == HK_TOKENS_NOT_YET_READY_FOR_DOWNLOAD {
		names = append(names, "HK_TOKENS_NOT_YET_READY_FOR_DOWNLOAD")
	}
	if c == HK_TOKEN_REQUEST_NOT_APPLICABLE_FOR_FACTORY {
		names = append(names, "HK_TOKEN_REQUEST_NOT_APPLICABLE_FOR_FACTORY")
	}
	if c == HK_UNDOCUMENTED_APPLE_ERROR {
		names = append(names, "HK_UNDOCUMENTED_APPLE_ERROR")
	}
	if c == INCORRECT_APP_SERVER_URL {
		names = append(names, "INCORRECT_APP_SERVER_URL")
	}
	if c == INCORRECT_VERIFICATION_CODE {
		names = append(names, "INCORRECT_VERIFICATION_CODE")
	}
	if c == IOT_ACTIVITY_NOT_FOUND {
		names = append(names, "IOT_ACTIVITY_NOT_FOUND")
	}
	if c == IOT_ASSERTION_EXCEPTION {
		names = append(names, "IOT_ASSERTION_EXCEPTION")
	}
	if c == IOT_BACKEND_SERVICE_UNAVAILABLE {
		names = append(names, "IOT_BACKEND_SERVICE_UNAVAILABLE")
	}
	if c == IOT_CACHE_OUT_OF_SERVICE {
		names = append(names, "IOT_CACHE_OUT_OF_SERVICE")
	}
	if c == IOT_CAPABILITY_NOT_SUPPORTED {
		names = append(names, "IOT_CAPABILITY_NOT_SUPPORTED")
	}
	if c == IOT_CIRCUIT_BREAKER_OPEN {
		names = append(names, "IOT_CIRCUIT_BREAKER_OPEN")
	}
	if c == IOT_DEPENDENT_SERVICE_NOT_INITIALIZED {
		names = append(names, "IOT_DEPENDENT_SERVICE_NOT_INITIALIZED")
	}
	if c == IOT_DESERIALIZATION_EXCEPTION {
		names = append(names, "IOT_DESERIALIZATION_EXCEPTION")
	}
	if c == IOT_DEVICE_ASSOCIATED_WITH_ANOTHER_PARENT_DEVICE {
		names = append(names, "IOT_DEVICE_ASSOCIATED_WITH_ANOTHER_PARENT_DEVICE")
	}
	if c == IOT_DEVICE_ERROR {
		names = append(names, "IOT_DEVICE_ERROR")
	}
	if c == IOT_DEVICE_LOGS_NOT_FOUND {
		names = append(names, "IOT_DEVICE_LOGS_NOT_FOUND")
	}
	if c == IOT_DEVICE_REGION_ERROR {
		names = append(names, "IOT_DEVICE_REGION_ERROR")
	}
	if c == IOT_EVENT_NOT_SUPPORTED_EXCEPTION {
		names = append(names, "IOT_EVENT_NOT_SUPPORTED_EXCEPTION")
	}
	if c == IOT_EVENT_RUNTIME_EXCEPTION {
		names = append(names, "IOT_EVENT_RUNTIME_EXCEPTION")
	}
	if c == IOT_EXTERNAL_NETWORK_SERVICE_DOES_NOT_EXIST {
		names = append(names, "IOT_EXTERNAL_NETWORK_SERVICE_DOES_NOT_EXIST")
	}
	if c == IOT_EXTERNAL_NETWORK_SERVICE_NOT_AVAILABLE {
		names = append(names, "IOT_EXTERNAL_NETWORK_SERVICE_NOT_AVAILABLE")
	}
	if c == IOT_FAILED_TO_INITIALIZE_CAPABILITIES {
		names = append(names, "IOT_FAILED_TO_INITIALIZE_CAPABILITIES")
	}
	if c == IOT_FIRMWARE_UPGRADE_FAILED {
		names = append(names, "IOT_FIRMWARE_UPGRADE_FAILED")
	}
	if c == IOT_FIRMWARE_UPGRADE_IN_PROGRESS {
		names = append(names, "IOT_FIRMWARE_UPGRADE_IN_PROGRESS")
	}
	if c == IOT_GENERAL_EXCEPTION {
		names = append(names, "IOT_GENERAL_EXCEPTION")
	}
	if c == IOT_GENERAL_SCENE_EXCEPTION {
		names = append(names, "IOT_GENERAL_SCENE_EXCEPTION")
	}
	if c == IOT_GOOGLE_WEAVE_MODEL_MANIFEST_ALREADY_EXISTS {
		names = append(names, "IOT_GOOGLE_WEAVE_MODEL_MANIFEST_ALREADY_EXISTS")
	}
	if c == IOT_GOOGLE_WEAVE_MODEL_MANIFEST_NOT_FOUND {
		names = append(names, "IOT_GOOGLE_WEAVE_MODEL_MANIFEST_NOT_FOUND")
	}
	if c == IOT_INITIALIZATION_FAILED {
		names = append(names, "IOT_INITIALIZATION_FAILED")
	}
	if c == IOT_INTEGRATION_EXCEPTION {
		names = append(names, "IOT_INTEGRATION_EXCEPTION")
	}
	if c == IOT_INVALID_CONFIGURATION {
		names = append(names, "IOT_INVALID_CONFIGURATION")
	}
	if c == IOT_INVALID_EVENT_EXCEPTION {
		names = append(names, "IOT_INVALID_EVENT_EXCEPTION")
	}
	if c == IOT_INVALID_MATCHER_EXCLUDE_PATH {
		names = append(names, "IOT_INVALID_MATCHER_EXCLUDE_PATH")
	}
	if c == IOT_INVALID_TIMEZONE_ID {
		names = append(names, "IOT_INVALID_TIMEZONE_ID")
	}
	if c == IOT_IO_EXCEPTION {
		names = append(names, "IOT_IO_EXCEPTION")
	}
	if c == IOT_JWT_REFRESH_TOKEN_DECODE_FAILED {
		names = append(names, "IOT_JWT_REFRESH_TOKEN_DECODE_FAILED")
	}
	if c == IOT_JWT_REFRESH_TOKEN_GENERATE_FAILED {
		names = append(names, "IOT_JWT_REFRESH_TOKEN_GENERATE_FAILED")
	}
	if c == IOT_JWT_TOKEN_DECODE_FAILED {
		names = append(names, "IOT_JWT_TOKEN_DECODE_FAILED")
	}
	if c == IOT_JWT_TOKEN_GENERATE_FAILED {
		names = append(names, "IOT_JWT_TOKEN_GENERATE_FAILED")
	}
	if c == IOT_JWT_TOKEN_INVALID {
		names = append(names, "IOT_JWT_TOKEN_INVALID")
	}
	if c == IOT_KINESIS_STREAM_PRODUCER_INIT_FAILED {
		names = append(names, "IOT_KINESIS_STREAM_PRODUCER_INIT_FAILED")
	}
	if c == IOT_LOCATION_NOT_LINKED {
		names = append(names, "IOT_LOCATION_NOT_LINKED")
	}
	if c == IOT_MESSAGEBROKER_INITIALIZATION_ERROR {
		names = append(names, "IOT_MESSAGEBROKER_INITIALIZATION_ERROR")
	}
	if c == IOT_OUT_OF_TIME_EXCEPTION {
		names = append(names, "IOT_OUT_OF_TIME_EXCEPTION")
	}
	if c == IOT_REGION_ENDPOINT_NOT_AVAILABLE {
		names = append(names, "IOT_REGION_ENDPOINT_NOT_AVAILABLE")
	}
	if c == IOT_REQUEST_VERSION_NOT_SUPPORTED {
		names = append(names, "IOT_REQUEST_VERSION_NOT_SUPPORTED")
	}
	if c == IOT_ROUTER_RULE_NOT_FOUND {
		names = append(names, "IOT_ROUTER_RULE_NOT_FOUND")
	}
	if c == IOT_SCENE_DEVICE_NOT_FOUND {
		names = append(names, "IOT_SCENE_DEVICE_NOT_FOUND")
	}
	if c == IOT_SCENE_NOT_FOUND {
		names = append(names, "IOT_SCENE_NOT_FOUND")
	}
	if c == IOT_SERIALIZATION_EXCEPTION {
		names = append(names, "IOT_SERIALIZATION_EXCEPTION")
	}
	if c == IOT_SHARED_VIDEOS_NOT_AVAILABLE {
		names = append(names, "IOT_SHARED_VIDEOS_NOT_AVAILABLE")
	}
	if c == IOT_SSL_INITIALIZATION_ERROR {
		names = append(names, "IOT_SSL_INITIALIZATION_ERROR")
	}
	if c == IOT_SYSTEM_SHUTDOWN_INITIATED {
		names = append(names, "IOT_SYSTEM_SHUTDOWN_INITIATED")
	}
	if c == IOT_TECHNICAL_EXCEPTION {
		names = append(names, "IOT_TECHNICAL_EXCEPTION")
	}
	if c == IOT_TIMEZONE_NOT_FOUND {
		names = append(names, "IOT_TIMEZONE_NOT_FOUND")
	}
	if c == IOT_TRY_LOCK_TIMEOUT {
		names = append(names, "IOT_TRY_LOCK_TIMEOUT")
	}
	if c == IOT_UNIDENTIFIED_JWT_TOKEN_ISSUER {
		names = append(names, "IOT_UNIDENTIFIED_JWT_TOKEN_ISSUER")
	}
	if c == IOT_UNKNOWN_DEVICE_TYPE {
		names = append(names, "IOT_UNKNOWN_DEVICE_TYPE")
	}
	if c == IOT_UNRECOGNIZED_VIDEO_ANALYTICS_CLASSIFICATION {
		names = append(names, "IOT_UNRECOGNIZED_VIDEO_ANALYTICS_CLASSIFICATION")
	}
	if c == K8S_CLIENT_INIT_WRONG_TYPE {
		names = append(names, "K8S_CLIENT_INIT_WRONG_TYPE")
	}
	if c == K8S_CLIENT_REQUEST_ERROR {
		names = append(names, "K8S_CLIENT_REQUEST_ERROR")
	}
	if c == KC_GENERIC_ERROR {
		names = append(names, "KC_GENERIC_ERROR")
	}
	if c == LIGHTING_EFFECTS_NOT_FOUND {
		names = append(names, "LIGHTING_EFFECTS_NOT_FOUND")
	}
	if c == LIGHTING_EFFECTS_PREDEFINED_EFFECT_ALREADY_EXISTS {
		names = append(names, "LIGHTING_EFFECTS_PREDEFINED_EFFECT_ALREADY_EXISTS")
	}
	if c == LIGHTING_EFFECTS_PREDEFINED_EFFECT_TEMPLATE_ALREADY_EXISTS {
		names = append(names, "LIGHTING_EFFECTS_PREDEFINED_EFFECT_TEMPLATE_ALREADY_EXISTS")
	}
	if c == LIGHTING_EFFECTS_PREDEFINED_EFFECT_TEMPLATE_NOT_FOUND {
		names = append(names, "LIGHTING_EFFECTS_PREDEFINED_EFFECT_TEMPLATE_NOT_FOUND")
	}
	if c == LIGHTING_EFFECTS_VALIDATION_FAILED {
		names = append(names, "LIGHTING_EFFECTS_VALIDATION_FAILED")
	}
	if c == MC_METHOD_NOT_ALLOWED {
		names = append(names, "MC_METHOD_NOT_ALLOWED")
	}
	if c == MC_RESOURCE_NOT_FOUND {
		names = append(names, "MC_RESOURCE_NOT_FOUND")
	}
	if c == MC_UNDOCUMENTED_ERROR {
		names = append(names, "MC_UNDOCUMENTED_ERROR")
	}
	if c == METRICS_PROVIDER_INVALID {
		names = append(names, "METRICS_PROVIDER_INVALID")
	}
	if c == METRICS_PROVIDER_NOT_FOUND {
		names = append(names, "METRICS_PROVIDER_NOT_FOUND")
	}
	if c == METRICS_PROVIDER_NOT_REGISTERED {
		names = append(names, "METRICS_PROVIDER_NOT_REGISTERED")
	}
	if c == MFA_PROCESS_CLOSED {
		names = append(names, "MFA_PROCESS_CLOSED")
	}
	if c == MFA_PROCESS_EXPIRED {
		names = append(names, "MFA_PROCESS_EXPIRED")
	}
	if c == PC_SEVER_ERROR {
		names = append(names, "PC_SEVER_ERROR")
	}
	if c == PC_SYSTEM_UPDATING {
		names = append(names, "PC_SYSTEM_UPDATING")
	}
	if c == PC_UNDOCUMENTED_ERROR {
		names = append(names, "PC_UNDOCUMENTED_ERROR")
	}
	if c == REQUEST_TOO_MUCH_CODE {
		names = append(names, "REQUEST_TOO_MUCH_CODE")
	}
	if c == SA_CONDITION_NOT_EVALUATED {
		names = append(names, "SA_CONDITION_NOT_EVALUATED")
	}
	if c == SA_EXCEPTION {
		names = append(names, "SA_EXCEPTION")
	}
	if c == SA_EXECUTION_ALREADY_IN_PROGRESS {
		names = append(names, "SA_EXECUTION_ALREADY_IN_PROGRESS")
	}
	if c == SA_EXECUTION_PLAN_GENERATION_FAILED {
		names = append(names, "SA_EXECUTION_PLAN_GENERATION_FAILED")
	}
	if c == SA_EXECUTION_PLAN_NOT_AVAILABLE {
		names = append(names, "SA_EXECUTION_PLAN_NOT_AVAILABLE")
	}
	if c == SA_FAILED_TO_SATISFY_CONDITION {
		names = append(names, "SA_FAILED_TO_SATISFY_CONDITION")
	}
	if c == SA_GEO_LOCATION_NOT_FOUND {
		names = append(names, "SA_GEO_LOCATION_NOT_FOUND")
	}
	if c == SA_INVALID_CONDITION {
		names = append(names, "SA_INVALID_CONDITION")
	}
	if c == SA_RANGE_KEY_MISSING {
		names = append(names, "SA_RANGE_KEY_MISSING")
	}
	if c == SA_RECOVERING_FAILED_EXECUTION {
		names = append(names, "SA_RECOVERING_FAILED_EXECUTION")
	}
	if c == SA_RULES_QUOTA_EXCEEDED {
		names = append(names, "SA_RULES_QUOTA_EXCEEDED")
	}
	if c == SA_SKIP_BY_CONFIG {
		names = append(names, "SA_SKIP_BY_CONFIG")
	}
	if c == SA_SMART_ACTION_NOT_INITIALIZED {
		names = append(names, "SA_SMART_ACTION_NOT_INITIALIZED")
	}
	if c == SA_TIMEZONE_NOT_FOUND {
		names = append(names, "SA_TIMEZONE_NOT_FOUND")
	}
	if c == SC_MODULE_NOT_REGISTERED {
		names = append(names, "SC_MODULE_NOT_REGISTERED")
	}
	if c == SC_PAYLOAD_NOT_FOUND {
		names = append(names, "SC_PAYLOAD_NOT_FOUND")
	}
	if c == SC_SCHEDULE_NOT_FOUND {
		names = append(names, "SC_SCHEDULE_NOT_FOUND")
	}
	if c == SGW_ACCOUNT_ALREADY_EXISTS {
		names = append(names, "SGW_ACCOUNT_ALREADY_EXISTS")
	}
	if c == SGW_ACCOUNT_ALREADY_INACTIVE {
		names = append(names, "SGW_ACCOUNT_ALREADY_INACTIVE")
	}
	if c == SGW_ACCOUNT_BAD_REQUEST {
		names = append(names, "SGW_ACCOUNT_BAD_REQUEST")
	}
	if c == SGW_ACCOUNT_BILLING_INFO_REQUIRED {
		names = append(names, "SGW_ACCOUNT_BILLING_INFO_REQUIRED")
	}
	if c == SGW_ACCOUNT_CLOSED {
		names = append(names, "SGW_ACCOUNT_CLOSED")
	}
	if c == SGW_ACCOUNT_IMMUTABLE_SUBSCRIPTION {
		names = append(names, "SGW_ACCOUNT_IMMUTABLE_SUBSCRIPTION")
	}
	if c == SGW_ACCOUNT_INTERNAL_SERVER_ERROR {
		names = append(names, "SGW_ACCOUNT_INTERNAL_SERVER_ERROR")
	}
	if c == SGW_ACCOUNT_INVALID_API_KEY {
		names = append(names, "SGW_ACCOUNT_INVALID_API_KEY")
	}
	if c == SGW_ACCOUNT_INVALID_API_VERSION {
		names = append(names, "SGW_ACCOUNT_INVALID_API_VERSION")
	}
	if c == SGW_ACCOUNT_INVALID_CONFIGURATION {
		names = append(names, "SGW_ACCOUNT_INVALID_CONFIGURATION")
	}
	if c == SGW_ACCOUNT_INVALID_CONTENT_TYPE {
		names = append(names, "SGW_ACCOUNT_INVALID_CONTENT_TYPE")
	}
	if c == SGW_ACCOUNT_INVALID_DATA {
		names = append(names, "SGW_ACCOUNT_INVALID_DATA")
	}
	if c == SGW_ACCOUNT_INVALID_PERMISSIONS {
		names = append(names, "SGW_ACCOUNT_INVALID_PERMISSIONS")
	}
	if c == SGW_ACCOUNT_INVALID_TOKEN {
		names = append(names, "SGW_ACCOUNT_INVALID_TOKEN")
	}
	if c == SGW_ACCOUNT_INVALID_TRANSITION {
		names = append(names, "SGW_ACCOUNT_INVALID_TRANSITION")
	}
	if c == SGW_ACCOUNT_MANDATORY_FIELD {
		names = append(names, "SGW_ACCOUNT_MANDATORY_FIELD")
	}
	if c == SGW_ACCOUNT_MISSING_FEATURE {
		names = append(names, "SGW_ACCOUNT_MISSING_FEATURE")
	}
	if c == SGW_ACCOUNT_NOT_CLOSED {
		names = append(names, "SGW_ACCOUNT_NOT_CLOSED")
	}
	if c == SGW_ACCOUNT_NOT_FOUND {
		names = append(names, "SGW_ACCOUNT_NOT_FOUND")
	}
	if c == SGW_ACCOUNT_RATE_LIMITED {
		names = append(names, "SGW_ACCOUNT_RATE_LIMITED")
	}
	if c == SGW_ACCOUNT_SIMULTANEOUS_REQUEST {
		names = append(names, "SGW_ACCOUNT_SIMULTANEOUS_REQUEST")
	}
	if c == SGW_ACCOUNT_STATE_INVALID {
		names = append(names, "SGW_ACCOUNT_STATE_INVALID")
	}
	if c == SGW_ACCOUNT_TRANSACTION {
		names = append(names, "SGW_ACCOUNT_TRANSACTION")
	}
	if c == SGW_ACCOUNT_UNACCEPTABLE_VALUE {
		names = append(names, "SGW_ACCOUNT_UNACCEPTABLE_VALUE")
	}
	if c == SGW_ACCOUNT_UNAUTHORIZED {
		names = append(names, "SGW_ACCOUNT_UNAUTHORIZED")
	}
	if c == SGW_ACCOUNT_UNAVAILABLE_IN_API_VERSION {
		names = append(names, "SGW_ACCOUNT_UNAVAILABLE_IN_API_VERSION")
	}
	if c == SGW_ACCOUNT_UNKNOWN_API_VERSION {
		names = append(names, "SGW_ACCOUNT_UNKNOWN_API_VERSION")
	}
	if c == SGW_ACCOUNT_VALIDATION {
		names = append(names, "SGW_ACCOUNT_VALIDATION")
	}
	if c == SGW_ACCOUNT_VALUE_NOT_A_NUMBER {
		names = append(names, "SGW_ACCOUNT_VALUE_NOT_A_NUMBER")
	}
	if c == SGW_ACCOUNT_VALUE_NOT_INCLUDED_IN_LIST {
		names = append(names, "SGW_ACCOUNT_VALUE_NOT_INCLUDED_IN_LIST")
	}
	if c == SGW_ACCOUNT_VERSION_NOT_FOUND {
		names = append(names, "SGW_ACCOUNT_VERSION_NOT_FOUND")
	}
	if c == SGW_BILLING_INFO_ALREADY_EXISTS {
		names = append(names, "SGW_BILLING_INFO_ALREADY_EXISTS")
	}
	if c == SGW_BILLING_INFO_BAD_REQUEST {
		names = append(names, "SGW_BILLING_INFO_BAD_REQUEST")
	}
	if c == SGW_BILLING_INFO_CONFIGURATION {
		names = append(names, "SGW_BILLING_INFO_CONFIGURATION")
	}
	if c == SGW_BILLING_INFO_IMMUTABLE_SUBSCRIPTION {
		names = append(names, "SGW_BILLING_INFO_IMMUTABLE_SUBSCRIPTION")
	}
	if c == SGW_BILLING_INFO_INTERNAL_SERVER_ERROR {
		names = append(names, "SGW_BILLING_INFO_INTERNAL_SERVER_ERROR")
	}
	if c == SGW_BILLING_INFO_INVALID_API_KEY {
		names = append(names, "SGW_BILLING_INFO_INVALID_API_KEY")
	}
	if c == SGW_BILLING_INFO_INVALID_API_VERSION {
		names = append(names, "SGW_BILLING_INFO_INVALID_API_VERSION")
	}
	if c == SGW_BILLING_INFO_INVALID_CONTENT_TYPE {
		names = append(names, "SGW_BILLING_INFO_INVALID_CONTENT_TYPE")
	}
	if c == SGW_BILLING_INFO_INVALID_DATA {
		names = append(names, "SGW_BILLING_INFO_INVALID_DATA")
	}
	if c == SGW_BILLING_INFO_INVALID_PERMISSIONS {
		names = append(names, "SGW_BILLING_INFO_INVALID_PERMISSIONS")
	}
	if c == SGW_BILLING_INFO_INVALID_TOKEN {
		names = append(names, "SGW_BILLING_INFO_INVALID_TOKEN")
	}
	if c == SGW_BILLING_INFO_MANDATORY_FIELD {
		names = append(names, "SGW_BILLING_INFO_MANDATORY_FIELD")
	}
	if c == SGW_BILLING_INFO_MISSING_FEATURE {
		names = append(names, "SGW_BILLING_INFO_MISSING_FEATURE")
	}
	if c == SGW_BILLING_INFO_NOT_FOUND {
		names = append(names, "SGW_BILLING_INFO_NOT_FOUND")
	}
	if c == SGW_BILLING_INFO_RATE_LIMITED {
		names = append(names, "SGW_BILLING_INFO_RATE_LIMITED")
	}
	if c == SGW_BILLING_INFO_REQUIRED {
		names = append(names, "SGW_BILLING_INFO_REQUIRED")
	}
	if c == SGW_BILLING_INFO_SIMULTANEOUS_REQUEST {
		names = append(names, "SGW_BILLING_INFO_SIMULTANEOUS_REQUEST")
	}
	if c == SGW_BILLING_INFO_TOKEN_INVALID {
		names = append(names, "SGW_BILLING_INFO_TOKEN_INVALID")
	}
	if c == SGW_BILLING_INFO_TRANSACTION {
		names = append(names, "SGW_BILLING_INFO_TRANSACTION")
	}
	if c == SGW_BILLING_INFO_UNAUTHORIZED {
		names = append(names, "SGW_BILLING_INFO_UNAUTHORIZED")
	}
	if c == SGW_BILLING_INFO_UNAVAILABLE_IN_API_VERSION {
		names = append(names, "SGW_BILLING_INFO_UNAVAILABLE_IN_API_VERSION")
	}
	if c == SGW_BILLING_INFO_UNKNOWN_API_VERSION {
		names = append(names, "SGW_BILLING_INFO_UNKNOWN_API_VERSION")
	}
	if c == SGW_BILLING_INFO_UPDATE_DECLINE {
		names = append(names, "SGW_BILLING_INFO_UPDATE_DECLINE")
	}
	if c == SGW_BILLING_INFO_VALIDATION {
		names = append(names, "SGW_BILLING_INFO_VALIDATION")
	}
	if c == SGW_BILLING_INFO_VALUE_NOT_A_NUMBER {
		names = append(names, "SGW_BILLING_INFO_VALUE_NOT_A_NUMBER")
	}
	if c == SGW_COUPON_NOT_FOUND {
		names = append(names, "SGW_COUPON_NOT_FOUND")
	}
	if c == SGW_COUPON_NOT_REDEEMABLE {
		names = append(names, "SGW_COUPON_NOT_REDEEMABLE")
	}
	if c == SGW_DEVICES_ALREADY_IN_MAX_BIND_NUMBER {
		names = append(names, "SGW_DEVICES_ALREADY_IN_MAX_BIND_NUMBER")
	}
	if c == SGW_DEVICE_ALREADY_IN_FREE_TRIAL {
		names = append(names, "SGW_DEVICE_ALREADY_IN_FREE_TRIAL")
	}
	if c == SGW_DEVICE_ALREADY_IN_PREMIUM_SUBSCRIPTION {
		names = append(names, "SGW_DEVICE_ALREADY_IN_PREMIUM_SUBSCRIPTION")
	}
	if c == SGW_FREE_TRIAL_PERIOD_INACTIVE {
		names = append(names, "SGW_FREE_TRIAL_PERIOD_INACTIVE")
	}
	if c == SGW_GENERIC_ERROR {
		names = append(names, "SGW_GENERIC_ERROR")
	}
	if c == SGW_INVALID_EMAIL {
		names = append(names, "SGW_INVALID_EMAIL")
	}
	if c == SGW_INVALID_PAYMENT_TOKEN {
		names = append(names, "SGW_INVALID_PAYMENT_TOKEN")
	}
	if c == SGW_INVALID_TRANSACTION {
		names = append(names, "SGW_INVALID_TRANSACTION")
	}
	if c == SGW_JOB_IN_PROGRESS {
		names = append(names, "SGW_JOB_IN_PROGRESS")
	}
	if c == SGW_NOT_SUPPORTED {
		names = append(names, "SGW_NOT_SUPPORTED")
	}
	if c == SGW_PAYMENT_ACH_TRANSACTIONS_NOT_SUPPORTED {
		names = append(names, "SGW_PAYMENT_ACH_TRANSACTIONS_NOT_SUPPORTED")
	}
	if c == SGW_PAYMENT_API_ERROR {
		names = append(names, "SGW_PAYMENT_API_ERROR")
	}
	if c == SGW_PAYMENT_APPROVED {
		names = append(names, "SGW_PAYMENT_APPROVED")
	}
	if c == SGW_PAYMENT_APPROVED_FRAUD_REVIEW {
		names = append(names, "SGW_PAYMENT_APPROVED_FRAUD_REVIEW")
	}
	if c == SGW_PAYMENT_AUTHORIZATION_ALREADY_CAPTURED {
		names = append(names, "SGW_PAYMENT_AUTHORIZATION_ALREADY_CAPTURED")
	}
	if c == SGW_PAYMENT_AUTHORIZATION_AMOUNT_DEPLETED {
		names = append(names, "SGW_PAYMENT_AUTHORIZATION_AMOUNT_DEPLETED")
	}
	if c == SGW_PAYMENT_AUTHORIZATION_EXPIRED {
		names = append(names, "SGW_PAYMENT_AUTHORIZATION_EXPIRED")
	}
	if c == SGW_PAYMENT_CALL_ISSUER {
		names = append(names, "SGW_PAYMENT_CALL_ISSUER")
	}
	if c == SGW_PAYMENT_CALL_ISSUER_UPDATE_CARDHOLDER_DATA {
		names = append(names, "SGW_PAYMENT_CALL_ISSUER_UPDATE_CARDHOLDER_DATA")
	}
	if c == SGW_PAYMENT_CANNOT_REFUND_UNSETTLED_TRANSACTIONS {
		names = append(names, "SGW_PAYMENT_CANNOT_REFUND_UNSETTLED_TRANSACTIONS")
	}
	if c == SGW_PAYMENT_CANNOT_VOID_PAYMENT_AUTHORIZATION {
		names = append(names, "SGW_PAYMENT_CANNOT_VOID_PAYMENT_AUTHORIZATION")
	}
	if c == SGW_PAYMENT_CARDHOLDER_REQUESTED_STOP {
		names = append(names, "SGW_PAYMENT_CARDHOLDER_REQUESTED_STOP")
	}
	if c == SGW_PAYMENT_CARD_NOT_ACTIVATED {
		names = append(names, "SGW_PAYMENT_CARD_NOT_ACTIVATED")
	}
	if c == SGW_PAYMENT_CARD_TYPE_NOT_ACCEPTED {
		names = append(names, "SGW_PAYMENT_CARD_TYPE_NOT_ACCEPTED")
	}
	if c == SGW_PAYMENT_CONTACT_GATEWAY {
		names = append(names, "SGW_PAYMENT_CONTACT_GATEWAY")
	}
	if c == SGW_PAYMENT_CURRENCY_NOT_SUPPORTED {
		names = append(names, "SGW_PAYMENT_CURRENCY_NOT_SUPPORTED")
	}
	if c == SGW_PAYMENT_CUSTOMER_CANCELED_TRANSACTION {
		names = append(names, "SGW_PAYMENT_CUSTOMER_CANCELED_TRANSACTION")
	}
	if c == SGW_PAYMENT_CVV_REQUIRED {
		names = append(names, "SGW_PAYMENT_CVV_REQUIRED")
	}
	if c == SGW_PAYMENT_DECLINED {
		names = append(names, "SGW_PAYMENT_DECLINED")
	}
	if c == SGW_PAYMENT_DECLINED_CARD_NUMBER {
		names = append(names, "SGW_PAYMENT_DECLINED_CARD_NUMBER")
	}
	if c == SGW_PAYMENT_DECLINED_EXCEPTION {
		names = append(names, "SGW_PAYMENT_DECLINED_EXCEPTION")
	}
	if c == SGW_PAYMENT_DECLINED_EXPIRATION_DATE {
		names = append(names, "SGW_PAYMENT_DECLINED_EXPIRATION_DATE")
	}
	if c == SGW_PAYMENT_DECLINED_MISSING_DATA {
		names = append(names, "SGW_PAYMENT_DECLINED_MISSING_DATA")
	}
	if c == SGW_PAYMENT_DECLINED_SECURITY_CODE {
		names = append(names, "SGW_PAYMENT_DECLINED_SECURITY_CODE")
	}
	if c == SGW_PAYMENT_DEPOSIT_REFERENCED_CHARGEBACK {
		names = append(names, "SGW_PAYMENT_DEPOSIT_REFERENCED_CHARGEBACK")
	}
	if c == SGW_PAYMENT_DUPLICATE_TRANSACTION {
		names = append(names, "SGW_PAYMENT_DUPLICATE_TRANSACTION")
	}
	if c == SGW_PAYMENT_EXCEEDS_DAILY_LIMIT {
		names = append(names, "SGW_PAYMENT_EXCEEDS_DAILY_LIMIT")
	}
	if c == SGW_PAYMENT_EXPIRED_CARD {
		names = append(names, "SGW_PAYMENT_EXPIRED_CARD")
	}
	if c == SGW_PAYMENT_FRAUD_ADDRESS {
		names = append(names, "SGW_PAYMENT_FRAUD_ADDRESS")
	}
	if c == SGW_PAYMENT_FRAUD_ADDRESS_RECURLY {
		names = append(names, "SGW_PAYMENT_FRAUD_ADDRESS_RECURLY")
	}
	if c == SGW_PAYMENT_FRAUD_ADVANCED_VERIFICATION {
		names = append(names, "SGW_PAYMENT_FRAUD_ADVANCED_VERIFICATION")
	}
	if c == SGW_PAYMENT_FRAUD_GATEWAY {
		names = append(names, "SGW_PAYMENT_FRAUD_GATEWAY")
	}
	if c == SGW_PAYMENT_FRAUD_GENERIC {
		names = append(names, "SGW_PAYMENT_FRAUD_GENERIC")
	}
	if c == SGW_PAYMENT_FRAUD_IP_ADDRESS {
		names = append(names, "SGW_PAYMENT_FRAUD_IP_ADDRESS")
	}
	if c == SGW_PAYMENT_FRAUD_MANUAL_DECISION {
		names = append(names, "SGW_PAYMENT_FRAUD_MANUAL_DECISION")
	}
	if c == SGW_PAYMENT_FRAUD_RISK_CHECK {
		names = append(names, "SGW_PAYMENT_FRAUD_RISK_CHECK")
	}
	if c == SGW_PAYMENT_FRAUD_SECURITY_CODE {
		names = append(names, "SGW_PAYMENT_FRAUD_SECURITY_CODE")
	}
	if c == SGW_PAYMENT_FRAUD_STOLEN_CARD {
		names = append(names, "SGW_PAYMENT_FRAUD_STOLEN_CARD")
	}
	if c == SGW_PAYMENT_FRAUD_TOO_MANY_ATTEMPTS {
		names = append(names, "SGW_PAYMENT_FRAUD_TOO_MANY_ATTEMPTS")
	}
	if c == SGW_PAYMENT_FRAUD_VELOCITY {
		names = append(names, "SGW_PAYMENT_FRAUD_VELOCITY")
	}
	if c == SGW_PAYMENT_GATEWAY_ERROR {
		names = append(names, "SGW_PAYMENT_GATEWAY_ERROR")
	}
	if c == SGW_PAYMENT_GATEWAY_TIMEOUT {
		names = append(names, "SGW_PAYMENT_GATEWAY_TIMEOUT")
	}
	if c == SGW_PAYMENT_GATEWAY_TOKEN_NOT_FOUND {
		names = append(names, "SGW_PAYMENT_GATEWAY_TOKEN_NOT_FOUND")
	}
	if c == SGW_PAYMENT_GATEWAY_UNAVAILABLE {
		names = append(names, "SGW_PAYMENT_GATEWAY_UNAVAILABLE")
	}
	if c == SGW_PAYMENT_INSUFFICIENT_FUNDS {
		names = append(names, "SGW_PAYMENT_INSUFFICIENT_FUNDS")
	}
	if c == SGW_PAYMENT_INVALID_ACCOUNT_NUMBER {
		names = append(names, "SGW_PAYMENT_INVALID_ACCOUNT_NUMBER")
	}
	if c == SGW_PAYMENT_INVALID_CARD_NUMBER {
		names = append(names, "SGW_PAYMENT_INVALID_CARD_NUMBER")
	}
	if c == SGW_PAYMENT_INVALID_DATA {
		names = append(names, "SGW_PAYMENT_INVALID_DATA")
	}
	if c == SGW_PAYMENT_INVALID_EMAIL {
		names = append(names, "SGW_PAYMENT_INVALID_EMAIL")
	}
	if c == SGW_PAYMENT_INVALID_GATEWAY_CONFIGURATION {
		names = append(names, "SGW_PAYMENT_INVALID_GATEWAY_CONFIGURATION")
	}
	if c == SGW_PAYMENT_INVALID_ISSUER {
		names = append(names, "SGW_PAYMENT_INVALID_ISSUER")
	}
	if c == SGW_PAYMENT_INVALID_LOGIN {
		names = append(names, "SGW_PAYMENT_INVALID_LOGIN")
	}
	if c == SGW_PAYMENT_INVALID_MERCHANT_TYPE {
		names = append(names, "SGW_PAYMENT_INVALID_MERCHANT_TYPE")
	}
	if c == SGW_PAYMENT_INVALID_TRANSACTION {
		names = append(names, "SGW_PAYMENT_INVALID_TRANSACTION")
	}
	if c == SGW_PAYMENT_ISSUER_UNAVAILABLE {
		names = append(names, "SGW_PAYMENT_ISSUER_UNAVAILABLE")
	}
	if c == SGW_PAYMENT_NO_BILLING_INFORMATION {
		names = append(names, "SGW_PAYMENT_NO_BILLING_INFORMATION")
	}
	if c == SGW_PAYMENT_NO_GATEWAY {
		names = append(names, "SGW_PAYMENT_NO_GATEWAY")
	}
	if c == SGW_PAYMENT_PARTIAL_CREDITS_NOT_SUPPORTED {
		names = append(names, "SGW_PAYMENT_PARTIAL_CREDITS_NOT_SUPPORTED")
	}
	if c == SGW_PAYMENT_PAYMENT_NOT_ACCEPTED {
		names = append(names, "SGW_PAYMENT_PAYMENT_NOT_ACCEPTED")
	}
	if c == SGW_PAYMENT_PAYPAL_ACCOUNT_ISSUE {
		names = append(names, "SGW_PAYMENT_PAYPAL_ACCOUNT_ISSUE")
	}
	if c == SGW_PAYMENT_PAYPAL_DECLINED_USE_ALTERNATE {
		names = append(names, "SGW_PAYMENT_PAYPAL_DECLINED_USE_ALTERNATE")
	}
	if c == SGW_PAYMENT_PAYPAL_HARD_DECLINE {
		names = append(names, "SGW_PAYMENT_PAYPAL_HARD_DECLINE")
	}
	if c == SGW_PAYMENT_PAYPAL_INVALID_BILLING_AGREEMENT {
		names = append(names, "SGW_PAYMENT_PAYPAL_INVALID_BILLING_AGREEMENT")
	}
	if c == SGW_PAYMENT_PAYPAL_PRIMARY_DECLINED {
		names = append(names, "SGW_PAYMENT_PAYPAL_PRIMARY_DECLINED")
	}
	if c == SGW_PAYMENT_PROCESSOR_UNAVAILABLE {
		names = append(names, "SGW_PAYMENT_PROCESSOR_UNAVAILABLE")
	}
	if c == SGW_PAYMENT_RECURLY_ERROR {
		names = append(names, "SGW_PAYMENT_RECURLY_ERROR")
	}
	if c == SGW_PAYMENT_RECURLY_FAILED_TO_GET_TOKEN {
		names = append(names, "SGW_PAYMENT_RECURLY_FAILED_TO_GET_TOKEN")
	}
	if c == SGW_PAYMENT_RECURLY_TOKEN_NOT_FOUND {
		names = append(names, "SGW_PAYMENT_RECURLY_TOKEN_NOT_FOUND")
	}
	if c == SGW_PAYMENT_RESTRICTED_CARD {
		names = append(names, "SGW_PAYMENT_RESTRICTED_CARD")
	}
	if c == SGW_PAYMENT_RESTRICTED_CARD_CHARGEBACK {
		names = append(names, "SGW_PAYMENT_RESTRICTED_CARD_CHARGEBACK")
	}
	if c == SGW_PAYMENT_SSL_ERROR {
		names = append(names, "SGW_PAYMENT_SSL_ERROR")
	}
	if c == SGW_PAYMENT_TEMPORARY_HOLD {
		names = append(names, "SGW_PAYMENT_TEMPORARY_HOLD")
	}
	if c == SGW_PAYMENT_THREE_D_SECURE_NOT_SUPPORTED {
		names = append(names, "SGW_PAYMENT_THREE_D_SECURE_NOT_SUPPORTED")
	}
	if c == SGW_PAYMENT_TOO_MANY_ATTEMPTS {
		names = append(names, "SGW_PAYMENT_TOO_MANY_ATTEMPTS")
	}
	if c == SGW_PAYMENT_TOTAL_CREDIT_EXCEEDS_CAPTURE {
		names = append(names, "SGW_PAYMENT_TOTAL_CREDIT_EXCEEDS_CAPTURE")
	}
	if c == SGW_PAYMENT_TRANSACTION_ALREADY_VOIDED {
		names = append(names, "SGW_PAYMENT_TRANSACTION_ALREADY_VOIDED")
	}
	if c == SGW_PAYMENT_TRANSACTION_CANNOT_BE_REFUNDED {
		names = append(names, "SGW_PAYMENT_TRANSACTION_CANNOT_BE_REFUNDED")
	}
	if c == SGW_PAYMENT_TRANSACTION_CANNOT_BE_VOIDED {
		names = append(names, "SGW_PAYMENT_TRANSACTION_CANNOT_BE_VOIDED")
	}
	if c == SGW_PAYMENT_TRANSACTION_FAILED_TO_SETTLE {
		names = append(names, "SGW_PAYMENT_TRANSACTION_FAILED_TO_SETTLE")
	}
	if c == SGW_PAYMENT_TRANSACTION_NOT_FOUND {
		names = append(names, "SGW_PAYMENT_TRANSACTION_NOT_FOUND")
	}
	if c == SGW_PAYMENT_TRANSACTION_SETTLED {
		names = append(names, "SGW_PAYMENT_TRANSACTION_SETTLED")
	}
	if c == SGW_PAYMENT_TRY_AGAIN {
		names = append(names, "SGW_PAYMENT_TRY_AGAIN")
	}
	if c == SGW_PAYMENT_UNKNOWN {
		names = append(names, "SGW_PAYMENT_UNKNOWN")
	}
	if c == SGW_PAYMENT_ZERO_DOLLAR_AUTH_NOT_SUPPORTED {
		names = append(names, "SGW_PAYMENT_ZERO_DOLLAR_AUTH_NOT_SUPPORTED")
	}
	if c == SGW_PLAN_ALREADY_EXISTS {
		names = append(names, "SGW_PLAN_ALREADY_EXISTS")
	}
	if c == SGW_PLAN_BAD_REQUEST {
		names = append(names, "SGW_PLAN_BAD_REQUEST")
	}
	if c == SGW_PLAN_DATE_IN_PAST {
		names = append(names, "SGW_PLAN_DATE_IN_PAST")
	}
	if c == SGW_PLAN_END_TIME_INVALID {
		names = append(names, "SGW_PLAN_END_TIME_INVALID")
	}
	if c == SGW_PLAN_FEATURE_ALREADY_EXISTS {
		names = append(names, "SGW_PLAN_FEATURE_ALREADY_EXISTS")
	}
	if c == SGW_PLAN_FEATURE_NOT_FOUND {
		names = append(names, "SGW_PLAN_FEATURE_NOT_FOUND")
	}
	if c == SGW_PLAN_FIELD_CANNOT_BE_BLANK {
		names = append(names, "SGW_PLAN_FIELD_CANNOT_BE_BLANK")
	}
	if c == SGW_PLAN_IMMUTABLE_SUBSCRIPTION {
		names = append(names, "SGW_PLAN_IMMUTABLE_SUBSCRIPTION")
	}
	if c == SGW_PLAN_INTERNAL_SERVER_ERROR {
		names = append(names, "SGW_PLAN_INTERNAL_SERVER_ERROR")
	}
	if c == SGW_PLAN_INVALID_API_KEY {
		names = append(names, "SGW_PLAN_INVALID_API_KEY")
	}
	if c == SGW_PLAN_INVALID_API_VERSION {
		names = append(names, "SGW_PLAN_INVALID_API_VERSION")
	}
	if c == SGW_PLAN_INVALID_CONFIGURATION {
		names = append(names, "SGW_PLAN_INVALID_CONFIGURATION")
	}
	if c == SGW_PLAN_INVALID_CONTENT_TYPE {
		names = append(names, "SGW_PLAN_INVALID_CONTENT_TYPE")
	}
	if c == SGW_PLAN_INVALID_DATA {
		names = append(names, "SGW_PLAN_INVALID_DATA")
	}
	if c == SGW_PLAN_INVALID_PERMISSIONS {
		names = append(names, "SGW_PLAN_INVALID_PERMISSIONS")
	}
	if c == SGW_PLAN_INVALID_TOKEN {
		names = append(names, "SGW_PLAN_INVALID_TOKEN")
	}
	if c == SGW_PLAN_INVALID_TRANSITION {
		names = append(names, "SGW_PLAN_INVALID_TRANSITION")
	}
	if c == SGW_PLAN_MANDATORY_FIELD {
		names = append(names, "SGW_PLAN_MANDATORY_FIELD")
	}
	if c == SGW_PLAN_MISSING_FEATURE {
		names = append(names, "SGW_PLAN_MISSING_FEATURE")
	}
	if c == SGW_PLAN_NOT_FOUND {
		names = append(names, "SGW_PLAN_NOT_FOUND")
	}
	if c == SGW_PLAN_RATE_LIMITED {
		names = append(names, "SGW_PLAN_RATE_LIMITED")
	}
	if c == SGW_PLAN_SIMULTANEOUS_REQUEST {
		names = append(names, "SGW_PLAN_SIMULTANEOUS_REQUEST")
	}
	if c == SGW_PLAN_TRANSACTION {
		names = append(names, "SGW_PLAN_TRANSACTION")
	}
	if c == SGW_PLAN_UNACCEPTABLE_VALUE {
		names = append(names, "SGW_PLAN_UNACCEPTABLE_VALUE")
	}
	if c == SGW_PLAN_UNAUTHORIZED {
		names = append(names, "SGW_PLAN_UNAUTHORIZED")
	}
	if c == SGW_PLAN_UNAVAILABLE_IN_API_VERSION {
		names = append(names, "SGW_PLAN_UNAVAILABLE_IN_API_VERSION")
	}
	if c == SGW_PLAN_UNKNOWN_API_VERSION {
		names = append(names, "SGW_PLAN_UNKNOWN_API_VERSION")
	}
	if c == SGW_PLAN_VALIDATION {
		names = append(names, "SGW_PLAN_VALIDATION")
	}
	if c == SGW_PLAN_VALUE_NOT_A_NUMBER {
		names = append(names, "SGW_PLAN_VALUE_NOT_A_NUMBER")
	}
	if c == SGW_PLAN_VALUE_NOT_INCLUDED_IN_LIST {
		names = append(names, "SGW_PLAN_VALUE_NOT_INCLUDED_IN_LIST")
	}
	if c == SGW_REFUND_FAILED {
		names = append(names, "SGW_REFUND_FAILED")
	}
	if c == SGW_SESSION_ALREADY_EXISTS {
		names = append(names, "SGW_SESSION_ALREADY_EXISTS")
	}
	if c == SGW_SESSION_NOT_FOUND {
		names = append(names, "SGW_SESSION_NOT_FOUND")
	}
	if c == SGW_SUBSCRIPTION_ACCOUNT_BILLING_INFO_REQUIRED {
		names = append(names, "SGW_SUBSCRIPTION_ACCOUNT_BILLING_INFO_REQUIRED")
	}
	if c == SGW_SUBSCRIPTION_ALREADY_EXISTS {
		names = append(names, "SGW_SUBSCRIPTION_ALREADY_EXISTS")
	}
	if c == SGW_SUBSCRIPTION_BAD_REQUEST {
		names = append(names, "SGW_SUBSCRIPTION_BAD_REQUEST")
	}
	if c == SGW_SUBSCRIPTION_IMMUTABLE_SUBSCRIPTION {
		names = append(names, "SGW_SUBSCRIPTION_IMMUTABLE_SUBSCRIPTION")
	}
	if c == SGW_SUBSCRIPTION_INACTIVE {
		names = append(names, "SGW_SUBSCRIPTION_INACTIVE")
	}
	if c == SGW_SUBSCRIPTION_INTERNAL_SERVER_ERROR {
		names = append(names, "SGW_SUBSCRIPTION_INTERNAL_SERVER_ERROR")
	}
	if c == SGW_SUBSCRIPTION_INVALID_API_KEY {
		names = append(names, "SGW_SUBSCRIPTION_INVALID_API_KEY")
	}
	if c == SGW_SUBSCRIPTION_INVALID_API_VERSION {
		names = append(names, "SGW_SUBSCRIPTION_INVALID_API_VERSION")
	}
	if c == SGW_SUBSCRIPTION_INVALID_CONFIGURATION {
		names = append(names, "SGW_SUBSCRIPTION_INVALID_CONFIGURATION")
	}
	if c == SGW_SUBSCRIPTION_INVALID_CONTENT_TYPE {
		names = append(names, "SGW_SUBSCRIPTION_INVALID_CONTENT_TYPE")
	}
	if c == SGW_SUBSCRIPTION_INVALID_DATA {
		names = append(names, "SGW_SUBSCRIPTION_INVALID_DATA")
	}
	if c == SGW_SUBSCRIPTION_INVALID_PERMISSIONS {
		names = append(names, "SGW_SUBSCRIPTION_INVALID_PERMISSIONS")
	}
	if c == SGW_SUBSCRIPTION_INVALID_TOKEN {
		names = append(names, "SGW_SUBSCRIPTION_INVALID_TOKEN")
	}
	if c == SGW_SUBSCRIPTION_INVALID_TRANSITION {
		names = append(names, "SGW_SUBSCRIPTION_INVALID_TRANSITION")
	}
	if c == SGW_SUBSCRIPTION_MANDATORY_FIELD {
		names = append(names, "SGW_SUBSCRIPTION_MANDATORY_FIELD")
	}
	if c == SGW_SUBSCRIPTION_MISSING_FEATURE {
		names = append(names, "SGW_SUBSCRIPTION_MISSING_FEATURE")
	}
	if c == SGW_SUBSCRIPTION_NOT_FOUND {
		names = append(names, "SGW_SUBSCRIPTION_NOT_FOUND")
	}
	if c == SGW_SUBSCRIPTION_NOT_MODIFIABLE {
		names = append(names, "SGW_SUBSCRIPTION_NOT_MODIFIABLE")
	}
	if c == SGW_SUBSCRIPTION_NOT_PREMIUM {
		names = append(names, "SGW_SUBSCRIPTION_NOT_PREMIUM")
	}
	if c == SGW_SUBSCRIPTION_RATE_LIMITED {
		names = append(names, "SGW_SUBSCRIPTION_RATE_LIMITED")
	}
	if c == SGW_SUBSCRIPTION_SIMULTANEOUS_REQUEST {
		names = append(names, "SGW_SUBSCRIPTION_SIMULTANEOUS_REQUEST")
	}
	if c == SGW_SUBSCRIPTION_STATE_INVALID {
		names = append(names, "SGW_SUBSCRIPTION_STATE_INVALID")
	}
	if c == SGW_SUBSCRIPTION_TRANSACTION {
		names = append(names, "SGW_SUBSCRIPTION_TRANSACTION")
	}
	if c == SGW_SUBSCRIPTION_TRANSACTION_DECLINE {
		names = append(names, "SGW_SUBSCRIPTION_TRANSACTION_DECLINE")
	}
	if c == SGW_SUBSCRIPTION_TRANSACTION_DECLINE_DUE_TO_3DS {
		names = append(names, "SGW_SUBSCRIPTION_TRANSACTION_DECLINE_DUE_TO_3DS")
	}
	if c == SGW_SUBSCRIPTION_UNAUTHORIZED {
		names = append(names, "SGW_SUBSCRIPTION_UNAUTHORIZED")
	}
	if c == SGW_SUBSCRIPTION_UNAVAILABLE_IN_API_VERSION {
		names = append(names, "SGW_SUBSCRIPTION_UNAVAILABLE_IN_API_VERSION")
	}
	if c == SGW_SUBSCRIPTION_UNKNOWN_API_VERSION {
		names = append(names, "SGW_SUBSCRIPTION_UNKNOWN_API_VERSION")
	}
	if c == SGW_SUBSCRIPTION_VALIDATION {
		names = append(names, "SGW_SUBSCRIPTION_VALIDATION")
	}
	if c == SGW_SUBSCRIPTION_VALUES_UNCHANGED {
		names = append(names, "SGW_SUBSCRIPTION_VALUES_UNCHANGED")
	}
	if c == SGW_SUBSCRIPTION_VALUE_NOT_A_NUMBER {
		names = append(names, "SGW_SUBSCRIPTION_VALUE_NOT_A_NUMBER")
	}
	if c == SGW_SUBSCRIPTION_VALUE_NOT_INCLUDED_IN_LIST {
		names = append(names, "SGW_SUBSCRIPTION_VALUE_NOT_INCLUDED_IN_LIST")
	}
	if c == SGW_TAX_ORDER_TRANSACTION_ID_NOT_FOUND {
		names = append(names, "SGW_TAX_ORDER_TRANSACTION_ID_NOT_FOUND")
	}
	if c == SGW_TAX_RECORD_ORDER_HAS_BEEN_CREATED {
		names = append(names, "SGW_TAX_RECORD_ORDER_HAS_BEEN_CREATED")
	}
	if c == SGW_TAX_RECORD_REFUND_ORDER_HAS_BEEN_CREATED {
		names = append(names, "SGW_TAX_RECORD_REFUND_ORDER_HAS_BEEN_CREATED")
	}
	if c == SGW_TAX_REFUND_ORDER_REFENRENCE_TRANSACTION_ID_NOT_FOUND {
		names = append(names, "SGW_TAX_REFUND_ORDER_REFENRENCE_TRANSACTION_ID_NOT_FOUND")
	}
	if c == SGW_TGW_BAD_REQUEST {
		names = append(names, "SGW_TGW_BAD_REQUEST")
	}
	if c == SGW_TGW_FORBIDDEN {
		names = append(names, "SGW_TGW_FORBIDDEN")
	}
	if c == SGW_TGW_GENERIC_EXCEPTION {
		names = append(names, "SGW_TGW_GENERIC_EXCEPTION")
	}
	if c == SGW_TGW_GONE {
		names = append(names, "SGW_TGW_GONE")
	}
	if c == SGW_TGW_INTERNAL_SERVER_ERROR {
		names = append(names, "SGW_TGW_INTERNAL_SERVER_ERROR")
	}
	if c == SGW_TGW_METHOD_NOT_ALLOWED {
		names = append(names, "SGW_TGW_METHOD_NOT_ALLOWED")
	}
	if c == SGW_TGW_NOT_ACCEPTABLE {
		names = append(names, "SGW_TGW_NOT_ACCEPTABLE")
	}
	if c == SGW_TGW_NOT_FOUND {
		names = append(names, "SGW_TGW_NOT_FOUND")
	}
	if c == SGW_TGW_SERVICE_UNAVAILABLE {
		names = append(names, "SGW_TGW_SERVICE_UNAVAILABLE")
	}
	if c == SGW_TGW_TOO_MANY_REQUESTS {
		names = append(names, "SGW_TGW_TOO_MANY_REQUESTS")
	}
	if c == SGW_TGW_UNAUTHORIZED {
		names = append(names, "SGW_TGW_UNAUTHORIZED")
	}
	if c == SGW_TGW_UNPROCESSABLE_ENTITY {
		names = append(names, "SGW_TGW_UNPROCESSABLE_ENTITY")
	}
	if c == SGW_USER_DATA_NOT_DELETED {
		names = append(names, "SGW_USER_DATA_NOT_DELETED")
	}
	if c == TERMINAL_NOT_BOUND {
		names = append(names, "TERMINAL_NOT_BOUND")
	}
	if c == TOO_MUCH_LOGIN_WITH_ONE_CODE {
		names = append(names, "TOO_MUCH_LOGIN_WITH_ONE_CODE")
	}
	if c == UNACCEPTABLE_SUBSCRIPTION_PLAN {
		names = append(names, "UNACCEPTABLE_SUBSCRIPTION_PLAN")
	}
	if c == URL_ACCOUNT_NOT_MATCH {
		names = append(names, "URL_ACCOUNT_NOT_MATCH")
	}
	if c == URL_INVALID_ENDPOINT {
		names = append(names, "URL_INVALID_ENDPOINT")
	}
	if c == URL_INVALID_TARGET_ERROR {
		names = append(names, "URL_INVALID_TARGET_ERROR")
	}
	if c == URL_IS_USED {
		names = append(names, "URL_IS_USED")
	}
	if c == URL_NOT_FOUND_OR_EXPIRED {
		names = append(names, "URL_NOT_FOUND_OR_EXPIRED")
	}
	if c == URL_SHORTURL_EXPIRED {
		names = append(names, "URL_SHORTURL_EXPIRED")
	}
	if c == URL_SHORTURL_NOT_FOUND {
		names = append(names, "URL_SHORTURL_NOT_FOUND")
	}
	if c == USER_PLACE_GEOFENCE_FAILED {
		names = append(names, "USER_PLACE_GEOFENCE_FAILED")
	}
	if c == USER_PLACE_GEOFENCE_SERVER_NOT_AVAILABLE {
		names = append(names, "USER_PLACE_GEOFENCE_SERVER_NOT_AVAILABLE")
	}
	if c == USER_PLACE_NOT_FOUND {
		names = append(names, "USER_PLACE_NOT_FOUND")
	}
	if c == USER_PROFILE_CONTAINS_INVALID_PLACE_ID {
		names = append(names, "USER_PROFILE_CONTAINS_INVALID_PLACE_ID")
	}
	if c == USER_PROFILE_GEOFENCE_FAILED {
		names = append(names, "USER_PROFILE_GEOFENCE_FAILED")
	}
	if c == USER_PROFILE_GEOFENCE_SERVER_NOT_AVAILABLE {
		names = append(names, "USER_PROFILE_GEOFENCE_SERVER_NOT_AVAILABLE")
	}
	if c == USER_PROFILE_NOT_FOUND {
		names = append(names, "USER_PROFILE_NOT_FOUND")
	}
	if c == USER_PROFILE_PLACE_ID_VALIDATION_FAILED {
		names = append(names, "USER_PROFILE_PLACE_ID_VALIDATION_FAILED")
	}
	if c == USER_PROFILE_TERMINAL_ASSOCIATION_EXISTS {
		names = append(names, "USER_PROFILE_TERMINAL_ASSOCIATION_EXISTS")
	}
	if c == USER_PROFILE_TERMINAL_ID_NOT_FOUND {
		names = append(names, "USER_PROFILE_TERMINAL_ID_NOT_FOUND")
	}
	if c == USER_PROFILE_TERMINAL_NOT_FOUND {
		names = append(names, "USER_PROFILE_TERMINAL_NOT_FOUND")
	}
	if c == VA_VIDEO_ANALYTICS_NOT_ENABLED {
		names = append(names, "VA_VIDEO_ANALYTICS_NOT_ENABLED")
	}
	if c == VA_VIDEO_ANALYTICS_WOWZA_ERROR {
		names = append(names, "VA_VIDEO_ANALYTICS_WOWZA_ERROR")
	}
	if c == VA_VIDEO_SUMMARY_AUTH_FAILED {
		names = append(names, "VA_VIDEO_SUMMARY_AUTH_FAILED")
	}
	if c == VA_VIDEO_SUMMARY_EVENT_FAILED {
		names = append(names, "VA_VIDEO_SUMMARY_EVENT_FAILED")
	}
	if c == VA_VIDEO_SUMMARY_INVALID_PAGINATOR {
		names = append(names, "VA_VIDEO_SUMMARY_INVALID_PAGINATOR")
	}
	if c == VA_VIDEO_SUMMARY_IN_PROGRESS {
		names = append(names, "VA_VIDEO_SUMMARY_IN_PROGRESS")
	}
	if c == VA_VIDEO_SUMMARY_NOT_ENABLED {
		names = append(names, "VA_VIDEO_SUMMARY_NOT_ENABLED")
	}
	if c == VA_VIDEO_SUMMARY_NOT_GENERATED {
		names = append(names, "VA_VIDEO_SUMMARY_NOT_GENERATED")
	}
	if c == VA_VIDEO_SUMMARY_WRONG_TYPE {
		names = append(names, "VA_VIDEO_SUMMARY_WRONG_TYPE")
	}
	if c == WEBRTC_CLIENT_EXISTS {
		names = append(names, "WEBRTC_CLIENT_EXISTS")
	}
	if c == WEBRTC_CLIENT_NOT_FOUND {
		names = append(names, "WEBRTC_CLIENT_NOT_FOUND")
	}
	if c == WEBRTC_CONFERENCE_NOT_FOUND {
		names = append(names, "WEBRTC_CONFERENCE_NOT_FOUND")
	}
	if c == WEBRTC_CONFERENCE_NOT_INITIALIZED {
		names = append(names, "WEBRTC_CONFERENCE_NOT_INITIALIZED")
	}
	if c == WEBRTC_KURENTO_SERVICE_UNAVAILABLE {
		names = append(names, "WEBRTC_KURENTO_SERVICE_UNAVAILABLE")
	}
	if c == WEBRTC_RECORDING_ALREADY {
		names = append(names, "WEBRTC_RECORDING_ALREADY")
	}
	if c == WEBRTC_RECORDING_NOT_READY {
		names = append(names, "WEBRTC_RECORDING_NOT_READY")
	}
	if c == WEBRTC_RECOVERING {
		names = append(names, "WEBRTC_RECOVERING")
	}
	if c == WEBSOCKET_ACCOUNT_NOT_BOUND {
		names = append(names, "WEBSOCKET_ACCOUNT_NOT_BOUND")
	}
	if c == WEBSOCKET_CONNECTION_INFO_NOT_FOUND {
		names = append(names, "WEBSOCKET_CONNECTION_INFO_NOT_FOUND")
	}
	if c == WEBSOCKET_FAILED_TO_CONNECT_TO_APIGATEWAY {
		names = append(names, "WEBSOCKET_FAILED_TO_CONNECT_TO_APIGATEWAY")
	}
	if c == WEBSOCKET_FAILED_TO_CREATE_CACHE_PROVIDER {
		names = append(names, "WEBSOCKET_FAILED_TO_CREATE_CACHE_PROVIDER")
	}
	if c == WEBSOCKET_FAILED_TO_NOTIFY_ALL {
		names = append(names, "WEBSOCKET_FAILED_TO_NOTIFY_ALL")
	}
	if c == WEBSOCKET_FAILED_TO_PROXY {
		names = append(names, "WEBSOCKET_FAILED_TO_PROXY")
	}
	if c == WEBSOCKET_PASSTHROUGH_TIMEOUT {
		names = append(names, "WEBSOCKET_PASSTHROUGH_TIMEOUT")
	}
	if c == WL_UNKNOWN_ROLE_TYPE {
		names = append(names, "WL_UNKNOWN_ROLE_TYPE")
	}
	if c == WL_UNKNOWN_STATUS_TYPE {
		names = append(names, "WL_UNKNOWN_STATUS_TYPE")
	}
	if c == ZB_UNDOCUMENTED_ERROR {
		names = append(names, "ZB_UNDOCUMENTED_ERROR")
	}
	return names
}
