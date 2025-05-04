import tensorflow as tf
import numpy as np

# Подгружаем модель
interpreter = tf.lite.Interpreter(model_path="model.tflite")
interpreter.allocate_tensors()
input_details = interpreter.get_input_details()
output_details = interpreter.get_output_details()


def run_analysis(value: float) -> dict:
    input_data = np.array([[value]], dtype=np.float32)
    interpreter.set_tensor(input_details[0]['index'], input_data)
    interpreter.invoke()
    output = interpreter.get_tensor(output_details[0]['index'])
    return {"prediction": float(output[0][0])}
