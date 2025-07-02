CREATE UNIQUE INDEX IF NOT EXISTS unique_dish_name_restaurant
ON dishes(name, restaurant_id);
