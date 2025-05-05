from sqlalchemy import Column, Integer, String, Float, DateTime
from sqlalchemy.orm import declarative_base
from datetime import datetime

Base = declarative_base()

class LabResult(Base):
    __tablename__ = "lab_results"
    id = Column(Integer, primary_key=True)
    patient_id = Column(String, index=True)
    type = Column(String)
    value = Column(Float)
    unit = Column(String)
    created_at = Column(DateTime, default=datetime.utcnow)
