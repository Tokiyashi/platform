--
-- PostgreSQL database dump
--

-- Dumped from database version 16.8 (Postgres.app)
-- Dumped by pg_dump version 17.4

-- Started on 2025-04-05 22:58:53 MSK

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- TOC entry 3695 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 219 (class 1259 OID 16483)
-- Name: courses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.courses (
    id bigint NOT NULL,
    title character varying(100) NOT NULL,
    description character varying(200),
    creator_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.courses OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 16482)
-- Name: courses_creator_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.courses_creator_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.courses_creator_id_seq OWNER TO postgres;

--
-- TOC entry 3696 (class 0 OID 0)
-- Dependencies: 218
-- Name: courses_creator_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.courses_creator_id_seq OWNED BY public.courses.creator_id;


--
-- TOC entry 217 (class 1259 OID 16481)
-- Name: courses_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.courses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.courses_id_seq OWNER TO postgres;

--
-- TOC entry 3697 (class 0 OID 0)
-- Dependencies: 217
-- Name: courses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.courses_id_seq OWNED BY public.courses.id;


--
-- TOC entry 228 (class 1259 OID 16538)
-- Name: files; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.files (
    id bigint NOT NULL,
    path character varying(200) NOT NULL,
    extension character varying(10) NOT NULL,
    type character varying(100) NOT NULL,
    message_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    author_id bigint NOT NULL
);


ALTER TABLE public.files OWNER TO postgres;

--
-- TOC entry 226 (class 1259 OID 16536)
-- Name: file_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.file_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.file_id_seq OWNER TO postgres;

--
-- TOC entry 3698 (class 0 OID 0)
-- Dependencies: 226
-- Name: file_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.file_id_seq OWNED BY public.files.id;


--
-- TOC entry 227 (class 1259 OID 16537)
-- Name: file_message_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.file_message_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.file_message_id_seq OWNER TO postgres;

--
-- TOC entry 3699 (class 0 OID 0)
-- Dependencies: 227
-- Name: file_message_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.file_message_id_seq OWNED BY public.files.message_id;


--
-- TOC entry 229 (class 1259 OID 16551)
-- Name: files_author_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.files_author_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.files_author_id_seq OWNER TO postgres;

--
-- TOC entry 3700 (class 0 OID 0)
-- Dependencies: 229
-- Name: files_author_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.files_author_id_seq OWNED BY public.files.author_id;


--
-- TOC entry 225 (class 1259 OID 16523)
-- Name: messages; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.messages (
    id bigint NOT NULL,
    text_content character varying(400) NOT NULL,
    section_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.messages OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 16521)
-- Name: message_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.message_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.message_id_seq OWNER TO postgres;

--
-- TOC entry 3701 (class 0 OID 0)
-- Dependencies: 223
-- Name: message_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.message_id_seq OWNED BY public.messages.id;


--
-- TOC entry 224 (class 1259 OID 16522)
-- Name: message_section_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.message_section_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.message_section_id_seq OWNER TO postgres;

--
-- TOC entry 3702 (class 0 OID 0)
-- Dependencies: 224
-- Name: message_section_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.message_section_id_seq OWNED BY public.messages.section_id;


--
-- TOC entry 222 (class 1259 OID 16508)
-- Name: sections; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sections (
    id bigint NOT NULL,
    title character varying(100) NOT NULL,
    description character varying(200),
    course_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.sections OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 16507)
-- Name: section_course_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.section_course_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.section_course_id_seq OWNER TO postgres;

--
-- TOC entry 3703 (class 0 OID 0)
-- Dependencies: 221
-- Name: section_course_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.section_course_id_seq OWNED BY public.sections.course_id;


--
-- TOC entry 220 (class 1259 OID 16506)
-- Name: section_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.section_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.section_id_seq OWNER TO postgres;

--
-- TOC entry 3704 (class 0 OID 0)
-- Dependencies: 220
-- Name: section_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.section_id_seq OWNED BY public.sections.id;


--
-- TOC entry 216 (class 1259 OID 16417)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    first_name character varying(50),
    last_name character varying(50),
    email character varying(50),
    role character varying(50),
    created_at timestamp without time zone DEFAULT now(),
    password character varying(50) NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 231 (class 1259 OID 16584)
-- Name: users_courses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users_courses (
    id bigint NOT NULL,
    user_id bigint,
    course_id bigint
);


