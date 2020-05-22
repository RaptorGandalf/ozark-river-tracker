set role postgres;

alter table river rename to rivers;
alter table gauge rename to gauges;
alter table metric rename to metrics;