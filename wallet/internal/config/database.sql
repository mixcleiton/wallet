-- Table: public.event_status

-- DROP TABLE IF EXISTS public.event_status;

CREATE TABLE IF NOT EXISTS public.event_status
(
    id integer NOT NULL DEFAULT nextval('event_status_id_seq'::regclass),
    name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    description character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT event_status_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.event_status
    OWNER to root;

-- Table: public.event_type

-- DROP TABLE IF EXISTS public.event_type;

CREATE TABLE IF NOT EXISTS public.event_type
(
    id integer NOT NULL DEFAULT nextval('event_type_id_seq'::regclass),
    name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    description character varying(255) COLLATE pg_catalog."default" NOT NULL,
    example character varying(255) COLLATE pg_catalog."default",
    CONSTRAINT event_type_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.event_type
    OWNER to root;


-- Table: public.extract

-- DROP TABLE IF EXISTS public."extract";

CREATE TABLE IF NOT EXISTS public."extract"
(
    id integer NOT NULL DEFAULT nextval('extract_id_seq'::regclass),
    created_at timestamp without time zone NOT NULL,
    event_id bigint NOT NULL,
    CONSTRAINT extract_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."extract"
    OWNER to root;


-- Table: public.wallet

-- DROP TABLE IF EXISTS public.wallet;

CREATE TABLE IF NOT EXISTS public.wallet
(
    id integer NOT NULL DEFAULT nextval('wallet_id_seq'::regclass),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone,
    "user" character varying(100) COLLATE pg_catalog."default" NOT NULL,
    document_number character varying(20) COLLATE pg_catalog."default" NOT NULL,
    id_uuid character varying(100) COLLATE pg_catalog."default" NOT NULL,
    saldo numeric(20,2) NOT NULL,
    CONSTRAINT wallet_pkey PRIMARY KEY (id),
    CONSTRAINT u_document_number UNIQUE (document_number),
    CONSTRAINT u_id_uuid UNIQUE (id_uuid)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.wallet
    OWNER to root;

-- Table: public.event

-- DROP TABLE IF EXISTS public.event;

CREATE TABLE IF NOT EXISTS public.event
(
    id integer NOT NULL DEFAULT nextval('event_id_seq'::regclass),
    wallet_id bigint NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone,
    status_id bigint NOT NULL,
    type_id bigint NOT NULL,
    description character varying(255) COLLATE pg_catalog."default" NOT NULL,
    id_uuid character varying(100) COLLATE pg_catalog."default" NOT NULL,
    value numeric(20,2) NOT NULL,
    event_id bigint,
    CONSTRAINT event_pkey PRIMARY KEY (id),
    CONSTRAINT fk_event FOREIGN KEY (event_id)
        REFERENCES public.event (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT fk_event_status FOREIGN KEY (status_id)
        REFERENCES public.event_status (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT fk_event_type FOREIGN KEY (type_id)
        REFERENCES public.event_type (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT wallet_id FOREIGN KEY (wallet_id)
        REFERENCES public.wallet (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.event
    OWNER to root;

INSERT INTO public.event_type(
	id, name, description, example)
	VALUES (5, 'ADIÇÃO', 'Ação de incrementar o saldo da carteira', 'Depósitos bancários, transferências de outras carteiras, recebimento de pagamentos');

INSERT INTO public.event_type(
	id, name, description, example)
	VALUES (6, 'RETIRADA', 'Ação de diminuir o saldo da carteira', ' Saques em caixas eletrônicos, transferências para outras contas, pagamentos de contas');

INSERT INTO public.event_type(
	id, name, description, example)
	VALUES (7, 'COMPRAS', 'Ação de utilizar o saldo da carteira para adquirir produtos ou serviços', 'Compras online, pagamentos em estabelecimentos físicos, assinaturas de serviços');

INSERT INTO public.event_type(
	id, name, description, example)
	VALUES (8, 'CANCELAMENTO', 'Ação de anular uma transação pendente ou em andamento', 'Cancelamento de uma compra online antes da entrega, cancelamento de uma assinatura');

INSERT INTO public.event_type(
	id, name, description, example)
	VALUES (9, 'ESTORNO', 'Ação de devolver um valor já creditado em uma conta, geralmente devido a uma compra cancelada, erro na transação ou disputa', 'Estorno de uma compra por defeito no produto, estorno de um pagamento duplicado');


INSERT INTO public.event_status(
	id, name, description)
	VALUES (1, 'AGUARDANDO_PROCESSAMENTO', 'evento que acabou de ser criado na base');

INSERT INTO public.event_status(
	id, name, description)
	VALUES (2, 'PROCESSADO', 'evento que foi processado e está aguardando alguma tratativa');

INSERT INTO public.event_status(
	id, name, description)
	VALUES (3, 'CONCLUIDO', 'evento que finalizou o fluxo de processamento');

INSERT INTO public.event_status(
	id, name, description)
	VALUES (4, 'CANCELADO', 'Evento cancelado por outro evento, exemplo um estorno');