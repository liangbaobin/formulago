// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SysApisColumns holds the columns for the "sys_apis" table.
	SysApisColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true, Comment: "primary key"},
		{Name: "created_at", Type: field.TypeTime, Comment: "created time"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "last update time"},
		{Name: "path", Type: field.TypeString, Comment: "API path | API 路径"},
		{Name: "description", Type: field.TypeString, Comment: "API description | API 描述"},
		{Name: "api_group", Type: field.TypeString, Comment: "API group | API 分组"},
		{Name: "method", Type: field.TypeString, Comment: "HTTP method | HTTP 请求类型", Default: "POST"},
	}
	// SysApisTable holds the schema information for the "sys_apis" table.
	SysApisTable = &schema.Table{
		Name:       "sys_apis",
		Columns:    SysApisColumns,
		PrimaryKey: []*schema.Column{SysApisColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "api_path_method",
				Unique:  true,
				Columns: []*schema.Column{SysApisColumns[3], SysApisColumns[6]},
			},
		},
	}
	// SysDictionariesColumns holds the columns for the "sys_dictionaries" table.
	SysDictionariesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true, Comment: "primary key"},
		{Name: "created_at", Type: field.TypeTime, Comment: "created time"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "last update time"},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Comment: "status 1 normal 0 ban | 状态 1 正常 0 禁用", Default: 1},
		{Name: "title", Type: field.TypeString, Comment: "the title shown in the ui | 展示名称 （建议配合i18n）"},
		{Name: "name", Type: field.TypeString, Unique: true, Comment: "the name of dictionary for search | 字典搜索名称"},
		{Name: "description", Type: field.TypeString, Comment: "the description of dictionary | 字典描述"},
	}
	// SysDictionariesTable holds the schema information for the "sys_dictionaries" table.
	SysDictionariesTable = &schema.Table{
		Name:       "sys_dictionaries",
		Columns:    SysDictionariesColumns,
		PrimaryKey: []*schema.Column{SysDictionariesColumns[0]},
	}
	// SysDictionaryDetailsColumns holds the columns for the "sys_dictionary_details" table.
	SysDictionaryDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true, Comment: "primary key"},
		{Name: "created_at", Type: field.TypeTime, Comment: "created time"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "last update time"},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Comment: "status 1 normal 0 ban | 状态 1 正常 0 禁用", Default: 1},
		{Name: "title", Type: field.TypeString, Comment: "the title shown in the ui | 展示名称 （建议配合i18n）"},
		{Name: "key", Type: field.TypeString, Comment: "key | 键"},
		{Name: "value", Type: field.TypeString, Comment: "value | 值"},
		{Name: "dictionary_id", Type: field.TypeUint64, Nullable: true},
	}
	// SysDictionaryDetailsTable holds the schema information for the "sys_dictionary_details" table.
	SysDictionaryDetailsTable = &schema.Table{
		Name:       "sys_dictionary_details",
		Columns:    SysDictionaryDetailsColumns,
		PrimaryKey: []*schema.Column{SysDictionaryDetailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_dictionary_details_sys_dictionaries_dictionary_details",
				Columns:    []*schema.Column{SysDictionaryDetailsColumns[7]},
				RefColumns: []*schema.Column{SysDictionariesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "dictionarydetail_key_dictionary_id",
				Unique:  true,
				Columns: []*schema.Column{SysDictionaryDetailsColumns[5], SysDictionaryDetailsColumns[7]},
			},
		},
	}
	// SysLogsColumns holds the columns for the "sys_logs" table.
	SysLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true, Comment: "primary key"},
		{Name: "created_at", Type: field.TypeTime, Comment: "created time"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "last update time"},
		{Name: "type", Type: field.TypeString, Comment: "type of log | 日志类型"},
		{Name: "method", Type: field.TypeString, Comment: "method of log | 日志请求方法"},
		{Name: "api", Type: field.TypeString, Comment: "api of log | 日志请求api"},
		{Name: "success", Type: field.TypeBool, Comment: "success of log | 日志请求是否成功"},
		{Name: "req_content", Type: field.TypeString, Nullable: true, Comment: "content of request log | 日志请求内容", SchemaType: map[string]string{"mysql": "varchar(512)"}},
		{Name: "resp_content", Type: field.TypeString, Nullable: true, Comment: "content of response log | 日志返回内容", SchemaType: map[string]string{"mysql": "varchar(512)"}},
		{Name: "ip", Type: field.TypeString, Nullable: true, Comment: "ip of log | 日志IP"},
		{Name: "user_agent", Type: field.TypeString, Nullable: true, Comment: "user_agent of log | 日志用户客户端"},
		{Name: "operator", Type: field.TypeString, Nullable: true, Comment: "operator of log | 日志操作者"},
		{Name: "time", Type: field.TypeInt, Nullable: true, Comment: "time of log(millisecond) | 日志时间(毫秒)"},
	}
	// SysLogsTable holds the schema information for the "sys_logs" table.
	SysLogsTable = &schema.Table{
		Name:       "sys_logs",
		Columns:    SysLogsColumns,
		PrimaryKey: []*schema.Column{SysLogsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "logs_api",
				Unique:  false,
				Columns: []*schema.Column{SysLogsColumns[5]},
			},
		},
	}
	// SysMenusColumns holds the columns for the "sys_menus" table.
	SysMenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true, Comment: "primary key"},
		{Name: "created_at", Type: field.TypeTime, Comment: "created time"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "last update time"},
		{Name: "menu_level", Type: field.TypeUint32, Comment: "menu level | 菜单层级"},
		{Name: "menu_type", Type: field.TypeUint32, Comment: "menu type | 菜单类型 0 目录 1 菜单 2 按钮"},
		{Name: "path", Type: field.TypeString, Nullable: true, Comment: "index path | 菜单路由路径", Default: ""},
		{Name: "name", Type: field.TypeString, Comment: "index name | 菜单名称"},
		{Name: "redirect", Type: field.TypeString, Nullable: true, Comment: "redirect path | 跳转路径 （外链）", Default: ""},
		{Name: "component", Type: field.TypeString, Nullable: true, Comment: "the path of vue file | 组件路径", Default: ""},
		{Name: "order_no", Type: field.TypeUint32, Comment: "sorting numbers | 排序编号", Default: 0},
		{Name: "disabled", Type: field.TypeBool, Nullable: true, Comment: "disable status | 是否停用", Default: false},
		{Name: "title", Type: field.TypeString, Comment: "menu name | 菜单显示标题"},
		{Name: "icon", Type: field.TypeString, Comment: "menu icon | 菜单图标"},
		{Name: "hide_menu", Type: field.TypeBool, Nullable: true, Comment: "hide menu | 是否隐藏菜单", Default: false},
		{Name: "hide_breadcrumb", Type: field.TypeBool, Nullable: true, Comment: "hide the breadcrumb | 隐藏面包屑", Default: false},
		{Name: "current_active_menu", Type: field.TypeString, Nullable: true, Comment: "set the active menu | 激活菜单", Default: ""},
		{Name: "ignore_keep_alive", Type: field.TypeBool, Nullable: true, Comment: "do not keep alive the tab | 取消页面缓存", Default: false},
		{Name: "hide_tab", Type: field.TypeBool, Nullable: true, Comment: "hide the tab header | 隐藏页头", Default: false},
		{Name: "frame_src", Type: field.TypeString, Nullable: true, Comment: "show iframe | 内嵌 iframe", Default: ""},
		{Name: "carry_param", Type: field.TypeBool, Nullable: true, Comment: "the route carries parameters or not | 携带参数", Default: false},
		{Name: "hide_children_in_menu", Type: field.TypeBool, Nullable: true, Comment: "hide children menu or not | 隐藏所有子菜单", Default: false},
		{Name: "affix", Type: field.TypeBool, Nullable: true, Comment: "affix tab | Tab 固定", Default: false},
		{Name: "dynamic_level", Type: field.TypeUint32, Nullable: true, Comment: "the maximum number of pages the router can open | 能打开的子TAB数", Default: 20},
		{Name: "real_path", Type: field.TypeString, Nullable: true, Comment: "the real path of the route without dynamic part | 菜单路由不包含参数部分", Default: ""},
		{Name: "parent_id", Type: field.TypeUint64, Nullable: true},
	}
	// SysMenusTable holds the schema information for the "sys_menus" table.
	SysMenusTable = &schema.Table{
		Name:       "sys_menus",
		Columns:    SysMenusColumns,
		PrimaryKey: []*schema.Column{SysMenusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_menus_sys_menus_children",
				Columns:    []*schema.Column{SysMenusColumns[24]},
				RefColumns: []*schema.Column{SysMenusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SysMenuParamsColumns holds the columns for the "sys_menu_params" table.
	SysMenuParamsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true, Comment: "primary key"},
		{Name: "created_at", Type: field.TypeTime, Comment: "created time"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "last update time"},
		{Name: "type", Type: field.TypeString, Comment: "pass parameters via params or query | 参数类型"},
		{Name: "key", Type: field.TypeString, Comment: "the key of parameters | 参数键"},
		{Name: "value", Type: field.TypeString, Comment: "the value of parameters | 参数值"},
		{Name: "menu_params", Type: field.TypeUint64, Nullable: true},
	}
	// SysMenuParamsTable holds the schema information for the "sys_menu_params" table.
	SysMenuParamsTable = &schema.Table{
		Name:       "sys_menu_params",
		Columns:    SysMenuParamsColumns,
		PrimaryKey: []*schema.Column{SysMenuParamsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_menu_params_sys_menus_params",
				Columns:    []*schema.Column{SysMenuParamsColumns[6]},
				RefColumns: []*schema.Column{SysMenusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SysOauthProvidersColumns holds the columns for the "sys_oauth_providers" table.
	SysOauthProvidersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true, Comment: "primary key"},
		{Name: "created_at", Type: field.TypeTime, Comment: "created time"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "last update time"},
		{Name: "name", Type: field.TypeString, Unique: true, Comment: "the provider's name | 提供商名称"},
		{Name: "app_id", Type: field.TypeString, Nullable: true, Comment: "the app id | 应用id"},
		{Name: "client_id", Type: field.TypeString, Comment: "the client id | 客户端 id"},
		{Name: "client_secret", Type: field.TypeString, Nullable: true, Comment: "the client secret | 客户端密钥"},
		{Name: "redirect_url", Type: field.TypeString, Comment: "the redirect url | 跳转地址"},
		{Name: "scopes", Type: field.TypeString, Nullable: true, Comment: "the scopes | 权限范围"},
		{Name: "auth_url", Type: field.TypeString, Comment: "the auth url of the provider | 认证地址"},
		{Name: "token_url", Type: field.TypeString, Nullable: true, Comment: "the token url of the provider | 获取 token地址"},
		{Name: "auth_style", Type: field.TypeUint64, Nullable: true, Comment: "the auth style, 0: auto detect; 1: third party login; 2: login with username and password"},
		{Name: "info_url", Type: field.TypeString, Nullable: true, Comment: "the URL to request user information by token | 用户信息请求地址"},
	}
	// SysOauthProvidersTable holds the schema information for the "sys_oauth_providers" table.
	SysOauthProvidersTable = &schema.Table{
		Name:       "sys_oauth_providers",
		Columns:    SysOauthProvidersColumns,
		PrimaryKey: []*schema.Column{SysOauthProvidersColumns[0]},
	}
	// SysRolesColumns holds the columns for the "sys_roles" table.
	SysRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true, Comment: "primary key"},
		{Name: "created_at", Type: field.TypeTime, Comment: "created time"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "last update time"},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Comment: "status 1 normal 0 ban | 状态 1 正常 0 禁用", Default: 1},
		{Name: "name", Type: field.TypeString, Comment: "role name | 角色名"},
		{Name: "value", Type: field.TypeString, Unique: true, Comment: "role value for permission control in front end | 角色值，用于前端权限控制"},
		{Name: "default_router", Type: field.TypeString, Comment: "default menu : dashboard | 默认登录页面", Default: "dashboard"},
		{Name: "remark", Type: field.TypeString, Comment: "remark | 备注", Default: ""},
		{Name: "order_no", Type: field.TypeUint32, Comment: "order number | 排序编号", Default: 0},
	}
	// SysRolesTable holds the schema information for the "sys_roles" table.
	SysRolesTable = &schema.Table{
		Name:       "sys_roles",
		Columns:    SysRolesColumns,
		PrimaryKey: []*schema.Column{SysRolesColumns[0]},
	}
	// SysTokensColumns holds the columns for the "sys_tokens" table.
	SysTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "created time"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "last update time"},
		{Name: "user_id", Type: field.TypeUint64, Unique: true, Comment: " User's ID | 用户的ID"},
		{Name: "token", Type: field.TypeString, Comment: "Token string | Token 字符串"},
		{Name: "source", Type: field.TypeString, Comment: "Log in source such as GitHub | Token 来源 （本地为core, 第三方如github等）"},
		{Name: "expired_at", Type: field.TypeTime, Comment: " Expire time | 过期时间"},
		{Name: "user_token", Type: field.TypeUint64, Unique: true, Nullable: true},
	}
	// SysTokensTable holds the schema information for the "sys_tokens" table.
	SysTokensTable = &schema.Table{
		Name:       "sys_tokens",
		Columns:    SysTokensColumns,
		PrimaryKey: []*schema.Column{SysTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_tokens_sys_users_token",
				Columns:    []*schema.Column{SysTokensColumns[7]},
				RefColumns: []*schema.Column{SysUsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "token_user_id",
				Unique:  false,
				Columns: []*schema.Column{SysTokensColumns[3]},
			},
			{
				Name:    "token_expired_at",
				Unique:  false,
				Columns: []*schema.Column{SysTokensColumns[6]},
			},
		},
	}
	// SysUsersColumns holds the columns for the "sys_users" table.
	SysUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true, Comment: "primary key"},
		{Name: "created_at", Type: field.TypeTime, Comment: "created time"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "last update time"},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Comment: "status 1 normal 0 ban | 状态 1 正常 0 禁用", Default: 1},
		{Name: "username", Type: field.TypeString, Unique: true, Comment: "user's login name | 登录名"},
		{Name: "password", Type: field.TypeString, Comment: "password | 密码"},
		{Name: "nickname", Type: field.TypeString, Unique: true, Comment: "nickname | 昵称"},
		{Name: "side_mode", Type: field.TypeString, Nullable: true, Comment: "template mode | 布局方式", Default: "dark"},
		{Name: "base_color", Type: field.TypeString, Nullable: true, Comment: "base color of template | 后台页面色调", Default: "#fff"},
		{Name: "active_color", Type: field.TypeString, Nullable: true, Comment: "active color of template | 当前激活的颜色设定", Default: "#1890ff"},
		{Name: "role_id", Type: field.TypeUint64, Nullable: true, Comment: "role id | 角色ID", Default: 2},
		{Name: "mobile", Type: field.TypeString, Unique: true, Comment: "mobile number | 手机号"},
		{Name: "email", Type: field.TypeString, Nullable: true, Comment: "email | 邮箱号"},
		{Name: "wecom", Type: field.TypeString, Nullable: true, Comment: "wecom | 企业微信号"},
		{Name: "avatar", Type: field.TypeString, Nullable: true, Comment: "avatar | 头像路径", Default: "", SchemaType: map[string]string{"mysql": "varchar(512)"}},
	}
	// SysUsersTable holds the schema information for the "sys_users" table.
	SysUsersTable = &schema.Table{
		Name:       "sys_users",
		Columns:    SysUsersColumns,
		PrimaryKey: []*schema.Column{SysUsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_username_email",
				Unique:  true,
				Columns: []*schema.Column{SysUsersColumns[4], SysUsersColumns[12]},
			},
		},
	}
	// RoleMenusColumns holds the columns for the "role_menus" table.
	RoleMenusColumns = []*schema.Column{
		{Name: "role_id", Type: field.TypeUint64},
		{Name: "menu_id", Type: field.TypeUint64},
	}
	// RoleMenusTable holds the schema information for the "role_menus" table.
	RoleMenusTable = &schema.Table{
		Name:       "role_menus",
		Columns:    RoleMenusColumns,
		PrimaryKey: []*schema.Column{RoleMenusColumns[0], RoleMenusColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_menus_role_id",
				Columns:    []*schema.Column{RoleMenusColumns[0]},
				RefColumns: []*schema.Column{SysRolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "role_menus_menu_id",
				Columns:    []*schema.Column{RoleMenusColumns[1]},
				RefColumns: []*schema.Column{SysMenusColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SysApisTable,
		SysDictionariesTable,
		SysDictionaryDetailsTable,
		SysLogsTable,
		SysMenusTable,
		SysMenuParamsTable,
		SysOauthProvidersTable,
		SysRolesTable,
		SysTokensTable,
		SysUsersTable,
		RoleMenusTable,
	}
)

func init() {
	SysApisTable.Annotation = &entsql.Annotation{
		Table: "sys_apis",
	}
	SysDictionariesTable.Annotation = &entsql.Annotation{
		Table: "sys_dictionaries",
	}
	SysDictionaryDetailsTable.ForeignKeys[0].RefTable = SysDictionariesTable
	SysDictionaryDetailsTable.Annotation = &entsql.Annotation{
		Table: "sys_dictionary_details",
	}
	SysLogsTable.Annotation = &entsql.Annotation{
		Table: "sys_logs",
	}
	SysMenusTable.ForeignKeys[0].RefTable = SysMenusTable
	SysMenusTable.Annotation = &entsql.Annotation{
		Table: "sys_menus",
	}
	SysMenuParamsTable.ForeignKeys[0].RefTable = SysMenusTable
	SysMenuParamsTable.Annotation = &entsql.Annotation{
		Table: "sys_menu_params",
	}
	SysOauthProvidersTable.Annotation = &entsql.Annotation{
		Table: "sys_oauth_providers",
	}
	SysRolesTable.Annotation = &entsql.Annotation{
		Table: "sys_roles",
	}
	SysTokensTable.ForeignKeys[0].RefTable = SysUsersTable
	SysTokensTable.Annotation = &entsql.Annotation{
		Table: "sys_tokens",
	}
	SysUsersTable.Annotation = &entsql.Annotation{
		Table: "sys_users",
	}
	RoleMenusTable.ForeignKeys[0].RefTable = SysRolesTable
	RoleMenusTable.ForeignKeys[1].RefTable = SysMenusTable
}
