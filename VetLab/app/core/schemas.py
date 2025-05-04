from pydantic import BaseModel
from datetime import datetime


class LabResultCreate(BaseModel):
    patient_id: str
    type: str
    value: float
    unit: str


class LabResultOut(LabResultCreate):
    id: int
    created_at: datetime
