CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    firebase_id VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE groups (
    group_id SERIAL PRIMARY KEY,
    group_name VARCHAR(100) NOT NULL,
    created_by INTEGER REFERENCES users(user_id) ON DELETE CASCADE
);


CREATE TABLE group_members (
    group_member_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(user_id) ON DELETE CASCADE,
    group_id INTEGER REFERENCES groups(group_id) ON DELETE CASCADE,
    joined_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE transactions (
    transaction_id SERIAL PRIMARY KEY,
    lender_id INTEGER REFERENCES users(user_id) ON DELETE CASCADE,
    borrower_id INTEGER REFERENCES users(user_id) ON DELETE CASCADE,
    group_id INTEGER REFERENCES groups(group_id) ON DELETE SET NULL,
    amount NUMERIC(10, 2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    purpose VARCHAR(255)
);


CREATE TABLE transaction_splits (
    transaction_split_id SERIAL PRIMARY KEY,
    transaction_id INTEGER REFERENCES transactions(transaction_id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users(user_id) ON DELETE CASCADE,
    amount NUMERIC(10, 2) NOT NULL
);


CREATE TABLE balances (
    balance_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(user_id) ON DELETE CASCADE,
    group_id INTEGER REFERENCES groups(group_id) ON DELETE SET NULL,
    owed_amount NUMERIC(10, 2) NOT NULL DEFAULT 0,
    lent_amount NUMERIC(10, 2) NOT NULL DEFAULT 0
);


CREATE TABLE requests (
    request_id SERIAL PRIMARY KEY,
    sender_id INTEGER REFERENCES users(user_id) ON DELETE CASCADE,
    receiver_id INTEGER REFERENCES users(user_id) ON DELETE SET NULL,
    group_id INTEGER REFERENCES groups(group_id) ON DELETE SET NULL,
    amount NUMERIC(10, 2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE settlements (
    settlement_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(user_id) ON DELETE CASCADE,
    counterparty_id INTEGER REFERENCES users(user_id) ON DELETE SET NULL,
    group_id INTEGER REFERENCES groups(group_id) ON DELETE SET NULL,
    amount NUMERIC(10, 2) NOT NULL,
    settlement_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- Add unique indexes on `users` table
CREATE UNIQUE INDEX idx_users_email ON users (email);
CREATE UNIQUE INDEX idx_users_firebase_id ON users (firebase_id);

-- Index on foreign keys for faster lookups
CREATE INDEX idx_group_members_user_id ON group_members (user_id);
CREATE INDEX idx_group_members_group_id ON group_members (group_id);

CREATE INDEX idx_transactions_lender_id ON transactions (lender_id);
CREATE INDEX idx_transactions_borrower_id ON transactions (borrower_id);

CREATE INDEX idx_transaction_splits_transaction_id ON transaction_splits (transaction_id);
CREATE INDEX idx_transaction_splits_user_id ON transaction_splits (user_id);

CREATE INDEX idx_balances_user_id ON balances (user_id);
CREATE INDEX idx_balances_group_id ON balances (group_id);



