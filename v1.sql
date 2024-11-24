CREATE TABLE "users"(
    "user_id" INTEGER NOT NULL,
    "email" VARCHAR(255) NULL,
    "firebase_id" VARCHAR(255) NULL,
    "name" VARCHAR(100) NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE
    "users" ADD PRIMARY KEY("user_id");
ALTER TABLE
    "users" ADD CONSTRAINT "users_email_unique" UNIQUE("email");
ALTER TABLE
    "users" ADD CONSTRAINT "users_firebase_id_unique" UNIQUE("firebase_id");
CREATE TABLE "groups"(
    "group_id" INTEGER NOT NULL,
    "group_name" VARCHAR(255) NULL,
    "created_by" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE
    "groups" ADD PRIMARY KEY("group_id");
CREATE TABLE "group_members"(
    "group_member_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "group_id" INTEGER NOT NULL,
    "joined_date" TIMESTAMP(0) WITHOUT TIME ZONE NULL DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE
    "group_members" ADD PRIMARY KEY("group_member_id");
CREATE TABLE "payment_methods"(
    "payment_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "payment_type" VARCHAR(50) NOT NULL,
    "upi_id" VARCHAR(100) NULL,
    "account_number" VARCHAR(20) NULL,
    "ifsc_code" VARCHAR(15) NULL,
    "wallet_provider" VARCHAR(50) NULL,
    "is_primary" BOOLEAN NULL DEFAULT 'DEFAULT FALSE',
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE
    "payment_methods" ADD PRIMARY KEY("payment_id");
CREATE INDEX "payment_methods_user_id_index" ON
    "payment_methods"("user_id");
CREATE TABLE "transactions"(
    "transaction_id" INTEGER NOT NULL,
    "lender_id" INTEGER NOT NULL,
    "borrower_id" INTEGER NOT NULL,
    "group_id" INTEGER NULL,
    "amount" DECIMAL(10, 2) NOT NULL,
    "status" VARCHAR(50) NULL,
    "purpose" VARCHAR(255) NULL,
    "payment_method_id" INTEGER NULL,
    "retry_count" INTEGER NULL,
    "failure_reason" TEXT NULL
);
ALTER TABLE
    "transactions" ADD PRIMARY KEY("transaction_id");
CREATE INDEX "transactions_lender_id_index" ON
    "transactions"("lender_id");
CREATE INDEX "transactions_borrower_id_index" ON
    "transactions"("borrower_id");
CREATE INDEX "transactions_group_id_index" ON
    "transactions"("group_id");
CREATE INDEX "transactions_payment_method_id_index" ON
    "transactions"("payment_method_id");
CREATE TABLE "transaction_splits"(
    "transaction_split_id" INTEGER NOT NULL,
    "transaction_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "amount" DECIMAL(10, 2) NOT NULL
);
ALTER TABLE
    "transaction_splits" ADD PRIMARY KEY("transaction_split_id");
CREATE INDEX "transaction_splits_transaction_id_index" ON
    "transaction_splits"("transaction_id");
CREATE INDEX "transaction_splits_user_id_index" ON
    "transaction_splits"("user_id");
CREATE TABLE "balances"(
    "balance_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "group_id" INTEGER NULL,
    "owed_amount" DECIMAL(10, 2) NULL,
    "lent_amount" DECIMAL(10, 2) NULL
);
ALTER TABLE
    "balances" ADD PRIMARY KEY("balance_id");
CREATE INDEX "balances_user_id_index" ON
    "balances"("user_id");
CREATE INDEX "balances_group_id_index" ON
    "balances"("group_id");
CREATE TABLE "requests"(
    "request_id" INTEGER NOT NULL,
    "sender_id" INTEGER NOT NULL,
    "receiver_id" INTEGER NOT NULL,
    "group_id" INTEGER NULL,
    "amount" DECIMAL(10, 2) NOT NULL,
    "status" VARCHAR(50) NULL DEFAULT 'pending',
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE
    "requests" ADD PRIMARY KEY("request_id");
CREATE TABLE "settlements"(
    "settlement_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "counterparty_id" INTEGER NOT NULL,
    "group_id" INTEGER NULL,
    "amount" DECIMAL(10, 2) NOT NULL,
    "settlement_date" TIMESTAMP(0) WITHOUT TIME ZONE NULL DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE
    "settlements" ADD PRIMARY KEY("settlement_id");
CREATE TABLE "transaction_logs"(
    "log_id" INTEGER NOT NULL,
    "transaction_id" INTEGER NOT NULL,
    "action" VARCHAR(100) NOT NULL,
    "status" VARCHAR(50) NULL,
    "details" TEXT NULL,
    "error_details" TEXT NULL,
    "timestamp" TIMESTAMP(0) WITHOUT TIME ZONE NULL DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE
    "transaction_logs" ADD PRIMARY KEY("log_id");
CREATE INDEX "transaction_logs_transaction_id_index" ON
    "transaction_logs"("transaction_id");
CREATE TABLE "payment_notifications"(
    "notification_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "transaction_id" INTEGER NOT NULL,
    "notification_type" VARCHAR(50) NOT NULL,
    "message" TEXT NOT NULL,
    "is_read" BOOLEAN NULL DEFAULT 'DEFAULT FALSE',
    "read_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NULL DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE
    "payment_notifications" ADD PRIMARY KEY("notification_id");
CREATE INDEX "payment_notifications_user_id_index" ON
    "payment_notifications"("user_id");
CREATE INDEX "payment_notifications_transaction_id_index" ON
    "payment_notifications"("transaction_id");
ALTER TABLE
    "payment_notifications" ADD CONSTRAINT "payment_notifications_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("user_id");
ALTER TABLE
    "transactions" ADD CONSTRAINT "transactions_group_id_foreign" FOREIGN KEY("group_id") REFERENCES "groups"("group_id");
ALTER TABLE
    "group_members" ADD CONSTRAINT "group_members_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("user_id");
ALTER TABLE
    "transaction_splits" ADD CONSTRAINT "transaction_splits_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("user_id");
ALTER TABLE
    "settlements" ADD CONSTRAINT "settlements_counterparty_id_foreign" FOREIGN KEY("counterparty_id") REFERENCES "users"("user_id");
ALTER TABLE
    "requests" ADD CONSTRAINT "requests_group_id_foreign" FOREIGN KEY("group_id") REFERENCES "groups"("group_id");
ALTER TABLE
    "transaction_splits" ADD CONSTRAINT "transaction_splits_transaction_id_foreign" FOREIGN KEY("transaction_id") REFERENCES "transactions"("transaction_id");
ALTER TABLE
    "payment_notifications" ADD CONSTRAINT "payment_notifications_transaction_id_foreign" FOREIGN KEY("transaction_id") REFERENCES "transactions"("transaction_id");
ALTER TABLE
    "settlements" ADD CONSTRAINT "settlements_group_id_foreign" FOREIGN KEY("group_id") REFERENCES "groups"("group_id");
ALTER TABLE
    "groups" ADD CONSTRAINT "groups_created_by_foreign" FOREIGN KEY("created_by") REFERENCES "users"("user_id");
ALTER TABLE
    "balances" ADD CONSTRAINT "balances_group_id_foreign" FOREIGN KEY("group_id") REFERENCES "groups"("group_id");
ALTER TABLE
    "balances" ADD CONSTRAINT "balances_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("user_id");
ALTER TABLE
    "settlements" ADD CONSTRAINT "settlements_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("user_id");
ALTER TABLE
    "transactions" ADD CONSTRAINT "transactions_lender_id_foreign" FOREIGN KEY("lender_id") REFERENCES "users"("user_id");
ALTER TABLE
    "group_members" ADD CONSTRAINT "group_members_group_id_foreign" FOREIGN KEY("group_id") REFERENCES "groups"("group_id");
ALTER TABLE
    "requests" ADD CONSTRAINT "requests_sender_id_foreign" FOREIGN KEY("sender_id") REFERENCES "users"("user_id");
ALTER TABLE
    "transactions" ADD CONSTRAINT "transactions_borrower_id_foreign" FOREIGN KEY("borrower_id") REFERENCES "users"("user_id");
ALTER TABLE
    "transaction_logs" ADD CONSTRAINT "transaction_logs_transaction_id_foreign" FOREIGN KEY("transaction_id") REFERENCES "transactions"("transaction_id");
ALTER TABLE
    "payment_methods" ADD CONSTRAINT "payment_methods_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("user_id");
ALTER TABLE
    "requests" ADD CONSTRAINT "requests_receiver_id_foreign" FOREIGN KEY("receiver_id") REFERENCES "users"("user_id");
ALTER TABLE
    "transactions" ADD CONSTRAINT "transactions_payment_method_id_foreign" FOREIGN KEY("payment_method_id") REFERENCES "payment_methods"("payment_id");