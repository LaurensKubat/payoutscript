FROM node:9

# normally you would use a seperate docker container for the database.
# this image is however intended for maximum hacking purposes, so we simply put everything in it.
RUN apt update
RUN apt install postgresql postgresql-contrib -y

# redis is used for the transaction pool. We build it ourselves, as that is the recommended way to run redis
RUN apt install redis-server -y


# Lerna grabs our dependencies for us. (it seems this one randomly fails sometimes when building the image)
RUN npm install --global lerna --loglevel verbose
RUN git clone -b develop https://github.com/ArkEcosystem/core.git
RUN (cd core && lerna bootstrap)

# enabling all APIs by default
#ENV ARK_API_DISABLED=false
ENV ARK_WEBHOOKS_ENABLED=true
ENV ARK_WEBHOOKS_API_ENABLED=true
ENV ARK_GRAPHQL_ENABLED=true
ENV ARK_JSON_RPC_ENABLED=true

# public API, this one is for developers
EXPOSE 4003

# webhook port
EXPOSE 4004

# JSON-RPC
EXPOSE 8080

# public graphql API, including graphiQL explorer
EXPOSE 4005

# internal API, for nodes to communicate
EXPOSE 4000

# postgresql port, if you want to directly query the DB
EXPOSE 5432

COPY entrypoint.sh /

RUN echo "listen_addresses = '*'" >> /etc/postgresql/9.4/main/postgresql.conf


# this will start an entire test network, including genesis block. To find the secrets, check out:
# https://github.com/ArkEcosystem/core/blob/develop/packages/core/lib/config/testnet/delegates.json
ENTRYPOINT ./entrypoint.sh