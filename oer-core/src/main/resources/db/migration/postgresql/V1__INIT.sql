                 --
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.12
-- Dumped by pg_dump version 9.6.12

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: oer_server_dev; Type: DATABASE; Schema: -; Owner: postgres
--

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: channel; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE IF NOT EXISTS public.channel (
    id bigint NOT NULL,
    adapter_family integer,
    channel_key integer,
    home_page character varying(255),
    name character varying(255),
    technical_id character varying(255)
);


--
-- Name: channel_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.channel_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: channel_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.channel_id_seq OWNED BY public.channel.id;


--
-- Name: image_link; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE IF NOT EXISTS public.image_link (
    id bigint NOT NULL,
    created_at timestamp without time zone NOT NULL,
    url character varying(2500)
);


--
-- Name: image_link_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.image_link_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: image_link_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.image_link_id_seq OWNED BY public.image_link.id;


--
-- Name: program_entry; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE IF NOT EXISTS public.program_entry (
    id bigint NOT NULL,
    adapter_family integer NOT NULL,
    created_at timestamp without time zone,
    description text,
    duration_in_minutes integer,
    end_date_time timestamp without time zone,
    home_page character varying(1000),
    start_date_time timestamp without time zone,
    technical_id character varying(500),
    title character varying(1000),
    updated_at timestamp without time zone,
    url character varying(1000),
    channel_id bigint NOT NULL
);


--
-- Name: program_entry_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.program_entry_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: program_entry_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.program_entry_id_seq OWNED BY public.program_entry.id;


--
-- Name: program_entry_image_links; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE IF NOT EXISTS public.program_entry_image_links (
    program_entry_id bigint NOT NULL,
    image_links_id bigint NOT NULL
);


--
-- Name: program_entry_tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE IF NOT EXISTS public.program_entry_tags (
    program_entry_id bigint NOT NULL,
    tags_id bigint NOT NULL
);


--
-- Name: tag; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE IF NOT EXISTS public.tag (
    id bigint NOT NULL,
    created_at timestamp without time zone NOT NULL,
    tag_name character varying(500)
);


--
-- Name: tag_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tag_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

--
-- Name: tag_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tag_id_seq OWNED BY public.tag.id;


--
-- Name: tv_show; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE IF NOT EXISTS public.tv_show (
    id bigint NOT NULL,
    adapter_family integer NOT NULL,
    additional_id character varying(1000),
    created_at timestamp without time zone,
    home_page character varying(1500),
    image_url character varying(1500),
    technical_id character varying(32) NOT NULL,
    title character varying(1000) NOT NULL,
    updated_at timestamp without time zone,
    url character varying(1500)
);


--
-- Name: tv_show_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tv_show_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: tv_show_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tv_show_id_seq OWNED BY public.tv_show.id;


--
-- Name: tv_show_related_program_entries; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE IF NOT EXISTS public.tv_show_related_program_entries (
    tv_show_id bigint NOT NULL,
    related_program_entry_id bigint NOT NULL
);


--
-- Name: tv_show_tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE IF NOT EXISTS public.tv_show_tags (
    tv_show_id bigint NOT NULL,
    tag_id bigint NOT NULL
);


--
-- Name: channel id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.channel ALTER COLUMN id SET DEFAULT nextval('public.channel_id_seq'::regclass);


--
-- Name: image_link id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.image_link ALTER COLUMN id SET DEFAULT nextval('public.image_link_id_seq'::regclass);


--
-- Name: program_entry id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.program_entry ALTER COLUMN id SET DEFAULT nextval('public.program_entry_id_seq'::regclass);


--
-- Name: tag id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tag ALTER COLUMN id SET DEFAULT nextval('public.tag_id_seq'::regclass);


--
-- Name: tv_show id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tv_show ALTER COLUMN id SET DEFAULT nextval('public.tv_show_id_seq'::regclass);


--
-- Name: channel channel_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.channel
    ADD CONSTRAINT channel_pkey PRIMARY KEY (id);


--
-- Name: image_link image_link_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.image_link
    ADD CONSTRAINT image_link_pkey PRIMARY KEY (id);


--
-- Name: program_entry program_entry_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.program_entry
    ADD CONSTRAINT program_entry_pkey PRIMARY KEY (id);


--
-- Name: tag tag_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tag
    ADD CONSTRAINT tag_pkey PRIMARY KEY (id);


--
-- Name: tv_show tv_show_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tv_show
    ADD CONSTRAINT tv_show_pkey PRIMARY KEY (id);


--
-- Name: program_entry UK_technical_id_adapter_family; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.program_entry
    ADD CONSTRAINT UK_technical_id_adapter_family UNIQUE (technical_id, adapter_family);


--
-- Name: tag UK_tag_name_index; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tag
    ADD CONSTRAINT UK_tag_name_index UNIQUE (tag_name);


--
-- Name: program_entry UK_technical_id_index; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.program_entry
    ADD CONSTRAINT UK_technical_id_index UNIQUE (technical_id);


--
-- Name: channel UK_channel_key_index; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.channel
    ADD CONSTRAINT UK_channel_key_index UNIQUE (channel_key);


--
-- Name: channel UK_channel_technical_id_index; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.channel
    ADD CONSTRAINT UK_channel_technical_id_index UNIQUE (technical_id);


--
-- Name: tv_show ukn7n8qfppeuky4u7tp0ahmau9h; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tv_show
    ADD CONSTRAINT UK_adapter_tid_index UNIQUE (adapter_family, technical_id);


--
-- Name: program_entry_tags fk29tmop492x2w7lrux944oc8vr; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.program_entry_tags
    ADD CONSTRAINT FK_tag_id_relation FOREIGN KEY (tags_id) REFERENCES public.tag(id);


--
-- Name: program_entry FK_channel_relation; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.program_entry
    ADD CONSTRAINT FK_channel_relation FOREIGN KEY (channel_id) REFERENCES public.channel(id);


--
-- Name: program_entry_image_links FK_program_entry_relation; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.program_entry_image_links
    ADD CONSTRAINT FK_program_entry_relation FOREIGN KEY (program_entry_id) REFERENCES public.program_entry(id);


--
-- Name: tv_show_tags FK_tv_show_tag_relation; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tv_show_tags
    ADD CONSTRAINT FK_tv_show_tag_relation FOREIGN KEY (tv_show_id) REFERENCES public.tv_show(id);


--
-- Name: tv_show_related_program_entries FK_tv_program_entry_relation; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tv_show_related_program_entries
    ADD CONSTRAINT FK_tv_program_entry_relation FOREIGN KEY (related_program_entry_id) REFERENCES public.program_entry(id);


--
-- Name: program_entry_tags FK_program_entry_tag_relation; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.program_entry_tags
    ADD CONSTRAINT FK_program_entry_tag_relation FOREIGN KEY (program_entry_id) REFERENCES public.program_entry(id);


--
-- Name: tv_show_tags FK_tag_show_relation; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tv_show_tags
    ADD CONSTRAINT FK_tag_show_relation FOREIGN KEY (tag_id) REFERENCES public.tag(id);


--
-- Name: tv_show_related_program_entries FK_tv_show_relation; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tv_show_related_program_entries
    ADD CONSTRAINT FK_tv_show_relation FOREIGN KEY (tv_show_id) REFERENCES public.tv_show(id);


--
-- Name: program_entry_image_links FK_tv_program_entry_relation; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.program_entry_image_links
    ADD CONSTRAINT FK_tv_program_entry_relation FOREIGN KEY (image_links_id) REFERENCES public.image_link(id);


--
-- PostgreSQL database dump complete
--