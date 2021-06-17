create table `user`
(
    id            int unsigned auto_increment primary key,
    user_name     varchar(32) not null comment '用户名',
    password      varchar(64) not null comment '密码',
    created_time  int unsigned default 0 null comment '创建时间',
    modified_time int unsigned default 0 null comment '修改时间',
    role          tinyint unsigned default 0 null comment '角色， 0为测试、1为开发',
    state         tinyint unsigned default 1 null comment '状态 0为禁用、1为启用'
) comment '用户表' charset = utf8;

create table `project`
(
    id            int unsigned auto_increment primary key,
    name          varchar(32) not null comment '项目名称',
    remake        varchar(100) comment '备注',
    created_by    int unsigned not null comment '创建人',
    modified_by   int unsignednull comment '修改人',
    created_time  int unsigned default 0 null comment '创建时间',
    modified_time int unsigned default 0 null comment '修改时间',
    state         tinyint unsigned default 1 null comment '状态 0为禁用、1为启用'
) comment '项目表' charset = utf8;

create table `project_user`
(
    `id`         int unsigned auto_increment primary key,
    `user_id`    int unsigned not null comment '用户id',
    `project_id` int unsigned not null comment '项目id'
)comment '项目所属用户' charset = utf8;

create table `environment`
(
    `id`            int unsigned auto_increment primary key,
    `name`          varchar(32)  not null comment '环境名称',
    `domain`        varchar(500) not null comment '环境域名',
    `headers`       text comment '环境请求头',
    `variables`     text comment '环境变量',
    `created_by`    int unsigned not null comment '创建人',
    `modified_by`   int unsignednull comment '修改人',
    `created_time`  int unsigned default 0 null comment '创建时间',
    `modified_time` int unsigned default 0 null comment '修改时间',
    `state`         tinyint unsigned default 1 null comment '状态 0为禁用、1为启用',
    `project_id`    int          not null comment '所属项目'
) comment '环境表' charset = utf8;

create table `module`
(
    `id`            int unsigned auto_increment primary key,
    `name`          varchar(32) not null comment '模块名称',
    `created_by`    int unsigned not null comment '创建人',
    `modified_by`   int unsignednull comment '修改人',
    `created_time`  int unsigned default 0 null comment '创建时间',
    `modified_time` int unsigned default 0 null comment '修改时间',
    `state`         tinyint unsigned default 1 null comment '状态 0为禁用、1为启用',
    `project_id`    int         not null comment '所属项目'
) comment '模块表' charset = utf8;

create table `interface`
(
    `id`            int unsigned auto_increment primary key,
    `name`          varchar(32)  not null comment '接口名称',
    `method`        varchar(12)  not null comment '请求方式',
    `url`           varchar(500) not null comment '请求地址',
    `created_by`    int unsigned not null comment '创建人',
    `modified_by`   int unsignednull comment '修改人',
    `created_time`  int unsigned default 0 null comment '创建时间',
    `modified_time` int unsigned default 0 null comment '修改时间',
    `state`         tinyint unsigned default 1 null comment '状态 0为禁用、1为启用',
    `project_id`    int          not null comment '所属项目',
    `module_id`     int          not null comment '所属模块'
)comment '接口表' charset = utf8;

create table `interface_user`
(
    `id`           int unsigned auto_increment primary key,
    `interface_id` int     not null comment '接口id',
    `user_id`      int     not null comment '用户id',
    `role`         tinyint not null comment '角色'
)comment '接口人员表' charset = utf8;

create table interface_case
(
    `id`            int unsigned auto_increment primary key,
    `name`          varchar(100) not null comment '用例名称',
    `type`          varchar(12)  not null comment 'body类型',
    `parameters`    text comment 'body参数',
    `headers`       text comment 'body' comment '请求头',
    `query`         text comment 'query参数',
    `asserts`       text comment '断言信息',
    `extract`       text comment '提取参数',
    `remark`        varchar(200) comment '备注',
    `interface_id`  int          not null comment '接口id',
    `env_id`        int          not null comment '环境id',
    `created_by`    int unsigned not null comment '创建人',
    `modified_by`   int unsigned null comment '修改人',
    `created_time`  int unsigned default 0 null comment '创建时间',
    `modified_time` int unsigned default 0 null comment '修改时间',
    `state`         tinyint unsigned default 1 null comment '状态 0为禁用、1为启用',
)comment '接口用例表' charset = utf8;

create table case_log
(
    id           int unsigned auto_increment primary key,
    level        varchar(32) not null comment '日志级别',
    msg          text        not null comment '日志消息',
    created_time int unsigned default 0 null comment '创建时间',
    report_id    int         not null comment '所属报告'
) comment '用例执行日志表';

create table case_result
(
    id                   int unsigned auto_increment primary key,
    case_name            varchar(100)  not null comment '用例名称',
    case_id              int           not null comment '用例id',
    interface_id         int           not null comment '接口id',
    suite_id             int           not null comment '集合id',
    method               varchar(12)   not null comment '请求方式',
    url                  varchar(1024) not null comment '请求url',
    result_type          int           not null comment '执行结果',
    request_headers      text comment '请求头',
    request_query        text comment '请求query参数',
    request_body_type    text comment '请求体类型',
    request_body         text comment '请求体',
    response_status_code int comment '响应状态码',
    response_body        text comment '响应体',
    response_headers     text comment '响应头',
    response_time        int comment '响应时间',
    response_asserts     text comment '断言信息',
    response_extracts    text comment '参数提取信息',
    exception            text comment '异常信息',
    project_id           int           not null comment '项目id',
    created_by           int           not null comment '创建人员',
    created_time         int unsigned default 0 null comment '创建时间'
)comment '用例执行表';