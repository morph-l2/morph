import os

from logging.config import dictConfig

log_level = os.getenv('LOG_LEVEL', 'INFO').upper()

log_config = {
    'version': 1,
    'loggers': {
        '': {
            'handlers': ['console'],
            'level': log_level
        },
    },
    'handlers': {
        'console': {
            'formatter': 'stderr',
            'class': 'logging.StreamHandler',
            'stream': 'ext://sys.stdout'
        }
    },
    'formatters': {
        'stderr': {
            'format': '[%(levelname)s|%(asctime)s] %(message)s',
            'datefmt': '%m-%d-%Y %I:%M:%S'
        }
    },
}

dictConfig(log_config)
