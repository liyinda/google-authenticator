package errno

var (
	OK = &Errno{Code: 20000, Message: "OK"}
	// Common error
	ErrInternalServer = &Errno{Code: 40001, Message: "Internal server error."}
	ErrBind           = &Errno{Code: 40002, Message: "Error occurred while binding the request body to the struct."}
	ErrForbidden      = &Errno{Code: 40003, Message: "You don't have permission to access"}
	ErrNotFound       = &Errno{Code: 40004, Message: "Not found!"}
	ErrNotLogin       = &Errno{Code: 40005, Message: "Not login!"}
	ErrToken          = &Errno{Code: 40006, Message: "check token failed!"}
	ErrQuery          = &Errno{Code: 40007, Message: "Error parsing Query"}
	ErrTpye           = &Errno{Code: 40008, Message: "Error parameter Type"}
	ErrStruct         = &Errno{Code: 40009, Message: "Error struct's json"}
	ErrSession        = &Errno{Code: 40010, Message: "check session failed!"}
	ErrLDAP           = &Errno{Code: 40010, Message: "check ldap failed!"}

	ErrProjectCreate           = &Errno{Code: 41001, Message: "Create Project failed."}
	ErrProjectGet              = &Errno{Code: 41002, Message: "Get Project  info failed."}
	ErrProjectUpdate           = &Errno{Code: 41003, Message: "Update Project  failed."}
	ErrProjectDelete           = &Errno{Code: 41004, Message: "Delete Project  failed."}
	ErrProjectList             = &Errno{Code: 41005, Message: "Get Project  List failed."}
	ErrProjectIsExists         = &Errno{Code: 41006, Message: "Project Is Exists."}
	ErrProjectNotExists        = &Errno{Code: 41007, Message: "Project Not Exists."}
	ErrProjectHistory          = &Errno{Code: 41008, Message: "Get Project History failed."}
	ErrProjectPodsInfo         = &Errno{Code: 41009, Message: "Get Project Pods Info failed."}
	ErrProjectEvents           = &Errno{Code: 41010, Message: "Get Project Events failed."}
	ErrProjectStop             = &Errno{Code: 41011, Message: "Stop Project failed."}
	ErrProjectPodsLog          = &Errno{Code: 41012, Message: "Get Project Pods Logs failed."}
	ErrProjectDomain           = &Errno{Code: 41013, Message: "Project Domain Collision"}
	ErrProjectPublish          = &Errno{Code: 41014, Message: "Publish Deployment  failed."}
	ErrProjectHavePodsRunning  = &Errno{Code: 41015, Message: "Have one or more pods running."}
	ErrProjectPodsStatusUpdate = &Errno{Code: 41016, Message: "Update Project Pods Status failed."}
	ErrProjectClusterUpdate    = &Errno{Code: 41017, Message: "Update Project Cluster failed."}
	ErrProjectHistoryAnalyze   = &Errno{Code: 41018, Message: "Get Project History analyze failed."}
	ErrProjectActive           = &Errno{Code: 41019, Message: "Get Project active failed."}

	ErrBuildCheck      = &Errno{Code: 42000, Message: "Check Build number failed, in 15 minutes, unfinished build count >= 3"}
	ErrBuildCreate     = &Errno{Code: 42001, Message: "Create Build failed."}
	ErrBuildGet        = &Errno{Code: 42002, Message: "Get Build  info failed."}
	ErrBuildUpdate     = &Errno{Code: 42003, Message: "Update Build  failed."}
	ErrBuildDelete     = &Errno{Code: 42004, Message: "Delete Build  failed."}
	ErrBuildList       = &Errno{Code: 42005, Message: "Get Build  List failed."}
	ErrBuildType       = &Errno{Code: 42006, Message: "Unknown Build  type."}
	ErrGit             = &Errno{Code: 42007, Message: "Clone project failed."}
	ErrGitBranch       = &Errno{Code: 42008, Message: "Get project branch failed."}
	ErrBuildK8s        = &Errno{Code: 42009, Message: "Build k8s project failed."}
	ErrBuildHost       = &Errno{Code: 42010, Message: "Build host project failed."}
	ErrBuildAnalyze    = &Errno{Code: 42011, Message: "Get Build analyze failed."}
	ErrBuildLock       = &Errno{Code: 42012, Message: "Service update now, plz wait a moment"}
	ErrBuildLockCreate = &Errno{Code: 42013, Message: "Create Build Lock failed."}
	ErrBuildLockGet    = &Errno{Code: 42014, Message: "Get Build Lock failed."}
	ErrBuildLockDelete = &Errno{Code: 42015, Message: "Delete Build Lock failed."}

	ErrUserCreate    = &Errno{Code: 4401, Message: "Create User failed."}
	ErrUserGet       = &Errno{Code: 4402, Message: "Get User info failed."}
	ErrUserUpdate    = &Errno{Code: 4403, Message: "Update User failed."}
	ErrUserDelete    = &Errno{Code: 4404, Message: "Delete User failed."}
	ErrUserList      = &Errno{Code: 4405, Message: "Get User List failed."}
	ErrUserTokenList = &Errno{Code: 4406, Message: "Get User Token List failed."}
	ErrUserTokenSet  = &Errno{Code: 4407, Message: "Set User Token failed."}
	ErrUserTokenDel  = &Errno{Code: 4408, Message: "Delete User Token failed."}

	ErrGroupCreate        = &Errno{Code: 45001, Message: "Create GroupName failed."}
	ErrGroupGet           = &Errno{Code: 45002, Message: "Get GroupName  info failed."}
	ErrGroupUpdate        = &Errno{Code: 45003, Message: "Update GroupName  failed."}
	ErrGroupDelete        = &Errno{Code: 45004, Message: "Delete GroupName  failed."}
	ErrGroupList          = &Errno{Code: 45005, Message: "Get GroupName  List failed."}
	ErrGroupIsExists      = &Errno{Code: 45006, Message: "GroupName Is Exists."}
	ErrGroupNotExists     = &Errno{Code: 45007, Message: "Group Not Exists."}
	ErrGroupRoleIsExists  = &Errno{Code: 45008, Message: "Group Role Is Exists."}
	ErrGroupRoleNotExists = &Errno{Code: 45009, Message: "Group Role Not Exists."}
	ErrGroupNotEmpty      = &Errno{Code: 45010, Message: "Group Not Empty."}
	ErrGroupHasChild      = &Errno{Code: 45011, Message: "Group Has Child Groups."}
	ErrGroupType          = &Errno{Code: 45011, Message: "Group Type Unknown."}

	ErrClusterIsExists  = &Errno{Code: 46001, Message: "Cluster Is Exists."}
	ErrClusterNotExists = &Errno{Code: 46002, Message: "Cluster Not Exists."}
	ErrClusterGet       = &Errno{Code: 46003, Message: "Get Cluster failed."}
	ErrClusterCreate    = &Errno{Code: 46004, Message: "Create Cluster failed."}
	ErrClusterUpdate    = &Errno{Code: 46005, Message: "Update Cluster failed."}

	ErrMonitoringCreate    = &Errno{Code: 47001, Message: "Create Monitoring failed."}
	ErrMonitoringGet       = &Errno{Code: 47002, Message: "Get Monitoring info failed."}
	ErrMonitoringUpdate    = &Errno{Code: 47003, Message: "Update Monitoring failed."}
	ErrMonitoringDelete    = &Errno{Code: 47004, Message: "Delete Monitoring failed."}
	ErrMonitoringList      = &Errno{Code: 47005, Message: "Get Monitoring List failed."}
	ErrMonitoringIsExists  = &Errno{Code: 47006, Message: "Monitoring Is Exists."}
	ErrMonitoringNotExists = &Errno{Code: 47007, Message: "Monitoring Not Exists."}

	ErrApplicationCreate    = &Errno{Code: 48001, Message: "Create Application failed."}
	ErrApplicationGet       = &Errno{Code: 48002, Message: "Get Application info failed."}
	ErrApplicationUpdate    = &Errno{Code: 48003, Message: "Update Application failed."}
	ErrApplicationDelete    = &Errno{Code: 48004, Message: "Delete Application failed."}
	ErrApplicationList      = &Errno{Code: 48005, Message: "Get Application List failed."}
	ErrApplicationIsExists  = &Errno{Code: 48006, Message: "Application Is Exists."}
	ErrApplicationNotExists = &Errno{Code: 480007, Message: "Application Not Exists."}
	ErrApplicationStatus    = &Errno{Code: 48008, Message: "Get Application Status failed."}
	ErrApplicationStop      = &Errno{Code: 48009, Message: "Stop Application failed."}
	ErrApplicationStart     = &Errno{Code: 48010, Message: "Start Application failed."}
	ErrStoreList            = &Errno{Code: 48011, Message: "Get store template list failed."}
	ErrStoreParam           = &Errno{Code: 48012, Message: "Get store template param failed."}

	ErrChangeLogCreate = &Errno{Code: 49001, Message: "Create Change Log failed."}
	ErrChangeLogGet    = &Errno{Code: 49002, Message: "Get Change Log failed."}
	ErrChangeLogUpdate = &Errno{Code: 49003, Message: "Update Change Log failed."}

	ErrCostEventsGet = &Errno{Code: 50001, Message: "Get Cost Events failed."}
)

