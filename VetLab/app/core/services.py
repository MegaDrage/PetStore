from app.core.schemas import LabResultCreate
from app.infrastructure.repository import LabResultRepository
from app.core.pdf_generator import generate_report
from app.core.analyzer import run_analysis


class LabResultService:
    def __init__(self, repo: LabResultRepository):
        self.repo = repo

    def add_result(self, data: LabResultCreate):
        return self.repo.create(data)

    def get_patient_results(self, patient_id: str):
        return self.repo.get_by_patient(patient_id)

    def generate_pdf_report(self, patient_id: str):
        results = self.repo.get_by_patient(patient_id)
        return generate_report(results, patient_id)

    def analyze_results(self, patient_id: str):
        results = self.repo.get_by_patient(patient_id)
        return [run_analysis(r.value) for r in results]
