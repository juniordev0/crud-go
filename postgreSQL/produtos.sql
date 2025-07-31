-- public.produtos definition

-- Drop table

-- DROP TABLE public.produtos;

CREATE TABLE public.produtos (
	id serial4 NOT NULL,
	nome varchar(100) NOT NULL,
	descricao text NULL,
	preco numeric(10, 2) NOT NULL,
	quantidade int4 NOT NULL,
	criado_em timestamp NULL,
	atualizado_em timestamp NULL,
	CONSTRAINT produtos_pkey PRIMARY KEY (id)
);