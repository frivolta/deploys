aws ecr get-login-password --region eu-west-1 | docker login --username AWS --password-stdin 120030055072.dkr.ecr.eu-west-1.amazonaws.com
docker tag frivolta/birdie-t2 120030055072.dkr.ecr.eu-west-1.amazonaws.com/deploys
