curl -u "email:token" \
  -X POST \
  -H "X-Atlassian-Token: nocheck" \
  -F "file=@report.csv" \
  -F "comment=File attached via REST API" \
  "https://baseurl/wiki/rest/api/content/65538/child/attachment" 


  curl -X POST -u <YourUsername>:<YourPersonalAccessToken> -H "X-Atlassian-Token: no-check" -F "file=@/path/to/your/file.png" https://your-confluence-site/wiki/rest/api/content/<PageID>/child/attachment

