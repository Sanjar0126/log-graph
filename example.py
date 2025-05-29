import sys
import time
import random
import logging

logger = logging.getLogger("query_logger")
logger.setLevel(logging.INFO)

handler = logging.StreamHandler(sys.stdout)
handler.setLevel(logging.INFO)

formatter = logging.Formatter('%(message)s')
handler.setFormatter(formatter)

logger.handlers = [handler]
logger.propagate = False

def run_queries(count=200):
    for _ in range(count):
        duration = round(random.uniform(0.1, 3.0), 2)
        duration2 = round(random.uniform(1.2, 7.0), 2)
        time.sleep(duration) 
        logger.info(f"[query]: {duration}s")
        logger.info(f"[query2]: {duration2}s")

if __name__ == "__main__":
    run_queries()