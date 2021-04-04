CREATE TABLE accounts (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    document_number VARCHAR(50) NOT NULL UNIQUE,
    available_credit_limit INTEGER NOT NULL,
    created_at TIMESTAMP
);

CREATE TABLE operations (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    description VARCHAR(50) NOT NULL,
    type VARCHAR(50) NOT NULL
);

CREATE TABLE transactions (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    account_id BIGINT NOT NULL,
    operation_id INT NOT NULL,
    amount INTEGER NOT NULL,
    balance INTEGER NOT NULL,
    created_at TIMESTAMP,

    FOREIGN KEY (account_id) REFERENCES accounts(id),
    FOREIGN KEY (operation_id) REFERENCES operations(id)
);

INSERT
    INTO
        `operations` (`description`, `type`)
    VALUES
        ('COMPRA A VISTA', 'DEBIT'),
        ('COMPRA PARCELADA', 'DEBIT'),
        ('SAQUE', 'DEBIT'),
        ('PAGAMENTO', 'CREDIT');