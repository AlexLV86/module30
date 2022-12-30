-- перед созданием удаляю все таблицы и создаю их заново
DROP TABLE IF EXISTS users, tasks, labels, tasks_labels;
DROP SEQUENCE IF EXISTS labels_id_seq, tasks_id_seq, users_id_seq;

CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL
);

CREATE TABLE labels (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL
);

CREATE TABLE tasks (
	id SERIAL PRIMARY KEY,
	opened BIGINT NOT NULL 
	DEFAULT extract(epoch from now()), -- дата создания задачи по умолчанию заполняется текущим временем
	closed BIGINT DEFAULT 0, -- дата закрытия задачи
	author_id INT REFERENCES users(id) ON DELETE CASCADE DEFAULT 0 , -- автор задачи из таблицы users
	assigned_id INT REFERENCES users(id) ON DELETE CASCADE DEFAULT 0 , -- ответственный по задаче из таблицы users
	title TEXT DEFAULT 'Без названия',
	content TEXT
);

CREATE TABLE tasks_labels (
	task_id INT REFERENCES tasks(id) ON DELETE CASCADE, -- ссылка на tasks
	label_id INT REFERENCES labels(id) ON DELETE CASCADE -- ссылка на labels
);

-- добавим пользователя по умолчанию
INSERT INTO users (id, name) VALUES (0, 'default');
