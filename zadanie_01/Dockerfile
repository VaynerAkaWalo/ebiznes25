FROM ubuntu:24.04

RUN apt update && apt-get install -y software-properties-common

RUN add-apt-repository ppa:deadsnakes/ppa && apt-get install -y python3.10

RUN update-alternatives --install /usr/bin/python3 python3 /usr/bin/python3.10 1

CMD python3 --version

RUN apt-get install -y unzip curl zip

RUN curl -s "https://get.sdkman.io" | bash

RUN /bin/bash -c "source $HOME/.sdkman/bin/sdkman-init.sh && sdk install java 8.0.452-zulu && sdk install kotlin 2.1.20"

ENV PATH=/root/.sdkman/candidates/java/current/bin:$PATH
ENV PATH=/root/.sdkman/candidates/kotlin/current/bin:$PATH

CMD java -version && kotlinc -version
