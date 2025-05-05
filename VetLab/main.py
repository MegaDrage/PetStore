from fastapi import FastAPI
from app.api import lab_results  # Импорт маршрутов

from app.config import settings
from app.core.models import Base
from app.infrastructure.db import engine

from app.config import settings

Base.metadata.create_all(bind=engine)

app = FastAPI(title=settings.app_name)

# Подключаем маршруты
app.include_router(lab_results.router)

# запускаем приложение
if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host=settings.host, port=settings.port)
