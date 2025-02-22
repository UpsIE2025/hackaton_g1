import sqlite3

class MessageStorage:
    def __init__(self, db_path="messages.db"):
        self.db_path = db_path
        self._create_table()

    def _create_table(self):
        """ Crea la tabla para almacenar mensajes si no existe """
        with sqlite3.connect(self.db_path) as conn:
            cursor = conn.cursor()
            cursor.execute("""
                CREATE TABLE IF NOT EXISTS messages (
                    id INTEGER PRIMARY KEY AUTOINCREMENT,
                    channel TEXT,
                    message TEXT,
                    delivered INTEGER DEFAULT 0
                )
            """)
            conn.commit()

    def save_message(self, channel, message):
        """ Guarda un mensaje en la base de datos """
        with sqlite3.connect(self.db_path) as conn:
            cursor = conn.cursor()
            cursor.execute("INSERT INTO messages (channel, message) VALUES (?, ?)", (channel, message))
            conn.commit()
            return cursor.lastrowid

    def get_undelivered_messages(self):
        """ Obtiene mensajes no entregados """
        with sqlite3.connect(self.db_path) as conn:
            cursor = conn.cursor()
            cursor.execute("SELECT id, channel, message FROM messages WHERE delivered = 0")
            return cursor.fetchall()

    def mark_as_delivered(self, message_id):
        """ Marca un mensaje como entregado en la base de datos """
        with sqlite3.connect(self.db_path) as conn:
            cursor = conn.cursor()
            cursor.execute("UPDATE messages SET delivered = 1 WHERE id = ?", (message_id,))
            conn.commit()
