#!/bin/bash

git clone https://github.com/CVN003/scstgateway.git scstgateway-1.0.3
tar cvfz /root/rpmbuild/SOURCES/scstgateway-1.0.3.tar.gz scstgateway-1.0.3
rpmbuild -ba scstgateway-1.0.3/scstgateway.spec