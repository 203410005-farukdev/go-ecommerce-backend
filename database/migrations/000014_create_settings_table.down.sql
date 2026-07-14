-- Down migration for settings table

DROP TRIGGER IF EXISTS trg_settings_updated_at ON settings;
DROP TABLE IF EXISTS settings;
