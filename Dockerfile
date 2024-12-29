# Use an official Ubuntu base image
FROM ubuntu:22.04

# Set environment variables
ENV DEBIAN_FRONTEND=noninteractive \
    LANG=C.UTF-8 \
    LC_ALL=C.UTF-8

# Update the system and install required dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    cmake \
    curl \
    git \
    python3 \
    python3-pip \
    libssl-dev \
    pkg-config \
    wget \
    unzip \
    clang \
    lld \
    nodejs \
    npm && \
    rm -rf /var/lib/apt/lists/*

# Install Rust and wasm32 target
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y && \
    . $HOME/.cargo/env && \
    rustup target add wasm32-unknown-unknown

# Install wasm-pack
RUN . $HOME/.cargo/env && \
    cargo install wasm-pack

# Install Emscripten
RUN git clone https://github.com/emscripten-core/emsdk.git /emsdk && \
    cd /emsdk && \
    ./emsdk install latest && \
    ./emsdk activate latest && \
    echo "source /emsdk/emsdk_env.sh" >> /root/.bashrc

# Install Wasmtime
RUN . $HOME/.cargo/env && \
    cargo install wasmtime-cli

# Install wasm-tools
RUN . $HOME/.cargo/env && \
    cargo install wasm-tools

# Install wac
RUN . $HOME/.cargo/env && \
    cargo install wac-cli

# Install wepl
RUN . $HOME/.cargo/env && \
    cargo install wepl

# Install componentize-py
RUN pip3 install componentize-py

# Install jco
RUN npm install -g @bytecodealliance/componentize-js @bytecodealliance/jco

# Install wkg
RUN . $HOME/.cargo/env && \
    cargo install wkg

RUN . $HOME/.cargo/env && \
    cargo install wit-bindgen-cli

RUN wget https://go.dev/dl/go1.23.4.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz

RUN wget https://github.com/tinygo-org/tinygo/releases/download/v0.34.0/tinygo_0.34.0_amd64.deb && \
    dpkg -i tinygo_0.34.0_amd64.deb

#Install WasmCloud Shell (wash)
RUN apt-get update && \
    curl -s https://packagecloud.io/install/repositories/wasmcloud/core/script.deb.sh | bash && \
    apt-get install -y --no-install-recommends wash && \
    rm -rf /var/lib/apt/lists/*

RUN . $HOME/.cargo/env && \
    cargo install wit-deps-cli

RUN rustup target add wasm32-wasip2

RUN apt-get update && \
    apt-get install -y --no-install-recommends less && \
    rm -rf /var/lib/apt/lists/*

# Add the necessary environment variables to the PATH
ENV PATH="/emsdk:/emsdk/upstream/emscripten:/emsdk/node/14.15.5_64bit/bin:/root/.cargo/bin:/usr/local/go/bin:/usr/local/bin:$PATH"

# Verify installations
RUN node -v && npm -v && \
    . $HOME/.cargo/env && \
    rustc --version && cargo --version && wasm-pack --version && \
    emcc --version && wasmtime --version && wasm-tools --version && wac --version && \
    wepl --version && componentize-py --version && wkg --version && wash --version

# Set up a working directory
WORKDIR /workspace

CMD ["/bin/bash"]
