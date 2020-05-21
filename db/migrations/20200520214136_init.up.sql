create table River (
    id UUID not null,
    name text not null,
    latitude double precision not null,
    longitude double precision not null,
    primary key (id)
);

create table Gauge (
    id UUID not null,
    name text not null,
    code text not null,
    river_id UUID not null,
    latitude double precision not null,
    longitude double precision not null,
    primary key (id),
    constraint Gauge_River_fk foreign key (river_id)
        references River (id)
);

create table Metric (
    id UUID not null,
    gauge_id UUID not null,
    gauge_type text not null,
    value double precision not null,
    recorded_date date not null,
    primary key (id)
);