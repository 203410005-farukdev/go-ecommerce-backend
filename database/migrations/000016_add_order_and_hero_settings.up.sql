-- Add ordering and hero/offer section settings
INSERT INTO settings (key, value) VALUES
  ('store_whatsapp', '8801700000000'),
  ('store_facebook', 'https://facebook.com/shopnexus'),
  ('order_whatsapp_enabled', 'true'),
  ('order_facebook_enabled', 'true'),
  ('hero_title', 'আপনার প্রিয় পণ্য এখন দরজায়'),
  ('hero_subtitle', 'Shop Nexus — বাংলাদেশের বিশ্বস্ত ই-কমার্স গন্তব্য। মানসম্মত পণ্য, দ্রুত ডেলিভারি, সহজ রিটার্ন।'),
  ('hero_image_url', ''),
  ('hero_btn_text', 'কেনাকাটা শুরু করুন'),
  ('hero_btn_link', '/shop'),
  ('offer_is_active', 'true'),
  ('offer_title', '২৫% পর্যন্ত ছাড়'),
  ('offer_subtitle', 'নির্বাচিত পণ্যে বিশেষ ছাড়। স্টক শেষ হওয়ার আগেই অর্ডার করুন।'),
  ('offer_btn_text', 'অফার দেখুন'),
  ('offer_btn_link', '/shop'),
  ('offer_background_color', 'linear-gradient(135deg, rgba(46,205,176,0.12), rgba(245,166,35,0.08))')
ON CONFLICT (key) DO NOTHING;
