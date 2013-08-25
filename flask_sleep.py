from flask import Flask
import time

app = Flask(__name__)

@app.route("/")
def hello():
    time.sleep(100)
    return "Hello World!"

if __name__ == "__main__":
    app.run(port=8000)
