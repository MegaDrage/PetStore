from reportlab.lib.pagesizes import A4
from reportlab.lib.units import cm
from reportlab.pdfgen import canvas
from reportlab.platypus import Image
import matplotlib.pyplot as plt
import io
import base64
from app.core.models import LabResult
from datetime import datetime


def generate_chart(results: list[LabResult]) -> io.BytesIO:
    values = [r.value for r in results]
    dates = [r.created_at.strftime('%Y-%m-%d') for r in results]

    plt.figure(figsize=(6, 3))
    plt.plot(dates, values, marker='o', color='blue')
    plt.title("Lab Result Dynamics")
    plt.xlabel("Date")
    plt.ylabel("Value")
    plt.grid(True)

    buf = io.BytesIO()
    plt.tight_layout()
    plt.savefig(buf, format='png')
    buf.seek(0)
    plt.close()
    return buf


def generate_report(results: list[LabResult], patient_id: str) -> str:
    buffer = io.BytesIO()
    c = canvas.Canvas(buffer, pagesize=A4)
    width, height = A4

    # Заголовок
    c.setFont("Helvetica-Bold", 16)
    c.drawString(2 * cm, height - 2 * cm, f"Lab Report for Patient: {patient_id}")

    # Таблица данных
    c.setFont("Helvetica", 10)
    y = height - 3 * cm
    c.drawString(2 * cm, y, "Date")
    c.drawString(6 * cm, y, "Type")
    c.drawString(10 * cm, y, "Value")
    c.drawString(14 * cm, y, "Unit")
    y -= 0.5 * cm

    for result in results:
        c.drawString(2 * cm, y, result.created_at.strftime('%Y-%m-%d'))
        c.drawString(6 * cm, y, result.type)
        c.drawString(10 * cm, y, f"{result.value:.2f}")
        c.drawString(14 * cm, y, result.unit)
        y -= 0.4 * cm
        if y < 5 * cm:
            c.showPage()
            y = height - 2 * cm

    # Вставка графика
    chart_buf = generate_chart(results)
    img = Image(chart_buf)
    img_path = "chart.png"
    with open(img_path, "wb") as f:
        f.write(chart_buf.getbuffer())
    c.drawImage(img_path, 2 * cm, 2 * cm, width - 4 * cm, 7 * cm)

    c.showPage()
    c.save()
    buffer.seek(0)

    # Возвращаем base64
    return base64.b64encode(buffer.getvalue()).decode()
