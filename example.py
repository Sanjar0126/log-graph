import logging
import random
import sys
import time
from datetime import datetime

# Setup logger
logger = logging.getLogger("api_logger")
logger.setLevel(logging.DEBUG)  # Enable all levels

handler = logging.StreamHandler(sys.stdout)
handler.flush = sys.stdout.flush
formatter = logging.Formatter('%(asctime)s - %(levelname)s - %(message)s')
handler.setFormatter(formatter)
logger.addHandler(handler)

api_endpoints = ["/api/login", "/api/data", "/api/logout", "/api/profile"]
users = ["user123", "admin456", "guest789"]

def simulate_request():
    path = random.choice(api_endpoints)
    user = random.choice(users)
    start_time = datetime.now()

    logger.info(f"Received request: path={path}, user={user}")
    
    response_time = round(random.uniform(0.1, 2.0), 3)
    time.sleep(response_time)

    if path == "/api/data" and random.random() < 0.1:
        logger.error(f"Failed to fetch data for user={user} (500 Internal Server Error)")
        return
    
    if response_time > 1.5:
        logger.debug(f"Slow response: path={path}, user={user}, duration={response_time}s")

    logger.info(f"Response sent: path={path}, user={user}, duration={response_time}s")

def run_server_simulation(request_count=20):
    for _ in range(request_count):
        simulate_request()

if __name__ == "__main__":
    run_server_simulation()
