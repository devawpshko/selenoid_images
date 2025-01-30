# Use existing base image
FROM browsers/base:7.4.2

ARG VERSION
ARG INSTALL_DIR=chrome

LABEL browser="google-chrome:${VERSION}"

# 1) Install required system libraries (add missing dependencies)
RUN apt-get update && apt-get install -y --no-install-recommends \
    libasound2 \
    unzip \
     libatk1.0-0 \
     libc6 \
     libcairo2 \
     libcups2 \
     libdbus-1-3 \
     libexpat1 \
     libfontconfig1 \
     libfreetype6 \
     libgcc1 \
     libgconf-2-4 \
     libgdk-pixbuf2.0-0 \
     libglib2.0-0 \
     libgtk2.0-0 \
     libgtk-3-0 \
     libnspr4 \
     libnss3 \
     libpango1.0-0 \
     libstdc++6 \
     libx11-6 \
     libx11-xcb1 \
     libxcb1 \
     libxcomposite1 \
     libxcursor1 \
     libxdamage1 \
     libxext6 \
     libxfixes3 \
     libxi6 \
     libxrandr2 \
     libxrender1 \
     libxss1 \
     libxtst6 \
     libgbm1 \
     libu2f-udev \
     libvulkan1 \
     ca-certificates \
     fonts-liberation \
     libappindicator3-1 \
     libnss3 \
     lsb-base \
     xdg-utils \
     libcurl4 \
     iproute2 \
    curl \
    && rm -rf /var/lib/apt/lists/*

# 2) Download Chrome for Testing
RUN curl -fSLo /tmp/chrome-linux64.zip \
    "https://edgedl.me.gvt1.com/edgedl/chrome/chrome-for-testing/${VERSION}/linux64/chrome-linux64.zip" \
    && unzip /tmp/chrome-linux64.zip -d /opt/google/ \
    && rm /tmp/chrome-linux64.zip

# 3) Rename extracted Chrome folder
RUN mv /opt/google/chrome-linux64 /opt/google/${INSTALL_DIR}

# 4) Modify the Chrome binary to add flags
RUN sed -i -e 's@exec -a "$0" "$HERE/chrome"@& --no-sandbox --disable-gpu@' \
    /opt/google/${INSTALL_DIR}/chrome || true

# 5) Adjust permissions for the Chrome sandbox
RUN if [ -f /opt/google/${INSTALL_DIR}/chrome_sandbox ]; then \
      chown root:root /opt/google/${INSTALL_DIR}/chrome_sandbox && \
      chmod 4755 /opt/google/${INSTALL_DIR}/chrome_sandbox; \
    fi

# 6) Create a symlink so "google-chrome" is on PATH
RUN ln -s /opt/google/${INSTALL_DIR}/chrome /usr/bin/google-chrome

# 7) Confirm Chrome installation (this was failing before)
RUN google-chrome --version

# Cleanup
RUN rm -Rf /tmp/*
