create table if not exists planets (
  id SERIAL primary key,
  planet_type varchar(45) not null,
  planet_short_name varchar(45) not null,
  zone int null,
  age bigint null
);
create unique index planet_short_name_UNIQUE on planets (planet_short_name ASC);
insert into
  planets (planet_type, planet_short_name, zone, age)
values
  ('ground', 'earth', 1, 4500000000),
  ('ground', 'mars', 1, 4650000000),
  ('ground', 'mercury', 1, 4500000000),
  ('ground', 'venus', 1, 4500000000);
CREATE TABLE public.planets_params (
    id serial NOT NULL,
    planet_id int4 NULL,
    planet_radius int8 NULL,
    CONSTRAINT planets_params_pkey PRIMARY KEY (id),
    CONSTRAINT planets_params_planet_id_fkey FOREIGN KEY (planet_id) REFERENCES planets(id) ON UPDATE CASCADE
  );