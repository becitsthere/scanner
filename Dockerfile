FROM 10.1.127.12:5000/neuvector/scanner_base

COPY stage /

LABEL neuvector.image="neuvector/scanner" \
      neuvector.role="scanner" \
      neuvector.vuln_db="vuln.xxxx"

ENTRYPOINT ["/usr/local/bin/monitor", "-s"]
