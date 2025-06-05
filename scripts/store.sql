--
-- PostgreSQL database dump
--

-- Dumped from database version 12.2
-- Dumped by pg_dump version 12.2

-- Started on 2025-06-05 11:27:08

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
-- TOC entry 6 (class 2615 OID 1063275)
-- Name: store; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA store;


ALTER SCHEMA store OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 213 (class 1259 OID 1063367)
-- Name: admins; Type: TABLE; Schema: store; Owner: postgres
--

CREATE TABLE store.admins (
    id bigint DEFAULT nextval(('store.admin_seq'::text)::regclass) NOT NULL,
    login character varying(64),
    salt character varying(64),
    pass character varying(64)
);


ALTER TABLE store.admins OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 1063365)
-- Name: admin_seq; Type: SEQUENCE; Schema: store; Owner: postgres
--

CREATE SEQUENCE store.admin_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE store.admin_seq OWNER TO postgres;

--
-- TOC entry 2923 (class 0 OID 0)
-- Dependencies: 212
-- Name: admin_seq; Type: SEQUENCE OWNED BY; Schema: store; Owner: postgres
--

ALTER SEQUENCE store.admin_seq OWNED BY store.admins.id;


--
-- TOC entry 215 (class 1259 OID 1063398)
-- Name: authors; Type: TABLE; Schema: store; Owner: postgres
--

CREATE TABLE store.authors (
    id integer DEFAULT nextval(('store.author_seq'::text)::regclass) NOT NULL,
    first_name character varying(100) NOT NULL,
    last_name character varying(100) NOT NULL,
    middle_name character varying(100),
    birth_date date,
    country character varying(100)
);


ALTER TABLE store.authors OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 1063396)
-- Name: author_seq; Type: SEQUENCE; Schema: store; Owner: postgres
--

CREATE SEQUENCE store.author_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE store.author_seq OWNER TO postgres;

--
-- TOC entry 2924 (class 0 OID 0)
-- Dependencies: 214
-- Name: author_seq; Type: SEQUENCE OWNED BY; Schema: store; Owner: postgres
--

ALTER SEQUENCE store.author_seq OWNED BY store.authors.id;


--
-- TOC entry 206 (class 1259 OID 1063291)
-- Name: books; Type: TABLE; Schema: store; Owner: postgres
--

CREATE TABLE store.books (
    id bigint DEFAULT nextval(('store.book_seq'::text)::regclass) NOT NULL,
    isbn character varying(20) NOT NULL,
    title character varying(128),
    descr text,
    price numeric(19,2),
    publisher_id bigint,
    author_id bigint,
    publication_year integer NOT NULL,
    genre character varying(64) NOT NULL
);
ALTER TABLE ONLY store.books ALTER COLUMN id SET STATISTICS 0;


ALTER TABLE store.books OWNER TO postgres;

--
-- TOC entry 205 (class 1259 OID 1063289)
-- Name: book_seq; Type: SEQUENCE; Schema: store; Owner: postgres
--

CREATE SEQUENCE store.book_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE store.book_seq OWNER TO postgres;

--
-- TOC entry 2925 (class 0 OID 0)
-- Dependencies: 205
-- Name: book_seq; Type: SEQUENCE OWNED BY; Schema: store; Owner: postgres
--

ALTER SEQUENCE store.book_seq OWNED BY store.books.id;


--
-- TOC entry 217 (class 1259 OID 1063424)
-- Name: clients; Type: TABLE; Schema: store; Owner: postgres
--

CREATE TABLE store.clients (
    id bigint DEFAULT nextval(('store.client_seq'::text)::regclass) NOT NULL,
    first_name character varying(100) NOT NULL,
    last_name character varying(100) NOT NULL,
    middle_name character varying(100),
    login character varying(64),
    salt character varying(64),
    password character varying(64)
);


ALTER TABLE store.clients OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 1063422)
-- Name: client_seq; Type: SEQUENCE; Schema: store; Owner: postgres
--

CREATE SEQUENCE store.client_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE store.client_seq OWNER TO postgres;

--
-- TOC entry 2926 (class 0 OID 0)
-- Dependencies: 216
-- Name: client_seq; Type: SEQUENCE OWNED BY; Schema: store; Owner: postgres
--

ALTER SEQUENCE store.client_seq OWNED BY store.clients.id;


--
-- TOC entry 209 (class 1259 OID 1063341)
-- Name: order_items; Type: TABLE; Schema: store; Owner: postgres
--

