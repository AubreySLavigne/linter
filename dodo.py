
DOIT_CONFIG = {'default_tasks': ['test']}


def task_test():
    """
    Run tests and display basic coverage
    """
    return {'actions': ['go test -cover ./...'],
            'verbosity': 2}


def task_coverage():
    """
    Run tests, then open the coverage report in a browser
    """
    return {'actions': ['go test ./... -coverprofile=coverage.out',
                        'go tool cover -html=coverage.out'],
            'verbosity': 2}
