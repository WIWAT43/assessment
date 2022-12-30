CREATE TABLE IF NOT EXISTS expenses
(
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    amount FLOAT  NOT NULL DEFAULT 0,
    note TEXT NOT NULL,
    tags TEXT[]
);

CREATE INDEX ON "expenses" ("id");