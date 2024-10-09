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
    id_uuid character varying(100) COLLATE pg_catalog."default" NOT NULL,
    wallet_id bigint NOT NULL,
    type_id bigint NOT NULL,
    value numeric(20,2) NOT NULL,
    status_id bigint NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone
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