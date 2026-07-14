-- Down migration for stock table

DROP TRIGGER IF EXISTS trg_stock_updated_at ON stock;
DROP INDEX IF EXISTS idx_stock_location;
DROP INDEX IF EXISTS idx_stock_sku;
DROP INDEX IF EXISTS idx_stock_product_id;
DROP TABLE IF EXISTS stock;
