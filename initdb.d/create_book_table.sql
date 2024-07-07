 CREATE TABLE IF NOT EXISTS `DB`.`Book` (
   `id` INT NOT NULL COMMENT 'Book ID',
   `title` VARCHAR(255) NOT NULL COMMENT 'Book title',
   `author` VARCHAR(255) NOT NULL COMMENT 'Book author',
   `rating` TINYINT COMMENT '10 point rating',
   `review` TEXT COMMENT 'review content',
   PRIMARY KEY (`id`))
 ENGINE = InnoDB;