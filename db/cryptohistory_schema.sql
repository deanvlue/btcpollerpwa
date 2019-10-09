CREATE TABLE sqlite_sequence(name,seq);
CREATE TABLE tahistory (
id integer primary key autoincrement,
source text null,
symbol text null,
high real null, last_price real null, created_at text null, book text null, volume real null, vwap real null, low real null, ask real null, bid real null, change24 real null);
