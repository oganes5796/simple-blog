CREATE TABLE users
(
    id            serial       not null unique,
    username      varchar(255) not null unique,
    role          varchar(255) not null,
    password      varchar(255) not null,
    created_at    TIMESTAMP DEFAULT NOW()
);

CREATE TABLE articles
(
    id            serial       not null unique,
    author_id     int          not null,
    title         varchar(255) not null,
    content       varchar(255) not null,
    created_at    TIMESTAMP DEFAULT NOW()
);

CREATE TABLE comments
(
    id            serial       not null unique,
    article_id    int          not null,
    reader_id     int          not null,
    content       varchar(255) not null,
    created_at    TIMESTAMP DEFAULT NOW()
);