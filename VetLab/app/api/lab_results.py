from fastapi import APIRouter, Depends, HTTPException
from app.core.schemas import LabResultCreate, LabResultOut
from app.infrastructure.db import get_db
from app.infrastructure.repository import LabResultRepository
from app.core.services import LabResultService
from sqlalchemy.orm import Session

router = APIRouter(prefix="/results", tags=["Lab Results"])


def get_service(db: Session = Depends(get_db)):
    return LabResultService(LabResultRepository(db))


@router.post("/", response_model=LabResultOut)
def add_result(data: LabResultCreate, service: LabResultService = Depends(get_service)):
    return service.add_result(data)


@router.get("/{patient_id}", response_model=list[LabResultOut])
def get_results(patient_id: str, service: LabResultService = Depends(get_service)):
    return service.get_patient_results(patient_id)


@router.get("/report/{patient_id}")
def get_pdf_report(patient_id: str, service: LabResultService = Depends(get_service)):
    pdf = service.generate_pdf_report(patient_id)
    return {"pdf_base64": pdf}


@router.get("/analyze/{patient_id}")
def analyze(patient_id: str, service: LabResultService = Depends(get_service)):
    return service.analyze_results(patient_id)
