CREATE TABLE tahistory (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    source TEXT NULL,
    symbol TEXT NULL,
    high REAL NULL,
    last_price REAL NULL,
    created_at TEXT NULL,
    book TEXT NULL,
    volume REAL NULL,
    vwap REAL NULL,
    low REAL NULL,
    ask REAL NULL,
    bid REAL NULL,
    change24 REAL NULL,
);