-- SEQUENCE: public.wallets_wallet_index_seq

-- DROP SEQUENCE IF EXISTS public.wallets_wallet_index_seq;

CREATE SEQUENCE IF NOT EXISTS public.wallets_wallet_index_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE public.wallets_wallet_index_seq
    OWNER TO postgres;

-- Table: public.wallets

-- DROP TABLE IF EXISTS public.wallets;

CREATE TABLE IF NOT EXISTS public.wallets
(
    wallet_index bigint NOT NULL DEFAULT nextval('wallets_wallet_index_seq'::regclass),
    wallet_address character varying COLLATE pg_catalog."default",
    created_at bigint,
    CONSTRAINT wallets_pkey PRIMARY KEY (wallet_index)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.wallets
    OWNER to postgres;