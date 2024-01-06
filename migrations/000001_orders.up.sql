create table orders
(
	id_orders serial not null constraint pk_orders primary key,
	data jsonb not null
); 