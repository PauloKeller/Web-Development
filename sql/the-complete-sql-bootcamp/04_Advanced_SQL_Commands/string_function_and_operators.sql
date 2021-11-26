SELECT first_name || ' ' || last_name AS full_name
FROM customer;

SELECT first_name, char_length(first_name)
FROM customer;

SELECT upper(first_name)
FROM customer;