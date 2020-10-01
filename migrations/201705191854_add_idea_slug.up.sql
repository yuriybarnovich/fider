ALTER TABLE ideas ADD slug varchar(1000);

UPDATE ideas SET slug = '';

ALTER TABLE ideas ALTER COLUMN slug SET NOT NULL;