CREATE TABLE store.order_items (
    order_id bigint NOT NULL,
    book_id bigint,
    item_price numeric(19,2),
    qty integer,
    CONSTRAINT order_items_chk CHECK ((qty > 0))
);
ALTER TABLE ONLY store.order_items ALTER COLUMN order_id SET STATISTICS 0;


ALTER TABLE store.order_items OWNER TO postgres;

--
-- TOC entry 208 (class 1259 OID 1063322)
-- Name: orders; Type: TABLE; Schema: store; Owner: postgres
--

CREATE TABLE store.orders (
    id bigint DEFAULT nextval(('store.order_seq'::text)::regclass) NOT NULL,
    client_id bigint,
    amount numeric(19,2),
    dt date DEFAULT now() NOT NULL,
    ship_name character varying(64),
    ship_address character varying(128),
    ship_city character varying(64),
    ship_zip_code character varying(32),
    ship_country character varying(64),
    status character(1) DEFAULT 'N'::bpchar NOT NULL,
    CONSTRAINT orders_chk CHECK ((status = ANY (ARRAY['N'::bpchar, 'P'::bpchar, 'S'::bpchar])))
);
ALTER TABLE ONLY store.orders ALTER COLUMN id SET STATISTICS 0;


ALTER TABLE store.orders OWNER TO postgres;

--
-- TOC entry 2927 (class 0 OID 0)
-- Dependencies: 208
-- Name: COLUMN orders.status; Type: COMMENT; Schema: store; Owner: postgres
--

COMMENT ON COLUMN store.orders.status IS 'N - Оформлен
P - Оплачен
S - Отправлен';


--
-- TOC entry 207 (class 1259 OID 1063320)
-- Name: order_seq; Type: SEQUENCE; Schema: store; Owner: postgres
--

CREATE SEQUENCE store.order_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE store.order_seq OWNER TO postgres;

--
-- TOC entry 2928 (class 0 OID 0)
-- Dependencies: 207
-- Name: order_seq; Type: SEQUENCE OWNED BY; Schema: store; Owner: postgres
--

ALTER SEQUENCE store.order_seq OWNED BY store.orders.id;


--
-- TOC entry 211 (class 1259 OID 1063350)
-- Name: publishers; Type: TABLE; Schema: store; Owner: postgres
--

CREATE TABLE store.publishers (
    id bigint DEFAULT nextval(('store.publisher_seq'::text)::regclass) NOT NULL,
    name character varying(128) NOT NULL,
    country character varying(32) NOT NULL,
    phone character varying(16) NOT NULL
);


ALTER TABLE store.publishers OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 1063348)
-- Name: publisher_seq; Type: SEQUENCE; Schema: store; Owner: postgres
--

CREATE SEQUENCE store.publisher_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE store.publisher_seq OWNER TO postgres;

--
-- TOC entry 2929 (class 0 OID 0)
-- Dependencies: 210
-- Name: publisher_seq; Type: SEQUENCE OWNED BY; Schema: store; Owner: postgres
--

ALTER SEQUENCE store.publisher_seq OWNED BY store.publishers.id;


--
-- TOC entry 220 (class 1259 OID 1063470)
-- Name: warehouse_books; Type: TABLE; Schema: store; Owner: postgres
--

CREATE TABLE store.warehouse_books (
    wrhs_id bigint NOT NULL,
    book_id bigint NOT NULL,
    qty integer DEFAULT 0 NOT NULL
);


ALTER TABLE store.warehouse_books OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 1063458)
-- Name: warehouses; Type: TABLE; Schema: store; Owner: postgres
--

CREATE TABLE store.warehouses (
    id bigint DEFAULT nextval(('store.warehouse_seq'::text)::regclass) NOT NULL,
    address character varying(255),
    capacity integer,
    CONSTRAINT warehouse_chk CHECK ((capacity > 0))
);


ALTER TABLE store.warehouses OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 1063456)
-- Name: warehouse_seq; Type: SEQUENCE; Schema: store; Owner: postgres
--

CREATE SEQUENCE store.warehouse_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE store.warehouse_seq OWNER TO postgres;

--
-- TOC entry 2930 (class 0 OID 0)
-- Dependencies: 218
-- Name: warehouse_seq; Type: SEQUENCE OWNED BY; Schema: store; Owner: postgres
--

ALTER SEQUENCE store.warehouse_seq OWNED BY store.warehouses.id;


--
-- TOC entry 2910 (class 0 OID 1063367)
-- Dependencies: 213
-- Data for Name: admins; Type: TABLE DATA; Schema: store; Owner: postgres
--

