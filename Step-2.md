# Step-2: 寝る前に一行妄想を登録するツールを作ってみよう！

## データベースの設定

```fantasies.sql
create database fantasy;
use fantasy;

create table `fantasies` (
    `id` int(11) not null auto_increment
    `fantasy` text not null,
    `created_at` datetime not null,
    primary key (`id`)
) engine=InnoDB default charset=utf8;

```

## 妄想を追加してみよう！


## 追加した妄想を見てみよう！


## 妄想を更新してみよう！


## 妄想を削除してみよう！