--
-- PostgreSQL database dump
--

-- Dumped from database version 16.6 (Ubuntu 16.6-0ubuntu0.24.04.1)
-- Dumped by pg_dump version 16.6 (Ubuntu 16.6-0ubuntu0.24.04.1)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: balances; Type: TABLE; Schema: public; Owner: hornet
--

CREATE TABLE public.balances (
    balance_id integer NOT NULL,
    user_id integer NOT NULL,
    group_id integer,
    owed_amount numeric(10,2) DEFAULT 0,
    lent_amount numeric(10,2) DEFAULT 0
);


ALTER TABLE public.balances OWNER TO hornet;

--
-- Name: balances_balance_id_seq; Type: SEQUENCE; Schema: public; Owner: hornet
--

CREATE SEQUENCE public.balances_balance_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.balances_balance_id_seq OWNER TO hornet;

--
-- Name: balances_balance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hornet
--

ALTER SEQUENCE public.balances_balance_id_seq OWNED BY public.balances.balance_id;


--
-- Name: group_members; Type: TABLE; Schema: public; Owner: hornet
--

CREATE TABLE public.group_members (
    group_member_id integer NOT NULL,
    user_id integer NOT NULL,
    group_id integer NOT NULL,
    joined_date timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.group_members OWNER TO hornet;

--
-- Name: group_members_group_member_id_seq; Type: SEQUENCE; Schema: public; Owner: hornet
--

CREATE SEQUENCE public.group_members_group_member_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.group_members_group_member_id_seq OWNER TO hornet;

--
-- Name: group_members_group_member_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hornet
--

ALTER SEQUENCE public.group_members_group_member_id_seq OWNED BY public.group_members.group_member_id;


--
-- Name: groups; Type: TABLE; Schema: public; Owner: hornet
--

CREATE TABLE public.groups (
    group_id integer NOT NULL,
    group_name character varying(255),
    created_by integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.groups OWNER TO hornet;

--
-- Name: groups_group_id_seq; Type: SEQUENCE; Schema: public; Owner: hornet
--

CREATE SEQUENCE public.groups_group_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.groups_group_id_seq OWNER TO hornet;

--
-- Name: groups_group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hornet
--

ALTER SEQUENCE public.groups_group_id_seq OWNED BY public.groups.group_id;


--
-- Name: payment_methods; Type: TABLE; Schema: public; Owner: hornet
--

CREATE TABLE public.payment_methods (
    payment_id integer NOT NULL,
    user_id integer NOT NULL,
    payment_type character varying(50) NOT NULL,
    upi_id character varying(100),
    account_number character varying(20),
    ifsc_code character varying(15),
    wallet_provider character varying(50),
    is_primary boolean DEFAULT false,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT payment_methods_payment_type_check CHECK (((payment_type)::text = ANY ((ARRAY['UPI'::character varying, 'Bank Account'::character varying])::text[])))
);


ALTER TABLE public.payment_methods OWNER TO hornet;

--
-- Name: payment_methods_payment_id_seq; Type: SEQUENCE; Schema: public; Owner: hornet
--

CREATE SEQUENCE public.payment_methods_payment_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.payment_methods_payment_id_seq OWNER TO hornet;

--
-- Name: payment_methods_payment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hornet
--

ALTER SEQUENCE public.payment_methods_payment_id_seq OWNED BY public.payment_methods.payment_id;


--
-- Name: payment_notifications; Type: TABLE; Schema: public; Owner: hornet
--

CREATE TABLE public.payment_notifications (
    notification_id integer NOT NULL,
    user_id integer NOT NULL,
    transaction_id integer NOT NULL,
    notification_type character varying(50) NOT NULL,
    message text NOT NULL,
    is_read boolean DEFAULT false,
    read_at timestamp without time zone,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.payment_notifications OWNER TO hornet;

--
-- Name: payment_notifications_notification_id_seq; Type: SEQUENCE; Schema: public; Owner: hornet
--

CREATE SEQUENCE public.payment_notifications_notification_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.payment_notifications_notification_id_seq OWNER TO hornet;

--
-- Name: payment_notifications_notification_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hornet
--

ALTER SEQUENCE public.payment_notifications_notification_id_seq OWNED BY public.payment_notifications.notification_id;


--
-- Name: requests; Type: TABLE; Schema: public; Owner: hornet
--

CREATE TABLE public.requests (
    request_id integer NOT NULL,
    sender_id integer NOT NULL,
    receiver_id integer NOT NULL,
    group_id integer,
    amount numeric(10,2) NOT NULL,
    status character varying(50) DEFAULT 'pending'::character varying,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.requests OWNER TO hornet;

--
-- Name: requests_request_id_seq; Type: SEQUENCE; Schema: public; Owner: hornet
--

CREATE SEQUENCE public.requests_request_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.requests_request_id_seq OWNER TO hornet;

--
-- Name: requests_request_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hornet
--

ALTER SEQUENCE public.requests_request_id_seq OWNED BY public.requests.request_id;


--
-- Name: settlements; Type: TABLE; Schema: public; Owner: hornet
--

CREATE TABLE public.settlements (
    settlement_id integer NOT NULL,
    user_id integer NOT NULL,
    counterparty_id integer NOT NULL,
    group_id integer,
    amount numeric(10,2) NOT NULL,
    settlement_date timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.settlements OWNER TO hornet;

--
-- Name: settlements_settlement_id_seq; Type: SEQUENCE; Schema: public; Owner: hornet
--

CREATE SEQUENCE public.settlements_settlement_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.settlements_settlement_id_seq OWNER TO hornet;

--
-- Name: settlements_settlement_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hornet
--

ALTER SEQUENCE public.settlements_settlement_id_seq OWNED BY public.settlements.settlement_id;


--
-- Name: transaction_logs; Type: TABLE; Schema: public; Owner: hornet
--

CREATE TABLE public.transaction_logs (
    log_id integer NOT NULL,
    transaction_id integer NOT NULL,
    action character varying(100) NOT NULL,
    status character varying(50),
    details text,
    error_details text,
    "timestamp" timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.transaction_logs OWNER TO hornet;

--
-- Name: transaction_logs_log_id_seq; Type: SEQUENCE; Schema: public; Owner: hornet
--

CREATE SEQUENCE public.transaction_logs_log_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.transaction_logs_log_id_seq OWNER TO hornet;

--
-- Name: transaction_logs_log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hornet
--

ALTER SEQUENCE public.transaction_logs_log_id_seq OWNED BY public.transaction_logs.log_id;


--
-- Name: transaction_splits; Type: TABLE; Schema: public; Owner: hornet
--

CREATE TABLE public.transaction_splits (
    transaction_split_id integer NOT NULL,
    transaction_id integer NOT NULL,
    user_id integer NOT NULL,
    amount numeric(10,2) NOT NULL
);


ALTER TABLE public.transaction_splits OWNER TO hornet;

--
-- Name: transaction_splits_transaction_split_id_seq; Type: SEQUENCE; Schema: public; Owner: hornet
--

CREATE SEQUENCE public.transaction_splits_transaction_split_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.transaction_splits_transaction_split_id_seq OWNER TO hornet;

--
-- Name: transaction_splits_transaction_split_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hornet
--

ALTER SEQUENCE public.transaction_splits_transaction_split_id_seq OWNED BY public.transaction_splits.transaction_split_id;


--
-- Name: transactions; Type: TABLE; Schema: public; Owner: hornet
--

CREATE TABLE public.transactions (
    transaction_id integer NOT NULL,
    lender_id integer NOT NULL,
    borrower_id integer NOT NULL,
    group_id integer,
    amount numeric(10,2) NOT NULL,
    status character varying(50),
    purpose character varying(255),
    payment_method_id integer,
    retry_count integer DEFAULT 0,
    failure_reason text,
    CONSTRAINT check_transaction_status CHECK (((status)::text = ANY ((ARRAY['pending'::character varying, 'successful'::character varying, 'failed'::character varying, 'retrying'::character varying])::text[])))
);


ALTER TABLE public.transactions OWNER TO hornet;

--
-- Name: transactions_transaction_id_seq; Type: SEQUENCE; Schema: public; Owner: hornet
--

CREATE SEQUENCE public.transactions_transaction_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.transactions_transaction_id_seq OWNER TO hornet;

--
-- Name: transactions_transaction_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hornet
--

ALTER SEQUENCE public.transactions_transaction_id_seq OWNED BY public.transactions.transaction_id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: hornet
--

CREATE TABLE public.users (
    user_id integer NOT NULL,
    email character varying(255),
    firebase_id character varying(255),
    name character varying(100),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    verified boolean DEFAULT false NOT NULL
);


ALTER TABLE public.users OWNER TO hornet;

--
-- Name: users_user_id_seq; Type: SEQUENCE; Schema: public; Owner: hornet
--

CREATE SEQUENCE public.users_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_user_id_seq OWNER TO hornet;

--
-- Name: users_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hornet
--

ALTER SEQUENCE public.users_user_id_seq OWNED BY public.users.user_id;


--
-- Name: balances balance_id; Type: DEFAULT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.balances ALTER COLUMN balance_id SET DEFAULT nextval('public.balances_balance_id_seq'::regclass);


--
-- Name: group_members group_member_id; Type: DEFAULT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.group_members ALTER COLUMN group_member_id SET DEFAULT nextval('public.group_members_group_member_id_seq'::regclass);


--
-- Name: groups group_id; Type: DEFAULT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.groups ALTER COLUMN group_id SET DEFAULT nextval('public.groups_group_id_seq'::regclass);


--
-- Name: payment_methods payment_id; Type: DEFAULT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.payment_methods ALTER COLUMN payment_id SET DEFAULT nextval('public.payment_methods_payment_id_seq'::regclass);


--
-- Name: payment_notifications notification_id; Type: DEFAULT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.payment_notifications ALTER COLUMN notification_id SET DEFAULT nextval('public.payment_notifications_notification_id_seq'::regclass);


--
-- Name: requests request_id; Type: DEFAULT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.requests ALTER COLUMN request_id SET DEFAULT nextval('public.requests_request_id_seq'::regclass);


--
-- Name: settlements settlement_id; Type: DEFAULT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.settlements ALTER COLUMN settlement_id SET DEFAULT nextval('public.settlements_settlement_id_seq'::regclass);


--
-- Name: transaction_logs log_id; Type: DEFAULT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.transaction_logs ALTER COLUMN log_id SET DEFAULT nextval('public.transaction_logs_log_id_seq'::regclass);


--
-- Name: transaction_splits transaction_split_id; Type: DEFAULT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.transaction_splits ALTER COLUMN transaction_split_id SET DEFAULT nextval('public.transaction_splits_transaction_split_id_seq'::regclass);


--
-- Name: transactions transaction_id; Type: DEFAULT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.transactions ALTER COLUMN transaction_id SET DEFAULT nextval('public.transactions_transaction_id_seq'::regclass);


--
-- Name: users user_id; Type: DEFAULT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.users ALTER COLUMN user_id SET DEFAULT nextval('public.users_user_id_seq'::regclass);


--
-- Data for Name: balances; Type: TABLE DATA; Schema: public; Owner: hornet
--

COPY public.balances (balance_id, user_id, group_id, owed_amount, lent_amount) FROM stdin;
\.


--
-- Data for Name: group_members; Type: TABLE DATA; Schema: public; Owner: hornet
--

COPY public.group_members (group_member_id, user_id, group_id, joined_date) FROM stdin;
\.


--
-- Data for Name: groups; Type: TABLE DATA; Schema: public; Owner: hornet
--

COPY public.groups (group_id, group_name, created_by, created_at) FROM stdin;
\.


--
-- Data for Name: payment_methods; Type: TABLE DATA; Schema: public; Owner: hornet
--

COPY public.payment_methods (payment_id, user_id, payment_type, upi_id, account_number, ifsc_code, wallet_provider, is_primary, created_at) FROM stdin;
\.


--
-- Data for Name: payment_notifications; Type: TABLE DATA; Schema: public; Owner: hornet
--

COPY public.payment_notifications (notification_id, user_id, transaction_id, notification_type, message, is_read, read_at, created_at) FROM stdin;
\.


--
-- Data for Name: requests; Type: TABLE DATA; Schema: public; Owner: hornet
--

COPY public.requests (request_id, sender_id, receiver_id, group_id, amount, status, created_at) FROM stdin;
\.


--
-- Data for Name: settlements; Type: TABLE DATA; Schema: public; Owner: hornet
--

COPY public.settlements (settlement_id, user_id, counterparty_id, group_id, amount, settlement_date) FROM stdin;
\.


--
-- Data for Name: transaction_logs; Type: TABLE DATA; Schema: public; Owner: hornet
--

COPY public.transaction_logs (log_id, transaction_id, action, status, details, error_details, "timestamp") FROM stdin;
\.


--
-- Data for Name: transaction_splits; Type: TABLE DATA; Schema: public; Owner: hornet
--

COPY public.transaction_splits (transaction_split_id, transaction_id, user_id, amount) FROM stdin;
\.


--
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: hornet
--

COPY public.transactions (transaction_id, lender_id, borrower_id, group_id, amount, status, purpose, payment_method_id, retry_count, failure_reason) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: hornet
--

COPY public.users (user_id, email, firebase_id, name, created_at, verified) FROM stdin;
1	kundupriyanka1608@gmail.com	xxxxxxxxx	Priyanka	2024-12-04 00:33:45.579099	f
\.


--
-- Name: balances_balance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hornet
--

SELECT pg_catalog.setval('public.balances_balance_id_seq', 1, false);


--
-- Name: group_members_group_member_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hornet
--

SELECT pg_catalog.setval('public.group_members_group_member_id_seq', 1, false);


--
-- Name: groups_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hornet
--

SELECT pg_catalog.setval('public.groups_group_id_seq', 1, false);


--
-- Name: payment_methods_payment_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hornet
--

SELECT pg_catalog.setval('public.payment_methods_payment_id_seq', 1, false);


--
-- Name: payment_notifications_notification_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hornet
--

SELECT pg_catalog.setval('public.payment_notifications_notification_id_seq', 1, false);


--
-- Name: requests_request_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hornet
--

SELECT pg_catalog.setval('public.requests_request_id_seq', 1, false);


--
-- Name: settlements_settlement_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hornet
--

SELECT pg_catalog.setval('public.settlements_settlement_id_seq', 1, false);


--
-- Name: transaction_logs_log_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hornet
--

SELECT pg_catalog.setval('public.transaction_logs_log_id_seq', 1, false);


--
-- Name: transaction_splits_transaction_split_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hornet
--

SELECT pg_catalog.setval('public.transaction_splits_transaction_split_id_seq', 1, false);


--
-- Name: transactions_transaction_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hornet
--

SELECT pg_catalog.setval('public.transactions_transaction_id_seq', 1, false);


--
-- Name: users_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hornet
--

SELECT pg_catalog.setval('public.users_user_id_seq', 1, true);


--
-- Name: balances balances_pkey; Type: CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.balances
    ADD CONSTRAINT balances_pkey PRIMARY KEY (balance_id);


--
-- Name: group_members group_members_pkey; Type: CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.group_members
    ADD CONSTRAINT group_members_pkey PRIMARY KEY (group_member_id);


--
-- Name: groups groups_pkey; Type: CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT groups_pkey PRIMARY KEY (group_id);


--
-- Name: payment_methods payment_methods_pkey; Type: CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.payment_methods
    ADD CONSTRAINT payment_methods_pkey PRIMARY KEY (payment_id);


--
-- Name: payment_notifications payment_notifications_pkey; Type: CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.payment_notifications
    ADD CONSTRAINT payment_notifications_pkey PRIMARY KEY (notification_id);


--
-- Name: requests requests_pkey; Type: CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.requests
    ADD CONSTRAINT requests_pkey PRIMARY KEY (request_id);


--
-- Name: settlements settlements_pkey; Type: CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.settlements
    ADD CONSTRAINT settlements_pkey PRIMARY KEY (settlement_id);


--
-- Name: transaction_logs transaction_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.transaction_logs
    ADD CONSTRAINT transaction_logs_pkey PRIMARY KEY (log_id);


--
-- Name: transaction_splits transaction_splits_pkey; Type: CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.transaction_splits
    ADD CONSTRAINT transaction_splits_pkey PRIMARY KEY (transaction_split_id);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (transaction_id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_firebase_id_key; Type: CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_firebase_id_key UNIQUE (firebase_id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);


--
-- Name: idx_balances_group_id; Type: INDEX; Schema: public; Owner: hornet
--

CREATE INDEX idx_balances_group_id ON public.balances USING btree (group_id);


--
-- Name: idx_balances_user_id; Type: INDEX; Schema: public; Owner: hornet
--

CREATE INDEX idx_balances_user_id ON public.balances USING btree (user_id);


--
-- Name: idx_payment_methods_user_id; Type: INDEX; Schema: public; Owner: hornet
--

CREATE INDEX idx_payment_methods_user_id ON public.payment_methods USING btree (user_id);


--
-- Name: idx_payment_notifications_transaction_id; Type: INDEX; Schema: public; Owner: hornet
--

CREATE INDEX idx_payment_notifications_transaction_id ON public.payment_notifications USING btree (transaction_id);


--
-- Name: idx_payment_notifications_user_id; Type: INDEX; Schema: public; Owner: hornet
--

CREATE INDEX idx_payment_notifications_user_id ON public.payment_notifications USING btree (user_id);


--
-- Name: idx_transaction_logs_transaction_id; Type: INDEX; Schema: public; Owner: hornet
--

CREATE INDEX idx_transaction_logs_transaction_id ON public.transaction_logs USING btree (transaction_id);


--
-- Name: idx_transaction_splits_transaction_id; Type: INDEX; Schema: public; Owner: hornet
--

CREATE INDEX idx_transaction_splits_transaction_id ON public.transaction_splits USING btree (transaction_id);


--
-- Name: idx_transaction_splits_user_id; Type: INDEX; Schema: public; Owner: hornet
--

CREATE INDEX idx_transaction_splits_user_id ON public.transaction_splits USING btree (user_id);


--
-- Name: idx_transactions_borrower_id; Type: INDEX; Schema: public; Owner: hornet
--

CREATE INDEX idx_transactions_borrower_id ON public.transactions USING btree (borrower_id);


--
-- Name: idx_transactions_group_id; Type: INDEX; Schema: public; Owner: hornet
--

CREATE INDEX idx_transactions_group_id ON public.transactions USING btree (group_id);


--
-- Name: idx_transactions_lender_id; Type: INDEX; Schema: public; Owner: hornet
--

CREATE INDEX idx_transactions_lender_id ON public.transactions USING btree (lender_id);


--
-- Name: idx_transactions_payment_method_id; Type: INDEX; Schema: public; Owner: hornet
--

CREATE INDEX idx_transactions_payment_method_id ON public.transactions USING btree (payment_method_id);


--
-- Name: balances balances_group_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.balances
    ADD CONSTRAINT balances_group_id_fkey FOREIGN KEY (group_id) REFERENCES public.groups(group_id);


--
-- Name: balances balances_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.balances
    ADD CONSTRAINT balances_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id);


--
-- Name: group_members group_members_group_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.group_members
    ADD CONSTRAINT group_members_group_id_fkey FOREIGN KEY (group_id) REFERENCES public.groups(group_id);


--
-- Name: group_members group_members_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.group_members
    ADD CONSTRAINT group_members_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id);


--
-- Name: groups groups_created_by_fkey; Type: FK CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT groups_created_by_fkey FOREIGN KEY (created_by) REFERENCES public.users(user_id);


--
-- Name: payment_methods payment_methods_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.payment_methods
    ADD CONSTRAINT payment_methods_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON DELETE CASCADE;


--
-- Name: payment_notifications payment_notifications_transaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.payment_notifications
    ADD CONSTRAINT payment_notifications_transaction_id_fkey FOREIGN KEY (transaction_id) REFERENCES public.transactions(transaction_id) ON DELETE CASCADE;


--
-- Name: payment_notifications payment_notifications_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.payment_notifications
    ADD CONSTRAINT payment_notifications_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON DELETE CASCADE;


--
-- Name: requests requests_group_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.requests
    ADD CONSTRAINT requests_group_id_fkey FOREIGN KEY (group_id) REFERENCES public.groups(group_id);


--
-- Name: requests requests_receiver_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.requests
    ADD CONSTRAINT requests_receiver_id_fkey FOREIGN KEY (receiver_id) REFERENCES public.users(user_id);


--
-- Name: requests requests_sender_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.requests
    ADD CONSTRAINT requests_sender_id_fkey FOREIGN KEY (sender_id) REFERENCES public.users(user_id);


--
-- Name: settlements settlements_counterparty_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.settlements
    ADD CONSTRAINT settlements_counterparty_id_fkey FOREIGN KEY (counterparty_id) REFERENCES public.users(user_id);


--
-- Name: settlements settlements_group_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.settlements
    ADD CONSTRAINT settlements_group_id_fkey FOREIGN KEY (group_id) REFERENCES public.groups(group_id);


--
-- Name: settlements settlements_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.settlements
    ADD CONSTRAINT settlements_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id);


--
-- Name: transaction_logs transaction_logs_transaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.transaction_logs
    ADD CONSTRAINT transaction_logs_transaction_id_fkey FOREIGN KEY (transaction_id) REFERENCES public.transactions(transaction_id) ON DELETE CASCADE;


--
-- Name: transaction_splits transaction_splits_transaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.transaction_splits
    ADD CONSTRAINT transaction_splits_transaction_id_fkey FOREIGN KEY (transaction_id) REFERENCES public.transactions(transaction_id);


--
-- Name: transaction_splits transaction_splits_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.transaction_splits
    ADD CONSTRAINT transaction_splits_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id);


--
-- Name: transactions transactions_borrower_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_borrower_id_fkey FOREIGN KEY (borrower_id) REFERENCES public.users(user_id);


--
-- Name: transactions transactions_group_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_group_id_fkey FOREIGN KEY (group_id) REFERENCES public.groups(group_id);


--
-- Name: transactions transactions_lender_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_lender_id_fkey FOREIGN KEY (lender_id) REFERENCES public.users(user_id);


--
-- Name: transactions transactions_payment_method_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: hornet
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_payment_method_id_fkey FOREIGN KEY (payment_method_id) REFERENCES public.payment_methods(payment_id);


--
-- PostgreSQL database dump complete
--

