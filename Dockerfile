FROM debian:stretch-slim
ADD employee_webservice /bin/employee_webservice
CMD ["bin/employee_webservice"]
