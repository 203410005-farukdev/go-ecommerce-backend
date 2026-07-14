-- Down migration for categories table

DROP TRIGGER IF EXISTS trg_categories_updated_at ON categories;
DROP INDEX IF EXISTS idx_categories_sort_order;
DROP INDEX IF EXISTS idx_categories_is_active;
DROP INDEX IF EXISTS idx_categories_slug;
DROP TABLE IF EXISTS categories;
