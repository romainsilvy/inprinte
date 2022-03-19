-- getProduct : 
SELECT name, description, price, AVG(stars_number) AS MOYENNE, picture.url, product_file.id_product 
FROM product 
INNER JOIN rate ON rate.id_product = product.id_product 
INNER JOIN product_picture ON product_picture.id_picture = product.id_product 
INNER JOIN picture ON picture.id_picture = product_picture.id_picture 
INNER JOIN product_file ON product_file.id_product = product.id_product
WHERE product.id_product = 1
GROUP BY product.id_product;

SELECT name, description, price, AVG(stars_number) AS MOYENNE, picture.url, product_file.id_product FROM product INNER JOIN rate ON rate.id_product = product.id_product INNER JOIN product_picture ON product_picture.id_picture = product.id_product INNER JOIN picture ON picture.id_picture = product_picture.id_picture INNER JOIN product_file ON product_file.id_product = product.id_productWHERE product.id_product = 1GROUP BY product.id_product;
