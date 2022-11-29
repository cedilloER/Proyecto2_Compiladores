/*
    comentario de multiples lineas en sql
    /* Esto no provoca error
    -- Esto no provoca error
*/

--Esto es un comentario de una linea 

create database USERS;

use USERS;

create table users (
    int id_users primary key, --Comentario 1
    varchar(16) _user_name,
    varchar(16) _password,
    varchar(25) _role 
);
/* hola ud */

insert into users values (1010, "elias", "elias22", "Administrador"); /* hola 
mundo 
curel */ insert into users /* esto es lo maximo */ values (1020, "orlin", "orlin34", "Cajero");
insert into users values (1030, "karen", "karen38", "Cliente"); --Comentario 2
insert into users values (1040, "litzy", "litzy88", "Cajero");

select * from users;

delete users where _user_name = "litzy";
/*
select _user_name, _role from users where _role = "Cajero"; * /