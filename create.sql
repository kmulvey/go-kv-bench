CREATE TABLE bench_results(
    BENCHNAME       TEXT        NOT NULL,
    BENCHTIME       DATETIME    NOT NULL,
    ALLOC_BYTES     LONG        NOT NULL,
    ALLOC_OP        LONG        NOT NULL,
    NS_OP           LONG        NOT NULL,
    PRIMARY KEY(BENCHNAME, BENCHTIME)
);
