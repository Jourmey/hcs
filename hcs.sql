create table agent
(
    id          int(11) unsigned auto_increment comment '主键id' primary key,
    host_name   varchar(10)  default ''                not null comment 'host_name',
    name        varchar(10)  default ''                not null comment '主机名',
    version     varchar(10)  default ''                not null comment '客户端版本',
    mark        varchar(255) default ''                not null comment '客户端备注',
    heart_time  timestamp    default CURRENT_TIMESTAMP not null comment '最后一次回应心跳的时间',
    create_time timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
);

create table task
(
    id          int(11) unsigned auto_increment comment '主键id' primary key,
    content     longtext charset utf8mb4             not null comment '任务配置内容',
    delete_flag tinyint(2) default 0                 not null comment '逻辑删除标记，默认0，删除1',
    create_time timestamp  default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time timestamp  default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
);

create table agent_task
(
    id          int(11) unsigned auto_increment comment '主键id' primary key,
    agent_id    int          default 0                 not null comment 'agentID',
    task_id     int          default 0                 not null comment 'TaskID',
    mark        varchar(255) default ''                not null comment '任务备注',
    status      tinyint(2)   default 0                 not null comment '任务状态 0:任务创建 1:任务运行中 2:任务执行失败 3:任务执行成功',
    reason      varchar(255) default ''                not null comment '任务状态原因',
    delete_flag tinyint(2)   default 0                 not null comment '逻辑删除标记，默认0，删除1',
    create_time timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
);

create index agent_id_index
    on agent_task (agent_id);

create index task_id_index
    on agent_task (task_id);

create index status_index
    on agent_task (status);