package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rijenth/gRPC/internal/config"
)

func InitDB() (*sql.DB, error) {
	cfg, err := config.LoadDatabaseConfig()

	if err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, fmt.Errorf("failed to open the database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	log.Println("Database connection established")

	if err := SeedUsersTable(db); err != nil {
		return nil, fmt.Errorf("failed to seed users table: %v", err)
	}

	return db, nil
}

func SeedUsersTable(db *sql.DB) error {
	log.Println("🔄 Seeding users table...")

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		date_of_birth DATE,
		address VARCHAR(255),
		phone_number VARCHAR(20),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		last_login TIMESTAMP NULL,
		is_active BOOLEAN DEFAULT TRUE,
		is_admin BOOLEAN DEFAULT FALSE,
		profile_picture VARCHAR(255),
		bio TEXT
	) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("error creating users table: %v", err)
	}

	deleteUsersQuery := `
	DELETE FROM users
	WHERE username IN ('alice_d', 'bob_m', 'claire_d', 'david_l', 'emma_m', 'francois_s', 'gabrielle_r', 'hugo_b', 'isabelle_m', 'julien_f');`
	_, err = db.Exec(deleteUsersQuery)
	if err != nil {
		return fmt.Errorf("error deleting existing users: %v", err)
	}

	insertUsersQuery := `
		INSERT INTO users (username, password, email, first_name, last_name, date_of_birth, address, phone_number, last_login, is_active, is_admin, profile_picture, bio) VALUES
		('alice_d', '$2a$10$.yyVfpcnOrsn28X6Bg5jveYeyhRsCLR/ayrs9ZhGZr5mXtwE2w/ZG', 'alice.dupont@example.com', 'Alice', 'Dupont', '1990-05-14', '123 Rue Principale, Paris', '0612345678', '2023-09-01 08:30:00', TRUE, FALSE, '/images/alice_d.png', 'Passionnee de photographie et de voyages.'),
		('bob_m', '$2a$10$.yyVfpcnOrsn28X6Bg5jveYeyhRsCLR/ayrs9ZhGZr5mXtwE2w/ZG', 'bob.martin@example.com', 'Bob', 'Martin', '1985-08-22', '456 Avenue des Champs, Lyon', '0623456789', '2023-09-02 09:15:00', TRUE, FALSE, '/images/bob_m.png', 'Amateur de cuisine et de musique jazz.'),
		('claire_d', '$2a$10$.yyVfpcnOrsn28X6Bg5jveYeyhRsCLR/ayrs9ZhGZr5mXtwE2w/ZG', 'claire.durand@example.com', 'Claire', 'Durand', '1992-12-05', '789 Boulevard du Centre, Marseille', '0634567890', '2023-09-03 10:00:00', TRUE, FALSE, '/images/claire_d.png', 'Lectrice avide et ecrivaine en herbe.'),
		('david_l', '$2a$10$.yyVfpcnOrsn28X6Bg5jveYeyhRsCLR/ayrs9ZhGZr5mXtwE2w/ZG', 'david.lefevre@example.com', 'David', 'Lefevre', '1988-03-18', '101 Rue de la Paix, Toulouse', '0645678901', '2023-09-04 10:45:00', TRUE, FALSE, '/images/david_l.png', 'Ingenieur logiciel et amateur de randonnee.'),
		('emma_m', '$2a$10$.yyVfpcnOrsn28X6Bg5jveYeyhRsCLR/ayrs9ZhGZr5mXtwE2w/ZG', 'emma.moreau@example.com', 'Emma', 'Moreau', '1995-07-30', '202 Avenue Victor Hugo, Nice', '0656789012', '2023-09-05 11:30:00', TRUE, FALSE, '/images/emma_m.png', 'Artiste peintre et fan de cinema classique.'),
		('francois_s', '$2a$10$.yyVfpcnOrsn28X6Bg5jveYeyhRsCLR/ayrs9ZhGZr5mXtwE2w/ZG', 'francois.simon@example.com', 'Francois', 'Simon', '1980-11-11', '303 Rue de la Liberte, Nantes', '0667890123', '2023-09-06 12:15:00', TRUE, FALSE, '/images/francois_s.png', 'Photographe professionnel et globe-trotter.'),
		('gabrielle_r', '$2a$10$.yyVfpcnOrsn28X6Bg5jveYeyhRsCLR/ayrs9ZhGZr5mXtwE2w/ZG', 'gabrielle.roux@example.com', 'Gabrielle', 'Roux', '1993-02-25', '404 Place de lEglise, Strasbourg', '0678901234', '2023-09-07 13:00:00', TRUE, FALSE, '/images/gabrielle_r.png', 'Danseuse classique et instructrice de yoga.'),
		('hugo_b', '$2a$10$.yyVfpcnOrsn28X6Bg5jveYeyhRsCLR/ayrs9ZhGZr5mXtwE2w/ZG', 'hugo.bonnet@example.com', 'Hugo', 'Bonnet', '1987-06-17', '505 Chemin des Ecoliers, Bordeaux', '0689012345', '2023-09-08 13:45:00', TRUE, FALSE, '/images/hugo_b.png', 'Historien passionne par les civilisations antiques.'),
		('isabelle_m', '$2a$10$.yyVfpcnOrsn28X6Bg5jveYeyhRsCLR/ayrs9ZhGZr5mXtwE2w/ZG', 'isabelle.mercier@example.com', 'Isabelle', 'Mercier', '1991-09-09', '606 Allee des Roses, Lille', '0690123456', '2023-09-09 14:30:00', TRUE, FALSE, '/images/isabelle_m.png', 'Chef patissiere amoureuse des desserts francais.'),
		('julien_f', '$2a$10$.yyVfpcnOrsn28X6Bg5jveYeyhRsCLR/ayrs9ZhGZr5mXtwE2w/ZG', 'julien.fournier@example.com', 'Julien', 'Fournier', '1983-04-02', '707 Route du Soleil, Rennes', '0601234567', '2023-09-10 15:15:00', TRUE, TRUE, '/images/julien_f.png', 'Administrateur systeme avec 10 ans d experience.');
		`

	_, err = db.Exec(insertUsersQuery)
	if err != nil {
		return fmt.Errorf("error inserting users: %v", err)
	}

	log.Println("Users table seeded successfully")
	return nil
}
