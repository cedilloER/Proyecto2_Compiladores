/*
    comentario de multiples lineas en sql
    /* Esto no provoca error
    -- Esto no provoca error
*/

--Esto es un comentario de una linea 

create database USERS;

use USERS;

create table users (
    int id_users primary key,
    varchar(16) _user_name,
    varchar(16) _password,
    varchar(25) _role
);

insert into users values (1010, "elias", "elias22", "Administrador");
insert into users values (1020, "orlin", "orlin34", "Cajero"); /* esto tambien */
insert into users values (1030, "karen", "karen38", "Cliente");
insert into users values (1040, "litzy", "litzy88", "Cajero"); -- Esto es una prueba


select * from users;

delete users where _user_name = "litzy";

select _user_name, _role from _role = "Cajero";