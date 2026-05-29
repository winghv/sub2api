-- Update untouched legacy brand defaults to Superwhv.
UPDATE settings
SET value = 'Superwhv', updated_at = NOW()
WHERE key = 'site_name'
  AND value IN ('', 'Sub2API');

UPDATE settings
SET value = '助力碳基进化 · Empowering Carbon-Based Evolution', updated_at = NOW()
WHERE key = 'site_subtitle'
  AND value IN ('', 'Subscription to API', 'Subscription to API Conversion Platform', '订阅转 API 转换平台');

INSERT INTO settings (key, value, updated_at)
VALUES ('site_subtitle', '助力碳基进化 · Empowering Carbon-Based Evolution', NOW())
ON CONFLICT (key) DO NOTHING;

UPDATE settings
SET value = '/superwhv-logo.png', updated_at = NOW()
WHERE key = 'site_logo'
  AND value IN ('', '/logo.png');

INSERT INTO settings (key, value, updated_at)
VALUES ('site_logo', '/superwhv-logo.png', NOW())
ON CONFLICT (key) DO NOTHING;
