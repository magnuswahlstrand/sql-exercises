#air -c ./.air.toml & \
#npx tailwind \
#  -i 'styles.css' \
#  -o 'public/styles.css' \
#  --watch & \
air -d -c ./.air.toml
#& \
#npx browser-sync start \
#  --files 'public/**/*.html, public/**/*.css, functions/**/*.gohtml' \
#  --port 3001 \
#  --proxy 'localhost:3000' \
#  --middleware 'function(req, res, next) { \
#    res.setHeader("Cache-Control", "no-cache, no-store, must-revalidate"); \
#    return next(); \
#  }'