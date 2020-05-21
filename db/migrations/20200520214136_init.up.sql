create table River (
    ID UUID not null,
    Name text not null,
    Latitude double precision not null,
    Longitude double precision not null,
    primary key (ID)
);

create table Gauge (
    ID UUID not null,
    Name text not null,
    Code text not null,
    RiverID UUID not null,
    Latitude double precision not null,
    Longitude double precision not null,
    primary key (ID),
    constraint Gauge_River_fk foreign key (RiverID)
        references River (ID)
);

create table Metric (
    ID UUID not null,
    GaugeType text not null,
    Value double precision not null,
    Date date not null,
    primary key (ID)
);