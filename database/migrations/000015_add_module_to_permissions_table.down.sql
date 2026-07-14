DROP INDEX IF EXISTS idx_permissions_module;
ALTER TABLE permissions DROP COLUMN IF EXISTS module;
