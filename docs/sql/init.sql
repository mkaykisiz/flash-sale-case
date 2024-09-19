-- init.sql

-- Create table for "auth"
CREATE TABLE users (
                      id SERIAL PRIMARY KEY,
                      username VARCHAR(50) NOT NULL,
                      password VARCHAR(50) NOT NULL
);

-- Insert sample data into "auth"
INSERT INTO users (username, password) VALUES
                                          ('admin', 'admin'),
                                          ('user1', 'password456');

-- Create table for "products"
CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          created_on INT NOT NULL,
                          modified_on INT NOT NULL,
                          deleted_on INT,
                          name VARCHAR(255) NOT NULL,
                          stock INT NOT NULL,
                          price FLOAT NOT NULL
);

-- Insert sample data into "products"
INSERT INTO products (created_on, modified_on, deleted_on, name, stock, price) VALUES
                                                                                       (EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW()), 0, 'Product 1', 100, 450.75),
                                                                                       (EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW()), 0, 'Product 2', 50, 450.75),
                                                                                       (EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW()), 0, 'Product 3', 0, 450.75);
-- Create table for "flash_sales"
CREATE TABLE flash_sales (
                             id SERIAL PRIMARY KEY,
                             created_on INT NOT NULL,
                             modified_on INT NOT NULL,
                             deleted_on INT,
                             product_id INTEGER NOT NULL,
                             discount_percent INTEGER NOT NULL CHECK (discount_percent >= 0 AND discount_percent <= 100),
                             stock INTEGER NOT NULL CHECK (stock >= 0),
                             start_time TIMESTAMP NOT NULL,
                             end_time TIMESTAMP NOT NULL,

                             CONSTRAINT fk_product
                                 FOREIGN KEY (product_id)
                                     REFERENCES products (id)
                                     ON UPDATE CASCADE
                                     ON DELETE SET NULL
);
-- Insert sample data into "flash_sales"
INSERT INTO flash_sales (created_on, modified_on, deleted_on, product_id, discount_percent, stock, start_time, end_time)
VALUES
    (EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW()), 0, 1, 20, 50, '2024-09-19 10:00:00', '2024-11-20 18:00:00'),
    (EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW()), 0, 2, 15, 30, '2024-09-19 09:00:00', '2024-11-21 17:00:00'),
    (EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW()), 0, 3, 50, 100, '2024-09-19 08:00:00', '2024-11-22 20:00:00');

-- Create table for "order"
CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        created_on INT,
                        modified_on INT,
                        deleted_on INT,
                        user_id INT REFERENCES users(id), -- Foreign key
                        flash_sale_id INT REFERENCES flash_sales(id), -- Foreign key
                        net_price FLOAT,
                        total_price FLOAT
);


-- Create table for "order_items"
CREATE TABLE order_items (
                             id SERIAL PRIMARY KEY,
                             created_on INT,
                             modified_on INT,
                             deleted_on INT,
                             order_id INT REFERENCES orders(id), -- Foreign key
                             product_id INT REFERENCES products(id), -- Foreign key
                             quantity INT,
                             price FLOAT,
                             discounted_price FLOAT
);

