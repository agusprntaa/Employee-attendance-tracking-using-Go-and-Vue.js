import numpy as np
from fastapi import FastAPI, File, UploadFile
from fastapi.responses import JSONResponse
import cv2
from insightface.app import FaceAnalysis

app = FastAPI()

print("Loading InsightFace model...")
face_app = FaceAnalysis(name="buffalo_l", providers=["CPUExecutionProvider"])
face_app.prepare(ctx_id=0, det_size=(640, 640))
print("Model loaded - siap menerima request")


def read_image(file_bytes: bytes):
    nparr = np.frombuffer(file_bytes, np.uint8)
    img = cv2.imdecode(nparr, cv2.IMREAD_COLOR)
    if img is None:
        raise ValueError("File bukan gambar valid")
    return img


@app.get("/health")
def health():
    return {"status": "ok", "model": "buffalo_l"}


@app.post("/detect")
async def detect_face(image: UploadFile = File(...)):
    try:
        img_bytes = await image.read()
        if not img_bytes:
            return JSONResponse(status_code=400, content={"face_count": 0, "error": "File kosong"})
        img = read_image(img_bytes)
        faces = face_app.get(img)
        return {"face_count": len(faces)}
    except ValueError as e:
        return JSONResponse(status_code=400, content={"face_count": 0, "error": str(e)})
    except Exception as e:
        return JSONResponse(status_code=500, content={"face_count": 0, "error": str(e)})


@app.post("/compare")
async def compare_faces(
    input_image: UploadFile = File(...),
    reference_image: UploadFile = File(...)
):
    try:
        input_bytes = await input_image.read()
        ref_bytes = await reference_image.read()

        if not input_bytes:
            return JSONResponse(status_code=400, content={"similarity": 0.0, "error": "input_image kosong"})
        if not ref_bytes:
            return JSONResponse(status_code=400, content={"similarity": 0.0, "error": "reference_image kosong"})

        input_img = read_image(input_bytes)
        ref_img = read_image(ref_bytes)

        input_faces = face_app.get(input_img)
        ref_faces = face_app.get(ref_img)

        if len(input_faces) == 0:
            return JSONResponse(status_code=400, content={"similarity": 0.0, "error": "Tidak ada wajah di foto input"})
        if len(ref_faces) == 0:
            return JSONResponse(status_code=400, content={"similarity": 0.0, "error": "Tidak ada wajah di foto referensi"})

        input_emb = input_faces[0].normed_embedding
        ref_emb = ref_faces[0].normed_embedding

        similarity = float(np.dot(input_emb, ref_emb))
        similarity = max(0.0, min(1.0, similarity))

        return {"similarity": round(similarity, 4)}

    except ValueError as e:
        return JSONResponse(status_code=400, content={"similarity": 0.0, "error": str(e)})
    except Exception as e:
        return JSONResponse(status_code=500, content={"similarity": 0.0, "error": str(e)})