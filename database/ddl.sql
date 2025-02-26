CREATE TABLE public.m_role (
	id int8 GENERATED BY DEFAULT AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL,
	"name" VARCHAR(255) NULL,
	created_at TIMESTAMP(6) NULL,
	updated_at TIMESTAMP(6) NULL,
	created_by int8 NULL,
	updated_by int8 NULL,
	created_by_name VARCHAR(255) NULL,
	updated_by_name VARCHAR(255) NULL,
	CONSTRAINT m_role_pkey PRIMARY KEY (id)
);

CREATE TABLE public.m_user (
	id int8 GENERATED BY DEFAULT AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL,
	"name" VARCHAR(255) NULL,
	username VARCHAR(255) NULL,
	"password" VARCHAR(255) NULL,
	role_id int8 NULL,
	created_at TIMESTAMP(6) NULL,
	updated_at TIMESTAMP(6) NULL,
	created_by int8 NULL,
	updated_by int8 NULL,
	created_by_name VARCHAR(255) NULL,
	updated_by_name VARCHAR(255) NULL,
	CONSTRAINT m_user_pkey PRIMARY KEY (id),
	CONSTRAINT user_role_id FOREIGN KEY (role_id) REFERENCES public.m_role (id)
);

CREATE TABLE public.t_refresh_token (
	id int8 GENERATED BY DEFAULT AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL,
	token VARCHAR(255) NULL,
	validity TIMESTAMP(6) NULL,
	user_id int8 NULL,
	CONSTRAINT t_refresh_token_pkey PRIMARY KEY (id),
	CONSTRAINT refresh_token_user_id FOREIGN KEY (user_id) REFERENCES public.m_user (id)
);

CREATE TABLE public.m_position (
	id int8 GENERATED BY DEFAULT AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL,
	"name" VARCHAR(255) NULL,
	created_at TIMESTAMP(6) NULL,
	updated_at TIMESTAMP(6) NULL,
	created_by int8 NULL,
	updated_by int8 NULL,
	created_by_name VARCHAR(255) NULL,
	updated_by_name VARCHAR(255) NULL,
	CONSTRAINT m_position_pkey PRIMARY KEY (id)
);

CREATE TABLE public.m_basic_salary (
	id int8 GENERATED BY DEFAULT AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL,
	salary_amount int8 NULL,
	total_year int4 NULL,
	position_id int8 NULL,
	created_at TIMESTAMP(6) NULL,
	updated_at TIMESTAMP(6) NULL,
	created_by int8 NULL,
	updated_by int8 NULL,
	created_by_name VARCHAR(255) NULL,
	updated_by_name VARCHAR(255) NULL,
	CONSTRAINT m_basic_salary_pkey PRIMARY KEY (id),
	CONSTRAINT basic_salary_position_id FOREIGN KEY (position_id) REFERENCES public.m_position (id)
);

CREATE TABLE public.m_allowance_type (
	id int8 GENERATED BY DEFAULT AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL,
	"name" VARCHAR(255) NULL,
	created_at TIMESTAMP(6) NULL,
	updated_at TIMESTAMP(6) NULL,
	created_by int8 NULL,
	updated_by int8 NULL,
	created_by_name VARCHAR(255) NULL,
	updated_by_name VARCHAR(255) NULL,
	CONSTRAINT m_allowance_type_pkey PRIMARY KEY (id)
);

CREATE TABLE public.m_allowance (
	id int8 GENERATED BY DEFAULT AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL,
	position_id int8 NULL,
	allowance_type_id int8 NULL,
	allowance_amount int8 NULL,
	created_at TIMESTAMP(6) NULL,
	updated_at TIMESTAMP(6) NULL,
	created_by int8 NULL,
	updated_by int8 NULL,
	created_by_name VARCHAR(255) NULL,
	updated_by_name VARCHAR(255) NULL,
	CONSTRAINT m_allowance_pkey PRIMARY KEY (id),
	CONSTRAINT allowance_position_id FOREIGN KEY (position_id) REFERENCES public.m_position (id),
	CONSTRAINT allowance_allowance_type_id FOREIGN KEY (allowance_type_id) REFERENCES public.m_allowance_type (id)
);

CREATE TABLE public.m_employee (
	id int8 GENERATED BY DEFAULT AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL,
	"name" VARCHAR(255) NULL,
	phone_number VARCHAR(255) NULL,
	email VARCHAR(255) NULL,
	"address" VARCHAR(255) NULL,
	work_status VARCHAR(255) NULL,
	bank_account_number VARCHAR(255) NULL,
	bank_account_name VARCHAR(255) NULL,
	npwp VARCHAR(255) NULL,
	date_of_birth DATE NULL,
	join_date DATE NULL,
	is_married BOOLEAN NULL,
	total_child int4 NULL,
	position_id int8 NULL,
	created_at TIMESTAMP(6) NULL,
	updated_at TIMESTAMP(6) NULL,
	created_by int8 NULL,
	updated_by int8 NULL,
	created_by_name VARCHAR(255) NULL,
	updated_by_name VARCHAR(255) NULL,
	CONSTRAINT m_employee_pkey PRIMARY KEY (id),
	CONSTRAINT employee_position_id FOREIGN KEY (position_id) REFERENCES public.m_position (id)
);