CREATE TABLE `users`
(
    id bigint auto_increment,
    PRIMARY KEY (`id`),
    email varchar(255) NOT NULL,
    UNIQUE (`email`)
    password varchar(255) NOT NULL
);

INSERT INTO `users` (email, password) VALUES ('nana.elvee@gmail.com', 'password'), 