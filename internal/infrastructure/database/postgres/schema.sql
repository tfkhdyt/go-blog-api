SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: -
--

-- *not* creating schema, since initdb creates it


--
-- Name: role; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.role AS ENUM (
    'user',
    'admin'
);


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: change_email_token; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.change_email_token (
    token text NOT NULL,
    new_email character varying(255) NOT NULL,
    expires_at timestamp without time zone NOT NULL,
    user_id integer
);


--
-- Name: refresh_token; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.refresh_token (
    token text NOT NULL,
    user_id integer
);


--
-- Name: reset_password_token; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.reset_password_token (
    token text NOT NULL,
    expires_at timestamp without time zone NOT NULL,
    user_id integer
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying(128) NOT NULL
);


--
-- Name: user; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public."user" (
    id integer NOT NULL,
    full_name character varying(50) NOT NULL,
    username character varying(16) NOT NULL,
    email character varying(255) NOT NULL,
    password text NOT NULL,
    role public.role DEFAULT 'user'::public.role,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_id_seq OWNED BY public."user".id;


--
-- Name: user id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."user" ALTER COLUMN id SET DEFAULT nextval('public.user_id_seq'::regclass);


--
-- Name: change_email_token change_email_token_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.change_email_token
    ADD CONSTRAINT change_email_token_pkey PRIMARY KEY (token);


--
-- Name: refresh_token refresh_token_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.refresh_token
    ADD CONSTRAINT refresh_token_pkey PRIMARY KEY (token);


--
-- Name: reset_password_token reset_password_token_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.reset_password_token
    ADD CONSTRAINT reset_password_token_pkey PRIMARY KEY (token);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: user user_email_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_email_key UNIQUE (email);


--
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- Name: user user_username_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_username_key UNIQUE (username);


--
-- Name: change_email_token fk_change_email_token_user; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.change_email_token
    ADD CONSTRAINT fk_change_email_token_user FOREIGN KEY (user_id) REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: refresh_token fk_refresh_token_user; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.refresh_token
    ADD CONSTRAINT fk_refresh_token_user FOREIGN KEY (user_id) REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: reset_password_token fk_reset_password_token_user; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.reset_password_token
    ADD CONSTRAINT fk_reset_password_token_user FOREIGN KEY (user_id) REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20230819005847'),
    ('20230819010128'),
    ('20230819010159'),
    ('20230819010234');
