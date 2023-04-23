CREATE TABLE IF NOT EXISTS urls (
    fullurl VARCHAR(100) NOT NULL UNIQUE,
    tinyurl VARCHAR(10) NOT NULL UNIQUE
);

CREATE INDEX IF NOT EXISTS urls_fullurl_idx ON urls USING HASH (fullurl);
CREATE INDEX IF NOT EXISTS urls_tinyurl_idx ON urls USING HASH (tinyurl);