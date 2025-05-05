# Используем официальный образ Python
FROM python:3.11.3-slim

# Устанавливаем зависимости
WORKDIR /app

# Копируем зависимости
COPY requirements.txt .

# Установка зависимостей
RUN pip install --no-cache-dir -r requirements.txt

# Копируем все файлы проекта
COPY . .

# Открываем порт
EXPOSE 8090

# Команда для запуска FastAPI
CMD ["python3", "main.py"]
