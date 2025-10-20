INSERT INTO apps (id, name, secret)
VALUES (1, 'url-shortener', 'shortener-secret')
ON CONFLICT DO NOTHING;