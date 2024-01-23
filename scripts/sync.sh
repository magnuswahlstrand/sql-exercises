npx browser-sync start \
  --files 'public/**/*.html, public/**/*.css, functions/**/*.gohtml' \
  --port 3001 \
  --reload-delay 3000 \
  --proxy 'localhost:3000'