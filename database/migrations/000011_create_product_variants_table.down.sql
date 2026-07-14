-- Down migration for product variants table

DROP TRIGGER IF EXISTS trg_product_variants_updated_at ON product_variants;
DROP INDEX IF EXISTS idx_product_variants_sort_order;
DROP INDEX IF EXISTS idx_product_variants_is_active;
DROP INDEX IF EXISTS idx_product_variants_sku;
DROP INDEX IF EXISTS idx_product_variants_product_id;
DROP TABLE IF EXISTS product_variants;
