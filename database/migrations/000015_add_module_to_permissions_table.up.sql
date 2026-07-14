ALTER TABLE permissions ADD COLUMN IF NOT EXISTS module VARCHAR(100) NOT NULL DEFAULT 'OTHER';
CREATE INDEX IF NOT EXISTS idx_permissions_module ON permissions(module);
