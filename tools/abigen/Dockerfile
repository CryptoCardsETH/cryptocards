FROM ethereum/client-go:alltools-stable 


RUN apk add --no-cache curl 

RUN curl -L -o /usr/bin/solc https://github.com/ethereum/solidity/releases/download/v0.4.17/solc-static-linux && chmod +x /usr/bin/solc

# run the codegen
ENTRYPOINT ["/bin/sh","/tools/abigen/docker_run.sh"]
CMD ["run"]
