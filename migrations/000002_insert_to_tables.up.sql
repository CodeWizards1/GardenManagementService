BEGIN;

-- Insert mock data into gardens table
INSERT INTO gardens (user_id, name, type, area_sqm)
VALUES
(uuid_generate_v4(), 'Sunny Balcony', 'balcony', 12.50),
(uuid_generate_v4(), 'Rooftop Oasis', 'rooftop', 45.00),
(uuid_generate_v4(), 'Indoor Jungle', 'indoor', 30.75),
(uuid_generate_v4(), 'Community Garden A', 'community', 100.00),
(uuid_generate_v4(), 'Backyard Haven', 'backyard', 75.25),
(uuid_generate_v4(), 'Urban Balcony', 'balcony', 10.00),
(uuid_generate_v4(), 'Skyline Rooftop', 'rooftop', 50.00),
(uuid_generate_v4(), 'Living Room Greenery', 'indoor', 25.00),
(uuid_generate_v4(), 'Community Garden B', 'community', 120.00),
(uuid_generate_v4(), 'Suburban Backyard', 'backyard', 80.00),
(uuid_generate_v4(), 'Cozy Balcony', 'balcony', 8.50),
(uuid_generate_v4(), 'Modern Rooftop', 'rooftop', 60.00),
(uuid_generate_v4(), 'Office Indoor Garden', 'indoor', 20.00),
(uuid_generate_v4(), 'Community Garden C', 'community', 110.00),
(uuid_generate_v4(), 'Large Backyard', 'backyard', 90.00);

-- Insert mock data into plants table
INSERT INTO plants (garden_id, species, quantity, planting_date, status)
SELECT id, 'Tomato', 5, CURRENT_DATE - interval '10 days', 'growing' FROM gardens LIMIT 1;
INSERT INTO plants (garden_id, species, quantity, planting_date, status)
SELECT id, 'Basil', 10, CURRENT_DATE - interval '20 days', 'planted' FROM gardens LIMIT 1 OFFSET 1;
INSERT INTO plants (garden_id, species, quantity, planting_date, status)
SELECT id, 'Rosemary', 3, CURRENT_DATE - interval '30 days', 'planned' FROM gardens LIMIT 1 OFFSET 2;
INSERT INTO plants (garden_id, species, quantity, planting_date, status)
SELECT id, 'Lettuce', 15, CURRENT_DATE - interval '40 days', 'harvesting' FROM gardens LIMIT 1 OFFSET 3;
INSERT INTO plants (garden_id, species, quantity, planting_date, status)
SELECT id, 'Carrot', 8, CURRENT_DATE - interval '50 days', 'dormant' FROM gardens LIMIT 1 OFFSET 4;
INSERT INTO plants (garden_id, species, quantity, planting_date, status)
SELECT id, 'Pepper', 6, CURRENT_DATE - interval '60 days', 'growing' FROM gardens LIMIT 1 OFFSET 5;
INSERT INTO plants (garden_id, species, quantity, planting_date, status)
SELECT id, 'Mint', 12, CURRENT_DATE - interval '70 days', 'planted' FROM gardens LIMIT 1 OFFSET 6;
INSERT INTO plants (garden_id, species, quantity, planting_date, status)
SELECT id, 'Cucumber', 4, CURRENT_DATE - interval '80 days', 'planned' FROM gardens LIMIT 1 OFFSET 7;
INSERT INTO plants (garden_id, species, quantity, planting_date, status)
SELECT id, 'Strawberry', 20, CURRENT_DATE - interval '90 days', 'harvesting' FROM gardens LIMIT 1 OFFSET 8;
INSERT INTO plants (garden_id, species, quantity, planting_date, status)
SELECT id, 'Blueberry', 7, CURRENT_DATE - interval '100 days', 'dormant' FROM gardens LIMIT 1 OFFSET 9;
INSERT INTO plants (garden_id, species, quantity, planting_date, status)
SELECT id, 'Lavender', 9, CURRENT_DATE - interval '110 days', 'growing' FROM gardens LIMIT 1 OFFSET 10;
INSERT INTO plants (garden_id, species, quantity, planting_date, status)
SELECT id, 'Thyme', 11, CURRENT_DATE - interval '120 days', 'planted' FROM gardens LIMIT 1 OFFSET 11;
INSERT INTO plants (garden_id, species, quantity, planting_date, status)
SELECT id, 'Oregano', 13, CURRENT_DATE - interval '130 days', 'planned' FROM gardens LIMIT 1 OFFSET 12;
INSERT INTO plants (garden_id, species, quantity, planting_date, status)
SELECT id, 'Parsley', 14, CURRENT_DATE - interval '140 days', 'harvesting' FROM gardens LIMIT 1 OFFSET 13;
INSERT INTO plants (garden_id, species, quantity, planting_date, status)
SELECT id, 'Chives', 2, CURRENT_DATE - interval '150 days', 'dormant' FROM gardens LIMIT 1 OFFSET 14;

-- Insert mock data into care_logs table
INSERT INTO care_logs (plant_id, action, notes)
SELECT id, 'Watering', 'Watered the plant thoroughly.' FROM plants LIMIT 1;
INSERT INTO care_logs (plant_id, action, notes)
SELECT id, 'Fertilizing', 'Added organic fertilizer.' FROM plants LIMIT 1 OFFSET 1;
INSERT INTO care_logs (plant_id, action, notes)
SELECT id, 'Pruning', 'Pruned the dead leaves.' FROM plants LIMIT 1 OFFSET 2;
INSERT INTO care_logs (plant_id, action, notes)
SELECT id, 'Weeding', 'Removed weeds around the plant.' FROM plants LIMIT 1 OFFSET 3;
INSERT INTO care_logs (plant_id, action, notes)
SELECT id, 'Harvesting', 'Harvested the ripe fruits.' FROM plants LIMIT 1 OFFSET 4;
INSERT INTO care_logs (plant_id, action, notes)
SELECT id, 'Watering', 'Watered the plant lightly.' FROM plants LIMIT 1 OFFSET 5;
INSERT INTO care_logs (plant_id, action, notes)
SELECT id, 'Fertilizing', 'Added compost.' FROM plants LIMIT 1 OFFSET 6;
INSERT INTO care_logs (plant_id, action, notes)
SELECT id, 'Pruning', 'Trimmed the branches.' FROM plants LIMIT 1 OFFSET 7;
INSERT INTO care_logs (plant_id, action, notes)
SELECT id, 'Weeding', 'Cleared the weeds.' FROM plants LIMIT 1 OFFSET 8;
INSERT INTO care_logs (plant_id, action, notes)
SELECT id, 'Harvesting', 'Collected the vegetables.' FROM plants LIMIT 1 OFFSET 9;
INSERT INTO care_logs (plant_id, action, notes)
SELECT id, 'Watering', 'Gave the plant a deep soak.' FROM plants LIMIT 1 OFFSET 10;
INSERT INTO care_logs (plant_id, action, notes)
SELECT id, 'Fertilizing', 'Applied liquid fertilizer.' FROM plants LIMIT 1 OFFSET 11;
INSERT INTO care_logs (plant_id, action, notes)
SELECT id, 'Pruning', 'Removed old stems.' FROM plants LIMIT 1 OFFSET 12;
INSERT INTO care_logs (plant_id, action, notes)
SELECT id, 'Weeding', 'Weeded the garden bed.' FROM plants LIMIT 1 OFFSET 13;
INSERT INTO care_logs (plant_id, action, notes)
SELECT id, 'Harvesting', 'Picked the herbs.' FROM plants LIMIT 1 OFFSET 14;

COMMIT;