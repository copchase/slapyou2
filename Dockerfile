FROM public.ecr.aws/lambda/provided:al2023-x86_64 as builder

RUN dnf update && \
    dnf install -y gzip tar && \
    curl -LO https://go.dev/dl/go1.22.2.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz

COPY . .

RUN export PATH=$PATH:/usr/local/go/bin && \
    go mod download && \
    go build


ENTRYPOINT ["cat", "slapyou2"]
