FROM golang:1.24.2

# Node.jsのインストール
RUN curl -fsSL https://deb.nodesource.com/setup_22.x | bash - \
    && apt-get install -y nodejs

RUN npm install -g @typespec/compiler