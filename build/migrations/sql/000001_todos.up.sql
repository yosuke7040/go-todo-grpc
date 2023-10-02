-- CreateTable
create table if not exists users (
  id integer unsigned auto_increment primary key,
  auth0_id varchar(100),
  name varchar(100),
  email varchar(100),
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp
);

create table if not exists todos (
  id integer unsigned auto_increment primary key,
  title varchar(100) not null,
  status int not null default 0,
  user_id integer unsigned not null,
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp,
  foreign key (user_id) references users(id)
);


-- Inster Data
insert into 
  users (auth0_id, name, email)
values
  ("hoge", "abe", "email@example.com"),
  ("fuga", "test-user", "hoge_mail@example.com");

insert into
  todos (title, status, user_id)
values
  ('test1', 0, 1),
  ('test2', 0, 1),
  ('test3', 1, 2);
