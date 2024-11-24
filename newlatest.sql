--- Create the 'users' table
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE,
    firebase_id VARCHAR(255) UNIQUE,
    name VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the 'groups' table
CREATE TABLE groups (
    group_id SERIAL PRIMARY KEY,
    group_name VARCHAR(255),
    created_by INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (created_by) REFERENCES users(user_id)
);

-- Create the 'group_members' table
CREATE TABLE group_members (
    group_member_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    group_id INT NOT NULL,
    joined_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (group_id) REFERENCES groups(group_id)
);

-- Create the 'payment_methods' table first so it can be referenced by the 'transactions' table
CREATE TABLE payment_methods (
    payment_id        SERIAL PRIMARY KEY,  -- Automatically increments
    user_id           INTEGER NOT NULL,
    payment_type      VARCHAR(50) NOT NULL CHECK (payment_type IN ('UPI', 'Bank Account')),
    upi_id            VARCHAR(100),  -- For UPI payment methods
    account_number    VARCHAR(20),   -- For Bank Account payment methods
    ifsc_code         VARCHAR(15),   -- For Bank Account IFSC Code
    wallet_provider   VARCHAR(50),   -- Optional, for UPI wallet provider name
    is_primary        BOOLEAN DEFAULT FALSE,  -- Whether it's the user's primary payment method
    created_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Timestamp when the method was created
);

-- Foreign key constraint linking to the users table for payment_methods
ALTER TABLE payment_methods
    ADD CONSTRAINT payment_methods_user_id_fkey FOREIGN KEY (user_id)
    REFERENCES users(user_id) ON DELETE CASCADE;

-- Index for user_id for fast lookups in payment_methods
CREATE INDEX idx_payment_methods_user_id ON payment_methods (user_id);

-- Create the 'transactions' table that references the payment_methods table
CREATE TABLE transactions (
    transaction_id SERIAL PRIMARY KEY,
    lender_id INT NOT NULL,
    borrower_id INT NOT NULL,
    group_id INT,
    amount DECIMAL(10, 2) NOT NULL,
    status VARCHAR(50),
    purpose VARCHAR(255),
    payment_method_id INT, -- Linking to payment method
    retry_count INT DEFAULT 0,  -- Keeps track of retry attempts
    failure_reason TEXT,       -- Describes the failure (e.g., insufficient funds, network error, etc.)
    FOREIGN KEY (lender_id) REFERENCES users(user_id),
    FOREIGN KEY (borrower_id) REFERENCES users(user_id),
    FOREIGN KEY (group_id) REFERENCES groups(group_id),
    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(payment_id), -- Foreign Key for payment method
    CONSTRAINT check_transaction_status CHECK (status IN ('pending', 'successful', 'failed', 'retrying'))
);

-- Create the 'transaction_splits' table
CREATE TABLE transaction_splits (
    transaction_split_id SERIAL PRIMARY KEY,
    transaction_id INT NOT NULL,
    user_id INT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (transaction_id) REFERENCES transactions(transaction_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

-- Create the 'balances' table
CREATE TABLE balances (
    balance_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    group_id INT,
    owed_amount DECIMAL(10, 2) DEFAULT 0,
    lent_amount DECIMAL(10, 2) DEFAULT 0,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (group_id) REFERENCES groups(group_id)
);

-- Create the 'requests' table
CREATE TABLE requests (
    request_id SERIAL PRIMARY KEY,
    sender_id INT NOT NULL,
    receiver_id INT NOT NULL,
    group_id INT,
    amount DECIMAL(10, 2) NOT NULL,
    status VARCHAR(50) DEFAULT 'pending', -- Status can be 'pending', 'accepted', 'rejected'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (sender_id) REFERENCES users(user_id),
    FOREIGN KEY (receiver_id) REFERENCES users(user_id),
    FOREIGN KEY (group_id) REFERENCES groups(group_id)
);

-- Create the 'settlements' table
CREATE TABLE settlements (
    settlement_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    counterparty_id INT NOT NULL,
    group_id INT,
    amount DECIMAL(10, 2) NOT NULL,
    settlement_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (counterparty_id) REFERENCES users(user_id),
    FOREIGN KEY (group_id) REFERENCES groups(group_id)
);

-- Create indexes for the 'transactions' table for performance
CREATE INDEX idx_transactions_lender_id ON transactions(lender_id);
CREATE INDEX idx_transactions_borrower_id ON transactions(borrower_id);
CREATE INDEX idx_transactions_group_id ON transactions(group_id);
CREATE INDEX idx_transactions_payment_method_id ON transactions(payment_method_id);

-- Create indexes for the 'transaction_splits' table for performance
CREATE INDEX idx_transaction_splits_transaction_id ON transaction_splits(transaction_id);
CREATE INDEX idx_transaction_splits_user_id ON transaction_splits(user_id);

-- Create indexes for the 'balances' table for performance
CREATE INDEX idx_balances_user_id ON balances(user_id);
CREATE INDEX idx_balances_group_id ON balances(group_id);

-- newtable idea for logging transaction actions
CREATE TABLE transaction_logs (
    log_id SERIAL PRIMARY KEY,
    transaction_id INT NOT NULL,                    -- Links to the transaction
    action VARCHAR(100) NOT NULL,                    -- Action taken (e.g., 'Payment Initiated', 'Payment Successful', 'Payment Failed')
    status VARCHAR(50),                              -- Status of the action (e.g., 'Pending', 'Success', 'Failed')
    details TEXT,                                    -- Additional details or error messages, if any
    error_details TEXT,                              -- Specific error details (for failed actions)
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,    -- Timestamp of the action
    FOREIGN KEY (transaction_id) REFERENCES transactions(transaction_id) ON DELETE CASCADE
);

-- Index for fast lookups
CREATE INDEX idx_transaction_logs_transaction_id ON transaction_logs (transaction_id);

-- Create the 'payment_notifications' table for user notifications
CREATE TABLE payment_notifications (
    notification_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,                             -- The user to whom the notification is sent
    transaction_id INT NOT NULL,                      -- Links to the transaction
    notification_type VARCHAR(50) NOT NULL,           -- Type of notification (e.g., 'Success', 'Failure', 'Pending')
    message TEXT NOT NULL,                            -- The message content of the notification
    is_read BOOLEAN DEFAULT FALSE,                    -- Whether the user has read the notification
    read_at TIMESTAMP,                                -- Timestamp of when the notification was read
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,   -- Timestamp when the notification was created
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (transaction_id) REFERENCES transactions(transaction_id) ON DELETE CASCADE
);

-- Indexes for fast lookups
CREATE INDEX idx_payment_notifications_user_id ON payment_notifications (user_id);
CREATE INDEX idx_payment_notifications_transaction_id ON payment_notifications (transaction_id);


