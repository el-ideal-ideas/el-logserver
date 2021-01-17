package db


var elLogTable = `
create table if not exists el_log (
    id bigint auto_increment primary key,
    ip varchar(45) not null,
    user_agent varchar(512) not null,
    app_name varchar(16) not null,
    log_type int not null,
    message text not null,
    info text default null,
    timestamp bigint
);
`

var insertLog = `
insert into el_log(ip, user_agent, app_name, log_type, message, info, timestamp)
values(?,?,?,?,?,?,?);
`

var getLog = `
select id, ip, user_agent, log_type, message, info, timestamp from el_log where app_name = ? order by id desc limit ? ?;
`

var cntLog = `
select count(1) from el_log where app_name = ?;
`
