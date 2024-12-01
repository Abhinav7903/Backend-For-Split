\# Database Schema Documentation

This document outlines the structure of the database tables for the financial transaction and group management system.

\#\# Tables

\#\#\# 1\. \*\*Users Table\*\*  
Stores user details.

| Column       | Type                     | Description                            |  
|--------------|--------------------------|----------------------------------------|  
| \`user\_id\`    | \`integer\`                | Primary key for the user              |  
| \`name\`       | \`character varying(100)\`  | Name of the user                      |  
| \`email\`      | \`character varying(100)\`  | Unique email address of the user      |  
| \`firebase\_id\`| \`character varying(255)\`  | Unique Firebase ID of the user        |

Indexes:  
\- Unique index on \`email\` and \`firebase\_id\`.

\---

\#\#\# 2\. \*\*Groups Table\*\*  
Stores group information.

| Column       | Type                     | Description                            |  
|--------------|--------------------------|----------------------------------------|  
| \`group\_id\`   | \`integer\`                | Primary key for the group             |  
| \`group\_name\` | \`character varying(100)\`  | Name of the group                     |  
| \`created\_by\` | \`integer\`                | Foreign key referencing \`users(user\_id)\` |

Indexes:  
\- Primary key on \`group\_id\`.

Foreign Key:  
\- \`created\_by\` references \`users(user\_id)\`.

\---

\#\#\# 3\. \*\*Group Members Table\*\*  
Manages relationships between users and groups.

| Column          | Type                     | Description                            |  
|-----------------|--------------------------|----------------------------------------|  
| \`group\_member\_id\` | \`integer\`              | Primary key for the group member      |  
| \`user\_id\`       | \`integer\`                | Foreign key referencing \`users(user\_id)\` |  
| \`group\_id\`      | \`integer\`                | Foreign key referencing \`groups(group\_id)\` |  
| \`joined\_date\`   | \`timestamp\`              | Date when the user joined the group   |

Indexes:  
\- Primary key on \`group\_member\_id\`.  
\- Indexes on \`group\_id\` and \`user\_id\`.

Foreign Keys:  
\- \`user\_id\` references \`users(user\_id)\`.  
\- \`group\_id\` references \`groups(group\_id)\`.

\---

\#\#\# 4\. \*\*Transactions Table\*\*  
Manages lending and borrowing activities.

| Column         | Type                     | Description                            |  
|----------------|--------------------------|----------------------------------------|  
| \`transaction\_id\`| \`integer\`               | Primary key for the transaction       |  
| \`lender\_id\`    | \`integer\`                | Foreign key referencing \`users(user\_id)\` (lender) |  
| \`borrower\_id\`  | \`integer\`                | Foreign key referencing \`users(user\_id)\` (borrower) |  
| \`group\_id\`     | \`integer\`                | Foreign key referencing \`groups(group\_id)\` (optional) |  
| \`amount\`       | \`numeric(10,2)\`          | Amount involved in the transaction    |  
| \`status\`       | \`character varying(50)\`  | Status of the transaction (e.g., pending, completed) |  
| \`timestamp\`    | \`timestamp\`              | Timestamp of the transaction          |  
| \`purpose\`      | \`character varying(255)\` | Purpose of the transaction            |  
| \`payment\_id\`   | \`integer\`                | Foreign key referencing \`payment\_methods(payment\_id)\` |

Indexes:  
\- Primary key on \`transaction\_id\`.  
\- Indexes on \`lender\_id\` and \`borrower\_id\`.

Foreign Keys:  
\- \`lender\_id\` references \`users(user\_id)\`.  
\- \`borrower\_id\` references \`users(user\_id)\`.  
\- \`group\_id\` references \`groups(group\_id)\`.  
\- \`payment\_id\` references \`payment\_methods(payment\_id)\`.

\---

\#\#\# 5\. \*\*Transaction Splits Table\*\*  
Breaks down transactions into amounts owed by individual participants.

| Column                | Type                     | Description                            |  
|-----------------------|--------------------------|----------------------------------------|  
| \`transaction\_split\_id\` | \`integer\`                | Primary key for the transaction split |  
| \`transaction\_id\`       | \`integer\`                | Foreign key referencing \`transactions(transaction\_id)\` |  
| \`user\_id\`              | \`integer\`                | Foreign key referencing \`users(user\_id)\` |  
| \`amount\`               | \`numeric(10,2)\`          | Amount the user owes                   |

Indexes:  
\- Primary key on \`transaction\_split\_id\`.  
\- Indexes on \`transaction\_id\` and \`user\_id\`.

Foreign Keys:  
\- \`transaction\_id\` references \`transactions(transaction\_id)\`.  
\- \`user\_id\` references \`users(user\_id)\`.

\---

\#\#\# 6\. \*\*Balances Table\*\*  
Tracks amounts owed and lent for each user in solo or group transactions.

| Column       | Type                     | Description                            |  
|--------------|--------------------------|----------------------------------------|  
| \`balance\_id\` | \`integer\`                | Primary key for the balance record    |  
| \`user\_id\`    | \`integer\`                | Foreign key referencing \`users(user\_id)\` |  
| \`group\_id\`   | \`integer\`                | Foreign key referencing \`groups(group\_id)\` (optional) |  
| \`owed\_amount\`| \`numeric(10,2)\`          | Amount the user owes                   |  
| \`lent\_amount\`| \`numeric(10,2)\`          | Amount the user has lent               |

Indexes:  
\- Primary key on \`balance\_id\`.  
\- Indexes on \`user\_id\` and \`group\_id\`.

Foreign Keys:  
\- \`user\_id\` references \`users(user\_id)\`.  
\- \`group\_id\` references \`groups(group\_id)\`.

\---

\#\#\# 7\. \*\*Requests Table\*\*  
Manages money requests between users.

| Column        | Type                     | Description                            |  
|---------------|--------------------------|----------------------------------------|  
| \`request\_id\`  | \`integer\`                | Primary key for the request           |  
| \`sender\_id\`   | \`integer\`                | Foreign key referencing \`users(user\_id)\` (sender) |  
| \`receiver\_id\` | \`integer\`                | Foreign key referencing \`users(user\_id)\` (receiver) |  
| \`group\_id\`    | \`integer\`                | Foreign key referencing \`groups(group\_id)\` (optional) |  
| \`amount\`      | \`numeric(10,2)\`          | Amount requested                      |  
| \`status\`      | \`character varying(50)\`  | Status of the request (e.g., pending, accepted) |  
| \`timestamp\`   | \`timestamp\`              | Timestamp of the request              |

Indexes:  
\- Primary key on \`request\_id\`.

Foreign Keys:  
\- \`sender\_id\` references \`users(user\_id)\`.  
\- \`receiver\_id\` references \`users(user\_id)\`.  
\- \`group\_id\` references \`groups(group\_id)\`.

\---

\#\#\# 8\. \*\*Settlements Table\*\*  
Tracks settlements of debts between users.

| Column         | Type                     | Description                            |  
|----------------|--------------------------|----------------------------------------|  
| \`settlement\_id\`| \`integer\`                | Primary key for the settlement        |  
| \`user\_id\`      | \`integer\`                | Foreign key referencing \`users(user\_id)\` (settler) |  
| \`counterparty\_id\`| \`integer\`              | Foreign key referencing \`users(user\_id)\` (counterparty) |  
| \`group\_id\`     | \`integer\`                | Foreign key referencing \`groups(group\_id)\` (optional) |  
| \`amount\`       | \`numeric(10,2)\`          | Amount settled                        |  
| \`settlement\_date\`| \`timestamp\`            | Date of settlement                    |  
| \`payment\_id\`   | \`integer\`                | Foreign key referencing \`payment\_methods(payment\_id)\` |

Indexes:  
\- Primary key on \`settlement\_id\`.

Foreign Keys:  
\- \`user\_id\` references \`users(user\_id)\`.  
\- \`counterparty\_id\` references \`users(user\_id)\`.  
\- \`group\_id\` references \`groups(group\_id)\`.  
\- \`payment\_id\` references \`payment\_methods(payment\_id)\`.

\---

\#\#\# 9\. \*\*Payment Methods Table\*\*  
Stores user payment methods.

| Column         | Type                     | Description                            |  
|----------------|--------------------------|----------------------------------------|  
| \`payment\_id\`   | \`integer\`                | Primary key for the payment method    |  
| \`user\_id\`      | \`integer\`                | Foreign key referencing \`users(user\_id)\` |  
| \`payment\_type\` | \`character varying(50)\`   | Type of payment method (e.g., UPI, bank account) |  
| \`upi\_id\`       | \`bytea\`                  | UPI ID (optional)                     |  
| \`account\_number\`| \`bytea\`                 | Account number (optional)             |  
| \`ifsc\_code\`    | \`character varying(15)\`   | IFSC code (optional)                  |  
| \`wallet\_provider\`| \`character varying(50)\` | Wallet provider name (optional)       |  
| \`is\_primary\`   | \`boolean\`                | Indicates if this is the primary payment method |  
| \`created\_at\`   | \`timestamp\`              | Timestamp when the payment method was added |

Indexes:  
\- Primary key on \`payment\_id\`.

Foreign Keys:  
\- \`user\_id\` references \`users(user\_id)\`.

\---

\#\# Relationships

1\. \*\*Users\*\* can have multiple \*\*Payment Methods\*\*, \*\*Requests\*\*, \*\*Settlements\*\*, and \*\*Transactions\*\*.  
2\. \*\*Groups\*\* can have multiple \*\*Group Members\*\*, \*\*Transactions\*\*, \*\*Requests\*\*, and \*\*Settlements\*\*.  
3\. \*\*Transactions\*\* can be linked to multiple \*\*Transaction Splits\*\*.  
4\. \*\*Payments\*\* are linked to \*\*Settlements\*\* and \*\*Transactions\*\* to track payment details.

\---

This database schema supports the management of users, their financial transactions (both individual and group-based), and the various payment methods they can use. The relationships between these tables ensure efficient tracking and management of debts, payments, and settlements.

