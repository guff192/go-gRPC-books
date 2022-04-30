INSERT INTO authors (first_name, last_name) VALUES ("Alexandr", "Pushkin");

INSERT INTO books (title, summary) VALUES ("Ruslan & Lyudmila", "Famous russian fairytale");

INSERT INTO books_authors (book_id, author_id) VALUES (
	(SELECT id FROM books WHERE title="Ruslan & Lyudmila"),
	(SELECT id FROM authors WHERE first_name="Alexandr" AND last_name="Pushkin")
);

INSERT INTO books (title, summary) VALUES ("Captain daughter", "Famous russian roman");

INSERT INTO books_authors (book_id, author_id) VALUES (
	(SELECT id FROM books WHERE title="Captain daughter"),
	(SELECT id FROM authors WHERE first_name="Alexandr" AND last_name="Pushkin")
);


INSERT INTO books (title, summary) VALUES
(
	"The Twelve Chairs",
	"Classic satirical novel by the Odessan Soviet authors Ilf and Petrov, published in 1928."
);

INSERT INTO authors (first_name, last_name) VALUES ("Ilya", "Ilf");

INSERT INTO authors (first_name, last_name) VALUES ("Yevgeny", "Petrov");

INSERT INTO books_authors (book_id, author_id) VALUES (
	(SELECT id FROM books WHERE title="The Twelve Chairs"),
	(SELECT id FROM authors WHERE first_name="Ilya" AND last_name="Ilf")
);

INSERT INTO books_authors (book_id, author_id) VALUES (
	(SELECT id FROM books WHERE title="The Twelve Chairs"),
	(SELECT id FROM authors WHERE first_name="Yevgeny" AND last_name="Petrov")
);
