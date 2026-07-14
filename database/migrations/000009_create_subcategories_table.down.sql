-- Down migration for subcategories table

DROP TRIGGER IF EXISTS trg_subcategories_updated_at ON subcategories;
DROP INDEX IF EXISTS idx_subcategories_sort_order;
DROP INDEX IF EXISTS idx_subcategories_is_active;
DROP INDEX IF EXISTS idx_subcategories_category_slug;
DROP INDEX IF EXISTS idx_subcategories_category_id;
DROP TABLE IF EXISTS subcategories;
