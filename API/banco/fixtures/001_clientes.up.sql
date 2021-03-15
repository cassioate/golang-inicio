create table cliente (
    uuid varchar(45),
    nome varchar(450),
    endereco varchar(450),
    cadastrado_em VARCHAR(100) DEFAULT to_char(current_timestamp, 'DD/MM/YYYY HH24:MI:SS'),
    atualizado_em VARCHAR(100) DEFAULT to_char(current_timestamp, 'DD/MM/YYYY HH24:MI:SS'),
    constraint id_pk primary key (uuid)
);