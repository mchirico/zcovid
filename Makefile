
docker-build:
	docker build --no-cache -t gcr.io/mchirico/zcovid:test -f Dockerfile .

push:
	docker push gcr.io/mchirico/zcovid:test

build:
	go build -v .

run:
	docker run --rm -it -p 3000:3000  gcr.io/mchirico/zcovid:test

deploy:
	docker build --no-cache -t gcr.io/mchirico/zcovid:test -f Dockerfile .
	docker push gcr.io/mchirico/zcovid:test
	gcloud run deploy zcovid  --image gcr.io/mchirico/zcovid:test --platform managed \
            --allow-unauthenticated --project mchirico \
            --region us-east1 --port 3000 --max-instances 1  --memory 256Mi
