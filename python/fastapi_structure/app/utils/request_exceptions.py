from fastapi.encoders import jsonable_encoder
from fastapi.exceptions import RequestValidationError
from starlette.exceptions import HTTPException
from starlette.requests import Request
from starlette.responses import JSONResponse
from starlette.status import HTTP_422_UNPROCESSABLE_ENTITY


async def http_exception_handler(req: Request, exc: HTTPException) -> JSONResponse:
    return JSONResponse({"detail": exc.detail}, status_code=exc.status_code)


async def request_validation_exception_handler(
    req: Request, exc: RequestValidationError
) -> JSONResponse:
    return JSONResponse(
        {"detail": jsonable_encoder(exc.errors())},
        status_code=HTTP_422_UNPROCESSABLE_ENTITY,
    )
