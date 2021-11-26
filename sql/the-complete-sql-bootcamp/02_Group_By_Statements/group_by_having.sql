SELECT customer_id, COUNT(amount) 
FROM payment p2
GROUP BY customer_id
HAVING COUNT(amount) > 40;

SELECT rating, AVG(rental_duration)
FROM film
GROUP BY rating
HAVING AVG(rental_duration) > 5;