require('esbuild').build({
  entryPoints: ['assets/javascript/app.ts', 'assets/css/style.css'],
  bundle: true,
  minify: true,
  sourcemap: true,
  target: ['es2015'],
  outfile: 'public/dist/app.js',
}).catch(() => process.exit(1))
