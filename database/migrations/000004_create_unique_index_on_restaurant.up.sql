CREATE UNIQUE INDEX IF NOT EXISTS unique_restaurant_name_address
ON restaurants(name, address)