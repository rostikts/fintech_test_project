CREATE TABLE IF NOT EXISTS service (
    id serial primary key,
    name varchar not null
);

CREATE TABLE IF NOT EXISTS payee (
    id serial primary key,
    name varchar not null,
    bank_mfo int not null,
    bank_account varchar not null
);

CREATE TABLE IF NOT EXISTS payment (
    id serial primary key ,
    type varchar,
    number varchar,
    narrative varchar
);

CREATE TABLE IF NOT EXISTS transaction (
    id serial,
    request_id int not null,
    terminal_id int not null,
    partner_object_id int not null,
    payment_id int not null,
    service_id int not null,
    payee_id int not null,
    amount_total int,
    amount_original int,
    commission_ps float,
    commission_client int,
    commission_provider float,
    date_input timestamp,
    date_post timestamp,
    status varchar,
    CONSTRAINT FK_payment FOREIGN KEY (payment_id) REFERENCES payment(id),
    CONSTRAINT FK_service FOREIGN KEY (service_id) REFERENCES service(id),
    CONSTRAINT FK_payee FOREIGN KEY (payee_id) REFERENCES payee(id)

);