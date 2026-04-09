CREATE TABLE students (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    first_name VARCHAR(100) NOT NULL,
    middle_name VARCHAR(100),
    last_name VARCHAR(100) NOT NULL,

    email VARCHAR(150) NOT NULL,

    guardian_name VARCHAR(150),
    guardian_relation VARCHAR(50),
    guardian_contact VARCHAR(20),

    class VARCHAR(20),
    address JSONB
);


CREATE TABLE rank_histories (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    student_id BIGINT NOT NULL,

    term VARCHAR(20) NOT NULL,
    year INT NOT NULL,

    rank INT NOT NULL,
    marks_attained INT,
    grade VARCHAR(5)
);


CREATE TABLE admins (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    email VARCHAR(150),
    password_hash VARCHAR(128) NOT NULL
);