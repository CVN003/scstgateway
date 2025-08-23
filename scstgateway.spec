Name:           scstgateway
Version:        1.0.3
Release:        1%{?dist}
Summary:        SCST Gateway service
License:        MIT
URL:            https://github.com/CVN003/scstgateway
Source0:        %{name}-%{version}.tar.gz
%define debug_package %{nil}
BuildRequires:  systemd-devel

%description
SCST Gateway service provides gRPC interface for managing storage targets.

%prep
%autosetup

%build
go mod tidy
go build -o %{name} ./main.go

%install

mkdir -p %{buildroot}/usr/local/scstgateway/{bin,log,etc}
install -Dm 0755 %{name} %{buildroot}/usr/local/scstgateway/bin/%{name}
install -Dm 0644 conf.json %{buildroot}/usr/local/scstgateway/etc/
install -Dm 0644 %{name}.service %{buildroot}%{_unitdir}/%{name}.service

%files
/usr/local/scstgateway/
/usr/local/scstgateway/bin/%{name}
/usr/local/scstgateway/log/
/usr/local/scstgateway/etc/conf.json
%{_unitdir}/%{name}.service

%post
systemctl daemon-reload
systemctl enable --now %{name}.service

%preun
if [ $1 -eq 0 ]; then
    systemctl stop %{name}.service
    systemctl disable %{name}.service
fi

systemctl daemon-reload

%changelog
* Aug 18 2025 cvn003<unixkeeper@163.com> - 1.0.1
- Initial package
- Change base directory to /usr/local/scstgateway
- Add GetLiveConfig - 1.0.2
- Aug 23 2025 Add build.sh - 1.0.3