ALTER TABLE public.users_courses OWNER TO postgres;

--
-- TOC entry 230 (class 1259 OID 16583)
-- Name: users_courses_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_courses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_courses_id_seq OWNER TO postgres;

--
-- TOC entry 3705 (class 0 OID 0)
-- Dependencies: 230
-- Name: users_courses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_courses_id_seq OWNED BY public.users_courses.id;


--
-- TOC entry 215 (class 1259 OID 16416)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 3706 (class 0 OID 0)
-- Dependencies: 215
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 3497 (class 2604 OID 16486)
-- Name: courses id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.courses ALTER COLUMN id SET DEFAULT nextval('public.courses_id_seq'::regclass);


--
-- TOC entry 3498 (class 2604 OID 16487)
-- Name: courses creator_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.courses ALTER COLUMN creator_id SET DEFAULT nextval('public.courses_creator_id_seq'::regclass);


--
-- TOC entry 3506 (class 2604 OID 16541)
-- Name: files id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.files ALTER COLUMN id SET DEFAULT nextval('public.file_id_seq'::regclass);


--
-- TOC entry 3507 (class 2604 OID 16542)
-- Name: files message_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.files ALTER COLUMN message_id SET DEFAULT nextval('public.file_message_id_seq'::regclass);


--
-- TOC entry 3509 (class 2604 OID 16607)
-- Name: files author_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.files ALTER COLUMN author_id SET DEFAULT nextval('public.files_author_id_seq'::regclass);


--
-- TOC entry 3503 (class 2604 OID 16526)
-- Name: messages id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.messages ALTER COLUMN id SET DEFAULT nextval('public.message_id_seq'::regclass);


--
-- TOC entry 3504 (class 2604 OID 16601)
-- Name: messages section_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.messages ALTER COLUMN section_id SET DEFAULT nextval('public.message_section_id_seq'::regclass);


--
-- TOC entry 3500 (class 2604 OID 16511)
-- Name: sections id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sections ALTER COLUMN id SET DEFAULT nextval('public.section_id_seq'::regclass);


--
-- TOC entry 3501 (class 2604 OID 16512)
-- Name: sections course_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sections ALTER COLUMN course_id SET DEFAULT nextval('public.section_course_id_seq'::regclass);


--
-- TOC entry 3495 (class 2604 OID 16420)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3510 (class 2604 OID 16587)
-- Name: users_courses id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_courses ALTER COLUMN id SET DEFAULT nextval('public.users_courses_id_seq'::regclass);


--
-- TOC entry 3677 (class 0 OID 16483)
-- Dependencies: 219
-- Data for Name: courses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.courses (id, title, description, creator_id, created_at) FROM stdin;
\.


--
-- TOC entry 3686 (class 0 OID 16538)
-- Dependencies: 228
-- Data for Name: files; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.files (id, path, extension, type, message_id, created_at, author_id) FROM stdin;
\.


--
-- TOC entry 3683 (class 0 OID 16523)
-- Dependencies: 225
-- Data for Name: messages; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.messages (id, text_content, section_id, created_at) FROM stdin;
\.


--
-- TOC entry 3680 (class 0 OID 16508)
-- Dependencies: 222
-- Data for Name: sections; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sections (id, title, description, course_id, created_at) FROM stdin;
\.


--
-- TOC entry 3674 (class 0 OID 16417)
-- Dependencies: 216
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, first_name, last_name, email, role, created_at, password) FROM stdin;
3	Klad	JormZ	vladislav.incorporated@sosal.com	student	2025-03-29 23:29:39.323459	5434
4	test		erere	\N	2025-03-30 17:41:53.269716	asds
5	2	3	1	\N	2025-03-30 19:30:16.842874	4
6	2	3	1	\N	2025-03-30 19:30:34.621603	4
7	2	3	1	\N	2025-03-30 19:56:33.730717	4
8	2	3	1	\N	2025-03-30 19:56:48.970603	
9	2	3	1	\N	2025-03-30 19:56:54.545056	
10		3		\N	2025-03-30 19:57:10.545231	
11				\N	2025-03-30 19:58:04.047304	
12				\N	2025-03-30 19:58:19.338194	
\.


