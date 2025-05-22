FROM python:3.13.3-slim

WORKDIR /app

COPY start.sh main.py requirements.txt /app/

RUN pip3 install --upgrade pip && pip install -r requirements.txt

EXPOSE 8000

CMD ["/bin/sh", "start.sh"]
# CMD ["pwd"]