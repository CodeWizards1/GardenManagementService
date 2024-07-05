BEGIN;

INSERT INTO gardens (user_id, name, type, area_sqm)
VALUES
(uuid_generate_v4(), 'Garden 1', 'balcony', 10.0),
(uuid_generate_v4(), 'Garden 2', 'rooftop', 20.0),
(uuid_generate_v4(), 'Garden 3', 'indoor', 5.0),
(uuid_generate_v4(), 'Garden 4', 'community', 50.0),
(uuid_generate_v4(), 'Garden 5', 'backyard', 30.0),
(uuid_generate_v4(), 'Garden 6', 'balcony', 15.0),
(uuid_generate_v4(), 'Garden 7', 'rooftop', 25.0),
(uuid_generate_v4(), 'Garden 8', 'indoor', 8.0),
(uuid_generate_v4(), 'Garden 9', 'community', 60.0),
(uuid_generate_v4(), 'Garden 10', 'backyard', 35.0),
(uuid_generate_v4(), 'Garden 11', 'balcony', 12.0),
(uuid_generate_v4(), 'Garden 12', 'rooftop', 22.0);

-- Insert mock data into plants table
INSERT INTO plants (garden_id, species, quantity, planting_date, status)
VALUES
((SELECT id FROM gardens OFFSET 0 LIMIT 1), 'Tomato', 10, '2023-01-01', 'planted'),
((SELECT id FROM gardens OFFSET 1 LIMIT 1), 'Lettuce', 20, '2023-02-01', 'growing'),
((SELECT id FROM gardens OFFSET 2 LIMIT 1), 'Carrot', 15, '2023-03-01', 'planned'),
((SELECT id FROM gardens OFFSET 3 LIMIT 1), 'Pepper', 12, '2023-04-01', 'harvesting'),
((SELECT id FROM gardens OFFSET 4 LIMIT 1), 'Cucumber', 18, '2023-05-01', 'dormant'),
((SELECT id FROM gardens OFFSET 5 LIMIT 1), 'Spinach', 22, '2023-06-01', 'planted'),
((SELECT id FROM gardens OFFSET 6 LIMIT 1), 'Broccoli', 25, '2023-07-01', 'growing'),
((SELECT id FROM gardens OFFSET 7 LIMIT 1), 'Kale', 30, '2023-08-01', 'planned'),
((SELECT id FROM gardens OFFSET 8 LIMIT 1), 'Radish', 35, '2023-09-01', 'harvesting'),
((SELECT id FROM gardens OFFSET 9 LIMIT 1), 'Beetroot', 40, '2023-10-01', 'dormant'),
((SELECT id FROM gardens OFFSET 10 LIMIT 1), 'Zucchini', 45, '2023-11-01', 'planted'),
((SELECT id FROM gardens OFFSET 11 LIMIT 1), 'Pumpkin', 50, '2023-12-01', 'growing');

-- Insert mock data into care_logs table
INSERT INTO care_logs (plant_id, action, notes)
VALUES
((SELECT id FROM plants OFFSET 0 LIMIT 1), 'Watered', 'Watered the plants in the morning'),
((SELECT id FROM plants OFFSET 1 LIMIT 1), 'Fertilized', 'Added organic fertilizer'),
((SELECT id FROM plants OFFSET 2 LIMIT 1), 'Pruned', 'Pruned the leaves'),
((SELECT id FROM plants OFFSET 3 LIMIT 1), 'Harvested', 'Harvested the ripe vegetables'),
((SELECT id FROM plants OFFSET 4 LIMIT 1), 'Watered', 'Watered the plants in the evening'),
((SELECT id FROM plants OFFSET 5 LIMIT 1), 'Fertilized', 'Added compost'),
((SELECT id FROM plants OFFSET 6 LIMIT 1), 'Pruned', 'Removed dead leaves'),
((SELECT id FROM plants OFFSET 7 LIMIT 1), 'Harvested', 'Collected the mature fruits'),
((SELECT id FROM plants OFFSET 8 LIMIT 1), 'Watered', 'Watered the plants in the afternoon'),
((SELECT id FROM plants OFFSET 9 LIMIT 1), 'Fertilized', 'Added liquid fertilizer'),
((SELECT id FROM plants OFFSET 10 LIMIT 1), 'Pruned', 'Trimmed the branches'),
((SELECT id FROM plants OFFSET 11 LIMIT 1), 'Harvested', 'Picked the vegetables');

COMMIT;