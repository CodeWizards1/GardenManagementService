BEGIN;

DROP TABLE IF EXISTS care_logs CASCADE;

DROP TABLE IF EXISTS plants CASCADE;

DROP TABLE IF EXISTS gardens CASCADE;

DROP TYPE plant_status CASCADE;

DROP TYPE garden_type CASCADE;

COMMIT;
