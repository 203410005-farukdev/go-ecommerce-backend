-- Remove ordering and hero/offer section settings
DELETE FROM settings WHERE key IN (
  'store_whatsapp',
  'store_facebook',
  'order_whatsapp_enabled',
  'order_facebook_enabled',
  'hero_title',
  'hero_subtitle',
  'hero_image_url',
  'hero_btn_text',
  'hero_btn_link',
  'offer_is_active',
  'offer_title',
  'offer_subtitle',
  'offer_btn_text',
  'offer_btn_link',
  'offer_background_color'
);
