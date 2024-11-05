-- Users Table
CREATE TABLE Users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    firebase_id VARCHAR(255) UNIQUE NOT NULL
);

-- Groups Table
CREATE TABLE Groups (
    group_id SERIAL PRIMARY KEY,
    group_name VARCHAR(100) NOT NULL,
    created_by INT REFERENCES Users(user_id) ON DELETE CASCADE
);
	

-- Group Members Table
CREATE TABLE Group_Members (
    group_member_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES Users(user_id) ON DELETE CASCADE,
    group_id INT REFERENCES Groups(group_id) ON DELETE CASCADE,
    joined_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- Transactions Table
CREATE TABLE Transactions (
    transaction_id SERIAL PRIMARY KEY,
    lender_id INT REFERENCES Users(user_id) ON DELETE CASCADE,
    borrower_id INT REFERENCES Users(user_id) ON DELETE CASCADE,
    group_id INT REFERENCES Groups(group_id) ON DELETE SET NULL,
    amount DECIMAL(10, 2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    purpose VARCHAR(255)
);


-- Transaction Splits Table
CREATE TABLE Transaction_Splits (
    transaction_split_id SERIAL PRIMARY KEY,
    transaction_id INT REFERENCES Transactions(transaction_id) ON DELETE CASCADE,
    user_id INT REFERENCES Users(user_id) ON DELETE CASCADE,
    amount DECIMAL(10, 2) NOT NULL
);


-- Requests Table
CREATE TABLE Requests (
    request_id SERIAL PRIMARY KEY,
    sender_id INT REFERENCES Users(user_id) ON DELETE CASCADE,
    receiver_id INT REFERENCES Users(user_id) ON DELETE SET NULL,
    group_id INT REFERENCES Groups(group_id) ON DELETE SET NULL,
    amount DECIMAL(10, 2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- Settlements Table
CREATE TABLE Settlements (
    settlement_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES Users(user_id) ON DELETE CASCADE,
    counterparty_id INT REFERENCES Users(user_id) ON DELETE SET NULL,
    group_id INT REFERENCES Groups(group_id) ON DELETE SET NULL,
    amount DECIMAL(10, 2) NOT NULL,
    settlement_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- Balances Table
CREATE TABLE Balances (
    balance_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES Users(user_id) ON DELETE CASCADE,
    group_id INT REFERENCES Groups(group_id) ON DELETE SET NULL,
    owed_amount DECIMAL(10, 2) NOT NULL DEFAULT 0,
    lent_amount DECIMAL(10, 2) NOT NULL DEFAULT 0
);

--optional
CREATE TABLE Owed_Amounts (
    owed_amount_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES Users(user_id) ON DELETE CASCADE,  -- The user who is owed
    debtor_id INT REFERENCES Users(user_id) ON DELETE CASCADE,  -- The user who owes
    amount DECIMAL(10, 2) NOT NULL,
    transaction_id INT REFERENCES Transactions(transaction_id) ON DELETE CASCADE,
    status VARCHAR(50) NOT NULL, -- e.g., pending, settled
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

