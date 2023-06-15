CREATE TABLE ingredientes (
    id serial4 NOT NULL,
    quantidade int4 NOT NULL,
    descricao text NOT NULL,
    receitas_id serial4 NOT NULL,
    CONSTRAINT ingredientes_pk PRIMARY KEY (id),
    CONSTRAINT ingredientes_fk FOREIGN KEY (receitas_id) REFERENCES public.receitas(id)
);