SELECT payment_id as my_payment_column FROM payment;

SELECT customer_id, SUM(amount) as total_spent
FROM payment
GROUP BY customer_id;