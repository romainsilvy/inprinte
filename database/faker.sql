-- CREATE ROLES
INSERT INTO role (role) 
VALUES ("admin");
INSERT INTO role (role) 
VALUES ("member");


-- CREATE GLOBAL CATEGORY
INSERT INTO global_category (global_category_name) VALUES ("informatique");
INSERT INTO global_category (global_category_name) VALUES ("jouets");
INSERT INTO global_category (global_category_name) VALUES ("nature");


-- CREATE CATEGORY
INSERT INTO category (id_global_category,category_name)VALUES ((SELECT id_global_category FROM global_category WHERE global_category_name = "informatique"),"accessoires de téléphone");
INSERT INTO category (id_global_category,category_name)VALUES ((SELECT id_global_category FROM global_category WHERE global_category_name = "informatique"),"accessoires de pc");

INSERT INTO category (id_global_category,category_name)VALUES ((SELECT id_global_category FROM global_category WHERE global_category_name = "jouets"),"voiture");
INSERT INTO category (id_global_category,category_name)VALUES ((SELECT id_global_category FROM global_category WHERE global_category_name = "jouets"),"bateau");

INSERT INTO category (id_global_category,category_name)VALUES ((SELECT id_global_category FROM global_category WHERE global_category_name = "nature"),"feuille d'érable");
INSERT INTO category (id_global_category,category_name)VALUES ((SELECT id_global_category FROM global_category WHERE global_category_name = "nature"),"merde de vegan");


-- CREATE AN USER AND HIS CART 
INSERT INTO user (first_name,last_name,email,password,phone,id_role,is_alive,street,country,state,zipCode,city)
VALUES("louis", "reynard", "louis@gmail.com", "louisensueur", "0673877878", (SELECT id_role FROM role WHERE role = "admin"), true, "rue de zeub","france","rhone alpes","69420","ampuis");
INSERT INTO `command` (id_user, date_command,command_state)
VALUES(LAST_INSERT_ID(), null,"cart");

INSERT INTO user (first_name,last_name,email,password,phone,id_role,is_alive,street,country,state,zipCode,city)
VALUES("isma", "touloulou", "isma@gmail.com", "goat", "0784848484", (SELECT id_role FROM role WHERE role = "member"), true, "rue de zeub","france","rhone alpes","69420","mes bourses");
INSERT INTO `command` (id_user, date_command,command_state)
VALUES(LAST_INSERT_ID(), null,"cart");

INSERT INTO user (first_name,last_name,email,password,phone,id_role,is_alive,street,country,state,zipCode,city)
VALUES("matt", "foot", "glandux@gmail.com", "zob", "0735353535", (SELECT id_role FROM role WHERE role = "member"), true, "rue de rsq","france","rhone alpes","69420","mes boursettes");
INSERT INTO `command` (id_user, date_command,command_state)
VALUES(LAST_INSERT_ID(), null,"cart");


-- CREATE A PRODUCT 
INSERT INTO product (name,description,price,id_category,pending_validation,id_user)VALUES ("porsche","neuve comme les boursettes d'isma, aucune usure, légèrement bleutée",10, (SELECT id_category FROM category WHERE category_name = "voiture"),false, 1);
INSERT INTO product_file (id_product,url) VALUES (LAST_INSERT_ID(),"product_url du premier produit");
INSERT INTO picture (url) VALUES ("picture_url du premier produit");
INSERT INTO product_picture (id_product, id_picture) VALUES((SELECT MAX(id_product) FROM product),LAST_INSERT_ID());

INSERT INTO product (name,description,price,id_category,pending_validation,id_user) VALUES ("pomme","juteuse comme ton fiak",10, 2,false, 2);
INSERT INTO product_file (id_product,url) VALUES (LAST_INSERT_ID(),"product_url du second produit");
INSERT INTO picture (url) VALUES ("picture_url du second produit");
INSERT INTO product_picture (id_product, id_picture) VALUES((SELECT MAX(id_product) FROM product),LAST_INSERT_ID());

INSERT INTO product (name,description,price,id_category,pending_validation,id_user) VALUES ("oignon","rond comme le tien",10, 3,false, 3);
INSERT INTO product_file (id_product,url) VALUES (LAST_INSERT_ID(),"product_url du troisième produit");
INSERT INTO picture (url) VALUES ("picture_url du troisième produit");
INSERT INTO product_picture (id_product, id_picture) VALUES((SELECT MAX(id_product) FROM product),LAST_INSERT_ID());


-- RATE A PRODUCT
INSERT INTO rate (id_product,id_user,stars_number)VALUES (1,1,1);
INSERT INTO rate (id_product,id_user,stars_number)VALUES (1,1,5);
INSERT INTO rate (id_product,id_user,stars_number)VALUES (2,2,1);
INSERT INTO rate (id_product,id_user,stars_number)VALUES (2,2,5);
INSERT INTO rate (id_product,id_user,stars_number)VALUES (3,3,1);
INSERT INTO rate (id_product,id_user,stars_number)VALUES (3,2,5);


-- ADD TO FAVORITE
INSERT INTO favorite (id_product,id_user) VALUES (1,1);
INSERT INTO favorite (id_product,id_user) VALUES (2,2);
INSERT INTO favorite (id_product,id_user) VALUES (3,3);
INSERT INTO favorite (id_product,id_user) VALUES (1,2);
INSERT INTO favorite (id_product,id_user) VALUES (2,3);
INSERT INTO favorite (id_product,id_user) VALUES (3,1);


-- ADD PRODUCT TO CART    
INSERT INTO command_line (id_product, id_command, quantity, command_line_state) VALUES (1,(SELECT id_command FROM command WHERE id_user = 1 AND command_state = "cart"), 3, "pending_printing");
INSERT INTO command_line (id_product, id_command, quantity, command_line_state) VALUES (2,(SELECT id_command FROM command WHERE id_user = 2 AND command_state = "cart"), 3, "pending_printing");
INSERT INTO command_line (id_product, id_command, quantity, command_line_state) VALUES (3,(SELECT id_command FROM command WHERE id_user = 3 AND command_state = "cart"), 3, "pending_printing");
INSERT INTO command_line (id_product, id_command, quantity, command_line_state) VALUES (3,(SELECT id_command FROM command WHERE id_user = 3 AND command_state = "cart"), 3, "pending_printing");


