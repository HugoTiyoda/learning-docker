SET timezone = 'America/Sao_Paulo';
CREATE SCHEMA IF NOT EXISTS estudo;

create table estudo.users(
   id SERIAL PRIMARY KEY,
   "name" VARCHAR(50) NOT NULL,
   "limit" INTEGER NOT NULL 
);

create table estudo.transacions (
	id SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL,
	"value" INTEGER NOT NULL,
    "type" CHAR(1) NOT NULL CHECK ("type" IN ('c', 'd')),
	"description" VARCHAR(10) NOT NULL,
	due_date TIMESTAMP NOT NULL,
	CONSTRAINT fk_users_transactions
		FOREIGN KEY (user_id) REFERENCES estudo.users(id)
);

create table estudo.balances (
	id SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL,
	"value" INTEGER NOT NULL,
	CONSTRAINT fk_users_balances
		FOREIGN KEY (user_id) REFERENCES estudo.users(id)
);

DO $$
BEGIN
  INSERT INTO estudo.users (name, limit)
  VALUES
    ('o barato sai caro', 1000 * 100),
    ('zan corp ltda', 800 * 100),
    ('les cruders', 10000 * 100),
    ('padaria joia de cocaia', 100000 * 100),
    ('kid mais', 5000 * 100);
END; $$