SELECT staff_id, COUNT(amount), SUM(amount) 
FROM payment
GROUP BY staff_id;

SELECT rating, ROUND(AVG(replacement_cost),2)
FROM film
GROUP BY rating;

SELECT customer_id, SUM(amount)  
FROM payment
GROUP BY customer_id
ORDER BY SUM(amount) DESC
LIMIT 5;