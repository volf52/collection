from fastapi import FastAPI
from fastapi.exceptions import RequestValidationError
from starlette.exceptions import HTTPException as StarletterHTTPException

from app.config.database import create_tables
from app.routers import foo as foo_router
from app.utils.app_exceptions import AppExceptionCase, app_exception_handler
from app.utils.request_exceptions import http_exception_handler, request_validation_exception_handler

create_tables()

app = FastAPI()


@app.exception_handler(StarletterHTTPException)
async def custom_http_exception_handler(req, e):
    return await http_exception_handler(req, e)


@app.exception_handler(RequestValidationError)
async def custom_validation_exception_handler(req, e):
    return await request_validation_exception_handler(req, e)


@app.exception_handler(AppExceptionCase)
async def custom_app_exception_handler(req, e):
    return await app_exception_handler(req, e)


app.include_router(foo_router.router)


@app.get("/")
async def root():
    return {"message": "Hello World"}
