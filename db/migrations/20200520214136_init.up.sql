create table river (
    id UUID not null,
    name text not null,
    latitude double precision not null,
    longitude double precision not null,
    primary key (id)
);

create table gauge (
    id UUID not null,
    name text not null,
    code text not null,
    river_id UUID not null,
    latitude double precision not null,
    longitude double precision not null,
    primary key (id),
    constraint gauge_river_fk foreign key (river_id)
        references river (id)
);

create table metric (
    id UUID not null,
    gauge_id UUID not null,
    gauge_type text not null,
    value double precision not null,
    recorded_date date not null,
    primary key (id),
    constraint metric_gauge_fk foreign key (gauge_id)
        references gauge (id)
);