--
-- TOC entry 3689 (class 0 OID 16584)
-- Dependencies: 231
-- Data for Name: users_courses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users_courses (id, user_id, course_id) FROM stdin;
\.


--
-- TOC entry 3707 (class 0 OID 0)
-- Dependencies: 218
-- Name: courses_creator_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.courses_creator_id_seq', 1, false);


--
-- TOC entry 3708 (class 0 OID 0)
-- Dependencies: 217
-- Name: courses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.courses_id_seq', 1, false);


--
-- TOC entry 3709 (class 0 OID 0)
-- Dependencies: 226
-- Name: file_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.file_id_seq', 1, false);


--
-- TOC entry 3710 (class 0 OID 0)
-- Dependencies: 227
-- Name: file_message_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.file_message_id_seq', 1, false);


--
-- TOC entry 3711 (class 0 OID 0)
-- Dependencies: 229
-- Name: files_author_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.files_author_id_seq', 1, false);


--
-- TOC entry 3712 (class 0 OID 0)
-- Dependencies: 223
-- Name: message_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.message_id_seq', 1, false);


--
-- TOC entry 3713 (class 0 OID 0)
-- Dependencies: 224
-- Name: message_section_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.message_section_id_seq', 1, false);


--
-- TOC entry 3714 (class 0 OID 0)
-- Dependencies: 221
-- Name: section_course_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.section_course_id_seq', 1, false);


--
-- TOC entry 3715 (class 0 OID 0)
-- Dependencies: 220
-- Name: section_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.section_id_seq', 1, false);


--
-- TOC entry 3716 (class 0 OID 0)
-- Dependencies: 230
-- Name: users_courses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_courses_id_seq', 1, false);


--
-- TOC entry 3717 (class 0 OID 0)
-- Dependencies: 215
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 12, true);


--
-- TOC entry 3514 (class 2606 OID 16490)
-- Name: courses courses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.courses
    ADD CONSTRAINT courses_pkey PRIMARY KEY (id);


--
-- TOC entry 3520 (class 2606 OID 16545)
-- Name: files file_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.files
    ADD CONSTRAINT file_pkey PRIMARY KEY (id);


--
-- TOC entry 3518 (class 2606 OID 16530)
-- Name: messages message_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT message_pkey PRIMARY KEY (id);


--
-- TOC entry 3516 (class 2606 OID 16515)
-- Name: sections section_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sections
    ADD CONSTRAINT section_pkey PRIMARY KEY (id);


--
-- TOC entry 3522 (class 2606 OID 16589)
-- Name: users_courses users_courses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_courses
    ADD CONSTRAINT users_courses_pkey PRIMARY KEY (id);


--
-- TOC entry 3512 (class 2606 OID 16480)
-- Name: users users_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_unique UNIQUE (id);


--
-- TOC entry 3523 (class 2606 OID 16491)
-- Name: courses courses_creator_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.courses
    ADD CONSTRAINT courses_creator_id_fkey FOREIGN KEY (creator_id) REFERENCES public.users(id);


--
-- TOC entry 3526 (class 2606 OID 16546)
-- Name: files file_message_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.files
    ADD CONSTRAINT file_message_id_fkey FOREIGN KEY (message_id) REFERENCES public.messages(id);


--
-- TOC entry 3527 (class 2606 OID 16608)
-- Name: files files_author_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.files
    ADD CONSTRAINT files_author_id_fkey FOREIGN KEY (author_id) REFERENCES public.users(id);


--
-- TOC entry 3525 (class 2606 OID 16602)
-- Name: messages message_section_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT message_section_id_fkey FOREIGN KEY (section_id) REFERENCES public.sections(id);


--
-- TOC entry 3524 (class 2606 OID 16516)
-- Name: sections section_course_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sections
    ADD CONSTRAINT section_course_id_fkey FOREIGN KEY (course_id) REFERENCES public.courses(id);


--
-- TOC entry 3528 (class 2606 OID 16595)
-- Name: users_courses users_courses_course_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_courses
    ADD CONSTRAINT users_courses_course_id_fkey FOREIGN KEY (course_id) REFERENCES public.courses(id);


--
-- TOC entry 3529 (class 2606 OID 16590)
-- Name: users_courses users_courses_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_courses
    ADD CONSTRAINT users_courses_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


-- Completed on 2025-04-05 22:58:53 MSK

--
-- PostgreSQL database dump complete
--

