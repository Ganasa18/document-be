CREATE TABLE IF NOT EXISTS user_access_menu (
    id SERIAL PRIMARY KEY, role_id INT, menu_id INT, "create" BOOLEAN DEFAULT FALSE, read BOOLEAN DEFAULT FALSE,
    update BOOLEAN DEFAULT FALSE, delete BOOLEAN DEFAULT FALSE, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, deleted_at TIMESTAMP DEFAULT NULL
);