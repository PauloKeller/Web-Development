SELECT customer_id + rental_id AS new_id
FROM payment;

SELECT round(AVG(amount), 2)
FROM payment;