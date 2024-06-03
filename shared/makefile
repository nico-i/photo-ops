dev:
	. env/bin/activate && cd ../.. && python -m services.$(SERVICE).src.presentation.grpc

venv:
	python -m venv env

install:
	. env/bin/activate && pip install --force-reinstall --no-deps -r requirements.txt

freeze:
	. env/bin/activate && pip freeze > requirements.txt

clean:
	rm -rf env