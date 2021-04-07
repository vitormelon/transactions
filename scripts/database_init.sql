CREATE TABLE `accounts` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `document_number` longtext,
    PRIMARY KEY (`id`)
);

CREATE TABLE `operation_types`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `description` longtext,
    PRIMARY KEY (`id`)
);

CREATE TABLE `transactions` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `account_id` bigint(20) unsigned DEFAULT NULL,
    `operation_type_id` bigint(20) unsigned DEFAULT NULL,
    `amount` float DEFAULT NULL,
    `created_at` datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `fk_transactions_operation_type` (`operation_type_id`),
    KEY `fk_transactions_account` (`account_id`),
    CONSTRAINT `fk_transactions_account` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`),
    CONSTRAINT `fk_transactions_operation_type` FOREIGN KEY (`operation_type_id`) REFERENCES `operation_types` (`id`)
);

INSERT
    INTO
        `operation_types` (`description`)
    VALUES
        ('COMPRA A VISTA'),
        ('COMPRA PARCELADA'),
        ('SAQUE'),
        ('PAGAMENTO');