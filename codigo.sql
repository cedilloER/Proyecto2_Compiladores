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
insert into users values (1020, "orlin", "orlin34", "Cajero");
insert into users values (1030, "karen", "karen38", "Cliente");
insert into users values (1040, "litzy", "litzy88", "Cajero");
select * from users;
delete users where _user_name = "litzy";
select _user_name, _role from _role = "Cajero";


Tokes 

id, token, lexema

1, create, CREATE
2, use, USE
3, database, DATABASE
4, table, table
5, int, INT 
6, varchar, VARCHAR
7, float, FLOAT
8, date, DATE
9, char, CHAR
10, boolean, BOOLEAN
11, insert, INSERT
12, into, INTO
13, values, VALUES
14, select, SELECT 
15, from, FROM 
16, delete, DELETE 
17, where, WHERE 




