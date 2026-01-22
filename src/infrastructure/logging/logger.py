import logging
import sys


def get_logger(log_level: int = logging.ERROR):
    logger = logging.getLogger()
    
    logger.setLevel(log_level)
    
    # Create handlers for logging to the standard output and a file
    stdout_handler = logging.StreamHandler(stream=sys.stdout)
    fmt = logging.Formatter(
    "%(name)s: %(asctime)s | %(levelname)s | %(filename)s:%(lineno)s | %(process)d >>> %(message)s"
    )

    # Set the log format on each handler
    stdout_handler.setFormatter(fmt)
    
    logger.addHandler(stdout_handler)
    
    return logger