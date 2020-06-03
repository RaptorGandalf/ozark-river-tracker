set role postgres;

alter table metrics rename column gauge_type to type;