COPY store.admins (id, login, salt, pass) FROM stdin;
1	admin	MMYch597vNqIIUjf4eUAmQAllQ6mUVIlveN7Zb1UWDwnqavp5JryszSDgKgcKMkY	5437e3dfe0348a08a078ef6a95d0fbca549fea506ab83aa9fb32c6624eeda361
\.


--
-- TOC entry 2912 (class 0 OID 1063398)
-- Dependencies: 215
-- Data for Name: authors; Type: TABLE DATA; Schema: store; Owner: postgres
--

COPY store.authors (id, first_name, last_name, middle_name, birth_date, country) FROM stdin;
1	Александр	Пушкин	Сергеевич	1799-06-06	Россия
2	Фёдор	Достоевский	Михайлович	1821-11-11	Россия
3	Лев	Толстой	Николаевич	1828-09-09	Россия
4	Эрих	Мария	Ремарк	1898-06-22	Германия
5	Джейн	Остин	Неизвестно	1775-12-16	Англия
6	Джордж	Оруэлл	Неизвестно	1903-06-25	Англия
7	Габриэль	Гарсиа	Маркес	1927-03-06	Колумбия
8	Артур	Конан	Дойл	1859-05-22	Англия
9	Рэй	Брэдбери	Дуглас	1920-08-22	США
10	Стивен	Кинг	Эдвин	1947-09-21	США
\.


--
-- TOC entry 2903 (class 0 OID 1063291)
-- Dependencies: 206
-- Data for Name: books; Type: TABLE DATA; Schema: store; Owner: postgres
--

COPY store.books (id, isbn, title, descr, price, publisher_id, author_id, publication_year, genre) FROM stdin;
1	978-5-9925-1230-4	Евгений Онегин	\N	500.00	1	1	1833	Роман в стихах
2	978-5-04-171716-2	Преступление и наказание	\N	600.00	2	2	1866	Роман
3	978-5-04-115995-5	Война и мир	\N	800.00	3	3	1869	Исторический роман
4	978-5-17-173239-4	Три товарища	\N	550.00	4	4	1936	Драма
5	978-5-04-171591-5	Гордость и предубеждение	\N	400.00	5	5	1813	Роман
6	978-5-17-148844-4	1984	\N	450.00	6	6	1949	Антиутопия
7	978-5-17-163296-0	Сто лет одиночества	\N	700.00	7	7	1967	Магический реализм
8	978-5-04-099148-8	Шерлок Холмс: Собрание рассказов	\N	650.00	8	8	1892	Детектив
9	978-5-04-116506-2	451 градус по Фаренгейту	\N	480.00	9	9	1953	Фантастика
10	978-5-17-112489-2	Сияние	…Проходят годы, десятилетия, но потрясающая история писателя Джека Торранса, его сынишки Дэнни, наделенного необычным даром, и поединка с темными силами, обитающими в роскошном отеле «Оверлук», по-прежнему завораживает и держит в неослабевающем напряжении читателей самого разного возраста…	550.00	10	10	1977	Хоррор
\.


--
-- TOC entry 2914 (class 0 OID 1063424)
-- Dependencies: 217
-- Data for Name: clients; Type: TABLE DATA; Schema: store; Owner: postgres
--

COPY store.clients (id, first_name, last_name, middle_name, login, salt, password) FROM stdin;
1	user	user	user	user	JUspx0w8JwwS0JrLKB8fBzcU4Vh6VZHbBVTeHS7V0PCQNBhaR22Qy0p733D9CD3s	d1ff5350aa4f00e5787a5887335bc220e9060f8651678707e4731f03e2b5ca70
\.


--
-- TOC entry 2906 (class 0 OID 1063341)
-- Dependencies: 209
-- Data for Name: order_items; Type: TABLE DATA; Schema: store; Owner: postgres
--

COPY store.order_items (order_id, book_id, item_price, qty) FROM stdin;
\.


--
-- TOC entry 2905 (class 0 OID 1063322)
-- Dependencies: 208
-- Data for Name: orders; Type: TABLE DATA; Schema: store; Owner: postgres
--

COPY store.orders (id, client_id, amount, dt, ship_name, ship_address, ship_city, ship_zip_code, ship_country, status) FROM stdin;
\.


--
-- TOC entry 2908 (class 0 OID 1063350)
-- Dependencies: 211
-- Data for Name: publishers; Type: TABLE DATA; Schema: store; Owner: postgres
--

