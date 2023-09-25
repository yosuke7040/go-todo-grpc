-- CreateTable
create table if not exists todos (
  id integer unsigned auto_increment primary key,
  title varchar(100) not null,
  status int not null default 0,
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp
);

-- Inster Data
insert into
  todos (title, status)
values
  ('test1', 0);

insert into
  todos (title, status)
values
  ('test2', 0);

insert into
  todos (title, status)
values
  ('test3', 1);