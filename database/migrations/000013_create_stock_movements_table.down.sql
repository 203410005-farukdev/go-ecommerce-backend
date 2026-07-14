-- Down migration for stock movements table

DROP INDEX IF EXISTS idx_stock_movements_reason;
DROP INDEX IF EXISTS idx_stock_movements_created_at;
DROP INDEX IF EXISTS idx_stock_movements_product_id;
DROP TABLE IF EXISTS stock_movements;
