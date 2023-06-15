CREATE TABLE receitas (
    id serial NOT NULL,
    titulo varchar(100) NOT NULL,
    preparo text NULL,
    rendimento text NULL,
    created_at timestamp NULL DEFAULT now(),
    CONSTRAINT receitas_pk PRIMARY KEY (id)
);