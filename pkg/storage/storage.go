package storage

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Хранилище данных.
type Storage struct {
	db *pgxpool.Pool
}

// Конструктор, принимает строку подключения к БД.
func New(constr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

// Задача.
type Task struct {
	ID           int
	Opened       int64
	Closed       int64
	AuthorID     int
	AssignedID   int
	Title        string
	Content      string
	AuthorName   string
	AssignedName string
	Labels       []string
}

// Пользователь
type User struct {
	ID   int
	Name string
}

// Users возвращает всех пользователей из БД.
func (s *Storage) Users(userID int) ([]User, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			id,
			name
		FROM users
		WHERE
			($1 = 0 OR id = $1)
		ORDER BY id;
	`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	var users []User
	// итерирование по результату выполнения запроса
	// и сканирование каждой строки в переменную
	for rows.Next() {
		var u User
		err = rows.Scan(
			&u.ID,
			&u.Name,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		users = append(users, u)

	}
	// ВАЖНО не забыть проверить rows.Err()
	return users, rows.Err()
}

// Tasks возвращает список задач из БД.
// Значение -1 вернет все задачи или всех авторов
func (s *Storage) Tasks(taskID, authorID int) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			id,
			opened,
			closed,
			author_id,
			assigned_id,
			title,
			content
		FROM tasks
		WHERE
			($1 = 0 OR id = $1) AND
			($2 = -1 OR author_id = $2)
		ORDER BY id;
	`,
		taskID,
		authorID,
	)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	// итерирование по результату выполнения запроса
	// и сканирование каждой строки в переменную
	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		tasks = append(tasks, t)

	}
	// ВАЖНО не забыть проверить rows.Err()
	return tasks, rows.Err()
}

// NewTask создаёт новую задачу и возвращает её id.
func (s *Storage) NewTask(t Task) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO tasks (title, content)
		VALUES ($1, $2) RETURNING id;
		`,
		t.Title,
		t.Content,
	).Scan(&id)
	return id, err
}

// DeleteTask удаляет задачу и возвращает ошибку .
func (s *Storage) DeleteTask(id int) error {
	_, err := s.db.Exec(context.Background(), `
		DELETE FROM tasks WHERE id=$1;
		`,
		id,
	)
	return err
}