const (
	SUCCESS        = 20000
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_EXIST_TAG       = 10001
	ERROR_EXIST_TAG_FAIL  = 10002
	ERROR_NOT_EXIST_TAG   = 10003
	ERROR_GET_TAGS_FAIL   = 10004
	ERROR_COUNT_TAG_FAIL  = 10005
	ERROR_ADD_TAG_FAIL    = 10006
	ERROR_EDIT_TAG_FAIL   = 10007
	ERROR_DELETE_TAG_FAIL = 10008
	ERROR_EXPORT_TAG_FAIL = 10009
	ERROR_IMPORT_TAG_FAIL = 10010

	ERROR_NOT_EXIST_ARTICLE        = 10011
	ERROR_CHECK_EXIST_ARTICLE_FAIL = 10012
	ERROR_ADD_ARTICLE_FAIL         = 10013
	ERROR_DELETE_ARTICLE_FAIL      = 10014
	ERROR_EDIT_ARTICLE_FAIL        = 10015
	ERROR_COUNT_ARTICLE_FAIL       = 10016
	ERROR_GET_ARTICLES_FAIL        = 10017
	ERROR_GET_ARTICLE_FAIL         = 10018
	ERROR_GEN_ARTICLE_POSTER_FAIL  = 10019
	ERROR_NOT_JSON                 = 10020
	ERROR_NOT_EXIST_USER           = 10021

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
	ERROR_AUTH_SESSION             = 20005

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 30001
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 30002
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003
)
