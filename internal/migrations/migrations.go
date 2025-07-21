package migrations

import (
	"errors"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func SafeUUIDMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20250721_safe_uuid_migration",
		Migrate: func(tx *gorm.DB) error {
			// 1. Добавляем новые колонки uuid с дефолтом генерации uuid
			if err := tx.Exec(`
				ALTER TABLE users ADD COLUMN new_id uuid DEFAULT gen_random_uuid();
				ALTER TABLE tokens ADD COLUMN new_user_id uuid;
				ALTER TABLE tokens ADD COLUMN new_id uuid DEFAULT gen_random_uuid();
			`).Error; err != nil {
				return err
			}

			// 2. Заполняем tokens.new_user_id по соответствию tokens.user_id -> users.id -> users.new_id
			if err := tx.Exec(`
				UPDATE tokens SET new_user_id = users.new_id
				FROM users WHERE tokens.user_id = users.id;
			`).Error; err != nil {
				return err
			}

			// 3. Проверяем что все new_user_id заполнены
			var count int64
			if err := tx.Raw(`SELECT COUNT(*) FROM tokens WHERE new_user_id IS NULL`).Scan(&count).Error; err != nil {
				return err
			}
			if count > 0 {
				return errors.New("not all tokens.new_user_id were populated")
			}

			// 4. Удаляем старые внешние ключи и первичные ключи
			if err := tx.Exec(`
				ALTER TABLE tokens DROP CONSTRAINT IF EXISTS tokens_user_id_fkey;
				ALTER TABLE tokens DROP CONSTRAINT IF EXISTS tokens_pkey;
				ALTER TABLE users DROP CONSTRAINT IF EXISTS users_pkey;
			`).Error; err != nil {
				return err
			}

			// 5. Удаляем старые колонки
			if err := tx.Exec(`
				ALTER TABLE tokens DROP COLUMN user_id;
				ALTER TABLE tokens DROP COLUMN id;
				ALTER TABLE users DROP COLUMN id;
			`).Error; err != nil {
				return err
			}

			// 6. Переименовываем новые колонки
			if err := tx.Exec(`
				ALTER TABLE users RENAME COLUMN new_id TO id;
				ALTER TABLE tokens RENAME COLUMN new_user_id TO user_id;
				ALTER TABLE tokens RENAME COLUMN new_id TO id;
			`).Error; err != nil {
				return err
			}

			// 7. Добавляем новые первичные ключи и внешние ключи
			if err := tx.Exec(`
				ALTER TABLE users ADD PRIMARY KEY (id);
				ALTER TABLE tokens ADD PRIMARY KEY (id);
				ALTER TABLE tokens ADD CONSTRAINT tokens_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
			`).Error; err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			// Откатить миграцию (опционально, в обратном порядке)
			return errors.New("rollback not implemented")
		},
	}
}