COPY store.publishers (id, name, country, phone) FROM stdin;
1	АСТ	Россия	+7 495 123-45-67
2	Эксмо	Россия	+7 495 234-56-78
3	Penguin Books	Англия	+44 20 1234 5678
4	HarperCollins	США	+1 212 555 7890
5	Random House	США	+1 212 555 1234
6	Macmillan	Англия	+44 20 9876 5432
7	Hachette	Франция	+33 1 2345 6789
8	Simon & Schuster	США	+1 212 555 6789
9	Alianza Editorial	Испания	+34 91 123 4567
10	Planeta	Испания	+34 91 234 5678
\.


--
-- TOC entry 2917 (class 0 OID 1063470)
-- Dependencies: 220
-- Data for Name: warehouse_books; Type: TABLE DATA; Schema: store; Owner: postgres
--

COPY store.warehouse_books (wrhs_id, book_id, qty) FROM stdin;
\.


--
-- TOC entry 2916 (class 0 OID 1063458)
-- Dependencies: 219
-- Data for Name: warehouses; Type: TABLE DATA; Schema: store; Owner: postgres
--

COPY store.warehouses (id, address, capacity) FROM stdin;
1	Москва, ул. Логистическая, д. 1	5000
2	Санкт-Петербург, проспект Складской, д. 7	4000
3	Екатеринбург, ул. Производственная, д. 12	3500
4	Казань, ул. Индустриальная, д. 5	3000
5	Новосибирск, ул. Товарная, д. 3	2500
6	Ростов-на-Дону, ул. Грузовая, д. 8	2800
7	Краснодар, ул. Логистическая, д. 14	2200
8	Самара, ул. Транспортная, д. 6	2700
9	Воронеж, ул. Хранения, д. 9	2400
10	Уфа, ул. Доставочная, д. 11	2600
\.


--
-- TOC entry 2931 (class 0 OID 0)
-- Dependencies: 212
-- Name: admin_seq; Type: SEQUENCE SET; Schema: store; Owner: postgres
--

SELECT pg_catalog.setval('store.admin_seq', 1, true);


--
-- TOC entry 2932 (class 0 OID 0)
-- Dependencies: 214
-- Name: author_seq; Type: SEQUENCE SET; Schema: store; Owner: postgres
--

SELECT pg_catalog.setval('store.author_seq', 10, true);


--
-- TOC entry 2933 (class 0 OID 0)
-- Dependencies: 205
-- Name: book_seq; Type: SEQUENCE SET; Schema: store; Owner: postgres
--

SELECT pg_catalog.setval('store.book_seq', 11, true);


--
-- TOC entry 2934 (class 0 OID 0)
-- Dependencies: 216
-- Name: client_seq; Type: SEQUENCE SET; Schema: store; Owner: postgres
--

SELECT pg_catalog.setval('store.client_seq', 1, true);


--
-- TOC entry 2935 (class 0 OID 0)
-- Dependencies: 207
-- Name: order_seq; Type: SEQUENCE SET; Schema: store; Owner: postgres
--

SELECT pg_catalog.setval('store.order_seq', 1, false);


--
-- TOC entry 2936 (class 0 OID 0)
-- Dependencies: 210
-- Name: publisher_seq; Type: SEQUENCE SET; Schema: store; Owner: postgres
--

SELECT pg_catalog.setval('store.publisher_seq', 10, true);


--
-- TOC entry 2937 (class 0 OID 0)
-- Dependencies: 218
-- Name: warehouse_seq; Type: SEQUENCE SET; Schema: store; Owner: postgres
--

SELECT pg_catalog.setval('store.warehouse_seq', 10, true);


--
-- TOC entry 2762 (class 2606 OID 1063372)
-- Name: admins admins_pkey; Type: CONSTRAINT; Schema: store; Owner: postgres
--

ALTER TABLE ONLY store.admins
    ADD CONSTRAINT admins_pkey PRIMARY KEY (id);


--
-- TOC entry 2764 (class 2606 OID 1063403)
-- Name: authors authors_pkey; Type: CONSTRAINT; Schema: store; Owner: postgres
--

ALTER TABLE ONLY store.authors
    ADD CONSTRAINT authors_pkey PRIMARY KEY (id);


--
-- TOC entry 2753 (class 2606 OID 1063299)
-- Name: books books_pkey; Type: CONSTRAINT; Schema: store; Owner: postgres
--

ALTER TABLE ONLY store.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (id);


