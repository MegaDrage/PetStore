from pydantic_settings import BaseSettings


class Settings(BaseSettings):
    app_name: str = "Lab Results Service"
    host: str = "0.0.0.0"
    port: int = 8090
    database_url: str = "postgresql://app_user:app_password@db:5432/app_db"

    class Config:
        env_file = ".env"


settings = Settings()
