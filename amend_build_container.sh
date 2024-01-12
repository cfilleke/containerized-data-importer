
exec into the build container locally and run this for the s390x cross-compile

     dnf install -y --setopt=install_weak_deps=False  --installroot /usr/s390x-linux-gnu/sys-root --forcearch s390x --releasever 9 \
          glibc-devel \
          glibc-static && \
      dnf install -y epel-release && \
      dnf install -y --setopt=install_weak_deps=False \
          gcc-s390x-linux-gnu && \
      dnf clean -y all; \

