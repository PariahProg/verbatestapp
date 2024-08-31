CREATE TABLE public.tasks (
	id integer GENERATED ALWAYS AS IDENTITY NOT NULL,
	title text NULL,
	description text NULL,
	due_date timestamp with time zone NOT NULL,
	created_at timestamp with time zone NOT NULL,
	updated_at timestamp with time zone NOT NULL,
	CONSTRAINT tasks_pk PRIMARY KEY (id)
);