
# Key Entities and Relationships

## Users Table (users)
- Central entity, representing the users.
- Relationships:
  - Links to balances, group_members, groups, payment_methods, payment_notifications, requests, settlements, transaction_splits, and transactions.

## Groups Table (groups)
- Represents groups that users can belong to.
- Relationships:
  - Links to group_members and balances to manage group members and group-specific balances.
  - Linked to transactions for group-based transactions.

## Group Members Table (group_members)
- Represents the membership of users in groups.
- Foreign keys:
  - user_id → users(user_id)
## Balances Table (balances)
- Tracks financial balances per user, possibly by group.
- Foreign keys:
  - user_id → users(user_id)
  - group_id → groups(group_id)

## Transactions Table (transactions)
- Central for managing financial transactions.
- Fields like lender_id, borrower_id, group_id, and payment_method_id link transactions to various entities.
- Relationships:
  - Links to users for lenders and borrowers.
  - Links to groups for group-specific transactions.
  - Links to payment_methods.

## Transaction Splits Table (transaction_splits)
Transaction Splits Table (transaction_splits):
- Tracks how a transaction is split among users.
- Foreign keys:
  - transaction_id → transactions(transaction_id)
  - user_id → users(user_id)

## Payment Methods Table (payment_methods)
- Represents user payment methods.
- Foreign key:
  - user_id → users(user_id)

## Payment Notifications Table (payment_notifications)
- Tracks notifications related to transactions.
- Foreign keys:
  - transaction_id → transactions(transaction_id)
  - user_id → users(user_id)

## Transaction Logs Table (transaction_logs)
- Tracks logs or changes related to transactions.
- Foreign key:
  - transaction_id → transactions(transaction_id)

## Requests Table (requests)
- Likely for payment or group requests.
- Foreign keys:
  - sender_id → users(user_id)
  - receiver_id → users(user_id)

## Settlements Table (settlements)
- Tracks settlements between users.
- Foreign keys:
  - user_id → users(user_id)
  - counterparty_id → users(user_id)
  - group_id → groups(group_id)
  - user_id and counterparty_id → users(user_id)
  - group_id → groups(group_id)
