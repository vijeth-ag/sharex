curl -u "email:token" \
  -X POST \
  -H "X-Atlassian-Token: nocheck" \
  -F "file=@report.csv" \
  -F "comment=File attached via REST API" \
  "https://baseurl/wiki/rest/api/content/65538/child/attachment" 
