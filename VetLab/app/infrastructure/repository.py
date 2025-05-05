from app.core.models import LabResult
from app.core.schemas import LabResultCreate
from sqlalchemy.orm import Session


class LabResultRepository:
    def __init__(self, db: Session):
        self.db = db

    def create(self, data: LabResultCreate) -> LabResult:
        lab_result = LabResult(**data.dict())
        self.db.add(lab_result)
        self.db.commit()
        self.db.refresh(lab_result)
        return lab_result

    def get_by_patient(self, patient_id: str):
        return self.db.query(LabResult).filter_by(patient_id=patient_id).all()
