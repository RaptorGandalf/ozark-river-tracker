set role postgres;

alter table metrics alter column recorded_date type TIMESTAMP;