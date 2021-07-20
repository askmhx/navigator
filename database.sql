# drop database  navigator;

create database if not exists navigator;

use navigator;

create table navigator.t_app_inf
(
    app_id     varchar(16)  not null primary key comment '应用id',
    app_key    varchar(32)  not null default '' comment '应用key',
    app_name   varchar(255) not null default '' comment '应用名称',
    department varchar(20)  not null default '' comment '部门',
    owner      varchar(16)  not null default '' comment '应用所有人',
    status     int          not null default 0 comment '状态(1启用/0未启用)',
    memo       varchar(255) comment '备注',
    created_at timestamp    not null default current_timestamp() comment '创建时间',
    created_by varchar(32)  not null comment '创建人',
    updated_at timestamp comment '更新时间',
    updated_by varchar(32)  comment '更新人'
) charset = utf8mb4 comment '应用信息表';

create unique index i_app_id_index on navigator.t_app_inf (app_id);


create table navigator.t_app_profile
(
    app_id     varchar(16)  not null comment '应用id',
    cluster    varchar(20)  not null default '' comment '集群名称',
    profile    varchar(19)  not null default '' comment '环境',
    notify_url varchar(100) comment '通知url',
    cid        varchar(20)  not null comment '配置标识',
    status     int          not null default 0 comment '状态(1启用/0未启用)',
    memo       varchar(255) comment '备注',
    created_at timestamp    not null default current_timestamp() comment '创建时间',
    created_by varchar(32)  not null comment '创建人',
    updated_at timestamp comment '更新时间',
    updated_by varchar(32)  comment '更新人'
) charset = utf8mb4 comment '应用环境表';

create index i_app_profile_id_index on navigator.t_app_profile (app_id);
create unique index i_app_profile_query_index on navigator.t_app_profile (app_id, cluster, profile);

create table navigator.t_app_config
(
    cid        varchar(20) not null comment '配置标识',
    config     text        not null comment '配置（文件）',
    status     int         not null default 0 comment '状态(1正式/0未启用)',
    memo       varchar(255) comment '备注',
    created_at timestamp   not null default current_timestamp() comment '创建时间',
    created_by varchar(32) not null comment '创建人',
    updated_at timestamp comment '更新时间',
    updated_by varchar(32) comment '更新人'
) charset = utf8mb4 comment '应用配置表';

create index i_app_config_id_index on navigator.t_app_config (cid);
create unique index i_app_config_query_index on navigator.t_app_config (cid, status, created_at);
