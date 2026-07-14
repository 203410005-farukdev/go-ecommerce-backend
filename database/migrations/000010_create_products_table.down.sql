-- Down migration for products table

DROP TRIGGER IF EXISTS trg_products_updated_at ON products;
DROP INDEX IF EXISTS idx_products_is_featured;
DROP INDEX IF EXISTS idx_products_is_active;
DROP INDEX IF EXISTS idx_products_slug;
DROP INDEX IF EXISTS idx_products_sku;
DROP INDEX IF EXISTS idx_products_subcategory_id;
DROP INDEX IF EXISTS idx_products_category_id;
DROP TABLE IF EXISTS products;