--
-- TOC entry 2767 (class 2606 OID 1063432)
-- Name: clients customers_pkey; Type: CONSTRAINT; Schema: store; Owner: postgres
--

ALTER TABLE ONLY store.clients
    ADD CONSTRAINT customers_pkey PRIMARY KEY (id);


--
-- TOC entry 2758 (class 2606 OID 1063347)
-- Name: order_items order_items_key; Type: CONSTRAINT; Schema: store; Owner: postgres
--

ALTER TABLE ONLY store.order_items
    ADD CONSTRAINT order_items_key UNIQUE (order_id, book_id);


--
-- TOC entry 2756 (class 2606 OID 1063328)
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: store; Owner: postgres
--

ALTER TABLE ONLY store.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);


--
-- TOC entry 2760 (class 2606 OID 1063355)
-- Name: publishers publishers_pkey; Type: CONSTRAINT; Schema: store; Owner: postgres
--

ALTER TABLE ONLY store.publishers
    ADD CONSTRAINT publishers_pkey PRIMARY KEY (id);


--
-- TOC entry 2771 (class 2606 OID 1063474)
-- Name: warehouse_books warehouse_books_idx; Type: CONSTRAINT; Schema: store; Owner: postgres
--

ALTER TABLE ONLY store.warehouse_books
    ADD CONSTRAINT warehouse_books_idx PRIMARY KEY (wrhs_id, book_id);


--
-- TOC entry 2769 (class 2606 OID 1063463)
-- Name: warehouses warehouse_pkey; Type: CONSTRAINT; Schema: store; Owner: postgres
--

ALTER TABLE ONLY store.warehouses
    ADD CONSTRAINT warehouse_pkey PRIMARY KEY (id);


--
-- TOC entry 2748 (class 1259 OID 1063305)
-- Name: books_idx; Type: INDEX; Schema: store; Owner: postgres
--

CREATE UNIQUE INDEX books_idx ON store.books USING btree (isbn);


--
-- TOC entry 2749 (class 1259 OID 1063303)
-- Name: books_idx1; Type: INDEX; Schema: store; Owner: postgres
--

CREATE INDEX books_idx1 ON store.books USING btree (title);


--
-- TOC entry 2750 (class 1259 OID 1063304)
-- Name: books_idx2; Type: INDEX; Schema: store; Owner: postgres
--

CREATE INDEX books_idx2 ON store.books USING btree (publisher_id);


--
-- TOC entry 2751 (class 1259 OID 1079840)
-- Name: books_idx3; Type: INDEX; Schema: store; Owner: postgres
--

CREATE INDEX books_idx3 ON store.books USING btree (genre);


--
-- TOC entry 2765 (class 1259 OID 1063433)
-- Name: customers_idx; Type: INDEX; Schema: store; Owner: postgres
--

CREATE UNIQUE INDEX customers_idx ON store.clients USING btree (login);


--
-- TOC entry 2754 (class 1259 OID 1063329)
-- Name: orders_idx; Type: INDEX; Schema: store; Owner: postgres
--

CREATE INDEX orders_idx ON store.orders USING btree (client_id);


--
-- TOC entry 2772 (class 2606 OID 1063356)
-- Name: books books_fk; Type: FK CONSTRAINT; Schema: store; Owner: postgres
--

ALTER TABLE ONLY store.books
    ADD CONSTRAINT books_fk FOREIGN KEY (publisher_id) REFERENCES store.publishers(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 2773 (class 2606 OID 1063442)
-- Name: orders orders_fk; Type: FK CONSTRAINT; Schema: store; Owner: postgres
--

ALTER TABLE ONLY store.orders
    ADD CONSTRAINT orders_fk FOREIGN KEY (client_id) REFERENCES store.clients(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 2774 (class 2606 OID 1063475)
-- Name: warehouse_books warehouse_books_fk; Type: FK CONSTRAINT; Schema: store; Owner: postgres
--

ALTER TABLE ONLY store.warehouse_books
    ADD CONSTRAINT warehouse_books_fk FOREIGN KEY (book_id) REFERENCES store.books(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 2775 (class 2606 OID 1063480)
-- Name: warehouse_books warehouse_books_fk1; Type: FK CONSTRAINT; Schema: store; Owner: postgres
--

ALTER TABLE ONLY store.warehouse_books
    ADD CONSTRAINT warehouse_books_fk1 FOREIGN KEY (wrhs_id) REFERENCES store.warehouses(id);


-- Completed on 2025-06-05 11:27:09

--
-- PostgreSQL database dump complete
--

