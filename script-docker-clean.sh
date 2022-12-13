
# ## remove todos os containers orphans
sudo docker-compose down --remove-orphans

# ## deleta todos os containers
sudo docker rm $(docker ps -a -q)
sudo docker container prune -f

# ## deleta todos os volumes
sudo docker volume prune -f

# ## delete todas as redes
sudo docker network prune -f

# ## deleta rodas a imagens
sudo docker rmi $(docker images -q)
sudo docker image prune -f
sudo docker image prune -a -f